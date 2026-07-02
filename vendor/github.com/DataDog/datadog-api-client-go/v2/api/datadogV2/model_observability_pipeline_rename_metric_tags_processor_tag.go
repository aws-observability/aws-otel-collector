// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineRenameMetricTagsProcessorTag Defines how to rename a tag on metric events.
type ObservabilityPipelineRenameMetricTagsProcessorTag struct {
	// The new tag key to assign in place of the original.
	RenameTo string `json:"rename_to"`
	// The original tag key on the metric event.
	Tag string `json:"tag"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineRenameMetricTagsProcessorTag instantiates a new ObservabilityPipelineRenameMetricTagsProcessorTag object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineRenameMetricTagsProcessorTag(renameTo string, tag string) *ObservabilityPipelineRenameMetricTagsProcessorTag {
	this := ObservabilityPipelineRenameMetricTagsProcessorTag{}
	this.RenameTo = renameTo
	this.Tag = tag
	return &this
}

// NewObservabilityPipelineRenameMetricTagsProcessorTagWithDefaults instantiates a new ObservabilityPipelineRenameMetricTagsProcessorTag object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineRenameMetricTagsProcessorTagWithDefaults() *ObservabilityPipelineRenameMetricTagsProcessorTag {
	this := ObservabilityPipelineRenameMetricTagsProcessorTag{}
	return &this
}

// GetRenameTo returns the RenameTo field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessorTag) GetRenameTo() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RenameTo
}

// GetRenameToOk returns a tuple with the RenameTo field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameMetricTagsProcessorTag) GetRenameToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RenameTo, true
}

// SetRenameTo sets field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessorTag) SetRenameTo(v string) {
	o.RenameTo = v
}

// GetTag returns the Tag field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessorTag) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameMetricTagsProcessorTag) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value.
func (o *ObservabilityPipelineRenameMetricTagsProcessorTag) SetTag(v string) {
	o.Tag = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineRenameMetricTagsProcessorTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rename_to"] = o.RenameTo
	toSerialize["tag"] = o.Tag

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineRenameMetricTagsProcessorTag) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RenameTo *string `json:"rename_to"`
		Tag      *string `json:"tag"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RenameTo == nil {
		return fmt.Errorf("required field rename_to missing")
	}
	if all.Tag == nil {
		return fmt.Errorf("required field tag missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rename_to", "tag"})
	} else {
		return err
	}
	o.RenameTo = *all.RenameTo
	o.Tag = *all.Tag

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
