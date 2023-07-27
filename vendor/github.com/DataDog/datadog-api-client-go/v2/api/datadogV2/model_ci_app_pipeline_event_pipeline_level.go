// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// CIAppPipelineEventPipelineLevel Used to distinguish between pipelines, stages, jobs, and steps.
type CIAppPipelineEventPipelineLevel string

// List of CIAppPipelineEventPipelineLevel.
const (
	CIAPPPIPELINEEVENTPIPELINELEVEL_PIPELINE CIAppPipelineEventPipelineLevel = "pipeline"
)

var allowedCIAppPipelineEventPipelineLevelEnumValues = []CIAppPipelineEventPipelineLevel{
	CIAPPPIPELINEEVENTPIPELINELEVEL_PIPELINE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventPipelineLevel) GetAllowedValues() []CIAppPipelineEventPipelineLevel {
	return allowedCIAppPipelineEventPipelineLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventPipelineLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventPipelineLevel(value)
	return nil
}

// NewCIAppPipelineEventPipelineLevelFromValue returns a pointer to a valid CIAppPipelineEventPipelineLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventPipelineLevelFromValue(v string) (*CIAppPipelineEventPipelineLevel, error) {
	ev := CIAppPipelineEventPipelineLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventPipelineLevel: valid values are %v", v, allowedCIAppPipelineEventPipelineLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventPipelineLevel) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventPipelineLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventPipelineLevel value.
func (v CIAppPipelineEventPipelineLevel) Ptr() *CIAppPipelineEventPipelineLevel {
	return &v
}

// NullableCIAppPipelineEventPipelineLevel handles when a null is used for CIAppPipelineEventPipelineLevel.
type NullableCIAppPipelineEventPipelineLevel struct {
	value *CIAppPipelineEventPipelineLevel
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppPipelineEventPipelineLevel) Get() *CIAppPipelineEventPipelineLevel {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppPipelineEventPipelineLevel) Set(val *CIAppPipelineEventPipelineLevel) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppPipelineEventPipelineLevel) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableCIAppPipelineEventPipelineLevel) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppPipelineEventPipelineLevel initializes the struct as if Set has been called.
func NewNullableCIAppPipelineEventPipelineLevel(val *CIAppPipelineEventPipelineLevel) *NullableCIAppPipelineEventPipelineLevel {
	return &NullableCIAppPipelineEventPipelineLevel{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppPipelineEventPipelineLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppPipelineEventPipelineLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
