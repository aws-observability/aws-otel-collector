// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsDecoderProcessor The decoder processor decodes any source attribute containing a
// base64/base16-encoded UTF-8/ASCII string back to its original value, storing the
// result in a target attribute.
type LogsDecoderProcessor struct {
	// The encoding used to represent the binary data.
	BinaryToTextEncoding LogsDecoderProcessorBinaryToTextEncoding `json:"binary_to_text_encoding"`
	// The original representation of input string.
	InputRepresentation LogsDecoderProcessorInputRepresentation `json:"input_representation"`
	// Whether the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
	// Name of the log attribute with the encoded data.
	Source string `json:"source"`
	// Name of the log attribute that contains the decoded data.
	Target string `json:"target"`
	// Type of logs decoder processor.
	Type LogsDecoderProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsDecoderProcessor instantiates a new LogsDecoderProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsDecoderProcessor(binaryToTextEncoding LogsDecoderProcessorBinaryToTextEncoding, inputRepresentation LogsDecoderProcessorInputRepresentation, source string, target string, typeVar LogsDecoderProcessorType) *LogsDecoderProcessor {
	this := LogsDecoderProcessor{}
	this.BinaryToTextEncoding = binaryToTextEncoding
	this.InputRepresentation = inputRepresentation
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	this.Source = source
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewLogsDecoderProcessorWithDefaults instantiates a new LogsDecoderProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsDecoderProcessorWithDefaults() *LogsDecoderProcessor {
	this := LogsDecoderProcessor{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var typeVar LogsDecoderProcessorType = LOGSDECODERPROCESSORTYPE_DECODER_PROCESSOR
	this.Type = typeVar
	return &this
}

// GetBinaryToTextEncoding returns the BinaryToTextEncoding field value.
func (o *LogsDecoderProcessor) GetBinaryToTextEncoding() LogsDecoderProcessorBinaryToTextEncoding {
	if o == nil {
		var ret LogsDecoderProcessorBinaryToTextEncoding
		return ret
	}
	return o.BinaryToTextEncoding
}

// GetBinaryToTextEncodingOk returns a tuple with the BinaryToTextEncoding field value
// and a boolean to check if the value has been set.
func (o *LogsDecoderProcessor) GetBinaryToTextEncodingOk() (*LogsDecoderProcessorBinaryToTextEncoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinaryToTextEncoding, true
}

// SetBinaryToTextEncoding sets field value.
func (o *LogsDecoderProcessor) SetBinaryToTextEncoding(v LogsDecoderProcessorBinaryToTextEncoding) {
	o.BinaryToTextEncoding = v
}

// GetInputRepresentation returns the InputRepresentation field value.
func (o *LogsDecoderProcessor) GetInputRepresentation() LogsDecoderProcessorInputRepresentation {
	if o == nil {
		var ret LogsDecoderProcessorInputRepresentation
		return ret
	}
	return o.InputRepresentation
}

// GetInputRepresentationOk returns a tuple with the InputRepresentation field value
// and a boolean to check if the value has been set.
func (o *LogsDecoderProcessor) GetInputRepresentationOk() (*LogsDecoderProcessorInputRepresentation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputRepresentation, true
}

// SetInputRepresentation sets field value.
func (o *LogsDecoderProcessor) SetInputRepresentation(v LogsDecoderProcessorInputRepresentation) {
	o.InputRepresentation = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsDecoderProcessor) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsDecoderProcessor) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsDecoderProcessor) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsDecoderProcessor) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsDecoderProcessor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsDecoderProcessor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsDecoderProcessor) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsDecoderProcessor) SetName(v string) {
	o.Name = &v
}

// GetSource returns the Source field value.
func (o *LogsDecoderProcessor) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LogsDecoderProcessor) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *LogsDecoderProcessor) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value.
func (o *LogsDecoderProcessor) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsDecoderProcessor) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *LogsDecoderProcessor) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *LogsDecoderProcessor) GetType() LogsDecoderProcessorType {
	if o == nil {
		var ret LogsDecoderProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsDecoderProcessor) GetTypeOk() (*LogsDecoderProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsDecoderProcessor) SetType(v LogsDecoderProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsDecoderProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["binary_to_text_encoding"] = o.BinaryToTextEncoding
	toSerialize["input_representation"] = o.InputRepresentation
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["source"] = o.Source
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsDecoderProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BinaryToTextEncoding *LogsDecoderProcessorBinaryToTextEncoding `json:"binary_to_text_encoding"`
		InputRepresentation  *LogsDecoderProcessorInputRepresentation  `json:"input_representation"`
		IsEnabled            *bool                                     `json:"is_enabled,omitempty"`
		Name                 *string                                   `json:"name,omitempty"`
		Source               *string                                   `json:"source"`
		Target               *string                                   `json:"target"`
		Type                 *LogsDecoderProcessorType                 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BinaryToTextEncoding == nil {
		return fmt.Errorf("required field binary_to_text_encoding missing")
	}
	if all.InputRepresentation == nil {
		return fmt.Errorf("required field input_representation missing")
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"binary_to_text_encoding", "input_representation", "is_enabled", "name", "source", "target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.BinaryToTextEncoding.IsValid() {
		hasInvalidField = true
	} else {
		o.BinaryToTextEncoding = *all.BinaryToTextEncoding
	}
	if !all.InputRepresentation.IsValid() {
		hasInvalidField = true
	} else {
		o.InputRepresentation = *all.InputRepresentation
	}
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
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
