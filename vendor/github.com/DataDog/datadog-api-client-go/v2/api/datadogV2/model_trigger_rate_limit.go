// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TriggerRateLimit Defines a rate limit for a trigger.
type TriggerRateLimit struct {
	// The `TriggerRateLimit` `count`.
	Count *int64 `json:"count,omitempty"`
	// The `TriggerRateLimit` `interval`. The expected format is the number of seconds ending with an s. For example, 1 day is 86400s
	Interval *string `json:"interval,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTriggerRateLimit instantiates a new TriggerRateLimit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTriggerRateLimit() *TriggerRateLimit {
	this := TriggerRateLimit{}
	return &this
}

// NewTriggerRateLimitWithDefaults instantiates a new TriggerRateLimit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTriggerRateLimitWithDefaults() *TriggerRateLimit {
	this := TriggerRateLimit{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TriggerRateLimit) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerRateLimit) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TriggerRateLimit) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *TriggerRateLimit) SetCount(v int64) {
	o.Count = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *TriggerRateLimit) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerRateLimit) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *TriggerRateLimit) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *TriggerRateLimit) SetInterval(v string) {
	o.Interval = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TriggerRateLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TriggerRateLimit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count    *int64  `json:"count,omitempty"`
		Interval *string `json:"interval,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "interval"})
	} else {
		return err
	}
	o.Count = all.Count
	o.Interval = all.Interval

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
