// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatuspageUrlSettingUpdateAttributes The Statuspage URL setting attributes for an update request.
type StatuspageUrlSettingUpdateAttributes struct {
	// Comma-separated list of custom tags to apply to events generated from this Statuspage URL.
	CustomTags *string `json:"custom_tags,omitempty"`
	// The Statuspage URL to monitor.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatuspageUrlSettingUpdateAttributes instantiates a new StatuspageUrlSettingUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatuspageUrlSettingUpdateAttributes() *StatuspageUrlSettingUpdateAttributes {
	this := StatuspageUrlSettingUpdateAttributes{}
	return &this
}

// NewStatuspageUrlSettingUpdateAttributesWithDefaults instantiates a new StatuspageUrlSettingUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatuspageUrlSettingUpdateAttributesWithDefaults() *StatuspageUrlSettingUpdateAttributes {
	this := StatuspageUrlSettingUpdateAttributes{}
	return &this
}

// GetCustomTags returns the CustomTags field value if set, zero value otherwise.
func (o *StatuspageUrlSettingUpdateAttributes) GetCustomTags() string {
	if o == nil || o.CustomTags == nil {
		var ret string
		return ret
	}
	return *o.CustomTags
}

// GetCustomTagsOk returns a tuple with the CustomTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatuspageUrlSettingUpdateAttributes) GetCustomTagsOk() (*string, bool) {
	if o == nil || o.CustomTags == nil {
		return nil, false
	}
	return o.CustomTags, true
}

// HasCustomTags returns a boolean if a field has been set.
func (o *StatuspageUrlSettingUpdateAttributes) HasCustomTags() bool {
	return o != nil && o.CustomTags != nil
}

// SetCustomTags gets a reference to the given string and assigns it to the CustomTags field.
func (o *StatuspageUrlSettingUpdateAttributes) SetCustomTags(v string) {
	o.CustomTags = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *StatuspageUrlSettingUpdateAttributes) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatuspageUrlSettingUpdateAttributes) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *StatuspageUrlSettingUpdateAttributes) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *StatuspageUrlSettingUpdateAttributes) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatuspageUrlSettingUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomTags != nil {
		toSerialize["custom_tags"] = o.CustomTags
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatuspageUrlSettingUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomTags *string `json:"custom_tags,omitempty"`
		Url        *string `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_tags", "url"})
	} else {
		return err
	}
	o.CustomTags = all.CustomTags
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
