// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationRuleAttributesVisibility The visibility of the notification rule.
type IncidentNotificationRuleAttributesVisibility string

// List of IncidentNotificationRuleAttributesVisibility.
const (
	INCIDENTNOTIFICATIONRULEATTRIBUTESVISIBILITY_ALL          IncidentNotificationRuleAttributesVisibility = "all"
	INCIDENTNOTIFICATIONRULEATTRIBUTESVISIBILITY_ORGANIZATION IncidentNotificationRuleAttributesVisibility = "organization"
	INCIDENTNOTIFICATIONRULEATTRIBUTESVISIBILITY_PRIVATE      IncidentNotificationRuleAttributesVisibility = "private"
)

var allowedIncidentNotificationRuleAttributesVisibilityEnumValues = []IncidentNotificationRuleAttributesVisibility{
	INCIDENTNOTIFICATIONRULEATTRIBUTESVISIBILITY_ALL,
	INCIDENTNOTIFICATIONRULEATTRIBUTESVISIBILITY_ORGANIZATION,
	INCIDENTNOTIFICATIONRULEATTRIBUTESVISIBILITY_PRIVATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentNotificationRuleAttributesVisibility) GetAllowedValues() []IncidentNotificationRuleAttributesVisibility {
	return allowedIncidentNotificationRuleAttributesVisibilityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentNotificationRuleAttributesVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentNotificationRuleAttributesVisibility(value)
	return nil
}

// NewIncidentNotificationRuleAttributesVisibilityFromValue returns a pointer to a valid IncidentNotificationRuleAttributesVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentNotificationRuleAttributesVisibilityFromValue(v string) (*IncidentNotificationRuleAttributesVisibility, error) {
	ev := IncidentNotificationRuleAttributesVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentNotificationRuleAttributesVisibility: valid values are %v", v, allowedIncidentNotificationRuleAttributesVisibilityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentNotificationRuleAttributesVisibility) IsValid() bool {
	for _, existing := range allowedIncidentNotificationRuleAttributesVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentNotificationRuleAttributesVisibility value.
func (v IncidentNotificationRuleAttributesVisibility) Ptr() *IncidentNotificationRuleAttributesVisibility {
	return &v
}
