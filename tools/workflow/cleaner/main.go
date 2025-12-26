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
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/apigw"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/aps"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/autoscaling"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/ebs"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/ec2"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/ecs"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/efs"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/iam"
	"github.com/aws-observability/aws-otel-collector/tools/workflow/cleaner/lambda"
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

	cleanerTypes   = []string{aps.Type, autoscaling.Type, ec2.Type, ecs.Type, efs.Type, iam.Type, launchconfig.Type, loadbalancer.Type, ebs.Type, apigw.Type, lambda.Type}
	cleanerOptions = strings.Join(cleanerTypes, delimiter)
)

func init() {
	flag.IntVar(&daysToKeep, "keep", defaultDaysToKeep, "Days to keep a resource before cleaning it.")
	flag.StringVar(&cleanersToRun, "clean", cleanerTypeAll,
		fmt.Sprintf("Determines which cleaners to run. (e.g. -clean ec2%siam)\nOptions: %s,%s", delimiter, cleanerTypeAll, cleanerOptions))
}

func main() {
	flag.Parse()

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config: %v", err)
	}

	keepDuration := -1 * time.Hour * 24 * time.Duration(daysToKeep)

	if strings.Contains(cleanersToRun, cleanerTypeAll) {
		cleanersToRun = cleanerOptions
	}

	// The date used to determine if a resource can be cleaned.
	// Anything older than the date will be removed.
	expirationDate := time.Now().UTC().Add(keepDuration)
	log.Printf("Expiration date set to %v", expirationDate)

	for _, cleaner := range strings.Split(cleanersToRun, delimiter) {
		if err := runCleaner(ctx, cfg, cleaner, expirationDate); err != nil {
			log.Fatalf("%v", err)
		}
	}

	log.Printf("Finished cleaning AWS resources")
}

func runCleaner(ctx context.Context, cfg aws.Config, cleaner string, expirationDate time.Time) error {
	switch cleaner {
	case aps.Type:
		return aps.Clean(ctx, cfg, expirationDate)
	case autoscaling.Type:
		return autoscaling.Clean(ctx, cfg, expirationDate)
	case ec2.Type:
		return ec2.Clean(ctx, cfg, expirationDate)
	case ecs.Type:
		return ecs.Clean(ctx, cfg, expirationDate)
	case efs.Type:
		return efs.Clean(ctx, cfg, expirationDate)
	case iam.Type:
		return iam.Clean(ctx, cfg, expirationDate)
	case lambda.Type:
		return lambda.Clean(ctx, cfg, expirationDate)
	case launchconfig.Type:
		return launchconfig.Clean(ctx, cfg, expirationDate)
	case loadbalancer.Type:
		return loadbalancer.Clean(ctx, cfg, expirationDate)
	case ebs.Type:
		return ebs.Clean(ctx, cfg, expirationDate)
	case apigw.Type:
		return apigw.Clean(ctx, cfg, expirationDate)
	default:
		return fmt.Errorf("invalid cleaner '%s' - see -h for options", cleaner)
	}
}
