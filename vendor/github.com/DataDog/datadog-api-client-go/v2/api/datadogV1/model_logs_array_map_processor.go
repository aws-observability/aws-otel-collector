// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayMapProcessor The array-map processor transforms each element of a source array by applying
// sub-processors in order and collecting the results into a target array.
// Results can be written to a new array, to the source array (in-place), or to
// an existing target array. Sub-processors can read from `$sourceElem.<field>`
// (object element field), bare `$sourceElem` (primitive element), or any parent
// log attribute path. Sub-processors write to `$targetElem.<field>` (object
// output field) or bare `$targetElem` (primitive output).
type LogsArrayMapProcessor struct {
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
	// When `false` and `source != target`, the source attribute is removed after
	// processing. Cannot be `false` when `source == target`.
	PreserveSource *bool `json:"preserve_source,omitempty"`
	// Sub-processors applied to each element. Allowed types: `attribute-remapper`,
	// `string-builder-processor`, `arithmetic-processor`, `category-processor`.
	Processors []LogsArrayMapSubProcessor `json:"processors"`
	// Attribute path of the source array. Elements are read-only via `$sourceElem`
	// inside sub-processors.
	Source string `json:"source"`
	// Attribute path of the output array. Sub-processors write to `$targetElem`
	// (or `$targetElem.<field>`) to build each output element.
	Target string `json:"target"`
	// Type of logs array-map processor.
	Type LogsArrayMapProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArrayMapProcessor instantiates a new LogsArrayMapProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArrayMapProcessor(processors []LogsArrayMapSubProcessor, source string, target string, typeVar LogsArrayMapProcessorType) *LogsArrayMapProcessor {
	this := LogsArrayMapProcessor{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var preserveSource bool = true
	this.PreserveSource = &preserveSource
	this.Processors = processors
	this.Source = source
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewLogsArrayMapProcessorWithDefaults instantiates a new LogsArrayMapProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArrayMapProcessorWithDefaults() *LogsArrayMapProcessor {
	this := LogsArrayMapProcessor{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var preserveSource bool = true
	this.PreserveSource = &preserveSource
	var typeVar LogsArrayMapProcessorType = LOGSARRAYMAPPROCESSORTYPE_ARRAY_MAP_PROCESSOR
	this.Type = typeVar
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsArrayMapProcessor) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayMapProcessor) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsArrayMapProcessor) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsArrayMapProcessor) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsArrayMapProcessor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayMapProcessor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsArrayMapProcessor) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsArrayMapProcessor) SetName(v string) {
	o.Name = &v
}

// GetPreserveSource returns the PreserveSource field value if set, zero value otherwise.
func (o *LogsArrayMapProcessor) GetPreserveSource() bool {
	if o == nil || o.PreserveSource == nil {
		var ret bool
		return ret
	}
	return *o.PreserveSource
}

// GetPreserveSourceOk returns a tuple with the PreserveSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayMapProcessor) GetPreserveSourceOk() (*bool, bool) {
	if o == nil || o.PreserveSource == nil {
		return nil, false
	}
	return o.PreserveSource, true
}

// HasPreserveSource returns a boolean if a field has been set.
func (o *LogsArrayMapProcessor) HasPreserveSource() bool {
	return o != nil && o.PreserveSource != nil
}

// SetPreserveSource gets a reference to the given bool and assigns it to the PreserveSource field.
func (o *LogsArrayMapProcessor) SetPreserveSource(v bool) {
	o.PreserveSource = &v
}

// GetProcessors returns the Processors field value.
func (o *LogsArrayMapProcessor) GetProcessors() []LogsArrayMapSubProcessor {
	if o == nil {
		var ret []LogsArrayMapSubProcessor
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapProcessor) GetProcessorsOk() (*[]LogsArrayMapSubProcessor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Processors, true
}

// SetProcessors sets field value.
func (o *LogsArrayMapProcessor) SetProcessors(v []LogsArrayMapSubProcessor) {
	o.Processors = v
}

// GetSource returns the Source field value.
func (o *LogsArrayMapProcessor) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapProcessor) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *LogsArrayMapProcessor) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value.
func (o *LogsArrayMapProcessor) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapProcessor) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *LogsArrayMapProcessor) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *LogsArrayMapProcessor) GetType() LogsArrayMapProcessorType {
	if o == nil {
		var ret LogsArrayMapProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapProcessor) GetTypeOk() (*LogsArrayMapProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArrayMapProcessor) SetType(v LogsArrayMapProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArrayMapProcessor) MarshalJSON() ([]byte, error) {
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
	if o.PreserveSource != nil {
		toSerialize["preserve_source"] = o.PreserveSource
	}
	toSerialize["processors"] = o.Processors
	toSerialize["source"] = o.Source
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArrayMapProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsEnabled      *bool                       `json:"is_enabled,omitempty"`
		Name           *string                     `json:"name,omitempty"`
		PreserveSource *bool                       `json:"preserve_source,omitempty"`
		Processors     *[]LogsArrayMapSubProcessor `json:"processors"`
		Source         *string                     `json:"source"`
		Target         *string                     `json:"target"`
		Type           *LogsArrayMapProcessorType  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Processors == nil {
		return fmt.Errorf("required field processors missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_enabled", "name", "preserve_source", "processors", "source", "target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
	o.PreserveSource = all.PreserveSource
	o.Processors = *all.Processors
	o.Source = *all.Source
	o.Target = *all.Target
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
