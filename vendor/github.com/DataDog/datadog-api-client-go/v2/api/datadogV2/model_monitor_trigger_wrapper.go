// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorTriggerWrapper Schema for a Monitor-based trigger.
type MonitorTriggerWrapper struct {
	// Trigger a workflow from a Monitor. For automatic triggering a handle must be configured and the workflow must be published.
	MonitorTrigger MonitorTrigger `json:"monitorTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorTriggerWrapper instantiates a new MonitorTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorTriggerWrapper(monitorTrigger MonitorTrigger) *MonitorTriggerWrapper {
	this := MonitorTriggerWrapper{}
	this.MonitorTrigger = monitorTrigger
	return &this
}

// NewMonitorTriggerWrapperWithDefaults instantiates a new MonitorTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorTriggerWrapperWithDefaults() *MonitorTriggerWrapper {
	this := MonitorTriggerWrapper{}
	return &this
}

// GetMonitorTrigger returns the MonitorTrigger field value.
func (o *MonitorTriggerWrapper) GetMonitorTrigger() MonitorTrigger {
	if o == nil {
		var ret MonitorTrigger
		return ret
	}
	return o.MonitorTrigger
}

// GetMonitorTriggerOk returns a tuple with the MonitorTrigger field value
// and a boolean to check if the value has been set.
func (o *MonitorTriggerWrapper) GetMonitorTriggerOk() (*MonitorTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorTrigger, true
}

// SetMonitorTrigger sets field value.
func (o *MonitorTriggerWrapper) SetMonitorTrigger(v MonitorTrigger) {
	o.MonitorTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *MonitorTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *MonitorTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *MonitorTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["monitorTrigger"] = o.MonitorTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MonitorTrigger *MonitorTrigger `json:"monitorTrigger"`
		StartStepNames []string        `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MonitorTrigger == nil {
		return fmt.Errorf("required field monitorTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"monitorTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.MonitorTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MonitorTrigger = *all.MonitorTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
