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
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const Type = "iam"

var (
	logger             = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)
	roleNamePrefixes   = []string{"terraform", "aoc-eks-assume-role", "fargate-profile-role", "lambda", "push-mode-sample-app--"}
	policyNamePrefixes = []string{"terraform", "lambda"}
)

func Clean(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	logger.Printf("Begin to clean IAM roles")
	client := iam.NewFromConfig(cfg)
	if err := deleteRoles(ctx, client, expirationDate); err != nil {
		return err
	}
	if err := deletePolicies(ctx, client, expirationDate); err != nil {
		return err
	}
	return nil
}

func deleteRoles(ctx context.Context, client *iam.Client, expirationDate time.Time) error {
	paginator := iam.NewListRolesPaginator(client, &iam.ListRolesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, role := range output.Roles {
			if hasPrefix(roleNamePrefixes, *role.RoleName) && expirationDate.After(*role.CreateDate) &&
				(role.RoleLastUsed == nil || role.RoleLastUsed.LastUsedDate == nil || expirationDate.After(*role.RoleLastUsed.LastUsedDate)) {
				logger.Printf("Trying to delete IAM role %s created on %v", *role.RoleName, *role.CreateDate)
				if err = detachPolicies(ctx, client, role.RoleName); err != nil {
					return err
				}
				if err = deleteProfiles(ctx, client, expirationDate, role.RoleName); err != nil {
					return err
				}

				_, err = client.DeleteRole(ctx, &iam.DeleteRoleInput{RoleName: role.RoleName})
				if err != nil {
					return err
				}
				logger.Printf("Deleted IAM role %s successfully", *role.RoleName)
			}
		}
	}
	return nil
}

func deletePolicies(ctx context.Context, client *iam.Client, expirationDate time.Time) error {
	paginator := iam.NewListPoliciesPaginator(client, &iam.ListPoliciesInput{Scope: types.PolicyScopeTypeLocal})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, policy := range output.Policies {
			if hasPrefix(policyNamePrefixes, *policy.PolicyName) && expirationDate.After(*policy.CreateDate) && *policy.AttachmentCount == 0 {
				logger.Printf("Trying to delete IAM policy %s created on %v", *policy.PolicyName, *policy.CreateDate)
				_, err = client.DeletePolicy(ctx, &iam.DeletePolicyInput{PolicyArn: policy.Arn})
				if err != nil {
					return err
				}
				logger.Printf("Deleted IAM policy %s successfully", *policy.PolicyName)
			}
		}
	}
	return nil
}

func detachPolicies(ctx context.Context, client *iam.Client, roleName *string) error {
	paginator := iam.NewListAttachedRolePoliciesPaginator(client, &iam.ListAttachedRolePoliciesInput{RoleName: roleName})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, policy := range output.AttachedPolicies {
			_, err = client.DetachRolePolicy(ctx, &iam.DetachRolePolicyInput{PolicyArn: policy.PolicyArn, RoleName: roleName})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func deleteProfiles(ctx context.Context, client *iam.Client, expirationDate time.Time, roleName *string) error {
	paginator := iam.NewListInstanceProfilesForRolePaginator(client, &iam.ListInstanceProfilesForRoleInput{RoleName: roleName})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, profile := range output.InstanceProfiles {
			_, err = client.RemoveRoleFromInstanceProfile(ctx, &iam.RemoveRoleFromInstanceProfileInput{InstanceProfileName: profile.InstanceProfileName, RoleName: roleName})
			if err != nil {
				return err
			}
			if len(profile.Roles) == 1 && expirationDate.After(*profile.CreateDate) {
				logger.Printf("Trying to delete instance profile %s attached to %s created on %v", *profile.InstanceProfileName, *roleName, *profile.CreateDate)
				_, err = client.DeleteInstanceProfile(ctx, &iam.DeleteInstanceProfileInput{InstanceProfileName: profile.InstanceProfileName})
				if err != nil {
					return err
				}
				logger.Printf("Deleted instance profile %s successfully", *profile.InstanceProfileName)
			}
		}
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
