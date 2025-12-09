// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImpactType Incident impact resource type.
type IncidentImpactType string

// List of IncidentImpactType.
const (
	INCIDENTIMPACTTYPE_INCIDENT_IMPACTS IncidentImpactType = "incident_impacts"
)

var allowedIncidentImpactTypeEnumValues = []IncidentImpactType{
	INCIDENTIMPACTTYPE_INCIDENT_IMPACTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentImpactType) GetAllowedValues() []IncidentImpactType {
	return allowedIncidentImpactTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentImpactType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentImpactType(value)
	return nil
}

// NewIncidentImpactTypeFromValue returns a pointer to a valid IncidentImpactType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentImpactTypeFromValue(v string) (*IncidentImpactType, error) {
	ev := IncidentImpactType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentImpactType: valid values are %v", v, allowedIncidentImpactTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentImpactType) IsValid() bool {
	for _, existing := range allowedIncidentImpactTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentImpactType value.
func (v IncidentImpactType) Ptr() *IncidentImpactType {
	return &v
}
