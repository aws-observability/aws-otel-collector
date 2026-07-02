// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationSchema Schema defining the labels for an annotation queue.
type LLMObsAnnotationSchema struct {
	// List of label schema definitions.
	LabelSchemas []LLMObsLabelSchema `json:"label_schemas"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationSchema instantiates a new LLMObsAnnotationSchema object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationSchema(labelSchemas []LLMObsLabelSchema) *LLMObsAnnotationSchema {
	this := LLMObsAnnotationSchema{}
	this.LabelSchemas = labelSchemas
	return &this
}

// NewLLMObsAnnotationSchemaWithDefaults instantiates a new LLMObsAnnotationSchema object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationSchemaWithDefaults() *LLMObsAnnotationSchema {
	this := LLMObsAnnotationSchema{}
	return &this
}

// GetLabelSchemas returns the LabelSchemas field value.
func (o *LLMObsAnnotationSchema) GetLabelSchemas() []LLMObsLabelSchema {
	if o == nil {
		var ret []LLMObsLabelSchema
		return ret
	}
	return o.LabelSchemas
}

// GetLabelSchemasOk returns a tuple with the LabelSchemas field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationSchema) GetLabelSchemasOk() (*[]LLMObsLabelSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelSchemas, true
}

// SetLabelSchemas sets field value.
func (o *LLMObsAnnotationSchema) SetLabelSchemas(v []LLMObsLabelSchema) {
	o.LabelSchemas = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["label_schemas"] = o.LabelSchemas

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationSchema) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LabelSchemas *[]LLMObsLabelSchema `json:"label_schemas"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LabelSchemas == nil {
		return fmt.Errorf("required field label_schemas missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"label_schemas"})
	} else {
		return err
	}
	o.LabelSchemas = *all.LabelSchemas

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
