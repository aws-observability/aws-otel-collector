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
	"testing"

	"go.opentelemetry.io/collector/service"

	"github.com/aws-observability/aws-otel-collector/pkg/defaultcomponents"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestNewCommandFlagSet(t *testing.T) {
	factories, _ := defaultcomponents.Components()
	params := service.CollectorSettings{
		Factories: factories,
	}

	validFlags := []string{"config", "set"}
	flagSet := newCommand(params).Flags()
	flagSet.VisitAll(func(f *pflag.Flag) {
		assert.Contains(t, validFlags, f.Name)
	})
}
