// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LayerAttributesInterval Defines how often the rotation repeats, using a combination of days and optional seconds. Should be at least 1 hour.
type LayerAttributesInterval struct {
	// The number of days in each rotation cycle.
	Days *int32 `json:"days,omitempty"`
	// Any additional seconds for the rotation cycle (up to 30 days).
	Seconds *int64 `json:"seconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLayerAttributesInterval instantiates a new LayerAttributesInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLayerAttributesInterval() *LayerAttributesInterval {
	this := LayerAttributesInterval{}
	return &this
}

// NewLayerAttributesIntervalWithDefaults instantiates a new LayerAttributesInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLayerAttributesIntervalWithDefaults() *LayerAttributesInterval {
	this := LayerAttributesInterval{}
	return &this
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *LayerAttributesInterval) GetDays() int32 {
	if o == nil || o.Days == nil {
		var ret int32
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerAttributesInterval) GetDaysOk() (*int32, bool) {
	if o == nil || o.Days == nil {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *LayerAttributesInterval) HasDays() bool {
	return o != nil && o.Days != nil
}

// SetDays gets a reference to the given int32 and assigns it to the Days field.
func (o *LayerAttributesInterval) SetDays(v int32) {
	o.Days = &v
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *LayerAttributesInterval) GetSeconds() int64 {
	if o == nil || o.Seconds == nil {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerAttributesInterval) GetSecondsOk() (*int64, bool) {
	if o == nil || o.Seconds == nil {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *LayerAttributesInterval) HasSeconds() bool {
	return o != nil && o.Seconds != nil
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *LayerAttributesInterval) SetSeconds(v int64) {
	o.Seconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LayerAttributesInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Days != nil {
		toSerialize["days"] = o.Days
	}
	if o.Seconds != nil {
		toSerialize["seconds"] = o.Seconds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LayerAttributesInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Days    *int32 `json:"days,omitempty"`
		Seconds *int64 `json:"seconds,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"days", "seconds"})
	} else {
		return err
	}
	o.Days = all.Days
	o.Seconds = all.Seconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
