// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentIncluded The definition of `DeploymentIncluded` object.
type DeploymentIncluded struct {
	// The definition of `DeploymentIncludedAttributes` object.
	Attributes *DeploymentIncludedAttributes `json:"attributes,omitempty"`
	// The `DeploymentIncluded` `id`.
	Id *string `json:"id,omitempty"`
	// The definition of `DeploymentIncludedMeta` object.
	Meta *DeploymentIncludedMeta `json:"meta,omitempty"`
	// The definition of `DeploymentIncludedType` object.
	Type *DeploymentIncludedType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentIncluded instantiates a new DeploymentIncluded object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentIncluded() *DeploymentIncluded {
	this := DeploymentIncluded{}
	var typeVar DeploymentIncludedType = DEPLOYMENTINCLUDEDTYPE_DEPLOYMENT
	this.Type = &typeVar
	return &this
}

// NewDeploymentIncludedWithDefaults instantiates a new DeploymentIncluded object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentIncludedWithDefaults() *DeploymentIncluded {
	this := DeploymentIncluded{}
	var typeVar DeploymentIncludedType = DEPLOYMENTINCLUDEDTYPE_DEPLOYMENT
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DeploymentIncluded) GetAttributes() DeploymentIncludedAttributes {
	if o == nil || o.Attributes == nil {
		var ret DeploymentIncludedAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentIncluded) GetAttributesOk() (*DeploymentIncludedAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DeploymentIncluded) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given DeploymentIncludedAttributes and assigns it to the Attributes field.
func (o *DeploymentIncluded) SetAttributes(v DeploymentIncludedAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeploymentIncluded) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentIncluded) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeploymentIncluded) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeploymentIncluded) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DeploymentIncluded) GetMeta() DeploymentIncludedMeta {
	if o == nil || o.Meta == nil {
		var ret DeploymentIncludedMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentIncluded) GetMetaOk() (*DeploymentIncludedMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DeploymentIncluded) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given DeploymentIncludedMeta and assigns it to the Meta field.
func (o *DeploymentIncluded) SetMeta(v DeploymentIncludedMeta) {
	o.Meta = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeploymentIncluded) GetType() DeploymentIncludedType {
	if o == nil || o.Type == nil {
		var ret DeploymentIncludedType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentIncluded) GetTypeOk() (*DeploymentIncludedType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeploymentIncluded) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given DeploymentIncludedType and assigns it to the Type field.
func (o *DeploymentIncluded) SetType(v DeploymentIncludedType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentIncluded) MarshalJSON() ([]byte, error) {
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
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
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
func (o *DeploymentIncluded) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DeploymentIncludedAttributes `json:"attributes,omitempty"`
		Id         *string                       `json:"id,omitempty"`
		Meta       *DeploymentIncludedMeta       `json:"meta,omitempty"`
		Type       *DeploymentIncludedType       `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "meta", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
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
