// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSchedulingOptions Options for scheduled rules. When this field is present, the rule runs based on the schedule. When absent, it runs real-time on ingested logs.
type SecurityMonitoringSchedulingOptions struct {
	// Schedule for the rule queries, written in RRULE syntax. See [RFC](https://icalendar.org/iCalendar-RFC-5545/3-8-5-3-recurrence-rule.html) for syntax reference.
	Rrule *string `json:"rrule,omitempty"`
	// Start date for the schedule, in ISO 8601 format without timezone.
	Start *string `json:"start,omitempty"`
	// Time zone of the start date, in the [tz database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) format.
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSchedulingOptions instantiates a new SecurityMonitoringSchedulingOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSchedulingOptions() *SecurityMonitoringSchedulingOptions {
	this := SecurityMonitoringSchedulingOptions{}
	return &this
}

// NewSecurityMonitoringSchedulingOptionsWithDefaults instantiates a new SecurityMonitoringSchedulingOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSchedulingOptionsWithDefaults() *SecurityMonitoringSchedulingOptions {
	this := SecurityMonitoringSchedulingOptions{}
	return &this
}

// GetRrule returns the Rrule field value if set, zero value otherwise.
func (o *SecurityMonitoringSchedulingOptions) GetRrule() string {
	if o == nil || o.Rrule == nil {
		var ret string
		return ret
	}
	return *o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSchedulingOptions) GetRruleOk() (*string, bool) {
	if o == nil || o.Rrule == nil {
		return nil, false
	}
	return o.Rrule, true
}

// HasRrule returns a boolean if a field has been set.
func (o *SecurityMonitoringSchedulingOptions) HasRrule() bool {
	return o != nil && o.Rrule != nil
}

// SetRrule gets a reference to the given string and assigns it to the Rrule field.
func (o *SecurityMonitoringSchedulingOptions) SetRrule(v string) {
	o.Rrule = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SecurityMonitoringSchedulingOptions) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSchedulingOptions) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SecurityMonitoringSchedulingOptions) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *SecurityMonitoringSchedulingOptions) SetStart(v string) {
	o.Start = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *SecurityMonitoringSchedulingOptions) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSchedulingOptions) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *SecurityMonitoringSchedulingOptions) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *SecurityMonitoringSchedulingOptions) SetTimezone(v string) {
	o.Timezone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSchedulingOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Rrule != nil {
		toSerialize["rrule"] = o.Rrule
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSchedulingOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rrule    *string `json:"rrule,omitempty"`
		Start    *string `json:"start,omitempty"`
		Timezone *string `json:"timezone,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rrule", "start", "timezone"})
	} else {
		return err
	}
	o.Rrule = all.Rrule
	o.Start = all.Start
	o.Timezone = all.Timezone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableSecurityMonitoringSchedulingOptions handles when a null is used for SecurityMonitoringSchedulingOptions.
type NullableSecurityMonitoringSchedulingOptions struct {
	value *SecurityMonitoringSchedulingOptions
	isSet bool
}

// Get returns the associated value.
func (v NullableSecurityMonitoringSchedulingOptions) Get() *SecurityMonitoringSchedulingOptions {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSecurityMonitoringSchedulingOptions) Set(val *SecurityMonitoringSchedulingOptions) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSecurityMonitoringSchedulingOptions) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableSecurityMonitoringSchedulingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSecurityMonitoringSchedulingOptions initializes the struct as if Set has been called.
func NewNullableSecurityMonitoringSchedulingOptions(val *SecurityMonitoringSchedulingOptions) *NullableSecurityMonitoringSchedulingOptions {
	return &NullableSecurityMonitoringSchedulingOptions{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSecurityMonitoringSchedulingOptions) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSecurityMonitoringSchedulingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
