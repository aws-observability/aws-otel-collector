// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestStats Test statistics for the flaky test.
type FlakyTestStats struct {
	// The failure rate percentage of the test for the past 7 days. This is the number of failed test runs divided by the total number of test runs (excluding skipped test runs).
	FailureRatePct datadog.NullableFloat64 `json:"failure_rate_pct,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFlakyTestStats instantiates a new FlakyTestStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlakyTestStats() *FlakyTestStats {
	this := FlakyTestStats{}
	return &this
}

// NewFlakyTestStatsWithDefaults instantiates a new FlakyTestStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlakyTestStatsWithDefaults() *FlakyTestStats {
	this := FlakyTestStats{}
	return &this
}

// GetFailureRatePct returns the FailureRatePct field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestStats) GetFailureRatePct() float64 {
	if o == nil || o.FailureRatePct.Get() == nil {
		var ret float64
		return ret
	}
	return *o.FailureRatePct.Get()
}

// GetFailureRatePctOk returns a tuple with the FailureRatePct field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestStats) GetFailureRatePctOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureRatePct.Get(), o.FailureRatePct.IsSet()
}

// HasFailureRatePct returns a boolean if a field has been set.
func (o *FlakyTestStats) HasFailureRatePct() bool {
	return o != nil && o.FailureRatePct.IsSet()
}

// SetFailureRatePct gets a reference to the given datadog.NullableFloat64 and assigns it to the FailureRatePct field.
func (o *FlakyTestStats) SetFailureRatePct(v float64) {
	o.FailureRatePct.Set(&v)
}

// SetFailureRatePctNil sets the value for FailureRatePct to be an explicit nil.
func (o *FlakyTestStats) SetFailureRatePctNil() {
	o.FailureRatePct.Set(nil)
}

// UnsetFailureRatePct ensures that no value is present for FailureRatePct, not even an explicit nil.
func (o *FlakyTestStats) UnsetFailureRatePct() {
	o.FailureRatePct.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o FlakyTestStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FailureRatePct.IsSet() {
		toSerialize["failure_rate_pct"] = o.FailureRatePct.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FlakyTestStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FailureRatePct datadog.NullableFloat64 `json:"failure_rate_pct,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"failure_rate_pct"})
	} else {
		return err
	}
	o.FailureRatePct = all.FailureRatePct

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
