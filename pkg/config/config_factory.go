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
	"fmt"
	"log"
	"os"

	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configmapprovider"
	"go.opentelemetry.io/collector/config/configunmarshaler"
	"go.opentelemetry.io/collector/service"
)

const (
	envKey = "AOT_CONFIG_CONTENT"
)

func GetConfigProvider() service.ConfigProvider {
	// aws-otel-collector supports loading yaml config from Env Var
	// including SSM parameter store for ECS use case
	if configContent, ok := os.LookupEnv(envKey); ok {
		log.Printf("Reading AOT config from environment: %v\n", configContent)
		return service.NewConfigProvider(
			[]string{fmt.Sprintf("env:%s", envKey)},
			map[string]configmapprovider.Provider{"env": configmapprovider.NewEnv()},
			[]config.MapConverterFunc{configmapprovider.NewExpandConverter()},
			configunmarshaler.NewDefault())
	}

	return service.NewDefaultConfigProvider([]string{getConfigFlag()}, getSetFlag())
}
