// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookStatus Publication status of the notebook. For now, always "published".
type NotebookStatus string

// List of NotebookStatus.
const (
	NOTEBOOKSTATUS_PUBLISHED NotebookStatus = "published"
)

var allowedNotebookStatusEnumValues = []NotebookStatus{
	NOTEBOOKSTATUS_PUBLISHED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotebookStatus) GetAllowedValues() []NotebookStatus {
	return allowedNotebookStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotebookStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotebookStatus(value)
	return nil
}

// NewNotebookStatusFromValue returns a pointer to a valid NotebookStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotebookStatusFromValue(v string) (*NotebookStatus, error) {
	ev := NotebookStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotebookStatus: valid values are %v", v, allowedNotebookStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotebookStatus) IsValid() bool {
	for _, existing := range allowedNotebookStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotebookStatus value.
func (v NotebookStatus) Ptr() *NotebookStatus {
	return &v
}
