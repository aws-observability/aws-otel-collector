// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3APISpecInterfaceDefinition The definition of `EntityV3APISpecInterfaceDefinition` object.
type EntityV3APISpecInterfaceDefinition struct {
	// The API definition.
	Definition interface{} `json:"definition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3APISpecInterfaceDefinition instantiates a new EntityV3APISpecInterfaceDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3APISpecInterfaceDefinition() *EntityV3APISpecInterfaceDefinition {
	this := EntityV3APISpecInterfaceDefinition{}
	return &this
}

// NewEntityV3APISpecInterfaceDefinitionWithDefaults instantiates a new EntityV3APISpecInterfaceDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3APISpecInterfaceDefinitionWithDefaults() *EntityV3APISpecInterfaceDefinition {
	this := EntityV3APISpecInterfaceDefinition{}
	return &this
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *EntityV3APISpecInterfaceDefinition) GetDefinition() interface{} {
	if o == nil || o.Definition == nil {
		var ret interface{}
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3APISpecInterfaceDefinition) GetDefinitionOk() (*interface{}, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return &o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *EntityV3APISpecInterfaceDefinition) HasDefinition() bool {
	return o != nil && o.Definition != nil
}

// SetDefinition gets a reference to the given interface{} and assigns it to the Definition field.
func (o *EntityV3APISpecInterfaceDefinition) SetDefinition(v interface{}) {
	o.Definition = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3APISpecInterfaceDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3APISpecInterfaceDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Definition interface{} `json:"definition,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Definition = all.Definition

	return nil
}
