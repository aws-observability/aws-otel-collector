// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventSystemAttributesCategory Event category identifying the type of event.
type EventSystemAttributesCategory string

// List of EventSystemAttributesCategory.
const (
	EVENTSYSTEMATTRIBUTESCATEGORY_CHANGE EventSystemAttributesCategory = "change"
	EVENTSYSTEMATTRIBUTESCATEGORY_ALERT  EventSystemAttributesCategory = "alert"
)

var allowedEventSystemAttributesCategoryEnumValues = []EventSystemAttributesCategory{
	EVENTSYSTEMATTRIBUTESCATEGORY_CHANGE,
	EVENTSYSTEMATTRIBUTESCATEGORY_ALERT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventSystemAttributesCategory) GetAllowedValues() []EventSystemAttributesCategory {
	return allowedEventSystemAttributesCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventSystemAttributesCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventSystemAttributesCategory(value)
	return nil
}

// NewEventSystemAttributesCategoryFromValue returns a pointer to a valid EventSystemAttributesCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventSystemAttributesCategoryFromValue(v string) (*EventSystemAttributesCategory, error) {
	ev := EventSystemAttributesCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventSystemAttributesCategory: valid values are %v", v, allowedEventSystemAttributesCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventSystemAttributesCategory) IsValid() bool {
	for _, existing := range allowedEventSystemAttributesCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventSystemAttributesCategory value.
func (v EventSystemAttributesCategory) Ptr() *EventSystemAttributesCategory {
	return &v
}
