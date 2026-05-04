// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpClientDestinationEncoding Encoding format for log events.
type ObservabilityPipelineHttpClientDestinationEncoding string

// List of ObservabilityPipelineHttpClientDestinationEncoding.
const (
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONENCODING_JSON ObservabilityPipelineHttpClientDestinationEncoding = "json"
)

var allowedObservabilityPipelineHttpClientDestinationEncodingEnumValues = []ObservabilityPipelineHttpClientDestinationEncoding{
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONENCODING_JSON,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineHttpClientDestinationEncoding) GetAllowedValues() []ObservabilityPipelineHttpClientDestinationEncoding {
	return allowedObservabilityPipelineHttpClientDestinationEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineHttpClientDestinationEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineHttpClientDestinationEncoding(value)
	return nil
}

// NewObservabilityPipelineHttpClientDestinationEncodingFromValue returns a pointer to a valid ObservabilityPipelineHttpClientDestinationEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineHttpClientDestinationEncodingFromValue(v string) (*ObservabilityPipelineHttpClientDestinationEncoding, error) {
	ev := ObservabilityPipelineHttpClientDestinationEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineHttpClientDestinationEncoding: valid values are %v", v, allowedObservabilityPipelineHttpClientDestinationEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineHttpClientDestinationEncoding) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineHttpClientDestinationEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineHttpClientDestinationEncoding value.
func (v ObservabilityPipelineHttpClientDestinationEncoding) Ptr() *ObservabilityPipelineHttpClientDestinationEncoding {
	return &v
}
