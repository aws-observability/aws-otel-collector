// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DueDateSeverity A severity level used to configure due date thresholds.
type DueDateSeverity string

// List of DueDateSeverity.
const (
	DUEDATESEVERITY_CRITICAL DueDateSeverity = "critical"
	DUEDATESEVERITY_HIGH     DueDateSeverity = "high"
	DUEDATESEVERITY_MEDIUM   DueDateSeverity = "medium"
	DUEDATESEVERITY_LOW      DueDateSeverity = "low"
	DUEDATESEVERITY_INFO     DueDateSeverity = "info"
	DUEDATESEVERITY_NONE     DueDateSeverity = "none"
	DUEDATESEVERITY_UNKNOWN  DueDateSeverity = "unknown"
)

var allowedDueDateSeverityEnumValues = []DueDateSeverity{
	DUEDATESEVERITY_CRITICAL,
	DUEDATESEVERITY_HIGH,
	DUEDATESEVERITY_MEDIUM,
	DUEDATESEVERITY_LOW,
	DUEDATESEVERITY_INFO,
	DUEDATESEVERITY_NONE,
	DUEDATESEVERITY_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DueDateSeverity) GetAllowedValues() []DueDateSeverity {
	return allowedDueDateSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DueDateSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DueDateSeverity(value)
	return nil
}

// NewDueDateSeverityFromValue returns a pointer to a valid DueDateSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDueDateSeverityFromValue(v string) (*DueDateSeverity, error) {
	ev := DueDateSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DueDateSeverity: valid values are %v", v, allowedDueDateSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DueDateSeverity) IsValid() bool {
	for _, existing := range allowedDueDateSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DueDateSeverity value.
func (v DueDateSeverity) Ptr() *DueDateSeverity {
	return &v
}
