// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsDowntimeTimeSlotRecurrenceResponse Recurrence settings returned in a Synthetics downtime time slot response.
type SyntheticsDowntimeTimeSlotRecurrenceResponse struct {
	// The recurrence frequency of a Synthetics downtime time slot.
	Frequency SyntheticsDowntimeFrequency `json:"frequency"`
	// The interval between recurrences, relative to the frequency.
	Interval int64 `json:"interval"`
	// A specific date and time used to define the start or end of a Synthetics downtime time slot.
	Until *SyntheticsDowntimeTimeSlotDate `json:"until,omitempty"`
	// Positions of the weekdays within a month for a monthly Synthetics downtime recurrence. Used in combination with `weekdays` to schedule occurrences such as "the first Monday of the month".
	WeekdayPositions []SyntheticsDowntimeWeekdayPosition `json:"weekdayPositions,omitempty"`
	// Days of the week for a Synthetics downtime recurrence schedule.
	Weekdays []SyntheticsDowntimeWeekday `json:"weekdays"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsDowntimeTimeSlotRecurrenceResponse instantiates a new SyntheticsDowntimeTimeSlotRecurrenceResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsDowntimeTimeSlotRecurrenceResponse(frequency SyntheticsDowntimeFrequency, interval int64, weekdays []SyntheticsDowntimeWeekday) *SyntheticsDowntimeTimeSlotRecurrenceResponse {
	this := SyntheticsDowntimeTimeSlotRecurrenceResponse{}
	this.Frequency = frequency
	this.Interval = interval
	this.Weekdays = weekdays
	return &this
}

// NewSyntheticsDowntimeTimeSlotRecurrenceResponseWithDefaults instantiates a new SyntheticsDowntimeTimeSlotRecurrenceResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsDowntimeTimeSlotRecurrenceResponseWithDefaults() *SyntheticsDowntimeTimeSlotRecurrenceResponse {
	this := SyntheticsDowntimeTimeSlotRecurrenceResponse{}
	return &this
}

// GetFrequency returns the Frequency field value.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) GetFrequency() SyntheticsDowntimeFrequency {
	if o == nil {
		var ret SyntheticsDowntimeFrequency
		return ret
	}
	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) GetFrequencyOk() (*SyntheticsDowntimeFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) SetFrequency(v SyntheticsDowntimeFrequency) {
	o.Frequency = v
}

// GetInterval returns the Interval field value.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) GetInterval() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) GetIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) SetInterval(v int64) {
	o.Interval = v
}

// GetUntil returns the Until field value if set, zero value otherwise.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) GetUntil() SyntheticsDowntimeTimeSlotDate {
	if o == nil || o.Until == nil {
		var ret SyntheticsDowntimeTimeSlotDate
		return ret
	}
	return *o.Until
}

// GetUntilOk returns a tuple with the Until field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) GetUntilOk() (*SyntheticsDowntimeTimeSlotDate, bool) {
	if o == nil || o.Until == nil {
		return nil, false
	}
	return o.Until, true
}

// HasUntil returns a boolean if a field has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) HasUntil() bool {
	return o != nil && o.Until != nil
}

// SetUntil gets a reference to the given SyntheticsDowntimeTimeSlotDate and assigns it to the Until field.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) SetUntil(v SyntheticsDowntimeTimeSlotDate) {
	o.Until = &v
}

// GetWeekdayPositions returns the WeekdayPositions field value if set, zero value otherwise.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) GetWeekdayPositions() []SyntheticsDowntimeWeekdayPosition {
	if o == nil || o.WeekdayPositions == nil {
		var ret []SyntheticsDowntimeWeekdayPosition
		return ret
	}
	return o.WeekdayPositions
}

// GetWeekdayPositionsOk returns a tuple with the WeekdayPositions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) GetWeekdayPositionsOk() (*[]SyntheticsDowntimeWeekdayPosition, bool) {
	if o == nil || o.WeekdayPositions == nil {
		return nil, false
	}
	return &o.WeekdayPositions, true
}

// HasWeekdayPositions returns a boolean if a field has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) HasWeekdayPositions() bool {
	return o != nil && o.WeekdayPositions != nil
}

// SetWeekdayPositions gets a reference to the given []SyntheticsDowntimeWeekdayPosition and assigns it to the WeekdayPositions field.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) SetWeekdayPositions(v []SyntheticsDowntimeWeekdayPosition) {
	o.WeekdayPositions = v
}

// GetWeekdays returns the Weekdays field value.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) GetWeekdays() []SyntheticsDowntimeWeekday {
	if o == nil {
		var ret []SyntheticsDowntimeWeekday
		return ret
	}
	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) GetWeekdaysOk() (*[]SyntheticsDowntimeWeekday, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weekdays, true
}

// SetWeekdays sets field value.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) SetWeekdays(v []SyntheticsDowntimeWeekday) {
	o.Weekdays = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsDowntimeTimeSlotRecurrenceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["frequency"] = o.Frequency
	toSerialize["interval"] = o.Interval
	if o.Until != nil {
		toSerialize["until"] = o.Until
	}
	if o.WeekdayPositions != nil {
		toSerialize["weekdayPositions"] = o.WeekdayPositions
	}
	toSerialize["weekdays"] = o.Weekdays

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsDowntimeTimeSlotRecurrenceResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Frequency        *SyntheticsDowntimeFrequency        `json:"frequency"`
		Interval         *int64                              `json:"interval"`
		Until            *SyntheticsDowntimeTimeSlotDate     `json:"until,omitempty"`
		WeekdayPositions []SyntheticsDowntimeWeekdayPosition `json:"weekdayPositions,omitempty"`
		Weekdays         *[]SyntheticsDowntimeWeekday        `json:"weekdays"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Frequency == nil {
		return fmt.Errorf("required field frequency missing")
	}
	if all.Interval == nil {
		return fmt.Errorf("required field interval missing")
	}
	if all.Weekdays == nil {
		return fmt.Errorf("required field weekdays missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"frequency", "interval", "until", "weekdayPositions", "weekdays"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Frequency.IsValid() {
		hasInvalidField = true
	} else {
		o.Frequency = *all.Frequency
	}
	o.Interval = *all.Interval
	if all.Until != nil && all.Until.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Until = all.Until
	o.WeekdayPositions = all.WeekdayPositions
	o.Weekdays = *all.Weekdays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
