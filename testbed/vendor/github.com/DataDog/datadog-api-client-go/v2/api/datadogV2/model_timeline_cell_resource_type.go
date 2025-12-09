// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimelineCellResourceType Timeline cell JSON:API resource type
type TimelineCellResourceType string

// List of TimelineCellResourceType.
const (
	TIMELINECELLRESOURCETYPE_TIMELINE_CELL TimelineCellResourceType = "timeline_cell"
)

var allowedTimelineCellResourceTypeEnumValues = []TimelineCellResourceType{
	TIMELINECELLRESOURCETYPE_TIMELINE_CELL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimelineCellResourceType) GetAllowedValues() []TimelineCellResourceType {
	return allowedTimelineCellResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimelineCellResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimelineCellResourceType(value)
	return nil
}

// NewTimelineCellResourceTypeFromValue returns a pointer to a valid TimelineCellResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimelineCellResourceTypeFromValue(v string) (*TimelineCellResourceType, error) {
	ev := TimelineCellResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimelineCellResourceType: valid values are %v", v, allowedTimelineCellResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimelineCellResourceType) IsValid() bool {
	for _, existing := range allowedTimelineCellResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimelineCellResourceType value.
func (v TimelineCellResourceType) Ptr() *TimelineCellResourceType {
	return &v
}
