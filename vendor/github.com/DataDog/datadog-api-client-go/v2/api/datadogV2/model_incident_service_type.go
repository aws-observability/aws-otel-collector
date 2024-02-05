// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentServiceType Incident service resource type.
type IncidentServiceType string

// List of IncidentServiceType.
const (
	INCIDENTSERVICETYPE_SERVICES IncidentServiceType = "services"
)

var allowedIncidentServiceTypeEnumValues = []IncidentServiceType{
	INCIDENTSERVICETYPE_SERVICES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentServiceType) GetAllowedValues() []IncidentServiceType {
	return allowedIncidentServiceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentServiceType(value)
	return nil
}

// NewIncidentServiceTypeFromValue returns a pointer to a valid IncidentServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentServiceTypeFromValue(v string) (*IncidentServiceType, error) {
	ev := IncidentServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentServiceType: valid values are %v", v, allowedIncidentServiceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentServiceType) IsValid() bool {
	for _, existing := range allowedIncidentServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentServiceType value.
func (v IncidentServiceType) Ptr() *IncidentServiceType {
	return &v
}
