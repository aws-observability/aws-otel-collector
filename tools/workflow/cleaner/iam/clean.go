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
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

const Type = "iam"

var (
	logger             = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)
	roleNamePrefixes   = []string{"terraform", "aoc-eks-assume-role", "fargate-profile-role", "lambda", "push-mode-sample-app--"}
	policyNamePrefixes = []string{"terraform", "lambda"}
)

func Clean(sess *session.Session, expirationDate time.Time) error {
	logger.Printf("Begin to clean IAM roles")
	client := iam.New(sess)
	if err := deleteRoles(client, expirationDate); err != nil {
		return err
	}
	if err := deletePolicies(client, expirationDate); err != nil {
		return err
	}
	return nil
}

func deleteRoles(client *iam.IAM, expirationDate time.Time) error {
	var marker *string
	for {
		lri := &iam.ListRolesInput{Marker: marker}
		lro, err := client.ListRoles(lri)
		if err != nil {
			return err
		}
		for _, role := range lro.Roles {
			if hasPrefix(roleNamePrefixes, *role.RoleName) && expirationDate.After(*role.CreateDate) &&
				(role.RoleLastUsed == nil || expirationDate.After(*role.RoleLastUsed.LastUsedDate)) {
				logger.Printf("Trying to delete IAM role %s created on %v", *role.RoleName, *role.CreateDate)
				if err = detachPolicies(client, role.RoleName); err != nil {
					return err
				}
				if err = deleteProfiles(client, expirationDate, role.RoleName); err != nil {
					return err
				}

				dri := &iam.DeleteRoleInput{RoleName: role.RoleName}
				if _, err = client.DeleteRole(dri); err != nil {
					return err
				}
				logger.Printf("Deleted IAM role %s successfully", *role.RoleName)
			}
		}
		if lro.Marker == nil {
			break
		}
		marker = lro.Marker
	}
	return nil
}

func deletePolicies(client *iam.IAM, expirationDate time.Time) error {
	var marker *string
	for {
		lpi := &iam.ListPoliciesInput{Marker: marker, Scope: aws.String(iam.PolicyScopeTypeLocal)}
		lpo, err := client.ListPolicies(lpi)
		if err != nil {
			return err
		}
		for _, policy := range lpo.Policies {
			if hasPrefix(policyNamePrefixes, *policy.PolicyName) && expirationDate.After(*policy.CreateDate) && *policy.AttachmentCount == 0 {
				logger.Printf("Trying to delete IAM policy %s created on %v", *policy.PolicyName, *policy.CreateDate)
				dpi := &iam.DeletePolicyInput{PolicyArn: policy.Arn}
				if _, err = client.DeletePolicy(dpi); err != nil {
					return err
				}
				logger.Printf("Deleted IAM policy %s successfully", *policy.PolicyName)
			}
		}
		if lpo.Marker == nil {
			break
		}
		marker = lpo.Marker
	}
	return nil
}

func detachPolicies(client *iam.IAM, roleName *string) error {
	var marker *string
	for {
		larpi := &iam.ListAttachedRolePoliciesInput{RoleName: roleName, Marker: marker}
		larpo, err := client.ListAttachedRolePolicies(larpi)
		if err != nil {
			return err
		}
		for _, policy := range larpo.AttachedPolicies {
			drpi := &iam.DetachRolePolicyInput{PolicyArn: policy.PolicyArn, RoleName: roleName}
			_, err = client.DetachRolePolicy(drpi)
			if err != nil {
				return err
			}
		}
		if larpo.Marker == nil {
			break
		}
		marker = larpo.Marker
	}
	return nil
}

func deleteProfiles(client *iam.IAM, expirationDate time.Time, roleName *string) error {
	var marker *string
	for {
		lipfri := &iam.ListInstanceProfilesForRoleInput{RoleName: roleName, Marker: marker}
		lipfro, err := client.ListInstanceProfilesForRole(lipfri)
		if err != nil {
			return err
		}
		for _, profile := range lipfro.InstanceProfiles {
			rrfipi := &iam.RemoveRoleFromInstanceProfileInput{InstanceProfileName: profile.InstanceProfileName, RoleName: roleName}
			if _, err = client.RemoveRoleFromInstanceProfile(rrfipi); err != nil {
				return err
			}
			if len(profile.Roles) == 1 && expirationDate.After(*profile.CreateDate) {
				logger.Printf("Trying to delete instance profile %s attached to %s created on %v", *profile.InstanceProfileName, *roleName, *profile.CreateDate)
				dpi := &iam.DeleteInstanceProfileInput{InstanceProfileName: profile.InstanceProfileName}
				if _, err = client.DeleteInstanceProfile(dpi); err != nil {
					return err
				}
				logger.Printf("Deleted instance profile %s successfully", *profile.InstanceProfileName)
			}
		}
		if lipfro.Marker == nil {
			break
		}
		marker = lipfro.Marker
	}
	return nil
}

func hasPrefix(prefixOptions []string, candidate string) bool {
	var result bool
	for _, prefix := range prefixOptions {
		if result = strings.HasPrefix(candidate, prefix); result {
			break
		}
	}
	return result
}
