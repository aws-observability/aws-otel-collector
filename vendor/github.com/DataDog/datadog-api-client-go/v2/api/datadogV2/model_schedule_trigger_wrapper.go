// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleTriggerWrapper Schema for a Schedule-based trigger.
type ScheduleTriggerWrapper struct {
	// Trigger a workflow from a Schedule. The workflow must be published.
	ScheduleTrigger ScheduleTrigger `json:"scheduleTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleTriggerWrapper instantiates a new ScheduleTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleTriggerWrapper(scheduleTrigger ScheduleTrigger) *ScheduleTriggerWrapper {
	this := ScheduleTriggerWrapper{}
	this.ScheduleTrigger = scheduleTrigger
	return &this
}

// NewScheduleTriggerWrapperWithDefaults instantiates a new ScheduleTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleTriggerWrapperWithDefaults() *ScheduleTriggerWrapper {
	this := ScheduleTriggerWrapper{}
	return &this
}

// GetScheduleTrigger returns the ScheduleTrigger field value.
func (o *ScheduleTriggerWrapper) GetScheduleTrigger() ScheduleTrigger {
	if o == nil {
		var ret ScheduleTrigger
		return ret
	}
	return o.ScheduleTrigger
}

// GetScheduleTriggerOk returns a tuple with the ScheduleTrigger field value
// and a boolean to check if the value has been set.
func (o *ScheduleTriggerWrapper) GetScheduleTriggerOk() (*ScheduleTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleTrigger, true
}

// SetScheduleTrigger sets field value.
func (o *ScheduleTriggerWrapper) SetScheduleTrigger(v ScheduleTrigger) {
	o.ScheduleTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *ScheduleTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *ScheduleTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *ScheduleTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["scheduleTrigger"] = o.ScheduleTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ScheduleTrigger *ScheduleTrigger `json:"scheduleTrigger"`
		StartStepNames  []string         `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ScheduleTrigger == nil {
		return fmt.Errorf("required field scheduleTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"scheduleTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ScheduleTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ScheduleTrigger = *all.ScheduleTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
