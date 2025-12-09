// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventCustomAttributesImpactedResourcesItemsType The type of the impacted resource.
type ChangeEventCustomAttributesImpactedResourcesItemsType string

// List of ChangeEventCustomAttributesImpactedResourcesItemsType.
const (
	CHANGEEVENTCUSTOMATTRIBUTESIMPACTEDRESOURCESITEMSTYPE_SERVICE ChangeEventCustomAttributesImpactedResourcesItemsType = "service"
)

var allowedChangeEventCustomAttributesImpactedResourcesItemsTypeEnumValues = []ChangeEventCustomAttributesImpactedResourcesItemsType{
	CHANGEEVENTCUSTOMATTRIBUTESIMPACTEDRESOURCESITEMSTYPE_SERVICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeEventCustomAttributesImpactedResourcesItemsType) GetAllowedValues() []ChangeEventCustomAttributesImpactedResourcesItemsType {
	return allowedChangeEventCustomAttributesImpactedResourcesItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeEventCustomAttributesImpactedResourcesItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeEventCustomAttributesImpactedResourcesItemsType(value)
	return nil
}

// NewChangeEventCustomAttributesImpactedResourcesItemsTypeFromValue returns a pointer to a valid ChangeEventCustomAttributesImpactedResourcesItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeEventCustomAttributesImpactedResourcesItemsTypeFromValue(v string) (*ChangeEventCustomAttributesImpactedResourcesItemsType, error) {
	ev := ChangeEventCustomAttributesImpactedResourcesItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeEventCustomAttributesImpactedResourcesItemsType: valid values are %v", v, allowedChangeEventCustomAttributesImpactedResourcesItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeEventCustomAttributesImpactedResourcesItemsType) IsValid() bool {
	for _, existing := range allowedChangeEventCustomAttributesImpactedResourcesItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeEventCustomAttributesImpactedResourcesItemsType value.
func (v ChangeEventCustomAttributesImpactedResourcesItemsType) Ptr() *ChangeEventCustomAttributesImpactedResourcesItemsType {
	return &v
}
