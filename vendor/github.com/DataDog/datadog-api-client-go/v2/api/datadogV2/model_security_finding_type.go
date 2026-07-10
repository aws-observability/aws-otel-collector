// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFindingType The type of security finding that the automation rule applies to.
type SecurityFindingType string

// List of SecurityFindingType.
const (
	SECURITYFINDINGTYPE_API_SECURITY                     SecurityFindingType = "api_security"
	SECURITYFINDINGTYPE_ATTACK_PATH                      SecurityFindingType = "attack_path"
	SECURITYFINDINGTYPE_HOST_AND_CONTAINER_VULNERABILITY SecurityFindingType = "host_and_container_vulnerability"
	SECURITYFINDINGTYPE_IAC_MISCONFIGURATION             SecurityFindingType = "iac_misconfiguration"
	SECURITYFINDINGTYPE_IDENTITY_RISK                    SecurityFindingType = "identity_risk"
	SECURITYFINDINGTYPE_LIBRARY_VULNERABILITY            SecurityFindingType = "library_vulnerability"
	SECURITYFINDINGTYPE_MISCONFIGURATION                 SecurityFindingType = "misconfiguration"
	SECURITYFINDINGTYPE_RUNTIME_CODE_VULNERABILITY       SecurityFindingType = "runtime_code_vulnerability"
	SECURITYFINDINGTYPE_SECRET                           SecurityFindingType = "secret"
	SECURITYFINDINGTYPE_STATIC_CODE_VULNERABILITY        SecurityFindingType = "static_code_vulnerability"
	SECURITYFINDINGTYPE_WORKLOAD_ACTIVITY                SecurityFindingType = "workload_activity"
)

var allowedSecurityFindingTypeEnumValues = []SecurityFindingType{
	SECURITYFINDINGTYPE_API_SECURITY,
	SECURITYFINDINGTYPE_ATTACK_PATH,
	SECURITYFINDINGTYPE_HOST_AND_CONTAINER_VULNERABILITY,
	SECURITYFINDINGTYPE_IAC_MISCONFIGURATION,
	SECURITYFINDINGTYPE_IDENTITY_RISK,
	SECURITYFINDINGTYPE_LIBRARY_VULNERABILITY,
	SECURITYFINDINGTYPE_MISCONFIGURATION,
	SECURITYFINDINGTYPE_RUNTIME_CODE_VULNERABILITY,
	SECURITYFINDINGTYPE_SECRET,
	SECURITYFINDINGTYPE_STATIC_CODE_VULNERABILITY,
	SECURITYFINDINGTYPE_WORKLOAD_ACTIVITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityFindingType) GetAllowedValues() []SecurityFindingType {
	return allowedSecurityFindingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityFindingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityFindingType(value)
	return nil
}

// NewSecurityFindingTypeFromValue returns a pointer to a valid SecurityFindingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityFindingTypeFromValue(v string) (*SecurityFindingType, error) {
	ev := SecurityFindingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityFindingType: valid values are %v", v, allowedSecurityFindingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityFindingType) IsValid() bool {
	for _, existing := range allowedSecurityFindingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityFindingType value.
func (v SecurityFindingType) Ptr() *SecurityFindingType {
	return &v
}
