// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricBulkTagConfigCreateAttributes Optional parameters for bulk creating metric tag configurations.
type MetricBulkTagConfigCreateAttributes struct {
	// A list of account emails to notify when the configuration is applied.
	Emails []string `json:"emails,omitempty"`
	// When set to true, the configuration will exclude the configured tags and include any other submitted tags.
	// When set to false, the configuration will include the configured tags and exclude any other submitted tags.
	// Defaults to false.
	ExcludeTagsMode *bool `json:"exclude_tags_mode,omitempty"`
	// A list of tag names to apply to the configuration.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricBulkTagConfigCreateAttributes instantiates a new MetricBulkTagConfigCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricBulkTagConfigCreateAttributes() *MetricBulkTagConfigCreateAttributes {
	this := MetricBulkTagConfigCreateAttributes{}
	return &this
}

// NewMetricBulkTagConfigCreateAttributesWithDefaults instantiates a new MetricBulkTagConfigCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricBulkTagConfigCreateAttributesWithDefaults() *MetricBulkTagConfigCreateAttributes {
	this := MetricBulkTagConfigCreateAttributes{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *MetricBulkTagConfigCreateAttributes) GetEmails() []string {
	if o == nil || o.Emails == nil {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricBulkTagConfigCreateAttributes) GetEmailsOk() (*[]string, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return &o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *MetricBulkTagConfigCreateAttributes) HasEmails() bool {
	return o != nil && o.Emails != nil
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *MetricBulkTagConfigCreateAttributes) SetEmails(v []string) {
	o.Emails = v
}

// GetExcludeTagsMode returns the ExcludeTagsMode field value if set, zero value otherwise.
func (o *MetricBulkTagConfigCreateAttributes) GetExcludeTagsMode() bool {
	if o == nil || o.ExcludeTagsMode == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeTagsMode
}

// GetExcludeTagsModeOk returns a tuple with the ExcludeTagsMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricBulkTagConfigCreateAttributes) GetExcludeTagsModeOk() (*bool, bool) {
	if o == nil || o.ExcludeTagsMode == nil {
		return nil, false
	}
	return o.ExcludeTagsMode, true
}

// HasExcludeTagsMode returns a boolean if a field has been set.
func (o *MetricBulkTagConfigCreateAttributes) HasExcludeTagsMode() bool {
	return o != nil && o.ExcludeTagsMode != nil
}

// SetExcludeTagsMode gets a reference to the given bool and assigns it to the ExcludeTagsMode field.
func (o *MetricBulkTagConfigCreateAttributes) SetExcludeTagsMode(v bool) {
	o.ExcludeTagsMode = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MetricBulkTagConfigCreateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricBulkTagConfigCreateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MetricBulkTagConfigCreateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MetricBulkTagConfigCreateAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricBulkTagConfigCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.ExcludeTagsMode != nil {
		toSerialize["exclude_tags_mode"] = o.ExcludeTagsMode
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
func (o *MetricBulkTagConfigCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Emails          []string `json:"emails,omitempty"`
		ExcludeTagsMode *bool    `json:"exclude_tags_mode,omitempty"`
		Tags            []string `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"emails", "exclude_tags_mode", "tags"})
	} else {
		return err
	}
	o.Emails = all.Emails
	o.ExcludeTagsMode = all.ExcludeTagsMode
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
