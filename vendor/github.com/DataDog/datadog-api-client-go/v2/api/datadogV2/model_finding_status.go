// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// FindingStatus The status of the finding.
type FindingStatus string

// List of FindingStatus.
const (
	FINDINGSTATUS_CRITICAL FindingStatus = "critical"
	FINDINGSTATUS_HIGH     FindingStatus = "high"
	FINDINGSTATUS_MEDIUM   FindingStatus = "medium"
	FINDINGSTATUS_LOW      FindingStatus = "low"
	FINDINGSTATUS_INFO     FindingStatus = "info"
)

var allowedFindingStatusEnumValues = []FindingStatus{
	FINDINGSTATUS_CRITICAL,
	FINDINGSTATUS_HIGH,
	FINDINGSTATUS_MEDIUM,
	FINDINGSTATUS_LOW,
	FINDINGSTATUS_INFO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FindingStatus) GetAllowedValues() []FindingStatus {
	return allowedFindingStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FindingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FindingStatus(value)
	return nil
}

// NewFindingStatusFromValue returns a pointer to a valid FindingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFindingStatusFromValue(v string) (*FindingStatus, error) {
	ev := FindingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FindingStatus: valid values are %v", v, allowedFindingStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FindingStatus) IsValid() bool {
	for _, existing := range allowedFindingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FindingStatus value.
func (v FindingStatus) Ptr() *FindingStatus {
	return &v
}

// NullableFindingStatus handles when a null is used for FindingStatus.
type NullableFindingStatus struct {
	value *FindingStatus
	isSet bool
}

// Get returns the associated value.
func (v NullableFindingStatus) Get() *FindingStatus {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableFindingStatus) Set(val *FindingStatus) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableFindingStatus) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableFindingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFindingStatus initializes the struct as if Set has been called.
func NewNullableFindingStatus(val *FindingStatus) *NullableFindingStatus {
	return &NullableFindingStatus{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableFindingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableFindingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
