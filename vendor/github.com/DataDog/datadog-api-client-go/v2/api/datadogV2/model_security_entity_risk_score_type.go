// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityEntityRiskScoreType Resource type
type SecurityEntityRiskScoreType string

// List of SecurityEntityRiskScoreType.
const (
	SECURITYENTITYRISKSCORETYPE_SECURITY_ENTITY_RISK_SCORE SecurityEntityRiskScoreType = "security_entity_risk_score"
)

var allowedSecurityEntityRiskScoreTypeEnumValues = []SecurityEntityRiskScoreType{
	SECURITYENTITYRISKSCORETYPE_SECURITY_ENTITY_RISK_SCORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityEntityRiskScoreType) GetAllowedValues() []SecurityEntityRiskScoreType {
	return allowedSecurityEntityRiskScoreTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityEntityRiskScoreType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityEntityRiskScoreType(value)
	return nil
}

// NewSecurityEntityRiskScoreTypeFromValue returns a pointer to a valid SecurityEntityRiskScoreType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityEntityRiskScoreTypeFromValue(v string) (*SecurityEntityRiskScoreType, error) {
	ev := SecurityEntityRiskScoreType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityEntityRiskScoreType: valid values are %v", v, allowedSecurityEntityRiskScoreTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityEntityRiskScoreType) IsValid() bool {
	for _, existing := range allowedSecurityEntityRiskScoreTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityEntityRiskScoreType value.
func (v SecurityEntityRiskScoreType) Ptr() *SecurityEntityRiskScoreType {
	return &v
}
