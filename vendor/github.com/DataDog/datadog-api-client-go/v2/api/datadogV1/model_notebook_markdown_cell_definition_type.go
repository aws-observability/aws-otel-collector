// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookMarkdownCellDefinitionType Type of the markdown cell.
type NotebookMarkdownCellDefinitionType string

// List of NotebookMarkdownCellDefinitionType.
const (
	NOTEBOOKMARKDOWNCELLDEFINITIONTYPE_MARKDOWN NotebookMarkdownCellDefinitionType = "markdown"
)

var allowedNotebookMarkdownCellDefinitionTypeEnumValues = []NotebookMarkdownCellDefinitionType{
	NOTEBOOKMARKDOWNCELLDEFINITIONTYPE_MARKDOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotebookMarkdownCellDefinitionType) GetAllowedValues() []NotebookMarkdownCellDefinitionType {
	return allowedNotebookMarkdownCellDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotebookMarkdownCellDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotebookMarkdownCellDefinitionType(value)
	return nil
}

// NewNotebookMarkdownCellDefinitionTypeFromValue returns a pointer to a valid NotebookMarkdownCellDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotebookMarkdownCellDefinitionTypeFromValue(v string) (*NotebookMarkdownCellDefinitionType, error) {
	ev := NotebookMarkdownCellDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotebookMarkdownCellDefinitionType: valid values are %v", v, allowedNotebookMarkdownCellDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotebookMarkdownCellDefinitionType) IsValid() bool {
	for _, existing := range allowedNotebookMarkdownCellDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotebookMarkdownCellDefinitionType value.
func (v NotebookMarkdownCellDefinitionType) Ptr() *NotebookMarkdownCellDefinitionType {
	return &v
}
