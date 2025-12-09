/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package arrow

// General configuration for traces. This configuration defines the different
// sorters used for attributes, spans, events, links, ... and the encoding used
// for parent IDs.

import (
	cfg "github.com/open-telemetry/otel-arrow/go/pkg/config"
	"github.com/open-telemetry/otel-arrow/go/pkg/otel/common/arrow"
)

type (
	Config struct {
		Global *cfg.Config

		Span  *SpanConfig
		Event *EventConfig
		Link  *LinkConfig
		Attrs *AttrsConfig
	}

	AttrsConfig struct {
		Resource *arrow.Attrs16Config
		Scope    *arrow.Attrs16Config
		Span     *arrow.Attrs16Config
		Event    *arrow.Attrs32Config
		Link     *arrow.Attrs32Config
	}

	SpanConfig struct {
		Sorter SpanSorter
	}

	EventConfig struct {
		Sorter EventSorter
	}

	LinkConfig struct {
		Sorter LinkSorter
	}
)

func DefaultConfig() *Config {
	return NewConfig(cfg.DefaultConfig())
}

func NewConfig(globalConf *cfg.Config) *Config {
	return &Config{
		Global: globalConf,
		Span: &SpanConfig{
			Sorter: SortSpansByResourceSpanIdScopeSpanIdNameTraceId(),
		},
		Event: &EventConfig{
			Sorter: SortEventsByNameParentId(),
		},
		Link: &LinkConfig{
			Sorter: SortLinksByTraceIdParentId(),
		},
		Attrs: &AttrsConfig{
			Resource: &arrow.Attrs16Config{
				Sorter: arrow.SortAttrs16ByTypeKeyValueParentId(),
			},
			Scope: &arrow.Attrs16Config{
				Sorter: arrow.SortAttrs16ByTypeKeyValueParentId(),
			},
			Span: &arrow.Attrs16Config{
				Sorter: arrow.SortAttrs16ByTypeKeyValueParentId(),
			},
			Event: &arrow.Attrs32Config{
				Sorter: arrow.SortAttrs32ByTypeKeyValueParentId(),
			},
			Link: &arrow.Attrs32Config{
				Sorter: arrow.SortAttrs32ByTypeKeyValueParentId(),
			},
		},
	}
}

func NewNoSortConfig(globalConf *cfg.Config) *Config {
	return &Config{
		Global: globalConf,
		Span: &SpanConfig{
			Sorter: UnsortedSpans(),
		},
		Event: &EventConfig{
			Sorter: UnsortedEvents(),
		},
		Link: &LinkConfig{
			Sorter: UnsortedLinks(),
		},
		Attrs: &AttrsConfig{
			Resource: &arrow.Attrs16Config{
				Sorter: arrow.UnsortedAttrs16(),
			},
			Scope: &arrow.Attrs16Config{
				Sorter: arrow.UnsortedAttrs16(),
			},
			Span: &arrow.Attrs16Config{
				Sorter: arrow.UnsortedAttrs16(),
			},
			Event: &arrow.Attrs32Config{
				Sorter: arrow.UnsortedAttrs32(),
			},
			Link: &arrow.Attrs32Config{
				Sorter: arrow.UnsortedAttrs32(),
			},
		},
	}
}
