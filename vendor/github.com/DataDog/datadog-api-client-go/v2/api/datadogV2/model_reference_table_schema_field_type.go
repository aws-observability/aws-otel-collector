// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReferenceTableSchemaFieldType The field type for reference table schema fields.
type ReferenceTableSchemaFieldType string

// List of ReferenceTableSchemaFieldType.
const (
	REFERENCETABLESCHEMAFIELDTYPE_STRING ReferenceTableSchemaFieldType = "STRING"
	REFERENCETABLESCHEMAFIELDTYPE_INT32  ReferenceTableSchemaFieldType = "INT32"
)

var allowedReferenceTableSchemaFieldTypeEnumValues = []ReferenceTableSchemaFieldType{
	REFERENCETABLESCHEMAFIELDTYPE_STRING,
	REFERENCETABLESCHEMAFIELDTYPE_INT32,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReferenceTableSchemaFieldType) GetAllowedValues() []ReferenceTableSchemaFieldType {
	return allowedReferenceTableSchemaFieldTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReferenceTableSchemaFieldType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReferenceTableSchemaFieldType(value)
	return nil
}

// NewReferenceTableSchemaFieldTypeFromValue returns a pointer to a valid ReferenceTableSchemaFieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReferenceTableSchemaFieldTypeFromValue(v string) (*ReferenceTableSchemaFieldType, error) {
	ev := ReferenceTableSchemaFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReferenceTableSchemaFieldType: valid values are %v", v, allowedReferenceTableSchemaFieldTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReferenceTableSchemaFieldType) IsValid() bool {
	for _, existing := range allowedReferenceTableSchemaFieldTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReferenceTableSchemaFieldType value.
func (v ReferenceTableSchemaFieldType) Ptr() *ReferenceTableSchemaFieldType {
	return &v
}
