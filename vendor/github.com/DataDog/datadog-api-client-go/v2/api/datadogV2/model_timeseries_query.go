// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesQuery - An individual timeseries query to one of the basic Datadog data sources.
type TimeseriesQuery struct {
	MetricsTimeseriesQuery *MetricsTimeseriesQuery
	EventsTimeseriesQuery  *EventsTimeseriesQuery

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MetricsTimeseriesQueryAsTimeseriesQuery is a convenience function that returns MetricsTimeseriesQuery wrapped in TimeseriesQuery.
func MetricsTimeseriesQueryAsTimeseriesQuery(v *MetricsTimeseriesQuery) TimeseriesQuery {
	return TimeseriesQuery{MetricsTimeseriesQuery: v}
}

// EventsTimeseriesQueryAsTimeseriesQuery is a convenience function that returns EventsTimeseriesQuery wrapped in TimeseriesQuery.
func EventsTimeseriesQueryAsTimeseriesQuery(v *EventsTimeseriesQuery) TimeseriesQuery {
	return TimeseriesQuery{EventsTimeseriesQuery: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TimeseriesQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MetricsTimeseriesQuery
	err = datadog.Unmarshal(data, &obj.MetricsTimeseriesQuery)
	if err == nil {
		if obj.MetricsTimeseriesQuery != nil && obj.MetricsTimeseriesQuery.UnparsedObject == nil {
			jsonMetricsTimeseriesQuery, _ := datadog.Marshal(obj.MetricsTimeseriesQuery)
			if string(jsonMetricsTimeseriesQuery) == "{}" { // empty struct
				obj.MetricsTimeseriesQuery = nil
			} else {
				match++
			}
		} else {
			obj.MetricsTimeseriesQuery = nil
		}
	} else {
		obj.MetricsTimeseriesQuery = nil
	}

	// try to unmarshal data into EventsTimeseriesQuery
	err = datadog.Unmarshal(data, &obj.EventsTimeseriesQuery)
	if err == nil {
		if obj.EventsTimeseriesQuery != nil && obj.EventsTimeseriesQuery.UnparsedObject == nil {
			jsonEventsTimeseriesQuery, _ := datadog.Marshal(obj.EventsTimeseriesQuery)
			if string(jsonEventsTimeseriesQuery) == "{}" { // empty struct
				obj.EventsTimeseriesQuery = nil
			} else {
				match++
			}
		} else {
			obj.EventsTimeseriesQuery = nil
		}
	} else {
		obj.EventsTimeseriesQuery = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MetricsTimeseriesQuery = nil
		obj.EventsTimeseriesQuery = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TimeseriesQuery) MarshalJSON() ([]byte, error) {
	if obj.MetricsTimeseriesQuery != nil {
		return datadog.Marshal(&obj.MetricsTimeseriesQuery)
	}

	if obj.EventsTimeseriesQuery != nil {
		return datadog.Marshal(&obj.EventsTimeseriesQuery)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TimeseriesQuery) GetActualInstance() interface{} {
	if obj.MetricsTimeseriesQuery != nil {
		return obj.MetricsTimeseriesQuery
	}

	if obj.EventsTimeseriesQuery != nil {
		return obj.EventsTimeseriesQuery
	}

	// all schemas are nil
	return nil
}
