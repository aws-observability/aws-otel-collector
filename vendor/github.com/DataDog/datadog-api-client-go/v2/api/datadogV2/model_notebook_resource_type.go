// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookResourceType Notebook resource type
type NotebookResourceType string

// List of NotebookResourceType.
const (
	NOTEBOOKRESOURCETYPE_NOTEBOOK NotebookResourceType = "notebook"
)

var allowedNotebookResourceTypeEnumValues = []NotebookResourceType{
	NOTEBOOKRESOURCETYPE_NOTEBOOK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotebookResourceType) GetAllowedValues() []NotebookResourceType {
	return allowedNotebookResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotebookResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotebookResourceType(value)
	return nil
}

// NewNotebookResourceTypeFromValue returns a pointer to a valid NotebookResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotebookResourceTypeFromValue(v string) (*NotebookResourceType, error) {
	ev := NotebookResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotebookResourceType: valid values are %v", v, allowedNotebookResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotebookResourceType) IsValid() bool {
	for _, existing := range allowedNotebookResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotebookResourceType value.
func (v NotebookResourceType) Ptr() *NotebookResourceType {
	return &v
}
