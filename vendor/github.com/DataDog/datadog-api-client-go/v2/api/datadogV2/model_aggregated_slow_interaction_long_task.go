// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedSlowInteractionLongTask Aggregated slow interaction with long task detection grouped by action and selector.
type AggregatedSlowInteractionLongTask struct {
	// Type of user interaction that triggered the slow response.
	ActionType string `json:"action_type"`
	// Average long task blocking duration in nanoseconds.
	AvgBlockingDuration int64 `json:"avg_blocking_duration"`
	// Average total interaction duration in nanoseconds.
	AvgDuration int64 `json:"avg_duration"`
	// Unique fingerprint identifying this detection group.
	Fingerprint string `json:"fingerprint"`
	// Impact score combining view frequency and blocking severity.
	ImpactScore float64 `json:"impact_score"`
	// Total number of detection instances across sampled views.
	InstanceCount int32 `json:"instance_count"`
	// CSS selector of the element that was interacted with.
	Selector datadog.NullableString `json:"selector"`
	// Normalized CSS selector with dynamic parts replaced.
	SelectorNormalized datadog.NullableString `json:"selector_normalized"`
	// Number of sampled views where this detection occurred.
	ViewOccurrences int32 `json:"view_occurrences"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedSlowInteractionLongTask instantiates a new AggregatedSlowInteractionLongTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedSlowInteractionLongTask(actionType string, avgBlockingDuration int64, avgDuration int64, fingerprint string, impactScore float64, instanceCount int32, selector datadog.NullableString, selectorNormalized datadog.NullableString, viewOccurrences int32) *AggregatedSlowInteractionLongTask {
	this := AggregatedSlowInteractionLongTask{}
	this.ActionType = actionType
	this.AvgBlockingDuration = avgBlockingDuration
	this.AvgDuration = avgDuration
	this.Fingerprint = fingerprint
	this.ImpactScore = impactScore
	this.InstanceCount = instanceCount
	this.Selector = selector
	this.SelectorNormalized = selectorNormalized
	this.ViewOccurrences = viewOccurrences
	return &this
}

// NewAggregatedSlowInteractionLongTaskWithDefaults instantiates a new AggregatedSlowInteractionLongTask object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedSlowInteractionLongTaskWithDefaults() *AggregatedSlowInteractionLongTask {
	this := AggregatedSlowInteractionLongTask{}
	return &this
}

// GetActionType returns the ActionType field value.
func (o *AggregatedSlowInteractionLongTask) GetActionType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowInteractionLongTask) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value.
func (o *AggregatedSlowInteractionLongTask) SetActionType(v string) {
	o.ActionType = v
}

// GetAvgBlockingDuration returns the AvgBlockingDuration field value.
func (o *AggregatedSlowInteractionLongTask) GetAvgBlockingDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgBlockingDuration
}

// GetAvgBlockingDurationOk returns a tuple with the AvgBlockingDuration field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowInteractionLongTask) GetAvgBlockingDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgBlockingDuration, true
}

// SetAvgBlockingDuration sets field value.
func (o *AggregatedSlowInteractionLongTask) SetAvgBlockingDuration(v int64) {
	o.AvgBlockingDuration = v
}

// GetAvgDuration returns the AvgDuration field value.
func (o *AggregatedSlowInteractionLongTask) GetAvgDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgDuration
}

// GetAvgDurationOk returns a tuple with the AvgDuration field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowInteractionLongTask) GetAvgDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgDuration, true
}

// SetAvgDuration sets field value.
func (o *AggregatedSlowInteractionLongTask) SetAvgDuration(v int64) {
	o.AvgDuration = v
}

// GetFingerprint returns the Fingerprint field value.
func (o *AggregatedSlowInteractionLongTask) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowInteractionLongTask) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *AggregatedSlowInteractionLongTask) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetImpactScore returns the ImpactScore field value.
func (o *AggregatedSlowInteractionLongTask) GetImpactScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ImpactScore
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowInteractionLongTask) GetImpactScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImpactScore, true
}

// SetImpactScore sets field value.
func (o *AggregatedSlowInteractionLongTask) SetImpactScore(v float64) {
	o.ImpactScore = v
}

// GetInstanceCount returns the InstanceCount field value.
func (o *AggregatedSlowInteractionLongTask) GetInstanceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowInteractionLongTask) GetInstanceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceCount, true
}

// SetInstanceCount sets field value.
func (o *AggregatedSlowInteractionLongTask) SetInstanceCount(v int32) {
	o.InstanceCount = v
}

// GetSelector returns the Selector field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *AggregatedSlowInteractionLongTask) GetSelector() string {
	if o == nil || o.Selector.Get() == nil {
		var ret string
		return ret
	}
	return *o.Selector.Get()
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AggregatedSlowInteractionLongTask) GetSelectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Selector.Get(), o.Selector.IsSet()
}

// SetSelector sets field value.
func (o *AggregatedSlowInteractionLongTask) SetSelector(v string) {
	o.Selector.Set(&v)
}

// GetSelectorNormalized returns the SelectorNormalized field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *AggregatedSlowInteractionLongTask) GetSelectorNormalized() string {
	if o == nil || o.SelectorNormalized.Get() == nil {
		var ret string
		return ret
	}
	return *o.SelectorNormalized.Get()
}

// GetSelectorNormalizedOk returns a tuple with the SelectorNormalized field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AggregatedSlowInteractionLongTask) GetSelectorNormalizedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectorNormalized.Get(), o.SelectorNormalized.IsSet()
}

// SetSelectorNormalized sets field value.
func (o *AggregatedSlowInteractionLongTask) SetSelectorNormalized(v string) {
	o.SelectorNormalized.Set(&v)
}

// GetViewOccurrences returns the ViewOccurrences field value.
func (o *AggregatedSlowInteractionLongTask) GetViewOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewOccurrences
}

// GetViewOccurrencesOk returns a tuple with the ViewOccurrences field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowInteractionLongTask) GetViewOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewOccurrences, true
}

// SetViewOccurrences sets field value.
func (o *AggregatedSlowInteractionLongTask) SetViewOccurrences(v int32) {
	o.ViewOccurrences = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedSlowInteractionLongTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action_type"] = o.ActionType
	toSerialize["avg_blocking_duration"] = o.AvgBlockingDuration
	toSerialize["avg_duration"] = o.AvgDuration
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["impact_score"] = o.ImpactScore
	toSerialize["instance_count"] = o.InstanceCount
	toSerialize["selector"] = o.Selector.Get()
	toSerialize["selector_normalized"] = o.SelectorNormalized.Get()
	toSerialize["view_occurrences"] = o.ViewOccurrences

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedSlowInteractionLongTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionType          *string                `json:"action_type"`
		AvgBlockingDuration *int64                 `json:"avg_blocking_duration"`
		AvgDuration         *int64                 `json:"avg_duration"`
		Fingerprint         *string                `json:"fingerprint"`
		ImpactScore         *float64               `json:"impact_score"`
		InstanceCount       *int32                 `json:"instance_count"`
		Selector            datadog.NullableString `json:"selector"`
		SelectorNormalized  datadog.NullableString `json:"selector_normalized"`
		ViewOccurrences     *int32                 `json:"view_occurrences"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActionType == nil {
		return fmt.Errorf("required field action_type missing")
	}
	if all.AvgBlockingDuration == nil {
		return fmt.Errorf("required field avg_blocking_duration missing")
	}
	if all.AvgDuration == nil {
		return fmt.Errorf("required field avg_duration missing")
	}
	if all.Fingerprint == nil {
		return fmt.Errorf("required field fingerprint missing")
	}
	if all.ImpactScore == nil {
		return fmt.Errorf("required field impact_score missing")
	}
	if all.InstanceCount == nil {
		return fmt.Errorf("required field instance_count missing")
	}
	if !all.Selector.IsSet() {
		return fmt.Errorf("required field selector missing")
	}
	if !all.SelectorNormalized.IsSet() {
		return fmt.Errorf("required field selector_normalized missing")
	}
	if all.ViewOccurrences == nil {
		return fmt.Errorf("required field view_occurrences missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action_type", "avg_blocking_duration", "avg_duration", "fingerprint", "impact_score", "instance_count", "selector", "selector_normalized", "view_occurrences"})
	} else {
		return err
	}
	o.ActionType = *all.ActionType
	o.AvgBlockingDuration = *all.AvgBlockingDuration
	o.AvgDuration = *all.AvgDuration
	o.Fingerprint = *all.Fingerprint
	o.ImpactScore = *all.ImpactScore
	o.InstanceCount = *all.InstanceCount
	o.Selector = all.Selector
	o.SelectorNormalized = all.SelectorNormalized
	o.ViewOccurrences = *all.ViewOccurrences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
