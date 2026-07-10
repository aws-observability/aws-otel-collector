// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineClickhouseDestinationFormat Insert format for events sent to ClickHouse.
// - `json_each_row`: Maps event fields to columns by name (ClickHouse `JSONEachRow`).
// - `json_as_object`: Inserts each event into a single `Object('json')` / `JSON` column (ClickHouse `JSONAsObject`).
// - `json_as_string`: Inserts each event into a single `String`-typed column as raw JSON (ClickHouse `JSONAsString`).
// - `arrow_stream`: Batches events using Apache Arrow IPC streaming format. Requires `batch_encoding`.
type ObservabilityPipelineClickhouseDestinationFormat string

// List of ObservabilityPipelineClickhouseDestinationFormat.
const (
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONFORMAT_JSON_EACH_ROW  ObservabilityPipelineClickhouseDestinationFormat = "json_each_row"
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONFORMAT_JSON_AS_OBJECT ObservabilityPipelineClickhouseDestinationFormat = "json_as_object"
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONFORMAT_JSON_AS_STRING ObservabilityPipelineClickhouseDestinationFormat = "json_as_string"
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONFORMAT_ARROW_STREAM   ObservabilityPipelineClickhouseDestinationFormat = "arrow_stream"
)

var allowedObservabilityPipelineClickhouseDestinationFormatEnumValues = []ObservabilityPipelineClickhouseDestinationFormat{
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONFORMAT_JSON_EACH_ROW,
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONFORMAT_JSON_AS_OBJECT,
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONFORMAT_JSON_AS_STRING,
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONFORMAT_ARROW_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineClickhouseDestinationFormat) GetAllowedValues() []ObservabilityPipelineClickhouseDestinationFormat {
	return allowedObservabilityPipelineClickhouseDestinationFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineClickhouseDestinationFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineClickhouseDestinationFormat(value)
	return nil
}

// NewObservabilityPipelineClickhouseDestinationFormatFromValue returns a pointer to a valid ObservabilityPipelineClickhouseDestinationFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineClickhouseDestinationFormatFromValue(v string) (*ObservabilityPipelineClickhouseDestinationFormat, error) {
	ev := ObservabilityPipelineClickhouseDestinationFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineClickhouseDestinationFormat: valid values are %v", v, allowedObservabilityPipelineClickhouseDestinationFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineClickhouseDestinationFormat) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineClickhouseDestinationFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineClickhouseDestinationFormat value.
func (v ObservabilityPipelineClickhouseDestinationFormat) Ptr() *ObservabilityPipelineClickhouseDestinationFormat {
	return &v
}
