// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateEnvironmentDataType The resource type.
type UpdateEnvironmentDataType string

// List of UpdateEnvironmentDataType.
const (
	UPDATEENVIRONMENTDATATYPE_ENVIRONMENTS UpdateEnvironmentDataType = "environments"
)

var allowedUpdateEnvironmentDataTypeEnumValues = []UpdateEnvironmentDataType{
	UPDATEENVIRONMENTDATATYPE_ENVIRONMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateEnvironmentDataType) GetAllowedValues() []UpdateEnvironmentDataType {
	return allowedUpdateEnvironmentDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateEnvironmentDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateEnvironmentDataType(value)
	return nil
}

// NewUpdateEnvironmentDataTypeFromValue returns a pointer to a valid UpdateEnvironmentDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateEnvironmentDataTypeFromValue(v string) (*UpdateEnvironmentDataType, error) {
	ev := UpdateEnvironmentDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateEnvironmentDataType: valid values are %v", v, allowedUpdateEnvironmentDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateEnvironmentDataType) IsValid() bool {
	for _, existing := range allowedUpdateEnvironmentDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateEnvironmentDataType value.
func (v UpdateEnvironmentDataType) Ptr() *UpdateEnvironmentDataType {
	return &v
}
