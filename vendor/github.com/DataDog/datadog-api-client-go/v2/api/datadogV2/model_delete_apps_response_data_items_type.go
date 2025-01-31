// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeleteAppsResponseDataItemsType The definition of `DeleteAppsResponseDataItemsType` object.
type DeleteAppsResponseDataItemsType string

// List of DeleteAppsResponseDataItemsType.
const (
	DELETEAPPSRESPONSEDATAITEMSTYPE_APPDEFINITIONS DeleteAppsResponseDataItemsType = "appDefinitions"
)

var allowedDeleteAppsResponseDataItemsTypeEnumValues = []DeleteAppsResponseDataItemsType{
	DELETEAPPSRESPONSEDATAITEMSTYPE_APPDEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeleteAppsResponseDataItemsType) GetAllowedValues() []DeleteAppsResponseDataItemsType {
	return allowedDeleteAppsResponseDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeleteAppsResponseDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeleteAppsResponseDataItemsType(value)
	return nil
}

// NewDeleteAppsResponseDataItemsTypeFromValue returns a pointer to a valid DeleteAppsResponseDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeleteAppsResponseDataItemsTypeFromValue(v string) (*DeleteAppsResponseDataItemsType, error) {
	ev := DeleteAppsResponseDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeleteAppsResponseDataItemsType: valid values are %v", v, allowedDeleteAppsResponseDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeleteAppsResponseDataItemsType) IsValid() bool {
	for _, existing := range allowedDeleteAppsResponseDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeleteAppsResponseDataItemsType value.
func (v DeleteAppsResponseDataItemsType) Ptr() *DeleteAppsResponseDataItemsType {
	return &v
}
