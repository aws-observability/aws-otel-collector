// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedLongTasksByInvokerType Aggregated long task statistics for a single invoker type.
type AggregatedLongTasksByInvokerType struct {
	// Number of sampled views where this invoker type had long tasks contributing to the criteria metric.
	CriteriaViewOccurrences *int32 `json:"criteria_view_occurrences,omitempty"`
	// Rank-product impact score combining view frequency and blocking time severity.
	ImpactScore *float64 `json:"impact_score,omitempty"`
	// Category of the long task invoker (for example, resolve-promise, user-callback).
	InvokerType string `json:"invoker_type"`
	// Statistical distributions of long task metrics computed per view across sampled views.
	StatsPerView LongTaskStatsPerView `json:"stats_per_view"`
	// Top invokers within this invoker type, sorted by impact score descending.
	TopInvokers []TopLongTaskInvoker `json:"top_invokers"`
	// Number of sampled views where this invoker type had any long tasks.
	ViewOccurrences int32 `json:"view_occurrences"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedLongTasksByInvokerType instantiates a new AggregatedLongTasksByInvokerType object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedLongTasksByInvokerType(invokerType string, statsPerView LongTaskStatsPerView, topInvokers []TopLongTaskInvoker, viewOccurrences int32) *AggregatedLongTasksByInvokerType {
	this := AggregatedLongTasksByInvokerType{}
	this.InvokerType = invokerType
	this.StatsPerView = statsPerView
	this.TopInvokers = topInvokers
	this.ViewOccurrences = viewOccurrences
	return &this
}

// NewAggregatedLongTasksByInvokerTypeWithDefaults instantiates a new AggregatedLongTasksByInvokerType object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedLongTasksByInvokerTypeWithDefaults() *AggregatedLongTasksByInvokerType {
	this := AggregatedLongTasksByInvokerType{}
	return &this
}

// GetCriteriaViewOccurrences returns the CriteriaViewOccurrences field value if set, zero value otherwise.
func (o *AggregatedLongTasksByInvokerType) GetCriteriaViewOccurrences() int32 {
	if o == nil || o.CriteriaViewOccurrences == nil {
		var ret int32
		return ret
	}
	return *o.CriteriaViewOccurrences
}

// GetCriteriaViewOccurrencesOk returns a tuple with the CriteriaViewOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksByInvokerType) GetCriteriaViewOccurrencesOk() (*int32, bool) {
	if o == nil || o.CriteriaViewOccurrences == nil {
		return nil, false
	}
	return o.CriteriaViewOccurrences, true
}

// HasCriteriaViewOccurrences returns a boolean if a field has been set.
func (o *AggregatedLongTasksByInvokerType) HasCriteriaViewOccurrences() bool {
	return o != nil && o.CriteriaViewOccurrences != nil
}

// SetCriteriaViewOccurrences gets a reference to the given int32 and assigns it to the CriteriaViewOccurrences field.
func (o *AggregatedLongTasksByInvokerType) SetCriteriaViewOccurrences(v int32) {
	o.CriteriaViewOccurrences = &v
}

// GetImpactScore returns the ImpactScore field value if set, zero value otherwise.
func (o *AggregatedLongTasksByInvokerType) GetImpactScore() float64 {
	if o == nil || o.ImpactScore == nil {
		var ret float64
		return ret
	}
	return *o.ImpactScore
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksByInvokerType) GetImpactScoreOk() (*float64, bool) {
	if o == nil || o.ImpactScore == nil {
		return nil, false
	}
	return o.ImpactScore, true
}

// HasImpactScore returns a boolean if a field has been set.
func (o *AggregatedLongTasksByInvokerType) HasImpactScore() bool {
	return o != nil && o.ImpactScore != nil
}

// SetImpactScore gets a reference to the given float64 and assigns it to the ImpactScore field.
func (o *AggregatedLongTasksByInvokerType) SetImpactScore(v float64) {
	o.ImpactScore = &v
}

// GetInvokerType returns the InvokerType field value.
func (o *AggregatedLongTasksByInvokerType) GetInvokerType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InvokerType
}

// GetInvokerTypeOk returns a tuple with the InvokerType field value
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksByInvokerType) GetInvokerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvokerType, true
}

// SetInvokerType sets field value.
func (o *AggregatedLongTasksByInvokerType) SetInvokerType(v string) {
	o.InvokerType = v
}

// GetStatsPerView returns the StatsPerView field value.
func (o *AggregatedLongTasksByInvokerType) GetStatsPerView() LongTaskStatsPerView {
	if o == nil {
		var ret LongTaskStatsPerView
		return ret
	}
	return o.StatsPerView
}

// GetStatsPerViewOk returns a tuple with the StatsPerView field value
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksByInvokerType) GetStatsPerViewOk() (*LongTaskStatsPerView, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatsPerView, true
}

// SetStatsPerView sets field value.
func (o *AggregatedLongTasksByInvokerType) SetStatsPerView(v LongTaskStatsPerView) {
	o.StatsPerView = v
}

// GetTopInvokers returns the TopInvokers field value.
func (o *AggregatedLongTasksByInvokerType) GetTopInvokers() []TopLongTaskInvoker {
	if o == nil {
		var ret []TopLongTaskInvoker
		return ret
	}
	return o.TopInvokers
}

// GetTopInvokersOk returns a tuple with the TopInvokers field value
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksByInvokerType) GetTopInvokersOk() (*[]TopLongTaskInvoker, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopInvokers, true
}

// SetTopInvokers sets field value.
func (o *AggregatedLongTasksByInvokerType) SetTopInvokers(v []TopLongTaskInvoker) {
	o.TopInvokers = v
}

// GetViewOccurrences returns the ViewOccurrences field value.
func (o *AggregatedLongTasksByInvokerType) GetViewOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewOccurrences
}

// GetViewOccurrencesOk returns a tuple with the ViewOccurrences field value
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksByInvokerType) GetViewOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewOccurrences, true
}

// SetViewOccurrences sets field value.
func (o *AggregatedLongTasksByInvokerType) SetViewOccurrences(v int32) {
	o.ViewOccurrences = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedLongTasksByInvokerType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CriteriaViewOccurrences != nil {
		toSerialize["criteria_view_occurrences"] = o.CriteriaViewOccurrences
	}
	if o.ImpactScore != nil {
		toSerialize["impact_score"] = o.ImpactScore
	}
	toSerialize["invoker_type"] = o.InvokerType
	toSerialize["stats_per_view"] = o.StatsPerView
	toSerialize["top_invokers"] = o.TopInvokers
	toSerialize["view_occurrences"] = o.ViewOccurrences

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedLongTasksByInvokerType) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CriteriaViewOccurrences *int32                `json:"criteria_view_occurrences,omitempty"`
		ImpactScore             *float64              `json:"impact_score,omitempty"`
		InvokerType             *string               `json:"invoker_type"`
		StatsPerView            *LongTaskStatsPerView `json:"stats_per_view"`
		TopInvokers             *[]TopLongTaskInvoker `json:"top_invokers"`
		ViewOccurrences         *int32                `json:"view_occurrences"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InvokerType == nil {
		return fmt.Errorf("required field invoker_type missing")
	}
	if all.StatsPerView == nil {
		return fmt.Errorf("required field stats_per_view missing")
	}
	if all.TopInvokers == nil {
		return fmt.Errorf("required field top_invokers missing")
	}
	if all.ViewOccurrences == nil {
		return fmt.Errorf("required field view_occurrences missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"criteria_view_occurrences", "impact_score", "invoker_type", "stats_per_view", "top_invokers", "view_occurrences"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CriteriaViewOccurrences = all.CriteriaViewOccurrences
	o.ImpactScore = all.ImpactScore
	o.InvokerType = *all.InvokerType
	if all.StatsPerView.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatsPerView = *all.StatsPerView
	o.TopInvokers = *all.TopInvokers
	o.ViewOccurrences = *all.ViewOccurrences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
