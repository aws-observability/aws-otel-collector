// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTableRequestDataType Reference table resource type.
type CreateTableRequestDataType string

// List of CreateTableRequestDataType.
const (
	CREATETABLEREQUESTDATATYPE_REFERENCE_TABLE CreateTableRequestDataType = "reference_table"
)

var allowedCreateTableRequestDataTypeEnumValues = []CreateTableRequestDataType{
	CREATETABLEREQUESTDATATYPE_REFERENCE_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateTableRequestDataType) GetAllowedValues() []CreateTableRequestDataType {
	return allowedCreateTableRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateTableRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateTableRequestDataType(value)
	return nil
}

// NewCreateTableRequestDataTypeFromValue returns a pointer to a valid CreateTableRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateTableRequestDataTypeFromValue(v string) (*CreateTableRequestDataType, error) {
	ev := CreateTableRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateTableRequestDataType: valid values are %v", v, allowedCreateTableRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateTableRequestDataType) IsValid() bool {
	for _, existing := range allowedCreateTableRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateTableRequestDataType value.
func (v CreateTableRequestDataType) Ptr() *CreateTableRequestDataType {
	return &v
}
