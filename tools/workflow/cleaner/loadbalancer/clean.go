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

package loadbalancer

import (
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

const (
	Type          = "loadbalancer"
	containLbName = "aoc-lb"
)

func Clean(sess *session.Session, keepDuration time.Duration) {
	elbv2client := elbv2.New(sess)

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
			if !time.Now().UTC().Add(keepDuration).After(*lb.CreatedTime) {
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
