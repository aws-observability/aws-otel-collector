// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTodoType Todo resource type.
type IncidentTodoType string

// List of IncidentTodoType.
const (
	INCIDENTTODOTYPE_INCIDENT_TODOS IncidentTodoType = "incident_todos"
)

var allowedIncidentTodoTypeEnumValues = []IncidentTodoType{
	INCIDENTTODOTYPE_INCIDENT_TODOS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTodoType) GetAllowedValues() []IncidentTodoType {
	return allowedIncidentTodoTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTodoType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTodoType(value)
	return nil
}

// NewIncidentTodoTypeFromValue returns a pointer to a valid IncidentTodoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTodoTypeFromValue(v string) (*IncidentTodoType, error) {
	ev := IncidentTodoType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTodoType: valid values are %v", v, allowedIncidentTodoTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTodoType) IsValid() bool {
	for _, existing := range allowedIncidentTodoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTodoType value.
func (v IncidentTodoType) Ptr() *IncidentTodoType {
	return &v
}
