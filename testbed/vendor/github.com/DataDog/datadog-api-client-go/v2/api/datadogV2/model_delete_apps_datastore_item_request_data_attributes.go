// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeleteAppsDatastoreItemRequestDataAttributes Attributes specifying which datastore item to delete by its primary key.
type DeleteAppsDatastoreItemRequestDataAttributes struct {
	// Optional unique identifier of the item to delete.
	Id *string `json:"id,omitempty"`
	// The primary key value that identifies the item to delete. Cannot exceed 256 characters.
	ItemKey string `json:"item_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeleteAppsDatastoreItemRequestDataAttributes instantiates a new DeleteAppsDatastoreItemRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeleteAppsDatastoreItemRequestDataAttributes(itemKey string) *DeleteAppsDatastoreItemRequestDataAttributes {
	this := DeleteAppsDatastoreItemRequestDataAttributes{}
	this.ItemKey = itemKey
	return &this
}

// NewDeleteAppsDatastoreItemRequestDataAttributesWithDefaults instantiates a new DeleteAppsDatastoreItemRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeleteAppsDatastoreItemRequestDataAttributesWithDefaults() *DeleteAppsDatastoreItemRequestDataAttributes {
	this := DeleteAppsDatastoreItemRequestDataAttributes{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteAppsDatastoreItemRequestDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAppsDatastoreItemRequestDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteAppsDatastoreItemRequestDataAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeleteAppsDatastoreItemRequestDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetItemKey returns the ItemKey field value.
func (o *DeleteAppsDatastoreItemRequestDataAttributes) GetItemKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ItemKey
}

// GetItemKeyOk returns a tuple with the ItemKey field value
// and a boolean to check if the value has been set.
func (o *DeleteAppsDatastoreItemRequestDataAttributes) GetItemKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemKey, true
}

// SetItemKey sets field value.
func (o *DeleteAppsDatastoreItemRequestDataAttributes) SetItemKey(v string) {
	o.ItemKey = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeleteAppsDatastoreItemRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["item_key"] = o.ItemKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeleteAppsDatastoreItemRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id      *string `json:"id,omitempty"`
		ItemKey *string `json:"item_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ItemKey == nil {
		return fmt.Errorf("required field item_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "item_key"})
	} else {
		return err
	}
	o.Id = all.Id
	o.ItemKey = *all.ItemKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
