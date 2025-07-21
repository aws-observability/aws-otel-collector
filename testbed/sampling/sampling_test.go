package sampling

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/aws-observability/aws-otel-collector/pkg/defaultcomponents"
	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/correctnesstests"
	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type NoopTestSummary struct{}

// NoopTestSummary implements testbed.TestResultsSummary
func (n *NoopTestSummary) Init(resultsDir string) {}

// Add results for one test.
func (n *NoopTestSummary) Add(testName string, result interface{}) {}

// Save the total results and close the file.
func (n *NoopTestSummary) Save() {}

var testResults testbed.TestResultsSummary = &NoopTestSummary{}

func createProcessorsConfiguration(policies string) []correctnesstests.ProcessorNameAndConfigBody {
	// TODO: It is painful to create configurations because of the go formatting.
	// Ideally it should be possible to create configurations from a file.
	processors := []correctnesstests.ProcessorNameAndConfigBody{
		{
			Name: "groupbytrace",
			Body: `
  groupbytrace:
`,
		},
		{
			Name: "tail_sampling",
			Body: fmt.Sprintf(`
  tail_sampling:
    decision_wait: 3s
    policies: %s`, policies),
		},
		{
			Name: "batch",
			Body: `
  batch:
    send_batch_size: 1024
`,
		},
	}
	return processors
}

func TestTailSamplingData(t *testing.T) {

	tests := []struct {
		name string
		// Content of the policies in the tail sampling processor
		policies string
		// Change the spans so that they be sampled by the policies
		spanCustomizer sampledSpanCustomizerFunc
	}{
		{
			"Attributes Sample",
			"[{name: attr_sample, type: string_attribute, string_attribute: {key: key, values: value}}]",
			func(span ptrace.Span) {
				span.Attributes().PutStr("key", "value")
			},
		},
		{
			"Attributes Sample regex",
			"[{name: attr_sample, type: string_attribute, string_attribute: {key: key, values: [value.*], enabled_regex_matching: true}}]",
			func(span ptrace.Span) {
				span.Attributes().PutStr("key", "valuefoo")
			},
		},
		{
			"Latency",
			"[{name: latency_sample, type: latency,  latency: {threshold_ms: 600}}]",
			func(span ptrace.Span) {
				startTime := time.Now()
				endTime := startTime.Add(610 * time.Millisecond)
				span.SetStartTimestamp(pcommon.NewTimestampFromTime(startTime))
				span.SetEndTimestamp(pcommon.NewTimestampFromTime(endTime))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			processors := createProcessorsConfiguration(test.policies)
			testWithSampledData(t, processors, test.spanCustomizer)
		})
	}
}

func testWithSampledData(
	t *testing.T,
	processors []correctnesstests.ProcessorNameAndConfigBody,
	customizer sampledSpanCustomizerFunc,
) {
	var resourceSpec testbed.ResourceSpec
	sender := correctnesstests.ConstructTraceSender(t, "otlp")
	receiver := correctnesstests.ConstructReceiver(t, "otlp")

	spanCustomizer := WithSampledSpanCustomizer(customizer)
	dataProvider := NewSamplingDataProvider(spanCustomizer)

	factories, err := defaultcomponents.Components()
	require.NoError(t, err, "default components resulted in: %v", err)

	runner := testbed.NewInProcessCollector(factories)

	validator := NewSamplingValidator(sender.ProtocolName(), receiver.ProtocolName(), dataProvider, t)

	config := correctnesstests.CreateConfigYaml(t, sender, receiver, nil, processors)

	log.Println(config)

	configCleanup, cfgErr := runner.PrepareConfig(t, config)

	require.NoError(t, cfgErr, "collector configuration resulted in: %v", cfgErr)

	defer configCleanup()

	tc := testbed.NewTestCase(
		t,
		dataProvider,
		sender,
		receiver,
		runner,
		validator,
		testResults,
		testbed.WithResourceLimits(resourceSpec),
	)
	defer tc.Stop()

	tc.EnableRecording()
	tc.StartBackend()
	tc.StartAgent()

	// the following parameters only affect the rate in which data is sent
	// Ultimately the data provider controls how much data to send using the
	// flag returned by GenerateTraces
	tc.StartLoad(testbed.LoadOptions{
		DataItemsPerSecond: 128,
		ItemsPerBatch:      1,
	})

	tc.Sleep(2 * time.Second)

	tc.StopLoad()

	tc.WaitForN(func() bool { return tc.MockBackend.DataItemsReceived() == dataProvider.sampledSpansGenerated() },
		10*time.Second, "all data items received")

	tc.StopAgent()
	tc.ValidateData()
}
