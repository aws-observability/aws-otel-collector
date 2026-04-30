// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImportVisibility The visibility of the incident.
type IncidentImportVisibility string

// List of IncidentImportVisibility.
const (
	INCIDENTIMPORTVISIBILITY_ORGANIZATION IncidentImportVisibility = "organization"
	INCIDENTIMPORTVISIBILITY_PRIVATE      IncidentImportVisibility = "private"
)

var allowedIncidentImportVisibilityEnumValues = []IncidentImportVisibility{
	INCIDENTIMPORTVISIBILITY_ORGANIZATION,
	INCIDENTIMPORTVISIBILITY_PRIVATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentImportVisibility) GetAllowedValues() []IncidentImportVisibility {
	return allowedIncidentImportVisibilityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentImportVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentImportVisibility(value)
	return nil
}

// NewIncidentImportVisibilityFromValue returns a pointer to a valid IncidentImportVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentImportVisibilityFromValue(v string) (*IncidentImportVisibility, error) {
	ev := IncidentImportVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentImportVisibility: valid values are %v", v, allowedIncidentImportVisibilityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentImportVisibility) IsValid() bool {
	for _, existing := range allowedIncidentImportVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentImportVisibility value.
func (v IncidentImportVisibility) Ptr() *IncidentImportVisibility {
	return &v
}
