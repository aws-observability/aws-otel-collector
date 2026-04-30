// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesQuery - An individual timeseries query to one of the basic Datadog data sources.
type TimeseriesQuery struct {
	MetricsTimeseriesQuery   *MetricsTimeseriesQuery
	EventsTimeseriesQuery    *EventsTimeseriesQuery
	ApmResourceStatsQuery    *ApmResourceStatsQuery
	ApmMetricsQuery          *ApmMetricsQuery
	ApmDependencyStatsQuery  *ApmDependencyStatsQuery
	SloQuery                 *SloQuery
	ProcessTimeseriesQuery   *ProcessTimeseriesQuery
	ContainerTimeseriesQuery *ContainerTimeseriesQuery

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

// ApmResourceStatsQueryAsTimeseriesQuery is a convenience function that returns ApmResourceStatsQuery wrapped in TimeseriesQuery.
func ApmResourceStatsQueryAsTimeseriesQuery(v *ApmResourceStatsQuery) TimeseriesQuery {
	return TimeseriesQuery{ApmResourceStatsQuery: v}
}

// ApmMetricsQueryAsTimeseriesQuery is a convenience function that returns ApmMetricsQuery wrapped in TimeseriesQuery.
func ApmMetricsQueryAsTimeseriesQuery(v *ApmMetricsQuery) TimeseriesQuery {
	return TimeseriesQuery{ApmMetricsQuery: v}
}

// ApmDependencyStatsQueryAsTimeseriesQuery is a convenience function that returns ApmDependencyStatsQuery wrapped in TimeseriesQuery.
func ApmDependencyStatsQueryAsTimeseriesQuery(v *ApmDependencyStatsQuery) TimeseriesQuery {
	return TimeseriesQuery{ApmDependencyStatsQuery: v}
}

// SloQueryAsTimeseriesQuery is a convenience function that returns SloQuery wrapped in TimeseriesQuery.
func SloQueryAsTimeseriesQuery(v *SloQuery) TimeseriesQuery {
	return TimeseriesQuery{SloQuery: v}
}

// ProcessTimeseriesQueryAsTimeseriesQuery is a convenience function that returns ProcessTimeseriesQuery wrapped in TimeseriesQuery.
func ProcessTimeseriesQueryAsTimeseriesQuery(v *ProcessTimeseriesQuery) TimeseriesQuery {
	return TimeseriesQuery{ProcessTimeseriesQuery: v}
}

// ContainerTimeseriesQueryAsTimeseriesQuery is a convenience function that returns ContainerTimeseriesQuery wrapped in TimeseriesQuery.
func ContainerTimeseriesQueryAsTimeseriesQuery(v *ContainerTimeseriesQuery) TimeseriesQuery {
	return TimeseriesQuery{ContainerTimeseriesQuery: v}
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

	// try to unmarshal data into ApmResourceStatsQuery
	err = datadog.Unmarshal(data, &obj.ApmResourceStatsQuery)
	if err == nil {
		if obj.ApmResourceStatsQuery != nil && obj.ApmResourceStatsQuery.UnparsedObject == nil {
			jsonApmResourceStatsQuery, _ := datadog.Marshal(obj.ApmResourceStatsQuery)
			if string(jsonApmResourceStatsQuery) == "{}" { // empty struct
				obj.ApmResourceStatsQuery = nil
			} else {
				match++
			}
		} else {
			obj.ApmResourceStatsQuery = nil
		}
	} else {
		obj.ApmResourceStatsQuery = nil
	}

	// try to unmarshal data into ApmMetricsQuery
	err = datadog.Unmarshal(data, &obj.ApmMetricsQuery)
	if err == nil {
		if obj.ApmMetricsQuery != nil && obj.ApmMetricsQuery.UnparsedObject == nil {
			jsonApmMetricsQuery, _ := datadog.Marshal(obj.ApmMetricsQuery)
			if string(jsonApmMetricsQuery) == "{}" { // empty struct
				obj.ApmMetricsQuery = nil
			} else {
				match++
			}
		} else {
			obj.ApmMetricsQuery = nil
		}
	} else {
		obj.ApmMetricsQuery = nil
	}

	// try to unmarshal data into ApmDependencyStatsQuery
	err = datadog.Unmarshal(data, &obj.ApmDependencyStatsQuery)
	if err == nil {
		if obj.ApmDependencyStatsQuery != nil && obj.ApmDependencyStatsQuery.UnparsedObject == nil {
			jsonApmDependencyStatsQuery, _ := datadog.Marshal(obj.ApmDependencyStatsQuery)
			if string(jsonApmDependencyStatsQuery) == "{}" { // empty struct
				obj.ApmDependencyStatsQuery = nil
			} else {
				match++
			}
		} else {
			obj.ApmDependencyStatsQuery = nil
		}
	} else {
		obj.ApmDependencyStatsQuery = nil
	}

	// try to unmarshal data into SloQuery
	err = datadog.Unmarshal(data, &obj.SloQuery)
	if err == nil {
		if obj.SloQuery != nil && obj.SloQuery.UnparsedObject == nil {
			jsonSloQuery, _ := datadog.Marshal(obj.SloQuery)
			if string(jsonSloQuery) == "{}" { // empty struct
				obj.SloQuery = nil
			} else {
				match++
			}
		} else {
			obj.SloQuery = nil
		}
	} else {
		obj.SloQuery = nil
	}

	// try to unmarshal data into ProcessTimeseriesQuery
	err = datadog.Unmarshal(data, &obj.ProcessTimeseriesQuery)
	if err == nil {
		if obj.ProcessTimeseriesQuery != nil && obj.ProcessTimeseriesQuery.UnparsedObject == nil {
			jsonProcessTimeseriesQuery, _ := datadog.Marshal(obj.ProcessTimeseriesQuery)
			if string(jsonProcessTimeseriesQuery) == "{}" { // empty struct
				obj.ProcessTimeseriesQuery = nil
			} else {
				match++
			}
		} else {
			obj.ProcessTimeseriesQuery = nil
		}
	} else {
		obj.ProcessTimeseriesQuery = nil
	}

	// try to unmarshal data into ContainerTimeseriesQuery
	err = datadog.Unmarshal(data, &obj.ContainerTimeseriesQuery)
	if err == nil {
		if obj.ContainerTimeseriesQuery != nil && obj.ContainerTimeseriesQuery.UnparsedObject == nil {
			jsonContainerTimeseriesQuery, _ := datadog.Marshal(obj.ContainerTimeseriesQuery)
			if string(jsonContainerTimeseriesQuery) == "{}" { // empty struct
				obj.ContainerTimeseriesQuery = nil
			} else {
				match++
			}
		} else {
			obj.ContainerTimeseriesQuery = nil
		}
	} else {
		obj.ContainerTimeseriesQuery = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MetricsTimeseriesQuery = nil
		obj.EventsTimeseriesQuery = nil
		obj.ApmResourceStatsQuery = nil
		obj.ApmMetricsQuery = nil
		obj.ApmDependencyStatsQuery = nil
		obj.SloQuery = nil
		obj.ProcessTimeseriesQuery = nil
		obj.ContainerTimeseriesQuery = nil
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

	if obj.ApmResourceStatsQuery != nil {
		return datadog.Marshal(&obj.ApmResourceStatsQuery)
	}

	if obj.ApmMetricsQuery != nil {
		return datadog.Marshal(&obj.ApmMetricsQuery)
	}

	if obj.ApmDependencyStatsQuery != nil {
		return datadog.Marshal(&obj.ApmDependencyStatsQuery)
	}

	if obj.SloQuery != nil {
		return datadog.Marshal(&obj.SloQuery)
	}

	if obj.ProcessTimeseriesQuery != nil {
		return datadog.Marshal(&obj.ProcessTimeseriesQuery)
	}

	if obj.ContainerTimeseriesQuery != nil {
		return datadog.Marshal(&obj.ContainerTimeseriesQuery)
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

	if obj.ApmResourceStatsQuery != nil {
		return obj.ApmResourceStatsQuery
	}

	if obj.ApmMetricsQuery != nil {
		return obj.ApmMetricsQuery
	}

	if obj.ApmDependencyStatsQuery != nil {
		return obj.ApmDependencyStatsQuery
	}

	if obj.SloQuery != nil {
		return obj.SloQuery
	}

	if obj.ProcessTimeseriesQuery != nil {
		return obj.ProcessTimeseriesQuery
	}

	if obj.ContainerTimeseriesQuery != nil {
		return obj.ContainerTimeseriesQuery
	}

	// all schemas are nil
	return nil
}
