// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName The definition of `SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName` object.
type SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName string

// List of SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName.
const (
	SYNTHETICSMOBILETESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_SHOW_ALL     SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName = "show_all"
	SYNTHETICSMOBILETESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_ALL     SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName = "hide_all"
	SYNTHETICSMOBILETESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_QUERY   SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName = "hide_query"
	SYNTHETICSMOBILETESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_HANDLES SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName = "hide_handles"
)

var allowedSyntheticsMobileTestOptionsMonitorOptionsNotificationPresetNameEnumValues = []SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName{
	SYNTHETICSMOBILETESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_SHOW_ALL,
	SYNTHETICSMOBILETESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_ALL,
	SYNTHETICSMOBILETESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_QUERY,
	SYNTHETICSMOBILETESTOPTIONSMONITOROPTIONSNOTIFICATIONPRESETNAME_HIDE_HANDLES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName) GetAllowedValues() []SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName {
	return allowedSyntheticsMobileTestOptionsMonitorOptionsNotificationPresetNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName(value)
	return nil
}

// NewSyntheticsMobileTestOptionsMonitorOptionsNotificationPresetNameFromValue returns a pointer to a valid SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsMobileTestOptionsMonitorOptionsNotificationPresetNameFromValue(v string) (*SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName, error) {
	ev := SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName: valid values are %v", v, allowedSyntheticsMobileTestOptionsMonitorOptionsNotificationPresetNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName) IsValid() bool {
	for _, existing := range allowedSyntheticsMobileTestOptionsMonitorOptionsNotificationPresetNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName value.
func (v SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName) Ptr() *SyntheticsMobileTestOptionsMonitorOptionsNotificationPresetName {
	return &v
}
