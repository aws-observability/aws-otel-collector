// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineParseGrokProcessor The `parse_grok` processor extracts structured fields from unstructured log messages using Grok patterns.
type ObservabilityPipelineParseGrokProcessor struct {
	// If set to `true`, disables the default Grok rules provided by Datadog.
	DisableLibraryRules *bool `json:"disable_library_rules,omitempty"`
	// A unique identifier for this processor.
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The list of Grok parsing rules. If multiple matching rules are provided, they are evaluated in order. The first successful match is applied.
	Rules []ObservabilityPipelineParseGrokProcessorRule `json:"rules"`
	// The processor type. The value should always be `parse_grok`.
	Type ObservabilityPipelineParseGrokProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineParseGrokProcessor instantiates a new ObservabilityPipelineParseGrokProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineParseGrokProcessor(id string, include string, inputs []string, rules []ObservabilityPipelineParseGrokProcessorRule, typeVar ObservabilityPipelineParseGrokProcessorType) *ObservabilityPipelineParseGrokProcessor {
	this := ObservabilityPipelineParseGrokProcessor{}
	var disableLibraryRules bool = false
	this.DisableLibraryRules = &disableLibraryRules
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Rules = rules
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineParseGrokProcessorWithDefaults instantiates a new ObservabilityPipelineParseGrokProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineParseGrokProcessorWithDefaults() *ObservabilityPipelineParseGrokProcessor {
	this := ObservabilityPipelineParseGrokProcessor{}
	var disableLibraryRules bool = false
	this.DisableLibraryRules = &disableLibraryRules
	var typeVar ObservabilityPipelineParseGrokProcessorType = OBSERVABILITYPIPELINEPARSEGROKPROCESSORTYPE_PARSE_GROK
	this.Type = typeVar
	return &this
}

// GetDisableLibraryRules returns the DisableLibraryRules field value if set, zero value otherwise.
func (o *ObservabilityPipelineParseGrokProcessor) GetDisableLibraryRules() bool {
	if o == nil || o.DisableLibraryRules == nil {
		var ret bool
		return ret
	}
	return *o.DisableLibraryRules
}

// GetDisableLibraryRulesOk returns a tuple with the DisableLibraryRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessor) GetDisableLibraryRulesOk() (*bool, bool) {
	if o == nil || o.DisableLibraryRules == nil {
		return nil, false
	}
	return o.DisableLibraryRules, true
}

// HasDisableLibraryRules returns a boolean if a field has been set.
func (o *ObservabilityPipelineParseGrokProcessor) HasDisableLibraryRules() bool {
	return o != nil && o.DisableLibraryRules != nil
}

// SetDisableLibraryRules gets a reference to the given bool and assigns it to the DisableLibraryRules field.
func (o *ObservabilityPipelineParseGrokProcessor) SetDisableLibraryRules(v bool) {
	o.DisableLibraryRules = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineParseGrokProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineParseGrokProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineParseGrokProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineParseGrokProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineParseGrokProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineParseGrokProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetRules returns the Rules field value.
func (o *ObservabilityPipelineParseGrokProcessor) GetRules() []ObservabilityPipelineParseGrokProcessorRule {
	if o == nil {
		var ret []ObservabilityPipelineParseGrokProcessorRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessor) GetRulesOk() (*[]ObservabilityPipelineParseGrokProcessorRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *ObservabilityPipelineParseGrokProcessor) SetRules(v []ObservabilityPipelineParseGrokProcessorRule) {
	o.Rules = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineParseGrokProcessor) GetType() ObservabilityPipelineParseGrokProcessorType {
	if o == nil {
		var ret ObservabilityPipelineParseGrokProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessor) GetTypeOk() (*ObservabilityPipelineParseGrokProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineParseGrokProcessor) SetType(v ObservabilityPipelineParseGrokProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineParseGrokProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisableLibraryRules != nil {
		toSerialize["disable_library_rules"] = o.DisableLibraryRules
	}
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	toSerialize["rules"] = o.Rules
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineParseGrokProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisableLibraryRules *bool                                          `json:"disable_library_rules,omitempty"`
		Id                  *string                                        `json:"id"`
		Include             *string                                        `json:"include"`
		Inputs              *[]string                                      `json:"inputs"`
		Rules               *[]ObservabilityPipelineParseGrokProcessorRule `json:"rules"`
		Type                *ObservabilityPipelineParseGrokProcessorType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"disable_library_rules", "id", "include", "inputs", "rules", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisableLibraryRules = all.DisableLibraryRules
	o.Id = *all.Id
	o.Include = *all.Include
	o.Inputs = *all.Inputs
	o.Rules = *all.Rules
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
