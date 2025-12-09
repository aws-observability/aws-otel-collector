// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BulkPutAppsDatastoreItemsRequestDataAttributes Configuration for bulk inserting multiple items into a datastore.
type BulkPutAppsDatastoreItemsRequestDataAttributes struct {
	// How to handle conflicts when inserting items that already exist in the datastore.
	ConflictMode *DatastoreItemConflictMode `json:"conflict_mode,omitempty"`
	// An array of items to add to the datastore, where each item is a set of key-value pairs representing the item's data. Up to 100 items can be updated in a single request.
	Values []map[string]interface{} `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBulkPutAppsDatastoreItemsRequestDataAttributes instantiates a new BulkPutAppsDatastoreItemsRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBulkPutAppsDatastoreItemsRequestDataAttributes(values []map[string]interface{}) *BulkPutAppsDatastoreItemsRequestDataAttributes {
	this := BulkPutAppsDatastoreItemsRequestDataAttributes{}
	this.Values = values
	return &this
}

// NewBulkPutAppsDatastoreItemsRequestDataAttributesWithDefaults instantiates a new BulkPutAppsDatastoreItemsRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBulkPutAppsDatastoreItemsRequestDataAttributesWithDefaults() *BulkPutAppsDatastoreItemsRequestDataAttributes {
	this := BulkPutAppsDatastoreItemsRequestDataAttributes{}
	return &this
}

// GetConflictMode returns the ConflictMode field value if set, zero value otherwise.
func (o *BulkPutAppsDatastoreItemsRequestDataAttributes) GetConflictMode() DatastoreItemConflictMode {
	if o == nil || o.ConflictMode == nil {
		var ret DatastoreItemConflictMode
		return ret
	}
	return *o.ConflictMode
}

// GetConflictModeOk returns a tuple with the ConflictMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkPutAppsDatastoreItemsRequestDataAttributes) GetConflictModeOk() (*DatastoreItemConflictMode, bool) {
	if o == nil || o.ConflictMode == nil {
		return nil, false
	}
	return o.ConflictMode, true
}

// HasConflictMode returns a boolean if a field has been set.
func (o *BulkPutAppsDatastoreItemsRequestDataAttributes) HasConflictMode() bool {
	return o != nil && o.ConflictMode != nil
}

// SetConflictMode gets a reference to the given DatastoreItemConflictMode and assigns it to the ConflictMode field.
func (o *BulkPutAppsDatastoreItemsRequestDataAttributes) SetConflictMode(v DatastoreItemConflictMode) {
	o.ConflictMode = &v
}

// GetValues returns the Values field value.
func (o *BulkPutAppsDatastoreItemsRequestDataAttributes) GetValues() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *BulkPutAppsDatastoreItemsRequestDataAttributes) GetValuesOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *BulkPutAppsDatastoreItemsRequestDataAttributes) SetValues(v []map[string]interface{}) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BulkPutAppsDatastoreItemsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConflictMode != nil {
		toSerialize["conflict_mode"] = o.ConflictMode
	}
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BulkPutAppsDatastoreItemsRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConflictMode *DatastoreItemConflictMode `json:"conflict_mode,omitempty"`
		Values       *[]map[string]interface{}  `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conflict_mode", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ConflictMode != nil && !all.ConflictMode.IsValid() {
		hasInvalidField = true
	} else {
		o.ConflictMode = all.ConflictMode
	}
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
