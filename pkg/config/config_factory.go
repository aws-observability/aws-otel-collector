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

	"go.opentelemetry.io/collector/config/configmapprovider"
)

const (
	envKey = "AOT_CONFIG_CONTENT"
)

func GetMapProvider() configmapprovider.Provider {
	// aws-otel-collector supports loading yaml config from Env Var
	// including SSM parameter store for ECS use case
	if configContent, ok := os.LookupEnv(envKey); ok {
		log.Printf("Reading AOT config from environment: %v\n", configContent)
		return configmapprovider.NewExpand(configmapprovider.NewInMemory(strings.NewReader(configContent)))
	}

	return configmapprovider.NewDefault(getConfigFlag(), getSetFlag())
}
