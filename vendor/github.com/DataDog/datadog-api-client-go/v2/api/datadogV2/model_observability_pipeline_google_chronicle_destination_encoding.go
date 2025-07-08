// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGoogleChronicleDestinationEncoding The encoding format for the logs sent to Chronicle.
type ObservabilityPipelineGoogleChronicleDestinationEncoding string

// List of ObservabilityPipelineGoogleChronicleDestinationEncoding.
const (
	OBSERVABILITYPIPELINEGOOGLECHRONICLEDESTINATIONENCODING_JSON        ObservabilityPipelineGoogleChronicleDestinationEncoding = "json"
	OBSERVABILITYPIPELINEGOOGLECHRONICLEDESTINATIONENCODING_RAW_MESSAGE ObservabilityPipelineGoogleChronicleDestinationEncoding = "raw_message"
)

var allowedObservabilityPipelineGoogleChronicleDestinationEncodingEnumValues = []ObservabilityPipelineGoogleChronicleDestinationEncoding{
	OBSERVABILITYPIPELINEGOOGLECHRONICLEDESTINATIONENCODING_JSON,
	OBSERVABILITYPIPELINEGOOGLECHRONICLEDESTINATIONENCODING_RAW_MESSAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGoogleChronicleDestinationEncoding) GetAllowedValues() []ObservabilityPipelineGoogleChronicleDestinationEncoding {
	return allowedObservabilityPipelineGoogleChronicleDestinationEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGoogleChronicleDestinationEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGoogleChronicleDestinationEncoding(value)
	return nil
}

// NewObservabilityPipelineGoogleChronicleDestinationEncodingFromValue returns a pointer to a valid ObservabilityPipelineGoogleChronicleDestinationEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGoogleChronicleDestinationEncodingFromValue(v string) (*ObservabilityPipelineGoogleChronicleDestinationEncoding, error) {
	ev := ObservabilityPipelineGoogleChronicleDestinationEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGoogleChronicleDestinationEncoding: valid values are %v", v, allowedObservabilityPipelineGoogleChronicleDestinationEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGoogleChronicleDestinationEncoding) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGoogleChronicleDestinationEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGoogleChronicleDestinationEncoding value.
func (v ObservabilityPipelineGoogleChronicleDestinationEncoding) Ptr() *ObservabilityPipelineGoogleChronicleDestinationEncoding {
	return &v
}
