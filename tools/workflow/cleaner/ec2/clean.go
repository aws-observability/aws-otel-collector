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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const Type = "ec2"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(sess *session.Session, expirationDate time.Time) error {
	ec2client := ec2.New(sess)

	if err := cleanSSHKeys(ec2client, expirationDate); err != nil {
		return err
	}
	return cleanInstances(ec2client, expirationDate)
}

func cleanSSHKeys(ec2client *ec2.EC2, expirationDate time.Time) error {
	logger.Printf("Begin to clean SSH Keys")
	tagFilter := &ec2.Filter{Name: aws.String("tag:ephemeral"), Values: []*string{aws.String("true")}}

	pairs, err := ec2client.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{Filters: []*ec2.Filter{tagFilter}})
	if err != nil {
		return fmt.Errorf("unable to describe SSH keys: %w", err)
	}

	for _, kpi := range pairs.KeyPairs {
		if expirationDate.After(*kpi.CreateTime) {
			logger.Printf("Try to delete key pair %q (%s) created on %v", *kpi.KeyName, *kpi.KeyPairId, *kpi.CreateTime)
			if _, err := ec2client.DeleteKeyPair(&ec2.DeleteKeyPairInput{KeyPairId: kpi.KeyPairId}); err != nil {
				return fmt.Errorf("error deleting key pair: %w", err)
			}
		}
	}

	return nil
}

func cleanInstances(ec2client *ec2.EC2, expirationDate time.Time) error {
	logger.Printf("Begin to clean EC2 Instances")
	// Get list of instance
	//State filter
	instanceStateFilter := ec2.Filter{Name: aws.String("instance-state-name"), Values: []*string{aws.String("running")}}

	//Tag filter
	instanceTagFilter := ec2.Filter{Name: aws.String("tag:Name"), Values: []*string{aws.String("Integ-test-Sample-App"), aws.String("Integ-test-aoc")}}

	//get instances to delete
	var deleteInstanceIds []*string
	var nextToken *string
	for {
		describeInstancesInput := ec2.DescribeInstancesInput{Filters: []*ec2.Filter{&instanceStateFilter, &instanceTagFilter}, NextToken: nextToken}
		describeInstancesOutput, err := ec2client.DescribeInstances(&describeInstancesInput)

		if err != nil {
			return err
		}

		for _, reservation := range describeInstancesOutput.Reservations {
			for _, instance := range reservation.Instances {
				//only delete instance if older than 5 days
				if expirationDate.After(*instance.LaunchTime) {
					logger.Printf("Try to delete instance %s tags %v launch-date %v", *instance.InstanceId, instance.Tags, instance.LaunchTime)
					deleteInstanceIds = append(deleteInstanceIds, instance.InstanceId)
				}
			}
		}
		if describeInstancesOutput.NextToken == nil {
			break
		}
		nextToken = describeInstancesOutput.NextToken
	}

	if len(deleteInstanceIds) < 1 {
		logger.Printf("No instances to delete")
		return nil
	}

	if _, err := ec2client.TerminateInstances(&ec2.TerminateInstancesInput{InstanceIds: deleteInstanceIds}); err != nil {
		return fmt.Errorf("error terminating instances: %w", err)
	}
	return nil
}
