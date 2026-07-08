// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ManagedOrgsData The managed organizations resource.
type ManagedOrgsData struct {
	// The UUID of the current organization.
	Id uuid.UUID `json:"id"`
	// Relationships of the managed organizations resource.
	Relationships ManagedOrgsRelationships `json:"relationships"`
	// The resource type for managed organizations.
	Type ManagedOrgsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewManagedOrgsData instantiates a new ManagedOrgsData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewManagedOrgsData(id uuid.UUID, relationships ManagedOrgsRelationships, typeVar ManagedOrgsType) *ManagedOrgsData {
	this := ManagedOrgsData{}
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewManagedOrgsDataWithDefaults instantiates a new ManagedOrgsData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewManagedOrgsDataWithDefaults() *ManagedOrgsData {
	this := ManagedOrgsData{}
	return &this
}

// GetId returns the Id field value.
func (o *ManagedOrgsData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ManagedOrgsData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ManagedOrgsData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *ManagedOrgsData) GetRelationships() ManagedOrgsRelationships {
	if o == nil {
		var ret ManagedOrgsRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *ManagedOrgsData) GetRelationshipsOk() (*ManagedOrgsRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *ManagedOrgsData) SetRelationships(v ManagedOrgsRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *ManagedOrgsData) GetType() ManagedOrgsType {
	if o == nil {
		var ret ManagedOrgsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManagedOrgsData) GetTypeOk() (*ManagedOrgsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ManagedOrgsData) SetType(v ManagedOrgsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ManagedOrgsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ManagedOrgsData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *uuid.UUID                `json:"id"`
		Relationships *ManagedOrgsRelationships `json:"relationships"`
		Type          *ManagedOrgsType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	if all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = *all.Relationships
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
