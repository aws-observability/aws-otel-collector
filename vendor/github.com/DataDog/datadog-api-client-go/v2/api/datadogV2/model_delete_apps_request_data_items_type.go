// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeleteAppsRequestDataItemsType The definition of `DeleteAppsRequestDataItemsType` object.
type DeleteAppsRequestDataItemsType string

// List of DeleteAppsRequestDataItemsType.
const (
	DELETEAPPSREQUESTDATAITEMSTYPE_APPDEFINITIONS DeleteAppsRequestDataItemsType = "appDefinitions"
)

var allowedDeleteAppsRequestDataItemsTypeEnumValues = []DeleteAppsRequestDataItemsType{
	DELETEAPPSREQUESTDATAITEMSTYPE_APPDEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeleteAppsRequestDataItemsType) GetAllowedValues() []DeleteAppsRequestDataItemsType {
	return allowedDeleteAppsRequestDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeleteAppsRequestDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeleteAppsRequestDataItemsType(value)
	return nil
}

// NewDeleteAppsRequestDataItemsTypeFromValue returns a pointer to a valid DeleteAppsRequestDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeleteAppsRequestDataItemsTypeFromValue(v string) (*DeleteAppsRequestDataItemsType, error) {
	ev := DeleteAppsRequestDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeleteAppsRequestDataItemsType: valid values are %v", v, allowedDeleteAppsRequestDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeleteAppsRequestDataItemsType) IsValid() bool {
	for _, existing := range allowedDeleteAppsRequestDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeleteAppsRequestDataItemsType value.
func (v DeleteAppsRequestDataItemsType) Ptr() *DeleteAppsRequestDataItemsType {
	return &v
}
