// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppRequestDataType The definition of `UpdateAppRequestDataType` object.
type UpdateAppRequestDataType string

// List of UpdateAppRequestDataType.
const (
	UPDATEAPPREQUESTDATATYPE_APPDEFINITIONS UpdateAppRequestDataType = "appDefinitions"
)

var allowedUpdateAppRequestDataTypeEnumValues = []UpdateAppRequestDataType{
	UPDATEAPPREQUESTDATATYPE_APPDEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateAppRequestDataType) GetAllowedValues() []UpdateAppRequestDataType {
	return allowedUpdateAppRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateAppRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateAppRequestDataType(value)
	return nil
}

// NewUpdateAppRequestDataTypeFromValue returns a pointer to a valid UpdateAppRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateAppRequestDataTypeFromValue(v string) (*UpdateAppRequestDataType, error) {
	ev := UpdateAppRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateAppRequestDataType: valid values are %v", v, allowedUpdateAppRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateAppRequestDataType) IsValid() bool {
	for _, existing := range allowedUpdateAppRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateAppRequestDataType value.
func (v UpdateAppRequestDataType) Ptr() *UpdateAppRequestDataType {
	return &v
}
