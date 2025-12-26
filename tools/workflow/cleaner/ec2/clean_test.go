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

package ec2

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/stretchr/testify/assert"
)

func TestShouldDeleteKeyPair(t *testing.T) {
	now := time.Now()
	oldTime := now.Add(-10 * 24 * time.Hour)  // 10 days ago
	newTime := now.Add(-1 * 24 * time.Hour)   // 1 day ago
	expirationDate := now.Add(-5 * 24 * time.Hour) // 5 days ago

	tests := []struct {
		name           string
		kpi            types.KeyPairInfo
		expirationDate time.Time
		expected       bool
	}{
		{
			name: "old key pair should be deleted",
			kpi: types.KeyPairInfo{
				KeyName:   aws.String("test-key"),
				KeyPairId: aws.String("key-123"),
				CreateTime: &oldTime,
			},
			expirationDate: expirationDate,
			expected:       true,
		},
		{
			name: "new key pair should not be deleted",
			kpi: types.KeyPairInfo{
				KeyName:   aws.String("test-key"),
				KeyPairId: aws.String("key-123"),
				CreateTime: &newTime,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "key pair with nil create time should not be deleted",
			kpi: types.KeyPairInfo{
				KeyName:   aws.String("test-key"),
				KeyPairId: aws.String("key-123"),
				CreateTime: nil,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldDeleteKeyPair(tc.kpi, tc.expirationDate)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestShouldDeleteInstance(t *testing.T) {
	now := time.Now()
	oldTime := now.Add(-10 * 24 * time.Hour)  // 10 days ago
	newTime := now.Add(-1 * 24 * time.Hour)   // 1 day ago
	expirationDate := now.Add(-5 * 24 * time.Hour) // 5 days ago

	tests := []struct {
		name           string
		instance       types.Instance
		expirationDate time.Time
		expected       bool
	}{
		{
			name: "old instance should be deleted",
			instance: types.Instance{
				InstanceId: aws.String("i-123"),
				LaunchTime: &oldTime,
			},
			expirationDate: expirationDate,
			expected:       true,
		},
		{
			name: "new instance should not be deleted",
			instance: types.Instance{
				InstanceId: aws.String("i-123"),
				LaunchTime: &newTime,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "instance with nil launch time should not be deleted",
			instance: types.Instance{
				InstanceId: aws.String("i-123"),
				LaunchTime: nil,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "instance launched exactly at expiration should not be deleted",
			instance: types.Instance{
				InstanceId: aws.String("i-123"),
				LaunchTime: &expirationDate,
			},
			expirationDate: expirationDate,
			expected:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldDeleteInstance(tc.instance, tc.expirationDate)
			assert.Equal(t, tc.expected, result)
		})
	}
}
