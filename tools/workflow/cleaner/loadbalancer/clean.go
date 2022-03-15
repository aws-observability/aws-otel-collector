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
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

const (
	Type          = "loadbalancer"
	containLbName = "aoc-lb"
)

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(sess *session.Session, expirationDate time.Time) error {
	logger.Printf("Begin to clean Load Balancer resources")
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
			return err
		}

		for _, lb := range describeLoadBalancerOutputs.LoadBalancers {
			//Skipping lb that does not contain aoc-lb string (relating to aws-otel-test-framework)
			if filterLbNameResult := strings.Contains(*lb.LoadBalancerName, containLbName); !filterLbNameResult {
				continue
			}

			//Skipping lb that does not older than 5 days
			if !expirationDate.After(*lb.CreatedTime) {
				continue
			}

			logger.Printf("Trying to delete lb %s with launch-date %v", *lb.LoadBalancerName, lb.CreatedTime)

			//Delete load balancer
			//Documentation: https://github.com/aws/aws-sdk-go/blob/main/service/elbv2/api.go#L829-L844
			deleteLoadBalancerInput := &elbv2.DeleteLoadBalancerInput{LoadBalancerArn: lb.LoadBalancerArn}
			if _, err = elbv2client.DeleteLoadBalancer(deleteLoadBalancerInput); err != nil {
				return err
			}
			logger.Printf("Deleted lb %s successfully", *lb.LoadBalancerName)
		}

		if describeLoadBalancerOutputs.NextMarker == nil {
			break
		}
		nextMarker = describeLoadBalancerOutputs.NextMarker
	}
	return nil
}
