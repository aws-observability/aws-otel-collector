// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedSlowFCPHighBytes Aggregated slow first contentful paint with high byte count detection.
type AggregatedSlowFCPHighBytes struct {
	// Average total bytes loaded before first contentful paint.
	AvgBytesBeforeFcpBytes int64 `json:"avg_bytes_before_fcp_bytes"`
	// Average first contentful paint time in milliseconds.
	AvgFirstContentfulPaintMs int64 `json:"avg_first_contentful_paint_ms"`
	// Average number of resources loaded before first contentful paint.
	AvgResourceCountBeforeFcp int64 `json:"avg_resource_count_before_fcp"`
	// Unique fingerprint identifying this detection group.
	Fingerprint string `json:"fingerprint"`
	// Impact score for this detection.
	ImpactScore float64 `json:"impact_score"`
	// Platform identifier for the affected views.
	Platform string `json:"platform"`
	// Number of sampled views where this detection occurred.
	ViewOccurrences int32 `json:"view_occurrences"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedSlowFCPHighBytes instantiates a new AggregatedSlowFCPHighBytes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedSlowFCPHighBytes(avgBytesBeforeFcpBytes int64, avgFirstContentfulPaintMs int64, avgResourceCountBeforeFcp int64, fingerprint string, impactScore float64, platform string, viewOccurrences int32) *AggregatedSlowFCPHighBytes {
	this := AggregatedSlowFCPHighBytes{}
	this.AvgBytesBeforeFcpBytes = avgBytesBeforeFcpBytes
	this.AvgFirstContentfulPaintMs = avgFirstContentfulPaintMs
	this.AvgResourceCountBeforeFcp = avgResourceCountBeforeFcp
	this.Fingerprint = fingerprint
	this.ImpactScore = impactScore
	this.Platform = platform
	this.ViewOccurrences = viewOccurrences
	return &this
}

// NewAggregatedSlowFCPHighBytesWithDefaults instantiates a new AggregatedSlowFCPHighBytes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedSlowFCPHighBytesWithDefaults() *AggregatedSlowFCPHighBytes {
	this := AggregatedSlowFCPHighBytes{}
	return &this
}

// GetAvgBytesBeforeFcpBytes returns the AvgBytesBeforeFcpBytes field value.
func (o *AggregatedSlowFCPHighBytes) GetAvgBytesBeforeFcpBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgBytesBeforeFcpBytes
}

// GetAvgBytesBeforeFcpBytesOk returns a tuple with the AvgBytesBeforeFcpBytes field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowFCPHighBytes) GetAvgBytesBeforeFcpBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgBytesBeforeFcpBytes, true
}

// SetAvgBytesBeforeFcpBytes sets field value.
func (o *AggregatedSlowFCPHighBytes) SetAvgBytesBeforeFcpBytes(v int64) {
	o.AvgBytesBeforeFcpBytes = v
}

// GetAvgFirstContentfulPaintMs returns the AvgFirstContentfulPaintMs field value.
func (o *AggregatedSlowFCPHighBytes) GetAvgFirstContentfulPaintMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgFirstContentfulPaintMs
}

// GetAvgFirstContentfulPaintMsOk returns a tuple with the AvgFirstContentfulPaintMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowFCPHighBytes) GetAvgFirstContentfulPaintMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgFirstContentfulPaintMs, true
}

// SetAvgFirstContentfulPaintMs sets field value.
func (o *AggregatedSlowFCPHighBytes) SetAvgFirstContentfulPaintMs(v int64) {
	o.AvgFirstContentfulPaintMs = v
}

// GetAvgResourceCountBeforeFcp returns the AvgResourceCountBeforeFcp field value.
func (o *AggregatedSlowFCPHighBytes) GetAvgResourceCountBeforeFcp() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgResourceCountBeforeFcp
}

// GetAvgResourceCountBeforeFcpOk returns a tuple with the AvgResourceCountBeforeFcp field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowFCPHighBytes) GetAvgResourceCountBeforeFcpOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgResourceCountBeforeFcp, true
}

// SetAvgResourceCountBeforeFcp sets field value.
func (o *AggregatedSlowFCPHighBytes) SetAvgResourceCountBeforeFcp(v int64) {
	o.AvgResourceCountBeforeFcp = v
}

// GetFingerprint returns the Fingerprint field value.
func (o *AggregatedSlowFCPHighBytes) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowFCPHighBytes) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *AggregatedSlowFCPHighBytes) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetImpactScore returns the ImpactScore field value.
func (o *AggregatedSlowFCPHighBytes) GetImpactScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ImpactScore
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowFCPHighBytes) GetImpactScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImpactScore, true
}

// SetImpactScore sets field value.
func (o *AggregatedSlowFCPHighBytes) SetImpactScore(v float64) {
	o.ImpactScore = v
}

// GetPlatform returns the Platform field value.
func (o *AggregatedSlowFCPHighBytes) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowFCPHighBytes) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value.
func (o *AggregatedSlowFCPHighBytes) SetPlatform(v string) {
	o.Platform = v
}

// GetViewOccurrences returns the ViewOccurrences field value.
func (o *AggregatedSlowFCPHighBytes) GetViewOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewOccurrences
}

// GetViewOccurrencesOk returns a tuple with the ViewOccurrences field value
// and a boolean to check if the value has been set.
func (o *AggregatedSlowFCPHighBytes) GetViewOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewOccurrences, true
}

// SetViewOccurrences sets field value.
func (o *AggregatedSlowFCPHighBytes) SetViewOccurrences(v int32) {
	o.ViewOccurrences = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedSlowFCPHighBytes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["avg_bytes_before_fcp_bytes"] = o.AvgBytesBeforeFcpBytes
	toSerialize["avg_first_contentful_paint_ms"] = o.AvgFirstContentfulPaintMs
	toSerialize["avg_resource_count_before_fcp"] = o.AvgResourceCountBeforeFcp
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["impact_score"] = o.ImpactScore
	toSerialize["platform"] = o.Platform
	toSerialize["view_occurrences"] = o.ViewOccurrences

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedSlowFCPHighBytes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvgBytesBeforeFcpBytes    *int64   `json:"avg_bytes_before_fcp_bytes"`
		AvgFirstContentfulPaintMs *int64   `json:"avg_first_contentful_paint_ms"`
		AvgResourceCountBeforeFcp *int64   `json:"avg_resource_count_before_fcp"`
		Fingerprint               *string  `json:"fingerprint"`
		ImpactScore               *float64 `json:"impact_score"`
		Platform                  *string  `json:"platform"`
		ViewOccurrences           *int32   `json:"view_occurrences"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AvgBytesBeforeFcpBytes == nil {
		return fmt.Errorf("required field avg_bytes_before_fcp_bytes missing")
	}
	if all.AvgFirstContentfulPaintMs == nil {
		return fmt.Errorf("required field avg_first_contentful_paint_ms missing")
	}
	if all.AvgResourceCountBeforeFcp == nil {
		return fmt.Errorf("required field avg_resource_count_before_fcp missing")
	}
	if all.Fingerprint == nil {
		return fmt.Errorf("required field fingerprint missing")
	}
	if all.ImpactScore == nil {
		return fmt.Errorf("required field impact_score missing")
	}
	if all.Platform == nil {
		return fmt.Errorf("required field platform missing")
	}
	if all.ViewOccurrences == nil {
		return fmt.Errorf("required field view_occurrences missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avg_bytes_before_fcp_bytes", "avg_first_contentful_paint_ms", "avg_resource_count_before_fcp", "fingerprint", "impact_score", "platform", "view_occurrences"})
	} else {
		return err
	}
	o.AvgBytesBeforeFcpBytes = *all.AvgBytesBeforeFcpBytes
	o.AvgFirstContentfulPaintMs = *all.AvgFirstContentfulPaintMs
	o.AvgResourceCountBeforeFcp = *all.AvgResourceCountBeforeFcp
	o.Fingerprint = *all.Fingerprint
	o.ImpactScore = *all.ImpactScore
	o.Platform = *all.Platform
	o.ViewOccurrences = *all.ViewOccurrences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
