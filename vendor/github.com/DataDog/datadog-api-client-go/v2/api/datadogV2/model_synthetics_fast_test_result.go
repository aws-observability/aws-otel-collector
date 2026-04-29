// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsFastTestResult Fast test result response. Returns `null` if the result is not yet available
// (the test is still running or timed out before completing).
type SyntheticsFastTestResult struct {
	// Fast test result data object (JSON:API format).
	Data *SyntheticsFastTestResultData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsFastTestResult instantiates a new SyntheticsFastTestResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsFastTestResult() *SyntheticsFastTestResult {
	this := SyntheticsFastTestResult{}
	return &this
}

// NewSyntheticsFastTestResultWithDefaults instantiates a new SyntheticsFastTestResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsFastTestResultWithDefaults() *SyntheticsFastTestResult {
	this := SyntheticsFastTestResult{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SyntheticsFastTestResult) GetData() SyntheticsFastTestResultData {
	if o == nil || o.Data == nil {
		var ret SyntheticsFastTestResultData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResult) GetDataOk() (*SyntheticsFastTestResultData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SyntheticsFastTestResult) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given SyntheticsFastTestResultData and assigns it to the Data field.
func (o *SyntheticsFastTestResult) SetData(v SyntheticsFastTestResultData) {
	o.Data = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsFastTestResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsFastTestResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *SyntheticsFastTestResultData `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableSyntheticsFastTestResult handles when a null is used for SyntheticsFastTestResult.
type NullableSyntheticsFastTestResult struct {
	value *SyntheticsFastTestResult
	isSet bool
}

// Get returns the associated value.
func (v NullableSyntheticsFastTestResult) Get() *SyntheticsFastTestResult {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSyntheticsFastTestResult) Set(val *SyntheticsFastTestResult) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSyntheticsFastTestResult) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableSyntheticsFastTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSyntheticsFastTestResult initializes the struct as if Set has been called.
func NewNullableSyntheticsFastTestResult(val *SyntheticsFastTestResult) *NullableSyntheticsFastTestResult {
	return &NullableSyntheticsFastTestResult{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSyntheticsFastTestResult) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSyntheticsFastTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
