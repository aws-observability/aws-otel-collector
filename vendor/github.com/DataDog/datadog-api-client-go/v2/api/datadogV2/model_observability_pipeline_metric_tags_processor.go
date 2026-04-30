// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricTagsProcessor The `metric_tags` processor filters metrics based on their tags using Datadog tag key patterns.
//
// **Supported pipeline types:** metrics
type ObservabilityPipelineMetricTagsProcessor struct {
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query that determines which metrics the processor targets.
	Include string `json:"include"`
	// A list of rules for filtering metric tags.
	Rules []ObservabilityPipelineMetricTagsProcessorRule `json:"rules"`
	// The processor type. The value should always be `metric_tags`.
	Type ObservabilityPipelineMetricTagsProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineMetricTagsProcessor instantiates a new ObservabilityPipelineMetricTagsProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineMetricTagsProcessor(enabled bool, id string, include string, rules []ObservabilityPipelineMetricTagsProcessorRule, typeVar ObservabilityPipelineMetricTagsProcessorType) *ObservabilityPipelineMetricTagsProcessor {
	this := ObservabilityPipelineMetricTagsProcessor{}
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.Rules = rules
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineMetricTagsProcessorWithDefaults instantiates a new ObservabilityPipelineMetricTagsProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineMetricTagsProcessorWithDefaults() *ObservabilityPipelineMetricTagsProcessor {
	this := ObservabilityPipelineMetricTagsProcessor{}
	var typeVar ObservabilityPipelineMetricTagsProcessorType = OBSERVABILITYPIPELINEMETRICTAGSPROCESSORTYPE_METRIC_TAGS
	this.Type = typeVar
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineMetricTagsProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricTagsProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineMetricTagsProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineMetricTagsProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineMetricTagsProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricTagsProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineMetricTagsProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineMetricTagsProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricTagsProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineMetricTagsProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineMetricTagsProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricTagsProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineMetricTagsProcessor) SetInclude(v string) {
	o.Include = v
}

// GetRules returns the Rules field value.
func (o *ObservabilityPipelineMetricTagsProcessor) GetRules() []ObservabilityPipelineMetricTagsProcessorRule {
	if o == nil {
		var ret []ObservabilityPipelineMetricTagsProcessorRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricTagsProcessor) GetRulesOk() (*[]ObservabilityPipelineMetricTagsProcessorRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *ObservabilityPipelineMetricTagsProcessor) SetRules(v []ObservabilityPipelineMetricTagsProcessorRule) {
	o.Rules = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineMetricTagsProcessor) GetType() ObservabilityPipelineMetricTagsProcessorType {
	if o == nil {
		var ret ObservabilityPipelineMetricTagsProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricTagsProcessor) GetTypeOk() (*ObservabilityPipelineMetricTagsProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineMetricTagsProcessor) SetType(v ObservabilityPipelineMetricTagsProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineMetricTagsProcessor) MarshalJSON() ([]byte, error) {
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
	toSerialize["rules"] = o.Rules
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineMetricTagsProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string                                         `json:"display_name,omitempty"`
		Enabled     *bool                                           `json:"enabled"`
		Id          *string                                         `json:"id"`
		Include     *string                                         `json:"include"`
		Rules       *[]ObservabilityPipelineMetricTagsProcessorRule `json:"rules"`
		Type        *ObservabilityPipelineMetricTagsProcessorType   `json:"type"`
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
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "enabled", "id", "include", "rules", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Include = *all.Include
	o.Rules = *all.Rules
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
