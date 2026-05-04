// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueInteractionResponseItem A single interaction result.
type LLMObsAnnotationQueueInteractionResponseItem struct {
	// Whether this interaction already existed in the queue.
	AlreadyExisted bool `json:"already_existed"`
	// Identifier of the content for this interaction.
	ContentId string `json:"content_id"`
	// Unique identifier of the interaction.
	Id string `json:"id"`
	// Type of interaction in an annotation queue.
	Type LLMObsInteractionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueInteractionResponseItem instantiates a new LLMObsAnnotationQueueInteractionResponseItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueInteractionResponseItem(alreadyExisted bool, contentId string, id string, typeVar LLMObsInteractionType) *LLMObsAnnotationQueueInteractionResponseItem {
	this := LLMObsAnnotationQueueInteractionResponseItem{}
	this.AlreadyExisted = alreadyExisted
	this.ContentId = contentId
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewLLMObsAnnotationQueueInteractionResponseItemWithDefaults instantiates a new LLMObsAnnotationQueueInteractionResponseItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueInteractionResponseItemWithDefaults() *LLMObsAnnotationQueueInteractionResponseItem {
	this := LLMObsAnnotationQueueInteractionResponseItem{}
	return &this
}

// GetAlreadyExisted returns the AlreadyExisted field value.
func (o *LLMObsAnnotationQueueInteractionResponseItem) GetAlreadyExisted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AlreadyExisted
}

// GetAlreadyExistedOk returns a tuple with the AlreadyExisted field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueInteractionResponseItem) GetAlreadyExistedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlreadyExisted, true
}

// SetAlreadyExisted sets field value.
func (o *LLMObsAnnotationQueueInteractionResponseItem) SetAlreadyExisted(v bool) {
	o.AlreadyExisted = v
}

// GetContentId returns the ContentId field value.
func (o *LLMObsAnnotationQueueInteractionResponseItem) GetContentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueInteractionResponseItem) GetContentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value.
func (o *LLMObsAnnotationQueueInteractionResponseItem) SetContentId(v string) {
	o.ContentId = v
}

// GetId returns the Id field value.
func (o *LLMObsAnnotationQueueInteractionResponseItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueInteractionResponseItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsAnnotationQueueInteractionResponseItem) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *LLMObsAnnotationQueueInteractionResponseItem) GetType() LLMObsInteractionType {
	if o == nil {
		var ret LLMObsInteractionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueInteractionResponseItem) GetTypeOk() (*LLMObsInteractionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsAnnotationQueueInteractionResponseItem) SetType(v LLMObsInteractionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueInteractionResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["already_existed"] = o.AlreadyExisted
	toSerialize["content_id"] = o.ContentId
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationQueueInteractionResponseItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlreadyExisted *bool                  `json:"already_existed"`
		ContentId      *string                `json:"content_id"`
		Id             *string                `json:"id"`
		Type           *LLMObsInteractionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AlreadyExisted == nil {
		return fmt.Errorf("required field already_existed missing")
	}
	if all.ContentId == nil {
		return fmt.Errorf("required field content_id missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"already_existed", "content_id", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AlreadyExisted = *all.AlreadyExisted
	o.ContentId = *all.ContentId
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
