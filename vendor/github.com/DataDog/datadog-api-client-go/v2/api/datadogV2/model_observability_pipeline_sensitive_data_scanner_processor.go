// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessor The `sensitive_data_scanner` processor detects and optionally redacts sensitive data in log events.
type ObservabilityPipelineSensitiveDataScannerProcessor struct {
	// The unique identifier for this component. Used to reference this component in other parts of the pipeline (e.g., as input to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// A list of rules for identifying and acting on sensitive data patterns.
	Rules []ObservabilityPipelineSensitiveDataScannerProcessorRule `json:"rules"`
	// The processor type. The value should always be `sensitive_data_scanner`.
	Type ObservabilityPipelineSensitiveDataScannerProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessor instantiates a new ObservabilityPipelineSensitiveDataScannerProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessor(id string, include string, inputs []string, rules []ObservabilityPipelineSensitiveDataScannerProcessorRule, typeVar ObservabilityPipelineSensitiveDataScannerProcessorType) *ObservabilityPipelineSensitiveDataScannerProcessor {
	this := ObservabilityPipelineSensitiveDataScannerProcessor{}
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Rules = rules
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessor {
	this := ObservabilityPipelineSensitiveDataScannerProcessor{}
	var typeVar ObservabilityPipelineSensitiveDataScannerProcessorType = OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORTYPE_SENSITIVE_DATA_SCANNER
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetRules returns the Rules field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) GetRules() []ObservabilityPipelineSensitiveDataScannerProcessorRule {
	if o == nil {
		var ret []ObservabilityPipelineSensitiveDataScannerProcessorRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) GetRulesOk() (*[]ObservabilityPipelineSensitiveDataScannerProcessorRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) SetRules(v []ObservabilityPipelineSensitiveDataScannerProcessorRule) {
	o.Rules = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) GetType() ObservabilityPipelineSensitiveDataScannerProcessorType {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) GetTypeOk() (*ObservabilityPipelineSensitiveDataScannerProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) SetType(v ObservabilityPipelineSensitiveDataScannerProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *ObservabilityPipelineSensitiveDataScannerProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id      *string                                                   `json:"id"`
		Include *string                                                   `json:"include"`
		Inputs  *[]string                                                 `json:"inputs"`
		Rules   *[]ObservabilityPipelineSensitiveDataScannerProcessorRule `json:"rules"`
		Type    *ObservabilityPipelineSensitiveDataScannerProcessorType   `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "include", "inputs", "rules", "type"})
	} else {
		return err
	}

	hasInvalidField := false
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
