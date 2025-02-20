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
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/otelcol"

	"go.opentelemetry.io/collector/component"
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

	// TODO : Remove after exporters are removed
	log.Printf("attn: users of the `datadog`, `logzio`, `sapm`, `signalfx` exporter components. please refer to " +
		"https://github.com/aws-observability/aws-otel-collector/issues/2734 in regards to an upcoming ADOT Collector " +
		"breaking change")

	logger.SetupErrorLogger()

	// set the collector config from extracfg file
	if extraConfig != nil {
		setCollectorConfigFromExtraCfg(extraConfig)
	}

	info := component.BuildInfo{
		Command:     "aws-otel-collector",
		Description: "AWS OTel Collector",
		Version:     version.Version,
	}

	flagSet, err := buildAndParseFlagSet(featuregate.GlobalRegistry())
	if err != nil {
		logFatal(err)
	}

	params := otelcol.CollectorSettings{
		Factories:              defaultcomponents.Components,
		BuildInfo:              info,
		LoggingOptions:         []zap.Option{logger.WrapCoreOpt()},
		ConfigProviderSettings: config.GetConfigProviderSettings(flagSet),
	}

	if err = run(params, flagSet); err != nil {
		logFatal(err)
	}
}

// We parse the flags manually here so that we can use feature gates when constructing
// our default component list. Flags also need to be parsed before creating the config provider.
func buildAndParseFlagSet(featgate *featuregate.Registry) (*flag.FlagSet, error) {
	flagSet := config.Flags(featgate)

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return nil, err
	}
	return flagSet, nil
}

func runInteractive(params otelcol.CollectorSettings, flagSet *flag.FlagSet) error {
	cmd := newCommand(params, flagSet)
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
		if err := os.Setenv(awsProfileKey, extraCfg.AwsProfile); err != nil {
			log.Printf("failed to set env var %s:%v\n", awsProfileKey, err)
		}
	}

	if extraCfg.AwsCredentialFile != "" {
		if err := os.Setenv(awsCredentialFileKey, extraCfg.AwsCredentialFile); err != nil {
			log.Printf("failed to set env var %s:%v\n", awsCredentialFileKey, err)
		}
	}
}

// newCommand constructs a new cobra.Command using the given settings.
func newCommand(params otelcol.CollectorSettings, flagSet *flag.FlagSet) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          params.BuildInfo.Command,
		Version:      params.BuildInfo.Version,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
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
