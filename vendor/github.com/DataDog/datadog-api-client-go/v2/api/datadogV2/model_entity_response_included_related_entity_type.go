// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedRelatedEntityType Related entity.
type EntityResponseIncludedRelatedEntityType string

// List of EntityResponseIncludedRelatedEntityType.
const (
	ENTITYRESPONSEINCLUDEDRELATEDENTITYTYPE_RELATED_ENTITY EntityResponseIncludedRelatedEntityType = "relatedEntity"
)

var allowedEntityResponseIncludedRelatedEntityTypeEnumValues = []EntityResponseIncludedRelatedEntityType{
	ENTITYRESPONSEINCLUDEDRELATEDENTITYTYPE_RELATED_ENTITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityResponseIncludedRelatedEntityType) GetAllowedValues() []EntityResponseIncludedRelatedEntityType {
	return allowedEntityResponseIncludedRelatedEntityTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityResponseIncludedRelatedEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityResponseIncludedRelatedEntityType(value)
	return nil
}

// NewEntityResponseIncludedRelatedEntityTypeFromValue returns a pointer to a valid EntityResponseIncludedRelatedEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityResponseIncludedRelatedEntityTypeFromValue(v string) (*EntityResponseIncludedRelatedEntityType, error) {
	ev := EntityResponseIncludedRelatedEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityResponseIncludedRelatedEntityType: valid values are %v", v, allowedEntityResponseIncludedRelatedEntityTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityResponseIncludedRelatedEntityType) IsValid() bool {
	for _, existing := range allowedEntityResponseIncludedRelatedEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityResponseIncludedRelatedEntityType value.
func (v EntityResponseIncludedRelatedEntityType) Ptr() *EntityResponseIncludedRelatedEntityType {
	return &v
}
