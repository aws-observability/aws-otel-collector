// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTypeType Incident type resource type.
type IncidentTypeType string

// List of IncidentTypeType.
const (
	INCIDENTTYPETYPE_INCIDENT_TYPES IncidentTypeType = "incident_types"
)

var allowedIncidentTypeTypeEnumValues = []IncidentTypeType{
	INCIDENTTYPETYPE_INCIDENT_TYPES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTypeType) GetAllowedValues() []IncidentTypeType {
	return allowedIncidentTypeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTypeType(value)
	return nil
}

// NewIncidentTypeTypeFromValue returns a pointer to a valid IncidentTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTypeTypeFromValue(v string) (*IncidentTypeType, error) {
	ev := IncidentTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTypeType: valid values are %v", v, allowedIncidentTypeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTypeType) IsValid() bool {
	for _, existing := range allowedIncidentTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTypeType value.
func (v IncidentTypeType) Ptr() *IncidentTypeType {
	return &v
}
