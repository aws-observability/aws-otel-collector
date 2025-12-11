// THIS IS A GENERATED FILE
// DO NOT EDIT
package process

import (
	bytes "bytes"
	protowire "google.golang.org/protobuf/encoding/protowire"
	io "io"
	math "math"
)

type ResCollectorBuilder struct {
	writer                     io.Writer
	buf                        bytes.Buffer
	scratch                    []byte
	resCollector_HeaderBuilder ResCollector_HeaderBuilder
	collectorStatusBuilder     CollectorStatusBuilder
}

func NewResCollectorBuilder(writer io.Writer) *ResCollectorBuilder {
	return &ResCollectorBuilder{
		writer: writer,
	}
}
func (x *ResCollectorBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResCollectorBuilder) SetHeader(cb func(w *ResCollector_HeaderBuilder)) {
	x.buf.Reset()
	x.resCollector_HeaderBuilder.writer = &x.buf
	x.resCollector_HeaderBuilder.scratch = x.scratch
	cb(&x.resCollector_HeaderBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ResCollectorBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResCollectorBuilder) SetStatus(cb func(w *CollectorStatusBuilder)) {
	x.buf.Reset()
	x.collectorStatusBuilder.writer = &x.buf
	x.collectorStatusBuilder.scratch = x.scratch
	cb(&x.collectorStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ResCollector_HeaderBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResCollector_HeaderBuilder(writer io.Writer) *ResCollector_HeaderBuilder {
	return &ResCollector_HeaderBuilder{
		writer: writer,
	}
}
func (x *ResCollector_HeaderBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResCollector_HeaderBuilder) SetType(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type CollectorProcBuilder struct {
	writer            io.Writer
	buf               bytes.Buffer
	scratch           []byte
	processBuilder    ProcessBuilder
	hostBuilder       HostBuilder
	systemInfoBuilder SystemInfoBuilder
	containerBuilder  ContainerBuilder
}

func NewCollectorProcBuilder(writer io.Writer) *CollectorProcBuilder {
	return &CollectorProcBuilder{
		writer: writer,
	}
}
func (x *CollectorProcBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorProcBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorProcBuilder) SetNetworkId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x5a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorProcBuilder) AddProcesses(cb func(w *ProcessBuilder)) {
	x.buf.Reset()
	x.processBuilder.writer = &x.buf
	x.processBuilder.scratch = x.scratch
	cb(&x.processBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcBuilder) SetInfo(cb func(w *SystemInfoBuilder)) {
	x.buf.Reset()
	x.systemInfoBuilder.writer = &x.buf
	x.systemInfoBuilder.scratch = x.scratch
	cb(&x.systemInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorProcBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorProcBuilder) AddContainers(cb func(w *ContainerBuilder)) {
	x.buf.Reset()
	x.containerBuilder.writer = &x.buf
	x.containerBuilder.scratch = x.scratch
	cb(&x.containerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcBuilder) SetContainerHostType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x60)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *CollectorProcBuilder) SetHintMask(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x70)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type CollectorProcDiscoveryBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	processDiscoveryBuilder ProcessDiscoveryBuilder
	hostBuilder             HostBuilder
}

func NewCollectorProcDiscoveryBuilder(writer io.Writer) *CollectorProcDiscoveryBuilder {
	return &CollectorProcDiscoveryBuilder{
		writer: writer,
	}
}
func (x *CollectorProcDiscoveryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorProcDiscoveryBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorProcDiscoveryBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorProcDiscoveryBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorProcDiscoveryBuilder) AddProcessDiscoveries(cb func(w *ProcessDiscoveryBuilder)) {
	x.buf.Reset()
	x.processDiscoveryBuilder.writer = &x.buf
	x.processDiscoveryBuilder.scratch = x.scratch
	cb(&x.processDiscoveryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcDiscoveryBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorRealTimeBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	processStatBuilder   ProcessStatBuilder
	containerStatBuilder ContainerStatBuilder
}

func NewCollectorRealTimeBuilder(writer io.Writer) *CollectorRealTimeBuilder {
	return &CollectorRealTimeBuilder{
		writer: writer,
	}
}
func (x *CollectorRealTimeBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorRealTimeBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) AddStats(cb func(w *ProcessStatBuilder)) {
	x.buf.Reset()
	x.processStatBuilder.writer = &x.buf
	x.processStatBuilder.scratch = x.scratch
	cb(&x.processStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorRealTimeBuilder) SetHostId(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) SetOrgId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) SetNumCpus(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) SetTotalMemory(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) AddContainerStats(cb func(w *ContainerStatBuilder)) {
	x.buf.Reset()
	x.containerStatBuilder.writer = &x.buf
	x.containerStatBuilder.scratch = x.scratch
	cb(&x.containerStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorRealTimeBuilder) SetContainerHostType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x58)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type CollectorContainerBuilder struct {
	writer            io.Writer
	buf               bytes.Buffer
	scratch           []byte
	systemInfoBuilder SystemInfoBuilder
	containerBuilder  ContainerBuilder
	hostBuilder       HostBuilder
}

func NewCollectorContainerBuilder(writer io.Writer) *CollectorContainerBuilder {
	return &CollectorContainerBuilder{
		writer: writer,
	}
}
func (x *CollectorContainerBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorContainerBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerBuilder) SetNetworkId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x5a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerBuilder) SetInfo(cb func(w *SystemInfoBuilder)) {
	x.buf.Reset()
	x.systemInfoBuilder.writer = &x.buf
	x.systemInfoBuilder.scratch = x.scratch
	cb(&x.systemInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorContainerBuilder) AddContainers(cb func(w *ContainerBuilder)) {
	x.buf.Reset()
	x.containerBuilder.writer = &x.buf
	x.containerBuilder.scratch = x.scratch
	cb(&x.containerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorContainerBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorContainerBuilder) SetContainerHostType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x48)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type CollectorContainerRealTimeBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	containerStatBuilder ContainerStatBuilder
}

func NewCollectorContainerRealTimeBuilder(writer io.Writer) *CollectorContainerRealTimeBuilder {
	return &CollectorContainerRealTimeBuilder{
		writer: writer,
	}
}
func (x *CollectorContainerRealTimeBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorContainerRealTimeBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) AddStats(cb func(w *ContainerStatBuilder)) {
	x.buf.Reset()
	x.containerStatBuilder.writer = &x.buf
	x.containerStatBuilder.scratch = x.scratch
	cb(&x.containerStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorContainerRealTimeBuilder) SetNumCpus(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) SetTotalMemory(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) SetHostId(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) SetContainerHostType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x40)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type CollectorReqStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCollectorReqStatusBuilder(writer io.Writer) *CollectorReqStatusBuilder {
	return &CollectorReqStatusBuilder{
		writer: writer,
	}
}
func (x *CollectorReqStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorReqStatusBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorECSTaskBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	eCSTaskBuilder      ECSTaskBuilder
	hostBuilder         HostBuilder
	systemInfoBuilder   SystemInfoBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorECSTaskBuilder(writer io.Writer) *CollectorECSTaskBuilder {
	return &CollectorECSTaskBuilder{
		writer: writer,
	}
}
func (x *CollectorECSTaskBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorECSTaskBuilder) SetAwsAccountID(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorECSTaskBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorECSTaskBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorECSTaskBuilder) SetRegion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorECSTaskBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorECSTaskBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorECSTaskBuilder) AddTasks(cb func(w *ECSTaskBuilder)) {
	x.buf.Reset()
	x.eCSTaskBuilder.writer = &x.buf
	x.eCSTaskBuilder.scratch = x.scratch
	cb(&x.eCSTaskBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorECSTaskBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorECSTaskBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x4a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorECSTaskBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorECSTaskBuilder) SetInfo(cb func(w *SystemInfoBuilder)) {
	x.buf.Reset()
	x.systemInfoBuilder.writer = &x.buf
	x.systemInfoBuilder.scratch = x.scratch
	cb(&x.systemInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorECSTaskBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x62)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ECSTaskBuilder struct {
	writer                                      io.Writer
	buf                                         bytes.Buffer
	scratch                                     []byte
	eCSTask_LimitsEntryBuilder                  ECSTask_LimitsEntryBuilder
	eCSTask_EphemeralStorageMetricsEntryBuilder ECSTask_EphemeralStorageMetricsEntryBuilder
	eCSContainerBuilder                         ECSContainerBuilder
	hostBuilder                                 HostBuilder
}

func NewECSTaskBuilder(writer io.Writer) *ECSTaskBuilder {
	return &ECSTaskBuilder{
		writer: writer,
	}
}
func (x *ECSTaskBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ECSTaskBuilder) SetArn(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetResourceVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetLaunchType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetDesiredStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetKnownStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetFamily(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetAvailabilityZone(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) AddLimits(cb func(w *ECSTask_LimitsEntryBuilder)) {
	x.buf.Reset()
	x.eCSTask_LimitsEntryBuilder.writer = &x.buf
	x.eCSTask_LimitsEntryBuilder.scratch = x.scratch
	cb(&x.eCSTask_LimitsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x4a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSTaskBuilder) AddEphemeralStorageMetrics(cb func(w *ECSTask_EphemeralStorageMetricsEntryBuilder)) {
	x.buf.Reset()
	x.eCSTask_EphemeralStorageMetricsEntryBuilder.writer = &x.buf
	x.eCSTask_EphemeralStorageMetricsEntryBuilder.scratch = x.scratch
	cb(&x.eCSTask_EphemeralStorageMetricsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSTaskBuilder) SetServiceName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x5a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetVpcId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x62)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetPullStartedAt(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x68)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetPullStoppedAt(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x70)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetExecutionStoppedAt(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x78)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) AddContainers(cb func(w *ECSContainerBuilder)) {
	x.buf.Reset()
	x.eCSContainerBuilder.writer = &x.buf
	x.eCSContainerBuilder.scratch = x.scratch
	cb(&x.eCSContainerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x82)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSTaskBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) AddEcsTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x92)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) AddContainerInstanceTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x9a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTaskBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa2)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSTaskBuilder) SetContainerInstanceArn(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xaa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ECSTask_LimitsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewECSTask_LimitsEntryBuilder(writer io.Writer) *ECSTask_LimitsEntryBuilder {
	return &ECSTask_LimitsEntryBuilder{
		writer: writer,
	}
}
func (x *ECSTask_LimitsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ECSTask_LimitsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTask_LimitsEntryBuilder) SetValue(v float64) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed64(x.scratch, math.Float64bits(v))
	x.writer.Write(x.scratch)
}

type ECSTask_EphemeralStorageMetricsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewECSTask_EphemeralStorageMetricsEntryBuilder(writer io.Writer) *ECSTask_EphemeralStorageMetricsEntryBuilder {
	return &ECSTask_EphemeralStorageMetricsEntryBuilder{
		writer: writer,
	}
}
func (x *ECSTask_EphemeralStorageMetricsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ECSTask_EphemeralStorageMetricsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSTask_EphemeralStorageMetricsEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ECSContainerBuilder struct {
	writer                              io.Writer
	buf                                 bytes.Buffer
	scratch                             []byte
	eCSContainerExitCodeBuilder         ECSContainerExitCodeBuilder
	eCSContainer_LogOptionsEntryBuilder ECSContainer_LogOptionsEntryBuilder
	eCSContainerPortBuilder             ECSContainerPortBuilder
	eCSContainerNetworkBuilder          ECSContainerNetworkBuilder
	eCSContainerVolumeBuilder           ECSContainerVolumeBuilder
	eCSContainerHealthBuilder           ECSContainerHealthBuilder
	eCSContainer_LimitsEntryBuilder     ECSContainer_LimitsEntryBuilder
}

func NewECSContainerBuilder(writer io.Writer) *ECSContainerBuilder {
	return &ECSContainerBuilder{
		writer: writer,
	}
}
func (x *ECSContainerBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ECSContainerBuilder) SetDockerID(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) SetDockerName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) SetImage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) SetImageID(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) SetCreatedAt(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) SetStartedAt(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) SetFinishedAt(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) SetDesiredStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) SetKnownStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) SetExitCode(cb func(w *ECSContainerExitCodeBuilder)) {
	x.buf.Reset()
	x.eCSContainerExitCodeBuilder.writer = &x.buf
	x.eCSContainerExitCodeBuilder.scratch = x.scratch
	cb(&x.eCSContainerExitCodeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSContainerBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x62)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) SetLogDriver(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x6a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) AddLogOptions(cb func(w *ECSContainer_LogOptionsEntryBuilder)) {
	x.buf.Reset()
	x.eCSContainer_LogOptionsEntryBuilder.writer = &x.buf
	x.eCSContainer_LogOptionsEntryBuilder.scratch = x.scratch
	cb(&x.eCSContainer_LogOptionsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x72)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSContainerBuilder) SetContainerArn(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x7a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) AddPorts(cb func(w *ECSContainerPortBuilder)) {
	x.buf.Reset()
	x.eCSContainerPortBuilder.writer = &x.buf
	x.eCSContainerPortBuilder.scratch = x.scratch
	cb(&x.eCSContainerPortBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x82)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSContainerBuilder) AddNetworks(cb func(w *ECSContainerNetworkBuilder)) {
	x.buf.Reset()
	x.eCSContainerNetworkBuilder.writer = &x.buf
	x.eCSContainerNetworkBuilder.scratch = x.scratch
	cb(&x.eCSContainerNetworkBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSContainerBuilder) AddVolumes(cb func(w *ECSContainerVolumeBuilder)) {
	x.buf.Reset()
	x.eCSContainerVolumeBuilder.writer = &x.buf
	x.eCSContainerVolumeBuilder.scratch = x.scratch
	cb(&x.eCSContainerVolumeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x92)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSContainerBuilder) SetHealth(cb func(w *ECSContainerHealthBuilder)) {
	x.buf.Reset()
	x.eCSContainerHealthBuilder.writer = &x.buf
	x.eCSContainerHealthBuilder.scratch = x.scratch
	cb(&x.eCSContainerHealthBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x9a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSContainerBuilder) AddLabels(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa2)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerBuilder) AddLimits(cb func(w *ECSContainer_LimitsEntryBuilder)) {
	x.buf.Reset()
	x.eCSContainer_LimitsEntryBuilder.writer = &x.buf
	x.eCSContainer_LimitsEntryBuilder.scratch = x.scratch
	cb(&x.eCSContainer_LimitsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xaa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSContainerBuilder) SetSnapshotter(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xb2)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ECSContainer_LogOptionsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewECSContainer_LogOptionsEntryBuilder(writer io.Writer) *ECSContainer_LogOptionsEntryBuilder {
	return &ECSContainer_LogOptionsEntryBuilder{
		writer: writer,
	}
}
func (x *ECSContainer_LogOptionsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ECSContainer_LogOptionsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainer_LogOptionsEntryBuilder) SetValue(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ECSContainer_LimitsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewECSContainer_LimitsEntryBuilder(writer io.Writer) *ECSContainer_LimitsEntryBuilder {
	return &ECSContainer_LimitsEntryBuilder{
		writer: writer,
	}
}
func (x *ECSContainer_LimitsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ECSContainer_LimitsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainer_LimitsEntryBuilder) SetValue(v float64) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed64(x.scratch, math.Float64bits(v))
	x.writer.Write(x.scratch)
}

type ECSContainerNetworkBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewECSContainerNetworkBuilder(writer io.Writer) *ECSContainerNetworkBuilder {
	return &ECSContainerNetworkBuilder{
		writer: writer,
	}
}
func (x *ECSContainerNetworkBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ECSContainerNetworkBuilder) SetNetworkMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerNetworkBuilder) AddIpv4Addresses(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerNetworkBuilder) AddIpv6Addresses(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ECSContainerPortBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewECSContainerPortBuilder(writer io.Writer) *ECSContainerPortBuilder {
	return &ECSContainerPortBuilder{
		writer: writer,
	}
}
func (x *ECSContainerPortBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ECSContainerPortBuilder) SetContainerPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ECSContainerPortBuilder) SetProtocol(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerPortBuilder) SetHostIp(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerPortBuilder) SetHostPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ECSContainerVolumeBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewECSContainerVolumeBuilder(writer io.Writer) *ECSContainerVolumeBuilder {
	return &ECSContainerVolumeBuilder{
		writer: writer,
	}
}
func (x *ECSContainerVolumeBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ECSContainerVolumeBuilder) SetDockerName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerVolumeBuilder) SetSource(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerVolumeBuilder) SetDestination(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ECSContainerHealthBuilder struct {
	writer                      io.Writer
	buf                         bytes.Buffer
	scratch                     []byte
	eCSContainerExitCodeBuilder ECSContainerExitCodeBuilder
}

func NewECSContainerHealthBuilder(writer io.Writer) *ECSContainerHealthBuilder {
	return &ECSContainerHealthBuilder{
		writer: writer,
	}
}
func (x *ECSContainerHealthBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ECSContainerHealthBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ECSContainerHealthBuilder) SetSince(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ECSContainerHealthBuilder) SetExitCode(cb func(w *ECSContainerExitCodeBuilder)) {
	x.buf.Reset()
	x.eCSContainerExitCodeBuilder.writer = &x.buf
	x.eCSContainerExitCodeBuilder.scratch = x.scratch
	cb(&x.eCSContainerExitCodeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ECSContainerHealthBuilder) SetOutput(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ECSContainerExitCodeBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewECSContainerExitCodeBuilder(writer io.Writer) *ECSContainerExitCodeBuilder {
	return &ECSContainerExitCodeBuilder{
		writer: writer,
	}
}
func (x *ECSContainerExitCodeBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ECSContainerExitCodeBuilder) SetExitCode(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type CollectorPodBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	podBuilder          PodBuilder
	hostBuilder         HostBuilder
	systemInfoBuilder   SystemInfoBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorPodBuilder(writer io.Writer) *CollectorPodBuilder {
	return &CollectorPodBuilder{
		writer: writer,
	}
}
func (x *CollectorPodBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorPodBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPodBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPodBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPodBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPodBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPodBuilder) AddPods(cb func(w *PodBuilder)) {
	x.buf.Reset()
	x.podBuilder.writer = &x.buf
	x.podBuilder.scratch = x.scratch
	cb(&x.podBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorPodBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorPodBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPodBuilder) SetInfo(cb func(w *SystemInfoBuilder)) {
	x.buf.Reset()
	x.systemInfoBuilder.writer = &x.buf
	x.systemInfoBuilder.scratch = x.scratch
	cb(&x.systemInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x4a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorPodBuilder) SetIsTerminated(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x50)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *CollectorPodBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorPodDisruptionBudgetBuilder struct {
	writer                     io.Writer
	buf                        bytes.Buffer
	scratch                    []byte
	podDisruptionBudgetBuilder PodDisruptionBudgetBuilder
	agentVersionBuilder        AgentVersionBuilder
}

func NewCollectorPodDisruptionBudgetBuilder(writer io.Writer) *CollectorPodDisruptionBudgetBuilder {
	return &CollectorPodDisruptionBudgetBuilder{
		writer: writer,
	}
}
func (x *CollectorPodDisruptionBudgetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorPodDisruptionBudgetBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPodDisruptionBudgetBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPodDisruptionBudgetBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPodDisruptionBudgetBuilder) AddPodDisruptionBudgets(cb func(w *PodDisruptionBudgetBuilder)) {
	x.buf.Reset()
	x.podDisruptionBudgetBuilder.writer = &x.buf
	x.podDisruptionBudgetBuilder.scratch = x.scratch
	cb(&x.podDisruptionBudgetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorPodDisruptionBudgetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPodDisruptionBudgetBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPodDisruptionBudgetBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorReplicaSetBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	replicaSetBuilder   ReplicaSetBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorReplicaSetBuilder(writer io.Writer) *CollectorReplicaSetBuilder {
	return &CollectorReplicaSetBuilder{
		writer: writer,
	}
}
func (x *CollectorReplicaSetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorReplicaSetBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorReplicaSetBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorReplicaSetBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorReplicaSetBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorReplicaSetBuilder) AddReplicaSets(cb func(w *ReplicaSetBuilder)) {
	x.buf.Reset()
	x.replicaSetBuilder.writer = &x.buf
	x.replicaSetBuilder.scratch = x.scratch
	cb(&x.replicaSetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorReplicaSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorReplicaSetBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorDeploymentBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	deploymentBuilder   DeploymentBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorDeploymentBuilder(writer io.Writer) *CollectorDeploymentBuilder {
	return &CollectorDeploymentBuilder{
		writer: writer,
	}
}
func (x *CollectorDeploymentBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorDeploymentBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorDeploymentBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorDeploymentBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorDeploymentBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorDeploymentBuilder) AddDeployments(cb func(w *DeploymentBuilder)) {
	x.buf.Reset()
	x.deploymentBuilder.writer = &x.buf
	x.deploymentBuilder.scratch = x.scratch
	cb(&x.deploymentBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorDeploymentBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorDeploymentBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorServiceBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	serviceBuilder      ServiceBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorServiceBuilder(writer io.Writer) *CollectorServiceBuilder {
	return &CollectorServiceBuilder{
		writer: writer,
	}
}
func (x *CollectorServiceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorServiceBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceBuilder) AddServices(cb func(w *ServiceBuilder)) {
	x.buf.Reset()
	x.serviceBuilder.writer = &x.buf
	x.serviceBuilder.scratch = x.scratch
	cb(&x.serviceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorServiceBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorNodeBuilder struct {
	writer                                     io.Writer
	buf                                        bytes.Buffer
	scratch                                    []byte
	nodeBuilder                                NodeBuilder
	collectorNode_HostAliasMappingEntryBuilder CollectorNode_HostAliasMappingEntryBuilder
	agentVersionBuilder                        AgentVersionBuilder
}

func NewCollectorNodeBuilder(writer io.Writer) *CollectorNodeBuilder {
	return &CollectorNodeBuilder{
		writer: writer,
	}
}
func (x *CollectorNodeBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorNodeBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNodeBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNodeBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorNodeBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorNodeBuilder) AddNodes(cb func(w *NodeBuilder)) {
	x.buf.Reset()
	x.nodeBuilder.writer = &x.buf
	x.nodeBuilder.scratch = x.scratch
	cb(&x.nodeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorNodeBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNodeBuilder) AddHostAliasMapping(cb func(w *CollectorNode_HostAliasMappingEntryBuilder)) {
	x.buf.Reset()
	x.collectorNode_HostAliasMappingEntryBuilder.writer = &x.buf
	x.collectorNode_HostAliasMappingEntryBuilder.scratch = x.scratch
	cb(&x.collectorNode_HostAliasMappingEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorNodeBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorNode_HostAliasMappingEntryBuilder struct {
	writer      io.Writer
	buf         bytes.Buffer
	scratch     []byte
	hostBuilder HostBuilder
}

func NewCollectorNode_HostAliasMappingEntryBuilder(writer io.Writer) *CollectorNode_HostAliasMappingEntryBuilder {
	return &CollectorNode_HostAliasMappingEntryBuilder{
		writer: writer,
	}
}
func (x *CollectorNode_HostAliasMappingEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorNode_HostAliasMappingEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNode_HostAliasMappingEntryBuilder) SetValue(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorClusterBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	clusterBuilder      ClusterBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorClusterBuilder(writer io.Writer) *CollectorClusterBuilder {
	return &CollectorClusterBuilder{
		writer: writer,
	}
}
func (x *CollectorClusterBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorClusterBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterBuilder) SetCluster(cb func(w *ClusterBuilder)) {
	x.buf.Reset()
	x.clusterBuilder.writer = &x.buf
	x.clusterBuilder.scratch = x.scratch
	cb(&x.clusterBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorClusterBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorManifestBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	manifestBuilder     ManifestBuilder
	agentVersionBuilder AgentVersionBuilder
	systemInfoBuilder   SystemInfoBuilder
}

func NewCollectorManifestBuilder(writer io.Writer) *CollectorManifestBuilder {
	return &CollectorManifestBuilder{
		writer: writer,
	}
}
func (x *CollectorManifestBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorManifestBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorManifestBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorManifestBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorManifestBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorManifestBuilder) AddManifests(cb func(w *ManifestBuilder)) {
	x.buf.Reset()
	x.manifestBuilder.writer = &x.buf
	x.manifestBuilder.scratch = x.scratch
	cb(&x.manifestBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorManifestBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorManifestBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorManifestBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorManifestBuilder) SetOriginCollector(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x48)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *CollectorManifestBuilder) SetSystemInfo(cb func(w *SystemInfoBuilder)) {
	x.buf.Reset()
	x.systemInfoBuilder.writer = &x.buf
	x.systemInfoBuilder.scratch = x.scratch
	cb(&x.systemInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorManifestCRDBuilder struct {
	writer                   io.Writer
	buf                      bytes.Buffer
	scratch                  []byte
	collectorManifestBuilder CollectorManifestBuilder
}

func NewCollectorManifestCRDBuilder(writer io.Writer) *CollectorManifestCRDBuilder {
	return &CollectorManifestCRDBuilder{
		writer: writer,
	}
}
func (x *CollectorManifestCRDBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorManifestCRDBuilder) SetManifest(cb func(w *CollectorManifestBuilder)) {
	x.buf.Reset()
	x.collectorManifestBuilder.writer = &x.buf
	x.collectorManifestBuilder.scratch = x.scratch
	cb(&x.collectorManifestBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorManifestCRDBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorManifestCRBuilder struct {
	writer                   io.Writer
	buf                      bytes.Buffer
	scratch                  []byte
	collectorManifestBuilder CollectorManifestBuilder
}

func NewCollectorManifestCRBuilder(writer io.Writer) *CollectorManifestCRBuilder {
	return &CollectorManifestCRBuilder{
		writer: writer,
	}
}
func (x *CollectorManifestCRBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorManifestCRBuilder) SetManifest(cb func(w *CollectorManifestBuilder)) {
	x.buf.Reset()
	x.collectorManifestBuilder.writer = &x.buf
	x.collectorManifestBuilder.scratch = x.scratch
	cb(&x.collectorManifestBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorManifestCRBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorNamespaceBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	namespaceBuilder    NamespaceBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorNamespaceBuilder(writer io.Writer) *CollectorNamespaceBuilder {
	return &CollectorNamespaceBuilder{
		writer: writer,
	}
}
func (x *CollectorNamespaceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorNamespaceBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNamespaceBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNamespaceBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorNamespaceBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorNamespaceBuilder) AddNamespaces(cb func(w *NamespaceBuilder)) {
	x.buf.Reset()
	x.namespaceBuilder.writer = &x.buf
	x.namespaceBuilder.scratch = x.scratch
	cb(&x.namespaceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorNamespaceBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNamespaceBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorJobBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	jobBuilder          JobBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorJobBuilder(writer io.Writer) *CollectorJobBuilder {
	return &CollectorJobBuilder{
		writer: writer,
	}
}
func (x *CollectorJobBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorJobBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorJobBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorJobBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorJobBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorJobBuilder) AddJobs(cb func(w *JobBuilder)) {
	x.buf.Reset()
	x.jobBuilder.writer = &x.buf
	x.jobBuilder.scratch = x.scratch
	cb(&x.jobBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorJobBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorJobBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorCronJobBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	cronJobBuilder      CronJobBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorCronJobBuilder(writer io.Writer) *CollectorCronJobBuilder {
	return &CollectorCronJobBuilder{
		writer: writer,
	}
}
func (x *CollectorCronJobBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorCronJobBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorCronJobBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorCronJobBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorCronJobBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorCronJobBuilder) AddCronJobs(cb func(w *CronJobBuilder)) {
	x.buf.Reset()
	x.cronJobBuilder.writer = &x.buf
	x.cronJobBuilder.scratch = x.scratch
	cb(&x.cronJobBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorCronJobBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorCronJobBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorDaemonSetBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	daemonSetBuilder    DaemonSetBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorDaemonSetBuilder(writer io.Writer) *CollectorDaemonSetBuilder {
	return &CollectorDaemonSetBuilder{
		writer: writer,
	}
}
func (x *CollectorDaemonSetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorDaemonSetBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorDaemonSetBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorDaemonSetBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorDaemonSetBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorDaemonSetBuilder) AddDaemonSets(cb func(w *DaemonSetBuilder)) {
	x.buf.Reset()
	x.daemonSetBuilder.writer = &x.buf
	x.daemonSetBuilder.scratch = x.scratch
	cb(&x.daemonSetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorDaemonSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorDaemonSetBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorStatefulSetBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	statefulSetBuilder  StatefulSetBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorStatefulSetBuilder(writer io.Writer) *CollectorStatefulSetBuilder {
	return &CollectorStatefulSetBuilder{
		writer: writer,
	}
}
func (x *CollectorStatefulSetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorStatefulSetBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorStatefulSetBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorStatefulSetBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorStatefulSetBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorStatefulSetBuilder) AddStatefulSets(cb func(w *StatefulSetBuilder)) {
	x.buf.Reset()
	x.statefulSetBuilder.writer = &x.buf
	x.statefulSetBuilder.scratch = x.scratch
	cb(&x.statefulSetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorStatefulSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorStatefulSetBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorPersistentVolumeBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	persistentVolumeBuilder PersistentVolumeBuilder
	agentVersionBuilder     AgentVersionBuilder
}

func NewCollectorPersistentVolumeBuilder(writer io.Writer) *CollectorPersistentVolumeBuilder {
	return &CollectorPersistentVolumeBuilder{
		writer: writer,
	}
}
func (x *CollectorPersistentVolumeBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorPersistentVolumeBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeBuilder) AddPersistentVolumes(cb func(w *PersistentVolumeBuilder)) {
	x.buf.Reset()
	x.persistentVolumeBuilder.writer = &x.buf
	x.persistentVolumeBuilder.scratch = x.scratch
	cb(&x.persistentVolumeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorPersistentVolumeBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorPersistentVolumeClaimBuilder struct {
	writer                       io.Writer
	buf                          bytes.Buffer
	scratch                      []byte
	persistentVolumeClaimBuilder PersistentVolumeClaimBuilder
	agentVersionBuilder          AgentVersionBuilder
}

func NewCollectorPersistentVolumeClaimBuilder(writer io.Writer) *CollectorPersistentVolumeClaimBuilder {
	return &CollectorPersistentVolumeClaimBuilder{
		writer: writer,
	}
}
func (x *CollectorPersistentVolumeClaimBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorPersistentVolumeClaimBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeClaimBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeClaimBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeClaimBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeClaimBuilder) AddPersistentVolumeClaims(cb func(w *PersistentVolumeClaimBuilder)) {
	x.buf.Reset()
	x.persistentVolumeClaimBuilder.writer = &x.buf
	x.persistentVolumeClaimBuilder.scratch = x.scratch
	cb(&x.persistentVolumeClaimBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorPersistentVolumeClaimBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeClaimBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorRoleBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	roleBuilder         RoleBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorRoleBuilder(writer io.Writer) *CollectorRoleBuilder {
	return &CollectorRoleBuilder{
		writer: writer,
	}
}
func (x *CollectorRoleBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorRoleBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBuilder) AddRoles(cb func(w *RoleBuilder)) {
	x.buf.Reset()
	x.roleBuilder.writer = &x.buf
	x.roleBuilder.scratch = x.scratch
	cb(&x.roleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorRoleBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorRoleBindingBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	roleBindingBuilder  RoleBindingBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorRoleBindingBuilder(writer io.Writer) *CollectorRoleBindingBuilder {
	return &CollectorRoleBindingBuilder{
		writer: writer,
	}
}
func (x *CollectorRoleBindingBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorRoleBindingBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBindingBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBindingBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBindingBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBindingBuilder) AddRoleBindings(cb func(w *RoleBindingBuilder)) {
	x.buf.Reset()
	x.roleBindingBuilder.writer = &x.buf
	x.roleBindingBuilder.scratch = x.scratch
	cb(&x.roleBindingBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorRoleBindingBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBindingBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorClusterRoleBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	clusterRoleBuilder  ClusterRoleBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorClusterRoleBuilder(writer io.Writer) *CollectorClusterRoleBuilder {
	return &CollectorClusterRoleBuilder{
		writer: writer,
	}
}
func (x *CollectorClusterRoleBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorClusterRoleBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBuilder) AddClusterRoles(cb func(w *ClusterRoleBuilder)) {
	x.buf.Reset()
	x.clusterRoleBuilder.writer = &x.buf
	x.clusterRoleBuilder.scratch = x.scratch
	cb(&x.clusterRoleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorClusterRoleBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorClusterRoleBindingBuilder struct {
	writer                    io.Writer
	buf                       bytes.Buffer
	scratch                   []byte
	clusterRoleBindingBuilder ClusterRoleBindingBuilder
	agentVersionBuilder       AgentVersionBuilder
}

func NewCollectorClusterRoleBindingBuilder(writer io.Writer) *CollectorClusterRoleBindingBuilder {
	return &CollectorClusterRoleBindingBuilder{
		writer: writer,
	}
}
func (x *CollectorClusterRoleBindingBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorClusterRoleBindingBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBindingBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBindingBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBindingBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBindingBuilder) AddClusterRoleBindings(cb func(w *ClusterRoleBindingBuilder)) {
	x.buf.Reset()
	x.clusterRoleBindingBuilder.writer = &x.buf
	x.clusterRoleBindingBuilder.scratch = x.scratch
	cb(&x.clusterRoleBindingBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorClusterRoleBindingBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBindingBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorServiceAccountBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	serviceAccountBuilder ServiceAccountBuilder
	agentVersionBuilder   AgentVersionBuilder
}

func NewCollectorServiceAccountBuilder(writer io.Writer) *CollectorServiceAccountBuilder {
	return &CollectorServiceAccountBuilder{
		writer: writer,
	}
}
func (x *CollectorServiceAccountBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorServiceAccountBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceAccountBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceAccountBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceAccountBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceAccountBuilder) AddServiceAccounts(cb func(w *ServiceAccountBuilder)) {
	x.buf.Reset()
	x.serviceAccountBuilder.writer = &x.buf
	x.serviceAccountBuilder.scratch = x.scratch
	cb(&x.serviceAccountBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorServiceAccountBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceAccountBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorIngressBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	ingressBuilder      IngressBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorIngressBuilder(writer io.Writer) *CollectorIngressBuilder {
	return &CollectorIngressBuilder{
		writer: writer,
	}
}
func (x *CollectorIngressBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorIngressBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorIngressBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorIngressBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorIngressBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorIngressBuilder) AddIngresses(cb func(w *IngressBuilder)) {
	x.buf.Reset()
	x.ingressBuilder.writer = &x.buf
	x.ingressBuilder.scratch = x.scratch
	cb(&x.ingressBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorIngressBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorIngressBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorVerticalPodAutoscalerBuilder struct {
	writer                       io.Writer
	buf                          bytes.Buffer
	scratch                      []byte
	verticalPodAutoscalerBuilder VerticalPodAutoscalerBuilder
	agentVersionBuilder          AgentVersionBuilder
}

func NewCollectorVerticalPodAutoscalerBuilder(writer io.Writer) *CollectorVerticalPodAutoscalerBuilder {
	return &CollectorVerticalPodAutoscalerBuilder{
		writer: writer,
	}
}
func (x *CollectorVerticalPodAutoscalerBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorVerticalPodAutoscalerBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorVerticalPodAutoscalerBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorVerticalPodAutoscalerBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorVerticalPodAutoscalerBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorVerticalPodAutoscalerBuilder) AddVerticalPodAutoscalers(cb func(w *VerticalPodAutoscalerBuilder)) {
	x.buf.Reset()
	x.verticalPodAutoscalerBuilder.writer = &x.buf
	x.verticalPodAutoscalerBuilder.scratch = x.scratch
	cb(&x.verticalPodAutoscalerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorVerticalPodAutoscalerBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorVerticalPodAutoscalerBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorHorizontalPodAutoscalerBuilder struct {
	writer                         io.Writer
	buf                            bytes.Buffer
	scratch                        []byte
	horizontalPodAutoscalerBuilder HorizontalPodAutoscalerBuilder
	agentVersionBuilder            AgentVersionBuilder
}

func NewCollectorHorizontalPodAutoscalerBuilder(writer io.Writer) *CollectorHorizontalPodAutoscalerBuilder {
	return &CollectorHorizontalPodAutoscalerBuilder{
		writer: writer,
	}
}
func (x *CollectorHorizontalPodAutoscalerBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorHorizontalPodAutoscalerBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorHorizontalPodAutoscalerBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorHorizontalPodAutoscalerBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorHorizontalPodAutoscalerBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorHorizontalPodAutoscalerBuilder) AddHorizontalPodAutoscalers(cb func(w *HorizontalPodAutoscalerBuilder)) {
	x.buf.Reset()
	x.horizontalPodAutoscalerBuilder.writer = &x.buf
	x.horizontalPodAutoscalerBuilder.scratch = x.scratch
	cb(&x.horizontalPodAutoscalerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorHorizontalPodAutoscalerBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorHorizontalPodAutoscalerBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorNetworkPolicyBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	networkPolicyBuilder NetworkPolicyBuilder
	agentVersionBuilder  AgentVersionBuilder
}

func NewCollectorNetworkPolicyBuilder(writer io.Writer) *CollectorNetworkPolicyBuilder {
	return &CollectorNetworkPolicyBuilder{
		writer: writer,
	}
}
func (x *CollectorNetworkPolicyBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorNetworkPolicyBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNetworkPolicyBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNetworkPolicyBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorNetworkPolicyBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorNetworkPolicyBuilder) AddNetworkPolicies(cb func(w *NetworkPolicyBuilder)) {
	x.buf.Reset()
	x.networkPolicyBuilder.writer = &x.buf
	x.networkPolicyBuilder.scratch = x.scratch
	cb(&x.networkPolicyBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorNetworkPolicyBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNetworkPolicyBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorLimitRangeBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	limitRangeBuilder   LimitRangeBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorLimitRangeBuilder(writer io.Writer) *CollectorLimitRangeBuilder {
	return &CollectorLimitRangeBuilder{
		writer: writer,
	}
}
func (x *CollectorLimitRangeBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorLimitRangeBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorLimitRangeBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorLimitRangeBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorLimitRangeBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorLimitRangeBuilder) AddLimitRanges(cb func(w *LimitRangeBuilder)) {
	x.buf.Reset()
	x.limitRangeBuilder.writer = &x.buf
	x.limitRangeBuilder.scratch = x.scratch
	cb(&x.limitRangeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorLimitRangeBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorLimitRangeBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorStorageClassBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	storageClassBuilder StorageClassBuilder
	agentVersionBuilder AgentVersionBuilder
}

func NewCollectorStorageClassBuilder(writer io.Writer) *CollectorStorageClassBuilder {
	return &CollectorStorageClassBuilder{
		writer: writer,
	}
}
func (x *CollectorStorageClassBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorStorageClassBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorStorageClassBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorStorageClassBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorStorageClassBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorStorageClassBuilder) AddStorageClasses(cb func(w *StorageClassBuilder)) {
	x.buf.Reset()
	x.storageClassBuilder.writer = &x.buf
	x.storageClassBuilder.scratch = x.scratch
	cb(&x.storageClassBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorStorageClassBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorStorageClassBuilder) SetAgentVersion(cb func(w *AgentVersionBuilder)) {
	x.buf.Reset()
	x.agentVersionBuilder.writer = &x.buf
	x.agentVersionBuilder.scratch = x.scratch
	cb(&x.agentVersionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type AgentVersionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewAgentVersionBuilder(writer io.Writer) *AgentVersionBuilder {
	return &AgentVersionBuilder{
		writer: writer,
	}
}
func (x *AgentVersionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *AgentVersionBuilder) SetMajor(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *AgentVersionBuilder) SetMinor(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *AgentVersionBuilder) SetPatch(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *AgentVersionBuilder) SetPre(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AgentVersionBuilder) SetMeta(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AgentVersionBuilder) SetCommit(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCollectorStatusBuilder(writer io.Writer) *CollectorStatusBuilder {
	return &CollectorStatusBuilder{
		writer: writer,
	}
}
func (x *CollectorStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorStatusBuilder) SetActiveClients(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorStatusBuilder) SetInterval(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type TracerMetadataBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewTracerMetadataBuilder(writer io.Writer) *TracerMetadataBuilder {
	return &TracerMetadataBuilder{
		writer: writer,
	}
}
func (x *TracerMetadataBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *TracerMetadataBuilder) SetRuntimeId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TracerMetadataBuilder) SetServiceName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PortInfoBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPortInfoBuilder(writer io.Writer) *PortInfoBuilder {
	return &PortInfoBuilder{
		writer: writer,
	}
}
func (x *PortInfoBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PortInfoBuilder) AddTcp(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PortInfoBuilder) AddUdp(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ProcessBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	hostBuilder             HostBuilder
	commandBuilder          CommandBuilder
	processUserBuilder      ProcessUserBuilder
	memoryStatBuilder       MemoryStatBuilder
	cPUStatBuilder          CPUStatBuilder
	containerBuilder        ContainerBuilder
	iOStatBuilder           IOStatBuilder
	processNetworksBuilder  ProcessNetworksBuilder
	portInfoBuilder         PortInfoBuilder
	serviceDiscoveryBuilder ServiceDiscoveryBuilder
}

func NewProcessBuilder(writer io.Writer) *ProcessBuilder {
	return &ProcessBuilder{
		writer: writer,
	}
}
func (x *ProcessBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProcessBuilder) SetKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetPid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetNsPid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetCommand(cb func(w *CommandBuilder)) {
	x.buf.Reset()
	x.commandBuilder.writer = &x.buf
	x.commandBuilder.scratch = x.scratch
	cb(&x.commandBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetUser(cb func(w *ProcessUserBuilder)) {
	x.buf.Reset()
	x.processUserBuilder.writer = &x.buf
	x.processUserBuilder.scratch = x.scratch
	cb(&x.processUserBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetMemory(cb func(w *MemoryStatBuilder)) {
	x.buf.Reset()
	x.memoryStatBuilder.writer = &x.buf
	x.memoryStatBuilder.scratch = x.scratch
	cb(&x.memoryStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetCpu(cb func(w *CPUStatBuilder)) {
	x.buf.Reset()
	x.cPUStatBuilder.writer = &x.buf
	x.cPUStatBuilder.scratch = x.scratch
	cb(&x.cPUStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetCreateTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetContainer(cb func(w *ContainerBuilder)) {
	x.buf.Reset()
	x.containerBuilder.writer = &x.buf
	x.containerBuilder.scratch = x.scratch
	cb(&x.containerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetOpenFdCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x58)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetState(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x60)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ProcessBuilder) SetIoStat(cb func(w *IOStatBuilder)) {
	x.buf.Reset()
	x.iOStatBuilder.writer = &x.buf
	x.iOStatBuilder.scratch = x.scratch
	cb(&x.iOStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x6a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetContainerId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x72)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetContainerKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x78)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetVoluntaryCtxSwitches(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x80)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetInvoluntaryCtxSwitches(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x88)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x92)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetContainerByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x9a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetNetworks(cb func(w *ProcessNetworksBuilder)) {
	x.buf.Reset()
	x.processNetworksBuilder.writer = &x.buf
	x.processNetworksBuilder.scratch = x.scratch
	cb(&x.processNetworksBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xaa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) AddProcessContext(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xb2)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xba)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetLanguage(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0xc0)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ProcessBuilder) SetPortInfo(cb func(w *PortInfoBuilder)) {
	x.buf.Reset()
	x.portInfoBuilder.writer = &x.buf
	x.portInfoBuilder.scratch = x.scratch
	cb(&x.portInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xca)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetServiceDiscovery(cb func(w *ServiceDiscoveryBuilder)) {
	x.buf.Reset()
	x.serviceDiscoveryBuilder.writer = &x.buf
	x.serviceDiscoveryBuilder.scratch = x.scratch
	cb(&x.serviceDiscoveryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xd2)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetInjectionState(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0xd8)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type ServiceDiscoveryBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	serviceNameBuilder    ServiceNameBuilder
	tracerMetadataBuilder TracerMetadataBuilder
	resourceBuilder       ResourceBuilder
}

func NewServiceDiscoveryBuilder(writer io.Writer) *ServiceDiscoveryBuilder {
	return &ServiceDiscoveryBuilder{
		writer: writer,
	}
}
func (x *ServiceDiscoveryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ServiceDiscoveryBuilder) SetGeneratedServiceName(cb func(w *ServiceNameBuilder)) {
	x.buf.Reset()
	x.serviceNameBuilder.writer = &x.buf
	x.serviceNameBuilder.scratch = x.scratch
	cb(&x.serviceNameBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceDiscoveryBuilder) SetDdServiceName(cb func(w *ServiceNameBuilder)) {
	x.buf.Reset()
	x.serviceNameBuilder.writer = &x.buf
	x.serviceNameBuilder.scratch = x.scratch
	cb(&x.serviceNameBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceDiscoveryBuilder) AddAdditionalGeneratedNames(cb func(w *ServiceNameBuilder)) {
	x.buf.Reset()
	x.serviceNameBuilder.writer = &x.buf
	x.serviceNameBuilder.scratch = x.scratch
	cb(&x.serviceNameBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceDiscoveryBuilder) AddTracerMetadata(cb func(w *TracerMetadataBuilder)) {
	x.buf.Reset()
	x.tracerMetadataBuilder.writer = &x.buf
	x.tracerMetadataBuilder.scratch = x.scratch
	cb(&x.tracerMetadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceDiscoveryBuilder) SetApmInstrumentation(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x28)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *ServiceDiscoveryBuilder) AddResources(cb func(w *ResourceBuilder)) {
	x.buf.Reset()
	x.resourceBuilder.writer = &x.buf
	x.resourceBuilder.scratch = x.scratch
	cb(&x.resourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ResourceBuilder struct {
	writer             io.Writer
	buf                bytes.Buffer
	scratch            []byte
	logResourceBuilder LogResourceBuilder
}

func NewResourceBuilder(writer io.Writer) *ResourceBuilder {
	return &ResourceBuilder{
		writer: writer,
	}
}
func (x *ResourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResourceBuilder) SetLogs(cb func(w *LogResourceBuilder)) {
	x.buf.Reset()
	x.logResourceBuilder.writer = &x.buf
	x.logResourceBuilder.scratch = x.scratch
	cb(&x.logResourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type LogResourceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewLogResourceBuilder(writer io.Writer) *LogResourceBuilder {
	return &LogResourceBuilder{
		writer: writer,
	}
}
func (x *LogResourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *LogResourceBuilder) SetPath(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ServiceNameBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewServiceNameBuilder(writer io.Writer) *ServiceNameBuilder {
	return &ServiceNameBuilder{
		writer: writer,
	}
}
func (x *ServiceNameBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ServiceNameBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceNameBuilder) SetSource(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x10)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type ProcessDiscoveryBuilder struct {
	writer             io.Writer
	buf                bytes.Buffer
	scratch            []byte
	hostBuilder        HostBuilder
	commandBuilder     CommandBuilder
	processUserBuilder ProcessUserBuilder
}

func NewProcessDiscoveryBuilder(writer io.Writer) *ProcessDiscoveryBuilder {
	return &ProcessDiscoveryBuilder{
		writer: writer,
	}
}
func (x *ProcessDiscoveryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProcessDiscoveryBuilder) SetPid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessDiscoveryBuilder) SetNsPid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessDiscoveryBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessDiscoveryBuilder) SetCommand(cb func(w *CommandBuilder)) {
	x.buf.Reset()
	x.commandBuilder.writer = &x.buf
	x.commandBuilder.scratch = x.scratch
	cb(&x.commandBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessDiscoveryBuilder) SetUser(cb func(w *ProcessUserBuilder)) {
	x.buf.Reset()
	x.processUserBuilder.writer = &x.buf
	x.processUserBuilder.scratch = x.scratch
	cb(&x.processUserBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessDiscoveryBuilder) SetCreateTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessDiscoveryBuilder) SetByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CommandBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCommandBuilder(writer io.Writer) *CommandBuilder {
	return &CommandBuilder{
		writer: writer,
	}
}
func (x *CommandBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CommandBuilder) AddArgs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CommandBuilder) SetCwd(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CommandBuilder) SetRoot(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CommandBuilder) SetOnDisk(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x28)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *CommandBuilder) SetPpid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CommandBuilder) SetPgroup(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CommandBuilder) SetExe(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CommandBuilder) SetComm(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ProcessUserBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewProcessUserBuilder(writer io.Writer) *ProcessUserBuilder {
	return &ProcessUserBuilder{
		writer: writer,
	}
}
func (x *ProcessUserBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProcessUserBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetUid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetGid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetEuid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetEgid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetSuid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetSgid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ProcessNetworksBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewProcessNetworksBuilder(writer io.Writer) *ProcessNetworksBuilder {
	return &ProcessNetworksBuilder{
		writer: writer,
	}
}
func (x *ProcessNetworksBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProcessNetworksBuilder) SetConnectionRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xd)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessNetworksBuilder) SetBytesRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x15)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}

type ContainerAddrBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewContainerAddrBuilder(writer io.Writer) *ContainerAddrBuilder {
	return &ContainerAddrBuilder{
		writer: writer,
	}
}
func (x *ContainerAddrBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ContainerAddrBuilder) SetIp(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerAddrBuilder) SetPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerAddrBuilder) SetProtocol(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x18)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type ContainerBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	hostBuilder          HostBuilder
	containerAddrBuilder ContainerAddrBuilder
}

func NewContainerBuilder(writer io.Writer) *ContainerBuilder {
	return &ContainerBuilder{
		writer: writer,
	}
}
func (x *ContainerBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ContainerBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetImage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetCpuLimit(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetMemoryLimit(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetState(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x40)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ContainerBuilder) SetHealth(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x48)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ContainerBuilder) SetCreated(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x50)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetRbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetWbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x65)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x68)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetNetRcvdPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x75)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetNetSentPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x7d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetNetRcvdBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x85)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetNetSentBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetUserPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x95)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetSystemPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x9d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetTotalPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa5)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetMemRss(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetMemCache(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xb0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xba)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerBuilder) SetStarted(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xc0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xca)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xd2)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) AddAddresses(cb func(w *ContainerAddrBuilder)) {
	x.buf.Reset()
	x.containerAddrBuilder.writer = &x.buf
	x.containerAddrBuilder.scratch = x.scratch
	cb(&x.containerAddrBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xda)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerBuilder) SetThreadCount(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xe0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetThreadLimit(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xe8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetMemUsage(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xf0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetCpuUsageNs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xfd)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetMemAccounted(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x100)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetCpuRequest(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x10d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetMemoryRequest(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x110)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetRepoDigest(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x11a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ProcessStatBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	memoryStatBuilder      MemoryStatBuilder
	cPUStatBuilder         CPUStatBuilder
	iOStatBuilder          IOStatBuilder
	processNetworksBuilder ProcessNetworksBuilder
}

func NewProcessStatBuilder(writer io.Writer) *ProcessStatBuilder {
	return &ProcessStatBuilder{
		writer: writer,
	}
}
func (x *ProcessStatBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProcessStatBuilder) SetPid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetCreateTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetMemory(cb func(w *MemoryStatBuilder)) {
	x.buf.Reset()
	x.memoryStatBuilder.writer = &x.buf
	x.memoryStatBuilder.scratch = x.scratch
	cb(&x.memoryStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessStatBuilder) SetCpu(cb func(w *CPUStatBuilder)) {
	x.buf.Reset()
	x.cPUStatBuilder.writer = &x.buf
	x.cPUStatBuilder.scratch = x.scratch
	cb(&x.cPUStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessStatBuilder) SetNice(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetThreads(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetOpenFdCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerState(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x58)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ProcessStatBuilder) SetProcessState(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x60)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ProcessStatBuilder) SetIoStat(cb func(w *IOStatBuilder)) {
	x.buf.Reset()
	x.iOStatBuilder.writer = &x.buf
	x.iOStatBuilder.scratch = x.scratch
	cb(&x.iOStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x9a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessStatBuilder) SetNetworks(cb func(w *ProcessNetworksBuilder)) {
	x.buf.Reset()
	x.processNetworksBuilder.writer = &x.buf
	x.processNetworksBuilder.scratch = x.scratch
	cb(&x.processNetworksBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xe2)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessStatBuilder) SetContainerHealth(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x78)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ProcessStatBuilder) SetContainerRbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x85)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerWbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x90)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerNetRcvdPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa5)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerNetSentPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xad)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerNetRcvdBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xb5)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerNetSentBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xbd)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetVoluntaryCtxSwitches(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xc0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetInvoluntaryCtxSwitches(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xc8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xd2)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessStatBuilder) SetContainerByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xda)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ProcStatsWithPermBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewProcStatsWithPermBuilder(writer io.Writer) *ProcStatsWithPermBuilder {
	return &ProcStatsWithPermBuilder{
		writer: writer,
	}
}
func (x *ProcStatsWithPermBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProcStatsWithPermBuilder) SetOpenFDCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcStatsWithPermBuilder) SetReadCount(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcStatsWithPermBuilder) SetWriteCount(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcStatsWithPermBuilder) SetReadBytes(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcStatsWithPermBuilder) SetWriteBytes(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ProcStatsWithPermByPIDBuilder struct {
	writer                                        io.Writer
	buf                                           bytes.Buffer
	scratch                                       []byte
	procStatsWithPermByPID_StatsByPIDEntryBuilder ProcStatsWithPermByPID_StatsByPIDEntryBuilder
}

func NewProcStatsWithPermByPIDBuilder(writer io.Writer) *ProcStatsWithPermByPIDBuilder {
	return &ProcStatsWithPermByPIDBuilder{
		writer: writer,
	}
}
func (x *ProcStatsWithPermByPIDBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProcStatsWithPermByPIDBuilder) AddStatsByPID(cb func(w *ProcStatsWithPermByPID_StatsByPIDEntryBuilder)) {
	x.buf.Reset()
	x.procStatsWithPermByPID_StatsByPIDEntryBuilder.writer = &x.buf
	x.procStatsWithPermByPID_StatsByPIDEntryBuilder.scratch = x.scratch
	cb(&x.procStatsWithPermByPID_StatsByPIDEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ProcStatsWithPermByPID_StatsByPIDEntryBuilder struct {
	writer                   io.Writer
	buf                      bytes.Buffer
	scratch                  []byte
	procStatsWithPermBuilder ProcStatsWithPermBuilder
}

func NewProcStatsWithPermByPID_StatsByPIDEntryBuilder(writer io.Writer) *ProcStatsWithPermByPID_StatsByPIDEntryBuilder {
	return &ProcStatsWithPermByPID_StatsByPIDEntryBuilder{
		writer: writer,
	}
}
func (x *ProcStatsWithPermByPID_StatsByPIDEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProcStatsWithPermByPID_StatsByPIDEntryBuilder) SetKey(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcStatsWithPermByPID_StatsByPIDEntryBuilder) SetValue(cb func(w *ProcStatsWithPermBuilder)) {
	x.buf.Reset()
	x.procStatsWithPermBuilder.writer = &x.buf
	x.procStatsWithPermBuilder.scratch = x.scratch
	cb(&x.procStatsWithPermBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ContainerStatBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewContainerStatBuilder(writer io.Writer) *ContainerStatBuilder {
	return &ContainerStatBuilder{
		writer: writer,
	}
}
func (x *ContainerStatBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ContainerStatBuilder) SetId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetUserPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x15)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetSystemPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetTotalPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x25)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetCpuLimit(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetMemRss(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetMemCache(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetMemLimit(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetRbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x4d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetWbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x55)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetNetRcvdPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetNetSentPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x65)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetNetRcvdBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x6d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetNetSentBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x75)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetState(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x78)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ContainerStatBuilder) SetHealth(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x80)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ContainerStatBuilder) SetKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x88)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetStarted(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x90)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x9a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerStatBuilder) SetThreadCount(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetThreadLimit(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetMemUsage(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xb0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetCpuUsageNs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xbd)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetMemAccounted(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xc0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetCpuRequest(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xcd)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetMemoryRequest(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xd0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type SystemInfoBuilder struct {
	writer         io.Writer
	buf            bytes.Buffer
	scratch        []byte
	oSInfoBuilder  OSInfoBuilder
	cPUInfoBuilder CPUInfoBuilder
}

func NewSystemInfoBuilder(writer io.Writer) *SystemInfoBuilder {
	return &SystemInfoBuilder{
		writer: writer,
	}
}
func (x *SystemInfoBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *SystemInfoBuilder) SetUuid(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SystemInfoBuilder) SetOs(cb func(w *OSInfoBuilder)) {
	x.buf.Reset()
	x.oSInfoBuilder.writer = &x.buf
	x.oSInfoBuilder.scratch = x.scratch
	cb(&x.oSInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *SystemInfoBuilder) AddCpus(cb func(w *CPUInfoBuilder)) {
	x.buf.Reset()
	x.cPUInfoBuilder.writer = &x.buf
	x.cPUInfoBuilder.scratch = x.scratch
	cb(&x.cPUInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *SystemInfoBuilder) SetTotalMemory(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type OSInfoBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewOSInfoBuilder(writer io.Writer) *OSInfoBuilder {
	return &OSInfoBuilder{
		writer: writer,
	}
}
func (x *OSInfoBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *OSInfoBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OSInfoBuilder) SetPlatform(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OSInfoBuilder) SetFamily(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OSInfoBuilder) SetVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OSInfoBuilder) SetKernelVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type IOStatBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewIOStatBuilder(writer io.Writer) *IOStatBuilder {
	return &IOStatBuilder{
		writer: writer,
	}
}
func (x *IOStatBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *IOStatBuilder) SetReadRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xd)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *IOStatBuilder) SetWriteRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x15)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *IOStatBuilder) SetReadBytesRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *IOStatBuilder) SetWriteBytesRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x25)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}

type MemoryStatBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewMemoryStatBuilder(writer io.Writer) *MemoryStatBuilder {
	return &MemoryStatBuilder{
		writer: writer,
	}
}
func (x *MemoryStatBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *MemoryStatBuilder) SetRss(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetVms(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetSwap(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetShared(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetText(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetLib(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetData(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetDirty(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type CPUStatBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	singleCPUStatBuilder SingleCPUStatBuilder
}

func NewCPUStatBuilder(writer io.Writer) *CPUStatBuilder {
	return &CPUStatBuilder{
		writer: writer,
	}
}
func (x *CPUStatBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CPUStatBuilder) SetLastCpu(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetTotalPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x15)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetUserPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1d)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetSystemPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x25)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetNumThreads(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) AddCpus(cb func(w *SingleCPUStatBuilder)) {
	x.buf.Reset()
	x.singleCPUStatBuilder.writer = &x.buf
	x.singleCPUStatBuilder.scratch = x.scratch
	cb(&x.singleCPUStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CPUStatBuilder) SetNice(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetUserTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetSystemTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type SingleCPUStatBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewSingleCPUStatBuilder(writer io.Writer) *SingleCPUStatBuilder {
	return &SingleCPUStatBuilder{
		writer: writer,
	}
}
func (x *SingleCPUStatBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *SingleCPUStatBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SingleCPUStatBuilder) SetTotalPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x15)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}

type CPUInfoBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCPUInfoBuilder(writer io.Writer) *CPUInfoBuilder {
	return &CPUInfoBuilder{
		writer: writer,
	}
}
func (x *CPUInfoBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CPUInfoBuilder) SetNumber(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetVendor(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetFamily(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetModel(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetPhysicalId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetCoreId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetCores(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetMhz(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetCacheSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type HostBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewHostBuilder(writer io.Writer) *HostBuilder {
	return &HostBuilder{
		writer: writer,
	}
}
func (x *HostBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HostBuilder) SetId(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetOrgId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) AddAllTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetNumCpus(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetTotalMemory(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetTagIndex(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetTagsModified(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x50)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ClusterBuilder struct {
	writer                                           io.Writer
	buf                                              bytes.Buffer
	scratch                                          []byte
	cluster_KubeletVersionsEntryBuilder              Cluster_KubeletVersionsEntryBuilder
	cluster_ApiServerVersionsEntryBuilder            Cluster_ApiServerVersionsEntryBuilder
	resourceMetricsBuilder                           ResourceMetricsBuilder
	cluster_ExtendedResourcesAllocatableEntryBuilder Cluster_ExtendedResourcesAllocatableEntryBuilder
	cluster_ExtendedResourcesCapacityEntryBuilder    Cluster_ExtendedResourcesCapacityEntryBuilder
	clusterNodeInfoBuilder                           ClusterNodeInfoBuilder
}

func NewClusterBuilder(writer io.Writer) *ClusterBuilder {
	return &ClusterBuilder{
		writer: writer,
	}
}
func (x *ClusterBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ClusterBuilder) SetNodeCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) AddKubeletVersions(cb func(w *Cluster_KubeletVersionsEntryBuilder)) {
	x.buf.Reset()
	x.cluster_KubeletVersionsEntryBuilder.writer = &x.buf
	x.cluster_KubeletVersionsEntryBuilder.scratch = x.scratch
	cb(&x.cluster_KubeletVersionsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterBuilder) AddApiServerVersions(cb func(w *Cluster_ApiServerVersionsEntryBuilder)) {
	x.buf.Reset()
	x.cluster_ApiServerVersionsEntryBuilder.writer = &x.buf
	x.cluster_ApiServerVersionsEntryBuilder.scratch = x.scratch
	cb(&x.cluster_ApiServerVersionsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterBuilder) SetPodCapacity(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetPodAllocatable(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetMemoryAllocatable(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetMemoryCapacity(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetCpuAllocatable(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetCpuCapacity(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetResourceVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetCreationTimestamp(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x58)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x62)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x6a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterBuilder) AddExtendedResourcesAllocatable(cb func(w *Cluster_ExtendedResourcesAllocatableEntryBuilder)) {
	x.buf.Reset()
	x.cluster_ExtendedResourcesAllocatableEntryBuilder.writer = &x.buf
	x.cluster_ExtendedResourcesAllocatableEntryBuilder.scratch = x.scratch
	cb(&x.cluster_ExtendedResourcesAllocatableEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x72)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterBuilder) AddExtendedResourcesCapacity(cb func(w *Cluster_ExtendedResourcesCapacityEntryBuilder)) {
	x.buf.Reset()
	x.cluster_ExtendedResourcesCapacityEntryBuilder.writer = &x.buf
	x.cluster_ExtendedResourcesCapacityEntryBuilder.scratch = x.scratch
	cb(&x.cluster_ExtendedResourcesCapacityEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x7a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterBuilder) AddNodesInfo(cb func(w *ClusterNodeInfoBuilder)) {
	x.buf.Reset()
	x.clusterNodeInfoBuilder.writer = &x.buf
	x.clusterNodeInfoBuilder.scratch = x.scratch
	cb(&x.clusterNodeInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x82)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type Cluster_KubeletVersionsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCluster_KubeletVersionsEntryBuilder(writer io.Writer) *Cluster_KubeletVersionsEntryBuilder {
	return &Cluster_KubeletVersionsEntryBuilder{
		writer: writer,
	}
}
func (x *Cluster_KubeletVersionsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Cluster_KubeletVersionsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *Cluster_KubeletVersionsEntryBuilder) SetValue(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type Cluster_ApiServerVersionsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCluster_ApiServerVersionsEntryBuilder(writer io.Writer) *Cluster_ApiServerVersionsEntryBuilder {
	return &Cluster_ApiServerVersionsEntryBuilder{
		writer: writer,
	}
}
func (x *Cluster_ApiServerVersionsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Cluster_ApiServerVersionsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *Cluster_ApiServerVersionsEntryBuilder) SetValue(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type Cluster_ExtendedResourcesAllocatableEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCluster_ExtendedResourcesAllocatableEntryBuilder(writer io.Writer) *Cluster_ExtendedResourcesAllocatableEntryBuilder {
	return &Cluster_ExtendedResourcesAllocatableEntryBuilder{
		writer: writer,
	}
}
func (x *Cluster_ExtendedResourcesAllocatableEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Cluster_ExtendedResourcesAllocatableEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *Cluster_ExtendedResourcesAllocatableEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type Cluster_ExtendedResourcesCapacityEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCluster_ExtendedResourcesCapacityEntryBuilder(writer io.Writer) *Cluster_ExtendedResourcesCapacityEntryBuilder {
	return &Cluster_ExtendedResourcesCapacityEntryBuilder{
		writer: writer,
	}
}
func (x *Cluster_ExtendedResourcesCapacityEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *Cluster_ExtendedResourcesCapacityEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *Cluster_ExtendedResourcesCapacityEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ClusterNodeInfoBuilder struct {
	writer                                          io.Writer
	buf                                             bytes.Buffer
	scratch                                         []byte
	clusterNodeInfo_ResourceAllocatableEntryBuilder ClusterNodeInfo_ResourceAllocatableEntryBuilder
	clusterNodeInfo_ResourceCapacityEntryBuilder    ClusterNodeInfo_ResourceCapacityEntryBuilder
}

func NewClusterNodeInfoBuilder(writer io.Writer) *ClusterNodeInfoBuilder {
	return &ClusterNodeInfoBuilder{
		writer: writer,
	}
}
func (x *ClusterNodeInfoBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ClusterNodeInfoBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterNodeInfoBuilder) SetRegion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterNodeInfoBuilder) SetInstanceType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterNodeInfoBuilder) SetOperatingSystem(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterNodeInfoBuilder) SetOperatingSystemImage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterNodeInfoBuilder) SetArchitecture(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterNodeInfoBuilder) SetKernelVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterNodeInfoBuilder) SetContainerRuntimeVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterNodeInfoBuilder) SetKubeletVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterNodeInfoBuilder) AddResourceAllocatable(cb func(w *ClusterNodeInfo_ResourceAllocatableEntryBuilder)) {
	x.buf.Reset()
	x.clusterNodeInfo_ResourceAllocatableEntryBuilder.writer = &x.buf
	x.clusterNodeInfo_ResourceAllocatableEntryBuilder.scratch = x.scratch
	cb(&x.clusterNodeInfo_ResourceAllocatableEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterNodeInfoBuilder) AddResourceCapacity(cb func(w *ClusterNodeInfo_ResourceCapacityEntryBuilder)) {
	x.buf.Reset()
	x.clusterNodeInfo_ResourceCapacityEntryBuilder.writer = &x.buf
	x.clusterNodeInfo_ResourceCapacityEntryBuilder.scratch = x.scratch
	cb(&x.clusterNodeInfo_ResourceCapacityEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ClusterNodeInfo_ResourceAllocatableEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewClusterNodeInfo_ResourceAllocatableEntryBuilder(writer io.Writer) *ClusterNodeInfo_ResourceAllocatableEntryBuilder {
	return &ClusterNodeInfo_ResourceAllocatableEntryBuilder{
		writer: writer,
	}
}
func (x *ClusterNodeInfo_ResourceAllocatableEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ClusterNodeInfo_ResourceAllocatableEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterNodeInfo_ResourceAllocatableEntryBuilder) SetValue(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ClusterNodeInfo_ResourceCapacityEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewClusterNodeInfo_ResourceCapacityEntryBuilder(writer io.Writer) *ClusterNodeInfo_ResourceCapacityEntryBuilder {
	return &ClusterNodeInfo_ResourceCapacityEntryBuilder{
		writer: writer,
	}
}
func (x *ClusterNodeInfo_ResourceCapacityEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ClusterNodeInfo_ResourceCapacityEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterNodeInfo_ResourceCapacityEntryBuilder) SetValue(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type MetadataBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	ownerReferenceBuilder OwnerReferenceBuilder
}

func NewMetadataBuilder(writer io.Writer) *MetadataBuilder {
	return &MetadataBuilder{
		writer: writer,
	}
}
func (x *MetadataBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *MetadataBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) SetNamespace(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) SetUid(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) SetCreationTimestamp(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) SetDeletionTimestamp(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) AddLabels(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) AddAnnotations(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) AddOwnerReferences(cb func(w *OwnerReferenceBuilder)) {
	x.buf.Reset()
	x.ownerReferenceBuilder.writer = &x.buf
	x.ownerReferenceBuilder.scratch = x.scratch
	cb(&x.ownerReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *MetadataBuilder) SetResourceVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) AddFinalizers(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) SetDeletionGracePeriodSeconds(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x58)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type OwnerReferenceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewOwnerReferenceBuilder(writer io.Writer) *OwnerReferenceBuilder {
	return &OwnerReferenceBuilder{
		writer: writer,
	}
}
func (x *OwnerReferenceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *OwnerReferenceBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OwnerReferenceBuilder) SetUid(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OwnerReferenceBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ObjectReferenceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewObjectReferenceBuilder(writer io.Writer) *ObjectReferenceBuilder {
	return &ObjectReferenceBuilder{
		writer: writer,
	}
}
func (x *ObjectReferenceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ObjectReferenceBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetNamespace(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetUid(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetApiVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetResourceVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetFieldPath(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ServicePortBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewServicePortBuilder(writer io.Writer) *ServicePortBuilder {
	return &ServicePortBuilder{
		writer: writer,
	}
}
func (x *ServicePortBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ServicePortBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServicePortBuilder) SetProtocol(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServicePortBuilder) SetPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ServicePortBuilder) SetTargetPort(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServicePortBuilder) SetNodePort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ServiceSessionAffinityConfigBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewServiceSessionAffinityConfigBuilder(writer io.Writer) *ServiceSessionAffinityConfigBuilder {
	return &ServiceSessionAffinityConfigBuilder{
		writer: writer,
	}
}
func (x *ServiceSessionAffinityConfigBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ServiceSessionAffinityConfigBuilder) SetClientIPTimeoutSeconds(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type NodeBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	metadataBuilder        MetadataBuilder
	taintBuilder           TaintBuilder
	nodeStatusBuilder      NodeStatusBuilder
	hostBuilder            HostBuilder
	resourceMetricsBuilder ResourceMetricsBuilder
}

func NewNodeBuilder(writer io.Writer) *NodeBuilder {
	return &NodeBuilder{
		writer: writer,
	}
}
func (x *NodeBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NodeBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeBuilder) SetPodCIDR(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeBuilder) AddPodCIDRs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeBuilder) SetUnschedulable(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *NodeBuilder) AddTaints(cb func(w *TaintBuilder)) {
	x.buf.Reset()
	x.taintBuilder.writer = &x.buf
	x.taintBuilder.scratch = x.scratch
	cb(&x.taintBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeBuilder) SetStatus(cb func(w *NodeStatusBuilder)) {
	x.buf.Reset()
	x.nodeStatusBuilder.writer = &x.buf
	x.nodeStatusBuilder.scratch = x.scratch
	cb(&x.nodeStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeBuilder) AddRoles(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeBuilder) SetProviderID(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x62)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type NodeStatusBuilder struct {
	writer                               io.Writer
	buf                                  bytes.Buffer
	scratch                              []byte
	nodeStatus_CapacityEntryBuilder      NodeStatus_CapacityEntryBuilder
	nodeStatus_AllocatableEntryBuilder   NodeStatus_AllocatableEntryBuilder
	nodeStatus_NodeAddressesEntryBuilder NodeStatus_NodeAddressesEntryBuilder
	nodeConditionBuilder                 NodeConditionBuilder
	containerImageBuilder                ContainerImageBuilder
}

func NewNodeStatusBuilder(writer io.Writer) *NodeStatusBuilder {
	return &NodeStatusBuilder{
		writer: writer,
	}
}
func (x *NodeStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NodeStatusBuilder) AddCapacity(cb func(w *NodeStatus_CapacityEntryBuilder)) {
	x.buf.Reset()
	x.nodeStatus_CapacityEntryBuilder.writer = &x.buf
	x.nodeStatus_CapacityEntryBuilder.scratch = x.scratch
	cb(&x.nodeStatus_CapacityEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeStatusBuilder) AddAllocatable(cb func(w *NodeStatus_AllocatableEntryBuilder)) {
	x.buf.Reset()
	x.nodeStatus_AllocatableEntryBuilder.writer = &x.buf
	x.nodeStatus_AllocatableEntryBuilder.scratch = x.scratch
	cb(&x.nodeStatus_AllocatableEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeStatusBuilder) AddNodeAddresses(cb func(w *NodeStatus_NodeAddressesEntryBuilder)) {
	x.buf.Reset()
	x.nodeStatus_NodeAddressesEntryBuilder.writer = &x.buf
	x.nodeStatus_NodeAddressesEntryBuilder.scratch = x.scratch
	cb(&x.nodeStatus_NodeAddressesEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeStatusBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetKubeletVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) AddConditions(cb func(w *NodeConditionBuilder)) {
	x.buf.Reset()
	x.nodeConditionBuilder.writer = &x.buf
	x.nodeConditionBuilder.scratch = x.scratch
	cb(&x.nodeConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeStatusBuilder) AddImages(cb func(w *ContainerImageBuilder)) {
	x.buf.Reset()
	x.containerImageBuilder.writer = &x.buf
	x.containerImageBuilder.scratch = x.scratch
	cb(&x.containerImageBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeStatusBuilder) SetKubeProxyVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetOperatingSystem(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetArchitecture(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetKernelVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x5a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetOsImage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x62)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetContainerRuntimeVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x6a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NodeStatus_CapacityEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNodeStatus_CapacityEntryBuilder(writer io.Writer) *NodeStatus_CapacityEntryBuilder {
	return &NodeStatus_CapacityEntryBuilder{
		writer: writer,
	}
}
func (x *NodeStatus_CapacityEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NodeStatus_CapacityEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatus_CapacityEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type NodeStatus_AllocatableEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNodeStatus_AllocatableEntryBuilder(writer io.Writer) *NodeStatus_AllocatableEntryBuilder {
	return &NodeStatus_AllocatableEntryBuilder{
		writer: writer,
	}
}
func (x *NodeStatus_AllocatableEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NodeStatus_AllocatableEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatus_AllocatableEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type NodeStatus_NodeAddressesEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNodeStatus_NodeAddressesEntryBuilder(writer io.Writer) *NodeStatus_NodeAddressesEntryBuilder {
	return &NodeStatus_NodeAddressesEntryBuilder{
		writer: writer,
	}
}
func (x *NodeStatus_NodeAddressesEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NodeStatus_NodeAddressesEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatus_NodeAddressesEntryBuilder) SetValue(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NodeConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNodeConditionBuilder(writer io.Writer) *NodeConditionBuilder {
	return &NodeConditionBuilder{
		writer: writer,
	}
}
func (x *NodeConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NodeConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *NodeConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ContainerImageBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewContainerImageBuilder(writer io.Writer) *ContainerImageBuilder {
	return &ContainerImageBuilder{
		writer: writer,
	}
}
func (x *ContainerImageBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ContainerImageBuilder) AddNames(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerImageBuilder) SetSizeBytes(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type TaintBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewTaintBuilder(writer io.Writer) *TaintBuilder {
	return &TaintBuilder{
		writer: writer,
	}
}
func (x *TaintBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *TaintBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TaintBuilder) SetValue(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TaintBuilder) SetEffect(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TaintBuilder) SetTimeAdded(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ServiceSpecBuilder struct {
	writer                              io.Writer
	buf                                 bytes.Buffer
	scratch                             []byte
	servicePortBuilder                  ServicePortBuilder
	labelSelectorRequirementBuilder     LabelSelectorRequirementBuilder
	serviceSessionAffinityConfigBuilder ServiceSessionAffinityConfigBuilder
}

func NewServiceSpecBuilder(writer io.Writer) *ServiceSpecBuilder {
	return &ServiceSpecBuilder{
		writer: writer,
	}
}
func (x *ServiceSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ServiceSpecBuilder) AddPorts(cb func(w *ServicePortBuilder)) {
	x.buf.Reset()
	x.servicePortBuilder.writer = &x.buf
	x.servicePortBuilder.scratch = x.scratch
	cb(&x.servicePortBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceSpecBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceSpecBuilder) SetClusterIP(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) AddExternalIPs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetSessionAffinity(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetLoadBalancerIP(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) AddLoadBalancerSourceRanges(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetExternalName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetExternalTrafficPolicy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetHealthCheckNodePort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x58)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetPublishNotReadyAddresses(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x60)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *ServiceSpecBuilder) SetSessionAffinityConfig(cb func(w *ServiceSessionAffinityConfigBuilder)) {
	x.buf.Reset()
	x.serviceSessionAffinityConfigBuilder.writer = &x.buf
	x.serviceSessionAffinityConfigBuilder.scratch = x.scratch
	cb(&x.serviceSessionAffinityConfigBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x6a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceSpecBuilder) SetIpFamily(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x72)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ServiceStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewServiceStatusBuilder(writer io.Writer) *ServiceStatusBuilder {
	return &ServiceStatusBuilder{
		writer: writer,
	}
}
func (x *ServiceStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ServiceStatusBuilder) AddLoadBalancerIngress(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ServiceBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	metadataBuilder        MetadataBuilder
	serviceSpecBuilder     ServiceSpecBuilder
	serviceStatusBuilder   ServiceStatusBuilder
	resourceMetricsBuilder ResourceMetricsBuilder
}

func NewServiceBuilder(writer io.Writer) *ServiceBuilder {
	return &ServiceBuilder{
		writer: writer,
	}
}
func (x *ServiceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ServiceBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceBuilder) SetSpec(cb func(w *ServiceSpecBuilder)) {
	x.buf.Reset()
	x.serviceSpecBuilder.writer = &x.buf
	x.serviceSpecBuilder.scratch = x.scratch
	cb(&x.serviceSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceBuilder) SetStatus(cb func(w *ServiceStatusBuilder)) {
	x.buf.Reset()
	x.serviceStatusBuilder.writer = &x.buf
	x.serviceStatusBuilder.scratch = x.scratch
	cb(&x.serviceStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type DeploymentConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewDeploymentConditionBuilder(writer io.Writer) *DeploymentConditionBuilder {
	return &DeploymentConditionBuilder{
		writer: writer,
	}
}
func (x *DeploymentConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DeploymentConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentConditionBuilder) SetLastUpdateTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type DeploymentBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	metadataBuilder                 MetadataBuilder
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceRequirementsBuilder     ResourceRequirementsBuilder
	resourceMetricsBuilder          ResourceMetricsBuilder
	deploymentConditionBuilder      DeploymentConditionBuilder
}

func NewDeploymentBuilder(writer io.Writer) *DeploymentBuilder {
	return &DeploymentBuilder{
		writer: writer,
	}
}
func (x *DeploymentBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DeploymentBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DeploymentBuilder) SetReplicasDesired(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetDeploymentStrategy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetMaxUnavailable(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetMaxSurge(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetPaused(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x30)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *DeploymentBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DeploymentBuilder) SetReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetUpdatedReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetReadyReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x50)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetAvailableReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x58)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetUnavailableReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x60)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetConditionMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x6a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x82)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DeploymentBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x72)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DeploymentBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x7a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DeploymentBuilder) AddConditions(cb func(w *DeploymentConditionBuilder)) {
	x.buf.Reset()
	x.deploymentConditionBuilder.writer = &x.buf
	x.deploymentConditionBuilder.scratch = x.scratch
	cb(&x.deploymentConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x92)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ReplicaSetConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewReplicaSetConditionBuilder(writer io.Writer) *ReplicaSetConditionBuilder {
	return &ReplicaSetConditionBuilder{
		writer: writer,
	}
}
func (x *ReplicaSetConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ReplicaSetConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ReplicaSetBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	metadataBuilder                 MetadataBuilder
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceRequirementsBuilder     ResourceRequirementsBuilder
	resourceMetricsBuilder          ResourceMetricsBuilder
	replicaSetConditionBuilder      ReplicaSetConditionBuilder
}

func NewReplicaSetBuilder(writer io.Writer) *ReplicaSetBuilder {
	return &ReplicaSetBuilder{
		writer: writer,
	}
}
func (x *ReplicaSetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ReplicaSetBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ReplicaSetBuilder) SetReplicasDesired(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ReplicaSetBuilder) SetReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) SetFullyLabeledReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) SetReadyReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) SetAvailableReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ReplicaSetBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ReplicaSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ReplicaSetBuilder) AddConditions(cb func(w *ReplicaSetConditionBuilder)) {
	x.buf.Reset()
	x.replicaSetConditionBuilder.writer = &x.buf
	x.replicaSetConditionBuilder.scratch = x.scratch
	cb(&x.replicaSetConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x62)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type LabelSelectorRequirementBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewLabelSelectorRequirementBuilder(writer io.Writer) *LabelSelectorRequirementBuilder {
	return &LabelSelectorRequirementBuilder{
		writer: writer,
	}
}
func (x *LabelSelectorRequirementBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *LabelSelectorRequirementBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LabelSelectorRequirementBuilder) SetOperator(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LabelSelectorRequirementBuilder) AddValues(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PodDisruptionBudgetBuilder struct {
	writer                           io.Writer
	buf                              bytes.Buffer
	scratch                          []byte
	metadataBuilder                  MetadataBuilder
	podDisruptionBudgetSpecBuilder   PodDisruptionBudgetSpecBuilder
	podDisruptionBudgetStatusBuilder PodDisruptionBudgetStatusBuilder
}

func NewPodDisruptionBudgetBuilder(writer io.Writer) *PodDisruptionBudgetBuilder {
	return &PodDisruptionBudgetBuilder{
		writer: writer,
	}
}
func (x *PodDisruptionBudgetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PodDisruptionBudgetBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodDisruptionBudgetBuilder) SetSpec(cb func(w *PodDisruptionBudgetSpecBuilder)) {
	x.buf.Reset()
	x.podDisruptionBudgetSpecBuilder.writer = &x.buf
	x.podDisruptionBudgetSpecBuilder.scratch = x.scratch
	cb(&x.podDisruptionBudgetSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodDisruptionBudgetBuilder) SetStatus(cb func(w *PodDisruptionBudgetStatusBuilder)) {
	x.buf.Reset()
	x.podDisruptionBudgetStatusBuilder.writer = &x.buf
	x.podDisruptionBudgetStatusBuilder.scratch = x.scratch
	cb(&x.podDisruptionBudgetStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodDisruptionBudgetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PodDisruptionBudgetSpecBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	intOrStringBuilder              IntOrStringBuilder
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
}

func NewPodDisruptionBudgetSpecBuilder(writer io.Writer) *PodDisruptionBudgetSpecBuilder {
	return &PodDisruptionBudgetSpecBuilder{
		writer: writer,
	}
}
func (x *PodDisruptionBudgetSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PodDisruptionBudgetSpecBuilder) SetMinAvailable(cb func(w *IntOrStringBuilder)) {
	x.buf.Reset()
	x.intOrStringBuilder.writer = &x.buf
	x.intOrStringBuilder.scratch = x.scratch
	cb(&x.intOrStringBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodDisruptionBudgetSpecBuilder) AddSelector(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodDisruptionBudgetSpecBuilder) SetMaxUnavailable(cb func(w *IntOrStringBuilder)) {
	x.buf.Reset()
	x.intOrStringBuilder.writer = &x.buf
	x.intOrStringBuilder.scratch = x.scratch
	cb(&x.intOrStringBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodDisruptionBudgetSpecBuilder) SetUnhealthyPodEvictionPolicy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PodDisruptionBudgetStatusBuilder struct {
	writer                                              io.Writer
	buf                                                 bytes.Buffer
	scratch                                             []byte
	podDisruptionBudgetStatus_DisruptedPodsEntryBuilder PodDisruptionBudgetStatus_DisruptedPodsEntryBuilder
	conditionBuilder                                    ConditionBuilder
}

func NewPodDisruptionBudgetStatusBuilder(writer io.Writer) *PodDisruptionBudgetStatusBuilder {
	return &PodDisruptionBudgetStatusBuilder{
		writer: writer,
	}
}
func (x *PodDisruptionBudgetStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PodDisruptionBudgetStatusBuilder) AddDisruptedPods(cb func(w *PodDisruptionBudgetStatus_DisruptedPodsEntryBuilder)) {
	x.buf.Reset()
	x.podDisruptionBudgetStatus_DisruptedPodsEntryBuilder.writer = &x.buf
	x.podDisruptionBudgetStatus_DisruptedPodsEntryBuilder.scratch = x.scratch
	cb(&x.podDisruptionBudgetStatus_DisruptedPodsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodDisruptionBudgetStatusBuilder) SetDisruptionsAllowed(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodDisruptionBudgetStatusBuilder) SetCurrentHealthy(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodDisruptionBudgetStatusBuilder) SetDesiredHealthy(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodDisruptionBudgetStatusBuilder) SetExpectedPods(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodDisruptionBudgetStatusBuilder) AddConditions(cb func(w *ConditionBuilder)) {
	x.buf.Reset()
	x.conditionBuilder.writer = &x.buf
	x.conditionBuilder.scratch = x.scratch
	cb(&x.conditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PodDisruptionBudgetStatus_DisruptedPodsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPodDisruptionBudgetStatus_DisruptedPodsEntryBuilder(writer io.Writer) *PodDisruptionBudgetStatus_DisruptedPodsEntryBuilder {
	return &PodDisruptionBudgetStatus_DisruptedPodsEntryBuilder{
		writer: writer,
	}
}
func (x *PodDisruptionBudgetStatus_DisruptedPodsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PodDisruptionBudgetStatus_DisruptedPodsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodDisruptionBudgetStatus_DisruptedPodsEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type IntOrStringBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewIntOrStringBuilder(writer io.Writer) *IntOrStringBuilder {
	return &IntOrStringBuilder{
		writer: writer,
	}
}
func (x *IntOrStringBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *IntOrStringBuilder) SetType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *IntOrStringBuilder) SetIntVal(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *IntOrStringBuilder) SetStrVal(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewConditionBuilder(writer io.Writer) *ConditionBuilder {
	return &ConditionBuilder{
		writer: writer,
	}
}
func (x *ConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PodBuilder struct {
	writer                      io.Writer
	buf                         bytes.Buffer
	scratch                     []byte
	metadataBuilder             MetadataBuilder
	containerStatusBuilder      ContainerStatusBuilder
	hostBuilder                 HostBuilder
	resourceRequirementsBuilder ResourceRequirementsBuilder
	resourceMetricsBuilder      ResourceMetricsBuilder
	podConditionBuilder         PodConditionBuilder
	nodeAffinityBuilder         NodeAffinityBuilder
}

func NewPodBuilder(writer io.Writer) *PodBuilder {
	return &PodBuilder{
		writer: writer,
	}
}
func (x *PodBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PodBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) SetIP(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetNominatedNodeName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetNodeName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetPhase(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetRestartCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) AddContainerStatuses(cb func(w *ContainerStatusBuilder)) {
	x.buf.Reset()
	x.containerStatusBuilder.writer = &x.buf
	x.containerStatusBuilder.scratch = x.scratch
	cb(&x.containerStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) AddInitContainerStatuses(cb func(w *ContainerStatusBuilder)) {
	x.buf.Reset()
	x.containerStatusBuilder.writer = &x.buf
	x.containerStatusBuilder.scratch = x.scratch
	cb(&x.containerStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x72)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) SetConditionMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x5a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x62)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x6a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) SetQOSClass(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x7a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetPriorityClass(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x82)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) SetStartTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x90)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetScheduledTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x98)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) AddConditions(cb func(w *PodConditionBuilder)) {
	x.buf.Reset()
	x.podConditionBuilder.writer = &x.buf
	x.podConditionBuilder.scratch = x.scratch
	cb(&x.podConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa2)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) SetNodeAffinity(cb func(w *NodeAffinityBuilder)) {
	x.buf.Reset()
	x.nodeAffinityBuilder.writer = &x.buf
	x.nodeAffinityBuilder.scratch = x.scratch
	cb(&x.nodeAffinityBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xaa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PodConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPodConditionBuilder(writer io.Writer) *PodConditionBuilder {
	return &PodConditionBuilder{
		writer: writer,
	}
}
func (x *PodConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PodConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodConditionBuilder) SetLastProbeTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ContainerStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewContainerStatusBuilder(writer io.Writer) *ContainerStatusBuilder {
	return &ContainerStatusBuilder{
		writer: writer,
	}
}
func (x *ContainerStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ContainerStatusBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerStatusBuilder) SetContainerID(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerStatusBuilder) SetReady(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x18)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *ContainerStatusBuilder) SetRestartCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatusBuilder) SetState(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerStatusBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerStatusBuilder) SetImage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerStatusBuilder) SetImageID(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NodeAffinityBuilder struct {
	writer                         io.Writer
	buf                            bytes.Buffer
	scratch                        []byte
	nodeSelectorBuilder            NodeSelectorBuilder
	preferredSchedulingTermBuilder PreferredSchedulingTermBuilder
}

func NewNodeAffinityBuilder(writer io.Writer) *NodeAffinityBuilder {
	return &NodeAffinityBuilder{
		writer: writer,
	}
}
func (x *NodeAffinityBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NodeAffinityBuilder) SetRequiredDuringSchedulingIgnoredDuringExecution(cb func(w *NodeSelectorBuilder)) {
	x.buf.Reset()
	x.nodeSelectorBuilder.writer = &x.buf
	x.nodeSelectorBuilder.scratch = x.scratch
	cb(&x.nodeSelectorBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeAffinityBuilder) AddPreferredDuringSchedulingIgnoredDuringExecution(cb func(w *PreferredSchedulingTermBuilder)) {
	x.buf.Reset()
	x.preferredSchedulingTermBuilder.writer = &x.buf
	x.preferredSchedulingTermBuilder.scratch = x.scratch
	cb(&x.preferredSchedulingTermBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type NodeSelectorBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	nodeSelectorTermBuilder NodeSelectorTermBuilder
}

func NewNodeSelectorBuilder(writer io.Writer) *NodeSelectorBuilder {
	return &NodeSelectorBuilder{
		writer: writer,
	}
}
func (x *NodeSelectorBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NodeSelectorBuilder) AddNodeSelectorTerms(cb func(w *NodeSelectorTermBuilder)) {
	x.buf.Reset()
	x.nodeSelectorTermBuilder.writer = &x.buf
	x.nodeSelectorTermBuilder.scratch = x.scratch
	cb(&x.nodeSelectorTermBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PreferredSchedulingTermBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	nodeSelectorTermBuilder NodeSelectorTermBuilder
}

func NewPreferredSchedulingTermBuilder(writer io.Writer) *PreferredSchedulingTermBuilder {
	return &PreferredSchedulingTermBuilder{
		writer: writer,
	}
}
func (x *PreferredSchedulingTermBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PreferredSchedulingTermBuilder) SetWeight(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PreferredSchedulingTermBuilder) SetPreference(cb func(w *NodeSelectorTermBuilder)) {
	x.buf.Reset()
	x.nodeSelectorTermBuilder.writer = &x.buf
	x.nodeSelectorTermBuilder.scratch = x.scratch
	cb(&x.nodeSelectorTermBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ManifestBuilder struct {
	writer      io.Writer
	buf         bytes.Buffer
	scratch     []byte
	hostBuilder HostBuilder
}

func NewManifestBuilder(writer io.Writer) *ManifestBuilder {
	return &ManifestBuilder{
		writer: writer,
	}
}
func (x *ManifestBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ManifestBuilder) SetType(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetResourceVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetUid(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetContent(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ManifestBuilder) SetContentType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetIsTerminated(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x40)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *ManifestBuilder) SetApiVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetNodeName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x5a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x62)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type NamespaceConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNamespaceConditionBuilder(writer io.Writer) *NamespaceConditionBuilder {
	return &NamespaceConditionBuilder{
		writer: writer,
	}
}
func (x *NamespaceConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NamespaceConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *NamespaceConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NamespaceBuilder struct {
	writer                    io.Writer
	buf                       bytes.Buffer
	scratch                   []byte
	metadataBuilder           MetadataBuilder
	namespaceConditionBuilder NamespaceConditionBuilder
}

func NewNamespaceBuilder(writer io.Writer) *NamespaceBuilder {
	return &NamespaceBuilder{
		writer: writer,
	}
}
func (x *NamespaceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NamespaceBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NamespaceBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceBuilder) SetConditionMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NamespaceBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceBuilder) AddConditions(cb func(w *NamespaceConditionBuilder)) {
	x.buf.Reset()
	x.namespaceConditionBuilder.writer = &x.buf
	x.namespaceConditionBuilder.scratch = x.scratch
	cb(&x.namespaceConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ResourceRequirementsBuilder struct {
	writer                                    io.Writer
	buf                                       bytes.Buffer
	scratch                                   []byte
	resourceRequirements_LimitsEntryBuilder   ResourceRequirements_LimitsEntryBuilder
	resourceRequirements_RequestsEntryBuilder ResourceRequirements_RequestsEntryBuilder
}

func NewResourceRequirementsBuilder(writer io.Writer) *ResourceRequirementsBuilder {
	return &ResourceRequirementsBuilder{
		writer: writer,
	}
}
func (x *ResourceRequirementsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResourceRequirementsBuilder) AddLimits(cb func(w *ResourceRequirements_LimitsEntryBuilder)) {
	x.buf.Reset()
	x.resourceRequirements_LimitsEntryBuilder.writer = &x.buf
	x.resourceRequirements_LimitsEntryBuilder.scratch = x.scratch
	cb(&x.resourceRequirements_LimitsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ResourceRequirementsBuilder) AddRequests(cb func(w *ResourceRequirements_RequestsEntryBuilder)) {
	x.buf.Reset()
	x.resourceRequirements_RequestsEntryBuilder.writer = &x.buf
	x.resourceRequirements_RequestsEntryBuilder.scratch = x.scratch
	cb(&x.resourceRequirements_RequestsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ResourceRequirementsBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceRequirementsBuilder) SetType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type ResourceRequirements_LimitsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResourceRequirements_LimitsEntryBuilder(writer io.Writer) *ResourceRequirements_LimitsEntryBuilder {
	return &ResourceRequirements_LimitsEntryBuilder{
		writer: writer,
	}
}
func (x *ResourceRequirements_LimitsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResourceRequirements_LimitsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceRequirements_LimitsEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ResourceRequirements_RequestsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResourceRequirements_RequestsEntryBuilder(writer io.Writer) *ResourceRequirements_RequestsEntryBuilder {
	return &ResourceRequirements_RequestsEntryBuilder{
		writer: writer,
	}
}
func (x *ResourceRequirements_RequestsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResourceRequirements_RequestsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceRequirements_RequestsEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ResourceMetricsBuilder struct {
	writer                                   io.Writer
	buf                                      bytes.Buffer
	scratch                                  []byte
	resourceMetrics_MetricValuesEntryBuilder ResourceMetrics_MetricValuesEntryBuilder
}

func NewResourceMetricsBuilder(writer io.Writer) *ResourceMetricsBuilder {
	return &ResourceMetricsBuilder{
		writer: writer,
	}
}
func (x *ResourceMetricsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResourceMetricsBuilder) AddMetricValues(cb func(w *ResourceMetrics_MetricValuesEntryBuilder)) {
	x.buf.Reset()
	x.resourceMetrics_MetricValuesEntryBuilder.writer = &x.buf
	x.resourceMetrics_MetricValuesEntryBuilder.scratch = x.scratch
	cb(&x.resourceMetrics_MetricValuesEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ResourceMetrics_MetricValuesEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResourceMetrics_MetricValuesEntryBuilder(writer io.Writer) *ResourceMetrics_MetricValuesEntryBuilder {
	return &ResourceMetrics_MetricValuesEntryBuilder{
		writer: writer,
	}
}
func (x *ResourceMetrics_MetricValuesEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResourceMetrics_MetricValuesEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceMetrics_MetricValuesEntryBuilder) SetValue(v float64) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed64(x.scratch, math.Float64bits(v))
	x.writer.Write(x.scratch)
}

type JobSpecBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceRequirementsBuilder     ResourceRequirementsBuilder
}

func NewJobSpecBuilder(writer io.Writer) *JobSpecBuilder {
	return &JobSpecBuilder{
		writer: writer,
	}
}
func (x *JobSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *JobSpecBuilder) SetParallelism(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobSpecBuilder) SetCompletions(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobSpecBuilder) SetActiveDeadlineSeconds(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobSpecBuilder) SetBackoffLimit(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobSpecBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *JobSpecBuilder) SetManualSelector(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x30)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *JobSpecBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type JobStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewJobStatusBuilder(writer io.Writer) *JobStatusBuilder {
	return &JobStatusBuilder{
		writer: writer,
	}
}
func (x *JobStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *JobStatusBuilder) SetConditionMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *JobStatusBuilder) SetStartTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobStatusBuilder) SetCompletionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobStatusBuilder) SetActive(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobStatusBuilder) SetSucceeded(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobStatusBuilder) SetFailed(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type JobConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewJobConditionBuilder(writer io.Writer) *JobConditionBuilder {
	return &JobConditionBuilder{
		writer: writer,
	}
}
func (x *JobConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *JobConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *JobConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *JobConditionBuilder) SetLastProbeTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *JobConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type JobBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	metadataBuilder     MetadataBuilder
	jobSpecBuilder      JobSpecBuilder
	jobStatusBuilder    JobStatusBuilder
	jobConditionBuilder JobConditionBuilder
}

func NewJobBuilder(writer io.Writer) *JobBuilder {
	return &JobBuilder{
		writer: writer,
	}
}
func (x *JobBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *JobBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *JobBuilder) SetSpec(cb func(w *JobSpecBuilder)) {
	x.buf.Reset()
	x.jobSpecBuilder.writer = &x.buf
	x.jobSpecBuilder.scratch = x.scratch
	cb(&x.jobSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *JobBuilder) SetStatus(cb func(w *JobStatusBuilder)) {
	x.buf.Reset()
	x.jobStatusBuilder.writer = &x.buf
	x.jobStatusBuilder.scratch = x.scratch
	cb(&x.jobStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *JobBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *JobBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *JobBuilder) AddConditions(cb func(w *JobConditionBuilder)) {
	x.buf.Reset()
	x.jobConditionBuilder.writer = &x.buf
	x.jobConditionBuilder.scratch = x.scratch
	cb(&x.jobConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CronJobSpecBuilder struct {
	writer                      io.Writer
	buf                         bytes.Buffer
	scratch                     []byte
	resourceRequirementsBuilder ResourceRequirementsBuilder
}

func NewCronJobSpecBuilder(writer io.Writer) *CronJobSpecBuilder {
	return &CronJobSpecBuilder{
		writer: writer,
	}
}
func (x *CronJobSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CronJobSpecBuilder) SetSchedule(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CronJobSpecBuilder) SetStartingDeadlineSeconds(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CronJobSpecBuilder) SetConcurrencyPolicy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CronJobSpecBuilder) SetSuspend(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *CronJobSpecBuilder) SetSuccessfulJobsHistoryLimit(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CronJobSpecBuilder) SetFailedJobsHistoryLimit(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CronJobSpecBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CronJobSpecBuilder) SetTimeZone(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CronJobStatusBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	objectReferenceBuilder ObjectReferenceBuilder
}

func NewCronJobStatusBuilder(writer io.Writer) *CronJobStatusBuilder {
	return &CronJobStatusBuilder{
		writer: writer,
	}
}
func (x *CronJobStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CronJobStatusBuilder) AddActive(cb func(w *ObjectReferenceBuilder)) {
	x.buf.Reset()
	x.objectReferenceBuilder.writer = &x.buf
	x.objectReferenceBuilder.scratch = x.scratch
	cb(&x.objectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CronJobStatusBuilder) SetLastScheduleTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CronJobStatusBuilder) SetLastSuccessfulTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type CronJobBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	metadataBuilder      MetadataBuilder
	cronJobSpecBuilder   CronJobSpecBuilder
	cronJobStatusBuilder CronJobStatusBuilder
}

func NewCronJobBuilder(writer io.Writer) *CronJobBuilder {
	return &CronJobBuilder{
		writer: writer,
	}
}
func (x *CronJobBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CronJobBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CronJobBuilder) SetSpec(cb func(w *CronJobSpecBuilder)) {
	x.buf.Reset()
	x.cronJobSpecBuilder.writer = &x.buf
	x.cronJobSpecBuilder.scratch = x.scratch
	cb(&x.cronJobSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CronJobBuilder) SetStatus(cb func(w *CronJobStatusBuilder)) {
	x.buf.Reset()
	x.cronJobStatusBuilder.writer = &x.buf
	x.cronJobStatusBuilder.scratch = x.scratch
	cb(&x.cronJobStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CronJobBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CronJobBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type DaemonSetSpecBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceRequirementsBuilder     ResourceRequirementsBuilder
}

func NewDaemonSetSpecBuilder(writer io.Writer) *DaemonSetSpecBuilder {
	return &DaemonSetSpecBuilder{
		writer: writer,
	}
}
func (x *DaemonSetSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DaemonSetSpecBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetSpecBuilder) SetDeploymentStrategy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetSpecBuilder) SetMaxUnavailable(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetSpecBuilder) SetMinReadySeconds(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetSpecBuilder) SetRevisionHistoryLimit(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetSpecBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type DaemonSetStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewDaemonSetStatusBuilder(writer io.Writer) *DaemonSetStatusBuilder {
	return &DaemonSetStatusBuilder{
		writer: writer,
	}
}
func (x *DaemonSetStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DaemonSetStatusBuilder) SetCurrentNumberScheduled(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetNumberMisscheduled(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetDesiredNumberScheduled(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetNumberReady(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetUpdatedNumberScheduled(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetNumberAvailable(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetNumberUnavailable(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type DaemonSetConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewDaemonSetConditionBuilder(writer io.Writer) *DaemonSetConditionBuilder {
	return &DaemonSetConditionBuilder{
		writer: writer,
	}
}
func (x *DaemonSetConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DaemonSetConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type DaemonSetBuilder struct {
	writer                    io.Writer
	buf                       bytes.Buffer
	scratch                   []byte
	metadataBuilder           MetadataBuilder
	daemonSetSpecBuilder      DaemonSetSpecBuilder
	daemonSetStatusBuilder    DaemonSetStatusBuilder
	resourceMetricsBuilder    ResourceMetricsBuilder
	daemonSetConditionBuilder DaemonSetConditionBuilder
}

func NewDaemonSetBuilder(writer io.Writer) *DaemonSetBuilder {
	return &DaemonSetBuilder{
		writer: writer,
	}
}
func (x *DaemonSetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DaemonSetBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetBuilder) SetSpec(cb func(w *DaemonSetSpecBuilder)) {
	x.buf.Reset()
	x.daemonSetSpecBuilder.writer = &x.buf
	x.daemonSetSpecBuilder.scratch = x.scratch
	cb(&x.daemonSetSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetBuilder) SetStatus(cb func(w *DaemonSetStatusBuilder)) {
	x.buf.Reset()
	x.daemonSetStatusBuilder.writer = &x.buf
	x.daemonSetStatusBuilder.scratch = x.scratch
	cb(&x.daemonSetStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetBuilder) AddConditions(cb func(w *DaemonSetConditionBuilder)) {
	x.buf.Reset()
	x.daemonSetConditionBuilder.writer = &x.buf
	x.daemonSetConditionBuilder.scratch = x.scratch
	cb(&x.daemonSetConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type StatefulSetSpecBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceRequirementsBuilder     ResourceRequirementsBuilder
}

func NewStatefulSetSpecBuilder(writer io.Writer) *StatefulSetSpecBuilder {
	return &StatefulSetSpecBuilder{
		writer: writer,
	}
}
func (x *StatefulSetSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *StatefulSetSpecBuilder) SetDesiredReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetSpecBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetSpecBuilder) SetServiceName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetSpecBuilder) SetPodManagementPolicy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetSpecBuilder) SetUpdateStrategy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetSpecBuilder) SetPartition(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetSpecBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type StatefulSetStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewStatefulSetStatusBuilder(writer io.Writer) *StatefulSetStatusBuilder {
	return &StatefulSetStatusBuilder{
		writer: writer,
	}
}
func (x *StatefulSetStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *StatefulSetStatusBuilder) SetReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetStatusBuilder) SetReadyReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetStatusBuilder) SetCurrentReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetStatusBuilder) SetUpdatedReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type StatefulSetConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewStatefulSetConditionBuilder(writer io.Writer) *StatefulSetConditionBuilder {
	return &StatefulSetConditionBuilder{
		writer: writer,
	}
}
func (x *StatefulSetConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *StatefulSetConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type StatefulSetBuilder struct {
	writer                      io.Writer
	buf                         bytes.Buffer
	scratch                     []byte
	metadataBuilder             MetadataBuilder
	statefulSetSpecBuilder      StatefulSetSpecBuilder
	statefulSetStatusBuilder    StatefulSetStatusBuilder
	resourceMetricsBuilder      ResourceMetricsBuilder
	statefulSetConditionBuilder StatefulSetConditionBuilder
}

func NewStatefulSetBuilder(writer io.Writer) *StatefulSetBuilder {
	return &StatefulSetBuilder{
		writer: writer,
	}
}
func (x *StatefulSetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *StatefulSetBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetBuilder) SetSpec(cb func(w *StatefulSetSpecBuilder)) {
	x.buf.Reset()
	x.statefulSetSpecBuilder.writer = &x.buf
	x.statefulSetSpecBuilder.scratch = x.scratch
	cb(&x.statefulSetSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetBuilder) SetStatus(cb func(w *StatefulSetStatusBuilder)) {
	x.buf.Reset()
	x.statefulSetStatusBuilder.writer = &x.buf
	x.statefulSetStatusBuilder.scratch = x.scratch
	cb(&x.statefulSetStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetBuilder) AddConditions(cb func(w *StatefulSetConditionBuilder)) {
	x.buf.Reset()
	x.statefulSetConditionBuilder.writer = &x.buf
	x.statefulSetConditionBuilder.scratch = x.scratch
	cb(&x.statefulSetConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PersistentVolumeBuilder struct {
	writer                        io.Writer
	buf                           bytes.Buffer
	scratch                       []byte
	metadataBuilder               MetadataBuilder
	persistentVolumeSpecBuilder   PersistentVolumeSpecBuilder
	persistentVolumeStatusBuilder PersistentVolumeStatusBuilder
}

func NewPersistentVolumeBuilder(writer io.Writer) *PersistentVolumeBuilder {
	return &PersistentVolumeBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PersistentVolumeBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeBuilder) SetSpec(cb func(w *PersistentVolumeSpecBuilder)) {
	x.buf.Reset()
	x.persistentVolumeSpecBuilder.writer = &x.buf
	x.persistentVolumeSpecBuilder.scratch = x.scratch
	cb(&x.persistentVolumeSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeBuilder) SetStatus(cb func(w *PersistentVolumeStatusBuilder)) {
	x.buf.Reset()
	x.persistentVolumeStatusBuilder.writer = &x.buf
	x.persistentVolumeStatusBuilder.scratch = x.scratch
	cb(&x.persistentVolumeStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PersistentVolumeSpecBuilder struct {
	writer                                    io.Writer
	buf                                       bytes.Buffer
	scratch                                   []byte
	persistentVolumeSpec_CapacityEntryBuilder PersistentVolumeSpec_CapacityEntryBuilder
	objectReferenceBuilder                    ObjectReferenceBuilder
	nodeSelectorTermBuilder                   NodeSelectorTermBuilder
	persistentVolumeSourceBuilder             PersistentVolumeSourceBuilder
}

func NewPersistentVolumeSpecBuilder(writer io.Writer) *PersistentVolumeSpecBuilder {
	return &PersistentVolumeSpecBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PersistentVolumeSpecBuilder) AddCapacity(cb func(w *PersistentVolumeSpec_CapacityEntryBuilder)) {
	x.buf.Reset()
	x.persistentVolumeSpec_CapacityEntryBuilder.writer = &x.buf
	x.persistentVolumeSpec_CapacityEntryBuilder.scratch = x.scratch
	cb(&x.persistentVolumeSpec_CapacityEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSpecBuilder) SetPersistentVolumeType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) AddAccessModes(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) SetClaimRef(cb func(w *ObjectReferenceBuilder)) {
	x.buf.Reset()
	x.objectReferenceBuilder.writer = &x.buf
	x.objectReferenceBuilder.scratch = x.scratch
	cb(&x.objectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSpecBuilder) SetPersistentVolumeReclaimPolicy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) SetStorageClassName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) AddMountOptions(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) SetVolumeMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) AddNodeAffinity(cb func(w *NodeSelectorTermBuilder)) {
	x.buf.Reset()
	x.nodeSelectorTermBuilder.writer = &x.buf
	x.nodeSelectorTermBuilder.scratch = x.scratch
	cb(&x.nodeSelectorTermBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x4a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSpecBuilder) SetPersistentVolumeSource(cb func(w *PersistentVolumeSourceBuilder)) {
	x.buf.Reset()
	x.persistentVolumeSourceBuilder.writer = &x.buf
	x.persistentVolumeSourceBuilder.scratch = x.scratch
	cb(&x.persistentVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PersistentVolumeSpec_CapacityEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPersistentVolumeSpec_CapacityEntryBuilder(writer io.Writer) *PersistentVolumeSpec_CapacityEntryBuilder {
	return &PersistentVolumeSpec_CapacityEntryBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeSpec_CapacityEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PersistentVolumeSpec_CapacityEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpec_CapacityEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type PersistentVolumeSourceBuilder struct {
	writer                                  io.Writer
	buf                                     bytes.Buffer
	scratch                                 []byte
	gCEPersistentDiskVolumeSourceBuilder    GCEPersistentDiskVolumeSourceBuilder
	aWSElasticBlockStoreVolumeSourceBuilder AWSElasticBlockStoreVolumeSourceBuilder
	azureFilePersistentVolumeSourceBuilder  AzureFilePersistentVolumeSourceBuilder
	azureDiskVolumeSourceBuilder            AzureDiskVolumeSourceBuilder
	cSIVolumeSourceBuilder                  CSIVolumeSourceBuilder
}

func NewPersistentVolumeSourceBuilder(writer io.Writer) *PersistentVolumeSourceBuilder {
	return &PersistentVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeSourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PersistentVolumeSourceBuilder) SetGcePersistentDisk(cb func(w *GCEPersistentDiskVolumeSourceBuilder)) {
	x.buf.Reset()
	x.gCEPersistentDiskVolumeSourceBuilder.writer = &x.buf
	x.gCEPersistentDiskVolumeSourceBuilder.scratch = x.scratch
	cb(&x.gCEPersistentDiskVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSourceBuilder) SetAwsElasticBlockStore(cb func(w *AWSElasticBlockStoreVolumeSourceBuilder)) {
	x.buf.Reset()
	x.aWSElasticBlockStoreVolumeSourceBuilder.writer = &x.buf
	x.aWSElasticBlockStoreVolumeSourceBuilder.scratch = x.scratch
	cb(&x.aWSElasticBlockStoreVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSourceBuilder) SetAzureFile(cb func(w *AzureFilePersistentVolumeSourceBuilder)) {
	x.buf.Reset()
	x.azureFilePersistentVolumeSourceBuilder.writer = &x.buf
	x.azureFilePersistentVolumeSourceBuilder.scratch = x.scratch
	cb(&x.azureFilePersistentVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSourceBuilder) SetAzureDisk(cb func(w *AzureDiskVolumeSourceBuilder)) {
	x.buf.Reset()
	x.azureDiskVolumeSourceBuilder.writer = &x.buf
	x.azureDiskVolumeSourceBuilder.scratch = x.scratch
	cb(&x.azureDiskVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSourceBuilder) SetCsi(cb func(w *CSIVolumeSourceBuilder)) {
	x.buf.Reset()
	x.cSIVolumeSourceBuilder.writer = &x.buf
	x.cSIVolumeSourceBuilder.scratch = x.scratch
	cb(&x.cSIVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type GCEPersistentDiskVolumeSourceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewGCEPersistentDiskVolumeSourceBuilder(writer io.Writer) *GCEPersistentDiskVolumeSourceBuilder {
	return &GCEPersistentDiskVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *GCEPersistentDiskVolumeSourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *GCEPersistentDiskVolumeSourceBuilder) SetPdName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *GCEPersistentDiskVolumeSourceBuilder) SetFsType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *GCEPersistentDiskVolumeSourceBuilder) SetPartition(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *GCEPersistentDiskVolumeSourceBuilder) SetReadOnly(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}

type AWSElasticBlockStoreVolumeSourceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewAWSElasticBlockStoreVolumeSourceBuilder(writer io.Writer) *AWSElasticBlockStoreVolumeSourceBuilder {
	return &AWSElasticBlockStoreVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *AWSElasticBlockStoreVolumeSourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *AWSElasticBlockStoreVolumeSourceBuilder) SetVolumeID(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AWSElasticBlockStoreVolumeSourceBuilder) SetFsType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AWSElasticBlockStoreVolumeSourceBuilder) SetPartition(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *AWSElasticBlockStoreVolumeSourceBuilder) SetReadOnly(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}

type AzureFilePersistentVolumeSourceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewAzureFilePersistentVolumeSourceBuilder(writer io.Writer) *AzureFilePersistentVolumeSourceBuilder {
	return &AzureFilePersistentVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *AzureFilePersistentVolumeSourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *AzureFilePersistentVolumeSourceBuilder) SetSecretName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureFilePersistentVolumeSourceBuilder) SetShareName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureFilePersistentVolumeSourceBuilder) SetReadOnly(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x18)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *AzureFilePersistentVolumeSourceBuilder) SetSecretNamespace(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type AzureDiskVolumeSourceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewAzureDiskVolumeSourceBuilder(writer io.Writer) *AzureDiskVolumeSourceBuilder {
	return &AzureDiskVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *AzureDiskVolumeSourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *AzureDiskVolumeSourceBuilder) SetDiskName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureDiskVolumeSourceBuilder) SetDiskURI(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureDiskVolumeSourceBuilder) SetCachingMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureDiskVolumeSourceBuilder) SetFsType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureDiskVolumeSourceBuilder) SetReadOnly(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x28)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *AzureDiskVolumeSourceBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CSIVolumeSourceBuilder struct {
	writer                                       io.Writer
	buf                                          bytes.Buffer
	scratch                                      []byte
	cSIVolumeSource_VolumeAttributesEntryBuilder CSIVolumeSource_VolumeAttributesEntryBuilder
	secretReferenceBuilder                       SecretReferenceBuilder
}

func NewCSIVolumeSourceBuilder(writer io.Writer) *CSIVolumeSourceBuilder {
	return &CSIVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *CSIVolumeSourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CSIVolumeSourceBuilder) SetDriver(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CSIVolumeSourceBuilder) SetVolumeHandle(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CSIVolumeSourceBuilder) SetReadOnly(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x18)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *CSIVolumeSourceBuilder) SetFsType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CSIVolumeSourceBuilder) AddVolumeAttributes(cb func(w *CSIVolumeSource_VolumeAttributesEntryBuilder)) {
	x.buf.Reset()
	x.cSIVolumeSource_VolumeAttributesEntryBuilder.writer = &x.buf
	x.cSIVolumeSource_VolumeAttributesEntryBuilder.scratch = x.scratch
	cb(&x.cSIVolumeSource_VolumeAttributesEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CSIVolumeSourceBuilder) SetControllerPublishSecretRef(cb func(w *SecretReferenceBuilder)) {
	x.buf.Reset()
	x.secretReferenceBuilder.writer = &x.buf
	x.secretReferenceBuilder.scratch = x.scratch
	cb(&x.secretReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CSIVolumeSourceBuilder) SetNodeStageSecretRef(cb func(w *SecretReferenceBuilder)) {
	x.buf.Reset()
	x.secretReferenceBuilder.writer = &x.buf
	x.secretReferenceBuilder.scratch = x.scratch
	cb(&x.secretReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CSIVolumeSourceBuilder) SetNodePublishSecretRef(cb func(w *SecretReferenceBuilder)) {
	x.buf.Reset()
	x.secretReferenceBuilder.writer = &x.buf
	x.secretReferenceBuilder.scratch = x.scratch
	cb(&x.secretReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CSIVolumeSourceBuilder) SetControllerExpandSecretRef(cb func(w *SecretReferenceBuilder)) {
	x.buf.Reset()
	x.secretReferenceBuilder.writer = &x.buf
	x.secretReferenceBuilder.scratch = x.scratch
	cb(&x.secretReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x4a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CSIVolumeSourceBuilder) SetNodeExpandSecretRef(cb func(w *SecretReferenceBuilder)) {
	x.buf.Reset()
	x.secretReferenceBuilder.writer = &x.buf
	x.secretReferenceBuilder.scratch = x.scratch
	cb(&x.secretReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CSIVolumeSource_VolumeAttributesEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCSIVolumeSource_VolumeAttributesEntryBuilder(writer io.Writer) *CSIVolumeSource_VolumeAttributesEntryBuilder {
	return &CSIVolumeSource_VolumeAttributesEntryBuilder{
		writer: writer,
	}
}
func (x *CSIVolumeSource_VolumeAttributesEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CSIVolumeSource_VolumeAttributesEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CSIVolumeSource_VolumeAttributesEntryBuilder) SetValue(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type SecretReferenceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewSecretReferenceBuilder(writer io.Writer) *SecretReferenceBuilder {
	return &SecretReferenceBuilder{
		writer: writer,
	}
}
func (x *SecretReferenceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *SecretReferenceBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SecretReferenceBuilder) SetNamespace(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PersistentVolumeStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPersistentVolumeStatusBuilder(writer io.Writer) *PersistentVolumeStatusBuilder {
	return &PersistentVolumeStatusBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PersistentVolumeStatusBuilder) SetPhase(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeStatusBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeStatusBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NodeSelectorTermBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
}

func NewNodeSelectorTermBuilder(writer io.Writer) *NodeSelectorTermBuilder {
	return &NodeSelectorTermBuilder{
		writer: writer,
	}
}
func (x *NodeSelectorTermBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NodeSelectorTermBuilder) AddMatchExpressions(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeSelectorTermBuilder) AddMatchFields(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PersistentVolumeClaimBuilder struct {
	writer                             io.Writer
	buf                                bytes.Buffer
	scratch                            []byte
	metadataBuilder                    MetadataBuilder
	persistentVolumeClaimSpecBuilder   PersistentVolumeClaimSpecBuilder
	persistentVolumeClaimStatusBuilder PersistentVolumeClaimStatusBuilder
}

func NewPersistentVolumeClaimBuilder(writer io.Writer) *PersistentVolumeClaimBuilder {
	return &PersistentVolumeClaimBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeClaimBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PersistentVolumeClaimBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimBuilder) SetSpec(cb func(w *PersistentVolumeClaimSpecBuilder)) {
	x.buf.Reset()
	x.persistentVolumeClaimSpecBuilder.writer = &x.buf
	x.persistentVolumeClaimSpecBuilder.scratch = x.scratch
	cb(&x.persistentVolumeClaimSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimBuilder) SetStatus(cb func(w *PersistentVolumeClaimStatusBuilder)) {
	x.buf.Reset()
	x.persistentVolumeClaimStatusBuilder.writer = &x.buf
	x.persistentVolumeClaimStatusBuilder.scratch = x.scratch
	cb(&x.persistentVolumeClaimStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PersistentVolumeClaimStatusBuilder struct {
	writer                                           io.Writer
	buf                                              bytes.Buffer
	scratch                                          []byte
	persistentVolumeClaimStatus_CapacityEntryBuilder PersistentVolumeClaimStatus_CapacityEntryBuilder
	persistentVolumeClaimConditionBuilder            PersistentVolumeClaimConditionBuilder
}

func NewPersistentVolumeClaimStatusBuilder(writer io.Writer) *PersistentVolumeClaimStatusBuilder {
	return &PersistentVolumeClaimStatusBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeClaimStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PersistentVolumeClaimStatusBuilder) SetPhase(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimStatusBuilder) AddAccessModes(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimStatusBuilder) AddCapacity(cb func(w *PersistentVolumeClaimStatus_CapacityEntryBuilder)) {
	x.buf.Reset()
	x.persistentVolumeClaimStatus_CapacityEntryBuilder.writer = &x.buf
	x.persistentVolumeClaimStatus_CapacityEntryBuilder.scratch = x.scratch
	cb(&x.persistentVolumeClaimStatus_CapacityEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimStatusBuilder) AddConditions(cb func(w *PersistentVolumeClaimConditionBuilder)) {
	x.buf.Reset()
	x.persistentVolumeClaimConditionBuilder.writer = &x.buf
	x.persistentVolumeClaimConditionBuilder.scratch = x.scratch
	cb(&x.persistentVolumeClaimConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PersistentVolumeClaimStatus_CapacityEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPersistentVolumeClaimStatus_CapacityEntryBuilder(writer io.Writer) *PersistentVolumeClaimStatus_CapacityEntryBuilder {
	return &PersistentVolumeClaimStatus_CapacityEntryBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeClaimStatus_CapacityEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PersistentVolumeClaimStatus_CapacityEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimStatus_CapacityEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type PersistentVolumeClaimSpecBuilder struct {
	writer                           io.Writer
	buf                              bytes.Buffer
	scratch                          []byte
	resourceRequirementsBuilder      ResourceRequirementsBuilder
	labelSelectorRequirementBuilder  LabelSelectorRequirementBuilder
	typedLocalObjectReferenceBuilder TypedLocalObjectReferenceBuilder
}

func NewPersistentVolumeClaimSpecBuilder(writer io.Writer) *PersistentVolumeClaimSpecBuilder {
	return &PersistentVolumeClaimSpecBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeClaimSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PersistentVolumeClaimSpecBuilder) AddAccessModes(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimSpecBuilder) SetResources(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimSpecBuilder) SetVolumeName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimSpecBuilder) AddSelector(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimSpecBuilder) SetStorageClassName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimSpecBuilder) SetVolumeMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimSpecBuilder) SetDataSource(cb func(w *TypedLocalObjectReferenceBuilder)) {
	x.buf.Reset()
	x.typedLocalObjectReferenceBuilder.writer = &x.buf
	x.typedLocalObjectReferenceBuilder.scratch = x.scratch
	cb(&x.typedLocalObjectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type TypedLocalObjectReferenceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewTypedLocalObjectReferenceBuilder(writer io.Writer) *TypedLocalObjectReferenceBuilder {
	return &TypedLocalObjectReferenceBuilder{
		writer: writer,
	}
}
func (x *TypedLocalObjectReferenceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *TypedLocalObjectReferenceBuilder) SetApiGroup(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TypedLocalObjectReferenceBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TypedLocalObjectReferenceBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PersistentVolumeClaimConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPersistentVolumeClaimConditionBuilder(writer io.Writer) *PersistentVolumeClaimConditionBuilder {
	return &PersistentVolumeClaimConditionBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeClaimConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PersistentVolumeClaimConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimConditionBuilder) SetLastProbeTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PolicyRuleBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPolicyRuleBuilder(writer io.Writer) *PolicyRuleBuilder {
	return &PolicyRuleBuilder{
		writer: writer,
	}
}
func (x *PolicyRuleBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PolicyRuleBuilder) AddVerbs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PolicyRuleBuilder) AddApiGroups(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PolicyRuleBuilder) AddResources(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PolicyRuleBuilder) AddResourceNames(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PolicyRuleBuilder) AddNonResourceURLs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type SubjectBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewSubjectBuilder(writer io.Writer) *SubjectBuilder {
	return &SubjectBuilder{
		writer: writer,
	}
}
func (x *SubjectBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *SubjectBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SubjectBuilder) SetApiGroup(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SubjectBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SubjectBuilder) SetNamespace(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type RoleBuilder struct {
	writer            io.Writer
	buf               bytes.Buffer
	scratch           []byte
	metadataBuilder   MetadataBuilder
	policyRuleBuilder PolicyRuleBuilder
}

func NewRoleBuilder(writer io.Writer) *RoleBuilder {
	return &RoleBuilder{
		writer: writer,
	}
}
func (x *RoleBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *RoleBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBuilder) AddRules(cb func(w *PolicyRuleBuilder)) {
	x.buf.Reset()
	x.policyRuleBuilder.writer = &x.buf
	x.policyRuleBuilder.scratch = x.scratch
	cb(&x.policyRuleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type RoleBindingBuilder struct {
	writer                           io.Writer
	buf                              bytes.Buffer
	scratch                          []byte
	metadataBuilder                  MetadataBuilder
	subjectBuilder                   SubjectBuilder
	typedLocalObjectReferenceBuilder TypedLocalObjectReferenceBuilder
}

func NewRoleBindingBuilder(writer io.Writer) *RoleBindingBuilder {
	return &RoleBindingBuilder{
		writer: writer,
	}
}
func (x *RoleBindingBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *RoleBindingBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBindingBuilder) AddSubjects(cb func(w *SubjectBuilder)) {
	x.buf.Reset()
	x.subjectBuilder.writer = &x.buf
	x.subjectBuilder.scratch = x.scratch
	cb(&x.subjectBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBindingBuilder) SetRoleRef(cb func(w *TypedLocalObjectReferenceBuilder)) {
	x.buf.Reset()
	x.typedLocalObjectReferenceBuilder.writer = &x.buf
	x.typedLocalObjectReferenceBuilder.scratch = x.scratch
	cb(&x.typedLocalObjectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBindingBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBindingBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ClusterRoleBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	metadataBuilder                 MetadataBuilder
	policyRuleBuilder               PolicyRuleBuilder
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceMetricsBuilder          ResourceMetricsBuilder
}

func NewClusterRoleBuilder(writer io.Writer) *ClusterRoleBuilder {
	return &ClusterRoleBuilder{
		writer: writer,
	}
}
func (x *ClusterRoleBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ClusterRoleBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBuilder) AddRules(cb func(w *PolicyRuleBuilder)) {
	x.buf.Reset()
	x.policyRuleBuilder.writer = &x.buf
	x.policyRuleBuilder.scratch = x.scratch
	cb(&x.policyRuleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBuilder) AddAggregationRules(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterRoleBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ClusterRoleBindingBuilder struct {
	writer                           io.Writer
	buf                              bytes.Buffer
	scratch                          []byte
	metadataBuilder                  MetadataBuilder
	subjectBuilder                   SubjectBuilder
	typedLocalObjectReferenceBuilder TypedLocalObjectReferenceBuilder
}

func NewClusterRoleBindingBuilder(writer io.Writer) *ClusterRoleBindingBuilder {
	return &ClusterRoleBindingBuilder{
		writer: writer,
	}
}
func (x *ClusterRoleBindingBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ClusterRoleBindingBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBindingBuilder) AddSubjects(cb func(w *SubjectBuilder)) {
	x.buf.Reset()
	x.subjectBuilder.writer = &x.buf
	x.subjectBuilder.scratch = x.scratch
	cb(&x.subjectBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBindingBuilder) SetRoleRef(cb func(w *TypedLocalObjectReferenceBuilder)) {
	x.buf.Reset()
	x.typedLocalObjectReferenceBuilder.writer = &x.buf
	x.typedLocalObjectReferenceBuilder.scratch = x.scratch
	cb(&x.typedLocalObjectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBindingBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBindingBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ServiceAccountBuilder struct {
	writer                           io.Writer
	buf                              bytes.Buffer
	scratch                          []byte
	metadataBuilder                  MetadataBuilder
	objectReferenceBuilder           ObjectReferenceBuilder
	typedLocalObjectReferenceBuilder TypedLocalObjectReferenceBuilder
}

func NewServiceAccountBuilder(writer io.Writer) *ServiceAccountBuilder {
	return &ServiceAccountBuilder{
		writer: writer,
	}
}
func (x *ServiceAccountBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ServiceAccountBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceAccountBuilder) AddSecrets(cb func(w *ObjectReferenceBuilder)) {
	x.buf.Reset()
	x.objectReferenceBuilder.writer = &x.buf
	x.objectReferenceBuilder.scratch = x.scratch
	cb(&x.objectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceAccountBuilder) AddImagePullSecrets(cb func(w *TypedLocalObjectReferenceBuilder)) {
	x.buf.Reset()
	x.typedLocalObjectReferenceBuilder.writer = &x.buf
	x.typedLocalObjectReferenceBuilder.scratch = x.scratch
	cb(&x.typedLocalObjectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceAccountBuilder) SetAutomountServiceAccountToken(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *ServiceAccountBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceAccountBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type IngressServiceBackendBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewIngressServiceBackendBuilder(writer io.Writer) *IngressServiceBackendBuilder {
	return &IngressServiceBackendBuilder{
		writer: writer,
	}
}
func (x *IngressServiceBackendBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *IngressServiceBackendBuilder) SetServiceName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *IngressServiceBackendBuilder) SetPortName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *IngressServiceBackendBuilder) SetPortNumber(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type IngressBackendBuilder struct {
	writer                           io.Writer
	buf                              bytes.Buffer
	scratch                          []byte
	ingressServiceBackendBuilder     IngressServiceBackendBuilder
	typedLocalObjectReferenceBuilder TypedLocalObjectReferenceBuilder
}

func NewIngressBackendBuilder(writer io.Writer) *IngressBackendBuilder {
	return &IngressBackendBuilder{
		writer: writer,
	}
}
func (x *IngressBackendBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *IngressBackendBuilder) SetService(cb func(w *IngressServiceBackendBuilder)) {
	x.buf.Reset()
	x.ingressServiceBackendBuilder.writer = &x.buf
	x.ingressServiceBackendBuilder.scratch = x.scratch
	cb(&x.ingressServiceBackendBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressBackendBuilder) SetResource(cb func(w *TypedLocalObjectReferenceBuilder)) {
	x.buf.Reset()
	x.typedLocalObjectReferenceBuilder.writer = &x.buf
	x.typedLocalObjectReferenceBuilder.scratch = x.scratch
	cb(&x.typedLocalObjectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type IngressTLSBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewIngressTLSBuilder(writer io.Writer) *IngressTLSBuilder {
	return &IngressTLSBuilder{
		writer: writer,
	}
}
func (x *IngressTLSBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *IngressTLSBuilder) AddHosts(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *IngressTLSBuilder) SetSecretName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type HTTPIngressPathBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	ingressBackendBuilder IngressBackendBuilder
}

func NewHTTPIngressPathBuilder(writer io.Writer) *HTTPIngressPathBuilder {
	return &HTTPIngressPathBuilder{
		writer: writer,
	}
}
func (x *HTTPIngressPathBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HTTPIngressPathBuilder) SetPath(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HTTPIngressPathBuilder) SetPathType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HTTPIngressPathBuilder) SetBackend(cb func(w *IngressBackendBuilder)) {
	x.buf.Reset()
	x.ingressBackendBuilder.writer = &x.buf
	x.ingressBackendBuilder.scratch = x.scratch
	cb(&x.ingressBackendBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type IngressRuleBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	hTTPIngressPathBuilder HTTPIngressPathBuilder
}

func NewIngressRuleBuilder(writer io.Writer) *IngressRuleBuilder {
	return &IngressRuleBuilder{
		writer: writer,
	}
}
func (x *IngressRuleBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *IngressRuleBuilder) SetHost(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *IngressRuleBuilder) AddHttpPaths(cb func(w *HTTPIngressPathBuilder)) {
	x.buf.Reset()
	x.hTTPIngressPathBuilder.writer = &x.buf
	x.hTTPIngressPathBuilder.scratch = x.scratch
	cb(&x.hTTPIngressPathBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type IngressSpecBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	ingressBackendBuilder IngressBackendBuilder
	ingressTLSBuilder     IngressTLSBuilder
	ingressRuleBuilder    IngressRuleBuilder
}

func NewIngressSpecBuilder(writer io.Writer) *IngressSpecBuilder {
	return &IngressSpecBuilder{
		writer: writer,
	}
}
func (x *IngressSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *IngressSpecBuilder) SetDefaultBackend(cb func(w *IngressBackendBuilder)) {
	x.buf.Reset()
	x.ingressBackendBuilder.writer = &x.buf
	x.ingressBackendBuilder.scratch = x.scratch
	cb(&x.ingressBackendBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressSpecBuilder) AddTls(cb func(w *IngressTLSBuilder)) {
	x.buf.Reset()
	x.ingressTLSBuilder.writer = &x.buf
	x.ingressTLSBuilder.scratch = x.scratch
	cb(&x.ingressTLSBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressSpecBuilder) AddRules(cb func(w *IngressRuleBuilder)) {
	x.buf.Reset()
	x.ingressRuleBuilder.writer = &x.buf
	x.ingressRuleBuilder.scratch = x.scratch
	cb(&x.ingressRuleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressSpecBuilder) SetIngressClassName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PortStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPortStatusBuilder(writer io.Writer) *PortStatusBuilder {
	return &PortStatusBuilder{
		writer: writer,
	}
}
func (x *PortStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PortStatusBuilder) SetPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PortStatusBuilder) SetProtocol(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PortStatusBuilder) SetError(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type LoadBalancerIngressBuilder struct {
	writer            io.Writer
	buf               bytes.Buffer
	scratch           []byte
	portStatusBuilder PortStatusBuilder
}

func NewLoadBalancerIngressBuilder(writer io.Writer) *LoadBalancerIngressBuilder {
	return &LoadBalancerIngressBuilder{
		writer: writer,
	}
}
func (x *LoadBalancerIngressBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *LoadBalancerIngressBuilder) SetIp(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LoadBalancerIngressBuilder) SetHostname(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LoadBalancerIngressBuilder) AddPorts(cb func(w *PortStatusBuilder)) {
	x.buf.Reset()
	x.portStatusBuilder.writer = &x.buf
	x.portStatusBuilder.scratch = x.scratch
	cb(&x.portStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type IngressStatusBuilder struct {
	writer                     io.Writer
	buf                        bytes.Buffer
	scratch                    []byte
	loadBalancerIngressBuilder LoadBalancerIngressBuilder
}

func NewIngressStatusBuilder(writer io.Writer) *IngressStatusBuilder {
	return &IngressStatusBuilder{
		writer: writer,
	}
}
func (x *IngressStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *IngressStatusBuilder) AddIngress(cb func(w *LoadBalancerIngressBuilder)) {
	x.buf.Reset()
	x.loadBalancerIngressBuilder.writer = &x.buf
	x.loadBalancerIngressBuilder.scratch = x.scratch
	cb(&x.loadBalancerIngressBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type IngressBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	metadataBuilder      MetadataBuilder
	ingressSpecBuilder   IngressSpecBuilder
	ingressStatusBuilder IngressStatusBuilder
}

func NewIngressBuilder(writer io.Writer) *IngressBuilder {
	return &IngressBuilder{
		writer: writer,
	}
}
func (x *IngressBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *IngressBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressBuilder) SetSpec(cb func(w *IngressSpecBuilder)) {
	x.buf.Reset()
	x.ingressSpecBuilder.writer = &x.buf
	x.ingressSpecBuilder.scratch = x.scratch
	cb(&x.ingressSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressBuilder) SetStatus(cb func(w *IngressStatusBuilder)) {
	x.buf.Reset()
	x.ingressStatusBuilder.writer = &x.buf
	x.ingressStatusBuilder.scratch = x.scratch
	cb(&x.ingressStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type KafkaStatsBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewKafkaStatsBuilder(writer io.Writer) *KafkaStatsBuilder {
	return &KafkaStatsBuilder{
		writer: writer,
	}
}
func (x *KafkaStatsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *KafkaStatsBuilder) SetCount(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *KafkaStatsBuilder) SetLatencies(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *KafkaStatsBuilder) SetFirstLatencySample(v float64) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x19)
	x.scratch = protowire.AppendFixed64(x.scratch, math.Float64bits(v))
	x.writer.Write(x.scratch)
}

type KafkaRequestHeaderBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewKafkaRequestHeaderBuilder(writer io.Writer) *KafkaRequestHeaderBuilder {
	return &KafkaRequestHeaderBuilder{
		writer: writer,
	}
}
func (x *KafkaRequestHeaderBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *KafkaRequestHeaderBuilder) SetRequest_type(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *KafkaRequestHeaderBuilder) SetRequest_version(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type KafkaAggregationBuilder struct {
	writer                                        io.Writer
	buf                                           bytes.Buffer
	scratch                                       []byte
	kafkaRequestHeaderBuilder                     KafkaRequestHeaderBuilder
	kafkaAggregation_StatsByErrorCodeEntryBuilder KafkaAggregation_StatsByErrorCodeEntryBuilder
}

func NewKafkaAggregationBuilder(writer io.Writer) *KafkaAggregationBuilder {
	return &KafkaAggregationBuilder{
		writer: writer,
	}
}
func (x *KafkaAggregationBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *KafkaAggregationBuilder) SetHeader(cb func(w *KafkaRequestHeaderBuilder)) {
	x.buf.Reset()
	x.kafkaRequestHeaderBuilder.writer = &x.buf
	x.kafkaRequestHeaderBuilder.scratch = x.scratch
	cb(&x.kafkaRequestHeaderBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *KafkaAggregationBuilder) SetTopic(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *KafkaAggregationBuilder) AddStatsByErrorCode(cb func(w *KafkaAggregation_StatsByErrorCodeEntryBuilder)) {
	x.buf.Reset()
	x.kafkaAggregation_StatsByErrorCodeEntryBuilder.writer = &x.buf
	x.kafkaAggregation_StatsByErrorCodeEntryBuilder.scratch = x.scratch
	cb(&x.kafkaAggregation_StatsByErrorCodeEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *KafkaAggregationBuilder) SetCount(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type KafkaAggregation_StatsByErrorCodeEntryBuilder struct {
	writer            io.Writer
	buf               bytes.Buffer
	scratch           []byte
	kafkaStatsBuilder KafkaStatsBuilder
}

func NewKafkaAggregation_StatsByErrorCodeEntryBuilder(writer io.Writer) *KafkaAggregation_StatsByErrorCodeEntryBuilder {
	return &KafkaAggregation_StatsByErrorCodeEntryBuilder{
		writer: writer,
	}
}
func (x *KafkaAggregation_StatsByErrorCodeEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *KafkaAggregation_StatsByErrorCodeEntryBuilder) SetKey(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *KafkaAggregation_StatsByErrorCodeEntryBuilder) SetValue(cb func(w *KafkaStatsBuilder)) {
	x.buf.Reset()
	x.kafkaStatsBuilder.writer = &x.buf
	x.kafkaStatsBuilder.scratch = x.scratch
	cb(&x.kafkaStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type DataStreamsAggregationsBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	kafkaAggregationBuilder KafkaAggregationBuilder
}

func NewDataStreamsAggregationsBuilder(writer io.Writer) *DataStreamsAggregationsBuilder {
	return &DataStreamsAggregationsBuilder{
		writer: writer,
	}
}
func (x *DataStreamsAggregationsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DataStreamsAggregationsBuilder) AddKafkaAggregations(cb func(w *KafkaAggregationBuilder)) {
	x.buf.Reset()
	x.kafkaAggregationBuilder.writer = &x.buf
	x.kafkaAggregationBuilder.scratch = x.scratch
	cb(&x.kafkaAggregationBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PostgresStatsBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPostgresStatsBuilder(writer io.Writer) *PostgresStatsBuilder {
	return &PostgresStatsBuilder{
		writer: writer,
	}
}
func (x *PostgresStatsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PostgresStatsBuilder) SetTableName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PostgresStatsBuilder) SetOperation(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x10)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *PostgresStatsBuilder) SetLatencies(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PostgresStatsBuilder) SetFirstLatencySample(v float64) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x21)
	x.scratch = protowire.AppendFixed64(x.scratch, math.Float64bits(v))
	x.writer.Write(x.scratch)
}
func (x *PostgresStatsBuilder) SetCount(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type RedisStatsBuilder struct {
	writer                              io.Writer
	buf                                 bytes.Buffer
	scratch                             []byte
	redisStats_ErrorToStatsEntryBuilder RedisStats_ErrorToStatsEntryBuilder
}

func NewRedisStatsBuilder(writer io.Writer) *RedisStatsBuilder {
	return &RedisStatsBuilder{
		writer: writer,
	}
}
func (x *RedisStatsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *RedisStatsBuilder) SetCommand(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *RedisStatsBuilder) SetKeyName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *RedisStatsBuilder) SetTruncated(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x18)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *RedisStatsBuilder) AddErrorToStats(cb func(w *RedisStats_ErrorToStatsEntryBuilder)) {
	x.buf.Reset()
	x.redisStats_ErrorToStatsEntryBuilder.writer = &x.buf
	x.redisStats_ErrorToStatsEntryBuilder.scratch = x.scratch
	cb(&x.redisStats_ErrorToStatsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type RedisStats_ErrorToStatsEntryBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	redisStatsEntryBuilder RedisStatsEntryBuilder
}

func NewRedisStats_ErrorToStatsEntryBuilder(writer io.Writer) *RedisStats_ErrorToStatsEntryBuilder {
	return &RedisStats_ErrorToStatsEntryBuilder{
		writer: writer,
	}
}
func (x *RedisStats_ErrorToStatsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *RedisStats_ErrorToStatsEntryBuilder) SetKey(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *RedisStats_ErrorToStatsEntryBuilder) SetValue(cb func(w *RedisStatsEntryBuilder)) {
	x.buf.Reset()
	x.redisStatsEntryBuilder.writer = &x.buf
	x.redisStatsEntryBuilder.scratch = x.scratch
	cb(&x.redisStatsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type RedisStatsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewRedisStatsEntryBuilder(writer io.Writer) *RedisStatsEntryBuilder {
	return &RedisStatsEntryBuilder{
		writer: writer,
	}
}
func (x *RedisStatsEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *RedisStatsEntryBuilder) SetLatencies(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RedisStatsEntryBuilder) SetFirstLatencySample(v float64) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed64(x.scratch, math.Float64bits(v))
	x.writer.Write(x.scratch)
}
func (x *RedisStatsEntryBuilder) SetCount(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type DatabaseStatsBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	postgresStatsBuilder PostgresStatsBuilder
	redisStatsBuilder    RedisStatsBuilder
}

func NewDatabaseStatsBuilder(writer io.Writer) *DatabaseStatsBuilder {
	return &DatabaseStatsBuilder{
		writer: writer,
	}
}
func (x *DatabaseStatsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DatabaseStatsBuilder) SetPostgres(cb func(w *PostgresStatsBuilder)) {
	x.buf.Reset()
	x.postgresStatsBuilder.writer = &x.buf
	x.postgresStatsBuilder.scratch = x.scratch
	cb(&x.postgresStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DatabaseStatsBuilder) SetRedis(cb func(w *RedisStatsBuilder)) {
	x.buf.Reset()
	x.redisStatsBuilder.writer = &x.buf
	x.redisStatsBuilder.scratch = x.scratch
	cb(&x.redisStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type DatabaseAggregationsBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	databaseStatsBuilder DatabaseStatsBuilder
}

func NewDatabaseAggregationsBuilder(writer io.Writer) *DatabaseAggregationsBuilder {
	return &DatabaseAggregationsBuilder{
		writer: writer,
	}
}
func (x *DatabaseAggregationsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DatabaseAggregationsBuilder) AddAggregations(cb func(w *DatabaseStatsBuilder)) {
	x.buf.Reset()
	x.databaseStatsBuilder.writer = &x.buf
	x.databaseStatsBuilder.scratch = x.scratch
	cb(&x.databaseStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HTTPAggregationsBuilder struct {
	writer           io.Writer
	buf              bytes.Buffer
	scratch          []byte
	hTTPStatsBuilder HTTPStatsBuilder
}

func NewHTTPAggregationsBuilder(writer io.Writer) *HTTPAggregationsBuilder {
	return &HTTPAggregationsBuilder{
		writer: writer,
	}
}
func (x *HTTPAggregationsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HTTPAggregationsBuilder) AddEndpointAggregations(cb func(w *HTTPStatsBuilder)) {
	x.buf.Reset()
	x.hTTPStatsBuilder.writer = &x.buf
	x.hTTPStatsBuilder.scratch = x.scratch
	cb(&x.hTTPStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HTTP2AggregationsBuilder struct {
	writer           io.Writer
	buf              bytes.Buffer
	scratch          []byte
	hTTPStatsBuilder HTTPStatsBuilder
}

func NewHTTP2AggregationsBuilder(writer io.Writer) *HTTP2AggregationsBuilder {
	return &HTTP2AggregationsBuilder{
		writer: writer,
	}
}
func (x *HTTP2AggregationsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HTTP2AggregationsBuilder) AddEndpointAggregations(cb func(w *HTTPStatsBuilder)) {
	x.buf.Reset()
	x.hTTPStatsBuilder.writer = &x.buf
	x.hTTPStatsBuilder.scratch = x.scratch
	cb(&x.hTTPStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HTTPStatsBuilder struct {
	writer                                  io.Writer
	buf                                     bytes.Buffer
	scratch                                 []byte
	hTTPStats_DataBuilder                   HTTPStats_DataBuilder
	hTTPStats_StatsByStatusCodeEntryBuilder HTTPStats_StatsByStatusCodeEntryBuilder
}

func NewHTTPStatsBuilder(writer io.Writer) *HTTPStatsBuilder {
	return &HTTPStatsBuilder{
		writer: writer,
	}
}
func (x *HTTPStatsBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HTTPStatsBuilder) SetPath(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HTTPStatsBuilder) SetMethod(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x28)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *HTTPStatsBuilder) SetFullPath(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x30)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *HTTPStatsBuilder) AddStatsByResponseStatus(cb func(w *HTTPStats_DataBuilder)) {
	x.buf.Reset()
	x.hTTPStats_DataBuilder.writer = &x.buf
	x.hTTPStats_DataBuilder.scratch = x.scratch
	cb(&x.hTTPStats_DataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HTTPStatsBuilder) AddStatsByStatusCode(cb func(w *HTTPStats_StatsByStatusCodeEntryBuilder)) {
	x.buf.Reset()
	x.hTTPStats_StatsByStatusCodeEntryBuilder.writer = &x.buf
	x.hTTPStats_StatsByStatusCodeEntryBuilder.scratch = x.scratch
	cb(&x.hTTPStats_StatsByStatusCodeEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HTTPStats_StatsByStatusCodeEntryBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	hTTPStats_DataBuilder HTTPStats_DataBuilder
}

func NewHTTPStats_StatsByStatusCodeEntryBuilder(writer io.Writer) *HTTPStats_StatsByStatusCodeEntryBuilder {
	return &HTTPStats_StatsByStatusCodeEntryBuilder{
		writer: writer,
	}
}
func (x *HTTPStats_StatsByStatusCodeEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HTTPStats_StatsByStatusCodeEntryBuilder) SetKey(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HTTPStats_StatsByStatusCodeEntryBuilder) SetValue(cb func(w *HTTPStats_DataBuilder)) {
	x.buf.Reset()
	x.hTTPStats_DataBuilder.writer = &x.buf
	x.hTTPStats_DataBuilder.scratch = x.scratch
	cb(&x.hTTPStats_DataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HTTPStats_DataBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewHTTPStats_DataBuilder(writer io.Writer) *HTTPStats_DataBuilder {
	return &HTTPStats_DataBuilder{
		writer: writer,
	}
}
func (x *HTTPStats_DataBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HTTPStats_DataBuilder) SetCount(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HTTPStats_DataBuilder) SetLatencies(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HTTPStats_DataBuilder) SetFirstLatencySample(v float64) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x21)
	x.scratch = protowire.AppendFixed64(x.scratch, math.Float64bits(v))
	x.writer.Write(x.scratch)
}

type DNSDatabaseEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewDNSDatabaseEntryBuilder(writer io.Writer) *DNSDatabaseEntryBuilder {
	return &DNSDatabaseEntryBuilder{
		writer: writer,
	}
}
func (x *DNSDatabaseEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *DNSDatabaseEntryBuilder) AddNameOffsets(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ResourceListBuilder struct {
	writer                                io.Writer
	buf                                   bytes.Buffer
	scratch                               []byte
	resourceList_MetricValuesEntryBuilder ResourceList_MetricValuesEntryBuilder
}

func NewResourceListBuilder(writer io.Writer) *ResourceListBuilder {
	return &ResourceListBuilder{
		writer: writer,
	}
}
func (x *ResourceListBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResourceListBuilder) AddMetricValues(cb func(w *ResourceList_MetricValuesEntryBuilder)) {
	x.buf.Reset()
	x.resourceList_MetricValuesEntryBuilder.writer = &x.buf
	x.resourceList_MetricValuesEntryBuilder.scratch = x.scratch
	cb(&x.resourceList_MetricValuesEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ResourceList_MetricValuesEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResourceList_MetricValuesEntryBuilder(writer io.Writer) *ResourceList_MetricValuesEntryBuilder {
	return &ResourceList_MetricValuesEntryBuilder{
		writer: writer,
	}
}
func (x *ResourceList_MetricValuesEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResourceList_MetricValuesEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceList_MetricValuesEntryBuilder) SetValue(v float64) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed64(x.scratch, math.Float64bits(v))
	x.writer.Write(x.scratch)
}

type VerticalPodAutoscalerBuilder struct {
	writer                                io.Writer
	buf                                   bytes.Buffer
	scratch                               []byte
	metadataBuilder                       MetadataBuilder
	verticalPodAutoscalerSpecBuilder      VerticalPodAutoscalerSpecBuilder
	verticalPodAutoscalerStatusBuilder    VerticalPodAutoscalerStatusBuilder
	verticalPodAutoscalerConditionBuilder VerticalPodAutoscalerConditionBuilder
}

func NewVerticalPodAutoscalerBuilder(writer io.Writer) *VerticalPodAutoscalerBuilder {
	return &VerticalPodAutoscalerBuilder{
		writer: writer,
	}
}
func (x *VerticalPodAutoscalerBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *VerticalPodAutoscalerBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerBuilder) SetSpec(cb func(w *VerticalPodAutoscalerSpecBuilder)) {
	x.buf.Reset()
	x.verticalPodAutoscalerSpecBuilder.writer = &x.buf
	x.verticalPodAutoscalerSpecBuilder.scratch = x.scratch
	cb(&x.verticalPodAutoscalerSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerBuilder) SetStatus(cb func(w *VerticalPodAutoscalerStatusBuilder)) {
	x.buf.Reset()
	x.verticalPodAutoscalerStatusBuilder.writer = &x.buf
	x.verticalPodAutoscalerStatusBuilder.scratch = x.scratch
	cb(&x.verticalPodAutoscalerStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerBuilder) AddConditions(cb func(w *VerticalPodAutoscalerConditionBuilder)) {
	x.buf.Reset()
	x.verticalPodAutoscalerConditionBuilder.writer = &x.buf
	x.verticalPodAutoscalerConditionBuilder.scratch = x.scratch
	cb(&x.verticalPodAutoscalerConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type VerticalPodAutoscalerConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewVerticalPodAutoscalerConditionBuilder(writer io.Writer) *VerticalPodAutoscalerConditionBuilder {
	return &VerticalPodAutoscalerConditionBuilder{
		writer: writer,
	}
}
func (x *VerticalPodAutoscalerConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *VerticalPodAutoscalerConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type VerticalPodAutoscalerSpecBuilder struct {
	writer                             io.Writer
	buf                                bytes.Buffer
	scratch                            []byte
	verticalPodAutoscalerTargetBuilder VerticalPodAutoscalerTargetBuilder
	containerResourcePolicyBuilder     ContainerResourcePolicyBuilder
}

func NewVerticalPodAutoscalerSpecBuilder(writer io.Writer) *VerticalPodAutoscalerSpecBuilder {
	return &VerticalPodAutoscalerSpecBuilder{
		writer: writer,
	}
}
func (x *VerticalPodAutoscalerSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *VerticalPodAutoscalerSpecBuilder) SetTarget(cb func(w *VerticalPodAutoscalerTargetBuilder)) {
	x.buf.Reset()
	x.verticalPodAutoscalerTargetBuilder.writer = &x.buf
	x.verticalPodAutoscalerTargetBuilder.scratch = x.scratch
	cb(&x.verticalPodAutoscalerTargetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerSpecBuilder) SetUpdateMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerSpecBuilder) AddResourcePolicies(cb func(w *ContainerResourcePolicyBuilder)) {
	x.buf.Reset()
	x.containerResourcePolicyBuilder.writer = &x.buf
	x.containerResourcePolicyBuilder.scratch = x.scratch
	cb(&x.containerResourcePolicyBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type VerticalPodAutoscalerTargetBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewVerticalPodAutoscalerTargetBuilder(writer io.Writer) *VerticalPodAutoscalerTargetBuilder {
	return &VerticalPodAutoscalerTargetBuilder{
		writer: writer,
	}
}
func (x *VerticalPodAutoscalerTargetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *VerticalPodAutoscalerTargetBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerTargetBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ContainerResourcePolicyBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	resourceListBuilder ResourceListBuilder
}

func NewContainerResourcePolicyBuilder(writer io.Writer) *ContainerResourcePolicyBuilder {
	return &ContainerResourcePolicyBuilder{
		writer: writer,
	}
}
func (x *ContainerResourcePolicyBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ContainerResourcePolicyBuilder) SetContainerName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerResourcePolicyBuilder) SetMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerResourcePolicyBuilder) SetMinAllowed(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerResourcePolicyBuilder) SetMaxAllowed(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerResourcePolicyBuilder) AddControlledResource(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerResourcePolicyBuilder) SetControlledValues(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type VerticalPodAutoscalerStatusBuilder struct {
	writer                         io.Writer
	buf                            bytes.Buffer
	scratch                        []byte
	containerRecommendationBuilder ContainerRecommendationBuilder
	vPAConditionBuilder            VPAConditionBuilder
}

func NewVerticalPodAutoscalerStatusBuilder(writer io.Writer) *VerticalPodAutoscalerStatusBuilder {
	return &VerticalPodAutoscalerStatusBuilder{
		writer: writer,
	}
}
func (x *VerticalPodAutoscalerStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *VerticalPodAutoscalerStatusBuilder) SetLastRecommendedDate(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerStatusBuilder) AddRecommendations(cb func(w *ContainerRecommendationBuilder)) {
	x.buf.Reset()
	x.containerRecommendationBuilder.writer = &x.buf
	x.containerRecommendationBuilder.scratch = x.scratch
	cb(&x.containerRecommendationBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerStatusBuilder) AddConditions(cb func(w *VPAConditionBuilder)) {
	x.buf.Reset()
	x.vPAConditionBuilder.writer = &x.buf
	x.vPAConditionBuilder.scratch = x.scratch
	cb(&x.vPAConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ContainerRecommendationBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	resourceListBuilder ResourceListBuilder
}

func NewContainerRecommendationBuilder(writer io.Writer) *ContainerRecommendationBuilder {
	return &ContainerRecommendationBuilder{
		writer: writer,
	}
}
func (x *ContainerRecommendationBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ContainerRecommendationBuilder) SetContainerName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerRecommendationBuilder) SetTarget(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerRecommendationBuilder) SetLowerBound(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerRecommendationBuilder) SetUpperBound(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerRecommendationBuilder) SetUncappedTarget(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type VPAConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewVPAConditionBuilder(writer io.Writer) *VPAConditionBuilder {
	return &VPAConditionBuilder{
		writer: writer,
	}
}
func (x *VPAConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *VPAConditionBuilder) SetConditionType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VPAConditionBuilder) SetConditionStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VPAConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *VPAConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VPAConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type HorizontalPodAutoscalerBuilder struct {
	writer                                  io.Writer
	buf                                     bytes.Buffer
	scratch                                 []byte
	metadataBuilder                         MetadataBuilder
	horizontalPodAutoscalerSpecBuilder      HorizontalPodAutoscalerSpecBuilder
	horizontalPodAutoscalerStatusBuilder    HorizontalPodAutoscalerStatusBuilder
	horizontalPodAutoscalerConditionBuilder HorizontalPodAutoscalerConditionBuilder
}

func NewHorizontalPodAutoscalerBuilder(writer io.Writer) *HorizontalPodAutoscalerBuilder {
	return &HorizontalPodAutoscalerBuilder{
		writer: writer,
	}
}
func (x *HorizontalPodAutoscalerBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HorizontalPodAutoscalerBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerBuilder) SetSpec(cb func(w *HorizontalPodAutoscalerSpecBuilder)) {
	x.buf.Reset()
	x.horizontalPodAutoscalerSpecBuilder.writer = &x.buf
	x.horizontalPodAutoscalerSpecBuilder.scratch = x.scratch
	cb(&x.horizontalPodAutoscalerSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerBuilder) SetStatus(cb func(w *HorizontalPodAutoscalerStatusBuilder)) {
	x.buf.Reset()
	x.horizontalPodAutoscalerStatusBuilder.writer = &x.buf
	x.horizontalPodAutoscalerStatusBuilder.scratch = x.scratch
	cb(&x.horizontalPodAutoscalerStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerBuilder) AddConditions(cb func(w *HorizontalPodAutoscalerConditionBuilder)) {
	x.buf.Reset()
	x.horizontalPodAutoscalerConditionBuilder.writer = &x.buf
	x.horizontalPodAutoscalerConditionBuilder.scratch = x.scratch
	cb(&x.horizontalPodAutoscalerConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HorizontalPodAutoscalerSpecBuilder struct {
	writer                                   io.Writer
	buf                                      bytes.Buffer
	scratch                                  []byte
	horizontalPodAutoscalerTargetBuilder     HorizontalPodAutoscalerTargetBuilder
	horizontalPodAutoscalerMetricSpecBuilder HorizontalPodAutoscalerMetricSpecBuilder
	horizontalPodAutoscalerBehaviorBuilder   HorizontalPodAutoscalerBehaviorBuilder
}

func NewHorizontalPodAutoscalerSpecBuilder(writer io.Writer) *HorizontalPodAutoscalerSpecBuilder {
	return &HorizontalPodAutoscalerSpecBuilder{
		writer: writer,
	}
}
func (x *HorizontalPodAutoscalerSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HorizontalPodAutoscalerSpecBuilder) SetTarget(cb func(w *HorizontalPodAutoscalerTargetBuilder)) {
	x.buf.Reset()
	x.horizontalPodAutoscalerTargetBuilder.writer = &x.buf
	x.horizontalPodAutoscalerTargetBuilder.scratch = x.scratch
	cb(&x.horizontalPodAutoscalerTargetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerSpecBuilder) SetMinReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerSpecBuilder) SetMaxReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerSpecBuilder) AddMetrics(cb func(w *HorizontalPodAutoscalerMetricSpecBuilder)) {
	x.buf.Reset()
	x.horizontalPodAutoscalerMetricSpecBuilder.writer = &x.buf
	x.horizontalPodAutoscalerMetricSpecBuilder.scratch = x.scratch
	cb(&x.horizontalPodAutoscalerMetricSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerSpecBuilder) SetBehavior(cb func(w *HorizontalPodAutoscalerBehaviorBuilder)) {
	x.buf.Reset()
	x.horizontalPodAutoscalerBehaviorBuilder.writer = &x.buf
	x.horizontalPodAutoscalerBehaviorBuilder.scratch = x.scratch
	cb(&x.horizontalPodAutoscalerBehaviorBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HorizontalPodAutoscalerTargetBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewHorizontalPodAutoscalerTargetBuilder(writer io.Writer) *HorizontalPodAutoscalerTargetBuilder {
	return &HorizontalPodAutoscalerTargetBuilder{
		writer: writer,
	}
}
func (x *HorizontalPodAutoscalerTargetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HorizontalPodAutoscalerTargetBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerTargetBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type HorizontalPodAutoscalerMetricSpecBuilder struct {
	writer                               io.Writer
	buf                                  bytes.Buffer
	scratch                              []byte
	objectMetricSourceBuilder            ObjectMetricSourceBuilder
	podsMetricSourceBuilder              PodsMetricSourceBuilder
	resourceMetricSourceBuilder          ResourceMetricSourceBuilder
	containerResourceMetricSourceBuilder ContainerResourceMetricSourceBuilder
	externalMetricSourceBuilder          ExternalMetricSourceBuilder
}

func NewHorizontalPodAutoscalerMetricSpecBuilder(writer io.Writer) *HorizontalPodAutoscalerMetricSpecBuilder {
	return &HorizontalPodAutoscalerMetricSpecBuilder{
		writer: writer,
	}
}
func (x *HorizontalPodAutoscalerMetricSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HorizontalPodAutoscalerMetricSpecBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerMetricSpecBuilder) SetObject(cb func(w *ObjectMetricSourceBuilder)) {
	x.buf.Reset()
	x.objectMetricSourceBuilder.writer = &x.buf
	x.objectMetricSourceBuilder.scratch = x.scratch
	cb(&x.objectMetricSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerMetricSpecBuilder) SetPods(cb func(w *PodsMetricSourceBuilder)) {
	x.buf.Reset()
	x.podsMetricSourceBuilder.writer = &x.buf
	x.podsMetricSourceBuilder.scratch = x.scratch
	cb(&x.podsMetricSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerMetricSpecBuilder) SetResource(cb func(w *ResourceMetricSourceBuilder)) {
	x.buf.Reset()
	x.resourceMetricSourceBuilder.writer = &x.buf
	x.resourceMetricSourceBuilder.scratch = x.scratch
	cb(&x.resourceMetricSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerMetricSpecBuilder) SetContainerResource(cb func(w *ContainerResourceMetricSourceBuilder)) {
	x.buf.Reset()
	x.containerResourceMetricSourceBuilder.writer = &x.buf
	x.containerResourceMetricSourceBuilder.scratch = x.scratch
	cb(&x.containerResourceMetricSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerMetricSpecBuilder) SetExternal(cb func(w *ExternalMetricSourceBuilder)) {
	x.buf.Reset()
	x.externalMetricSourceBuilder.writer = &x.buf
	x.externalMetricSourceBuilder.scratch = x.scratch
	cb(&x.externalMetricSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ObjectMetricSourceBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	objectReferenceBuilder  ObjectReferenceBuilder
	metricTargetBuilder     MetricTargetBuilder
	metricIdentifierBuilder MetricIdentifierBuilder
}

func NewObjectMetricSourceBuilder(writer io.Writer) *ObjectMetricSourceBuilder {
	return &ObjectMetricSourceBuilder{
		writer: writer,
	}
}
func (x *ObjectMetricSourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ObjectMetricSourceBuilder) SetDescribedObject(cb func(w *ObjectReferenceBuilder)) {
	x.buf.Reset()
	x.objectReferenceBuilder.writer = &x.buf
	x.objectReferenceBuilder.scratch = x.scratch
	cb(&x.objectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ObjectMetricSourceBuilder) SetTarget(cb func(w *MetricTargetBuilder)) {
	x.buf.Reset()
	x.metricTargetBuilder.writer = &x.buf
	x.metricTargetBuilder.scratch = x.scratch
	cb(&x.metricTargetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ObjectMetricSourceBuilder) SetMetric(cb func(w *MetricIdentifierBuilder)) {
	x.buf.Reset()
	x.metricIdentifierBuilder.writer = &x.buf
	x.metricIdentifierBuilder.scratch = x.scratch
	cb(&x.metricIdentifierBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type MetricTargetBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewMetricTargetBuilder(writer io.Writer) *MetricTargetBuilder {
	return &MetricTargetBuilder{
		writer: writer,
	}
}
func (x *MetricTargetBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *MetricTargetBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetricTargetBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type MetricIdentifierBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
}

func NewMetricIdentifierBuilder(writer io.Writer) *MetricIdentifierBuilder {
	return &MetricIdentifierBuilder{
		writer: writer,
	}
}
func (x *MetricIdentifierBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *MetricIdentifierBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetricIdentifierBuilder) AddLabelSelector(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PodsMetricSourceBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	metricIdentifierBuilder MetricIdentifierBuilder
	metricTargetBuilder     MetricTargetBuilder
}

func NewPodsMetricSourceBuilder(writer io.Writer) *PodsMetricSourceBuilder {
	return &PodsMetricSourceBuilder{
		writer: writer,
	}
}
func (x *PodsMetricSourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PodsMetricSourceBuilder) SetMetric(cb func(w *MetricIdentifierBuilder)) {
	x.buf.Reset()
	x.metricIdentifierBuilder.writer = &x.buf
	x.metricIdentifierBuilder.scratch = x.scratch
	cb(&x.metricIdentifierBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodsMetricSourceBuilder) SetTarget(cb func(w *MetricTargetBuilder)) {
	x.buf.Reset()
	x.metricTargetBuilder.writer = &x.buf
	x.metricTargetBuilder.scratch = x.scratch
	cb(&x.metricTargetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ResourceMetricSourceBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	metricTargetBuilder MetricTargetBuilder
}

func NewResourceMetricSourceBuilder(writer io.Writer) *ResourceMetricSourceBuilder {
	return &ResourceMetricSourceBuilder{
		writer: writer,
	}
}
func (x *ResourceMetricSourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResourceMetricSourceBuilder) SetResourceName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceMetricSourceBuilder) SetTarget(cb func(w *MetricTargetBuilder)) {
	x.buf.Reset()
	x.metricTargetBuilder.writer = &x.buf
	x.metricTargetBuilder.scratch = x.scratch
	cb(&x.metricTargetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ContainerResourceMetricSourceBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	metricTargetBuilder MetricTargetBuilder
}

func NewContainerResourceMetricSourceBuilder(writer io.Writer) *ContainerResourceMetricSourceBuilder {
	return &ContainerResourceMetricSourceBuilder{
		writer: writer,
	}
}
func (x *ContainerResourceMetricSourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ContainerResourceMetricSourceBuilder) SetResourceName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerResourceMetricSourceBuilder) SetTarget(cb func(w *MetricTargetBuilder)) {
	x.buf.Reset()
	x.metricTargetBuilder.writer = &x.buf
	x.metricTargetBuilder.scratch = x.scratch
	cb(&x.metricTargetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerResourceMetricSourceBuilder) SetContainer(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ExternalMetricSourceBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	metricIdentifierBuilder MetricIdentifierBuilder
	metricTargetBuilder     MetricTargetBuilder
}

func NewExternalMetricSourceBuilder(writer io.Writer) *ExternalMetricSourceBuilder {
	return &ExternalMetricSourceBuilder{
		writer: writer,
	}
}
func (x *ExternalMetricSourceBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ExternalMetricSourceBuilder) SetMetric(cb func(w *MetricIdentifierBuilder)) {
	x.buf.Reset()
	x.metricIdentifierBuilder.writer = &x.buf
	x.metricIdentifierBuilder.scratch = x.scratch
	cb(&x.metricIdentifierBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ExternalMetricSourceBuilder) SetTarget(cb func(w *MetricTargetBuilder)) {
	x.buf.Reset()
	x.metricTargetBuilder.writer = &x.buf
	x.metricTargetBuilder.scratch = x.scratch
	cb(&x.metricTargetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HorizontalPodAutoscalerBehaviorBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	hPAScalingRulesBuilder HPAScalingRulesBuilder
}

func NewHorizontalPodAutoscalerBehaviorBuilder(writer io.Writer) *HorizontalPodAutoscalerBehaviorBuilder {
	return &HorizontalPodAutoscalerBehaviorBuilder{
		writer: writer,
	}
}
func (x *HorizontalPodAutoscalerBehaviorBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HorizontalPodAutoscalerBehaviorBuilder) SetScaleUp(cb func(w *HPAScalingRulesBuilder)) {
	x.buf.Reset()
	x.hPAScalingRulesBuilder.writer = &x.buf
	x.hPAScalingRulesBuilder.scratch = x.scratch
	cb(&x.hPAScalingRulesBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerBehaviorBuilder) SetScaleDown(cb func(w *HPAScalingRulesBuilder)) {
	x.buf.Reset()
	x.hPAScalingRulesBuilder.writer = &x.buf
	x.hPAScalingRulesBuilder.scratch = x.scratch
	cb(&x.hPAScalingRulesBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HPAScalingRulesBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	hPAScalingPolicyBuilder HPAScalingPolicyBuilder
}

func NewHPAScalingRulesBuilder(writer io.Writer) *HPAScalingRulesBuilder {
	return &HPAScalingRulesBuilder{
		writer: writer,
	}
}
func (x *HPAScalingRulesBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HPAScalingRulesBuilder) SetStabilizationWindowSeconds(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HPAScalingRulesBuilder) SetSelectPolicy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HPAScalingRulesBuilder) AddPolicies(cb func(w *HPAScalingPolicyBuilder)) {
	x.buf.Reset()
	x.hPAScalingPolicyBuilder.writer = &x.buf
	x.hPAScalingPolicyBuilder.scratch = x.scratch
	cb(&x.hPAScalingPolicyBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HPAScalingPolicyBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewHPAScalingPolicyBuilder(writer io.Writer) *HPAScalingPolicyBuilder {
	return &HPAScalingPolicyBuilder{
		writer: writer,
	}
}
func (x *HPAScalingPolicyBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HPAScalingPolicyBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HPAScalingPolicyBuilder) SetValue(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HPAScalingPolicyBuilder) SetPeriodSeconds(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type HorizontalPodAutoscalerStatusBuilder struct {
	writer                                     io.Writer
	buf                                        bytes.Buffer
	scratch                                    []byte
	horizontalPodAutoscalerMetricStatusBuilder HorizontalPodAutoscalerMetricStatusBuilder
}

func NewHorizontalPodAutoscalerStatusBuilder(writer io.Writer) *HorizontalPodAutoscalerStatusBuilder {
	return &HorizontalPodAutoscalerStatusBuilder{
		writer: writer,
	}
}
func (x *HorizontalPodAutoscalerStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HorizontalPodAutoscalerStatusBuilder) SetObservedGeneration(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerStatusBuilder) SetLastScaleTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerStatusBuilder) SetCurrentReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerStatusBuilder) SetDesiredReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerStatusBuilder) AddCurrentMetrics(cb func(w *HorizontalPodAutoscalerMetricStatusBuilder)) {
	x.buf.Reset()
	x.horizontalPodAutoscalerMetricStatusBuilder.writer = &x.buf
	x.horizontalPodAutoscalerMetricStatusBuilder.scratch = x.scratch
	cb(&x.horizontalPodAutoscalerMetricStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HorizontalPodAutoscalerConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewHorizontalPodAutoscalerConditionBuilder(writer io.Writer) *HorizontalPodAutoscalerConditionBuilder {
	return &HorizontalPodAutoscalerConditionBuilder{
		writer: writer,
	}
}
func (x *HorizontalPodAutoscalerConditionBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HorizontalPodAutoscalerConditionBuilder) SetConditionType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerConditionBuilder) SetConditionStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type HorizontalPodAutoscalerMetricStatusBuilder struct {
	writer                               io.Writer
	buf                                  bytes.Buffer
	scratch                              []byte
	objectMetricStatusBuilder            ObjectMetricStatusBuilder
	podsMetricStatusBuilder              PodsMetricStatusBuilder
	resourceMetricStatusBuilder          ResourceMetricStatusBuilder
	containerResourceMetricStatusBuilder ContainerResourceMetricStatusBuilder
	externalMetricStatusBuilder          ExternalMetricStatusBuilder
}

func NewHorizontalPodAutoscalerMetricStatusBuilder(writer io.Writer) *HorizontalPodAutoscalerMetricStatusBuilder {
	return &HorizontalPodAutoscalerMetricStatusBuilder{
		writer: writer,
	}
}
func (x *HorizontalPodAutoscalerMetricStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *HorizontalPodAutoscalerMetricStatusBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HorizontalPodAutoscalerMetricStatusBuilder) SetObject(cb func(w *ObjectMetricStatusBuilder)) {
	x.buf.Reset()
	x.objectMetricStatusBuilder.writer = &x.buf
	x.objectMetricStatusBuilder.scratch = x.scratch
	cb(&x.objectMetricStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerMetricStatusBuilder) SetPods(cb func(w *PodsMetricStatusBuilder)) {
	x.buf.Reset()
	x.podsMetricStatusBuilder.writer = &x.buf
	x.podsMetricStatusBuilder.scratch = x.scratch
	cb(&x.podsMetricStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerMetricStatusBuilder) SetResource(cb func(w *ResourceMetricStatusBuilder)) {
	x.buf.Reset()
	x.resourceMetricStatusBuilder.writer = &x.buf
	x.resourceMetricStatusBuilder.scratch = x.scratch
	cb(&x.resourceMetricStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerMetricStatusBuilder) SetContainerResource(cb func(w *ContainerResourceMetricStatusBuilder)) {
	x.buf.Reset()
	x.containerResourceMetricStatusBuilder.writer = &x.buf
	x.containerResourceMetricStatusBuilder.scratch = x.scratch
	cb(&x.containerResourceMetricStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HorizontalPodAutoscalerMetricStatusBuilder) SetExternal(cb func(w *ExternalMetricStatusBuilder)) {
	x.buf.Reset()
	x.externalMetricStatusBuilder.writer = &x.buf
	x.externalMetricStatusBuilder.scratch = x.scratch
	cb(&x.externalMetricStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ObjectMetricStatusBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	objectReferenceBuilder  ObjectReferenceBuilder
	metricIdentifierBuilder MetricIdentifierBuilder
}

func NewObjectMetricStatusBuilder(writer io.Writer) *ObjectMetricStatusBuilder {
	return &ObjectMetricStatusBuilder{
		writer: writer,
	}
}
func (x *ObjectMetricStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ObjectMetricStatusBuilder) SetDescribedObject(cb func(w *ObjectReferenceBuilder)) {
	x.buf.Reset()
	x.objectReferenceBuilder.writer = &x.buf
	x.objectReferenceBuilder.scratch = x.scratch
	cb(&x.objectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ObjectMetricStatusBuilder) SetCurrent(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ObjectMetricStatusBuilder) SetMetric(cb func(w *MetricIdentifierBuilder)) {
	x.buf.Reset()
	x.metricIdentifierBuilder.writer = &x.buf
	x.metricIdentifierBuilder.scratch = x.scratch
	cb(&x.metricIdentifierBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PodsMetricStatusBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	metricIdentifierBuilder MetricIdentifierBuilder
}

func NewPodsMetricStatusBuilder(writer io.Writer) *PodsMetricStatusBuilder {
	return &PodsMetricStatusBuilder{
		writer: writer,
	}
}
func (x *PodsMetricStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *PodsMetricStatusBuilder) SetMetric(cb func(w *MetricIdentifierBuilder)) {
	x.buf.Reset()
	x.metricIdentifierBuilder.writer = &x.buf
	x.metricIdentifierBuilder.scratch = x.scratch
	cb(&x.metricIdentifierBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodsMetricStatusBuilder) SetCurrent(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ResourceMetricStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResourceMetricStatusBuilder(writer io.Writer) *ResourceMetricStatusBuilder {
	return &ResourceMetricStatusBuilder{
		writer: writer,
	}
}
func (x *ResourceMetricStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ResourceMetricStatusBuilder) SetResourceName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceMetricStatusBuilder) SetCurrent(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ContainerResourceMetricStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewContainerResourceMetricStatusBuilder(writer io.Writer) *ContainerResourceMetricStatusBuilder {
	return &ContainerResourceMetricStatusBuilder{
		writer: writer,
	}
}
func (x *ContainerResourceMetricStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ContainerResourceMetricStatusBuilder) SetResourceName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerResourceMetricStatusBuilder) SetCurrent(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerResourceMetricStatusBuilder) SetContainer(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ExternalMetricStatusBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	metricIdentifierBuilder MetricIdentifierBuilder
}

func NewExternalMetricStatusBuilder(writer io.Writer) *ExternalMetricStatusBuilder {
	return &ExternalMetricStatusBuilder{
		writer: writer,
	}
}
func (x *ExternalMetricStatusBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ExternalMetricStatusBuilder) SetMetric(cb func(w *MetricIdentifierBuilder)) {
	x.buf.Reset()
	x.metricIdentifierBuilder.writer = &x.buf
	x.metricIdentifierBuilder.scratch = x.scratch
	cb(&x.metricIdentifierBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ExternalMetricStatusBuilder) SetCurrent(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type NetworkPolicyBuilder struct {
	writer                   io.Writer
	buf                      bytes.Buffer
	scratch                  []byte
	metadataBuilder          MetadataBuilder
	networkPolicySpecBuilder NetworkPolicySpecBuilder
}

func NewNetworkPolicyBuilder(writer io.Writer) *NetworkPolicyBuilder {
	return &NetworkPolicyBuilder{
		writer: writer,
	}
}
func (x *NetworkPolicyBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NetworkPolicyBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NetworkPolicyBuilder) SetSpec(cb func(w *NetworkPolicySpecBuilder)) {
	x.buf.Reset()
	x.networkPolicySpecBuilder.writer = &x.buf
	x.networkPolicySpecBuilder.scratch = x.scratch
	cb(&x.networkPolicySpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NetworkPolicyBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NetworkPolicyBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NetworkPolicySpecBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	networkPolicyIngressRuleBuilder NetworkPolicyIngressRuleBuilder
	networkPolicyEgressRuleBuilder  NetworkPolicyEgressRuleBuilder
}

func NewNetworkPolicySpecBuilder(writer io.Writer) *NetworkPolicySpecBuilder {
	return &NetworkPolicySpecBuilder{
		writer: writer,
	}
}
func (x *NetworkPolicySpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NetworkPolicySpecBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NetworkPolicySpecBuilder) AddIngress(cb func(w *NetworkPolicyIngressRuleBuilder)) {
	x.buf.Reset()
	x.networkPolicyIngressRuleBuilder.writer = &x.buf
	x.networkPolicyIngressRuleBuilder.scratch = x.scratch
	cb(&x.networkPolicyIngressRuleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NetworkPolicySpecBuilder) AddEgress(cb func(w *NetworkPolicyEgressRuleBuilder)) {
	x.buf.Reset()
	x.networkPolicyEgressRuleBuilder.writer = &x.buf
	x.networkPolicyEgressRuleBuilder.scratch = x.scratch
	cb(&x.networkPolicyEgressRuleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NetworkPolicySpecBuilder) AddPolicyTypes(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NetworkPolicyIPBlockBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNetworkPolicyIPBlockBuilder(writer io.Writer) *NetworkPolicyIPBlockBuilder {
	return &NetworkPolicyIPBlockBuilder{
		writer: writer,
	}
}
func (x *NetworkPolicyIPBlockBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NetworkPolicyIPBlockBuilder) SetCidr(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NetworkPolicyIPBlockBuilder) AddExcept(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NetworkPolicyIngressRuleBuilder struct {
	writer                   io.Writer
	buf                      bytes.Buffer
	scratch                  []byte
	networkPolicyPortBuilder NetworkPolicyPortBuilder
	networkPolicyPeerBuilder NetworkPolicyPeerBuilder
}

func NewNetworkPolicyIngressRuleBuilder(writer io.Writer) *NetworkPolicyIngressRuleBuilder {
	return &NetworkPolicyIngressRuleBuilder{
		writer: writer,
	}
}
func (x *NetworkPolicyIngressRuleBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NetworkPolicyIngressRuleBuilder) AddPorts(cb func(w *NetworkPolicyPortBuilder)) {
	x.buf.Reset()
	x.networkPolicyPortBuilder.writer = &x.buf
	x.networkPolicyPortBuilder.scratch = x.scratch
	cb(&x.networkPolicyPortBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NetworkPolicyIngressRuleBuilder) AddFrom(cb func(w *NetworkPolicyPeerBuilder)) {
	x.buf.Reset()
	x.networkPolicyPeerBuilder.writer = &x.buf
	x.networkPolicyPeerBuilder.scratch = x.scratch
	cb(&x.networkPolicyPeerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type NetworkPolicyEgressRuleBuilder struct {
	writer                   io.Writer
	buf                      bytes.Buffer
	scratch                  []byte
	networkPolicyPortBuilder NetworkPolicyPortBuilder
	networkPolicyPeerBuilder NetworkPolicyPeerBuilder
}

func NewNetworkPolicyEgressRuleBuilder(writer io.Writer) *NetworkPolicyEgressRuleBuilder {
	return &NetworkPolicyEgressRuleBuilder{
		writer: writer,
	}
}
func (x *NetworkPolicyEgressRuleBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NetworkPolicyEgressRuleBuilder) AddPorts(cb func(w *NetworkPolicyPortBuilder)) {
	x.buf.Reset()
	x.networkPolicyPortBuilder.writer = &x.buf
	x.networkPolicyPortBuilder.scratch = x.scratch
	cb(&x.networkPolicyPortBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NetworkPolicyEgressRuleBuilder) AddTo(cb func(w *NetworkPolicyPeerBuilder)) {
	x.buf.Reset()
	x.networkPolicyPeerBuilder.writer = &x.buf
	x.networkPolicyPeerBuilder.scratch = x.scratch
	cb(&x.networkPolicyPeerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type NetworkPolicyPeerBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	networkPolicyIPBlockBuilder     NetworkPolicyIPBlockBuilder
}

func NewNetworkPolicyPeerBuilder(writer io.Writer) *NetworkPolicyPeerBuilder {
	return &NetworkPolicyPeerBuilder{
		writer: writer,
	}
}
func (x *NetworkPolicyPeerBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NetworkPolicyPeerBuilder) AddPodSelector(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NetworkPolicyPeerBuilder) AddNamespaceSelector(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NetworkPolicyPeerBuilder) SetIpBlock(cb func(w *NetworkPolicyIPBlockBuilder)) {
	x.buf.Reset()
	x.networkPolicyIPBlockBuilder.writer = &x.buf
	x.networkPolicyIPBlockBuilder.scratch = x.scratch
	cb(&x.networkPolicyIPBlockBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NetworkPolicyPeerBuilder) SetHasPodSelector(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *NetworkPolicyPeerBuilder) SetHasNamespaceSelector(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x28)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}

type NetworkPolicyPortBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNetworkPolicyPortBuilder(writer io.Writer) *NetworkPolicyPortBuilder {
	return &NetworkPolicyPortBuilder{
		writer: writer,
	}
}
func (x *NetworkPolicyPortBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *NetworkPolicyPortBuilder) SetProtocol(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NetworkPolicyPortBuilder) SetPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *NetworkPolicyPortBuilder) SetEndPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type LimitRangeBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	metadataBuilder       MetadataBuilder
	limitRangeSpecBuilder LimitRangeSpecBuilder
}

func NewLimitRangeBuilder(writer io.Writer) *LimitRangeBuilder {
	return &LimitRangeBuilder{
		writer: writer,
	}
}
func (x *LimitRangeBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *LimitRangeBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *LimitRangeBuilder) SetSpec(cb func(w *LimitRangeSpecBuilder)) {
	x.buf.Reset()
	x.limitRangeSpecBuilder.writer = &x.buf
	x.limitRangeSpecBuilder.scratch = x.scratch
	cb(&x.limitRangeSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *LimitRangeBuilder) AddLimitTypes(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LimitRangeBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type LimitRangeSpecBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	limitRangeItemBuilder LimitRangeItemBuilder
}

func NewLimitRangeSpecBuilder(writer io.Writer) *LimitRangeSpecBuilder {
	return &LimitRangeSpecBuilder{
		writer: writer,
	}
}
func (x *LimitRangeSpecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *LimitRangeSpecBuilder) AddLimits(cb func(w *LimitRangeItemBuilder)) {
	x.buf.Reset()
	x.limitRangeItemBuilder.writer = &x.buf
	x.limitRangeItemBuilder.scratch = x.scratch
	cb(&x.limitRangeItemBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type LimitRangeItemBuilder struct {
	writer                                          io.Writer
	buf                                             bytes.Buffer
	scratch                                         []byte
	limitRangeItem_DefaultEntryBuilder              LimitRangeItem_DefaultEntryBuilder
	limitRangeItem_DefaultRequestEntryBuilder       LimitRangeItem_DefaultRequestEntryBuilder
	limitRangeItem_MaxEntryBuilder                  LimitRangeItem_MaxEntryBuilder
	limitRangeItem_MinEntryBuilder                  LimitRangeItem_MinEntryBuilder
	limitRangeItem_MaxLimitRequestRatioEntryBuilder LimitRangeItem_MaxLimitRequestRatioEntryBuilder
}

func NewLimitRangeItemBuilder(writer io.Writer) *LimitRangeItemBuilder {
	return &LimitRangeItemBuilder{
		writer: writer,
	}
}
func (x *LimitRangeItemBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *LimitRangeItemBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LimitRangeItemBuilder) AddDefault(cb func(w *LimitRangeItem_DefaultEntryBuilder)) {
	x.buf.Reset()
	x.limitRangeItem_DefaultEntryBuilder.writer = &x.buf
	x.limitRangeItem_DefaultEntryBuilder.scratch = x.scratch
	cb(&x.limitRangeItem_DefaultEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *LimitRangeItemBuilder) AddDefaultRequest(cb func(w *LimitRangeItem_DefaultRequestEntryBuilder)) {
	x.buf.Reset()
	x.limitRangeItem_DefaultRequestEntryBuilder.writer = &x.buf
	x.limitRangeItem_DefaultRequestEntryBuilder.scratch = x.scratch
	cb(&x.limitRangeItem_DefaultRequestEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *LimitRangeItemBuilder) AddMax(cb func(w *LimitRangeItem_MaxEntryBuilder)) {
	x.buf.Reset()
	x.limitRangeItem_MaxEntryBuilder.writer = &x.buf
	x.limitRangeItem_MaxEntryBuilder.scratch = x.scratch
	cb(&x.limitRangeItem_MaxEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *LimitRangeItemBuilder) AddMin(cb func(w *LimitRangeItem_MinEntryBuilder)) {
	x.buf.Reset()
	x.limitRangeItem_MinEntryBuilder.writer = &x.buf
	x.limitRangeItem_MinEntryBuilder.scratch = x.scratch
	cb(&x.limitRangeItem_MinEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *LimitRangeItemBuilder) AddMaxLimitRequestRatio(cb func(w *LimitRangeItem_MaxLimitRequestRatioEntryBuilder)) {
	x.buf.Reset()
	x.limitRangeItem_MaxLimitRequestRatioEntryBuilder.writer = &x.buf
	x.limitRangeItem_MaxLimitRequestRatioEntryBuilder.scratch = x.scratch
	cb(&x.limitRangeItem_MaxLimitRequestRatioEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type LimitRangeItem_DefaultEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewLimitRangeItem_DefaultEntryBuilder(writer io.Writer) *LimitRangeItem_DefaultEntryBuilder {
	return &LimitRangeItem_DefaultEntryBuilder{
		writer: writer,
	}
}
func (x *LimitRangeItem_DefaultEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *LimitRangeItem_DefaultEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LimitRangeItem_DefaultEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type LimitRangeItem_DefaultRequestEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewLimitRangeItem_DefaultRequestEntryBuilder(writer io.Writer) *LimitRangeItem_DefaultRequestEntryBuilder {
	return &LimitRangeItem_DefaultRequestEntryBuilder{
		writer: writer,
	}
}
func (x *LimitRangeItem_DefaultRequestEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *LimitRangeItem_DefaultRequestEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LimitRangeItem_DefaultRequestEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type LimitRangeItem_MaxEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewLimitRangeItem_MaxEntryBuilder(writer io.Writer) *LimitRangeItem_MaxEntryBuilder {
	return &LimitRangeItem_MaxEntryBuilder{
		writer: writer,
	}
}
func (x *LimitRangeItem_MaxEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *LimitRangeItem_MaxEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LimitRangeItem_MaxEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type LimitRangeItem_MinEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewLimitRangeItem_MinEntryBuilder(writer io.Writer) *LimitRangeItem_MinEntryBuilder {
	return &LimitRangeItem_MinEntryBuilder{
		writer: writer,
	}
}
func (x *LimitRangeItem_MinEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *LimitRangeItem_MinEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LimitRangeItem_MinEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type LimitRangeItem_MaxLimitRequestRatioEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewLimitRangeItem_MaxLimitRequestRatioEntryBuilder(writer io.Writer) *LimitRangeItem_MaxLimitRequestRatioEntryBuilder {
	return &LimitRangeItem_MaxLimitRequestRatioEntryBuilder{
		writer: writer,
	}
}
func (x *LimitRangeItem_MaxLimitRequestRatioEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *LimitRangeItem_MaxLimitRequestRatioEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LimitRangeItem_MaxLimitRequestRatioEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type StorageClassBuilder struct {
	writer                              io.Writer
	buf                                 bytes.Buffer
	scratch                             []byte
	metadataBuilder                     MetadataBuilder
	storageClass_ParametersEntryBuilder StorageClass_ParametersEntryBuilder
	storageClassTopologiesBuilder       StorageClassTopologiesBuilder
}

func NewStorageClassBuilder(writer io.Writer) *StorageClassBuilder {
	return &StorageClassBuilder{
		writer: writer,
	}
}
func (x *StorageClassBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *StorageClassBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StorageClassBuilder) SetProvisioner(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StorageClassBuilder) AddParameters(cb func(w *StorageClass_ParametersEntryBuilder)) {
	x.buf.Reset()
	x.storageClass_ParametersEntryBuilder.writer = &x.buf
	x.storageClass_ParametersEntryBuilder.scratch = x.scratch
	cb(&x.storageClass_ParametersEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StorageClassBuilder) SetReclaimPolicy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StorageClassBuilder) AddMountOptions(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StorageClassBuilder) SetAllowVolumeExpansion(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x30)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *StorageClassBuilder) SetAllowedTopologies(cb func(w *StorageClassTopologiesBuilder)) {
	x.buf.Reset()
	x.storageClassTopologiesBuilder.writer = &x.buf
	x.storageClassTopologiesBuilder.scratch = x.scratch
	cb(&x.storageClassTopologiesBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StorageClassBuilder) SetVolumeBindingMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StorageClassBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type StorageClass_ParametersEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewStorageClass_ParametersEntryBuilder(writer io.Writer) *StorageClass_ParametersEntryBuilder {
	return &StorageClass_ParametersEntryBuilder{
		writer: writer,
	}
}
func (x *StorageClass_ParametersEntryBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *StorageClass_ParametersEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StorageClass_ParametersEntryBuilder) SetValue(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type StorageClassTopologiesBuilder struct {
	writer                       io.Writer
	buf                          bytes.Buffer
	scratch                      []byte
	topologyLabelSelectorBuilder TopologyLabelSelectorBuilder
}

func NewStorageClassTopologiesBuilder(writer io.Writer) *StorageClassTopologiesBuilder {
	return &StorageClassTopologiesBuilder{
		writer: writer,
	}
}
func (x *StorageClassTopologiesBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *StorageClassTopologiesBuilder) AddLabelSelectors(cb func(w *TopologyLabelSelectorBuilder)) {
	x.buf.Reset()
	x.topologyLabelSelectorBuilder.writer = &x.buf
	x.topologyLabelSelectorBuilder.scratch = x.scratch
	cb(&x.topologyLabelSelectorBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type TopologyLabelSelectorBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewTopologyLabelSelectorBuilder(writer io.Writer) *TopologyLabelSelectorBuilder {
	return &TopologyLabelSelectorBuilder{
		writer: writer,
	}
}
func (x *TopologyLabelSelectorBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *TopologyLabelSelectorBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TopologyLabelSelectorBuilder) AddValues(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
