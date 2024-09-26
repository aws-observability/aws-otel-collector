// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardUpdateRequestGlobalTime Timeframe setting for the shared dashboard.
type SharedDashboardUpdateRequestGlobalTime struct {
	// Dashboard global time live_span selection
	LiveSpan *DashboardGlobalTimeLiveSpan `json:"live_span,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSharedDashboardUpdateRequestGlobalTime instantiates a new SharedDashboardUpdateRequestGlobalTime object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboardUpdateRequestGlobalTime() *SharedDashboardUpdateRequestGlobalTime {
	this := SharedDashboardUpdateRequestGlobalTime{}
	return &this
}

// NewSharedDashboardUpdateRequestGlobalTimeWithDefaults instantiates a new SharedDashboardUpdateRequestGlobalTime object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardUpdateRequestGlobalTimeWithDefaults() *SharedDashboardUpdateRequestGlobalTime {
	this := SharedDashboardUpdateRequestGlobalTime{}
	return &this
}

// GetLiveSpan returns the LiveSpan field value if set, zero value otherwise.
func (o *SharedDashboardUpdateRequestGlobalTime) GetLiveSpan() DashboardGlobalTimeLiveSpan {
	if o == nil || o.LiveSpan == nil {
		var ret DashboardGlobalTimeLiveSpan
		return ret
	}
	return *o.LiveSpan
}

// GetLiveSpanOk returns a tuple with the LiveSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDashboardUpdateRequestGlobalTime) GetLiveSpanOk() (*DashboardGlobalTimeLiveSpan, bool) {
	if o == nil || o.LiveSpan == nil {
		return nil, false
	}
	return o.LiveSpan, true
}

// HasLiveSpan returns a boolean if a field has been set.
func (o *SharedDashboardUpdateRequestGlobalTime) HasLiveSpan() bool {
	return o != nil && o.LiveSpan != nil
}

// SetLiveSpan gets a reference to the given DashboardGlobalTimeLiveSpan and assigns it to the LiveSpan field.
func (o *SharedDashboardUpdateRequestGlobalTime) SetLiveSpan(v DashboardGlobalTimeLiveSpan) {
	o.LiveSpan = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboardUpdateRequestGlobalTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.LiveSpan != nil {
		toSerialize["live_span"] = o.LiveSpan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboardUpdateRequestGlobalTime) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LiveSpan *DashboardGlobalTimeLiveSpan `json:"live_span,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"live_span"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.LiveSpan != nil && !all.LiveSpan.IsValid() {
		hasInvalidField = true
	} else {
		o.LiveSpan = all.LiveSpan
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableSharedDashboardUpdateRequestGlobalTime handles when a null is used for SharedDashboardUpdateRequestGlobalTime.
type NullableSharedDashboardUpdateRequestGlobalTime struct {
	value *SharedDashboardUpdateRequestGlobalTime
	isSet bool
}

// Get returns the associated value.
func (v NullableSharedDashboardUpdateRequestGlobalTime) Get() *SharedDashboardUpdateRequestGlobalTime {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSharedDashboardUpdateRequestGlobalTime) Set(val *SharedDashboardUpdateRequestGlobalTime) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSharedDashboardUpdateRequestGlobalTime) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableSharedDashboardUpdateRequestGlobalTime) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSharedDashboardUpdateRequestGlobalTime initializes the struct as if Set has been called.
func NewNullableSharedDashboardUpdateRequestGlobalTime(val *SharedDashboardUpdateRequestGlobalTime) *NullableSharedDashboardUpdateRequestGlobalTime {
	return &NullableSharedDashboardUpdateRequestGlobalTime{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSharedDashboardUpdateRequestGlobalTime) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSharedDashboardUpdateRequestGlobalTime) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
