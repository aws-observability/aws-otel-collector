// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleDataRelationshipsLayersDataItems Relates a layer to this schedule, identified by `id` and `type` (must be `layers`).
type ScheduleDataRelationshipsLayersDataItems struct {
	// The unique identifier of the layer in this relationship.
	Id string `json:"id"`
	// Layers resource type.
	Type ScheduleDataRelationshipsLayersDataItemsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleDataRelationshipsLayersDataItems instantiates a new ScheduleDataRelationshipsLayersDataItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleDataRelationshipsLayersDataItems(id string, typeVar ScheduleDataRelationshipsLayersDataItemsType) *ScheduleDataRelationshipsLayersDataItems {
	this := ScheduleDataRelationshipsLayersDataItems{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewScheduleDataRelationshipsLayersDataItemsWithDefaults instantiates a new ScheduleDataRelationshipsLayersDataItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleDataRelationshipsLayersDataItemsWithDefaults() *ScheduleDataRelationshipsLayersDataItems {
	this := ScheduleDataRelationshipsLayersDataItems{}
	var typeVar ScheduleDataRelationshipsLayersDataItemsType = SCHEDULEDATARELATIONSHIPSLAYERSDATAITEMSTYPE_LAYERS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *ScheduleDataRelationshipsLayersDataItems) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScheduleDataRelationshipsLayersDataItems) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ScheduleDataRelationshipsLayersDataItems) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ScheduleDataRelationshipsLayersDataItems) GetType() ScheduleDataRelationshipsLayersDataItemsType {
	if o == nil {
		var ret ScheduleDataRelationshipsLayersDataItemsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleDataRelationshipsLayersDataItems) GetTypeOk() (*ScheduleDataRelationshipsLayersDataItemsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ScheduleDataRelationshipsLayersDataItems) SetType(v ScheduleDataRelationshipsLayersDataItemsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleDataRelationshipsLayersDataItems) MarshalJSON() ([]byte, error) {
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
func (o *ScheduleDataRelationshipsLayersDataItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                                       `json:"id"`
		Type *ScheduleDataRelationshipsLayersDataItemsType `json:"type"`
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
