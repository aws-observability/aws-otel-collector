// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlResourceType JSON:API resource type for a governance control.
type GovernanceControlResourceType string

// List of GovernanceControlResourceType.
const (
	GOVERNANCECONTROLRESOURCETYPE_GOVERNANCE_CONTROL GovernanceControlResourceType = "governance_control"
)

var allowedGovernanceControlResourceTypeEnumValues = []GovernanceControlResourceType{
	GOVERNANCECONTROLRESOURCETYPE_GOVERNANCE_CONTROL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GovernanceControlResourceType) GetAllowedValues() []GovernanceControlResourceType {
	return allowedGovernanceControlResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GovernanceControlResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GovernanceControlResourceType(value)
	return nil
}

// NewGovernanceControlResourceTypeFromValue returns a pointer to a valid GovernanceControlResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGovernanceControlResourceTypeFromValue(v string) (*GovernanceControlResourceType, error) {
	ev := GovernanceControlResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GovernanceControlResourceType: valid values are %v", v, allowedGovernanceControlResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GovernanceControlResourceType) IsValid() bool {
	for _, existing := range allowedGovernanceControlResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GovernanceControlResourceType value.
func (v GovernanceControlResourceType) Ptr() *GovernanceControlResourceType {
	return &v
}
