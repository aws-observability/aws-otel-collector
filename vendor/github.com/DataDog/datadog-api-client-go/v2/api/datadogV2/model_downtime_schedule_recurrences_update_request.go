// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeScheduleRecurrencesUpdateRequest A recurring downtime schedule definition.
type DowntimeScheduleRecurrencesUpdateRequest struct {
	// A list of downtime recurrences.
	Recurrences []DowntimeScheduleRecurrenceCreateUpdateRequest `json:"recurrences,omitempty"`
	// The timezone in which to schedule the downtime.
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDowntimeScheduleRecurrencesUpdateRequest instantiates a new DowntimeScheduleRecurrencesUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeScheduleRecurrencesUpdateRequest() *DowntimeScheduleRecurrencesUpdateRequest {
	this := DowntimeScheduleRecurrencesUpdateRequest{}
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// NewDowntimeScheduleRecurrencesUpdateRequestWithDefaults instantiates a new DowntimeScheduleRecurrencesUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeScheduleRecurrencesUpdateRequestWithDefaults() *DowntimeScheduleRecurrencesUpdateRequest {
	this := DowntimeScheduleRecurrencesUpdateRequest{}
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// GetRecurrences returns the Recurrences field value if set, zero value otherwise.
func (o *DowntimeScheduleRecurrencesUpdateRequest) GetRecurrences() []DowntimeScheduleRecurrenceCreateUpdateRequest {
	if o == nil || o.Recurrences == nil {
		var ret []DowntimeScheduleRecurrenceCreateUpdateRequest
		return ret
	}
	return o.Recurrences
}

// GetRecurrencesOk returns a tuple with the Recurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeScheduleRecurrencesUpdateRequest) GetRecurrencesOk() (*[]DowntimeScheduleRecurrenceCreateUpdateRequest, bool) {
	if o == nil || o.Recurrences == nil {
		return nil, false
	}
	return &o.Recurrences, true
}

// HasRecurrences returns a boolean if a field has been set.
func (o *DowntimeScheduleRecurrencesUpdateRequest) HasRecurrences() bool {
	return o != nil && o.Recurrences != nil
}

// SetRecurrences gets a reference to the given []DowntimeScheduleRecurrenceCreateUpdateRequest and assigns it to the Recurrences field.
func (o *DowntimeScheduleRecurrencesUpdateRequest) SetRecurrences(v []DowntimeScheduleRecurrenceCreateUpdateRequest) {
	o.Recurrences = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *DowntimeScheduleRecurrencesUpdateRequest) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeScheduleRecurrencesUpdateRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *DowntimeScheduleRecurrencesUpdateRequest) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *DowntimeScheduleRecurrencesUpdateRequest) SetTimezone(v string) {
	o.Timezone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeScheduleRecurrencesUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Recurrences != nil {
		toSerialize["recurrences"] = o.Recurrences
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeScheduleRecurrencesUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Recurrences []DowntimeScheduleRecurrenceCreateUpdateRequest `json:"recurrences,omitempty"`
		Timezone    *string                                         `json:"timezone,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Recurrences = all.Recurrences
	o.Timezone = all.Timezone

	return nil
}
