// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDisplayBlockInteractionItem An interaction whose rendered content is supplied directly as a list
// of display blocks. The server generates `content_id` deterministically
// from the block list.
type LLMObsDisplayBlockInteractionItem struct {
	// List of content blocks that make up a `display_block` interaction.
	// Must contain at least one block.
	DisplayBlock []LLMObsContentBlock `json:"display_block"`
	// Type discriminator for a `display_block` interaction.
	Type LLMObsDisplayBlockInteractionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDisplayBlockInteractionItem instantiates a new LLMObsDisplayBlockInteractionItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDisplayBlockInteractionItem(displayBlock []LLMObsContentBlock, typeVar LLMObsDisplayBlockInteractionType) *LLMObsDisplayBlockInteractionItem {
	this := LLMObsDisplayBlockInteractionItem{}
	this.DisplayBlock = displayBlock
	this.Type = typeVar
	return &this
}

// NewLLMObsDisplayBlockInteractionItemWithDefaults instantiates a new LLMObsDisplayBlockInteractionItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDisplayBlockInteractionItemWithDefaults() *LLMObsDisplayBlockInteractionItem {
	this := LLMObsDisplayBlockInteractionItem{}
	return &this
}

// GetDisplayBlock returns the DisplayBlock field value.
func (o *LLMObsDisplayBlockInteractionItem) GetDisplayBlock() []LLMObsContentBlock {
	if o == nil {
		var ret []LLMObsContentBlock
		return ret
	}
	return o.DisplayBlock
}

// GetDisplayBlockOk returns a tuple with the DisplayBlock field value
// and a boolean to check if the value has been set.
func (o *LLMObsDisplayBlockInteractionItem) GetDisplayBlockOk() (*[]LLMObsContentBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayBlock, true
}

// SetDisplayBlock sets field value.
func (o *LLMObsDisplayBlockInteractionItem) SetDisplayBlock(v []LLMObsContentBlock) {
	o.DisplayBlock = v
}

// GetType returns the Type field value.
func (o *LLMObsDisplayBlockInteractionItem) GetType() LLMObsDisplayBlockInteractionType {
	if o == nil {
		var ret LLMObsDisplayBlockInteractionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsDisplayBlockInteractionItem) GetTypeOk() (*LLMObsDisplayBlockInteractionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsDisplayBlockInteractionItem) SetType(v LLMObsDisplayBlockInteractionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDisplayBlockInteractionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["display_block"] = o.DisplayBlock
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDisplayBlockInteractionItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayBlock *[]LLMObsContentBlock              `json:"display_block"`
		Type         *LLMObsDisplayBlockInteractionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DisplayBlock == nil {
		return fmt.Errorf("required field display_block missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_block", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayBlock = *all.DisplayBlock
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
