// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORAIncidentType JSON:API type for DORA incident events.
type DORAIncidentType string

// List of DORAIncidentType.
const (
	DORAINCIDENTTYPE_DORA_INCIDENT DORAIncidentType = "dora_incident"
)

var allowedDORAIncidentTypeEnumValues = []DORAIncidentType{
	DORAINCIDENTTYPE_DORA_INCIDENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DORAIncidentType) GetAllowedValues() []DORAIncidentType {
	return allowedDORAIncidentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DORAIncidentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DORAIncidentType(value)
	return nil
}

// NewDORAIncidentTypeFromValue returns a pointer to a valid DORAIncidentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDORAIncidentTypeFromValue(v string) (*DORAIncidentType, error) {
	ev := DORAIncidentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DORAIncidentType: valid values are %v", v, allowedDORAIncidentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DORAIncidentType) IsValid() bool {
	for _, existing := range allowedDORAIncidentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DORAIncidentType value.
func (v DORAIncidentType) Ptr() *DORAIncidentType {
	return &v
}
