// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricAllTagsAttributes Object containing the definition of a metric's indexed and ingested tags.
type MetricAllTagsAttributes struct {
	// List of ingested tags that are not indexed.
	IngestedTags []string `json:"ingested_tags,omitempty"`
	// List of indexed tags.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricAllTagsAttributes instantiates a new MetricAllTagsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricAllTagsAttributes() *MetricAllTagsAttributes {
	this := MetricAllTagsAttributes{}
	return &this
}

// NewMetricAllTagsAttributesWithDefaults instantiates a new MetricAllTagsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricAllTagsAttributesWithDefaults() *MetricAllTagsAttributes {
	this := MetricAllTagsAttributes{}
	return &this
}

// GetIngestedTags returns the IngestedTags field value if set, zero value otherwise.
func (o *MetricAllTagsAttributes) GetIngestedTags() []string {
	if o == nil || o.IngestedTags == nil {
		var ret []string
		return ret
	}
	return o.IngestedTags
}

// GetIngestedTagsOk returns a tuple with the IngestedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAllTagsAttributes) GetIngestedTagsOk() (*[]string, bool) {
	if o == nil || o.IngestedTags == nil {
		return nil, false
	}
	return &o.IngestedTags, true
}

// HasIngestedTags returns a boolean if a field has been set.
func (o *MetricAllTagsAttributes) HasIngestedTags() bool {
	return o != nil && o.IngestedTags != nil
}

// SetIngestedTags gets a reference to the given []string and assigns it to the IngestedTags field.
func (o *MetricAllTagsAttributes) SetIngestedTags(v []string) {
	o.IngestedTags = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MetricAllTagsAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAllTagsAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MetricAllTagsAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MetricAllTagsAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricAllTagsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IngestedTags != nil {
		toSerialize["ingested_tags"] = o.IngestedTags
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricAllTagsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IngestedTags []string `json:"ingested_tags,omitempty"`
		Tags         []string `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ingested_tags", "tags"})
	} else {
		return err
	}
	o.IngestedTags = all.IngestedTags
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
