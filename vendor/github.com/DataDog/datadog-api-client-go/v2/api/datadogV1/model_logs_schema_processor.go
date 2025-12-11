// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSchemaProcessor A processor that has additional validations and checks for a given schema. Currently supported schema types include OCSF.
type LogsSchemaProcessor struct {
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// The `LogsSchemaProcessor` `mappers`.
	Mappers []LogsSchemaMapper `json:"mappers"`
	// Name of the processor.
	Name string `json:"name"`
	// Configuration of the schema data to use.
	Schema LogsSchemaData `json:"schema"`
	// Type of logs schema processor.
	Type LogsSchemaProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsSchemaProcessor instantiates a new LogsSchemaProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsSchemaProcessor(mappers []LogsSchemaMapper, name string, schema LogsSchemaData, typeVar LogsSchemaProcessorType) *LogsSchemaProcessor {
	this := LogsSchemaProcessor{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	this.Mappers = mappers
	this.Name = name
	this.Schema = schema
	this.Type = typeVar
	return &this
}

// NewLogsSchemaProcessorWithDefaults instantiates a new LogsSchemaProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsSchemaProcessorWithDefaults() *LogsSchemaProcessor {
	this := LogsSchemaProcessor{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var typeVar LogsSchemaProcessorType = LOGSSCHEMAPROCESSORTYPE_SCHEMA_PROCESSOR
	this.Type = typeVar
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsSchemaProcessor) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsSchemaProcessor) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsSchemaProcessor) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsSchemaProcessor) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetMappers returns the Mappers field value.
func (o *LogsSchemaProcessor) GetMappers() []LogsSchemaMapper {
	if o == nil {
		var ret []LogsSchemaMapper
		return ret
	}
	return o.Mappers
}

// GetMappersOk returns a tuple with the Mappers field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaProcessor) GetMappersOk() (*[]LogsSchemaMapper, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mappers, true
}

// SetMappers sets field value.
func (o *LogsSchemaProcessor) SetMappers(v []LogsSchemaMapper) {
	o.Mappers = v
}

// GetName returns the Name field value.
func (o *LogsSchemaProcessor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaProcessor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LogsSchemaProcessor) SetName(v string) {
	o.Name = v
}

// GetSchema returns the Schema field value.
func (o *LogsSchemaProcessor) GetSchema() LogsSchemaData {
	if o == nil {
		var ret LogsSchemaData
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaProcessor) GetSchemaOk() (*LogsSchemaData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value.
func (o *LogsSchemaProcessor) SetSchema(v LogsSchemaData) {
	o.Schema = v
}

// GetType returns the Type field value.
func (o *LogsSchemaProcessor) GetType() LogsSchemaProcessorType {
	if o == nil {
		var ret LogsSchemaProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaProcessor) GetTypeOk() (*LogsSchemaProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsSchemaProcessor) SetType(v LogsSchemaProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsSchemaProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	toSerialize["mappers"] = o.Mappers
	toSerialize["name"] = o.Name
	toSerialize["schema"] = o.Schema
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsSchemaProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsEnabled *bool                    `json:"is_enabled,omitempty"`
		Mappers   *[]LogsSchemaMapper      `json:"mappers"`
		Name      *string                  `json:"name"`
		Schema    *LogsSchemaData          `json:"schema"`
		Type      *LogsSchemaProcessorType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mappers == nil {
		return fmt.Errorf("required field mappers missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Schema == nil {
		return fmt.Errorf("required field schema missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_enabled", "mappers", "name", "schema", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.IsEnabled = all.IsEnabled
	o.Mappers = *all.Mappers
	o.Name = *all.Name
	if all.Schema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schema = *all.Schema
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
