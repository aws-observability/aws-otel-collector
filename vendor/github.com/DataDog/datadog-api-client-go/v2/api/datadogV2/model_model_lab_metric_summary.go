// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabMetricSummary Summary statistics for a metric recorded during a Model Lab run.
type ModelLabMetricSummary struct {
	// The total number of recorded values.
	Count int64 `json:"count"`
	// The first step at which the metric was recorded.
	FirstStep datadog.NullableInt64 `json:"first_step,omitempty"`
	// The metric name.
	Key string `json:"key"`
	// The last step at which the metric was recorded.
	LastStep datadog.NullableInt64 `json:"last_step,omitempty"`
	// The most recently recorded value.
	Latest datadog.NullableFloat64 `json:"latest,omitempty"`
	// The maximum recorded value.
	Max datadog.NullableFloat64 `json:"max,omitempty"`
	// The mean of recorded values.
	Mean datadog.NullableFloat64 `json:"mean,omitempty"`
	// The minimum recorded value.
	Min datadog.NullableFloat64 `json:"min,omitempty"`
	// The standard deviation of recorded values.
	Stddev datadog.NullableFloat64 `json:"stddev,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabMetricSummary instantiates a new ModelLabMetricSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabMetricSummary(count int64, key string) *ModelLabMetricSummary {
	this := ModelLabMetricSummary{}
	this.Count = count
	this.Key = key
	return &this
}

// NewModelLabMetricSummaryWithDefaults instantiates a new ModelLabMetricSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabMetricSummaryWithDefaults() *ModelLabMetricSummary {
	this := ModelLabMetricSummary{}
	return &this
}

// GetCount returns the Count field value.
func (o *ModelLabMetricSummary) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ModelLabMetricSummary) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *ModelLabMetricSummary) SetCount(v int64) {
	o.Count = v
}

// GetFirstStep returns the FirstStep field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabMetricSummary) GetFirstStep() int64 {
	if o == nil || o.FirstStep.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FirstStep.Get()
}

// GetFirstStepOk returns a tuple with the FirstStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabMetricSummary) GetFirstStepOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstStep.Get(), o.FirstStep.IsSet()
}

// HasFirstStep returns a boolean if a field has been set.
func (o *ModelLabMetricSummary) HasFirstStep() bool {
	return o != nil && o.FirstStep.IsSet()
}

// SetFirstStep gets a reference to the given datadog.NullableInt64 and assigns it to the FirstStep field.
func (o *ModelLabMetricSummary) SetFirstStep(v int64) {
	o.FirstStep.Set(&v)
}

// SetFirstStepNil sets the value for FirstStep to be an explicit nil.
func (o *ModelLabMetricSummary) SetFirstStepNil() {
	o.FirstStep.Set(nil)
}

// UnsetFirstStep ensures that no value is present for FirstStep, not even an explicit nil.
func (o *ModelLabMetricSummary) UnsetFirstStep() {
	o.FirstStep.Unset()
}

// GetKey returns the Key field value.
func (o *ModelLabMetricSummary) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ModelLabMetricSummary) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *ModelLabMetricSummary) SetKey(v string) {
	o.Key = v
}

// GetLastStep returns the LastStep field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabMetricSummary) GetLastStep() int64 {
	if o == nil || o.LastStep.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastStep.Get()
}

// GetLastStepOk returns a tuple with the LastStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabMetricSummary) GetLastStepOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastStep.Get(), o.LastStep.IsSet()
}

// HasLastStep returns a boolean if a field has been set.
func (o *ModelLabMetricSummary) HasLastStep() bool {
	return o != nil && o.LastStep.IsSet()
}

// SetLastStep gets a reference to the given datadog.NullableInt64 and assigns it to the LastStep field.
func (o *ModelLabMetricSummary) SetLastStep(v int64) {
	o.LastStep.Set(&v)
}

// SetLastStepNil sets the value for LastStep to be an explicit nil.
func (o *ModelLabMetricSummary) SetLastStepNil() {
	o.LastStep.Set(nil)
}

// UnsetLastStep ensures that no value is present for LastStep, not even an explicit nil.
func (o *ModelLabMetricSummary) UnsetLastStep() {
	o.LastStep.Unset()
}

// GetLatest returns the Latest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabMetricSummary) GetLatest() float64 {
	if o == nil || o.Latest.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Latest.Get()
}

// GetLatestOk returns a tuple with the Latest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabMetricSummary) GetLatestOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latest.Get(), o.Latest.IsSet()
}

// HasLatest returns a boolean if a field has been set.
func (o *ModelLabMetricSummary) HasLatest() bool {
	return o != nil && o.Latest.IsSet()
}

// SetLatest gets a reference to the given datadog.NullableFloat64 and assigns it to the Latest field.
func (o *ModelLabMetricSummary) SetLatest(v float64) {
	o.Latest.Set(&v)
}

// SetLatestNil sets the value for Latest to be an explicit nil.
func (o *ModelLabMetricSummary) SetLatestNil() {
	o.Latest.Set(nil)
}

// UnsetLatest ensures that no value is present for Latest, not even an explicit nil.
func (o *ModelLabMetricSummary) UnsetLatest() {
	o.Latest.Unset()
}

// GetMax returns the Max field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabMetricSummary) GetMax() float64 {
	if o == nil || o.Max.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Max.Get()
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabMetricSummary) GetMaxOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Max.Get(), o.Max.IsSet()
}

// HasMax returns a boolean if a field has been set.
func (o *ModelLabMetricSummary) HasMax() bool {
	return o != nil && o.Max.IsSet()
}

// SetMax gets a reference to the given datadog.NullableFloat64 and assigns it to the Max field.
func (o *ModelLabMetricSummary) SetMax(v float64) {
	o.Max.Set(&v)
}

// SetMaxNil sets the value for Max to be an explicit nil.
func (o *ModelLabMetricSummary) SetMaxNil() {
	o.Max.Set(nil)
}

// UnsetMax ensures that no value is present for Max, not even an explicit nil.
func (o *ModelLabMetricSummary) UnsetMax() {
	o.Max.Unset()
}

// GetMean returns the Mean field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabMetricSummary) GetMean() float64 {
	if o == nil || o.Mean.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Mean.Get()
}

// GetMeanOk returns a tuple with the Mean field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabMetricSummary) GetMeanOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mean.Get(), o.Mean.IsSet()
}

// HasMean returns a boolean if a field has been set.
func (o *ModelLabMetricSummary) HasMean() bool {
	return o != nil && o.Mean.IsSet()
}

// SetMean gets a reference to the given datadog.NullableFloat64 and assigns it to the Mean field.
func (o *ModelLabMetricSummary) SetMean(v float64) {
	o.Mean.Set(&v)
}

// SetMeanNil sets the value for Mean to be an explicit nil.
func (o *ModelLabMetricSummary) SetMeanNil() {
	o.Mean.Set(nil)
}

// UnsetMean ensures that no value is present for Mean, not even an explicit nil.
func (o *ModelLabMetricSummary) UnsetMean() {
	o.Mean.Unset()
}

// GetMin returns the Min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabMetricSummary) GetMin() float64 {
	if o == nil || o.Min.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Min.Get()
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabMetricSummary) GetMinOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Min.Get(), o.Min.IsSet()
}

// HasMin returns a boolean if a field has been set.
func (o *ModelLabMetricSummary) HasMin() bool {
	return o != nil && o.Min.IsSet()
}

// SetMin gets a reference to the given datadog.NullableFloat64 and assigns it to the Min field.
func (o *ModelLabMetricSummary) SetMin(v float64) {
	o.Min.Set(&v)
}

// SetMinNil sets the value for Min to be an explicit nil.
func (o *ModelLabMetricSummary) SetMinNil() {
	o.Min.Set(nil)
}

// UnsetMin ensures that no value is present for Min, not even an explicit nil.
func (o *ModelLabMetricSummary) UnsetMin() {
	o.Min.Unset()
}

// GetStddev returns the Stddev field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabMetricSummary) GetStddev() float64 {
	if o == nil || o.Stddev.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Stddev.Get()
}

// GetStddevOk returns a tuple with the Stddev field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabMetricSummary) GetStddevOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stddev.Get(), o.Stddev.IsSet()
}

// HasStddev returns a boolean if a field has been set.
func (o *ModelLabMetricSummary) HasStddev() bool {
	return o != nil && o.Stddev.IsSet()
}

// SetStddev gets a reference to the given datadog.NullableFloat64 and assigns it to the Stddev field.
func (o *ModelLabMetricSummary) SetStddev(v float64) {
	o.Stddev.Set(&v)
}

// SetStddevNil sets the value for Stddev to be an explicit nil.
func (o *ModelLabMetricSummary) SetStddevNil() {
	o.Stddev.Set(nil)
}

// UnsetStddev ensures that no value is present for Stddev, not even an explicit nil.
func (o *ModelLabMetricSummary) UnsetStddev() {
	o.Stddev.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabMetricSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count
	if o.FirstStep.IsSet() {
		toSerialize["first_step"] = o.FirstStep.Get()
	}
	toSerialize["key"] = o.Key
	if o.LastStep.IsSet() {
		toSerialize["last_step"] = o.LastStep.Get()
	}
	if o.Latest.IsSet() {
		toSerialize["latest"] = o.Latest.Get()
	}
	if o.Max.IsSet() {
		toSerialize["max"] = o.Max.Get()
	}
	if o.Mean.IsSet() {
		toSerialize["mean"] = o.Mean.Get()
	}
	if o.Min.IsSet() {
		toSerialize["min"] = o.Min.Get()
	}
	if o.Stddev.IsSet() {
		toSerialize["stddev"] = o.Stddev.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabMetricSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count     *int64                  `json:"count"`
		FirstStep datadog.NullableInt64   `json:"first_step,omitempty"`
		Key       *string                 `json:"key"`
		LastStep  datadog.NullableInt64   `json:"last_step,omitempty"`
		Latest    datadog.NullableFloat64 `json:"latest,omitempty"`
		Max       datadog.NullableFloat64 `json:"max,omitempty"`
		Mean      datadog.NullableFloat64 `json:"mean,omitempty"`
		Min       datadog.NullableFloat64 `json:"min,omitempty"`
		Stddev    datadog.NullableFloat64 `json:"stddev,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "first_step", "key", "last_step", "latest", "max", "mean", "min", "stddev"})
	} else {
		return err
	}
	o.Count = *all.Count
	o.FirstStep = all.FirstStep
	o.Key = *all.Key
	o.LastStep = all.LastStep
	o.Latest = all.Latest
	o.Max = all.Max
	o.Mean = all.Mean
	o.Min = all.Min
	o.Stddev = all.Stddev

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
