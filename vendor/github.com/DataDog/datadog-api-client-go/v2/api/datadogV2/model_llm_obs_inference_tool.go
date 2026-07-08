// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsInferenceTool A tool definition available to the model during inference.
type LLMObsInferenceTool struct {
	// A function definition for a tool available to the model.
	Function LLMObsInferenceFunction `json:"function"`
	// The type of tool.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsInferenceTool instantiates a new LLMObsInferenceTool object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsInferenceTool(function LLMObsInferenceFunction, typeVar string) *LLMObsInferenceTool {
	this := LLMObsInferenceTool{}
	this.Function = function
	this.Type = typeVar
	return &this
}

// NewLLMObsInferenceToolWithDefaults instantiates a new LLMObsInferenceTool object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsInferenceToolWithDefaults() *LLMObsInferenceTool {
	this := LLMObsInferenceTool{}
	return &this
}

// GetFunction returns the Function field value.
func (o *LLMObsInferenceTool) GetFunction() LLMObsInferenceFunction {
	if o == nil {
		var ret LLMObsInferenceFunction
		return ret
	}
	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceTool) GetFunctionOk() (*LLMObsInferenceFunction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value.
func (o *LLMObsInferenceTool) SetFunction(v LLMObsInferenceFunction) {
	o.Function = v
}

// GetType returns the Type field value.
func (o *LLMObsInferenceTool) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceTool) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsInferenceTool) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsInferenceTool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["function"] = o.Function
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsInferenceTool) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Function *LLMObsInferenceFunction `json:"function"`
		Type     *string                  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Function == nil {
		return fmt.Errorf("required field function missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"function", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Function.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Function = *all.Function
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
