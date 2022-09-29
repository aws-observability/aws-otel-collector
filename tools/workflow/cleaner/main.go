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

	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/apigw"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/aps"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/autoscaling"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/ebs"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/ec2"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/ecs"
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

	cleanerTypes   = []string{aps.Type, autoscaling.Type, ec2.Type, ecs.Type, efs.Type, iam.Type, launchconfig.Type, loadbalancer.Type, ebs.Type, apigw.Type}
	cleanerOptions = strings.Join(cleanerTypes, delimiter)
)

func init() {
	flag.IntVar(&daysToKeep, "keep", defaultDaysToKeep, "Days to keep a resource before cleaning it.")
	flag.StringVar(&cleanersToRun, "clean", cleanerTypeAll,
		fmt.Sprintf("Determines which cleaners to run. (e.g. -clean ec2%siam)\nOptions: %s,%s", delimiter, cleanerTypeAll, cleanerOptions))
}

func main() {
	flag.Parse()

	sess := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))

	keepDuration := -1 * time.Hour * 24 * time.Duration(daysToKeep)

	if strings.Contains(cleanersToRun, cleanerTypeAll) {
		cleanersToRun = cleanerOptions
	}

	// The date used to determine if a resource can be cleaned.
	// Anything older than the date will be removed.
	expirationDate := time.Now().UTC().Add(keepDuration)
	log.Printf("Expiration date set to %v", expirationDate)

	for _, cleaner := range strings.Split(cleanersToRun, delimiter) {
		switch cleaner {
		case aps.Type:
			if err := aps.Clean(sess, expirationDate); err != nil {
				log.Printf("%v", err)
			}
		case autoscaling.Type:
			if err := autoscaling.Clean(sess, expirationDate); err != nil {
				log.Printf("%v", err)
			}
		case ec2.Type:
			if err := ec2.Clean(sess, expirationDate); err != nil {
				log.Printf("%v", err)
			}
		case ecs.Type:
			if err := ecs.Clean(sess, expirationDate); err != nil {
				log.Printf("%v", err)
			}
		case efs.Type:
			if err := efs.Clean(sess, expirationDate); err != nil {
				log.Printf("%v", err)
			}
		case iam.Type:
			if err := iam.Clean(sess, expirationDate); err != nil {
				log.Printf("%v", err)
			}
		case launchconfig.Type:
			if err := launchconfig.Clean(sess, expirationDate); err != nil {
				log.Printf("%v", err)
			}
		case loadbalancer.Type:
			if err := loadbalancer.Clean(sess, expirationDate); err != nil {
				log.Printf("%v", err)
			}
		case ebs.Type:
			if err := ebs.Clean(sess, expirationDate); err != nil {
				log.Printf("%v", err)
			}
		case apigw.Type:
			if err := apigw.Clean(sess, expirationDate); err != nil {
				log.Printf("%v", err)
			}
		default:
			log.Printf("Skipping invalid cleaner '%s'. Please see -h for options.", cleaner)
		}
	}

	log.Printf("Finished cleaning AWS resources")
}
