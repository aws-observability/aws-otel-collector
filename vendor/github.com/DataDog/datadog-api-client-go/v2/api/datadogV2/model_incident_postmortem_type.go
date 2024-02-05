// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPostmortemType Incident postmortem resource type.
type IncidentPostmortemType string

// List of IncidentPostmortemType.
const (
	INCIDENTPOSTMORTEMTYPE_INCIDENT_POSTMORTEMS IncidentPostmortemType = "incident_postmortems"
)

var allowedIncidentPostmortemTypeEnumValues = []IncidentPostmortemType{
	INCIDENTPOSTMORTEMTYPE_INCIDENT_POSTMORTEMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentPostmortemType) GetAllowedValues() []IncidentPostmortemType {
	return allowedIncidentPostmortemTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentPostmortemType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentPostmortemType(value)
	return nil
}

// NewIncidentPostmortemTypeFromValue returns a pointer to a valid IncidentPostmortemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentPostmortemTypeFromValue(v string) (*IncidentPostmortemType, error) {
	ev := IncidentPostmortemType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentPostmortemType: valid values are %v", v, allowedIncidentPostmortemTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentPostmortemType) IsValid() bool {
	for _, existing := range allowedIncidentPostmortemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentPostmortemType value.
func (v IncidentPostmortemType) Ptr() *IncidentPostmortemType {
	return &v
}
