// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentAttachmentLinkAttachmentType The type of link attachment attributes.
type IncidentAttachmentLinkAttachmentType string

// List of IncidentAttachmentLinkAttachmentType.
const (
	INCIDENTATTACHMENTLINKATTACHMENTTYPE_LINK IncidentAttachmentLinkAttachmentType = "link"
)

var allowedIncidentAttachmentLinkAttachmentTypeEnumValues = []IncidentAttachmentLinkAttachmentType{
	INCIDENTATTACHMENTLINKATTACHMENTTYPE_LINK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentAttachmentLinkAttachmentType) GetAllowedValues() []IncidentAttachmentLinkAttachmentType {
	return allowedIncidentAttachmentLinkAttachmentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentAttachmentLinkAttachmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentAttachmentLinkAttachmentType(value)
	return nil
}

// NewIncidentAttachmentLinkAttachmentTypeFromValue returns a pointer to a valid IncidentAttachmentLinkAttachmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentAttachmentLinkAttachmentTypeFromValue(v string) (*IncidentAttachmentLinkAttachmentType, error) {
	ev := IncidentAttachmentLinkAttachmentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentAttachmentLinkAttachmentType: valid values are %v", v, allowedIncidentAttachmentLinkAttachmentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentAttachmentLinkAttachmentType) IsValid() bool {
	for _, existing := range allowedIncidentAttachmentLinkAttachmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentAttachmentLinkAttachmentType value.
func (v IncidentAttachmentLinkAttachmentType) Ptr() *IncidentAttachmentLinkAttachmentType {
	return &v
}
