// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomConnection The definition of `CustomConnection` object.
type CustomConnection struct {
	// The definition of `CustomConnectionAttributes` object.
	Attributes *CustomConnectionAttributes `json:"attributes,omitempty"`
	// The `CustomConnection` `id`.
	Id *string `json:"id,omitempty"`
	// The definition of `CustomConnectionType` object.
	Type *CustomConnectionType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomConnection instantiates a new CustomConnection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomConnection() *CustomConnection {
	this := CustomConnection{}
	var typeVar CustomConnectionType = CUSTOMCONNECTIONTYPE_CUSTOM_CONNECTIONS
	this.Type = &typeVar
	return &this
}

// NewCustomConnectionWithDefaults instantiates a new CustomConnection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomConnectionWithDefaults() *CustomConnection {
	this := CustomConnection{}
	var typeVar CustomConnectionType = CUSTOMCONNECTIONTYPE_CUSTOM_CONNECTIONS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomConnection) GetAttributes() CustomConnectionAttributes {
	if o == nil || o.Attributes == nil {
		var ret CustomConnectionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomConnection) GetAttributesOk() (*CustomConnectionAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomConnection) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CustomConnectionAttributes and assigns it to the Attributes field.
func (o *CustomConnection) SetAttributes(v CustomConnectionAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomConnection) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomConnection) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomConnection) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomConnection) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomConnection) GetType() CustomConnectionType {
	if o == nil || o.Type == nil {
		var ret CustomConnectionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomConnection) GetTypeOk() (*CustomConnectionType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomConnection) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given CustomConnectionType and assigns it to the Type field.
func (o *CustomConnection) SetType(v CustomConnectionType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *CustomConnection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CustomConnectionAttributes `json:"attributes,omitempty"`
		Id         *string                     `json:"id,omitempty"`
		Type       *CustomConnectionType       `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
