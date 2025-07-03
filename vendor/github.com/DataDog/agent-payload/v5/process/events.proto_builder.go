// THIS IS A GENERATED FILE
// DO NOT EDIT
package process

import (
	bytes "bytes"
	protowire "google.golang.org/protobuf/encoding/protowire"
	io "io"
)

type CollectorProcEventBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	hostBuilder         HostBuilder
	systemInfoBuilder   SystemInfoBuilder
	processEventBuilder ProcessEventBuilder
}

func NewCollectorProcEventBuilder(writer io.Writer) *CollectorProcEventBuilder {
	return &CollectorProcEventBuilder{
		writer: writer,
	}
}
func (x *CollectorProcEventBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *CollectorProcEventBuilder) SetHostname(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorProcEventBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcEventBuilder) SetInfo(cb func(w *SystemInfoBuilder)) {
	x.buf.Reset()
	x.systemInfoBuilder.writer = &x.buf
	x.systemInfoBuilder.scratch = x.scratch
	cb(&x.systemInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcEventBuilder) AddEvents(cb func(w *ProcessEventBuilder)) {
	x.buf.Reset()
	x.processEventBuilder.writer = &x.buf
	x.processEventBuilder.scratch = x.scratch
	cb(&x.processEventBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcEventBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorProcEventBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ProcessEventBuilder struct {
	writer             io.Writer
	buf                bytes.Buffer
	scratch            []byte
	commandBuilder     CommandBuilder
	processUserBuilder ProcessUserBuilder
	hostBuilder        HostBuilder
	processExecBuilder ProcessExecBuilder
	processExitBuilder ProcessExitBuilder
}

func NewProcessEventBuilder(writer io.Writer) *ProcessEventBuilder {
	return &ProcessEventBuilder{
		writer: writer,
	}
}
func (x *ProcessEventBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProcessEventBuilder) SetType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ProcessEventBuilder) SetCollectionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessEventBuilder) SetPid(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessEventBuilder) SetCommand(cb func(w *CommandBuilder)) {
	x.buf.Reset()
	x.commandBuilder.writer = &x.buf
	x.commandBuilder.scratch = x.scratch
	cb(&x.commandBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessEventBuilder) SetUser(cb func(w *ProcessUserBuilder)) {
	x.buf.Reset()
	x.processUserBuilder.writer = &x.buf
	x.processUserBuilder.scratch = x.scratch
	cb(&x.processUserBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessEventBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessEventBuilder) SetContainerId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ProcessEventBuilder) SetExec(cb func(w *ProcessExecBuilder)) {
	x.buf.Reset()
	x.processExecBuilder.writer = &x.buf
	x.processExecBuilder.scratch = x.scratch
	cb(&x.processExecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessEventBuilder) SetExit(cb func(w *ProcessExitBuilder)) {
	x.buf.Reset()
	x.processExitBuilder.writer = &x.buf
	x.processExitBuilder.scratch = x.scratch
	cb(&x.processExitBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x4a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ProcessExecBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewProcessExecBuilder(writer io.Writer) *ProcessExecBuilder {
	return &ProcessExecBuilder{
		writer: writer,
	}
}
func (x *ProcessExecBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProcessExecBuilder) SetForkTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessExecBuilder) SetExecTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ProcessExitBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewProcessExitBuilder(writer io.Writer) *ProcessExitBuilder {
	return &ProcessExitBuilder{
		writer: writer,
	}
}
func (x *ProcessExitBuilder) Reset(writer io.Writer) {
	x.buf.Reset()
	x.writer = writer
}
func (x *ProcessExitBuilder) SetExecTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessExitBuilder) SetExitTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessExitBuilder) SetExitCode(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
