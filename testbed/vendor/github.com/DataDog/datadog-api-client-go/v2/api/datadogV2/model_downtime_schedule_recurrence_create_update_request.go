// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeScheduleRecurrenceCreateUpdateRequest An object defining the recurrence of the downtime.
type DowntimeScheduleRecurrenceCreateUpdateRequest struct {
	// The length of the downtime. Must begin with an integer and end with one of 'm', 'h', d', or 'w'.
	Duration string `json:"duration"`
	// The `RRULE` standard for defining recurring events.
	// For example, to have a recurring event on the first day of each month, set the type to `rrule` and set the `FREQ` to `MONTHLY` and `BYMONTHDAY` to `1`.
	// Most common `rrule` options from the [iCalendar Spec](https://tools.ietf.org/html/rfc5545) are supported.
	//
	// **Note**: Attributes specifying the duration in `RRULE` are not supported (for example, `DTSTART`, `DTEND`, `DURATION`).
	// More examples available in this [downtime guide](https://docs.datadoghq.com/monitors/guide/suppress-alert-with-downtimes/?tab=api).
	Rrule string `json:"rrule"`
	// ISO-8601 Datetime to start the downtime. Must not include a UTC offset. If not provided, the
	// downtime starts the moment it is created.
	Start datadog.NullableString `json:"start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeScheduleRecurrenceCreateUpdateRequest instantiates a new DowntimeScheduleRecurrenceCreateUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeScheduleRecurrenceCreateUpdateRequest(duration string, rrule string) *DowntimeScheduleRecurrenceCreateUpdateRequest {
	this := DowntimeScheduleRecurrenceCreateUpdateRequest{}
	this.Duration = duration
	this.Rrule = rrule
	return &this
}

// NewDowntimeScheduleRecurrenceCreateUpdateRequestWithDefaults instantiates a new DowntimeScheduleRecurrenceCreateUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeScheduleRecurrenceCreateUpdateRequestWithDefaults() *DowntimeScheduleRecurrenceCreateUpdateRequest {
	this := DowntimeScheduleRecurrenceCreateUpdateRequest{}
	return &this
}

// GetDuration returns the Duration field value.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) SetDuration(v string) {
	o.Duration = v
}

// GetRrule returns the Rrule field value.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) GetRrule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value
// and a boolean to check if the value has been set.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) GetRruleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rrule, true
}

// SetRrule sets field value.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) SetRrule(v string) {
	o.Rrule = v
}

// GetStart returns the Start field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) GetStart() string {
	if o == nil || o.Start.Get() == nil {
		var ret string
		return ret
	}
	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// HasStart returns a boolean if a field has been set.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) HasStart() bool {
	return o != nil && o.Start.IsSet()
}

// SetStart gets a reference to the given datadog.NullableString and assigns it to the Start field.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) SetStart(v string) {
	o.Start.Set(&v)
}

// SetStartNil sets the value for Start to be an explicit nil.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) SetStartNil() {
	o.Start.Set(nil)
}

// UnsetStart ensures that no value is present for Start, not even an explicit nil.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) UnsetStart() {
	o.Start.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeScheduleRecurrenceCreateUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["duration"] = o.Duration
	toSerialize["rrule"] = o.Rrule
	if o.Start.IsSet() {
		toSerialize["start"] = o.Start.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeScheduleRecurrenceCreateUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration *string                `json:"duration"`
		Rrule    *string                `json:"rrule"`
		Start    datadog.NullableString `json:"start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Duration == nil {
		return fmt.Errorf("required field duration missing")
	}
	if all.Rrule == nil {
		return fmt.Errorf("required field rrule missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration", "rrule", "start"})
	} else {
		return err
	}
	o.Duration = *all.Duration
	o.Rrule = *all.Rrule
	o.Start = all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
