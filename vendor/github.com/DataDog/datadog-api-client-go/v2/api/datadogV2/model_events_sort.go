// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsSort The sort parameters when querying events.
type EventsSort string

// List of EventsSort.
const (
	EVENTSSORT_TIMESTAMP_ASCENDING  EventsSort = "timestamp"
	EVENTSSORT_TIMESTAMP_DESCENDING EventsSort = "-timestamp"
)

var allowedEventsSortEnumValues = []EventsSort{
	EVENTSSORT_TIMESTAMP_ASCENDING,
	EVENTSSORT_TIMESTAMP_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventsSort) GetAllowedValues() []EventsSort {
	return allowedEventsSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventsSort(value)
	return nil
}

// NewEventsSortFromValue returns a pointer to a valid EventsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventsSortFromValue(v string) (*EventsSort, error) {
	ev := EventsSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventsSort: valid values are %v", v, allowedEventsSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventsSort) IsValid() bool {
	for _, existing := range allowedEventsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventsSort value.
func (v EventsSort) Ptr() *EventsSort {
	return &v
}
