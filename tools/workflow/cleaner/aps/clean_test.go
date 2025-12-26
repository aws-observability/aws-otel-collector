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

package aps

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/stretchr/testify/assert"
)

func TestShouldDelete(t *testing.T) {
	tests := []struct {
		name     string
		ws       types.WorkspaceSummary
		expected bool
	}{
		{
			name: "ephemeral tag set to true",
			ws: types.WorkspaceSummary{
				WorkspaceId: aws.String("ws-123"),
				Tags:        map[string]string{"ephemeral": "true"},
			},
			expected: true,
		},
		{
			name: "ephemeral tag set to false",
			ws: types.WorkspaceSummary{
				WorkspaceId: aws.String("ws-123"),
				Tags:        map[string]string{"ephemeral": "false"},
			},
			expected: false,
		},
		{
			name: "no tags and no alias",
			ws: types.WorkspaceSummary{
				WorkspaceId: aws.String("ws-123"),
				Tags:        map[string]string{},
				Alias:       nil,
			},
			expected: true,
		},
		{
			name: "no tags but has alias",
			ws: types.WorkspaceSummary{
				WorkspaceId: aws.String("ws-123"),
				Tags:        map[string]string{},
				Alias:       aws.String("my-workspace"),
			},
			expected: false,
		},
		{
			name: "no tags and empty alias",
			ws: types.WorkspaceSummary{
				WorkspaceId: aws.String("ws-123"),
				Tags:        map[string]string{},
				Alias:       aws.String(""),
			},
			expected: true,
		},
		{
			name: "has other tags but no ephemeral tag and no alias",
			ws: types.WorkspaceSummary{
				WorkspaceId: aws.String("ws-123"),
				Tags:        map[string]string{"team": "observability"},
				Alias:       nil,
			},
			expected: false,
		},
		{
			name: "has other tags and alias",
			ws: types.WorkspaceSummary{
				WorkspaceId: aws.String("ws-123"),
				Tags:        map[string]string{"team": "observability"},
				Alias:       aws.String("my-workspace"),
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldDelete(tc.ws)
			assert.Equal(t, tc.expected, result)
		})
	}
}
