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

package main // import "aws-observability.io/collector/cmd/awscollector"

import (
	"fmt"
	"log"
	"os"

	"github.com/aws-observability/aws-otel-collector/pkg/config"
	"github.com/aws-observability/aws-otel-collector/pkg/defaultcomponents"
	"github.com/aws-observability/aws-otel-collector/pkg/extraconfig"
	"github.com/aws-observability/aws-otel-collector/pkg/logger"
	"github.com/aws-observability/aws-otel-collector/tools/version"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"
	"go.uber.org/zap"
)

// aws-otel-collector is built upon opentelemetry-collector.
// in main() funtion, aws team has customized logging and configuration handling
// logic and it only supports the selected components which have been verified by AWS
// from opentelemetry-collector list
func main() {
	// get extra config
	extraConfig := getExtraConfig()

	logger.SetupErrorLogger()

	factories, err := defaultcomponents.Components()
	if err != nil {
		log.Fatalf("failed to build components: %v", err)
	}

	// init cfgFactory
	cfgFactory := config.GetCfgFactory()

	// init lumberFunc for zap logger
	lumberHook := logger.GetLumberHook()

	// set the collector config from extracfg file
	if extraConfig != nil {
		setCollectorConfigFromExtraCfg(extraConfig)
	}

	info := component.ApplicationStartInfo{
		ExeName:  "aws-otel-collector",
		LongName: "AWS OTel Collector",
		Version:  version.Version,
		GitHash:  version.GitHash,
	}

	if err := run(service.Parameters{
		Factories:            factories,
		ApplicationStartInfo: info,
		ConfigFactory:        cfgFactory,
		LoggingOptions:       []zap.Option{zap.Hooks(lumberHook)}}); err != nil {
		log.Fatal(err)
	}

}

func runInteractive(params service.Parameters) error {
	app, err := service.New(params)
	if err != nil {
		return fmt.Errorf("failed to construct the application: %w", err)
	}

	err = app.Run()
	if err != nil {
		return fmt.Errorf("application run finished with error: %w", err)
	}

	return nil
}

func getExtraConfig() *extraconfig.ExtraConfig {
	extraConfig, err := extraconfig.GetExtraConfig()
	if err != nil {
		log.Printf("find no extra config, skip it, err: %v", err)
		return nil
	}
	return extraConfig
}

func setCollectorConfigFromExtraCfg(extraCfg *extraconfig.ExtraConfig) {
	if extraCfg.LoggingLevel != "" {
		logger.SetLogLevel(extraCfg.LoggingLevel)
	}
	if extraCfg.AwsProfile != "" {
		os.Setenv("AWS_PROFILE", extraCfg.AwsProfile)
	}
	if extraCfg.AwsCredentialFile != "" {
		os.Setenv("AWS_SHARED_CREDENTIALS_FILE", extraCfg.AwsCredentialFile)
	}
}
