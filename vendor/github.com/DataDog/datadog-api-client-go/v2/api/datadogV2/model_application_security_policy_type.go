// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityPolicyType The type of the resource. The value should always be `policy`.
type ApplicationSecurityPolicyType string

// List of ApplicationSecurityPolicyType.
const (
	APPLICATIONSECURITYPOLICYTYPE_POLICY ApplicationSecurityPolicyType = "policy"
)

var allowedApplicationSecurityPolicyTypeEnumValues = []ApplicationSecurityPolicyType{
	APPLICATIONSECURITYPOLICYTYPE_POLICY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApplicationSecurityPolicyType) GetAllowedValues() []ApplicationSecurityPolicyType {
	return allowedApplicationSecurityPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApplicationSecurityPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApplicationSecurityPolicyType(value)
	return nil
}

// NewApplicationSecurityPolicyTypeFromValue returns a pointer to a valid ApplicationSecurityPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApplicationSecurityPolicyTypeFromValue(v string) (*ApplicationSecurityPolicyType, error) {
	ev := ApplicationSecurityPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApplicationSecurityPolicyType: valid values are %v", v, allowedApplicationSecurityPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApplicationSecurityPolicyType) IsValid() bool {
	for _, existing := range allowedApplicationSecurityPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSecurityPolicyType value.
func (v ApplicationSecurityPolicyType) Ptr() *ApplicationSecurityPolicyType {
	return &v
}
