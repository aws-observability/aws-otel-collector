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
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
)

const Type = "autoscaling"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	logger.Printf("Begin to clean ECS AutoScaling")

	autoscalingclient := autoscaling.NewFromConfig(cfg)

	// Get list of autoscaling group
	//Autoscaling's tag filter
	autoscalingTagFilter := types.Filter{Name: aws.String("tag:Component"), Values: []string{"aoc"}}

	//Allow to load all the auto scaling groups since the default respond is paginated auto scaling groups.
	//Look into the documentations and read the starting-token for more details
	//Documentation: https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-auto-scaling-groups.html#options
	paginator := autoscaling.NewDescribeAutoScalingGroupsPaginator(autoscalingclient, &autoscaling.DescribeAutoScalingGroupsInput{
		Filters: []types.Filter{autoscalingTagFilter},
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		for _, asg := range output.AutoScalingGroups {
			if !shouldDeleteAutoScalingGroup(asg, expirationDate) {
				continue
			}
			logger.Printf("deleting asg %s", *asg.AutoScalingGroupName)

			_, err = autoscalingclient.DeleteAutoScalingGroup(ctx, &autoscaling.DeleteAutoScalingGroupInput{
				AutoScalingGroupName: asg.AutoScalingGroupName,
				ForceDelete:          aws.Bool(true),
			})

			if err != nil {
				return err
			}

			logger.Printf("Deleted asg %s successfully", *asg.AutoScalingGroupName)
		}
	}
	return nil
}

// shouldDeleteAutoScalingGroup returns true if the ASG is older than the expiration date.
func shouldDeleteAutoScalingGroup(asg types.AutoScalingGroup, expirationDate time.Time) bool {
	return asg.CreatedTime != nil && expirationDate.After(*asg.CreatedTime)
}
