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
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

const (
	Type          = "loadbalancer"
	containLbName = "aoc-lb"
)

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	logger.Printf("Begin to clean Load Balancer resources")
	elbv2client := elasticloadbalancingv2.NewFromConfig(cfg)

	//Allow to load all the load balancers since the default respond is paginated load balancers.
	//Look into the documentations and read the starting-token for more details
	//Documentation: https://docs.aws.amazon.com/cli/latest/reference/elbv2/describe-load-balancers.html#options
	paginator := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(elbv2client, &elasticloadbalancingv2.DescribeLoadBalancersInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		for _, lb := range output.LoadBalancers {
			if !shouldDeleteLoadBalancer(lb, expirationDate) {
				continue
			}

			logger.Printf("Trying to delete lb %s with launch-date %v", *lb.LoadBalancerName, lb.CreatedTime)

			//Delete load balancer
			//Documentation: https://github.com/aws/aws-sdk-go/blob/main/service/elbv2/api.go#L829-L844
			_, err = elbv2client.DeleteLoadBalancer(ctx, &elasticloadbalancingv2.DeleteLoadBalancerInput{LoadBalancerArn: lb.LoadBalancerArn})
			if err != nil {
				return err
			}
			logger.Printf("Deleted lb %s successfully", *lb.LoadBalancerName)
		}
	}
	return nil
}

// shouldDeleteLoadBalancer returns true if the load balancer name contains "aoc-lb" and is older than the expiration date.
func shouldDeleteLoadBalancer(lb types.LoadBalancer, expirationDate time.Time) bool {
	if lb.LoadBalancerName == nil || lb.CreatedTime == nil {
		return false
	}
	return strings.Contains(*lb.LoadBalancerName, containLbName) && expirationDate.After(*lb.CreatedTime)
}
