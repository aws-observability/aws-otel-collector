// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImportRelatedObject Object related to an incident that can be included in the response.
type IncidentImportRelatedObject string

// List of IncidentImportRelatedObject.
const (
	INCIDENTIMPORTRELATEDOBJECT_LAST_MODIFIED_BY_USER IncidentImportRelatedObject = "last_modified_by_user"
	INCIDENTIMPORTRELATEDOBJECT_CREATED_BY_USER       IncidentImportRelatedObject = "created_by_user"
	INCIDENTIMPORTRELATEDOBJECT_COMMANDER_USER        IncidentImportRelatedObject = "commander_user"
	INCIDENTIMPORTRELATEDOBJECT_DECLARED_BY_USER      IncidentImportRelatedObject = "declared_by_user"
	INCIDENTIMPORTRELATEDOBJECT_INCIDENT_TYPE         IncidentImportRelatedObject = "incident_type"
)

var allowedIncidentImportRelatedObjectEnumValues = []IncidentImportRelatedObject{
	INCIDENTIMPORTRELATEDOBJECT_LAST_MODIFIED_BY_USER,
	INCIDENTIMPORTRELATEDOBJECT_CREATED_BY_USER,
	INCIDENTIMPORTRELATEDOBJECT_COMMANDER_USER,
	INCIDENTIMPORTRELATEDOBJECT_DECLARED_BY_USER,
	INCIDENTIMPORTRELATEDOBJECT_INCIDENT_TYPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentImportRelatedObject) GetAllowedValues() []IncidentImportRelatedObject {
	return allowedIncidentImportRelatedObjectEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentImportRelatedObject) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentImportRelatedObject(value)
	return nil
}

// NewIncidentImportRelatedObjectFromValue returns a pointer to a valid IncidentImportRelatedObject
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentImportRelatedObjectFromValue(v string) (*IncidentImportRelatedObject, error) {
	ev := IncidentImportRelatedObject(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentImportRelatedObject: valid values are %v", v, allowedIncidentImportRelatedObjectEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentImportRelatedObject) IsValid() bool {
	for _, existing := range allowedIncidentImportRelatedObjectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentImportRelatedObject value.
func (v IncidentImportRelatedObject) Ptr() *IncidentImportRelatedObject {
	return &v
}
