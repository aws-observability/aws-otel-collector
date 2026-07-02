// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TriggerWorkflowAutomationActionType Indicates that the action triggers a Workflow Automation.
type TriggerWorkflowAutomationActionType string

// List of TriggerWorkflowAutomationActionType.
const (
	TRIGGERWORKFLOWAUTOMATIONACTIONTYPE_TRIGGER_WORKFLOW_AUTOMATION TriggerWorkflowAutomationActionType = "workflow"
)

var allowedTriggerWorkflowAutomationActionTypeEnumValues = []TriggerWorkflowAutomationActionType{
	TRIGGERWORKFLOWAUTOMATIONACTIONTYPE_TRIGGER_WORKFLOW_AUTOMATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TriggerWorkflowAutomationActionType) GetAllowedValues() []TriggerWorkflowAutomationActionType {
	return allowedTriggerWorkflowAutomationActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TriggerWorkflowAutomationActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TriggerWorkflowAutomationActionType(value)
	return nil
}

// NewTriggerWorkflowAutomationActionTypeFromValue returns a pointer to a valid TriggerWorkflowAutomationActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTriggerWorkflowAutomationActionTypeFromValue(v string) (*TriggerWorkflowAutomationActionType, error) {
	ev := TriggerWorkflowAutomationActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TriggerWorkflowAutomationActionType: valid values are %v", v, allowedTriggerWorkflowAutomationActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TriggerWorkflowAutomationActionType) IsValid() bool {
	for _, existing := range allowedTriggerWorkflowAutomationActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerWorkflowAutomationActionType value.
func (v TriggerWorkflowAutomationActionType) Ptr() *TriggerWorkflowAutomationActionType {
	return &v
}
