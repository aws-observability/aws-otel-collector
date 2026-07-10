// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaxSessionDurationUpdateAttributes Attributes for the maximum session duration update request.
type MaxSessionDurationUpdateAttributes struct {
	// The maximum session duration, in seconds.
	MaxSessionDuration int64 `json:"max_session_duration"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMaxSessionDurationUpdateAttributes instantiates a new MaxSessionDurationUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMaxSessionDurationUpdateAttributes(maxSessionDuration int64) *MaxSessionDurationUpdateAttributes {
	this := MaxSessionDurationUpdateAttributes{}
	this.MaxSessionDuration = maxSessionDuration
	return &this
}

// NewMaxSessionDurationUpdateAttributesWithDefaults instantiates a new MaxSessionDurationUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMaxSessionDurationUpdateAttributesWithDefaults() *MaxSessionDurationUpdateAttributes {
	this := MaxSessionDurationUpdateAttributes{}
	return &this
}

// GetMaxSessionDuration returns the MaxSessionDuration field value.
func (o *MaxSessionDurationUpdateAttributes) GetMaxSessionDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxSessionDuration
}

// GetMaxSessionDurationOk returns a tuple with the MaxSessionDuration field value
// and a boolean to check if the value has been set.
func (o *MaxSessionDurationUpdateAttributes) GetMaxSessionDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSessionDuration, true
}

// SetMaxSessionDuration sets field value.
func (o *MaxSessionDurationUpdateAttributes) SetMaxSessionDuration(v int64) {
	o.MaxSessionDuration = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MaxSessionDurationUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["max_session_duration"] = o.MaxSessionDuration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MaxSessionDurationUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxSessionDuration *int64 `json:"max_session_duration"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MaxSessionDuration == nil {
		return fmt.Errorf("required field max_session_duration missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max_session_duration"})
	} else {
		return err
	}
	o.MaxSessionDuration = *all.MaxSessionDuration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
