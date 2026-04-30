// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueInteractionsDataAttributesRequest Attributes for adding interactions to an annotation queue.
type LLMObsAnnotationQueueInteractionsDataAttributesRequest struct {
	// List of interactions to add to the queue. Must contain at least one item.
	Interactions []LLMObsAnnotationQueueInteractionItem `json:"interactions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueInteractionsDataAttributesRequest instantiates a new LLMObsAnnotationQueueInteractionsDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueInteractionsDataAttributesRequest(interactions []LLMObsAnnotationQueueInteractionItem) *LLMObsAnnotationQueueInteractionsDataAttributesRequest {
	this := LLMObsAnnotationQueueInteractionsDataAttributesRequest{}
	this.Interactions = interactions
	return &this
}

// NewLLMObsAnnotationQueueInteractionsDataAttributesRequestWithDefaults instantiates a new LLMObsAnnotationQueueInteractionsDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueInteractionsDataAttributesRequestWithDefaults() *LLMObsAnnotationQueueInteractionsDataAttributesRequest {
	this := LLMObsAnnotationQueueInteractionsDataAttributesRequest{}
	return &this
}

// GetInteractions returns the Interactions field value.
func (o *LLMObsAnnotationQueueInteractionsDataAttributesRequest) GetInteractions() []LLMObsAnnotationQueueInteractionItem {
	if o == nil {
		var ret []LLMObsAnnotationQueueInteractionItem
		return ret
	}
	return o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueInteractionsDataAttributesRequest) GetInteractionsOk() (*[]LLMObsAnnotationQueueInteractionItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactions, true
}

// SetInteractions sets field value.
func (o *LLMObsAnnotationQueueInteractionsDataAttributesRequest) SetInteractions(v []LLMObsAnnotationQueueInteractionItem) {
	o.Interactions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueInteractionsDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["interactions"] = o.Interactions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationQueueInteractionsDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interactions *[]LLMObsAnnotationQueueInteractionItem `json:"interactions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Interactions == nil {
		return fmt.Errorf("required field interactions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"interactions"})
	} else {
		return err
	}
	o.Interactions = *all.Interactions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
