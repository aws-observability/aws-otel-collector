// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Query - A data query used by an app. This can take the form of an external action, a data transformation, or a state variable.
type Query struct {
	ActionQuery   *ActionQuery
	DataTransform *DataTransform
	StateVariable *StateVariable

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ActionQueryAsQuery is a convenience function that returns ActionQuery wrapped in Query.
func ActionQueryAsQuery(v *ActionQuery) Query {
	return Query{ActionQuery: v}
}

// DataTransformAsQuery is a convenience function that returns DataTransform wrapped in Query.
func DataTransformAsQuery(v *DataTransform) Query {
	return Query{DataTransform: v}
}

// StateVariableAsQuery is a convenience function that returns StateVariable wrapped in Query.
func StateVariableAsQuery(v *StateVariable) Query {
	return Query{StateVariable: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *Query) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActionQuery
	err = datadog.Unmarshal(data, &obj.ActionQuery)
	if err == nil {
		if obj.ActionQuery != nil && obj.ActionQuery.UnparsedObject == nil {
			jsonActionQuery, _ := datadog.Marshal(obj.ActionQuery)
			if string(jsonActionQuery) == "{}" { // empty struct
				obj.ActionQuery = nil
			} else {
				match++
			}
		} else {
			obj.ActionQuery = nil
		}
	} else {
		obj.ActionQuery = nil
	}

	// try to unmarshal data into DataTransform
	err = datadog.Unmarshal(data, &obj.DataTransform)
	if err == nil {
		if obj.DataTransform != nil && obj.DataTransform.UnparsedObject == nil {
			jsonDataTransform, _ := datadog.Marshal(obj.DataTransform)
			if string(jsonDataTransform) == "{}" { // empty struct
				obj.DataTransform = nil
			} else {
				match++
			}
		} else {
			obj.DataTransform = nil
		}
	} else {
		obj.DataTransform = nil
	}

	// try to unmarshal data into StateVariable
	err = datadog.Unmarshal(data, &obj.StateVariable)
	if err == nil {
		if obj.StateVariable != nil && obj.StateVariable.UnparsedObject == nil {
			jsonStateVariable, _ := datadog.Marshal(obj.StateVariable)
			if string(jsonStateVariable) == "{}" { // empty struct
				obj.StateVariable = nil
			} else {
				match++
			}
		} else {
			obj.StateVariable = nil
		}
	} else {
		obj.StateVariable = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ActionQuery = nil
		obj.DataTransform = nil
		obj.StateVariable = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj Query) MarshalJSON() ([]byte, error) {
	if obj.ActionQuery != nil {
		return datadog.Marshal(&obj.ActionQuery)
	}

	if obj.DataTransform != nil {
		return datadog.Marshal(&obj.DataTransform)
	}

	if obj.StateVariable != nil {
		return datadog.Marshal(&obj.StateVariable)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *Query) GetActualInstance() interface{} {
	if obj.ActionQuery != nil {
		return obj.ActionQuery
	}

	if obj.DataTransform != nil {
		return obj.DataTransform
	}

	if obj.StateVariable != nil {
		return obj.StateVariable
	}

	// all schemas are nil
	return nil
}
