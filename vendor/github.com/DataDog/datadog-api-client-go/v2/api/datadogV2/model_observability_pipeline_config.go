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
	// The type of data being ingested. Defaults to `logs` if not specified.
	PipelineType *ObservabilityPipelineConfigPipelineType `json:"pipeline_type,omitempty"`
	// A list of processor groups that transform or enrich log data.
	ProcessorGroups []ObservabilityPipelineConfigProcessorGroup `json:"processor_groups,omitempty"`
	// A list of processor groups that transform or enrich log data.
	//
	// **Deprecated:** This field is deprecated, you should now use the processor_groups field.
	// Deprecated
	Processors []ObservabilityPipelineConfigProcessorGroup `json:"processors,omitempty"`
	// A list of configured data sources for the pipeline.
	Sources []ObservabilityPipelineConfigSourceItem `json:"sources"`
	// Set to `true` to continue using the legacy search syntax while migrating filter queries. After migrating all queries to the new syntax, set to `false`.
	// The legacy syntax is deprecated and will eventually be removed.
	// Requires Observability Pipelines Worker 2.11 or later.
	// Only applies to `logs` pipelines. This field is ignored for `metrics` pipelines.
	// See [Upgrade Your Filter Queries to the New Search Syntax](https://docs.datadoghq.com/observability_pipelines/guide/upgrade_your_filter_queries_to_the_new_search_syntax/) for more information.
	UseLegacySearchSyntax *bool `json:"use_legacy_search_syntax,omitempty"`
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
	var pipelineType ObservabilityPipelineConfigPipelineType = OBSERVABILITYPIPELINECONFIGPIPELINETYPE_LOGS
	this.PipelineType = &pipelineType
	this.Sources = sources
	return &this
}

// NewObservabilityPipelineConfigWithDefaults instantiates a new ObservabilityPipelineConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineConfigWithDefaults() *ObservabilityPipelineConfig {
	this := ObservabilityPipelineConfig{}
	var pipelineType ObservabilityPipelineConfigPipelineType = OBSERVABILITYPIPELINECONFIGPIPELINETYPE_LOGS
	this.PipelineType = &pipelineType
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

// GetPipelineType returns the PipelineType field value if set, zero value otherwise.
func (o *ObservabilityPipelineConfig) GetPipelineType() ObservabilityPipelineConfigPipelineType {
	if o == nil || o.PipelineType == nil {
		var ret ObservabilityPipelineConfigPipelineType
		return ret
	}
	return *o.PipelineType
}

// GetPipelineTypeOk returns a tuple with the PipelineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineConfig) GetPipelineTypeOk() (*ObservabilityPipelineConfigPipelineType, bool) {
	if o == nil || o.PipelineType == nil {
		return nil, false
	}
	return o.PipelineType, true
}

// HasPipelineType returns a boolean if a field has been set.
func (o *ObservabilityPipelineConfig) HasPipelineType() bool {
	return o != nil && o.PipelineType != nil
}

// SetPipelineType gets a reference to the given ObservabilityPipelineConfigPipelineType and assigns it to the PipelineType field.
func (o *ObservabilityPipelineConfig) SetPipelineType(v ObservabilityPipelineConfigPipelineType) {
	o.PipelineType = &v
}

// GetProcessorGroups returns the ProcessorGroups field value if set, zero value otherwise.
func (o *ObservabilityPipelineConfig) GetProcessorGroups() []ObservabilityPipelineConfigProcessorGroup {
	if o == nil || o.ProcessorGroups == nil {
		var ret []ObservabilityPipelineConfigProcessorGroup
		return ret
	}
	return o.ProcessorGroups
}

// GetProcessorGroupsOk returns a tuple with the ProcessorGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineConfig) GetProcessorGroupsOk() (*[]ObservabilityPipelineConfigProcessorGroup, bool) {
	if o == nil || o.ProcessorGroups == nil {
		return nil, false
	}
	return &o.ProcessorGroups, true
}

// HasProcessorGroups returns a boolean if a field has been set.
func (o *ObservabilityPipelineConfig) HasProcessorGroups() bool {
	return o != nil && o.ProcessorGroups != nil
}

// SetProcessorGroups gets a reference to the given []ObservabilityPipelineConfigProcessorGroup and assigns it to the ProcessorGroups field.
func (o *ObservabilityPipelineConfig) SetProcessorGroups(v []ObservabilityPipelineConfigProcessorGroup) {
	o.ProcessorGroups = v
}

// GetProcessors returns the Processors field value if set, zero value otherwise.
// Deprecated
func (o *ObservabilityPipelineConfig) GetProcessors() []ObservabilityPipelineConfigProcessorGroup {
	if o == nil || o.Processors == nil {
		var ret []ObservabilityPipelineConfigProcessorGroup
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ObservabilityPipelineConfig) GetProcessorsOk() (*[]ObservabilityPipelineConfigProcessorGroup, bool) {
	if o == nil || o.Processors == nil {
		return nil, false
	}
	return &o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *ObservabilityPipelineConfig) HasProcessors() bool {
	return o != nil && o.Processors != nil
}

// SetProcessors gets a reference to the given []ObservabilityPipelineConfigProcessorGroup and assigns it to the Processors field.
// Deprecated
func (o *ObservabilityPipelineConfig) SetProcessors(v []ObservabilityPipelineConfigProcessorGroup) {
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

// GetUseLegacySearchSyntax returns the UseLegacySearchSyntax field value if set, zero value otherwise.
func (o *ObservabilityPipelineConfig) GetUseLegacySearchSyntax() bool {
	if o == nil || o.UseLegacySearchSyntax == nil {
		var ret bool
		return ret
	}
	return *o.UseLegacySearchSyntax
}

// GetUseLegacySearchSyntaxOk returns a tuple with the UseLegacySearchSyntax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineConfig) GetUseLegacySearchSyntaxOk() (*bool, bool) {
	if o == nil || o.UseLegacySearchSyntax == nil {
		return nil, false
	}
	return o.UseLegacySearchSyntax, true
}

// HasUseLegacySearchSyntax returns a boolean if a field has been set.
func (o *ObservabilityPipelineConfig) HasUseLegacySearchSyntax() bool {
	return o != nil && o.UseLegacySearchSyntax != nil
}

// SetUseLegacySearchSyntax gets a reference to the given bool and assigns it to the UseLegacySearchSyntax field.
func (o *ObservabilityPipelineConfig) SetUseLegacySearchSyntax(v bool) {
	o.UseLegacySearchSyntax = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["destinations"] = o.Destinations
	if o.PipelineType != nil {
		toSerialize["pipeline_type"] = o.PipelineType
	}
	if o.ProcessorGroups != nil {
		toSerialize["processor_groups"] = o.ProcessorGroups
	}
	if o.Processors != nil {
		toSerialize["processors"] = o.Processors
	}
	toSerialize["sources"] = o.Sources
	if o.UseLegacySearchSyntax != nil {
		toSerialize["use_legacy_search_syntax"] = o.UseLegacySearchSyntax
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Destinations          *[]ObservabilityPipelineConfigDestinationItem `json:"destinations"`
		PipelineType          *ObservabilityPipelineConfigPipelineType      `json:"pipeline_type,omitempty"`
		ProcessorGroups       []ObservabilityPipelineConfigProcessorGroup   `json:"processor_groups,omitempty"`
		Processors            []ObservabilityPipelineConfigProcessorGroup   `json:"processors,omitempty"`
		Sources               *[]ObservabilityPipelineConfigSourceItem      `json:"sources"`
		UseLegacySearchSyntax *bool                                         `json:"use_legacy_search_syntax,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"destinations", "pipeline_type", "processor_groups", "processors", "sources", "use_legacy_search_syntax"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Destinations = *all.Destinations
	if all.PipelineType != nil && !all.PipelineType.IsValid() {
		hasInvalidField = true
	} else {
		o.PipelineType = all.PipelineType
	}
	o.ProcessorGroups = all.ProcessorGroups
	o.Processors = all.Processors
	o.Sources = *all.Sources
	o.UseLegacySearchSyntax = all.UseLegacySearchSyntax

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
