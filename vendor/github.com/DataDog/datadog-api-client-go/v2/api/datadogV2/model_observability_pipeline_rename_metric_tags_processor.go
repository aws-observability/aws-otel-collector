// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineRenameMetricTagsProcessor The `rename_metric_tags` processor changes the keys of tags on metrics.
//
// **Supported pipeline types:** metrics
type ObservabilityPipelineRenameMetricTagsProcessor struct {
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which metrics this processor targets.
	Include string `json:"include"`
	// A list of rename rules specifying which tag keys to rename on each metric.
	Tags []ObservabilityPipelineRenameMetricTagsProcessorTag `json:"tags"`
	// The processor type. The value must be `rename_metric_tags`.
	Type ObservabilityPipelineRenameMetricTagsProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineRenameMetricTagsProcessor instantiates a new ObservabilityPipelineRenameMetricTagsProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineRenameMetricTagsProcessor(enabled bool, id string, include string, tags []ObservabilityPipelineRenameMetricTagsProcessorTag, typeVar ObservabilityPipelineRenameMetricTagsProcessorType) *ObservabilityPipelineRenameMetricTagsProcessor {
	this := ObservabilityPipelineRenameMetricTagsProcessor{}
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.Tags = tags
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineRenameMetricTagsProcessorWithDefaults instantiates a new ObservabilityPipelineRenameMetricTagsProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineRenameMetricTagsProcessorWithDefaults() *ObservabilityPipelineRenameMetricTagsProcessor {
	this := ObservabilityPipelineRenameMetricTagsProcessor{}
	var typeVar ObservabilityPipelineRenameMetricTagsProcessorType = OBSERVABILITYPIPELINERENAMEMETRICTAGSPROCESSORTYPE_RENAME_METRIC_TAGS
	this.Type = typeVar
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) SetInclude(v string) {
	o.Include = v
}

// GetTags returns the Tags field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetTags() []ObservabilityPipelineRenameMetricTagsProcessorTag {
	if o == nil {
		var ret []ObservabilityPipelineRenameMetricTagsProcessorTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetTagsOk() (*[]ObservabilityPipelineRenameMetricTagsProcessorTag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) SetTags(v []ObservabilityPipelineRenameMetricTagsProcessorTag) {
	o.Tags = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetType() ObservabilityPipelineRenameMetricTagsProcessorType {
	if o == nil {
		var ret ObservabilityPipelineRenameMetricTagsProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) GetTypeOk() (*ObservabilityPipelineRenameMetricTagsProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) SetType(v ObservabilityPipelineRenameMetricTagsProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineRenameMetricTagsProcessor) MarshalJSON() ([]byte, error) {
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
	toSerialize["tags"] = o.Tags
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineRenameMetricTagsProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string                                              `json:"display_name,omitempty"`
		Enabled     *bool                                                `json:"enabled"`
		Id          *string                                              `json:"id"`
		Include     *string                                              `json:"include"`
		Tags        *[]ObservabilityPipelineRenameMetricTagsProcessorTag `json:"tags"`
		Type        *ObservabilityPipelineRenameMetricTagsProcessorType  `json:"type"`
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
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "enabled", "id", "include", "tags", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Include = *all.Include
	o.Tags = *all.Tags
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
