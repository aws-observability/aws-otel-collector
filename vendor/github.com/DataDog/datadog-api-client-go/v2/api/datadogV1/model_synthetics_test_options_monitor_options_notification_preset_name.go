// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestOptionsMonitorOptionsNotificationPresetName The name of the preset for the notification for the monitor.
type SyntheticsTestOptionsMonitorOptionsNotificationPresetName string

// List of SyntheticsTestOptionsMonitorOptionsNotificationPresetName.
const (
	SYNTHETICSTESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_SHOW_ALL     SyntheticsTestOptionsMonitorOptionsNotificationPresetName = "show_all"
	SYNTHETICSTESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_ALL     SyntheticsTestOptionsMonitorOptionsNotificationPresetName = "hide_all"
	SYNTHETICSTESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_QUERY   SyntheticsTestOptionsMonitorOptionsNotificationPresetName = "hide_query"
	SYNTHETICSTESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_HANDLES SyntheticsTestOptionsMonitorOptionsNotificationPresetName = "hide_handles"
)

var allowedSyntheticsTestOptionsMonitorOptionsNotificationPresetNameEnumValues = []SyntheticsTestOptionsMonitorOptionsNotificationPresetName{
	SYNTHETICSTESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_SHOW_ALL,
	SYNTHETICSTESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_ALL,
	SYNTHETICSTESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_QUERY,
	SYNTHETICSTESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_HANDLES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestOptionsMonitorOptionsNotificationPresetName) GetAllowedValues() []SyntheticsTestOptionsMonitorOptionsNotificationPresetName {
	return allowedSyntheticsTestOptionsMonitorOptionsNotificationPresetNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestOptionsMonitorOptionsNotificationPresetName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestOptionsMonitorOptionsNotificationPresetName(value)
	return nil
}

// NewSyntheticsTestOptionsMonitorOptionsNotificationPresetNameFromValue returns a pointer to a valid SyntheticsTestOptionsMonitorOptionsNotificationPresetName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestOptionsMonitorOptionsNotificationPresetNameFromValue(v string) (*SyntheticsTestOptionsMonitorOptionsNotificationPresetName, error) {
	ev := SyntheticsTestOptionsMonitorOptionsNotificationPresetName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestOptionsMonitorOptionsNotificationPresetName: valid values are %v", v, allowedSyntheticsTestOptionsMonitorOptionsNotificationPresetNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestOptionsMonitorOptionsNotificationPresetName) IsValid() bool {
	for _, existing := range allowedSyntheticsTestOptionsMonitorOptionsNotificationPresetNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestOptionsMonitorOptionsNotificationPresetName value.
func (v SyntheticsTestOptionsMonitorOptionsNotificationPresetName) Ptr() *SyntheticsTestOptionsMonitorOptionsNotificationPresetName {
	return &v
}
