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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

const Type = "autoscaling"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(sess *session.Session, expirationDate time.Time) error {
	logger.Printf("Begin to clean ECS AutoScaling")

	autoscalingclient := autoscaling.New(sess)

	// Get list of autoscaling group
	//Autoscaling's tag filter
	autoscalingTagFilter := autoscaling.Filter{Name: aws.String("tag:Component"), Values: []*string{aws.String("aoc")}}

	//Allow to load all the auto scaling groups since the default respond is paginated auto scaling groups.
	//Look into the documentations and read the starting-token for more details
	//Documentation: https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-auto-scaling-groups.html#options
	var nextToken *string

	for {
		describeAutoScalingInputs := &autoscaling.DescribeAutoScalingGroupsInput{Filters: []*autoscaling.Filter{&autoscalingTagFilter}, NextToken: nextToken}
		describeAutoScalingOutputs, err := autoscalingclient.DescribeAutoScalingGroups(describeAutoScalingInputs)

		if err != nil {
			return err
		}

		for _, asg := range describeAutoScalingOutputs.AutoScalingGroups {
			//Skipping asg that does not older than 5 days
			if !expirationDate.After(*asg.CreatedTime) {
				continue
			}

			deleteAutoScalingGroupInput := &autoscaling.DeleteAutoScalingGroupInput{
				AutoScalingGroupName: asg.AutoScalingGroupName,
				ForceDelete:          aws.Bool(true),
			}

			_, err = autoscalingclient.DeleteAutoScalingGroup(deleteAutoScalingGroupInput)

			if err != nil {
				return err
			}

			deleteLaunchConfigurationInput := &autoscaling.DeleteLaunchConfigurationInput{
				LaunchConfigurationName: asg.LaunchConfigurationName,
			}

			if _, err = autoscalingclient.DeleteLaunchConfiguration(deleteLaunchConfigurationInput); err != nil {
				return err
			}

			logger.Printf("Deleted asg %s successfully", *asg.AutoScalingGroupName)
		}

		if describeAutoScalingOutputs.NextToken == nil {
			break
		}

		nextToken = describeAutoScalingOutputs.NextToken
	}
	return nil
}
