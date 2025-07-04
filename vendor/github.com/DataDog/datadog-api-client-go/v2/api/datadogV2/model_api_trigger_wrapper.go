// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APITriggerWrapper Schema for an API-based trigger.
type APITriggerWrapper struct {
	// Trigger a workflow from an API request. The workflow must be published.
	ApiTrigger APITrigger `json:"apiTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAPITriggerWrapper instantiates a new APITriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAPITriggerWrapper(apiTrigger APITrigger) *APITriggerWrapper {
	this := APITriggerWrapper{}
	this.ApiTrigger = apiTrigger
	return &this
}

// NewAPITriggerWrapperWithDefaults instantiates a new APITriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAPITriggerWrapperWithDefaults() *APITriggerWrapper {
	this := APITriggerWrapper{}
	return &this
}

// GetApiTrigger returns the ApiTrigger field value.
func (o *APITriggerWrapper) GetApiTrigger() APITrigger {
	if o == nil {
		var ret APITrigger
		return ret
	}
	return o.ApiTrigger
}

// GetApiTriggerOk returns a tuple with the ApiTrigger field value
// and a boolean to check if the value has been set.
func (o *APITriggerWrapper) GetApiTriggerOk() (*APITrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiTrigger, true
}

// SetApiTrigger sets field value.
func (o *APITriggerWrapper) SetApiTrigger(v APITrigger) {
	o.ApiTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *APITriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APITriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *APITriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *APITriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o APITriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["apiTrigger"] = o.ApiTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *APITriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiTrigger     *APITrigger `json:"apiTrigger"`
		StartStepNames []string    `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiTrigger == nil {
		return fmt.Errorf("required field apiTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"apiTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ApiTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ApiTrigger = *all.ApiTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
