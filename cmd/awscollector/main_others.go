//go:build !windows
// +build !windows

/*
 * Copyright The OpenTelemetry Authors
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
	"log"

	"go.opentelemetry.io/collector/otelcol"

	"github.com/aws-observability/aws-otel-collector/pkg/extraconfig"
	"github.com/aws-observability/aws-otel-collector/pkg/userutils"
)

func run(params otelcol.CollectorSettings, flagSet *flag.FlagSet) error {
	// Try to switch user when the collector is running on a host.
	// For container the user and group is determined by the deployed manifest.
	if !extraconfig.IsRunningInContainer() {
		// avoid running as 'root' user on Linux
		_, err := userutils.ChangeUser()
		if err != nil {
			log.Printf("E! Failed to ChangeUser: %v ", err)
			return err
		}
	}
	return runInteractive(params, flagSet)
}

func logFatal(err error) {
	log.Fatal(err)
}
