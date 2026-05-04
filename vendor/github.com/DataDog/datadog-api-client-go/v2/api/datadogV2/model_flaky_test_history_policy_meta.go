// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestHistoryPolicyMeta Metadata about the policy that triggered this status change.
type FlakyTestHistoryPolicyMeta struct {
	// Branches where the test was flaky at the time of the status change.
	Branches datadog.NullableList[string] `json:"branches,omitempty"`
	// Configuration parameters of the policy that triggered this status change.
	Config *FlakyTestHistoryPolicyMetaConfig `json:"config,omitempty"`
	// The number of days the test has been active at the time of the status change.
	DaysActive datadog.NullableInt32 `json:"days_active,omitempty"`
	// The number of days since the test last exhibited flakiness.
	DaysWithoutFlake datadog.NullableInt32 `json:"days_without_flake,omitempty"`
	// The failure rate of the test at the time of the status change.
	FailureRate datadog.NullableFloat64 `json:"failure_rate,omitempty"`
	// The previous state of the test.
	State datadog.NullableString `json:"state,omitempty"`
	// The total number of test runs at the time of the status change.
	TotalRuns datadog.NullableInt32 `json:"total_runs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFlakyTestHistoryPolicyMeta instantiates a new FlakyTestHistoryPolicyMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlakyTestHistoryPolicyMeta() *FlakyTestHistoryPolicyMeta {
	this := FlakyTestHistoryPolicyMeta{}
	return &this
}

// NewFlakyTestHistoryPolicyMetaWithDefaults instantiates a new FlakyTestHistoryPolicyMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlakyTestHistoryPolicyMetaWithDefaults() *FlakyTestHistoryPolicyMeta {
	this := FlakyTestHistoryPolicyMeta{}
	return &this
}

// GetBranches returns the Branches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMeta) GetBranches() []string {
	if o == nil || o.Branches.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Branches.Get()
}

// GetBranchesOk returns a tuple with the Branches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMeta) GetBranchesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branches.Get(), o.Branches.IsSet()
}

// HasBranches returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMeta) HasBranches() bool {
	return o != nil && o.Branches.IsSet()
}

// SetBranches gets a reference to the given datadog.NullableList[string] and assigns it to the Branches field.
func (o *FlakyTestHistoryPolicyMeta) SetBranches(v []string) {
	o.Branches.Set(&v)
}

// SetBranchesNil sets the value for Branches to be an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) SetBranchesNil() {
	o.Branches.Set(nil)
}

// UnsetBranches ensures that no value is present for Branches, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) UnsetBranches() {
	o.Branches.Unset()
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *FlakyTestHistoryPolicyMeta) GetConfig() FlakyTestHistoryPolicyMetaConfig {
	if o == nil || o.Config == nil {
		var ret FlakyTestHistoryPolicyMetaConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestHistoryPolicyMeta) GetConfigOk() (*FlakyTestHistoryPolicyMetaConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMeta) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given FlakyTestHistoryPolicyMetaConfig and assigns it to the Config field.
func (o *FlakyTestHistoryPolicyMeta) SetConfig(v FlakyTestHistoryPolicyMetaConfig) {
	o.Config = &v
}

// GetDaysActive returns the DaysActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMeta) GetDaysActive() int32 {
	if o == nil || o.DaysActive.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DaysActive.Get()
}

// GetDaysActiveOk returns a tuple with the DaysActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMeta) GetDaysActiveOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DaysActive.Get(), o.DaysActive.IsSet()
}

// HasDaysActive returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMeta) HasDaysActive() bool {
	return o != nil && o.DaysActive.IsSet()
}

// SetDaysActive gets a reference to the given datadog.NullableInt32 and assigns it to the DaysActive field.
func (o *FlakyTestHistoryPolicyMeta) SetDaysActive(v int32) {
	o.DaysActive.Set(&v)
}

// SetDaysActiveNil sets the value for DaysActive to be an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) SetDaysActiveNil() {
	o.DaysActive.Set(nil)
}

// UnsetDaysActive ensures that no value is present for DaysActive, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) UnsetDaysActive() {
	o.DaysActive.Unset()
}

// GetDaysWithoutFlake returns the DaysWithoutFlake field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMeta) GetDaysWithoutFlake() int32 {
	if o == nil || o.DaysWithoutFlake.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DaysWithoutFlake.Get()
}

// GetDaysWithoutFlakeOk returns a tuple with the DaysWithoutFlake field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMeta) GetDaysWithoutFlakeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DaysWithoutFlake.Get(), o.DaysWithoutFlake.IsSet()
}

// HasDaysWithoutFlake returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMeta) HasDaysWithoutFlake() bool {
	return o != nil && o.DaysWithoutFlake.IsSet()
}

// SetDaysWithoutFlake gets a reference to the given datadog.NullableInt32 and assigns it to the DaysWithoutFlake field.
func (o *FlakyTestHistoryPolicyMeta) SetDaysWithoutFlake(v int32) {
	o.DaysWithoutFlake.Set(&v)
}

// SetDaysWithoutFlakeNil sets the value for DaysWithoutFlake to be an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) SetDaysWithoutFlakeNil() {
	o.DaysWithoutFlake.Set(nil)
}

// UnsetDaysWithoutFlake ensures that no value is present for DaysWithoutFlake, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) UnsetDaysWithoutFlake() {
	o.DaysWithoutFlake.Unset()
}

// GetFailureRate returns the FailureRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMeta) GetFailureRate() float64 {
	if o == nil || o.FailureRate.Get() == nil {
		var ret float64
		return ret
	}
	return *o.FailureRate.Get()
}

// GetFailureRateOk returns a tuple with the FailureRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMeta) GetFailureRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureRate.Get(), o.FailureRate.IsSet()
}

// HasFailureRate returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMeta) HasFailureRate() bool {
	return o != nil && o.FailureRate.IsSet()
}

// SetFailureRate gets a reference to the given datadog.NullableFloat64 and assigns it to the FailureRate field.
func (o *FlakyTestHistoryPolicyMeta) SetFailureRate(v float64) {
	o.FailureRate.Set(&v)
}

// SetFailureRateNil sets the value for FailureRate to be an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) SetFailureRateNil() {
	o.FailureRate.Set(nil)
}

// UnsetFailureRate ensures that no value is present for FailureRate, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) UnsetFailureRate() {
	o.FailureRate.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMeta) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMeta) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMeta) HasState() bool {
	return o != nil && o.State.IsSet()
}

// SetState gets a reference to the given datadog.NullableString and assigns it to the State field.
func (o *FlakyTestHistoryPolicyMeta) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) UnsetState() {
	o.State.Unset()
}

// GetTotalRuns returns the TotalRuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestHistoryPolicyMeta) GetTotalRuns() int32 {
	if o == nil || o.TotalRuns.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TotalRuns.Get()
}

// GetTotalRunsOk returns a tuple with the TotalRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestHistoryPolicyMeta) GetTotalRunsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalRuns.Get(), o.TotalRuns.IsSet()
}

// HasTotalRuns returns a boolean if a field has been set.
func (o *FlakyTestHistoryPolicyMeta) HasTotalRuns() bool {
	return o != nil && o.TotalRuns.IsSet()
}

// SetTotalRuns gets a reference to the given datadog.NullableInt32 and assigns it to the TotalRuns field.
func (o *FlakyTestHistoryPolicyMeta) SetTotalRuns(v int32) {
	o.TotalRuns.Set(&v)
}

// SetTotalRunsNil sets the value for TotalRuns to be an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) SetTotalRunsNil() {
	o.TotalRuns.Set(nil)
}

// UnsetTotalRuns ensures that no value is present for TotalRuns, not even an explicit nil.
func (o *FlakyTestHistoryPolicyMeta) UnsetTotalRuns() {
	o.TotalRuns.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o FlakyTestHistoryPolicyMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Branches.IsSet() {
		toSerialize["branches"] = o.Branches.Get()
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.DaysActive.IsSet() {
		toSerialize["days_active"] = o.DaysActive.Get()
	}
	if o.DaysWithoutFlake.IsSet() {
		toSerialize["days_without_flake"] = o.DaysWithoutFlake.Get()
	}
	if o.FailureRate.IsSet() {
		toSerialize["failure_rate"] = o.FailureRate.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.TotalRuns.IsSet() {
		toSerialize["total_runs"] = o.TotalRuns.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FlakyTestHistoryPolicyMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Branches         datadog.NullableList[string]      `json:"branches,omitempty"`
		Config           *FlakyTestHistoryPolicyMetaConfig `json:"config,omitempty"`
		DaysActive       datadog.NullableInt32             `json:"days_active,omitempty"`
		DaysWithoutFlake datadog.NullableInt32             `json:"days_without_flake,omitempty"`
		FailureRate      datadog.NullableFloat64           `json:"failure_rate,omitempty"`
		State            datadog.NullableString            `json:"state,omitempty"`
		TotalRuns        datadog.NullableInt32             `json:"total_runs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"branches", "config", "days_active", "days_without_flake", "failure_rate", "state", "total_runs"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Branches = all.Branches
	if all.Config != nil && all.Config.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Config = all.Config
	o.DaysActive = all.DaysActive
	o.DaysWithoutFlake = all.DaysWithoutFlake
	o.FailureRate = all.FailureRate
	o.State = all.State
	o.TotalRuns = all.TotalRuns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
