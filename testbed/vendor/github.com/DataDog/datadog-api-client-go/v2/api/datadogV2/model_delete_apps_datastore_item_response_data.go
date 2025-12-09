// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeleteAppsDatastoreItemResponseData Data containing the identifier of the datastore item that was successfully deleted.
type DeleteAppsDatastoreItemResponseData struct {
	// The unique identifier of the item that was deleted.
	Id *string `json:"id,omitempty"`
	// The resource type for datastore items.
	Type DatastoreItemsDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeleteAppsDatastoreItemResponseData instantiates a new DeleteAppsDatastoreItemResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeleteAppsDatastoreItemResponseData(typeVar DatastoreItemsDataType) *DeleteAppsDatastoreItemResponseData {
	this := DeleteAppsDatastoreItemResponseData{}
	this.Type = typeVar
	return &this
}

// NewDeleteAppsDatastoreItemResponseDataWithDefaults instantiates a new DeleteAppsDatastoreItemResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeleteAppsDatastoreItemResponseDataWithDefaults() *DeleteAppsDatastoreItemResponseData {
	this := DeleteAppsDatastoreItemResponseData{}
	var typeVar DatastoreItemsDataType = DATASTOREITEMSDATATYPE_ITEMS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteAppsDatastoreItemResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAppsDatastoreItemResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteAppsDatastoreItemResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeleteAppsDatastoreItemResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *DeleteAppsDatastoreItemResponseData) GetType() DatastoreItemsDataType {
	if o == nil {
		var ret DatastoreItemsDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeleteAppsDatastoreItemResponseData) GetTypeOk() (*DatastoreItemsDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DeleteAppsDatastoreItemResponseData) SetType(v DatastoreItemsDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeleteAppsDatastoreItemResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeleteAppsDatastoreItemResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                 `json:"id,omitempty"`
		Type *DatastoreItemsDataType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.Id = all.Id
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
