// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionJSONSchemaTargetTarget Composed target for `validatesJSONSchema` operator.
type SyntheticsAssertionJSONSchemaTargetTarget struct {
	// The JSON Schema to assert.
	JsonSchema *string `json:"jsonSchema,omitempty"`
	// The JSON Schema meta-schema version used in the assertion.
	MetaSchema *SyntheticsAssertionJSONSchemaMetaSchema `json:"metaSchema,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsAssertionJSONSchemaTargetTarget instantiates a new SyntheticsAssertionJSONSchemaTargetTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsAssertionJSONSchemaTargetTarget() *SyntheticsAssertionJSONSchemaTargetTarget {
	this := SyntheticsAssertionJSONSchemaTargetTarget{}
	return &this
}

// NewSyntheticsAssertionJSONSchemaTargetTargetWithDefaults instantiates a new SyntheticsAssertionJSONSchemaTargetTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsAssertionJSONSchemaTargetTargetWithDefaults() *SyntheticsAssertionJSONSchemaTargetTarget {
	this := SyntheticsAssertionJSONSchemaTargetTarget{}
	return &this
}

// GetJsonSchema returns the JsonSchema field value if set, zero value otherwise.
func (o *SyntheticsAssertionJSONSchemaTargetTarget) GetJsonSchema() string {
	if o == nil || o.JsonSchema == nil {
		var ret string
		return ret
	}
	return *o.JsonSchema
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionJSONSchemaTargetTarget) GetJsonSchemaOk() (*string, bool) {
	if o == nil || o.JsonSchema == nil {
		return nil, false
	}
	return o.JsonSchema, true
}

// HasJsonSchema returns a boolean if a field has been set.
func (o *SyntheticsAssertionJSONSchemaTargetTarget) HasJsonSchema() bool {
	return o != nil && o.JsonSchema != nil
}

// SetJsonSchema gets a reference to the given string and assigns it to the JsonSchema field.
func (o *SyntheticsAssertionJSONSchemaTargetTarget) SetJsonSchema(v string) {
	o.JsonSchema = &v
}

// GetMetaSchema returns the MetaSchema field value if set, zero value otherwise.
func (o *SyntheticsAssertionJSONSchemaTargetTarget) GetMetaSchema() SyntheticsAssertionJSONSchemaMetaSchema {
	if o == nil || o.MetaSchema == nil {
		var ret SyntheticsAssertionJSONSchemaMetaSchema
		return ret
	}
	return *o.MetaSchema
}

// GetMetaSchemaOk returns a tuple with the MetaSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionJSONSchemaTargetTarget) GetMetaSchemaOk() (*SyntheticsAssertionJSONSchemaMetaSchema, bool) {
	if o == nil || o.MetaSchema == nil {
		return nil, false
	}
	return o.MetaSchema, true
}

// HasMetaSchema returns a boolean if a field has been set.
func (o *SyntheticsAssertionJSONSchemaTargetTarget) HasMetaSchema() bool {
	return o != nil && o.MetaSchema != nil
}

// SetMetaSchema gets a reference to the given SyntheticsAssertionJSONSchemaMetaSchema and assigns it to the MetaSchema field.
func (o *SyntheticsAssertionJSONSchemaTargetTarget) SetMetaSchema(v SyntheticsAssertionJSONSchemaMetaSchema) {
	o.MetaSchema = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsAssertionJSONSchemaTargetTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.JsonSchema != nil {
		toSerialize["jsonSchema"] = o.JsonSchema
	}
	if o.MetaSchema != nil {
		toSerialize["metaSchema"] = o.MetaSchema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsAssertionJSONSchemaTargetTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		JsonSchema *string                                  `json:"jsonSchema,omitempty"`
		MetaSchema *SyntheticsAssertionJSONSchemaMetaSchema `json:"metaSchema,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"jsonSchema", "metaSchema"})
	} else {
		return err
	}

	hasInvalidField := false
	o.JsonSchema = all.JsonSchema
	if all.MetaSchema != nil && !all.MetaSchema.IsValid() {
		hasInvalidField = true
	} else {
		o.MetaSchema = all.MetaSchema
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
