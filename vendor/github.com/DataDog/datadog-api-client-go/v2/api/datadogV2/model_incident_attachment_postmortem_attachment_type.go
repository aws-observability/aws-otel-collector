// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentAttachmentPostmortemAttachmentType The type of postmortem attachment attributes.
type IncidentAttachmentPostmortemAttachmentType string

// List of IncidentAttachmentPostmortemAttachmentType.
const (
	INCIDENTATTACHMENTPOSTMORTEMATTACHMENTTYPE_POSTMORTEM IncidentAttachmentPostmortemAttachmentType = "postmortem"
)

var allowedIncidentAttachmentPostmortemAttachmentTypeEnumValues = []IncidentAttachmentPostmortemAttachmentType{
	INCIDENTATTACHMENTPOSTMORTEMATTACHMENTTYPE_POSTMORTEM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentAttachmentPostmortemAttachmentType) GetAllowedValues() []IncidentAttachmentPostmortemAttachmentType {
	return allowedIncidentAttachmentPostmortemAttachmentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentAttachmentPostmortemAttachmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentAttachmentPostmortemAttachmentType(value)
	return nil
}

// NewIncidentAttachmentPostmortemAttachmentTypeFromValue returns a pointer to a valid IncidentAttachmentPostmortemAttachmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentAttachmentPostmortemAttachmentTypeFromValue(v string) (*IncidentAttachmentPostmortemAttachmentType, error) {
	ev := IncidentAttachmentPostmortemAttachmentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentAttachmentPostmortemAttachmentType: valid values are %v", v, allowedIncidentAttachmentPostmortemAttachmentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentAttachmentPostmortemAttachmentType) IsValid() bool {
	for _, existing := range allowedIncidentAttachmentPostmortemAttachmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentAttachmentPostmortemAttachmentType value.
func (v IncidentAttachmentPostmortemAttachmentType) Ptr() *IncidentAttachmentPostmortemAttachmentType {
	return &v
}
