// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineFluentBitSourceType The source type. The value should always be `fluent_bit`.
type ObservabilityPipelineFluentBitSourceType string

// List of ObservabilityPipelineFluentBitSourceType.
const (
	OBSERVABILITYPIPELINEFLUENTBITSOURCETYPE_FLUENT_BIT ObservabilityPipelineFluentBitSourceType = "fluent_bit"
)

var allowedObservabilityPipelineFluentBitSourceTypeEnumValues = []ObservabilityPipelineFluentBitSourceType{
	OBSERVABILITYPIPELINEFLUENTBITSOURCETYPE_FLUENT_BIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineFluentBitSourceType) GetAllowedValues() []ObservabilityPipelineFluentBitSourceType {
	return allowedObservabilityPipelineFluentBitSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineFluentBitSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineFluentBitSourceType(value)
	return nil
}

// NewObservabilityPipelineFluentBitSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineFluentBitSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineFluentBitSourceTypeFromValue(v string) (*ObservabilityPipelineFluentBitSourceType, error) {
	ev := ObservabilityPipelineFluentBitSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineFluentBitSourceType: valid values are %v", v, allowedObservabilityPipelineFluentBitSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineFluentBitSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineFluentBitSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineFluentBitSourceType value.
func (v ObservabilityPipelineFluentBitSourceType) Ptr() *ObservabilityPipelineFluentBitSourceType {
	return &v
}
