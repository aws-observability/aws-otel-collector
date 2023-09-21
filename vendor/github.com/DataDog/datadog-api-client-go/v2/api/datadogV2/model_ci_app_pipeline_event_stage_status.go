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

// CIAppPipelineEventStageStatus The final status of the stage.
type CIAppPipelineEventStageStatus string

// List of CIAppPipelineEventStageStatus.
const (
	CIAPPPIPELINEEVENTSTAGESTATUS_SUCCESS  CIAppPipelineEventStageStatus = "success"
	CIAPPPIPELINEEVENTSTAGESTATUS_ERROR    CIAppPipelineEventStageStatus = "error"
	CIAPPPIPELINEEVENTSTAGESTATUS_CANCELED CIAppPipelineEventStageStatus = "canceled"
	CIAPPPIPELINEEVENTSTAGESTATUS_SKIPPED  CIAppPipelineEventStageStatus = "skipped"
)

var allowedCIAppPipelineEventStageStatusEnumValues = []CIAppPipelineEventStageStatus{
	CIAPPPIPELINEEVENTSTAGESTATUS_SUCCESS,
	CIAPPPIPELINEEVENTSTAGESTATUS_ERROR,
	CIAPPPIPELINEEVENTSTAGESTATUS_CANCELED,
	CIAPPPIPELINEEVENTSTAGESTATUS_SKIPPED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventStageStatus) GetAllowedValues() []CIAppPipelineEventStageStatus {
	return allowedCIAppPipelineEventStageStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventStageStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventStageStatus(value)
	return nil
}

// NewCIAppPipelineEventStageStatusFromValue returns a pointer to a valid CIAppPipelineEventStageStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventStageStatusFromValue(v string) (*CIAppPipelineEventStageStatus, error) {
	ev := CIAppPipelineEventStageStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventStageStatus: valid values are %v", v, allowedCIAppPipelineEventStageStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventStageStatus) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventStageStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventStageStatus value.
func (v CIAppPipelineEventStageStatus) Ptr() *CIAppPipelineEventStageStatus {
	return &v
}
<<<<<<< HEAD

// NullableCIAppPipelineEventStageStatus handles when a null is used for CIAppPipelineEventStageStatus.
type NullableCIAppPipelineEventStageStatus struct {
	value *CIAppPipelineEventStageStatus
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppPipelineEventStageStatus) Get() *CIAppPipelineEventStageStatus {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppPipelineEventStageStatus) Set(val *CIAppPipelineEventStageStatus) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppPipelineEventStageStatus) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableCIAppPipelineEventStageStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppPipelineEventStageStatus initializes the struct as if Set has been called.
func NewNullableCIAppPipelineEventStageStatus(val *CIAppPipelineEventStageStatus) *NullableCIAppPipelineEventStageStatus {
	return &NullableCIAppPipelineEventStageStatus{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppPipelineEventStageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppPipelineEventStageStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
