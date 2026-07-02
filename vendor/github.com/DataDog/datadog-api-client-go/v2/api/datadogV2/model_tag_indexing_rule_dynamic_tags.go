// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleDynamicTags Configuration for including dynamically queried tags.
type TagIndexingRuleDynamicTags struct {
	// Window in seconds for evaluating queried tags.
	QueriedTagsWindowSeconds *int64 `json:"queried_tags_window_seconds,omitempty"`
	// When true, tags from related assets are included.
	RelatedAssetTags *bool `json:"related_asset_tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleDynamicTags instantiates a new TagIndexingRuleDynamicTags object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleDynamicTags() *TagIndexingRuleDynamicTags {
	this := TagIndexingRuleDynamicTags{}
	return &this
}

// NewTagIndexingRuleDynamicTagsWithDefaults instantiates a new TagIndexingRuleDynamicTags object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleDynamicTagsWithDefaults() *TagIndexingRuleDynamicTags {
	this := TagIndexingRuleDynamicTags{}
	return &this
}

// GetQueriedTagsWindowSeconds returns the QueriedTagsWindowSeconds field value if set, zero value otherwise.
func (o *TagIndexingRuleDynamicTags) GetQueriedTagsWindowSeconds() int64 {
	if o == nil || o.QueriedTagsWindowSeconds == nil {
		var ret int64
		return ret
	}
	return *o.QueriedTagsWindowSeconds
}

// GetQueriedTagsWindowSecondsOk returns a tuple with the QueriedTagsWindowSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleDynamicTags) GetQueriedTagsWindowSecondsOk() (*int64, bool) {
	if o == nil || o.QueriedTagsWindowSeconds == nil {
		return nil, false
	}
	return o.QueriedTagsWindowSeconds, true
}

// HasQueriedTagsWindowSeconds returns a boolean if a field has been set.
func (o *TagIndexingRuleDynamicTags) HasQueriedTagsWindowSeconds() bool {
	return o != nil && o.QueriedTagsWindowSeconds != nil
}

// SetQueriedTagsWindowSeconds gets a reference to the given int64 and assigns it to the QueriedTagsWindowSeconds field.
func (o *TagIndexingRuleDynamicTags) SetQueriedTagsWindowSeconds(v int64) {
	o.QueriedTagsWindowSeconds = &v
}

// GetRelatedAssetTags returns the RelatedAssetTags field value if set, zero value otherwise.
func (o *TagIndexingRuleDynamicTags) GetRelatedAssetTags() bool {
	if o == nil || o.RelatedAssetTags == nil {
		var ret bool
		return ret
	}
	return *o.RelatedAssetTags
}

// GetRelatedAssetTagsOk returns a tuple with the RelatedAssetTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleDynamicTags) GetRelatedAssetTagsOk() (*bool, bool) {
	if o == nil || o.RelatedAssetTags == nil {
		return nil, false
	}
	return o.RelatedAssetTags, true
}

// HasRelatedAssetTags returns a boolean if a field has been set.
func (o *TagIndexingRuleDynamicTags) HasRelatedAssetTags() bool {
	return o != nil && o.RelatedAssetTags != nil
}

// SetRelatedAssetTags gets a reference to the given bool and assigns it to the RelatedAssetTags field.
func (o *TagIndexingRuleDynamicTags) SetRelatedAssetTags(v bool) {
	o.RelatedAssetTags = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleDynamicTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.QueriedTagsWindowSeconds != nil {
		toSerialize["queried_tags_window_seconds"] = o.QueriedTagsWindowSeconds
	}
	if o.RelatedAssetTags != nil {
		toSerialize["related_asset_tags"] = o.RelatedAssetTags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagIndexingRuleDynamicTags) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueriedTagsWindowSeconds *int64 `json:"queried_tags_window_seconds,omitempty"`
		RelatedAssetTags         *bool  `json:"related_asset_tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"queried_tags_window_seconds", "related_asset_tags"})
	} else {
		return err
	}
	o.QueriedTagsWindowSeconds = all.QueriedTagsWindowSeconds
	o.RelatedAssetTags = all.RelatedAssetTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
