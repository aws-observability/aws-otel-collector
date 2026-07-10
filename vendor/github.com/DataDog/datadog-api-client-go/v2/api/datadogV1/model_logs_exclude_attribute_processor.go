// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsExcludeAttributeProcessor Use this processor to remove an attribute from a log during processing.
// The processor strips the specified attribute from the log event, which is useful
// when the attribute contains sensitive data or is no longer needed downstream.
type LogsExcludeAttributeProcessor struct {
	// Name of the log attribute to remove from the log event.
	AttributeToExclude string `json:"attribute_to_exclude"`
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
	// Type of logs exclude attribute processor.
	Type LogsExcludeAttributeProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsExcludeAttributeProcessor instantiates a new LogsExcludeAttributeProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsExcludeAttributeProcessor(attributeToExclude string, typeVar LogsExcludeAttributeProcessorType) *LogsExcludeAttributeProcessor {
	this := LogsExcludeAttributeProcessor{}
	this.AttributeToExclude = attributeToExclude
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	this.Type = typeVar
	return &this
}

// NewLogsExcludeAttributeProcessorWithDefaults instantiates a new LogsExcludeAttributeProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsExcludeAttributeProcessorWithDefaults() *LogsExcludeAttributeProcessor {
	this := LogsExcludeAttributeProcessor{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var typeVar LogsExcludeAttributeProcessorType = LOGSEXCLUDEATTRIBUTEPROCESSORTYPE_EXCLUDE_ATTRIBUTE
	this.Type = typeVar
	return &this
}

// GetAttributeToExclude returns the AttributeToExclude field value.
func (o *LogsExcludeAttributeProcessor) GetAttributeToExclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AttributeToExclude
}

// GetAttributeToExcludeOk returns a tuple with the AttributeToExclude field value
// and a boolean to check if the value has been set.
func (o *LogsExcludeAttributeProcessor) GetAttributeToExcludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeToExclude, true
}

// SetAttributeToExclude sets field value.
func (o *LogsExcludeAttributeProcessor) SetAttributeToExclude(v string) {
	o.AttributeToExclude = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsExcludeAttributeProcessor) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsExcludeAttributeProcessor) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsExcludeAttributeProcessor) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsExcludeAttributeProcessor) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsExcludeAttributeProcessor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsExcludeAttributeProcessor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsExcludeAttributeProcessor) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsExcludeAttributeProcessor) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value.
func (o *LogsExcludeAttributeProcessor) GetType() LogsExcludeAttributeProcessorType {
	if o == nil {
		var ret LogsExcludeAttributeProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsExcludeAttributeProcessor) GetTypeOk() (*LogsExcludeAttributeProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsExcludeAttributeProcessor) SetType(v LogsExcludeAttributeProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsExcludeAttributeProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attribute_to_exclude"] = o.AttributeToExclude
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsExcludeAttributeProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AttributeToExclude *string                            `json:"attribute_to_exclude"`
		IsEnabled          *bool                              `json:"is_enabled,omitempty"`
		Name               *string                            `json:"name,omitempty"`
		Type               *LogsExcludeAttributeProcessorType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AttributeToExclude == nil {
		return fmt.Errorf("required field attribute_to_exclude missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute_to_exclude", "is_enabled", "name", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AttributeToExclude = *all.AttributeToExclude
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
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
