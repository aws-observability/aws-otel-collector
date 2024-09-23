// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeScheduleRecurrencesResponse A recurring downtime schedule definition.
type DowntimeScheduleRecurrencesResponse struct {
	// The most recent actual start and end dates for a recurring downtime. For a canceled downtime,
	// this is the previously occurring downtime. For active downtimes, this is the ongoing downtime, and for scheduled
	// downtimes it is the upcoming downtime.
	CurrentDowntime *DowntimeScheduleCurrentDowntimeResponse `json:"current_downtime,omitempty"`
	// A list of downtime recurrences.
	Recurrences []DowntimeScheduleRecurrenceResponse `json:"recurrences"`
	// The timezone in which to schedule the downtime. This affects recurring start and end dates.
	// Must match `display_timezone`.
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeScheduleRecurrencesResponse instantiates a new DowntimeScheduleRecurrencesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeScheduleRecurrencesResponse(recurrences []DowntimeScheduleRecurrenceResponse) *DowntimeScheduleRecurrencesResponse {
	this := DowntimeScheduleRecurrencesResponse{}
	this.Recurrences = recurrences
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// NewDowntimeScheduleRecurrencesResponseWithDefaults instantiates a new DowntimeScheduleRecurrencesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeScheduleRecurrencesResponseWithDefaults() *DowntimeScheduleRecurrencesResponse {
	this := DowntimeScheduleRecurrencesResponse{}
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// GetCurrentDowntime returns the CurrentDowntime field value if set, zero value otherwise.
func (o *DowntimeScheduleRecurrencesResponse) GetCurrentDowntime() DowntimeScheduleCurrentDowntimeResponse {
	if o == nil || o.CurrentDowntime == nil {
		var ret DowntimeScheduleCurrentDowntimeResponse
		return ret
	}
	return *o.CurrentDowntime
}

// GetCurrentDowntimeOk returns a tuple with the CurrentDowntime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeScheduleRecurrencesResponse) GetCurrentDowntimeOk() (*DowntimeScheduleCurrentDowntimeResponse, bool) {
	if o == nil || o.CurrentDowntime == nil {
		return nil, false
	}
	return o.CurrentDowntime, true
}

// HasCurrentDowntime returns a boolean if a field has been set.
func (o *DowntimeScheduleRecurrencesResponse) HasCurrentDowntime() bool {
	return o != nil && o.CurrentDowntime != nil
}

// SetCurrentDowntime gets a reference to the given DowntimeScheduleCurrentDowntimeResponse and assigns it to the CurrentDowntime field.
func (o *DowntimeScheduleRecurrencesResponse) SetCurrentDowntime(v DowntimeScheduleCurrentDowntimeResponse) {
	o.CurrentDowntime = &v
}

// GetRecurrences returns the Recurrences field value.
func (o *DowntimeScheduleRecurrencesResponse) GetRecurrences() []DowntimeScheduleRecurrenceResponse {
	if o == nil {
		var ret []DowntimeScheduleRecurrenceResponse
		return ret
	}
	return o.Recurrences
}

// GetRecurrencesOk returns a tuple with the Recurrences field value
// and a boolean to check if the value has been set.
func (o *DowntimeScheduleRecurrencesResponse) GetRecurrencesOk() (*[]DowntimeScheduleRecurrenceResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurrences, true
}

// SetRecurrences sets field value.
func (o *DowntimeScheduleRecurrencesResponse) SetRecurrences(v []DowntimeScheduleRecurrenceResponse) {
	o.Recurrences = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *DowntimeScheduleRecurrencesResponse) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeScheduleRecurrencesResponse) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *DowntimeScheduleRecurrencesResponse) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *DowntimeScheduleRecurrencesResponse) SetTimezone(v string) {
	o.Timezone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeScheduleRecurrencesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CurrentDowntime != nil {
		toSerialize["current_downtime"] = o.CurrentDowntime
	}
	toSerialize["recurrences"] = o.Recurrences
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeScheduleRecurrencesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CurrentDowntime *DowntimeScheduleCurrentDowntimeResponse `json:"current_downtime,omitempty"`
		Recurrences     *[]DowntimeScheduleRecurrenceResponse    `json:"recurrences"`
		Timezone        *string                                  `json:"timezone,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Recurrences == nil {
		return fmt.Errorf("required field recurrences missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"current_downtime", "recurrences", "timezone"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CurrentDowntime != nil && all.CurrentDowntime.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CurrentDowntime = all.CurrentDowntime
	o.Recurrences = *all.Recurrences
	o.Timezone = all.Timezone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
