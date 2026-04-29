// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseNotificationRuleCreateAttributes Notification rule creation attributes
type CaseNotificationRuleCreateAttributes struct {
	// Whether the notification rule is enabled
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Query to filter cases for this notification rule
	Query *string `json:"query,omitempty"`
	// List of notification recipients
	Recipients []CaseNotificationRuleRecipient `json:"recipients"`
	// List of triggers for this notification rule
	Triggers []CaseNotificationRuleTrigger `json:"triggers"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseNotificationRuleCreateAttributes instantiates a new CaseNotificationRuleCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseNotificationRuleCreateAttributes(recipients []CaseNotificationRuleRecipient, triggers []CaseNotificationRuleTrigger) *CaseNotificationRuleCreateAttributes {
	this := CaseNotificationRuleCreateAttributes{}
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	this.Recipients = recipients
	this.Triggers = triggers
	return &this
}

// NewCaseNotificationRuleCreateAttributesWithDefaults instantiates a new CaseNotificationRuleCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseNotificationRuleCreateAttributesWithDefaults() *CaseNotificationRuleCreateAttributes {
	this := CaseNotificationRuleCreateAttributes{}
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *CaseNotificationRuleCreateAttributes) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleCreateAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *CaseNotificationRuleCreateAttributes) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *CaseNotificationRuleCreateAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *CaseNotificationRuleCreateAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleCreateAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *CaseNotificationRuleCreateAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *CaseNotificationRuleCreateAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetRecipients returns the Recipients field value.
func (o *CaseNotificationRuleCreateAttributes) GetRecipients() []CaseNotificationRuleRecipient {
	if o == nil {
		var ret []CaseNotificationRuleRecipient
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleCreateAttributes) GetRecipientsOk() (*[]CaseNotificationRuleRecipient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value.
func (o *CaseNotificationRuleCreateAttributes) SetRecipients(v []CaseNotificationRuleRecipient) {
	o.Recipients = v
}

// GetTriggers returns the Triggers field value.
func (o *CaseNotificationRuleCreateAttributes) GetTriggers() []CaseNotificationRuleTrigger {
	if o == nil {
		var ret []CaseNotificationRuleTrigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleCreateAttributes) GetTriggersOk() (*[]CaseNotificationRuleTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Triggers, true
}

// SetTriggers sets field value.
func (o *CaseNotificationRuleCreateAttributes) SetTriggers(v []CaseNotificationRuleTrigger) {
	o.Triggers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseNotificationRuleCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	toSerialize["recipients"] = o.Recipients
	toSerialize["triggers"] = o.Triggers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseNotificationRuleCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsEnabled  *bool                            `json:"is_enabled,omitempty"`
		Query      *string                          `json:"query,omitempty"`
		Recipients *[]CaseNotificationRuleRecipient `json:"recipients"`
		Triggers   *[]CaseNotificationRuleTrigger   `json:"triggers"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Recipients == nil {
		return fmt.Errorf("required field recipients missing")
	}
	if all.Triggers == nil {
		return fmt.Errorf("required field triggers missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_enabled", "query", "recipients", "triggers"})
	} else {
		return err
	}
	o.IsEnabled = all.IsEnabled
	o.Query = all.Query
	o.Recipients = *all.Recipients
	o.Triggers = *all.Triggers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
