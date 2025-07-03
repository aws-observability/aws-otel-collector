// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutputSchema A list of output parameters for the workflow.
type OutputSchema struct {
	// The `OutputSchema` `parameters`.
	Parameters []OutputSchemaParameters `json:"parameters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOutputSchema instantiates a new OutputSchema object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOutputSchema() *OutputSchema {
	this := OutputSchema{}
	return &this
}

// NewOutputSchemaWithDefaults instantiates a new OutputSchema object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOutputSchemaWithDefaults() *OutputSchema {
	this := OutputSchema{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *OutputSchema) GetParameters() []OutputSchemaParameters {
	if o == nil || o.Parameters == nil {
		var ret []OutputSchemaParameters
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputSchema) GetParametersOk() (*[]OutputSchemaParameters, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *OutputSchema) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given []OutputSchemaParameters and assigns it to the Parameters field.
func (o *OutputSchema) SetParameters(v []OutputSchemaParameters) {
	o.Parameters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OutputSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OutputSchema) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Parameters []OutputSchemaParameters `json:"parameters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"parameters"})
	} else {
		return err
	}
	o.Parameters = all.Parameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
