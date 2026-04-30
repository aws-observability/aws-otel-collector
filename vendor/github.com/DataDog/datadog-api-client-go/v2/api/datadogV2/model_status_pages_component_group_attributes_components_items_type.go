// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesComponentGroupAttributesComponentsItemsType The type of the component.
type StatusPagesComponentGroupAttributesComponentsItemsType string

// List of StatusPagesComponentGroupAttributesComponentsItemsType.
const (
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSTYPE_COMPONENT StatusPagesComponentGroupAttributesComponentsItemsType = "component"
)

var allowedStatusPagesComponentGroupAttributesComponentsItemsTypeEnumValues = []StatusPagesComponentGroupAttributesComponentsItemsType{
	STATUSPAGESCOMPONENTGROUPATTRIBUTESCOMPONENTSITEMSTYPE_COMPONENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StatusPagesComponentGroupAttributesComponentsItemsType) GetAllowedValues() []StatusPagesComponentGroupAttributesComponentsItemsType {
	return allowedStatusPagesComponentGroupAttributesComponentsItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StatusPagesComponentGroupAttributesComponentsItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StatusPagesComponentGroupAttributesComponentsItemsType(value)
	return nil
}

// NewStatusPagesComponentGroupAttributesComponentsItemsTypeFromValue returns a pointer to a valid StatusPagesComponentGroupAttributesComponentsItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatusPagesComponentGroupAttributesComponentsItemsTypeFromValue(v string) (*StatusPagesComponentGroupAttributesComponentsItemsType, error) {
	ev := StatusPagesComponentGroupAttributesComponentsItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StatusPagesComponentGroupAttributesComponentsItemsType: valid values are %v", v, allowedStatusPagesComponentGroupAttributesComponentsItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StatusPagesComponentGroupAttributesComponentsItemsType) IsValid() bool {
	for _, existing := range allowedStatusPagesComponentGroupAttributesComponentsItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusPagesComponentGroupAttributesComponentsItemsType value.
func (v StatusPagesComponentGroupAttributesComponentsItemsType) Ptr() *StatusPagesComponentGroupAttributesComponentsItemsType {
	return &v
}
