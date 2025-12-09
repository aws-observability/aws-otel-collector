// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSEventBridgeDeleteResponseData Amazon EventBridge delete response data.
type AWSEventBridgeDeleteResponseData struct {
	// The EventBridge source delete response attributes.
	Attributes AWSEventBridgeDeleteResponseAttributes `json:"attributes"`
	// The ID of the Amazon EventBridge list response data.
	Id *string `json:"id,omitempty"`
	// Amazon EventBridge resource type.
	Type AWSEventBridgeType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSEventBridgeDeleteResponseData instantiates a new AWSEventBridgeDeleteResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSEventBridgeDeleteResponseData(attributes AWSEventBridgeDeleteResponseAttributes, typeVar AWSEventBridgeType) *AWSEventBridgeDeleteResponseData {
	this := AWSEventBridgeDeleteResponseData{}
	this.Attributes = attributes
	var id string = "delete_event_bridge"
	this.Id = &id
	this.Type = typeVar
	return &this
}

// NewAWSEventBridgeDeleteResponseDataWithDefaults instantiates a new AWSEventBridgeDeleteResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSEventBridgeDeleteResponseDataWithDefaults() *AWSEventBridgeDeleteResponseData {
	this := AWSEventBridgeDeleteResponseData{}
	var id string = "delete_event_bridge"
	this.Id = &id
	var typeVar AWSEventBridgeType = AWSEVENTBRIDGETYPE_EVENT_BRIDGE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AWSEventBridgeDeleteResponseData) GetAttributes() AWSEventBridgeDeleteResponseAttributes {
	if o == nil {
		var ret AWSEventBridgeDeleteResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeDeleteResponseData) GetAttributesOk() (*AWSEventBridgeDeleteResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AWSEventBridgeDeleteResponseData) SetAttributes(v AWSEventBridgeDeleteResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AWSEventBridgeDeleteResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeDeleteResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AWSEventBridgeDeleteResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AWSEventBridgeDeleteResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *AWSEventBridgeDeleteResponseData) GetType() AWSEventBridgeType {
	if o == nil {
		var ret AWSEventBridgeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeDeleteResponseData) GetTypeOk() (*AWSEventBridgeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AWSEventBridgeDeleteResponseData) SetType(v AWSEventBridgeType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSEventBridgeDeleteResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSEventBridgeDeleteResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AWSEventBridgeDeleteResponseAttributes `json:"attributes"`
		Id         *string                                 `json:"id,omitempty"`
		Type       *AWSEventBridgeType                     `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = all.Id
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
