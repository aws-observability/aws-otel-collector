// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleAnomalyDetectionOptions Options on anomaly detection method.
type SecurityMonitoringRuleAnomalyDetectionOptions struct {
	// Duration in seconds of the time buckets used to aggregate events matched by the rule.
	// Must be greater than or equal to 300.
	BucketDuration *SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration `json:"bucketDuration,omitempty"`
	// An optional parameter that sets how permissive anomaly detection is.
	// Higher values require higher deviations before triggering a signal.
	DetectionTolerance *SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance `json:"detectionTolerance,omitempty"`
	// When set to true, Datadog uses previous values that fall within the defined learning window to construct the baseline, enabling the system to establish an accurate baseline more rapidly rather than relying solely on gradual learning over time.
	InstantaneousBaseline *bool `json:"instantaneousBaseline,omitempty"`
	// Learning duration in hours. Anomaly detection waits for at least this amount of historical data before it starts evaluating.
	LearningDuration *SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration `json:"learningDuration,omitempty"`
	// An optional override baseline to apply while the rule is in the learning period. Must be greater than or equal to 0.
	LearningPeriodBaseline *int64 `json:"learningPeriodBaseline,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleAnomalyDetectionOptions instantiates a new SecurityMonitoringRuleAnomalyDetectionOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleAnomalyDetectionOptions() *SecurityMonitoringRuleAnomalyDetectionOptions {
	this := SecurityMonitoringRuleAnomalyDetectionOptions{}
	return &this
}

// NewSecurityMonitoringRuleAnomalyDetectionOptionsWithDefaults instantiates a new SecurityMonitoringRuleAnomalyDetectionOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleAnomalyDetectionOptionsWithDefaults() *SecurityMonitoringRuleAnomalyDetectionOptions {
	this := SecurityMonitoringRuleAnomalyDetectionOptions{}
	return &this
}

// GetBucketDuration returns the BucketDuration field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) GetBucketDuration() SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration {
	if o == nil || o.BucketDuration == nil {
		var ret SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration
		return ret
	}
	return *o.BucketDuration
}

// GetBucketDurationOk returns a tuple with the BucketDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) GetBucketDurationOk() (*SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration, bool) {
	if o == nil || o.BucketDuration == nil {
		return nil, false
	}
	return o.BucketDuration, true
}

// HasBucketDuration returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) HasBucketDuration() bool {
	return o != nil && o.BucketDuration != nil
}

// SetBucketDuration gets a reference to the given SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration and assigns it to the BucketDuration field.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) SetBucketDuration(v SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration) {
	o.BucketDuration = &v
}

// GetDetectionTolerance returns the DetectionTolerance field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) GetDetectionTolerance() SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance {
	if o == nil || o.DetectionTolerance == nil {
		var ret SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance
		return ret
	}
	return *o.DetectionTolerance
}

// GetDetectionToleranceOk returns a tuple with the DetectionTolerance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) GetDetectionToleranceOk() (*SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance, bool) {
	if o == nil || o.DetectionTolerance == nil {
		return nil, false
	}
	return o.DetectionTolerance, true
}

// HasDetectionTolerance returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) HasDetectionTolerance() bool {
	return o != nil && o.DetectionTolerance != nil
}

// SetDetectionTolerance gets a reference to the given SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance and assigns it to the DetectionTolerance field.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) SetDetectionTolerance(v SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance) {
	o.DetectionTolerance = &v
}

// GetInstantaneousBaseline returns the InstantaneousBaseline field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) GetInstantaneousBaseline() bool {
	if o == nil || o.InstantaneousBaseline == nil {
		var ret bool
		return ret
	}
	return *o.InstantaneousBaseline
}

// GetInstantaneousBaselineOk returns a tuple with the InstantaneousBaseline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) GetInstantaneousBaselineOk() (*bool, bool) {
	if o == nil || o.InstantaneousBaseline == nil {
		return nil, false
	}
	return o.InstantaneousBaseline, true
}

// HasInstantaneousBaseline returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) HasInstantaneousBaseline() bool {
	return o != nil && o.InstantaneousBaseline != nil
}

// SetInstantaneousBaseline gets a reference to the given bool and assigns it to the InstantaneousBaseline field.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) SetInstantaneousBaseline(v bool) {
	o.InstantaneousBaseline = &v
}

// GetLearningDuration returns the LearningDuration field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) GetLearningDuration() SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration {
	if o == nil || o.LearningDuration == nil {
		var ret SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration
		return ret
	}
	return *o.LearningDuration
}

// GetLearningDurationOk returns a tuple with the LearningDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) GetLearningDurationOk() (*SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration, bool) {
	if o == nil || o.LearningDuration == nil {
		return nil, false
	}
	return o.LearningDuration, true
}

// HasLearningDuration returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) HasLearningDuration() bool {
	return o != nil && o.LearningDuration != nil
}

// SetLearningDuration gets a reference to the given SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration and assigns it to the LearningDuration field.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) SetLearningDuration(v SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration) {
	o.LearningDuration = &v
}

// GetLearningPeriodBaseline returns the LearningPeriodBaseline field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) GetLearningPeriodBaseline() int64 {
	if o == nil || o.LearningPeriodBaseline == nil {
		var ret int64
		return ret
	}
	return *o.LearningPeriodBaseline
}

// GetLearningPeriodBaselineOk returns a tuple with the LearningPeriodBaseline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) GetLearningPeriodBaselineOk() (*int64, bool) {
	if o == nil || o.LearningPeriodBaseline == nil {
		return nil, false
	}
	return o.LearningPeriodBaseline, true
}

// HasLearningPeriodBaseline returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) HasLearningPeriodBaseline() bool {
	return o != nil && o.LearningPeriodBaseline != nil
}

// SetLearningPeriodBaseline gets a reference to the given int64 and assigns it to the LearningPeriodBaseline field.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) SetLearningPeriodBaseline(v int64) {
	o.LearningPeriodBaseline = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleAnomalyDetectionOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BucketDuration != nil {
		toSerialize["bucketDuration"] = o.BucketDuration
	}
	if o.DetectionTolerance != nil {
		toSerialize["detectionTolerance"] = o.DetectionTolerance
	}
	if o.InstantaneousBaseline != nil {
		toSerialize["instantaneousBaseline"] = o.InstantaneousBaseline
	}
	if o.LearningDuration != nil {
		toSerialize["learningDuration"] = o.LearningDuration
	}
	if o.LearningPeriodBaseline != nil {
		toSerialize["learningPeriodBaseline"] = o.LearningPeriodBaseline
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleAnomalyDetectionOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketDuration         *SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration     `json:"bucketDuration,omitempty"`
		DetectionTolerance     *SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance `json:"detectionTolerance,omitempty"`
		InstantaneousBaseline  *bool                                                            `json:"instantaneousBaseline,omitempty"`
		LearningDuration       *SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration   `json:"learningDuration,omitempty"`
		LearningPeriodBaseline *int64                                                           `json:"learningPeriodBaseline,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucketDuration", "detectionTolerance", "instantaneousBaseline", "learningDuration", "learningPeriodBaseline"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.BucketDuration != nil && !all.BucketDuration.IsValid() {
		hasInvalidField = true
	} else {
		o.BucketDuration = all.BucketDuration
	}
	if all.DetectionTolerance != nil && !all.DetectionTolerance.IsValid() {
		hasInvalidField = true
	} else {
		o.DetectionTolerance = all.DetectionTolerance
	}
	o.InstantaneousBaseline = all.InstantaneousBaseline
	if all.LearningDuration != nil && !all.LearningDuration.IsValid() {
		hasInvalidField = true
	} else {
		o.LearningDuration = all.LearningDuration
	}
	o.LearningPeriodBaseline = all.LearningPeriodBaseline

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
