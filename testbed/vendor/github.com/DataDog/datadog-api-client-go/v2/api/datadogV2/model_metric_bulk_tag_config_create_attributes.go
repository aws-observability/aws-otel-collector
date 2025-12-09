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
	// When provided, all tags that have been actively queried are
	// configured (and, therefore, remain queryable) for each metric that
	// matches the given prefix.  Minimum value is 1 second, and maximum
	// value is 7,776,000 seconds (90 days).
	IncludeActivelyQueriedTagsWindow *float64 `json:"include_actively_queried_tags_window,omitempty"`
	// When set to true, the configuration overrides any existing
	// configurations for the given metric with the new set of tags in this
	// configuration request. If false, old configurations are kept and
	// are merged with the set of tags in this configuration request.
	// Defaults to true.
	OverrideExistingConfigurations *bool `json:"override_existing_configurations,omitempty"`
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

// GetIncludeActivelyQueriedTagsWindow returns the IncludeActivelyQueriedTagsWindow field value if set, zero value otherwise.
func (o *MetricBulkTagConfigCreateAttributes) GetIncludeActivelyQueriedTagsWindow() float64 {
	if o == nil || o.IncludeActivelyQueriedTagsWindow == nil {
		var ret float64
		return ret
	}
	return *o.IncludeActivelyQueriedTagsWindow
}

// GetIncludeActivelyQueriedTagsWindowOk returns a tuple with the IncludeActivelyQueriedTagsWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricBulkTagConfigCreateAttributes) GetIncludeActivelyQueriedTagsWindowOk() (*float64, bool) {
	if o == nil || o.IncludeActivelyQueriedTagsWindow == nil {
		return nil, false
	}
	return o.IncludeActivelyQueriedTagsWindow, true
}

// HasIncludeActivelyQueriedTagsWindow returns a boolean if a field has been set.
func (o *MetricBulkTagConfigCreateAttributes) HasIncludeActivelyQueriedTagsWindow() bool {
	return o != nil && o.IncludeActivelyQueriedTagsWindow != nil
}

// SetIncludeActivelyQueriedTagsWindow gets a reference to the given float64 and assigns it to the IncludeActivelyQueriedTagsWindow field.
func (o *MetricBulkTagConfigCreateAttributes) SetIncludeActivelyQueriedTagsWindow(v float64) {
	o.IncludeActivelyQueriedTagsWindow = &v
}

// GetOverrideExistingConfigurations returns the OverrideExistingConfigurations field value if set, zero value otherwise.
func (o *MetricBulkTagConfigCreateAttributes) GetOverrideExistingConfigurations() bool {
	if o == nil || o.OverrideExistingConfigurations == nil {
		var ret bool
		return ret
	}
	return *o.OverrideExistingConfigurations
}

// GetOverrideExistingConfigurationsOk returns a tuple with the OverrideExistingConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricBulkTagConfigCreateAttributes) GetOverrideExistingConfigurationsOk() (*bool, bool) {
	if o == nil || o.OverrideExistingConfigurations == nil {
		return nil, false
	}
	return o.OverrideExistingConfigurations, true
}

// HasOverrideExistingConfigurations returns a boolean if a field has been set.
func (o *MetricBulkTagConfigCreateAttributes) HasOverrideExistingConfigurations() bool {
	return o != nil && o.OverrideExistingConfigurations != nil
}

// SetOverrideExistingConfigurations gets a reference to the given bool and assigns it to the OverrideExistingConfigurations field.
func (o *MetricBulkTagConfigCreateAttributes) SetOverrideExistingConfigurations(v bool) {
	o.OverrideExistingConfigurations = &v
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
	if o.IncludeActivelyQueriedTagsWindow != nil {
		toSerialize["include_actively_queried_tags_window"] = o.IncludeActivelyQueriedTagsWindow
	}
	if o.OverrideExistingConfigurations != nil {
		toSerialize["override_existing_configurations"] = o.OverrideExistingConfigurations
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
		Emails                           []string `json:"emails,omitempty"`
		ExcludeTagsMode                  *bool    `json:"exclude_tags_mode,omitempty"`
		IncludeActivelyQueriedTagsWindow *float64 `json:"include_actively_queried_tags_window,omitempty"`
		OverrideExistingConfigurations   *bool    `json:"override_existing_configurations,omitempty"`
		Tags                             []string `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"emails", "exclude_tags_mode", "include_actively_queried_tags_window", "override_existing_configurations", "tags"})
	} else {
		return err
	}
	o.Emails = all.Emails
	o.ExcludeTagsMode = all.ExcludeTagsMode
	o.IncludeActivelyQueriedTagsWindow = all.IncludeActivelyQueriedTagsWindow
	o.OverrideExistingConfigurations = all.OverrideExistingConfigurations
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
