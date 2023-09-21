// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelcol // import "go.opentelemetry.io/collector/otelcol"

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"go.opentelemetry.io/collector/component"
)

<<<<<<< HEAD
type componentsOutput struct {
	BuildInfo  component.BuildInfo
	Receivers  []component.Type
	Processors []component.Type
	Exporters  []component.Type
	Connectors []component.Type
	Extensions []component.Type
=======
type componentWithStability struct {
	Name      component.Type
	Stability map[string]string
}

type componentsOutput struct {
	BuildInfo  component.BuildInfo
	Receivers  []componentWithStability
	Processors []componentWithStability
	Exporters  []componentWithStability
	Connectors []componentWithStability
	Extensions []componentWithStability
>>>>>>> main
}

// newComponentsCommand constructs a new components command using the given CollectorSettings.
func newComponentsCommand(set CollectorSettings) *cobra.Command {
	return &cobra.Command{
		Use:   "components",
		Short: "Outputs available components in this collector distribution",
<<<<<<< HEAD
=======
		Long:  "Outputs available components in this collector distribution including their stability levels. The output format is not stable and can change between releases.",
>>>>>>> main
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			components := componentsOutput{}
			for con := range set.Factories.Connectors {
<<<<<<< HEAD
				components.Connectors = append(components.Connectors, con)
			}
			for ext := range set.Factories.Extensions {
				components.Extensions = append(components.Extensions, ext)
			}
			for prs := range set.Factories.Processors {
				components.Processors = append(components.Processors, prs)
			}
			for rcv := range set.Factories.Receivers {
				components.Receivers = append(components.Receivers, rcv)
			}
			for exp := range set.Factories.Exporters {
				components.Exporters = append(components.Exporters, exp)
=======
				components.Connectors = append(components.Connectors, componentWithStability{
					Name: con,
					Stability: map[string]string{
						"logs-to-logs":    set.Factories.Connectors[con].LogsToLogsStability().String(),
						"logs-to-metrics": set.Factories.Connectors[con].LogsToMetricsStability().String(),
						"logs-to-traces":  set.Factories.Connectors[con].LogsToTracesStability().String(),

						"metrics-to-logs":    set.Factories.Connectors[con].MetricsToLogsStability().String(),
						"metrics-to-metrics": set.Factories.Connectors[con].MetricsToMetricsStability().String(),
						"metrics-to-traces":  set.Factories.Connectors[con].MetricsToTracesStability().String(),

						"traces-to-logs":    set.Factories.Connectors[con].TracesToLogsStability().String(),
						"traces-to-metrics": set.Factories.Connectors[con].TracesToMetricsStability().String(),
						"traces-to-traces":  set.Factories.Connectors[con].TracesToTracesStability().String(),
					},
				})
			}
			for ext := range set.Factories.Extensions {
				components.Extensions = append(components.Extensions, componentWithStability{
					Name: ext,
					Stability: map[string]string{
						"extension": set.Factories.Extensions[ext].ExtensionStability().String(),
					},
				})
			}
			for prs := range set.Factories.Processors {
				components.Processors = append(components.Processors, componentWithStability{
					Name: prs,
					Stability: map[string]string{
						"logs":    set.Factories.Processors[prs].LogsProcessorStability().String(),
						"metrics": set.Factories.Processors[prs].MetricsProcessorStability().String(),
						"traces":  set.Factories.Processors[prs].TracesProcessorStability().String(),
					},
				})
			}
			for rcv := range set.Factories.Receivers {
				components.Receivers = append(components.Receivers, componentWithStability{
					Name: rcv,
					Stability: map[string]string{
						"logs":    set.Factories.Receivers[rcv].LogsReceiverStability().String(),
						"metrics": set.Factories.Receivers[rcv].MetricsReceiverStability().String(),
						"traces":  set.Factories.Receivers[rcv].TracesReceiverStability().String(),
					},
				})
			}
			for exp := range set.Factories.Exporters {
				components.Exporters = append(components.Exporters, componentWithStability{
					Name: exp,
					Stability: map[string]string{
						"logs":    set.Factories.Exporters[exp].LogsExporterStability().String(),
						"metrics": set.Factories.Exporters[exp].MetricsExporterStability().String(),
						"traces":  set.Factories.Exporters[exp].TracesExporterStability().String(),
					},
				})
>>>>>>> main
			}
			components.BuildInfo = set.BuildInfo
			yamlData, err := yaml.Marshal(components)
			if err != nil {
				return err
			}
			fmt.Fprint(cmd.OutOrStdout(), string(yamlData))
			return nil
		},
	}
}
