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
			if err = autoscaling.Clean(sess, keepDuration); err != nil {
				log.Printf("%v", err)
			}
		case ec2.Type:
			if err = ec2.Clean(sess, keepDuration); err != nil {
				log.Printf("%v", err)
			}
		case efs.Type:
			if err = efs.Clean(sess, keepDuration); err != nil {
				log.Printf("%v", err)
			}
		case iam.Type:
			if err = iam.Clean(sess, keepDuration); err != nil {
				log.Printf("%v", err)
			}
		case launchconfig.Type:
			if err = launchconfig.Clean(sess, keepDuration); err != nil {
				log.Printf("%v", err)
			}
		case loadbalancer.Type:
			if err = loadbalancer.Clean(sess, keepDuration); err != nil {
				log.Printf("%v", err)
			}
		default:
			log.Printf("Skipping invalid cleaner '%s'. Please see -h for options.", cleaner)
		}
	}

	log.Printf("Finished cleaning AWS resources")
}
