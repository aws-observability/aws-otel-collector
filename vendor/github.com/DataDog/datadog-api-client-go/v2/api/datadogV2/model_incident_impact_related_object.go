// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImpactRelatedObject A reference to a resource related to an incident impact.
type IncidentImpactRelatedObject string

// List of IncidentImpactRelatedObject.
const (
	INCIDENTIMPACTRELATEDOBJECT_INCIDENT              IncidentImpactRelatedObject = "incident"
	INCIDENTIMPACTRELATEDOBJECT_CREATED_BY_USER       IncidentImpactRelatedObject = "created_by_user"
	INCIDENTIMPACTRELATEDOBJECT_LAST_MODIFIED_BY_USER IncidentImpactRelatedObject = "last_modified_by_user"
)

var allowedIncidentImpactRelatedObjectEnumValues = []IncidentImpactRelatedObject{
	INCIDENTIMPACTRELATEDOBJECT_INCIDENT,
	INCIDENTIMPACTRELATEDOBJECT_CREATED_BY_USER,
	INCIDENTIMPACTRELATEDOBJECT_LAST_MODIFIED_BY_USER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentImpactRelatedObject) GetAllowedValues() []IncidentImpactRelatedObject {
	return allowedIncidentImpactRelatedObjectEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentImpactRelatedObject) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentImpactRelatedObject(value)
	return nil
}

// NewIncidentImpactRelatedObjectFromValue returns a pointer to a valid IncidentImpactRelatedObject
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentImpactRelatedObjectFromValue(v string) (*IncidentImpactRelatedObject, error) {
	ev := IncidentImpactRelatedObject(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentImpactRelatedObject: valid values are %v", v, allowedIncidentImpactRelatedObjectEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentImpactRelatedObject) IsValid() bool {
	for _, existing := range allowedIncidentImpactRelatedObjectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentImpactRelatedObject value.
func (v IncidentImpactRelatedObject) Ptr() *IncidentImpactRelatedObject {
	return &v
}
