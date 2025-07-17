// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookTriggerWrapper Schema for a Notebook-based trigger.
type NotebookTriggerWrapper struct {
	// Trigger a workflow from a Notebook.
	NotebookTrigger interface{} `json:"notebookTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotebookTriggerWrapper instantiates a new NotebookTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookTriggerWrapper(notebookTrigger interface{}) *NotebookTriggerWrapper {
	this := NotebookTriggerWrapper{}
	this.NotebookTrigger = notebookTrigger
	return &this
}

// NewNotebookTriggerWrapperWithDefaults instantiates a new NotebookTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookTriggerWrapperWithDefaults() *NotebookTriggerWrapper {
	this := NotebookTriggerWrapper{}
	return &this
}

// GetNotebookTrigger returns the NotebookTrigger field value.
func (o *NotebookTriggerWrapper) GetNotebookTrigger() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NotebookTrigger
}

// GetNotebookTriggerOk returns a tuple with the NotebookTrigger field value
// and a boolean to check if the value has been set.
func (o *NotebookTriggerWrapper) GetNotebookTriggerOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotebookTrigger, true
}

// SetNotebookTrigger sets field value.
func (o *NotebookTriggerWrapper) SetNotebookTrigger(v interface{}) {
	o.NotebookTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *NotebookTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *NotebookTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *NotebookTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["notebookTrigger"] = o.NotebookTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NotebookTrigger *interface{} `json:"notebookTrigger"`
		StartStepNames  []string     `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.NotebookTrigger == nil {
		return fmt.Errorf("required field notebookTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"notebookTrigger", "startStepNames"})
	} else {
		return err
	}
	o.NotebookTrigger = *all.NotebookTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
