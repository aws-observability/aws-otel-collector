// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestPipelineStats CI pipeline related statistics for the flaky test. This information is only available if test runs are associated with CI pipeline events from CI Visibility.
type FlakyTestPipelineStats struct {
	// The number of pipelines that failed due to this test for the past 7 days. This is computed as the sum of failed CI pipeline events associated with test runs where the flaky test failed.
	FailedPipelines datadog.NullableInt64 `json:"failed_pipelines,omitempty"`
	// The total time lost by CI pipelines due to this flaky test in milliseconds. This is computed as the sum of the duration of failed CI pipeline events associated with test runs where the flaky test failed.
	TotalLostTimeMs datadog.NullableInt64 `json:"total_lost_time_ms,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFlakyTestPipelineStats instantiates a new FlakyTestPipelineStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlakyTestPipelineStats() *FlakyTestPipelineStats {
	this := FlakyTestPipelineStats{}
	return &this
}

// NewFlakyTestPipelineStatsWithDefaults instantiates a new FlakyTestPipelineStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlakyTestPipelineStatsWithDefaults() *FlakyTestPipelineStats {
	this := FlakyTestPipelineStats{}
	return &this
}

// GetFailedPipelines returns the FailedPipelines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestPipelineStats) GetFailedPipelines() int64 {
	if o == nil || o.FailedPipelines.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FailedPipelines.Get()
}

// GetFailedPipelinesOk returns a tuple with the FailedPipelines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestPipelineStats) GetFailedPipelinesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedPipelines.Get(), o.FailedPipelines.IsSet()
}

// HasFailedPipelines returns a boolean if a field has been set.
func (o *FlakyTestPipelineStats) HasFailedPipelines() bool {
	return o != nil && o.FailedPipelines.IsSet()
}

// SetFailedPipelines gets a reference to the given datadog.NullableInt64 and assigns it to the FailedPipelines field.
func (o *FlakyTestPipelineStats) SetFailedPipelines(v int64) {
	o.FailedPipelines.Set(&v)
}

// SetFailedPipelinesNil sets the value for FailedPipelines to be an explicit nil.
func (o *FlakyTestPipelineStats) SetFailedPipelinesNil() {
	o.FailedPipelines.Set(nil)
}

// UnsetFailedPipelines ensures that no value is present for FailedPipelines, not even an explicit nil.
func (o *FlakyTestPipelineStats) UnsetFailedPipelines() {
	o.FailedPipelines.Unset()
}

// GetTotalLostTimeMs returns the TotalLostTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestPipelineStats) GetTotalLostTimeMs() int64 {
	if o == nil || o.TotalLostTimeMs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalLostTimeMs.Get()
}

// GetTotalLostTimeMsOk returns a tuple with the TotalLostTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestPipelineStats) GetTotalLostTimeMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalLostTimeMs.Get(), o.TotalLostTimeMs.IsSet()
}

// HasTotalLostTimeMs returns a boolean if a field has been set.
func (o *FlakyTestPipelineStats) HasTotalLostTimeMs() bool {
	return o != nil && o.TotalLostTimeMs.IsSet()
}

// SetTotalLostTimeMs gets a reference to the given datadog.NullableInt64 and assigns it to the TotalLostTimeMs field.
func (o *FlakyTestPipelineStats) SetTotalLostTimeMs(v int64) {
	o.TotalLostTimeMs.Set(&v)
}

// SetTotalLostTimeMsNil sets the value for TotalLostTimeMs to be an explicit nil.
func (o *FlakyTestPipelineStats) SetTotalLostTimeMsNil() {
	o.TotalLostTimeMs.Set(nil)
}

// UnsetTotalLostTimeMs ensures that no value is present for TotalLostTimeMs, not even an explicit nil.
func (o *FlakyTestPipelineStats) UnsetTotalLostTimeMs() {
	o.TotalLostTimeMs.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o FlakyTestPipelineStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FailedPipelines.IsSet() {
		toSerialize["failed_pipelines"] = o.FailedPipelines.Get()
	}
	if o.TotalLostTimeMs.IsSet() {
		toSerialize["total_lost_time_ms"] = o.TotalLostTimeMs.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FlakyTestPipelineStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FailedPipelines datadog.NullableInt64 `json:"failed_pipelines,omitempty"`
		TotalLostTimeMs datadog.NullableInt64 `json:"total_lost_time_ms,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"failed_pipelines", "total_lost_time_ms"})
	} else {
		return err
	}
	o.FailedPipelines = all.FailedPipelines
	o.TotalLostTimeMs = all.TotalLostTimeMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
