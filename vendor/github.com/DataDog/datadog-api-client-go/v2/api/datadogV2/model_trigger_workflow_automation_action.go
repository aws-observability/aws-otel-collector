// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TriggerWorkflowAutomationAction Triggers a Workflow Automation.
type TriggerWorkflowAutomationAction struct {
	// The handle of the Workflow Automation to trigger.
	Handle string `json:"handle"`
	// Indicates that the action triggers a Workflow Automation.
	Type TriggerWorkflowAutomationActionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTriggerWorkflowAutomationAction instantiates a new TriggerWorkflowAutomationAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTriggerWorkflowAutomationAction(handle string, typeVar TriggerWorkflowAutomationActionType) *TriggerWorkflowAutomationAction {
	this := TriggerWorkflowAutomationAction{}
	this.Handle = handle
	this.Type = typeVar
	return &this
}

// NewTriggerWorkflowAutomationActionWithDefaults instantiates a new TriggerWorkflowAutomationAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTriggerWorkflowAutomationActionWithDefaults() *TriggerWorkflowAutomationAction {
	this := TriggerWorkflowAutomationAction{}
	var typeVar TriggerWorkflowAutomationActionType = TRIGGERWORKFLOWAUTOMATIONACTIONTYPE_TRIGGER_WORKFLOW_AUTOMATION
	this.Type = typeVar
	return &this
}

// GetHandle returns the Handle field value.
func (o *TriggerWorkflowAutomationAction) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *TriggerWorkflowAutomationAction) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *TriggerWorkflowAutomationAction) SetHandle(v string) {
	o.Handle = v
}

// GetType returns the Type field value.
func (o *TriggerWorkflowAutomationAction) GetType() TriggerWorkflowAutomationActionType {
	if o == nil {
		var ret TriggerWorkflowAutomationActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerWorkflowAutomationAction) GetTypeOk() (*TriggerWorkflowAutomationActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TriggerWorkflowAutomationAction) SetType(v TriggerWorkflowAutomationActionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TriggerWorkflowAutomationAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["handle"] = o.Handle
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TriggerWorkflowAutomationAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle *string                              `json:"handle"`
		Type   *TriggerWorkflowAutomationActionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Handle = *all.Handle
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
