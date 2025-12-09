// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimelineCellType Timeline cell content type
type TimelineCellType string

// List of TimelineCellType.
const (
	TIMELINECELLTYPE_COMMENT TimelineCellType = "COMMENT"
)

var allowedTimelineCellTypeEnumValues = []TimelineCellType{
	TIMELINECELLTYPE_COMMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimelineCellType) GetAllowedValues() []TimelineCellType {
	return allowedTimelineCellTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimelineCellType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimelineCellType(value)
	return nil
}

// NewTimelineCellTypeFromValue returns a pointer to a valid TimelineCellType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimelineCellTypeFromValue(v string) (*TimelineCellType, error) {
	ev := TimelineCellType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimelineCellType: valid values are %v", v, allowedTimelineCellTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimelineCellType) IsValid() bool {
	for _, existing := range allowedTimelineCellTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimelineCellType value.
func (v TimelineCellType) Ptr() *TimelineCellType {
	return &v
}
