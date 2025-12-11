// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationRuleAttributes The notification rule's attributes.
type IncidentNotificationRuleAttributes struct {
	// The conditions that trigger this notification rule.
	Conditions []IncidentNotificationRuleConditionsItems `json:"conditions"`
	// Timestamp when the notification rule was created.
	Created time.Time `json:"created"`
	// Whether the notification rule is enabled.
	Enabled bool `json:"enabled"`
	// The notification handles (targets) for this rule.
	Handles []string `json:"handles"`
	// Timestamp when the notification rule was last modified.
	Modified time.Time `json:"modified"`
	// List of incident fields that trigger re-notification when changed.
	RenotifyOn []string `json:"renotify_on,omitempty"`
	// The trigger event for this notification rule.
	Trigger string `json:"trigger"`
	// The visibility of the notification rule.
	Visibility IncidentNotificationRuleAttributesVisibility `json:"visibility"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationRuleAttributes instantiates a new IncidentNotificationRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationRuleAttributes(conditions []IncidentNotificationRuleConditionsItems, created time.Time, enabled bool, handles []string, modified time.Time, trigger string, visibility IncidentNotificationRuleAttributesVisibility) *IncidentNotificationRuleAttributes {
	this := IncidentNotificationRuleAttributes{}
	this.Conditions = conditions
	this.Created = created
	this.Enabled = enabled
	this.Handles = handles
	this.Modified = modified
	this.Trigger = trigger
	this.Visibility = visibility
	return &this
}

// NewIncidentNotificationRuleAttributesWithDefaults instantiates a new IncidentNotificationRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationRuleAttributesWithDefaults() *IncidentNotificationRuleAttributes {
	this := IncidentNotificationRuleAttributes{}
	return &this
}

// GetConditions returns the Conditions field value.
func (o *IncidentNotificationRuleAttributes) GetConditions() []IncidentNotificationRuleConditionsItems {
	if o == nil {
		var ret []IncidentNotificationRuleConditionsItems
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleAttributes) GetConditionsOk() (*[]IncidentNotificationRuleConditionsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value.
func (o *IncidentNotificationRuleAttributes) SetConditions(v []IncidentNotificationRuleConditionsItems) {
	o.Conditions = v
}

// GetCreated returns the Created field value.
func (o *IncidentNotificationRuleAttributes) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentNotificationRuleAttributes) SetCreated(v time.Time) {
	o.Created = v
}

// GetEnabled returns the Enabled field value.
func (o *IncidentNotificationRuleAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *IncidentNotificationRuleAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetHandles returns the Handles field value.
func (o *IncidentNotificationRuleAttributes) GetHandles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Handles
}

// GetHandlesOk returns a tuple with the Handles field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleAttributes) GetHandlesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handles, true
}

// SetHandles sets field value.
func (o *IncidentNotificationRuleAttributes) SetHandles(v []string) {
	o.Handles = v
}

// GetModified returns the Modified field value.
func (o *IncidentNotificationRuleAttributes) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentNotificationRuleAttributes) SetModified(v time.Time) {
	o.Modified = v
}

// GetRenotifyOn returns the RenotifyOn field value if set, zero value otherwise.
func (o *IncidentNotificationRuleAttributes) GetRenotifyOn() []string {
	if o == nil || o.RenotifyOn == nil {
		var ret []string
		return ret
	}
	return o.RenotifyOn
}

// GetRenotifyOnOk returns a tuple with the RenotifyOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleAttributes) GetRenotifyOnOk() (*[]string, bool) {
	if o == nil || o.RenotifyOn == nil {
		return nil, false
	}
	return &o.RenotifyOn, true
}

// HasRenotifyOn returns a boolean if a field has been set.
func (o *IncidentNotificationRuleAttributes) HasRenotifyOn() bool {
	return o != nil && o.RenotifyOn != nil
}

// SetRenotifyOn gets a reference to the given []string and assigns it to the RenotifyOn field.
func (o *IncidentNotificationRuleAttributes) SetRenotifyOn(v []string) {
	o.RenotifyOn = v
}

// GetTrigger returns the Trigger field value.
func (o *IncidentNotificationRuleAttributes) GetTrigger() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleAttributes) GetTriggerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value.
func (o *IncidentNotificationRuleAttributes) SetTrigger(v string) {
	o.Trigger = v
}

// GetVisibility returns the Visibility field value.
func (o *IncidentNotificationRuleAttributes) GetVisibility() IncidentNotificationRuleAttributesVisibility {
	if o == nil {
		var ret IncidentNotificationRuleAttributesVisibility
		return ret
	}
	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleAttributes) GetVisibilityOk() (*IncidentNotificationRuleAttributesVisibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value.
func (o *IncidentNotificationRuleAttributes) SetVisibility(v IncidentNotificationRuleAttributesVisibility) {
	o.Visibility = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["conditions"] = o.Conditions
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["handles"] = o.Handles
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.RenotifyOn != nil {
		toSerialize["renotify_on"] = o.RenotifyOn
	}
	toSerialize["trigger"] = o.Trigger
	toSerialize["visibility"] = o.Visibility

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentNotificationRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Conditions *[]IncidentNotificationRuleConditionsItems    `json:"conditions"`
		Created    *time.Time                                    `json:"created"`
		Enabled    *bool                                         `json:"enabled"`
		Handles    *[]string                                     `json:"handles"`
		Modified   *time.Time                                    `json:"modified"`
		RenotifyOn []string                                      `json:"renotify_on,omitempty"`
		Trigger    *string                                       `json:"trigger"`
		Visibility *IncidentNotificationRuleAttributesVisibility `json:"visibility"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Conditions == nil {
		return fmt.Errorf("required field conditions missing")
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Handles == nil {
		return fmt.Errorf("required field handles missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	if all.Trigger == nil {
		return fmt.Errorf("required field trigger missing")
	}
	if all.Visibility == nil {
		return fmt.Errorf("required field visibility missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditions", "created", "enabled", "handles", "modified", "renotify_on", "trigger", "visibility"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Conditions = *all.Conditions
	o.Created = *all.Created
	o.Enabled = *all.Enabled
	o.Handles = *all.Handles
	o.Modified = *all.Modified
	o.RenotifyOn = all.RenotifyOn
	o.Trigger = *all.Trigger
	if !all.Visibility.IsValid() {
		hasInvalidField = true
	} else {
		o.Visibility = *all.Visibility
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
