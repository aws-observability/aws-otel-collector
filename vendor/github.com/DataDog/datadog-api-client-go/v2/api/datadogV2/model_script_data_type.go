// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScriptDataType The definition of `ScriptDataType` object.
type ScriptDataType string

// List of ScriptDataType.
const (
	SCRIPTDATATYPE_SCRIPTS ScriptDataType = "scripts"
)

var allowedScriptDataTypeEnumValues = []ScriptDataType{
	SCRIPTDATATYPE_SCRIPTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScriptDataType) GetAllowedValues() []ScriptDataType {
	return allowedScriptDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScriptDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScriptDataType(value)
	return nil
}

// NewScriptDataTypeFromValue returns a pointer to a valid ScriptDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScriptDataTypeFromValue(v string) (*ScriptDataType, error) {
	ev := ScriptDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScriptDataType: valid values are %v", v, allowedScriptDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScriptDataType) IsValid() bool {
	for _, existing := range allowedScriptDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScriptDataType value.
func (v ScriptDataType) Ptr() *ScriptDataType {
	return &v
}
