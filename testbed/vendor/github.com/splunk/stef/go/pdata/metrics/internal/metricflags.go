package internal

import "github.com/splunk/stef/go/otel/oteltef"

type MetricFlags struct {
	Temporality oteltef.AggregationTemporality
	Monotonic   bool
}
