// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsAggregationValue Standard aggregation types for events-based queries.
type EventsAggregationValue string

// List of EventsAggregationValue.
const (
	EVENTSAGGREGATIONVALUE_AVG           EventsAggregationValue = "avg"
	EVENTSAGGREGATIONVALUE_CARDINALITY   EventsAggregationValue = "cardinality"
	EVENTSAGGREGATIONVALUE_COUNT         EventsAggregationValue = "count"
	EVENTSAGGREGATIONVALUE_DELTA         EventsAggregationValue = "delta"
	EVENTSAGGREGATIONVALUE_EARLIEST      EventsAggregationValue = "earliest"
	EVENTSAGGREGATIONVALUE_LATEST        EventsAggregationValue = "latest"
	EVENTSAGGREGATIONVALUE_MAX           EventsAggregationValue = "max"
	EVENTSAGGREGATIONVALUE_MEDIAN        EventsAggregationValue = "median"
	EVENTSAGGREGATIONVALUE_MIN           EventsAggregationValue = "min"
	EVENTSAGGREGATIONVALUE_MOST_FREQUENT EventsAggregationValue = "most_frequent"
	EVENTSAGGREGATIONVALUE_SUM           EventsAggregationValue = "sum"
)

var allowedEventsAggregationValueEnumValues = []EventsAggregationValue{
	EVENTSAGGREGATIONVALUE_AVG,
	EVENTSAGGREGATIONVALUE_CARDINALITY,
	EVENTSAGGREGATIONVALUE_COUNT,
	EVENTSAGGREGATIONVALUE_DELTA,
	EVENTSAGGREGATIONVALUE_EARLIEST,
	EVENTSAGGREGATIONVALUE_LATEST,
	EVENTSAGGREGATIONVALUE_MAX,
	EVENTSAGGREGATIONVALUE_MEDIAN,
	EVENTSAGGREGATIONVALUE_MIN,
	EVENTSAGGREGATIONVALUE_MOST_FREQUENT,
	EVENTSAGGREGATIONVALUE_SUM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventsAggregationValue) GetAllowedValues() []EventsAggregationValue {
	return allowedEventsAggregationValueEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventsAggregationValue) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventsAggregationValue(value)
	return nil
}

// NewEventsAggregationValueFromValue returns a pointer to a valid EventsAggregationValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventsAggregationValueFromValue(v string) (*EventsAggregationValue, error) {
	ev := EventsAggregationValue(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventsAggregationValue: valid values are %v", v, allowedEventsAggregationValueEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventsAggregationValue) IsValid() bool {
	for _, existing := range allowedEventsAggregationValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventsAggregationValue value.
func (v EventsAggregationValue) Ptr() *EventsAggregationValue {
	return &v
}
