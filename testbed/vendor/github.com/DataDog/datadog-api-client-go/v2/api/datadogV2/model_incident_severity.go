// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSeverity The incident severity.
type IncidentSeverity string

// List of IncidentSeverity.
const (
	INCIDENTSEVERITY_UNKNOWN IncidentSeverity = "UNKNOWN"
	INCIDENTSEVERITY_SEV_0   IncidentSeverity = "SEV-0"
	INCIDENTSEVERITY_SEV_1   IncidentSeverity = "SEV-1"
	INCIDENTSEVERITY_SEV_2   IncidentSeverity = "SEV-2"
	INCIDENTSEVERITY_SEV_3   IncidentSeverity = "SEV-3"
	INCIDENTSEVERITY_SEV_4   IncidentSeverity = "SEV-4"
	INCIDENTSEVERITY_SEV_5   IncidentSeverity = "SEV-5"
)

var allowedIncidentSeverityEnumValues = []IncidentSeverity{
	INCIDENTSEVERITY_UNKNOWN,
	INCIDENTSEVERITY_SEV_0,
	INCIDENTSEVERITY_SEV_1,
	INCIDENTSEVERITY_SEV_2,
	INCIDENTSEVERITY_SEV_3,
	INCIDENTSEVERITY_SEV_4,
	INCIDENTSEVERITY_SEV_5,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentSeverity) GetAllowedValues() []IncidentSeverity {
	return allowedIncidentSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentSeverity(value)
	return nil
}

// NewIncidentSeverityFromValue returns a pointer to a valid IncidentSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentSeverityFromValue(v string) (*IncidentSeverity, error) {
	ev := IncidentSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentSeverity: valid values are %v", v, allowedIncidentSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentSeverity) IsValid() bool {
	for _, existing := range allowedIncidentSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentSeverity value.
func (v IncidentSeverity) Ptr() *IncidentSeverity {
	return &v
}
