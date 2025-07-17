// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeRestriction Defines a single time restriction rule with start and end times and the applicable weekdays.
type TimeRestriction struct {
	// A day of the week.
	EndDay *Weekday `json:"end_day,omitempty"`
	// Specifies the ending time for this restriction.
	EndTime *string `json:"end_time,omitempty"`
	// A day of the week.
	StartDay *Weekday `json:"start_day,omitempty"`
	// Specifies the starting time for this restriction.
	StartTime *string `json:"start_time,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeRestriction instantiates a new TimeRestriction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeRestriction() *TimeRestriction {
	this := TimeRestriction{}
	return &this
}

// NewTimeRestrictionWithDefaults instantiates a new TimeRestriction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeRestrictionWithDefaults() *TimeRestriction {
	this := TimeRestriction{}
	return &this
}

// GetEndDay returns the EndDay field value if set, zero value otherwise.
func (o *TimeRestriction) GetEndDay() Weekday {
	if o == nil || o.EndDay == nil {
		var ret Weekday
		return ret
	}
	return *o.EndDay
}

// GetEndDayOk returns a tuple with the EndDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRestriction) GetEndDayOk() (*Weekday, bool) {
	if o == nil || o.EndDay == nil {
		return nil, false
	}
	return o.EndDay, true
}

// HasEndDay returns a boolean if a field has been set.
func (o *TimeRestriction) HasEndDay() bool {
	return o != nil && o.EndDay != nil
}

// SetEndDay gets a reference to the given Weekday and assigns it to the EndDay field.
func (o *TimeRestriction) SetEndDay(v Weekday) {
	o.EndDay = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TimeRestriction) GetEndTime() string {
	if o == nil || o.EndTime == nil {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRestriction) GetEndTimeOk() (*string, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TimeRestriction) HasEndTime() bool {
	return o != nil && o.EndTime != nil
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *TimeRestriction) SetEndTime(v string) {
	o.EndTime = &v
}

// GetStartDay returns the StartDay field value if set, zero value otherwise.
func (o *TimeRestriction) GetStartDay() Weekday {
	if o == nil || o.StartDay == nil {
		var ret Weekday
		return ret
	}
	return *o.StartDay
}

// GetStartDayOk returns a tuple with the StartDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRestriction) GetStartDayOk() (*Weekday, bool) {
	if o == nil || o.StartDay == nil {
		return nil, false
	}
	return o.StartDay, true
}

// HasStartDay returns a boolean if a field has been set.
func (o *TimeRestriction) HasStartDay() bool {
	return o != nil && o.StartDay != nil
}

// SetStartDay gets a reference to the given Weekday and assigns it to the StartDay field.
func (o *TimeRestriction) SetStartDay(v Weekday) {
	o.StartDay = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TimeRestriction) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRestriction) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TimeRestriction) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *TimeRestriction) SetStartTime(v string) {
	o.StartTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeRestriction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EndDay != nil {
		toSerialize["end_day"] = o.EndDay
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	if o.StartDay != nil {
		toSerialize["start_day"] = o.StartDay
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeRestriction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndDay    *Weekday `json:"end_day,omitempty"`
		EndTime   *string  `json:"end_time,omitempty"`
		StartDay  *Weekday `json:"start_day,omitempty"`
		StartTime *string  `json:"start_time,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end_day", "end_time", "start_day", "start_time"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.EndDay != nil && !all.EndDay.IsValid() {
		hasInvalidField = true
	} else {
		o.EndDay = all.EndDay
	}
	o.EndTime = all.EndTime
	if all.StartDay != nil && !all.StartDay.IsValid() {
		hasInvalidField = true
	} else {
		o.StartDay = all.StartDay
	}
	o.StartTime = all.StartTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
