// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseDataType Entity resource type.
type EntityResponseDataType string

// List of EntityResponseDataType.
const (
	ENTITYRESPONSEDATATYPE_ENTITY EntityResponseDataType = "entity"
)

var allowedEntityResponseDataTypeEnumValues = []EntityResponseDataType{
	ENTITYRESPONSEDATATYPE_ENTITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityResponseDataType) GetAllowedValues() []EntityResponseDataType {
	return allowedEntityResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityResponseDataType(value)
	return nil
}

// NewEntityResponseDataTypeFromValue returns a pointer to a valid EntityResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityResponseDataTypeFromValue(v string) (*EntityResponseDataType, error) {
	ev := EntityResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityResponseDataType: valid values are %v", v, allowedEntityResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityResponseDataType) IsValid() bool {
	for _, existing := range allowedEntityResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityResponseDataType value.
func (v EntityResponseDataType) Ptr() *EntityResponseDataType {
	return &v
}
