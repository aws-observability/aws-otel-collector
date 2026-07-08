// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsInferenceErrorResponse Error details returned when an inference provider returns an error.
type LLMObsInferenceErrorResponse struct {
	// A human-readable description of the error.
	Message string `json:"message"`
	// The provider-specific error type.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsInferenceErrorResponse instantiates a new LLMObsInferenceErrorResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsInferenceErrorResponse(message string, typeVar string) *LLMObsInferenceErrorResponse {
	this := LLMObsInferenceErrorResponse{}
	this.Message = message
	this.Type = typeVar
	return &this
}

// NewLLMObsInferenceErrorResponseWithDefaults instantiates a new LLMObsInferenceErrorResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsInferenceErrorResponseWithDefaults() *LLMObsInferenceErrorResponse {
	this := LLMObsInferenceErrorResponse{}
	return &this
}

// GetMessage returns the Message field value.
func (o *LLMObsInferenceErrorResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *LLMObsInferenceErrorResponse) SetMessage(v string) {
	o.Message = v
}

// GetType returns the Type field value.
func (o *LLMObsInferenceErrorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsInferenceErrorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsInferenceErrorResponse) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsInferenceErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["message"] = o.Message
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsInferenceErrorResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message *string `json:"message"`
		Type    *string `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "type"})
	} else {
		return err
	}
	o.Message = *all.Message
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
