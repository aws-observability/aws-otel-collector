// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// CIAppPipelineEventStepLevel Used to distinguish between pipelines, stages, jobs and steps.
type CIAppPipelineEventStepLevel string

// List of CIAppPipelineEventStepLevel.
const (
	CIAPPPIPELINEEVENTSTEPLEVEL_STEP CIAppPipelineEventStepLevel = "step"
)

var allowedCIAppPipelineEventStepLevelEnumValues = []CIAppPipelineEventStepLevel{
	CIAPPPIPELINEEVENTSTEPLEVEL_STEP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventStepLevel) GetAllowedValues() []CIAppPipelineEventStepLevel {
	return allowedCIAppPipelineEventStepLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventStepLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventStepLevel(value)
	return nil
}

// NewCIAppPipelineEventStepLevelFromValue returns a pointer to a valid CIAppPipelineEventStepLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventStepLevelFromValue(v string) (*CIAppPipelineEventStepLevel, error) {
	ev := CIAppPipelineEventStepLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventStepLevel: valid values are %v", v, allowedCIAppPipelineEventStepLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventStepLevel) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventStepLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventStepLevel value.
func (v CIAppPipelineEventStepLevel) Ptr() *CIAppPipelineEventStepLevel {
	return &v
}

// NullableCIAppPipelineEventStepLevel handles when a null is used for CIAppPipelineEventStepLevel.
type NullableCIAppPipelineEventStepLevel struct {
	value *CIAppPipelineEventStepLevel
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppPipelineEventStepLevel) Get() *CIAppPipelineEventStepLevel {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppPipelineEventStepLevel) Set(val *CIAppPipelineEventStepLevel) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppPipelineEventStepLevel) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableCIAppPipelineEventStepLevel) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppPipelineEventStepLevel initializes the struct as if Set has been called.
func NewNullableCIAppPipelineEventStepLevel(val *CIAppPipelineEventStepLevel) *NullableCIAppPipelineEventStepLevel {
	return &NullableCIAppPipelineEventStepLevel{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppPipelineEventStepLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppPipelineEventStepLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
