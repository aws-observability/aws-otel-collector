// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationTemplateType Notification templates resource type.
type IncidentNotificationTemplateType string

// List of IncidentNotificationTemplateType.
const (
	INCIDENTNOTIFICATIONTEMPLATETYPE_NOTIFICATION_TEMPLATES IncidentNotificationTemplateType = "notification_templates"
)

var allowedIncidentNotificationTemplateTypeEnumValues = []IncidentNotificationTemplateType{
	INCIDENTNOTIFICATIONTEMPLATETYPE_NOTIFICATION_TEMPLATES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentNotificationTemplateType) GetAllowedValues() []IncidentNotificationTemplateType {
	return allowedIncidentNotificationTemplateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentNotificationTemplateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentNotificationTemplateType(value)
	return nil
}

// NewIncidentNotificationTemplateTypeFromValue returns a pointer to a valid IncidentNotificationTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentNotificationTemplateTypeFromValue(v string) (*IncidentNotificationTemplateType, error) {
	ev := IncidentNotificationTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentNotificationTemplateType: valid values are %v", v, allowedIncidentNotificationTemplateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentNotificationTemplateType) IsValid() bool {
	for _, existing := range allowedIncidentNotificationTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentNotificationTemplateType value.
func (v IncidentNotificationTemplateType) Ptr() *IncidentNotificationTemplateType {
	return &v
}
