// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalVariableJsonPatchType Global variable JSON Patch type.
type GlobalVariableJsonPatchType string

// List of GlobalVariableJsonPatchType.
const (
	GLOBALVARIABLEJSONPATCHTYPE_GLOBAL_VARIABLES_JSON_PATCH GlobalVariableJsonPatchType = "global_variables_json_patch"
)

var allowedGlobalVariableJsonPatchTypeEnumValues = []GlobalVariableJsonPatchType{
	GLOBALVARIABLEJSONPATCHTYPE_GLOBAL_VARIABLES_JSON_PATCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GlobalVariableJsonPatchType) GetAllowedValues() []GlobalVariableJsonPatchType {
	return allowedGlobalVariableJsonPatchTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GlobalVariableJsonPatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GlobalVariableJsonPatchType(value)
	return nil
}

// NewGlobalVariableJsonPatchTypeFromValue returns a pointer to a valid GlobalVariableJsonPatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGlobalVariableJsonPatchTypeFromValue(v string) (*GlobalVariableJsonPatchType, error) {
	ev := GlobalVariableJsonPatchType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GlobalVariableJsonPatchType: valid values are %v", v, allowedGlobalVariableJsonPatchTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GlobalVariableJsonPatchType) IsValid() bool {
	for _, existing := range allowedGlobalVariableJsonPatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GlobalVariableJsonPatchType value.
func (v GlobalVariableJsonPatchType) Ptr() *GlobalVariableJsonPatchType {
	return &v
}
