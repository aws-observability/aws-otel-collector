// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionQueryDefinition - A formula and function query.
type FormulaAndFunctionQueryDefinition struct {
	FormulaAndFunctionMetricQueryDefinition                   *FormulaAndFunctionMetricQueryDefinition
	FormulaAndFunctionEventQueryDefinition                    *FormulaAndFunctionEventQueryDefinition
	FormulaAndFunctionProcessQueryDefinition                  *FormulaAndFunctionProcessQueryDefinition
	FormulaAndFunctionApmDependencyStatsQueryDefinition       *FormulaAndFunctionApmDependencyStatsQueryDefinition
	FormulaAndFunctionApmResourceStatsQueryDefinition         *FormulaAndFunctionApmResourceStatsQueryDefinition
	FormulaAndFunctionApmMetricsQueryDefinition               *FormulaAndFunctionApmMetricsQueryDefinition
	FormulaAndFunctionSLOQueryDefinition                      *FormulaAndFunctionSLOQueryDefinition
	FormulaAndFunctionCloudCostQueryDefinition                *FormulaAndFunctionCloudCostQueryDefinition
	FormulaAndFunctionProductAnalyticsExtendedQueryDefinition *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition
	FormulaAndFunctionUserJourneyQueryDefinition              *FormulaAndFunctionUserJourneyQueryDefinition
	FormulaAndFunctionRetentionQueryDefinition                *FormulaAndFunctionRetentionQueryDefinition

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

// FormulaAndFunctionApmMetricsQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionApmMetricsQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionApmMetricsQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionApmMetricsQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionApmMetricsQueryDefinition: v}
}

// FormulaAndFunctionSLOQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionSLOQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionSLOQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionSLOQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionSLOQueryDefinition: v}
}

// FormulaAndFunctionCloudCostQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionCloudCostQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionCloudCostQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionCloudCostQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionCloudCostQueryDefinition: v}
}

// FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionProductAnalyticsExtendedQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionProductAnalyticsExtendedQueryDefinition: v}
}

// FormulaAndFunctionUserJourneyQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionUserJourneyQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionUserJourneyQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionUserJourneyQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionUserJourneyQueryDefinition: v}
}

// FormulaAndFunctionRetentionQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionRetentionQueryDefinition wrapped in FormulaAndFunctionQueryDefinition.
func FormulaAndFunctionRetentionQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionRetentionQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionRetentionQueryDefinition: v}
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

	// try to unmarshal data into FormulaAndFunctionApmMetricsQueryDefinition
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionApmMetricsQueryDefinition)
	if err == nil {
		if obj.FormulaAndFunctionApmMetricsQueryDefinition != nil && obj.FormulaAndFunctionApmMetricsQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionApmMetricsQueryDefinition, _ := datadog.Marshal(obj.FormulaAndFunctionApmMetricsQueryDefinition)
			if string(jsonFormulaAndFunctionApmMetricsQueryDefinition) == "{}" { // empty struct
				obj.FormulaAndFunctionApmMetricsQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionApmMetricsQueryDefinition = nil
		}
	} else {
		obj.FormulaAndFunctionApmMetricsQueryDefinition = nil
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

	// try to unmarshal data into FormulaAndFunctionProductAnalyticsExtendedQueryDefinition
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition)
	if err == nil {
		if obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition != nil && obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionProductAnalyticsExtendedQueryDefinition, _ := datadog.Marshal(obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition)
			if string(jsonFormulaAndFunctionProductAnalyticsExtendedQueryDefinition) == "{}" { // empty struct
				obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition = nil
		}
	} else {
		obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition = nil
	}

	// try to unmarshal data into FormulaAndFunctionUserJourneyQueryDefinition
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionUserJourneyQueryDefinition)
	if err == nil {
		if obj.FormulaAndFunctionUserJourneyQueryDefinition != nil && obj.FormulaAndFunctionUserJourneyQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionUserJourneyQueryDefinition, _ := datadog.Marshal(obj.FormulaAndFunctionUserJourneyQueryDefinition)
			if string(jsonFormulaAndFunctionUserJourneyQueryDefinition) == "{}" { // empty struct
				obj.FormulaAndFunctionUserJourneyQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionUserJourneyQueryDefinition = nil
		}
	} else {
		obj.FormulaAndFunctionUserJourneyQueryDefinition = nil
	}

	// try to unmarshal data into FormulaAndFunctionRetentionQueryDefinition
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionRetentionQueryDefinition)
	if err == nil {
		if obj.FormulaAndFunctionRetentionQueryDefinition != nil && obj.FormulaAndFunctionRetentionQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionRetentionQueryDefinition, _ := datadog.Marshal(obj.FormulaAndFunctionRetentionQueryDefinition)
			if string(jsonFormulaAndFunctionRetentionQueryDefinition) == "{}" { // empty struct
				obj.FormulaAndFunctionRetentionQueryDefinition = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionRetentionQueryDefinition = nil
		}
	} else {
		obj.FormulaAndFunctionRetentionQueryDefinition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FormulaAndFunctionMetricQueryDefinition = nil
		obj.FormulaAndFunctionEventQueryDefinition = nil
		obj.FormulaAndFunctionProcessQueryDefinition = nil
		obj.FormulaAndFunctionApmDependencyStatsQueryDefinition = nil
		obj.FormulaAndFunctionApmResourceStatsQueryDefinition = nil
		obj.FormulaAndFunctionApmMetricsQueryDefinition = nil
		obj.FormulaAndFunctionSLOQueryDefinition = nil
		obj.FormulaAndFunctionCloudCostQueryDefinition = nil
		obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition = nil
		obj.FormulaAndFunctionUserJourneyQueryDefinition = nil
		obj.FormulaAndFunctionRetentionQueryDefinition = nil
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

	if obj.FormulaAndFunctionApmMetricsQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionApmMetricsQueryDefinition)
	}

	if obj.FormulaAndFunctionSLOQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionSLOQueryDefinition)
	}

	if obj.FormulaAndFunctionCloudCostQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionCloudCostQueryDefinition)
	}

	if obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition)
	}

	if obj.FormulaAndFunctionUserJourneyQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionUserJourneyQueryDefinition)
	}

	if obj.FormulaAndFunctionRetentionQueryDefinition != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionRetentionQueryDefinition)
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

	if obj.FormulaAndFunctionApmMetricsQueryDefinition != nil {
		return obj.FormulaAndFunctionApmMetricsQueryDefinition
	}

	if obj.FormulaAndFunctionSLOQueryDefinition != nil {
		return obj.FormulaAndFunctionSLOQueryDefinition
	}

	if obj.FormulaAndFunctionCloudCostQueryDefinition != nil {
		return obj.FormulaAndFunctionCloudCostQueryDefinition
	}

	if obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition != nil {
		return obj.FormulaAndFunctionProductAnalyticsExtendedQueryDefinition
	}

	if obj.FormulaAndFunctionUserJourneyQueryDefinition != nil {
		return obj.FormulaAndFunctionUserJourneyQueryDefinition
	}

	if obj.FormulaAndFunctionRetentionQueryDefinition != nil {
		return obj.FormulaAndFunctionRetentionQueryDefinition
	}

	// all schemas are nil
	return nil
}
