// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDedupeProcessor The `dedupe` processor removes duplicate fields in log events.
type ObservabilityPipelineDedupeProcessor struct {
	// A list of log field paths to check for duplicates.
	Fields []string `json:"fields"`
	// The unique identifier for this processor.
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the input for this processor.
	Inputs []string `json:"inputs"`
	// The deduplication mode to apply to the fields.
	Mode ObservabilityPipelineDedupeProcessorMode `json:"mode"`
	// The processor type. The value should always be `dedupe`.
	Type ObservabilityPipelineDedupeProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineDedupeProcessor instantiates a new ObservabilityPipelineDedupeProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineDedupeProcessor(fields []string, id string, include string, inputs []string, mode ObservabilityPipelineDedupeProcessorMode, typeVar ObservabilityPipelineDedupeProcessorType) *ObservabilityPipelineDedupeProcessor {
	this := ObservabilityPipelineDedupeProcessor{}
	this.Fields = fields
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Mode = mode
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineDedupeProcessorWithDefaults instantiates a new ObservabilityPipelineDedupeProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineDedupeProcessorWithDefaults() *ObservabilityPipelineDedupeProcessor {
	this := ObservabilityPipelineDedupeProcessor{}
	var typeVar ObservabilityPipelineDedupeProcessorType = OBSERVABILITYPIPELINEDEDUPEPROCESSORTYPE_DEDUPE
	this.Type = typeVar
	return &this
}

// GetFields returns the Fields field value.
func (o *ObservabilityPipelineDedupeProcessor) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDedupeProcessor) GetFieldsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value.
func (o *ObservabilityPipelineDedupeProcessor) SetFields(v []string) {
	o.Fields = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineDedupeProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDedupeProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineDedupeProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineDedupeProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDedupeProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineDedupeProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineDedupeProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDedupeProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineDedupeProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetMode returns the Mode field value.
func (o *ObservabilityPipelineDedupeProcessor) GetMode() ObservabilityPipelineDedupeProcessorMode {
	if o == nil {
		var ret ObservabilityPipelineDedupeProcessorMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDedupeProcessor) GetModeOk() (*ObservabilityPipelineDedupeProcessorMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ObservabilityPipelineDedupeProcessor) SetMode(v ObservabilityPipelineDedupeProcessorMode) {
	o.Mode = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineDedupeProcessor) GetType() ObservabilityPipelineDedupeProcessorType {
	if o == nil {
		var ret ObservabilityPipelineDedupeProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDedupeProcessor) GetTypeOk() (*ObservabilityPipelineDedupeProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineDedupeProcessor) SetType(v ObservabilityPipelineDedupeProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineDedupeProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fields"] = o.Fields
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	toSerialize["mode"] = o.Mode
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineDedupeProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields  *[]string                                 `json:"fields"`
		Id      *string                                   `json:"id"`
		Include *string                                   `json:"include"`
		Inputs  *[]string                                 `json:"inputs"`
		Mode    *ObservabilityPipelineDedupeProcessorMode `json:"mode"`
		Type    *ObservabilityPipelineDedupeProcessorType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Fields == nil {
		return fmt.Errorf("required field fields missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "id", "include", "inputs", "mode", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Fields = *all.Fields
	o.Id = *all.Id
	o.Include = *all.Include
	o.Inputs = *all.Inputs
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
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
