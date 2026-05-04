// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallTriggerWrapper Schema for an On-Call-based trigger.
type OnCallTriggerWrapper struct {
	// Trigger a workflow from an On-Call Page or On-Call Handover. For automatic triggering a handle must be configured and the workflow must be published.
	OnCallTrigger OnCallTrigger `json:"onCallTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallTriggerWrapper instantiates a new OnCallTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallTriggerWrapper(onCallTrigger OnCallTrigger) *OnCallTriggerWrapper {
	this := OnCallTriggerWrapper{}
	this.OnCallTrigger = onCallTrigger
	return &this
}

// NewOnCallTriggerWrapperWithDefaults instantiates a new OnCallTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallTriggerWrapperWithDefaults() *OnCallTriggerWrapper {
	this := OnCallTriggerWrapper{}
	return &this
}

// GetOnCallTrigger returns the OnCallTrigger field value.
func (o *OnCallTriggerWrapper) GetOnCallTrigger() OnCallTrigger {
	if o == nil {
		var ret OnCallTrigger
		return ret
	}
	return o.OnCallTrigger
}

// GetOnCallTriggerOk returns a tuple with the OnCallTrigger field value
// and a boolean to check if the value has been set.
func (o *OnCallTriggerWrapper) GetOnCallTriggerOk() (*OnCallTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnCallTrigger, true
}

// SetOnCallTrigger sets field value.
func (o *OnCallTriggerWrapper) SetOnCallTrigger(v OnCallTrigger) {
	o.OnCallTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *OnCallTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *OnCallTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *OnCallTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["onCallTrigger"] = o.OnCallTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OnCallTrigger  *OnCallTrigger `json:"onCallTrigger"`
		StartStepNames []string       `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OnCallTrigger == nil {
		return fmt.Errorf("required field onCallTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"onCallTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OnCallTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OnCallTrigger = *all.OnCallTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
