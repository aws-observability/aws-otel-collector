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

package main // import "aws-observability.io/collector/cmd/awscollector"

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws-observability/aws-otel-collector/pkg/config"
	"github.com/aws-observability/aws-otel-collector/pkg/defaultcomponents"
	"github.com/aws-observability/aws-otel-collector/pkg/extraconfig"
	"github.com/aws-observability/aws-otel-collector/pkg/logger"
	"github.com/aws-observability/aws-otel-collector/tools/version"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"
	"go.uber.org/zap"
)

// aws-otel-collector is built upon opentelemetry-collector.
// in main() function, aws team has customized logging and configuration handling
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

	var zapHooks []zap.Option

	// init lumberFunc for zap logger
	lumberHook := logger.GetLumberHook()
	if lumberHook != nil {
		zapHooks = append(zapHooks, zap.Hooks(lumberHook))
	}

	// init an ECS Error Logger and a go routine for error reporting when the STATUS_MESSAGE_FILE_PATH is set
	if os.Getenv("STATUS_MESSAGE_FILE_PATH") != "" {
		ecsErrorLogger := logger.NewECSErrorLogger()
		ecsErrorHook := ecsErrorLogger.GetErrorHook()
		zapHooks = append(zapHooks, zap.Hooks(ecsErrorHook))
		ecsErrorLogger.Ticker = time.NewTicker(time.Minute)
		go ecsErrorLogger.Run()
	}

	// set the collector config from extracfg file
	if extraConfig != nil {
		setCollectorConfigFromExtraCfg(extraConfig)
	}

	info := component.BuildInfo{
		Command:     "aws-otel-collector",
		Description: "AWS OTel Collector",
		Version:     version.Version,
	}

	params := service.CollectorSettings{
		Factories: factories,
		BuildInfo: info,
	}
	if len(zapHooks) != 0 {
		params.LoggingOptions = zapHooks
	}
	if err := run(params); err != nil {
		logFatal(err)
	}

}

func runInteractive(params service.CollectorSettings) error {
	cmd := newCommand(params)
	err := cmd.Execute()
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

// newCommand constructs a new cobra.Command using the given settings.
func newCommand(params service.CollectorSettings) *cobra.Command {
	// build the Command we will use that only has config/set flags
	rootCmd := &cobra.Command{
		Use:          params.BuildInfo.Command,
		Version:      params.BuildInfo.Version,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Initialize provider after flags have been set
			params.ConfigMapProvider = config.GetMapProvider()
			col, err := service.New(params)
			if err != nil {
				return fmt.Errorf("failed to construct the application: %w", err)
			}
			return col.Run(cmd.Context())
		},
	}

	rootCmd.Flags().AddGoFlagSet(config.Flags())
	return rootCmd
}
