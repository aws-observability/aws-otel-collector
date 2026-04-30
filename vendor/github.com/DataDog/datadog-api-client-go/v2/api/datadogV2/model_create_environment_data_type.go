// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateEnvironmentDataType The resource type.
type CreateEnvironmentDataType string

// List of CreateEnvironmentDataType.
const (
	CREATEENVIRONMENTDATATYPE_ENVIRONMENTS CreateEnvironmentDataType = "environments"
)

var allowedCreateEnvironmentDataTypeEnumValues = []CreateEnvironmentDataType{
	CREATEENVIRONMENTDATATYPE_ENVIRONMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateEnvironmentDataType) GetAllowedValues() []CreateEnvironmentDataType {
	return allowedCreateEnvironmentDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateEnvironmentDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateEnvironmentDataType(value)
	return nil
}

// NewCreateEnvironmentDataTypeFromValue returns a pointer to a valid CreateEnvironmentDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateEnvironmentDataTypeFromValue(v string) (*CreateEnvironmentDataType, error) {
	ev := CreateEnvironmentDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateEnvironmentDataType: valid values are %v", v, allowedCreateEnvironmentDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateEnvironmentDataType) IsValid() bool {
	for _, existing := range allowedCreateEnvironmentDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateEnvironmentDataType value.
func (v CreateEnvironmentDataType) Ptr() *CreateEnvironmentDataType {
	return &v
}
