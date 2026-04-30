// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestHistoryPolicyMetaConfig Configuration parameters of the policy that triggered this status change.
type FlakyTestHistoryPolicyMetaConfig struct {
	// The branches considered by the policy.
	Branches datadog.NullableList[string] `json:"branches,omitempty"`
	// The number of days a test must have been active for the policy to trigger.
	DaysActive datadog.NullableInt32 `json:"days_active,omitempty"`
	// The failure rate threshold for the policy to trigger.
	FailureRate datadog.NullableFloat64 `json:"failure_rate,omitempty"`
	// Branches excluded from the policy evaluation.
	ForgetBranches datadog.NullableList[string] `json:"forget_branches,omitempty"`
	// The minimum number of test runs required for the policy to trigger.
	RequiredRuns datadog.NullableInt32 `json:"required_runs,omitempty"`
	// The target state the policy transitions the test from.
	State datadog.NullableString `json:"state,omitempty"`
	// Test services excluded from the policy evaluation.
	TestServices datadog.NullableList[string] `json:"test_services,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFlakyTestHistoryPolicyMetaConfig instantiates a new FlakyTestHistoryPolicyMetaConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlakyTestHistoryPolicyMetaConfig() *FlakyTestHistoryPolicyMetaConfig {
	this := FlakyTestHistoryPolicyMetaConfig{}
	return &this
}

// NewFlakyTestHistoryPolicyMetaConfigWithDefaults instantiates a new FlakyTestHistoryPolicyMetaConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlakyTestHistoryPolicyMetaConfigWithDefaults() *FlakyTestHistoryPolicyMetaConfig {
	this := FlakyTestHistoryPolicyMetaConfig{}
	return &this
}

// GetBranches returns the Branches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMetaConfig) GetBranches() []string {
	if o == nil || o.Branches.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Branches.Get()
}

// GetBranchesOk returns a tuple with the Branches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMetaConfig) GetBranchesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branches.Get(), o.Branches.IsSet()
}

// HasBranches returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMetaConfig) HasBranches() bool {
	return o != nil && o.Branches.IsSet()
}

// SetBranches gets a reference to the given datadog.NullableList[string] and assigns it to the Branches field.
func (o *FlakyTestHistoryPolicyMetaConfig) SetBranches(v []string) {
	o.Branches.Set(&v)
}

// SetBranchesNil sets the value for Branches to be an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) SetBranchesNil() {
	o.Branches.Set(nil)
}

// UnsetBranches ensures that no value is present for Branches, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) UnsetBranches() {
	o.Branches.Unset()
}

// GetDaysActive returns the DaysActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMetaConfig) GetDaysActive() int32 {
	if o == nil || o.DaysActive.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DaysActive.Get()
}

// GetDaysActiveOk returns a tuple with the DaysActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMetaConfig) GetDaysActiveOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DaysActive.Get(), o.DaysActive.IsSet()
}

// HasDaysActive returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMetaConfig) HasDaysActive() bool {
	return o != nil && o.DaysActive.IsSet()
}

// SetDaysActive gets a reference to the given datadog.NullableInt32 and assigns it to the DaysActive field.
func (o *FlakyTestHistoryPolicyMetaConfig) SetDaysActive(v int32) {
	o.DaysActive.Set(&v)
}

// SetDaysActiveNil sets the value for DaysActive to be an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) SetDaysActiveNil() {
	o.DaysActive.Set(nil)
}

// UnsetDaysActive ensures that no value is present for DaysActive, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) UnsetDaysActive() {
	o.DaysActive.Unset()
}

// GetFailureRate returns the FailureRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMetaConfig) GetFailureRate() float64 {
	if o == nil || o.FailureRate.Get() == nil {
		var ret float64
		return ret
	}
	return *o.FailureRate.Get()
}

// GetFailureRateOk returns a tuple with the FailureRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMetaConfig) GetFailureRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureRate.Get(), o.FailureRate.IsSet()
}

// HasFailureRate returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMetaConfig) HasFailureRate() bool {
	return o != nil && o.FailureRate.IsSet()
}

// SetFailureRate gets a reference to the given datadog.NullableFloat64 and assigns it to the FailureRate field.
func (o *FlakyTestHistoryPolicyMetaConfig) SetFailureRate(v float64) {
	o.FailureRate.Set(&v)
}

// SetFailureRateNil sets the value for FailureRate to be an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) SetFailureRateNil() {
	o.FailureRate.Set(nil)
}

// UnsetFailureRate ensures that no value is present for FailureRate, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) UnsetFailureRate() {
	o.FailureRate.Unset()
}

// GetForgetBranches returns the ForgetBranches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMetaConfig) GetForgetBranches() []string {
	if o == nil || o.ForgetBranches.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ForgetBranches.Get()
}

// GetForgetBranchesOk returns a tuple with the ForgetBranches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMetaConfig) GetForgetBranchesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForgetBranches.Get(), o.ForgetBranches.IsSet()
}

// HasForgetBranches returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMetaConfig) HasForgetBranches() bool {
	return o != nil && o.ForgetBranches.IsSet()
}

// SetForgetBranches gets a reference to the given datadog.NullableList[string] and assigns it to the ForgetBranches field.
func (o *FlakyTestHistoryPolicyMetaConfig) SetForgetBranches(v []string) {
	o.ForgetBranches.Set(&v)
}

// SetForgetBranchesNil sets the value for ForgetBranches to be an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) SetForgetBranchesNil() {
	o.ForgetBranches.Set(nil)
}

// UnsetForgetBranches ensures that no value is present for ForgetBranches, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) UnsetForgetBranches() {
	o.ForgetBranches.Unset()
}

// GetRequiredRuns returns the RequiredRuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMetaConfig) GetRequiredRuns() int32 {
	if o == nil || o.RequiredRuns.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RequiredRuns.Get()
}

// GetRequiredRunsOk returns a tuple with the RequiredRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMetaConfig) GetRequiredRunsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredRuns.Get(), o.RequiredRuns.IsSet()
}

// HasRequiredRuns returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMetaConfig) HasRequiredRuns() bool {
	return o != nil && o.RequiredRuns.IsSet()
}

// SetRequiredRuns gets a reference to the given datadog.NullableInt32 and assigns it to the RequiredRuns field.
func (o *FlakyTestHistoryPolicyMetaConfig) SetRequiredRuns(v int32) {
	o.RequiredRuns.Set(&v)
}

// SetRequiredRunsNil sets the value for RequiredRuns to be an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) SetRequiredRunsNil() {
	o.RequiredRuns.Set(nil)
}

// UnsetRequiredRuns ensures that no value is present for RequiredRuns, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) UnsetRequiredRuns() {
	o.RequiredRuns.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMetaConfig) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMetaConfig) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMetaConfig) HasState() bool {
	return o != nil && o.State.IsSet()
}

// SetState gets a reference to the given datadog.NullableString and assigns it to the State field.
func (o *FlakyTestHistoryPolicyMetaConfig) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) UnsetState() {
	o.State.Unset()
}

// GetTestServices returns the TestServices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMetaConfig) GetTestServices() []string {
	if o == nil || o.TestServices.Get() == nil {
		var ret []string
		return ret
	}
	return *o.TestServices.Get()
}

// GetTestServicesOk returns a tuple with the TestServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMetaConfig) GetTestServicesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestServices.Get(), o.TestServices.IsSet()
}

// HasTestServices returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMetaConfig) HasTestServices() bool {
	return o != nil && o.TestServices.IsSet()
}

// SetTestServices gets a reference to the given datadog.NullableList[string] and assigns it to the TestServices field.
func (o *FlakyTestHistoryPolicyMetaConfig) SetTestServices(v []string) {
	o.TestServices.Set(&v)
}

// SetTestServicesNil sets the value for TestServices to be an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) SetTestServicesNil() {
	o.TestServices.Set(nil)
}

// UnsetTestServices ensures that no value is present for TestServices, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMetaConfig) UnsetTestServices() {
	o.TestServices.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o FlakyTestHistoryPolicyMetaConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Branches.IsSet() {
		toSerialize["branches"] = o.Branches.Get()
	}
	if o.DaysActive.IsSet() {
		toSerialize["days_active"] = o.DaysActive.Get()
	}
	if o.FailureRate.IsSet() {
		toSerialize["failure_rate"] = o.FailureRate.Get()
	}
	if o.ForgetBranches.IsSet() {
		toSerialize["forget_branches"] = o.ForgetBranches.Get()
	}
	if o.RequiredRuns.IsSet() {
		toSerialize["required_runs"] = o.RequiredRuns.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.TestServices.IsSet() {
		toSerialize["test_services"] = o.TestServices.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FlakyTestHistoryPolicyMetaConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Branches       datadog.NullableList[string] `json:"branches,omitempty"`
		DaysActive     datadog.NullableInt32        `json:"days_active,omitempty"`
		FailureRate    datadog.NullableFloat64      `json:"failure_rate,omitempty"`
		ForgetBranches datadog.NullableList[string] `json:"forget_branches,omitempty"`
		RequiredRuns   datadog.NullableInt32        `json:"required_runs,omitempty"`
		State          datadog.NullableString       `json:"state,omitempty"`
		TestServices   datadog.NullableList[string] `json:"test_services,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"branches", "days_active", "failure_rate", "forget_branches", "required_runs", "state", "test_services"})
	} else {
		return err
	}
	o.Branches = all.Branches
	o.DaysActive = all.DaysActive
	o.FailureRate = all.FailureRate
	o.ForgetBranches = all.ForgetBranches
	o.RequiredRuns = all.RequiredRuns
	o.State = all.State
	o.TestServices = all.TestServices

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
