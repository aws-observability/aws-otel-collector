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
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

const (
	Type          = "ecs"
	clusterPrefix = "aoc-testing"
	taskDefPrefix = "taskdef"
)

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)
var errServiceTooNewError = errors.New("service is not past the expiration date")

func Clean(sess *session.Session, expirationDate time.Time) error {
	logger.Printf("Begin to clean ECS Task Definitions")
	client := ecs.New(sess)

	var nextToken *string
	for {
		lci := &ecs.ListClustersInput{NextToken: nextToken}
		lco, err := client.ListClusters(lci)
		if err != nil {
			return err
		}
		dci := &ecs.DescribeClustersInput{Clusters: lco.ClusterArns}
		dco, err := client.DescribeClusters(dci)
		if err != nil {
			return err
		}
		for _, cluster := range dco.Clusters {
			if strings.HasPrefix(*cluster.ClusterName, clusterPrefix) {
				/**
				I would prefer to use multi errors here instead of exiting immediately. This would allow us to
				build a list of errors and still iterate through the clusters rather than failing fast.
				This should be a follow-on fix.
				As of now: failing fast gives us better insight into edge cases where the resource cleaner may be
				failing rather than having to log dive a successful run.
				TODO: use mutli errs and don't fail fast
				*/
				if *cluster.ActiveServicesCount > 0 {
					err := deleteServices(client, expirationDate, cluster.ClusterArn)
					if err != nil && errors.Is(err, errServiceTooNewError) {
						logger.Printf("skipping %s deletion: cluster service not past expiration date",
							*cluster.ClusterName)
						continue
					} else if err != nil {
						return fmt.Errorf("unable to remove all services in cluster %s due to %w", *cluster.ClusterName, err)
					}
				}
				logger.Printf("Trying to delete cluster %s", *cluster.ClusterName)
				dlci := &ecs.DeleteClusterInput{Cluster: cluster.ClusterArn}
				if _, err = client.DeleteCluster(dlci); err != nil {
					return fmt.Errorf("unable to delete cluster %s due to %w", *cluster.ClusterName, err)
				} else {
					logger.Printf("Deleted cluster %s successfully", *cluster.ClusterName)
				}
			}
		}
		if lco.NextToken == nil {
			break
		}
		nextToken = lco.NextToken
	}
	return nil
}

func deleteServices(client *ecs.ECS, expirationDate time.Time, clusterArn *string) error {
	var nextToken *string
	var remainingServices int
	for {
		lsi := &ecs.ListServicesInput{Cluster: clusterArn, NextToken: nextToken}
		lso, err := client.ListServices(lsi)
		if err != nil {
			return err
		}
		dsi := &ecs.DescribeServicesInput{Cluster: clusterArn, Services: lso.ServiceArns}
		dso, err := client.DescribeServices(dsi)
		if err != nil {
			return err
		}
		remainingServices += len(dso.Services)
		for _, service := range dso.Services {
			if expirationDate.After(*service.CreatedAt) {
				logger.Printf("Trying to delete service %s created on %v", *service.ServiceName, *service.CreatedAt)
				if service.TaskDefinition != nil {
					if err = deleteTaskDefinition(client, expirationDate, service.TaskDefinition); err != nil {
						return err
					}
				}
				dlsi := &ecs.DeleteServiceInput{Cluster: clusterArn, Service: service.ServiceName, Force: aws.Bool(true)}
				if _, err = client.DeleteService(dlsi); err != nil {
					return err
				}
				logger.Printf("Deleted service %s successfully", *service.ServiceName)
				remainingServices--
			} else {
				return errServiceTooNewError
			}

		}
		if lso.NextToken == nil {
			break
		}
		nextToken = lso.NextToken
	}
	if remainingServices > 0 {
		return fmt.Errorf("failed to delete all services")
	}
	return nil
}

func deleteTaskDefinition(client *ecs.ECS, expirationDate time.Time, taskDefinitionArn *string) error {
	dtdi := &ecs.DescribeTaskDefinitionInput{TaskDefinition: taskDefinitionArn}
	dtdo, err := client.DescribeTaskDefinition(dtdi)
	if err != nil {
		return err
	}
	taskDefinition := dtdo.TaskDefinition
	if strings.HasPrefix(*taskDefinition.Family, taskDefPrefix) && expirationDate.After(*taskDefinition.RegisteredAt) {
		logger.Printf("Trying to de-register task definition %s registered on %v", *taskDefinition.Family, *taskDefinition.RegisteredAt)
		drtdi := &ecs.DeregisterTaskDefinitionInput{TaskDefinition: taskDefinitionArn}
		if _, err = client.DeregisterTaskDefinition(drtdi); err != nil {
			return err
		}
		logger.Printf("De-registered task definition %s successfully", *taskDefinition.Family)
	}
	return nil
}
