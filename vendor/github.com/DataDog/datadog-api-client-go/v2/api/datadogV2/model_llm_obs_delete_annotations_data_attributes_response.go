// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDeleteAnnotationsDataAttributesResponse Attributes of the annotation deletion response.
type LLMObsDeleteAnnotationsDataAttributesResponse struct {
	// IDs of the successfully deleted annotations.
	AnnotationIds []string `json:"annotation_ids"`
	// Errors for annotations that could not be deleted.
	Errors []LLMObsDeleteAnnotationError `json:"errors"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDeleteAnnotationsDataAttributesResponse instantiates a new LLMObsDeleteAnnotationsDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDeleteAnnotationsDataAttributesResponse(annotationIds []string, errors []LLMObsDeleteAnnotationError) *LLMObsDeleteAnnotationsDataAttributesResponse {
	this := LLMObsDeleteAnnotationsDataAttributesResponse{}
	this.AnnotationIds = annotationIds
	this.Errors = errors
	return &this
}

// NewLLMObsDeleteAnnotationsDataAttributesResponseWithDefaults instantiates a new LLMObsDeleteAnnotationsDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDeleteAnnotationsDataAttributesResponseWithDefaults() *LLMObsDeleteAnnotationsDataAttributesResponse {
	this := LLMObsDeleteAnnotationsDataAttributesResponse{}
	return &this
}

// GetAnnotationIds returns the AnnotationIds field value.
func (o *LLMObsDeleteAnnotationsDataAttributesResponse) GetAnnotationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AnnotationIds
}

// GetAnnotationIdsOk returns a tuple with the AnnotationIds field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeleteAnnotationsDataAttributesResponse) GetAnnotationIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnnotationIds, true
}

// SetAnnotationIds sets field value.
func (o *LLMObsDeleteAnnotationsDataAttributesResponse) SetAnnotationIds(v []string) {
	o.AnnotationIds = v
}

// GetErrors returns the Errors field value.
func (o *LLMObsDeleteAnnotationsDataAttributesResponse) GetErrors() []LLMObsDeleteAnnotationError {
	if o == nil {
		var ret []LLMObsDeleteAnnotationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeleteAnnotationsDataAttributesResponse) GetErrorsOk() (*[]LLMObsDeleteAnnotationError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value.
func (o *LLMObsDeleteAnnotationsDataAttributesResponse) SetErrors(v []LLMObsDeleteAnnotationError) {
	o.Errors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDeleteAnnotationsDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotation_ids"] = o.AnnotationIds
	toSerialize["errors"] = o.Errors

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDeleteAnnotationsDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotationIds *[]string                      `json:"annotation_ids"`
		Errors        *[]LLMObsDeleteAnnotationError `json:"errors"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AnnotationIds == nil {
		return fmt.Errorf("required field annotation_ids missing")
	}
	if all.Errors == nil {
		return fmt.Errorf("required field errors missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotation_ids", "errors"})
	} else {
		return err
	}
	o.AnnotationIds = *all.AnnotationIds
	o.Errors = *all.Errors

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
