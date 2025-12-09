// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListConnectionsResponseDataAttributesConnectionsItemsJoin
type ListConnectionsResponseDataAttributesConnectionsItemsJoin struct {
	//
	Attribute *string `json:"attribute,omitempty"`
	//
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListConnectionsResponseDataAttributesConnectionsItemsJoin instantiates a new ListConnectionsResponseDataAttributesConnectionsItemsJoin object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListConnectionsResponseDataAttributesConnectionsItemsJoin() *ListConnectionsResponseDataAttributesConnectionsItemsJoin {
	this := ListConnectionsResponseDataAttributesConnectionsItemsJoin{}
	return &this
}

// NewListConnectionsResponseDataAttributesConnectionsItemsJoinWithDefaults instantiates a new ListConnectionsResponseDataAttributesConnectionsItemsJoin object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListConnectionsResponseDataAttributesConnectionsItemsJoinWithDefaults() *ListConnectionsResponseDataAttributesConnectionsItemsJoin {
	this := ListConnectionsResponseDataAttributesConnectionsItemsJoin{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsJoin) GetAttribute() string {
	if o == nil || o.Attribute == nil {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsJoin) GetAttributeOk() (*string, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsJoin) HasAttribute() bool {
	return o != nil && o.Attribute != nil
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsJoin) SetAttribute(v string) {
	o.Attribute = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsJoin) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsJoin) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsJoin) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsJoin) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListConnectionsResponseDataAttributesConnectionsItemsJoin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attribute != nil {
		toSerialize["attribute"] = o.Attribute
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsJoin) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute *string `json:"attribute,omitempty"`
		Type      *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute", "type"})
	} else {
		return err
	}
	o.Attribute = all.Attribute
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
