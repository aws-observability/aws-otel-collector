// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedRawSchemaAttributes Included raw schema attributes.
type EntityResponseIncludedRawSchemaAttributes struct {
	// Schema from user input in base64 encoding.
	RawSchema *string `json:"rawSchema,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityResponseIncludedRawSchemaAttributes instantiates a new EntityResponseIncludedRawSchemaAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityResponseIncludedRawSchemaAttributes() *EntityResponseIncludedRawSchemaAttributes {
	this := EntityResponseIncludedRawSchemaAttributes{}
	return &this
}

// NewEntityResponseIncludedRawSchemaAttributesWithDefaults instantiates a new EntityResponseIncludedRawSchemaAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityResponseIncludedRawSchemaAttributesWithDefaults() *EntityResponseIncludedRawSchemaAttributes {
	this := EntityResponseIncludedRawSchemaAttributes{}
	return &this
}

// GetRawSchema returns the RawSchema field value if set, zero value otherwise.
func (o *EntityResponseIncludedRawSchemaAttributes) GetRawSchema() string {
	if o == nil || o.RawSchema == nil {
		var ret string
		return ret
	}
	return *o.RawSchema
}

// GetRawSchemaOk returns a tuple with the RawSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRawSchemaAttributes) GetRawSchemaOk() (*string, bool) {
	if o == nil || o.RawSchema == nil {
		return nil, false
	}
	return o.RawSchema, true
}

// HasRawSchema returns a boolean if a field has been set.
func (o *EntityResponseIncludedRawSchemaAttributes) HasRawSchema() bool {
	return o != nil && o.RawSchema != nil
}

// SetRawSchema gets a reference to the given string and assigns it to the RawSchema field.
func (o *EntityResponseIncludedRawSchemaAttributes) SetRawSchema(v string) {
	o.RawSchema = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityResponseIncludedRawSchemaAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RawSchema != nil {
		toSerialize["rawSchema"] = o.RawSchema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityResponseIncludedRawSchemaAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RawSchema *string `json:"rawSchema,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rawSchema"})
	} else {
		return err
	}
	o.RawSchema = all.RawSchema

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
