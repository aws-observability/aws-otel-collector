// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationFlakyTestsManagementPoliciesBranchRule Branch filtering rule for a Flaky Tests Management policy.
type TestOptimizationFlakyTestsManagementPoliciesBranchRule struct {
	// List of branches to which the policy applies.
	Branches []string `json:"branches,omitempty"`
	// Whether this branch rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of branches excluded from the policy.
	ExcludedBranches []string `json:"excluded_branches,omitempty"`
	// List of test services excluded from the policy.
	ExcludedTestServices []string `json:"excluded_test_services,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestOptimizationFlakyTestsManagementPoliciesBranchRule instantiates a new TestOptimizationFlakyTestsManagementPoliciesBranchRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestOptimizationFlakyTestsManagementPoliciesBranchRule() *TestOptimizationFlakyTestsManagementPoliciesBranchRule {
	this := TestOptimizationFlakyTestsManagementPoliciesBranchRule{}
	return &this
}

// NewTestOptimizationFlakyTestsManagementPoliciesBranchRuleWithDefaults instantiates a new TestOptimizationFlakyTestsManagementPoliciesBranchRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestOptimizationFlakyTestsManagementPoliciesBranchRuleWithDefaults() *TestOptimizationFlakyTestsManagementPoliciesBranchRule {
	this := TestOptimizationFlakyTestsManagementPoliciesBranchRule{}
	return &this
}

// GetBranches returns the Branches field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) GetBranches() []string {
	if o == nil || o.Branches == nil {
		var ret []string
		return ret
	}
	return o.Branches
}

// GetBranchesOk returns a tuple with the Branches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) GetBranchesOk() (*[]string, bool) {
	if o == nil || o.Branches == nil {
		return nil, false
	}
	return &o.Branches, true
}

// HasBranches returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) HasBranches() bool {
	return o != nil && o.Branches != nil
}

// SetBranches gets a reference to the given []string and assigns it to the Branches field.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) SetBranches(v []string) {
	o.Branches = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExcludedBranches returns the ExcludedBranches field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) GetExcludedBranches() []string {
	if o == nil || o.ExcludedBranches == nil {
		var ret []string
		return ret
	}
	return o.ExcludedBranches
}

// GetExcludedBranchesOk returns a tuple with the ExcludedBranches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) GetExcludedBranchesOk() (*[]string, bool) {
	if o == nil || o.ExcludedBranches == nil {
		return nil, false
	}
	return &o.ExcludedBranches, true
}

// HasExcludedBranches returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) HasExcludedBranches() bool {
	return o != nil && o.ExcludedBranches != nil
}

// SetExcludedBranches gets a reference to the given []string and assigns it to the ExcludedBranches field.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) SetExcludedBranches(v []string) {
	o.ExcludedBranches = v
}

// GetExcludedTestServices returns the ExcludedTestServices field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) GetExcludedTestServices() []string {
	if o == nil || o.ExcludedTestServices == nil {
		var ret []string
		return ret
	}
	return o.ExcludedTestServices
}

// GetExcludedTestServicesOk returns a tuple with the ExcludedTestServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) GetExcludedTestServicesOk() (*[]string, bool) {
	if o == nil || o.ExcludedTestServices == nil {
		return nil, false
	}
	return &o.ExcludedTestServices, true
}

// HasExcludedTestServices returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) HasExcludedTestServices() bool {
	return o != nil && o.ExcludedTestServices != nil
}

// SetExcludedTestServices gets a reference to the given []string and assigns it to the ExcludedTestServices field.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) SetExcludedTestServices(v []string) {
	o.ExcludedTestServices = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationFlakyTestsManagementPoliciesBranchRule) MarshalJSON() ([]byte, error) {
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
	if o.ExcludedBranches != nil {
		toSerialize["excluded_branches"] = o.ExcludedBranches
	}
	if o.ExcludedTestServices != nil {
		toSerialize["excluded_test_services"] = o.ExcludedTestServices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationFlakyTestsManagementPoliciesBranchRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Branches             []string `json:"branches,omitempty"`
		Enabled              *bool    `json:"enabled,omitempty"`
		ExcludedBranches     []string `json:"excluded_branches,omitempty"`
		ExcludedTestServices []string `json:"excluded_test_services,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"branches", "enabled", "excluded_branches", "excluded_test_services"})
	} else {
		return err
	}
	o.Branches = all.Branches
	o.Enabled = all.Enabled
	o.ExcludedBranches = all.ExcludedBranches
	o.ExcludedTestServices = all.ExcludedTestServices

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
