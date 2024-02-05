// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeIncludedMonitorType Monitor resource type.
type DowntimeIncludedMonitorType string

// List of DowntimeIncludedMonitorType.
const (
	DOWNTIMEINCLUDEDMONITORTYPE_MONITORS DowntimeIncludedMonitorType = "monitors"
)

var allowedDowntimeIncludedMonitorTypeEnumValues = []DowntimeIncludedMonitorType{
	DOWNTIMEINCLUDEDMONITORTYPE_MONITORS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DowntimeIncludedMonitorType) GetAllowedValues() []DowntimeIncludedMonitorType {
	return allowedDowntimeIncludedMonitorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DowntimeIncludedMonitorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DowntimeIncludedMonitorType(value)
	return nil
}

// NewDowntimeIncludedMonitorTypeFromValue returns a pointer to a valid DowntimeIncludedMonitorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDowntimeIncludedMonitorTypeFromValue(v string) (*DowntimeIncludedMonitorType, error) {
	ev := DowntimeIncludedMonitorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DowntimeIncludedMonitorType: valid values are %v", v, allowedDowntimeIncludedMonitorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DowntimeIncludedMonitorType) IsValid() bool {
	for _, existing := range allowedDowntimeIncludedMonitorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DowntimeIncludedMonitorType value.
func (v DowntimeIncludedMonitorType) Ptr() *DowntimeIncludedMonitorType {
	return &v
}
