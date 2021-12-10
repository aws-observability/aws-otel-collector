package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"time"
)

func main() {
	// set up aws go sdk ec2 client
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
				if time.Now().UTC().Add(-1 * time.Hour * 24 * 5).After(*instance.LaunchTime) {
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
		log.Print("No instances to delete")
		return
	}

	terminateInstancesInput := ec2.TerminateInstancesInput{InstanceIds: deleteInstanceIds}
	_, err = ec2client.TerminateInstances(&terminateInstancesInput)
	if err != nil {
		log.Fatalf("Failed to terminate instance %v", err)
	}
}
