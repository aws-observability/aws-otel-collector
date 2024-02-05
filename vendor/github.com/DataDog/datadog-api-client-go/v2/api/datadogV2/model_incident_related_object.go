// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRelatedObject Object related to an incident.
type IncidentRelatedObject string

// List of IncidentRelatedObject.
const (
	INCIDENTRELATEDOBJECT_USERS       IncidentRelatedObject = "users"
	INCIDENTRELATEDOBJECT_ATTACHMENTS IncidentRelatedObject = "attachments"
)

var allowedIncidentRelatedObjectEnumValues = []IncidentRelatedObject{
	INCIDENTRELATEDOBJECT_USERS,
	INCIDENTRELATEDOBJECT_ATTACHMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentRelatedObject) GetAllowedValues() []IncidentRelatedObject {
	return allowedIncidentRelatedObjectEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentRelatedObject) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentRelatedObject(value)
	return nil
}

// NewIncidentRelatedObjectFromValue returns a pointer to a valid IncidentRelatedObject
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentRelatedObjectFromValue(v string) (*IncidentRelatedObject, error) {
	ev := IncidentRelatedObject(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentRelatedObject: valid values are %v", v, allowedIncidentRelatedObjectEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentRelatedObject) IsValid() bool {
	for _, existing := range allowedIncidentRelatedObjectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentRelatedObject value.
func (v IncidentRelatedObject) Ptr() *IncidentRelatedObject {
	return &v
}
