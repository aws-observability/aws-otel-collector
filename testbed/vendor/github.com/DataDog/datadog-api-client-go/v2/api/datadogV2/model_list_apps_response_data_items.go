// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAppsResponseDataItems An app definition object. This contains only basic information about the app such as ID, name, and tags.
type ListAppsResponseDataItems struct {
	// Basic information about the app such as name, description, and tags.
	Attributes ListAppsResponseDataItemsAttributes `json:"attributes"`
	// The ID of the app.
	Id uuid.UUID `json:"id"`
	// Metadata of an app.
	Meta *AppMeta `json:"meta,omitempty"`
	// The app's publication information.
	Relationships *ListAppsResponseDataItemsRelationships `json:"relationships,omitempty"`
	// The app definition type.
	Type AppDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListAppsResponseDataItems instantiates a new ListAppsResponseDataItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListAppsResponseDataItems(attributes ListAppsResponseDataItemsAttributes, id uuid.UUID, typeVar AppDefinitionType) *ListAppsResponseDataItems {
	this := ListAppsResponseDataItems{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewListAppsResponseDataItemsWithDefaults instantiates a new ListAppsResponseDataItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListAppsResponseDataItemsWithDefaults() *ListAppsResponseDataItems {
	this := ListAppsResponseDataItems{}
	var typeVar AppDefinitionType = APPDEFINITIONTYPE_APPDEFINITIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ListAppsResponseDataItems) GetAttributes() ListAppsResponseDataItemsAttributes {
	if o == nil {
		var ret ListAppsResponseDataItemsAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ListAppsResponseDataItems) GetAttributesOk() (*ListAppsResponseDataItemsAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ListAppsResponseDataItems) SetAttributes(v ListAppsResponseDataItemsAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *ListAppsResponseDataItems) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListAppsResponseDataItems) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ListAppsResponseDataItems) SetId(v uuid.UUID) {
	o.Id = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListAppsResponseDataItems) GetMeta() AppMeta {
	if o == nil || o.Meta == nil {
		var ret AppMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponseDataItems) GetMetaOk() (*AppMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListAppsResponseDataItems) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given AppMeta and assigns it to the Meta field.
func (o *ListAppsResponseDataItems) SetMeta(v AppMeta) {
	o.Meta = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ListAppsResponseDataItems) GetRelationships() ListAppsResponseDataItemsRelationships {
	if o == nil || o.Relationships == nil {
		var ret ListAppsResponseDataItemsRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponseDataItems) GetRelationshipsOk() (*ListAppsResponseDataItemsRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ListAppsResponseDataItems) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ListAppsResponseDataItemsRelationships and assigns it to the Relationships field.
func (o *ListAppsResponseDataItems) SetRelationships(v ListAppsResponseDataItemsRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *ListAppsResponseDataItems) GetType() AppDefinitionType {
	if o == nil {
		var ret AppDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListAppsResponseDataItems) GetTypeOk() (*AppDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ListAppsResponseDataItems) SetType(v AppDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListAppsResponseDataItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
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
func (o *ListAppsResponseDataItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ListAppsResponseDataItemsAttributes    `json:"attributes"`
		Id            *uuid.UUID                              `json:"id"`
		Meta          *AppMeta                                `json:"meta,omitempty"`
		Relationships *ListAppsResponseDataItemsRelationships `json:"relationships,omitempty"`
		Type          *AppDefinitionType                      `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "meta", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
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
