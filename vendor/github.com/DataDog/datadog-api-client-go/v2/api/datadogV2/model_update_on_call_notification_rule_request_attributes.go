// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateOnCallNotificationRuleRequestAttributes Attributes for creating or modifying an on-call notification rule.
type UpdateOnCallNotificationRuleRequestAttributes struct {
	// Specifies the category a notification rule will apply to
	Category *OnCallNotificationRuleCategory `json:"category,omitempty"`
	// Defines the configuration for a channel associated with a notification rule
	ChannelSettings *OnCallNotificationRuleChannelSettings `json:"channel_settings,omitempty"`
	// The number of minutes that will elapse before this rule is evaluated.  0 indicates immediate evaluation
	DelayMinutes *int64 `json:"delay_minutes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateOnCallNotificationRuleRequestAttributes instantiates a new UpdateOnCallNotificationRuleRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateOnCallNotificationRuleRequestAttributes() *UpdateOnCallNotificationRuleRequestAttributes {
	this := UpdateOnCallNotificationRuleRequestAttributes{}
	var category OnCallNotificationRuleCategory = ONCALLNOTIFICATIONRULECATEGORY_HIGH_URGENCY
	this.Category = &category
	return &this
}

// NewUpdateOnCallNotificationRuleRequestAttributesWithDefaults instantiates a new UpdateOnCallNotificationRuleRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateOnCallNotificationRuleRequestAttributesWithDefaults() *UpdateOnCallNotificationRuleRequestAttributes {
	this := UpdateOnCallNotificationRuleRequestAttributes{}
	var category OnCallNotificationRuleCategory = ONCALLNOTIFICATIONRULECATEGORY_HIGH_URGENCY
	this.Category = &category
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *UpdateOnCallNotificationRuleRequestAttributes) GetCategory() OnCallNotificationRuleCategory {
	if o == nil || o.Category == nil {
		var ret OnCallNotificationRuleCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOnCallNotificationRuleRequestAttributes) GetCategoryOk() (*OnCallNotificationRuleCategory, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *UpdateOnCallNotificationRuleRequestAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given OnCallNotificationRuleCategory and assigns it to the Category field.
func (o *UpdateOnCallNotificationRuleRequestAttributes) SetCategory(v OnCallNotificationRuleCategory) {
	o.Category = &v
}

// GetChannelSettings returns the ChannelSettings field value if set, zero value otherwise.
func (o *UpdateOnCallNotificationRuleRequestAttributes) GetChannelSettings() OnCallNotificationRuleChannelSettings {
	if o == nil || o.ChannelSettings == nil {
		var ret OnCallNotificationRuleChannelSettings
		return ret
	}
	return *o.ChannelSettings
}

// GetChannelSettingsOk returns a tuple with the ChannelSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOnCallNotificationRuleRequestAttributes) GetChannelSettingsOk() (*OnCallNotificationRuleChannelSettings, bool) {
	if o == nil || o.ChannelSettings == nil {
		return nil, false
	}
	return o.ChannelSettings, true
}

// HasChannelSettings returns a boolean if a field has been set.
func (o *UpdateOnCallNotificationRuleRequestAttributes) HasChannelSettings() bool {
	return o != nil && o.ChannelSettings != nil
}

// SetChannelSettings gets a reference to the given OnCallNotificationRuleChannelSettings and assigns it to the ChannelSettings field.
func (o *UpdateOnCallNotificationRuleRequestAttributes) SetChannelSettings(v OnCallNotificationRuleChannelSettings) {
	o.ChannelSettings = &v
}

// GetDelayMinutes returns the DelayMinutes field value if set, zero value otherwise.
func (o *UpdateOnCallNotificationRuleRequestAttributes) GetDelayMinutes() int64 {
	if o == nil || o.DelayMinutes == nil {
		var ret int64
		return ret
	}
	return *o.DelayMinutes
}

// GetDelayMinutesOk returns a tuple with the DelayMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOnCallNotificationRuleRequestAttributes) GetDelayMinutesOk() (*int64, bool) {
	if o == nil || o.DelayMinutes == nil {
		return nil, false
	}
	return o.DelayMinutes, true
}

// HasDelayMinutes returns a boolean if a field has been set.
func (o *UpdateOnCallNotificationRuleRequestAttributes) HasDelayMinutes() bool {
	return o != nil && o.DelayMinutes != nil
}

// SetDelayMinutes gets a reference to the given int64 and assigns it to the DelayMinutes field.
func (o *UpdateOnCallNotificationRuleRequestAttributes) SetDelayMinutes(v int64) {
	o.DelayMinutes = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateOnCallNotificationRuleRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.ChannelSettings != nil {
		toSerialize["channel_settings"] = o.ChannelSettings
	}
	if o.DelayMinutes != nil {
		toSerialize["delay_minutes"] = o.DelayMinutes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateOnCallNotificationRuleRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category        *OnCallNotificationRuleCategory        `json:"category,omitempty"`
		ChannelSettings *OnCallNotificationRuleChannelSettings `json:"channel_settings,omitempty"`
		DelayMinutes    *int64                                 `json:"delay_minutes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "channel_settings", "delay_minutes"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Category != nil && !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = all.Category
	}
	o.ChannelSettings = all.ChannelSettings
	o.DelayMinutes = all.DelayMinutes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
