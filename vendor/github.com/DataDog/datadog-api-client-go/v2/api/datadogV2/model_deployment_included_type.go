// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentIncludedType The definition of `DeploymentIncludedType` object.
type DeploymentIncludedType string

// List of DeploymentIncludedType.
const (
	DEPLOYMENTINCLUDEDTYPE_DEPLOYMENT DeploymentIncludedType = "deployment"
)

var allowedDeploymentIncludedTypeEnumValues = []DeploymentIncludedType{
	DEPLOYMENTINCLUDEDTYPE_DEPLOYMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentIncludedType) GetAllowedValues() []DeploymentIncludedType {
	return allowedDeploymentIncludedTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentIncludedType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentIncludedType(value)
	return nil
}

// NewDeploymentIncludedTypeFromValue returns a pointer to a valid DeploymentIncludedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentIncludedTypeFromValue(v string) (*DeploymentIncludedType, error) {
	ev := DeploymentIncludedType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentIncludedType: valid values are %v", v, allowedDeploymentIncludedTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentIncludedType) IsValid() bool {
	for _, existing := range allowedDeploymentIncludedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentIncludedType value.
func (v DeploymentIncludedType) Ptr() *DeploymentIncludedType {
	return &v
}
