// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedMobileScrollFriction Aggregated mobile scroll friction detection at view level.
type AggregatedMobileScrollFriction struct {
	// Average number of frozen frames during scroll interactions.
	AvgScrollFrozenFrameCount int32 `json:"avg_scroll_frozen_frame_count"`
	// Unique fingerprint identifying this detection group.
	Fingerprint string `json:"fingerprint"`
	// Impact score for this detection.
	ImpactScore float64 `json:"impact_score"`
	// Number of sampled views where this detection occurred.
	ViewOccurrences int32 `json:"view_occurrences"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedMobileScrollFriction instantiates a new AggregatedMobileScrollFriction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedMobileScrollFriction(avgScrollFrozenFrameCount int32, fingerprint string, impactScore float64, viewOccurrences int32) *AggregatedMobileScrollFriction {
	this := AggregatedMobileScrollFriction{}
	this.AvgScrollFrozenFrameCount = avgScrollFrozenFrameCount
	this.Fingerprint = fingerprint
	this.ImpactScore = impactScore
	this.ViewOccurrences = viewOccurrences
	return &this
}

// NewAggregatedMobileScrollFrictionWithDefaults instantiates a new AggregatedMobileScrollFriction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedMobileScrollFrictionWithDefaults() *AggregatedMobileScrollFriction {
	this := AggregatedMobileScrollFriction{}
	return &this
}

// GetAvgScrollFrozenFrameCount returns the AvgScrollFrozenFrameCount field value.
func (o *AggregatedMobileScrollFriction) GetAvgScrollFrozenFrameCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.AvgScrollFrozenFrameCount
}

// GetAvgScrollFrozenFrameCountOk returns a tuple with the AvgScrollFrozenFrameCount field value
// and a boolean to check if the value has been set.
func (o *AggregatedMobileScrollFriction) GetAvgScrollFrozenFrameCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgScrollFrozenFrameCount, true
}

// SetAvgScrollFrozenFrameCount sets field value.
func (o *AggregatedMobileScrollFriction) SetAvgScrollFrozenFrameCount(v int32) {
	o.AvgScrollFrozenFrameCount = v
}

// GetFingerprint returns the Fingerprint field value.
func (o *AggregatedMobileScrollFriction) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *AggregatedMobileScrollFriction) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *AggregatedMobileScrollFriction) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetImpactScore returns the ImpactScore field value.
func (o *AggregatedMobileScrollFriction) GetImpactScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ImpactScore
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value
// and a boolean to check if the value has been set.
func (o *AggregatedMobileScrollFriction) GetImpactScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImpactScore, true
}

// SetImpactScore sets field value.
func (o *AggregatedMobileScrollFriction) SetImpactScore(v float64) {
	o.ImpactScore = v
}

// GetViewOccurrences returns the ViewOccurrences field value.
func (o *AggregatedMobileScrollFriction) GetViewOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewOccurrences
}

// GetViewOccurrencesOk returns a tuple with the ViewOccurrences field value
// and a boolean to check if the value has been set.
func (o *AggregatedMobileScrollFriction) GetViewOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewOccurrences, true
}

// SetViewOccurrences sets field value.
func (o *AggregatedMobileScrollFriction) SetViewOccurrences(v int32) {
	o.ViewOccurrences = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedMobileScrollFriction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["avg_scroll_frozen_frame_count"] = o.AvgScrollFrozenFrameCount
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["impact_score"] = o.ImpactScore
	toSerialize["view_occurrences"] = o.ViewOccurrences

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedMobileScrollFriction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvgScrollFrozenFrameCount *int32   `json:"avg_scroll_frozen_frame_count"`
		Fingerprint               *string  `json:"fingerprint"`
		ImpactScore               *float64 `json:"impact_score"`
		ViewOccurrences           *int32   `json:"view_occurrences"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AvgScrollFrozenFrameCount == nil {
		return fmt.Errorf("required field avg_scroll_frozen_frame_count missing")
	}
	if all.Fingerprint == nil {
		return fmt.Errorf("required field fingerprint missing")
	}
	if all.ImpactScore == nil {
		return fmt.Errorf("required field impact_score missing")
	}
	if all.ViewOccurrences == nil {
		return fmt.Errorf("required field view_occurrences missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avg_scroll_frozen_frame_count", "fingerprint", "impact_score", "view_occurrences"})
	} else {
		return err
	}
	o.AvgScrollFrozenFrameCount = *all.AvgScrollFrozenFrameCount
	o.Fingerprint = *all.Fingerprint
	o.ImpactScore = *all.ImpactScore
	o.ViewOccurrences = *all.ViewOccurrences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
