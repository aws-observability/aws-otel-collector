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
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/stretchr/testify/assert"
)

func TestShouldConsiderCluster(t *testing.T) {
	tests := []struct {
		name     string
		cluster  types.Cluster
		expected bool
	}{
		{
			name: "cluster with aoc-testing prefix should be considered",
			cluster: types.Cluster{
				ClusterName: aws.String("aoc-testing-12345"),
				ClusterArn:  aws.String("arn:aws:ecs:us-east-1:123456789012:cluster/aoc-testing-12345"),
			},
			expected: true,
		},
		{
			name: "cluster without aoc-testing prefix should not be considered",
			cluster: types.Cluster{
				ClusterName: aws.String("production-cluster"),
				ClusterArn:  aws.String("arn:aws:ecs:us-east-1:123456789012:cluster/production-cluster"),
			},
			expected: false,
		},
		{
			name: "cluster with nil name should not be considered",
			cluster: types.Cluster{
				ClusterName: nil,
				ClusterArn:  aws.String("arn:aws:ecs:us-east-1:123456789012:cluster/test"),
			},
			expected: false,
		},
		{
			name: "cluster with aoc-testing in middle should not be considered",
			cluster: types.Cluster{
				ClusterName: aws.String("my-aoc-testing-cluster"),
				ClusterArn:  aws.String("arn:aws:ecs:us-east-1:123456789012:cluster/my-aoc-testing-cluster"),
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldConsiderCluster(tc.cluster)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestShouldDeleteService(t *testing.T) {
	now := time.Now()
	oldTime := now.Add(-10 * 24 * time.Hour)  // 10 days ago
	newTime := now.Add(-1 * 24 * time.Hour)   // 1 day ago
	expirationDate := now.Add(-5 * 24 * time.Hour) // 5 days ago

	tests := []struct {
		name           string
		service        types.Service
		expirationDate time.Time
		expected       bool
	}{
		{
			name: "old service should be deleted",
			service: types.Service{
				ServiceName: aws.String("test-service"),
				CreatedAt:   &oldTime,
			},
			expirationDate: expirationDate,
			expected:       true,
		},
		{
			name: "new service should not be deleted",
			service: types.Service{
				ServiceName: aws.String("test-service"),
				CreatedAt:   &newTime,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "service with nil created time should not be deleted",
			service: types.Service{
				ServiceName: aws.String("test-service"),
				CreatedAt:   nil,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "service created exactly at expiration should not be deleted",
			service: types.Service{
				ServiceName: aws.String("test-service"),
				CreatedAt:   &expirationDate,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldDeleteService(tc.service, tc.expirationDate)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestShouldDeregisterTaskDefinition(t *testing.T) {
	now := time.Now()
	oldTime := now.Add(-10 * 24 * time.Hour)  // 10 days ago
	newTime := now.Add(-1 * 24 * time.Hour)   // 1 day ago
	expirationDate := now.Add(-5 * 24 * time.Hour) // 5 days ago

	tests := []struct {
		name           string
		taskDef        types.TaskDefinition
		expirationDate time.Time
		expected       bool
	}{
		{
			name: "old task def with taskdef prefix should be deregistered",
			taskDef: types.TaskDefinition{
				Family:       aws.String("taskdef-test-123"),
				RegisteredAt: &oldTime,
			},
			expirationDate: expirationDate,
			expected:       true,
		},
		{
			name: "new task def with taskdef prefix should not be deregistered",
			taskDef: types.TaskDefinition{
				Family:       aws.String("taskdef-test-123"),
				RegisteredAt: &newTime,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "old task def without taskdef prefix should not be deregistered",
			taskDef: types.TaskDefinition{
				Family:       aws.String("production-service"),
				RegisteredAt: &oldTime,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "task def with nil family should not be deregistered",
			taskDef: types.TaskDefinition{
				Family:       nil,
				RegisteredAt: &oldTime,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "task def with nil registered time should not be deregistered",
			taskDef: types.TaskDefinition{
				Family:       aws.String("taskdef-test-123"),
				RegisteredAt: nil,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "task def registered exactly at expiration should not be deregistered",
			taskDef: types.TaskDefinition{
				Family:       aws.String("taskdef-test-123"),
				RegisteredAt: &expirationDate,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldDeregisterTaskDefinition(tc.taskDef, tc.expirationDate)
			assert.Equal(t, tc.expected, result)
		})
	}
}
