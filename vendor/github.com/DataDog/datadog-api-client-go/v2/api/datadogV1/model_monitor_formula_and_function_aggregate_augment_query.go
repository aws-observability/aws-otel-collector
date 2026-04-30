// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionAggregateAugmentQuery - Augment query for aggregate augmented queries. Can be an events query or a reference table query.
type MonitorFormulaAndFunctionAggregateAugmentQuery struct {
	MonitorFormulaAndFunctionEventQueryDefinition          *MonitorFormulaAndFunctionEventQueryDefinition
	MonitorFormulaAndFunctionReferenceTableQueryDefinition *MonitorFormulaAndFunctionReferenceTableQueryDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MonitorFormulaAndFunctionEventQueryDefinitionAsMonitorFormulaAndFunctionAggregateAugmentQuery is a convenience function that returns MonitorFormulaAndFunctionEventQueryDefinition wrapped in MonitorFormulaAndFunctionAggregateAugmentQuery.
func MonitorFormulaAndFunctionEventQueryDefinitionAsMonitorFormulaAndFunctionAggregateAugmentQuery(v *MonitorFormulaAndFunctionEventQueryDefinition) MonitorFormulaAndFunctionAggregateAugmentQuery {
	return MonitorFormulaAndFunctionAggregateAugmentQuery{MonitorFormulaAndFunctionEventQueryDefinition: v}
}

// MonitorFormulaAndFunctionReferenceTableQueryDefinitionAsMonitorFormulaAndFunctionAggregateAugmentQuery is a convenience function that returns MonitorFormulaAndFunctionReferenceTableQueryDefinition wrapped in MonitorFormulaAndFunctionAggregateAugmentQuery.
func MonitorFormulaAndFunctionReferenceTableQueryDefinitionAsMonitorFormulaAndFunctionAggregateAugmentQuery(v *MonitorFormulaAndFunctionReferenceTableQueryDefinition) MonitorFormulaAndFunctionAggregateAugmentQuery {
	return MonitorFormulaAndFunctionAggregateAugmentQuery{MonitorFormulaAndFunctionReferenceTableQueryDefinition: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *MonitorFormulaAndFunctionAggregateAugmentQuery) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into MonitorFormulaAndFunctionReferenceTableQueryDefinition
	err = datadog.Unmarshal(data, &obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition)
	if err == nil {
		if obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition != nil && obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition.UnparsedObject == nil {
			jsonMonitorFormulaAndFunctionReferenceTableQueryDefinition, _ := datadog.Marshal(obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition)
			if string(jsonMonitorFormulaAndFunctionReferenceTableQueryDefinition) == "{}" { // empty struct
				obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition = nil
		}
	} else {
		obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MonitorFormulaAndFunctionEventQueryDefinition = nil
		obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj MonitorFormulaAndFunctionAggregateAugmentQuery) MarshalJSON() ([]byte, error) {
	if obj.MonitorFormulaAndFunctionEventQueryDefinition != nil {
		return datadog.Marshal(&obj.MonitorFormulaAndFunctionEventQueryDefinition)
	}

	if obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition != nil {
		return datadog.Marshal(&obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *MonitorFormulaAndFunctionAggregateAugmentQuery) GetActualInstance() interface{} {
	if obj.MonitorFormulaAndFunctionEventQueryDefinition != nil {
		return obj.MonitorFormulaAndFunctionEventQueryDefinition
	}

	if obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition != nil {
		return obj.MonitorFormulaAndFunctionReferenceTableQueryDefinition
	}

	// all schemas are nil
	return nil
}
