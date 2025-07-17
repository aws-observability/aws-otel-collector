// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppTriggerWrapper Schema for an App-based trigger.
type AppTriggerWrapper struct {
	// Trigger a workflow from an App.
	AppTrigger interface{} `json:"appTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAppTriggerWrapper instantiates a new AppTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAppTriggerWrapper(appTrigger interface{}) *AppTriggerWrapper {
	this := AppTriggerWrapper{}
	this.AppTrigger = appTrigger
	return &this
}

// NewAppTriggerWrapperWithDefaults instantiates a new AppTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAppTriggerWrapperWithDefaults() *AppTriggerWrapper {
	this := AppTriggerWrapper{}
	return &this
}

// GetAppTrigger returns the AppTrigger field value.
func (o *AppTriggerWrapper) GetAppTrigger() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AppTrigger
}

// GetAppTriggerOk returns a tuple with the AppTrigger field value
// and a boolean to check if the value has been set.
func (o *AppTriggerWrapper) GetAppTriggerOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppTrigger, true
}

// SetAppTrigger sets field value.
func (o *AppTriggerWrapper) SetAppTrigger(v interface{}) {
	o.AppTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *AppTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *AppTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *AppTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AppTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["appTrigger"] = o.AppTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AppTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppTrigger     *interface{} `json:"appTrigger"`
		StartStepNames []string     `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AppTrigger == nil {
		return fmt.Errorf("required field appTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"appTrigger", "startStepNames"})
	} else {
		return err
	}
	o.AppTrigger = *all.AppTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
