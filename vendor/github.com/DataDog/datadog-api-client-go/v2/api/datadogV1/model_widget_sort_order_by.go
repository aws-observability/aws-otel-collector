// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetSortOrderBy - The item to sort the widget by.
type WidgetSortOrderBy struct {
	WidgetFormulaSort *WidgetFormulaSort
	WidgetGroupSort   *WidgetGroupSort

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// WidgetFormulaSortAsWidgetSortOrderBy is a convenience function that returns WidgetFormulaSort wrapped in WidgetSortOrderBy.
func WidgetFormulaSortAsWidgetSortOrderBy(v *WidgetFormulaSort) WidgetSortOrderBy {
	return WidgetSortOrderBy{WidgetFormulaSort: v}
}

// WidgetGroupSortAsWidgetSortOrderBy is a convenience function that returns WidgetGroupSort wrapped in WidgetSortOrderBy.
func WidgetGroupSortAsWidgetSortOrderBy(v *WidgetGroupSort) WidgetSortOrderBy {
	return WidgetSortOrderBy{WidgetGroupSort: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *WidgetSortOrderBy) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into WidgetFormulaSort
	err = datadog.Unmarshal(data, &obj.WidgetFormulaSort)
	if err == nil {
		if obj.WidgetFormulaSort != nil && obj.WidgetFormulaSort.UnparsedObject == nil {
			jsonWidgetFormulaSort, _ := datadog.Marshal(obj.WidgetFormulaSort)
			if string(jsonWidgetFormulaSort) == "{}" { // empty struct
				obj.WidgetFormulaSort = nil
			} else {
				match++
			}
		} else {
			obj.WidgetFormulaSort = nil
		}
	} else {
		obj.WidgetFormulaSort = nil
	}

	// try to unmarshal data into WidgetGroupSort
	err = datadog.Unmarshal(data, &obj.WidgetGroupSort)
	if err == nil {
		if obj.WidgetGroupSort != nil && obj.WidgetGroupSort.UnparsedObject == nil {
			jsonWidgetGroupSort, _ := datadog.Marshal(obj.WidgetGroupSort)
			if string(jsonWidgetGroupSort) == "{}" { // empty struct
				obj.WidgetGroupSort = nil
			} else {
				match++
			}
		} else {
			obj.WidgetGroupSort = nil
		}
	} else {
		obj.WidgetGroupSort = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.WidgetFormulaSort = nil
		obj.WidgetGroupSort = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj WidgetSortOrderBy) MarshalJSON() ([]byte, error) {
	if obj.WidgetFormulaSort != nil {
		return datadog.Marshal(&obj.WidgetFormulaSort)
	}

	if obj.WidgetGroupSort != nil {
		return datadog.Marshal(&obj.WidgetGroupSort)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *WidgetSortOrderBy) GetActualInstance() interface{} {
	if obj.WidgetFormulaSort != nil {
		return obj.WidgetFormulaSort
	}

	if obj.WidgetGroupSort != nil {
		return obj.WidgetGroupSort
	}

	// all schemas are nil
	return nil
}
