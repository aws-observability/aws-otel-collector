// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeployAppResponseDataType The definition of `DeployAppResponseDataType` object.
type DeployAppResponseDataType string

// List of DeployAppResponseDataType.
const (
	DEPLOYAPPRESPONSEDATATYPE_DEPLOYMENT DeployAppResponseDataType = "deployment"
)

var allowedDeployAppResponseDataTypeEnumValues = []DeployAppResponseDataType{
	DEPLOYAPPRESPONSEDATATYPE_DEPLOYMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeployAppResponseDataType) GetAllowedValues() []DeployAppResponseDataType {
	return allowedDeployAppResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeployAppResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeployAppResponseDataType(value)
	return nil
}

// NewDeployAppResponseDataTypeFromValue returns a pointer to a valid DeployAppResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeployAppResponseDataTypeFromValue(v string) (*DeployAppResponseDataType, error) {
	ev := DeployAppResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeployAppResponseDataType: valid values are %v", v, allowedDeployAppResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeployAppResponseDataType) IsValid() bool {
	for _, existing := range allowedDeployAppResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeployAppResponseDataType value.
func (v DeployAppResponseDataType) Ptr() *DeployAppResponseDataType {
	return &v
}
