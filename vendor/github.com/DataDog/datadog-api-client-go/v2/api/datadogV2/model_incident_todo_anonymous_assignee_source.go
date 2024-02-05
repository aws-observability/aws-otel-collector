// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTodoAnonymousAssigneeSource The source of the anonymous assignee.
type IncidentTodoAnonymousAssigneeSource string

// List of IncidentTodoAnonymousAssigneeSource.
const (
	INCIDENTTODOANONYMOUSASSIGNEESOURCE_SLACK           IncidentTodoAnonymousAssigneeSource = "slack"
	INCIDENTTODOANONYMOUSASSIGNEESOURCE_MICROSOFT_TEAMS IncidentTodoAnonymousAssigneeSource = "microsoft_teams"
)

var allowedIncidentTodoAnonymousAssigneeSourceEnumValues = []IncidentTodoAnonymousAssigneeSource{
	INCIDENTTODOANONYMOUSASSIGNEESOURCE_SLACK,
	INCIDENTTODOANONYMOUSASSIGNEESOURCE_MICROSOFT_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTodoAnonymousAssigneeSource) GetAllowedValues() []IncidentTodoAnonymousAssigneeSource {
	return allowedIncidentTodoAnonymousAssigneeSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTodoAnonymousAssigneeSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTodoAnonymousAssigneeSource(value)
	return nil
}

// NewIncidentTodoAnonymousAssigneeSourceFromValue returns a pointer to a valid IncidentTodoAnonymousAssigneeSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTodoAnonymousAssigneeSourceFromValue(v string) (*IncidentTodoAnonymousAssigneeSource, error) {
	ev := IncidentTodoAnonymousAssigneeSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTodoAnonymousAssigneeSource: valid values are %v", v, allowedIncidentTodoAnonymousAssigneeSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTodoAnonymousAssigneeSource) IsValid() bool {
	for _, existing := range allowedIncidentTodoAnonymousAssigneeSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTodoAnonymousAssigneeSource value.
func (v IncidentTodoAnonymousAssigneeSource) Ptr() *IncidentTodoAnonymousAssigneeSource {
	return &v
}
