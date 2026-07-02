// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopLongTaskInvoker A top long task invoker within an invoker type.
type TopLongTaskInvoker struct {
	// Number of sampled views where this invoker had long tasks contributing to the criteria metric.
	CriteriaViewOccurrences *int32 `json:"criteria_view_occurrences,omitempty"`
	// Cleaned source file path for the invoker script.
	File datadog.NullableString `json:"file"`
	// Rank-product impact score combining view frequency and blocking time severity.
	ImpactScore *float64 `json:"impact_score,omitempty"`
	// Name of the invoker function or script.
	Invoker string `json:"invoker"`
	// Statistical distributions of long task metrics computed per view across sampled views.
	StatsPerView LongTaskStatsPerView `json:"stats_per_view"`
	// Number of sampled views where this invoker had any long tasks.
	ViewOccurrences int32 `json:"view_occurrences"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTopLongTaskInvoker instantiates a new TopLongTaskInvoker object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopLongTaskInvoker(file datadog.NullableString, invoker string, statsPerView LongTaskStatsPerView, viewOccurrences int32) *TopLongTaskInvoker {
	this := TopLongTaskInvoker{}
	this.File = file
	this.Invoker = invoker
	this.StatsPerView = statsPerView
	this.ViewOccurrences = viewOccurrences
	return &this
}

// NewTopLongTaskInvokerWithDefaults instantiates a new TopLongTaskInvoker object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopLongTaskInvokerWithDefaults() *TopLongTaskInvoker {
	this := TopLongTaskInvoker{}
	return &this
}

// GetCriteriaViewOccurrences returns the CriteriaViewOccurrences field value if set, zero value otherwise.
func (o *TopLongTaskInvoker) GetCriteriaViewOccurrences() int32 {
	if o == nil || o.CriteriaViewOccurrences == nil {
		var ret int32
		return ret
	}
	return *o.CriteriaViewOccurrences
}

// GetCriteriaViewOccurrencesOk returns a tuple with the CriteriaViewOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopLongTaskInvoker) GetCriteriaViewOccurrencesOk() (*int32, bool) {
	if o == nil || o.CriteriaViewOccurrences == nil {
		return nil, false
	}
	return o.CriteriaViewOccurrences, true
}

// HasCriteriaViewOccurrences returns a boolean if a field has been set.
func (o *TopLongTaskInvoker) HasCriteriaViewOccurrences() bool {
	return o != nil && o.CriteriaViewOccurrences != nil
}

// SetCriteriaViewOccurrences gets a reference to the given int32 and assigns it to the CriteriaViewOccurrences field.
func (o *TopLongTaskInvoker) SetCriteriaViewOccurrences(v int32) {
	o.CriteriaViewOccurrences = &v
}

// GetFile returns the File field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *TopLongTaskInvoker) GetFile() string {
	if o == nil || o.File.Get() == nil {
		var ret string
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TopLongTaskInvoker) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// SetFile sets field value.
func (o *TopLongTaskInvoker) SetFile(v string) {
	o.File.Set(&v)
}

// GetImpactScore returns the ImpactScore field value if set, zero value otherwise.
func (o *TopLongTaskInvoker) GetImpactScore() float64 {
	if o == nil || o.ImpactScore == nil {
		var ret float64
		return ret
	}
	return *o.ImpactScore
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopLongTaskInvoker) GetImpactScoreOk() (*float64, bool) {
	if o == nil || o.ImpactScore == nil {
		return nil, false
	}
	return o.ImpactScore, true
}

// HasImpactScore returns a boolean if a field has been set.
func (o *TopLongTaskInvoker) HasImpactScore() bool {
	return o != nil && o.ImpactScore != nil
}

// SetImpactScore gets a reference to the given float64 and assigns it to the ImpactScore field.
func (o *TopLongTaskInvoker) SetImpactScore(v float64) {
	o.ImpactScore = &v
}

// GetInvoker returns the Invoker field value.
func (o *TopLongTaskInvoker) GetInvoker() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Invoker
}

// GetInvokerOk returns a tuple with the Invoker field value
// and a boolean to check if the value has been set.
func (o *TopLongTaskInvoker) GetInvokerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invoker, true
}

// SetInvoker sets field value.
func (o *TopLongTaskInvoker) SetInvoker(v string) {
	o.Invoker = v
}

// GetStatsPerView returns the StatsPerView field value.
func (o *TopLongTaskInvoker) GetStatsPerView() LongTaskStatsPerView {
	if o == nil {
		var ret LongTaskStatsPerView
		return ret
	}
	return o.StatsPerView
}

// GetStatsPerViewOk returns a tuple with the StatsPerView field value
// and a boolean to check if the value has been set.
func (o *TopLongTaskInvoker) GetStatsPerViewOk() (*LongTaskStatsPerView, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatsPerView, true
}

// SetStatsPerView sets field value.
func (o *TopLongTaskInvoker) SetStatsPerView(v LongTaskStatsPerView) {
	o.StatsPerView = v
}

// GetViewOccurrences returns the ViewOccurrences field value.
func (o *TopLongTaskInvoker) GetViewOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewOccurrences
}

// GetViewOccurrencesOk returns a tuple with the ViewOccurrences field value
// and a boolean to check if the value has been set.
func (o *TopLongTaskInvoker) GetViewOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewOccurrences, true
}

// SetViewOccurrences sets field value.
func (o *TopLongTaskInvoker) SetViewOccurrences(v int32) {
	o.ViewOccurrences = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopLongTaskInvoker) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CriteriaViewOccurrences != nil {
		toSerialize["criteria_view_occurrences"] = o.CriteriaViewOccurrences
	}
	toSerialize["file"] = o.File.Get()
	if o.ImpactScore != nil {
		toSerialize["impact_score"] = o.ImpactScore
	}
	toSerialize["invoker"] = o.Invoker
	toSerialize["stats_per_view"] = o.StatsPerView
	toSerialize["view_occurrences"] = o.ViewOccurrences

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopLongTaskInvoker) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CriteriaViewOccurrences *int32                 `json:"criteria_view_occurrences,omitempty"`
		File                    datadog.NullableString `json:"file"`
		ImpactScore             *float64               `json:"impact_score,omitempty"`
		Invoker                 *string                `json:"invoker"`
		StatsPerView            *LongTaskStatsPerView  `json:"stats_per_view"`
		ViewOccurrences         *int32                 `json:"view_occurrences"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.File.IsSet() {
		return fmt.Errorf("required field file missing")
	}
	if all.Invoker == nil {
		return fmt.Errorf("required field invoker missing")
	}
	if all.StatsPerView == nil {
		return fmt.Errorf("required field stats_per_view missing")
	}
	if all.ViewOccurrences == nil {
		return fmt.Errorf("required field view_occurrences missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"criteria_view_occurrences", "file", "impact_score", "invoker", "stats_per_view", "view_occurrences"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CriteriaViewOccurrences = all.CriteriaViewOccurrences
	o.File = all.File
	o.ImpactScore = all.ImpactScore
	o.Invoker = *all.Invoker
	if all.StatsPerView.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatsPerView = *all.StatsPerView
	o.ViewOccurrences = *all.ViewOccurrences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
