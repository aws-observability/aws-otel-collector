// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowTriggerWrapper Schema for a Workflow-based trigger.
type WorkflowTriggerWrapper struct {
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// Trigger a workflow from the Datadog UI. Only required if no other trigger exists.
	WorkflowTrigger interface{} `json:"workflowTrigger"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowTriggerWrapper instantiates a new WorkflowTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowTriggerWrapper(workflowTrigger interface{}) *WorkflowTriggerWrapper {
	this := WorkflowTriggerWrapper{}
	this.WorkflowTrigger = workflowTrigger
	return &this
}

// NewWorkflowTriggerWrapperWithDefaults instantiates a new WorkflowTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowTriggerWrapperWithDefaults() *WorkflowTriggerWrapper {
	this := WorkflowTriggerWrapper{}
	return &this
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *WorkflowTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *WorkflowTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *WorkflowTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// GetWorkflowTrigger returns the WorkflowTrigger field value.
func (o *WorkflowTriggerWrapper) GetWorkflowTrigger() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WorkflowTrigger
}

// GetWorkflowTriggerOk returns a tuple with the WorkflowTrigger field value
// and a boolean to check if the value has been set.
func (o *WorkflowTriggerWrapper) GetWorkflowTriggerOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowTrigger, true
}

// SetWorkflowTrigger sets field value.
func (o *WorkflowTriggerWrapper) SetWorkflowTrigger(v interface{}) {
	o.WorkflowTrigger = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}
	toSerialize["workflowTrigger"] = o.WorkflowTrigger

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StartStepNames  []string     `json:"startStepNames,omitempty"`
		WorkflowTrigger *interface{} `json:"workflowTrigger"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.WorkflowTrigger == nil {
		return fmt.Errorf("required field workflowTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"startStepNames", "workflowTrigger"})
	} else {
		return err
	}
	o.StartStepNames = all.StartStepNames
	o.WorkflowTrigger = *all.WorkflowTrigger

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
