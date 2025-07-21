// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleAttributes Attributes of the monitor notification rule.
type MonitorNotificationRuleAttributes struct {
	// Filter used to associate the notification rule with monitors.
	Filter *MonitorNotificationRuleFilter `json:"filter,omitempty"`
	// The name of the monitor notification rule.
	Name string `json:"name"`
	// A list of recipients to notify. Uses the same format as the monitor `message` field. Must not start with an '@'.
	Recipients []string `json:"recipients"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorNotificationRuleAttributes instantiates a new MonitorNotificationRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorNotificationRuleAttributes(name string, recipients []string) *MonitorNotificationRuleAttributes {
	this := MonitorNotificationRuleAttributes{}
	this.Name = name
	this.Recipients = recipients
	return &this
}

// NewMonitorNotificationRuleAttributesWithDefaults instantiates a new MonitorNotificationRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorNotificationRuleAttributesWithDefaults() *MonitorNotificationRuleAttributes {
	this := MonitorNotificationRuleAttributes{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MonitorNotificationRuleAttributes) GetFilter() MonitorNotificationRuleFilter {
	if o == nil || o.Filter == nil {
		var ret MonitorNotificationRuleFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleAttributes) GetFilterOk() (*MonitorNotificationRuleFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MonitorNotificationRuleAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given MonitorNotificationRuleFilter and assigns it to the Filter field.
func (o *MonitorNotificationRuleAttributes) SetFilter(v MonitorNotificationRuleFilter) {
	o.Filter = &v
}

// GetName returns the Name field value.
func (o *MonitorNotificationRuleAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MonitorNotificationRuleAttributes) SetName(v string) {
	o.Name = v
}

// GetRecipients returns the Recipients field value.
func (o *MonitorNotificationRuleAttributes) GetRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleAttributes) GetRecipientsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value.
func (o *MonitorNotificationRuleAttributes) SetRecipients(v []string) {
	o.Recipients = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorNotificationRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["name"] = o.Name
	toSerialize["recipients"] = o.Recipients
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorNotificationRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter     *MonitorNotificationRuleFilter `json:"filter,omitempty"`
		Name       *string                        `json:"name"`
		Recipients *[]string                      `json:"recipients"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Recipients == nil {
		return fmt.Errorf("required field recipients missing")
	}
	o.Filter = all.Filter
	o.Name = *all.Name
	o.Recipients = *all.Recipients

	return nil
}
