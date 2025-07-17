// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineReduceProcessor The `reduce` processor aggregates and merges logs based on matching keys and merge strategies.
type ObservabilityPipelineReduceProcessor struct {
	// A list of fields used to group log events for merging.
	GroupBy []string `json:"group_by"`
	// The unique identifier for this processor.
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the input for this processor.
	Inputs []string `json:"inputs"`
	// List of merge strategies defining how values from grouped events should be combined.
	MergeStrategies []ObservabilityPipelineReduceProcessorMergeStrategy `json:"merge_strategies"`
	// The processor type. The value should always be `reduce`.
	Type ObservabilityPipelineReduceProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineReduceProcessor instantiates a new ObservabilityPipelineReduceProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineReduceProcessor(groupBy []string, id string, include string, inputs []string, mergeStrategies []ObservabilityPipelineReduceProcessorMergeStrategy, typeVar ObservabilityPipelineReduceProcessorType) *ObservabilityPipelineReduceProcessor {
	this := ObservabilityPipelineReduceProcessor{}
	this.GroupBy = groupBy
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.MergeStrategies = mergeStrategies
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineReduceProcessorWithDefaults instantiates a new ObservabilityPipelineReduceProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineReduceProcessorWithDefaults() *ObservabilityPipelineReduceProcessor {
	this := ObservabilityPipelineReduceProcessor{}
	var typeVar ObservabilityPipelineReduceProcessorType = OBSERVABILITYPIPELINEREDUCEPROCESSORTYPE_REDUCE
	this.Type = typeVar
	return &this
}

// GetGroupBy returns the GroupBy field value.
func (o *ObservabilityPipelineReduceProcessor) GetGroupBy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineReduceProcessor) GetGroupByOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// SetGroupBy sets field value.
func (o *ObservabilityPipelineReduceProcessor) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineReduceProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineReduceProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineReduceProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineReduceProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineReduceProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineReduceProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineReduceProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineReduceProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineReduceProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetMergeStrategies returns the MergeStrategies field value.
func (o *ObservabilityPipelineReduceProcessor) GetMergeStrategies() []ObservabilityPipelineReduceProcessorMergeStrategy {
	if o == nil {
		var ret []ObservabilityPipelineReduceProcessorMergeStrategy
		return ret
	}
	return o.MergeStrategies
}

// GetMergeStrategiesOk returns a tuple with the MergeStrategies field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineReduceProcessor) GetMergeStrategiesOk() (*[]ObservabilityPipelineReduceProcessorMergeStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MergeStrategies, true
}

// SetMergeStrategies sets field value.
func (o *ObservabilityPipelineReduceProcessor) SetMergeStrategies(v []ObservabilityPipelineReduceProcessorMergeStrategy) {
	o.MergeStrategies = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineReduceProcessor) GetType() ObservabilityPipelineReduceProcessorType {
	if o == nil {
		var ret ObservabilityPipelineReduceProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineReduceProcessor) GetTypeOk() (*ObservabilityPipelineReduceProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineReduceProcessor) SetType(v ObservabilityPipelineReduceProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineReduceProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["group_by"] = o.GroupBy
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	toSerialize["merge_strategies"] = o.MergeStrategies
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineReduceProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupBy         *[]string                                            `json:"group_by"`
		Id              *string                                              `json:"id"`
		Include         *string                                              `json:"include"`
		Inputs          *[]string                                            `json:"inputs"`
		MergeStrategies *[]ObservabilityPipelineReduceProcessorMergeStrategy `json:"merge_strategies"`
		Type            *ObservabilityPipelineReduceProcessorType            `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GroupBy == nil {
		return fmt.Errorf("required field group_by missing")
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
	if all.MergeStrategies == nil {
		return fmt.Errorf("required field merge_strategies missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group_by", "id", "include", "inputs", "merge_strategies", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.GroupBy = *all.GroupBy
	o.Id = *all.Id
	o.Include = *all.Include
	o.Inputs = *all.Inputs
	o.MergeStrategies = *all.MergeStrategies
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
