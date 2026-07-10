// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineTagCardinalityLimitProcessor The `tag_cardinality_limit` processor caps the number of distinct tag value combinations on metrics, dropping tags or events once the limit is exceeded.
//
// **Supported pipeline types:** metrics
type ObservabilityPipelineTagCardinalityLimitProcessor struct {
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which metrics this processor targets.
	Include string `json:"include"`
	// The action to take when the cardinality limit is exceeded.
	LimitExceededAction ObservabilityPipelineTagCardinalityLimitProcessorAction `json:"limit_exceeded_action"`
	// A list of per-metric cardinality overrides that take precedence over the default `value_limit`.
	PerMetricLimits []ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit `json:"per_metric_limits,omitempty"`
	// The processor type. The value must be `tag_cardinality_limit`.
	Type ObservabilityPipelineTagCardinalityLimitProcessorType `json:"type"`
	// The default maximum number of distinct tag value combinations allowed per metric.
	ValueLimit int64 `json:"value_limit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineTagCardinalityLimitProcessor instantiates a new ObservabilityPipelineTagCardinalityLimitProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineTagCardinalityLimitProcessor(enabled bool, id string, include string, limitExceededAction ObservabilityPipelineTagCardinalityLimitProcessorAction, typeVar ObservabilityPipelineTagCardinalityLimitProcessorType, valueLimit int64) *ObservabilityPipelineTagCardinalityLimitProcessor {
	this := ObservabilityPipelineTagCardinalityLimitProcessor{}
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.LimitExceededAction = limitExceededAction
	this.Type = typeVar
	this.ValueLimit = valueLimit
	return &this
}

// NewObservabilityPipelineTagCardinalityLimitProcessorWithDefaults instantiates a new ObservabilityPipelineTagCardinalityLimitProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineTagCardinalityLimitProcessorWithDefaults() *ObservabilityPipelineTagCardinalityLimitProcessor {
	this := ObservabilityPipelineTagCardinalityLimitProcessor{}
	var typeVar ObservabilityPipelineTagCardinalityLimitProcessorType = OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORTYPE_TAG_CARDINALITY_LIMIT
	this.Type = typeVar
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) SetInclude(v string) {
	o.Include = v
}

// GetLimitExceededAction returns the LimitExceededAction field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetLimitExceededAction() ObservabilityPipelineTagCardinalityLimitProcessorAction {
	if o == nil {
		var ret ObservabilityPipelineTagCardinalityLimitProcessorAction
		return ret
	}
	return o.LimitExceededAction
}

// GetLimitExceededActionOk returns a tuple with the LimitExceededAction field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetLimitExceededActionOk() (*ObservabilityPipelineTagCardinalityLimitProcessorAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitExceededAction, true
}

// SetLimitExceededAction sets field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) SetLimitExceededAction(v ObservabilityPipelineTagCardinalityLimitProcessorAction) {
	o.LimitExceededAction = v
}

// GetPerMetricLimits returns the PerMetricLimits field value if set, zero value otherwise.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetPerMetricLimits() []ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit {
	if o == nil || o.PerMetricLimits == nil {
		var ret []ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit
		return ret
	}
	return o.PerMetricLimits
}

// GetPerMetricLimitsOk returns a tuple with the PerMetricLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetPerMetricLimitsOk() (*[]ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit, bool) {
	if o == nil || o.PerMetricLimits == nil {
		return nil, false
	}
	return &o.PerMetricLimits, true
}

// HasPerMetricLimits returns a boolean if a field has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) HasPerMetricLimits() bool {
	return o != nil && o.PerMetricLimits != nil
}

// SetPerMetricLimits gets a reference to the given []ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit and assigns it to the PerMetricLimits field.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) SetPerMetricLimits(v []ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit) {
	o.PerMetricLimits = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetType() ObservabilityPipelineTagCardinalityLimitProcessorType {
	if o == nil {
		var ret ObservabilityPipelineTagCardinalityLimitProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetTypeOk() (*ObservabilityPipelineTagCardinalityLimitProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) SetType(v ObservabilityPipelineTagCardinalityLimitProcessorType) {
	o.Type = v
}

// GetValueLimit returns the ValueLimit field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetValueLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ValueLimit
}

// GetValueLimitOk returns a tuple with the ValueLimit field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) GetValueLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueLimit, true
}

// SetValueLimit sets field value.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) SetValueLimit(v int64) {
	o.ValueLimit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineTagCardinalityLimitProcessor) MarshalJSON() ([]byte, error) {
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
	toSerialize["limit_exceeded_action"] = o.LimitExceededAction
	if o.PerMetricLimits != nil {
		toSerialize["per_metric_limits"] = o.PerMetricLimits
	}
	toSerialize["type"] = o.Type
	toSerialize["value_limit"] = o.ValueLimit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineTagCardinalityLimitProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName         *string                                                           `json:"display_name,omitempty"`
		Enabled             *bool                                                             `json:"enabled"`
		Id                  *string                                                           `json:"id"`
		Include             *string                                                           `json:"include"`
		LimitExceededAction *ObservabilityPipelineTagCardinalityLimitProcessorAction          `json:"limit_exceeded_action"`
		PerMetricLimits     []ObservabilityPipelineTagCardinalityLimitProcessorPerMetricLimit `json:"per_metric_limits,omitempty"`
		Type                *ObservabilityPipelineTagCardinalityLimitProcessorType            `json:"type"`
		ValueLimit          *int64                                                            `json:"value_limit"`
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
	if all.LimitExceededAction == nil {
		return fmt.Errorf("required field limit_exceeded_action missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.ValueLimit == nil {
		return fmt.Errorf("required field value_limit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "enabled", "id", "include", "limit_exceeded_action", "per_metric_limits", "type", "value_limit"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Include = *all.Include
	if !all.LimitExceededAction.IsValid() {
		hasInvalidField = true
	} else {
		o.LimitExceededAction = *all.LimitExceededAction
	}
	o.PerMetricLimits = all.PerMetricLimits
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.ValueLimit = *all.ValueLimit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
