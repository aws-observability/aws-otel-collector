// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package graph // import "go.opentelemetry.io/collector/service/internal/graph"

import (
	"context"
	"fmt"
	"hash/fnv"
	"strings"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentprofiles"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/connector/connectorprofiles"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumerprofiles"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/internal/fanoutconsumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/service/internal/builders"
	"go.opentelemetry.io/collector/service/internal/capabilityconsumer"
	"go.opentelemetry.io/collector/service/internal/components"
)

const (
	receiverSeed      = "receiver"
	processorSeed     = "processor"
	exporterSeed      = "exporter"
	connectorSeed     = "connector"
	capabilitiesSeed  = "capabilities"
	fanOutToExporters = "fanout_to_exporters"
)

// baseConsumer redeclared here since not public in consumer package. May consider to make that public.
type baseConsumer interface {
	Capabilities() consumer.Capabilities
}

type nodeID int64

func (n nodeID) ID() int64 {
	return int64(n)
}

func newNodeID(parts ...string) nodeID {
	h := fnv.New64a()
	h.Write([]byte(strings.Join(parts, "|")))
	return nodeID(h.Sum64())
}

type consumerNode interface {
	getConsumer() baseConsumer
}

// A receiver instance can be shared by multiple pipelines of the same type.
// Therefore, nodeID is derived from "pipeline type" and "component ID".
type receiverNode struct {
	nodeID
	componentID  component.ID
	pipelineType component.DataType
	component.Component
}

func newReceiverNode(pipelineType component.DataType, recvID component.ID) *receiverNode {
	return &receiverNode{
		nodeID:       newNodeID(receiverSeed, pipelineType.String(), recvID.String()),
		componentID:  recvID,
		pipelineType: pipelineType,
	}
}

func (n *receiverNode) buildComponent(ctx context.Context,
	tel component.TelemetrySettings,
	info component.BuildInfo,
	builder *builders.ReceiverBuilder,
	nexts []baseConsumer,
) error {
	tel.Logger = components.ReceiverLogger(tel.Logger, n.componentID, n.pipelineType)
	set := receiver.Settings{ID: n.componentID, TelemetrySettings: tel, BuildInfo: info}
	var err error
	switch n.pipelineType {
	case component.DataTypeTraces:
		var consumers []consumer.Traces
		for _, next := range nexts {
			consumers = append(consumers, next.(consumer.Traces))
		}
		n.Component, err = builder.CreateTraces(ctx, set, fanoutconsumer.NewTraces(consumers))
	case component.DataTypeMetrics:
		var consumers []consumer.Metrics
		for _, next := range nexts {
			consumers = append(consumers, next.(consumer.Metrics))
		}
		n.Component, err = builder.CreateMetrics(ctx, set, fanoutconsumer.NewMetrics(consumers))
	case component.DataTypeLogs:
		var consumers []consumer.Logs
		for _, next := range nexts {
			consumers = append(consumers, next.(consumer.Logs))
		}
		n.Component, err = builder.CreateLogs(ctx, set, fanoutconsumer.NewLogs(consumers))
	case componentprofiles.DataTypeProfiles:
		var consumers []consumerprofiles.Profiles
		for _, next := range nexts {
			consumers = append(consumers, next.(consumerprofiles.Profiles))
		}
		n.Component, err = builder.CreateProfiles(ctx, set, fanoutconsumer.NewProfiles(consumers))
	default:
		return fmt.Errorf("error creating receiver %q for data type %q is not supported", set.ID, n.pipelineType)
	}
	if err != nil {
		return fmt.Errorf("failed to create %q receiver for data type %q: %w", set.ID, n.pipelineType, err)
	}
	return nil
}

var _ consumerNode = (*processorNode)(nil)

// Every processor instance is unique to one pipeline.
// Therefore, nodeID is derived from "pipeline ID" and "component ID".
type processorNode struct {
	nodeID
	componentID component.ID
	pipelineID  component.ID
	component.Component
}

func newProcessorNode(pipelineID, procID component.ID) *processorNode {
	return &processorNode{
		nodeID:      newNodeID(processorSeed, pipelineID.String(), procID.String()),
		componentID: procID,
		pipelineID:  pipelineID,
	}
}

func (n *processorNode) getConsumer() baseConsumer {
	return n.Component.(baseConsumer)
}

func (n *processorNode) buildComponent(ctx context.Context,
	tel component.TelemetrySettings,
	info component.BuildInfo,
	builder *builders.ProcessorBuilder,
	next baseConsumer,
) error {
	tel.Logger = components.ProcessorLogger(tel.Logger, n.componentID, n.pipelineID)
	set := processor.Settings{ID: n.componentID, TelemetrySettings: tel, BuildInfo: info}
	var err error
	switch n.pipelineID.Type() {
	case component.DataTypeTraces:
		n.Component, err = builder.CreateTraces(ctx, set, next.(consumer.Traces))
	case component.DataTypeMetrics:
		n.Component, err = builder.CreateMetrics(ctx, set, next.(consumer.Metrics))
	case component.DataTypeLogs:
		n.Component, err = builder.CreateLogs(ctx, set, next.(consumer.Logs))
	case componentprofiles.DataTypeProfiles:
		n.Component, err = builder.CreateProfiles(ctx, set, next.(consumerprofiles.Profiles))
	default:
		return fmt.Errorf("error creating processor %q in pipeline %q, data type %q is not supported", set.ID, n.pipelineID, n.pipelineID.Type())
	}
	if err != nil {
		return fmt.Errorf("failed to create %q processor, in pipeline %q: %w", set.ID, n.pipelineID, err)
	}
	return nil
}

var _ consumerNode = (*exporterNode)(nil)

// An exporter instance can be shared by multiple pipelines of the same type.
// Therefore, nodeID is derived from "pipeline type" and "component ID".
type exporterNode struct {
	nodeID
	componentID  component.ID
	pipelineType component.DataType
	component.Component
}

func newExporterNode(pipelineType component.DataType, exprID component.ID) *exporterNode {
	return &exporterNode{
		nodeID:       newNodeID(exporterSeed, pipelineType.String(), exprID.String()),
		componentID:  exprID,
		pipelineType: pipelineType,
	}
}

func (n *exporterNode) getConsumer() baseConsumer {
	return n.Component.(baseConsumer)
}

func (n *exporterNode) buildComponent(
	ctx context.Context,
	tel component.TelemetrySettings,
	info component.BuildInfo,
	builder *builders.ExporterBuilder,
) error {
	tel.Logger = components.ExporterLogger(tel.Logger, n.componentID, n.pipelineType)
	set := exporter.Settings{ID: n.componentID, TelemetrySettings: tel, BuildInfo: info}
	var err error
	switch n.pipelineType {
	case component.DataTypeTraces:
		n.Component, err = builder.CreateTraces(ctx, set)
	case component.DataTypeMetrics:
		n.Component, err = builder.CreateMetrics(ctx, set)
	case component.DataTypeLogs:
		n.Component, err = builder.CreateLogs(ctx, set)
	case componentprofiles.DataTypeProfiles:
		n.Component, err = builder.CreateProfiles(ctx, set)
	default:
		return fmt.Errorf("error creating exporter %q for data type %q is not supported", set.ID, n.pipelineType)
	}
	if err != nil {
		return fmt.Errorf("failed to create %q exporter for data type %q: %w", set.ID, n.pipelineType, err)
	}
	return nil
}

var _ consumerNode = (*connectorNode)(nil)

// A connector instance connects one pipeline type to one other pipeline type.
// Therefore, nodeID is derived from "exporter pipeline type", "receiver pipeline type", and "component ID".
type connectorNode struct {
	nodeID
	componentID      component.ID
	exprPipelineType component.DataType
	rcvrPipelineType component.DataType
	component.Component
	baseConsumer
}

func newConnectorNode(exprPipelineType, rcvrPipelineType component.DataType, connID component.ID) *connectorNode {
	return &connectorNode{
		nodeID:           newNodeID(connectorSeed, connID.String(), exprPipelineType.String(), rcvrPipelineType.String()),
		componentID:      connID,
		exprPipelineType: exprPipelineType,
		rcvrPipelineType: rcvrPipelineType,
	}
}

func (n *connectorNode) getConsumer() baseConsumer {
	return n.baseConsumer
}

func (n *connectorNode) buildComponent(
	ctx context.Context,
	tel component.TelemetrySettings,
	info component.BuildInfo,
	builder *builders.ConnectorBuilder,
	nexts []baseConsumer,
) error {
	tel.Logger = components.ConnectorLogger(tel.Logger, n.componentID, n.exprPipelineType, n.rcvrPipelineType)
	set := connector.Settings{ID: n.componentID, TelemetrySettings: tel, BuildInfo: info}

	switch n.rcvrPipelineType {
	case component.DataTypeTraces:
		capability := consumer.Capabilities{MutatesData: false}
		consumers := make(map[component.ID]consumer.Traces, len(nexts))
		for _, next := range nexts {
			consumers[next.(*capabilitiesNode).pipelineID] = next.(consumer.Traces)
			capability.MutatesData = capability.MutatesData || next.Capabilities().MutatesData
		}
		next := connector.NewTracesRouter(consumers)

		switch n.exprPipelineType {
		case component.DataTypeTraces:
			conn, err := builder.CreateTracesToTraces(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component = conn
			// When connecting pipelines of the same data type, the connector must
			// inherit the capabilities of pipelines in which it is acting as a receiver.
			// Since the incoming and outgoing data types are the same, we must also consider
			// that the connector itself may MutatesData.
			capability.MutatesData = capability.MutatesData || conn.Capabilities().MutatesData
			n.baseConsumer = capabilityconsumer.NewTraces(conn, capability)
		case component.DataTypeMetrics:
			conn, err := builder.CreateMetricsToTraces(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		case component.DataTypeLogs:
			conn, err := builder.CreateLogsToTraces(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		case componentprofiles.DataTypeProfiles:
			conn, err := builder.CreateProfilesToTraces(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		}

	case component.DataTypeMetrics:
		capability := consumer.Capabilities{MutatesData: false}
		consumers := make(map[component.ID]consumer.Metrics, len(nexts))
		for _, next := range nexts {
			consumers[next.(*capabilitiesNode).pipelineID] = next.(consumer.Metrics)
			capability.MutatesData = capability.MutatesData || next.Capabilities().MutatesData
		}
		next := connector.NewMetricsRouter(consumers)

		switch n.exprPipelineType {
		case component.DataTypeTraces:
			conn, err := builder.CreateTracesToMetrics(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		case component.DataTypeMetrics:
			conn, err := builder.CreateMetricsToMetrics(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component = conn
			// When connecting pipelines of the same data type, the connector must
			// inherit the capabilities of pipelines in which it is acting as a receiver.
			// Since the incoming and outgoing data types are the same, we must also consider
			// that the connector itself may MutatesData.
			capability.MutatesData = capability.MutatesData || conn.Capabilities().MutatesData
			n.baseConsumer = capabilityconsumer.NewMetrics(conn, capability)
		case component.DataTypeLogs:
			conn, err := builder.CreateLogsToMetrics(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		case componentprofiles.DataTypeProfiles:
			conn, err := builder.CreateProfilesToMetrics(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		}
	case component.DataTypeLogs:
		capability := consumer.Capabilities{MutatesData: false}
		consumers := make(map[component.ID]consumer.Logs, len(nexts))
		for _, next := range nexts {
			consumers[next.(*capabilitiesNode).pipelineID] = next.(consumer.Logs)
			capability.MutatesData = capability.MutatesData || next.Capabilities().MutatesData
		}
		next := connector.NewLogsRouter(consumers)

		switch n.exprPipelineType {
		case component.DataTypeTraces:
			conn, err := builder.CreateTracesToLogs(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		case component.DataTypeMetrics:
			conn, err := builder.CreateMetricsToLogs(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		case component.DataTypeLogs:
			conn, err := builder.CreateLogsToLogs(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component = conn
			// When connecting pipelines of the same data type, the connector must
			// inherit the capabilities of pipelines in which it is acting as a receiver.
			// Since the incoming and outgoing data types are the same, we must also consider
			// that the connector itself may MutatesData.
			capability.MutatesData = capability.MutatesData || conn.Capabilities().MutatesData
			n.baseConsumer = capabilityconsumer.NewLogs(conn, capability)
		case componentprofiles.DataTypeProfiles:
			conn, err := builder.CreateProfilesToLogs(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		}
	case componentprofiles.DataTypeProfiles:
		capability := consumer.Capabilities{MutatesData: false}
		consumers := make(map[component.ID]consumerprofiles.Profiles, len(nexts))
		for _, next := range nexts {
			consumers[next.(*capabilitiesNode).pipelineID] = next.(consumerprofiles.Profiles)
			capability.MutatesData = capability.MutatesData || next.Capabilities().MutatesData
		}
		next := connectorprofiles.NewProfilesRouter(consumers)

		switch n.exprPipelineType {
		case component.DataTypeTraces:
			conn, err := builder.CreateTracesToProfiles(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		case component.DataTypeMetrics:
			conn, err := builder.CreateMetricsToProfiles(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		case component.DataTypeLogs:
			conn, err := builder.CreateLogsToProfiles(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component, n.baseConsumer = conn, conn
		case componentprofiles.DataTypeProfiles:
			conn, err := builder.CreateProfilesToProfiles(ctx, set, next)
			if err != nil {
				return err
			}
			n.Component = conn
			// When connecting pipelines of the same data type, the connector must
			// inherit the capabilities of pipelines in which it is acting as a receiver.
			// Since the incoming and outgoing data types are the same, we must also consider
			// that the connector itself may MutatesData.
			capability.MutatesData = capability.MutatesData || conn.Capabilities().MutatesData
			n.baseConsumer = capabilityconsumer.NewProfiles(conn, capability)
		}
	}
	return nil
}

var _ consumerNode = (*capabilitiesNode)(nil)

// Every pipeline has a "virtual" capabilities node immediately after the receiver(s).
// There are two purposes for this node:
// 1. Present aggregated capabilities to receivers, such as whether the pipeline mutates data.
// 2. Present a consistent "first consumer" for each pipeline.
// The nodeID is derived from "pipeline ID".
type capabilitiesNode struct {
	nodeID
	pipelineID component.ID
	baseConsumer
	consumer.ConsumeTracesFunc
	consumer.ConsumeMetricsFunc
	consumer.ConsumeLogsFunc
	consumerprofiles.ConsumeProfilesFunc
}

func newCapabilitiesNode(pipelineID component.ID) *capabilitiesNode {
	return &capabilitiesNode{
		nodeID:     newNodeID(capabilitiesSeed, pipelineID.String()),
		pipelineID: pipelineID,
	}
}

func (n *capabilitiesNode) getConsumer() baseConsumer {
	return n
}

var _ consumerNode = (*fanOutNode)(nil)

// Each pipeline has one fan-out node before exporters.
// Therefore, nodeID is derived from "pipeline ID".
type fanOutNode struct {
	nodeID
	pipelineID component.ID
	baseConsumer
}

func newFanOutNode(pipelineID component.ID) *fanOutNode {
	return &fanOutNode{
		nodeID:     newNodeID(fanOutToExporters, pipelineID.String()),
		pipelineID: pipelineID,
	}
}

func (n *fanOutNode) getConsumer() baseConsumer {
	return n.baseConsumer
}
