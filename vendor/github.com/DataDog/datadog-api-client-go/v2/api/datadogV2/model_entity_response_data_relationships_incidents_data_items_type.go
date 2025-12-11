// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseDataRelationshipsIncidentsDataItemsType Incident resource type.
type EntityResponseDataRelationshipsIncidentsDataItemsType string

// List of EntityResponseDataRelationshipsIncidentsDataItemsType.
const (
	ENTITYRESPONSEDATARELATIONSHIPSINCIDENTSDATAITEMSTYPE_INCIDENT EntityResponseDataRelationshipsIncidentsDataItemsType = "incident"
)

var allowedEntityResponseDataRelationshipsIncidentsDataItemsTypeEnumValues = []EntityResponseDataRelationshipsIncidentsDataItemsType{
	ENTITYRESPONSEDATARELATIONSHIPSINCIDENTSDATAITEMSTYPE_INCIDENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityResponseDataRelationshipsIncidentsDataItemsType) GetAllowedValues() []EntityResponseDataRelationshipsIncidentsDataItemsType {
	return allowedEntityResponseDataRelationshipsIncidentsDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityResponseDataRelationshipsIncidentsDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityResponseDataRelationshipsIncidentsDataItemsType(value)
	return nil
}

// NewEntityResponseDataRelationshipsIncidentsDataItemsTypeFromValue returns a pointer to a valid EntityResponseDataRelationshipsIncidentsDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityResponseDataRelationshipsIncidentsDataItemsTypeFromValue(v string) (*EntityResponseDataRelationshipsIncidentsDataItemsType, error) {
	ev := EntityResponseDataRelationshipsIncidentsDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityResponseDataRelationshipsIncidentsDataItemsType: valid values are %v", v, allowedEntityResponseDataRelationshipsIncidentsDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityResponseDataRelationshipsIncidentsDataItemsType) IsValid() bool {
	for _, existing := range allowedEntityResponseDataRelationshipsIncidentsDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityResponseDataRelationshipsIncidentsDataItemsType value.
func (v EntityResponseDataRelationshipsIncidentsDataItemsType) Ptr() *EntityResponseDataRelationshipsIncidentsDataItemsType {
	return &v
}
