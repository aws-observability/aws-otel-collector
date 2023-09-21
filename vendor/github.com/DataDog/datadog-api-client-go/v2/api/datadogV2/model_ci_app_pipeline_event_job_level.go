// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
	"fmt"

	"github.com/goccy/go-json"
>>>>>>> main
)

// CIAppPipelineEventJobLevel Used to distinguish between pipelines, stages, jobs, and steps.
type CIAppPipelineEventJobLevel string

// List of CIAppPipelineEventJobLevel.
const (
	CIAPPPIPELINEEVENTJOBLEVEL_JOB CIAppPipelineEventJobLevel = "job"
)

var allowedCIAppPipelineEventJobLevelEnumValues = []CIAppPipelineEventJobLevel{
	CIAPPPIPELINEEVENTJOBLEVEL_JOB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventJobLevel) GetAllowedValues() []CIAppPipelineEventJobLevel {
	return allowedCIAppPipelineEventJobLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventJobLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventJobLevel(value)
	return nil
}

// NewCIAppPipelineEventJobLevelFromValue returns a pointer to a valid CIAppPipelineEventJobLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventJobLevelFromValue(v string) (*CIAppPipelineEventJobLevel, error) {
	ev := CIAppPipelineEventJobLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventJobLevel: valid values are %v", v, allowedCIAppPipelineEventJobLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventJobLevel) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventJobLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventJobLevel value.
func (v CIAppPipelineEventJobLevel) Ptr() *CIAppPipelineEventJobLevel {
	return &v
}
<<<<<<< HEAD

// NullableCIAppPipelineEventJobLevel handles when a null is used for CIAppPipelineEventJobLevel.
type NullableCIAppPipelineEventJobLevel struct {
	value *CIAppPipelineEventJobLevel
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppPipelineEventJobLevel) Get() *CIAppPipelineEventJobLevel {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppPipelineEventJobLevel) Set(val *CIAppPipelineEventJobLevel) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppPipelineEventJobLevel) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableCIAppPipelineEventJobLevel) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppPipelineEventJobLevel initializes the struct as if Set has been called.
func NewNullableCIAppPipelineEventJobLevel(val *CIAppPipelineEventJobLevel) *NullableCIAppPipelineEventJobLevel {
	return &NullableCIAppPipelineEventJobLevel{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppPipelineEventJobLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppPipelineEventJobLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
