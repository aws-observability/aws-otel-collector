// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedRawSchemaType Raw schema type.
type EntityResponseIncludedRawSchemaType string

// List of EntityResponseIncludedRawSchemaType.
const (
	ENTITYRESPONSEINCLUDEDRAWSCHEMATYPE_RAW_SCHEMA EntityResponseIncludedRawSchemaType = "rawSchema"
)

var allowedEntityResponseIncludedRawSchemaTypeEnumValues = []EntityResponseIncludedRawSchemaType{
	ENTITYRESPONSEINCLUDEDRAWSCHEMATYPE_RAW_SCHEMA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityResponseIncludedRawSchemaType) GetAllowedValues() []EntityResponseIncludedRawSchemaType {
	return allowedEntityResponseIncludedRawSchemaTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityResponseIncludedRawSchemaType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityResponseIncludedRawSchemaType(value)
	return nil
}

// NewEntityResponseIncludedRawSchemaTypeFromValue returns a pointer to a valid EntityResponseIncludedRawSchemaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityResponseIncludedRawSchemaTypeFromValue(v string) (*EntityResponseIncludedRawSchemaType, error) {
	ev := EntityResponseIncludedRawSchemaType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityResponseIncludedRawSchemaType: valid values are %v", v, allowedEntityResponseIncludedRawSchemaTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityResponseIncludedRawSchemaType) IsValid() bool {
	for _, existing := range allowedEntityResponseIncludedRawSchemaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityResponseIncludedRawSchemaType value.
func (v EntityResponseIncludedRawSchemaType) Ptr() *EntityResponseIncludedRawSchemaType {
	return &v
}
