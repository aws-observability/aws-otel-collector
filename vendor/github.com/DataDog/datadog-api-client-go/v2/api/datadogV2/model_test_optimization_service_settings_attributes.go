// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationServiceSettingsAttributes Attributes for Test Optimization service settings.
type TestOptimizationServiceSettingsAttributes struct {
	// Whether Auto Test Retries are enabled for this service.
	AutoTestRetriesEnabled *bool `json:"auto_test_retries_enabled,omitempty"`
	// Whether the Auto Test Retries setting is overridden at the service level.
	AutoTestRetriesEnabledIsOverridden *bool `json:"auto_test_retries_enabled_is_overridden,omitempty"`
	// Whether Code Coverage is enabled for this service.
	CodeCoverageEnabled *bool `json:"code_coverage_enabled,omitempty"`
	// Whether the Code Coverage setting is overridden at the service level.
	CodeCoverageEnabledIsOverridden *bool `json:"code_coverage_enabled_is_overridden,omitempty"`
	// Whether Early Flake Detection is enabled for this service.
	EarlyFlakeDetectionEnabled *bool `json:"early_flake_detection_enabled,omitempty"`
	// Whether the Early Flake Detection setting is overridden at the service level.
	EarlyFlakeDetectionEnabledIsOverridden *bool `json:"early_flake_detection_enabled_is_overridden,omitempty"`
	// The environment name.
	Env *string `json:"env,omitempty"`
	// Whether Failed Test Replay is enabled for this service.
	FailedTestReplayEnabled *bool `json:"failed_test_replay_enabled,omitempty"`
	// Whether the Failed Test Replay setting is overridden at the service level.
	FailedTestReplayEnabledIsOverridden *bool `json:"failed_test_replay_enabled_is_overridden,omitempty"`
	// Whether PR Comments are enabled. This value reflects the repository-level setting and cannot be overridden at the service level.
	PrCommentsEnabled *bool `json:"pr_comments_enabled,omitempty"`
	// The repository identifier.
	RepositoryId *string `json:"repository_id,omitempty"`
	// The service name.
	ServiceName *string `json:"service_name,omitempty"`
	// Whether Test Impact Analysis is enabled for this service.
	TestImpactAnalysisEnabled *bool `json:"test_impact_analysis_enabled,omitempty"`
	// Whether the Test Impact Analysis setting is overridden at the service level.
	TestImpactAnalysisEnabledIsOverridden *bool `json:"test_impact_analysis_enabled_is_overridden,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestOptimizationServiceSettingsAttributes instantiates a new TestOptimizationServiceSettingsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestOptimizationServiceSettingsAttributes() *TestOptimizationServiceSettingsAttributes {
	this := TestOptimizationServiceSettingsAttributes{}
	return &this
}

// NewTestOptimizationServiceSettingsAttributesWithDefaults instantiates a new TestOptimizationServiceSettingsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestOptimizationServiceSettingsAttributesWithDefaults() *TestOptimizationServiceSettingsAttributes {
	this := TestOptimizationServiceSettingsAttributes{}
	return &this
}

// GetAutoTestRetriesEnabled returns the AutoTestRetriesEnabled field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetAutoTestRetriesEnabled() bool {
	if o == nil || o.AutoTestRetriesEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutoTestRetriesEnabled
}

// GetAutoTestRetriesEnabledOk returns a tuple with the AutoTestRetriesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetAutoTestRetriesEnabledOk() (*bool, bool) {
	if o == nil || o.AutoTestRetriesEnabled == nil {
		return nil, false
	}
	return o.AutoTestRetriesEnabled, true
}

// HasAutoTestRetriesEnabled returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasAutoTestRetriesEnabled() bool {
	return o != nil && o.AutoTestRetriesEnabled != nil
}

// SetAutoTestRetriesEnabled gets a reference to the given bool and assigns it to the AutoTestRetriesEnabled field.
func (o *TestOptimizationServiceSettingsAttributes) SetAutoTestRetriesEnabled(v bool) {
	o.AutoTestRetriesEnabled = &v
}

// GetAutoTestRetriesEnabledIsOverridden returns the AutoTestRetriesEnabledIsOverridden field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetAutoTestRetriesEnabledIsOverridden() bool {
	if o == nil || o.AutoTestRetriesEnabledIsOverridden == nil {
		var ret bool
		return ret
	}
	return *o.AutoTestRetriesEnabledIsOverridden
}

// GetAutoTestRetriesEnabledIsOverriddenOk returns a tuple with the AutoTestRetriesEnabledIsOverridden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetAutoTestRetriesEnabledIsOverriddenOk() (*bool, bool) {
	if o == nil || o.AutoTestRetriesEnabledIsOverridden == nil {
		return nil, false
	}
	return o.AutoTestRetriesEnabledIsOverridden, true
}

// HasAutoTestRetriesEnabledIsOverridden returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasAutoTestRetriesEnabledIsOverridden() bool {
	return o != nil && o.AutoTestRetriesEnabledIsOverridden != nil
}

// SetAutoTestRetriesEnabledIsOverridden gets a reference to the given bool and assigns it to the AutoTestRetriesEnabledIsOverridden field.
func (o *TestOptimizationServiceSettingsAttributes) SetAutoTestRetriesEnabledIsOverridden(v bool) {
	o.AutoTestRetriesEnabledIsOverridden = &v
}

// GetCodeCoverageEnabled returns the CodeCoverageEnabled field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetCodeCoverageEnabled() bool {
	if o == nil || o.CodeCoverageEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CodeCoverageEnabled
}

// GetCodeCoverageEnabledOk returns a tuple with the CodeCoverageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetCodeCoverageEnabledOk() (*bool, bool) {
	if o == nil || o.CodeCoverageEnabled == nil {
		return nil, false
	}
	return o.CodeCoverageEnabled, true
}

// HasCodeCoverageEnabled returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasCodeCoverageEnabled() bool {
	return o != nil && o.CodeCoverageEnabled != nil
}

// SetCodeCoverageEnabled gets a reference to the given bool and assigns it to the CodeCoverageEnabled field.
func (o *TestOptimizationServiceSettingsAttributes) SetCodeCoverageEnabled(v bool) {
	o.CodeCoverageEnabled = &v
}

// GetCodeCoverageEnabledIsOverridden returns the CodeCoverageEnabledIsOverridden field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetCodeCoverageEnabledIsOverridden() bool {
	if o == nil || o.CodeCoverageEnabledIsOverridden == nil {
		var ret bool
		return ret
	}
	return *o.CodeCoverageEnabledIsOverridden
}

// GetCodeCoverageEnabledIsOverriddenOk returns a tuple with the CodeCoverageEnabledIsOverridden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetCodeCoverageEnabledIsOverriddenOk() (*bool, bool) {
	if o == nil || o.CodeCoverageEnabledIsOverridden == nil {
		return nil, false
	}
	return o.CodeCoverageEnabledIsOverridden, true
}

// HasCodeCoverageEnabledIsOverridden returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasCodeCoverageEnabledIsOverridden() bool {
	return o != nil && o.CodeCoverageEnabledIsOverridden != nil
}

// SetCodeCoverageEnabledIsOverridden gets a reference to the given bool and assigns it to the CodeCoverageEnabledIsOverridden field.
func (o *TestOptimizationServiceSettingsAttributes) SetCodeCoverageEnabledIsOverridden(v bool) {
	o.CodeCoverageEnabledIsOverridden = &v
}

// GetEarlyFlakeDetectionEnabled returns the EarlyFlakeDetectionEnabled field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetEarlyFlakeDetectionEnabled() bool {
	if o == nil || o.EarlyFlakeDetectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EarlyFlakeDetectionEnabled
}

// GetEarlyFlakeDetectionEnabledOk returns a tuple with the EarlyFlakeDetectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetEarlyFlakeDetectionEnabledOk() (*bool, bool) {
	if o == nil || o.EarlyFlakeDetectionEnabled == nil {
		return nil, false
	}
	return o.EarlyFlakeDetectionEnabled, true
}

// HasEarlyFlakeDetectionEnabled returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasEarlyFlakeDetectionEnabled() bool {
	return o != nil && o.EarlyFlakeDetectionEnabled != nil
}

// SetEarlyFlakeDetectionEnabled gets a reference to the given bool and assigns it to the EarlyFlakeDetectionEnabled field.
func (o *TestOptimizationServiceSettingsAttributes) SetEarlyFlakeDetectionEnabled(v bool) {
	o.EarlyFlakeDetectionEnabled = &v
}

// GetEarlyFlakeDetectionEnabledIsOverridden returns the EarlyFlakeDetectionEnabledIsOverridden field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetEarlyFlakeDetectionEnabledIsOverridden() bool {
	if o == nil || o.EarlyFlakeDetectionEnabledIsOverridden == nil {
		var ret bool
		return ret
	}
	return *o.EarlyFlakeDetectionEnabledIsOverridden
}

// GetEarlyFlakeDetectionEnabledIsOverriddenOk returns a tuple with the EarlyFlakeDetectionEnabledIsOverridden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetEarlyFlakeDetectionEnabledIsOverriddenOk() (*bool, bool) {
	if o == nil || o.EarlyFlakeDetectionEnabledIsOverridden == nil {
		return nil, false
	}
	return o.EarlyFlakeDetectionEnabledIsOverridden, true
}

// HasEarlyFlakeDetectionEnabledIsOverridden returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasEarlyFlakeDetectionEnabledIsOverridden() bool {
	return o != nil && o.EarlyFlakeDetectionEnabledIsOverridden != nil
}

// SetEarlyFlakeDetectionEnabledIsOverridden gets a reference to the given bool and assigns it to the EarlyFlakeDetectionEnabledIsOverridden field.
func (o *TestOptimizationServiceSettingsAttributes) SetEarlyFlakeDetectionEnabledIsOverridden(v bool) {
	o.EarlyFlakeDetectionEnabledIsOverridden = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *TestOptimizationServiceSettingsAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetFailedTestReplayEnabled returns the FailedTestReplayEnabled field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetFailedTestReplayEnabled() bool {
	if o == nil || o.FailedTestReplayEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FailedTestReplayEnabled
}

// GetFailedTestReplayEnabledOk returns a tuple with the FailedTestReplayEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetFailedTestReplayEnabledOk() (*bool, bool) {
	if o == nil || o.FailedTestReplayEnabled == nil {
		return nil, false
	}
	return o.FailedTestReplayEnabled, true
}

// HasFailedTestReplayEnabled returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasFailedTestReplayEnabled() bool {
	return o != nil && o.FailedTestReplayEnabled != nil
}

// SetFailedTestReplayEnabled gets a reference to the given bool and assigns it to the FailedTestReplayEnabled field.
func (o *TestOptimizationServiceSettingsAttributes) SetFailedTestReplayEnabled(v bool) {
	o.FailedTestReplayEnabled = &v
}

// GetFailedTestReplayEnabledIsOverridden returns the FailedTestReplayEnabledIsOverridden field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetFailedTestReplayEnabledIsOverridden() bool {
	if o == nil || o.FailedTestReplayEnabledIsOverridden == nil {
		var ret bool
		return ret
	}
	return *o.FailedTestReplayEnabledIsOverridden
}

// GetFailedTestReplayEnabledIsOverriddenOk returns a tuple with the FailedTestReplayEnabledIsOverridden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetFailedTestReplayEnabledIsOverriddenOk() (*bool, bool) {
	if o == nil || o.FailedTestReplayEnabledIsOverridden == nil {
		return nil, false
	}
	return o.FailedTestReplayEnabledIsOverridden, true
}

// HasFailedTestReplayEnabledIsOverridden returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasFailedTestReplayEnabledIsOverridden() bool {
	return o != nil && o.FailedTestReplayEnabledIsOverridden != nil
}

// SetFailedTestReplayEnabledIsOverridden gets a reference to the given bool and assigns it to the FailedTestReplayEnabledIsOverridden field.
func (o *TestOptimizationServiceSettingsAttributes) SetFailedTestReplayEnabledIsOverridden(v bool) {
	o.FailedTestReplayEnabledIsOverridden = &v
}

// GetPrCommentsEnabled returns the PrCommentsEnabled field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetPrCommentsEnabled() bool {
	if o == nil || o.PrCommentsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PrCommentsEnabled
}

// GetPrCommentsEnabledOk returns a tuple with the PrCommentsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetPrCommentsEnabledOk() (*bool, bool) {
	if o == nil || o.PrCommentsEnabled == nil {
		return nil, false
	}
	return o.PrCommentsEnabled, true
}

// HasPrCommentsEnabled returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasPrCommentsEnabled() bool {
	return o != nil && o.PrCommentsEnabled != nil
}

// SetPrCommentsEnabled gets a reference to the given bool and assigns it to the PrCommentsEnabled field.
func (o *TestOptimizationServiceSettingsAttributes) SetPrCommentsEnabled(v bool) {
	o.PrCommentsEnabled = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetRepositoryId() string {
	if o == nil || o.RepositoryId == nil {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasRepositoryId() bool {
	return o != nil && o.RepositoryId != nil
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *TestOptimizationServiceSettingsAttributes) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasServiceName() bool {
	return o != nil && o.ServiceName != nil
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *TestOptimizationServiceSettingsAttributes) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetTestImpactAnalysisEnabled returns the TestImpactAnalysisEnabled field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetTestImpactAnalysisEnabled() bool {
	if o == nil || o.TestImpactAnalysisEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TestImpactAnalysisEnabled
}

// GetTestImpactAnalysisEnabledOk returns a tuple with the TestImpactAnalysisEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetTestImpactAnalysisEnabledOk() (*bool, bool) {
	if o == nil || o.TestImpactAnalysisEnabled == nil {
		return nil, false
	}
	return o.TestImpactAnalysisEnabled, true
}

// HasTestImpactAnalysisEnabled returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasTestImpactAnalysisEnabled() bool {
	return o != nil && o.TestImpactAnalysisEnabled != nil
}

// SetTestImpactAnalysisEnabled gets a reference to the given bool and assigns it to the TestImpactAnalysisEnabled field.
func (o *TestOptimizationServiceSettingsAttributes) SetTestImpactAnalysisEnabled(v bool) {
	o.TestImpactAnalysisEnabled = &v
}

// GetTestImpactAnalysisEnabledIsOverridden returns the TestImpactAnalysisEnabledIsOverridden field value if set, zero value otherwise.
func (o *TestOptimizationServiceSettingsAttributes) GetTestImpactAnalysisEnabledIsOverridden() bool {
	if o == nil || o.TestImpactAnalysisEnabledIsOverridden == nil {
		var ret bool
		return ret
	}
	return *o.TestImpactAnalysisEnabledIsOverridden
}

// GetTestImpactAnalysisEnabledIsOverriddenOk returns a tuple with the TestImpactAnalysisEnabledIsOverridden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationServiceSettingsAttributes) GetTestImpactAnalysisEnabledIsOverriddenOk() (*bool, bool) {
	if o == nil || o.TestImpactAnalysisEnabledIsOverridden == nil {
		return nil, false
	}
	return o.TestImpactAnalysisEnabledIsOverridden, true
}

// HasTestImpactAnalysisEnabledIsOverridden returns a boolean if a field has been set.
func (o *TestOptimizationServiceSettingsAttributes) HasTestImpactAnalysisEnabledIsOverridden() bool {
	return o != nil && o.TestImpactAnalysisEnabledIsOverridden != nil
}

// SetTestImpactAnalysisEnabledIsOverridden gets a reference to the given bool and assigns it to the TestImpactAnalysisEnabledIsOverridden field.
func (o *TestOptimizationServiceSettingsAttributes) SetTestImpactAnalysisEnabledIsOverridden(v bool) {
	o.TestImpactAnalysisEnabledIsOverridden = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationServiceSettingsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoTestRetriesEnabled != nil {
		toSerialize["auto_test_retries_enabled"] = o.AutoTestRetriesEnabled
	}
	if o.AutoTestRetriesEnabledIsOverridden != nil {
		toSerialize["auto_test_retries_enabled_is_overridden"] = o.AutoTestRetriesEnabledIsOverridden
	}
	if o.CodeCoverageEnabled != nil {
		toSerialize["code_coverage_enabled"] = o.CodeCoverageEnabled
	}
	if o.CodeCoverageEnabledIsOverridden != nil {
		toSerialize["code_coverage_enabled_is_overridden"] = o.CodeCoverageEnabledIsOverridden
	}
	if o.EarlyFlakeDetectionEnabled != nil {
		toSerialize["early_flake_detection_enabled"] = o.EarlyFlakeDetectionEnabled
	}
	if o.EarlyFlakeDetectionEnabledIsOverridden != nil {
		toSerialize["early_flake_detection_enabled_is_overridden"] = o.EarlyFlakeDetectionEnabledIsOverridden
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.FailedTestReplayEnabled != nil {
		toSerialize["failed_test_replay_enabled"] = o.FailedTestReplayEnabled
	}
	if o.FailedTestReplayEnabledIsOverridden != nil {
		toSerialize["failed_test_replay_enabled_is_overridden"] = o.FailedTestReplayEnabledIsOverridden
	}
	if o.PrCommentsEnabled != nil {
		toSerialize["pr_comments_enabled"] = o.PrCommentsEnabled
	}
	if o.RepositoryId != nil {
		toSerialize["repository_id"] = o.RepositoryId
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.TestImpactAnalysisEnabled != nil {
		toSerialize["test_impact_analysis_enabled"] = o.TestImpactAnalysisEnabled
	}
	if o.TestImpactAnalysisEnabledIsOverridden != nil {
		toSerialize["test_impact_analysis_enabled_is_overridden"] = o.TestImpactAnalysisEnabledIsOverridden
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationServiceSettingsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoTestRetriesEnabled                 *bool   `json:"auto_test_retries_enabled,omitempty"`
		AutoTestRetriesEnabledIsOverridden     *bool   `json:"auto_test_retries_enabled_is_overridden,omitempty"`
		CodeCoverageEnabled                    *bool   `json:"code_coverage_enabled,omitempty"`
		CodeCoverageEnabledIsOverridden        *bool   `json:"code_coverage_enabled_is_overridden,omitempty"`
		EarlyFlakeDetectionEnabled             *bool   `json:"early_flake_detection_enabled,omitempty"`
		EarlyFlakeDetectionEnabledIsOverridden *bool   `json:"early_flake_detection_enabled_is_overridden,omitempty"`
		Env                                    *string `json:"env,omitempty"`
		FailedTestReplayEnabled                *bool   `json:"failed_test_replay_enabled,omitempty"`
		FailedTestReplayEnabledIsOverridden    *bool   `json:"failed_test_replay_enabled_is_overridden,omitempty"`
		PrCommentsEnabled                      *bool   `json:"pr_comments_enabled,omitempty"`
		RepositoryId                           *string `json:"repository_id,omitempty"`
		ServiceName                            *string `json:"service_name,omitempty"`
		TestImpactAnalysisEnabled              *bool   `json:"test_impact_analysis_enabled,omitempty"`
		TestImpactAnalysisEnabledIsOverridden  *bool   `json:"test_impact_analysis_enabled_is_overridden,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_test_retries_enabled", "auto_test_retries_enabled_is_overridden", "code_coverage_enabled", "code_coverage_enabled_is_overridden", "early_flake_detection_enabled", "early_flake_detection_enabled_is_overridden", "env", "failed_test_replay_enabled", "failed_test_replay_enabled_is_overridden", "pr_comments_enabled", "repository_id", "service_name", "test_impact_analysis_enabled", "test_impact_analysis_enabled_is_overridden"})
	} else {
		return err
	}
	o.AutoTestRetriesEnabled = all.AutoTestRetriesEnabled
	o.AutoTestRetriesEnabledIsOverridden = all.AutoTestRetriesEnabledIsOverridden
	o.CodeCoverageEnabled = all.CodeCoverageEnabled
	o.CodeCoverageEnabledIsOverridden = all.CodeCoverageEnabledIsOverridden
	o.EarlyFlakeDetectionEnabled = all.EarlyFlakeDetectionEnabled
	o.EarlyFlakeDetectionEnabledIsOverridden = all.EarlyFlakeDetectionEnabledIsOverridden
	o.Env = all.Env
	o.FailedTestReplayEnabled = all.FailedTestReplayEnabled
	o.FailedTestReplayEnabledIsOverridden = all.FailedTestReplayEnabledIsOverridden
	o.PrCommentsEnabled = all.PrCommentsEnabled
	o.RepositoryId = all.RepositoryId
	o.ServiceName = all.ServiceName
	o.TestImpactAnalysisEnabled = all.TestImpactAnalysisEnabled
	o.TestImpactAnalysisEnabledIsOverridden = all.TestImpactAnalysisEnabledIsOverridden

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
