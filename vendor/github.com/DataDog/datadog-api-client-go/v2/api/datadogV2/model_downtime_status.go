// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeStatus The current status of the downtime.
type DowntimeStatus string

// List of DowntimeStatus.
const (
	DOWNTIMESTATUS_ACTIVE    DowntimeStatus = "active"
	DOWNTIMESTATUS_CANCELED  DowntimeStatus = "canceled"
	DOWNTIMESTATUS_ENDED     DowntimeStatus = "ended"
	DOWNTIMESTATUS_SCHEDULED DowntimeStatus = "scheduled"
)

var allowedDowntimeStatusEnumValues = []DowntimeStatus{
	DOWNTIMESTATUS_ACTIVE,
	DOWNTIMESTATUS_CANCELED,
	DOWNTIMESTATUS_ENDED,
	DOWNTIMESTATUS_SCHEDULED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DowntimeStatus) GetAllowedValues() []DowntimeStatus {
	return allowedDowntimeStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DowntimeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DowntimeStatus(value)
	return nil
}

// NewDowntimeStatusFromValue returns a pointer to a valid DowntimeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDowntimeStatusFromValue(v string) (*DowntimeStatus, error) {
	ev := DowntimeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DowntimeStatus: valid values are %v", v, allowedDowntimeStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DowntimeStatus) IsValid() bool {
	for _, existing := range allowedDowntimeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DowntimeStatus value.
func (v DowntimeStatus) Ptr() *DowntimeStatus {
	return &v
}
