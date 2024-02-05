// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookUpdateCell - Updating a notebook can either insert new cell(s) or update existing cell(s) by including the cell `id`.
// To delete existing cell(s), simply omit it from the list of cells.
type NotebookUpdateCell struct {
	NotebookCellCreateRequest *NotebookCellCreateRequest
	NotebookCellUpdateRequest *NotebookCellUpdateRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NotebookCellCreateRequestAsNotebookUpdateCell is a convenience function that returns NotebookCellCreateRequest wrapped in NotebookUpdateCell.
func NotebookCellCreateRequestAsNotebookUpdateCell(v *NotebookCellCreateRequest) NotebookUpdateCell {
	return NotebookUpdateCell{NotebookCellCreateRequest: v}
}

// NotebookCellUpdateRequestAsNotebookUpdateCell is a convenience function that returns NotebookCellUpdateRequest wrapped in NotebookUpdateCell.
func NotebookCellUpdateRequestAsNotebookUpdateCell(v *NotebookCellUpdateRequest) NotebookUpdateCell {
	return NotebookUpdateCell{NotebookCellUpdateRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *NotebookUpdateCell) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotebookCellCreateRequest
	err = datadog.Unmarshal(data, &obj.NotebookCellCreateRequest)
	if err == nil {
		if obj.NotebookCellCreateRequest != nil && obj.NotebookCellCreateRequest.UnparsedObject == nil {
			jsonNotebookCellCreateRequest, _ := datadog.Marshal(obj.NotebookCellCreateRequest)
			if string(jsonNotebookCellCreateRequest) == "{}" { // empty struct
				obj.NotebookCellCreateRequest = nil
			} else {
				match++
			}
		} else {
			obj.NotebookCellCreateRequest = nil
		}
	} else {
		obj.NotebookCellCreateRequest = nil
	}

	// try to unmarshal data into NotebookCellUpdateRequest
	err = datadog.Unmarshal(data, &obj.NotebookCellUpdateRequest)
	if err == nil {
		if obj.NotebookCellUpdateRequest != nil && obj.NotebookCellUpdateRequest.UnparsedObject == nil {
			jsonNotebookCellUpdateRequest, _ := datadog.Marshal(obj.NotebookCellUpdateRequest)
			if string(jsonNotebookCellUpdateRequest) == "{}" { // empty struct
				obj.NotebookCellUpdateRequest = nil
			} else {
				match++
			}
		} else {
			obj.NotebookCellUpdateRequest = nil
		}
	} else {
		obj.NotebookCellUpdateRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NotebookCellCreateRequest = nil
		obj.NotebookCellUpdateRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj NotebookUpdateCell) MarshalJSON() ([]byte, error) {
	if obj.NotebookCellCreateRequest != nil {
		return datadog.Marshal(&obj.NotebookCellCreateRequest)
	}

	if obj.NotebookCellUpdateRequest != nil {
		return datadog.Marshal(&obj.NotebookCellUpdateRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *NotebookUpdateCell) GetActualInstance() interface{} {
	if obj.NotebookCellCreateRequest != nil {
		return obj.NotebookCellCreateRequest
	}

	if obj.NotebookCellUpdateRequest != nil {
		return obj.NotebookCellUpdateRequest
	}

	// all schemas are nil
	return nil
}
