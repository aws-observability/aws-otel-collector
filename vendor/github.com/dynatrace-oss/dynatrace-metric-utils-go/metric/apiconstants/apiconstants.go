package apiconstants

const (
	defaultOneAgentEndpoint = "http://localhost:14499/metrics/ingest"
	payloadLinesLimit       = 1000
)

// GetDefaultOneAgentEndpoint returns the default OneAgent metrics ingest endpoint.
// See the Dynatrace documentation (https://www.dynatrace.com/support/help/how-to-use-dynatrace/metrics/metric-ingestion/ingestion-methods/local-api/) for more information.
func GetDefaultOneAgentEndpoint() string {
	return defaultOneAgentEndpoint
}

// GetMetricPayloadLinesLimit returns the maximum number of lines per POST request accepted by the ingest endpoint.
func GetPayloadLinesLimit() int {
	return payloadLinesLimit
}
