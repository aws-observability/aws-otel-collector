// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityTriggerWrapper Schema for a Security-based trigger.
type SecurityTriggerWrapper struct {
	// Trigger a workflow from a Security Signal or Finding. For automatic triggering a handle must be configured and the workflow must be published.
	SecurityTrigger SecurityTrigger `json:"securityTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityTriggerWrapper instantiates a new SecurityTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityTriggerWrapper(securityTrigger SecurityTrigger) *SecurityTriggerWrapper {
	this := SecurityTriggerWrapper{}
	this.SecurityTrigger = securityTrigger
	return &this
}

// NewSecurityTriggerWrapperWithDefaults instantiates a new SecurityTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityTriggerWrapperWithDefaults() *SecurityTriggerWrapper {
	this := SecurityTriggerWrapper{}
	return &this
}

// GetSecurityTrigger returns the SecurityTrigger field value.
func (o *SecurityTriggerWrapper) GetSecurityTrigger() SecurityTrigger {
	if o == nil {
		var ret SecurityTrigger
		return ret
	}
	return o.SecurityTrigger
}

// GetSecurityTriggerOk returns a tuple with the SecurityTrigger field value
// and a boolean to check if the value has been set.
func (o *SecurityTriggerWrapper) GetSecurityTriggerOk() (*SecurityTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityTrigger, true
}

// SetSecurityTrigger sets field value.
func (o *SecurityTriggerWrapper) SetSecurityTrigger(v SecurityTrigger) {
	o.SecurityTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *SecurityTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *SecurityTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *SecurityTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["securityTrigger"] = o.SecurityTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SecurityTrigger *SecurityTrigger `json:"securityTrigger"`
		StartStepNames  []string         `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SecurityTrigger == nil {
		return fmt.Errorf("required field securityTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"securityTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.SecurityTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SecurityTrigger = *all.SecurityTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
