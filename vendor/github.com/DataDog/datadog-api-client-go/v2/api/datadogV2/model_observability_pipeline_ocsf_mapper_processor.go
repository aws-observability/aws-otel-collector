// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOcsfMapperProcessor The `ocsf_mapper` processor transforms logs into the OCSF schema using a predefined mapping configuration.
type ObservabilityPipelineOcsfMapperProcessor struct {
	// The unique identifier for this component. Used to reference this component in other parts of the pipeline.
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the `input` for this processor.
	Inputs []string `json:"inputs"`
	// A list of mapping rules to convert events to the OCSF format.
	Mappings []ObservabilityPipelineOcsfMapperProcessorMapping `json:"mappings"`
	// The processor type. The value should always be `ocsf_mapper`.
	Type ObservabilityPipelineOcsfMapperProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineOcsfMapperProcessor instantiates a new ObservabilityPipelineOcsfMapperProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineOcsfMapperProcessor(id string, include string, inputs []string, mappings []ObservabilityPipelineOcsfMapperProcessorMapping, typeVar ObservabilityPipelineOcsfMapperProcessorType) *ObservabilityPipelineOcsfMapperProcessor {
	this := ObservabilityPipelineOcsfMapperProcessor{}
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Mappings = mappings
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineOcsfMapperProcessorWithDefaults instantiates a new ObservabilityPipelineOcsfMapperProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineOcsfMapperProcessorWithDefaults() *ObservabilityPipelineOcsfMapperProcessor {
	this := ObservabilityPipelineOcsfMapperProcessor{}
	var typeVar ObservabilityPipelineOcsfMapperProcessorType = OBSERVABILITYPIPELINEOCSFMAPPERPROCESSORTYPE_OCSF_MAPPER
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineOcsfMapperProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMapperProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineOcsfMapperProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineOcsfMapperProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMapperProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineOcsfMapperProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineOcsfMapperProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMapperProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineOcsfMapperProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetMappings returns the Mappings field value.
func (o *ObservabilityPipelineOcsfMapperProcessor) GetMappings() []ObservabilityPipelineOcsfMapperProcessorMapping {
	if o == nil {
		var ret []ObservabilityPipelineOcsfMapperProcessorMapping
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMapperProcessor) GetMappingsOk() (*[]ObservabilityPipelineOcsfMapperProcessorMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mappings, true
}

// SetMappings sets field value.
func (o *ObservabilityPipelineOcsfMapperProcessor) SetMappings(v []ObservabilityPipelineOcsfMapperProcessorMapping) {
	o.Mappings = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineOcsfMapperProcessor) GetType() ObservabilityPipelineOcsfMapperProcessorType {
	if o == nil {
		var ret ObservabilityPipelineOcsfMapperProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMapperProcessor) GetTypeOk() (*ObservabilityPipelineOcsfMapperProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineOcsfMapperProcessor) SetType(v ObservabilityPipelineOcsfMapperProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineOcsfMapperProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	toSerialize["mappings"] = o.Mappings
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineOcsfMapperProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id       *string                                            `json:"id"`
		Include  *string                                            `json:"include"`
		Inputs   *[]string                                          `json:"inputs"`
		Mappings *[]ObservabilityPipelineOcsfMapperProcessorMapping `json:"mappings"`
		Type     *ObservabilityPipelineOcsfMapperProcessorType      `json:"type"`
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
	if all.Mappings == nil {
		return fmt.Errorf("required field mappings missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "include", "inputs", "mappings", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Include = *all.Include
	o.Inputs = *all.Inputs
	o.Mappings = *all.Mappings
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
