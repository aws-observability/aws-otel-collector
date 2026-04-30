// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionQueryDefinition - A formula and function query.
type MonitorFormulaAndFunctionQueryDefinition struct {
	MonitorFormulaAndFunctionEventQueryDefinition              *MonitorFormulaAndFunctionEventQueryDefinition
	MonitorFormulaAndFunctionCostQueryDefinition               *MonitorFormulaAndFunctionCostQueryDefinition
	MonitorFormulaAndFunctionDataQualityQueryDefinition        *MonitorFormulaAndFunctionDataQualityQueryDefinition
	MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition
	MonitorFormulaAndFunctionAggregateFilteredQueryDefinition  *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MonitorFormulaAndFunctionEventQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition is a convenience function that returns MonitorFormulaAndFunctionEventQueryDefinition wrapped in MonitorFormulaAndFunctionQueryDefinition.
func MonitorFormulaAndFunctionEventQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition(v *MonitorFormulaAndFunctionEventQueryDefinition) MonitorFormulaAndFunctionQueryDefinition {
	return MonitorFormulaAndFunctionQueryDefinition{MonitorFormulaAndFunctionEventQueryDefinition: v}
}

// MonitorFormulaAndFunctionCostQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition is a convenience function that returns MonitorFormulaAndFunctionCostQueryDefinition wrapped in MonitorFormulaAndFunctionQueryDefinition.
func MonitorFormulaAndFunctionCostQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition(v *MonitorFormulaAndFunctionCostQueryDefinition) MonitorFormulaAndFunctionQueryDefinition {
	return MonitorFormulaAndFunctionQueryDefinition{MonitorFormulaAndFunctionCostQueryDefinition: v}
}

// MonitorFormulaAndFunctionDataQualityQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition is a convenience function that returns MonitorFormulaAndFunctionDataQualityQueryDefinition wrapped in MonitorFormulaAndFunctionQueryDefinition.
func MonitorFormulaAndFunctionDataQualityQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition(v *MonitorFormulaAndFunctionDataQualityQueryDefinition) MonitorFormulaAndFunctionQueryDefinition {
	return MonitorFormulaAndFunctionQueryDefinition{MonitorFormulaAndFunctionDataQualityQueryDefinition: v}
}

// MonitorFormulaAndFunctionAggregateAugmentedQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition is a convenience function that returns MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition wrapped in MonitorFormulaAndFunctionQueryDefinition.
func MonitorFormulaAndFunctionAggregateAugmentedQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition(v *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) MonitorFormulaAndFunctionQueryDefinition {
	return MonitorFormulaAndFunctionQueryDefinition{MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition: v}
}

// MonitorFormulaAndFunctionAggregateFilteredQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition is a convenience function that returns MonitorFormulaAndFunctionAggregateFilteredQueryDefinition wrapped in MonitorFormulaAndFunctionQueryDefinition.
func MonitorFormulaAndFunctionAggregateFilteredQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition(v *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) MonitorFormulaAndFunctionQueryDefinition {
	return MonitorFormulaAndFunctionQueryDefinition{MonitorFormulaAndFunctionAggregateFilteredQueryDefinition: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *MonitorFormulaAndFunctionQueryDefinition) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into MonitorFormulaAndFunctionCostQueryDefinition
	err = datadog.Unmarshal(data, &obj.MonitorFormulaAndFunctionCostQueryDefinition)
	if err == nil {
		if obj.MonitorFormulaAndFunctionCostQueryDefinition != nil && obj.MonitorFormulaAndFunctionCostQueryDefinition.UnparsedObject == nil {
			jsonMonitorFormulaAndFunctionCostQueryDefinition, _ := datadog.Marshal(obj.MonitorFormulaAndFunctionCostQueryDefinition)
			if string(jsonMonitorFormulaAndFunctionCostQueryDefinition) == "{}" { // empty struct
				obj.MonitorFormulaAndFunctionCostQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.MonitorFormulaAndFunctionCostQueryDefinition = nil
		}
	} else {
		obj.MonitorFormulaAndFunctionCostQueryDefinition = nil
	}

	// try to unmarshal data into MonitorFormulaAndFunctionDataQualityQueryDefinition
	err = datadog.Unmarshal(data, &obj.MonitorFormulaAndFunctionDataQualityQueryDefinition)
	if err == nil {
		if obj.MonitorFormulaAndFunctionDataQualityQueryDefinition != nil && obj.MonitorFormulaAndFunctionDataQualityQueryDefinition.UnparsedObject == nil {
			jsonMonitorFormulaAndFunctionDataQualityQueryDefinition, _ := datadog.Marshal(obj.MonitorFormulaAndFunctionDataQualityQueryDefinition)
			if string(jsonMonitorFormulaAndFunctionDataQualityQueryDefinition) == "{}" { // empty struct
				obj.MonitorFormulaAndFunctionDataQualityQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.MonitorFormulaAndFunctionDataQualityQueryDefinition = nil
		}
	} else {
		obj.MonitorFormulaAndFunctionDataQualityQueryDefinition = nil
	}

	// try to unmarshal data into MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition
	err = datadog.Unmarshal(data, &obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition)
	if err == nil {
		if obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition != nil && obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition.UnparsedObject == nil {
			jsonMonitorFormulaAndFunctionAggregateAugmentedQueryDefinition, _ := datadog.Marshal(obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition)
			if string(jsonMonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) == "{}" { // empty struct
				obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition = nil
		}
	} else {
		obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition = nil
	}

	// try to unmarshal data into MonitorFormulaAndFunctionAggregateFilteredQueryDefinition
	err = datadog.Unmarshal(data, &obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition)
	if err == nil {
		if obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition != nil && obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition.UnparsedObject == nil {
			jsonMonitorFormulaAndFunctionAggregateFilteredQueryDefinition, _ := datadog.Marshal(obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition)
			if string(jsonMonitorFormulaAndFunctionAggregateFilteredQueryDefinition) == "{}" { // empty struct
				obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition = nil
		}
	} else {
		obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MonitorFormulaAndFunctionEventQueryDefinition = nil
		obj.MonitorFormulaAndFunctionCostQueryDefinition = nil
		obj.MonitorFormulaAndFunctionDataQualityQueryDefinition = nil
		obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition = nil
		obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj MonitorFormulaAndFunctionQueryDefinition) MarshalJSON() ([]byte, error) {
	if obj.MonitorFormulaAndFunctionEventQueryDefinition != nil {
		return datadog.Marshal(&obj.MonitorFormulaAndFunctionEventQueryDefinition)
	}

	if obj.MonitorFormulaAndFunctionCostQueryDefinition != nil {
		return datadog.Marshal(&obj.MonitorFormulaAndFunctionCostQueryDefinition)
	}

	if obj.MonitorFormulaAndFunctionDataQualityQueryDefinition != nil {
		return datadog.Marshal(&obj.MonitorFormulaAndFunctionDataQualityQueryDefinition)
	}

	if obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition != nil {
		return datadog.Marshal(&obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition)
	}

	if obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition != nil {
		return datadog.Marshal(&obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *MonitorFormulaAndFunctionQueryDefinition) GetActualInstance() interface{} {
	if obj.MonitorFormulaAndFunctionEventQueryDefinition != nil {
		return obj.MonitorFormulaAndFunctionEventQueryDefinition
	}

	if obj.MonitorFormulaAndFunctionCostQueryDefinition != nil {
		return obj.MonitorFormulaAndFunctionCostQueryDefinition
	}

	if obj.MonitorFormulaAndFunctionDataQualityQueryDefinition != nil {
		return obj.MonitorFormulaAndFunctionDataQualityQueryDefinition
	}

	if obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition != nil {
		return obj.MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition
	}

	if obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition != nil {
		return obj.MonitorFormulaAndFunctionAggregateFilteredQueryDefinition
	}

	// all schemas are nil
	return nil
}
