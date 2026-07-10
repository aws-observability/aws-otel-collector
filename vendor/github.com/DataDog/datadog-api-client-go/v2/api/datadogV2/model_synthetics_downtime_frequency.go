// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsDowntimeFrequency The recurrence frequency of a Synthetics downtime time slot.
type SyntheticsDowntimeFrequency string

// List of SyntheticsDowntimeFrequency.
const (
	SYNTHETICSDOWNTIMEFREQUENCY_DAILY   SyntheticsDowntimeFrequency = "DAILY"
	SYNTHETICSDOWNTIMEFREQUENCY_WEEKLY  SyntheticsDowntimeFrequency = "WEEKLY"
	SYNTHETICSDOWNTIMEFREQUENCY_MONTHLY SyntheticsDowntimeFrequency = "MONTHLY"
	SYNTHETICSDOWNTIMEFREQUENCY_YEARLY  SyntheticsDowntimeFrequency = "YEARLY"
)

var allowedSyntheticsDowntimeFrequencyEnumValues = []SyntheticsDowntimeFrequency{
	SYNTHETICSDOWNTIMEFREQUENCY_DAILY,
	SYNTHETICSDOWNTIMEFREQUENCY_WEEKLY,
	SYNTHETICSDOWNTIMEFREQUENCY_MONTHLY,
	SYNTHETICSDOWNTIMEFREQUENCY_YEARLY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsDowntimeFrequency) GetAllowedValues() []SyntheticsDowntimeFrequency {
	return allowedSyntheticsDowntimeFrequencyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsDowntimeFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsDowntimeFrequency(value)
	return nil
}

// NewSyntheticsDowntimeFrequencyFromValue returns a pointer to a valid SyntheticsDowntimeFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsDowntimeFrequencyFromValue(v string) (*SyntheticsDowntimeFrequency, error) {
	ev := SyntheticsDowntimeFrequency(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsDowntimeFrequency: valid values are %v", v, allowedSyntheticsDowntimeFrequencyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsDowntimeFrequency) IsValid() bool {
	for _, existing := range allowedSyntheticsDowntimeFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsDowntimeFrequency value.
func (v SyntheticsDowntimeFrequency) Ptr() *SyntheticsDowntimeFrequency {
	return &v
}
