// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOTimeSliceSpec A time-slice SLI specification.
type SLOTimeSliceSpec struct {
	// The time-slice condition, composed of 3 parts: 1. the metric timeseries query, 2. the comparator,
	// and 3. the threshold. Optionally, a fourth part, the query interval, can be provided.
	TimeSlice SLOTimeSliceCondition `json:"time_slice"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSLOTimeSliceSpec instantiates a new SLOTimeSliceSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOTimeSliceSpec(timeSlice SLOTimeSliceCondition) *SLOTimeSliceSpec {
	this := SLOTimeSliceSpec{}
	this.TimeSlice = timeSlice
	return &this
}

// NewSLOTimeSliceSpecWithDefaults instantiates a new SLOTimeSliceSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOTimeSliceSpecWithDefaults() *SLOTimeSliceSpec {
	this := SLOTimeSliceSpec{}
	return &this
}

// GetTimeSlice returns the TimeSlice field value.
func (o *SLOTimeSliceSpec) GetTimeSlice() SLOTimeSliceCondition {
	if o == nil {
		var ret SLOTimeSliceCondition
		return ret
	}
	return o.TimeSlice
}

// GetTimeSliceOk returns a tuple with the TimeSlice field value
// and a boolean to check if the value has been set.
func (o *SLOTimeSliceSpec) GetTimeSliceOk() (*SLOTimeSliceCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeSlice, true
}

// SetTimeSlice sets field value.
func (o *SLOTimeSliceSpec) SetTimeSlice(v SLOTimeSliceCondition) {
	o.TimeSlice = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOTimeSliceSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["time_slice"] = o.TimeSlice
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOTimeSliceSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TimeSlice *SLOTimeSliceCondition `json:"time_slice"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TimeSlice == nil {
		return fmt.Errorf("required field time_slice missing")
	}

	hasInvalidField := false
	if all.TimeSlice.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeSlice = *all.TimeSlice

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
