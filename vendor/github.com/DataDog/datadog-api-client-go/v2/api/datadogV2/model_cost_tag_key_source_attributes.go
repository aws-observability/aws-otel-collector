// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagKeySourceAttributes Attributes of a Cloud Cost Management tag source.
type CostTagKeySourceAttributes struct {
	// The tag key name.
	TagKey string `json:"tag_key"`
	// Origins where this tag key was observed (for example, `aws-user-defined`).
	TagSources []string `json:"tag_sources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostTagKeySourceAttributes instantiates a new CostTagKeySourceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostTagKeySourceAttributes(tagKey string, tagSources []string) *CostTagKeySourceAttributes {
	this := CostTagKeySourceAttributes{}
	this.TagKey = tagKey
	this.TagSources = tagSources
	return &this
}

// NewCostTagKeySourceAttributesWithDefaults instantiates a new CostTagKeySourceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostTagKeySourceAttributesWithDefaults() *CostTagKeySourceAttributes {
	this := CostTagKeySourceAttributes{}
	return &this
}

// GetTagKey returns the TagKey field value.
func (o *CostTagKeySourceAttributes) GetTagKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
func (o *CostTagKeySourceAttributes) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKey, true
}

// SetTagKey sets field value.
func (o *CostTagKeySourceAttributes) SetTagKey(v string) {
	o.TagKey = v
}

// GetTagSources returns the TagSources field value.
func (o *CostTagKeySourceAttributes) GetTagSources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagSources
}

// GetTagSourcesOk returns a tuple with the TagSources field value
// and a boolean to check if the value has been set.
func (o *CostTagKeySourceAttributes) GetTagSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagSources, true
}

// SetTagSources sets field value.
func (o *CostTagKeySourceAttributes) SetTagSources(v []string) {
	o.TagSources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostTagKeySourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["tag_key"] = o.TagKey
	toSerialize["tag_sources"] = o.TagSources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostTagKeySourceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TagKey     *string   `json:"tag_key"`
		TagSources *[]string `json:"tag_sources"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TagKey == nil {
		return fmt.Errorf("required field tag_key missing")
	}
	if all.TagSources == nil {
		return fmt.Errorf("required field tag_sources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tag_key", "tag_sources"})
	} else {
		return err
	}
	o.TagKey = *all.TagKey
	o.TagSources = *all.TagSources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
