// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest Attributes for deleting interactions from an annotation queue.
type LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest struct {
	// List of interaction IDs to delete. Must contain at least one item.
	InteractionIds []string `json:"interaction_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest instantiates a new LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest(interactionIds []string) *LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest {
	this := LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest{}
	this.InteractionIds = interactionIds
	return &this
}

// NewLLMObsDeleteAnnotationQueueInteractionsDataAttributesRequestWithDefaults instantiates a new LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDeleteAnnotationQueueInteractionsDataAttributesRequestWithDefaults() *LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest {
	this := LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest{}
	return &this
}

// GetInteractionIds returns the InteractionIds field value.
func (o *LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest) GetInteractionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.InteractionIds
}

// GetInteractionIdsOk returns a tuple with the InteractionIds field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest) GetInteractionIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InteractionIds, true
}

// SetInteractionIds sets field value.
func (o *LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest) SetInteractionIds(v []string) {
	o.InteractionIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["interaction_ids"] = o.InteractionIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDeleteAnnotationQueueInteractionsDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InteractionIds *[]string `json:"interaction_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InteractionIds == nil {
		return fmt.Errorf("required field interaction_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"interaction_ids"})
	} else {
		return err
	}
	o.InteractionIds = *all.InteractionIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
