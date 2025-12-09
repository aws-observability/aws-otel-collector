// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleCondition Conditions for `conditional_recipients`.
type MonitorNotificationRuleCondition struct {
	// A list of recipients to notify. Uses the same format as the monitor `message` field. Must not start with an '@'.
	Recipients []string `json:"recipients"`
	// The scope to which the monitor applied.
	Scope string `json:"scope"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorNotificationRuleCondition instantiates a new MonitorNotificationRuleCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorNotificationRuleCondition(recipients []string, scope string) *MonitorNotificationRuleCondition {
	this := MonitorNotificationRuleCondition{}
	this.Recipients = recipients
	this.Scope = scope
	return &this
}

// NewMonitorNotificationRuleConditionWithDefaults instantiates a new MonitorNotificationRuleCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorNotificationRuleConditionWithDefaults() *MonitorNotificationRuleCondition {
	this := MonitorNotificationRuleCondition{}
	return &this
}

// GetRecipients returns the Recipients field value.
func (o *MonitorNotificationRuleCondition) GetRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleCondition) GetRecipientsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value.
func (o *MonitorNotificationRuleCondition) SetRecipients(v []string) {
	o.Recipients = v
}

// GetScope returns the Scope field value.
func (o *MonitorNotificationRuleCondition) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleCondition) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *MonitorNotificationRuleCondition) SetScope(v string) {
	o.Scope = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorNotificationRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["recipients"] = o.Recipients
	toSerialize["scope"] = o.Scope

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorNotificationRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Recipients *[]string `json:"recipients"`
		Scope      *string   `json:"scope"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Recipients == nil {
		return fmt.Errorf("required field recipients missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"recipients", "scope"})
	} else {
		return err
	}
	o.Recipients = *all.Recipients
	o.Scope = *all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
