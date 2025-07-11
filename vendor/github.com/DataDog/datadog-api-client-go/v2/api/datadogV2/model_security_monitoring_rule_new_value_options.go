// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleNewValueOptions Options on new value detection method.
type SecurityMonitoringRuleNewValueOptions struct {
	// The duration in days after which a learned value is forgotten.
	ForgetAfter *SecurityMonitoringRuleNewValueOptionsForgetAfter `json:"forgetAfter,omitempty"`
	// The duration in days during which values are learned, and after which signals will be generated for values that
	// weren't learned. If set to 0, a signal will be generated for all new values after the first value is learned.
	LearningDuration *SecurityMonitoringRuleNewValueOptionsLearningDuration `json:"learningDuration,omitempty"`
	// The learning method used to determine when signals should be generated for values that weren't learned.
	LearningMethod *SecurityMonitoringRuleNewValueOptionsLearningMethod `json:"learningMethod,omitempty"`
	// A number of occurrences after which signals will be generated for values that weren't learned.
	LearningThreshold *SecurityMonitoringRuleNewValueOptionsLearningThreshold `json:"learningThreshold,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleNewValueOptions instantiates a new SecurityMonitoringRuleNewValueOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleNewValueOptions() *SecurityMonitoringRuleNewValueOptions {
	this := SecurityMonitoringRuleNewValueOptions{}
	var learningDuration SecurityMonitoringRuleNewValueOptionsLearningDuration = SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_ZERO_DAYS
	this.LearningDuration = &learningDuration
	var learningMethod SecurityMonitoringRuleNewValueOptionsLearningMethod = SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGMETHOD_DURATION
	this.LearningMethod = &learningMethod
	var learningThreshold SecurityMonitoringRuleNewValueOptionsLearningThreshold = SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGTHRESHOLD_ZERO_OCCURRENCES
	this.LearningThreshold = &learningThreshold
	return &this
}

// NewSecurityMonitoringRuleNewValueOptionsWithDefaults instantiates a new SecurityMonitoringRuleNewValueOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleNewValueOptionsWithDefaults() *SecurityMonitoringRuleNewValueOptions {
	this := SecurityMonitoringRuleNewValueOptions{}
	var learningDuration SecurityMonitoringRuleNewValueOptionsLearningDuration = SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_ZERO_DAYS
	this.LearningDuration = &learningDuration
	var learningMethod SecurityMonitoringRuleNewValueOptionsLearningMethod = SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGMETHOD_DURATION
	this.LearningMethod = &learningMethod
	var learningThreshold SecurityMonitoringRuleNewValueOptionsLearningThreshold = SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGTHRESHOLD_ZERO_OCCURRENCES
	this.LearningThreshold = &learningThreshold
	return &this
}

// GetForgetAfter returns the ForgetAfter field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleNewValueOptions) GetForgetAfter() SecurityMonitoringRuleNewValueOptionsForgetAfter {
	if o == nil || o.ForgetAfter == nil {
		var ret SecurityMonitoringRuleNewValueOptionsForgetAfter
		return ret
	}
	return *o.ForgetAfter
}

// GetForgetAfterOk returns a tuple with the ForgetAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleNewValueOptions) GetForgetAfterOk() (*SecurityMonitoringRuleNewValueOptionsForgetAfter, bool) {
	if o == nil || o.ForgetAfter == nil {
		return nil, false
	}
	return o.ForgetAfter, true
}

// HasForgetAfter returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleNewValueOptions) HasForgetAfter() bool {
	return o != nil && o.ForgetAfter != nil
}

// SetForgetAfter gets a reference to the given SecurityMonitoringRuleNewValueOptionsForgetAfter and assigns it to the ForgetAfter field.
func (o *SecurityMonitoringRuleNewValueOptions) SetForgetAfter(v SecurityMonitoringRuleNewValueOptionsForgetAfter) {
	o.ForgetAfter = &v
}

// GetLearningDuration returns the LearningDuration field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleNewValueOptions) GetLearningDuration() SecurityMonitoringRuleNewValueOptionsLearningDuration {
	if o == nil || o.LearningDuration == nil {
		var ret SecurityMonitoringRuleNewValueOptionsLearningDuration
		return ret
	}
	return *o.LearningDuration
}

// GetLearningDurationOk returns a tuple with the LearningDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleNewValueOptions) GetLearningDurationOk() (*SecurityMonitoringRuleNewValueOptionsLearningDuration, bool) {
	if o == nil || o.LearningDuration == nil {
		return nil, false
	}
	return o.LearningDuration, true
}

// HasLearningDuration returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleNewValueOptions) HasLearningDuration() bool {
	return o != nil && o.LearningDuration != nil
}

// SetLearningDuration gets a reference to the given SecurityMonitoringRuleNewValueOptionsLearningDuration and assigns it to the LearningDuration field.
func (o *SecurityMonitoringRuleNewValueOptions) SetLearningDuration(v SecurityMonitoringRuleNewValueOptionsLearningDuration) {
	o.LearningDuration = &v
}

// GetLearningMethod returns the LearningMethod field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleNewValueOptions) GetLearningMethod() SecurityMonitoringRuleNewValueOptionsLearningMethod {
	if o == nil || o.LearningMethod == nil {
		var ret SecurityMonitoringRuleNewValueOptionsLearningMethod
		return ret
	}
	return *o.LearningMethod
}

// GetLearningMethodOk returns a tuple with the LearningMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleNewValueOptions) GetLearningMethodOk() (*SecurityMonitoringRuleNewValueOptionsLearningMethod, bool) {
	if o == nil || o.LearningMethod == nil {
		return nil, false
	}
	return o.LearningMethod, true
}

// HasLearningMethod returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleNewValueOptions) HasLearningMethod() bool {
	return o != nil && o.LearningMethod != nil
}

// SetLearningMethod gets a reference to the given SecurityMonitoringRuleNewValueOptionsLearningMethod and assigns it to the LearningMethod field.
func (o *SecurityMonitoringRuleNewValueOptions) SetLearningMethod(v SecurityMonitoringRuleNewValueOptionsLearningMethod) {
	o.LearningMethod = &v
}

// GetLearningThreshold returns the LearningThreshold field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleNewValueOptions) GetLearningThreshold() SecurityMonitoringRuleNewValueOptionsLearningThreshold {
	if o == nil || o.LearningThreshold == nil {
		var ret SecurityMonitoringRuleNewValueOptionsLearningThreshold
		return ret
	}
	return *o.LearningThreshold
}

// GetLearningThresholdOk returns a tuple with the LearningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleNewValueOptions) GetLearningThresholdOk() (*SecurityMonitoringRuleNewValueOptionsLearningThreshold, bool) {
	if o == nil || o.LearningThreshold == nil {
		return nil, false
	}
	return o.LearningThreshold, true
}

// HasLearningThreshold returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleNewValueOptions) HasLearningThreshold() bool {
	return o != nil && o.LearningThreshold != nil
}

// SetLearningThreshold gets a reference to the given SecurityMonitoringRuleNewValueOptionsLearningThreshold and assigns it to the LearningThreshold field.
func (o *SecurityMonitoringRuleNewValueOptions) SetLearningThreshold(v SecurityMonitoringRuleNewValueOptionsLearningThreshold) {
	o.LearningThreshold = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleNewValueOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ForgetAfter != nil {
		toSerialize["forgetAfter"] = o.ForgetAfter
	}
	if o.LearningDuration != nil {
		toSerialize["learningDuration"] = o.LearningDuration
	}
	if o.LearningMethod != nil {
		toSerialize["learningMethod"] = o.LearningMethod
	}
	if o.LearningThreshold != nil {
		toSerialize["learningThreshold"] = o.LearningThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleNewValueOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ForgetAfter       *SecurityMonitoringRuleNewValueOptionsForgetAfter       `json:"forgetAfter,omitempty"`
		LearningDuration  *SecurityMonitoringRuleNewValueOptionsLearningDuration  `json:"learningDuration,omitempty"`
		LearningMethod    *SecurityMonitoringRuleNewValueOptionsLearningMethod    `json:"learningMethod,omitempty"`
		LearningThreshold *SecurityMonitoringRuleNewValueOptionsLearningThreshold `json:"learningThreshold,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"forgetAfter", "learningDuration", "learningMethod", "learningThreshold"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ForgetAfter != nil && !all.ForgetAfter.IsValid() {
		hasInvalidField = true
	} else {
		o.ForgetAfter = all.ForgetAfter
	}
	if all.LearningDuration != nil && !all.LearningDuration.IsValid() {
		hasInvalidField = true
	} else {
		o.LearningDuration = all.LearningDuration
	}
	if all.LearningMethod != nil && !all.LearningMethod.IsValid() {
		hasInvalidField = true
	} else {
		o.LearningMethod = all.LearningMethod
	}
	if all.LearningThreshold != nil && !all.LearningThreshold.IsValid() {
		hasInvalidField = true
	} else {
		o.LearningThreshold = all.LearningThreshold
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
