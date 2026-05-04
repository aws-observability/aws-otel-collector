// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalVariableType Global variable type.
type GlobalVariableType string

// List of GlobalVariableType.
const (
	GLOBALVARIABLETYPE_GLOBAL_VARIABLES GlobalVariableType = "global_variables"
)

var allowedGlobalVariableTypeEnumValues = []GlobalVariableType{
	GLOBALVARIABLETYPE_GLOBAL_VARIABLES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GlobalVariableType) GetAllowedValues() []GlobalVariableType {
	return allowedGlobalVariableTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GlobalVariableType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GlobalVariableType(value)
	return nil
}

// NewGlobalVariableTypeFromValue returns a pointer to a valid GlobalVariableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGlobalVariableTypeFromValue(v string) (*GlobalVariableType, error) {
	ev := GlobalVariableType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GlobalVariableType: valid values are %v", v, allowedGlobalVariableTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GlobalVariableType) IsValid() bool {
	for _, existing := range allowedGlobalVariableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GlobalVariableType value.
func (v GlobalVariableType) Ptr() *GlobalVariableType {
	return &v
}
