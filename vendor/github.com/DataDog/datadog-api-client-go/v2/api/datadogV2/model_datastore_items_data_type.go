// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatastoreItemsDataType The resource type for datastore items.
type DatastoreItemsDataType string

// List of DatastoreItemsDataType.
const (
	DATASTOREITEMSDATATYPE_ITEMS DatastoreItemsDataType = "items"
)

var allowedDatastoreItemsDataTypeEnumValues = []DatastoreItemsDataType{
	DATASTOREITEMSDATATYPE_ITEMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatastoreItemsDataType) GetAllowedValues() []DatastoreItemsDataType {
	return allowedDatastoreItemsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatastoreItemsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatastoreItemsDataType(value)
	return nil
}

// NewDatastoreItemsDataTypeFromValue returns a pointer to a valid DatastoreItemsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatastoreItemsDataTypeFromValue(v string) (*DatastoreItemsDataType, error) {
	ev := DatastoreItemsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatastoreItemsDataType: valid values are %v", v, allowedDatastoreItemsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatastoreItemsDataType) IsValid() bool {
	for _, existing := range allowedDatastoreItemsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatastoreItemsDataType value.
func (v DatastoreItemsDataType) Ptr() *DatastoreItemsDataType {
	return &v
}
