// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleResponseAttributes Attributes of the monitor notification rule.
type MonitorNotificationRuleResponseAttributes struct {
	// Use conditional recipients to define different recipients for different situations.
	ConditionalRecipients *MonitorNotificationRuleConditionalRecipients `json:"conditional_recipients,omitempty"`
	// Creation time of the monitor notification rule.
	Created *time.Time `json:"created,omitempty"`
	// Filter used to associate the notification rule with monitors.
	Filter *MonitorNotificationRuleFilter `json:"filter,omitempty"`
	// Time the monitor notification rule was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// The name of the monitor notification rule.
	Name *string `json:"name,omitempty"`
	// A list of recipients to notify. Uses the same format as the monitor `message` field. Must not start with an '@'.
	Recipients []string `json:"recipients,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorNotificationRuleResponseAttributes instantiates a new MonitorNotificationRuleResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorNotificationRuleResponseAttributes() *MonitorNotificationRuleResponseAttributes {
	this := MonitorNotificationRuleResponseAttributes{}
	return &this
}

// NewMonitorNotificationRuleResponseAttributesWithDefaults instantiates a new MonitorNotificationRuleResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorNotificationRuleResponseAttributesWithDefaults() *MonitorNotificationRuleResponseAttributes {
	this := MonitorNotificationRuleResponseAttributes{}
	return &this
}

// GetConditionalRecipients returns the ConditionalRecipients field value if set, zero value otherwise.
func (o *MonitorNotificationRuleResponseAttributes) GetConditionalRecipients() MonitorNotificationRuleConditionalRecipients {
	if o == nil || o.ConditionalRecipients == nil {
		var ret MonitorNotificationRuleConditionalRecipients
		return ret
	}
	return *o.ConditionalRecipients
}

// GetConditionalRecipientsOk returns a tuple with the ConditionalRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleResponseAttributes) GetConditionalRecipientsOk() (*MonitorNotificationRuleConditionalRecipients, bool) {
	if o == nil || o.ConditionalRecipients == nil {
		return nil, false
	}
	return o.ConditionalRecipients, true
}

// HasConditionalRecipients returns a boolean if a field has been set.
func (o *MonitorNotificationRuleResponseAttributes) HasConditionalRecipients() bool {
	return o != nil && o.ConditionalRecipients != nil
}

// SetConditionalRecipients gets a reference to the given MonitorNotificationRuleConditionalRecipients and assigns it to the ConditionalRecipients field.
func (o *MonitorNotificationRuleResponseAttributes) SetConditionalRecipients(v MonitorNotificationRuleConditionalRecipients) {
	o.ConditionalRecipients = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *MonitorNotificationRuleResponseAttributes) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleResponseAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *MonitorNotificationRuleResponseAttributes) HasCreated() bool {
	return o != nil && o.Created != nil
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *MonitorNotificationRuleResponseAttributes) SetCreated(v time.Time) {
	o.Created = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MonitorNotificationRuleResponseAttributes) GetFilter() MonitorNotificationRuleFilter {
	if o == nil || o.Filter == nil {
		var ret MonitorNotificationRuleFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleResponseAttributes) GetFilterOk() (*MonitorNotificationRuleFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MonitorNotificationRuleResponseAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given MonitorNotificationRuleFilter and assigns it to the Filter field.
func (o *MonitorNotificationRuleResponseAttributes) SetFilter(v MonitorNotificationRuleFilter) {
	o.Filter = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *MonitorNotificationRuleResponseAttributes) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleResponseAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *MonitorNotificationRuleResponseAttributes) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *MonitorNotificationRuleResponseAttributes) SetModified(v time.Time) {
	o.Modified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MonitorNotificationRuleResponseAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MonitorNotificationRuleResponseAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MonitorNotificationRuleResponseAttributes) SetName(v string) {
	o.Name = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *MonitorNotificationRuleResponseAttributes) GetRecipients() []string {
	if o == nil || o.Recipients == nil {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleResponseAttributes) GetRecipientsOk() (*[]string, bool) {
	if o == nil || o.Recipients == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *MonitorNotificationRuleResponseAttributes) HasRecipients() bool {
	return o != nil && o.Recipients != nil
}

// SetRecipients gets a reference to the given []string and assigns it to the Recipients field.
func (o *MonitorNotificationRuleResponseAttributes) SetRecipients(v []string) {
	o.Recipients = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorNotificationRuleResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConditionalRecipients != nil {
		toSerialize["conditional_recipients"] = o.ConditionalRecipients
	}
	if o.Created != nil {
		if o.Created.Nanosecond() == 0 {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Modified != nil {
		if o.Modified.Nanosecond() == 0 {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Recipients != nil {
		toSerialize["recipients"] = o.Recipients
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorNotificationRuleResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalRecipients *MonitorNotificationRuleConditionalRecipients `json:"conditional_recipients,omitempty"`
		Created               *time.Time                                    `json:"created,omitempty"`
		Filter                *MonitorNotificationRuleFilter                `json:"filter,omitempty"`
		Modified              *time.Time                                    `json:"modified,omitempty"`
		Name                  *string                                       `json:"name,omitempty"`
		Recipients            []string                                      `json:"recipients,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditional_recipients", "created", "filter", "modified", "name", "recipients"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ConditionalRecipients != nil && all.ConditionalRecipients.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConditionalRecipients = all.ConditionalRecipients
	o.Created = all.Created
	o.Filter = all.Filter
	o.Modified = all.Modified
	o.Name = all.Name
	o.Recipients = all.Recipients

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
