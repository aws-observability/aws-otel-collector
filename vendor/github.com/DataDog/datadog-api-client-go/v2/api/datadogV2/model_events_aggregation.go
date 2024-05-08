// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsAggregation The type of aggregation that can be performed on events-based queries.
type EventsAggregation string

// List of EventsAggregation.
const (
	EVENTSAGGREGATION_COUNT       EventsAggregation = "count"
	EVENTSAGGREGATION_CARDINALITY EventsAggregation = "cardinality"
	EVENTSAGGREGATION_PC75        EventsAggregation = "pc75"
	EVENTSAGGREGATION_PC90        EventsAggregation = "pc90"
	EVENTSAGGREGATION_PC95        EventsAggregation = "pc95"
	EVENTSAGGREGATION_PC98        EventsAggregation = "pc98"
	EVENTSAGGREGATION_PC99        EventsAggregation = "pc99"
	EVENTSAGGREGATION_SUM         EventsAggregation = "sum"
	EVENTSAGGREGATION_MIN         EventsAggregation = "min"
	EVENTSAGGREGATION_MAX         EventsAggregation = "max"
	EVENTSAGGREGATION_AVG         EventsAggregation = "avg"
)

var allowedEventsAggregationEnumValues = []EventsAggregation{
	EVENTSAGGREGATION_COUNT,
	EVENTSAGGREGATION_CARDINALITY,
	EVENTSAGGREGATION_PC75,
	EVENTSAGGREGATION_PC90,
	EVENTSAGGREGATION_PC95,
	EVENTSAGGREGATION_PC98,
	EVENTSAGGREGATION_PC99,
	EVENTSAGGREGATION_SUM,
	EVENTSAGGREGATION_MIN,
	EVENTSAGGREGATION_MAX,
	EVENTSAGGREGATION_AVG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventsAggregation) GetAllowedValues() []EventsAggregation {
	return allowedEventsAggregationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventsAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventsAggregation(value)
	return nil
}

// NewEventsAggregationFromValue returns a pointer to a valid EventsAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventsAggregationFromValue(v string) (*EventsAggregation, error) {
	ev := EventsAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventsAggregation: valid values are %v", v, allowedEventsAggregationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventsAggregation) IsValid() bool {
	for _, existing := range allowedEventsAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventsAggregation value.
func (v EventsAggregation) Ptr() *EventsAggregation {
	return &v
}
