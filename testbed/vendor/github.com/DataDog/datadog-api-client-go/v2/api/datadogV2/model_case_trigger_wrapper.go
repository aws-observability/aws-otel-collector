// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseTriggerWrapper Schema for a Case-based trigger.
type CaseTriggerWrapper struct {
	// Trigger a workflow from a Case. For automatic triggering a handle must be configured and the workflow must be published.
	CaseTrigger CaseTrigger `json:"caseTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseTriggerWrapper instantiates a new CaseTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseTriggerWrapper(caseTrigger CaseTrigger) *CaseTriggerWrapper {
	this := CaseTriggerWrapper{}
	this.CaseTrigger = caseTrigger
	return &this
}

// NewCaseTriggerWrapperWithDefaults instantiates a new CaseTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseTriggerWrapperWithDefaults() *CaseTriggerWrapper {
	this := CaseTriggerWrapper{}
	return &this
}

// GetCaseTrigger returns the CaseTrigger field value.
func (o *CaseTriggerWrapper) GetCaseTrigger() CaseTrigger {
	if o == nil {
		var ret CaseTrigger
		return ret
	}
	return o.CaseTrigger
}

// GetCaseTriggerOk returns a tuple with the CaseTrigger field value
// and a boolean to check if the value has been set.
func (o *CaseTriggerWrapper) GetCaseTriggerOk() (*CaseTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaseTrigger, true
}

// SetCaseTrigger sets field value.
func (o *CaseTriggerWrapper) SetCaseTrigger(v CaseTrigger) {
	o.CaseTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *CaseTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *CaseTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *CaseTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["caseTrigger"] = o.CaseTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaseTrigger    *CaseTrigger `json:"caseTrigger"`
		StartStepNames []string     `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CaseTrigger == nil {
		return fmt.Errorf("required field caseTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"caseTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CaseTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CaseTrigger = *all.CaseTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
