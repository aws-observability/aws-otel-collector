// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedLowCacheHitRate Aggregated low cache hit rate detection at view level.
type AggregatedLowCacheHitRate struct {
	// Average cache hit rate across affected views.
	AvgCacheHitRate float64 `json:"avg_cache_hit_rate"`
	// Average total download size of uncached resources in bytes.
	AvgResourceDownloadSizeBytes int64 `json:"avg_resource_download_size_bytes"`
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

// NewAggregatedLowCacheHitRate instantiates a new AggregatedLowCacheHitRate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedLowCacheHitRate(avgCacheHitRate float64, avgResourceDownloadSizeBytes int64, fingerprint string, impactScore float64, viewOccurrences int32) *AggregatedLowCacheHitRate {
	this := AggregatedLowCacheHitRate{}
	this.AvgCacheHitRate = avgCacheHitRate
	this.AvgResourceDownloadSizeBytes = avgResourceDownloadSizeBytes
	this.Fingerprint = fingerprint
	this.ImpactScore = impactScore
	this.ViewOccurrences = viewOccurrences
	return &this
}

// NewAggregatedLowCacheHitRateWithDefaults instantiates a new AggregatedLowCacheHitRate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedLowCacheHitRateWithDefaults() *AggregatedLowCacheHitRate {
	this := AggregatedLowCacheHitRate{}
	return &this
}

// GetAvgCacheHitRate returns the AvgCacheHitRate field value.
func (o *AggregatedLowCacheHitRate) GetAvgCacheHitRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AvgCacheHitRate
}

// GetAvgCacheHitRateOk returns a tuple with the AvgCacheHitRate field value
// and a boolean to check if the value has been set.
func (o *AggregatedLowCacheHitRate) GetAvgCacheHitRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgCacheHitRate, true
}

// SetAvgCacheHitRate sets field value.
func (o *AggregatedLowCacheHitRate) SetAvgCacheHitRate(v float64) {
	o.AvgCacheHitRate = v
}

// GetAvgResourceDownloadSizeBytes returns the AvgResourceDownloadSizeBytes field value.
func (o *AggregatedLowCacheHitRate) GetAvgResourceDownloadSizeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgResourceDownloadSizeBytes
}

// GetAvgResourceDownloadSizeBytesOk returns a tuple with the AvgResourceDownloadSizeBytes field value
// and a boolean to check if the value has been set.
func (o *AggregatedLowCacheHitRate) GetAvgResourceDownloadSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgResourceDownloadSizeBytes, true
}

// SetAvgResourceDownloadSizeBytes sets field value.
func (o *AggregatedLowCacheHitRate) SetAvgResourceDownloadSizeBytes(v int64) {
	o.AvgResourceDownloadSizeBytes = v
}

// GetFingerprint returns the Fingerprint field value.
func (o *AggregatedLowCacheHitRate) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *AggregatedLowCacheHitRate) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *AggregatedLowCacheHitRate) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetImpactScore returns the ImpactScore field value.
func (o *AggregatedLowCacheHitRate) GetImpactScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ImpactScore
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value
// and a boolean to check if the value has been set.
func (o *AggregatedLowCacheHitRate) GetImpactScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImpactScore, true
}

// SetImpactScore sets field value.
func (o *AggregatedLowCacheHitRate) SetImpactScore(v float64) {
	o.ImpactScore = v
}

// GetViewOccurrences returns the ViewOccurrences field value.
func (o *AggregatedLowCacheHitRate) GetViewOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewOccurrences
}

// GetViewOccurrencesOk returns a tuple with the ViewOccurrences field value
// and a boolean to check if the value has been set.
func (o *AggregatedLowCacheHitRate) GetViewOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewOccurrences, true
}

// SetViewOccurrences sets field value.
func (o *AggregatedLowCacheHitRate) SetViewOccurrences(v int32) {
	o.ViewOccurrences = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedLowCacheHitRate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["avg_cache_hit_rate"] = o.AvgCacheHitRate
	toSerialize["avg_resource_download_size_bytes"] = o.AvgResourceDownloadSizeBytes
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["impact_score"] = o.ImpactScore
	toSerialize["view_occurrences"] = o.ViewOccurrences

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedLowCacheHitRate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvgCacheHitRate              *float64 `json:"avg_cache_hit_rate"`
		AvgResourceDownloadSizeBytes *int64   `json:"avg_resource_download_size_bytes"`
		Fingerprint                  *string  `json:"fingerprint"`
		ImpactScore                  *float64 `json:"impact_score"`
		ViewOccurrences              *int32   `json:"view_occurrences"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AvgCacheHitRate == nil {
		return fmt.Errorf("required field avg_cache_hit_rate missing")
	}
	if all.AvgResourceDownloadSizeBytes == nil {
		return fmt.Errorf("required field avg_resource_download_size_bytes missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"avg_cache_hit_rate", "avg_resource_download_size_bytes", "fingerprint", "impact_score", "view_occurrences"})
	} else {
		return err
	}
	o.AvgCacheHitRate = *all.AvgCacheHitRate
	o.AvgResourceDownloadSizeBytes = *all.AvgResourceDownloadSizeBytes
	o.Fingerprint = *all.Fingerprint
	o.ImpactScore = *all.ImpactScore
	o.ViewOccurrences = *all.ViewOccurrences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
