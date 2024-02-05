// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionQueryDefinition - A formula and function query.
type FormulaAndFunctionQueryDefinition struct {
	FormulaAndFunctionMetricQueryDefinition             *FormulaAndFunctionMetricQueryDefinition
	FormulaAndFunctionEventQueryDefinition              *FormulaAndFunctionEventQueryDefinition
	FormulaAndFunctionProcessQueryDefinition            *FormulaAndFunctionProcessQueryDefinition
	FormulaAndFunctionApmDependencyStatsQueryDefinition *FormulaAndFunctionApmDependencyStatsQueryDefinition
	FormulaAndFunctionApmResourceStatsQueryDefinition   *FormulaAndFunctionApmResourceStatsQueryDefinition
	FormulaAndFunctionSLOQueryDefinition                *FormulaAndFunctionSLOQueryDefinition
	FormulaAndFunctionCloudCostQueryDefinition          *FormulaAndFunctionCloudCostQueryDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FormulaAndFunctionMetricQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionMetricQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionMetricQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionMetricQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionMetricQueryDefinition: v}
}

// FormulaAndFunctionEventQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionEventQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionEventQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionEventQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionEventQueryDefinition: v}
}

// FormulaAndFunctionProcessQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionProcessQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionProcessQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionProcessQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionProcessQueryDefinition: v}
}

// FormulaAndFunctionApmDependencyStatsQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionApmDependencyStatsQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionApmDependencyStatsQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionApmDependencyStatsQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionApmDependencyStatsQueryDefinition: v}
}

// FormulaAndFunctionApmResourceStatsQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionApmResourceStatsQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionApmResourceStatsQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionApmResourceStatsQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionApmResourceStatsQueryDefinition: v}
}

// FormulaAndFunctionSLOQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionSLOQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionSLOQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionSLOQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionSLOQueryDefinition: v}
}

// FormulaAndFunctionCloudCostQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionCloudCostQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionCloudCostQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionCloudCostQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionCloudCostQueryDefinition: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FormulaAndFunctionQueryDefinition) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into FormulaAndFunctionProcessQueryDefinition
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionProcessQueryDefinition)
	if err == nil {
		if obj.FormulaAndFunctionProcessQueryDefinition != nil && obj.FormulaAndFunctionProcessQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionProcessQueryDefinition, _ := datadog.Marshal(obj.FormulaAndFunctionProcessQueryDefinition)
			if string(jsonFormulaAndFunctionProcessQueryDefinition) == "{}" { // empty struct
				obj.FormulaAndFunctionProcessQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionProcessQueryDefinition = nil
		}
	} else {
		obj.FormulaAndFunctionProcessQueryDefinition = nil
	}

	// try to unmarshal data into FormulaAndFunctionApmDependencyStatsQueryDefinition
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionApmDependencyStatsQueryDefinition)
	if err == nil {
		if obj.FormulaAndFunctionApmDependencyStatsQueryDefinition != nil && obj.FormulaAndFunctionApmDependencyStatsQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionApmDependencyStatsQueryDefinition, _ := datadog.Marshal(obj.FormulaAndFunctionApmDependencyStatsQueryDefinition)
			if string(jsonFormulaAndFunctionApmDependencyStatsQueryDefinition) == "{}" { // empty struct
				obj.FormulaAndFunctionApmDependencyStatsQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionApmDependencyStatsQueryDefinition = nil
		}
	} else {
		obj.FormulaAndFunctionApmDependencyStatsQueryDefinition = nil
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

	// try to unmarshal data into FormulaAndFunctionSLOQueryDefinition
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionSLOQueryDefinition)
	if err == nil {
		if obj.FormulaAndFunctionSLOQueryDefinition != nil && obj.FormulaAndFunctionSLOQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionSLOQueryDefinition, _ := datadog.Marshal(obj.FormulaAndFunctionSLOQueryDefinition)
			if string(jsonFormulaAndFunctionSLOQueryDefinition) == "{}" { // empty struct
				obj.FormulaAndFunctionSLOQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionSLOQueryDefinition = nil
		}
	} else {
		obj.FormulaAndFunctionSLOQueryDefinition = nil
	}

	// try to unmarshal data into FormulaAndFunctionCloudCostQueryDefinition
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionCloudCostQueryDefinition)
	if err == nil {
		if obj.FormulaAndFunctionCloudCostQueryDefinition != nil && obj.FormulaAndFunctionCloudCostQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionCloudCostQueryDefinition, _ := datadog.Marshal(obj.FormulaAndFunctionCloudCostQueryDefinition)
			if string(jsonFormulaAndFunctionCloudCostQueryDefinition) == "{}" { // empty struct
				obj.FormulaAndFunctionCloudCostQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionCloudCostQueryDefinition = nil
		}
	} else {
		obj.FormulaAndFunctionCloudCostQueryDefinition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FormulaAndFunctionMetricQueryDefinition = nil
		obj.FormulaAndFunctionEventQueryDefinition = nil
		obj.FormulaAndFunctionProcessQueryDefinition = nil
		obj.FormulaAndFunctionApmDependencyStatsQueryDefinition = nil
		obj.FormulaAndFunctionApmResourceStatsQueryDefinition = nil
		obj.FormulaAndFunctionSLOQueryDefinition = nil
		obj.FormulaAndFunctionCloudCostQueryDefinition = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FormulaAndFunctionQueryDefinition) MarshalJSON() ([]byte, error) {
	if obj.FormulaAndFunctionMetricQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionMetricQueryDefinition)
	}

	if obj.FormulaAndFunctionEventQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionEventQueryDefinition)
	}

	if obj.FormulaAndFunctionProcessQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionProcessQueryDefinition)
	}

	if obj.FormulaAndFunctionApmDependencyStatsQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionApmDependencyStatsQueryDefinition)
	}

	if obj.FormulaAndFunctionApmResourceStatsQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionApmResourceStatsQueryDefinition)
	}

	if obj.FormulaAndFunctionSLOQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionSLOQueryDefinition)
	}

	if obj.FormulaAndFunctionCloudCostQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionCloudCostQueryDefinition)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FormulaAndFunctionQueryDefinition) GetActualInstance() interface{} {
	if obj.FormulaAndFunctionMetricQueryDefinition != nil {
		return obj.FormulaAndFunctionMetricQueryDefinition
	}

	if obj.FormulaAndFunctionEventQueryDefinition != nil {
		return obj.FormulaAndFunctionEventQueryDefinition
	}

	if obj.FormulaAndFunctionProcessQueryDefinition != nil {
		return obj.FormulaAndFunctionProcessQueryDefinition
	}

	if obj.FormulaAndFunctionApmDependencyStatsQueryDefinition != nil {
		return obj.FormulaAndFunctionApmDependencyStatsQueryDefinition
	}

	if obj.FormulaAndFunctionApmResourceStatsQueryDefinition != nil {
		return obj.FormulaAndFunctionApmResourceStatsQueryDefinition
	}

	if obj.FormulaAndFunctionSLOQueryDefinition != nil {
		return obj.FormulaAndFunctionSLOQueryDefinition
	}

	if obj.FormulaAndFunctionCloudCostQueryDefinition != nil {
		return obj.FormulaAndFunctionCloudCostQueryDefinition
	}

	// all schemas are nil
	return nil
}
