// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/multierr"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal"
)

var _ exporter.Traces = (*traceExporterImp)(nil)

type exporterTraces map[*wrappedExporter]ptrace.Traces

type traceExporterImp struct {
	loadBalancer *loadBalancer
	routingKey   routingKey

	stopped    bool
	shutdownWg sync.WaitGroup
}

// Create new traces exporter
func newTracesExporter(params exporter.CreateSettings, cfg component.Config) (*traceExporterImp, error) {
	exporterFactory := otlpexporter.NewFactory()

	lb, err := newLoadBalancer(params, cfg, func(ctx context.Context, endpoint string) (component.Component, error) {
		oCfg := buildExporterConfig(cfg.(*Config), endpoint)
		return exporterFactory.CreateTracesExporter(ctx, params, &oCfg)
	})
	if err != nil {
		return nil, err
	}

	traceExporter := traceExporterImp{loadBalancer: lb, routingKey: traceIDRouting}

	switch cfg.(*Config).RoutingKey {
	case "service":
		traceExporter.routingKey = svcRouting
	case "traceID", "":
	default:
		return nil, fmt.Errorf("unsupported routing_key: %s", cfg.(*Config).RoutingKey)
	}
	return &traceExporter, nil
}

func buildExporterConfig(cfg *Config, endpoint string) otlpexporter.Config {
	oCfg := cfg.Protocol.OTLP
	oCfg.Endpoint = endpoint
	return oCfg
}

func (e *traceExporterImp) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (e *traceExporterImp) Start(ctx context.Context, host component.Host) error {
	return e.loadBalancer.Start(ctx, host)
}

func (e *traceExporterImp) Shutdown(ctx context.Context) error {
	err := e.loadBalancer.Shutdown(ctx)
	e.stopped = true
	e.shutdownWg.Wait()
	return err
}

func (e *traceExporterImp) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	batches := batchpersignal.SplitTraces(td)

	exporterSegregatedTraces := make(exporterTraces)
	endpoints := make(map[*wrappedExporter]string)
	for _, batch := range batches {
		routingID, err := routingIdentifiersFromTraces(batch, e.routingKey)
		if err != nil {
			return err
		}

		for rid := range routingID {
			exp, endpoint, err := e.loadBalancer.exporterAndEndpoint([]byte(rid))
			if err != nil {
				return err
			}

			_, ok := exporterSegregatedTraces[exp]
			if !ok {
				exp.consumeWG.Add(1)
				exporterSegregatedTraces[exp] = ptrace.NewTraces()
			}
			exporterSegregatedTraces[exp] = mergeTraces(exporterSegregatedTraces[exp], batch)

			endpoints[exp] = endpoint
		}
	}

	var errs error

	for exp, td := range exporterSegregatedTraces {
		start := time.Now()
		err := exp.ConsumeTraces(ctx, td)
		exp.consumeWG.Done()
		errs = multierr.Append(errs, err)
		duration := time.Since(start)

		if err == nil {
			_ = stats.RecordWithTags(
				ctx,
				[]tag.Mutator{tag.Upsert(endpointTagKey, endpoints[exp]), successTrueMutator},
				mBackendLatency.M(duration.Milliseconds()))
		} else {
			_ = stats.RecordWithTags(
				ctx,
				[]tag.Mutator{tag.Upsert(endpointTagKey, endpoints[exp]), successFalseMutator},
				mBackendLatency.M(duration.Milliseconds()))
		}
	}

	return errs
}

func routingIdentifiersFromTraces(td ptrace.Traces, key routingKey) (map[string]bool, error) {
	ids := make(map[string]bool)
	rs := td.ResourceSpans()
	if rs.Len() == 0 {
		return nil, errors.New("empty resource spans")
	}

	ils := rs.At(0).ScopeSpans()
	if ils.Len() == 0 {
		return nil, errors.New("empty scope spans")
	}

	spans := ils.At(0).Spans()
	if spans.Len() == 0 {
		return nil, errors.New("empty spans")
	}

	if key == svcRouting {
		for i := 0; i < rs.Len(); i++ {
			svc, ok := rs.At(i).Resource().Attributes().Get("service.name")
			if !ok {
				return nil, errors.New("unable to get service name")
			}
			ids[svc.Str()] = true
		}
		return ids, nil
	}
	tid := spans.At(0).TraceID()
	ids[string(tid[:])] = true
	return ids, nil
}
