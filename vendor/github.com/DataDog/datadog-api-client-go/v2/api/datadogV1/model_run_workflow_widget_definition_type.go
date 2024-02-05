// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunWorkflowWidgetDefinitionType Type of the run workflow widget.
type RunWorkflowWidgetDefinitionType string

// List of RunWorkflowWidgetDefinitionType.
const (
	RUNWORKFLOWWIDGETDEFINITIONTYPE_RUN_WORKFLOW RunWorkflowWidgetDefinitionType = "run_workflow"
)

var allowedRunWorkflowWidgetDefinitionTypeEnumValues = []RunWorkflowWidgetDefinitionType{
	RUNWORKFLOWWIDGETDEFINITIONTYPE_RUN_WORKFLOW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RunWorkflowWidgetDefinitionType) GetAllowedValues() []RunWorkflowWidgetDefinitionType {
	return allowedRunWorkflowWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RunWorkflowWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RunWorkflowWidgetDefinitionType(value)
	return nil
}

// NewRunWorkflowWidgetDefinitionTypeFromValue returns a pointer to a valid RunWorkflowWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRunWorkflowWidgetDefinitionTypeFromValue(v string) (*RunWorkflowWidgetDefinitionType, error) {
	ev := RunWorkflowWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RunWorkflowWidgetDefinitionType: valid values are %v", v, allowedRunWorkflowWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RunWorkflowWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedRunWorkflowWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunWorkflowWidgetDefinitionType value.
func (v RunWorkflowWidgetDefinitionType) Ptr() *RunWorkflowWidgetDefinitionType {
	return &v
}
