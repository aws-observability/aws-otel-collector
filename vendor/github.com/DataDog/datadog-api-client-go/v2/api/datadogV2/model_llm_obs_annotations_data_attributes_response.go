// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationsDataAttributesResponse Attributes of the annotations response.
type LLMObsAnnotationsDataAttributesResponse struct {
	// Successfully created or updated annotations.
	Annotations []LLMObsAnnotationItemResponse `json:"annotations"`
	// Partial errors for annotations that could not be processed.
	Errors []LLMObsAnnotationError `json:"errors,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationsDataAttributesResponse instantiates a new LLMObsAnnotationsDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationsDataAttributesResponse(annotations []LLMObsAnnotationItemResponse) *LLMObsAnnotationsDataAttributesResponse {
	this := LLMObsAnnotationsDataAttributesResponse{}
	this.Annotations = annotations
	return &this
}

// NewLLMObsAnnotationsDataAttributesResponseWithDefaults instantiates a new LLMObsAnnotationsDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationsDataAttributesResponseWithDefaults() *LLMObsAnnotationsDataAttributesResponse {
	this := LLMObsAnnotationsDataAttributesResponse{}
	return &this
}

// GetAnnotations returns the Annotations field value.
func (o *LLMObsAnnotationsDataAttributesResponse) GetAnnotations() []LLMObsAnnotationItemResponse {
	if o == nil {
		var ret []LLMObsAnnotationItemResponse
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationsDataAttributesResponse) GetAnnotationsOk() (*[]LLMObsAnnotationItemResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// SetAnnotations sets field value.
func (o *LLMObsAnnotationsDataAttributesResponse) SetAnnotations(v []LLMObsAnnotationItemResponse) {
	o.Annotations = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *LLMObsAnnotationsDataAttributesResponse) GetErrors() []LLMObsAnnotationError {
	if o == nil || o.Errors == nil {
		var ret []LLMObsAnnotationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationsDataAttributesResponse) GetErrorsOk() (*[]LLMObsAnnotationError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *LLMObsAnnotationsDataAttributesResponse) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given []LLMObsAnnotationError and assigns it to the Errors field.
func (o *LLMObsAnnotationsDataAttributesResponse) SetErrors(v []LLMObsAnnotationError) {
	o.Errors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationsDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotations"] = o.Annotations
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationsDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Annotations *[]LLMObsAnnotationItemResponse `json:"annotations"`
		Errors      []LLMObsAnnotationError         `json:"errors,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Annotations == nil {
		return fmt.Errorf("required field annotations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotations", "errors"})
	} else {
		return err
	}
	o.Annotations = *all.Annotations
	o.Errors = all.Errors

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
