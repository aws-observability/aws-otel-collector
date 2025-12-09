// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppsDatastoreItemRequestDataAttributes Attributes for updating a datastore item, including the item key and changes to apply.
type UpdateAppsDatastoreItemRequestDataAttributes struct {
	// The unique identifier of the item being updated.
	Id *string `json:"id,omitempty"`
	// Changes to apply to a datastore item using set operations.
	ItemChanges UpdateAppsDatastoreItemRequestDataAttributesItemChanges `json:"item_changes"`
	// The primary key that identifies the item to update. Cannot exceed 256 characters.
	ItemKey string `json:"item_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateAppsDatastoreItemRequestDataAttributes instantiates a new UpdateAppsDatastoreItemRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateAppsDatastoreItemRequestDataAttributes(itemChanges UpdateAppsDatastoreItemRequestDataAttributesItemChanges, itemKey string) *UpdateAppsDatastoreItemRequestDataAttributes {
	this := UpdateAppsDatastoreItemRequestDataAttributes{}
	this.ItemChanges = itemChanges
	this.ItemKey = itemKey
	return &this
}

// NewUpdateAppsDatastoreItemRequestDataAttributesWithDefaults instantiates a new UpdateAppsDatastoreItemRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateAppsDatastoreItemRequestDataAttributesWithDefaults() *UpdateAppsDatastoreItemRequestDataAttributes {
	this := UpdateAppsDatastoreItemRequestDataAttributes{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateAppsDatastoreItemRequestDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAppsDatastoreItemRequestDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateAppsDatastoreItemRequestDataAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateAppsDatastoreItemRequestDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetItemChanges returns the ItemChanges field value.
func (o *UpdateAppsDatastoreItemRequestDataAttributes) GetItemChanges() UpdateAppsDatastoreItemRequestDataAttributesItemChanges {
	if o == nil {
		var ret UpdateAppsDatastoreItemRequestDataAttributesItemChanges
		return ret
	}
	return o.ItemChanges
}

// GetItemChangesOk returns a tuple with the ItemChanges field value
// and a boolean to check if the value has been set.
func (o *UpdateAppsDatastoreItemRequestDataAttributes) GetItemChangesOk() (*UpdateAppsDatastoreItemRequestDataAttributesItemChanges, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemChanges, true
}

// SetItemChanges sets field value.
func (o *UpdateAppsDatastoreItemRequestDataAttributes) SetItemChanges(v UpdateAppsDatastoreItemRequestDataAttributesItemChanges) {
	o.ItemChanges = v
}

// GetItemKey returns the ItemKey field value.
func (o *UpdateAppsDatastoreItemRequestDataAttributes) GetItemKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ItemKey
}

// GetItemKeyOk returns a tuple with the ItemKey field value
// and a boolean to check if the value has been set.
func (o *UpdateAppsDatastoreItemRequestDataAttributes) GetItemKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemKey, true
}

// SetItemKey sets field value.
func (o *UpdateAppsDatastoreItemRequestDataAttributes) SetItemKey(v string) {
	o.ItemKey = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateAppsDatastoreItemRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["item_changes"] = o.ItemChanges
	toSerialize["item_key"] = o.ItemKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateAppsDatastoreItemRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string                                                  `json:"id,omitempty"`
		ItemChanges *UpdateAppsDatastoreItemRequestDataAttributesItemChanges `json:"item_changes"`
		ItemKey     *string                                                  `json:"item_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ItemChanges == nil {
		return fmt.Errorf("required field item_changes missing")
	}
	if all.ItemKey == nil {
		return fmt.Errorf("required field item_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "item_changes", "item_key"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	if all.ItemChanges.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ItemChanges = *all.ItemChanges
	o.ItemKey = *all.ItemKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
