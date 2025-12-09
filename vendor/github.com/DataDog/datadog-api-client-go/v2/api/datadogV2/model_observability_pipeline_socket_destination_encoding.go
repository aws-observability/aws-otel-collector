// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketDestinationEncoding Encoding format for log events.
type ObservabilityPipelineSocketDestinationEncoding string

// List of ObservabilityPipelineSocketDestinationEncoding.
const (
	OBSERVABILITYPIPELINESOCKETDESTINATIONENCODING_JSON        ObservabilityPipelineSocketDestinationEncoding = "json"
	OBSERVABILITYPIPELINESOCKETDESTINATIONENCODING_RAW_MESSAGE ObservabilityPipelineSocketDestinationEncoding = "raw_message"
)

var allowedObservabilityPipelineSocketDestinationEncodingEnumValues = []ObservabilityPipelineSocketDestinationEncoding{
	OBSERVABILITYPIPELINESOCKETDESTINATIONENCODING_JSON,
	OBSERVABILITYPIPELINESOCKETDESTINATIONENCODING_RAW_MESSAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSocketDestinationEncoding) GetAllowedValues() []ObservabilityPipelineSocketDestinationEncoding {
	return allowedObservabilityPipelineSocketDestinationEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSocketDestinationEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSocketDestinationEncoding(value)
	return nil
}

// NewObservabilityPipelineSocketDestinationEncodingFromValue returns a pointer to a valid ObservabilityPipelineSocketDestinationEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSocketDestinationEncodingFromValue(v string) (*ObservabilityPipelineSocketDestinationEncoding, error) {
	ev := ObservabilityPipelineSocketDestinationEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSocketDestinationEncoding: valid values are %v", v, allowedObservabilityPipelineSocketDestinationEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSocketDestinationEncoding) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSocketDestinationEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSocketDestinationEncoding value.
func (v ObservabilityPipelineSocketDestinationEncoding) Ptr() *ObservabilityPipelineSocketDestinationEncoding {
	return &v
}
