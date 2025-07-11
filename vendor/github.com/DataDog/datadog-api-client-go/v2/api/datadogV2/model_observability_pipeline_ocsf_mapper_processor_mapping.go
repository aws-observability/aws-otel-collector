// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOcsfMapperProcessorMapping Defines how specific events are transformed to OCSF using a mapping configuration.
type ObservabilityPipelineOcsfMapperProcessorMapping struct {
	// A Datadog search query used to select the logs that this mapping should apply to.
	Include string `json:"include"`
	// Defines a single mapping rule for transforming logs into the OCSF schema.
	Mapping ObservabilityPipelineOcsfMapperProcessorMappingMapping `json:"mapping"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineOcsfMapperProcessorMapping instantiates a new ObservabilityPipelineOcsfMapperProcessorMapping object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineOcsfMapperProcessorMapping(include string, mapping ObservabilityPipelineOcsfMapperProcessorMappingMapping) *ObservabilityPipelineOcsfMapperProcessorMapping {
	this := ObservabilityPipelineOcsfMapperProcessorMapping{}
	this.Include = include
	this.Mapping = mapping
	return &this
}

// NewObservabilityPipelineOcsfMapperProcessorMappingWithDefaults instantiates a new ObservabilityPipelineOcsfMapperProcessorMapping object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineOcsfMapperProcessorMappingWithDefaults() *ObservabilityPipelineOcsfMapperProcessorMapping {
	this := ObservabilityPipelineOcsfMapperProcessorMapping{}
	return &this
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineOcsfMapperProcessorMapping) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMapperProcessorMapping) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineOcsfMapperProcessorMapping) SetInclude(v string) {
	o.Include = v
}

// GetMapping returns the Mapping field value.
func (o *ObservabilityPipelineOcsfMapperProcessorMapping) GetMapping() ObservabilityPipelineOcsfMapperProcessorMappingMapping {
	if o == nil {
		var ret ObservabilityPipelineOcsfMapperProcessorMappingMapping
		return ret
	}
	return o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMapperProcessorMapping) GetMappingOk() (*ObservabilityPipelineOcsfMapperProcessorMappingMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapping, true
}

// SetMapping sets field value.
func (o *ObservabilityPipelineOcsfMapperProcessorMapping) SetMapping(v ObservabilityPipelineOcsfMapperProcessorMappingMapping) {
	o.Mapping = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineOcsfMapperProcessorMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["include"] = o.Include
	toSerialize["mapping"] = o.Mapping

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineOcsfMapperProcessorMapping) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Include *string                                                 `json:"include"`
		Mapping *ObservabilityPipelineOcsfMapperProcessorMappingMapping `json:"mapping"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.Mapping == nil {
		return fmt.Errorf("required field mapping missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include", "mapping"})
	} else {
		return err
	}
	o.Include = *all.Include
	o.Mapping = *all.Mapping

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
