// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventCreateRequest An event object.
type EventCreateRequest struct {
	// Event attributes.
	Attributes EventPayload `json:"attributes"`
	// Entity type.
	Type EventCreateRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventCreateRequest instantiates a new EventCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventCreateRequest(attributes EventPayload, typeVar EventCreateRequestType) *EventCreateRequest {
	this := EventCreateRequest{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewEventCreateRequestWithDefaults instantiates a new EventCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventCreateRequestWithDefaults() *EventCreateRequest {
	this := EventCreateRequest{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *EventCreateRequest) GetAttributes() EventPayload {
	if o == nil {
		var ret EventPayload
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetAttributesOk() (*EventPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *EventCreateRequest) SetAttributes(v EventPayload) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *EventCreateRequest) GetType() EventCreateRequestType {
	if o == nil {
		var ret EventCreateRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetTypeOk() (*EventCreateRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EventCreateRequest) SetType(v EventCreateRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *EventPayload           `json:"attributes"`
		Type       *EventCreateRequestType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
