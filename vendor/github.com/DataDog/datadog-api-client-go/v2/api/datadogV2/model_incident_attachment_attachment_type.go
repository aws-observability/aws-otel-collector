// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentAttachmentAttachmentType The type of the incident attachment attributes.
type IncidentAttachmentAttachmentType string

// List of IncidentAttachmentAttachmentType.
const (
	INCIDENTATTACHMENTATTACHMENTTYPE_LINK       IncidentAttachmentAttachmentType = "link"
	INCIDENTATTACHMENTATTACHMENTTYPE_POSTMORTEM IncidentAttachmentAttachmentType = "postmortem"
)

var allowedIncidentAttachmentAttachmentTypeEnumValues = []IncidentAttachmentAttachmentType{
	INCIDENTATTACHMENTATTACHMENTTYPE_LINK,
	INCIDENTATTACHMENTATTACHMENTTYPE_POSTMORTEM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentAttachmentAttachmentType) GetAllowedValues() []IncidentAttachmentAttachmentType {
	return allowedIncidentAttachmentAttachmentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentAttachmentAttachmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentAttachmentAttachmentType(value)
	return nil
}

// NewIncidentAttachmentAttachmentTypeFromValue returns a pointer to a valid IncidentAttachmentAttachmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentAttachmentAttachmentTypeFromValue(v string) (*IncidentAttachmentAttachmentType, error) {
	ev := IncidentAttachmentAttachmentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentAttachmentAttachmentType: valid values are %v", v, allowedIncidentAttachmentAttachmentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentAttachmentAttachmentType) IsValid() bool {
	for _, existing := range allowedIncidentAttachmentAttachmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentAttachmentAttachmentType value.
func (v IncidentAttachmentAttachmentType) Ptr() *IncidentAttachmentAttachmentType {
	return &v
}
