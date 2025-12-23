// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfiguredScheduleTargetType Indicates that the resource is of type `schedule_target`.
type ConfiguredScheduleTargetType string

// List of ConfiguredScheduleTargetType.
const (
	CONFIGUREDSCHEDULETARGETTYPE_SCHEDULE_TARGET ConfiguredScheduleTargetType = "schedule_target"
)

var allowedConfiguredScheduleTargetTypeEnumValues = []ConfiguredScheduleTargetType{
	CONFIGUREDSCHEDULETARGETTYPE_SCHEDULE_TARGET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ConfiguredScheduleTargetType) GetAllowedValues() []ConfiguredScheduleTargetType {
	return allowedConfiguredScheduleTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ConfiguredScheduleTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ConfiguredScheduleTargetType(value)
	return nil
}

// NewConfiguredScheduleTargetTypeFromValue returns a pointer to a valid ConfiguredScheduleTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewConfiguredScheduleTargetTypeFromValue(v string) (*ConfiguredScheduleTargetType, error) {
	ev := ConfiguredScheduleTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConfiguredScheduleTargetType: valid values are %v", v, allowedConfiguredScheduleTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ConfiguredScheduleTargetType) IsValid() bool {
	for _, existing := range allowedConfiguredScheduleTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConfiguredScheduleTargetType value.
func (v ConfiguredScheduleTargetType) Ptr() *ConfiguredScheduleTargetType {
	return &v
}
