// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SignalsProblemsDetections Grouped detection results by detection type.
type SignalsProblemsDetections struct {
	// Detected high frozen frame rate issues.
	HighFrozenFrameRates []AggregatedHighFrozenFrameRate `json:"high_frozen_frame_rates,omitempty"`
	// Detected high script evaluation issues.
	HighScriptEvaluations []AggregatedHighScriptEval `json:"high_script_evaluations,omitempty"`
	// Detected low cache hit rate issues.
	LowCacheHitRates []AggregatedLowCacheHitRate `json:"low_cache_hit_rates,omitempty"`
	// Detected mobile scroll friction issues.
	MobileScrollFrictions []AggregatedMobileScrollFriction `json:"mobile_scroll_frictions,omitempty"`
	// Detected slow first contentful paint with high byte count issues.
	SlowFcpHighBytes []AggregatedSlowFCPHighBytes `json:"slow_fcp_high_bytes,omitempty"`
	// Detected slow interaction with long task issues.
	SlowInteractionLongTasks []AggregatedSlowInteractionLongTask `json:"slow_interaction_long_tasks,omitempty"`
	// Detected uncompressed resource issues.
	UncompressedResources []AggregatedUncompressedResource `json:"uncompressed_resources,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSignalsProblemsDetections instantiates a new SignalsProblemsDetections object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSignalsProblemsDetections() *SignalsProblemsDetections {
	this := SignalsProblemsDetections{}
	return &this
}

// NewSignalsProblemsDetectionsWithDefaults instantiates a new SignalsProblemsDetections object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSignalsProblemsDetectionsWithDefaults() *SignalsProblemsDetections {
	this := SignalsProblemsDetections{}
	return &this
}

// GetHighFrozenFrameRates returns the HighFrozenFrameRates field value if set, zero value otherwise.
func (o *SignalsProblemsDetections) GetHighFrozenFrameRates() []AggregatedHighFrozenFrameRate {
	if o == nil || o.HighFrozenFrameRates == nil {
		var ret []AggregatedHighFrozenFrameRate
		return ret
	}
	return o.HighFrozenFrameRates
}

// GetHighFrozenFrameRatesOk returns a tuple with the HighFrozenFrameRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsProblemsDetections) GetHighFrozenFrameRatesOk() (*[]AggregatedHighFrozenFrameRate, bool) {
	if o == nil || o.HighFrozenFrameRates == nil {
		return nil, false
	}
	return &o.HighFrozenFrameRates, true
}

// HasHighFrozenFrameRates returns a boolean if a field has been set.
func (o *SignalsProblemsDetections) HasHighFrozenFrameRates() bool {
	return o != nil && o.HighFrozenFrameRates != nil
}

// SetHighFrozenFrameRates gets a reference to the given []AggregatedHighFrozenFrameRate and assigns it to the HighFrozenFrameRates field.
func (o *SignalsProblemsDetections) SetHighFrozenFrameRates(v []AggregatedHighFrozenFrameRate) {
	o.HighFrozenFrameRates = v
}

// GetHighScriptEvaluations returns the HighScriptEvaluations field value if set, zero value otherwise.
func (o *SignalsProblemsDetections) GetHighScriptEvaluations() []AggregatedHighScriptEval {
	if o == nil || o.HighScriptEvaluations == nil {
		var ret []AggregatedHighScriptEval
		return ret
	}
	return o.HighScriptEvaluations
}

// GetHighScriptEvaluationsOk returns a tuple with the HighScriptEvaluations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsProblemsDetections) GetHighScriptEvaluationsOk() (*[]AggregatedHighScriptEval, bool) {
	if o == nil || o.HighScriptEvaluations == nil {
		return nil, false
	}
	return &o.HighScriptEvaluations, true
}

// HasHighScriptEvaluations returns a boolean if a field has been set.
func (o *SignalsProblemsDetections) HasHighScriptEvaluations() bool {
	return o != nil && o.HighScriptEvaluations != nil
}

// SetHighScriptEvaluations gets a reference to the given []AggregatedHighScriptEval and assigns it to the HighScriptEvaluations field.
func (o *SignalsProblemsDetections) SetHighScriptEvaluations(v []AggregatedHighScriptEval) {
	o.HighScriptEvaluations = v
}

// GetLowCacheHitRates returns the LowCacheHitRates field value if set, zero value otherwise.
func (o *SignalsProblemsDetections) GetLowCacheHitRates() []AggregatedLowCacheHitRate {
	if o == nil || o.LowCacheHitRates == nil {
		var ret []AggregatedLowCacheHitRate
		return ret
	}
	return o.LowCacheHitRates
}

// GetLowCacheHitRatesOk returns a tuple with the LowCacheHitRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsProblemsDetections) GetLowCacheHitRatesOk() (*[]AggregatedLowCacheHitRate, bool) {
	if o == nil || o.LowCacheHitRates == nil {
		return nil, false
	}
	return &o.LowCacheHitRates, true
}

// HasLowCacheHitRates returns a boolean if a field has been set.
func (o *SignalsProblemsDetections) HasLowCacheHitRates() bool {
	return o != nil && o.LowCacheHitRates != nil
}

// SetLowCacheHitRates gets a reference to the given []AggregatedLowCacheHitRate and assigns it to the LowCacheHitRates field.
func (o *SignalsProblemsDetections) SetLowCacheHitRates(v []AggregatedLowCacheHitRate) {
	o.LowCacheHitRates = v
}

// GetMobileScrollFrictions returns the MobileScrollFrictions field value if set, zero value otherwise.
func (o *SignalsProblemsDetections) GetMobileScrollFrictions() []AggregatedMobileScrollFriction {
	if o == nil || o.MobileScrollFrictions == nil {
		var ret []AggregatedMobileScrollFriction
		return ret
	}
	return o.MobileScrollFrictions
}

// GetMobileScrollFrictionsOk returns a tuple with the MobileScrollFrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsProblemsDetections) GetMobileScrollFrictionsOk() (*[]AggregatedMobileScrollFriction, bool) {
	if o == nil || o.MobileScrollFrictions == nil {
		return nil, false
	}
	return &o.MobileScrollFrictions, true
}

// HasMobileScrollFrictions returns a boolean if a field has been set.
func (o *SignalsProblemsDetections) HasMobileScrollFrictions() bool {
	return o != nil && o.MobileScrollFrictions != nil
}

// SetMobileScrollFrictions gets a reference to the given []AggregatedMobileScrollFriction and assigns it to the MobileScrollFrictions field.
func (o *SignalsProblemsDetections) SetMobileScrollFrictions(v []AggregatedMobileScrollFriction) {
	o.MobileScrollFrictions = v
}

// GetSlowFcpHighBytes returns the SlowFcpHighBytes field value if set, zero value otherwise.
func (o *SignalsProblemsDetections) GetSlowFcpHighBytes() []AggregatedSlowFCPHighBytes {
	if o == nil || o.SlowFcpHighBytes == nil {
		var ret []AggregatedSlowFCPHighBytes
		return ret
	}
	return o.SlowFcpHighBytes
}

// GetSlowFcpHighBytesOk returns a tuple with the SlowFcpHighBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsProblemsDetections) GetSlowFcpHighBytesOk() (*[]AggregatedSlowFCPHighBytes, bool) {
	if o == nil || o.SlowFcpHighBytes == nil {
		return nil, false
	}
	return &o.SlowFcpHighBytes, true
}

// HasSlowFcpHighBytes returns a boolean if a field has been set.
func (o *SignalsProblemsDetections) HasSlowFcpHighBytes() bool {
	return o != nil && o.SlowFcpHighBytes != nil
}

// SetSlowFcpHighBytes gets a reference to the given []AggregatedSlowFCPHighBytes and assigns it to the SlowFcpHighBytes field.
func (o *SignalsProblemsDetections) SetSlowFcpHighBytes(v []AggregatedSlowFCPHighBytes) {
	o.SlowFcpHighBytes = v
}

// GetSlowInteractionLongTasks returns the SlowInteractionLongTasks field value if set, zero value otherwise.
func (o *SignalsProblemsDetections) GetSlowInteractionLongTasks() []AggregatedSlowInteractionLongTask {
	if o == nil || o.SlowInteractionLongTasks == nil {
		var ret []AggregatedSlowInteractionLongTask
		return ret
	}
	return o.SlowInteractionLongTasks
}

// GetSlowInteractionLongTasksOk returns a tuple with the SlowInteractionLongTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsProblemsDetections) GetSlowInteractionLongTasksOk() (*[]AggregatedSlowInteractionLongTask, bool) {
	if o == nil || o.SlowInteractionLongTasks == nil {
		return nil, false
	}
	return &o.SlowInteractionLongTasks, true
}

// HasSlowInteractionLongTasks returns a boolean if a field has been set.
func (o *SignalsProblemsDetections) HasSlowInteractionLongTasks() bool {
	return o != nil && o.SlowInteractionLongTasks != nil
}

// SetSlowInteractionLongTasks gets a reference to the given []AggregatedSlowInteractionLongTask and assigns it to the SlowInteractionLongTasks field.
func (o *SignalsProblemsDetections) SetSlowInteractionLongTasks(v []AggregatedSlowInteractionLongTask) {
	o.SlowInteractionLongTasks = v
}

// GetUncompressedResources returns the UncompressedResources field value if set, zero value otherwise.
func (o *SignalsProblemsDetections) GetUncompressedResources() []AggregatedUncompressedResource {
	if o == nil || o.UncompressedResources == nil {
		var ret []AggregatedUncompressedResource
		return ret
	}
	return o.UncompressedResources
}

// GetUncompressedResourcesOk returns a tuple with the UncompressedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsProblemsDetections) GetUncompressedResourcesOk() (*[]AggregatedUncompressedResource, bool) {
	if o == nil || o.UncompressedResources == nil {
		return nil, false
	}
	return &o.UncompressedResources, true
}

// HasUncompressedResources returns a boolean if a field has been set.
func (o *SignalsProblemsDetections) HasUncompressedResources() bool {
	return o != nil && o.UncompressedResources != nil
}

// SetUncompressedResources gets a reference to the given []AggregatedUncompressedResource and assigns it to the UncompressedResources field.
func (o *SignalsProblemsDetections) SetUncompressedResources(v []AggregatedUncompressedResource) {
	o.UncompressedResources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SignalsProblemsDetections) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HighFrozenFrameRates != nil {
		toSerialize["high_frozen_frame_rates"] = o.HighFrozenFrameRates
	}
	if o.HighScriptEvaluations != nil {
		toSerialize["high_script_evaluations"] = o.HighScriptEvaluations
	}
	if o.LowCacheHitRates != nil {
		toSerialize["low_cache_hit_rates"] = o.LowCacheHitRates
	}
	if o.MobileScrollFrictions != nil {
		toSerialize["mobile_scroll_frictions"] = o.MobileScrollFrictions
	}
	if o.SlowFcpHighBytes != nil {
		toSerialize["slow_fcp_high_bytes"] = o.SlowFcpHighBytes
	}
	if o.SlowInteractionLongTasks != nil {
		toSerialize["slow_interaction_long_tasks"] = o.SlowInteractionLongTasks
	}
	if o.UncompressedResources != nil {
		toSerialize["uncompressed_resources"] = o.UncompressedResources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SignalsProblemsDetections) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HighFrozenFrameRates     []AggregatedHighFrozenFrameRate     `json:"high_frozen_frame_rates,omitempty"`
		HighScriptEvaluations    []AggregatedHighScriptEval          `json:"high_script_evaluations,omitempty"`
		LowCacheHitRates         []AggregatedLowCacheHitRate         `json:"low_cache_hit_rates,omitempty"`
		MobileScrollFrictions    []AggregatedMobileScrollFriction    `json:"mobile_scroll_frictions,omitempty"`
		SlowFcpHighBytes         []AggregatedSlowFCPHighBytes        `json:"slow_fcp_high_bytes,omitempty"`
		SlowInteractionLongTasks []AggregatedSlowInteractionLongTask `json:"slow_interaction_long_tasks,omitempty"`
		UncompressedResources    []AggregatedUncompressedResource    `json:"uncompressed_resources,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"high_frozen_frame_rates", "high_script_evaluations", "low_cache_hit_rates", "mobile_scroll_frictions", "slow_fcp_high_bytes", "slow_interaction_long_tasks", "uncompressed_resources"})
	} else {
		return err
	}
	o.HighFrozenFrameRates = all.HighFrozenFrameRates
	o.HighScriptEvaluations = all.HighScriptEvaluations
	o.LowCacheHitRates = all.LowCacheHitRates
	o.MobileScrollFrictions = all.MobileScrollFrictions
	o.SlowFcpHighBytes = all.SlowFcpHighBytes
	o.SlowInteractionLongTasks = all.SlowInteractionLongTasks
	o.UncompressedResources = all.UncompressedResources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
