// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseNotificationRuleResourceType Notification rule resource type
type CaseNotificationRuleResourceType string

// List of CaseNotificationRuleResourceType.
const (
	CASENOTIFICATIONRULERESOURCETYPE_NOTIFICATION_RULE CaseNotificationRuleResourceType = "notification_rule"
)

var allowedCaseNotificationRuleResourceTypeEnumValues = []CaseNotificationRuleResourceType{
	CASENOTIFICATIONRULERESOURCETYPE_NOTIFICATION_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseNotificationRuleResourceType) GetAllowedValues() []CaseNotificationRuleResourceType {
	return allowedCaseNotificationRuleResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseNotificationRuleResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseNotificationRuleResourceType(value)
	return nil
}

// NewCaseNotificationRuleResourceTypeFromValue returns a pointer to a valid CaseNotificationRuleResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseNotificationRuleResourceTypeFromValue(v string) (*CaseNotificationRuleResourceType, error) {
	ev := CaseNotificationRuleResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseNotificationRuleResourceType: valid values are %v", v, allowedCaseNotificationRuleResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseNotificationRuleResourceType) IsValid() bool {
	for _, existing := range allowedCaseNotificationRuleResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseNotificationRuleResourceType value.
func (v CaseNotificationRuleResourceType) Ptr() *CaseNotificationRuleResourceType {
	return &v
}
