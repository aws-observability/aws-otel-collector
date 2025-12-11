// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseDataRelationshipsOncallsDataItemsType Oncall resource type.
type EntityResponseDataRelationshipsOncallsDataItemsType string

// List of EntityResponseDataRelationshipsOncallsDataItemsType.
const (
	ENTITYRESPONSEDATARELATIONSHIPSONCALLSDATAITEMSTYPE_ONCALL EntityResponseDataRelationshipsOncallsDataItemsType = "oncall"
)

var allowedEntityResponseDataRelationshipsOncallsDataItemsTypeEnumValues = []EntityResponseDataRelationshipsOncallsDataItemsType{
	ENTITYRESPONSEDATARELATIONSHIPSONCALLSDATAITEMSTYPE_ONCALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityResponseDataRelationshipsOncallsDataItemsType) GetAllowedValues() []EntityResponseDataRelationshipsOncallsDataItemsType {
	return allowedEntityResponseDataRelationshipsOncallsDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityResponseDataRelationshipsOncallsDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityResponseDataRelationshipsOncallsDataItemsType(value)
	return nil
}

// NewEntityResponseDataRelationshipsOncallsDataItemsTypeFromValue returns a pointer to a valid EntityResponseDataRelationshipsOncallsDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityResponseDataRelationshipsOncallsDataItemsTypeFromValue(v string) (*EntityResponseDataRelationshipsOncallsDataItemsType, error) {
	ev := EntityResponseDataRelationshipsOncallsDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityResponseDataRelationshipsOncallsDataItemsType: valid values are %v", v, allowedEntityResponseDataRelationshipsOncallsDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityResponseDataRelationshipsOncallsDataItemsType) IsValid() bool {
	for _, existing := range allowedEntityResponseDataRelationshipsOncallsDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityResponseDataRelationshipsOncallsDataItemsType value.
func (v EntityResponseDataRelationshipsOncallsDataItemsType) Ptr() *EntityResponseDataRelationshipsOncallsDataItemsType {
	return &v
}
