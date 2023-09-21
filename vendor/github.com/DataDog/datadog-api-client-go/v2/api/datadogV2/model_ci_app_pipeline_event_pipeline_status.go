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

// CIAppPipelineEventPipelineStatus The final status of the pipeline.
type CIAppPipelineEventPipelineStatus string

// List of CIAppPipelineEventPipelineStatus.
const (
	CIAPPPIPELINEEVENTPIPELINESTATUS_SUCCESS  CIAppPipelineEventPipelineStatus = "success"
	CIAPPPIPELINEEVENTPIPELINESTATUS_ERROR    CIAppPipelineEventPipelineStatus = "error"
	CIAPPPIPELINEEVENTPIPELINESTATUS_CANCELED CIAppPipelineEventPipelineStatus = "canceled"
	CIAPPPIPELINEEVENTPIPELINESTATUS_SKIPPED  CIAppPipelineEventPipelineStatus = "skipped"
	CIAPPPIPELINEEVENTPIPELINESTATUS_BLOCKED  CIAppPipelineEventPipelineStatus = "blocked"
)

var allowedCIAppPipelineEventPipelineStatusEnumValues = []CIAppPipelineEventPipelineStatus{
	CIAPPPIPELINEEVENTPIPELINESTATUS_SUCCESS,
	CIAPPPIPELINEEVENTPIPELINESTATUS_ERROR,
	CIAPPPIPELINEEVENTPIPELINESTATUS_CANCELED,
	CIAPPPIPELINEEVENTPIPELINESTATUS_SKIPPED,
	CIAPPPIPELINEEVENTPIPELINESTATUS_BLOCKED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventPipelineStatus) GetAllowedValues() []CIAppPipelineEventPipelineStatus {
	return allowedCIAppPipelineEventPipelineStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventPipelineStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventPipelineStatus(value)
	return nil
}

// NewCIAppPipelineEventPipelineStatusFromValue returns a pointer to a valid CIAppPipelineEventPipelineStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventPipelineStatusFromValue(v string) (*CIAppPipelineEventPipelineStatus, error) {
	ev := CIAppPipelineEventPipelineStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventPipelineStatus: valid values are %v", v, allowedCIAppPipelineEventPipelineStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventPipelineStatus) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventPipelineStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventPipelineStatus value.
func (v CIAppPipelineEventPipelineStatus) Ptr() *CIAppPipelineEventPipelineStatus {
	return &v
}
<<<<<<< HEAD

// NullableCIAppPipelineEventPipelineStatus handles when a null is used for CIAppPipelineEventPipelineStatus.
type NullableCIAppPipelineEventPipelineStatus struct {
	value *CIAppPipelineEventPipelineStatus
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppPipelineEventPipelineStatus) Get() *CIAppPipelineEventPipelineStatus {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppPipelineEventPipelineStatus) Set(val *CIAppPipelineEventPipelineStatus) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppPipelineEventPipelineStatus) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableCIAppPipelineEventPipelineStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppPipelineEventPipelineStatus initializes the struct as if Set has been called.
func NewNullableCIAppPipelineEventPipelineStatus(val *CIAppPipelineEventPipelineStatus) *NullableCIAppPipelineEventPipelineStatus {
	return &NullableCIAppPipelineEventPipelineStatus{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppPipelineEventPipelineStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppPipelineEventPipelineStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
