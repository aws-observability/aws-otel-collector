// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BatchUpsertRowsRequestData Row resource containing a single row identifier and its column values.
type BatchUpsertRowsRequestData struct {
	// Attributes containing row data values for row creation or update operations.
	Attributes *BatchUpsertRowsRequestDataAttributes `json:"attributes,omitempty"`
	//
	Id string `json:"id"`
	// Row resource type.
	Type TableRowResourceDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBatchUpsertRowsRequestData instantiates a new BatchUpsertRowsRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBatchUpsertRowsRequestData(id string, typeVar TableRowResourceDataType) *BatchUpsertRowsRequestData {
	this := BatchUpsertRowsRequestData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewBatchUpsertRowsRequestDataWithDefaults instantiates a new BatchUpsertRowsRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBatchUpsertRowsRequestDataWithDefaults() *BatchUpsertRowsRequestData {
	this := BatchUpsertRowsRequestData{}
	var typeVar TableRowResourceDataType = TABLEROWRESOURCEDATATYPE_ROW
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BatchUpsertRowsRequestData) GetAttributes() BatchUpsertRowsRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret BatchUpsertRowsRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpsertRowsRequestData) GetAttributesOk() (*BatchUpsertRowsRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BatchUpsertRowsRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given BatchUpsertRowsRequestDataAttributes and assigns it to the Attributes field.
func (o *BatchUpsertRowsRequestData) SetAttributes(v BatchUpsertRowsRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *BatchUpsertRowsRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BatchUpsertRowsRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *BatchUpsertRowsRequestData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *BatchUpsertRowsRequestData) GetType() TableRowResourceDataType {
	if o == nil {
		var ret TableRowResourceDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BatchUpsertRowsRequestData) GetTypeOk() (*TableRowResourceDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *BatchUpsertRowsRequestData) SetType(v TableRowResourceDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BatchUpsertRowsRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BatchUpsertRowsRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *BatchUpsertRowsRequestDataAttributes `json:"attributes,omitempty"`
		Id         *string                               `json:"id"`
		Type       *TableRowResourceDataType             `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
