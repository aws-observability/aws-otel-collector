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

package loadbalancer

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	classictypes "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/stretchr/testify/assert"
)

func TestShouldDeleteLoadBalancer(t *testing.T) {
	now := time.Now()
	oldTime := now.Add(-10 * 24 * time.Hour)  // 10 days ago
	newTime := now.Add(-1 * 24 * time.Hour)   // 1 day ago
	expirationDate := now.Add(-5 * 24 * time.Hour) // 5 days ago

	tests := []struct {
		name           string
		lb             types.LoadBalancer
		expirationDate time.Time
		expected       bool
	}{
		{
			name: "old aoc-lb should be deleted",
			lb: types.LoadBalancer{
				LoadBalancerName: aws.String("my-aoc-lb-123"),
				CreatedTime:      &oldTime,
			},
			expirationDate: expirationDate,
			expected:       true,
		},
		{
			name: "new aoc-lb should not be deleted",
			lb: types.LoadBalancer{
				LoadBalancerName: aws.String("my-aoc-lb-123"),
				CreatedTime:      &newTime,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "old non-aoc lb should not be deleted",
			lb: types.LoadBalancer{
				LoadBalancerName: aws.String("production-lb"),
				CreatedTime:      &oldTime,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "lb with nil name should not be deleted",
			lb: types.LoadBalancer{
				LoadBalancerName: nil,
				CreatedTime:      &oldTime,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "lb with nil created time should not be deleted",
			lb: types.LoadBalancer{
				LoadBalancerName: aws.String("my-aoc-lb-123"),
				CreatedTime:      nil,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldDeleteLoadBalancer(tc.lb, tc.expirationDate)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestShouldDeleteClassicLoadBalancer(t *testing.T) {
	now := time.Now()
	oldTime := now.Add(-10 * 24 * time.Hour)
	newTime := now.Add(-1 * 24 * time.Hour)
	expirationDate := now.Add(-5 * 24 * time.Hour)

	tests := []struct {
		name           string
		lb             classictypes.LoadBalancerDescription
		expirationDate time.Time
		expected       bool
	}{
		{
			name: "old zero-backend lb (k8s leak) should be deleted",
			lb: classictypes.LoadBalancerDescription{
				LoadBalancerName: aws.String("a1b2c3d4e5f6g7h8i9j0"),
				CreatedTime:      &oldTime,
				Instances:        nil,
			},
			expirationDate: expirationDate,
			expected:       true,
		},
		{
			name: "young zero-backend lb should not be deleted",
			lb: classictypes.LoadBalancerDescription{
				LoadBalancerName: aws.String("a1b2c3d4e5f6g7h8i9j0"),
				CreatedTime:      &newTime,
				Instances:        nil,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "old lb with backends should not be deleted",
			lb: classictypes.LoadBalancerDescription{
				LoadBalancerName: aws.String("production-elb"),
				CreatedTime:      &oldTime,
				Instances: []classictypes.Instance{
					{InstanceId: aws.String("i-0123456789abcdef0")},
				},
			},
			expirationDate: expirationDate,
			expected:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldDeleteClassicLoadBalancer(tc.lb, tc.expirationDate)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestShouldDeleteTargetGroup(t *testing.T) {
	tests := []struct {
		name     string
		tg       types.TargetGroup
		expected bool
	}{
		{
			name: "aoc-lb orphan tg should be deleted",
			tg: types.TargetGroup{
				TargetGroupName:  aws.String("aoc-lbtg-0124db2f4f513571"),
				LoadBalancerArns: []string{},
			},
			expected: true,
		},
		{
			name: "aoc-lb tg with parent lb should not be deleted",
			tg: types.TargetGroup{
				TargetGroupName:  aws.String("aoc-lbtg-0124db2f4f513571"),
				LoadBalancerArns: []string{"arn:aws:elasticloadbalancing:us-west-2:111122223333:loadbalancer/app/aoc-lb-x/abc"},
			},
			expected: false,
		},
		{
			name: "non-aoc orphan tg should not be deleted",
			tg: types.TargetGroup{
				TargetGroupName:  aws.String("production-tg"),
				LoadBalancerArns: []string{},
			},
			expected: false,
		},
		{
			name: "tg with nil name should not be deleted",
			tg: types.TargetGroup{
				TargetGroupName:  nil,
				LoadBalancerArns: []string{},
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldDeleteTargetGroup(tc.tg)
			assert.Equal(t, tc.expected, result)
		})
	}
}
