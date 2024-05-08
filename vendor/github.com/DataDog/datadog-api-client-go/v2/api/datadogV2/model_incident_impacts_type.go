// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImpactsType The incident impacts type.
type IncidentImpactsType string

// List of IncidentImpactsType.
const (
	INCIDENTIMPACTSTYPE_INCIDENT_IMPACTS IncidentImpactsType = "incident_impacts"
)

var allowedIncidentImpactsTypeEnumValues = []IncidentImpactsType{
	INCIDENTIMPACTSTYPE_INCIDENT_IMPACTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentImpactsType) GetAllowedValues() []IncidentImpactsType {
	return allowedIncidentImpactsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentImpactsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentImpactsType(value)
	return nil
}

// NewIncidentImpactsTypeFromValue returns a pointer to a valid IncidentImpactsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentImpactsTypeFromValue(v string) (*IncidentImpactsType, error) {
	ev := IncidentImpactsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentImpactsType: valid values are %v", v, allowedIncidentImpactsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentImpactsType) IsValid() bool {
	for _, existing := range allowedIncidentImpactsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentImpactsType value.
func (v IncidentImpactsType) Ptr() *IncidentImpactsType {
	return &v
}
