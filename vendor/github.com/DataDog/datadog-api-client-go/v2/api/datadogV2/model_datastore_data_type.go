// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatastoreDataType The resource type for datastores.
type DatastoreDataType string

// List of DatastoreDataType.
const (
	DATASTOREDATATYPE_DATASTORES DatastoreDataType = "datastores"
)

var allowedDatastoreDataTypeEnumValues = []DatastoreDataType{
	DATASTOREDATATYPE_DATASTORES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatastoreDataType) GetAllowedValues() []DatastoreDataType {
	return allowedDatastoreDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatastoreDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatastoreDataType(value)
	return nil
}

// NewDatastoreDataTypeFromValue returns a pointer to a valid DatastoreDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatastoreDataTypeFromValue(v string) (*DatastoreDataType, error) {
	ev := DatastoreDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatastoreDataType: valid values are %v", v, allowedDatastoreDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatastoreDataType) IsValid() bool {
	for _, existing := range allowedDatastoreDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatastoreDataType value.
func (v DatastoreDataType) Ptr() *DatastoreDataType {
	return &v
}
