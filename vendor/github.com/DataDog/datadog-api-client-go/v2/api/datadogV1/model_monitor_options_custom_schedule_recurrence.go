// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorOptionsCustomScheduleRecurrence Configuration for a recurrence set on the monitor options for custom schedule.
type MonitorOptionsCustomScheduleRecurrence struct {
	// Defines the recurrence rule (RRULE) for a given schedule.
	Rrule *string `json:"rrule,omitempty"`
	// Defines the start date and time of the recurring schedule.
	Start *string `json:"start,omitempty"`
	// Defines the timezone the schedule runs on.
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorOptionsCustomScheduleRecurrence instantiates a new MonitorOptionsCustomScheduleRecurrence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorOptionsCustomScheduleRecurrence() *MonitorOptionsCustomScheduleRecurrence {
	this := MonitorOptionsCustomScheduleRecurrence{}
	return &this
}

// NewMonitorOptionsCustomScheduleRecurrenceWithDefaults instantiates a new MonitorOptionsCustomScheduleRecurrence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorOptionsCustomScheduleRecurrenceWithDefaults() *MonitorOptionsCustomScheduleRecurrence {
	this := MonitorOptionsCustomScheduleRecurrence{}
	return &this
}

// GetRrule returns the Rrule field value if set, zero value otherwise.
func (o *MonitorOptionsCustomScheduleRecurrence) GetRrule() string {
	if o == nil || o.Rrule == nil {
		var ret string
		return ret
	}
	return *o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorOptionsCustomScheduleRecurrence) GetRruleOk() (*string, bool) {
	if o == nil || o.Rrule == nil {
		return nil, false
	}
	return o.Rrule, true
}

// HasRrule returns a boolean if a field has been set.
func (o *MonitorOptionsCustomScheduleRecurrence) HasRrule() bool {
	return o != nil && o.Rrule != nil
}

// SetRrule gets a reference to the given string and assigns it to the Rrule field.
func (o *MonitorOptionsCustomScheduleRecurrence) SetRrule(v string) {
	o.Rrule = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *MonitorOptionsCustomScheduleRecurrence) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorOptionsCustomScheduleRecurrence) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *MonitorOptionsCustomScheduleRecurrence) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *MonitorOptionsCustomScheduleRecurrence) SetStart(v string) {
	o.Start = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *MonitorOptionsCustomScheduleRecurrence) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorOptionsCustomScheduleRecurrence) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *MonitorOptionsCustomScheduleRecurrence) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *MonitorOptionsCustomScheduleRecurrence) SetTimezone(v string) {
	o.Timezone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorOptionsCustomScheduleRecurrence) MarshalJSON() ([]byte, error) {
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
func (o *MonitorOptionsCustomScheduleRecurrence) UnmarshalJSON(bytes []byte) (err error) {
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
