// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSumoLogicDestinationType The destination type. The value should always be `sumo_logic`.
type ObservabilityPipelineSumoLogicDestinationType string

// List of ObservabilityPipelineSumoLogicDestinationType.
const (
	OBSERVABILITYPIPELINESUMOLOGICDESTINATIONTYPE_SUMO_LOGIC ObservabilityPipelineSumoLogicDestinationType = "sumo_logic"
)

var allowedObservabilityPipelineSumoLogicDestinationTypeEnumValues = []ObservabilityPipelineSumoLogicDestinationType{
	OBSERVABILITYPIPELINESUMOLOGICDESTINATIONTYPE_SUMO_LOGIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSumoLogicDestinationType) GetAllowedValues() []ObservabilityPipelineSumoLogicDestinationType {
	return allowedObservabilityPipelineSumoLogicDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSumoLogicDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSumoLogicDestinationType(value)
	return nil
}

// NewObservabilityPipelineSumoLogicDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineSumoLogicDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSumoLogicDestinationTypeFromValue(v string) (*ObservabilityPipelineSumoLogicDestinationType, error) {
	ev := ObservabilityPipelineSumoLogicDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSumoLogicDestinationType: valid values are %v", v, allowedObservabilityPipelineSumoLogicDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSumoLogicDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSumoLogicDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSumoLogicDestinationType value.
func (v ObservabilityPipelineSumoLogicDestinationType) Ptr() *ObservabilityPipelineSumoLogicDestinationType {
	return &v
}
