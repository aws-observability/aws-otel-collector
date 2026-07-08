// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueLabelSchemaAttributes Attributes of an annotation queue label schema.
type LLMObsAnnotationQueueLabelSchemaAttributes struct {
	// Schema defining the labels for an annotation queue.
	AnnotationSchema LLMObsAnnotationSchema `json:"annotation_schema"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueLabelSchemaAttributes instantiates a new LLMObsAnnotationQueueLabelSchemaAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueLabelSchemaAttributes(annotationSchema LLMObsAnnotationSchema) *LLMObsAnnotationQueueLabelSchemaAttributes {
	this := LLMObsAnnotationQueueLabelSchemaAttributes{}
	this.AnnotationSchema = annotationSchema
	return &this
}

// NewLLMObsAnnotationQueueLabelSchemaAttributesWithDefaults instantiates a new LLMObsAnnotationQueueLabelSchemaAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueLabelSchemaAttributesWithDefaults() *LLMObsAnnotationQueueLabelSchemaAttributes {
	this := LLMObsAnnotationQueueLabelSchemaAttributes{}
	return &this
}

// GetAnnotationSchema returns the AnnotationSchema field value.
func (o *LLMObsAnnotationQueueLabelSchemaAttributes) GetAnnotationSchema() LLMObsAnnotationSchema {
	if o == nil {
		var ret LLMObsAnnotationSchema
		return ret
	}
	return o.AnnotationSchema
}

// GetAnnotationSchemaOk returns a tuple with the AnnotationSchema field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueLabelSchemaAttributes) GetAnnotationSchemaOk() (*LLMObsAnnotationSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnnotationSchema, true
}

// SetAnnotationSchema sets field value.
func (o *LLMObsAnnotationQueueLabelSchemaAttributes) SetAnnotationSchema(v LLMObsAnnotationSchema) {
	o.AnnotationSchema = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueLabelSchemaAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotation_schema"] = o.AnnotationSchema

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationQueueLabelSchemaAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotationSchema *LLMObsAnnotationSchema `json:"annotation_schema"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AnnotationSchema == nil {
		return fmt.Errorf("required field annotation_schema missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotation_schema"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AnnotationSchema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnnotationSchema = *all.AnnotationSchema

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
