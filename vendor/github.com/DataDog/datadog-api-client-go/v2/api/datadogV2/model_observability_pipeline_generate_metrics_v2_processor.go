// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGenerateMetricsV2Processor The `generate_metrics` processor creates custom metrics from logs.
// Metrics can be counters, gauges, or distributions and optionally grouped by log fields.
// The generated metrics must be routed to a metrics destination using the input `<processor-id>.metrics`.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineGenerateMetricsV2Processor struct {
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this component. Used to reference this component in other parts of the pipeline.
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include *string `json:"include,omitempty"`
	// Configuration for generating individual metrics.
	Metrics []ObservabilityPipelineGeneratedMetric `json:"metrics,omitempty"`
	// The processor type. Always `generate_metrics`.
	Type ObservabilityPipelineGenerateMetricsV2ProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineGenerateMetricsV2Processor instantiates a new ObservabilityPipelineGenerateMetricsV2Processor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineGenerateMetricsV2Processor(enabled bool, id string, typeVar ObservabilityPipelineGenerateMetricsV2ProcessorType) *ObservabilityPipelineGenerateMetricsV2Processor {
	this := ObservabilityPipelineGenerateMetricsV2Processor{}
	this.Enabled = enabled
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineGenerateMetricsV2ProcessorWithDefaults instantiates a new ObservabilityPipelineGenerateMetricsV2Processor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineGenerateMetricsV2ProcessorWithDefaults() *ObservabilityPipelineGenerateMetricsV2Processor {
	this := ObservabilityPipelineGenerateMetricsV2Processor{}
	var typeVar ObservabilityPipelineGenerateMetricsV2ProcessorType = OBSERVABILITYPIPELINEGENERATEMETRICSV2PROCESSORTYPE_GENERATE_METRICS
	this.Type = typeVar
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetInclude() string {
	if o == nil || o.Include == nil {
		var ret string
		return ret
	}
	return *o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetIncludeOk() (*string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) HasInclude() bool {
	return o != nil && o.Include != nil
}

// SetInclude gets a reference to the given string and assigns it to the Include field.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) SetInclude(v string) {
	o.Include = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetMetrics() []ObservabilityPipelineGeneratedMetric {
	if o == nil || o.Metrics == nil {
		var ret []ObservabilityPipelineGeneratedMetric
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetMetricsOk() (*[]ObservabilityPipelineGeneratedMetric, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []ObservabilityPipelineGeneratedMetric and assigns it to the Metrics field.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) SetMetrics(v []ObservabilityPipelineGeneratedMetric) {
	o.Metrics = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetType() ObservabilityPipelineGenerateMetricsV2ProcessorType {
	if o == nil {
		var ret ObservabilityPipelineGenerateMetricsV2ProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) GetTypeOk() (*ObservabilityPipelineGenerateMetricsV2ProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) SetType(v ObservabilityPipelineGenerateMetricsV2ProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineGenerateMetricsV2Processor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineGenerateMetricsV2Processor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string                                              `json:"display_name,omitempty"`
		Enabled     *bool                                                `json:"enabled"`
		Id          *string                                              `json:"id"`
		Include     *string                                              `json:"include,omitempty"`
		Metrics     []ObservabilityPipelineGeneratedMetric               `json:"metrics,omitempty"`
		Type        *ObservabilityPipelineGenerateMetricsV2ProcessorType `json:"type"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "enabled", "id", "include", "metrics", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Include = all.Include
	o.Metrics = all.Metrics
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
