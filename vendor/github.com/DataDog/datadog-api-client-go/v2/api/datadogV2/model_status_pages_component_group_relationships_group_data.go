// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesComponentGroupRelationshipsGroupData The data object identifying the parent group of a component group.
type StatusPagesComponentGroupRelationshipsGroupData struct {
	// The ID of the parent group.
	Id uuid.UUID `json:"id"`
	// Components resource type.
	Type StatusPagesComponentGroupType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPagesComponentGroupRelationshipsGroupData instantiates a new StatusPagesComponentGroupRelationshipsGroupData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPagesComponentGroupRelationshipsGroupData(id uuid.UUID, typeVar StatusPagesComponentGroupType) *StatusPagesComponentGroupRelationshipsGroupData {
	this := StatusPagesComponentGroupRelationshipsGroupData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewStatusPagesComponentGroupRelationshipsGroupDataWithDefaults instantiates a new StatusPagesComponentGroupRelationshipsGroupData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPagesComponentGroupRelationshipsGroupDataWithDefaults() *StatusPagesComponentGroupRelationshipsGroupData {
	this := StatusPagesComponentGroupRelationshipsGroupData{}
	var typeVar StatusPagesComponentGroupType = STATUSPAGESCOMPONENTGROUPTYPE_COMPONENTS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *StatusPagesComponentGroupRelationshipsGroupData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentGroupRelationshipsGroupData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *StatusPagesComponentGroupRelationshipsGroupData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *StatusPagesComponentGroupRelationshipsGroupData) GetType() StatusPagesComponentGroupType {
	if o == nil {
		var ret StatusPagesComponentGroupType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentGroupRelationshipsGroupData) GetTypeOk() (*StatusPagesComponentGroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StatusPagesComponentGroupRelationshipsGroupData) SetType(v StatusPagesComponentGroupType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPagesComponentGroupRelationshipsGroupData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatusPagesComponentGroupRelationshipsGroupData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *uuid.UUID                     `json:"id"`
		Type *StatusPagesComponentGroupType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
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

// NullableStatusPagesComponentGroupRelationshipsGroupData handles when a null is used for StatusPagesComponentGroupRelationshipsGroupData.
type NullableStatusPagesComponentGroupRelationshipsGroupData struct {
	value *StatusPagesComponentGroupRelationshipsGroupData
	isSet bool
}

// Get returns the associated value.
func (v NullableStatusPagesComponentGroupRelationshipsGroupData) Get() *StatusPagesComponentGroupRelationshipsGroupData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableStatusPagesComponentGroupRelationshipsGroupData) Set(val *StatusPagesComponentGroupRelationshipsGroupData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableStatusPagesComponentGroupRelationshipsGroupData) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableStatusPagesComponentGroupRelationshipsGroupData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableStatusPagesComponentGroupRelationshipsGroupData initializes the struct as if Set has been called.
func NewNullableStatusPagesComponentGroupRelationshipsGroupData(val *StatusPagesComponentGroupRelationshipsGroupData) *NullableStatusPagesComponentGroupRelationshipsGroupData {
	return &NullableStatusPagesComponentGroupRelationshipsGroupData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableStatusPagesComponentGroupRelationshipsGroupData) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableStatusPagesComponentGroupRelationshipsGroupData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
