// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BulkDeleteAppsDatastoreItemsRequestDataType Items resource type.
type BulkDeleteAppsDatastoreItemsRequestDataType string

// List of BulkDeleteAppsDatastoreItemsRequestDataType.
const (
	BULKDELETEAPPSDATASTOREITEMSREQUESTDATATYPE_ITEMS BulkDeleteAppsDatastoreItemsRequestDataType = "items"
)

var allowedBulkDeleteAppsDatastoreItemsRequestDataTypeEnumValues = []BulkDeleteAppsDatastoreItemsRequestDataType{
	BULKDELETEAPPSDATASTOREITEMSREQUESTDATATYPE_ITEMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BulkDeleteAppsDatastoreItemsRequestDataType) GetAllowedValues() []BulkDeleteAppsDatastoreItemsRequestDataType {
	return allowedBulkDeleteAppsDatastoreItemsRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BulkDeleteAppsDatastoreItemsRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BulkDeleteAppsDatastoreItemsRequestDataType(value)
	return nil
}

// NewBulkDeleteAppsDatastoreItemsRequestDataTypeFromValue returns a pointer to a valid BulkDeleteAppsDatastoreItemsRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBulkDeleteAppsDatastoreItemsRequestDataTypeFromValue(v string) (*BulkDeleteAppsDatastoreItemsRequestDataType, error) {
	ev := BulkDeleteAppsDatastoreItemsRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BulkDeleteAppsDatastoreItemsRequestDataType: valid values are %v", v, allowedBulkDeleteAppsDatastoreItemsRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BulkDeleteAppsDatastoreItemsRequestDataType) IsValid() bool {
	for _, existing := range allowedBulkDeleteAppsDatastoreItemsRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BulkDeleteAppsDatastoreItemsRequestDataType value.
func (v BulkDeleteAppsDatastoreItemsRequestDataType) Ptr() *BulkDeleteAppsDatastoreItemsRequestDataType {
	return &v
}
