// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormTriggerWrapper Schema for a Form-based trigger.
type FormTriggerWrapper struct {
	// Trigger a workflow from a Form.
	FormTrigger FormTrigger `json:"formTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormTriggerWrapper instantiates a new FormTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormTriggerWrapper(formTrigger FormTrigger) *FormTriggerWrapper {
	this := FormTriggerWrapper{}
	this.FormTrigger = formTrigger
	return &this
}

// NewFormTriggerWrapperWithDefaults instantiates a new FormTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormTriggerWrapperWithDefaults() *FormTriggerWrapper {
	this := FormTriggerWrapper{}
	return &this
}

// GetFormTrigger returns the FormTrigger field value.
func (o *FormTriggerWrapper) GetFormTrigger() FormTrigger {
	if o == nil {
		var ret FormTrigger
		return ret
	}
	return o.FormTrigger
}

// GetFormTriggerOk returns a tuple with the FormTrigger field value
// and a boolean to check if the value has been set.
func (o *FormTriggerWrapper) GetFormTriggerOk() (*FormTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormTrigger, true
}

// SetFormTrigger sets field value.
func (o *FormTriggerWrapper) SetFormTrigger(v FormTrigger) {
	o.FormTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *FormTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *FormTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *FormTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["formTrigger"] = o.FormTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FormTrigger    *FormTrigger `json:"formTrigger"`
		StartStepNames []string     `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FormTrigger == nil {
		return fmt.Errorf("required field formTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"formTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.FormTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FormTrigger = *all.FormTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
