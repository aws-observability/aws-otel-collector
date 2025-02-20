// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScalarColumn - A single column in a scalar query response.
type ScalarColumn struct {
	GroupScalarColumn *GroupScalarColumn
	DataScalarColumn  *DataScalarColumn

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GroupScalarColumnAsScalarColumn is a convenience function that returns GroupScalarColumn wrapped in ScalarColumn.
func GroupScalarColumnAsScalarColumn(v *GroupScalarColumn) ScalarColumn {
	return ScalarColumn{GroupScalarColumn: v}
}

// DataScalarColumnAsScalarColumn is a convenience function that returns DataScalarColumn wrapped in ScalarColumn.
func DataScalarColumnAsScalarColumn(v *DataScalarColumn) ScalarColumn {
	return ScalarColumn{DataScalarColumn: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ScalarColumn) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GroupScalarColumn
	err = datadog.Unmarshal(data, &obj.GroupScalarColumn)
	if err == nil {
		if obj.GroupScalarColumn != nil && obj.GroupScalarColumn.UnparsedObject == nil {
			jsonGroupScalarColumn, _ := datadog.Marshal(obj.GroupScalarColumn)
			if string(jsonGroupScalarColumn) == "{}" && string(data) != "{}" { // empty struct
				obj.GroupScalarColumn = nil
			} else {
				match++
			}
		} else {
			obj.GroupScalarColumn = nil
		}
	} else {
		obj.GroupScalarColumn = nil
	}

	// try to unmarshal data into DataScalarColumn
	err = datadog.Unmarshal(data, &obj.DataScalarColumn)
	if err == nil {
		if obj.DataScalarColumn != nil && obj.DataScalarColumn.UnparsedObject == nil {
			jsonDataScalarColumn, _ := datadog.Marshal(obj.DataScalarColumn)
			if string(jsonDataScalarColumn) == "{}" && string(data) != "{}" { // empty struct
				obj.DataScalarColumn = nil
			} else {
				match++
			}
		} else {
			obj.DataScalarColumn = nil
		}
	} else {
		obj.DataScalarColumn = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GroupScalarColumn = nil
		obj.DataScalarColumn = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ScalarColumn) MarshalJSON() ([]byte, error) {
	if obj.GroupScalarColumn != nil {
		return datadog.Marshal(&obj.GroupScalarColumn)
	}

	if obj.DataScalarColumn != nil {
		return datadog.Marshal(&obj.DataScalarColumn)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ScalarColumn) GetActualInstance() interface{} {
	if obj.GroupScalarColumn != nil {
		return obj.GroupScalarColumn
	}

	if obj.DataScalarColumn != nil {
		return obj.DataScalarColumn
	}

	// all schemas are nil
	return nil
}
