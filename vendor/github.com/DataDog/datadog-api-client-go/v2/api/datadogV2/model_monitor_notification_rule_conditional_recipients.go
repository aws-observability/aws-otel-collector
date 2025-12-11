// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleConditionalRecipients Use conditional recipients to define different recipients for different situations.
type MonitorNotificationRuleConditionalRecipients struct {
	// Conditions of the notification rule.
	Conditions []MonitorNotificationRuleCondition `json:"conditions"`
	// A list of recipients to notify. Uses the same format as the monitor `message` field. Must not start with an '@'.
	FallbackRecipients []string `json:"fallback_recipients,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorNotificationRuleConditionalRecipients instantiates a new MonitorNotificationRuleConditionalRecipients object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorNotificationRuleConditionalRecipients(conditions []MonitorNotificationRuleCondition) *MonitorNotificationRuleConditionalRecipients {
	this := MonitorNotificationRuleConditionalRecipients{}
	this.Conditions = conditions
	return &this
}

// NewMonitorNotificationRuleConditionalRecipientsWithDefaults instantiates a new MonitorNotificationRuleConditionalRecipients object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorNotificationRuleConditionalRecipientsWithDefaults() *MonitorNotificationRuleConditionalRecipients {
	this := MonitorNotificationRuleConditionalRecipients{}
	return &this
}

// GetConditions returns the Conditions field value.
func (o *MonitorNotificationRuleConditionalRecipients) GetConditions() []MonitorNotificationRuleCondition {
	if o == nil {
		var ret []MonitorNotificationRuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleConditionalRecipients) GetConditionsOk() (*[]MonitorNotificationRuleCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value.
func (o *MonitorNotificationRuleConditionalRecipients) SetConditions(v []MonitorNotificationRuleCondition) {
	o.Conditions = v
}

// GetFallbackRecipients returns the FallbackRecipients field value if set, zero value otherwise.
func (o *MonitorNotificationRuleConditionalRecipients) GetFallbackRecipients() []string {
	if o == nil || o.FallbackRecipients == nil {
		var ret []string
		return ret
	}
	return o.FallbackRecipients
}

// GetFallbackRecipientsOk returns a tuple with the FallbackRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleConditionalRecipients) GetFallbackRecipientsOk() (*[]string, bool) {
	if o == nil || o.FallbackRecipients == nil {
		return nil, false
	}
	return &o.FallbackRecipients, true
}

// HasFallbackRecipients returns a boolean if a field has been set.
func (o *MonitorNotificationRuleConditionalRecipients) HasFallbackRecipients() bool {
	return o != nil && o.FallbackRecipients != nil
}

// SetFallbackRecipients gets a reference to the given []string and assigns it to the FallbackRecipients field.
func (o *MonitorNotificationRuleConditionalRecipients) SetFallbackRecipients(v []string) {
	o.FallbackRecipients = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorNotificationRuleConditionalRecipients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["conditions"] = o.Conditions
	if o.FallbackRecipients != nil {
		toSerialize["fallback_recipients"] = o.FallbackRecipients
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorNotificationRuleConditionalRecipients) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Conditions         *[]MonitorNotificationRuleCondition `json:"conditions"`
		FallbackRecipients []string                            `json:"fallback_recipients,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Conditions == nil {
		return fmt.Errorf("required field conditions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditions", "fallback_recipients"})
	} else {
		return err
	}
	o.Conditions = *all.Conditions
	o.FallbackRecipients = all.FallbackRecipients

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
