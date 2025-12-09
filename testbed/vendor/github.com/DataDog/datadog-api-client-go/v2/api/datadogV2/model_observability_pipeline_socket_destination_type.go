// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketDestinationType The destination type. The value should always be `socket`.
type ObservabilityPipelineSocketDestinationType string

// List of ObservabilityPipelineSocketDestinationType.
const (
	OBSERVABILITYPIPELINESOCKETDESTINATIONTYPE_SOCKET ObservabilityPipelineSocketDestinationType = "socket"
)

var allowedObservabilityPipelineSocketDestinationTypeEnumValues = []ObservabilityPipelineSocketDestinationType{
	OBSERVABILITYPIPELINESOCKETDESTINATIONTYPE_SOCKET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSocketDestinationType) GetAllowedValues() []ObservabilityPipelineSocketDestinationType {
	return allowedObservabilityPipelineSocketDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSocketDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSocketDestinationType(value)
	return nil
}

// NewObservabilityPipelineSocketDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineSocketDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSocketDestinationTypeFromValue(v string) (*ObservabilityPipelineSocketDestinationType, error) {
	ev := ObservabilityPipelineSocketDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSocketDestinationType: valid values are %v", v, allowedObservabilityPipelineSocketDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSocketDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSocketDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSocketDestinationType value.
func (v ObservabilityPipelineSocketDestinationType) Ptr() *ObservabilityPipelineSocketDestinationType {
	return &v
}
