// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableRowResourceIdentifier Row resource containing a single row identifier.
type TableRowResourceIdentifier struct {
	// The primary key value that uniquely identifies the row to delete.
	Id string `json:"id"`
	// Row resource type.
	Type TableRowResourceDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTableRowResourceIdentifier instantiates a new TableRowResourceIdentifier object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableRowResourceIdentifier(id string, typeVar TableRowResourceDataType) *TableRowResourceIdentifier {
	this := TableRowResourceIdentifier{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewTableRowResourceIdentifierWithDefaults instantiates a new TableRowResourceIdentifier object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableRowResourceIdentifierWithDefaults() *TableRowResourceIdentifier {
	this := TableRowResourceIdentifier{}
	var typeVar TableRowResourceDataType = TABLEROWRESOURCEDATATYPE_ROW
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *TableRowResourceIdentifier) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TableRowResourceIdentifier) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TableRowResourceIdentifier) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *TableRowResourceIdentifier) GetType() TableRowResourceDataType {
	if o == nil {
		var ret TableRowResourceDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TableRowResourceIdentifier) GetTypeOk() (*TableRowResourceDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TableRowResourceIdentifier) SetType(v TableRowResourceDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableRowResourceIdentifier) MarshalJSON() ([]byte, error) {
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
func (o *TableRowResourceIdentifier) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                   `json:"id"`
		Type *TableRowResourceDataType `json:"type"`
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
