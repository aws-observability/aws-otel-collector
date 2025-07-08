// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedSchemaType Schema type.
type EntityResponseIncludedSchemaType string

// List of EntityResponseIncludedSchemaType.
const (
	ENTITYRESPONSEINCLUDEDSCHEMATYPE_SCHEMA EntityResponseIncludedSchemaType = "schema"
)

var allowedEntityResponseIncludedSchemaTypeEnumValues = []EntityResponseIncludedSchemaType{
	ENTITYRESPONSEINCLUDEDSCHEMATYPE_SCHEMA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityResponseIncludedSchemaType) GetAllowedValues() []EntityResponseIncludedSchemaType {
	return allowedEntityResponseIncludedSchemaTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityResponseIncludedSchemaType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityResponseIncludedSchemaType(value)
	return nil
}

// NewEntityResponseIncludedSchemaTypeFromValue returns a pointer to a valid EntityResponseIncludedSchemaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityResponseIncludedSchemaTypeFromValue(v string) (*EntityResponseIncludedSchemaType, error) {
	ev := EntityResponseIncludedSchemaType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityResponseIncludedSchemaType: valid values are %v", v, allowedEntityResponseIncludedSchemaTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityResponseIncludedSchemaType) IsValid() bool {
	for _, existing := range allowedEntityResponseIncludedSchemaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityResponseIncludedSchemaType value.
func (v EntityResponseIncludedSchemaType) Ptr() *EntityResponseIncludedSchemaType {
	return &v
}
