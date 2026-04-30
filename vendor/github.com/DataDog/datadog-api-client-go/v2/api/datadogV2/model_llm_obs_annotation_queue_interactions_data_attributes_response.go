// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueInteractionsDataAttributesResponse Attributes of the interaction addition response.
type LLMObsAnnotationQueueInteractionsDataAttributesResponse struct {
	// List of interactions that were processed.
	Interactions []LLMObsAnnotationQueueInteractionResponseItem `json:"interactions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueInteractionsDataAttributesResponse instantiates a new LLMObsAnnotationQueueInteractionsDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueInteractionsDataAttributesResponse(interactions []LLMObsAnnotationQueueInteractionResponseItem) *LLMObsAnnotationQueueInteractionsDataAttributesResponse {
	this := LLMObsAnnotationQueueInteractionsDataAttributesResponse{}
	this.Interactions = interactions
	return &this
}

// NewLLMObsAnnotationQueueInteractionsDataAttributesResponseWithDefaults instantiates a new LLMObsAnnotationQueueInteractionsDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueInteractionsDataAttributesResponseWithDefaults() *LLMObsAnnotationQueueInteractionsDataAttributesResponse {
	this := LLMObsAnnotationQueueInteractionsDataAttributesResponse{}
	return &this
}

// GetInteractions returns the Interactions field value.
func (o *LLMObsAnnotationQueueInteractionsDataAttributesResponse) GetInteractions() []LLMObsAnnotationQueueInteractionResponseItem {
	if o == nil {
		var ret []LLMObsAnnotationQueueInteractionResponseItem
		return ret
	}
	return o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueInteractionsDataAttributesResponse) GetInteractionsOk() (*[]LLMObsAnnotationQueueInteractionResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactions, true
}

// SetInteractions sets field value.
func (o *LLMObsAnnotationQueueInteractionsDataAttributesResponse) SetInteractions(v []LLMObsAnnotationQueueInteractionResponseItem) {
	o.Interactions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueInteractionsDataAttributesResponse) MarshalJSON() ([]byte, error) {
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
func (o *LLMObsAnnotationQueueInteractionsDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interactions *[]LLMObsAnnotationQueueInteractionResponseItem `json:"interactions"`
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
