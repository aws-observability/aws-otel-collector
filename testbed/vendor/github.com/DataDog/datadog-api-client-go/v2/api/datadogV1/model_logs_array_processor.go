// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayProcessor A processor for extracting, aggregating, or transforming values from JSON arrays within your logs.
// Supported operations are:
// - Select value from matching element
// - Compute array length
// - Append a value to an array
type LogsArrayProcessor struct {
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
	// Configuration of the array processor operation to perform.
	Operation LogsArrayProcessorOperation `json:"operation"`
	// Type of logs array processor.
	Type LogsArrayProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArrayProcessor instantiates a new LogsArrayProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArrayProcessor(operation LogsArrayProcessorOperation, typeVar LogsArrayProcessorType) *LogsArrayProcessor {
	this := LogsArrayProcessor{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	this.Operation = operation
	this.Type = typeVar
	return &this
}

// NewLogsArrayProcessorWithDefaults instantiates a new LogsArrayProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArrayProcessorWithDefaults() *LogsArrayProcessor {
	this := LogsArrayProcessor{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var typeVar LogsArrayProcessorType = LOGSARRAYPROCESSORTYPE_ARRAY_PROCESSOR
	this.Type = typeVar
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsArrayProcessor) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessor) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsArrayProcessor) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsArrayProcessor) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsArrayProcessor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsArrayProcessor) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsArrayProcessor) SetName(v string) {
	o.Name = &v
}

// GetOperation returns the Operation field value.
func (o *LogsArrayProcessor) GetOperation() LogsArrayProcessorOperation {
	if o == nil {
		var ret LogsArrayProcessorOperation
		return ret
	}
	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessor) GetOperationOk() (*LogsArrayProcessorOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value.
func (o *LogsArrayProcessor) SetOperation(v LogsArrayProcessorOperation) {
	o.Operation = v
}

// GetType returns the Type field value.
func (o *LogsArrayProcessor) GetType() LogsArrayProcessorType {
	if o == nil {
		var ret LogsArrayProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArrayProcessor) GetTypeOk() (*LogsArrayProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArrayProcessor) SetType(v LogsArrayProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArrayProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["operation"] = o.Operation
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArrayProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsEnabled *bool                        `json:"is_enabled,omitempty"`
		Name      *string                      `json:"name,omitempty"`
		Operation *LogsArrayProcessorOperation `json:"operation"`
		Type      *LogsArrayProcessorType      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Operation == nil {
		return fmt.Errorf("required field operation missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_enabled", "name", "operation", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
	o.Operation = *all.Operation
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
