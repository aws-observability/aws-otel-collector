// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseDataRelationshipsRelatedEntitiesDataItems
type EntityResponseDataRelationshipsRelatedEntitiesDataItems struct {
	//
	Id string `json:"id"`
	// Related entity resource type.
	Type EntityResponseDataRelationshipsRelatedEntitiesDataItemsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityResponseDataRelationshipsRelatedEntitiesDataItems instantiates a new EntityResponseDataRelationshipsRelatedEntitiesDataItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityResponseDataRelationshipsRelatedEntitiesDataItems(id string, typeVar EntityResponseDataRelationshipsRelatedEntitiesDataItemsType) *EntityResponseDataRelationshipsRelatedEntitiesDataItems {
	this := EntityResponseDataRelationshipsRelatedEntitiesDataItems{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewEntityResponseDataRelationshipsRelatedEntitiesDataItemsWithDefaults instantiates a new EntityResponseDataRelationshipsRelatedEntitiesDataItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityResponseDataRelationshipsRelatedEntitiesDataItemsWithDefaults() *EntityResponseDataRelationshipsRelatedEntitiesDataItems {
	this := EntityResponseDataRelationshipsRelatedEntitiesDataItems{}
	var typeVar EntityResponseDataRelationshipsRelatedEntitiesDataItemsType = ENTITYRESPONSEDATARELATIONSHIPSRELATEDENTITIESDATAITEMSTYPE_RELATEDENTITY
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *EntityResponseDataRelationshipsRelatedEntitiesDataItems) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntityResponseDataRelationshipsRelatedEntitiesDataItems) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *EntityResponseDataRelationshipsRelatedEntitiesDataItems) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *EntityResponseDataRelationshipsRelatedEntitiesDataItems) GetType() EntityResponseDataRelationshipsRelatedEntitiesDataItemsType {
	if o == nil {
		var ret EntityResponseDataRelationshipsRelatedEntitiesDataItemsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntityResponseDataRelationshipsRelatedEntitiesDataItems) GetTypeOk() (*EntityResponseDataRelationshipsRelatedEntitiesDataItemsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EntityResponseDataRelationshipsRelatedEntitiesDataItems) SetType(v EntityResponseDataRelationshipsRelatedEntitiesDataItemsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityResponseDataRelationshipsRelatedEntitiesDataItems) MarshalJSON() ([]byte, error) {
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
func (o *EntityResponseDataRelationshipsRelatedEntitiesDataItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                                                      `json:"id"`
		Type *EntityResponseDataRelationshipsRelatedEntitiesDataItemsType `json:"type"`
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
