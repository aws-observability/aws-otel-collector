// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScalarQuery - An individual scalar query to one of the basic Datadog data sources.
type ScalarQuery struct {
	MetricsScalarQuery      *MetricsScalarQuery
	EventsScalarQuery       *EventsScalarQuery
	ApmResourceStatsQuery   *ApmResourceStatsQuery
	ApmMetricsQuery         *ApmMetricsQuery
	ApmDependencyStatsQuery *ApmDependencyStatsQuery
	SloQuery                *SloQuery
	ProcessScalarQuery      *ProcessScalarQuery
	ContainerScalarQuery    *ContainerScalarQuery

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

// ApmResourceStatsQueryAsScalarQuery is a convenience function that returns ApmResourceStatsQuery wrapped in ScalarQuery.
func ApmResourceStatsQueryAsScalarQuery(v *ApmResourceStatsQuery) ScalarQuery {
	return ScalarQuery{ApmResourceStatsQuery: v}
}

// ApmMetricsQueryAsScalarQuery is a convenience function that returns ApmMetricsQuery wrapped in ScalarQuery.
func ApmMetricsQueryAsScalarQuery(v *ApmMetricsQuery) ScalarQuery {
	return ScalarQuery{ApmMetricsQuery: v}
}

// ApmDependencyStatsQueryAsScalarQuery is a convenience function that returns ApmDependencyStatsQuery wrapped in ScalarQuery.
func ApmDependencyStatsQueryAsScalarQuery(v *ApmDependencyStatsQuery) ScalarQuery {
	return ScalarQuery{ApmDependencyStatsQuery: v}
}

// SloQueryAsScalarQuery is a convenience function that returns SloQuery wrapped in ScalarQuery.
func SloQueryAsScalarQuery(v *SloQuery) ScalarQuery {
	return ScalarQuery{SloQuery: v}
}

// ProcessScalarQueryAsScalarQuery is a convenience function that returns ProcessScalarQuery wrapped in ScalarQuery.
func ProcessScalarQueryAsScalarQuery(v *ProcessScalarQuery) ScalarQuery {
	return ScalarQuery{ProcessScalarQuery: v}
}

// ContainerScalarQueryAsScalarQuery is a convenience function that returns ContainerScalarQuery wrapped in ScalarQuery.
func ContainerScalarQueryAsScalarQuery(v *ContainerScalarQuery) ScalarQuery {
	return ScalarQuery{ContainerScalarQuery: v}
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

	// try to unmarshal data into ProcessScalarQuery
	err = datadog.Unmarshal(data, &obj.ProcessScalarQuery)
	if err == nil {
		if obj.ProcessScalarQuery != nil && obj.ProcessScalarQuery.UnparsedObject == nil {
			jsonProcessScalarQuery, _ := datadog.Marshal(obj.ProcessScalarQuery)
			if string(jsonProcessScalarQuery) == "{}" { // empty struct
				obj.ProcessScalarQuery = nil
			} else {
				match++
			}
		} else {
			obj.ProcessScalarQuery = nil
		}
	} else {
		obj.ProcessScalarQuery = nil
	}

	// try to unmarshal data into ContainerScalarQuery
	err = datadog.Unmarshal(data, &obj.ContainerScalarQuery)
	if err == nil {
		if obj.ContainerScalarQuery != nil && obj.ContainerScalarQuery.UnparsedObject == nil {
			jsonContainerScalarQuery, _ := datadog.Marshal(obj.ContainerScalarQuery)
			if string(jsonContainerScalarQuery) == "{}" { // empty struct
				obj.ContainerScalarQuery = nil
			} else {
				match++
			}
		} else {
			obj.ContainerScalarQuery = nil
		}
	} else {
		obj.ContainerScalarQuery = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MetricsScalarQuery = nil
		obj.EventsScalarQuery = nil
		obj.ApmResourceStatsQuery = nil
		obj.ApmMetricsQuery = nil
		obj.ApmDependencyStatsQuery = nil
		obj.SloQuery = nil
		obj.ProcessScalarQuery = nil
		obj.ContainerScalarQuery = nil
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

	if obj.ProcessScalarQuery != nil {
		return datadog.Marshal(&obj.ProcessScalarQuery)
	}

	if obj.ContainerScalarQuery != nil {
		return datadog.Marshal(&obj.ContainerScalarQuery)
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

	if obj.ProcessScalarQuery != nil {
		return obj.ProcessScalarQuery
	}

	if obj.ContainerScalarQuery != nil {
		return obj.ContainerScalarQuery
	}

	// all schemas are nil
	return nil
}
