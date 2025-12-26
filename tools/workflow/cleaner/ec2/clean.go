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
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const Type = "ec2"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	ec2client := ec2.NewFromConfig(cfg)

	if err := cleanSSHKeys(ctx, ec2client, expirationDate); err != nil {
		return err
	}
	return cleanInstances(ctx, ec2client, expirationDate)
}

func cleanSSHKeys(ctx context.Context, ec2client *ec2.Client, expirationDate time.Time) error {
	logger.Printf("Begin to clean SSH Keys")
	tagFilter := types.Filter{Name: aws.String("tag:ephemeral"), Values: []string{"true"}}

	pairs, err := ec2client.DescribeKeyPairs(ctx, &ec2.DescribeKeyPairsInput{Filters: []types.Filter{tagFilter}})
	if err != nil {
		return fmt.Errorf("unable to describe SSH keys: %w", err)
	}

	for _, kpi := range pairs.KeyPairs {
		if shouldDeleteKeyPair(kpi, expirationDate) {
			logger.Printf("Try to delete key pair %q (%s) created on %v", *kpi.KeyName, *kpi.KeyPairId, *kpi.CreateTime)
			if _, err := ec2client.DeleteKeyPair(ctx, &ec2.DeleteKeyPairInput{KeyPairId: kpi.KeyPairId}); err != nil {
				return fmt.Errorf("error deleting key pair: %w", err)
			}
		}
	}

	return nil
}

func cleanInstances(ctx context.Context, ec2client *ec2.Client, expirationDate time.Time) error {
	logger.Printf("Begin to clean EC2 Instances")
	// Get list of instance
	//State filter
	instanceStateFilter := types.Filter{Name: aws.String("instance-state-name"), Values: []string{"running"}}

	//Tag filter
	instanceTagFilter := types.Filter{Name: aws.String("tag:Name"), Values: []string{"Integ-test-Sample-App", "Integ-test-aoc"}}

	//get instances to delete
	var deleteInstanceIds []string
	paginator := ec2.NewDescribeInstancesPaginator(ec2client, &ec2.DescribeInstancesInput{
		Filters: []types.Filter{instanceStateFilter, instanceTagFilter},
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		for _, reservation := range output.Reservations {
			for _, instance := range reservation.Instances {
				if shouldDeleteInstance(instance, expirationDate) {
					logger.Printf("Try to delete instance %s tags %v launch-date %v", *instance.InstanceId, instance.Tags, instance.LaunchTime)
					deleteInstanceIds = append(deleteInstanceIds, *instance.InstanceId)
				}
			}
		}
	}

	if len(deleteInstanceIds) < 1 {
		logger.Printf("No instances to delete")
		return nil
	}

	if _, err := ec2client.TerminateInstances(ctx, &ec2.TerminateInstancesInput{InstanceIds: deleteInstanceIds}); err != nil {
		return fmt.Errorf("error terminating instances: %w", err)
	}
	return nil
}

// shouldDeleteKeyPair returns true if the key pair is older than the expiration date.
func shouldDeleteKeyPair(kpi types.KeyPairInfo, expirationDate time.Time) bool {
	return kpi.CreateTime != nil && expirationDate.After(*kpi.CreateTime)
}

// shouldDeleteInstance returns true if the instance is older than the expiration date.
func shouldDeleteInstance(instance types.Instance, expirationDate time.Time) bool {
	return instance.LaunchTime != nil && expirationDate.After(*instance.LaunchTime)
}
