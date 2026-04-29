// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionEventQueryGroupByConfig - Group by configuration for a formula and functions events query. Accepts either a list of facet objects or a flat object that specifies a list of facet fields.
type FormulaAndFunctionEventQueryGroupByConfig struct {
	FormulaAndFunctionEventQueryGroupByList   *[]FormulaAndFunctionEventQueryGroupBy
	FormulaAndFunctionEventQueryGroupByFields *FormulaAndFunctionEventQueryGroupByFields

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FormulaAndFunctionEventQueryGroupByListAsFormulaAndFunctionEventQueryGroupByConfig is a convenience function that returns []FormulaAndFunctionEventQueryGroupBy wrapped in FormulaAndFunctionEventQueryGroupByConfig.
func FormulaAndFunctionEventQueryGroupByListAsFormulaAndFunctionEventQueryGroupByConfig(v *[]FormulaAndFunctionEventQueryGroupBy) FormulaAndFunctionEventQueryGroupByConfig {
	return FormulaAndFunctionEventQueryGroupByConfig{FormulaAndFunctionEventQueryGroupByList: v}
}

// FormulaAndFunctionEventQueryGroupByFieldsAsFormulaAndFunctionEventQueryGroupByConfig is a convenience function that returns FormulaAndFunctionEventQueryGroupByFields wrapped in FormulaAndFunctionEventQueryGroupByConfig.
func FormulaAndFunctionEventQueryGroupByFieldsAsFormulaAndFunctionEventQueryGroupByConfig(v *FormulaAndFunctionEventQueryGroupByFields) FormulaAndFunctionEventQueryGroupByConfig {
	return FormulaAndFunctionEventQueryGroupByConfig{FormulaAndFunctionEventQueryGroupByFields: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FormulaAndFunctionEventQueryGroupByConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FormulaAndFunctionEventQueryGroupByList
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionEventQueryGroupByList)
	if err == nil {
		if obj.FormulaAndFunctionEventQueryGroupByList != nil {
			jsonFormulaAndFunctionEventQueryGroupByList, _ := datadog.Marshal(obj.FormulaAndFunctionEventQueryGroupByList)
			if string(jsonFormulaAndFunctionEventQueryGroupByList) == "{}" && string(data) != "{}" { // empty struct
				obj.FormulaAndFunctionEventQueryGroupByList = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionEventQueryGroupByList = nil
		}
	} else {
		obj.FormulaAndFunctionEventQueryGroupByList = nil
	}

	// try to unmarshal data into FormulaAndFunctionEventQueryGroupByFields
	err = datadog.Unmarshal(data, &obj.FormulaAndFunctionEventQueryGroupByFields)
	if err == nil {
		if obj.FormulaAndFunctionEventQueryGroupByFields != nil && obj.FormulaAndFunctionEventQueryGroupByFields.UnparsedObject == nil {
			jsonFormulaAndFunctionEventQueryGroupByFields, _ := datadog.Marshal(obj.FormulaAndFunctionEventQueryGroupByFields)
			if string(jsonFormulaAndFunctionEventQueryGroupByFields) == "{}" { // empty struct
				obj.FormulaAndFunctionEventQueryGroupByFields = nil
			} else {
				match++
			}
		} else {
			obj.FormulaAndFunctionEventQueryGroupByFields = nil
		}
	} else {
		obj.FormulaAndFunctionEventQueryGroupByFields = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FormulaAndFunctionEventQueryGroupByList = nil
		obj.FormulaAndFunctionEventQueryGroupByFields = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FormulaAndFunctionEventQueryGroupByConfig) MarshalJSON() ([]byte, error) {
	if obj.FormulaAndFunctionEventQueryGroupByList != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionEventQueryGroupByList)
	}

	if obj.FormulaAndFunctionEventQueryGroupByFields != nil {
		return datadog.Marshal(&obj.FormulaAndFunctionEventQueryGroupByFields)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FormulaAndFunctionEventQueryGroupByConfig) GetActualInstance() interface{} {
	if obj.FormulaAndFunctionEventQueryGroupByList != nil {
		return obj.FormulaAndFunctionEventQueryGroupByList
	}

	if obj.FormulaAndFunctionEventQueryGroupByFields != nil {
		return obj.FormulaAndFunctionEventQueryGroupByFields
	}

	// all schemas are nil
	return nil
}
