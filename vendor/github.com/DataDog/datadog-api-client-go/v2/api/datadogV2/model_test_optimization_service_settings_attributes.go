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
	// Whether Code Coverage is enabled for this service.
	CodeCoverageEnabled *bool `json:"code_coverage_enabled,omitempty"`
	// Whether Early Flake Detection is enabled for this service.
	EarlyFlakeDetectionEnabled *bool `json:"early_flake_detection_enabled,omitempty"`
	// The environment name.
	Env *string `json:"env,omitempty"`
	// Whether Failed Test Replay is enabled for this service.
	FailedTestReplayEnabled *bool `json:"failed_test_replay_enabled,omitempty"`
	// Whether PR Comments are enabled for this service.
	PrCommentsEnabled *bool `json:"pr_comments_enabled,omitempty"`
	// The repository identifier.
	RepositoryId *string `json:"repository_id,omitempty"`
	// The service name.
	ServiceName *string `json:"service_name,omitempty"`
	// Whether Test Impact Analysis is enabled for this service.
	TestImpactAnalysisEnabled *bool `json:"test_impact_analysis_enabled,omitempty"`
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

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationServiceSettingsAttributes) MarshalJSON() ([]byte, error) {
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
	if o.RepositoryId != nil {
		toSerialize["repository_id"] = o.RepositoryId
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.TestImpactAnalysisEnabled != nil {
		toSerialize["test_impact_analysis_enabled"] = o.TestImpactAnalysisEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationServiceSettingsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoTestRetriesEnabled     *bool   `json:"auto_test_retries_enabled,omitempty"`
		CodeCoverageEnabled        *bool   `json:"code_coverage_enabled,omitempty"`
		EarlyFlakeDetectionEnabled *bool   `json:"early_flake_detection_enabled,omitempty"`
		Env                        *string `json:"env,omitempty"`
		FailedTestReplayEnabled    *bool   `json:"failed_test_replay_enabled,omitempty"`
		PrCommentsEnabled          *bool   `json:"pr_comments_enabled,omitempty"`
		RepositoryId               *string `json:"repository_id,omitempty"`
		ServiceName                *string `json:"service_name,omitempty"`
		TestImpactAnalysisEnabled  *bool   `json:"test_impact_analysis_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.RepositoryId = all.RepositoryId
	o.ServiceName = all.ServiceName
	o.TestImpactAnalysisEnabled = all.TestImpactAnalysisEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
