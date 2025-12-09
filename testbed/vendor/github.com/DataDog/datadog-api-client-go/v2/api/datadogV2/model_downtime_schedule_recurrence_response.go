// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeScheduleRecurrenceResponse An RRULE-based recurring downtime.
type DowntimeScheduleRecurrenceResponse struct {
	// The length of the downtime. Must begin with an integer and end with one of 'm', 'h', d', or 'w'.
	Duration *string `json:"duration,omitempty"`
	// The `RRULE` standard for defining recurring events.
	// For example, to have a recurring event on the first day of each month, set the type to `rrule` and set the `FREQ` to `MONTHLY` and `BYMONTHDAY` to `1`.
	// Most common `rrule` options from the [iCalendar Spec](https://tools.ietf.org/html/rfc5545) are supported.
	//
	// **Note**: Attributes specifying the duration in `RRULE` are not supported (for example, `DTSTART`, `DTEND`, `DURATION`).
	// More examples available in this [downtime guide](https://docs.datadoghq.com/monitors/guide/suppress-alert-with-downtimes/?tab=api).
	Rrule *string `json:"rrule,omitempty"`
	// ISO-8601 Datetime to start the downtime. Must not include a UTC offset. If not provided, the
	// downtime starts the moment it is created.
	Start *string `json:"start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeScheduleRecurrenceResponse instantiates a new DowntimeScheduleRecurrenceResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeScheduleRecurrenceResponse() *DowntimeScheduleRecurrenceResponse {
	this := DowntimeScheduleRecurrenceResponse{}
	return &this
}

// NewDowntimeScheduleRecurrenceResponseWithDefaults instantiates a new DowntimeScheduleRecurrenceResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeScheduleRecurrenceResponseWithDefaults() *DowntimeScheduleRecurrenceResponse {
	this := DowntimeScheduleRecurrenceResponse{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DowntimeScheduleRecurrenceResponse) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeScheduleRecurrenceResponse) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DowntimeScheduleRecurrenceResponse) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *DowntimeScheduleRecurrenceResponse) SetDuration(v string) {
	o.Duration = &v
}

// GetRrule returns the Rrule field value if set, zero value otherwise.
func (o *DowntimeScheduleRecurrenceResponse) GetRrule() string {
	if o == nil || o.Rrule == nil {
		var ret string
		return ret
	}
	return *o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeScheduleRecurrenceResponse) GetRruleOk() (*string, bool) {
	if o == nil || o.Rrule == nil {
		return nil, false
	}
	return o.Rrule, true
}

// HasRrule returns a boolean if a field has been set.
func (o *DowntimeScheduleRecurrenceResponse) HasRrule() bool {
	return o != nil && o.Rrule != nil
}

// SetRrule gets a reference to the given string and assigns it to the Rrule field.
func (o *DowntimeScheduleRecurrenceResponse) SetRrule(v string) {
	o.Rrule = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *DowntimeScheduleRecurrenceResponse) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeScheduleRecurrenceResponse) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *DowntimeScheduleRecurrenceResponse) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *DowntimeScheduleRecurrenceResponse) SetStart(v string) {
	o.Start = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeScheduleRecurrenceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Rrule != nil {
		toSerialize["rrule"] = o.Rrule
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeScheduleRecurrenceResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration *string `json:"duration,omitempty"`
		Rrule    *string `json:"rrule,omitempty"`
		Start    *string `json:"start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration", "rrule", "start"})
	} else {
		return err
	}
	o.Duration = all.Duration
	o.Rrule = all.Rrule
	o.Start = all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
