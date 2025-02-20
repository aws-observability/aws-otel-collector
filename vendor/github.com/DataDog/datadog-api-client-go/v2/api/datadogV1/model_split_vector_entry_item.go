// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitVectorEntryItem The split graph list contains a graph for each value of the split dimension.
type SplitVectorEntryItem struct {
	// The tag key.
	TagKey string `json:"tag_key"`
	// The tag values.
	TagValues []string `json:"tag_values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSplitVectorEntryItem instantiates a new SplitVectorEntryItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitVectorEntryItem(tagKey string, tagValues []string) *SplitVectorEntryItem {
	this := SplitVectorEntryItem{}
	this.TagKey = tagKey
	this.TagValues = tagValues
	return &this
}

// NewSplitVectorEntryItemWithDefaults instantiates a new SplitVectorEntryItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitVectorEntryItemWithDefaults() *SplitVectorEntryItem {
	this := SplitVectorEntryItem{}
	return &this
}

// GetTagKey returns the TagKey field value.
func (o *SplitVectorEntryItem) GetTagKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
func (o *SplitVectorEntryItem) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKey, true
}

// SetTagKey sets field value.
func (o *SplitVectorEntryItem) SetTagKey(v string) {
	o.TagKey = v
}

// GetTagValues returns the TagValues field value.
func (o *SplitVectorEntryItem) GetTagValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value
// and a boolean to check if the value has been set.
func (o *SplitVectorEntryItem) GetTagValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagValues, true
}

// SetTagValues sets field value.
func (o *SplitVectorEntryItem) SetTagValues(v []string) {
	o.TagValues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitVectorEntryItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["tag_key"] = o.TagKey
	toSerialize["tag_values"] = o.TagValues

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplitVectorEntryItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TagKey    *string   `json:"tag_key"`
		TagValues *[]string `json:"tag_values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TagKey == nil {
		return fmt.Errorf("required field tag_key missing")
	}
	if all.TagValues == nil {
		return fmt.Errorf("required field tag_values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tag_key", "tag_values"})
	} else {
		return err
	}
	o.TagKey = *all.TagKey
	o.TagValues = *all.TagValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
