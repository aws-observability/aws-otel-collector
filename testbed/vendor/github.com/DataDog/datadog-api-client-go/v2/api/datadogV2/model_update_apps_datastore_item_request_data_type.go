// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppsDatastoreItemRequestDataType The resource type for datastore items.
type UpdateAppsDatastoreItemRequestDataType string

// List of UpdateAppsDatastoreItemRequestDataType.
const (
	UPDATEAPPSDATASTOREITEMREQUESTDATATYPE_ITEMS UpdateAppsDatastoreItemRequestDataType = "items"
)

var allowedUpdateAppsDatastoreItemRequestDataTypeEnumValues = []UpdateAppsDatastoreItemRequestDataType{
	UPDATEAPPSDATASTOREITEMREQUESTDATATYPE_ITEMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateAppsDatastoreItemRequestDataType) GetAllowedValues() []UpdateAppsDatastoreItemRequestDataType {
	return allowedUpdateAppsDatastoreItemRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateAppsDatastoreItemRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateAppsDatastoreItemRequestDataType(value)
	return nil
}

// NewUpdateAppsDatastoreItemRequestDataTypeFromValue returns a pointer to a valid UpdateAppsDatastoreItemRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateAppsDatastoreItemRequestDataTypeFromValue(v string) (*UpdateAppsDatastoreItemRequestDataType, error) {
	ev := UpdateAppsDatastoreItemRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateAppsDatastoreItemRequestDataType: valid values are %v", v, allowedUpdateAppsDatastoreItemRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateAppsDatastoreItemRequestDataType) IsValid() bool {
	for _, existing := range allowedUpdateAppsDatastoreItemRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateAppsDatastoreItemRequestDataType value.
func (v UpdateAppsDatastoreItemRequestDataType) Ptr() *UpdateAppsDatastoreItemRequestDataType {
	return &v
}
