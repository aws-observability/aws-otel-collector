// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tracer_payload.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// TraceChunk represents a list of spans with the same trace ID. In other words, a chunk of a trace.
type TraceChunk struct {
	// priority specifies sampling priority of the trace.
	Priority int32 `protobuf:"varint,1,opt,name=priority,proto3" json:"priority" msg:"priority"`
	// origin specifies origin product ("lambda", "rum", etc.) of the trace.
	Origin string `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin" msg:"origin"`
	// spans specifies list of containing spans.
	Spans []*Span `protobuf:"bytes,3,rep,name=spans,proto3" json:"spans" msg:"spans"`
	// tags specifies tags common in all `spans`.
	Tags map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags" msg:"tags" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// droppedTrace specifies whether the trace was dropped by samplers or not.
	DroppedTrace bool `protobuf:"varint,5,opt,name=droppedTrace,proto3" json:"dropped_trace" msg:"dropped_trace"`
}

func (m *TraceChunk) Reset()         { *m = TraceChunk{} }
func (m *TraceChunk) String() string { return proto.CompactTextString(m) }
func (*TraceChunk) ProtoMessage()    {}
func (*TraceChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02adc68a2cbcd51, []int{0}
}
func (m *TraceChunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceChunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceChunk.Merge(m, src)
}
func (m *TraceChunk) XXX_Size() int {
	return m.Size()
}
func (m *TraceChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceChunk.DiscardUnknown(m)
}

var xxx_messageInfo_TraceChunk proto.InternalMessageInfo

func (m *TraceChunk) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TraceChunk) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *TraceChunk) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *TraceChunk) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TraceChunk) GetDroppedTrace() bool {
	if m != nil {
		return m.DroppedTrace
	}
	return false
}

// TracerPayload represents a payload the trace agent receives from tracers.
type TracerPayload struct {
	// containerID specifies the ID of the container where the tracer is running on.
	ContainerID string `protobuf:"bytes,1,opt,name=containerID,proto3" json:"container_id" msg:"container_id"`
	// languageName specifies language of the tracer.
	LanguageName string `protobuf:"bytes,2,opt,name=languageName,proto3" json:"language_name" msg:"language_name"`
	// languageVersion specifies language version of the tracer.
	LanguageVersion string `protobuf:"bytes,3,opt,name=languageVersion,proto3" json:"language_version" msg:"language_version"`
	// tracerVersion specifies version of the tracer.
	TracerVersion string `protobuf:"bytes,4,opt,name=tracerVersion,proto3" json:"tracer_version" msg:"tracer_version"`
	// runtimeID specifies V4 UUID representation of a tracer session.
	RuntimeID string `protobuf:"bytes,5,opt,name=runtimeID,proto3" json:"runtime_id" msg:"runtime_id"`
	// chunks specifies list of containing trace chunks.
	Chunks []*TraceChunk `protobuf:"bytes,6,rep,name=chunks,proto3" json:"chunks" msg:"chunks"`
	// tags specifies tags common in all `chunks`.
	Tags map[string]string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags" msg:"tags" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// env specifies `env` tag that set with the tracer.
	Env string `protobuf:"bytes,8,opt,name=env,proto3" json:"env" msg:"env"`
	// hostname specifies hostname of where the tracer is running.
	Hostname string `protobuf:"bytes,9,opt,name=hostname,proto3" json:"hostname" msg:"hostname"`
	// version specifies `version` tag that set with the tracer.
	AppVersion string `protobuf:"bytes,10,opt,name=appVersion,proto3" json:"app_version" msg:"app_version"`
}

func (m *TracerPayload) Reset()         { *m = TracerPayload{} }
func (m *TracerPayload) String() string { return proto.CompactTextString(m) }
func (*TracerPayload) ProtoMessage()    {}
func (*TracerPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02adc68a2cbcd51, []int{1}
}
func (m *TracerPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TracerPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TracerPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TracerPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TracerPayload.Merge(m, src)
}
func (m *TracerPayload) XXX_Size() int {
	return m.Size()
}
func (m *TracerPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_TracerPayload.DiscardUnknown(m)
}

var xxx_messageInfo_TracerPayload proto.InternalMessageInfo

func (m *TracerPayload) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *TracerPayload) GetLanguageName() string {
	if m != nil {
		return m.LanguageName
	}
	return ""
}

func (m *TracerPayload) GetLanguageVersion() string {
	if m != nil {
		return m.LanguageVersion
	}
	return ""
}

func (m *TracerPayload) GetTracerVersion() string {
	if m != nil {
		return m.TracerVersion
	}
	return ""
}

func (m *TracerPayload) GetRuntimeID() string {
	if m != nil {
		return m.RuntimeID
	}
	return ""
}

func (m *TracerPayload) GetChunks() []*TraceChunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func (m *TracerPayload) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TracerPayload) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *TracerPayload) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *TracerPayload) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*TraceChunk)(nil), "pb.TraceChunk")
	proto.RegisterMapType((map[string]string)(nil), "pb.TraceChunk.TagsEntry")
	proto.RegisterType((*TracerPayload)(nil), "pb.TracerPayload")
	proto.RegisterMapType((map[string]string)(nil), "pb.TracerPayload.TagsEntry")
}

func init() { proto.RegisterFile("tracer_payload.proto", fileDescriptor_f02adc68a2cbcd51) }

var fileDescriptor_f02adc68a2cbcd51 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x6b, 0xdb, 0x3e,
	0x18, 0xc7, 0xeb, 0xa6, 0xc9, 0x2f, 0x56, 0xda, 0xfe, 0x3a, 0x2d, 0x14, 0x93, 0x81, 0x15, 0xc4,
	0x28, 0x61, 0xb0, 0x14, 0xba, 0xc3, 0x4a, 0x4f, 0xc3, 0xcb, 0xd8, 0x0a, 0x63, 0x14, 0xad, 0x8c,
	0xdd, 0x82, 0x92, 0x78, 0xae, 0x69, 0x23, 0x0b, 0xd9, 0x0e, 0xf4, 0x3c, 0x76, 0xdf, 0xcb, 0xda,
	0xb1, 0xc7, 0x9d, 0xc4, 0x68, 0x6f, 0x3e, 0xfa, 0x15, 0x0c, 0x3f, 0xb2, 0x9d, 0x78, 0xa7, 0xb1,
	0x9b, 0x9e, 0xcf, 0xa3, 0xef, 0x37, 0x7e, 0xfe, 0x28, 0xa8, 0x9f, 0x28, 0x3e, 0xf7, 0xd5, 0x54,
	0xf2, 0xdb, 0x9b, 0x88, 0x2f, 0xc6, 0x52, 0x45, 0x49, 0x84, 0xb7, 0xe5, 0x6c, 0xf0, 0x3c, 0x08,
	0x93, 0xab, 0x74, 0x36, 0x9e, 0x47, 0xcb, 0xe3, 0x20, 0x0a, 0xa2, 0x63, 0x48, 0xcd, 0xd2, 0x2f,
	0x10, 0x41, 0x00, 0x27, 0x23, 0x19, 0xa0, 0x58, 0x72, 0x61, 0xce, 0xf4, 0x6b, 0x0b, 0xa1, 0xcb,
	0xc2, 0xf7, 0xf5, 0x55, 0x2a, 0xae, 0xf1, 0x19, 0xea, 0x4a, 0x15, 0x46, 0x2a, 0x4c, 0x6e, 0x1d,
	0x6b, 0x68, 0x8d, 0xda, 0x9e, 0x9b, 0x69, 0x52, 0xb3, 0x5c, 0x93, 0xfd, 0x65, 0x1c, 0x9c, 0xd1,
	0x0a, 0x50, 0x56, 0xe7, 0xf0, 0x09, 0xea, 0x44, 0x2a, 0x0c, 0x42, 0xe1, 0x6c, 0x0f, 0xad, 0x91,
	0xed, 0x0d, 0x32, 0x4d, 0x4a, 0x92, 0x6b, 0xb2, 0x0b, 0x3a, 0x13, 0x52, 0x56, 0x72, 0x7c, 0x8a,
	0xda, 0xc5, 0xc7, 0xc4, 0x4e, 0x6b, 0xd8, 0x1a, 0xf5, 0x4e, 0xba, 0x63, 0x39, 0x1b, 0x7f, 0x94,
	0x5c, 0x78, 0x4e, 0xa6, 0x89, 0x49, 0xe5, 0x9a, 0xf4, 0x40, 0x0b, 0x11, 0x65, 0x86, 0xe2, 0x09,
	0xda, 0x49, 0x78, 0x10, 0x3b, 0x3b, 0x20, 0x74, 0x0a, 0xe1, 0xba, 0x8e, 0xf1, 0x25, 0x0f, 0xe2,
	0x37, 0x22, 0x51, 0xb7, 0xde, 0x61, 0xa6, 0x09, 0xdc, 0xcc, 0x35, 0x41, 0xe0, 0x53, 0x04, 0x94,
	0x01, 0xc3, 0xef, 0xd1, 0xee, 0x42, 0x45, 0x52, 0xfa, 0x0b, 0x10, 0x3b, 0xed, 0xa1, 0x35, 0xea,
	0x7a, 0xa3, 0x4c, 0x93, 0xbd, 0x92, 0x4f, 0xa1, 0xeb, 0xb9, 0x26, 0x8f, 0x41, 0xdc, 0xa0, 0x94,
	0x35, 0xd4, 0x83, 0x97, 0xc8, 0xae, 0x7f, 0x18, 0x1f, 0xa0, 0xd6, 0xb5, 0x6f, 0xba, 0x68, 0xb3,
	0xe2, 0x88, 0xfb, 0xa8, 0xbd, 0xe2, 0x37, 0xa9, 0x6f, 0xfa, 0xc3, 0x4c, 0x70, 0xb6, 0x7d, 0x6a,
	0xd1, 0x6f, 0x1d, 0xb4, 0x07, 0x16, 0xea, 0xc2, 0x0c, 0x17, 0xbf, 0x43, 0xbd, 0x79, 0x24, 0x12,
	0x1e, 0x0a, 0x5f, 0x9d, 0x4f, 0x8c, 0x8b, 0x77, 0x94, 0x69, 0xb2, 0x5b, 0xe3, 0x69, 0xb8, 0xc8,
	0x35, 0xc1, 0xf0, 0x59, 0x9b, 0x90, 0xb2, 0x4d, 0x69, 0x51, 0xe2, 0x0d, 0x17, 0x41, 0xca, 0x03,
	0xff, 0x03, 0x5f, 0x96, 0x3f, 0x6e, 0x4a, 0xac, 0xf8, 0x54, 0xf0, 0xe5, 0xba, 0xc4, 0x06, 0xa5,
	0xac, 0xa1, 0xc6, 0x9f, 0xd1, 0xff, 0x55, 0xfc, 0xc9, 0x57, 0x71, 0x18, 0x09, 0xa7, 0x05, 0x86,
	0xe3, 0x4c, 0x93, 0x83, 0x5a, 0xba, 0x32, 0xb9, 0x5c, 0x93, 0xc3, 0xa6, 0x67, 0x99, 0xa0, 0xec,
	0x4f, 0x1b, 0x7c, 0x81, 0xf6, 0xcc, 0x82, 0x57, 0xbe, 0x3b, 0xe0, 0xfb, 0x2c, 0xd3, 0x64, 0xbf,
	0xdc, 0xfc, 0xb5, 0x6b, 0xdf, 0x4c, 0xb2, 0x81, 0x29, 0x6b, 0x1a, 0xe0, 0x57, 0xc8, 0x56, 0xa9,
	0x48, 0xc2, 0xa5, 0x7f, 0x3e, 0x81, 0xc9, 0xda, 0x1e, 0xcd, 0x34, 0x41, 0x25, 0x34, 0xfd, 0x3b,
	0x00, 0xa7, 0x35, 0xa2, 0x6c, 0x2d, 0xc2, 0x1e, 0xea, 0xcc, 0x8b, 0x7d, 0x8a, 0x9d, 0x0e, 0xac,
	0xd9, 0x7e, 0x73, 0xcd, 0xcc, 0x8a, 0x9b, 0x1b, 0xf5, 0x8a, 0x9b, 0x90, 0xb2, 0x92, 0xe3, 0xb7,
	0xe5, 0xa2, 0xfe, 0x07, 0x0e, 0x4f, 0x6a, 0x87, 0x6a, 0xd4, 0x7f, 0xbd, 0xab, 0x47, 0xa8, 0xe5,
	0x8b, 0x95, 0xd3, 0x85, 0x42, 0xfa, 0x99, 0x26, 0x45, 0x98, 0x6b, 0x62, 0xc3, 0x4d, 0x5f, 0xac,
	0x28, 0x2b, 0x48, 0xf1, 0x86, 0xaf, 0xa2, 0x38, 0x29, 0xa6, 0xe7, 0xd8, 0x70, 0x19, 0xde, 0x70,
	0xc5, 0xea, 0x37, 0x5c, 0x01, 0xca, 0xea, 0x1c, 0x9e, 0x20, 0xc4, 0xa5, 0xac, 0x26, 0x80, 0x40,
	0xfd, 0x34, 0xd3, 0xa4, 0xc7, 0xa5, 0xdc, 0x68, 0xff, 0x23, 0x30, 0xd8, 0x60, 0x94, 0x6d, 0xe8,
	0xfe, 0xf9, 0x1d, 0x78, 0xce, 0x8f, 0x7b, 0xd7, 0xba, 0xbb, 0x77, 0xad, 0x5f, 0xf7, 0xae, 0xf5,
	0xfd, 0xc1, 0xdd, 0xba, 0x7b, 0x70, 0xb7, 0x7e, 0x3e, 0xb8, 0x5b, 0xb3, 0x0e, 0xfc, 0x5d, 0xbd,
	0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x17, 0x87, 0x13, 0xf6, 0x05, 0x05, 0x00, 0x00,
}

func (m *TraceChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceChunk) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Priority != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTracerPayload(dAtA, i, uint64(m.Priority))
	}
	if len(m.Origin) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTracerPayload(dAtA, i, uint64(len(m.Origin)))
		i += copy(dAtA[i:], m.Origin)
	}
	if len(m.Spans) > 0 {
		for _, msg := range m.Spans {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintTracerPayload(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Tags) > 0 {
		for k, _ := range m.Tags {
			dAtA[i] = 0x22
			i++
			v := m.Tags[k]
			mapSize := 1 + len(k) + sovTracerPayload(uint64(len(k))) + 1 + len(v) + sovTracerPayload(uint64(len(v)))
			i = encodeVarintTracerPayload(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintTracerPayload(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintTracerPayload(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.DroppedTrace {
		dAtA[i] = 0x28
		i++
		if m.DroppedTrace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *TracerPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TracerPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ContainerID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTracerPayload(dAtA, i, uint64(len(m.ContainerID)))
		i += copy(dAtA[i:], m.ContainerID)
	}
	if len(m.LanguageName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTracerPayload(dAtA, i, uint64(len(m.LanguageName)))
		i += copy(dAtA[i:], m.LanguageName)
	}
	if len(m.LanguageVersion) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTracerPayload(dAtA, i, uint64(len(m.LanguageVersion)))
		i += copy(dAtA[i:], m.LanguageVersion)
	}
	if len(m.TracerVersion) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTracerPayload(dAtA, i, uint64(len(m.TracerVersion)))
		i += copy(dAtA[i:], m.TracerVersion)
	}
	if len(m.RuntimeID) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTracerPayload(dAtA, i, uint64(len(m.RuntimeID)))
		i += copy(dAtA[i:], m.RuntimeID)
	}
	if len(m.Chunks) > 0 {
		for _, msg := range m.Chunks {
			dAtA[i] = 0x32
			i++
			i = encodeVarintTracerPayload(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Tags) > 0 {
		for k, _ := range m.Tags {
			dAtA[i] = 0x3a
			i++
			v := m.Tags[k]
			mapSize := 1 + len(k) + sovTracerPayload(uint64(len(k))) + 1 + len(v) + sovTracerPayload(uint64(len(v)))
			i = encodeVarintTracerPayload(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintTracerPayload(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintTracerPayload(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Env) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintTracerPayload(dAtA, i, uint64(len(m.Env)))
		i += copy(dAtA[i:], m.Env)
	}
	if len(m.Hostname) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintTracerPayload(dAtA, i, uint64(len(m.Hostname)))
		i += copy(dAtA[i:], m.Hostname)
	}
	if len(m.AppVersion) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintTracerPayload(dAtA, i, uint64(len(m.AppVersion)))
		i += copy(dAtA[i:], m.AppVersion)
	}
	return i, nil
}

func encodeVarintTracerPayload(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TraceChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Priority != 0 {
		n += 1 + sovTracerPayload(uint64(m.Priority))
	}
	l = len(m.Origin)
	if l > 0 {
		n += 1 + l + sovTracerPayload(uint64(l))
	}
	if len(m.Spans) > 0 {
		for _, e := range m.Spans {
			l = e.Size()
			n += 1 + l + sovTracerPayload(uint64(l))
		}
	}
	if len(m.Tags) > 0 {
		for k, v := range m.Tags {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTracerPayload(uint64(len(k))) + 1 + len(v) + sovTracerPayload(uint64(len(v)))
			n += mapEntrySize + 1 + sovTracerPayload(uint64(mapEntrySize))
		}
	}
	if m.DroppedTrace {
		n += 2
	}
	return n
}

func (m *TracerPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContainerID)
	if l > 0 {
		n += 1 + l + sovTracerPayload(uint64(l))
	}
	l = len(m.LanguageName)
	if l > 0 {
		n += 1 + l + sovTracerPayload(uint64(l))
	}
	l = len(m.LanguageVersion)
	if l > 0 {
		n += 1 + l + sovTracerPayload(uint64(l))
	}
	l = len(m.TracerVersion)
	if l > 0 {
		n += 1 + l + sovTracerPayload(uint64(l))
	}
	l = len(m.RuntimeID)
	if l > 0 {
		n += 1 + l + sovTracerPayload(uint64(l))
	}
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.Size()
			n += 1 + l + sovTracerPayload(uint64(l))
		}
	}
	if len(m.Tags) > 0 {
		for k, v := range m.Tags {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTracerPayload(uint64(len(k))) + 1 + len(v) + sovTracerPayload(uint64(len(v)))
			n += mapEntrySize + 1 + sovTracerPayload(uint64(mapEntrySize))
		}
	}
	l = len(m.Env)
	if l > 0 {
		n += 1 + l + sovTracerPayload(uint64(l))
	}
	l = len(m.Hostname)
	if l > 0 {
		n += 1 + l + sovTracerPayload(uint64(l))
	}
	l = len(m.AppVersion)
	if l > 0 {
		n += 1 + l + sovTracerPayload(uint64(l))
	}
	return n
}

func sovTracerPayload(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTracerPayload(x uint64) (n int) {
	return sovTracerPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TraceChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracerPayload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TraceChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Origin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spans = append(m.Spans, &Span{})
			if err := m.Spans[len(m.Spans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tags == nil {
				m.Tags = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTracerPayload
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTracerPayload
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTracerPayload
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTracerPayload
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTracerPayload
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTracerPayload
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthTracerPayload
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTracerPayload(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTracerPayload
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Tags[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DroppedTrace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DroppedTrace = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTracerPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TracerPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracerPayload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TracerPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TracerPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LanguageName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LanguageName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LanguageVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LanguageVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TracerVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TracerVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RuntimeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks, &TraceChunk{})
			if err := m.Chunks[len(m.Chunks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tags == nil {
				m.Tags = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTracerPayload
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTracerPayload
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTracerPayload
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTracerPayload
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTracerPayload
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTracerPayload
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthTracerPayload
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTracerPayload(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTracerPayload
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Tags[mapkey] = mapvalue
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hostname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hostname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracerPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracerPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTracerPayload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTracerPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTracerPayload
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTracerPayload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTracerPayload
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTracerPayload
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTracerPayload
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTracerPayload(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTracerPayload
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTracerPayload = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTracerPayload   = fmt.Errorf("proto: integer overflow")
)