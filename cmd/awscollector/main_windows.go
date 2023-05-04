//go:build windows
// +build windows

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
	"os"

	"github.com/pkg/errors"
	"golang.org/x/sys/windows/svc"

	"go.opentelemetry.io/collector/otelcol"
)

func run(params otelcol.CollectorSettings, flagSet *flag.FlagSet) error {
	isService, err := svc.IsWindowsService()
	if err != nil {
		return errors.Wrap(err, "failed to determine if we are running in an interactive session")
	}

	if isService {
		return runService(params)
	} else {
		return runInteractive(params, flagSet)
	}
}

func runService(params otelcol.CollectorSettings) error {
	if err := svc.Run("", otelcol.NewSvcHandler(params)); err != nil {
		return errors.Wrap(err, "failed to start service")
	}

	return nil
}

// Has to set exit code 0 for the fatal errors when run the collector as service in Windows
// to prevent infinite reboot by Windows service
// https://github.com/aws-observability/aws-otel-collector/issues/340
func logFatal(err error) {
	log.Println(err.Error())
	os.Exit(0)
}
