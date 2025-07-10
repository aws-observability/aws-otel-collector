// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberFormatUnitScaleType The type of unit scale.
type NumberFormatUnitScaleType string

// List of NumberFormatUnitScaleType.
const (
	NUMBERFORMATUNITSCALETYPE_CANONICAL_UNIT NumberFormatUnitScaleType = "canonical_unit"
)

var allowedNumberFormatUnitScaleTypeEnumValues = []NumberFormatUnitScaleType{
	NUMBERFORMATUNITSCALETYPE_CANONICAL_UNIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NumberFormatUnitScaleType) GetAllowedValues() []NumberFormatUnitScaleType {
	return allowedNumberFormatUnitScaleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NumberFormatUnitScaleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NumberFormatUnitScaleType(value)
	return nil
}

// NewNumberFormatUnitScaleTypeFromValue returns a pointer to a valid NumberFormatUnitScaleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNumberFormatUnitScaleTypeFromValue(v string) (*NumberFormatUnitScaleType, error) {
	ev := NumberFormatUnitScaleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NumberFormatUnitScaleType: valid values are %v", v, allowedNumberFormatUnitScaleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NumberFormatUnitScaleType) IsValid() bool {
	for _, existing := range allowedNumberFormatUnitScaleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NumberFormatUnitScaleType value.
func (v NumberFormatUnitScaleType) Ptr() *NumberFormatUnitScaleType {
	return &v
}
