// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAppsResponseDataItemsType The definition of `ListAppsResponseDataItemsType` object.
type ListAppsResponseDataItemsType string

// List of ListAppsResponseDataItemsType.
const (
	LISTAPPSRESPONSEDATAITEMSTYPE_APPDEFINITIONS ListAppsResponseDataItemsType = "appDefinitions"
)

var allowedListAppsResponseDataItemsTypeEnumValues = []ListAppsResponseDataItemsType{
	LISTAPPSRESPONSEDATAITEMSTYPE_APPDEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListAppsResponseDataItemsType) GetAllowedValues() []ListAppsResponseDataItemsType {
	return allowedListAppsResponseDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListAppsResponseDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListAppsResponseDataItemsType(value)
	return nil
}

// NewListAppsResponseDataItemsTypeFromValue returns a pointer to a valid ListAppsResponseDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListAppsResponseDataItemsTypeFromValue(v string) (*ListAppsResponseDataItemsType, error) {
	ev := ListAppsResponseDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListAppsResponseDataItemsType: valid values are %v", v, allowedListAppsResponseDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListAppsResponseDataItemsType) IsValid() bool {
	for _, existing := range allowedListAppsResponseDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListAppsResponseDataItemsType value.
func (v ListAppsResponseDataItemsType) Ptr() *ListAppsResponseDataItemsType {
	return &v
}
