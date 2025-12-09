// THIS IS A GENERATED FILE
// DO NOT EDIT
package process

import (
	bytes "bytes"
	protowire "google.golang.org/protobuf/encoding/protowire"
	io "io"
)

type CollectorConnectionsBuilder struct {
	writer                                                       io.Writer
	buf                                                          bytes.Buffer
	scratch                                                      []byte
	connectionBuilder                                            ConnectionBuilder
	collectorConnections_ResolvedResourcesEntryBuilder           CollectorConnections_ResolvedResourcesEntryBuilder
	collectorConnections_ContainerForPidEntryBuilder             CollectorConnections_ContainerForPidEntryBuilder
	collectorConnectionsTelemetryBuilder                         CollectorConnectionsTelemetryBuilder
	collectorConnections_ConnTelemetryMapEntryBuilder            CollectorConnections_ConnTelemetryMapEntryBuilder
	collectorConnections_CompilationTelemetryByAssetEntryBuilder CollectorConnections_CompilationTelemetryByAssetEntryBuilder
	collectorConnections_CORETelemetryByAssetEntryBuilder        CollectorConnections_CORETelemetryByAssetEntryBuilder
	routeBuilder                                                 RouteBuilder
	routeMetadataBuilder                                         RouteMetadataBuilder
	agentConfigurationBuilder                                    AgentConfigurationBuilder
	collectorConnections_ResolvedHostsByNameEntryBuilder         CollectorConnections_ResolvedHostsByNameEntryBuilder
	collectorConnections_ResolvedPublicIpsEntryBuilder           CollectorConnections_ResolvedPublicIpsEntryBuilder
}

func NewCollectorConnectionsBuilder(writer io.Writer) *CollectorConnectionsBuilder {
	return &CollectorConnectionsBuilder{
		writer: writer,
	}
}
func (x *CollectorConnectionsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorConnectionsBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) SetNetworkId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x62)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) AddConnections(cb func(w *ConnectionBuilder)) {
	x.buf.Reset()
	x.connectionBuilder.writer = &x.buf
	x.connectionBuilder.scratch = x.scratch
	cb(&x.connectionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) AddResolvedResources(cb func(w *CollectorConnections_ResolvedResourcesEntryBuilder)) {
	x.buf.Reset()
	x.collectorConnections_ResolvedResourcesEntryBuilder.writer = &x.buf
	x.collectorConnections_ResolvedResourcesEntryBuilder.scratch = x.scratch
	cb(&x.collectorConnections_ResolvedResourcesEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) AddContainerForPid(cb func(w *CollectorConnections_ContainerForPidEntryBuilder)) {
	x.buf.Reset()
	x.collectorConnections_ContainerForPidEntryBuilder.writer = &x.buf
	x.collectorConnections_ContainerForPidEntryBuilder.scratch = x.scratch
	cb(&x.collectorConnections_ContainerForPidEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) SetEncodedTags(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) SetEncodedConnectionsTags(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x132)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) SetContainerHostType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x78)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *CollectorConnectionsBuilder) SetConnTelemetry(cb func(w *CollectorConnectionsTelemetryBuilder)) {
	x.buf.Reset()
	x.collectorConnectionsTelemetryBuilder.writer = &x.buf
	x.collectorConnectionsTelemetryBuilder.scratch = x.scratch
	cb(&x.collectorConnectionsTelemetryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x82)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) AddConnTelemetryMap(cb func(w *CollectorConnections_ConnTelemetryMapEntryBuilder)) {
	x.buf.Reset()
	x.collectorConnections_ConnTelemetryMapEntryBuilder.writer = &x.buf
	x.collectorConnections_ConnTelemetryMapEntryBuilder.scratch = x.scratch
	cb(&x.collectorConnections_ConnTelemetryMapEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x13a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) SetArchitecture(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) SetKernelVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x92)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) SetPlatform(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x9a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) SetPlatformVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa2)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) AddCompilationTelemetryByAsset(cb func(w *CollectorConnections_CompilationTelemetryByAssetEntryBuilder)) {
	x.buf.Reset()
	x.collectorConnections_CompilationTelemetryByAssetEntryBuilder.writer = &x.buf
	x.collectorConnections_CompilationTelemetryByAssetEntryBuilder.scratch = x.scratch
	cb(&x.collectorConnections_CompilationTelemetryByAssetEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xaa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) SetKernelHeaderFetchResult(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x148)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *CollectorConnectionsBuilder) AddCORETelemetryByAsset(cb func(w *CollectorConnections_CORETelemetryByAssetEntryBuilder)) {
	x.buf.Reset()
	x.collectorConnections_CORETelemetryByAssetEntryBuilder.writer = &x.buf
	x.collectorConnections_CORETelemetryByAssetEntryBuilder.scratch = x.scratch
	cb(&x.collectorConnections_CORETelemetryByAssetEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x152)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) AddPrebuiltEBPFAssets(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x162)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) AddRoutes(cb func(w *RouteBuilder)) {
	x.buf.Reset()
	x.routeBuilder.writer = &x.buf
	x.routeBuilder.scratch = x.scratch
	cb(&x.routeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xfa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) AddRouteMetadata(cb func(w *RouteMetadataBuilder)) {
	x.buf.Reset()
	x.routeMetadataBuilder.writer = &x.buf
	x.routeMetadataBuilder.scratch = x.scratch
	cb(&x.routeMetadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x112)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) SetAgentConfiguration(cb func(w *AgentConfigurationBuilder)) {
	x.buf.Reset()
	x.agentConfigurationBuilder.writer = &x.buf
	x.agentConfigurationBuilder.scratch = x.scratch
	cb(&x.agentConfigurationBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) SetEncodedDNS(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x72)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) AddDomains(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xf2)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) SetEncodedDomainDatabase(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x122)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) SetEncodedDnsLookups(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) AddResolvedHostsByName(cb func(w *CollectorConnections_ResolvedHostsByNameEntryBuilder)) {
	x.buf.Reset()
	x.collectorConnections_ResolvedHostsByNameEntryBuilder.writer = &x.buf
	x.collectorConnections_ResolvedHostsByNameEntryBuilder.scratch = x.scratch
	cb(&x.collectorConnections_ResolvedHostsByNameEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x142)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) SetEcsTask(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x16a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) AddResolvedPublicIps(cb func(w *CollectorConnections_ResolvedPublicIpsEntryBuilder)) {
	x.buf.Reset()
	x.collectorConnections_ResolvedPublicIpsEntryBuilder.writer = &x.buf
	x.collectorConnections_ResolvedPublicIpsEntryBuilder.scratch = x.scratch
	cb(&x.collectorConnections_ResolvedPublicIpsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x172)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorConnectionsBuilder) SetHostTagsIndex(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x178)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsBuilder) AddResolvConfs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x182)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorConnections_ResolvedResourcesEntryBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	resourceMetadataBuilder ResourceMetadataBuilder
}

func NewCollectorConnections_ResolvedResourcesEntryBuilder(writer io.Writer) *CollectorConnections_ResolvedResourcesEntryBuilder {
	return &CollectorConnections_ResolvedResourcesEntryBuilder{
		writer: writer,
	}
}
func (x *CollectorConnections_ResolvedResourcesEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorConnections_ResolvedResourcesEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnections_ResolvedResourcesEntryBuilder) SetValue(cb func(w *ResourceMetadataBuilder)) {
	x.buf.Reset()
	x.resourceMetadataBuilder.writer = &x.buf
	x.resourceMetadataBuilder.scratch = x.scratch
	cb(&x.resourceMetadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorConnections_ContainerForPidEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCollectorConnections_ContainerForPidEntryBuilder(writer io.Writer) *CollectorConnections_ContainerForPidEntryBuilder {
	return &CollectorConnections_ContainerForPidEntryBuilder{
		writer: writer,
	}
}
func (x *CollectorConnections_ContainerForPidEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorConnections_ContainerForPidEntryBuilder) SetKey(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnections_ContainerForPidEntryBuilder) SetValue(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorConnections_ConnTelemetryMapEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCollectorConnections_ConnTelemetryMapEntryBuilder(writer io.Writer) *CollectorConnections_ConnTelemetryMapEntryBuilder {
	return &CollectorConnections_ConnTelemetryMapEntryBuilder{
		writer: writer,
	}
}
func (x *CollectorConnections_ConnTelemetryMapEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorConnections_ConnTelemetryMapEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnections_ConnTelemetryMapEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type CollectorConnections_CompilationTelemetryByAssetEntryBuilder struct {
	writer                             io.Writer
	buf                                bytes.Buffer
	scratch                            []byte
	runtimeCompilationTelemetryBuilder RuntimeCompilationTelemetryBuilder
}

func NewCollectorConnections_CompilationTelemetryByAssetEntryBuilder(writer io.Writer) *CollectorConnections_CompilationTelemetryByAssetEntryBuilder {
	return &CollectorConnections_CompilationTelemetryByAssetEntryBuilder{
		writer: writer,
	}
}
func (x *CollectorConnections_CompilationTelemetryByAssetEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorConnections_CompilationTelemetryByAssetEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnections_CompilationTelemetryByAssetEntryBuilder) SetValue(cb func(w *RuntimeCompilationTelemetryBuilder)) {
	x.buf.Reset()
	x.runtimeCompilationTelemetryBuilder.writer = &x.buf
	x.runtimeCompilationTelemetryBuilder.scratch = x.scratch
	cb(&x.runtimeCompilationTelemetryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorConnections_CORETelemetryByAssetEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCollectorConnections_CORETelemetryByAssetEntryBuilder(writer io.Writer) *CollectorConnections_CORETelemetryByAssetEntryBuilder {
	return &CollectorConnections_CORETelemetryByAssetEntryBuilder{
		writer: writer,
	}
}
func (x *CollectorConnections_CORETelemetryByAssetEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorConnections_CORETelemetryByAssetEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnections_CORETelemetryByAssetEntryBuilder) SetValue(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x10)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type CollectorConnections_ResolvedHostsByNameEntryBuilder struct {
	writer      io.Writer
	buf         bytes.Buffer
	scratch     []byte
	hostBuilder HostBuilder
}

func NewCollectorConnections_ResolvedHostsByNameEntryBuilder(writer io.Writer) *CollectorConnections_ResolvedHostsByNameEntryBuilder {
	return &CollectorConnections_ResolvedHostsByNameEntryBuilder{
		writer: writer,
	}
}
func (x *CollectorConnections_ResolvedHostsByNameEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorConnections_ResolvedHostsByNameEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnections_ResolvedHostsByNameEntryBuilder) SetValue(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorConnections_ResolvedPublicIpsEntryBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	publicIpMetadataBuilder PublicIpMetadataBuilder
}

func NewCollectorConnections_ResolvedPublicIpsEntryBuilder(writer io.Writer) *CollectorConnections_ResolvedPublicIpsEntryBuilder {
	return &CollectorConnections_ResolvedPublicIpsEntryBuilder{
		writer: writer,
	}
}
func (x *CollectorConnections_ResolvedPublicIpsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorConnections_ResolvedPublicIpsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorConnections_ResolvedPublicIpsEntryBuilder) SetValue(cb func(w *PublicIpMetadataBuilder)) {
	x.buf.Reset()
	x.publicIpMetadataBuilder.writer = &x.buf
	x.publicIpMetadataBuilder.scratch = x.scratch
	cb(&x.publicIpMetadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ConnectionsBuilder struct {
	writer                                              io.Writer
	buf                                                 bytes.Buffer
	scratch                                             []byte
	connectionBuilder                                   ConnectionBuilder
	connections_DnsEntryBuilder                         Connections_DnsEntryBuilder
	connectionsTelemetryBuilder                         ConnectionsTelemetryBuilder
	routeBuilder                                        RouteBuilder
	connections_CompilationTelemetryByAssetEntryBuilder Connections_CompilationTelemetryByAssetEntryBuilder
	agentConfigurationBuilder                           AgentConfigurationBuilder
	connections_ConnTelemetryMapEntryBuilder            Connections_ConnTelemetryMapEntryBuilder
	connections_CORETelemetryByAssetEntryBuilder        Connections_CORETelemetryByAssetEntryBuilder
}

func NewConnectionsBuilder(writer io.Writer) *ConnectionsBuilder {
	return &ConnectionsBuilder{
		writer: writer,
	}
}
func (x *ConnectionsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ConnectionsBuilder) AddConns(cb func(w *ConnectionBuilder)) {
	x.buf.Reset()
	x.connectionBuilder.writer = &x.buf
	x.connectionBuilder.scratch = x.scratch
	cb(&x.connectionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionsBuilder) AddDns(cb func(w *Connections_DnsEntryBuilder)) {
	x.buf.Reset()
	x.connections_DnsEntryBuilder.writer = &x.buf
	x.connections_DnsEntryBuilder.scratch = x.scratch
	cb(&x.connections_DnsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionsBuilder) SetConnTelemetry(cb func(w *ConnectionsTelemetryBuilder)) {
	x.buf.Reset()
	x.connectionsTelemetryBuilder.writer = &x.buf
	x.connectionsTelemetryBuilder.scratch = x.scratch
	cb(&x.connectionsTelemetryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionsBuilder) AddDomains(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ConnectionsBuilder) AddRoutes(cb func(w *RouteBuilder)) {
	x.buf.Reset()
	x.routeBuilder.writer = &x.buf
	x.routeBuilder.scratch = x.scratch
	cb(&x.routeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionsBuilder) AddCompilationTelemetryByAsset(cb func(w *Connections_CompilationTelemetryByAssetEntryBuilder)) {
	x.buf.Reset()
	x.connections_CompilationTelemetryByAssetEntryBuilder.writer = &x.buf
	x.connections_CompilationTelemetryByAssetEntryBuilder.scratch = x.scratch
	cb(&x.connections_CompilationTelemetryByAssetEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionsBuilder) SetAgentConfiguration(cb func(w *AgentConfigurationBuilder)) {
	x.buf.Reset()
	x.agentConfigurationBuilder.writer = &x.buf
	x.agentConfigurationBuilder.scratch = x.scratch
	cb(&x.agentConfigurationBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionsBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ConnectionsBuilder) AddConnTelemetryMap(cb func(w *Connections_ConnTelemetryMapEntryBuilder)) {
	x.buf.Reset()
	x.connections_ConnTelemetryMapEntryBuilder.writer = &x.buf
	x.connections_ConnTelemetryMapEntryBuilder.scratch = x.scratch
	cb(&x.connections_ConnTelemetryMapEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x4a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionsBuilder) SetKernelHeaderFetchResult(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x50)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ConnectionsBuilder) AddCORETelemetryByAsset(cb func(w *Connections_CORETelemetryByAssetEntryBuilder)) {
	x.buf.Reset()
	x.connections_CORETelemetryByAssetEntryBuilder.writer = &x.buf
	x.connections_CORETelemetryByAssetEntryBuilder.scratch = x.scratch
	cb(&x.connections_CORETelemetryByAssetEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionsBuilder) AddPrebuiltEBPFAssets(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x62)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ConnectionsBuilder) AddResolvConfs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x6a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type Connections_DnsEntryBuilder struct {
	writer          io.Writer
	buf             bytes.Buffer
	scratch         []byte
	dNSEntryBuilder DNSEntryBuilder
}

func NewConnections_DnsEntryBuilder(writer io.Writer) *Connections_DnsEntryBuilder {
	return &Connections_DnsEntryBuilder{
		writer: writer,
	}
}
func (x *Connections_DnsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Connections_DnsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *Connections_DnsEntryBuilder) SetValue(cb func(w *DNSEntryBuilder)) {
	x.buf.Reset()
	x.dNSEntryBuilder.writer = &x.buf
	x.dNSEntryBuilder.scratch = x.scratch
	cb(&x.dNSEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type Connections_CompilationTelemetryByAssetEntryBuilder struct {
	writer                             io.Writer
	buf                                bytes.Buffer
	scratch                            []byte
	runtimeCompilationTelemetryBuilder RuntimeCompilationTelemetryBuilder
}

func NewConnections_CompilationTelemetryByAssetEntryBuilder(writer io.Writer) *Connections_CompilationTelemetryByAssetEntryBuilder {
	return &Connections_CompilationTelemetryByAssetEntryBuilder{
		writer: writer,
	}
}
func (x *Connections_CompilationTelemetryByAssetEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Connections_CompilationTelemetryByAssetEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *Connections_CompilationTelemetryByAssetEntryBuilder) SetValue(cb func(w *RuntimeCompilationTelemetryBuilder)) {
	x.buf.Reset()
	x.runtimeCompilationTelemetryBuilder.writer = &x.buf
	x.runtimeCompilationTelemetryBuilder.scratch = x.scratch
	cb(&x.runtimeCompilationTelemetryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type Connections_ConnTelemetryMapEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewConnections_ConnTelemetryMapEntryBuilder(writer io.Writer) *Connections_ConnTelemetryMapEntryBuilder {
	return &Connections_ConnTelemetryMapEntryBuilder{
		writer: writer,
	}
}
func (x *Connections_ConnTelemetryMapEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Connections_ConnTelemetryMapEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *Connections_ConnTelemetryMapEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type Connections_CORETelemetryByAssetEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewConnections_CORETelemetryByAssetEntryBuilder(writer io.Writer) *Connections_CORETelemetryByAssetEntryBuilder {
	return &Connections_CORETelemetryByAssetEntryBuilder{
		writer: writer,
	}
}
func (x *Connections_CORETelemetryByAssetEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Connections_CORETelemetryByAssetEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *Connections_CORETelemetryByAssetEntryBuilder) SetValue(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x10)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type ConnectionBuilder struct {
	writer                                                   io.Writer
	buf                                                      bytes.Buffer
	scratch                                                  []byte
	addrBuilder                                              AddrBuilder
	protocolStackBuilder                                     ProtocolStackBuilder
	iPTranslationBuilder                                     IPTranslationBuilder
	connection_DnsCountByRcodeEntryBuilder                   Connection_DnsCountByRcodeEntryBuilder
	connection_DnsStatsByDomainEntryBuilder                  Connection_DnsStatsByDomainEntryBuilder
	connection_DnsStatsByDomainByQueryTypeEntryBuilder       Connection_DnsStatsByDomainByQueryTypeEntryBuilder
	connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder Connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder
	connection_TcpFailuresByErrCodeEntryBuilder              Connection_TcpFailuresByErrCodeEntryBuilder
}

func NewConnectionBuilder(writer io.Writer) *ConnectionBuilder {
	return &ConnectionBuilder{
		writer: writer,
	}
}
func (x *ConnectionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ConnectionBuilder) SetPid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetLaddr(cb func(w *AddrBuilder)) {
	x.buf.Reset()
	x.addrBuilder.writer = &x.buf
	x.addrBuilder.scratch = x.scratch
	cb(&x.addrBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) SetRaddr(cb func(w *AddrBuilder)) {
	x.buf.Reset()
	x.addrBuilder.writer = &x.buf
	x.addrBuilder.scratch = x.scratch
	cb(&x.addrBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) SetFamily(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x50)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ConnectionBuilder) SetType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x58)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ConnectionBuilder) SetIsLocalPortEphemeral(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x148)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ConnectionBuilder) SetLastBytesSent(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x80)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetLastBytesReceived(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x88)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetLastRetransmits(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x90)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetDirection(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x98)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ConnectionBuilder) SetLastPacketsSent(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x130)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetLastPacketsReceived(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x138)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetProtocol(cb func(w *ProtocolStackBuilder)) {
	x.buf.Reset()
	x.protocolStackBuilder.writer = &x.buf
	x.protocolStackBuilder.scratch = x.scratch
	cb(&x.protocolStackBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x182)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) SetNetNS(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetRemoteNetworkId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x102)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetIpTranslation(cb func(w *IPTranslationBuilder)) {
	x.buf.Reset()
	x.iPTranslationBuilder.writer = &x.buf
	x.iPTranslationBuilder.scratch = x.scratch
	cb(&x.iPTranslationBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xaa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) SetRtt(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xb0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetRttVar(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xb8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetIntraHost(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0xc0)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *ConnectionBuilder) SetDnsSuccessfulResponses(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xc8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetDnsFailedResponses(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xd0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetDnsTimeouts(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xd8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetDnsSuccessLatencySum(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xe0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetDnsFailureLatencySum(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xe8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) AddDnsCountByRcode(cb func(w *Connection_DnsCountByRcodeEntryBuilder)) {
	x.buf.Reset()
	x.connection_DnsCountByRcodeEntryBuilder.writer = &x.buf
	x.connection_DnsCountByRcodeEntryBuilder.scratch = x.scratch
	cb(&x.connection_DnsCountByRcodeEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x10a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) SetLastTcpEstablished(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xf0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetLastTcpClosed(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xf8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) AddDnsStatsByDomain(cb func(w *Connection_DnsStatsByDomainEntryBuilder)) {
	x.buf.Reset()
	x.connection_DnsStatsByDomainEntryBuilder.writer = &x.buf
	x.connection_DnsStatsByDomainEntryBuilder.scratch = x.scratch
	cb(&x.connection_DnsStatsByDomainEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x112)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) AddDnsStatsByDomainByQueryType(cb func(w *Connection_DnsStatsByDomainByQueryTypeEntryBuilder)) {
	x.buf.Reset()
	x.connection_DnsStatsByDomainByQueryTypeEntryBuilder.writer = &x.buf
	x.connection_DnsStatsByDomainByQueryTypeEntryBuilder.scratch = x.scratch
	cb(&x.connection_DnsStatsByDomainByQueryTypeEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x152)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) AddDnsStatsByDomainOffsetByQueryType(cb func(w *Connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder)) {
	x.buf.Reset()
	x.connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder.writer = &x.buf
	x.connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder.scratch = x.scratch
	cb(&x.connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x15a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) SetRouteIdx(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x120)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetRouteTargetIdx(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x140)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetHttpAggregations(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) AddTags(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x160)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetTagsIdx(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x168)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetTagsChecksum(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x178)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetStateIndex(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x170)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetDataStreamsAggregations(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x18a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) SetHttp2Aggregations(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x192)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) SetDatabaseAggregations(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x19a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) AddTcpFailuresByErrCode(cb func(w *Connection_TcpFailuresByErrCodeEntryBuilder)) {
	x.buf.Reset()
	x.connection_TcpFailuresByErrCodeEntryBuilder.writer = &x.buf
	x.connection_TcpFailuresByErrCodeEntryBuilder.scratch = x.scratch
	cb(&x.connection_TcpFailuresByErrCodeEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a2)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ConnectionBuilder) SetRemoteEcsTask(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1aa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetLocalContainerTagsIndex(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1b0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionBuilder) SetSystemProbeConn(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1b8)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *ConnectionBuilder) SetResolvConfIdx(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1c0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type Connection_DnsCountByRcodeEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewConnection_DnsCountByRcodeEntryBuilder(writer io.Writer) *Connection_DnsCountByRcodeEntryBuilder {
	return &Connection_DnsCountByRcodeEntryBuilder{
		writer: writer,
	}
}
func (x *Connection_DnsCountByRcodeEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Connection_DnsCountByRcodeEntryBuilder) SetKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *Connection_DnsCountByRcodeEntryBuilder) SetValue(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type Connection_DnsStatsByDomainEntryBuilder struct {
	writer          io.Writer
	buf             bytes.Buffer
	scratch         []byte
	dNSStatsBuilder DNSStatsBuilder
}

func NewConnection_DnsStatsByDomainEntryBuilder(writer io.Writer) *Connection_DnsStatsByDomainEntryBuilder {
	return &Connection_DnsStatsByDomainEntryBuilder{
		writer: writer,
	}
}
func (x *Connection_DnsStatsByDomainEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Connection_DnsStatsByDomainEntryBuilder) SetKey(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *Connection_DnsStatsByDomainEntryBuilder) SetValue(cb func(w *DNSStatsBuilder)) {
	x.buf.Reset()
	x.dNSStatsBuilder.writer = &x.buf
	x.dNSStatsBuilder.scratch = x.scratch
	cb(&x.dNSStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type Connection_DnsStatsByDomainByQueryTypeEntryBuilder struct {
	writer                     io.Writer
	buf                        bytes.Buffer
	scratch                    []byte
	dNSStatsByQueryTypeBuilder DNSStatsByQueryTypeBuilder
}

func NewConnection_DnsStatsByDomainByQueryTypeEntryBuilder(writer io.Writer) *Connection_DnsStatsByDomainByQueryTypeEntryBuilder {
	return &Connection_DnsStatsByDomainByQueryTypeEntryBuilder{
		writer: writer,
	}
}
func (x *Connection_DnsStatsByDomainByQueryTypeEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Connection_DnsStatsByDomainByQueryTypeEntryBuilder) SetKey(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *Connection_DnsStatsByDomainByQueryTypeEntryBuilder) SetValue(cb func(w *DNSStatsByQueryTypeBuilder)) {
	x.buf.Reset()
	x.dNSStatsByQueryTypeBuilder.writer = &x.buf
	x.dNSStatsByQueryTypeBuilder.scratch = x.scratch
	cb(&x.dNSStatsByQueryTypeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type Connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder struct {
	writer                     io.Writer
	buf                        bytes.Buffer
	scratch                    []byte
	dNSStatsByQueryTypeBuilder DNSStatsByQueryTypeBuilder
}

func NewConnection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder(writer io.Writer) *Connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder {
	return &Connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder{
		writer: writer,
	}
}
func (x *Connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder) SetKey(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *Connection_DnsStatsByDomainOffsetByQueryTypeEntryBuilder) SetValue(cb func(w *DNSStatsByQueryTypeBuilder)) {
	x.buf.Reset()
	x.dNSStatsByQueryTypeBuilder.writer = &x.buf
	x.dNSStatsByQueryTypeBuilder.scratch = x.scratch
	cb(&x.dNSStatsByQueryTypeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type Connection_TcpFailuresByErrCodeEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewConnection_TcpFailuresByErrCodeEntryBuilder(writer io.Writer) *Connection_TcpFailuresByErrCodeEntryBuilder {
	return &Connection_TcpFailuresByErrCodeEntryBuilder{
		writer: writer,
	}
}
func (x *Connection_TcpFailuresByErrCodeEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Connection_TcpFailuresByErrCodeEntryBuilder) SetKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *Connection_TcpFailuresByErrCodeEntryBuilder) SetValue(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ResourceMetadataBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResourceMetadataBuilder(writer io.Writer) *ResourceMetadataBuilder {
	return &ResourceMetadataBuilder{
		writer: writer,
	}
}
func (x *ResourceMetadataBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResourceMetadataBuilder) SetId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceMetadataBuilder) SetByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ResourceMetadataBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceMetadataBuilder) SetTagIndex(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ResourceMetadataBuilder) SetTagsModified(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ConnectionsTelemetryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewConnectionsTelemetryBuilder(writer io.Writer) *ConnectionsTelemetryBuilder {
	return &ConnectionsTelemetryBuilder{
		writer: writer,
	}
}
func (x *ConnectionsTelemetryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ConnectionsTelemetryBuilder) SetMonotonicKprobesTriggered(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionsTelemetryBuilder) SetMonotonicKprobesMissed(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionsTelemetryBuilder) SetMonotonicConntrackRegisters(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionsTelemetryBuilder) SetMonotonicConntrackRegistersDropped(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionsTelemetryBuilder) SetMonotonicDnsPacketsProcessed(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionsTelemetryBuilder) SetMonotonicConnsClosed(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionsTelemetryBuilder) SetConnsBpfMapSize(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionsTelemetryBuilder) SetMonotonicUdpSendsProcessed(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionsTelemetryBuilder) SetMonotonicUdpSendsMissed(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionsTelemetryBuilder) SetConntrackSamplingPercent(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x50)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConnectionsTelemetryBuilder) SetDnsStatsDropped(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x58)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type CollectorConnectionsTelemetryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCollectorConnectionsTelemetryBuilder(writer io.Writer) *CollectorConnectionsTelemetryBuilder {
	return &CollectorConnectionsTelemetryBuilder{
		writer: writer,
	}
}
func (x *CollectorConnectionsTelemetryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorConnectionsTelemetryBuilder) SetKprobesTriggered(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsTelemetryBuilder) SetKprobesMissed(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsTelemetryBuilder) SetConntrackRegisters(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsTelemetryBuilder) SetConntrackRegistersDropped(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsTelemetryBuilder) SetDnsPacketsProcessed(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsTelemetryBuilder) SetConnsClosed(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsTelemetryBuilder) SetConnsBpfMapSize(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsTelemetryBuilder) SetUdpSendsProcessed(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsTelemetryBuilder) SetUdpSendsMissed(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsTelemetryBuilder) SetConntrackSamplingPercent(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x50)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorConnectionsTelemetryBuilder) SetDnsStatsDropped(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x58)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type RuntimeCompilationTelemetryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewRuntimeCompilationTelemetryBuilder(writer io.Writer) *RuntimeCompilationTelemetryBuilder {
	return &RuntimeCompilationTelemetryBuilder{
		writer: writer,
	}
}
func (x *RuntimeCompilationTelemetryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *RuntimeCompilationTelemetryBuilder) SetRuntimeCompilationEnabled(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *RuntimeCompilationTelemetryBuilder) SetRuntimeCompilationResult(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x10)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *RuntimeCompilationTelemetryBuilder) SetRuntimeCompilationDuration(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *RuntimeCompilationTelemetryBuilder) SetKernelHeaderFetchResult(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type AgentConfigurationBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewAgentConfigurationBuilder(writer io.Writer) *AgentConfigurationBuilder {
	return &AgentConfigurationBuilder{
		writer: writer,
	}
}
func (x *AgentConfigurationBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *AgentConfigurationBuilder) SetNpmEnabled(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *AgentConfigurationBuilder) SetUsmEnabled(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x10)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *AgentConfigurationBuilder) SetDsmEnabled(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x18)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *AgentConfigurationBuilder) SetCcmEnabled(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *AgentConfigurationBuilder) SetCsmEnabled(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x28)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}

type RouteBuilder struct {
	writer           io.Writer
	buf              bytes.Buffer
	scratch          []byte
	subnetBuilder    SubnetBuilder
	interfaceBuilder InterfaceBuilder
}

func NewRouteBuilder(writer io.Writer) *RouteBuilder {
	return &RouteBuilder{
		writer: writer,
	}
}
func (x *RouteBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *RouteBuilder) SetSubnet(cb func(w *SubnetBuilder)) {
	x.buf.Reset()
	x.subnetBuilder.writer = &x.buf
	x.subnetBuilder.scratch = x.scratch
	cb(&x.subnetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RouteBuilder) SetInterface(cb func(w *InterfaceBuilder)) {
	x.buf.Reset()
	x.interfaceBuilder.writer = &x.buf
	x.interfaceBuilder.scratch = x.scratch
	cb(&x.interfaceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type InterfaceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewInterfaceBuilder(writer io.Writer) *InterfaceBuilder {
	return &InterfaceBuilder{
		writer: writer,
	}
}
func (x *InterfaceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *InterfaceBuilder) SetHardwareAddr(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type SubnetBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewSubnetBuilder(writer io.Writer) *SubnetBuilder {
	return &SubnetBuilder{
		writer: writer,
	}
}
func (x *SubnetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *SubnetBuilder) SetAlias(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type RouteMetadataBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewRouteMetadataBuilder(writer io.Writer) *RouteMetadataBuilder {
	return &RouteMetadataBuilder{
		writer: writer,
	}
}
func (x *RouteMetadataBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *RouteMetadataBuilder) SetAlias(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *RouteMetadataBuilder) SetTagIndex(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *RouteMetadataBuilder) SetTagsModified(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *RouteMetadataBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type IPTranslationBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewIPTranslationBuilder(writer io.Writer) *IPTranslationBuilder {
	return &IPTranslationBuilder{
		writer: writer,
	}
}
func (x *IPTranslationBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *IPTranslationBuilder) SetReplSrcIP(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *IPTranslationBuilder) SetReplDstIP(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *IPTranslationBuilder) SetReplSrcPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *IPTranslationBuilder) SetReplDstPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type AddrBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewAddrBuilder(writer io.Writer) *AddrBuilder {
	return &AddrBuilder{
		writer: writer,
	}
}
func (x *AddrBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *AddrBuilder) SetIp(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AddrBuilder) SetPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *AddrBuilder) SetContainerId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AddrBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ProtocolStackBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewProtocolStackBuilder(writer io.Writer) *ProtocolStackBuilder {
	return &ProtocolStackBuilder{
		writer: writer,
	}
}
func (x *ProtocolStackBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProtocolStackBuilder) AddStack(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type DNSEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewDNSEntryBuilder(writer io.Writer) *DNSEntryBuilder {
	return &DNSEntryBuilder{
		writer: writer,
	}
}
func (x *DNSEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DNSEntryBuilder) AddNames(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type DNSStatsBuilder struct {
	writer                               io.Writer
	buf                                  bytes.Buffer
	scratch                              []byte
	dNSStats_DnsCountByRcodeEntryBuilder DNSStats_DnsCountByRcodeEntryBuilder
}

func NewDNSStatsBuilder(writer io.Writer) *DNSStatsBuilder {
	return &DNSStatsBuilder{
		writer: writer,
	}
}
func (x *DNSStatsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DNSStatsBuilder) SetDnsTimeouts(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DNSStatsBuilder) SetDnsSuccessLatencySum(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DNSStatsBuilder) SetDnsFailureLatencySum(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DNSStatsBuilder) AddDnsCountByRcode(cb func(w *DNSStats_DnsCountByRcodeEntryBuilder)) {
	x.buf.Reset()
	x.dNSStats_DnsCountByRcodeEntryBuilder.writer = &x.buf
	x.dNSStats_DnsCountByRcodeEntryBuilder.scratch = x.scratch
	cb(&x.dNSStats_DnsCountByRcodeEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type DNSStats_DnsCountByRcodeEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewDNSStats_DnsCountByRcodeEntryBuilder(writer io.Writer) *DNSStats_DnsCountByRcodeEntryBuilder {
	return &DNSStats_DnsCountByRcodeEntryBuilder{
		writer: writer,
	}
}
func (x *DNSStats_DnsCountByRcodeEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DNSStats_DnsCountByRcodeEntryBuilder) SetKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DNSStats_DnsCountByRcodeEntryBuilder) SetValue(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type DNSStatsByQueryTypeBuilder struct {
	writer                                              io.Writer
	buf                                                 bytes.Buffer
	scratch                                             []byte
	dNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder DNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder
}

func NewDNSStatsByQueryTypeBuilder(writer io.Writer) *DNSStatsByQueryTypeBuilder {
	return &DNSStatsByQueryTypeBuilder{
		writer: writer,
	}
}
func (x *DNSStatsByQueryTypeBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DNSStatsByQueryTypeBuilder) AddDnsStatsByQueryType(cb func(w *DNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder)) {
	x.buf.Reset()
	x.dNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder.writer = &x.buf
	x.dNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder.scratch = x.scratch
	cb(&x.dNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type DNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder struct {
	writer          io.Writer
	buf             bytes.Buffer
	scratch         []byte
	dNSStatsBuilder DNSStatsBuilder
}

func NewDNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder(writer io.Writer) *DNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder {
	return &DNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder{
		writer: writer,
	}
}
func (x *DNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder) SetKey(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DNSStatsByQueryType_DnsStatsByQueryTypeEntryBuilder) SetValue(cb func(w *DNSStatsBuilder)) {
	x.buf.Reset()
	x.dNSStatsBuilder.writer = &x.buf
	x.dNSStatsBuilder.scratch = x.scratch
	cb(&x.dNSStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PublicIpMetadataBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPublicIpMetadataBuilder(writer io.Writer) *PublicIpMetadataBuilder {
	return &PublicIpMetadataBuilder{
		writer: writer,
	}
}
func (x *PublicIpMetadataBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PublicIpMetadataBuilder) SetIp(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PublicIpMetadataBuilder) SetCloudProvider(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PublicIpMetadataBuilder) SetRegion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PublicIpMetadataBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
