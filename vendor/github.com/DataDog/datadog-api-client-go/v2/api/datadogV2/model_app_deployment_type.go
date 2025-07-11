// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppDeploymentType The deployment type.
type AppDeploymentType string

// List of AppDeploymentType.
const (
	APPDEPLOYMENTTYPE_DEPLOYMENT AppDeploymentType = "deployment"
)

var allowedAppDeploymentTypeEnumValues = []AppDeploymentType{
	APPDEPLOYMENTTYPE_DEPLOYMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppDeploymentType) GetAllowedValues() []AppDeploymentType {
	return allowedAppDeploymentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppDeploymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppDeploymentType(value)
	return nil
}

// NewAppDeploymentTypeFromValue returns a pointer to a valid AppDeploymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppDeploymentTypeFromValue(v string) (*AppDeploymentType, error) {
	ev := AppDeploymentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppDeploymentType: valid values are %v", v, allowedAppDeploymentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppDeploymentType) IsValid() bool {
	for _, existing := range allowedAppDeploymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppDeploymentType value.
func (v AppDeploymentType) Ptr() *AppDeploymentType {
	return &v
}
