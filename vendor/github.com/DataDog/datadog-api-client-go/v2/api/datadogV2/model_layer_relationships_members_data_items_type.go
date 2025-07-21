// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LayerRelationshipsMembersDataItemsType Members resource type.
type LayerRelationshipsMembersDataItemsType string

// List of LayerRelationshipsMembersDataItemsType.
const (
	LAYERRELATIONSHIPSMEMBERSDATAITEMSTYPE_MEMBERS LayerRelationshipsMembersDataItemsType = "members"
)

var allowedLayerRelationshipsMembersDataItemsTypeEnumValues = []LayerRelationshipsMembersDataItemsType{
	LAYERRELATIONSHIPSMEMBERSDATAITEMSTYPE_MEMBERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LayerRelationshipsMembersDataItemsType) GetAllowedValues() []LayerRelationshipsMembersDataItemsType {
	return allowedLayerRelationshipsMembersDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LayerRelationshipsMembersDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LayerRelationshipsMembersDataItemsType(value)
	return nil
}

// NewLayerRelationshipsMembersDataItemsTypeFromValue returns a pointer to a valid LayerRelationshipsMembersDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLayerRelationshipsMembersDataItemsTypeFromValue(v string) (*LayerRelationshipsMembersDataItemsType, error) {
	ev := LayerRelationshipsMembersDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LayerRelationshipsMembersDataItemsType: valid values are %v", v, allowedLayerRelationshipsMembersDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LayerRelationshipsMembersDataItemsType) IsValid() bool {
	for _, existing := range allowedLayerRelationshipsMembersDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LayerRelationshipsMembersDataItemsType value.
func (v LayerRelationshipsMembersDataItemsType) Ptr() *LayerRelationshipsMembersDataItemsType {
	return &v
}
