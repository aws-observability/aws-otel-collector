// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedSchemaAttributes Included schema.
type EntityResponseIncludedSchemaAttributes struct {
	// Entity schema v3.
	Schema *EntityV3 `json:"schema,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityResponseIncludedSchemaAttributes instantiates a new EntityResponseIncludedSchemaAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityResponseIncludedSchemaAttributes() *EntityResponseIncludedSchemaAttributes {
	this := EntityResponseIncludedSchemaAttributes{}
	return &this
}

// NewEntityResponseIncludedSchemaAttributesWithDefaults instantiates a new EntityResponseIncludedSchemaAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityResponseIncludedSchemaAttributesWithDefaults() *EntityResponseIncludedSchemaAttributes {
	this := EntityResponseIncludedSchemaAttributes{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *EntityResponseIncludedSchemaAttributes) GetSchema() EntityV3 {
	if o == nil || o.Schema == nil {
		var ret EntityV3
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedSchemaAttributes) GetSchemaOk() (*EntityV3, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *EntityResponseIncludedSchemaAttributes) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given EntityV3 and assigns it to the Schema field.
func (o *EntityResponseIncludedSchemaAttributes) SetSchema(v EntityV3) {
	o.Schema = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityResponseIncludedSchemaAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityResponseIncludedSchemaAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Schema *EntityV3 `json:"schema,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"schema"})
	} else {
		return err
	}
	o.Schema = all.Schema

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
