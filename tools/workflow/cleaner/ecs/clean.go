/*
 * Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package ecs

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const (
	Type          = "ecs"
	clusterPrefix = "aoc-testing"
	taskDefPrefix = "taskdef"
)

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)
var errServiceTooNewError = errors.New("service is not past the expiration date")

func Clean(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	logger.Printf("Begin to clean ECS Task Definitions")
	client := ecs.NewFromConfig(cfg)

	paginator := ecs.NewListClustersPaginator(client, &ecs.ListClustersInput{})

	for paginator.HasMorePages() {
		lco, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		if len(lco.ClusterArns) == 0 {
			continue
		}

		dco, err := client.DescribeClusters(ctx, &ecs.DescribeClustersInput{Clusters: lco.ClusterArns})
		if err != nil {
			return err
		}
		for _, cluster := range dco.Clusters {
			if shouldConsiderCluster(cluster) {
				/**
				I would prefer to use multi errors here instead of exiting immediately. This would allow us to
				build a list of errors and still iterate through the clusters rather than failing fast.
				This should be a follow-on fix.
				As of now: failing fast gives us better insight into edge cases where the resource cleaner may be
				failing rather than having to log dive a successful run.
				TODO: use mutli errs and don't fail fast
				*/
				if cluster.ActiveServicesCount > 0 {
					err := deleteServices(ctx, client, expirationDate, cluster.ClusterArn)
					if err != nil && errors.Is(err, errServiceTooNewError) {
						logger.Printf("skipping %s deletion: cluster service not past expiration date",
							*cluster.ClusterName)
						continue
					} else if err != nil {
						return fmt.Errorf("unable to remove all services in cluster %s due to %w", *cluster.ClusterName, err)
					}
				}
				logger.Printf("Trying to delete cluster %s", *cluster.ClusterName)
				_, err = client.DeleteCluster(ctx, &ecs.DeleteClusterInput{Cluster: cluster.ClusterArn})
				if err != nil {
					return fmt.Errorf("unable to delete cluster %s due to %w", *cluster.ClusterName, err)
				} else {
					logger.Printf("Deleted cluster %s successfully", *cluster.ClusterName)
				}
			}
		}
	}
	return nil
}

func deleteServices(ctx context.Context, client *ecs.Client, expirationDate time.Time, clusterArn *string) error {
	var remainingServices int
	paginator := ecs.NewListServicesPaginator(client, &ecs.ListServicesInput{Cluster: clusterArn})

	for paginator.HasMorePages() {
		lso, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		if len(lso.ServiceArns) == 0 {
			continue
		}

		dso, err := client.DescribeServices(ctx, &ecs.DescribeServicesInput{Cluster: clusterArn, Services: lso.ServiceArns})
		if err != nil {
			return err
		}
		remainingServices += len(dso.Services)
		for _, service := range dso.Services {
			if shouldDeleteService(service, expirationDate) {
				logger.Printf("Trying to delete service %s created on %v", *service.ServiceName, *service.CreatedAt)
				if service.TaskDefinition != nil {
					if err = deleteTaskDefinition(ctx, client, expirationDate, service.TaskDefinition); err != nil {
						return err
					}
				}
				_, err = client.DeleteService(ctx, &ecs.DeleteServiceInput{Cluster: clusterArn, Service: service.ServiceName, Force: aws.Bool(true)})
				if err != nil {
					return err
				}
				logger.Printf("Deleted service %s successfully", *service.ServiceName)
				remainingServices--
			} else {
				return errServiceTooNewError
			}

		}
	}
	if remainingServices > 0 {
		return fmt.Errorf("failed to delete all services")
	}
	return nil
}

func deleteTaskDefinition(ctx context.Context, client *ecs.Client, expirationDate time.Time, taskDefinitionArn *string) error {
	dtdo, err := client.DescribeTaskDefinition(ctx, &ecs.DescribeTaskDefinitionInput{TaskDefinition: taskDefinitionArn})
	if err != nil {
		return err
	}
	taskDefinition := dtdo.TaskDefinition
	if shouldDeregisterTaskDefinition(*taskDefinition, expirationDate) {
		logger.Printf("Trying to de-register task definition %s registered on %v", *taskDefinition.Family, *taskDefinition.RegisteredAt)
		_, err = client.DeregisterTaskDefinition(ctx, &ecs.DeregisterTaskDefinitionInput{TaskDefinition: taskDefinitionArn})
		if err != nil {
			return err
		}
		logger.Printf("De-registered task definition %s successfully", *taskDefinition.Family)
	}
	return nil
}

// shouldConsiderCluster returns true if the cluster name starts with the expected prefix.
func shouldConsiderCluster(cluster types.Cluster) bool {
	return cluster.ClusterName != nil && strings.HasPrefix(*cluster.ClusterName, clusterPrefix)
}

// shouldDeleteService returns true if the service is older than the expiration date.
func shouldDeleteService(service types.Service, expirationDate time.Time) bool {
	return service.CreatedAt != nil && expirationDate.After(*service.CreatedAt)
}

// shouldDeregisterTaskDefinition returns true if the task definition family starts with the expected prefix
// and is older than the expiration date.
func shouldDeregisterTaskDefinition(taskDef types.TaskDefinition, expirationDate time.Time) bool {
	if taskDef.Family == nil || taskDef.RegisteredAt == nil {
		return false
	}
	return strings.HasPrefix(*taskDef.Family, taskDefPrefix) && expirationDate.After(*taskDef.RegisteredAt)
}
