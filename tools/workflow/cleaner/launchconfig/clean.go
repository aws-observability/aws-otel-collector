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

package launchconfig

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	astypes "github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const (
	Type                             = "launchconfig"
	containLcName                    = "cluster-aoc-testing"
	requiredIAMInstanceProfileLength = 6
)

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	logger.Printf("Begin to clean ECS Launch Configuration")
	ec2client := ec2.NewFromConfig(cfg)
	autoscalingclient := autoscaling.NewFromConfig(cfg)

	//Allow to load all the launch configurations since the default respond is paginated launch configurations.
	//Look into the documentations and read the starting-token for more details
	//Documentation: https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-launch-configurations.html#options
	paginator := autoscaling.NewDescribeLaunchConfigurationsPaginator(autoscalingclient, &autoscaling.DescribeLaunchConfigurationsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		for _, lc := range output.LaunchConfigurations {
			if !shouldConsiderLaunchConfig(lc, expirationDate) {
				continue
			}

			//Split the instance profile to get the testing ID associate with the testing framework
			splitLcIamInstanceProfile := strings.Split(*lc.IamInstanceProfile, "-")

			if len(splitLcIamInstanceProfile) != requiredIAMInstanceProfileLength {
				continue
			}

			//Check if the ec2 instance which is created by ecs is still exist with this launch configuration
			//In order to confirm if these load configurations are still being used

			testingId := splitLcIamInstanceProfile[5]
			ec2InstanceName := "cluster-worker-aoc-testing-" + testingId

			//State filter
			instanceStateFilter := ec2types.Filter{Name: aws.String("instance-state-name"), Values: []string{"running"}}
			//Tag filter
			instanceTagFilter := ec2types.Filter{Name: aws.String("tag:Name"), Values: []string{ec2InstanceName}}

			describeInstancesOutput, err := ec2client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
				Filters: []ec2types.Filter{instanceStateFilter, instanceTagFilter},
			})

			if err != nil {
				return err
			}

			//Confirm whether the ec2 associated with the ecs is still exist or not
			if len(describeInstancesOutput.Reservations) != 0 {
				continue
			}

			logger.Printf("Trying to delete lc %s with launch-date %v", *lc.LaunchConfigurationName, lc.CreatedTime)

			_, err = autoscalingclient.DeleteLaunchConfiguration(ctx, &autoscaling.DeleteLaunchConfigurationInput{
				LaunchConfigurationName: lc.LaunchConfigurationName,
			})

			if err != nil {
				return err
			}

			logger.Printf("Deleted launch configuration %s successfully", *lc.LaunchConfigurationName)
		}
	}
	return nil
}

// shouldConsiderLaunchConfig returns true if the launch config name contains the expected prefix
// and is older than the expiration date.
func shouldConsiderLaunchConfig(lc astypes.LaunchConfiguration, expirationDate time.Time) bool {
	if lc.LaunchConfigurationName == nil || lc.CreatedTime == nil {
		return false
	}
	return strings.Contains(*lc.LaunchConfigurationName, containLcName) && expirationDate.After(*lc.CreatedTime)
}
