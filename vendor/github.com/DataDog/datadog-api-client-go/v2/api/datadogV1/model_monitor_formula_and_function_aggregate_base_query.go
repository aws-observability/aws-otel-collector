// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionAggregateBaseQuery - Base query for aggregate queries. Can be an events query or a metrics query.
type MonitorFormulaAndFunctionAggregateBaseQuery struct {
	MonitorFormulaAndFunctionEventQueryDefinition   *MonitorFormulaAndFunctionEventQueryDefinition
	MonitorFormulaAndFunctionMetricsQueryDefinition *MonitorFormulaAndFunctionMetricsQueryDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MonitorFormulaAndFunctionEventQueryDefinitionAsMonitorFormulaAndFunctionAggregateBaseQuery is a convenience function that returns MonitorFormulaAndFunctionEventQueryDefinition wrapped in MonitorFormulaAndFunctionAggregateBaseQuery.
func MonitorFormulaAndFunctionEventQueryDefinitionAsMonitorFormulaAndFunctionAggregateBaseQuery(v *MonitorFormulaAndFunctionEventQueryDefinition) MonitorFormulaAndFunctionAggregateBaseQuery {
	return MonitorFormulaAndFunctionAggregateBaseQuery{MonitorFormulaAndFunctionEventQueryDefinition: v}
}

// MonitorFormulaAndFunctionMetricsQueryDefinitionAsMonitorFormulaAndFunctionAggregateBaseQuery is a convenience function that returns MonitorFormulaAndFunctionMetricsQueryDefinition wrapped in MonitorFormulaAndFunctionAggregateBaseQuery.
func MonitorFormulaAndFunctionMetricsQueryDefinitionAsMonitorFormulaAndFunctionAggregateBaseQuery(v *MonitorFormulaAndFunctionMetricsQueryDefinition) MonitorFormulaAndFunctionAggregateBaseQuery {
	return MonitorFormulaAndFunctionAggregateBaseQuery{MonitorFormulaAndFunctionMetricsQueryDefinition: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *MonitorFormulaAndFunctionAggregateBaseQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MonitorFormulaAndFunctionEventQueryDefinition
	err = datadog.Unmarshal(data, &obj.MonitorFormulaAndFunctionEventQueryDefinition)
	if err == nil {
		if obj.MonitorFormulaAndFunctionEventQueryDefinition != nil && obj.MonitorFormulaAndFunctionEventQueryDefinition.UnparsedObject == nil {
			jsonMonitorFormulaAndFunctionEventQueryDefinition, _ := datadog.Marshal(obj.MonitorFormulaAndFunctionEventQueryDefinition)
			if string(jsonMonitorFormulaAndFunctionEventQueryDefinition) == "{}" { // empty struct
				obj.MonitorFormulaAndFunctionEventQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.MonitorFormulaAndFunctionEventQueryDefinition = nil
		}
	} else {
		obj.MonitorFormulaAndFunctionEventQueryDefinition = nil
	}

	// try to unmarshal data into MonitorFormulaAndFunctionMetricsQueryDefinition
	err = datadog.Unmarshal(data, &obj.MonitorFormulaAndFunctionMetricsQueryDefinition)
	if err == nil {
		if obj.MonitorFormulaAndFunctionMetricsQueryDefinition != nil && obj.MonitorFormulaAndFunctionMetricsQueryDefinition.UnparsedObject == nil {
			jsonMonitorFormulaAndFunctionMetricsQueryDefinition, _ := datadog.Marshal(obj.MonitorFormulaAndFunctionMetricsQueryDefinition)
			if string(jsonMonitorFormulaAndFunctionMetricsQueryDefinition) == "{}" { // empty struct
				obj.MonitorFormulaAndFunctionMetricsQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.MonitorFormulaAndFunctionMetricsQueryDefinition = nil
		}
	} else {
		obj.MonitorFormulaAndFunctionMetricsQueryDefinition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MonitorFormulaAndFunctionEventQueryDefinition = nil
		obj.MonitorFormulaAndFunctionMetricsQueryDefinition = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj MonitorFormulaAndFunctionAggregateBaseQuery) MarshalJSON() ([]byte, error) {
	if obj.MonitorFormulaAndFunctionEventQueryDefinition != nil {
		return datadog.Marshal(&obj.MonitorFormulaAndFunctionEventQueryDefinition)
	}

	if obj.MonitorFormulaAndFunctionMetricsQueryDefinition != nil {
		return datadog.Marshal(&obj.MonitorFormulaAndFunctionMetricsQueryDefinition)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *MonitorFormulaAndFunctionAggregateBaseQuery) GetActualInstance() interface{} {
	if obj.MonitorFormulaAndFunctionEventQueryDefinition != nil {
		return obj.MonitorFormulaAndFunctionEventQueryDefinition
	}

	if obj.MonitorFormulaAndFunctionMetricsQueryDefinition != nil {
		return obj.MonitorFormulaAndFunctionMetricsQueryDefinition
	}

	// all schemas are nil
	return nil
}
