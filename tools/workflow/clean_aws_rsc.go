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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"log"
	"strings"
	"time"
)
const (
	containLbString = "aoc-lb"
	pastDayDelete = 5
	pastDayDeleteCalculation = -1 * time.Hour * 24 * pastDayDelete //Currently, deleting resources over 5 days
)

func main() {

	log.Printf("Beging terminating EC2 Instances")
	terminateEc2InstancesError := terminateEc2Instances()
	if (terminateEc2InstancesError != nil){
		log.Fatalf("Exit terminating EC2 Instances")
	}

	log.Printf("Begin destroy Load Balancer resources")
	destroyLoadBalancerError := destroyLoadBalancerResource()

	if (destroyLoadBalancerError != nil){
		log.Fatalf("Exit destroy Load Balancer Resources")
	}

	log.Printf("Finish destroy AWS resources")
}

func terminateEc2Instances() error {
	// set up aws go sdk ec2 client
	testSession := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
			}))

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
			log.Printf("Failed to get instance for error %v", err)
			return err
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
		return nil
	}

	terminateInstancesInput := ec2.TerminateInstancesInput{InstanceIds: deleteInstanceIds}
	_, err := ec2client.TerminateInstances(&terminateInstancesInput)
	if err != nil {
		log.Printf("Failed to terminate instances %v because of %v",deleteInstanceIds, err)
		return err
	}

	return nil
}

func destroyLoadBalancerResource() error {
	// Set up aws go sdk session
	//Enable default region and credentials
	//Documents: https://docs.aws.amazon.com/ja_jp/sdk-for-go/v1/developer-guide/configuring-sdk.html
	session := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
			}))

	svc := elbv2.New(session)

	//Allow to load all the load balancers since the default respond is paginated load balancers.
	//Look into the documentations and read the starting-token for more details
	//Documentation: https://docs.aws.amazon.com/cli/latest/reference/elbv2/describe-load-balancers.html#options
	var nextMarker *string

	for {
		//ELB Go SDK currently does not support filter tag or filter wildcard. Only supports with matching name
		//Documentation: https://github.com/aws/aws-sdk-go/blob/02266ed24221ac21bb37d6ac614d1ced95407556/service/elbv2/api.go#L5879-L5895
		describeLoadBalancerInputs := &elbv2.DescribeLoadBalancersInput{Marker: nextMarker}
		describeLoadBalancerOutputs, err := svc.DescribeLoadBalancers(describeLoadBalancerInputs)

		if err != nil {
			log.Printf("Failed to get metadata for load balancer because of %v", err)
			return err
		}

		for _, lb := range describeLoadBalancerOutputs.LoadBalancers {

			//Skipping lb that does not contain aoc-lb string (relating to aws-otel-test-framework)
			if filterLbNameResult := strings.Contains(*lb.LoadBalancerName, containLbString); !filterLbNameResult {
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
			_, err = svc.DeleteLoadBalancer(deleteLoadBalancerInput)

			if err != nil {
				log.Printf("Failed to delete lb %s because of %v", *lb.LoadBalancerName,err)
				return err
			}
			log.Printf("Delete lb %s successfully", *lb.LoadBalancerName)
		}

		if describeLoadBalancerOutputs.NextMarker == nil {
			break
		}
		nextMarker = describeLoadBalancerOutputs.NextMarker
	}
	return nil
}
