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

// CIAppPipelineEventStepStatus The final status of the step.
type CIAppPipelineEventStepStatus string

// List of CIAppPipelineEventStepStatus.
const (
	CIAPPPIPELINEEVENTSTEPSTATUS_SUCCESS CIAppPipelineEventStepStatus = "success"
	CIAPPPIPELINEEVENTSTEPSTATUS_ERROR   CIAppPipelineEventStepStatus = "error"
)

var allowedCIAppPipelineEventStepStatusEnumValues = []CIAppPipelineEventStepStatus{
	CIAPPPIPELINEEVENTSTEPSTATUS_SUCCESS,
	CIAPPPIPELINEEVENTSTEPSTATUS_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventStepStatus) GetAllowedValues() []CIAppPipelineEventStepStatus {
	return allowedCIAppPipelineEventStepStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventStepStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventStepStatus(value)
	return nil
}

// NewCIAppPipelineEventStepStatusFromValue returns a pointer to a valid CIAppPipelineEventStepStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventStepStatusFromValue(v string) (*CIAppPipelineEventStepStatus, error) {
	ev := CIAppPipelineEventStepStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventStepStatus: valid values are %v", v, allowedCIAppPipelineEventStepStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventStepStatus) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventStepStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventStepStatus value.
func (v CIAppPipelineEventStepStatus) Ptr() *CIAppPipelineEventStepStatus {
	return &v
}
<<<<<<< HEAD

// NullableCIAppPipelineEventStepStatus handles when a null is used for CIAppPipelineEventStepStatus.
type NullableCIAppPipelineEventStepStatus struct {
	value *CIAppPipelineEventStepStatus
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppPipelineEventStepStatus) Get() *CIAppPipelineEventStepStatus {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppPipelineEventStepStatus) Set(val *CIAppPipelineEventStepStatus) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppPipelineEventStepStatus) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableCIAppPipelineEventStepStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppPipelineEventStepStatus initializes the struct as if Set has been called.
func NewNullableCIAppPipelineEventStepStatus(val *CIAppPipelineEventStepStatus) *NullableCIAppPipelineEventStepStatus {
	return &NullableCIAppPipelineEventStepStatus{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppPipelineEventStepStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppPipelineEventStepStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
