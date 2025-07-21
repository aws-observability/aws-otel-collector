// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineFluentdSourceType The source type. The value should always be `fluentd.
type ObservabilityPipelineFluentdSourceType string

// List of ObservabilityPipelineFluentdSourceType.
const (
	OBSERVABILITYPIPELINEFLUENTDSOURCETYPE_FLUENTD ObservabilityPipelineFluentdSourceType = "fluentd"
)

var allowedObservabilityPipelineFluentdSourceTypeEnumValues = []ObservabilityPipelineFluentdSourceType{
	OBSERVABILITYPIPELINEFLUENTDSOURCETYPE_FLUENTD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineFluentdSourceType) GetAllowedValues() []ObservabilityPipelineFluentdSourceType {
	return allowedObservabilityPipelineFluentdSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineFluentdSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineFluentdSourceType(value)
	return nil
}

// NewObservabilityPipelineFluentdSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineFluentdSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineFluentdSourceTypeFromValue(v string) (*ObservabilityPipelineFluentdSourceType, error) {
	ev := ObservabilityPipelineFluentdSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineFluentdSourceType: valid values are %v", v, allowedObservabilityPipelineFluentdSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineFluentdSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineFluentdSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineFluentdSourceType value.
func (v ObservabilityPipelineFluentdSourceType) Ptr() *ObservabilityPipelineFluentdSourceType {
	return &v
}
