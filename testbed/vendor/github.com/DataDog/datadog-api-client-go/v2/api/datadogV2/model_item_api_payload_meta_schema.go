// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ItemApiPayloadMetaSchema Schema information about the datastore, including its primary key and field definitions.
type ItemApiPayloadMetaSchema struct {
	// An array describing the columns available in this datastore.
	Fields []ItemApiPayloadMetaSchemaField `json:"fields,omitempty"`
	// The name of the primary key column for this datastore.
	PrimaryKey *string `json:"primary_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewItemApiPayloadMetaSchema instantiates a new ItemApiPayloadMetaSchema object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewItemApiPayloadMetaSchema() *ItemApiPayloadMetaSchema {
	this := ItemApiPayloadMetaSchema{}
	return &this
}

// NewItemApiPayloadMetaSchemaWithDefaults instantiates a new ItemApiPayloadMetaSchema object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewItemApiPayloadMetaSchemaWithDefaults() *ItemApiPayloadMetaSchema {
	this := ItemApiPayloadMetaSchema{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ItemApiPayloadMetaSchema) GetFields() []ItemApiPayloadMetaSchemaField {
	if o == nil || o.Fields == nil {
		var ret []ItemApiPayloadMetaSchemaField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadMetaSchema) GetFieldsOk() (*[]ItemApiPayloadMetaSchemaField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ItemApiPayloadMetaSchema) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given []ItemApiPayloadMetaSchemaField and assigns it to the Fields field.
func (o *ItemApiPayloadMetaSchema) SetFields(v []ItemApiPayloadMetaSchemaField) {
	o.Fields = v
}

// GetPrimaryKey returns the PrimaryKey field value if set, zero value otherwise.
func (o *ItemApiPayloadMetaSchema) GetPrimaryKey() string {
	if o == nil || o.PrimaryKey == nil {
		var ret string
		return ret
	}
	return *o.PrimaryKey
}

// GetPrimaryKeyOk returns a tuple with the PrimaryKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadMetaSchema) GetPrimaryKeyOk() (*string, bool) {
	if o == nil || o.PrimaryKey == nil {
		return nil, false
	}
	return o.PrimaryKey, true
}

// HasPrimaryKey returns a boolean if a field has been set.
func (o *ItemApiPayloadMetaSchema) HasPrimaryKey() bool {
	return o != nil && o.PrimaryKey != nil
}

// SetPrimaryKey gets a reference to the given string and assigns it to the PrimaryKey field.
func (o *ItemApiPayloadMetaSchema) SetPrimaryKey(v string) {
	o.PrimaryKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ItemApiPayloadMetaSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.PrimaryKey != nil {
		toSerialize["primary_key"] = o.PrimaryKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ItemApiPayloadMetaSchema) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields     []ItemApiPayloadMetaSchemaField `json:"fields,omitempty"`
		PrimaryKey *string                         `json:"primary_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "primary_key"})
	} else {
		return err
	}
	o.Fields = all.Fields
	o.PrimaryKey = all.PrimaryKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
