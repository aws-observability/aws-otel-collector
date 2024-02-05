// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMApplicationListType RUM application list type.
type RUMApplicationListType string

// List of RUMApplicationListType.
const (
	RUMAPPLICATIONLISTTYPE_RUM_APPLICATION RUMApplicationListType = "rum_application"
)

var allowedRUMApplicationListTypeEnumValues = []RUMApplicationListType{
	RUMAPPLICATIONLISTTYPE_RUM_APPLICATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMApplicationListType) GetAllowedValues() []RUMApplicationListType {
	return allowedRUMApplicationListTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMApplicationListType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMApplicationListType(value)
	return nil
}

// NewRUMApplicationListTypeFromValue returns a pointer to a valid RUMApplicationListType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMApplicationListTypeFromValue(v string) (*RUMApplicationListType, error) {
	ev := RUMApplicationListType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMApplicationListType: valid values are %v", v, allowedRUMApplicationListTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMApplicationListType) IsValid() bool {
	for _, existing := range allowedRUMApplicationListTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMApplicationListType value.
func (v RUMApplicationListType) Ptr() *RUMApplicationListType {
	return &v
}
