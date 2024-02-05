// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentAttachmentType The incident attachment resource type.
type IncidentAttachmentType string

// List of IncidentAttachmentType.
const (
	INCIDENTATTACHMENTTYPE_INCIDENT_ATTACHMENTS IncidentAttachmentType = "incident_attachments"
)

var allowedIncidentAttachmentTypeEnumValues = []IncidentAttachmentType{
	INCIDENTATTACHMENTTYPE_INCIDENT_ATTACHMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentAttachmentType) GetAllowedValues() []IncidentAttachmentType {
	return allowedIncidentAttachmentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentAttachmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentAttachmentType(value)
	return nil
}

// NewIncidentAttachmentTypeFromValue returns a pointer to a valid IncidentAttachmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentAttachmentTypeFromValue(v string) (*IncidentAttachmentType, error) {
	ev := IncidentAttachmentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentAttachmentType: valid values are %v", v, allowedIncidentAttachmentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentAttachmentType) IsValid() bool {
	for _, existing := range allowedIncidentAttachmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentAttachmentType value.
func (v IncidentAttachmentType) Ptr() *IncidentAttachmentType {
	return &v
}
