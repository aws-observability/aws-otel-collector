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
	"os"

	"github.com/spf13/viper"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/service"
)

// GetCfgFactory returns AOC/Otel config
func GetCfgFactory() func(otelViper *viper.Viper, f component.Factories) (*configmodels.Config, error) {
	return func(otelViper *viper.Viper, f component.Factories) (*configmodels.Config, error) {
		// AOC supports loading yaml config from SSM parameter store
		if ssmConfigContent, ok := os.LookupEnv("AOT_CONFIG_CONTENT"); ok &&
			os.Getenv("RUN_IN_CONTAINER") == "True" {
			fmt.Printf("Reading json config from from environment: %v\n", ssmConfigContent)
			return readConfigString(otelViper, f, ssmConfigContent)
		}

		// use OTel yaml config from input
		otelCfg, err := service.FileLoaderConfigFactory(otelViper, f)
		if err != nil {
			return nil, err
		}
		return otelCfg, nil
	}
}

// readConfigString set AOC/Otel config from SSM parameter store
func readConfigString(v *viper.Viper,
	factories component.Factories,
	configContent string) (*configmodels.Config, error) {
	v.SetConfigType("yaml")
	var configBytes = []byte(configContent)
	err := v.ReadConfig(bytes.NewBuffer(configBytes))
	if err != nil {
		return nil, fmt.Errorf("error loading SSM config file %v", err)
	}
	return config.Load(v, factories)
}
