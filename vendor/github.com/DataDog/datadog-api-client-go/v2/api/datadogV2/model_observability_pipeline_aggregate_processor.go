// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAggregateProcessor The `aggregate` processor combines metrics that share the same name and tags into a single metric over a configurable interval.
//
// **Supported pipeline types:** metrics
type ObservabilityPipelineAggregateProcessor struct {
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which metrics this processor targets.
	Include string `json:"include"`
	// The interval, in seconds, over which metrics are aggregated.
	IntervalSecs int64 `json:"interval_secs"`
	// The aggregation mode applied to metrics that share the same name and tags within the interval.
	Mode ObservabilityPipelineAggregateProcessorMode `json:"mode"`
	// The processor type. The value must be `aggregate`.
	Type ObservabilityPipelineAggregateProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAggregateProcessor instantiates a new ObservabilityPipelineAggregateProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAggregateProcessor(enabled bool, id string, include string, intervalSecs int64, mode ObservabilityPipelineAggregateProcessorMode, typeVar ObservabilityPipelineAggregateProcessorType) *ObservabilityPipelineAggregateProcessor {
	this := ObservabilityPipelineAggregateProcessor{}
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.IntervalSecs = intervalSecs
	this.Mode = mode
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineAggregateProcessorWithDefaults instantiates a new ObservabilityPipelineAggregateProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAggregateProcessorWithDefaults() *ObservabilityPipelineAggregateProcessor {
	this := ObservabilityPipelineAggregateProcessor{}
	var typeVar ObservabilityPipelineAggregateProcessorType = OBSERVABILITYPIPELINEAGGREGATEPROCESSORTYPE_AGGREGATE
	this.Type = typeVar
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineAggregateProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAggregateProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineAggregateProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineAggregateProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineAggregateProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAggregateProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineAggregateProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineAggregateProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAggregateProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineAggregateProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineAggregateProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAggregateProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineAggregateProcessor) SetInclude(v string) {
	o.Include = v
}

// GetIntervalSecs returns the IntervalSecs field value.
func (o *ObservabilityPipelineAggregateProcessor) GetIntervalSecs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.IntervalSecs
}

// GetIntervalSecsOk returns a tuple with the IntervalSecs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAggregateProcessor) GetIntervalSecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntervalSecs, true
}

// SetIntervalSecs sets field value.
func (o *ObservabilityPipelineAggregateProcessor) SetIntervalSecs(v int64) {
	o.IntervalSecs = v
}

// GetMode returns the Mode field value.
func (o *ObservabilityPipelineAggregateProcessor) GetMode() ObservabilityPipelineAggregateProcessorMode {
	if o == nil {
		var ret ObservabilityPipelineAggregateProcessorMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAggregateProcessor) GetModeOk() (*ObservabilityPipelineAggregateProcessorMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ObservabilityPipelineAggregateProcessor) SetMode(v ObservabilityPipelineAggregateProcessorMode) {
	o.Mode = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineAggregateProcessor) GetType() ObservabilityPipelineAggregateProcessorType {
	if o == nil {
		var ret ObservabilityPipelineAggregateProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAggregateProcessor) GetTypeOk() (*ObservabilityPipelineAggregateProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineAggregateProcessor) SetType(v ObservabilityPipelineAggregateProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAggregateProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["interval_secs"] = o.IntervalSecs
	toSerialize["mode"] = o.Mode
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAggregateProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName  *string                                      `json:"display_name,omitempty"`
		Enabled      *bool                                        `json:"enabled"`
		Id           *string                                      `json:"id"`
		Include      *string                                      `json:"include"`
		IntervalSecs *int64                                       `json:"interval_secs"`
		Mode         *ObservabilityPipelineAggregateProcessorMode `json:"mode"`
		Type         *ObservabilityPipelineAggregateProcessorType `json:"type"`
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
	if all.IntervalSecs == nil {
		return fmt.Errorf("required field interval_secs missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "enabled", "id", "include", "interval_secs", "mode", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Include = *all.Include
	o.IntervalSecs = *all.IntervalSecs
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
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
