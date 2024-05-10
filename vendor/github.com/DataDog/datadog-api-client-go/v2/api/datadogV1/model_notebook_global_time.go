// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookGlobalTime - Notebook global timeframe.
type NotebookGlobalTime struct {
	NotebookRelativeTime *NotebookRelativeTime
	NotebookAbsoluteTime *NotebookAbsoluteTime

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NotebookRelativeTimeAsNotebookGlobalTime is a convenience function that returns NotebookRelativeTime wrapped in NotebookGlobalTime.
func NotebookRelativeTimeAsNotebookGlobalTime(v *NotebookRelativeTime) NotebookGlobalTime {
	return NotebookGlobalTime{NotebookRelativeTime: v}
}

// NotebookAbsoluteTimeAsNotebookGlobalTime is a convenience function that returns NotebookAbsoluteTime wrapped in NotebookGlobalTime.
func NotebookAbsoluteTimeAsNotebookGlobalTime(v *NotebookAbsoluteTime) NotebookGlobalTime {
	return NotebookGlobalTime{NotebookAbsoluteTime: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *NotebookGlobalTime) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotebookRelativeTime
	err = datadog.Unmarshal(data, &obj.NotebookRelativeTime)
	if err == nil {
		if obj.NotebookRelativeTime != nil && obj.NotebookRelativeTime.UnparsedObject == nil {
			jsonNotebookRelativeTime, _ := datadog.Marshal(obj.NotebookRelativeTime)
			if string(jsonNotebookRelativeTime) == "{}" { // empty struct
				obj.NotebookRelativeTime = nil
			} else {
				match++
			}
		} else {
			obj.NotebookRelativeTime = nil
		}
	} else {
		obj.NotebookRelativeTime = nil
	}

	// try to unmarshal data into NotebookAbsoluteTime
	err = datadog.Unmarshal(data, &obj.NotebookAbsoluteTime)
	if err == nil {
		if obj.NotebookAbsoluteTime != nil && obj.NotebookAbsoluteTime.UnparsedObject == nil {
			jsonNotebookAbsoluteTime, _ := datadog.Marshal(obj.NotebookAbsoluteTime)
			if string(jsonNotebookAbsoluteTime) == "{}" { // empty struct
				obj.NotebookAbsoluteTime = nil
			} else {
				match++
			}
		} else {
			obj.NotebookAbsoluteTime = nil
		}
	} else {
		obj.NotebookAbsoluteTime = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NotebookRelativeTime = nil
		obj.NotebookAbsoluteTime = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj NotebookGlobalTime) MarshalJSON() ([]byte, error) {
	if obj.NotebookRelativeTime != nil {
		return datadog.Marshal(&obj.NotebookRelativeTime)
	}

	if obj.NotebookAbsoluteTime != nil {
		return datadog.Marshal(&obj.NotebookAbsoluteTime)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *NotebookGlobalTime) GetActualInstance() interface{} {
	if obj.NotebookRelativeTime != nil {
		return obj.NotebookRelativeTime
	}

	if obj.NotebookAbsoluteTime != nil {
		return obj.NotebookAbsoluteTime
	}

	// all schemas are nil
	return nil
}
