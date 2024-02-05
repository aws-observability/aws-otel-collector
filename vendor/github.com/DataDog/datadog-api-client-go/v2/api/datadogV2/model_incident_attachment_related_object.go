// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentAttachmentRelatedObject The object related to an incident attachment.
type IncidentAttachmentRelatedObject string

// List of IncidentAttachmentRelatedObject.
const (
	INCIDENTATTACHMENTRELATEDOBJECT_USERS IncidentAttachmentRelatedObject = "users"
)

var allowedIncidentAttachmentRelatedObjectEnumValues = []IncidentAttachmentRelatedObject{
	INCIDENTATTACHMENTRELATEDOBJECT_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentAttachmentRelatedObject) GetAllowedValues() []IncidentAttachmentRelatedObject {
	return allowedIncidentAttachmentRelatedObjectEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentAttachmentRelatedObject) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentAttachmentRelatedObject(value)
	return nil
}

// NewIncidentAttachmentRelatedObjectFromValue returns a pointer to a valid IncidentAttachmentRelatedObject
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentAttachmentRelatedObjectFromValue(v string) (*IncidentAttachmentRelatedObject, error) {
	ev := IncidentAttachmentRelatedObject(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentAttachmentRelatedObject: valid values are %v", v, allowedIncidentAttachmentRelatedObjectEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentAttachmentRelatedObject) IsValid() bool {
	for _, existing := range allowedIncidentAttachmentRelatedObjectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentAttachmentRelatedObject value.
func (v IncidentAttachmentRelatedObject) Ptr() *IncidentAttachmentRelatedObject {
	return &v
}
