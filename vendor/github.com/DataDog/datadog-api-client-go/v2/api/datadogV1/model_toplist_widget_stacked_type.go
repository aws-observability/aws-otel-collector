// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistWidgetStackedType Top list widget stacked display type.
type ToplistWidgetStackedType string

// List of ToplistWidgetStackedType.
const (
	TOPLISTWIDGETSTACKEDTYPE_STACKED ToplistWidgetStackedType = "stacked"
)

var allowedToplistWidgetStackedTypeEnumValues = []ToplistWidgetStackedType{
	TOPLISTWIDGETSTACKEDTYPE_STACKED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistWidgetStackedType) GetAllowedValues() []ToplistWidgetStackedType {
	return allowedToplistWidgetStackedTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistWidgetStackedType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistWidgetStackedType(value)
	return nil
}

// NewToplistWidgetStackedTypeFromValue returns a pointer to a valid ToplistWidgetStackedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistWidgetStackedTypeFromValue(v string) (*ToplistWidgetStackedType, error) {
	ev := ToplistWidgetStackedType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistWidgetStackedType: valid values are %v", v, allowedToplistWidgetStackedTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistWidgetStackedType) IsValid() bool {
	for _, existing := range allowedToplistWidgetStackedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistWidgetStackedType value.
func (v ToplistWidgetStackedType) Ptr() *ToplistWidgetStackedType {
	return &v
}
