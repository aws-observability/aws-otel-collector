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

package iam

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		name          string
		prefixOptions []string
		candidate     string
		expected      bool
	}{
		{
			name:          "matches terraform prefix",
			prefixOptions: roleNamePrefixes,
			candidate:     "terraform-20241223-abcd",
			expected:      true,
		},
		{
			name:          "matches aoc-eks-assume-role prefix",
			prefixOptions: roleNamePrefixes,
			candidate:     "aoc-eks-assume-role-12345",
			expected:      true,
		},
		{
			name:          "matches fargate-profile-role prefix",
			prefixOptions: roleNamePrefixes,
			candidate:     "fargate-profile-role-test",
			expected:      true,
		},
		{
			name:          "matches lambda prefix",
			prefixOptions: roleNamePrefixes,
			candidate:     "lambda-execution-role",
			expected:      true,
		},
		{
			name:          "matches push-mode-sample-app prefix",
			prefixOptions: roleNamePrefixes,
			candidate:     "push-mode-sample-app--12345",
			expected:      true,
		},
		{
			name:          "does not match any prefix",
			prefixOptions: roleNamePrefixes,
			candidate:     "some-random-role",
			expected:      false,
		},
		{
			name:          "empty candidate",
			prefixOptions: roleNamePrefixes,
			candidate:     "",
			expected:      false,
		},
		{
			name:          "exact match of prefix",
			prefixOptions: []string{"test"},
			candidate:     "test",
			expected:      true,
		},
		{
			name:          "policy prefix terraform",
			prefixOptions: policyNamePrefixes,
			candidate:     "terraform-policy-123",
			expected:      true,
		},
		{
			name:          "policy prefix lambda",
			prefixOptions: policyNamePrefixes,
			candidate:     "lambda-execution-policy",
			expected:      true,
		},
		{
			name:          "policy does not match prefix",
			prefixOptions: policyNamePrefixes,
			candidate:     "custom-policy",
			expected:      false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := hasPrefix(tc.prefixOptions, tc.candidate)
			assert.Equal(t, tc.expected, result)
		})
	}
}
