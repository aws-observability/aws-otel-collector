// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDeleteAnnotationsDataAttributesRequest Attributes for deleting annotations.
type LLMObsDeleteAnnotationsDataAttributesRequest struct {
	// IDs of the annotations to delete. Must contain at least one item.
	AnnotationIds []string `json:"annotation_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDeleteAnnotationsDataAttributesRequest instantiates a new LLMObsDeleteAnnotationsDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDeleteAnnotationsDataAttributesRequest(annotationIds []string) *LLMObsDeleteAnnotationsDataAttributesRequest {
	this := LLMObsDeleteAnnotationsDataAttributesRequest{}
	this.AnnotationIds = annotationIds
	return &this
}

// NewLLMObsDeleteAnnotationsDataAttributesRequestWithDefaults instantiates a new LLMObsDeleteAnnotationsDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDeleteAnnotationsDataAttributesRequestWithDefaults() *LLMObsDeleteAnnotationsDataAttributesRequest {
	this := LLMObsDeleteAnnotationsDataAttributesRequest{}
	return &this
}

// GetAnnotationIds returns the AnnotationIds field value.
func (o *LLMObsDeleteAnnotationsDataAttributesRequest) GetAnnotationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AnnotationIds
}

// GetAnnotationIdsOk returns a tuple with the AnnotationIds field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeleteAnnotationsDataAttributesRequest) GetAnnotationIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnnotationIds, true
}

// SetAnnotationIds sets field value.
func (o *LLMObsDeleteAnnotationsDataAttributesRequest) SetAnnotationIds(v []string) {
	o.AnnotationIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDeleteAnnotationsDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotation_ids"] = o.AnnotationIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDeleteAnnotationsDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotationIds *[]string `json:"annotation_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AnnotationIds == nil {
		return fmt.Errorf("required field annotation_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotation_ids"})
	} else {
		return err
	}
	o.AnnotationIds = *all.AnnotationIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
