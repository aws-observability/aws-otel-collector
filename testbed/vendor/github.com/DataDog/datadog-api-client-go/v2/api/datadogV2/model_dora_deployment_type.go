// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentType JSON:API type for DORA deployment events.
type DORADeploymentType string

// List of DORADeploymentType.
const (
	DORADEPLOYMENTTYPE_DORA_DEPLOYMENT DORADeploymentType = "dora_deployment"
)

var allowedDORADeploymentTypeEnumValues = []DORADeploymentType{
	DORADEPLOYMENTTYPE_DORA_DEPLOYMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DORADeploymentType) GetAllowedValues() []DORADeploymentType {
	return allowedDORADeploymentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DORADeploymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DORADeploymentType(value)
	return nil
}

// NewDORADeploymentTypeFromValue returns a pointer to a valid DORADeploymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDORADeploymentTypeFromValue(v string) (*DORADeploymentType, error) {
	ev := DORADeploymentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DORADeploymentType: valid values are %v", v, allowedDORADeploymentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DORADeploymentType) IsValid() bool {
	for _, existing := range allowedDORADeploymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DORADeploymentType value.
func (v DORADeploymentType) Ptr() *DORADeploymentType {
	return &v
}
