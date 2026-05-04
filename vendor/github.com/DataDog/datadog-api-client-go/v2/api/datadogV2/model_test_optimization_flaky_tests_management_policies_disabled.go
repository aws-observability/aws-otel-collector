// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationFlakyTestsManagementPoliciesDisabled Configuration for the disabled Flaky Tests Management policy.
type TestOptimizationFlakyTestsManagementPoliciesDisabled struct {
	// Automatic disable triggering rule based on a time window and test status.
	AutoDisableRule *TestOptimizationFlakyTestsManagementPoliciesAutoDisableRule `json:"auto_disable_rule,omitempty"`
	// Branch filtering rule for a Flaky Tests Management policy.
	BranchRule *TestOptimizationFlakyTestsManagementPoliciesBranchRule `json:"branch_rule,omitempty"`
	// Whether the disabled policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Failure-rate-based rule for the disabled policy.
	FailureRateRule *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule `json:"failure_rate_rule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestOptimizationFlakyTestsManagementPoliciesDisabled instantiates a new TestOptimizationFlakyTestsManagementPoliciesDisabled object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestOptimizationFlakyTestsManagementPoliciesDisabled() *TestOptimizationFlakyTestsManagementPoliciesDisabled {
	this := TestOptimizationFlakyTestsManagementPoliciesDisabled{}
	return &this
}

// NewTestOptimizationFlakyTestsManagementPoliciesDisabledWithDefaults instantiates a new TestOptimizationFlakyTestsManagementPoliciesDisabled object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestOptimizationFlakyTestsManagementPoliciesDisabledWithDefaults() *TestOptimizationFlakyTestsManagementPoliciesDisabled {
	this := TestOptimizationFlakyTestsManagementPoliciesDisabled{}
	return &this
}

// GetAutoDisableRule returns the AutoDisableRule field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) GetAutoDisableRule() TestOptimizationFlakyTestsManagementPoliciesAutoDisableRule {
	if o == nil || o.AutoDisableRule == nil {
		var ret TestOptimizationFlakyTestsManagementPoliciesAutoDisableRule
		return ret
	}
	return *o.AutoDisableRule
}

// GetAutoDisableRuleOk returns a tuple with the AutoDisableRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) GetAutoDisableRuleOk() (*TestOptimizationFlakyTestsManagementPoliciesAutoDisableRule, bool) {
	if o == nil || o.AutoDisableRule == nil {
		return nil, false
	}
	return o.AutoDisableRule, true
}

// HasAutoDisableRule returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) HasAutoDisableRule() bool {
	return o != nil && o.AutoDisableRule != nil
}

// SetAutoDisableRule gets a reference to the given TestOptimizationFlakyTestsManagementPoliciesAutoDisableRule and assigns it to the AutoDisableRule field.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) SetAutoDisableRule(v TestOptimizationFlakyTestsManagementPoliciesAutoDisableRule) {
	o.AutoDisableRule = &v
}

// GetBranchRule returns the BranchRule field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) GetBranchRule() TestOptimizationFlakyTestsManagementPoliciesBranchRule {
	if o == nil || o.BranchRule == nil {
		var ret TestOptimizationFlakyTestsManagementPoliciesBranchRule
		return ret
	}
	return *o.BranchRule
}

// GetBranchRuleOk returns a tuple with the BranchRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) GetBranchRuleOk() (*TestOptimizationFlakyTestsManagementPoliciesBranchRule, bool) {
	if o == nil || o.BranchRule == nil {
		return nil, false
	}
	return o.BranchRule, true
}

// HasBranchRule returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) HasBranchRule() bool {
	return o != nil && o.BranchRule != nil
}

// SetBranchRule gets a reference to the given TestOptimizationFlakyTestsManagementPoliciesBranchRule and assigns it to the BranchRule field.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) SetBranchRule(v TestOptimizationFlakyTestsManagementPoliciesBranchRule) {
	o.BranchRule = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFailureRateRule returns the FailureRateRule field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) GetFailureRateRule() TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule {
	if o == nil || o.FailureRateRule == nil {
		var ret TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule
		return ret
	}
	return *o.FailureRateRule
}

// GetFailureRateRuleOk returns a tuple with the FailureRateRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) GetFailureRateRuleOk() (*TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule, bool) {
	if o == nil || o.FailureRateRule == nil {
		return nil, false
	}
	return o.FailureRateRule, true
}

// HasFailureRateRule returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) HasFailureRateRule() bool {
	return o != nil && o.FailureRateRule != nil
}

// SetFailureRateRule gets a reference to the given TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule and assigns it to the FailureRateRule field.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) SetFailureRateRule(v TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) {
	o.FailureRateRule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationFlakyTestsManagementPoliciesDisabled) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoDisableRule != nil {
		toSerialize["auto_disable_rule"] = o.AutoDisableRule
	}
	if o.BranchRule != nil {
		toSerialize["branch_rule"] = o.BranchRule
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.FailureRateRule != nil {
		toSerialize["failure_rate_rule"] = o.FailureRateRule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabled) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoDisableRule *TestOptimizationFlakyTestsManagementPoliciesAutoDisableRule         `json:"auto_disable_rule,omitempty"`
		BranchRule      *TestOptimizationFlakyTestsManagementPoliciesBranchRule              `json:"branch_rule,omitempty"`
		Enabled         *bool                                                                `json:"enabled,omitempty"`
		FailureRateRule *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule `json:"failure_rate_rule,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_disable_rule", "branch_rule", "enabled", "failure_rate_rule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AutoDisableRule != nil && all.AutoDisableRule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AutoDisableRule = all.AutoDisableRule
	if all.BranchRule != nil && all.BranchRule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BranchRule = all.BranchRule
	o.Enabled = all.Enabled
	if all.FailureRateRule != nil && all.FailureRateRule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FailureRateRule = all.FailureRateRule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
