// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MuteDataType Mute resource type.
type MuteDataType string

// List of MuteDataType.
const (
	MUTEDATATYPE_MUTE MuteDataType = "mute"
)

var allowedMuteDataTypeEnumValues = []MuteDataType{
	MUTEDATATYPE_MUTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MuteDataType) GetAllowedValues() []MuteDataType {
	return allowedMuteDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MuteDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MuteDataType(value)
	return nil
}

// NewMuteDataTypeFromValue returns a pointer to a valid MuteDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMuteDataTypeFromValue(v string) (*MuteDataType, error) {
	ev := MuteDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MuteDataType: valid values are %v", v, allowedMuteDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MuteDataType) IsValid() bool {
	for _, existing := range allowedMuteDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MuteDataType value.
func (v MuteDataType) Ptr() *MuteDataType {
	return &v
}
