// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventCategory Event category identifying the type of event.
type EventCategory string

// List of EventCategory.
const (
	EVENTCATEGORY_CHANGE EventCategory = "change"
	EVENTCATEGORY_ALERT  EventCategory = "alert"
)

var allowedEventCategoryEnumValues = []EventCategory{
	EVENTCATEGORY_CHANGE,
	EVENTCATEGORY_ALERT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventCategory) GetAllowedValues() []EventCategory {
	return allowedEventCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventCategory(value)
	return nil
}

// NewEventCategoryFromValue returns a pointer to a valid EventCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventCategoryFromValue(v string) (*EventCategory, error) {
	ev := EventCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventCategory: valid values are %v", v, allowedEventCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventCategory) IsValid() bool {
	for _, existing := range allowedEventCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventCategory value.
func (v EventCategory) Ptr() *EventCategory {
	return &v
}
