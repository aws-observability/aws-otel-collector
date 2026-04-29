// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule Failure-rate-based rule for the disabled policy.
type TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule struct {
	// List of branches to which this rule applies.
	Branches []string `json:"branches,omitempty"`
	// Whether this failure rate rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Minimum number of runs required before the rule is evaluated. Must be greater than or equal to 0.
	MinRuns *int64 `json:"min_runs,omitempty"`
	// Test status that the disable policy applies to.
	// Must be either `active` or `quarantined`.
	Status *TestOptimizationFlakyTestsManagementPoliciesDisabledStatus `json:"status,omitempty"`
	// Failure rate threshold (0.0–1.0) above which the rule triggers.
	Threshold *float64 `json:"threshold,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule instantiates a new TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule() *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule {
	this := TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule{}
	return &this
}

// NewTestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRuleWithDefaults instantiates a new TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRuleWithDefaults() *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule {
	this := TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule{}
	return &this
}

// GetBranches returns the Branches field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) GetBranches() []string {
	if o == nil || o.Branches == nil {
		var ret []string
		return ret
	}
	return o.Branches
}

// GetBranchesOk returns a tuple with the Branches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) GetBranchesOk() (*[]string, bool) {
	if o == nil || o.Branches == nil {
		return nil, false
	}
	return &o.Branches, true
}

// HasBranches returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) HasBranches() bool {
	return o != nil && o.Branches != nil
}

// SetBranches gets a reference to the given []string and assigns it to the Branches field.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) SetBranches(v []string) {
	o.Branches = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMinRuns returns the MinRuns field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) GetMinRuns() int64 {
	if o == nil || o.MinRuns == nil {
		var ret int64
		return ret
	}
	return *o.MinRuns
}

// GetMinRunsOk returns a tuple with the MinRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) GetMinRunsOk() (*int64, bool) {
	if o == nil || o.MinRuns == nil {
		return nil, false
	}
	return o.MinRuns, true
}

// HasMinRuns returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) HasMinRuns() bool {
	return o != nil && o.MinRuns != nil
}

// SetMinRuns gets a reference to the given int64 and assigns it to the MinRuns field.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) SetMinRuns(v int64) {
	o.MinRuns = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) GetStatus() TestOptimizationFlakyTestsManagementPoliciesDisabledStatus {
	if o == nil || o.Status == nil {
		var ret TestOptimizationFlakyTestsManagementPoliciesDisabledStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) GetStatusOk() (*TestOptimizationFlakyTestsManagementPoliciesDisabledStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given TestOptimizationFlakyTestsManagementPoliciesDisabledStatus and assigns it to the Status field.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) SetStatus(v TestOptimizationFlakyTestsManagementPoliciesDisabledStatus) {
	o.Status = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) GetThreshold() float64 {
	if o == nil || o.Threshold == nil {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) GetThresholdOk() (*float64, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) HasThreshold() bool {
	return o != nil && o.Threshold != nil
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) SetThreshold(v float64) {
	o.Threshold = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Branches != nil {
		toSerialize["branches"] = o.Branches
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.MinRuns != nil {
		toSerialize["min_runs"] = o.MinRuns
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Branches  []string                                                    `json:"branches,omitempty"`
		Enabled   *bool                                                       `json:"enabled,omitempty"`
		MinRuns   *int64                                                      `json:"min_runs,omitempty"`
		Status    *TestOptimizationFlakyTestsManagementPoliciesDisabledStatus `json:"status,omitempty"`
		Threshold *float64                                                    `json:"threshold,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"branches", "enabled", "min_runs", "status", "threshold"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Branches = all.Branches
	o.Enabled = all.Enabled
	o.MinRuns = all.MinRuns
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Threshold = all.Threshold

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
