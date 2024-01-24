// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistWidgetFlatType Top list widget flat display type.
type ToplistWidgetFlatType string

// List of ToplistWidgetFlatType.
const (
	TOPLISTWIDGETFLATTYPE_FLAT ToplistWidgetFlatType = "flat"
)

var allowedToplistWidgetFlatTypeEnumValues = []ToplistWidgetFlatType{
	TOPLISTWIDGETFLATTYPE_FLAT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistWidgetFlatType) GetAllowedValues() []ToplistWidgetFlatType {
	return allowedToplistWidgetFlatTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistWidgetFlatType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistWidgetFlatType(value)
	return nil
}

// NewToplistWidgetFlatTypeFromValue returns a pointer to a valid ToplistWidgetFlatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistWidgetFlatTypeFromValue(v string) (*ToplistWidgetFlatType, error) {
	ev := ToplistWidgetFlatType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistWidgetFlatType: valid values are %v", v, allowedToplistWidgetFlatTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistWidgetFlatType) IsValid() bool {
	for _, existing := range allowedToplistWidgetFlatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistWidgetFlatType value.
func (v ToplistWidgetFlatType) Ptr() *ToplistWidgetFlatType {
	return &v
}
