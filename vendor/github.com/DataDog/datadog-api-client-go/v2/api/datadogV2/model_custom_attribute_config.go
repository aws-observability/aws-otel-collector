// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomAttributeConfig The definition of `CustomAttributeConfig` object.
type CustomAttributeConfig struct {
	// Custom attribute resource attributes
	Attributes *CustomAttributeConfigResourceAttributes `json:"attributes,omitempty"`
	// Custom attribute configs identifier
	Id *string `json:"id,omitempty"`
	// Custom attributes config JSON:API resource type
	Type *CustomAttributeConfigResourceType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomAttributeConfig instantiates a new CustomAttributeConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomAttributeConfig() *CustomAttributeConfig {
	this := CustomAttributeConfig{}
	var typeVar CustomAttributeConfigResourceType = CUSTOMATTRIBUTECONFIGRESOURCETYPE_CUSTOM_ATTRIBUTE
	this.Type = &typeVar
	return &this
}

// NewCustomAttributeConfigWithDefaults instantiates a new CustomAttributeConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomAttributeConfigWithDefaults() *CustomAttributeConfig {
	this := CustomAttributeConfig{}
	var typeVar CustomAttributeConfigResourceType = CUSTOMATTRIBUTECONFIGRESOURCETYPE_CUSTOM_ATTRIBUTE
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomAttributeConfig) GetAttributes() CustomAttributeConfigResourceAttributes {
	if o == nil || o.Attributes == nil {
		var ret CustomAttributeConfigResourceAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfig) GetAttributesOk() (*CustomAttributeConfigResourceAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomAttributeConfig) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CustomAttributeConfigResourceAttributes and assigns it to the Attributes field.
func (o *CustomAttributeConfig) SetAttributes(v CustomAttributeConfigResourceAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomAttributeConfig) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfig) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomAttributeConfig) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomAttributeConfig) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomAttributeConfig) GetType() CustomAttributeConfigResourceType {
	if o == nil || o.Type == nil {
		var ret CustomAttributeConfigResourceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfig) GetTypeOk() (*CustomAttributeConfigResourceType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomAttributeConfig) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given CustomAttributeConfigResourceType and assigns it to the Type field.
func (o *CustomAttributeConfig) SetType(v CustomAttributeConfigResourceType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomAttributeConfig) MarshalJSON() ([]byte, error) {
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
func (o *CustomAttributeConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CustomAttributeConfigResourceAttributes `json:"attributes,omitempty"`
		Id         *string                                  `json:"id,omitempty"`
		Type       *CustomAttributeConfigResourceType       `json:"type,omitempty"`
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
