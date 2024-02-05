// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionWidgetHistogramRequestQuery - Query definition for Distribution Widget Histogram Request
type DistributionWidgetHistogramRequestQuery struct {
	FormulaAndFunctionMetricQueryDefinition           *FormulaAndFunctionMetricQueryDefinition
	FormulaAndFunctionEventQueryDefinition            *FormulaAndFunctionEventQueryDefinition
	FormulaAndFunctionApmResourceStatsQueryDefinition *FormulaAndFunctionApmResourceStatsQueryDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FormulaAndFunctionMetricQueryDefinitionAsDistributionWidgetHistogramRequestQuery is a convenience function that returns FormulaAndFunctionMetricQueryDefinition wrapped in DistributionWidgetHistogramRequestQuery.
func FormulaAndFunctionMetricQueryDefinitionAsDistributionWidgetHistogramRequestQuery(v *FormulaAndFunctionMetricQueryDefinition) DistributionWidgetHistogramRequestQuery {
	return DistributionWidgetHistogramRequestQuery{FormulaAndFunctionMetricQueryDefinition: v}
}

// FormulaAndFunctionEventQueryDefinitionAsDistributionWidgetHistogramRequestQuery is a convenience function that returns FormulaAndFunctionEventQueryDefinition wrapped in DistributionWidgetHistogramRequestQuery.
func FormulaAndFunctionEventQueryDefinitionAsDistributionWidgetHistogramRequestQuery(v *FormulaAndFunctionEventQueryDefinition) DistributionWidgetHistogramRequestQuery {
	return DistributionWidgetHistogramRequestQuery{FormulaAndFunctionEventQueryDefinition: v}
}

// FormulaAndFunctionApmResourceStatsQueryDefinitionAsDistributionWidgetHistogramRequestQuery is a convenience function that returns FormulaAndFunctionApmResourceStatsQueryDefinition wrapped in DistributionWidgetHistogramRequestQuery.
func FormulaAndFunctionApmResourceStatsQueryDefinitionAsDistributionWidgetHistogramRequestQuery(v *FormulaAndFunctionApmResourceStatsQueryDefinition) DistributionWidgetHistogramRequestQuery {
	return DistributionWidgetHistogramRequestQuery{FormulaAndFunctionApmResourceStatsQueryDefinition: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DistributionWidgetHistogramRequestQuery) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FormulaAndFunctionMetricQueryDefinition
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionMetricQueryDefinition)
	if err == nil {
		if obj.FormulaAndFunctionMetricQueryDefinition != nil && obj.FormulaAndFunctionMetricQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionMetricQueryDefinition, _ := datadog.Marshal(obj.FormulaAndFunctionMetricQueryDefinition)
			if string(jsonFormulaAndFunctionMetricQueryDefinition) == "{}" { // empty struct
				obj.FormulaAndFunctionMetricQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionMetricQueryDefinition = nil
		}
	} else {
		obj.FormulaAndFunctionMetricQueryDefinition = nil
	}

	// try to unmarshal data into FormulaAndFunctionEventQueryDefinition
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionEventQueryDefinition)
	if err == nil {
		if obj.FormulaAndFunctionEventQueryDefinition != nil && obj.FormulaAndFunctionEventQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionEventQueryDefinition, _ := datadog.Marshal(obj.FormulaAndFunctionEventQueryDefinition)
			if string(jsonFormulaAndFunctionEventQueryDefinition) == "{}" { // empty struct
				obj.FormulaAndFunctionEventQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionEventQueryDefinition = nil
		}
	} else {
		obj.FormulaAndFunctionEventQueryDefinition = nil
	}

	// try to unmarshal data into FormulaAndFunctionApmResourceStatsQueryDefinition
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionApmResourceStatsQueryDefinition)
	if err == nil {
		if obj.FormulaAndFunctionApmResourceStatsQueryDefinition != nil && obj.FormulaAndFunctionApmResourceStatsQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionApmResourceStatsQueryDefinition, _ := datadog.Marshal(obj.FormulaAndFunctionApmResourceStatsQueryDefinition)
			if string(jsonFormulaAndFunctionApmResourceStatsQueryDefinition) == "{}" { // empty struct
				obj.FormulaAndFunctionApmResourceStatsQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionApmResourceStatsQueryDefinition = nil
		}
	} else {
		obj.FormulaAndFunctionApmResourceStatsQueryDefinition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FormulaAndFunctionMetricQueryDefinition = nil
		obj.FormulaAndFunctionEventQueryDefinition = nil
		obj.FormulaAndFunctionApmResourceStatsQueryDefinition = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DistributionWidgetHistogramRequestQuery) MarshalJSON() ([]byte, error) {
	if obj.FormulaAndFunctionMetricQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionMetricQueryDefinition)
	}

	if obj.FormulaAndFunctionEventQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionEventQueryDefinition)
	}

	if obj.FormulaAndFunctionApmResourceStatsQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionApmResourceStatsQueryDefinition)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DistributionWidgetHistogramRequestQuery) GetActualInstance() interface{} {
	if obj.FormulaAndFunctionMetricQueryDefinition != nil {
		return obj.FormulaAndFunctionMetricQueryDefinition
	}

	if obj.FormulaAndFunctionEventQueryDefinition != nil {
		return obj.FormulaAndFunctionEventQueryDefinition
	}

	if obj.FormulaAndFunctionApmResourceStatsQueryDefinition != nil {
		return obj.FormulaAndFunctionApmResourceStatsQueryDefinition
	}

	// all schemas are nil
	return nil
}
