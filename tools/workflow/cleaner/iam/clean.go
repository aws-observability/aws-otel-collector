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
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

const Type = "iam"

var roleNamePrefixes = []string{"terraform", "aoc-eks-assume-role", "fargate-profile-role"}

func Clean(sess *session.Session, keepDuration time.Duration) {
	client := iam.New(sess)

	var rolesMarker *string

	for {
		lri := &iam.ListRolesInput{Marker: rolesMarker}
		lro, err := client.ListRoles(lri)

		if err != nil {
			log.Fatalf("Failed to get roles because of %v", err)
		}

		for _, role := range lro.Roles {
			var hasPrefix bool
			for _, prefix := range roleNamePrefixes {
				if hasPrefix = strings.HasPrefix(*role.RoleName, prefix); hasPrefix {
					break
				}
			}
			expirationDate := time.Now().UTC().Add(keepDuration)
			if hasPrefix && expirationDate.After(*role.CreateDate) && (role.RoleLastUsed == nil || expirationDate.After(*role.RoleLastUsed.LastUsedDate)) {
				var policiesMarker *string
				for {
					larpi := &iam.ListAttachedRolePoliciesInput{RoleName: role.RoleName, Marker: policiesMarker}
					larpo, err := client.ListAttachedRolePolicies(larpi)

					if err != nil {
						log.Fatalf("Failed to get policies for %s because of %v", *role.RoleName, err)
					}

					for _, policy := range larpo.AttachedPolicies {
						drpi := &iam.DetachRolePolicyInput{PolicyArn: policy.PolicyArn, RoleName: role.RoleName}
						_, err = client.DetachRolePolicy(drpi)
						if err != nil {
							log.Fatalf("Failed to detach policy %s from %s because of %v", *policy.PolicyName, *role.RoleName, err)
						}
					}

					if larpo.Marker == nil {
						break
					}
					policiesMarker = larpo.Marker
				}

				dri := &iam.DeleteRoleInput{RoleName: role.RoleName}
				_, err = client.DeleteRole(dri)
				if err != nil {
					log.Fatalf("Failed to delete %s because of %v", *role.RoleName, err)
				}
			}
		}
		if lro.Marker == nil {
			break
		}
		rolesMarker = lro.Marker
	}
}
