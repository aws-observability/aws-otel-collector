// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedHighFrozenFrameRate Aggregated high frozen frame rate detection at view level.
type AggregatedHighFrozenFrameRate struct {
	// Average frozen frame rate as a fraction of total frames.
	AvgFrozenFrameRate float64 `json:"avg_frozen_frame_rate"`
	// Average segment duration in nanoseconds.
	AvgSegmentDuration int64 `json:"avg_segment_duration"`
	// Average total frozen duration in nanoseconds.
	AvgTotalFrozenDuration int64 `json:"avg_total_frozen_duration"`
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

// NewAggregatedHighFrozenFrameRate instantiates a new AggregatedHighFrozenFrameRate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedHighFrozenFrameRate(avgFrozenFrameRate float64, avgSegmentDuration int64, avgTotalFrozenDuration int64, fingerprint string, impactScore float64, viewOccurrences int32) *AggregatedHighFrozenFrameRate {
	this := AggregatedHighFrozenFrameRate{}
	this.AvgFrozenFrameRate = avgFrozenFrameRate
	this.AvgSegmentDuration = avgSegmentDuration
	this.AvgTotalFrozenDuration = avgTotalFrozenDuration
	this.Fingerprint = fingerprint
	this.ImpactScore = impactScore
	this.ViewOccurrences = viewOccurrences
	return &this
}

// NewAggregatedHighFrozenFrameRateWithDefaults instantiates a new AggregatedHighFrozenFrameRate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedHighFrozenFrameRateWithDefaults() *AggregatedHighFrozenFrameRate {
	this := AggregatedHighFrozenFrameRate{}
	return &this
}

// GetAvgFrozenFrameRate returns the AvgFrozenFrameRate field value.
func (o *AggregatedHighFrozenFrameRate) GetAvgFrozenFrameRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AvgFrozenFrameRate
}

// GetAvgFrozenFrameRateOk returns a tuple with the AvgFrozenFrameRate field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighFrozenFrameRate) GetAvgFrozenFrameRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgFrozenFrameRate, true
}

// SetAvgFrozenFrameRate sets field value.
func (o *AggregatedHighFrozenFrameRate) SetAvgFrozenFrameRate(v float64) {
	o.AvgFrozenFrameRate = v
}

// GetAvgSegmentDuration returns the AvgSegmentDuration field value.
func (o *AggregatedHighFrozenFrameRate) GetAvgSegmentDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgSegmentDuration
}

// GetAvgSegmentDurationOk returns a tuple with the AvgSegmentDuration field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighFrozenFrameRate) GetAvgSegmentDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgSegmentDuration, true
}

// SetAvgSegmentDuration sets field value.
func (o *AggregatedHighFrozenFrameRate) SetAvgSegmentDuration(v int64) {
	o.AvgSegmentDuration = v
}

// GetAvgTotalFrozenDuration returns the AvgTotalFrozenDuration field value.
func (o *AggregatedHighFrozenFrameRate) GetAvgTotalFrozenDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgTotalFrozenDuration
}

// GetAvgTotalFrozenDurationOk returns a tuple with the AvgTotalFrozenDuration field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighFrozenFrameRate) GetAvgTotalFrozenDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgTotalFrozenDuration, true
}

// SetAvgTotalFrozenDuration sets field value.
func (o *AggregatedHighFrozenFrameRate) SetAvgTotalFrozenDuration(v int64) {
	o.AvgTotalFrozenDuration = v
}

// GetFingerprint returns the Fingerprint field value.
func (o *AggregatedHighFrozenFrameRate) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighFrozenFrameRate) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *AggregatedHighFrozenFrameRate) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetImpactScore returns the ImpactScore field value.
func (o *AggregatedHighFrozenFrameRate) GetImpactScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ImpactScore
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighFrozenFrameRate) GetImpactScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImpactScore, true
}

// SetImpactScore sets field value.
func (o *AggregatedHighFrozenFrameRate) SetImpactScore(v float64) {
	o.ImpactScore = v
}

// GetViewOccurrences returns the ViewOccurrences field value.
func (o *AggregatedHighFrozenFrameRate) GetViewOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewOccurrences
}

// GetViewOccurrencesOk returns a tuple with the ViewOccurrences field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighFrozenFrameRate) GetViewOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewOccurrences, true
}

// SetViewOccurrences sets field value.
func (o *AggregatedHighFrozenFrameRate) SetViewOccurrences(v int32) {
	o.ViewOccurrences = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedHighFrozenFrameRate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["avg_frozen_frame_rate"] = o.AvgFrozenFrameRate
	toSerialize["avg_segment_duration"] = o.AvgSegmentDuration
	toSerialize["avg_total_frozen_duration"] = o.AvgTotalFrozenDuration
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["impact_score"] = o.ImpactScore
	toSerialize["view_occurrences"] = o.ViewOccurrences

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedHighFrozenFrameRate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvgFrozenFrameRate     *float64 `json:"avg_frozen_frame_rate"`
		AvgSegmentDuration     *int64   `json:"avg_segment_duration"`
		AvgTotalFrozenDuration *int64   `json:"avg_total_frozen_duration"`
		Fingerprint            *string  `json:"fingerprint"`
		ImpactScore            *float64 `json:"impact_score"`
		ViewOccurrences        *int32   `json:"view_occurrences"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AvgFrozenFrameRate == nil {
		return fmt.Errorf("required field avg_frozen_frame_rate missing")
	}
	if all.AvgSegmentDuration == nil {
		return fmt.Errorf("required field avg_segment_duration missing")
	}
	if all.AvgTotalFrozenDuration == nil {
		return fmt.Errorf("required field avg_total_frozen_duration missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"avg_frozen_frame_rate", "avg_segment_duration", "avg_total_frozen_duration", "fingerprint", "impact_score", "view_occurrences"})
	} else {
		return err
	}
	o.AvgFrozenFrameRate = *all.AvgFrozenFrameRate
	o.AvgSegmentDuration = *all.AvgSegmentDuration
	o.AvgTotalFrozenDuration = *all.AvgTotalFrozenDuration
	o.Fingerprint = *all.Fingerprint
	o.ImpactScore = *all.ImpactScore
	o.ViewOccurrences = *all.ViewOccurrences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
