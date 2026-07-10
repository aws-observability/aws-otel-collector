// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomForecastEntryTagFilter A tag filter that scopes a custom forecast entry to specific resource tags.
type CustomForecastEntryTagFilter struct {
	// The tag key to filter on.
	TagKey string `json:"tag_key"`
	// The tag value to filter on.
	TagValue string `json:"tag_value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomForecastEntryTagFilter instantiates a new CustomForecastEntryTagFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomForecastEntryTagFilter(tagKey string, tagValue string) *CustomForecastEntryTagFilter {
	this := CustomForecastEntryTagFilter{}
	this.TagKey = tagKey
	this.TagValue = tagValue
	return &this
}

// NewCustomForecastEntryTagFilterWithDefaults instantiates a new CustomForecastEntryTagFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomForecastEntryTagFilterWithDefaults() *CustomForecastEntryTagFilter {
	this := CustomForecastEntryTagFilter{}
	return &this
}

// GetTagKey returns the TagKey field value.
func (o *CustomForecastEntryTagFilter) GetTagKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
func (o *CustomForecastEntryTagFilter) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKey, true
}

// SetTagKey sets field value.
func (o *CustomForecastEntryTagFilter) SetTagKey(v string) {
	o.TagKey = v
}

// GetTagValue returns the TagValue field value.
func (o *CustomForecastEntryTagFilter) GetTagValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagValue
}

// GetTagValueOk returns a tuple with the TagValue field value
// and a boolean to check if the value has been set.
func (o *CustomForecastEntryTagFilter) GetTagValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagValue, true
}

// SetTagValue sets field value.
func (o *CustomForecastEntryTagFilter) SetTagValue(v string) {
	o.TagValue = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomForecastEntryTagFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["tag_key"] = o.TagKey
	toSerialize["tag_value"] = o.TagValue

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomForecastEntryTagFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TagKey   *string `json:"tag_key"`
		TagValue *string `json:"tag_value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TagKey == nil {
		return fmt.Errorf("required field tag_key missing")
	}
	if all.TagValue == nil {
		return fmt.Errorf("required field tag_value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tag_key", "tag_value"})
	} else {
		return err
	}
	o.TagKey = *all.TagKey
	o.TagValue = *all.TagValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
