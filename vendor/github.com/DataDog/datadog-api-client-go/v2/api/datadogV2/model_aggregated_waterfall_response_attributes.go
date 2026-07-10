// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedWaterfallResponseAttributes Attributes of an aggregated waterfall response.
type AggregatedWaterfallResponseAttributes struct {
	// The RUM application ID that was analyzed.
	ApplicationId string `json:"application_id"`
	// Performance criteria to filter view instances by a metric threshold.
	Criteria *AggregatedWaterfallPerformanceCriteria `json:"criteria,omitempty"`
	// Start of the analyzed time range as a Unix timestamp in seconds.
	From int64 `json:"from"`
	// Network resources in chronological waterfall order.
	Resources []AggregatedResource `json:"resources"`
	// List of RUM view IDs sampled for this aggregation, capped at 50.
	SampledViewIds []string `json:"sampled_view_ids"`
	// End of the analyzed time range as a Unix timestamp in seconds.
	To int64 `json:"to"`
	// Overall cache hit rate across all sampled views.
	TotalCacheHitRatePct float64 `json:"total_cache_hit_rate_pct"`
	// Number of view instances included in the analysis.
	ViewCount int32 `json:"view_count"`
	// The RUM view name that was analyzed.
	ViewName string `json:"view_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedWaterfallResponseAttributes instantiates a new AggregatedWaterfallResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedWaterfallResponseAttributes(applicationId string, from int64, resources []AggregatedResource, sampledViewIds []string, to int64, totalCacheHitRatePct float64, viewCount int32, viewName string) *AggregatedWaterfallResponseAttributes {
	this := AggregatedWaterfallResponseAttributes{}
	this.ApplicationId = applicationId
	this.From = from
	this.Resources = resources
	this.SampledViewIds = sampledViewIds
	this.To = to
	this.TotalCacheHitRatePct = totalCacheHitRatePct
	this.ViewCount = viewCount
	this.ViewName = viewName
	return &this
}

// NewAggregatedWaterfallResponseAttributesWithDefaults instantiates a new AggregatedWaterfallResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedWaterfallResponseAttributesWithDefaults() *AggregatedWaterfallResponseAttributes {
	this := AggregatedWaterfallResponseAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *AggregatedWaterfallResponseAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallResponseAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *AggregatedWaterfallResponseAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *AggregatedWaterfallResponseAttributes) GetCriteria() AggregatedWaterfallPerformanceCriteria {
	if o == nil || o.Criteria == nil {
		var ret AggregatedWaterfallPerformanceCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallResponseAttributes) GetCriteriaOk() (*AggregatedWaterfallPerformanceCriteria, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *AggregatedWaterfallResponseAttributes) HasCriteria() bool {
	return o != nil && o.Criteria != nil
}

// SetCriteria gets a reference to the given AggregatedWaterfallPerformanceCriteria and assigns it to the Criteria field.
func (o *AggregatedWaterfallResponseAttributes) SetCriteria(v AggregatedWaterfallPerformanceCriteria) {
	o.Criteria = &v
}

// GetFrom returns the From field value.
func (o *AggregatedWaterfallResponseAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallResponseAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *AggregatedWaterfallResponseAttributes) SetFrom(v int64) {
	o.From = v
}

// GetResources returns the Resources field value.
func (o *AggregatedWaterfallResponseAttributes) GetResources() []AggregatedResource {
	if o == nil {
		var ret []AggregatedResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallResponseAttributes) GetResourcesOk() (*[]AggregatedResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value.
func (o *AggregatedWaterfallResponseAttributes) SetResources(v []AggregatedResource) {
	o.Resources = v
}

// GetSampledViewIds returns the SampledViewIds field value.
func (o *AggregatedWaterfallResponseAttributes) GetSampledViewIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SampledViewIds
}

// GetSampledViewIdsOk returns a tuple with the SampledViewIds field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallResponseAttributes) GetSampledViewIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampledViewIds, true
}

// SetSampledViewIds sets field value.
func (o *AggregatedWaterfallResponseAttributes) SetSampledViewIds(v []string) {
	o.SampledViewIds = v
}

// GetTo returns the To field value.
func (o *AggregatedWaterfallResponseAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallResponseAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *AggregatedWaterfallResponseAttributes) SetTo(v int64) {
	o.To = v
}

// GetTotalCacheHitRatePct returns the TotalCacheHitRatePct field value.
func (o *AggregatedWaterfallResponseAttributes) GetTotalCacheHitRatePct() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.TotalCacheHitRatePct
}

// GetTotalCacheHitRatePctOk returns a tuple with the TotalCacheHitRatePct field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallResponseAttributes) GetTotalCacheHitRatePctOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCacheHitRatePct, true
}

// SetTotalCacheHitRatePct sets field value.
func (o *AggregatedWaterfallResponseAttributes) SetTotalCacheHitRatePct(v float64) {
	o.TotalCacheHitRatePct = v
}

// GetViewCount returns the ViewCount field value.
func (o *AggregatedWaterfallResponseAttributes) GetViewCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewCount
}

// GetViewCountOk returns a tuple with the ViewCount field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallResponseAttributes) GetViewCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewCount, true
}

// SetViewCount sets field value.
func (o *AggregatedWaterfallResponseAttributes) SetViewCount(v int32) {
	o.ViewCount = v
}

// GetViewName returns the ViewName field value.
func (o *AggregatedWaterfallResponseAttributes) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallResponseAttributes) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value.
func (o *AggregatedWaterfallResponseAttributes) SetViewName(v string) {
	o.ViewName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedWaterfallResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_id"] = o.ApplicationId
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	toSerialize["from"] = o.From
	toSerialize["resources"] = o.Resources
	toSerialize["sampled_view_ids"] = o.SampledViewIds
	toSerialize["to"] = o.To
	toSerialize["total_cache_hit_rate_pct"] = o.TotalCacheHitRatePct
	toSerialize["view_count"] = o.ViewCount
	toSerialize["view_name"] = o.ViewName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedWaterfallResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId        *string                                 `json:"application_id"`
		Criteria             *AggregatedWaterfallPerformanceCriteria `json:"criteria,omitempty"`
		From                 *int64                                  `json:"from"`
		Resources            *[]AggregatedResource                   `json:"resources"`
		SampledViewIds       *[]string                               `json:"sampled_view_ids"`
		To                   *int64                                  `json:"to"`
		TotalCacheHitRatePct *float64                                `json:"total_cache_hit_rate_pct"`
		ViewCount            *int32                                  `json:"view_count"`
		ViewName             *string                                 `json:"view_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Resources == nil {
		return fmt.Errorf("required field resources missing")
	}
	if all.SampledViewIds == nil {
		return fmt.Errorf("required field sampled_view_ids missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	if all.TotalCacheHitRatePct == nil {
		return fmt.Errorf("required field total_cache_hit_rate_pct missing")
	}
	if all.ViewCount == nil {
		return fmt.Errorf("required field view_count missing")
	}
	if all.ViewName == nil {
		return fmt.Errorf("required field view_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "criteria", "from", "resources", "sampled_view_ids", "to", "total_cache_hit_rate_pct", "view_count", "view_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = *all.ApplicationId
	if all.Criteria != nil && all.Criteria.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Criteria = all.Criteria
	o.From = *all.From
	o.Resources = *all.Resources
	o.SampledViewIds = *all.SampledViewIds
	o.To = *all.To
	o.TotalCacheHitRatePct = *all.TotalCacheHitRatePct
	o.ViewCount = *all.ViewCount
	o.ViewName = *all.ViewName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
