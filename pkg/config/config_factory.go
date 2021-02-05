/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/service"
)

// GetCfgFactory returns aws-otel-collector config
func GetCfgFactory() func(otelViper *viper.Viper, cmd *cobra.Command, f component.Factories) (*configmodels.Config, error) {
	return func(otelViper *viper.Viper, cmd *cobra.Command, f component.Factories) (*configmodels.Config, error) {
		// aws-otel-collector supports loading yaml config from Env Var
		// including SSM parameter store for ECS use case
		if configContent, ok := os.LookupEnv("AOT_CONFIG_CONTENT"); ok {
			log.Printf("Reading AOT config from from environment: %v\n", configContent)
			return readConfigString(otelViper, f, configContent)
		}

		// use OTel yaml config from input
		otelCfg, err := service.FileLoaderConfigFactory(otelViper, cmd, f)
		if err != nil {
			log.Printf("Config file is missing or invalid, %s", err)
			os.Exit(0)
		}
		return otelCfg, nil
	}
}

// readConfigString set aws-otel-collector config from env var
func readConfigString(v *viper.Viper,
	factories component.Factories,
	configContent string) (*configmodels.Config, error) {
	v.SetConfigType("yaml")
	var configBytes = []byte(configContent)
	err := v.ReadConfig(bytes.NewBuffer(configBytes))
	if err != nil {
		return nil, fmt.Errorf("error loading config %v", err)
	}
	return config.Load(v, factories)
}
