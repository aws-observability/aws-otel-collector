// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsDowntimeTimeSlotRecurrenceRequest Recurrence settings for a Synthetics downtime time slot.
type SyntheticsDowntimeTimeSlotRecurrenceRequest struct {
	// A specific date and time used to define the start or end of a Synthetics downtime time slot.
	End *SyntheticsDowntimeTimeSlotDate `json:"end,omitempty"`
	// The recurrence frequency of a Synthetics downtime time slot.
	Frequency SyntheticsDowntimeFrequency `json:"frequency"`
	// The interval between recurrences, relative to the frequency.
	Interval *int64 `json:"interval,omitempty"`
	// Positions of the weekdays within a month for a monthly Synthetics downtime recurrence. Used in combination with `weekdays` to schedule occurrences such as "the first Monday of the month".
	WeekdayPositions []SyntheticsDowntimeWeekdayPosition `json:"weekdayPositions,omitempty"`
	// Days of the week for a Synthetics downtime recurrence schedule.
	Weekdays []SyntheticsDowntimeWeekday `json:"weekdays,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsDowntimeTimeSlotRecurrenceRequest instantiates a new SyntheticsDowntimeTimeSlotRecurrenceRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsDowntimeTimeSlotRecurrenceRequest(frequency SyntheticsDowntimeFrequency) *SyntheticsDowntimeTimeSlotRecurrenceRequest {
	this := SyntheticsDowntimeTimeSlotRecurrenceRequest{}
	this.Frequency = frequency
	return &this
}

// NewSyntheticsDowntimeTimeSlotRecurrenceRequestWithDefaults instantiates a new SyntheticsDowntimeTimeSlotRecurrenceRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsDowntimeTimeSlotRecurrenceRequestWithDefaults() *SyntheticsDowntimeTimeSlotRecurrenceRequest {
	this := SyntheticsDowntimeTimeSlotRecurrenceRequest{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) GetEnd() SyntheticsDowntimeTimeSlotDate {
	if o == nil || o.End == nil {
		var ret SyntheticsDowntimeTimeSlotDate
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) GetEndOk() (*SyntheticsDowntimeTimeSlotDate, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given SyntheticsDowntimeTimeSlotDate and assigns it to the End field.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) SetEnd(v SyntheticsDowntimeTimeSlotDate) {
	o.End = &v
}

// GetFrequency returns the Frequency field value.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) GetFrequency() SyntheticsDowntimeFrequency {
	if o == nil {
		var ret SyntheticsDowntimeFrequency
		return ret
	}
	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) GetFrequencyOk() (*SyntheticsDowntimeFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) SetFrequency(v SyntheticsDowntimeFrequency) {
	o.Frequency = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) GetInterval() int64 {
	if o == nil || o.Interval == nil {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) GetIntervalOk() (*int64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) SetInterval(v int64) {
	o.Interval = &v
}

// GetWeekdayPositions returns the WeekdayPositions field value if set, zero value otherwise.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) GetWeekdayPositions() []SyntheticsDowntimeWeekdayPosition {
	if o == nil || o.WeekdayPositions == nil {
		var ret []SyntheticsDowntimeWeekdayPosition
		return ret
	}
	return o.WeekdayPositions
}

// GetWeekdayPositionsOk returns a tuple with the WeekdayPositions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) GetWeekdayPositionsOk() (*[]SyntheticsDowntimeWeekdayPosition, bool) {
	if o == nil || o.WeekdayPositions == nil {
		return nil, false
	}
	return &o.WeekdayPositions, true
}

// HasWeekdayPositions returns a boolean if a field has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) HasWeekdayPositions() bool {
	return o != nil && o.WeekdayPositions != nil
}

// SetWeekdayPositions gets a reference to the given []SyntheticsDowntimeWeekdayPosition and assigns it to the WeekdayPositions field.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) SetWeekdayPositions(v []SyntheticsDowntimeWeekdayPosition) {
	o.WeekdayPositions = v
}

// GetWeekdays returns the Weekdays field value if set, zero value otherwise.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) GetWeekdays() []SyntheticsDowntimeWeekday {
	if o == nil || o.Weekdays == nil {
		var ret []SyntheticsDowntimeWeekday
		return ret
	}
	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) GetWeekdaysOk() (*[]SyntheticsDowntimeWeekday, bool) {
	if o == nil || o.Weekdays == nil {
		return nil, false
	}
	return &o.Weekdays, true
}

// HasWeekdays returns a boolean if a field has been set.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) HasWeekdays() bool {
	return o != nil && o.Weekdays != nil
}

// SetWeekdays gets a reference to the given []SyntheticsDowntimeWeekday and assigns it to the Weekdays field.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) SetWeekdays(v []SyntheticsDowntimeWeekday) {
	o.Weekdays = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsDowntimeTimeSlotRecurrenceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	toSerialize["frequency"] = o.Frequency
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.WeekdayPositions != nil {
		toSerialize["weekdayPositions"] = o.WeekdayPositions
	}
	if o.Weekdays != nil {
		toSerialize["weekdays"] = o.Weekdays
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsDowntimeTimeSlotRecurrenceRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End              *SyntheticsDowntimeTimeSlotDate     `json:"end,omitempty"`
		Frequency        *SyntheticsDowntimeFrequency        `json:"frequency"`
		Interval         *int64                              `json:"interval,omitempty"`
		WeekdayPositions []SyntheticsDowntimeWeekdayPosition `json:"weekdayPositions,omitempty"`
		Weekdays         []SyntheticsDowntimeWeekday         `json:"weekdays,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Frequency == nil {
		return fmt.Errorf("required field frequency missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "frequency", "interval", "weekdayPositions", "weekdays"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.End != nil && all.End.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.End = all.End
	if !all.Frequency.IsValid() {
		hasInvalidField = true
	} else {
		o.Frequency = *all.Frequency
	}
	o.Interval = all.Interval
	o.WeekdayPositions = all.WeekdayPositions
	o.Weekdays = all.Weekdays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
