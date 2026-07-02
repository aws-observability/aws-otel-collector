// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationError A partial error for a single annotation that could not be processed.
type LLMObsAnnotationError struct {
	// ID of the annotation that failed, if applicable.
	AnnotationId *string `json:"annotation_id,omitempty"`
	// Error message.
	Error string `json:"error"`
	// ID of the interaction that failed.
	InteractionId string `json:"interaction_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationError instantiates a new LLMObsAnnotationError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationError(error string, interactionId string) *LLMObsAnnotationError {
	this := LLMObsAnnotationError{}
	this.Error = error
	this.InteractionId = interactionId
	return &this
}

// NewLLMObsAnnotationErrorWithDefaults instantiates a new LLMObsAnnotationError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationErrorWithDefaults() *LLMObsAnnotationError {
	this := LLMObsAnnotationError{}
	return &this
}

// GetAnnotationId returns the AnnotationId field value if set, zero value otherwise.
func (o *LLMObsAnnotationError) GetAnnotationId() string {
	if o == nil || o.AnnotationId == nil {
		var ret string
		return ret
	}
	return *o.AnnotationId
}

// GetAnnotationIdOk returns a tuple with the AnnotationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationError) GetAnnotationIdOk() (*string, bool) {
	if o == nil || o.AnnotationId == nil {
		return nil, false
	}
	return o.AnnotationId, true
}

// HasAnnotationId returns a boolean if a field has been set.
func (o *LLMObsAnnotationError) HasAnnotationId() bool {
	return o != nil && o.AnnotationId != nil
}

// SetAnnotationId gets a reference to the given string and assigns it to the AnnotationId field.
func (o *LLMObsAnnotationError) SetAnnotationId(v string) {
	o.AnnotationId = &v
}

// GetError returns the Error field value.
func (o *LLMObsAnnotationError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value.
func (o *LLMObsAnnotationError) SetError(v string) {
	o.Error = v
}

// GetInteractionId returns the InteractionId field value.
func (o *LLMObsAnnotationError) GetInteractionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InteractionId
}

// GetInteractionIdOk returns a tuple with the InteractionId field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationError) GetInteractionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InteractionId, true
}

// SetInteractionId sets field value.
func (o *LLMObsAnnotationError) SetInteractionId(v string) {
	o.InteractionId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnnotationId != nil {
		toSerialize["annotation_id"] = o.AnnotationId
	}
	toSerialize["error"] = o.Error
	toSerialize["interaction_id"] = o.InteractionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotationId  *string `json:"annotation_id,omitempty"`
		Error         *string `json:"error"`
		InteractionId *string `json:"interaction_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Error == nil {
		return fmt.Errorf("required field error missing")
	}
	if all.InteractionId == nil {
		return fmt.Errorf("required field interaction_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotation_id", "error", "interaction_id"})
	} else {
		return err
	}
	o.AnnotationId = all.AnnotationId
	o.Error = *all.Error
	o.InteractionId = *all.InteractionId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
