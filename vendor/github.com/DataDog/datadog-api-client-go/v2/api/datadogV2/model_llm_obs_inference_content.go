// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsInferenceContent A structured content block within a message.
type LLMObsInferenceContent struct {
	// The content block type.
	Type string `json:"type"`
	// The typed value of a message content block.
	Value LLMObsInferenceContentValue `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsInferenceContent instantiates a new LLMObsInferenceContent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsInferenceContent(typeVar string, value LLMObsInferenceContentValue) *LLMObsInferenceContent {
	this := LLMObsInferenceContent{}
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewLLMObsInferenceContentWithDefaults instantiates a new LLMObsInferenceContent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsInferenceContentWithDefaults() *LLMObsInferenceContent {
	this := LLMObsInferenceContent{}
	return &this
}

// GetType returns the Type field value.
func (o *LLMObsInferenceContent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceContent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsInferenceContent) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *LLMObsInferenceContent) GetValue() LLMObsInferenceContentValue {
	if o == nil {
		var ret LLMObsInferenceContentValue
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceContent) GetValueOk() (*LLMObsInferenceContentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *LLMObsInferenceContent) SetValue(v LLMObsInferenceContentValue) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsInferenceContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsInferenceContent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *string                      `json:"type"`
		Value *LLMObsInferenceContentValue `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Type = *all.Type
	if all.Value.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
