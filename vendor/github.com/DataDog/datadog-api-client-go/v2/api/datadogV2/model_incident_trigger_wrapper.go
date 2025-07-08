// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTriggerWrapper Schema for an Incident-based trigger.
type IncidentTriggerWrapper struct {
	// Trigger a workflow from an Incident. For automatic triggering a handle must be configured and the workflow must be published.
	IncidentTrigger IncidentTrigger `json:"incidentTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTriggerWrapper instantiates a new IncidentTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTriggerWrapper(incidentTrigger IncidentTrigger) *IncidentTriggerWrapper {
	this := IncidentTriggerWrapper{}
	this.IncidentTrigger = incidentTrigger
	return &this
}

// NewIncidentTriggerWrapperWithDefaults instantiates a new IncidentTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTriggerWrapperWithDefaults() *IncidentTriggerWrapper {
	this := IncidentTriggerWrapper{}
	return &this
}

// GetIncidentTrigger returns the IncidentTrigger field value.
func (o *IncidentTriggerWrapper) GetIncidentTrigger() IncidentTrigger {
	if o == nil {
		var ret IncidentTrigger
		return ret
	}
	return o.IncidentTrigger
}

// GetIncidentTriggerOk returns a tuple with the IncidentTrigger field value
// and a boolean to check if the value has been set.
func (o *IncidentTriggerWrapper) GetIncidentTriggerOk() (*IncidentTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentTrigger, true
}

// SetIncidentTrigger sets field value.
func (o *IncidentTriggerWrapper) SetIncidentTrigger(v IncidentTrigger) {
	o.IncidentTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *IncidentTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *IncidentTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *IncidentTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["incidentTrigger"] = o.IncidentTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentTrigger *IncidentTrigger `json:"incidentTrigger"`
		StartStepNames  []string         `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncidentTrigger == nil {
		return fmt.Errorf("required field incidentTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incidentTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IncidentTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentTrigger = *all.IncidentTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
