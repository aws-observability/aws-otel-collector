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

	"go.opentelemetry.io/collector/service"

	"go.opentelemetry.io/collector/config"

	"go.opentelemetry.io/collector/config/mapconverter/expandmapconverter"
	"go.opentelemetry.io/collector/config/mapprovider/envmapprovider"
	"go.opentelemetry.io/collector/config/mapprovider/filemapprovider"
	"go.opentelemetry.io/collector/config/mapprovider/yamlmapprovider"
)

const (
	envKey = "AOT_CONFIG_CONTENT"
)

func GetConfigProvider() service.ConfigProvider {
	// aws-otel-collector supports loading yaml config from Env Var
	// including SSM parameter store for ECS use case
	loc := getConfigFlag()
	if configContent, ok := os.LookupEnv(envKey); ok {
		log.Printf("Reading AOT config from environment: %v\n", configContent)
		loc = []string{"env:" + envKey}
	}

	providers := []config.MapProvider{filemapprovider.New(), envmapprovider.New(), yamlmapprovider.New()}

	ret := make(map[string]config.MapProvider, len(providers))
	for _, provider := range providers {
		ret[provider.Scheme()] = provider
	}

	settings := service.ConfigProviderSettings{
		Locations:     loc,
		MapProviders:  ret,
		MapConverters: []config.MapConverterFunc{expandmapconverter.New()},
	}

	config_provider, err := service.NewConfigProvider(settings)

	getSetFlag()

	log.Printf("Err on creating Config Provider: %v\n", err)

	return config_provider

}
