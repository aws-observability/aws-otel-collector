// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationRelationships Relation relationships.
type RelationRelationships struct {
	// Relation to entity.
	FromEntity *RelationToEntity `json:"fromEntity,omitempty"`
	// Relation to entity.
	ToEntity *RelationToEntity `json:"toEntity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRelationRelationships instantiates a new RelationRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRelationRelationships() *RelationRelationships {
	this := RelationRelationships{}
	return &this
}

// NewRelationRelationshipsWithDefaults instantiates a new RelationRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRelationRelationshipsWithDefaults() *RelationRelationships {
	this := RelationRelationships{}
	return &this
}

// GetFromEntity returns the FromEntity field value if set, zero value otherwise.
func (o *RelationRelationships) GetFromEntity() RelationToEntity {
	if o == nil || o.FromEntity == nil {
		var ret RelationToEntity
		return ret
	}
	return *o.FromEntity
}

// GetFromEntityOk returns a tuple with the FromEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationRelationships) GetFromEntityOk() (*RelationToEntity, bool) {
	if o == nil || o.FromEntity == nil {
		return nil, false
	}
	return o.FromEntity, true
}

// HasFromEntity returns a boolean if a field has been set.
func (o *RelationRelationships) HasFromEntity() bool {
	return o != nil && o.FromEntity != nil
}

// SetFromEntity gets a reference to the given RelationToEntity and assigns it to the FromEntity field.
func (o *RelationRelationships) SetFromEntity(v RelationToEntity) {
	o.FromEntity = &v
}

// GetToEntity returns the ToEntity field value if set, zero value otherwise.
func (o *RelationRelationships) GetToEntity() RelationToEntity {
	if o == nil || o.ToEntity == nil {
		var ret RelationToEntity
		return ret
	}
	return *o.ToEntity
}

// GetToEntityOk returns a tuple with the ToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationRelationships) GetToEntityOk() (*RelationToEntity, bool) {
	if o == nil || o.ToEntity == nil {
		return nil, false
	}
	return o.ToEntity, true
}

// HasToEntity returns a boolean if a field has been set.
func (o *RelationRelationships) HasToEntity() bool {
	return o != nil && o.ToEntity != nil
}

// SetToEntity gets a reference to the given RelationToEntity and assigns it to the ToEntity field.
func (o *RelationRelationships) SetToEntity(v RelationToEntity) {
	o.ToEntity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RelationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FromEntity != nil {
		toSerialize["fromEntity"] = o.FromEntity
	}
	if o.ToEntity != nil {
		toSerialize["toEntity"] = o.ToEntity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RelationRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FromEntity *RelationToEntity `json:"fromEntity,omitempty"`
		ToEntity   *RelationToEntity `json:"toEntity,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fromEntity", "toEntity"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.FromEntity != nil && all.FromEntity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FromEntity = all.FromEntity
	if all.ToEntity != nil && all.ToEntity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ToEntity = all.ToEntity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
