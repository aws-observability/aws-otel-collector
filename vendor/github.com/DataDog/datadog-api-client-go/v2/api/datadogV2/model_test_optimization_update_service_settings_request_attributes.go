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
// Setting a field to `null` is a no-op. To reset a setting to inherit from the repository level, use the corresponding `<setting>_inherit` field.
type TestOptimizationUpdateServiceSettingsRequestAttributes struct {
	// Whether Auto Test Retries are enabled for this service. Setting to `null` is a no-op; use `auto_test_retries_enabled_inherit` to reset to repository-level inheritance.
	AutoTestRetriesEnabled *bool `json:"auto_test_retries_enabled,omitempty"`
	// When `true`, resets the Auto Test Retries setting to inherit from the repository level.
	AutoTestRetriesEnabledInherit *bool `json:"auto_test_retries_enabled_inherit,omitempty"`
	// Whether Code Coverage is enabled for this service. Setting to `null` is a no-op; use `code_coverage_enabled_inherit` to reset to repository-level inheritance.
	CodeCoverageEnabled *bool `json:"code_coverage_enabled,omitempty"`
	// When `true`, resets the Code Coverage setting to inherit from the repository level.
	CodeCoverageEnabledInherit *bool `json:"code_coverage_enabled_inherit,omitempty"`
	// Whether Early Flake Detection is enabled for this service. Setting to `null` is a no-op; use `early_flake_detection_enabled_inherit` to reset to repository-level inheritance.
	EarlyFlakeDetectionEnabled *bool `json:"early_flake_detection_enabled,omitempty"`
	// When `true`, resets the Early Flake Detection setting to inherit from the repository level.
	EarlyFlakeDetectionEnabledInherit *bool `json:"early_flake_detection_enabled_inherit,omitempty"`
	// The environment name. If omitted, defaults to `none`.
	Env *string `json:"env,omitempty"`
	// Whether Failed Test Replay is enabled for this service. Setting to `null` is a no-op; use `failed_test_replay_enabled_inherit` to reset to repository-level inheritance.
	FailedTestReplayEnabled *bool `json:"failed_test_replay_enabled,omitempty"`
	// When `true`, resets the Failed Test Replay setting to inherit from the repository level.
	FailedTestReplayEnabledInherit *bool `json:"failed_test_replay_enabled_inherit,omitempty"`
	// This field is ignored. PR Comments cannot be overridden at the service level.
	PrCommentsEnabled *bool `json:"pr_comments_enabled,omitempty"`
	// The repository identifier.
	RepositoryId string `json:"repository_id"`
	// The service name.
	ServiceName string `json:"service_name"`
	// Whether Test Impact Analysis is enabled for this service. Setting to `null` is a no-op; use `test_impact_analysis_enabled_inherit` to reset to repository-level inheritance.
	TestImpactAnalysisEnabled *bool `json:"test_impact_analysis_enabled,omitempty"`
	// When `true`, resets the Test Impact Analysis setting to inherit from the repository level.
	TestImpactAnalysisEnabledInherit *bool `json:"test_impact_analysis_enabled_inherit,omitempty"`
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

// GetAutoTestRetriesEnabledInherit returns the AutoTestRetriesEnabledInherit field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetAutoTestRetriesEnabledInherit() bool {
	if o == nil || o.AutoTestRetriesEnabledInherit == nil {
		var ret bool
		return ret
	}
	return *o.AutoTestRetriesEnabledInherit
}

// GetAutoTestRetriesEnabledInheritOk returns a tuple with the AutoTestRetriesEnabledInherit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetAutoTestRetriesEnabledInheritOk() (*bool, bool) {
	if o == nil || o.AutoTestRetriesEnabledInherit == nil {
		return nil, false
	}
	return o.AutoTestRetriesEnabledInherit, true
}

// HasAutoTestRetriesEnabledInherit returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasAutoTestRetriesEnabledInherit() bool {
	return o != nil && o.AutoTestRetriesEnabledInherit != nil
}

// SetAutoTestRetriesEnabledInherit gets a reference to the given bool and assigns it to the AutoTestRetriesEnabledInherit field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetAutoTestRetriesEnabledInherit(v bool) {
	o.AutoTestRetriesEnabledInherit = &v
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

// GetCodeCoverageEnabledInherit returns the CodeCoverageEnabledInherit field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetCodeCoverageEnabledInherit() bool {
	if o == nil || o.CodeCoverageEnabledInherit == nil {
		var ret bool
		return ret
	}
	return *o.CodeCoverageEnabledInherit
}

// GetCodeCoverageEnabledInheritOk returns a tuple with the CodeCoverageEnabledInherit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetCodeCoverageEnabledInheritOk() (*bool, bool) {
	if o == nil || o.CodeCoverageEnabledInherit == nil {
		return nil, false
	}
	return o.CodeCoverageEnabledInherit, true
}

// HasCodeCoverageEnabledInherit returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasCodeCoverageEnabledInherit() bool {
	return o != nil && o.CodeCoverageEnabledInherit != nil
}

// SetCodeCoverageEnabledInherit gets a reference to the given bool and assigns it to the CodeCoverageEnabledInherit field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetCodeCoverageEnabledInherit(v bool) {
	o.CodeCoverageEnabledInherit = &v
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

// GetEarlyFlakeDetectionEnabledInherit returns the EarlyFlakeDetectionEnabledInherit field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetEarlyFlakeDetectionEnabledInherit() bool {
	if o == nil || o.EarlyFlakeDetectionEnabledInherit == nil {
		var ret bool
		return ret
	}
	return *o.EarlyFlakeDetectionEnabledInherit
}

// GetEarlyFlakeDetectionEnabledInheritOk returns a tuple with the EarlyFlakeDetectionEnabledInherit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetEarlyFlakeDetectionEnabledInheritOk() (*bool, bool) {
	if o == nil || o.EarlyFlakeDetectionEnabledInherit == nil {
		return nil, false
	}
	return o.EarlyFlakeDetectionEnabledInherit, true
}

// HasEarlyFlakeDetectionEnabledInherit returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasEarlyFlakeDetectionEnabledInherit() bool {
	return o != nil && o.EarlyFlakeDetectionEnabledInherit != nil
}

// SetEarlyFlakeDetectionEnabledInherit gets a reference to the given bool and assigns it to the EarlyFlakeDetectionEnabledInherit field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetEarlyFlakeDetectionEnabledInherit(v bool) {
	o.EarlyFlakeDetectionEnabledInherit = &v
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

// GetFailedTestReplayEnabledInherit returns the FailedTestReplayEnabledInherit field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetFailedTestReplayEnabledInherit() bool {
	if o == nil || o.FailedTestReplayEnabledInherit == nil {
		var ret bool
		return ret
	}
	return *o.FailedTestReplayEnabledInherit
}

// GetFailedTestReplayEnabledInheritOk returns a tuple with the FailedTestReplayEnabledInherit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetFailedTestReplayEnabledInheritOk() (*bool, bool) {
	if o == nil || o.FailedTestReplayEnabledInherit == nil {
		return nil, false
	}
	return o.FailedTestReplayEnabledInherit, true
}

// HasFailedTestReplayEnabledInherit returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasFailedTestReplayEnabledInherit() bool {
	return o != nil && o.FailedTestReplayEnabledInherit != nil
}

// SetFailedTestReplayEnabledInherit gets a reference to the given bool and assigns it to the FailedTestReplayEnabledInherit field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetFailedTestReplayEnabledInherit(v bool) {
	o.FailedTestReplayEnabledInherit = &v
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

// GetTestImpactAnalysisEnabledInherit returns the TestImpactAnalysisEnabledInherit field value if set, zero value otherwise.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetTestImpactAnalysisEnabledInherit() bool {
	if o == nil || o.TestImpactAnalysisEnabledInherit == nil {
		var ret bool
		return ret
	}
	return *o.TestImpactAnalysisEnabledInherit
}

// GetTestImpactAnalysisEnabledInheritOk returns a tuple with the TestImpactAnalysisEnabledInherit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) GetTestImpactAnalysisEnabledInheritOk() (*bool, bool) {
	if o == nil || o.TestImpactAnalysisEnabledInherit == nil {
		return nil, false
	}
	return o.TestImpactAnalysisEnabledInherit, true
}

// HasTestImpactAnalysisEnabledInherit returns a boolean if a field has been set.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) HasTestImpactAnalysisEnabledInherit() bool {
	return o != nil && o.TestImpactAnalysisEnabledInherit != nil
}

// SetTestImpactAnalysisEnabledInherit gets a reference to the given bool and assigns it to the TestImpactAnalysisEnabledInherit field.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) SetTestImpactAnalysisEnabledInherit(v bool) {
	o.TestImpactAnalysisEnabledInherit = &v
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
	if o.AutoTestRetriesEnabledInherit != nil {
		toSerialize["auto_test_retries_enabled_inherit"] = o.AutoTestRetriesEnabledInherit
	}
	if o.CodeCoverageEnabled != nil {
		toSerialize["code_coverage_enabled"] = o.CodeCoverageEnabled
	}
	if o.CodeCoverageEnabledInherit != nil {
		toSerialize["code_coverage_enabled_inherit"] = o.CodeCoverageEnabledInherit
	}
	if o.EarlyFlakeDetectionEnabled != nil {
		toSerialize["early_flake_detection_enabled"] = o.EarlyFlakeDetectionEnabled
	}
	if o.EarlyFlakeDetectionEnabledInherit != nil {
		toSerialize["early_flake_detection_enabled_inherit"] = o.EarlyFlakeDetectionEnabledInherit
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.FailedTestReplayEnabled != nil {
		toSerialize["failed_test_replay_enabled"] = o.FailedTestReplayEnabled
	}
	if o.FailedTestReplayEnabledInherit != nil {
		toSerialize["failed_test_replay_enabled_inherit"] = o.FailedTestReplayEnabledInherit
	}
	if o.PrCommentsEnabled != nil {
		toSerialize["pr_comments_enabled"] = o.PrCommentsEnabled
	}
	toSerialize["repository_id"] = o.RepositoryId
	toSerialize["service_name"] = o.ServiceName
	if o.TestImpactAnalysisEnabled != nil {
		toSerialize["test_impact_analysis_enabled"] = o.TestImpactAnalysisEnabled
	}
	if o.TestImpactAnalysisEnabledInherit != nil {
		toSerialize["test_impact_analysis_enabled_inherit"] = o.TestImpactAnalysisEnabledInherit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationUpdateServiceSettingsRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoTestRetriesEnabled            *bool   `json:"auto_test_retries_enabled,omitempty"`
		AutoTestRetriesEnabledInherit     *bool   `json:"auto_test_retries_enabled_inherit,omitempty"`
		CodeCoverageEnabled               *bool   `json:"code_coverage_enabled,omitempty"`
		CodeCoverageEnabledInherit        *bool   `json:"code_coverage_enabled_inherit,omitempty"`
		EarlyFlakeDetectionEnabled        *bool   `json:"early_flake_detection_enabled,omitempty"`
		EarlyFlakeDetectionEnabledInherit *bool   `json:"early_flake_detection_enabled_inherit,omitempty"`
		Env                               *string `json:"env,omitempty"`
		FailedTestReplayEnabled           *bool   `json:"failed_test_replay_enabled,omitempty"`
		FailedTestReplayEnabledInherit    *bool   `json:"failed_test_replay_enabled_inherit,omitempty"`
		PrCommentsEnabled                 *bool   `json:"pr_comments_enabled,omitempty"`
		RepositoryId                      *string `json:"repository_id"`
		ServiceName                       *string `json:"service_name"`
		TestImpactAnalysisEnabled         *bool   `json:"test_impact_analysis_enabled,omitempty"`
		TestImpactAnalysisEnabledInherit  *bool   `json:"test_impact_analysis_enabled_inherit,omitempty"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_test_retries_enabled", "auto_test_retries_enabled_inherit", "code_coverage_enabled", "code_coverage_enabled_inherit", "early_flake_detection_enabled", "early_flake_detection_enabled_inherit", "env", "failed_test_replay_enabled", "failed_test_replay_enabled_inherit", "pr_comments_enabled", "repository_id", "service_name", "test_impact_analysis_enabled", "test_impact_analysis_enabled_inherit"})
	} else {
		return err
	}
	o.AutoTestRetriesEnabled = all.AutoTestRetriesEnabled
	o.AutoTestRetriesEnabledInherit = all.AutoTestRetriesEnabledInherit
	o.CodeCoverageEnabled = all.CodeCoverageEnabled
	o.CodeCoverageEnabledInherit = all.CodeCoverageEnabledInherit
	o.EarlyFlakeDetectionEnabled = all.EarlyFlakeDetectionEnabled
	o.EarlyFlakeDetectionEnabledInherit = all.EarlyFlakeDetectionEnabledInherit
	o.Env = all.Env
	o.FailedTestReplayEnabled = all.FailedTestReplayEnabled
	o.FailedTestReplayEnabledInherit = all.FailedTestReplayEnabledInherit
	o.PrCommentsEnabled = all.PrCommentsEnabled
	o.RepositoryId = *all.RepositoryId
	o.ServiceName = *all.ServiceName
	o.TestImpactAnalysisEnabled = all.TestImpactAnalysisEnabled
	o.TestImpactAnalysisEnabledInherit = all.TestImpactAnalysisEnabledInherit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
