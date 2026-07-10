// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAddMetricTagsProcessor The `add_metric_tags` processor adds static tags to metrics.
//
// **Supported pipeline types:** metrics
type ObservabilityPipelineAddMetricTagsProcessor struct {
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which metrics this processor targets.
	Include string `json:"include"`
	// A list of static tags (key-value pairs) added to each metric processed by this component.
	Tags []ObservabilityPipelineFieldValue `json:"tags"`
	// The processor type. The value must be `add_metric_tags`.
	Type ObservabilityPipelineAddMetricTagsProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAddMetricTagsProcessor instantiates a new ObservabilityPipelineAddMetricTagsProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAddMetricTagsProcessor(enabled bool, id string, include string, tags []ObservabilityPipelineFieldValue, typeVar ObservabilityPipelineAddMetricTagsProcessorType) *ObservabilityPipelineAddMetricTagsProcessor {
	this := ObservabilityPipelineAddMetricTagsProcessor{}
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.Tags = tags
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineAddMetricTagsProcessorWithDefaults instantiates a new ObservabilityPipelineAddMetricTagsProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAddMetricTagsProcessorWithDefaults() *ObservabilityPipelineAddMetricTagsProcessor {
	this := ObservabilityPipelineAddMetricTagsProcessor{}
	var typeVar ObservabilityPipelineAddMetricTagsProcessorType = OBSERVABILITYPIPELINEADDMETRICTAGSPROCESSORTYPE_ADD_METRIC_TAGS
	this.Type = typeVar
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineAddMetricTagsProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineAddMetricTagsProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineAddMetricTagsProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineAddMetricTagsProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineAddMetricTagsProcessor) SetInclude(v string) {
	o.Include = v
}

// GetTags returns the Tags field value.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetTags() []ObservabilityPipelineFieldValue {
	if o == nil {
		var ret []ObservabilityPipelineFieldValue
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetTagsOk() (*[]ObservabilityPipelineFieldValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *ObservabilityPipelineAddMetricTagsProcessor) SetTags(v []ObservabilityPipelineFieldValue) {
	o.Tags = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetType() ObservabilityPipelineAddMetricTagsProcessorType {
	if o == nil {
		var ret ObservabilityPipelineAddMetricTagsProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddMetricTagsProcessor) GetTypeOk() (*ObservabilityPipelineAddMetricTagsProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineAddMetricTagsProcessor) SetType(v ObservabilityPipelineAddMetricTagsProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAddMetricTagsProcessor) MarshalJSON() ([]byte, error) {
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
func (o *ObservabilityPipelineAddMetricTagsProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string                                          `json:"display_name,omitempty"`
		Enabled     *bool                                            `json:"enabled"`
		Id          *string                                          `json:"id"`
		Include     *string                                          `json:"include"`
		Tags        *[]ObservabilityPipelineFieldValue               `json:"tags"`
		Type        *ObservabilityPipelineAddMetricTagsProcessorType `json:"type"`
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
