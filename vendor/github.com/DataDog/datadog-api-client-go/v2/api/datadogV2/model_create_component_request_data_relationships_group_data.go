// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateComponentRequestDataRelationshipsGroupData The data object identifying the group to create the component within.
type CreateComponentRequestDataRelationshipsGroupData struct {
	// The ID of the group.
	Id uuid.UUID `json:"id"`
	// Components resource type.
	Type StatusPagesComponentGroupType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateComponentRequestDataRelationshipsGroupData instantiates a new CreateComponentRequestDataRelationshipsGroupData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateComponentRequestDataRelationshipsGroupData(id uuid.UUID, typeVar StatusPagesComponentGroupType) *CreateComponentRequestDataRelationshipsGroupData {
	this := CreateComponentRequestDataRelationshipsGroupData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCreateComponentRequestDataRelationshipsGroupDataWithDefaults instantiates a new CreateComponentRequestDataRelationshipsGroupData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateComponentRequestDataRelationshipsGroupDataWithDefaults() *CreateComponentRequestDataRelationshipsGroupData {
	this := CreateComponentRequestDataRelationshipsGroupData{}
	var typeVar StatusPagesComponentGroupType = STATUSPAGESCOMPONENTGROUPTYPE_COMPONENTS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *CreateComponentRequestDataRelationshipsGroupData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateComponentRequestDataRelationshipsGroupData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CreateComponentRequestDataRelationshipsGroupData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CreateComponentRequestDataRelationshipsGroupData) GetType() StatusPagesComponentGroupType {
	if o == nil {
		var ret StatusPagesComponentGroupType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateComponentRequestDataRelationshipsGroupData) GetTypeOk() (*StatusPagesComponentGroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateComponentRequestDataRelationshipsGroupData) SetType(v StatusPagesComponentGroupType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateComponentRequestDataRelationshipsGroupData) MarshalJSON() ([]byte, error) {
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
func (o *CreateComponentRequestDataRelationshipsGroupData) UnmarshalJSON(bytes []byte) (err error) {
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

// NullableCreateComponentRequestDataRelationshipsGroupData handles when a null is used for CreateComponentRequestDataRelationshipsGroupData.
type NullableCreateComponentRequestDataRelationshipsGroupData struct {
	value *CreateComponentRequestDataRelationshipsGroupData
	isSet bool
}

// Get returns the associated value.
func (v NullableCreateComponentRequestDataRelationshipsGroupData) Get() *CreateComponentRequestDataRelationshipsGroupData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCreateComponentRequestDataRelationshipsGroupData) Set(val *CreateComponentRequestDataRelationshipsGroupData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCreateComponentRequestDataRelationshipsGroupData) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableCreateComponentRequestDataRelationshipsGroupData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCreateComponentRequestDataRelationshipsGroupData initializes the struct as if Set has been called.
func NewNullableCreateComponentRequestDataRelationshipsGroupData(val *CreateComponentRequestDataRelationshipsGroupData) *NullableCreateComponentRequestDataRelationshipsGroupData {
	return &NullableCreateComponentRequestDataRelationshipsGroupData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCreateComponentRequestDataRelationshipsGroupData) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCreateComponentRequestDataRelationshipsGroupData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
