package main

import (
	"time"
	"log"
	"strings"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"

)

const (
	containLbString = "aoc-lb"
	pastDayDelete = 5
)

func main() {
	// Set up aws go sdk session
	//Enable default region and credentials
	//Documents: https://docs.aws.amazon.com/ja_jp/sdk-for-go/v1/developer-guide/configuring-sdk.html
	session := session.Must(
					session.NewSessionWithOptions(
						session.Options{
							SharedConfigState: session.SharedConfigEnable,
	}))

	svc := elbv2.New(session)

	//ELB Go SDK currently does not support filter tag or filter wildcard. Only supports with matching name
	//Documentation: https://github.com/aws/aws-sdk-go/blob/02266ed24221ac21bb37d6ac614d1ced95407556/service/elbv2/api.go#L5879-L5895

	describeLoadBalancerInputs := &elbv2.DescribeLoadBalancersInput{}


	describeLoadBalancerOutputs, describeErr := svc.DescribeLoadBalancers(describeLoadBalancerInputs)

	if describeErr != nil {
		log.Fatalf("Failed to get metadata for load balancer because of %v", describeErr)
	}

	for _, lb := range describeLoadBalancerOutputs.LoadBalancers {

		//Skipping lb that does not contain aoc-lb string (relating to aws-otel-test-framework)
		if filterLbNameResult := strings.Contains(*lb.LoadBalancerName, containLbString); !filterLbNameResult {
			log.Printf("Skipping lb %s because of not containing %s", *lb.LoadBalancerName, containLbString)
			continue
		}

		//Skipping lb that does not older than 5 days
		if !time.Now().UTC().Add(-1 * time.Hour * 24 * pastDayDelete).After(*lb.CreatedTime) {
			log.Printf("Skipping lb %s with launch-date %v", *lb.LoadBalancerName, lb.CreatedTime)
			continue
		}

		log.Printf("Trying to delete lb %s with launch-date %v", *lb.LoadBalancerName, lb.CreatedTime)

		//Delete load balancer
		//Documentation: https://github.com/aws/aws-sdk-go/blob/main/service/elbv2/api.go#L829-L844
		deleteLoadBalancerInput := &elbv2.DeleteLoadBalancerInput{
			LoadBalancerArn: lb.LoadBalancerArn,
		}
		_, deleteLoadBalancerError := svc.DeleteLoadBalancer(deleteLoadBalancerInput)

		if deleteLoadBalancerError != nil {
			log.Fatalf("Failed to delete lb %s because of %v", *lb.LoadBalancerName,deleteLoadBalancerError)
		}

		log.Printf("Delete lb %s successfully", *lb.LoadBalancerName)

	}

}




