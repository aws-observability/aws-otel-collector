// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScalarQuery - An individual scalar query to one of the basic Datadog data sources.
type ScalarQuery struct {
	MetricsScalarQuery *MetricsScalarQuery
	EventsScalarQuery  *EventsScalarQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MetricsScalarQueryAsScalarQuery is a convenience function that returns MetricsScalarQuery wrapped in ScalarQuery.
func MetricsScalarQueryAsScalarQuery(v *MetricsScalarQuery) ScalarQuery {
	return ScalarQuery{MetricsScalarQuery: v}
}

// EventsScalarQueryAsScalarQuery is a convenience function that returns EventsScalarQuery wrapped in ScalarQuery.
func EventsScalarQueryAsScalarQuery(v *EventsScalarQuery) ScalarQuery {
	return ScalarQuery{EventsScalarQuery: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ScalarQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MetricsScalarQuery
	err = datadog.Unmarshal(data, &obj.MetricsScalarQuery)
	if err == nil {
		if obj.MetricsScalarQuery != nil && obj.MetricsScalarQuery.UnparsedObject == nil {
			jsonMetricsScalarQuery, _ := datadog.Marshal(obj.MetricsScalarQuery)
			if string(jsonMetricsScalarQuery) == "{}" { // empty struct
				obj.MetricsScalarQuery = nil
			} else {
				match++
			}
		} else {
			obj.MetricsScalarQuery = nil
		}
	} else {
		obj.MetricsScalarQuery = nil
	}

	// try to unmarshal data into EventsScalarQuery
	err = datadog.Unmarshal(data, &obj.EventsScalarQuery)
	if err == nil {
		if obj.EventsScalarQuery != nil && obj.EventsScalarQuery.UnparsedObject == nil {
			jsonEventsScalarQuery, _ := datadog.Marshal(obj.EventsScalarQuery)
			if string(jsonEventsScalarQuery) == "{}" { // empty struct
				obj.EventsScalarQuery = nil
			} else {
				match++
			}
		} else {
			obj.EventsScalarQuery = nil
		}
	} else {
		obj.EventsScalarQuery = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MetricsScalarQuery = nil
		obj.EventsScalarQuery = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ScalarQuery) MarshalJSON() ([]byte, error) {
	if obj.MetricsScalarQuery != nil {
		return datadog.Marshal(&obj.MetricsScalarQuery)
	}

	if obj.EventsScalarQuery != nil {
		return datadog.Marshal(&obj.EventsScalarQuery)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ScalarQuery) GetActualInstance() interface{} {
	if obj.MetricsScalarQuery != nil {
		return obj.MetricsScalarQuery
	}

	if obj.EventsScalarQuery != nil {
		return obj.EventsScalarQuery
	}

	// all schemas are nil
	return nil
}
