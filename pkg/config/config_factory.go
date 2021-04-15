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

package config

import (
	"log"
	"os"
	"strings"

	"go.opentelemetry.io/collector/service/parserprovider"
)

const (
	envKey = "AOT_CONFIG_CONTENT"
)

func GetParserProvider() parserprovider.ParserProvider {
	// aws-otel-collector supports loading yaml config from Env Var
	// including SSM parameter store for ECS use case
	if configContent, ok := os.LookupEnv(envKey); ok {
		log.Printf("Reading AOT config from from environment: %v\n", configContent)
		return parserprovider.NewInMemory(strings.NewReader(configContent))
	}
	return parserprovider.Default()
}

// TryFileConfig forces loading file based config.
// TODO(pingleig): It's a hack for https://github.com/aws-observability/aws-otel-collector/issues/340.
func TryFileConfig() error {
	// We only use env when running in container, in that case we should go with normal error handling.
	if _, ok := os.LookupEnv(envKey); ok {
		return nil
	}

	// Default loads config file specified in cli arg and we use that on EC2
	// for composite agent (i.e. cloudwatch agent + otel collector).
	// Because we don't want to restart on windows so we exit 0 for invalid config.
	// https://github.com/aws-observability/aws-otel-collector/pull/286#issuecomment-769512414
	provider := parserprovider.Default()
	_, err := provider.Get()
	return err
}
