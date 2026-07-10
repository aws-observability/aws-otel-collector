// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DueDateFrom The reference point from which the due date is calculated. When `fix_available` is selected but not applicable to the finding type, `first_seen` is used instead.
type DueDateFrom string

// List of DueDateFrom.
const (
	DUEDATEFROM_FIRST_SEEN    DueDateFrom = "first_seen"
	DUEDATEFROM_FIX_AVAILABLE DueDateFrom = "fix_available"
)

var allowedDueDateFromEnumValues = []DueDateFrom{
	DUEDATEFROM_FIRST_SEEN,
	DUEDATEFROM_FIX_AVAILABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DueDateFrom) GetAllowedValues() []DueDateFrom {
	return allowedDueDateFromEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DueDateFrom) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DueDateFrom(value)
	return nil
}

// NewDueDateFromFromValue returns a pointer to a valid DueDateFrom
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDueDateFromFromValue(v string) (*DueDateFrom, error) {
	ev := DueDateFrom(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DueDateFrom: valid values are %v", v, allowedDueDateFromEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DueDateFrom) IsValid() bool {
	for _, existing := range allowedDueDateFromEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DueDateFrom value.
func (v DueDateFrom) Ptr() *DueDateFrom {
	return &v
}
