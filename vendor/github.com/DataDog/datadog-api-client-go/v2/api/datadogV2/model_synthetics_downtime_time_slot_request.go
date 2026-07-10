// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsDowntimeTimeSlotRequest A time slot for a Synthetics downtime create or update request.
type SyntheticsDowntimeTimeSlotRequest struct {
	// The duration of the time slot in seconds, between 60 and 604800.
	Duration int64 `json:"duration"`
	// An optional label for the time slot.
	Name *string `json:"name,omitempty"`
	// Recurrence settings for a Synthetics downtime time slot.
	Recurrence *SyntheticsDowntimeTimeSlotRecurrenceRequest `json:"recurrence,omitempty"`
	// A specific date and time used to define the start or end of a Synthetics downtime time slot.
	Start SyntheticsDowntimeTimeSlotDate `json:"start"`
	// The IANA timezone name for the time slot.
	Timezone string `json:"timezone"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsDowntimeTimeSlotRequest instantiates a new SyntheticsDowntimeTimeSlotRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsDowntimeTimeSlotRequest(duration int64, start SyntheticsDowntimeTimeSlotDate, timezone string) *SyntheticsDowntimeTimeSlotRequest {
	this := SyntheticsDowntimeTimeSlotRequest{}
	this.Duration = duration
	this.Start = start
	this.Timezone = timezone
	return &this
}

// NewSyntheticsDowntimeTimeSlotRequestWithDefaults instantiates a new SyntheticsDowntimeTimeSlotRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsDowntimeTimeSlotRequestWithDefaults() *SyntheticsDowntimeTimeSlotRequest {
	this := SyntheticsDowntimeTimeSlotRequest{}
	return &this
}

// GetDuration returns the Duration field value.
func (o *SyntheticsDowntimeTimeSlotRequest) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRequest) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value.
func (o *SyntheticsDowntimeTimeSlotRequest) SetDuration(v int64) {
	o.Duration = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsDowntimeTimeSlotRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsDowntimeTimeSlotRequest) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsDowntimeTimeSlotRequest) SetName(v string) {
	o.Name = &v
}

// GetRecurrence returns the Recurrence field value if set, zero value otherwise.
func (o *SyntheticsDowntimeTimeSlotRequest) GetRecurrence() SyntheticsDowntimeTimeSlotRecurrenceRequest {
	if o == nil || o.Recurrence == nil {
		var ret SyntheticsDowntimeTimeSlotRecurrenceRequest
		return ret
	}
	return *o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRequest) GetRecurrenceOk() (*SyntheticsDowntimeTimeSlotRecurrenceRequest, bool) {
	if o == nil || o.Recurrence == nil {
		return nil, false
	}
	return o.Recurrence, true
}

// HasRecurrence returns a boolean if a field has been set.
func (o *SyntheticsDowntimeTimeSlotRequest) HasRecurrence() bool {
	return o != nil && o.Recurrence != nil
}

// SetRecurrence gets a reference to the given SyntheticsDowntimeTimeSlotRecurrenceRequest and assigns it to the Recurrence field.
func (o *SyntheticsDowntimeTimeSlotRequest) SetRecurrence(v SyntheticsDowntimeTimeSlotRecurrenceRequest) {
	o.Recurrence = &v
}

// GetStart returns the Start field value.
func (o *SyntheticsDowntimeTimeSlotRequest) GetStart() SyntheticsDowntimeTimeSlotDate {
	if o == nil {
		var ret SyntheticsDowntimeTimeSlotDate
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRequest) GetStartOk() (*SyntheticsDowntimeTimeSlotDate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *SyntheticsDowntimeTimeSlotRequest) SetStart(v SyntheticsDowntimeTimeSlotDate) {
	o.Start = v
}

// GetTimezone returns the Timezone field value.
func (o *SyntheticsDowntimeTimeSlotRequest) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeTimeSlotRequest) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value.
func (o *SyntheticsDowntimeTimeSlotRequest) SetTimezone(v string) {
	o.Timezone = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsDowntimeTimeSlotRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["duration"] = o.Duration
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Recurrence != nil {
		toSerialize["recurrence"] = o.Recurrence
	}
	toSerialize["start"] = o.Start
	toSerialize["timezone"] = o.Timezone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsDowntimeTimeSlotRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration   *int64                                       `json:"duration"`
		Name       *string                                      `json:"name,omitempty"`
		Recurrence *SyntheticsDowntimeTimeSlotRecurrenceRequest `json:"recurrence,omitempty"`
		Start      *SyntheticsDowntimeTimeSlotDate              `json:"start"`
		Timezone   *string                                      `json:"timezone"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Duration == nil {
		return fmt.Errorf("required field duration missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	if all.Timezone == nil {
		return fmt.Errorf("required field timezone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration", "name", "recurrence", "start", "timezone"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Duration = *all.Duration
	o.Name = all.Name
	if all.Recurrence != nil && all.Recurrence.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Recurrence = all.Recurrence
	if all.Start.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Start = *all.Start
	o.Timezone = *all.Timezone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
