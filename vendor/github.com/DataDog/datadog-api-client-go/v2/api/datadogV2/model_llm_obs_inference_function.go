// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsInferenceFunction A function definition for a tool available to the model.
type LLMObsInferenceFunction struct {
	// A description of what the function does.
	Description *string `json:"description,omitempty"`
	// The name of the function.
	Name string `json:"name"`
	// JSON schema describing the function parameters.
	Parameters map[string]interface{} `json:"parameters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsInferenceFunction instantiates a new LLMObsInferenceFunction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsInferenceFunction(name string, parameters map[string]interface{}) *LLMObsInferenceFunction {
	this := LLMObsInferenceFunction{}
	this.Name = name
	this.Parameters = parameters
	return &this
}

// NewLLMObsInferenceFunctionWithDefaults instantiates a new LLMObsInferenceFunction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsInferenceFunctionWithDefaults() *LLMObsInferenceFunction {
	this := LLMObsInferenceFunction{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsInferenceFunction) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceFunction) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsInferenceFunction) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsInferenceFunction) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value.
func (o *LLMObsInferenceFunction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceFunction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsInferenceFunction) SetName(v string) {
	o.Name = v
}

// GetParameters returns the Parameters field value.
func (o *LLMObsInferenceFunction) GetParameters() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceFunction) GetParametersOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value.
func (o *LLMObsInferenceFunction) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsInferenceFunction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	toSerialize["parameters"] = o.Parameters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsInferenceFunction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                 `json:"description,omitempty"`
		Name        *string                 `json:"name"`
		Parameters  *map[string]interface{} `json:"parameters"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Parameters == nil {
		return fmt.Errorf("required field parameters missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "parameters"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Name = *all.Name
	o.Parameters = *all.Parameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
