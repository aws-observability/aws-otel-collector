// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BulkDeleteAppsDatastoreItemsRequestDataAttributes Attributes of request data to delete items from a datastore.
type BulkDeleteAppsDatastoreItemsRequestDataAttributes struct {
	// List of primary keys identifying items to delete from datastore. Up to 100 items can be deleted in a single request.
	ItemKeys []string `json:"item_keys,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBulkDeleteAppsDatastoreItemsRequestDataAttributes instantiates a new BulkDeleteAppsDatastoreItemsRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBulkDeleteAppsDatastoreItemsRequestDataAttributes() *BulkDeleteAppsDatastoreItemsRequestDataAttributes {
	this := BulkDeleteAppsDatastoreItemsRequestDataAttributes{}
	return &this
}

// NewBulkDeleteAppsDatastoreItemsRequestDataAttributesWithDefaults instantiates a new BulkDeleteAppsDatastoreItemsRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBulkDeleteAppsDatastoreItemsRequestDataAttributesWithDefaults() *BulkDeleteAppsDatastoreItemsRequestDataAttributes {
	this := BulkDeleteAppsDatastoreItemsRequestDataAttributes{}
	return &this
}

// GetItemKeys returns the ItemKeys field value if set, zero value otherwise.
func (o *BulkDeleteAppsDatastoreItemsRequestDataAttributes) GetItemKeys() []string {
	if o == nil || o.ItemKeys == nil {
		var ret []string
		return ret
	}
	return o.ItemKeys
}

// GetItemKeysOk returns a tuple with the ItemKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkDeleteAppsDatastoreItemsRequestDataAttributes) GetItemKeysOk() (*[]string, bool) {
	if o == nil || o.ItemKeys == nil {
		return nil, false
	}
	return &o.ItemKeys, true
}

// HasItemKeys returns a boolean if a field has been set.
func (o *BulkDeleteAppsDatastoreItemsRequestDataAttributes) HasItemKeys() bool {
	return o != nil && o.ItemKeys != nil
}

// SetItemKeys gets a reference to the given []string and assigns it to the ItemKeys field.
func (o *BulkDeleteAppsDatastoreItemsRequestDataAttributes) SetItemKeys(v []string) {
	o.ItemKeys = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BulkDeleteAppsDatastoreItemsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ItemKeys != nil {
		toSerialize["item_keys"] = o.ItemKeys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BulkDeleteAppsDatastoreItemsRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ItemKeys []string `json:"item_keys,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"item_keys"})
	} else {
		return err
	}
	o.ItemKeys = all.ItemKeys

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
