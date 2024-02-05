// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookCellResourceType Type of the Notebook Cell resource.
type NotebookCellResourceType string

// List of NotebookCellResourceType.
const (
	NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS NotebookCellResourceType = "notebook_cells"
)

var allowedNotebookCellResourceTypeEnumValues = []NotebookCellResourceType{
	NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotebookCellResourceType) GetAllowedValues() []NotebookCellResourceType {
	return allowedNotebookCellResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotebookCellResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotebookCellResourceType(value)
	return nil
}

// NewNotebookCellResourceTypeFromValue returns a pointer to a valid NotebookCellResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotebookCellResourceTypeFromValue(v string) (*NotebookCellResourceType, error) {
	ev := NotebookCellResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotebookCellResourceType: valid values are %v", v, allowedNotebookCellResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotebookCellResourceType) IsValid() bool {
	for _, existing := range allowedNotebookCellResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotebookCellResourceType value.
func (v NotebookCellResourceType) Ptr() *NotebookCellResourceType {
	return &v
}
