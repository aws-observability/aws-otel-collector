// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedResource Aggregated performance statistics for a single network resource across sampled view instances.
type AggregatedResource struct {
	// Average total duration in milliseconds.
	AvgDurationMs float64 `json:"avg_duration_ms"`
	// Average start time relative to view start in milliseconds.
	AvgStartTimeMs float64 `json:"avg_start_time_ms"`
	// Cache hit rate as a percentage.
	CacheHitRatePct float64 `json:"cache_hit_rate_pct"`
	// Number of requests served from cache.
	CachedCount int32 `json:"cached_count"`
	// Number of requests downloaded from the network.
	DownloadedCount int32 `json:"downloaded_count"`
	// 75th percentile duration across all view names in the application, present when include_global_appearance is true.
	GlobalP75DurationMs *float64 `json:"global_p75_duration_ms,omitempty"`
	// Number of distinct view names in the application that load this resource, present when include_global_appearance is true.
	GlobalViewNameCount *int32 `json:"global_view_name_count,omitempty"`
	// Percentage of distinct view names in the application that load this resource, present when include_global_appearance is true.
	GlobalViewNamePct *float64 `json:"global_view_name_pct,omitempty"`
	// HTTP method for the resource request.
	HttpMethod datadog.NullableString `json:"http_method"`
	// Percentage of sampled view instances that loaded this resource.
	LoadFrequencyPct float64 `json:"load_frequency_pct"`
	// Maximum duration in milliseconds.
	MaxDurationMs float64 `json:"max_duration_ms"`
	// Median duration in milliseconds.
	MedianDurationMs float64 `json:"median_duration_ms"`
	// Minimum duration in milliseconds.
	MinDurationMs float64 `json:"min_duration_ms"`
	// 75th percentile duration in milliseconds.
	P75DurationMs float64 `json:"p75_duration_ms"`
	// 95th percentile duration in milliseconds.
	P95DurationMs float64 `json:"p95_duration_ms"`
	// Resource type (JS, CSS, image, fetch, XHR, document, and so on).
	ResourceType datadog.NullableString `json:"resource_type"`
	// URL path group used to aggregate similar resources.
	ResourceUrlPathGroup string `json:"resource_url_path_group"`
	// Average timing breakdown per network phase for a resource.
	TimingBreakdown AggregatedResourceTimingBreakdown `json:"timing_breakdown"`
	// Total number of requests for this resource across all sampled views.
	TotalRequests int32 `json:"total_requests"`
	// Number of sampled view instances that loaded this resource.
	ViewsWithResource int32 `json:"views_with_resource"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedResource instantiates a new AggregatedResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedResource(avgDurationMs float64, avgStartTimeMs float64, cacheHitRatePct float64, cachedCount int32, downloadedCount int32, httpMethod datadog.NullableString, loadFrequencyPct float64, maxDurationMs float64, medianDurationMs float64, minDurationMs float64, p75DurationMs float64, p95DurationMs float64, resourceType datadog.NullableString, resourceUrlPathGroup string, timingBreakdown AggregatedResourceTimingBreakdown, totalRequests int32, viewsWithResource int32) *AggregatedResource {
	this := AggregatedResource{}
	this.AvgDurationMs = avgDurationMs
	this.AvgStartTimeMs = avgStartTimeMs
	this.CacheHitRatePct = cacheHitRatePct
	this.CachedCount = cachedCount
	this.DownloadedCount = downloadedCount
	this.HttpMethod = httpMethod
	this.LoadFrequencyPct = loadFrequencyPct
	this.MaxDurationMs = maxDurationMs
	this.MedianDurationMs = medianDurationMs
	this.MinDurationMs = minDurationMs
	this.P75DurationMs = p75DurationMs
	this.P95DurationMs = p95DurationMs
	this.ResourceType = resourceType
	this.ResourceUrlPathGroup = resourceUrlPathGroup
	this.TimingBreakdown = timingBreakdown
	this.TotalRequests = totalRequests
	this.ViewsWithResource = viewsWithResource
	return &this
}

// NewAggregatedResourceWithDefaults instantiates a new AggregatedResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedResourceWithDefaults() *AggregatedResource {
	this := AggregatedResource{}
	return &this
}

// GetAvgDurationMs returns the AvgDurationMs field value.
func (o *AggregatedResource) GetAvgDurationMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AvgDurationMs
}

// GetAvgDurationMsOk returns a tuple with the AvgDurationMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetAvgDurationMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgDurationMs, true
}

// SetAvgDurationMs sets field value.
func (o *AggregatedResource) SetAvgDurationMs(v float64) {
	o.AvgDurationMs = v
}

// GetAvgStartTimeMs returns the AvgStartTimeMs field value.
func (o *AggregatedResource) GetAvgStartTimeMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AvgStartTimeMs
}

// GetAvgStartTimeMsOk returns a tuple with the AvgStartTimeMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetAvgStartTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgStartTimeMs, true
}

// SetAvgStartTimeMs sets field value.
func (o *AggregatedResource) SetAvgStartTimeMs(v float64) {
	o.AvgStartTimeMs = v
}

// GetCacheHitRatePct returns the CacheHitRatePct field value.
func (o *AggregatedResource) GetCacheHitRatePct() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.CacheHitRatePct
}

// GetCacheHitRatePctOk returns a tuple with the CacheHitRatePct field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetCacheHitRatePctOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheHitRatePct, true
}

// SetCacheHitRatePct sets field value.
func (o *AggregatedResource) SetCacheHitRatePct(v float64) {
	o.CacheHitRatePct = v
}

// GetCachedCount returns the CachedCount field value.
func (o *AggregatedResource) GetCachedCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.CachedCount
}

// GetCachedCountOk returns a tuple with the CachedCount field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetCachedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CachedCount, true
}

// SetCachedCount sets field value.
func (o *AggregatedResource) SetCachedCount(v int32) {
	o.CachedCount = v
}

// GetDownloadedCount returns the DownloadedCount field value.
func (o *AggregatedResource) GetDownloadedCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.DownloadedCount
}

// GetDownloadedCountOk returns a tuple with the DownloadedCount field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetDownloadedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadedCount, true
}

// SetDownloadedCount sets field value.
func (o *AggregatedResource) SetDownloadedCount(v int32) {
	o.DownloadedCount = v
}

// GetGlobalP75DurationMs returns the GlobalP75DurationMs field value if set, zero value otherwise.
func (o *AggregatedResource) GetGlobalP75DurationMs() float64 {
	if o == nil || o.GlobalP75DurationMs == nil {
		var ret float64
		return ret
	}
	return *o.GlobalP75DurationMs
}

// GetGlobalP75DurationMsOk returns a tuple with the GlobalP75DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetGlobalP75DurationMsOk() (*float64, bool) {
	if o == nil || o.GlobalP75DurationMs == nil {
		return nil, false
	}
	return o.GlobalP75DurationMs, true
}

// HasGlobalP75DurationMs returns a boolean if a field has been set.
func (o *AggregatedResource) HasGlobalP75DurationMs() bool {
	return o != nil && o.GlobalP75DurationMs != nil
}

// SetGlobalP75DurationMs gets a reference to the given float64 and assigns it to the GlobalP75DurationMs field.
func (o *AggregatedResource) SetGlobalP75DurationMs(v float64) {
	o.GlobalP75DurationMs = &v
}

// GetGlobalViewNameCount returns the GlobalViewNameCount field value if set, zero value otherwise.
func (o *AggregatedResource) GetGlobalViewNameCount() int32 {
	if o == nil || o.GlobalViewNameCount == nil {
		var ret int32
		return ret
	}
	return *o.GlobalViewNameCount
}

// GetGlobalViewNameCountOk returns a tuple with the GlobalViewNameCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetGlobalViewNameCountOk() (*int32, bool) {
	if o == nil || o.GlobalViewNameCount == nil {
		return nil, false
	}
	return o.GlobalViewNameCount, true
}

// HasGlobalViewNameCount returns a boolean if a field has been set.
func (o *AggregatedResource) HasGlobalViewNameCount() bool {
	return o != nil && o.GlobalViewNameCount != nil
}

// SetGlobalViewNameCount gets a reference to the given int32 and assigns it to the GlobalViewNameCount field.
func (o *AggregatedResource) SetGlobalViewNameCount(v int32) {
	o.GlobalViewNameCount = &v
}

// GetGlobalViewNamePct returns the GlobalViewNamePct field value if set, zero value otherwise.
func (o *AggregatedResource) GetGlobalViewNamePct() float64 {
	if o == nil || o.GlobalViewNamePct == nil {
		var ret float64
		return ret
	}
	return *o.GlobalViewNamePct
}

// GetGlobalViewNamePctOk returns a tuple with the GlobalViewNamePct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetGlobalViewNamePctOk() (*float64, bool) {
	if o == nil || o.GlobalViewNamePct == nil {
		return nil, false
	}
	return o.GlobalViewNamePct, true
}

// HasGlobalViewNamePct returns a boolean if a field has been set.
func (o *AggregatedResource) HasGlobalViewNamePct() bool {
	return o != nil && o.GlobalViewNamePct != nil
}

// SetGlobalViewNamePct gets a reference to the given float64 and assigns it to the GlobalViewNamePct field.
func (o *AggregatedResource) SetGlobalViewNamePct(v float64) {
	o.GlobalViewNamePct = &v
}

// GetHttpMethod returns the HttpMethod field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *AggregatedResource) GetHttpMethod() string {
	if o == nil || o.HttpMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.HttpMethod.Get()
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AggregatedResource) GetHttpMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpMethod.Get(), o.HttpMethod.IsSet()
}

// SetHttpMethod sets field value.
func (o *AggregatedResource) SetHttpMethod(v string) {
	o.HttpMethod.Set(&v)
}

// GetLoadFrequencyPct returns the LoadFrequencyPct field value.
func (o *AggregatedResource) GetLoadFrequencyPct() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.LoadFrequencyPct
}

// GetLoadFrequencyPctOk returns a tuple with the LoadFrequencyPct field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetLoadFrequencyPctOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadFrequencyPct, true
}

// SetLoadFrequencyPct sets field value.
func (o *AggregatedResource) SetLoadFrequencyPct(v float64) {
	o.LoadFrequencyPct = v
}

// GetMaxDurationMs returns the MaxDurationMs field value.
func (o *AggregatedResource) GetMaxDurationMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MaxDurationMs
}

// GetMaxDurationMsOk returns a tuple with the MaxDurationMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetMaxDurationMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxDurationMs, true
}

// SetMaxDurationMs sets field value.
func (o *AggregatedResource) SetMaxDurationMs(v float64) {
	o.MaxDurationMs = v
}

// GetMedianDurationMs returns the MedianDurationMs field value.
func (o *AggregatedResource) GetMedianDurationMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MedianDurationMs
}

// GetMedianDurationMsOk returns a tuple with the MedianDurationMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetMedianDurationMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MedianDurationMs, true
}

// SetMedianDurationMs sets field value.
func (o *AggregatedResource) SetMedianDurationMs(v float64) {
	o.MedianDurationMs = v
}

// GetMinDurationMs returns the MinDurationMs field value.
func (o *AggregatedResource) GetMinDurationMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MinDurationMs
}

// GetMinDurationMsOk returns a tuple with the MinDurationMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetMinDurationMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinDurationMs, true
}

// SetMinDurationMs sets field value.
func (o *AggregatedResource) SetMinDurationMs(v float64) {
	o.MinDurationMs = v
}

// GetP75DurationMs returns the P75DurationMs field value.
func (o *AggregatedResource) GetP75DurationMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.P75DurationMs
}

// GetP75DurationMsOk returns a tuple with the P75DurationMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetP75DurationMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.P75DurationMs, true
}

// SetP75DurationMs sets field value.
func (o *AggregatedResource) SetP75DurationMs(v float64) {
	o.P75DurationMs = v
}

// GetP95DurationMs returns the P95DurationMs field value.
func (o *AggregatedResource) GetP95DurationMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.P95DurationMs
}

// GetP95DurationMsOk returns a tuple with the P95DurationMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetP95DurationMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.P95DurationMs, true
}

// SetP95DurationMs sets field value.
func (o *AggregatedResource) SetP95DurationMs(v float64) {
	o.P95DurationMs = v
}

// GetResourceType returns the ResourceType field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *AggregatedResource) GetResourceType() string {
	if o == nil || o.ResourceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceType.Get()
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AggregatedResource) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType.Get(), o.ResourceType.IsSet()
}

// SetResourceType sets field value.
func (o *AggregatedResource) SetResourceType(v string) {
	o.ResourceType.Set(&v)
}

// GetResourceUrlPathGroup returns the ResourceUrlPathGroup field value.
func (o *AggregatedResource) GetResourceUrlPathGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceUrlPathGroup
}

// GetResourceUrlPathGroupOk returns a tuple with the ResourceUrlPathGroup field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetResourceUrlPathGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceUrlPathGroup, true
}

// SetResourceUrlPathGroup sets field value.
func (o *AggregatedResource) SetResourceUrlPathGroup(v string) {
	o.ResourceUrlPathGroup = v
}

// GetTimingBreakdown returns the TimingBreakdown field value.
func (o *AggregatedResource) GetTimingBreakdown() AggregatedResourceTimingBreakdown {
	if o == nil {
		var ret AggregatedResourceTimingBreakdown
		return ret
	}
	return o.TimingBreakdown
}

// GetTimingBreakdownOk returns a tuple with the TimingBreakdown field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetTimingBreakdownOk() (*AggregatedResourceTimingBreakdown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimingBreakdown, true
}

// SetTimingBreakdown sets field value.
func (o *AggregatedResource) SetTimingBreakdown(v AggregatedResourceTimingBreakdown) {
	o.TimingBreakdown = v
}

// GetTotalRequests returns the TotalRequests field value.
func (o *AggregatedResource) GetTotalRequests() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.TotalRequests
}

// GetTotalRequestsOk returns a tuple with the TotalRequests field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetTotalRequestsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRequests, true
}

// SetTotalRequests sets field value.
func (o *AggregatedResource) SetTotalRequests(v int32) {
	o.TotalRequests = v
}

// GetViewsWithResource returns the ViewsWithResource field value.
func (o *AggregatedResource) GetViewsWithResource() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewsWithResource
}

// GetViewsWithResourceOk returns a tuple with the ViewsWithResource field value
// and a boolean to check if the value has been set.
func (o *AggregatedResource) GetViewsWithResourceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewsWithResource, true
}

// SetViewsWithResource sets field value.
func (o *AggregatedResource) SetViewsWithResource(v int32) {
	o.ViewsWithResource = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["avg_duration_ms"] = o.AvgDurationMs
	toSerialize["avg_start_time_ms"] = o.AvgStartTimeMs
	toSerialize["cache_hit_rate_pct"] = o.CacheHitRatePct
	toSerialize["cached_count"] = o.CachedCount
	toSerialize["downloaded_count"] = o.DownloadedCount
	if o.GlobalP75DurationMs != nil {
		toSerialize["global_p75_duration_ms"] = o.GlobalP75DurationMs
	}
	if o.GlobalViewNameCount != nil {
		toSerialize["global_view_name_count"] = o.GlobalViewNameCount
	}
	if o.GlobalViewNamePct != nil {
		toSerialize["global_view_name_pct"] = o.GlobalViewNamePct
	}
	toSerialize["http_method"] = o.HttpMethod.Get()
	toSerialize["load_frequency_pct"] = o.LoadFrequencyPct
	toSerialize["max_duration_ms"] = o.MaxDurationMs
	toSerialize["median_duration_ms"] = o.MedianDurationMs
	toSerialize["min_duration_ms"] = o.MinDurationMs
	toSerialize["p75_duration_ms"] = o.P75DurationMs
	toSerialize["p95_duration_ms"] = o.P95DurationMs
	toSerialize["resource_type"] = o.ResourceType.Get()
	toSerialize["resource_url_path_group"] = o.ResourceUrlPathGroup
	toSerialize["timing_breakdown"] = o.TimingBreakdown
	toSerialize["total_requests"] = o.TotalRequests
	toSerialize["views_with_resource"] = o.ViewsWithResource

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvgDurationMs        *float64                           `json:"avg_duration_ms"`
		AvgStartTimeMs       *float64                           `json:"avg_start_time_ms"`
		CacheHitRatePct      *float64                           `json:"cache_hit_rate_pct"`
		CachedCount          *int32                             `json:"cached_count"`
		DownloadedCount      *int32                             `json:"downloaded_count"`
		GlobalP75DurationMs  *float64                           `json:"global_p75_duration_ms,omitempty"`
		GlobalViewNameCount  *int32                             `json:"global_view_name_count,omitempty"`
		GlobalViewNamePct    *float64                           `json:"global_view_name_pct,omitempty"`
		HttpMethod           datadog.NullableString             `json:"http_method"`
		LoadFrequencyPct     *float64                           `json:"load_frequency_pct"`
		MaxDurationMs        *float64                           `json:"max_duration_ms"`
		MedianDurationMs     *float64                           `json:"median_duration_ms"`
		MinDurationMs        *float64                           `json:"min_duration_ms"`
		P75DurationMs        *float64                           `json:"p75_duration_ms"`
		P95DurationMs        *float64                           `json:"p95_duration_ms"`
		ResourceType         datadog.NullableString             `json:"resource_type"`
		ResourceUrlPathGroup *string                            `json:"resource_url_path_group"`
		TimingBreakdown      *AggregatedResourceTimingBreakdown `json:"timing_breakdown"`
		TotalRequests        *int32                             `json:"total_requests"`
		ViewsWithResource    *int32                             `json:"views_with_resource"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AvgDurationMs == nil {
		return fmt.Errorf("required field avg_duration_ms missing")
	}
	if all.AvgStartTimeMs == nil {
		return fmt.Errorf("required field avg_start_time_ms missing")
	}
	if all.CacheHitRatePct == nil {
		return fmt.Errorf("required field cache_hit_rate_pct missing")
	}
	if all.CachedCount == nil {
		return fmt.Errorf("required field cached_count missing")
	}
	if all.DownloadedCount == nil {
		return fmt.Errorf("required field downloaded_count missing")
	}
	if !all.HttpMethod.IsSet() {
		return fmt.Errorf("required field http_method missing")
	}
	if all.LoadFrequencyPct == nil {
		return fmt.Errorf("required field load_frequency_pct missing")
	}
	if all.MaxDurationMs == nil {
		return fmt.Errorf("required field max_duration_ms missing")
	}
	if all.MedianDurationMs == nil {
		return fmt.Errorf("required field median_duration_ms missing")
	}
	if all.MinDurationMs == nil {
		return fmt.Errorf("required field min_duration_ms missing")
	}
	if all.P75DurationMs == nil {
		return fmt.Errorf("required field p75_duration_ms missing")
	}
	if all.P95DurationMs == nil {
		return fmt.Errorf("required field p95_duration_ms missing")
	}
	if !all.ResourceType.IsSet() {
		return fmt.Errorf("required field resource_type missing")
	}
	if all.ResourceUrlPathGroup == nil {
		return fmt.Errorf("required field resource_url_path_group missing")
	}
	if all.TimingBreakdown == nil {
		return fmt.Errorf("required field timing_breakdown missing")
	}
	if all.TotalRequests == nil {
		return fmt.Errorf("required field total_requests missing")
	}
	if all.ViewsWithResource == nil {
		return fmt.Errorf("required field views_with_resource missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avg_duration_ms", "avg_start_time_ms", "cache_hit_rate_pct", "cached_count", "downloaded_count", "global_p75_duration_ms", "global_view_name_count", "global_view_name_pct", "http_method", "load_frequency_pct", "max_duration_ms", "median_duration_ms", "min_duration_ms", "p75_duration_ms", "p95_duration_ms", "resource_type", "resource_url_path_group", "timing_breakdown", "total_requests", "views_with_resource"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AvgDurationMs = *all.AvgDurationMs
	o.AvgStartTimeMs = *all.AvgStartTimeMs
	o.CacheHitRatePct = *all.CacheHitRatePct
	o.CachedCount = *all.CachedCount
	o.DownloadedCount = *all.DownloadedCount
	o.GlobalP75DurationMs = all.GlobalP75DurationMs
	o.GlobalViewNameCount = all.GlobalViewNameCount
	o.GlobalViewNamePct = all.GlobalViewNamePct
	o.HttpMethod = all.HttpMethod
	o.LoadFrequencyPct = *all.LoadFrequencyPct
	o.MaxDurationMs = *all.MaxDurationMs
	o.MedianDurationMs = *all.MedianDurationMs
	o.MinDurationMs = *all.MinDurationMs
	o.P75DurationMs = *all.P75DurationMs
	o.P95DurationMs = *all.P95DurationMs
	o.ResourceType = all.ResourceType
	o.ResourceUrlPathGroup = *all.ResourceUrlPathGroup
	if all.TimingBreakdown.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimingBreakdown = *all.TimingBreakdown
	o.TotalRequests = *all.TotalRequests
	o.ViewsWithResource = *all.ViewsWithResource

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
