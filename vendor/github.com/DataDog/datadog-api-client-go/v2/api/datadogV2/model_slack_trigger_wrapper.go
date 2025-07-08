// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SlackTriggerWrapper Schema for a Slack-based trigger.
type SlackTriggerWrapper struct {
	// Trigger a workflow from Slack. The workflow must be published.
	SlackTrigger interface{} `json:"slackTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSlackTriggerWrapper instantiates a new SlackTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSlackTriggerWrapper(slackTrigger interface{}) *SlackTriggerWrapper {
	this := SlackTriggerWrapper{}
	this.SlackTrigger = slackTrigger
	return &this
}

// NewSlackTriggerWrapperWithDefaults instantiates a new SlackTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSlackTriggerWrapperWithDefaults() *SlackTriggerWrapper {
	this := SlackTriggerWrapper{}
	return &this
}

// GetSlackTrigger returns the SlackTrigger field value.
func (o *SlackTriggerWrapper) GetSlackTrigger() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SlackTrigger
}

// GetSlackTriggerOk returns a tuple with the SlackTrigger field value
// and a boolean to check if the value has been set.
func (o *SlackTriggerWrapper) GetSlackTriggerOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlackTrigger, true
}

// SetSlackTrigger sets field value.
func (o *SlackTriggerWrapper) SetSlackTrigger(v interface{}) {
	o.SlackTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *SlackTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *SlackTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *SlackTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SlackTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["slackTrigger"] = o.SlackTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SlackTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SlackTrigger   *interface{} `json:"slackTrigger"`
		StartStepNames []string     `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SlackTrigger == nil {
		return fmt.Errorf("required field slackTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"slackTrigger", "startStepNames"})
	} else {
		return err
	}
	o.SlackTrigger = *all.SlackTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
