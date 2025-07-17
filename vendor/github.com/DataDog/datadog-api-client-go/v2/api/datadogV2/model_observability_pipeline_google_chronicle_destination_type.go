// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGoogleChronicleDestinationType The destination type. The value should always be `google_chronicle`.
type ObservabilityPipelineGoogleChronicleDestinationType string

// List of ObservabilityPipelineGoogleChronicleDestinationType.
const (
	OBSERVABILITYPIPELINEGOOGLECHRONICLEDESTINATIONTYPE_GOOGLE_CHRONICLE ObservabilityPipelineGoogleChronicleDestinationType = "google_chronicle"
)

var allowedObservabilityPipelineGoogleChronicleDestinationTypeEnumValues = []ObservabilityPipelineGoogleChronicleDestinationType{
	OBSERVABILITYPIPELINEGOOGLECHRONICLEDESTINATIONTYPE_GOOGLE_CHRONICLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGoogleChronicleDestinationType) GetAllowedValues() []ObservabilityPipelineGoogleChronicleDestinationType {
	return allowedObservabilityPipelineGoogleChronicleDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGoogleChronicleDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGoogleChronicleDestinationType(value)
	return nil
}

// NewObservabilityPipelineGoogleChronicleDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineGoogleChronicleDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGoogleChronicleDestinationTypeFromValue(v string) (*ObservabilityPipelineGoogleChronicleDestinationType, error) {
	ev := ObservabilityPipelineGoogleChronicleDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGoogleChronicleDestinationType: valid values are %v", v, allowedObservabilityPipelineGoogleChronicleDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGoogleChronicleDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGoogleChronicleDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGoogleChronicleDestinationType value.
func (v ObservabilityPipelineGoogleChronicleDestinationType) Ptr() *ObservabilityPipelineGoogleChronicleDestinationType {
	return &v
}
