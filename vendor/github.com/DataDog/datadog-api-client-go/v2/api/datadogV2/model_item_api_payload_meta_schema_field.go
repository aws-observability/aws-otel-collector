// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ItemApiPayloadMetaSchemaField Information about a specific column in the datastore schema.
type ItemApiPayloadMetaSchemaField struct {
	// The name of this column in the datastore.
	Name string `json:"name"`
	// The data type of this column. For example, 'string', 'number', or 'boolean'.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewItemApiPayloadMetaSchemaField instantiates a new ItemApiPayloadMetaSchemaField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewItemApiPayloadMetaSchemaField(name string, typeVar string) *ItemApiPayloadMetaSchemaField {
	this := ItemApiPayloadMetaSchemaField{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewItemApiPayloadMetaSchemaFieldWithDefaults instantiates a new ItemApiPayloadMetaSchemaField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewItemApiPayloadMetaSchemaFieldWithDefaults() *ItemApiPayloadMetaSchemaField {
	this := ItemApiPayloadMetaSchemaField{}
	return &this
}

// GetName returns the Name field value.
func (o *ItemApiPayloadMetaSchemaField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadMetaSchemaField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ItemApiPayloadMetaSchemaField) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *ItemApiPayloadMetaSchemaField) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadMetaSchemaField) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ItemApiPayloadMetaSchemaField) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ItemApiPayloadMetaSchemaField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ItemApiPayloadMetaSchemaField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name"`
		Type *string `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "type"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
