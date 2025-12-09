/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package arrow

// General configuration for logs. This configuration defines the different
// sorters used for attributes, log records, ... and the encoding used for
// parent IDs.

import (
	cfg "github.com/open-telemetry/otel-arrow/go/pkg/config"
	"github.com/open-telemetry/otel-arrow/go/pkg/otel/common/arrow"
)

type (
	Config struct {
		Global *cfg.Config

		Log   *LogConfig
		Attrs *AttrsConfig
	}

	AttrsConfig struct {
		Resource *arrow.Attrs16Config
		Scope    *arrow.Attrs16Config
		Log      *arrow.Attrs16Config
	}

	LogConfig struct {
		Sorter LogSorter
	}
)

func DefaultConfig() *Config {
	return NewConfig(cfg.DefaultConfig())
}

func NewConfig(globalConf *cfg.Config) *Config {
	return &Config{
		Global: globalConf,
		Log: &LogConfig{
			Sorter: SortLogsByResourceLogsIDScopeLogsIDTraceID(),
		},
		Attrs: &AttrsConfig{
			Resource: &arrow.Attrs16Config{
				Sorter: arrow.SortAttrs16ByTypeKeyValueParentId(),
			},
			Scope: &arrow.Attrs16Config{
				Sorter: arrow.SortAttrs16ByTypeKeyValueParentId(),
			},
			Log: &arrow.Attrs16Config{
				Sorter: arrow.SortAttrs16ByTypeKeyValueParentId(),
			},
		},
	}
}

func NewNoSortConfig(globalConf *cfg.Config) *Config {
	return &Config{
		Global: globalConf,
		Log: &LogConfig{
			Sorter: UnsortedLogs(),
		},
		Attrs: &AttrsConfig{
			Resource: &arrow.Attrs16Config{
				Sorter: arrow.UnsortedAttrs16(),
			},
			Scope: &arrow.Attrs16Config{
				Sorter: arrow.UnsortedAttrs16(),
			},
			Log: &arrow.Attrs16Config{
				Sorter: arrow.UnsortedAttrs16(),
			},
		},
	}
}
