// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
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
	err := datadog.Unmarshal(src, &value)
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
