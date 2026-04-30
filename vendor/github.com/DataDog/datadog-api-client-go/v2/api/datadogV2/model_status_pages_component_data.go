// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesComponentData The data object for a component.
type StatusPagesComponentData struct {
	// The attributes of a component.
	Attributes *StatusPagesComponentDataAttributes `json:"attributes,omitempty"`
	// The ID of the component.
	Id *uuid.UUID `json:"id,omitempty"`
	// The relationships of a component.
	Relationships *StatusPagesComponentDataRelationships `json:"relationships,omitempty"`
	// Components resource type.
	Type StatusPagesComponentGroupType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPagesComponentData instantiates a new StatusPagesComponentData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPagesComponentData(typeVar StatusPagesComponentGroupType) *StatusPagesComponentData {
	this := StatusPagesComponentData{}
	this.Type = typeVar
	return &this
}

// NewStatusPagesComponentDataWithDefaults instantiates a new StatusPagesComponentData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPagesComponentDataWithDefaults() *StatusPagesComponentData {
	this := StatusPagesComponentData{}
	var typeVar StatusPagesComponentGroupType = STATUSPAGESCOMPONENTGROUPTYPE_COMPONENTS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *StatusPagesComponentData) GetAttributes() StatusPagesComponentDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret StatusPagesComponentDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentData) GetAttributesOk() (*StatusPagesComponentDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *StatusPagesComponentData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given StatusPagesComponentDataAttributes and assigns it to the Attributes field.
func (o *StatusPagesComponentData) SetAttributes(v StatusPagesComponentDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StatusPagesComponentData) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StatusPagesComponentData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *StatusPagesComponentData) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StatusPagesComponentData) GetRelationships() StatusPagesComponentDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret StatusPagesComponentDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentData) GetRelationshipsOk() (*StatusPagesComponentDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StatusPagesComponentData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given StatusPagesComponentDataRelationships and assigns it to the Relationships field.
func (o *StatusPagesComponentData) SetRelationships(v StatusPagesComponentDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *StatusPagesComponentData) GetType() StatusPagesComponentGroupType {
	if o == nil {
		var ret StatusPagesComponentGroupType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentData) GetTypeOk() (*StatusPagesComponentGroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StatusPagesComponentData) SetType(v StatusPagesComponentGroupType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPagesComponentData) MarshalJSON() ([]byte, error) {
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
func (o *StatusPagesComponentData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *StatusPagesComponentDataAttributes    `json:"attributes,omitempty"`
		Id            *uuid.UUID                             `json:"id,omitempty"`
		Relationships *StatusPagesComponentDataRelationships `json:"relationships,omitempty"`
		Type          *StatusPagesComponentGroupType         `json:"type"`
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
