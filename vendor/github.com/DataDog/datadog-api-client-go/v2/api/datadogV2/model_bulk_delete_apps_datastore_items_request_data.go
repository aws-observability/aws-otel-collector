// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BulkDeleteAppsDatastoreItemsRequestData Data wrapper containing the data needed to delete items from a datastore.
type BulkDeleteAppsDatastoreItemsRequestData struct {
	// Attributes of request data to delete items from a datastore.
	Attributes *BulkDeleteAppsDatastoreItemsRequestDataAttributes `json:"attributes,omitempty"`
	// ID for the datastore of the items to delete.
	Id *string `json:"id,omitempty"`
	// Items resource type.
	Type BulkDeleteAppsDatastoreItemsRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBulkDeleteAppsDatastoreItemsRequestData instantiates a new BulkDeleteAppsDatastoreItemsRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBulkDeleteAppsDatastoreItemsRequestData(typeVar BulkDeleteAppsDatastoreItemsRequestDataType) *BulkDeleteAppsDatastoreItemsRequestData {
	this := BulkDeleteAppsDatastoreItemsRequestData{}
	this.Type = typeVar
	return &this
}

// NewBulkDeleteAppsDatastoreItemsRequestDataWithDefaults instantiates a new BulkDeleteAppsDatastoreItemsRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBulkDeleteAppsDatastoreItemsRequestDataWithDefaults() *BulkDeleteAppsDatastoreItemsRequestData {
	this := BulkDeleteAppsDatastoreItemsRequestData{}
	var typeVar BulkDeleteAppsDatastoreItemsRequestDataType = BULKDELETEAPPSDATASTOREITEMSREQUESTDATATYPE_ITEMS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BulkDeleteAppsDatastoreItemsRequestData) GetAttributes() BulkDeleteAppsDatastoreItemsRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret BulkDeleteAppsDatastoreItemsRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkDeleteAppsDatastoreItemsRequestData) GetAttributesOk() (*BulkDeleteAppsDatastoreItemsRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BulkDeleteAppsDatastoreItemsRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given BulkDeleteAppsDatastoreItemsRequestDataAttributes and assigns it to the Attributes field.
func (o *BulkDeleteAppsDatastoreItemsRequestData) SetAttributes(v BulkDeleteAppsDatastoreItemsRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BulkDeleteAppsDatastoreItemsRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkDeleteAppsDatastoreItemsRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BulkDeleteAppsDatastoreItemsRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BulkDeleteAppsDatastoreItemsRequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *BulkDeleteAppsDatastoreItemsRequestData) GetType() BulkDeleteAppsDatastoreItemsRequestDataType {
	if o == nil {
		var ret BulkDeleteAppsDatastoreItemsRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BulkDeleteAppsDatastoreItemsRequestData) GetTypeOk() (*BulkDeleteAppsDatastoreItemsRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *BulkDeleteAppsDatastoreItemsRequestData) SetType(v BulkDeleteAppsDatastoreItemsRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BulkDeleteAppsDatastoreItemsRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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
func (o *BulkDeleteAppsDatastoreItemsRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *BulkDeleteAppsDatastoreItemsRequestDataAttributes `json:"attributes,omitempty"`
		Id         *string                                            `json:"id,omitempty"`
		Type       *BulkDeleteAppsDatastoreItemsRequestDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
