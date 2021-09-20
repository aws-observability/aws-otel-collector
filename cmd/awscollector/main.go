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

	"github.com/aws-observability/aws-otel-collector/pkg/config"
	"github.com/aws-observability/aws-otel-collector/pkg/defaultcomponents"
	"github.com/aws-observability/aws-otel-collector/pkg/extraconfig"
	"github.com/aws-observability/aws-otel-collector/pkg/logger"
	"github.com/aws-observability/aws-otel-collector/tools/version"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"
	"go.opentelemetry.io/collector/service/parserprovider"
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

	// init cfgFactory
	cfgFactory := config.GetParserProvider()

	// init lumberFunc for zap logger
	lumberHook := logger.GetLumberHook()

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
		Factories:      factories,
		BuildInfo:      info,
		ParserProvider: cfgFactory,
	}
	if lumberHook != nil {
		params.LoggingOptions = []zap.Option{zap.Hooks(lumberHook)}
	}
	if err := run(params); err != nil {
		logFatal(err)
	}

}

func runInteractive(params service.CollectorSettings) error {
	app, err := service.New(params)
	if err != nil {
		return fmt.Errorf("failed to construct the application: %w", err)
	}

	cmd := newCommand(app, params)
	err = cmd.Execute()
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

// newCommand constructs a new cobra.Command using the given Collector and settings.
func newCommand(col *service.Collector, params service.CollectorSettings) *cobra.Command {
	// create a Command with the upstream default FlagSet
	// so that we can parse them to ensure default values are set
	coreCmd := service.NewCommand(col)
	err := coreCmd.Flags().Parse([]string{})
	if err != nil {
		panic(fmt.Sprintf("error setting default flag values: %s", err))
	}

	// build the Command we will use that only has Flags
	// for the ParserProvider
	rootCmd := &cobra.Command{
		Use:          params.BuildInfo.Command,
		Version:      params.BuildInfo.Version,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return col.Run(cmd.Context())
		},
	}

	flagSet := new(flag.FlagSet)
	addFlagsFns := []func(*flag.FlagSet){
		parserprovider.Flags,
	}
	for _, addFlags := range addFlagsFns {
		addFlags(flagSet)
	}

	rootCmd.Flags().AddGoFlagSet(flagSet)
	return rootCmd
}
