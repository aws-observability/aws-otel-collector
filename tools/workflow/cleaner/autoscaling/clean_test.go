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

package autoscaling

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/stretchr/testify/assert"
)

func TestShouldDeleteAutoScalingGroup(t *testing.T) {
	now := time.Now()
	oldTime := now.Add(-10 * 24 * time.Hour)  // 10 days ago
	newTime := now.Add(-1 * 24 * time.Hour)   // 1 day ago
	expirationDate := now.Add(-5 * 24 * time.Hour) // 5 days ago

	tests := []struct {
		name           string
		asg            types.AutoScalingGroup
		expirationDate time.Time
		expected       bool
	}{
		{
			name: "old ASG should be deleted",
			asg: types.AutoScalingGroup{
				AutoScalingGroupName: aws.String("aoc-asg-123"),
				CreatedTime:          &oldTime,
			},
			expirationDate: expirationDate,
			expected:       true,
		},
		{
			name: "new ASG should not be deleted",
			asg: types.AutoScalingGroup{
				AutoScalingGroupName: aws.String("aoc-asg-123"),
				CreatedTime:          &newTime,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "ASG with nil created time should not be deleted",
			asg: types.AutoScalingGroup{
				AutoScalingGroupName: aws.String("aoc-asg-123"),
				CreatedTime:          nil,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "ASG created exactly at expiration should not be deleted",
			asg: types.AutoScalingGroup{
				AutoScalingGroupName: aws.String("aoc-asg-123"),
				CreatedTime:          &expirationDate,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldDeleteAutoScalingGroup(tc.asg, tc.expirationDate)
			assert.Equal(t, tc.expected, result)
		})
	}
}
