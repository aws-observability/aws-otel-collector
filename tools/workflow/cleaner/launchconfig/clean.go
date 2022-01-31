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
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	Type                             = "launchconfig"
	containLcName                    = "cluster-aoc-testing"
	requiredIAMInstanceProfileLength = 6
)

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(sess *session.Session, expirationDate time.Time) error {
	logger.Printf("Begin to clean ECS Launch Configuration")
	ec2client := ec2.New(sess)
	autoscalingclient := autoscaling.New(sess)

	//Allow to load all the launch configurations since the default respond is paginated launch configurations.
	//Look into the documentations and read the starting-token for more details
	//Documentation: https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-launch-configurations.html#options
	var nextToken *string

	for {
		//Launch Configuration Go SDK currently does not support filter tag or filter wildcard. Only supports with matching name
		//Documentation: https://github.com/aws/aws-sdk-go/blob/02266ed24221ac21bb37d6ac614d1ced95407556/service/autoscaling/api.go#L9577-L9593
		describeLaunchConfigurationInputs := &autoscaling.DescribeLaunchConfigurationsInput{NextToken: nextToken}
		describeLaunchConfigurationOutputs, err := autoscalingclient.DescribeLaunchConfigurations(describeLaunchConfigurationInputs)

		if err != nil {
			return err
		}

		for _, lc := range describeLaunchConfigurationOutputs.LaunchConfigurations {
			//Skipping lc that does not contain cluster-aoc-testing string (relating to aws-otel-test-framework)
			if filterLcNameResult := strings.Contains(*lc.LaunchConfigurationName, containLcName); !filterLcNameResult {
				continue
			}

			//Skipping lc that does not older than 5 days
			if !expirationDate.After(*lc.CreatedTime) {
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
			instanceStateFilter := ec2.Filter{Name: aws.String("instance-state-name"), Values: []*string{aws.String("running")}}
			//Tag filter
			instanceTagFilter := ec2.Filter{Name: aws.String("tag:Name"), Values: []*string{aws.String(ec2InstanceName)}}

			describeInstancesInput := ec2.DescribeInstancesInput{Filters: []*ec2.Filter{&instanceStateFilter, &instanceTagFilter}, NextToken: nextToken}
			describeInstancesOutput, err := ec2client.DescribeInstances(&describeInstancesInput)

			if err != nil {
				return err
			}

			//Confirm whether the ec2 associated with the ecs is still exist or not
			if len(describeInstancesOutput.Reservations) != 0 {
				continue
			}

			logger.Printf("Trying to delete lc %s with launch-date %v", *lc.LaunchConfigurationName, lc.CreatedTime)

			deleteLaunchConfigurationInput := &autoscaling.DeleteLaunchConfigurationInput{
				LaunchConfigurationName: lc.LaunchConfigurationName,
			}

			if _, err = autoscalingclient.DeleteLaunchConfiguration(deleteLaunchConfigurationInput); err != nil {
				return err
			}

			logger.Printf("Deleted launch configuration %s successfully", *lc.LaunchConfigurationName)
		}

		if describeLaunchConfigurationOutputs.NextToken == nil {
			break
		}

		nextToken = describeLaunchConfigurationOutputs.NextToken
	}
	return nil
}
