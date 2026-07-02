// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationsDataAttributesRequest Attributes for creating or updating annotations.
type LLMObsAnnotationsDataAttributesRequest struct {
	// List of annotations to create or update. Must contain at least one item.
	Annotations []LLMObsUpsertAnnotationItem `json:"annotations"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationsDataAttributesRequest instantiates a new LLMObsAnnotationsDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationsDataAttributesRequest(annotations []LLMObsUpsertAnnotationItem) *LLMObsAnnotationsDataAttributesRequest {
	this := LLMObsAnnotationsDataAttributesRequest{}
	this.Annotations = annotations
	return &this
}

// NewLLMObsAnnotationsDataAttributesRequestWithDefaults instantiates a new LLMObsAnnotationsDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationsDataAttributesRequestWithDefaults() *LLMObsAnnotationsDataAttributesRequest {
	this := LLMObsAnnotationsDataAttributesRequest{}
	return &this
}

// GetAnnotations returns the Annotations field value.
func (o *LLMObsAnnotationsDataAttributesRequest) GetAnnotations() []LLMObsUpsertAnnotationItem {
	if o == nil {
		var ret []LLMObsUpsertAnnotationItem
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationsDataAttributesRequest) GetAnnotationsOk() (*[]LLMObsUpsertAnnotationItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// SetAnnotations sets field value.
func (o *LLMObsAnnotationsDataAttributesRequest) SetAnnotations(v []LLMObsUpsertAnnotationItem) {
	o.Annotations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationsDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotations"] = o.Annotations

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationsDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Annotations *[]LLMObsUpsertAnnotationItem `json:"annotations"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Annotations == nil {
		return fmt.Errorf("required field annotations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotations"})
	} else {
		return err
	}
	o.Annotations = *all.Annotations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
