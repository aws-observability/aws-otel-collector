// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationFlakyTestsManagementPoliciesQuarantined Configuration for the quarantined Flaky Tests Management policy.
type TestOptimizationFlakyTestsManagementPoliciesQuarantined struct {
	// Automatic quarantine triggering rule based on a time window.
	AutoQuarantineRule *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule `json:"auto_quarantine_rule,omitempty"`
	// Branch filtering rule for a Flaky Tests Management policy.
	BranchRule *TestOptimizationFlakyTestsManagementPoliciesBranchRule `json:"branch_rule,omitempty"`
	// Whether the quarantined policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Failure-rate-based rule for the quarantined policy.
	FailureRateRule *TestOptimizationFlakyTestsManagementPoliciesQuarantinedFailureRateRule `json:"failure_rate_rule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestOptimizationFlakyTestsManagementPoliciesQuarantined instantiates a new TestOptimizationFlakyTestsManagementPoliciesQuarantined object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestOptimizationFlakyTestsManagementPoliciesQuarantined() *TestOptimizationFlakyTestsManagementPoliciesQuarantined {
	this := TestOptimizationFlakyTestsManagementPoliciesQuarantined{}
	return &this
}

// NewTestOptimizationFlakyTestsManagementPoliciesQuarantinedWithDefaults instantiates a new TestOptimizationFlakyTestsManagementPoliciesQuarantined object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestOptimizationFlakyTestsManagementPoliciesQuarantinedWithDefaults() *TestOptimizationFlakyTestsManagementPoliciesQuarantined {
	this := TestOptimizationFlakyTestsManagementPoliciesQuarantined{}
	return &this
}

// GetAutoQuarantineRule returns the AutoQuarantineRule field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) GetAutoQuarantineRule() TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule {
	if o == nil || o.AutoQuarantineRule == nil {
		var ret TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule
		return ret
	}
	return *o.AutoQuarantineRule
}

// GetAutoQuarantineRuleOk returns a tuple with the AutoQuarantineRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) GetAutoQuarantineRuleOk() (*TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule, bool) {
	if o == nil || o.AutoQuarantineRule == nil {
		return nil, false
	}
	return o.AutoQuarantineRule, true
}

// HasAutoQuarantineRule returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) HasAutoQuarantineRule() bool {
	return o != nil && o.AutoQuarantineRule != nil
}

// SetAutoQuarantineRule gets a reference to the given TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule and assigns it to the AutoQuarantineRule field.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) SetAutoQuarantineRule(v TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule) {
	o.AutoQuarantineRule = &v
}

// GetBranchRule returns the BranchRule field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) GetBranchRule() TestOptimizationFlakyTestsManagementPoliciesBranchRule {
	if o == nil || o.BranchRule == nil {
		var ret TestOptimizationFlakyTestsManagementPoliciesBranchRule
		return ret
	}
	return *o.BranchRule
}

// GetBranchRuleOk returns a tuple with the BranchRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) GetBranchRuleOk() (*TestOptimizationFlakyTestsManagementPoliciesBranchRule, bool) {
	if o == nil || o.BranchRule == nil {
		return nil, false
	}
	return o.BranchRule, true
}

// HasBranchRule returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) HasBranchRule() bool {
	return o != nil && o.BranchRule != nil
}

// SetBranchRule gets a reference to the given TestOptimizationFlakyTestsManagementPoliciesBranchRule and assigns it to the BranchRule field.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) SetBranchRule(v TestOptimizationFlakyTestsManagementPoliciesBranchRule) {
	o.BranchRule = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFailureRateRule returns the FailureRateRule field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) GetFailureRateRule() TestOptimizationFlakyTestsManagementPoliciesQuarantinedFailureRateRule {
	if o == nil || o.FailureRateRule == nil {
		var ret TestOptimizationFlakyTestsManagementPoliciesQuarantinedFailureRateRule
		return ret
	}
	return *o.FailureRateRule
}

// GetFailureRateRuleOk returns a tuple with the FailureRateRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) GetFailureRateRuleOk() (*TestOptimizationFlakyTestsManagementPoliciesQuarantinedFailureRateRule, bool) {
	if o == nil || o.FailureRateRule == nil {
		return nil, false
	}
	return o.FailureRateRule, true
}

// HasFailureRateRule returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) HasFailureRateRule() bool {
	return o != nil && o.FailureRateRule != nil
}

// SetFailureRateRule gets a reference to the given TestOptimizationFlakyTestsManagementPoliciesQuarantinedFailureRateRule and assigns it to the FailureRateRule field.
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) SetFailureRateRule(v TestOptimizationFlakyTestsManagementPoliciesQuarantinedFailureRateRule) {
	o.FailureRateRule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationFlakyTestsManagementPoliciesQuarantined) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoQuarantineRule != nil {
		toSerialize["auto_quarantine_rule"] = o.AutoQuarantineRule
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
func (o *TestOptimizationFlakyTestsManagementPoliciesQuarantined) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoQuarantineRule *TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule         `json:"auto_quarantine_rule,omitempty"`
		BranchRule         *TestOptimizationFlakyTestsManagementPoliciesBranchRule                 `json:"branch_rule,omitempty"`
		Enabled            *bool                                                                   `json:"enabled,omitempty"`
		FailureRateRule    *TestOptimizationFlakyTestsManagementPoliciesQuarantinedFailureRateRule `json:"failure_rate_rule,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_quarantine_rule", "branch_rule", "enabled", "failure_rate_rule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AutoQuarantineRule != nil && all.AutoQuarantineRule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AutoQuarantineRule = all.AutoQuarantineRule
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
