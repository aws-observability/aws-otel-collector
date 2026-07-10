// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatuspageUrlSettingCreateAttributes The Statuspage URL setting attributes for a create request.
type StatuspageUrlSettingCreateAttributes struct {
	// Comma-separated list of custom tags to apply to events generated from this Statuspage URL.
	CustomTags string `json:"custom_tags"`
	// The Statuspage URL to monitor. Must be a `status.io` or `statuspage.com` URL.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatuspageUrlSettingCreateAttributes instantiates a new StatuspageUrlSettingCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatuspageUrlSettingCreateAttributes(customTags string, url string) *StatuspageUrlSettingCreateAttributes {
	this := StatuspageUrlSettingCreateAttributes{}
	this.CustomTags = customTags
	this.Url = url
	return &this
}

// NewStatuspageUrlSettingCreateAttributesWithDefaults instantiates a new StatuspageUrlSettingCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatuspageUrlSettingCreateAttributesWithDefaults() *StatuspageUrlSettingCreateAttributes {
	this := StatuspageUrlSettingCreateAttributes{}
	return &this
}

// GetCustomTags returns the CustomTags field value.
func (o *StatuspageUrlSettingCreateAttributes) GetCustomTags() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CustomTags
}

// GetCustomTagsOk returns a tuple with the CustomTags field value
// and a boolean to check if the value has been set.
func (o *StatuspageUrlSettingCreateAttributes) GetCustomTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomTags, true
}

// SetCustomTags sets field value.
func (o *StatuspageUrlSettingCreateAttributes) SetCustomTags(v string) {
	o.CustomTags = v
}

// GetUrl returns the Url field value.
func (o *StatuspageUrlSettingCreateAttributes) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *StatuspageUrlSettingCreateAttributes) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *StatuspageUrlSettingCreateAttributes) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatuspageUrlSettingCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["custom_tags"] = o.CustomTags
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatuspageUrlSettingCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomTags *string `json:"custom_tags"`
		Url        *string `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CustomTags == nil {
		return fmt.Errorf("required field custom_tags missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_tags", "url"})
	} else {
		return err
	}
	o.CustomTags = *all.CustomTags
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
