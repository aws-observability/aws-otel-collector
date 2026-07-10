// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueLabelSchemaData Data object for an annotation queue label schema.
type LLMObsAnnotationQueueLabelSchemaData struct {
	// Attributes of an annotation queue label schema.
	Attributes LLMObsAnnotationQueueLabelSchemaAttributes `json:"attributes"`
	// Unique identifier of the annotation queue.
	Id string `json:"id"`
	// Resource type of an LLM Observability annotation queue.
	Type LLMObsAnnotationQueueType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueLabelSchemaData instantiates a new LLMObsAnnotationQueueLabelSchemaData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueLabelSchemaData(attributes LLMObsAnnotationQueueLabelSchemaAttributes, id string, typeVar LLMObsAnnotationQueueType) *LLMObsAnnotationQueueLabelSchemaData {
	this := LLMObsAnnotationQueueLabelSchemaData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewLLMObsAnnotationQueueLabelSchemaDataWithDefaults instantiates a new LLMObsAnnotationQueueLabelSchemaData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueLabelSchemaDataWithDefaults() *LLMObsAnnotationQueueLabelSchemaData {
	this := LLMObsAnnotationQueueLabelSchemaData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *LLMObsAnnotationQueueLabelSchemaData) GetAttributes() LLMObsAnnotationQueueLabelSchemaAttributes {
	if o == nil {
		var ret LLMObsAnnotationQueueLabelSchemaAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueLabelSchemaData) GetAttributesOk() (*LLMObsAnnotationQueueLabelSchemaAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *LLMObsAnnotationQueueLabelSchemaData) SetAttributes(v LLMObsAnnotationQueueLabelSchemaAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *LLMObsAnnotationQueueLabelSchemaData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueLabelSchemaData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsAnnotationQueueLabelSchemaData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *LLMObsAnnotationQueueLabelSchemaData) GetType() LLMObsAnnotationQueueType {
	if o == nil {
		var ret LLMObsAnnotationQueueType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueLabelSchemaData) GetTypeOk() (*LLMObsAnnotationQueueType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsAnnotationQueueLabelSchemaData) SetType(v LLMObsAnnotationQueueType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueLabelSchemaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationQueueLabelSchemaData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *LLMObsAnnotationQueueLabelSchemaAttributes `json:"attributes"`
		Id         *string                                     `json:"id"`
		Type       *LLMObsAnnotationQueueType                  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
