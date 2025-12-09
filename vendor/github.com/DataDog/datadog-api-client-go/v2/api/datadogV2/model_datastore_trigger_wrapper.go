// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatastoreTriggerWrapper Schema for a Datastore-based trigger.
type DatastoreTriggerWrapper struct {
	// Trigger a workflow from a Datastore. For automatic triggering a handle must be configured and the workflow must be published.
	DatastoreTrigger DatastoreTrigger `json:"datastoreTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatastoreTriggerWrapper instantiates a new DatastoreTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatastoreTriggerWrapper(datastoreTrigger DatastoreTrigger) *DatastoreTriggerWrapper {
	this := DatastoreTriggerWrapper{}
	this.DatastoreTrigger = datastoreTrigger
	return &this
}

// NewDatastoreTriggerWrapperWithDefaults instantiates a new DatastoreTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatastoreTriggerWrapperWithDefaults() *DatastoreTriggerWrapper {
	this := DatastoreTriggerWrapper{}
	return &this
}

// GetDatastoreTrigger returns the DatastoreTrigger field value.
func (o *DatastoreTriggerWrapper) GetDatastoreTrigger() DatastoreTrigger {
	if o == nil {
		var ret DatastoreTrigger
		return ret
	}
	return o.DatastoreTrigger
}

// GetDatastoreTriggerOk returns a tuple with the DatastoreTrigger field value
// and a boolean to check if the value has been set.
func (o *DatastoreTriggerWrapper) GetDatastoreTriggerOk() (*DatastoreTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatastoreTrigger, true
}

// SetDatastoreTrigger sets field value.
func (o *DatastoreTriggerWrapper) SetDatastoreTrigger(v DatastoreTrigger) {
	o.DatastoreTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *DatastoreTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *DatastoreTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *DatastoreTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatastoreTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["datastoreTrigger"] = o.DatastoreTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatastoreTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatastoreTrigger *DatastoreTrigger `json:"datastoreTrigger"`
		StartStepNames   []string          `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatastoreTrigger == nil {
		return fmt.Errorf("required field datastoreTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"datastoreTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DatastoreTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatastoreTrigger = *all.DatastoreTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
