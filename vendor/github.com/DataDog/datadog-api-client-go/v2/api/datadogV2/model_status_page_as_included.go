// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPageAsIncluded The included status page resource.
type StatusPageAsIncluded struct {
	// The attributes of a status page.
	Attributes *StatusPageAsIncludedAttributes `json:"attributes,omitempty"`
	// The ID of the status page.
	Id *uuid.UUID `json:"id,omitempty"`
	// The relationships of a status page.
	Relationships *StatusPageAsIncludedRelationships `json:"relationships,omitempty"`
	// Status pages resource type.
	Type StatusPageDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPageAsIncluded instantiates a new StatusPageAsIncluded object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPageAsIncluded(typeVar StatusPageDataType) *StatusPageAsIncluded {
	this := StatusPageAsIncluded{}
	this.Type = typeVar
	return &this
}

// NewStatusPageAsIncludedWithDefaults instantiates a new StatusPageAsIncluded object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPageAsIncludedWithDefaults() *StatusPageAsIncluded {
	this := StatusPageAsIncluded{}
	var typeVar StatusPageDataType = STATUSPAGEDATATYPE_STATUS_PAGES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *StatusPageAsIncluded) GetAttributes() StatusPageAsIncludedAttributes {
	if o == nil || o.Attributes == nil {
		var ret StatusPageAsIncludedAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageAsIncluded) GetAttributesOk() (*StatusPageAsIncludedAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *StatusPageAsIncluded) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given StatusPageAsIncludedAttributes and assigns it to the Attributes field.
func (o *StatusPageAsIncluded) SetAttributes(v StatusPageAsIncludedAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StatusPageAsIncluded) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageAsIncluded) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StatusPageAsIncluded) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *StatusPageAsIncluded) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StatusPageAsIncluded) GetRelationships() StatusPageAsIncludedRelationships {
	if o == nil || o.Relationships == nil {
		var ret StatusPageAsIncludedRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageAsIncluded) GetRelationshipsOk() (*StatusPageAsIncludedRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StatusPageAsIncluded) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given StatusPageAsIncludedRelationships and assigns it to the Relationships field.
func (o *StatusPageAsIncluded) SetRelationships(v StatusPageAsIncludedRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *StatusPageAsIncluded) GetType() StatusPageDataType {
	if o == nil {
		var ret StatusPageDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StatusPageAsIncluded) GetTypeOk() (*StatusPageDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StatusPageAsIncluded) SetType(v StatusPageDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPageAsIncluded) MarshalJSON() ([]byte, error) {
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
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatusPageAsIncluded) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *StatusPageAsIncludedAttributes    `json:"attributes,omitempty"`
		Id            *uuid.UUID                         `json:"id,omitempty"`
		Relationships *StatusPageAsIncludedRelationships `json:"relationships,omitempty"`
		Type          *StatusPageDataType                `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
