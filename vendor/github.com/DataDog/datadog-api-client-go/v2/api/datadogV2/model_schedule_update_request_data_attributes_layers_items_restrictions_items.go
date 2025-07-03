// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems Defines a time restriction object for a layer within a schedule update, including
// start and end days, as well as times.
type ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems struct {
	// Defines the day of the week on which the time restriction ends.
	EndDay *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay `json:"end_day,omitempty"`
	// The time at which this restriction ends (hh:mm:ss).
	EndTime *string `json:"end_time,omitempty"`
	// Defines the day of the week on which the time restriction starts.
	StartDay *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay `json:"start_day,omitempty"`
	// The time at which this restriction starts (hh:mm:ss).
	StartTime *string `json:"start_time,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems instantiates a new ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems() *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems {
	this := ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems{}
	return &this
}

// NewScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsWithDefaults instantiates a new ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsWithDefaults() *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems {
	this := ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems{}
	return &this
}

// GetEndDay returns the EndDay field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) GetEndDay() ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay {
	if o == nil || o.EndDay == nil {
		var ret ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay
		return ret
	}
	return *o.EndDay
}

// GetEndDayOk returns a tuple with the EndDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) GetEndDayOk() (*ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay, bool) {
	if o == nil || o.EndDay == nil {
		return nil, false
	}
	return o.EndDay, true
}

// HasEndDay returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) HasEndDay() bool {
	return o != nil && o.EndDay != nil
}

// SetEndDay gets a reference to the given ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay and assigns it to the EndDay field.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) SetEndDay(v ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay) {
	o.EndDay = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) GetEndTime() string {
	if o == nil || o.EndTime == nil {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) GetEndTimeOk() (*string, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) HasEndTime() bool {
	return o != nil && o.EndTime != nil
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) SetEndTime(v string) {
	o.EndTime = &v
}

// GetStartDay returns the StartDay field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) GetStartDay() ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay {
	if o == nil || o.StartDay == nil {
		var ret ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay
		return ret
	}
	return *o.StartDay
}

// GetStartDayOk returns a tuple with the StartDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) GetStartDayOk() (*ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay, bool) {
	if o == nil || o.StartDay == nil {
		return nil, false
	}
	return o.StartDay, true
}

// HasStartDay returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) HasStartDay() bool {
	return o != nil && o.StartDay != nil
}

// SetStartDay gets a reference to the given ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay and assigns it to the StartDay field.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) SetStartDay(v ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay) {
	o.StartDay = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) SetStartTime(v string) {
	o.StartTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) MarshalJSON() ([]byte, error) {
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
func (o *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndDay    *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsEndDay   `json:"end_day,omitempty"`
		EndTime   *string                                                                  `json:"end_time,omitempty"`
		StartDay  *ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItemsStartDay `json:"start_day,omitempty"`
		StartTime *string                                                                  `json:"start_time,omitempty"`
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
