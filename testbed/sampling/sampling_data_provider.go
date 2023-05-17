package sampling

import (
	"encoding/binary"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// samplingDataProvider is an implementation of DataProvider for use in trace tests
// This data provider will generator some traces that fall within the provided option set
// and some data that does not. This is useful for testing filter and sampling processors.
type samplingDataProvider struct {
	// Used by the load generator to report how many itens were created
	dataItemsGenerated *atomic.Uint64
	// Number of traces to generate in a run
	numTracesToGenerate int
	// Number of spans per trace to generate in a run
	spansPerTrace   int
	traceIDSequence atomic.Uint64
	// Total traces generated
	tracesGenerated int
	// we will use this to ensure that the received spans matches the expected sampled spans
	sampledSpansGenerated atomic.Uint64
	// Controls if a trace should be sampled
	shouldSample shouldSampleFunc
	// Customize spans that need to be sampled so that they can properly filtered
	sampledSpanCustomizer sampledSpanCustomizerFunc
	// Stores all the trances that were sampled
	expectedSampledTraces []ptrace.Traces
}

type sampledSpanCustomizerFunc func(span ptrace.Span)

type shouldSampleFunc func(traceIDSequence uint64) bool

// SamplingDataProviderOption defines a PollutedDataProvider option.
type SamplingDataProviderOption func(t *samplingDataProvider)

// WithShouldSampleFunction allows you to control if a trace should be sampled.
func WithSamplingFunction(shouldSample shouldSampleFunc) SamplingDataProviderOption {
	return func(tc *samplingDataProvider) {
		tc.shouldSample = shouldSample
	}
}

// WithSampledSpanCustomizer allows you to customize the spans that are sampled.
func WithSampledSpanCustomizer(customizer sampledSpanCustomizerFunc) SamplingDataProviderOption {
	return func(tc *samplingDataProvider) {
		tc.sampledSpanCustomizer = customizer
	}
}

// WithExpectedSampledTraces allows you to specify the number of spans per trace
func WithSpansPerTrace(spansPerTrace int) SamplingDataProviderOption {
	return func(tc *samplingDataProvider) {
		tc.spansPerTrace = spansPerTrace
	}
}

// WithExpectedSampledTraces allows you to specify the total number of traces generated in a run
func WithNumTracesToGenerate(numTracesToGenerate int) SamplingDataProviderOption {
	return func(tc *samplingDataProvider) {
		tc.numTracesToGenerate = numTracesToGenerate
	}
}

// NewSamplingDataProvider creates a new instance of samplingDataProvider
func NewSamplingDataProvider(opts ...SamplingDataProviderOption) *samplingDataProvider {
	sdp := samplingDataProvider{
		// Defaults. These are arbitrary values.
		spansPerTrace:       6,
		numTracesToGenerate: 50,
	}

	for _, opt := range opts {
		opt(&sdp)
	}

	return &sdp
}

// SetLoadGeneratorCounters sets the data providers items generated field to have a ref to the load generators counter
func (sdp *samplingDataProvider) SetLoadGeneratorCounters(dataItemsGenerated *atomic.Uint64) {
	sdp.dataItemsGenerated = dataItemsGenerated
}

// GenerateTraces the boolean in this return reflects whether we are done generating traces or not.
func (sdp *samplingDataProvider) GenerateTraces() (ptrace.Traces, bool) {
	if sdp.tracesGenerated >= sdp.numTracesToGenerate {
		return ptrace.NewTraces(), true
	}
	sdp.tracesGenerated += 1

	traceData := sdp.generateSampledTrace()

	return traceData, false

}

func (sdp *samplingDataProvider) ToBeSampledSpansGenerated() uint64 {
	return sdp.sampledSpansGenerated.Load()
}

func (sdp *samplingDataProvider) generateSampledTrace() ptrace.Traces {
	traceData := ptrace.NewTraces()

	spans := traceData.ResourceSpans().AppendEmpty().ScopeSpans().AppendEmpty().Spans()
	spans.EnsureCapacity(sdp.spansPerTrace)

	traceID := sdp.traceIDSequence.Add(1)
	shouldSample := sdp.shouldSample(traceID)
	// TODO: should we be using ItemsPerBatch to determine how many spans are in a trace?
	for i := 0; i < sdp.spansPerTrace; i++ {

		startTime := time.Now()
		endTime := startTime.Add(time.Millisecond)

		spanID := sdp.dataItemsGenerated.Add(1)

		span := spans.AppendEmpty()

		// Create a span.
		span.SetTraceID(uInt64ToTraceID(0, traceID))
		span.SetSpanID(uInt64ToSpanID(spanID))
		span.SetName("dirty-trace-span")
		span.SetKind(ptrace.SpanKindClient)
		attrs := span.Attributes()

		attrs.PutInt("dirty_trace.span_seq_num", int64(spanID))
		attrs.PutInt("dirty_trace.trace_seq_num", int64(traceID))

		span.SetStartTimestamp(pcommon.NewTimestampFromTime(startTime))
		span.SetEndTimestamp(pcommon.NewTimestampFromTime(endTime))

		if shouldSample {
			sdp.sampledSpansGenerated.Add(1)
			sdp.sampledSpanCustomizer(span)
		}
	}

	if shouldSample {
		sdp.expectedSampledTraces = append(sdp.expectedSampledTraces, traceData)
	}

	return traceData
}

// GenerateMetrics TODO
func (sdp *samplingDataProvider) GenerateMetrics() (pmetric.Metrics, bool) {
	return pmetric.NewMetrics(), false
}

// GenerateLogs TODO
func (sdp *samplingDataProvider) GenerateLogs() (plog.Logs, bool) {
	return plog.NewLogs(), true
}

// uInt64ToTraceID converts the pair of uint64 representation of a TraceID to pcommon.TraceID.
func uInt64ToTraceID(high, low uint64) pcommon.TraceID {
	traceID := [16]byte{}
	binary.BigEndian.PutUint64(traceID[:8], high)
	binary.BigEndian.PutUint64(traceID[8:], low)
	return traceID
}

// uInt64ToSpanID converts the uint64 representation of a SpanID to pcommon.SpanID.
func uInt64ToSpanID(id uint64) pcommon.SpanID {
	spanID := [8]byte{}
	binary.BigEndian.PutUint64(spanID[:], id)
	return pcommon.SpanID(spanID)
}
