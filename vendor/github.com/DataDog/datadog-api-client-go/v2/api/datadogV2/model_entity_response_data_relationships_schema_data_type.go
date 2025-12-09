// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseDataRelationshipsSchemaDataType Schema resource type.
type EntityResponseDataRelationshipsSchemaDataType string

// List of EntityResponseDataRelationshipsSchemaDataType.
const (
	ENTITYRESPONSEDATARELATIONSHIPSSCHEMADATATYPE_SCHEMA EntityResponseDataRelationshipsSchemaDataType = "schema"
)

var allowedEntityResponseDataRelationshipsSchemaDataTypeEnumValues = []EntityResponseDataRelationshipsSchemaDataType{
	ENTITYRESPONSEDATARELATIONSHIPSSCHEMADATATYPE_SCHEMA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityResponseDataRelationshipsSchemaDataType) GetAllowedValues() []EntityResponseDataRelationshipsSchemaDataType {
	return allowedEntityResponseDataRelationshipsSchemaDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityResponseDataRelationshipsSchemaDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityResponseDataRelationshipsSchemaDataType(value)
	return nil
}

// NewEntityResponseDataRelationshipsSchemaDataTypeFromValue returns a pointer to a valid EntityResponseDataRelationshipsSchemaDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityResponseDataRelationshipsSchemaDataTypeFromValue(v string) (*EntityResponseDataRelationshipsSchemaDataType, error) {
	ev := EntityResponseDataRelationshipsSchemaDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityResponseDataRelationshipsSchemaDataType: valid values are %v", v, allowedEntityResponseDataRelationshipsSchemaDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityResponseDataRelationshipsSchemaDataType) IsValid() bool {
	for _, existing := range allowedEntityResponseDataRelationshipsSchemaDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityResponseDataRelationshipsSchemaDataType value.
func (v EntityResponseDataRelationshipsSchemaDataType) Ptr() *EntityResponseDataRelationshipsSchemaDataType {
	return &v
}
