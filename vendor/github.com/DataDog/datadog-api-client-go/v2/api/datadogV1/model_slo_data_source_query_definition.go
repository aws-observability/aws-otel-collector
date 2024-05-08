// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLODataSourceQueryDefinition - A formula and function query.
type SLODataSourceQueryDefinition struct {
	FormulaAndFunctionMetricQueryDefinition *FormulaAndFunctionMetricQueryDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FormulaAndFunctionMetricQueryDefinitionAsSLODataSourceQueryDefinition is a convenience function that returns FormulaAndFunctionMetricQueryDefinition wrapped in SLODataSourceQueryDefinition.
func FormulaAndFunctionMetricQueryDefinitionAsSLODataSourceQueryDefinition(v *FormulaAndFunctionMetricQueryDefinition) SLODataSourceQueryDefinition {
	return SLODataSourceQueryDefinition{FormulaAndFunctionMetricQueryDefinition: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SLODataSourceQueryDefinition) UnmarshalJSON(data []byte) error {
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

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FormulaAndFunctionMetricQueryDefinition = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SLODataSourceQueryDefinition) MarshalJSON() ([]byte, error) {
	if obj.FormulaAndFunctionMetricQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionMetricQueryDefinition)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SLODataSourceQueryDefinition) GetActualInstance() interface{} {
	if obj.FormulaAndFunctionMetricQueryDefinition != nil {
		return obj.FormulaAndFunctionMetricQueryDefinition
	}

	// all schemas are nil
	return nil
}
