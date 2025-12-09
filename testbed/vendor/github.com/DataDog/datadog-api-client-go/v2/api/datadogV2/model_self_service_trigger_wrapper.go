// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SelfServiceTriggerWrapper Schema for a Self Service-based trigger.
type SelfServiceTriggerWrapper struct {
	// Trigger a workflow from Self Service.
	SelfServiceTrigger interface{} `json:"selfServiceTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSelfServiceTriggerWrapper instantiates a new SelfServiceTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSelfServiceTriggerWrapper(selfServiceTrigger interface{}) *SelfServiceTriggerWrapper {
	this := SelfServiceTriggerWrapper{}
	this.SelfServiceTrigger = selfServiceTrigger
	return &this
}

// NewSelfServiceTriggerWrapperWithDefaults instantiates a new SelfServiceTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSelfServiceTriggerWrapperWithDefaults() *SelfServiceTriggerWrapper {
	this := SelfServiceTriggerWrapper{}
	return &this
}

// GetSelfServiceTrigger returns the SelfServiceTrigger field value.
func (o *SelfServiceTriggerWrapper) GetSelfServiceTrigger() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SelfServiceTrigger
}

// GetSelfServiceTriggerOk returns a tuple with the SelfServiceTrigger field value
// and a boolean to check if the value has been set.
func (o *SelfServiceTriggerWrapper) GetSelfServiceTriggerOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfServiceTrigger, true
}

// SetSelfServiceTrigger sets field value.
func (o *SelfServiceTriggerWrapper) SetSelfServiceTrigger(v interface{}) {
	o.SelfServiceTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *SelfServiceTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *SelfServiceTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *SelfServiceTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SelfServiceTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["selfServiceTrigger"] = o.SelfServiceTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SelfServiceTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SelfServiceTrigger *interface{} `json:"selfServiceTrigger"`
		StartStepNames     []string     `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SelfServiceTrigger == nil {
		return fmt.Errorf("required field selfServiceTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"selfServiceTrigger", "startStepNames"})
	} else {
		return err
	}
	o.SelfServiceTrigger = *all.SelfServiceTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
