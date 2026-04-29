// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalIncidentSettingsType Global incident settings resource type
type GlobalIncidentSettingsType string

// List of GlobalIncidentSettingsType.
const (
	GLOBALINCIDENTSETTINGSTYPE_INCIDENTS_GLOBAL_SETTINGS GlobalIncidentSettingsType = "incidents_global_settings"
)

var allowedGlobalIncidentSettingsTypeEnumValues = []GlobalIncidentSettingsType{
	GLOBALINCIDENTSETTINGSTYPE_INCIDENTS_GLOBAL_SETTINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GlobalIncidentSettingsType) GetAllowedValues() []GlobalIncidentSettingsType {
	return allowedGlobalIncidentSettingsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GlobalIncidentSettingsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GlobalIncidentSettingsType(value)
	return nil
}

// NewGlobalIncidentSettingsTypeFromValue returns a pointer to a valid GlobalIncidentSettingsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGlobalIncidentSettingsTypeFromValue(v string) (*GlobalIncidentSettingsType, error) {
	ev := GlobalIncidentSettingsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GlobalIncidentSettingsType: valid values are %v", v, allowedGlobalIncidentSettingsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GlobalIncidentSettingsType) IsValid() bool {
	for _, existing := range allowedGlobalIncidentSettingsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GlobalIncidentSettingsType value.
func (v GlobalIncidentSettingsType) Ptr() *GlobalIncidentSettingsType {
	return &v
}
