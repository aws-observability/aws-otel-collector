// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigProcessorGroup A group of processors.
type ObservabilityPipelineConfigProcessorGroup struct {
	// Whether this processor group is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for the processor group.
	Id string `json:"id"`
	// Conditional expression for when this processor group should execute.
	Include string `json:"include"`
	// A list of IDs for components whose output is used as the input for this processor group.
	Inputs []string `json:"inputs"`
	// Processors applied sequentially within this group. Events flow through each processor in order.
	Processors []ObservabilityPipelineConfigProcessorItem `json:"processors"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineConfigProcessorGroup instantiates a new ObservabilityPipelineConfigProcessorGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineConfigProcessorGroup(enabled bool, id string, include string, inputs []string, processors []ObservabilityPipelineConfigProcessorItem) *ObservabilityPipelineConfigProcessorGroup {
	this := ObservabilityPipelineConfigProcessorGroup{}
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Processors = processors
	return &this
}

// NewObservabilityPipelineConfigProcessorGroupWithDefaults instantiates a new ObservabilityPipelineConfigProcessorGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineConfigProcessorGroupWithDefaults() *ObservabilityPipelineConfigProcessorGroup {
	this := ObservabilityPipelineConfigProcessorGroup{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineConfigProcessorGroup) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineConfigProcessorGroup) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineConfigProcessorGroup) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineConfigProcessorGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineConfigProcessorGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineConfigProcessorGroup) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineConfigProcessorGroup) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineConfigProcessorGroup) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineConfigProcessorGroup) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineConfigProcessorGroup) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineConfigProcessorGroup) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineConfigProcessorGroup) SetInputs(v []string) {
	o.Inputs = v
}

// GetProcessors returns the Processors field value.
func (o *ObservabilityPipelineConfigProcessorGroup) GetProcessors() []ObservabilityPipelineConfigProcessorItem {
	if o == nil {
		var ret []ObservabilityPipelineConfigProcessorItem
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineConfigProcessorGroup) GetProcessorsOk() (*[]ObservabilityPipelineConfigProcessorItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Processors, true
}

// SetProcessors sets field value.
func (o *ObservabilityPipelineConfigProcessorGroup) SetProcessors(v []ObservabilityPipelineConfigProcessorItem) {
	o.Processors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineConfigProcessorGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	toSerialize["processors"] = o.Processors

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineConfigProcessorGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled    *bool                                       `json:"enabled"`
		Id         *string                                     `json:"id"`
		Include    *string                                     `json:"include"`
		Inputs     *[]string                                   `json:"inputs"`
		Processors *[]ObservabilityPipelineConfigProcessorItem `json:"processors"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
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
	if all.Processors == nil {
		return fmt.Errorf("required field processors missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "id", "include", "inputs", "processors"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Include = *all.Include
	o.Inputs = *all.Inputs
	o.Processors = *all.Processors

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
