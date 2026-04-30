// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityEntityRiskScoreAttributesSeverity Severity level based on risk score
type SecurityEntityRiskScoreAttributesSeverity string

// List of SecurityEntityRiskScoreAttributesSeverity.
const (
	SECURITYENTITYRISKSCOREATTRIBUTESSEVERITY_CRITICAL SecurityEntityRiskScoreAttributesSeverity = "critical"
	SECURITYENTITYRISKSCOREATTRIBUTESSEVERITY_HIGH     SecurityEntityRiskScoreAttributesSeverity = "high"
	SECURITYENTITYRISKSCOREATTRIBUTESSEVERITY_MEDIUM   SecurityEntityRiskScoreAttributesSeverity = "medium"
	SECURITYENTITYRISKSCOREATTRIBUTESSEVERITY_LOW      SecurityEntityRiskScoreAttributesSeverity = "low"
	SECURITYENTITYRISKSCOREATTRIBUTESSEVERITY_INFO     SecurityEntityRiskScoreAttributesSeverity = "info"
)

var allowedSecurityEntityRiskScoreAttributesSeverityEnumValues = []SecurityEntityRiskScoreAttributesSeverity{
	SECURITYENTITYRISKSCOREATTRIBUTESSEVERITY_CRITICAL,
	SECURITYENTITYRISKSCOREATTRIBUTESSEVERITY_HIGH,
	SECURITYENTITYRISKSCOREATTRIBUTESSEVERITY_MEDIUM,
	SECURITYENTITYRISKSCOREATTRIBUTESSEVERITY_LOW,
	SECURITYENTITYRISKSCOREATTRIBUTESSEVERITY_INFO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityEntityRiskScoreAttributesSeverity) GetAllowedValues() []SecurityEntityRiskScoreAttributesSeverity {
	return allowedSecurityEntityRiskScoreAttributesSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityEntityRiskScoreAttributesSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityEntityRiskScoreAttributesSeverity(value)
	return nil
}

// NewSecurityEntityRiskScoreAttributesSeverityFromValue returns a pointer to a valid SecurityEntityRiskScoreAttributesSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityEntityRiskScoreAttributesSeverityFromValue(v string) (*SecurityEntityRiskScoreAttributesSeverity, error) {
	ev := SecurityEntityRiskScoreAttributesSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityEntityRiskScoreAttributesSeverity: valid values are %v", v, allowedSecurityEntityRiskScoreAttributesSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityEntityRiskScoreAttributesSeverity) IsValid() bool {
	for _, existing := range allowedSecurityEntityRiskScoreAttributesSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityEntityRiskScoreAttributesSeverity value.
func (v SecurityEntityRiskScoreAttributesSeverity) Ptr() *SecurityEntityRiskScoreAttributesSeverity {
	return &v
}
