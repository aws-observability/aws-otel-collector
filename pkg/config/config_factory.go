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
	"flag"
	"log"
	"os"

	"github.com/open-telemetry/opentelemetry-collector-contrib/confmap/provider/s3provider"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/converter/expandconverter"
	"go.opentelemetry.io/collector/confmap/provider/envprovider"
	"go.opentelemetry.io/collector/confmap/provider/fileprovider"
	"go.opentelemetry.io/collector/confmap/provider/httpprovider"
	"go.opentelemetry.io/collector/confmap/provider/httpsprovider"
	"go.opentelemetry.io/collector/confmap/provider/yamlprovider"
	"go.opentelemetry.io/collector/otelcol"
)

const (
	envKey = "AOT_CONFIG_CONTENT"
)

func GetConfigProvider(flags *flag.FlagSet) otelcol.ConfigProvider {
	// aws-otel-collector supports loading yaml config from Env Var
	// including SSM parameter store for ECS use case
	loc := getConfigFlag(flags)
	if _, ok := os.LookupEnv(envKey); ok {
		loc = []string{"env:" + envKey}
	}

	// generate the MapProviders for the Config Provider Settings
	providers := []confmap.Provider{
		fileprovider.NewWithSettings(confmap.ProviderSettings{}),
		envprovider.NewWithSettings(confmap.ProviderSettings{}),
		yamlprovider.NewWithSettings(confmap.ProviderSettings{}),
		httpprovider.NewWithSettings(confmap.ProviderSettings{}),
		httpsprovider.NewWithSettings(confmap.ProviderSettings{}),
		s3provider.New(),
	}

	mapProviders := make(map[string]confmap.Provider, len(providers))
	for _, provider := range providers {
		mapProviders[provider.Scheme()] = provider
	}

	// create Config Provider Settings
	settings := otelcol.ConfigProviderSettings{
		ResolverSettings: confmap.ResolverSettings{
			URIs:       loc,
			Providers:  mapProviders,
			Converters: []confmap.Converter{expandconverter.New(confmap.ConverterSettings{})},
		},
	}

	// get New config Provider
	config_provider, err := otelcol.NewConfigProvider(settings)

	if err != nil {
		log.Panicf("Err on creating Config Provider: %v\n", err)
	}

	return config_provider
}
