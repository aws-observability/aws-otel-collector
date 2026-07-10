// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSpanToolDefinition A tool definition available to an LLM span.
type LLMObsSpanToolDefinition struct {
	// Description of what the tool does.
	Description *string `json:"description,omitempty"`
	// Name of the tool.
	Name *string `json:"name,omitempty"`
	// JSON schema describing the tool's input parameters.
	Schema map[string]interface{} `json:"schema,omitempty"`
	// Version of the tool definition.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsSpanToolDefinition instantiates a new LLMObsSpanToolDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsSpanToolDefinition() *LLMObsSpanToolDefinition {
	this := LLMObsSpanToolDefinition{}
	return &this
}

// NewLLMObsSpanToolDefinitionWithDefaults instantiates a new LLMObsSpanToolDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsSpanToolDefinitionWithDefaults() *LLMObsSpanToolDefinition {
	this := LLMObsSpanToolDefinition{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsSpanToolDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanToolDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsSpanToolDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsSpanToolDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LLMObsSpanToolDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanToolDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LLMObsSpanToolDefinition) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LLMObsSpanToolDefinition) SetName(v string) {
	o.Name = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *LLMObsSpanToolDefinition) GetSchema() map[string]interface{} {
	if o == nil || o.Schema == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanToolDefinition) GetSchemaOk() (*map[string]interface{}, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return &o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *LLMObsSpanToolDefinition) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given map[string]interface{} and assigns it to the Schema field.
func (o *LLMObsSpanToolDefinition) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LLMObsSpanToolDefinition) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanToolDefinition) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LLMObsSpanToolDefinition) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LLMObsSpanToolDefinition) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsSpanToolDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsSpanToolDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                `json:"description,omitempty"`
		Name        *string                `json:"name,omitempty"`
		Schema      map[string]interface{} `json:"schema,omitempty"`
		Version     *string                `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "schema", "version"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Name = all.Name
	o.Schema = all.Schema
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
