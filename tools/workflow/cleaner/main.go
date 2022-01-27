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
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/autoscaling"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/ec2"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/efs"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/iam"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/launchconfig"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/loadbalancer"
)

const (
	defaultDaysToKeep = 5
	cleanerTypeAll    = "all"
	delimiter         = ","
)

var (
	// flags
	daysToKeep    int
	cleanersToRun string

	cleanerTypes   = []string{autoscaling.Type, ec2.Type, efs.Type, iam.Type, launchconfig.Type, loadbalancer.Type}
	cleanerOptions = strings.Join(cleanerTypes, delimiter)
)

func init() {
	flag.IntVar(&daysToKeep, "keep", defaultDaysToKeep, "Days to keep a resource before cleaning it.")
	flag.StringVar(&cleanersToRun, "clean", cleanerTypeAll,
		fmt.Sprintf("Determines which cleaners to run. (e.g. -clean ec2%siam)\nOptions: %s,%s", delimiter, cleanerTypeAll, cleanerOptions))
}

func main() {
	flag.Parse()

	sess, err := session.NewSession()

	if err != nil {
		log.Fatalf("Error creating session %v", err)
	}

	keepDuration := -1 * time.Hour * 24 * time.Duration(daysToKeep)

	if strings.Contains(cleanersToRun, cleanerTypeAll) {
		cleanersToRun = cleanerOptions
	}

	for _, cleaner := range strings.Split(cleanersToRun, delimiter) {
		switch cleaner {
		case autoscaling.Type:
			log.Printf("Begin to clean ECS AutoScaling")
			autoscaling.Clean(sess, keepDuration)
		case ec2.Type:
			log.Printf("Begin to clean EC2 Instances")
			ec2.Clean(sess, keepDuration)
		case efs.Type:
			log.Printf("Begin to clean EFS resources")
			efs.Clean(sess, keepDuration)
		case iam.Type:
			log.Printf("Begin to clean IAM roles")
			iam.Clean(sess, keepDuration)
		case launchconfig.Type:
			log.Printf("Begin to clean ECS Launch Configuration")
			launchconfig.Clean(sess, keepDuration)
		case loadbalancer.Type:
			log.Printf("Begin to clean Load Balancer resources")
			loadbalancer.Clean(sess, keepDuration)
		default:
			log.Printf("Skipping invalid cleaner '%s'. Please see -h for options.", cleaner)
		}
	}

	log.Printf("Finished cleaning AWS resources")
}
