// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGooglePubSubDestinationEncoding Encoding format for log events.
type ObservabilityPipelineGooglePubSubDestinationEncoding string

// List of ObservabilityPipelineGooglePubSubDestinationEncoding.
const (
	OBSERVABILITYPIPELINEGOOGLEPUBSUBDESTINATIONENCODING_JSON        ObservabilityPipelineGooglePubSubDestinationEncoding = "json"
	OBSERVABILITYPIPELINEGOOGLEPUBSUBDESTINATIONENCODING_RAW_MESSAGE ObservabilityPipelineGooglePubSubDestinationEncoding = "raw_message"
)

var allowedObservabilityPipelineGooglePubSubDestinationEncodingEnumValues = []ObservabilityPipelineGooglePubSubDestinationEncoding{
	OBSERVABILITYPIPELINEGOOGLEPUBSUBDESTINATIONENCODING_JSON,
	OBSERVABILITYPIPELINEGOOGLEPUBSUBDESTINATIONENCODING_RAW_MESSAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGooglePubSubDestinationEncoding) GetAllowedValues() []ObservabilityPipelineGooglePubSubDestinationEncoding {
	return allowedObservabilityPipelineGooglePubSubDestinationEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGooglePubSubDestinationEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGooglePubSubDestinationEncoding(value)
	return nil
}

// NewObservabilityPipelineGooglePubSubDestinationEncodingFromValue returns a pointer to a valid ObservabilityPipelineGooglePubSubDestinationEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGooglePubSubDestinationEncodingFromValue(v string) (*ObservabilityPipelineGooglePubSubDestinationEncoding, error) {
	ev := ObservabilityPipelineGooglePubSubDestinationEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGooglePubSubDestinationEncoding: valid values are %v", v, allowedObservabilityPipelineGooglePubSubDestinationEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGooglePubSubDestinationEncoding) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGooglePubSubDestinationEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGooglePubSubDestinationEncoding value.
func (v ObservabilityPipelineGooglePubSubDestinationEncoding) Ptr() *ObservabilityPipelineGooglePubSubDestinationEncoding {
	return &v
}
