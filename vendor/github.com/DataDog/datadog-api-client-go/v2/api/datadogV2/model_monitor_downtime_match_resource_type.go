// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorDowntimeMatchResourceType Monitor Downtime Match resource type.
type MonitorDowntimeMatchResourceType string

// List of MonitorDowntimeMatchResourceType.
const (
	MONITORDOWNTIMEMATCHRESOURCETYPE_DOWNTIME_MATCH MonitorDowntimeMatchResourceType = "downtime_match"
)

var allowedMonitorDowntimeMatchResourceTypeEnumValues = []MonitorDowntimeMatchResourceType{
	MONITORDOWNTIMEMATCHRESOURCETYPE_DOWNTIME_MATCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorDowntimeMatchResourceType) GetAllowedValues() []MonitorDowntimeMatchResourceType {
	return allowedMonitorDowntimeMatchResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorDowntimeMatchResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorDowntimeMatchResourceType(value)
	return nil
}

// NewMonitorDowntimeMatchResourceTypeFromValue returns a pointer to a valid MonitorDowntimeMatchResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorDowntimeMatchResourceTypeFromValue(v string) (*MonitorDowntimeMatchResourceType, error) {
	ev := MonitorDowntimeMatchResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorDowntimeMatchResourceType: valid values are %v", v, allowedMonitorDowntimeMatchResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorDowntimeMatchResourceType) IsValid() bool {
	for _, existing := range allowedMonitorDowntimeMatchResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorDowntimeMatchResourceType value.
func (v MonitorDowntimeMatchResourceType) Ptr() *MonitorDowntimeMatchResourceType {
	return &v
}
