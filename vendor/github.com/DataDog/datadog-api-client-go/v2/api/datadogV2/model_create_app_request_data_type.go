// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppRequestDataType The definition of `CreateAppRequestDataType` object.
type CreateAppRequestDataType string

// List of CreateAppRequestDataType.
const (
	CREATEAPPREQUESTDATATYPE_APPDEFINITIONS CreateAppRequestDataType = "appDefinitions"
)

var allowedCreateAppRequestDataTypeEnumValues = []CreateAppRequestDataType{
	CREATEAPPREQUESTDATATYPE_APPDEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateAppRequestDataType) GetAllowedValues() []CreateAppRequestDataType {
	return allowedCreateAppRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateAppRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateAppRequestDataType(value)
	return nil
}

// NewCreateAppRequestDataTypeFromValue returns a pointer to a valid CreateAppRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateAppRequestDataTypeFromValue(v string) (*CreateAppRequestDataType, error) {
	ev := CreateAppRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateAppRequestDataType: valid values are %v", v, allowedCreateAppRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateAppRequestDataType) IsValid() bool {
	for _, existing := range allowedCreateAppRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateAppRequestDataType value.
func (v CreateAppRequestDataType) Ptr() *CreateAppRequestDataType {
	return &v
}
