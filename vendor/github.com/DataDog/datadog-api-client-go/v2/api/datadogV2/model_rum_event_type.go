// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMEventType Type of the event.
type RUMEventType string

// List of RUMEventType.
const (
	RUMEVENTTYPE_RUM RUMEventType = "rum"
)

var allowedRUMEventTypeEnumValues = []RUMEventType{
	RUMEVENTTYPE_RUM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMEventType) GetAllowedValues() []RUMEventType {
	return allowedRUMEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMEventType(value)
	return nil
}

// NewRUMEventTypeFromValue returns a pointer to a valid RUMEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMEventTypeFromValue(v string) (*RUMEventType, error) {
	ev := RUMEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMEventType: valid values are %v", v, allowedRUMEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMEventType) IsValid() bool {
	for _, existing := range allowedRUMEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMEventType value.
func (v RUMEventType) Ptr() *RUMEventType {
	return &v
}
