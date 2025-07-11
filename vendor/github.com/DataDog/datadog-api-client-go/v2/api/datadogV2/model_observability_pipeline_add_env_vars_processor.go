// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAddEnvVarsProcessor The `add_env_vars` processor adds environment variable values to log events.
type ObservabilityPipelineAddEnvVarsProcessor struct {
	// The unique identifier for this component. Used to reference this processor in the pipeline.
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the input for this processor.
	Inputs []string `json:"inputs"`
	// The processor type. The value should always be `add_env_vars`.
	Type ObservabilityPipelineAddEnvVarsProcessorType `json:"type"`
	// A list of environment variable mappings to apply to log fields.
	Variables []ObservabilityPipelineAddEnvVarsProcessorVariable `json:"variables"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAddEnvVarsProcessor instantiates a new ObservabilityPipelineAddEnvVarsProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAddEnvVarsProcessor(id string, include string, inputs []string, typeVar ObservabilityPipelineAddEnvVarsProcessorType, variables []ObservabilityPipelineAddEnvVarsProcessorVariable) *ObservabilityPipelineAddEnvVarsProcessor {
	this := ObservabilityPipelineAddEnvVarsProcessor{}
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Type = typeVar
	this.Variables = variables
	return &this
}

// NewObservabilityPipelineAddEnvVarsProcessorWithDefaults instantiates a new ObservabilityPipelineAddEnvVarsProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAddEnvVarsProcessorWithDefaults() *ObservabilityPipelineAddEnvVarsProcessor {
	this := ObservabilityPipelineAddEnvVarsProcessor{}
	var typeVar ObservabilityPipelineAddEnvVarsProcessorType = OBSERVABILITYPIPELINEADDENVVARSPROCESSORTYPE_ADD_ENV_VARS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineAddEnvVarsProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddEnvVarsProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineAddEnvVarsProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineAddEnvVarsProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddEnvVarsProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineAddEnvVarsProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineAddEnvVarsProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddEnvVarsProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineAddEnvVarsProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineAddEnvVarsProcessor) GetType() ObservabilityPipelineAddEnvVarsProcessorType {
	if o == nil {
		var ret ObservabilityPipelineAddEnvVarsProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddEnvVarsProcessor) GetTypeOk() (*ObservabilityPipelineAddEnvVarsProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineAddEnvVarsProcessor) SetType(v ObservabilityPipelineAddEnvVarsProcessorType) {
	o.Type = v
}

// GetVariables returns the Variables field value.
func (o *ObservabilityPipelineAddEnvVarsProcessor) GetVariables() []ObservabilityPipelineAddEnvVarsProcessorVariable {
	if o == nil {
		var ret []ObservabilityPipelineAddEnvVarsProcessorVariable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddEnvVarsProcessor) GetVariablesOk() (*[]ObservabilityPipelineAddEnvVarsProcessorVariable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variables, true
}

// SetVariables sets field value.
func (o *ObservabilityPipelineAddEnvVarsProcessor) SetVariables(v []ObservabilityPipelineAddEnvVarsProcessorVariable) {
	o.Variables = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAddEnvVarsProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	toSerialize["type"] = o.Type
	toSerialize["variables"] = o.Variables

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAddEnvVarsProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id        *string                                             `json:"id"`
		Include   *string                                             `json:"include"`
		Inputs    *[]string                                           `json:"inputs"`
		Type      *ObservabilityPipelineAddEnvVarsProcessorType       `json:"type"`
		Variables *[]ObservabilityPipelineAddEnvVarsProcessorVariable `json:"variables"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Variables == nil {
		return fmt.Errorf("required field variables missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "include", "inputs", "type", "variables"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Include = *all.Include
	o.Inputs = *all.Inputs
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Variables = *all.Variables

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
