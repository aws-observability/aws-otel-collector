// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeScheduleOneTimeCreateUpdateRequest A one-time downtime definition.
type DowntimeScheduleOneTimeCreateUpdateRequest struct {
	// ISO-8601 Datetime to end the downtime. Must include a UTC offset of zero. If not provided, the
	// downtime continues forever.
	End datadog.NullableTime `json:"end,omitempty"`
	// ISO-8601 Datetime to start the downtime. Must include a UTC offset of zero. If not provided, the
	// downtime starts the moment it is created.
	Start datadog.NullableTime `json:"start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDowntimeScheduleOneTimeCreateUpdateRequest instantiates a new DowntimeScheduleOneTimeCreateUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeScheduleOneTimeCreateUpdateRequest() *DowntimeScheduleOneTimeCreateUpdateRequest {
	this := DowntimeScheduleOneTimeCreateUpdateRequest{}
	return &this
}

// NewDowntimeScheduleOneTimeCreateUpdateRequestWithDefaults instantiates a new DowntimeScheduleOneTimeCreateUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeScheduleOneTimeCreateUpdateRequestWithDefaults() *DowntimeScheduleOneTimeCreateUpdateRequest {
	this := DowntimeScheduleOneTimeCreateUpdateRequest{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) GetEnd() time.Time {
	if o == nil || o.End.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) HasEnd() bool {
	return o != nil && o.End.IsSet()
}

// SetEnd gets a reference to the given datadog.NullableTime and assigns it to the End field.
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) SetEnd(v time.Time) {
	o.End.Set(&v)
}

// SetEndNil sets the value for End to be an explicit nil.
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil.
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) UnsetEnd() {
	o.End.Unset()
}

// GetStart returns the Start field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) GetStart() time.Time {
	if o == nil || o.Start.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// HasStart returns a boolean if a field has been set.
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) HasStart() bool {
	return o != nil && o.Start.IsSet()
}

// SetStart gets a reference to the given datadog.NullableTime and assigns it to the Start field.
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) SetStart(v time.Time) {
	o.Start.Set(&v)
}

// SetStartNil sets the value for Start to be an explicit nil.
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) SetStartNil() {
	o.Start.Set(nil)
}

// UnsetStart ensures that no value is present for Start, not even an explicit nil.
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) UnsetStart() {
	o.Start.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeScheduleOneTimeCreateUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	if o.Start.IsSet() {
		toSerialize["start"] = o.Start.Get()
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeScheduleOneTimeCreateUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End   datadog.NullableTime `json:"end,omitempty"`
		Start datadog.NullableTime `json:"start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.End = all.End
	o.Start = all.Start

	return nil
}
