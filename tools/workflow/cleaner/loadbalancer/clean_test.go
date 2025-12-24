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
