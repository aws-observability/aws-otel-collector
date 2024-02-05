// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsSortType The type of sort to use on the calculated value.
type EventsSortType string

// List of EventsSortType.
const (
	EVENTSSORTTYPE_ALPHABETICAL EventsSortType = "alphabetical"
	EVENTSSORTTYPE_MEASURE      EventsSortType = "measure"
)

var allowedEventsSortTypeEnumValues = []EventsSortType{
	EVENTSSORTTYPE_ALPHABETICAL,
	EVENTSSORTTYPE_MEASURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventsSortType) GetAllowedValues() []EventsSortType {
	return allowedEventsSortTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventsSortType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventsSortType(value)
	return nil
}

// NewEventsSortTypeFromValue returns a pointer to a valid EventsSortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventsSortTypeFromValue(v string) (*EventsSortType, error) {
	ev := EventsSortType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventsSortType: valid values are %v", v, allowedEventsSortTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventsSortType) IsValid() bool {
	for _, existing := range allowedEventsSortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventsSortType value.
func (v EventsSortType) Ptr() *EventsSortType {
	return &v
}
