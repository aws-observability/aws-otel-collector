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
	"os"

	"github.com/open-telemetry/opentelemetry-collector-contrib/confmap/provider/s3provider"
	"go.opentelemetry.io/collector/confmap"
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

func GetConfigProviderSettings(flags *flag.FlagSet) otelcol.ConfigProviderSettings {
	// aws-otel-collector supports loading yaml config from Env Var
	// including SSM parameter store for ECS use case
	loc := getConfigFlag(flags)
	if _, ok := os.LookupEnv(envKey); ok {
		loc = []string{"env:" + envKey}
	}

	// generate the MapProviders for the Config Provider Settings
	providers := []confmap.ProviderFactory{
		fileprovider.NewFactory(),
		envprovider.NewFactory(),
		yamlprovider.NewFactory(),
		httpprovider.NewFactory(),
		httpsprovider.NewFactory(),
		s3provider.NewFactory(),
	}

	// create Config Provider Settings
	configProviderSettings := otelcol.ConfigProviderSettings{
		ResolverSettings: confmap.ResolverSettings{
			URIs:              loc,
			ProviderFactories: providers,
			DefaultScheme:     "env",
		},
	}

	return configProviderSettings
}
