// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCustomProcessor The `custom_processor` processor transforms events using [Vector Remap Language (VRL)](https://vector.dev/docs/reference/vrl/) scripts with advanced filtering capabilities.
type ObservabilityPipelineCustomProcessor struct {
	// Whether this processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this processor.
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets. This field should always be set to `*` for the custom_processor processor.
	Include string `json:"include"`
	// Array of VRL remap rules.
	Remaps []ObservabilityPipelineCustomProcessorRemap `json:"remaps"`
	// The processor type. The value should always be `custom_processor`.
	Type ObservabilityPipelineCustomProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineCustomProcessor instantiates a new ObservabilityPipelineCustomProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineCustomProcessor(enabled bool, id string, include string, remaps []ObservabilityPipelineCustomProcessorRemap, typeVar ObservabilityPipelineCustomProcessorType) *ObservabilityPipelineCustomProcessor {
	this := ObservabilityPipelineCustomProcessor{}
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.Remaps = remaps
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineCustomProcessorWithDefaults instantiates a new ObservabilityPipelineCustomProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineCustomProcessorWithDefaults() *ObservabilityPipelineCustomProcessor {
	this := ObservabilityPipelineCustomProcessor{}
	var include string = "*"
	this.Include = include
	var typeVar ObservabilityPipelineCustomProcessorType = OBSERVABILITYPIPELINECUSTOMPROCESSORTYPE_CUSTOM_PROCESSOR
	this.Type = typeVar
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineCustomProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineCustomProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineCustomProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineCustomProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineCustomProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineCustomProcessor) SetInclude(v string) {
	o.Include = v
}

// GetRemaps returns the Remaps field value.
func (o *ObservabilityPipelineCustomProcessor) GetRemaps() []ObservabilityPipelineCustomProcessorRemap {
	if o == nil {
		var ret []ObservabilityPipelineCustomProcessorRemap
		return ret
	}
	return o.Remaps
}

// GetRemapsOk returns a tuple with the Remaps field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessor) GetRemapsOk() (*[]ObservabilityPipelineCustomProcessorRemap, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remaps, true
}

// SetRemaps sets field value.
func (o *ObservabilityPipelineCustomProcessor) SetRemaps(v []ObservabilityPipelineCustomProcessorRemap) {
	o.Remaps = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineCustomProcessor) GetType() ObservabilityPipelineCustomProcessorType {
	if o == nil {
		var ret ObservabilityPipelineCustomProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessor) GetTypeOk() (*ObservabilityPipelineCustomProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineCustomProcessor) SetType(v ObservabilityPipelineCustomProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineCustomProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["remaps"] = o.Remaps
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineCustomProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool                                        `json:"enabled"`
		Id      *string                                      `json:"id"`
		Include *string                                      `json:"include"`
		Remaps  *[]ObservabilityPipelineCustomProcessorRemap `json:"remaps"`
		Type    *ObservabilityPipelineCustomProcessorType    `json:"type"`
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
	if all.Remaps == nil {
		return fmt.Errorf("required field remaps missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "id", "include", "remaps", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Include = *all.Include
	o.Remaps = *all.Remaps
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
