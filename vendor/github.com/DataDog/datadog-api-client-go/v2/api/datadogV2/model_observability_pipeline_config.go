// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfig Specifies the pipeline's configuration, including its sources, processors, and destinations.
type ObservabilityPipelineConfig struct {
	// A list of destination components where processed logs are sent.
	Destinations []ObservabilityPipelineConfigDestinationItem `json:"destinations"`
	// A list of processors that transform or enrich log data.
	Processors []ObservabilityPipelineConfigProcessorItem `json:"processors,omitempty"`
	// A list of configured data sources for the pipeline.
	Sources []ObservabilityPipelineConfigSourceItem `json:"sources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineConfig instantiates a new ObservabilityPipelineConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineConfig(destinations []ObservabilityPipelineConfigDestinationItem, sources []ObservabilityPipelineConfigSourceItem) *ObservabilityPipelineConfig {
	this := ObservabilityPipelineConfig{}
	this.Destinations = destinations
	this.Sources = sources
	return &this
}

// NewObservabilityPipelineConfigWithDefaults instantiates a new ObservabilityPipelineConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineConfigWithDefaults() *ObservabilityPipelineConfig {
	this := ObservabilityPipelineConfig{}
	return &this
}

// GetDestinations returns the Destinations field value.
func (o *ObservabilityPipelineConfig) GetDestinations() []ObservabilityPipelineConfigDestinationItem {
	if o == nil {
		var ret []ObservabilityPipelineConfigDestinationItem
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineConfig) GetDestinationsOk() (*[]ObservabilityPipelineConfigDestinationItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destinations, true
}

// SetDestinations sets field value.
func (o *ObservabilityPipelineConfig) SetDestinations(v []ObservabilityPipelineConfigDestinationItem) {
	o.Destinations = v
}

// GetProcessors returns the Processors field value if set, zero value otherwise.
func (o *ObservabilityPipelineConfig) GetProcessors() []ObservabilityPipelineConfigProcessorItem {
	if o == nil || o.Processors == nil {
		var ret []ObservabilityPipelineConfigProcessorItem
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineConfig) GetProcessorsOk() (*[]ObservabilityPipelineConfigProcessorItem, bool) {
	if o == nil || o.Processors == nil {
		return nil, false
	}
	return &o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *ObservabilityPipelineConfig) HasProcessors() bool {
	return o != nil && o.Processors != nil
}

// SetProcessors gets a reference to the given []ObservabilityPipelineConfigProcessorItem and assigns it to the Processors field.
func (o *ObservabilityPipelineConfig) SetProcessors(v []ObservabilityPipelineConfigProcessorItem) {
	o.Processors = v
}

// GetSources returns the Sources field value.
func (o *ObservabilityPipelineConfig) GetSources() []ObservabilityPipelineConfigSourceItem {
	if o == nil {
		var ret []ObservabilityPipelineConfigSourceItem
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineConfig) GetSourcesOk() (*[]ObservabilityPipelineConfigSourceItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *ObservabilityPipelineConfig) SetSources(v []ObservabilityPipelineConfigSourceItem) {
	o.Sources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["destinations"] = o.Destinations
	if o.Processors != nil {
		toSerialize["processors"] = o.Processors
	}
	toSerialize["sources"] = o.Sources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Destinations *[]ObservabilityPipelineConfigDestinationItem `json:"destinations"`
		Processors   []ObservabilityPipelineConfigProcessorItem    `json:"processors,omitempty"`
		Sources      *[]ObservabilityPipelineConfigSourceItem      `json:"sources"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Destinations == nil {
		return fmt.Errorf("required field destinations missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"destinations", "processors", "sources"})
	} else {
		return err
	}
	o.Destinations = *all.Destinations
	o.Processors = all.Processors
	o.Sources = *all.Sources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
