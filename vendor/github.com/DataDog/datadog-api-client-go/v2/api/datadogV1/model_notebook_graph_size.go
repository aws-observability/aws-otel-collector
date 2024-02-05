// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookGraphSize The size of the graph.
type NotebookGraphSize string

// List of NotebookGraphSize.
const (
	NOTEBOOKGRAPHSIZE_EXTRA_SMALL NotebookGraphSize = "xs"
	NOTEBOOKGRAPHSIZE_SMALL       NotebookGraphSize = "s"
	NOTEBOOKGRAPHSIZE_MEDIUM      NotebookGraphSize = "m"
	NOTEBOOKGRAPHSIZE_LARGE       NotebookGraphSize = "l"
	NOTEBOOKGRAPHSIZE_EXTRA_LARGE NotebookGraphSize = "xl"
)

var allowedNotebookGraphSizeEnumValues = []NotebookGraphSize{
	NOTEBOOKGRAPHSIZE_EXTRA_SMALL,
	NOTEBOOKGRAPHSIZE_SMALL,
	NOTEBOOKGRAPHSIZE_MEDIUM,
	NOTEBOOKGRAPHSIZE_LARGE,
	NOTEBOOKGRAPHSIZE_EXTRA_LARGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotebookGraphSize) GetAllowedValues() []NotebookGraphSize {
	return allowedNotebookGraphSizeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotebookGraphSize) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotebookGraphSize(value)
	return nil
}

// NewNotebookGraphSizeFromValue returns a pointer to a valid NotebookGraphSize
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotebookGraphSizeFromValue(v string) (*NotebookGraphSize, error) {
	ev := NotebookGraphSize(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotebookGraphSize: valid values are %v", v, allowedNotebookGraphSizeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotebookGraphSize) IsValid() bool {
	for _, existing := range allowedNotebookGraphSizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotebookGraphSize value.
func (v NotebookGraphSize) Ptr() *NotebookGraphSize {
	return &v
}
