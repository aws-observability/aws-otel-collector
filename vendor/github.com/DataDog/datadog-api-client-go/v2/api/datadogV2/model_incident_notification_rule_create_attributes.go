// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationRuleCreateAttributes The attributes for creating a notification rule.
type IncidentNotificationRuleCreateAttributes struct {
	// The conditions that trigger this notification rule.
	Conditions []IncidentNotificationRuleConditionsItems `json:"conditions"`
	// Whether the notification rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The notification handles (targets) for this rule.
	Handles []string `json:"handles"`
	// List of incident fields that trigger re-notification when changed.
	RenotifyOn []string `json:"renotify_on,omitempty"`
	// The trigger event for this notification rule.
	Trigger string `json:"trigger"`
	// The visibility of the notification rule.
	Visibility *IncidentNotificationRuleCreateAttributesVisibility `json:"visibility,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationRuleCreateAttributes instantiates a new IncidentNotificationRuleCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationRuleCreateAttributes(conditions []IncidentNotificationRuleConditionsItems, handles []string, trigger string) *IncidentNotificationRuleCreateAttributes {
	this := IncidentNotificationRuleCreateAttributes{}
	this.Conditions = conditions
	var enabled bool = false
	this.Enabled = &enabled
	this.Handles = handles
	this.Trigger = trigger
	return &this
}

// NewIncidentNotificationRuleCreateAttributesWithDefaults instantiates a new IncidentNotificationRuleCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationRuleCreateAttributesWithDefaults() *IncidentNotificationRuleCreateAttributes {
	this := IncidentNotificationRuleCreateAttributes{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetConditions returns the Conditions field value.
func (o *IncidentNotificationRuleCreateAttributes) GetConditions() []IncidentNotificationRuleConditionsItems {
	if o == nil {
		var ret []IncidentNotificationRuleConditionsItems
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleCreateAttributes) GetConditionsOk() (*[]IncidentNotificationRuleConditionsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value.
func (o *IncidentNotificationRuleCreateAttributes) SetConditions(v []IncidentNotificationRuleConditionsItems) {
	o.Conditions = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IncidentNotificationRuleCreateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IncidentNotificationRuleCreateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IncidentNotificationRuleCreateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHandles returns the Handles field value.
func (o *IncidentNotificationRuleCreateAttributes) GetHandles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Handles
}

// GetHandlesOk returns a tuple with the Handles field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleCreateAttributes) GetHandlesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handles, true
}

// SetHandles sets field value.
func (o *IncidentNotificationRuleCreateAttributes) SetHandles(v []string) {
	o.Handles = v
}

// GetRenotifyOn returns the RenotifyOn field value if set, zero value otherwise.
func (o *IncidentNotificationRuleCreateAttributes) GetRenotifyOn() []string {
	if o == nil || o.RenotifyOn == nil {
		var ret []string
		return ret
	}
	return o.RenotifyOn
}

// GetRenotifyOnOk returns a tuple with the RenotifyOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleCreateAttributes) GetRenotifyOnOk() (*[]string, bool) {
	if o == nil || o.RenotifyOn == nil {
		return nil, false
	}
	return &o.RenotifyOn, true
}

// HasRenotifyOn returns a boolean if a field has been set.
func (o *IncidentNotificationRuleCreateAttributes) HasRenotifyOn() bool {
	return o != nil && o.RenotifyOn != nil
}

// SetRenotifyOn gets a reference to the given []string and assigns it to the RenotifyOn field.
func (o *IncidentNotificationRuleCreateAttributes) SetRenotifyOn(v []string) {
	o.RenotifyOn = v
}

// GetTrigger returns the Trigger field value.
func (o *IncidentNotificationRuleCreateAttributes) GetTrigger() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleCreateAttributes) GetTriggerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value.
func (o *IncidentNotificationRuleCreateAttributes) SetTrigger(v string) {
	o.Trigger = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *IncidentNotificationRuleCreateAttributes) GetVisibility() IncidentNotificationRuleCreateAttributesVisibility {
	if o == nil || o.Visibility == nil {
		var ret IncidentNotificationRuleCreateAttributesVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleCreateAttributes) GetVisibilityOk() (*IncidentNotificationRuleCreateAttributesVisibility, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *IncidentNotificationRuleCreateAttributes) HasVisibility() bool {
	return o != nil && o.Visibility != nil
}

// SetVisibility gets a reference to the given IncidentNotificationRuleCreateAttributesVisibility and assigns it to the Visibility field.
func (o *IncidentNotificationRuleCreateAttributes) SetVisibility(v IncidentNotificationRuleCreateAttributesVisibility) {
	o.Visibility = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationRuleCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["conditions"] = o.Conditions
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["handles"] = o.Handles
	if o.RenotifyOn != nil {
		toSerialize["renotify_on"] = o.RenotifyOn
	}
	toSerialize["trigger"] = o.Trigger
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentNotificationRuleCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Conditions *[]IncidentNotificationRuleConditionsItems          `json:"conditions"`
		Enabled    *bool                                               `json:"enabled,omitempty"`
		Handles    *[]string                                           `json:"handles"`
		RenotifyOn []string                                            `json:"renotify_on,omitempty"`
		Trigger    *string                                             `json:"trigger"`
		Visibility *IncidentNotificationRuleCreateAttributesVisibility `json:"visibility,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Conditions == nil {
		return fmt.Errorf("required field conditions missing")
	}
	if all.Handles == nil {
		return fmt.Errorf("required field handles missing")
	}
	if all.Trigger == nil {
		return fmt.Errorf("required field trigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditions", "enabled", "handles", "renotify_on", "trigger", "visibility"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Conditions = *all.Conditions
	o.Enabled = all.Enabled
	o.Handles = *all.Handles
	o.RenotifyOn = all.RenotifyOn
	o.Trigger = *all.Trigger
	if all.Visibility != nil && !all.Visibility.IsValid() {
		hasInvalidField = true
	} else {
		o.Visibility = all.Visibility
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
