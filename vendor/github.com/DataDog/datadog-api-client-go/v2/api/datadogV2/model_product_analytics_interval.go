// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsInterval An interval definition in a timeseries response.
type ProductAnalyticsInterval struct {
	// The duration of each time bucket in milliseconds.
	Milliseconds *int64 `json:"milliseconds,omitempty"`
	// The start of this interval as an epoch timestamp in milliseconds.
	StartTime *int64 `json:"start_time,omitempty"`
	// Epoch timestamps (in milliseconds) for each bucket in this interval.
	Times []int64 `json:"times,omitempty"`
	// The interval type (e.g., fixed or auto-computed bucket size).
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsInterval instantiates a new ProductAnalyticsInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsInterval() *ProductAnalyticsInterval {
	this := ProductAnalyticsInterval{}
	return &this
}

// NewProductAnalyticsIntervalWithDefaults instantiates a new ProductAnalyticsInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsIntervalWithDefaults() *ProductAnalyticsInterval {
	this := ProductAnalyticsInterval{}
	return &this
}

// GetMilliseconds returns the Milliseconds field value if set, zero value otherwise.
func (o *ProductAnalyticsInterval) GetMilliseconds() int64 {
	if o == nil || o.Milliseconds == nil {
		var ret int64
		return ret
	}
	return *o.Milliseconds
}

// GetMillisecondsOk returns a tuple with the Milliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsInterval) GetMillisecondsOk() (*int64, bool) {
	if o == nil || o.Milliseconds == nil {
		return nil, false
	}
	return o.Milliseconds, true
}

// HasMilliseconds returns a boolean if a field has been set.
func (o *ProductAnalyticsInterval) HasMilliseconds() bool {
	return o != nil && o.Milliseconds != nil
}

// SetMilliseconds gets a reference to the given int64 and assigns it to the Milliseconds field.
func (o *ProductAnalyticsInterval) SetMilliseconds(v int64) {
	o.Milliseconds = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ProductAnalyticsInterval) GetStartTime() int64 {
	if o == nil || o.StartTime == nil {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsInterval) GetStartTimeOk() (*int64, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ProductAnalyticsInterval) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *ProductAnalyticsInterval) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetTimes returns the Times field value if set, zero value otherwise.
func (o *ProductAnalyticsInterval) GetTimes() []int64 {
	if o == nil || o.Times == nil {
		var ret []int64
		return ret
	}
	return o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsInterval) GetTimesOk() (*[]int64, bool) {
	if o == nil || o.Times == nil {
		return nil, false
	}
	return &o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *ProductAnalyticsInterval) HasTimes() bool {
	return o != nil && o.Times != nil
}

// SetTimes gets a reference to the given []int64 and assigns it to the Times field.
func (o *ProductAnalyticsInterval) SetTimes(v []int64) {
	o.Times = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProductAnalyticsInterval) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsInterval) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProductAnalyticsInterval) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProductAnalyticsInterval) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Milliseconds != nil {
		toSerialize["milliseconds"] = o.Milliseconds
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.Times != nil {
		toSerialize["times"] = o.Times
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Milliseconds *int64  `json:"milliseconds,omitempty"`
		StartTime    *int64  `json:"start_time,omitempty"`
		Times        []int64 `json:"times,omitempty"`
		Type         *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"milliseconds", "start_time", "times", "type"})
	} else {
		return err
	}
	o.Milliseconds = all.Milliseconds
	o.StartTime = all.StartTime
	o.Times = all.Times
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
