// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberFormatUnitCustomType The type of custom unit.
type NumberFormatUnitCustomType string

// List of NumberFormatUnitCustomType.
const (
	NUMBERFORMATUNITCUSTOMTYPE_CUSTOM_UNIT_LABEL NumberFormatUnitCustomType = "custom_unit_label"
)

var allowedNumberFormatUnitCustomTypeEnumValues = []NumberFormatUnitCustomType{
	NUMBERFORMATUNITCUSTOMTYPE_CUSTOM_UNIT_LABEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NumberFormatUnitCustomType) GetAllowedValues() []NumberFormatUnitCustomType {
	return allowedNumberFormatUnitCustomTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NumberFormatUnitCustomType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NumberFormatUnitCustomType(value)
	return nil
}

// NewNumberFormatUnitCustomTypeFromValue returns a pointer to a valid NumberFormatUnitCustomType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNumberFormatUnitCustomTypeFromValue(v string) (*NumberFormatUnitCustomType, error) {
	ev := NumberFormatUnitCustomType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NumberFormatUnitCustomType: valid values are %v", v, allowedNumberFormatUnitCustomTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NumberFormatUnitCustomType) IsValid() bool {
	for _, existing := range allowedNumberFormatUnitCustomTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NumberFormatUnitCustomType value.
func (v NumberFormatUnitCustomType) Ptr() *NumberFormatUnitCustomType {
	return &v
}
