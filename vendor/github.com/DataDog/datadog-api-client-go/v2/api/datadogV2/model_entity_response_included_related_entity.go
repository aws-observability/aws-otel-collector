// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedRelatedEntity Included related entity.
type EntityResponseIncludedRelatedEntity struct {
	// Related entity attributes.
	Attributes *EntityResponseIncludedRelatedEntityAttributes `json:"attributes,omitempty"`
	// Entity UUID.
	Id *string `json:"id,omitempty"`
	// Included related entity meta.
	Meta *EntityResponseIncludedRelatedEntityMeta `json:"meta,omitempty"`
	// Related entity.
	Type *EntityResponseIncludedRelatedEntityType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityResponseIncludedRelatedEntity instantiates a new EntityResponseIncludedRelatedEntity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityResponseIncludedRelatedEntity() *EntityResponseIncludedRelatedEntity {
	this := EntityResponseIncludedRelatedEntity{}
	return &this
}

// NewEntityResponseIncludedRelatedEntityWithDefaults instantiates a new EntityResponseIncludedRelatedEntity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityResponseIncludedRelatedEntityWithDefaults() *EntityResponseIncludedRelatedEntity {
	this := EntityResponseIncludedRelatedEntity{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntity) GetAttributes() EntityResponseIncludedRelatedEntityAttributes {
	if o == nil || o.Attributes == nil {
		var ret EntityResponseIncludedRelatedEntityAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntity) GetAttributesOk() (*EntityResponseIncludedRelatedEntityAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntity) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given EntityResponseIncludedRelatedEntityAttributes and assigns it to the Attributes field.
func (o *EntityResponseIncludedRelatedEntity) SetAttributes(v EntityResponseIncludedRelatedEntityAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntity) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntityResponseIncludedRelatedEntity) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntity) GetMeta() EntityResponseIncludedRelatedEntityMeta {
	if o == nil || o.Meta == nil {
		var ret EntityResponseIncludedRelatedEntityMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntity) GetMetaOk() (*EntityResponseIncludedRelatedEntityMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntity) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given EntityResponseIncludedRelatedEntityMeta and assigns it to the Meta field.
func (o *EntityResponseIncludedRelatedEntity) SetMeta(v EntityResponseIncludedRelatedEntityMeta) {
	o.Meta = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntity) GetType() EntityResponseIncludedRelatedEntityType {
	if o == nil || o.Type == nil {
		var ret EntityResponseIncludedRelatedEntityType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntity) GetTypeOk() (*EntityResponseIncludedRelatedEntityType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntity) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given EntityResponseIncludedRelatedEntityType and assigns it to the Type field.
func (o *EntityResponseIncludedRelatedEntity) SetType(v EntityResponseIncludedRelatedEntityType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityResponseIncludedRelatedEntity) MarshalJSON() ([]byte, error) {
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
func (o *EntityResponseIncludedRelatedEntity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *EntityResponseIncludedRelatedEntityAttributes `json:"attributes,omitempty"`
		Id         *string                                        `json:"id,omitempty"`
		Meta       *EntityResponseIncludedRelatedEntityMeta       `json:"meta,omitempty"`
		Type       *EntityResponseIncludedRelatedEntityType       `json:"type,omitempty"`
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
