// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueInteractionsDataResponse Data object for the interaction addition response.
type LLMObsAnnotationQueueInteractionsDataResponse struct {
	// Attributes of the interaction addition response.
	Attributes LLMObsAnnotationQueueInteractionsDataAttributesResponse `json:"attributes"`
	// The queue ID the interactions were added to.
	Id string `json:"id"`
	// Resource type for annotation queue interactions.
	Type LLMObsAnnotationQueueInteractionsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueInteractionsDataResponse instantiates a new LLMObsAnnotationQueueInteractionsDataResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueInteractionsDataResponse(attributes LLMObsAnnotationQueueInteractionsDataAttributesResponse, id string, typeVar LLMObsAnnotationQueueInteractionsType) *LLMObsAnnotationQueueInteractionsDataResponse {
	this := LLMObsAnnotationQueueInteractionsDataResponse{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewLLMObsAnnotationQueueInteractionsDataResponseWithDefaults instantiates a new LLMObsAnnotationQueueInteractionsDataResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueInteractionsDataResponseWithDefaults() *LLMObsAnnotationQueueInteractionsDataResponse {
	this := LLMObsAnnotationQueueInteractionsDataResponse{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *LLMObsAnnotationQueueInteractionsDataResponse) GetAttributes() LLMObsAnnotationQueueInteractionsDataAttributesResponse {
	if o == nil {
		var ret LLMObsAnnotationQueueInteractionsDataAttributesResponse
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueInteractionsDataResponse) GetAttributesOk() (*LLMObsAnnotationQueueInteractionsDataAttributesResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *LLMObsAnnotationQueueInteractionsDataResponse) SetAttributes(v LLMObsAnnotationQueueInteractionsDataAttributesResponse) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *LLMObsAnnotationQueueInteractionsDataResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueInteractionsDataResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsAnnotationQueueInteractionsDataResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *LLMObsAnnotationQueueInteractionsDataResponse) GetType() LLMObsAnnotationQueueInteractionsType {
	if o == nil {
		var ret LLMObsAnnotationQueueInteractionsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueInteractionsDataResponse) GetTypeOk() (*LLMObsAnnotationQueueInteractionsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsAnnotationQueueInteractionsDataResponse) SetType(v LLMObsAnnotationQueueInteractionsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueInteractionsDataResponse) MarshalJSON() ([]byte, error) {
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
func (o *LLMObsAnnotationQueueInteractionsDataResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *LLMObsAnnotationQueueInteractionsDataAttributesResponse `json:"attributes"`
		Id         *string                                                  `json:"id"`
		Type       *LLMObsAnnotationQueueInteractionsType                   `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
