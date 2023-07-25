// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"fmt"
)

// NotifyEndType A notification end type.
type NotifyEndType string

// List of NotifyEndType.
const (
	NOTIFYENDTYPE_CANCELED NotifyEndType = "canceled"
	NOTIFYENDTYPE_EXPIRED  NotifyEndType = "expired"
)

var allowedNotifyEndTypeEnumValues = []NotifyEndType{
	NOTIFYENDTYPE_CANCELED,
	NOTIFYENDTYPE_EXPIRED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotifyEndType) GetAllowedValues() []NotifyEndType {
	return allowedNotifyEndTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotifyEndType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotifyEndType(value)
	return nil
}

// NewNotifyEndTypeFromValue returns a pointer to a valid NotifyEndType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotifyEndTypeFromValue(v string) (*NotifyEndType, error) {
	ev := NotifyEndType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotifyEndType: valid values are %v", v, allowedNotifyEndTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotifyEndType) IsValid() bool {
	for _, existing := range allowedNotifyEndTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotifyEndType value.
func (v NotifyEndType) Ptr() *NotifyEndType {
	return &v
}

// NullableNotifyEndType handles when a null is used for NotifyEndType.
type NullableNotifyEndType struct {
	value *NotifyEndType
	isSet bool
}

// Get returns the associated value.
func (v NullableNotifyEndType) Get() *NotifyEndType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableNotifyEndType) Set(val *NotifyEndType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableNotifyEndType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableNotifyEndType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNotifyEndType initializes the struct as if Set has been called.
func NewNullableNotifyEndType(val *NotifyEndType) *NullableNotifyEndType {
	return &NullableNotifyEndType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableNotifyEndType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableNotifyEndType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
