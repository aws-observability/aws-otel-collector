// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsAggregation - The type of aggregation that can be performed on events-based queries.
type EventsAggregation struct {
	EventsAggregationValue      *EventsAggregationValue
	EventsAggregationPercentile *string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// EventsAggregationValueAsEventsAggregation is a convenience function that returns EventsAggregationValue wrapped in EventsAggregation.
func EventsAggregationValueAsEventsAggregation(v *EventsAggregationValue) EventsAggregation {
	return EventsAggregation{EventsAggregationValue: v}
}

// EventsAggregationPercentileAsEventsAggregation is a convenience function that returns string wrapped in EventsAggregation.
func EventsAggregationPercentileAsEventsAggregation(v *string) EventsAggregation {
	return EventsAggregation{EventsAggregationPercentile: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *EventsAggregation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EventsAggregationValue
	err = datadog.Unmarshal(data, &obj.EventsAggregationValue)
	if err == nil {
		if obj.EventsAggregationValue != nil {
			jsonEventsAggregationValue, _ := datadog.Marshal(obj.EventsAggregationValue)
			if string(jsonEventsAggregationValue) == "{}" && string(data) != "{}" { // empty struct
				obj.EventsAggregationValue = nil
			} else {
				match++
			}
		} else {
			obj.EventsAggregationValue = nil
		}
	} else {
		obj.EventsAggregationValue = nil
	}

	// try to unmarshal data into EventsAggregationPercentile
	err = datadog.Unmarshal(data, &obj.EventsAggregationPercentile)
	if err == nil {
		if obj.EventsAggregationPercentile != nil {
			jsonEventsAggregationPercentile, _ := datadog.Marshal(obj.EventsAggregationPercentile)
			if string(jsonEventsAggregationPercentile) == "{}" { // empty struct
				obj.EventsAggregationPercentile = nil
			} else {
				match++
			}
		} else {
			obj.EventsAggregationPercentile = nil
		}
	} else {
		obj.EventsAggregationPercentile = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.EventsAggregationValue = nil
		obj.EventsAggregationPercentile = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj EventsAggregation) MarshalJSON() ([]byte, error) {
	if obj.EventsAggregationValue != nil {
		return datadog.Marshal(&obj.EventsAggregationValue)
	}

	if obj.EventsAggregationPercentile != nil {
		return datadog.Marshal(&obj.EventsAggregationPercentile)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *EventsAggregation) GetActualInstance() interface{} {
	if obj.EventsAggregationValue != nil {
		return obj.EventsAggregationValue
	}

	if obj.EventsAggregationPercentile != nil {
		return obj.EventsAggregationPercentile
	}

	// all schemas are nil
	return nil
}
