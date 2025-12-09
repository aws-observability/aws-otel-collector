// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestAttributes Attributes of a flaky test.
type FlakyTestAttributes struct {
	// Unique identifier for the attempt to fix this flaky test. Use this ID in the Git commit message in order to trigger the attempt to fix workflow.
	//
	// When the workflow is triggered the test is automatically retried by the tracer a certain number of configurable times. When all retries pass, the test is automatically marked as fixed in Flaky Test Management.
	// Test runs are tagged with @test.test_management.attempt_to_fix_passed and @test.test_management.is_attempt_to_fix when the attempt to fix workflow is triggered.
	AttemptToFixId *string `json:"attempt_to_fix_id,omitempty"`
	// The name of the test's code owners as inferred from the repository configuration.
	Codeowners []string `json:"codeowners,omitempty"`
	// List of environments where this test has been flaky.
	Envs []string `json:"envs,omitempty"`
	// The branch name where the test exhibited flakiness for the first time.
	FirstFlakedBranch *string `json:"first_flaked_branch,omitempty"`
	// The commit SHA where the test exhibited flakiness for the first time.
	FirstFlakedSha *string `json:"first_flaked_sha,omitempty"`
	// Unix timestamp when the test exhibited flakiness for the first time.
	FirstFlakedTs *int64 `json:"first_flaked_ts,omitempty"`
	// The category of a flaky test.
	FlakyCategory datadog.NullableString `json:"flaky_category,omitempty"`
	// The current state of the flaky test.
	FlakyState *FlakyTestAttributesFlakyState `json:"flaky_state,omitempty"`
	// The branch name where the test exhibited flakiness for the last time.
	LastFlakedBranch *string `json:"last_flaked_branch,omitempty"`
	// The commit SHA where the test exhibited flakiness for the last time.
	LastFlakedSha *string `json:"last_flaked_sha,omitempty"`
	// Unix timestamp when the test exhibited flakiness for the last time.
	LastFlakedTs *int64 `json:"last_flaked_ts,omitempty"`
	// The name of the test module. The definition of module changes slightly per language:
	// - In .NET, a test module groups every test that is run under the same unit test project.
	// - In Swift, a test module groups every test that is run for a given bundle.
	// - In JavaScript, the test modules map one-to-one to test sessions.
	// - In Java, a test module groups every test that is run by the same Maven Surefire/Failsafe or Gradle Test task execution.
	// - In Python, a test module groups every test that is run under the same `.py` file as part of a test suite, which is typically managed by a framework like `unittest` or `pytest`.
	// - In Ruby, a test module groups every test that is run within the same test file, which is typically managed by a framework like `RSpec` or `Minitest`.
	Module datadog.NullableString `json:"module,omitempty"`
	// The test name. A concise name for a test case. Defined in the test itself.
	Name *string `json:"name,omitempty"`
	// CI pipeline related statistics for the flaky test. This information is only available if test runs are associated with CI pipeline events from CI Visibility.
	PipelineStats *FlakyTestPipelineStats `json:"pipeline_stats,omitempty"`
	// List of test service names where this test has been flaky.
	//
	// A test service is a group of tests associated with a project or repository. It contains all the individual tests for your code, optionally organized into test suites, which are like folders for your tests.
	Services []string `json:"services,omitempty"`
	// The name of the test suite. A group of tests exercising the same unit of code depending on your language and testing framework.
	Suite *string `json:"suite,omitempty"`
	// Metadata about the latest failed test run of the flaky test.
	TestRunMetadata *FlakyTestRunMetadata `json:"test_run_metadata,omitempty"`
	// Test statistics for the flaky test.
	TestStats *FlakyTestStats `json:"test_stats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFlakyTestAttributes instantiates a new FlakyTestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlakyTestAttributes() *FlakyTestAttributes {
	this := FlakyTestAttributes{}
	return &this
}

// NewFlakyTestAttributesWithDefaults instantiates a new FlakyTestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlakyTestAttributesWithDefaults() *FlakyTestAttributes {
	this := FlakyTestAttributes{}
	return &this
}

// GetAttemptToFixId returns the AttemptToFixId field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetAttemptToFixId() string {
	if o == nil || o.AttemptToFixId == nil {
		var ret string
		return ret
	}
	return *o.AttemptToFixId
}

// GetAttemptToFixIdOk returns a tuple with the AttemptToFixId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetAttemptToFixIdOk() (*string, bool) {
	if o == nil || o.AttemptToFixId == nil {
		return nil, false
	}
	return o.AttemptToFixId, true
}

// HasAttemptToFixId returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasAttemptToFixId() bool {
	return o != nil && o.AttemptToFixId != nil
}

// SetAttemptToFixId gets a reference to the given string and assigns it to the AttemptToFixId field.
func (o *FlakyTestAttributes) SetAttemptToFixId(v string) {
	o.AttemptToFixId = &v
}

// GetCodeowners returns the Codeowners field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetCodeowners() []string {
	if o == nil || o.Codeowners == nil {
		var ret []string
		return ret
	}
	return o.Codeowners
}

// GetCodeownersOk returns a tuple with the Codeowners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetCodeownersOk() (*[]string, bool) {
	if o == nil || o.Codeowners == nil {
		return nil, false
	}
	return &o.Codeowners, true
}

// HasCodeowners returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasCodeowners() bool {
	return o != nil && o.Codeowners != nil
}

// SetCodeowners gets a reference to the given []string and assigns it to the Codeowners field.
func (o *FlakyTestAttributes) SetCodeowners(v []string) {
	o.Codeowners = v
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetEnvs() []string {
	if o == nil || o.Envs == nil {
		var ret []string
		return ret
	}
	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetEnvsOk() (*[]string, bool) {
	if o == nil || o.Envs == nil {
		return nil, false
	}
	return &o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasEnvs() bool {
	return o != nil && o.Envs != nil
}

// SetEnvs gets a reference to the given []string and assigns it to the Envs field.
func (o *FlakyTestAttributes) SetEnvs(v []string) {
	o.Envs = v
}

// GetFirstFlakedBranch returns the FirstFlakedBranch field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetFirstFlakedBranch() string {
	if o == nil || o.FirstFlakedBranch == nil {
		var ret string
		return ret
	}
	return *o.FirstFlakedBranch
}

// GetFirstFlakedBranchOk returns a tuple with the FirstFlakedBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetFirstFlakedBranchOk() (*string, bool) {
	if o == nil || o.FirstFlakedBranch == nil {
		return nil, false
	}
	return o.FirstFlakedBranch, true
}

// HasFirstFlakedBranch returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasFirstFlakedBranch() bool {
	return o != nil && o.FirstFlakedBranch != nil
}

// SetFirstFlakedBranch gets a reference to the given string and assigns it to the FirstFlakedBranch field.
func (o *FlakyTestAttributes) SetFirstFlakedBranch(v string) {
	o.FirstFlakedBranch = &v
}

// GetFirstFlakedSha returns the FirstFlakedSha field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetFirstFlakedSha() string {
	if o == nil || o.FirstFlakedSha == nil {
		var ret string
		return ret
	}
	return *o.FirstFlakedSha
}

// GetFirstFlakedShaOk returns a tuple with the FirstFlakedSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetFirstFlakedShaOk() (*string, bool) {
	if o == nil || o.FirstFlakedSha == nil {
		return nil, false
	}
	return o.FirstFlakedSha, true
}

// HasFirstFlakedSha returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasFirstFlakedSha() bool {
	return o != nil && o.FirstFlakedSha != nil
}

// SetFirstFlakedSha gets a reference to the given string and assigns it to the FirstFlakedSha field.
func (o *FlakyTestAttributes) SetFirstFlakedSha(v string) {
	o.FirstFlakedSha = &v
}

// GetFirstFlakedTs returns the FirstFlakedTs field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetFirstFlakedTs() int64 {
	if o == nil || o.FirstFlakedTs == nil {
		var ret int64
		return ret
	}
	return *o.FirstFlakedTs
}

// GetFirstFlakedTsOk returns a tuple with the FirstFlakedTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetFirstFlakedTsOk() (*int64, bool) {
	if o == nil || o.FirstFlakedTs == nil {
		return nil, false
	}
	return o.FirstFlakedTs, true
}

// HasFirstFlakedTs returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasFirstFlakedTs() bool {
	return o != nil && o.FirstFlakedTs != nil
}

// SetFirstFlakedTs gets a reference to the given int64 and assigns it to the FirstFlakedTs field.
func (o *FlakyTestAttributes) SetFirstFlakedTs(v int64) {
	o.FirstFlakedTs = &v
}

// GetFlakyCategory returns the FlakyCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestAttributes) GetFlakyCategory() string {
	if o == nil || o.FlakyCategory.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlakyCategory.Get()
}

// GetFlakyCategoryOk returns a tuple with the FlakyCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestAttributes) GetFlakyCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlakyCategory.Get(), o.FlakyCategory.IsSet()
}

// HasFlakyCategory returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasFlakyCategory() bool {
	return o != nil && o.FlakyCategory.IsSet()
}

// SetFlakyCategory gets a reference to the given datadog.NullableString and assigns it to the FlakyCategory field.
func (o *FlakyTestAttributes) SetFlakyCategory(v string) {
	o.FlakyCategory.Set(&v)
}

// SetFlakyCategoryNil sets the value for FlakyCategory to be an explicit nil.
func (o *FlakyTestAttributes) SetFlakyCategoryNil() {
	o.FlakyCategory.Set(nil)
}

// UnsetFlakyCategory ensures that no value is present for FlakyCategory, not even an explicit nil.
func (o *FlakyTestAttributes) UnsetFlakyCategory() {
	o.FlakyCategory.Unset()
}

// GetFlakyState returns the FlakyState field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetFlakyState() FlakyTestAttributesFlakyState {
	if o == nil || o.FlakyState == nil {
		var ret FlakyTestAttributesFlakyState
		return ret
	}
	return *o.FlakyState
}

// GetFlakyStateOk returns a tuple with the FlakyState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetFlakyStateOk() (*FlakyTestAttributesFlakyState, bool) {
	if o == nil || o.FlakyState == nil {
		return nil, false
	}
	return o.FlakyState, true
}

// HasFlakyState returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasFlakyState() bool {
	return o != nil && o.FlakyState != nil
}

// SetFlakyState gets a reference to the given FlakyTestAttributesFlakyState and assigns it to the FlakyState field.
func (o *FlakyTestAttributes) SetFlakyState(v FlakyTestAttributesFlakyState) {
	o.FlakyState = &v
}

// GetLastFlakedBranch returns the LastFlakedBranch field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetLastFlakedBranch() string {
	if o == nil || o.LastFlakedBranch == nil {
		var ret string
		return ret
	}
	return *o.LastFlakedBranch
}

// GetLastFlakedBranchOk returns a tuple with the LastFlakedBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetLastFlakedBranchOk() (*string, bool) {
	if o == nil || o.LastFlakedBranch == nil {
		return nil, false
	}
	return o.LastFlakedBranch, true
}

// HasLastFlakedBranch returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasLastFlakedBranch() bool {
	return o != nil && o.LastFlakedBranch != nil
}

// SetLastFlakedBranch gets a reference to the given string and assigns it to the LastFlakedBranch field.
func (o *FlakyTestAttributes) SetLastFlakedBranch(v string) {
	o.LastFlakedBranch = &v
}

// GetLastFlakedSha returns the LastFlakedSha field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetLastFlakedSha() string {
	if o == nil || o.LastFlakedSha == nil {
		var ret string
		return ret
	}
	return *o.LastFlakedSha
}

// GetLastFlakedShaOk returns a tuple with the LastFlakedSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetLastFlakedShaOk() (*string, bool) {
	if o == nil || o.LastFlakedSha == nil {
		return nil, false
	}
	return o.LastFlakedSha, true
}

// HasLastFlakedSha returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasLastFlakedSha() bool {
	return o != nil && o.LastFlakedSha != nil
}

// SetLastFlakedSha gets a reference to the given string and assigns it to the LastFlakedSha field.
func (o *FlakyTestAttributes) SetLastFlakedSha(v string) {
	o.LastFlakedSha = &v
}

// GetLastFlakedTs returns the LastFlakedTs field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetLastFlakedTs() int64 {
	if o == nil || o.LastFlakedTs == nil {
		var ret int64
		return ret
	}
	return *o.LastFlakedTs
}

// GetLastFlakedTsOk returns a tuple with the LastFlakedTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetLastFlakedTsOk() (*int64, bool) {
	if o == nil || o.LastFlakedTs == nil {
		return nil, false
	}
	return o.LastFlakedTs, true
}

// HasLastFlakedTs returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasLastFlakedTs() bool {
	return o != nil && o.LastFlakedTs != nil
}

// SetLastFlakedTs gets a reference to the given int64 and assigns it to the LastFlakedTs field.
func (o *FlakyTestAttributes) SetLastFlakedTs(v int64) {
	o.LastFlakedTs = &v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestAttributes) GetModule() string {
	if o == nil || o.Module.Get() == nil {
		var ret string
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestAttributes) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasModule() bool {
	return o != nil && o.Module.IsSet()
}

// SetModule gets a reference to the given datadog.NullableString and assigns it to the Module field.
func (o *FlakyTestAttributes) SetModule(v string) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil.
func (o *FlakyTestAttributes) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil.
func (o *FlakyTestAttributes) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FlakyTestAttributes) SetName(v string) {
	o.Name = &v
}

// GetPipelineStats returns the PipelineStats field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetPipelineStats() FlakyTestPipelineStats {
	if o == nil || o.PipelineStats == nil {
		var ret FlakyTestPipelineStats
		return ret
	}
	return *o.PipelineStats
}

// GetPipelineStatsOk returns a tuple with the PipelineStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetPipelineStatsOk() (*FlakyTestPipelineStats, bool) {
	if o == nil || o.PipelineStats == nil {
		return nil, false
	}
	return o.PipelineStats, true
}

// HasPipelineStats returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasPipelineStats() bool {
	return o != nil && o.PipelineStats != nil
}

// SetPipelineStats gets a reference to the given FlakyTestPipelineStats and assigns it to the PipelineStats field.
func (o *FlakyTestAttributes) SetPipelineStats(v FlakyTestPipelineStats) {
	o.PipelineStats = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *FlakyTestAttributes) SetServices(v []string) {
	o.Services = v
}

// GetSuite returns the Suite field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetSuite() string {
	if o == nil || o.Suite == nil {
		var ret string
		return ret
	}
	return *o.Suite
}

// GetSuiteOk returns a tuple with the Suite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetSuiteOk() (*string, bool) {
	if o == nil || o.Suite == nil {
		return nil, false
	}
	return o.Suite, true
}

// HasSuite returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasSuite() bool {
	return o != nil && o.Suite != nil
}

// SetSuite gets a reference to the given string and assigns it to the Suite field.
func (o *FlakyTestAttributes) SetSuite(v string) {
	o.Suite = &v
}

// GetTestRunMetadata returns the TestRunMetadata field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetTestRunMetadata() FlakyTestRunMetadata {
	if o == nil || o.TestRunMetadata == nil {
		var ret FlakyTestRunMetadata
		return ret
	}
	return *o.TestRunMetadata
}

// GetTestRunMetadataOk returns a tuple with the TestRunMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetTestRunMetadataOk() (*FlakyTestRunMetadata, bool) {
	if o == nil || o.TestRunMetadata == nil {
		return nil, false
	}
	return o.TestRunMetadata, true
}

// HasTestRunMetadata returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasTestRunMetadata() bool {
	return o != nil && o.TestRunMetadata != nil
}

// SetTestRunMetadata gets a reference to the given FlakyTestRunMetadata and assigns it to the TestRunMetadata field.
func (o *FlakyTestAttributes) SetTestRunMetadata(v FlakyTestRunMetadata) {
	o.TestRunMetadata = &v
}

// GetTestStats returns the TestStats field value if set, zero value otherwise.
func (o *FlakyTestAttributes) GetTestStats() FlakyTestStats {
	if o == nil || o.TestStats == nil {
		var ret FlakyTestStats
		return ret
	}
	return *o.TestStats
}

// GetTestStatsOk returns a tuple with the TestStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestAttributes) GetTestStatsOk() (*FlakyTestStats, bool) {
	if o == nil || o.TestStats == nil {
		return nil, false
	}
	return o.TestStats, true
}

// HasTestStats returns a boolean if a field has been set.
func (o *FlakyTestAttributes) HasTestStats() bool {
	return o != nil && o.TestStats != nil
}

// SetTestStats gets a reference to the given FlakyTestStats and assigns it to the TestStats field.
func (o *FlakyTestAttributes) SetTestStats(v FlakyTestStats) {
	o.TestStats = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FlakyTestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AttemptToFixId != nil {
		toSerialize["attempt_to_fix_id"] = o.AttemptToFixId
	}
	if o.Codeowners != nil {
		toSerialize["codeowners"] = o.Codeowners
	}
	if o.Envs != nil {
		toSerialize["envs"] = o.Envs
	}
	if o.FirstFlakedBranch != nil {
		toSerialize["first_flaked_branch"] = o.FirstFlakedBranch
	}
	if o.FirstFlakedSha != nil {
		toSerialize["first_flaked_sha"] = o.FirstFlakedSha
	}
	if o.FirstFlakedTs != nil {
		toSerialize["first_flaked_ts"] = o.FirstFlakedTs
	}
	if o.FlakyCategory.IsSet() {
		toSerialize["flaky_category"] = o.FlakyCategory.Get()
	}
	if o.FlakyState != nil {
		toSerialize["flaky_state"] = o.FlakyState
	}
	if o.LastFlakedBranch != nil {
		toSerialize["last_flaked_branch"] = o.LastFlakedBranch
	}
	if o.LastFlakedSha != nil {
		toSerialize["last_flaked_sha"] = o.LastFlakedSha
	}
	if o.LastFlakedTs != nil {
		toSerialize["last_flaked_ts"] = o.LastFlakedTs
	}
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PipelineStats != nil {
		toSerialize["pipeline_stats"] = o.PipelineStats
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Suite != nil {
		toSerialize["suite"] = o.Suite
	}
	if o.TestRunMetadata != nil {
		toSerialize["test_run_metadata"] = o.TestRunMetadata
	}
	if o.TestStats != nil {
		toSerialize["test_stats"] = o.TestStats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FlakyTestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AttemptToFixId    *string                        `json:"attempt_to_fix_id,omitempty"`
		Codeowners        []string                       `json:"codeowners,omitempty"`
		Envs              []string                       `json:"envs,omitempty"`
		FirstFlakedBranch *string                        `json:"first_flaked_branch,omitempty"`
		FirstFlakedSha    *string                        `json:"first_flaked_sha,omitempty"`
		FirstFlakedTs     *int64                         `json:"first_flaked_ts,omitempty"`
		FlakyCategory     datadog.NullableString         `json:"flaky_category,omitempty"`
		FlakyState        *FlakyTestAttributesFlakyState `json:"flaky_state,omitempty"`
		LastFlakedBranch  *string                        `json:"last_flaked_branch,omitempty"`
		LastFlakedSha     *string                        `json:"last_flaked_sha,omitempty"`
		LastFlakedTs      *int64                         `json:"last_flaked_ts,omitempty"`
		Module            datadog.NullableString         `json:"module,omitempty"`
		Name              *string                        `json:"name,omitempty"`
		PipelineStats     *FlakyTestPipelineStats        `json:"pipeline_stats,omitempty"`
		Services          []string                       `json:"services,omitempty"`
		Suite             *string                        `json:"suite,omitempty"`
		TestRunMetadata   *FlakyTestRunMetadata          `json:"test_run_metadata,omitempty"`
		TestStats         *FlakyTestStats                `json:"test_stats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attempt_to_fix_id", "codeowners", "envs", "first_flaked_branch", "first_flaked_sha", "first_flaked_ts", "flaky_category", "flaky_state", "last_flaked_branch", "last_flaked_sha", "last_flaked_ts", "module", "name", "pipeline_stats", "services", "suite", "test_run_metadata", "test_stats"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AttemptToFixId = all.AttemptToFixId
	o.Codeowners = all.Codeowners
	o.Envs = all.Envs
	o.FirstFlakedBranch = all.FirstFlakedBranch
	o.FirstFlakedSha = all.FirstFlakedSha
	o.FirstFlakedTs = all.FirstFlakedTs
	o.FlakyCategory = all.FlakyCategory
	if all.FlakyState != nil && !all.FlakyState.IsValid() {
		hasInvalidField = true
	} else {
		o.FlakyState = all.FlakyState
	}
	o.LastFlakedBranch = all.LastFlakedBranch
	o.LastFlakedSha = all.LastFlakedSha
	o.LastFlakedTs = all.LastFlakedTs
	o.Module = all.Module
	o.Name = all.Name
	if all.PipelineStats != nil && all.PipelineStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PipelineStats = all.PipelineStats
	o.Services = all.Services
	o.Suite = all.Suite
	if all.TestRunMetadata != nil && all.TestRunMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TestRunMetadata = all.TestRunMetadata
	if all.TestStats != nil && all.TestStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TestStats = all.TestStats

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
