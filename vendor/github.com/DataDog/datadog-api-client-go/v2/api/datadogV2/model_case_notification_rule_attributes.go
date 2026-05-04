// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseNotificationRuleAttributes Notification rule attributes
type CaseNotificationRuleAttributes struct {
	// Whether the notification rule is enabled
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Query to filter cases for this notification rule
	Query *string `json:"query,omitempty"`
	// List of notification recipients
	Recipients []CaseNotificationRuleRecipient `json:"recipients,omitempty"`
	// List of triggers for this notification rule
	Triggers []CaseNotificationRuleTrigger `json:"triggers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseNotificationRuleAttributes instantiates a new CaseNotificationRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseNotificationRuleAttributes() *CaseNotificationRuleAttributes {
	this := CaseNotificationRuleAttributes{}
	return &this
}

// NewCaseNotificationRuleAttributesWithDefaults instantiates a new CaseNotificationRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseNotificationRuleAttributesWithDefaults() *CaseNotificationRuleAttributes {
	this := CaseNotificationRuleAttributes{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *CaseNotificationRuleAttributes) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *CaseNotificationRuleAttributes) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *CaseNotificationRuleAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *CaseNotificationRuleAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *CaseNotificationRuleAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *CaseNotificationRuleAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *CaseNotificationRuleAttributes) GetRecipients() []CaseNotificationRuleRecipient {
	if o == nil || o.Recipients == nil {
		var ret []CaseNotificationRuleRecipient
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleAttributes) GetRecipientsOk() (*[]CaseNotificationRuleRecipient, bool) {
	if o == nil || o.Recipients == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *CaseNotificationRuleAttributes) HasRecipients() bool {
	return o != nil && o.Recipients != nil
}

// SetRecipients gets a reference to the given []CaseNotificationRuleRecipient and assigns it to the Recipients field.
func (o *CaseNotificationRuleAttributes) SetRecipients(v []CaseNotificationRuleRecipient) {
	o.Recipients = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *CaseNotificationRuleAttributes) GetTriggers() []CaseNotificationRuleTrigger {
	if o == nil || o.Triggers == nil {
		var ret []CaseNotificationRuleTrigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleAttributes) GetTriggersOk() (*[]CaseNotificationRuleTrigger, bool) {
	if o == nil || o.Triggers == nil {
		return nil, false
	}
	return &o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *CaseNotificationRuleAttributes) HasTriggers() bool {
	return o != nil && o.Triggers != nil
}

// SetTriggers gets a reference to the given []CaseNotificationRuleTrigger and assigns it to the Triggers field.
func (o *CaseNotificationRuleAttributes) SetTriggers(v []CaseNotificationRuleTrigger) {
	o.Triggers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseNotificationRuleAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Recipients != nil {
		toSerialize["recipients"] = o.Recipients
	}
	if o.Triggers != nil {
		toSerialize["triggers"] = o.Triggers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseNotificationRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsEnabled  *bool                           `json:"is_enabled,omitempty"`
		Query      *string                         `json:"query,omitempty"`
		Recipients []CaseNotificationRuleRecipient `json:"recipients,omitempty"`
		Triggers   []CaseNotificationRuleTrigger   `json:"triggers,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_enabled", "query", "recipients", "triggers"})
	} else {
		return err
	}
	o.IsEnabled = all.IsEnabled
	o.Query = all.Query
	o.Recipients = all.Recipients
	o.Triggers = all.Triggers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
