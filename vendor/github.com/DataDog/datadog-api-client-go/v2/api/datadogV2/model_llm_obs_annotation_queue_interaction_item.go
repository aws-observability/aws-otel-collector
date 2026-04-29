// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueInteractionItem A single interaction to add to an annotation queue.
type LLMObsAnnotationQueueInteractionItem struct {
	// Identifier of the content (such as trace ID) for this interaction.
	ContentId string `json:"content_id"`
	// Type of interaction in an annotation queue.
	Type LLMObsInteractionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueInteractionItem instantiates a new LLMObsAnnotationQueueInteractionItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueInteractionItem(contentId string, typeVar LLMObsInteractionType) *LLMObsAnnotationQueueInteractionItem {
	this := LLMObsAnnotationQueueInteractionItem{}
	this.ContentId = contentId
	this.Type = typeVar
	return &this
}

// NewLLMObsAnnotationQueueInteractionItemWithDefaults instantiates a new LLMObsAnnotationQueueInteractionItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueInteractionItemWithDefaults() *LLMObsAnnotationQueueInteractionItem {
	this := LLMObsAnnotationQueueInteractionItem{}
	return &this
}

// GetContentId returns the ContentId field value.
func (o *LLMObsAnnotationQueueInteractionItem) GetContentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueInteractionItem) GetContentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value.
func (o *LLMObsAnnotationQueueInteractionItem) SetContentId(v string) {
	o.ContentId = v
}

// GetType returns the Type field value.
func (o *LLMObsAnnotationQueueInteractionItem) GetType() LLMObsInteractionType {
	if o == nil {
		var ret LLMObsInteractionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueInteractionItem) GetTypeOk() (*LLMObsInteractionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsAnnotationQueueInteractionItem) SetType(v LLMObsInteractionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueInteractionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["content_id"] = o.ContentId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationQueueInteractionItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContentId *string                `json:"content_id"`
		Type      *LLMObsInteractionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ContentId == nil {
		return fmt.Errorf("required field content_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content_id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ContentId = *all.ContentId
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
