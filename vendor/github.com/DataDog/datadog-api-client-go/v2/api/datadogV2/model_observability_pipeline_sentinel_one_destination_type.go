// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSentinelOneDestinationType The destination type. The value should always be `sentinel_one`.
type ObservabilityPipelineSentinelOneDestinationType string

// List of ObservabilityPipelineSentinelOneDestinationType.
const (
	OBSERVABILITYPIPELINESENTINELONEDESTINATIONTYPE_SENTINEL_ONE ObservabilityPipelineSentinelOneDestinationType = "sentinel_one"
)

var allowedObservabilityPipelineSentinelOneDestinationTypeEnumValues = []ObservabilityPipelineSentinelOneDestinationType{
	OBSERVABILITYPIPELINESENTINELONEDESTINATIONTYPE_SENTINEL_ONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSentinelOneDestinationType) GetAllowedValues() []ObservabilityPipelineSentinelOneDestinationType {
	return allowedObservabilityPipelineSentinelOneDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSentinelOneDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSentinelOneDestinationType(value)
	return nil
}

// NewObservabilityPipelineSentinelOneDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineSentinelOneDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSentinelOneDestinationTypeFromValue(v string) (*ObservabilityPipelineSentinelOneDestinationType, error) {
	ev := ObservabilityPipelineSentinelOneDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSentinelOneDestinationType: valid values are %v", v, allowedObservabilityPipelineSentinelOneDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSentinelOneDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSentinelOneDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSentinelOneDestinationType value.
func (v ObservabilityPipelineSentinelOneDestinationType) Ptr() *ObservabilityPipelineSentinelOneDestinationType {
	return &v
}
