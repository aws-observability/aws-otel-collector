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

package main

import (
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

const (
	containLcName 			 = "cluster-aoc-testing"
	containLbName            = "aoc-lb"
	pastDayDelete            = 5
	pastDayDeleteCalculation = -1 * time.Hour * 24 * pastDayDelete //Currently, deleting resources over 5 days
)

func main() {
	log.Printf("Begin to terminate EC2 Instances")
	terminateEc2Instances()

	log.Printf("Begin to destroy ECS's AutoScaling")
	destroyECSAutoScaling()

	log.Printf("Begin to destroy ECS's Launch Configuration")
	destroyECSLaunchConfiguration()
	
	log.Printf("Begin to destroy Load Balancer resources")
	destroyLoadBalancerResource()

	log.Printf("Finish destroy AWS resources")
}

func terminateEc2Instances() {
	//Set up aws go sdk ec2 client
	testSession, err := session.NewSession()

	if err != nil {
		log.Fatalf("Error creating session %v", err)
	}

	ec2client := ec2.New(testSession)

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
			log.Fatalf("Failed to get instance for error %v", err)
		}

		for _, reservation := range describeInstancesOutput.Reservations {
			for _, instance := range reservation.Instances {
				//only delete instance if older than 5 days
				if time.Now().UTC().Add(pastDayDeleteCalculation).After(*instance.LaunchTime) {
					log.Printf("Try to delete instance %s tags %v launch-date %v", *instance.InstanceId, instance.Tags, instance.LaunchTime)
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
		log.Printf("No instances to delete")
		return
	}

	terminateInstancesInput := ec2.TerminateInstancesInput{InstanceIds: deleteInstanceIds}
	_, err = ec2client.TerminateInstances(&terminateInstancesInput)
	if err != nil {
		log.Fatalf("Failed to terminate instances %v because of %v", deleteInstanceIds, err)
	}
}

func destroyECSAutoScaling(){
	//Set up aws go sdk ec2 client
	testSession, err := session.NewSession()

	if err != nil {
		log.Fatalf("Error creating session %v", err)
	}

	autoscalingclient := autoscaling.New(testSession)

	// Get list of autoscaling group
	//Autoscaling's tag filter
	autoscalingTagFilter := autoscaling.Filter{Name: aws.String("tag:Component"), Values: []*string{aws.String("aoc")}}

	//Allow to load all the launch configurations since the default respond is paginated launch configurations.
	//Look into the documentations and read the starting-token for more details
	//Documentation: https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-auto-scaling-groups.html#options
	var nextToken *string

	for {
		describeAutoScalingInputs := &autoscaling.DescribeAutoScalingGroupsInput{Filters:[]*autoscaling.Filter{&autoscalingTagFilter} ,NextToken: nextToken}
		describeAutoScalingOutputs, err := autoscalingclient.DescribeAutoScalingGroups(describeAutoScalingInputs)

		if err != nil {
			log.Fatalf("Failed to get metadata from launch configuration because of %v", err)
		}

		for _, asg := range describeAutoScalingOutputs.AutoScalingGroups {
			//Skipping lc that does not older than 5 days
			if !time.Now().UTC().Add(pastDayDeleteCalculation).After(*asg.CreatedTime) {
				continue
			}

			deleteAutoScalingGroupInput := &autoscaling.DeleteAutoScalingGroupInput{
				AutoScalingGroupName: asg.AutoScalingGroupName,
				ForceDelete:          aws.Bool(true),
			}

			_, err = autoscalingclient.DeleteAutoScalingGroup(deleteAutoScalingGroupInput)

			if err != nil {
				log.Fatalf("Failed to delete asg %s because of %v", *asg.AutoScalingGroupName, err)
			}

			deleteLaunchConfigurationInput := &autoscaling.DeleteLaunchConfigurationInput{
				LaunchConfigurationName: asg.LaunchConfigurationName,
			}

			_, err = autoscalingclient.DeleteLaunchConfiguration(deleteLaunchConfigurationInput)

			if err != nil {
				log.Fatalf("Failed to delete lc of asg %s  because of %v", *asg.AutoScalingGroupName, err)
			}
			
			log.Printf("Delete asg %s successfully", *asg.AutoScalingGroupName)
		}

		if describeAutoScalingOutputs.NextToken == nil {
			break
		}

		nextToken = describeAutoScalingOutputs.NextToken
	}
}

func destroyECSLaunchConfiguration(){
	//Set up aws go sdk ec2 client
	testSession, err := session.NewSession()

	if err != nil {
		log.Fatalf("Error creating session %v", err)
	}

	ec2client := ec2.New(testSession)
	autoscalingclient := autoscaling.New(testSession)
	
	
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
			log.Fatalf("Failed to get metadata from launch configuration because of %v", err)
		}

		for _, lc := range describeLaunchConfigurationOutputs.LaunchConfigurations {

			//Skipping lc that does not contain cluster-aoc-testing string (relating to aws-otel-test-framework)
			if filterLcNameResult := strings.Contains(*lc.LaunchConfigurationName, containLcName); !filterLcNameResult {
				continue
			}

			//Skipping lc that does not older than 5 days
			if !time.Now().UTC().Add(pastDayDeleteCalculation).After(*lc.CreatedTime) {
				continue
			}

			//Split the instance profile to get the testing ID associate with the testing framework
			splitLcIamInstanceProfile := strings.Split(*lc.IamInstanceProfile, "-")

			if len(splitLcIamInstanceProfile) != 6 {
				continue
			}

			//Check if the ec2 instance which is created by ecs is still exist with this launch configuration
			//In order to confirm if these load configurations are still being used

			var testingId string = splitLcIamInstanceProfile[5]
			var ec2InstanceName string = "cluster-worker-aoc-testing-" + testingId

			//State filter
			instanceStateFilter := ec2.Filter{Name: aws.String("instance-state-name"), Values: []*string{aws.String("running")}}
			//Tag filter
			instanceTagFilter := ec2.Filter{Name: aws.String("tag:Name"), Values: []*string{aws.String(ec2InstanceName)}}

			describeInstancesInput := ec2.DescribeInstancesInput{Filters: []*ec2.Filter{&instanceStateFilter, &instanceTagFilter}, NextToken: nextToken}
			describeInstancesOutput, err := ec2client.DescribeInstances(&describeInstancesInput)

			if err != nil {
				log.Fatalf("Failed to get metadata from EC2 instance", err)
			}

			if  len(describeInstancesOutput.Reservations) != 0 {
				continue
			}

			log.Printf("Trying to delete lc %s with launch-date %v", *lc.LaunchConfigurationName, lc.CreatedTime)

			deleteLaunchConfigurationInput := &autoscaling.DeleteLaunchConfigurationInput{
				LaunchConfigurationName: lc.LaunchConfigurationName,
			}

			_, err = autoscalingclient.DeleteLaunchConfiguration(deleteLaunchConfigurationInput)

			if err != nil {
				log.Fatalf("Failed to delete lc %s because of %v", *lc.LaunchConfigurationName, err)
			}

			log.Printf("Delete lc %s successfully", *lc.LaunchConfigurationName)

		}

		if describeLaunchConfigurationOutputs.NextToken == nil {
			break
		}
		
		nextToken = describeLaunchConfigurationOutputs.NextToken
	}

}

func destroyLoadBalancerResource() {
	//Set up aws go sdk ec2 client
	testSession, err := session.NewSession()

	if err != nil {
		log.Fatalf("Error creating session %v", err)
	}

	elbv2client := elbv2.New(testSession)

	//Allow to load all the load balancers since the default respond is paginated load balancers.
	//Look into the documentations and read the starting-token for more details
	//Documentation: https://docs.aws.amazon.com/cli/latest/reference/elbv2/describe-load-balancers.html#options
	var nextMarker *string

	for {
		//ELB Go SDK currently does not support filter tag or filter wildcard. Only supports with matching name
		//Documentation: https://github.com/aws/aws-sdk-go/blob/02266ed24221ac21bb37d6ac614d1ced95407556/service/elbv2/api.go#L5879-L5895
		describeLoadBalancerInputs := &elbv2.DescribeLoadBalancersInput{Marker: nextMarker}
		describeLoadBalancerOutputs, err := elbv2client.DescribeLoadBalancers(describeLoadBalancerInputs)

		if err != nil {
			log.Fatalf("Failed to get metadata from load balancer because of %v", err)
		}
		
		for _, lb := range describeLoadBalancerOutputs.LoadBalancers {

			//Skipping lb that does not contain aoc-lb string (relating to aws-otel-test-framework)
			if filterLbNameResult := strings.Contains(*lb.LoadBalancerName, containLbName); !filterLbNameResult {
				continue
			}

			//Skipping lb that does not older than 5 days
			if !time.Now().UTC().Add(pastDayDeleteCalculation).After(*lb.CreatedTime) {
				continue
			}

			log.Printf("Trying to delete lb %s with launch-date %v", *lb.LoadBalancerName, lb.CreatedTime)

			//Delete load balancer
			//Documentation: https://github.com/aws/aws-sdk-go/blob/main/service/elbv2/api.go#L829-L844
			deleteLoadBalancerInput := &elbv2.DeleteLoadBalancerInput{
				LoadBalancerArn: lb.LoadBalancerArn,
			}
			_, err = elbv2client.DeleteLoadBalancer(deleteLoadBalancerInput)

			if err != nil {
				log.Fatalf("Failed to delete lb %s because of %v", *lb.LoadBalancerName, err)
			}
			log.Printf("Delete lb %s successfully", *lb.LoadBalancerName)
		}

		if describeLoadBalancerOutputs.NextMarker == nil {
			break
		}
		nextMarker = describeLoadBalancerOutputs.NextMarker
	}
}
