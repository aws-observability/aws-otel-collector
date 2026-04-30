// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestOptionsMonitorOptions Object containing the options for a Synthetic test as a monitor
// (for example, renotification).
type SyntheticsTestOptionsMonitorOptions struct {
	// Message to include in the escalation notification.
	EscalationMessage *string `json:"escalation_message,omitempty"`
	// The name of the preset for the notification for the monitor.
	NotificationPresetName *SyntheticsTestOptionsMonitorOptionsNotificationPresetName `json:"notification_preset_name,omitempty"`
	// Time interval before renotifying if the test is still failing
	// (in minutes).
	RenotifyInterval *int64 `json:"renotify_interval,omitempty"`
	// The number of times to renotify if the test is still failing.
	RenotifyOccurrences *int64 `json:"renotify_occurrences,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestOptionsMonitorOptions instantiates a new SyntheticsTestOptionsMonitorOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestOptionsMonitorOptions() *SyntheticsTestOptionsMonitorOptions {
	this := SyntheticsTestOptionsMonitorOptions{}
	return &this
}

// NewSyntheticsTestOptionsMonitorOptionsWithDefaults instantiates a new SyntheticsTestOptionsMonitorOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestOptionsMonitorOptionsWithDefaults() *SyntheticsTestOptionsMonitorOptions {
	this := SyntheticsTestOptionsMonitorOptions{}
	return &this
}

// GetEscalationMessage returns the EscalationMessage field value if set, zero value otherwise.
func (o *SyntheticsTestOptionsMonitorOptions) GetEscalationMessage() string {
	if o == nil || o.EscalationMessage == nil {
		var ret string
		return ret
	}
	return *o.EscalationMessage
}

// GetEscalationMessageOk returns a tuple with the EscalationMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptionsMonitorOptions) GetEscalationMessageOk() (*string, bool) {
	if o == nil || o.EscalationMessage == nil {
		return nil, false
	}
	return o.EscalationMessage, true
}

// HasEscalationMessage returns a boolean if a field has been set.
func (o *SyntheticsTestOptionsMonitorOptions) HasEscalationMessage() bool {
	return o != nil && o.EscalationMessage != nil
}

// SetEscalationMessage gets a reference to the given string and assigns it to the EscalationMessage field.
func (o *SyntheticsTestOptionsMonitorOptions) SetEscalationMessage(v string) {
	o.EscalationMessage = &v
}

// GetNotificationPresetName returns the NotificationPresetName field value if set, zero value otherwise.
func (o *SyntheticsTestOptionsMonitorOptions) GetNotificationPresetName() SyntheticsTestOptionsMonitorOptionsNotificationPresetName {
	if o == nil || o.NotificationPresetName == nil {
		var ret SyntheticsTestOptionsMonitorOptionsNotificationPresetName
		return ret
	}
	return *o.NotificationPresetName
}

// GetNotificationPresetNameOk returns a tuple with the NotificationPresetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptionsMonitorOptions) GetNotificationPresetNameOk() (*SyntheticsTestOptionsMonitorOptionsNotificationPresetName, bool) {
	if o == nil || o.NotificationPresetName == nil {
		return nil, false
	}
	return o.NotificationPresetName, true
}

// HasNotificationPresetName returns a boolean if a field has been set.
func (o *SyntheticsTestOptionsMonitorOptions) HasNotificationPresetName() bool {
	return o != nil && o.NotificationPresetName != nil
}

// SetNotificationPresetName gets a reference to the given SyntheticsTestOptionsMonitorOptionsNotificationPresetName and assigns it to the NotificationPresetName field.
func (o *SyntheticsTestOptionsMonitorOptions) SetNotificationPresetName(v SyntheticsTestOptionsMonitorOptionsNotificationPresetName) {
	o.NotificationPresetName = &v
}

// GetRenotifyInterval returns the RenotifyInterval field value if set, zero value otherwise.
func (o *SyntheticsTestOptionsMonitorOptions) GetRenotifyInterval() int64 {
	if o == nil || o.RenotifyInterval == nil {
		var ret int64
		return ret
	}
	return *o.RenotifyInterval
}

// GetRenotifyIntervalOk returns a tuple with the RenotifyInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptionsMonitorOptions) GetRenotifyIntervalOk() (*int64, bool) {
	if o == nil || o.RenotifyInterval == nil {
		return nil, false
	}
	return o.RenotifyInterval, true
}

// HasRenotifyInterval returns a boolean if a field has been set.
func (o *SyntheticsTestOptionsMonitorOptions) HasRenotifyInterval() bool {
	return o != nil && o.RenotifyInterval != nil
}

// SetRenotifyInterval gets a reference to the given int64 and assigns it to the RenotifyInterval field.
func (o *SyntheticsTestOptionsMonitorOptions) SetRenotifyInterval(v int64) {
	o.RenotifyInterval = &v
}

// GetRenotifyOccurrences returns the RenotifyOccurrences field value if set, zero value otherwise.
func (o *SyntheticsTestOptionsMonitorOptions) GetRenotifyOccurrences() int64 {
	if o == nil || o.RenotifyOccurrences == nil {
		var ret int64
		return ret
	}
	return *o.RenotifyOccurrences
}

// GetRenotifyOccurrencesOk returns a tuple with the RenotifyOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptionsMonitorOptions) GetRenotifyOccurrencesOk() (*int64, bool) {
	if o == nil || o.RenotifyOccurrences == nil {
		return nil, false
	}
	return o.RenotifyOccurrences, true
}

// HasRenotifyOccurrences returns a boolean if a field has been set.
func (o *SyntheticsTestOptionsMonitorOptions) HasRenotifyOccurrences() bool {
	return o != nil && o.RenotifyOccurrences != nil
}

// SetRenotifyOccurrences gets a reference to the given int64 and assigns it to the RenotifyOccurrences field.
func (o *SyntheticsTestOptionsMonitorOptions) SetRenotifyOccurrences(v int64) {
	o.RenotifyOccurrences = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestOptionsMonitorOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EscalationMessage != nil {
		toSerialize["escalation_message"] = o.EscalationMessage
	}
	if o.NotificationPresetName != nil {
		toSerialize["notification_preset_name"] = o.NotificationPresetName
	}
	if o.RenotifyInterval != nil {
		toSerialize["renotify_interval"] = o.RenotifyInterval
	}
	if o.RenotifyOccurrences != nil {
		toSerialize["renotify_occurrences"] = o.RenotifyOccurrences
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestOptionsMonitorOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EscalationMessage      *string                                                    `json:"escalation_message,omitempty"`
		NotificationPresetName *SyntheticsTestOptionsMonitorOptionsNotificationPresetName `json:"notification_preset_name,omitempty"`
		RenotifyInterval       *int64                                                     `json:"renotify_interval,omitempty"`
		RenotifyOccurrences    *int64                                                     `json:"renotify_occurrences,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"escalation_message", "notification_preset_name", "renotify_interval", "renotify_occurrences"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EscalationMessage = all.EscalationMessage
	if all.NotificationPresetName != nil && !all.NotificationPresetName.IsValid() {
		hasInvalidField = true
	} else {
		o.NotificationPresetName = all.NotificationPresetName
	}
	o.RenotifyInterval = all.RenotifyInterval
	o.RenotifyOccurrences = all.RenotifyOccurrences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
