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
				if *cluster.ActiveServicesCount > 0 {
					allServicesDeleted, err := deleteServices(client, expirationDate, cluster.ClusterArn)

					if allServicesDeleted && err == nil {
						logger.Printf("Trying to delete cluster %s", *cluster.ClusterName)
						dlci := &ecs.DeleteClusterInput{Cluster: cluster.ClusterArn}
						if _, err = client.DeleteCluster(dlci); err != nil {
							logger.Printf("Unable to delete cluster %s due to %v", *cluster.ClusterName, err)
						} else {
							logger.Printf("Deleted cluster %s successfully", *cluster.ClusterName)
						}
					} else {
						logger.Printf("Unable to remove all services. Skipping cluster deletion %s due to %v", *cluster.ClusterName, err)
					}
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

func deleteServices(client *ecs.ECS, expirationDate time.Time, clusterArn *string) (bool, error) {
	var nextToken *string
	var remainingServices int
	for {
		lsi := &ecs.ListServicesInput{Cluster: clusterArn, NextToken: nextToken}
		lso, err := client.ListServices(lsi)
		if err != nil {
			return false, err
		}
		dsi := &ecs.DescribeServicesInput{Cluster: clusterArn, Services: lso.ServiceArns}
		dso, err := client.DescribeServices(dsi)
		if err != nil {
			return false, err
		}
		remainingServices += len(dso.Services)
		for _, service := range dso.Services {
			if expirationDate.After(*service.CreatedAt) {
				logger.Printf("Trying to delete service %s created on %v", *service.ServiceName, *service.CreatedAt)
				if service.TaskDefinition != nil {
					if err = deleteTaskDefinition(client, expirationDate, service.TaskDefinition); err != nil {
						return false, err
					}
				}
				dlsi := &ecs.DeleteServiceInput{Cluster: clusterArn, Service: service.ServiceName, Force: aws.Bool(true)}
				if _, err = client.DeleteService(dlsi); err != nil {
					return false, err
				}
				logger.Printf("Deleted service %s successfully", *service.ServiceName)
				remainingServices--
			}
		}
		if lso.NextToken == nil {
			break
		}
		nextToken = lso.NextToken
	}
	return remainingServices == 0, nil
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
