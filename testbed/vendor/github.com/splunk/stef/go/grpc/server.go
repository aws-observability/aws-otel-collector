package stefgrpc

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/splunk/stef/go/pkg/schema"

	"github.com/splunk/stef/go/grpc/internal"
	"github.com/splunk/stef/go/grpc/stef_proto"
	"github.com/splunk/stef/go/grpc/types"
)

type grpcMsgSource interface {
	// recvMsg receives a STEF/gRPC message containing STEF bytes.
	recvMsg() (tefBytes []byte, isEndOfChunk bool, err error)
}

type GrpcReader interface {
	Read(p []byte) (n int, err error)
	Stats() GrpcReaderStats
}

type GrpcReaderStats struct {
	MessagesReceived uint64
	BytesReceived    uint64
}

type chunkAssembler struct {
	source    grpcMsgSource
	buf       []byte
	readIndex int

	statsMux sync.RWMutex
	stats    GrpcReaderStats
}

var _ io.Reader = (*chunkAssembler)(nil)

func (g *chunkAssembler) recvMsg() (chunkBytes []byte, err error) {
	var chunkBuf []byte
	for {
		bytes, isEndOfChunk, err := g.source.recvMsg()
		if err != nil {
			return nil, err
		}
		if chunkBuf == nil {
			chunkBuf = bytes
		} else {
			chunkBuf = append(chunkBuf, bytes...)
		}

		if isEndOfChunk {
			// These bytes are ending a chunk. Return what is accumulated.
			g.statsMux.Lock()
			g.stats.MessagesReceived++
			g.stats.BytesReceived += uint64(len(chunkBuf))
			g.statsMux.Unlock()
			return chunkBuf, nil
		}
	}
}

func (g *chunkAssembler) Read(p []byte) (n int, err error) {
	if g.readIndex >= len(g.buf) {
		data, err := g.recvMsg()
		if err != nil {
			return 0, err
		}
		g.buf = data
		g.readIndex = 0
	}
	n = copy(p, g.buf[g.readIndex:])
	g.readIndex += n
	return n, nil
}

func (g *chunkAssembler) Stats() GrpcReaderStats {
	g.statsMux.RLock()
	defer g.statsMux.RUnlock()
	return g.stats
}

func newChunkAssembler(source grpcMsgSource) GrpcReader {
	return &chunkAssembler{source: source}
}

type grpcChunkSource struct {
	serverStream     stef_proto.STEFDestination_StreamServer
	message          stef_proto.STEFServerMessage
	serverResponse   *stef_proto.STEFServerMessage_Response
	messagesReceived uint64
}

func newGrpcChunkSource(serverStream stef_proto.STEFDestination_StreamServer) *grpcChunkSource {
	s := &grpcChunkSource{
		serverStream: serverStream,
	}
	s.serverResponse = &stef_proto.STEFServerMessage_Response{}
	s.message = stef_proto.STEFServerMessage{
		Message: s.serverResponse,
	}

	return s
}

func (r *grpcChunkSource) recvMsg() (tefBytes []byte, isEndOfChunk bool, err error) {
	response, err := r.serverStream.Recv()
	if err != nil {
		return nil, false, err
	}
	r.messagesReceived++
	return response.StefBytes, response.IsEndOfChunk, nil
}

func (r *grpcChunkSource) SendDataResponse(response *stef_proto.STEFDataResponse) error {
	r.serverResponse.Response = response
	return r.serverStream.Send(&r.message)
}

type StreamServer struct {
	stef_proto.UnimplementedSTEFDestinationServer

	logger       types.Logger
	serverSchema *schema.WireSchema
	maxDictBytes uint64
	callbacks    Callbacks
}

var _ stef_proto.STEFDestinationServer = (*StreamServer)(nil)

// Callbacks is a set of callbacks that StreamServer calls.
type Callbacks struct {
	// OnStream is called when a new stream is opened to the server.
	// The reader is ready to be passed as an input to a STEF reader.
	// If the callback returns an error, the stream will be closed.
	// If the callback is nil, the default behavior is to close the stream immediately.
	OnStream func(reader GrpcReader, stream STEFStream) error
}

func defaultOnStream(GrpcReader, STEFStream) error {
	return nil
}

// SetDefaults ensures that callbacks that are not initialized by the
// user are set to default functions.
func (c *Callbacks) SetDefaults() {
	if c.OnStream == nil {
		c.OnStream = defaultOnStream
	}
}

// STEFStream represents a STEF-over-gRPC stream.
type STEFStream interface {
	// SendDataResponse sends a response to the client. Must not be called concurrently.
	SendDataResponse(response *stef_proto.STEFDataResponse) error
}

// ServerSettings contains configuration options for creating a new STEF-over-gRPC stream server.
type ServerSettings struct {
	// Logger for logging server events and errors. If nil, a no-op logger will be used.
	// The logger should be safe for concurrent use across multiple goroutines.
	Logger types.Logger

	// ServerSchema defines the wire schema that the server will use for STEF message decoding.
	// This schema must be compatible with the data being received from clients. If the schema
	// is incompatible the client will refuse to send data. Must not be nil.
	ServerSchema *schema.WireSchema

	// MaxDictBytes sets the maximum size in bytes for dictionary data that the server will accept
	// per STEF stream. This limit helps prevent memory exhaustion.
	// A value of 0 means no limit is enforced and is not recommended.
	MaxDictBytes uint64

	// Callbacks contains the callback functions that will be invoked during stream processing.
	Callbacks Callbacks
}

// NewStreamServer creates a new STEF-over-gRPC stream server with the given settings.
func NewStreamServer(settings ServerSettings) *StreamServer {
	if settings.Logger == nil {
		settings.Logger = internal.NopLogger{}
	}
	s := &StreamServer{
		logger:       settings.Logger,
		serverSchema: settings.ServerSchema,
		maxDictBytes: settings.MaxDictBytes,
		callbacks:    settings.Callbacks,
	}
	s.callbacks.SetDefaults()
	return s
}

func (s *StreamServer) Stream(server stef_proto.STEFDestination_StreamServer) error {
	// Receive the first message from the client.
	clientMsg, err := server.Recv()
	if err != nil {
		return fmt.Errorf("failed to receive a message from the client: %w", err)
	}
	if clientMsg.FirstMessage == nil {
		return fmt.Errorf("FirstMessage is nil")
	}
	if clientMsg.FirstMessage.RootStructName == "" {
		return fmt.Errorf("RootStructName is unspecified")
	}

	// Send capabilities message to the client.

	// Prepare the schema bytes.
	var schemaBytes bytes.Buffer
	err = s.serverSchema.Serialize(&schemaBytes)
	if err != nil {
		return fmt.Errorf("could not marshal server schema: %w", err)
	}

	message := stef_proto.STEFServerMessage{
		Message: &stef_proto.STEFServerMessage_Capabilities{
			Capabilities: &stef_proto.STEFDestinationCapabilities{
				DictionaryLimits: &stef_proto.STEFDictionaryLimits{MaxDictBytes: s.maxDictBytes},
				Schema:           schemaBytes.Bytes(),
			},
		},
	}
	err = server.Send(&message)
	if err != nil {
		s.logger.Errorf(context.Background(), "cannot send message to the client: %v", err)
		return err
	}

	grpcStream := newGrpcChunkSource(server)
	reader := newChunkAssembler(grpcStream)
	return s.callbacks.OnStream(reader, grpcStream)
}
