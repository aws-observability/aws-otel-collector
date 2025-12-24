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

package ebs

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/stretchr/testify/assert"
)

func TestShouldDeleteVolume(t *testing.T) {
	now := time.Now()
	oldTime := now.Add(-10 * 24 * time.Hour)  // 10 days ago
	newTime := now.Add(-1 * 24 * time.Hour)   // 1 day ago
	expirationDate := now.Add(-5 * 24 * time.Hour) // 5 days ago

	tests := []struct {
		name           string
		volume         types.Volume
		expirationDate time.Time
		expected       bool
	}{
		{
			name: "old unattached volume should be deleted",
			volume: types.Volume{
				VolumeId:    aws.String("vol-123"),
				CreateTime:  &oldTime,
				Attachments: []types.VolumeAttachment{},
			},
			expirationDate: expirationDate,
			expected:       true,
		},
		{
			name: "new unattached volume should not be deleted",
			volume: types.Volume{
				VolumeId:    aws.String("vol-123"),
				CreateTime:  &newTime,
				Attachments: []types.VolumeAttachment{},
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "old attached volume should not be deleted",
			volume: types.Volume{
				VolumeId:   aws.String("vol-123"),
				CreateTime: &oldTime,
				Attachments: []types.VolumeAttachment{
					{InstanceId: aws.String("i-123")},
				},
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "new attached volume should not be deleted",
			volume: types.Volume{
				VolumeId:   aws.String("vol-123"),
				CreateTime: &newTime,
				Attachments: []types.VolumeAttachment{
					{InstanceId: aws.String("i-123")},
				},
			},
			expirationDate: expirationDate,
			expected:       false,
		},
		{
			name: "volume with nil create time should not be deleted",
			volume: types.Volume{
				VolumeId:    aws.String("vol-123"),
				CreateTime:  nil,
				Attachments: []types.VolumeAttachment{},
			},
			expirationDate: expirationDate,
			expected:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldDeleteVolume(tc.volume, tc.expirationDate)
			assert.Equal(t, tc.expected, result)
		})
	}
}
