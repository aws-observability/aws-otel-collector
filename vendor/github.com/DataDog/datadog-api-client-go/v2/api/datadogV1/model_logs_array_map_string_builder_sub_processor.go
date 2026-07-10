// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayMapStringBuilderSubProcessor A string builder sub-processor for use inside an array-map processor.
// Unlike the top-level string builder processor, `is_enabled` is not supported.
type LogsArrayMapStringBuilderSubProcessor struct {
	// Replace missing attribute values with an empty string.
	IsReplaceMissing *bool `json:"is_replace_missing,omitempty"`
	// Name of the sub-processor.
	Name *string `json:"name,omitempty"`
	// Target attribute path for the result.
	Target string `json:"target"`
	// Formula with one or more attributes and raw text.
	Template string `json:"template"`
	// Type of logs string builder processor.
	Type LogsStringBuilderProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArrayMapStringBuilderSubProcessor instantiates a new LogsArrayMapStringBuilderSubProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArrayMapStringBuilderSubProcessor(target string, template string, typeVar LogsStringBuilderProcessorType) *LogsArrayMapStringBuilderSubProcessor {
	this := LogsArrayMapStringBuilderSubProcessor{}
	var isReplaceMissing bool = false
	this.IsReplaceMissing = &isReplaceMissing
	this.Target = target
	this.Template = template
	this.Type = typeVar
	return &this
}

// NewLogsArrayMapStringBuilderSubProcessorWithDefaults instantiates a new LogsArrayMapStringBuilderSubProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArrayMapStringBuilderSubProcessorWithDefaults() *LogsArrayMapStringBuilderSubProcessor {
	this := LogsArrayMapStringBuilderSubProcessor{}
	var isReplaceMissing bool = false
	this.IsReplaceMissing = &isReplaceMissing
	var typeVar LogsStringBuilderProcessorType = LOGSSTRINGBUILDERPROCESSORTYPE_STRING_BUILDER_PROCESSOR
	this.Type = typeVar
	return &this
}

// GetIsReplaceMissing returns the IsReplaceMissing field value if set, zero value otherwise.
func (o *LogsArrayMapStringBuilderSubProcessor) GetIsReplaceMissing() bool {
	if o == nil || o.IsReplaceMissing == nil {
		var ret bool
		return ret
	}
	return *o.IsReplaceMissing
}

// GetIsReplaceMissingOk returns a tuple with the IsReplaceMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayMapStringBuilderSubProcessor) GetIsReplaceMissingOk() (*bool, bool) {
	if o == nil || o.IsReplaceMissing == nil {
		return nil, false
	}
	return o.IsReplaceMissing, true
}

// HasIsReplaceMissing returns a boolean if a field has been set.
func (o *LogsArrayMapStringBuilderSubProcessor) HasIsReplaceMissing() bool {
	return o != nil && o.IsReplaceMissing != nil
}

// SetIsReplaceMissing gets a reference to the given bool and assigns it to the IsReplaceMissing field.
func (o *LogsArrayMapStringBuilderSubProcessor) SetIsReplaceMissing(v bool) {
	o.IsReplaceMissing = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsArrayMapStringBuilderSubProcessor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayMapStringBuilderSubProcessor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsArrayMapStringBuilderSubProcessor) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsArrayMapStringBuilderSubProcessor) SetName(v string) {
	o.Name = &v
}

// GetTarget returns the Target field value.
func (o *LogsArrayMapStringBuilderSubProcessor) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapStringBuilderSubProcessor) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *LogsArrayMapStringBuilderSubProcessor) SetTarget(v string) {
	o.Target = v
}

// GetTemplate returns the Template field value.
func (o *LogsArrayMapStringBuilderSubProcessor) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapStringBuilderSubProcessor) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value.
func (o *LogsArrayMapStringBuilderSubProcessor) SetTemplate(v string) {
	o.Template = v
}

// GetType returns the Type field value.
func (o *LogsArrayMapStringBuilderSubProcessor) GetType() LogsStringBuilderProcessorType {
	if o == nil {
		var ret LogsStringBuilderProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapStringBuilderSubProcessor) GetTypeOk() (*LogsStringBuilderProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArrayMapStringBuilderSubProcessor) SetType(v LogsStringBuilderProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArrayMapStringBuilderSubProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsReplaceMissing != nil {
		toSerialize["is_replace_missing"] = o.IsReplaceMissing
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["target"] = o.Target
	toSerialize["template"] = o.Template
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArrayMapStringBuilderSubProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsReplaceMissing *bool                           `json:"is_replace_missing,omitempty"`
		Name             *string                         `json:"name,omitempty"`
		Target           *string                         `json:"target"`
		Template         *string                         `json:"template"`
		Type             *LogsStringBuilderProcessorType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Template == nil {
		return fmt.Errorf("required field template missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_replace_missing", "name", "target", "template", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.IsReplaceMissing = all.IsReplaceMissing
	o.Name = all.Name
	o.Target = *all.Target
	o.Template = *all.Template
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
