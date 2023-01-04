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

	"go.opentelemetry.io/collector/otelcol"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/featuregate"
	"go.uber.org/zap"

	"github.com/aws-observability/aws-otel-collector/pkg/config"
	"github.com/aws-observability/aws-otel-collector/pkg/defaultcomponents"
	"github.com/aws-observability/aws-otel-collector/pkg/extraconfig"
	"github.com/aws-observability/aws-otel-collector/pkg/logger"
	"github.com/aws-observability/aws-otel-collector/tools/version"
)

const (
	awsProfileKey        = "AWS_PROFILE"
	awsCredentialFileKey = "AWS_SHARED_CREDENTIALS_FILE" //nolint:gosec // this is a false positive for G101: Potential hardcoded credentials
)

// aws-otel-collector is built upon opentelemetry-collector.
// in main() function, aws team has customized logging and configuration handling
// logic and it only supports the selected components which have been verified by AWS
// from opentelemetry-collector list
func main() {
	// get extra config
	extraConfig, err := extraconfig.GetExtraConfig()
	if err != nil {
		log.Printf("found no extra config, skip it, err: %v", err)
	}

	logger.SetupErrorLogger()

	factories, err := defaultcomponents.Components()
	if err != nil {
		log.Fatalf("failed to build components: %v", err)
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

	params := otelcol.CollectorSettings{
		Factories:      factories,
		BuildInfo:      info,
		LoggingOptions: []zap.Option{logger.WrapCoreOpt()},
	}

	if err = run(params); err != nil {
		logFatal(err)
	}
}

func runInteractive(params otelcol.CollectorSettings) error {
	cmd := newCommand(params)
	err := cmd.Execute()
	if err != nil {
		return fmt.Errorf("application run finished with error: %w", err)
	}

	return nil
}

func setCollectorConfigFromExtraCfg(extraCfg *extraconfig.ExtraConfig) {
	if extraCfg.LoggingLevel != "" {
		logger.SetLogLevel(extraCfg.LoggingLevel)
	}
	if extraCfg.AwsProfile != "" {
		os.Setenv(awsProfileKey, extraCfg.AwsProfile)
	}
	if extraCfg.AwsCredentialFile != "" {
		os.Setenv(awsCredentialFileKey, extraCfg.AwsCredentialFile)
	}
}

// newCommand constructs a new cobra.Command using the given settings.
func newCommand(params otelcol.CollectorSettings) *cobra.Command {
	flagSet := config.Flags()
	// build the Command we will use that only has config/set flags
	rootCmd := &cobra.Command{
		Use:          params.BuildInfo.Command,
		Version:      params.BuildInfo.Version,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := featuregate.GetRegistry().Apply(config.GetFeatureGatesFlag(flagSet)); err != nil {
				return err
			}
			// Initialize provider after flags have been set
			params.ConfigProvider = config.GetConfigProvider(flagSet)
			col, err := otelcol.NewCollector(params)
			if err != nil {
				return fmt.Errorf("failed to construct the application: %w", err)
			}
			return col.Run(cmd.Context())
		},
	}

	rootCmd.Flags().AddGoFlagSet(flagSet)
	return rootCmd
}
