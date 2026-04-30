// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationUpdateServiceSettingsRequestAttributes Attributes for updating Test Optimization service settings.
// All non-required fields are optional; only provided fields will be updated.
type TestOptimizationUpdateServiceSettingsRequestAttributes struct {
	// Whether Auto Test Retries are enabled for this service.
	AutoTestRetriesEnabled *bool `json:"auto_test_retries_enabled,omitempty"`
	// Whether Code Coverage is enabled for this service.
	CodeCoverageEnabled *bool `json:"code_coverage_enabled,omitempty"`
	// Whether Early Flake Detection is enabled for this service.
	EarlyFlakeDetectionEnabled *bool `json:"early_flake_detection_enabled,omitempty"`
	// The environment name. If omitted, defaults to `none`.
	Env *string `json:"env,omitempty"`
	// Whether Failed Test Replay is enabled for this service.
	FailedTestReplayEnabled *bool `json:"failed_test_replay_enabled,omitempty"`
	// Whether PR Comments are enabled for this service.
	PrCommentsEnabled *bool `json:"pr_comments_enabled,omitempty"`
	// The repository identifier.
	RepositoryId string `json:"repository_id"`
	// The service name.
	ServiceName string `json:"service_name"`
	// Whether Test Impact Analysis is enabled for this service.
	TestImpactAnalysisEnabled *bool `json:"test_impact_analysis_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestOptimizationUpdateServiceSettingsRequestAttributes instantiates a new TestOptimizationUpdateServiceSettingsRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestOptimizationUpdateServiceSettingsRequestAttributes(repositoryId string, serviceName string) *TestOptimizationUpdateServiceSettingsRequestAttributes {
	this := TestOptimizationUpdateServiceSettingsRequestAttributes{}
	this.RepositoryId = repositoryId
	this.ServiceName = serviceName
	return &this
}

// NewTestOptimizationUpdateServiceSettingsRequestAttributesWithDefaults instantiates a new TestOptimizationUpdateServiceSettingsRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestOptimizationUpdateServiceSettingsRequestAttributesWithDefaults() *TestOptimizationUpdateServiceSettingsRequestAttributes {
	this := TestOptimizationUpdateServiceSettingsRequestAttributes{}
	return &this
}

// GetAutoTestRetriesEnabled returns the AutoTestRetriesEnabled field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetAutoTestRetriesEnabled() bool {
	if o == nil || o.AutoTestRetriesEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutoTestRetriesEnabled
}

// GetAutoTestRetriesEnabledOk returns a tuple with the AutoTestRetriesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetAutoTestRetriesEnabledOk() (*bool, bool) {
	if o == nil || o.AutoTestRetriesEnabled == nil {
		return nil, false
	}
	return o.AutoTestRetriesEnabled, true
}

// HasAutoTestRetriesEnabled returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasAutoTestRetriesEnabled() bool {
	return o != nil && o.AutoTestRetriesEnabled != nil
}

// SetAutoTestRetriesEnabled gets a reference to the given bool and assigns it to the AutoTestRetriesEnabled field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetAutoTestRetriesEnabled(v bool) {
	o.AutoTestRetriesEnabled = &v
}

// GetCodeCoverageEnabled returns the CodeCoverageEnabled field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetCodeCoverageEnabled() bool {
	if o == nil || o.CodeCoverageEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CodeCoverageEnabled
}

// GetCodeCoverageEnabledOk returns a tuple with the CodeCoverageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetCodeCoverageEnabledOk() (*bool, bool) {
	if o == nil || o.CodeCoverageEnabled == nil {
		return nil, false
	}
	return o.CodeCoverageEnabled, true
}

// HasCodeCoverageEnabled returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasCodeCoverageEnabled() bool {
	return o != nil && o.CodeCoverageEnabled != nil
}

// SetCodeCoverageEnabled gets a reference to the given bool and assigns it to the CodeCoverageEnabled field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetCodeCoverageEnabled(v bool) {
	o.CodeCoverageEnabled = &v
}

// GetEarlyFlakeDetectionEnabled returns the EarlyFlakeDetectionEnabled field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetEarlyFlakeDetectionEnabled() bool {
	if o == nil || o.EarlyFlakeDetectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EarlyFlakeDetectionEnabled
}

// GetEarlyFlakeDetectionEnabledOk returns a tuple with the EarlyFlakeDetectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetEarlyFlakeDetectionEnabledOk() (*bool, bool) {
	if o == nil || o.EarlyFlakeDetectionEnabled == nil {
		return nil, false
	}
	return o.EarlyFlakeDetectionEnabled, true
}

// HasEarlyFlakeDetectionEnabled returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasEarlyFlakeDetectionEnabled() bool {
	return o != nil && o.EarlyFlakeDetectionEnabled != nil
}

// SetEarlyFlakeDetectionEnabled gets a reference to the given bool and assigns it to the EarlyFlakeDetectionEnabled field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetEarlyFlakeDetectionEnabled(v bool) {
	o.EarlyFlakeDetectionEnabled = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetFailedTestReplayEnabled returns the FailedTestReplayEnabled field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetFailedTestReplayEnabled() bool {
	if o == nil || o.FailedTestReplayEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FailedTestReplayEnabled
}

// GetFailedTestReplayEnabledOk returns a tuple with the FailedTestReplayEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetFailedTestReplayEnabledOk() (*bool, bool) {
	if o == nil || o.FailedTestReplayEnabled == nil {
		return nil, false
	}
	return o.FailedTestReplayEnabled, true
}

// HasFailedTestReplayEnabled returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasFailedTestReplayEnabled() bool {
	return o != nil && o.FailedTestReplayEnabled != nil
}

// SetFailedTestReplayEnabled gets a reference to the given bool and assigns it to the FailedTestReplayEnabled field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetFailedTestReplayEnabled(v bool) {
	o.FailedTestReplayEnabled = &v
}

// GetPrCommentsEnabled returns the PrCommentsEnabled field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetPrCommentsEnabled() bool {
	if o == nil || o.PrCommentsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PrCommentsEnabled
}

// GetPrCommentsEnabledOk returns a tuple with the PrCommentsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetPrCommentsEnabledOk() (*bool, bool) {
	if o == nil || o.PrCommentsEnabled == nil {
		return nil, false
	}
	return o.PrCommentsEnabled, true
}

// HasPrCommentsEnabled returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasPrCommentsEnabled() bool {
	return o != nil && o.PrCommentsEnabled != nil
}

// SetPrCommentsEnabled gets a reference to the given bool and assigns it to the PrCommentsEnabled field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetPrCommentsEnabled(v bool) {
	o.PrCommentsEnabled = &v
}

// GetRepositoryId returns the RepositoryId field value.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetRepositoryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryId, true
}

// SetRepositoryId sets field value.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetRepositoryId(v string) {
	o.RepositoryId = v
}

// GetServiceName returns the ServiceName field value.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetServiceName(v string) {
	o.ServiceName = v
}

// GetTestImpactAnalysisEnabled returns the TestImpactAnalysisEnabled field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetTestImpactAnalysisEnabled() bool {
	if o == nil || o.TestImpactAnalysisEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TestImpactAnalysisEnabled
}

// GetTestImpactAnalysisEnabledOk returns a tuple with the TestImpactAnalysisEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetTestImpactAnalysisEnabledOk() (*bool, bool) {
	if o == nil || o.TestImpactAnalysisEnabled == nil {
		return nil, false
	}
	return o.TestImpactAnalysisEnabled, true
}

// HasTestImpactAnalysisEnabled returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasTestImpactAnalysisEnabled() bool {
	return o != nil && o.TestImpactAnalysisEnabled != nil
}

// SetTestImpactAnalysisEnabled gets a reference to the given bool and assigns it to the TestImpactAnalysisEnabled field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetTestImpactAnalysisEnabled(v bool) {
	o.TestImpactAnalysisEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationUpdateServiceSettingsRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoTestRetriesEnabled != nil {
		toSerialize["auto_test_retries_enabled"] = o.AutoTestRetriesEnabled
	}
	if o.CodeCoverageEnabled != nil {
		toSerialize["code_coverage_enabled"] = o.CodeCoverageEnabled
	}
	if o.EarlyFlakeDetectionEnabled != nil {
		toSerialize["early_flake_detection_enabled"] = o.EarlyFlakeDetectionEnabled
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.FailedTestReplayEnabled != nil {
		toSerialize["failed_test_replay_enabled"] = o.FailedTestReplayEnabled
	}
	if o.PrCommentsEnabled != nil {
		toSerialize["pr_comments_enabled"] = o.PrCommentsEnabled
	}
	toSerialize["repository_id"] = o.RepositoryId
	toSerialize["service_name"] = o.ServiceName
	if o.TestImpactAnalysisEnabled != nil {
		toSerialize["test_impact_analysis_enabled"] = o.TestImpactAnalysisEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoTestRetriesEnabled     *bool   `json:"auto_test_retries_enabled,omitempty"`
		CodeCoverageEnabled        *bool   `json:"code_coverage_enabled,omitempty"`
		EarlyFlakeDetectionEnabled *bool   `json:"early_flake_detection_enabled,omitempty"`
		Env                        *string `json:"env,omitempty"`
		FailedTestReplayEnabled    *bool   `json:"failed_test_replay_enabled,omitempty"`
		PrCommentsEnabled          *bool   `json:"pr_comments_enabled,omitempty"`
		RepositoryId               *string `json:"repository_id"`
		ServiceName                *string `json:"service_name"`
		TestImpactAnalysisEnabled  *bool   `json:"test_impact_analysis_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RepositoryId == nil {
		return fmt.Errorf("required field repository_id missing")
	}
	if all.ServiceName == nil {
		return fmt.Errorf("required field service_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_test_retries_enabled", "code_coverage_enabled", "early_flake_detection_enabled", "env", "failed_test_replay_enabled", "pr_comments_enabled", "repository_id", "service_name", "test_impact_analysis_enabled"})
	} else {
		return err
	}
	o.AutoTestRetriesEnabled = all.AutoTestRetriesEnabled
	o.CodeCoverageEnabled = all.CodeCoverageEnabled
	o.EarlyFlakeDetectionEnabled = all.EarlyFlakeDetectionEnabled
	o.Env = all.Env
	o.FailedTestReplayEnabled = all.FailedTestReplayEnabled
	o.PrCommentsEnabled = all.PrCommentsEnabled
	o.RepositoryId = *all.RepositoryId
	o.ServiceName = *all.ServiceName
	o.TestImpactAnalysisEnabled = all.TestImpactAnalysisEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
