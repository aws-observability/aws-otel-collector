// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BlueprintDataType The resource type for a blueprint.
type BlueprintDataType string

// List of BlueprintDataType.
const (
	BLUEPRINTDATATYPE_BLUEPRINT BlueprintDataType = "blueprint"
)

var allowedBlueprintDataTypeEnumValues = []BlueprintDataType{
	BLUEPRINTDATATYPE_BLUEPRINT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BlueprintDataType) GetAllowedValues() []BlueprintDataType {
	return allowedBlueprintDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BlueprintDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BlueprintDataType(value)
	return nil
}

// NewBlueprintDataTypeFromValue returns a pointer to a valid BlueprintDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBlueprintDataTypeFromValue(v string) (*BlueprintDataType, error) {
	ev := BlueprintDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BlueprintDataType: valid values are %v", v, allowedBlueprintDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BlueprintDataType) IsValid() bool {
	for _, existing := range allowedBlueprintDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BlueprintDataType value.
func (v BlueprintDataType) Ptr() *BlueprintDataType {
	return &v
}
