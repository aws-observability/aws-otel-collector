// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookAbsoluteTime Absolute timeframe.
type NotebookAbsoluteTime struct {
	// The end time.
	End time.Time `json:"end"`
	// Indicates whether the timeframe should be shifted to end at the current time.
	Live *bool `json:"live,omitempty"`
	// The start time.
	Start time.Time `json:"start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotebookAbsoluteTime instantiates a new NotebookAbsoluteTime object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookAbsoluteTime(end time.Time, start time.Time) *NotebookAbsoluteTime {
	this := NotebookAbsoluteTime{}
	this.End = end
	this.Start = start
	return &this
}

// NewNotebookAbsoluteTimeWithDefaults instantiates a new NotebookAbsoluteTime object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookAbsoluteTimeWithDefaults() *NotebookAbsoluteTime {
	this := NotebookAbsoluteTime{}
	return &this
}

// GetEnd returns the End field value.
func (o *NotebookAbsoluteTime) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *NotebookAbsoluteTime) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *NotebookAbsoluteTime) SetEnd(v time.Time) {
	o.End = v
}

// GetLive returns the Live field value if set, zero value otherwise.
func (o *NotebookAbsoluteTime) GetLive() bool {
	if o == nil || o.Live == nil {
		var ret bool
		return ret
	}
	return *o.Live
}

// GetLiveOk returns a tuple with the Live field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookAbsoluteTime) GetLiveOk() (*bool, bool) {
	if o == nil || o.Live == nil {
		return nil, false
	}
	return o.Live, true
}

// HasLive returns a boolean if a field has been set.
func (o *NotebookAbsoluteTime) HasLive() bool {
	return o != nil && o.Live != nil
}

// SetLive gets a reference to the given bool and assigns it to the Live field.
func (o *NotebookAbsoluteTime) SetLive(v bool) {
	o.Live = &v
}

// GetStart returns the Start field value.
func (o *NotebookAbsoluteTime) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *NotebookAbsoluteTime) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *NotebookAbsoluteTime) SetStart(v time.Time) {
	o.Start = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookAbsoluteTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End.Nanosecond() == 0 {
		toSerialize["end"] = o.End.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["end"] = o.End.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Live != nil {
		toSerialize["live"] = o.Live
	}
	if o.Start.Nanosecond() == 0 {
		toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookAbsoluteTime) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End   *time.Time `json:"end"`
		Live  *bool      `json:"live,omitempty"`
		Start *time.Time `json:"start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "live", "start"})
	} else {
		return err
	}
	o.End = *all.End
	o.Live = all.Live
	o.Start = *all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
