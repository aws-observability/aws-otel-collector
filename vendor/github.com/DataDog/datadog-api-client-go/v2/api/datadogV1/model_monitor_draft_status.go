// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorDraftStatus Indicates whether the monitor is in a draft or published state.
//
// `draft`: The monitor appears as Draft and does not send notifications.
// `published`: The monitor is active and evaluates conditions and notify as configured.
//
// This field is in preview. The draft value is only available to customers with the feature enabled.
type MonitorDraftStatus string

// List of MonitorDraftStatus.
const (
	MONITORDRAFTSTATUS_DRAFT     MonitorDraftStatus = "draft"
	MONITORDRAFTSTATUS_PUBLISHED MonitorDraftStatus = "published"
)

var allowedMonitorDraftStatusEnumValues = []MonitorDraftStatus{
	MONITORDRAFTSTATUS_DRAFT,
	MONITORDRAFTSTATUS_PUBLISHED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorDraftStatus) GetAllowedValues() []MonitorDraftStatus {
	return allowedMonitorDraftStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorDraftStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorDraftStatus(value)
	return nil
}

// NewMonitorDraftStatusFromValue returns a pointer to a valid MonitorDraftStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorDraftStatusFromValue(v string) (*MonitorDraftStatus, error) {
	ev := MonitorDraftStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorDraftStatus: valid values are %v", v, allowedMonitorDraftStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorDraftStatus) IsValid() bool {
	for _, existing := range allowedMonitorDraftStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorDraftStatus value.
func (v MonitorDraftStatus) Ptr() *MonitorDraftStatus {
	return &v
}
