// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowDataType The definition of `WorkflowDataType` object.
type WorkflowDataType string

// List of WorkflowDataType.
const (
	WORKFLOWDATATYPE_WORKFLOWS WorkflowDataType = "workflows"
)

var allowedWorkflowDataTypeEnumValues = []WorkflowDataType{
	WORKFLOWDATATYPE_WORKFLOWS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WorkflowDataType) GetAllowedValues() []WorkflowDataType {
	return allowedWorkflowDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowDataType(value)
	return nil
}

// NewWorkflowDataTypeFromValue returns a pointer to a valid WorkflowDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowDataTypeFromValue(v string) (*WorkflowDataType, error) {
	ev := WorkflowDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowDataType: valid values are %v", v, allowedWorkflowDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowDataType) IsValid() bool {
	for _, existing := range allowedWorkflowDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowDataType value.
func (v WorkflowDataType) Ptr() *WorkflowDataType {
	return &v
}
