// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkHecDestinationEncoding Encoding format for log events.
type ObservabilityPipelineSplunkHecDestinationEncoding string

// List of ObservabilityPipelineSplunkHecDestinationEncoding.
const (
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONENCODING_JSON        ObservabilityPipelineSplunkHecDestinationEncoding = "json"
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONENCODING_RAW_MESSAGE ObservabilityPipelineSplunkHecDestinationEncoding = "raw_message"
)

var allowedObservabilityPipelineSplunkHecDestinationEncodingEnumValues = []ObservabilityPipelineSplunkHecDestinationEncoding{
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONENCODING_JSON,
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONENCODING_RAW_MESSAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSplunkHecDestinationEncoding) GetAllowedValues() []ObservabilityPipelineSplunkHecDestinationEncoding {
	return allowedObservabilityPipelineSplunkHecDestinationEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSplunkHecDestinationEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSplunkHecDestinationEncoding(value)
	return nil
}

// NewObservabilityPipelineSplunkHecDestinationEncodingFromValue returns a pointer to a valid ObservabilityPipelineSplunkHecDestinationEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSplunkHecDestinationEncodingFromValue(v string) (*ObservabilityPipelineSplunkHecDestinationEncoding, error) {
	ev := ObservabilityPipelineSplunkHecDestinationEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSplunkHecDestinationEncoding: valid values are %v", v, allowedObservabilityPipelineSplunkHecDestinationEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSplunkHecDestinationEncoding) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSplunkHecDestinationEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSplunkHecDestinationEncoding value.
func (v ObservabilityPipelineSplunkHecDestinationEncoding) Ptr() *ObservabilityPipelineSplunkHecDestinationEncoding {
	return &v
}
