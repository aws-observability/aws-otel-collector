// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDeleteAnnotationError A partial error for a single annotation that could not be deleted.
type LLMObsDeleteAnnotationError struct {
	// ID of the annotation that could not be deleted.
	AnnotationId string `json:"annotation_id"`
	// Error message.
	Error string `json:"error"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDeleteAnnotationError instantiates a new LLMObsDeleteAnnotationError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDeleteAnnotationError(annotationId string, error string) *LLMObsDeleteAnnotationError {
	this := LLMObsDeleteAnnotationError{}
	this.AnnotationId = annotationId
	this.Error = error
	return &this
}

// NewLLMObsDeleteAnnotationErrorWithDefaults instantiates a new LLMObsDeleteAnnotationError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDeleteAnnotationErrorWithDefaults() *LLMObsDeleteAnnotationError {
	this := LLMObsDeleteAnnotationError{}
	return &this
}

// GetAnnotationId returns the AnnotationId field value.
func (o *LLMObsDeleteAnnotationError) GetAnnotationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AnnotationId
}

// GetAnnotationIdOk returns a tuple with the AnnotationId field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeleteAnnotationError) GetAnnotationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnnotationId, true
}

// SetAnnotationId sets field value.
func (o *LLMObsDeleteAnnotationError) SetAnnotationId(v string) {
	o.AnnotationId = v
}

// GetError returns the Error field value.
func (o *LLMObsDeleteAnnotationError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeleteAnnotationError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value.
func (o *LLMObsDeleteAnnotationError) SetError(v string) {
	o.Error = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDeleteAnnotationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotation_id"] = o.AnnotationId
	toSerialize["error"] = o.Error

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDeleteAnnotationError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotationId *string `json:"annotation_id"`
		Error        *string `json:"error"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AnnotationId == nil {
		return fmt.Errorf("required field annotation_id missing")
	}
	if all.Error == nil {
		return fmt.Errorf("required field error missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotation_id", "error"})
	} else {
		return err
	}
	o.AnnotationId = *all.AnnotationId
	o.Error = *all.Error

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
