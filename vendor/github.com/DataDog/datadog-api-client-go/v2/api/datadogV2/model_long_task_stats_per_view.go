// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LongTaskStatsPerView Statistical distributions of long task metrics computed per view across sampled views.
type LongTaskStatsPerView struct {
	// Statistical distribution (average, min, max) of a long task metric across sampled views.
	FcpBlockingTimeMs *LongTaskMetricStats `json:"fcp_blocking_time_ms,omitempty"`
	// Statistical distribution (average, min, max) of a long task metric across sampled views.
	FcpCount *LongTaskMetricStats `json:"fcp_count,omitempty"`
	// Statistical distribution (average, min, max) of a long task metric across sampled views.
	InpOverlapBlockingTimeMs *LongTaskMetricStats `json:"inp_overlap_blocking_time_ms,omitempty"`
	// Statistical distribution (average, min, max) of a long task metric across sampled views.
	InpOverlapCount *LongTaskMetricStats `json:"inp_overlap_count,omitempty"`
	// Statistical distribution (average, min, max) of a long task metric across sampled views.
	LcpBlockingTimeMs *LongTaskMetricStats `json:"lcp_blocking_time_ms,omitempty"`
	// Statistical distribution (average, min, max) of a long task metric across sampled views.
	LcpCount *LongTaskMetricStats `json:"lcp_count,omitempty"`
	// Statistical distribution (average, min, max) of a long task metric across sampled views.
	LoadingTimeBlockingTimeMs *LongTaskMetricStats `json:"loading_time_blocking_time_ms,omitempty"`
	// Statistical distribution (average, min, max) of a long task metric across sampled views.
	LoadingTimeCount *LongTaskMetricStats `json:"loading_time_count,omitempty"`
	// Statistical distribution (average, min, max) of a long task metric across sampled views.
	TotalBlockingTimeMs *LongTaskMetricStats `json:"total_blocking_time_ms,omitempty"`
	// Statistical distribution (average, min, max) of a long task metric across sampled views.
	TotalCount *LongTaskMetricStats `json:"total_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLongTaskStatsPerView instantiates a new LongTaskStatsPerView object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLongTaskStatsPerView() *LongTaskStatsPerView {
	this := LongTaskStatsPerView{}
	return &this
}

// NewLongTaskStatsPerViewWithDefaults instantiates a new LongTaskStatsPerView object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLongTaskStatsPerViewWithDefaults() *LongTaskStatsPerView {
	this := LongTaskStatsPerView{}
	return &this
}

// GetFcpBlockingTimeMs returns the FcpBlockingTimeMs field value if set, zero value otherwise.
func (o *LongTaskStatsPerView) GetFcpBlockingTimeMs() LongTaskMetricStats {
	if o == nil || o.FcpBlockingTimeMs == nil {
		var ret LongTaskMetricStats
		return ret
	}
	return *o.FcpBlockingTimeMs
}

// GetFcpBlockingTimeMsOk returns a tuple with the FcpBlockingTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongTaskStatsPerView) GetFcpBlockingTimeMsOk() (*LongTaskMetricStats, bool) {
	if o == nil || o.FcpBlockingTimeMs == nil {
		return nil, false
	}
	return o.FcpBlockingTimeMs, true
}

// HasFcpBlockingTimeMs returns a boolean if a field has been set.
func (o *LongTaskStatsPerView) HasFcpBlockingTimeMs() bool {
	return o != nil && o.FcpBlockingTimeMs != nil
}

// SetFcpBlockingTimeMs gets a reference to the given LongTaskMetricStats and assigns it to the FcpBlockingTimeMs field.
func (o *LongTaskStatsPerView) SetFcpBlockingTimeMs(v LongTaskMetricStats) {
	o.FcpBlockingTimeMs = &v
}

// GetFcpCount returns the FcpCount field value if set, zero value otherwise.
func (o *LongTaskStatsPerView) GetFcpCount() LongTaskMetricStats {
	if o == nil || o.FcpCount == nil {
		var ret LongTaskMetricStats
		return ret
	}
	return *o.FcpCount
}

// GetFcpCountOk returns a tuple with the FcpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongTaskStatsPerView) GetFcpCountOk() (*LongTaskMetricStats, bool) {
	if o == nil || o.FcpCount == nil {
		return nil, false
	}
	return o.FcpCount, true
}

// HasFcpCount returns a boolean if a field has been set.
func (o *LongTaskStatsPerView) HasFcpCount() bool {
	return o != nil && o.FcpCount != nil
}

// SetFcpCount gets a reference to the given LongTaskMetricStats and assigns it to the FcpCount field.
func (o *LongTaskStatsPerView) SetFcpCount(v LongTaskMetricStats) {
	o.FcpCount = &v
}

// GetInpOverlapBlockingTimeMs returns the InpOverlapBlockingTimeMs field value if set, zero value otherwise.
func (o *LongTaskStatsPerView) GetInpOverlapBlockingTimeMs() LongTaskMetricStats {
	if o == nil || o.InpOverlapBlockingTimeMs == nil {
		var ret LongTaskMetricStats
		return ret
	}
	return *o.InpOverlapBlockingTimeMs
}

// GetInpOverlapBlockingTimeMsOk returns a tuple with the InpOverlapBlockingTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongTaskStatsPerView) GetInpOverlapBlockingTimeMsOk() (*LongTaskMetricStats, bool) {
	if o == nil || o.InpOverlapBlockingTimeMs == nil {
		return nil, false
	}
	return o.InpOverlapBlockingTimeMs, true
}

// HasInpOverlapBlockingTimeMs returns a boolean if a field has been set.
func (o *LongTaskStatsPerView) HasInpOverlapBlockingTimeMs() bool {
	return o != nil && o.InpOverlapBlockingTimeMs != nil
}

// SetInpOverlapBlockingTimeMs gets a reference to the given LongTaskMetricStats and assigns it to the InpOverlapBlockingTimeMs field.
func (o *LongTaskStatsPerView) SetInpOverlapBlockingTimeMs(v LongTaskMetricStats) {
	o.InpOverlapBlockingTimeMs = &v
}

// GetInpOverlapCount returns the InpOverlapCount field value if set, zero value otherwise.
func (o *LongTaskStatsPerView) GetInpOverlapCount() LongTaskMetricStats {
	if o == nil || o.InpOverlapCount == nil {
		var ret LongTaskMetricStats
		return ret
	}
	return *o.InpOverlapCount
}

// GetInpOverlapCountOk returns a tuple with the InpOverlapCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongTaskStatsPerView) GetInpOverlapCountOk() (*LongTaskMetricStats, bool) {
	if o == nil || o.InpOverlapCount == nil {
		return nil, false
	}
	return o.InpOverlapCount, true
}

// HasInpOverlapCount returns a boolean if a field has been set.
func (o *LongTaskStatsPerView) HasInpOverlapCount() bool {
	return o != nil && o.InpOverlapCount != nil
}

// SetInpOverlapCount gets a reference to the given LongTaskMetricStats and assigns it to the InpOverlapCount field.
func (o *LongTaskStatsPerView) SetInpOverlapCount(v LongTaskMetricStats) {
	o.InpOverlapCount = &v
}

// GetLcpBlockingTimeMs returns the LcpBlockingTimeMs field value if set, zero value otherwise.
func (o *LongTaskStatsPerView) GetLcpBlockingTimeMs() LongTaskMetricStats {
	if o == nil || o.LcpBlockingTimeMs == nil {
		var ret LongTaskMetricStats
		return ret
	}
	return *o.LcpBlockingTimeMs
}

// GetLcpBlockingTimeMsOk returns a tuple with the LcpBlockingTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongTaskStatsPerView) GetLcpBlockingTimeMsOk() (*LongTaskMetricStats, bool) {
	if o == nil || o.LcpBlockingTimeMs == nil {
		return nil, false
	}
	return o.LcpBlockingTimeMs, true
}

// HasLcpBlockingTimeMs returns a boolean if a field has been set.
func (o *LongTaskStatsPerView) HasLcpBlockingTimeMs() bool {
	return o != nil && o.LcpBlockingTimeMs != nil
}

// SetLcpBlockingTimeMs gets a reference to the given LongTaskMetricStats and assigns it to the LcpBlockingTimeMs field.
func (o *LongTaskStatsPerView) SetLcpBlockingTimeMs(v LongTaskMetricStats) {
	o.LcpBlockingTimeMs = &v
}

// GetLcpCount returns the LcpCount field value if set, zero value otherwise.
func (o *LongTaskStatsPerView) GetLcpCount() LongTaskMetricStats {
	if o == nil || o.LcpCount == nil {
		var ret LongTaskMetricStats
		return ret
	}
	return *o.LcpCount
}

// GetLcpCountOk returns a tuple with the LcpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongTaskStatsPerView) GetLcpCountOk() (*LongTaskMetricStats, bool) {
	if o == nil || o.LcpCount == nil {
		return nil, false
	}
	return o.LcpCount, true
}

// HasLcpCount returns a boolean if a field has been set.
func (o *LongTaskStatsPerView) HasLcpCount() bool {
	return o != nil && o.LcpCount != nil
}

// SetLcpCount gets a reference to the given LongTaskMetricStats and assigns it to the LcpCount field.
func (o *LongTaskStatsPerView) SetLcpCount(v LongTaskMetricStats) {
	o.LcpCount = &v
}

// GetLoadingTimeBlockingTimeMs returns the LoadingTimeBlockingTimeMs field value if set, zero value otherwise.
func (o *LongTaskStatsPerView) GetLoadingTimeBlockingTimeMs() LongTaskMetricStats {
	if o == nil || o.LoadingTimeBlockingTimeMs == nil {
		var ret LongTaskMetricStats
		return ret
	}
	return *o.LoadingTimeBlockingTimeMs
}

// GetLoadingTimeBlockingTimeMsOk returns a tuple with the LoadingTimeBlockingTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongTaskStatsPerView) GetLoadingTimeBlockingTimeMsOk() (*LongTaskMetricStats, bool) {
	if o == nil || o.LoadingTimeBlockingTimeMs == nil {
		return nil, false
	}
	return o.LoadingTimeBlockingTimeMs, true
}

// HasLoadingTimeBlockingTimeMs returns a boolean if a field has been set.
func (o *LongTaskStatsPerView) HasLoadingTimeBlockingTimeMs() bool {
	return o != nil && o.LoadingTimeBlockingTimeMs != nil
}

// SetLoadingTimeBlockingTimeMs gets a reference to the given LongTaskMetricStats and assigns it to the LoadingTimeBlockingTimeMs field.
func (o *LongTaskStatsPerView) SetLoadingTimeBlockingTimeMs(v LongTaskMetricStats) {
	o.LoadingTimeBlockingTimeMs = &v
}

// GetLoadingTimeCount returns the LoadingTimeCount field value if set, zero value otherwise.
func (o *LongTaskStatsPerView) GetLoadingTimeCount() LongTaskMetricStats {
	if o == nil || o.LoadingTimeCount == nil {
		var ret LongTaskMetricStats
		return ret
	}
	return *o.LoadingTimeCount
}

// GetLoadingTimeCountOk returns a tuple with the LoadingTimeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongTaskStatsPerView) GetLoadingTimeCountOk() (*LongTaskMetricStats, bool) {
	if o == nil || o.LoadingTimeCount == nil {
		return nil, false
	}
	return o.LoadingTimeCount, true
}

// HasLoadingTimeCount returns a boolean if a field has been set.
func (o *LongTaskStatsPerView) HasLoadingTimeCount() bool {
	return o != nil && o.LoadingTimeCount != nil
}

// SetLoadingTimeCount gets a reference to the given LongTaskMetricStats and assigns it to the LoadingTimeCount field.
func (o *LongTaskStatsPerView) SetLoadingTimeCount(v LongTaskMetricStats) {
	o.LoadingTimeCount = &v
}

// GetTotalBlockingTimeMs returns the TotalBlockingTimeMs field value if set, zero value otherwise.
func (o *LongTaskStatsPerView) GetTotalBlockingTimeMs() LongTaskMetricStats {
	if o == nil || o.TotalBlockingTimeMs == nil {
		var ret LongTaskMetricStats
		return ret
	}
	return *o.TotalBlockingTimeMs
}

// GetTotalBlockingTimeMsOk returns a tuple with the TotalBlockingTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongTaskStatsPerView) GetTotalBlockingTimeMsOk() (*LongTaskMetricStats, bool) {
	if o == nil || o.TotalBlockingTimeMs == nil {
		return nil, false
	}
	return o.TotalBlockingTimeMs, true
}

// HasTotalBlockingTimeMs returns a boolean if a field has been set.
func (o *LongTaskStatsPerView) HasTotalBlockingTimeMs() bool {
	return o != nil && o.TotalBlockingTimeMs != nil
}

// SetTotalBlockingTimeMs gets a reference to the given LongTaskMetricStats and assigns it to the TotalBlockingTimeMs field.
func (o *LongTaskStatsPerView) SetTotalBlockingTimeMs(v LongTaskMetricStats) {
	o.TotalBlockingTimeMs = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *LongTaskStatsPerView) GetTotalCount() LongTaskMetricStats {
	if o == nil || o.TotalCount == nil {
		var ret LongTaskMetricStats
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongTaskStatsPerView) GetTotalCountOk() (*LongTaskMetricStats, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *LongTaskStatsPerView) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given LongTaskMetricStats and assigns it to the TotalCount field.
func (o *LongTaskStatsPerView) SetTotalCount(v LongTaskMetricStats) {
	o.TotalCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LongTaskStatsPerView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FcpBlockingTimeMs != nil {
		toSerialize["fcp_blocking_time_ms"] = o.FcpBlockingTimeMs
	}
	if o.FcpCount != nil {
		toSerialize["fcp_count"] = o.FcpCount
	}
	if o.InpOverlapBlockingTimeMs != nil {
		toSerialize["inp_overlap_blocking_time_ms"] = o.InpOverlapBlockingTimeMs
	}
	if o.InpOverlapCount != nil {
		toSerialize["inp_overlap_count"] = o.InpOverlapCount
	}
	if o.LcpBlockingTimeMs != nil {
		toSerialize["lcp_blocking_time_ms"] = o.LcpBlockingTimeMs
	}
	if o.LcpCount != nil {
		toSerialize["lcp_count"] = o.LcpCount
	}
	if o.LoadingTimeBlockingTimeMs != nil {
		toSerialize["loading_time_blocking_time_ms"] = o.LoadingTimeBlockingTimeMs
	}
	if o.LoadingTimeCount != nil {
		toSerialize["loading_time_count"] = o.LoadingTimeCount
	}
	if o.TotalBlockingTimeMs != nil {
		toSerialize["total_blocking_time_ms"] = o.TotalBlockingTimeMs
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LongTaskStatsPerView) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FcpBlockingTimeMs         *LongTaskMetricStats `json:"fcp_blocking_time_ms,omitempty"`
		FcpCount                  *LongTaskMetricStats `json:"fcp_count,omitempty"`
		InpOverlapBlockingTimeMs  *LongTaskMetricStats `json:"inp_overlap_blocking_time_ms,omitempty"`
		InpOverlapCount           *LongTaskMetricStats `json:"inp_overlap_count,omitempty"`
		LcpBlockingTimeMs         *LongTaskMetricStats `json:"lcp_blocking_time_ms,omitempty"`
		LcpCount                  *LongTaskMetricStats `json:"lcp_count,omitempty"`
		LoadingTimeBlockingTimeMs *LongTaskMetricStats `json:"loading_time_blocking_time_ms,omitempty"`
		LoadingTimeCount          *LongTaskMetricStats `json:"loading_time_count,omitempty"`
		TotalBlockingTimeMs       *LongTaskMetricStats `json:"total_blocking_time_ms,omitempty"`
		TotalCount                *LongTaskMetricStats `json:"total_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fcp_blocking_time_ms", "fcp_count", "inp_overlap_blocking_time_ms", "inp_overlap_count", "lcp_blocking_time_ms", "lcp_count", "loading_time_blocking_time_ms", "loading_time_count", "total_blocking_time_ms", "total_count"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.FcpBlockingTimeMs != nil && all.FcpBlockingTimeMs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FcpBlockingTimeMs = all.FcpBlockingTimeMs
	if all.FcpCount != nil && all.FcpCount.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FcpCount = all.FcpCount
	if all.InpOverlapBlockingTimeMs != nil && all.InpOverlapBlockingTimeMs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.InpOverlapBlockingTimeMs = all.InpOverlapBlockingTimeMs
	if all.InpOverlapCount != nil && all.InpOverlapCount.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.InpOverlapCount = all.InpOverlapCount
	if all.LcpBlockingTimeMs != nil && all.LcpBlockingTimeMs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LcpBlockingTimeMs = all.LcpBlockingTimeMs
	if all.LcpCount != nil && all.LcpCount.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LcpCount = all.LcpCount
	if all.LoadingTimeBlockingTimeMs != nil && all.LoadingTimeBlockingTimeMs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LoadingTimeBlockingTimeMs = all.LoadingTimeBlockingTimeMs
	if all.LoadingTimeCount != nil && all.LoadingTimeCount.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LoadingTimeCount = all.LoadingTimeCount
	if all.TotalBlockingTimeMs != nil && all.TotalBlockingTimeMs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TotalBlockingTimeMs = all.TotalBlockingTimeMs
	if all.TotalCount != nil && all.TotalCount.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TotalCount = all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
