// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsDowntimeTimeSlotDate A specific date and time used to define the start or end of a Synthetics downtime time slot.
type SyntheticsDowntimeTimeSlotDate struct {
	// The day component of the date (1-31).
	Day int64 `json:"day"`
	// The hour component of the time (0-23).
	Hour int64 `json:"hour"`
	// The minute component of the time (0-59).
	Minute int64 `json:"minute"`
	// The month component of the date (1-12).
	Month int64 `json:"month"`
	// The year component of the date.
	Year int64 `json:"year"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsDowntimeTimeSlotDate instantiates a new SyntheticsDowntimeTimeSlotDate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsDowntimeTimeSlotDate(day int64, hour int64, minute int64, month int64, year int64) *SyntheticsDowntimeTimeSlotDate {
	this := SyntheticsDowntimeTimeSlotDate{}
	this.Day = day
	this.Hour = hour
	this.Minute = minute
	this.Month = month
	this.Year = year
	return &this
}

// NewSyntheticsDowntimeTimeSlotDateWithDefaults instantiates a new SyntheticsDowntimeTimeSlotDate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsDowntimeTimeSlotDateWithDefaults() *SyntheticsDowntimeTimeSlotDate {
	this := SyntheticsDowntimeTimeSlotDate{}
	return &this
}

// GetDay returns the Day field value.
func (o *SyntheticsDowntimeTimeSlotDate) GetDay() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotDate) GetDayOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value.
func (o *SyntheticsDowntimeTimeSlotDate) SetDay(v int64) {
	o.Day = v
}

// GetHour returns the Hour field value.
func (o *SyntheticsDowntimeTimeSlotDate) GetHour() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Hour
}

// GetHourOk returns a tuple with the Hour field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotDate) GetHourOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hour, true
}

// SetHour sets field value.
func (o *SyntheticsDowntimeTimeSlotDate) SetHour(v int64) {
	o.Hour = v
}

// GetMinute returns the Minute field value.
func (o *SyntheticsDowntimeTimeSlotDate) GetMinute() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotDate) GetMinuteOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minute, true
}

// SetMinute sets field value.
func (o *SyntheticsDowntimeTimeSlotDate) SetMinute(v int64) {
	o.Minute = v
}

// GetMonth returns the Month field value.
func (o *SyntheticsDowntimeTimeSlotDate) GetMonth() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotDate) GetMonthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value.
func (o *SyntheticsDowntimeTimeSlotDate) SetMonth(v int64) {
	o.Month = v
}

// GetYear returns the Year field value.
func (o *SyntheticsDowntimeTimeSlotDate) GetYear() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotDate) GetYearOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value.
func (o *SyntheticsDowntimeTimeSlotDate) SetYear(v int64) {
	o.Year = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsDowntimeTimeSlotDate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["day"] = o.Day
	toSerialize["hour"] = o.Hour
	toSerialize["minute"] = o.Minute
	toSerialize["month"] = o.Month
	toSerialize["year"] = o.Year

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsDowntimeTimeSlotDate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Day    *int64 `json:"day"`
		Hour   *int64 `json:"hour"`
		Minute *int64 `json:"minute"`
		Month  *int64 `json:"month"`
		Year   *int64 `json:"year"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Day == nil {
		return fmt.Errorf("required field day missing")
	}
	if all.Hour == nil {
		return fmt.Errorf("required field hour missing")
	}
	if all.Minute == nil {
		return fmt.Errorf("required field minute missing")
	}
	if all.Month == nil {
		return fmt.Errorf("required field month missing")
	}
	if all.Year == nil {
		return fmt.Errorf("required field year missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"day", "hour", "minute", "month", "year"})
	} else {
		return err
	}
	o.Day = *all.Day
	o.Hour = *all.Hour
	o.Minute = *all.Minute
	o.Month = *all.Month
	o.Year = *all.Year

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
