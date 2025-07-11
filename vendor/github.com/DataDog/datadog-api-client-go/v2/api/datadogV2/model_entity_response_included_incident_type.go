// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedIncidentType Incident description.
type EntityResponseIncludedIncidentType string

// List of EntityResponseIncludedIncidentType.
const (
	ENTITYRESPONSEINCLUDEDINCIDENTTYPE_INCIDENT EntityResponseIncludedIncidentType = "incident"
)

var allowedEntityResponseIncludedIncidentTypeEnumValues = []EntityResponseIncludedIncidentType{
	ENTITYRESPONSEINCLUDEDINCIDENTTYPE_INCIDENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityResponseIncludedIncidentType) GetAllowedValues() []EntityResponseIncludedIncidentType {
	return allowedEntityResponseIncludedIncidentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityResponseIncludedIncidentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityResponseIncludedIncidentType(value)
	return nil
}

// NewEntityResponseIncludedIncidentTypeFromValue returns a pointer to a valid EntityResponseIncludedIncidentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityResponseIncludedIncidentTypeFromValue(v string) (*EntityResponseIncludedIncidentType, error) {
	ev := EntityResponseIncludedIncidentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityResponseIncludedIncidentType: valid values are %v", v, allowedEntityResponseIncludedIncidentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityResponseIncludedIncidentType) IsValid() bool {
	for _, existing := range allowedEntityResponseIncludedIncidentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityResponseIncludedIncidentType value.
func (v EntityResponseIncludedIncidentType) Ptr() *EntityResponseIncludedIncidentType {
	return &v
}
