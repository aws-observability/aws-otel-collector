// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationRuleCreateAttributesVisibility The visibility of the notification rule.
type IncidentNotificationRuleCreateAttributesVisibility string

// List of IncidentNotificationRuleCreateAttributesVisibility.
const (
	INCIDENTNOTIFICATIONRULECREATEATTRIBUTESVISIBILITY_ALL          IncidentNotificationRuleCreateAttributesVisibility = "all"
	INCIDENTNOTIFICATIONRULECREATEATTRIBUTESVISIBILITY_ORGANIZATION IncidentNotificationRuleCreateAttributesVisibility = "organization"
	INCIDENTNOTIFICATIONRULECREATEATTRIBUTESVISIBILITY_PRIVATE      IncidentNotificationRuleCreateAttributesVisibility = "private"
)

var allowedIncidentNotificationRuleCreateAttributesVisibilityEnumValues = []IncidentNotificationRuleCreateAttributesVisibility{
	INCIDENTNOTIFICATIONRULECREATEATTRIBUTESVISIBILITY_ALL,
	INCIDENTNOTIFICATIONRULECREATEATTRIBUTESVISIBILITY_ORGANIZATION,
	INCIDENTNOTIFICATIONRULECREATEATTRIBUTESVISIBILITY_PRIVATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentNotificationRuleCreateAttributesVisibility) GetAllowedValues() []IncidentNotificationRuleCreateAttributesVisibility {
	return allowedIncidentNotificationRuleCreateAttributesVisibilityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentNotificationRuleCreateAttributesVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentNotificationRuleCreateAttributesVisibility(value)
	return nil
}

// NewIncidentNotificationRuleCreateAttributesVisibilityFromValue returns a pointer to a valid IncidentNotificationRuleCreateAttributesVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentNotificationRuleCreateAttributesVisibilityFromValue(v string) (*IncidentNotificationRuleCreateAttributesVisibility, error) {
	ev := IncidentNotificationRuleCreateAttributesVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentNotificationRuleCreateAttributesVisibility: valid values are %v", v, allowedIncidentNotificationRuleCreateAttributesVisibilityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentNotificationRuleCreateAttributesVisibility) IsValid() bool {
	for _, existing := range allowedIncidentNotificationRuleCreateAttributesVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentNotificationRuleCreateAttributesVisibility value.
func (v IncidentNotificationRuleCreateAttributesVisibility) Ptr() *IncidentNotificationRuleCreateAttributesVisibility {
	return &v
}
