// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseDataRelationshipsRawSchemaDataType Raw schema resource type.
type EntityResponseDataRelationshipsRawSchemaDataType string

// List of EntityResponseDataRelationshipsRawSchemaDataType.
const (
	ENTITYRESPONSEDATARELATIONSHIPSRAWSCHEMADATATYPE_RAWSCHEMA EntityResponseDataRelationshipsRawSchemaDataType = "rawSchema"
)

var allowedEntityResponseDataRelationshipsRawSchemaDataTypeEnumValues = []EntityResponseDataRelationshipsRawSchemaDataType{
	ENTITYRESPONSEDATARELATIONSHIPSRAWSCHEMADATATYPE_RAWSCHEMA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityResponseDataRelationshipsRawSchemaDataType) GetAllowedValues() []EntityResponseDataRelationshipsRawSchemaDataType {
	return allowedEntityResponseDataRelationshipsRawSchemaDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityResponseDataRelationshipsRawSchemaDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityResponseDataRelationshipsRawSchemaDataType(value)
	return nil
}

// NewEntityResponseDataRelationshipsRawSchemaDataTypeFromValue returns a pointer to a valid EntityResponseDataRelationshipsRawSchemaDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityResponseDataRelationshipsRawSchemaDataTypeFromValue(v string) (*EntityResponseDataRelationshipsRawSchemaDataType, error) {
	ev := EntityResponseDataRelationshipsRawSchemaDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityResponseDataRelationshipsRawSchemaDataType: valid values are %v", v, allowedEntityResponseDataRelationshipsRawSchemaDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityResponseDataRelationshipsRawSchemaDataType) IsValid() bool {
	for _, existing := range allowedEntityResponseDataRelationshipsRawSchemaDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityResponseDataRelationshipsRawSchemaDataType value.
func (v EntityResponseDataRelationshipsRawSchemaDataType) Ptr() *EntityResponseDataRelationshipsRawSchemaDataType {
	return &v
}
