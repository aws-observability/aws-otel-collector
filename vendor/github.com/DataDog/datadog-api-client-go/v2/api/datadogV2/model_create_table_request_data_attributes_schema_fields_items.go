// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTableRequestDataAttributesSchemaFieldsItems A single field (column) in the reference table schema to be created.
type CreateTableRequestDataAttributesSchemaFieldsItems struct {
	// The field name.
	Name string `json:"name"`
	// The field type for reference table schema fields.
	Type ReferenceTableSchemaFieldType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateTableRequestDataAttributesSchemaFieldsItems instantiates a new CreateTableRequestDataAttributesSchemaFieldsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTableRequestDataAttributesSchemaFieldsItems(name string, typeVar ReferenceTableSchemaFieldType) *CreateTableRequestDataAttributesSchemaFieldsItems {
	this := CreateTableRequestDataAttributesSchemaFieldsItems{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewCreateTableRequestDataAttributesSchemaFieldsItemsWithDefaults instantiates a new CreateTableRequestDataAttributesSchemaFieldsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTableRequestDataAttributesSchemaFieldsItemsWithDefaults() *CreateTableRequestDataAttributesSchemaFieldsItems {
	this := CreateTableRequestDataAttributesSchemaFieldsItems{}
	return &this
}

// GetName returns the Name field value.
func (o *CreateTableRequestDataAttributesSchemaFieldsItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesSchemaFieldsItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateTableRequestDataAttributesSchemaFieldsItems) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *CreateTableRequestDataAttributesSchemaFieldsItems) GetType() ReferenceTableSchemaFieldType {
	if o == nil {
		var ret ReferenceTableSchemaFieldType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesSchemaFieldsItems) GetTypeOk() (*ReferenceTableSchemaFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateTableRequestDataAttributesSchemaFieldsItems) SetType(v ReferenceTableSchemaFieldType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTableRequestDataAttributesSchemaFieldsItems) MarshalJSON() ([]byte, error) {
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
func (o *CreateTableRequestDataAttributesSchemaFieldsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string                        `json:"name"`
		Type *ReferenceTableSchemaFieldType `json:"type"`
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

	hasInvalidField := false
	o.Name = *all.Name
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
