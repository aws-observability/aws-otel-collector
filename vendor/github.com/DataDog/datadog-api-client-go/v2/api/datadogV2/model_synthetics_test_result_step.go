// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultStep A step result from a browser, mobile, or multistep API test.
type SyntheticsTestResultStep struct {
	// Whether the test continues when this step fails.
	AllowFailure *bool `json:"allow_failure,omitempty"`
	// Inner API test definition for browser `runApiTest` steps.
	ApiTest map[string]interface{} `json:"api_test,omitempty"`
	// Assertion result for a browser or mobile step.
	AssertionResult *SyntheticsTestResultStepAssertionResult `json:"assertion_result,omitempty"`
	// Assertion results produced by the step.
	Assertions []SyntheticsTestResultAssertionResult `json:"assertions,omitempty"`
	// URLs of requests blocked during the step.
	BlockedRequestsUrls []string `json:"blocked_requests_urls,omitempty"`
	// Bounding box of an element on the page.
	Bounds *SyntheticsTestResultBounds `json:"bounds,omitempty"`
	// Browser errors captured during the step.
	BrowserErrors []SyntheticsTestResultBrowserError `json:"browser_errors,omitempty"`
	// Storage bucket keys for artifacts produced during a step or test.
	BucketKeys *SyntheticsTestResultBucketKeys `json:"bucket_keys,omitempty"`
	// CDN resources encountered during the step.
	CdnResources []SyntheticsTestResultCdnResource `json:"cdn_resources,omitempty"`
	// Click type performed in a browser step.
	ClickType *string `json:"click_type,omitempty"`
	// Compressed JSON descriptor for the step (internal format).
	CompressedJsonDescriptor *string `json:"compressed_json_descriptor,omitempty"`
	// Request configuration executed by this step (API test steps).
	Config map[string]interface{} `json:"config,omitempty"`
	// Human-readable description of the step.
	Description *string `json:"description,omitempty"`
	// Duration of the step in milliseconds.
	Duration *float64 `json:"duration,omitempty"`
	// Description of the element interacted with by the step.
	ElementDescription *string `json:"element_description,omitempty"`
	// Element locator updates produced during a step.
	ElementUpdates *SyntheticsTestResultStepElementUpdates `json:"element_updates,omitempty"`
	// A variable used or extracted during a test.
	ExtractedValue *SyntheticsTestResultVariable `json:"extracted_value,omitempty"`
	// Details about the failure of a Synthetic test.
	Failure *SyntheticsTestResultFailure `json:"failure,omitempty"`
	// HTTP results produced by an MCP step.
	HttpResults []SyntheticsTestResultAssertionResult `json:"http_results,omitempty"`
	// Identifier of the step.
	Id *string `json:"id,omitempty"`
	// Whether this step is critical for the test outcome.
	IsCritical *bool `json:"is_critical,omitempty"`
	// Whether the step uses a custom JavaScript assertion.
	JavascriptCustomAssertionCode *bool `json:"javascript_custom_assertion_code,omitempty"`
	// Time taken to locate the element in milliseconds.
	LocateElementDuration *float64 `json:"locate_element_duration,omitempty"`
	// Name of the step.
	Name *string `json:"name,omitempty"`
	// Details of the outgoing request made during the test execution.
	Request *SyntheticsTestResultRequestInfo `json:"request,omitempty"`
	// Details of the response received during the test execution.
	Response *SyntheticsTestResultResponseInfo `json:"response,omitempty"`
	// Retry results for the step.
	Retries []SyntheticsTestResultStep `json:"retries,omitempty"`
	// Number of times this step was retried.
	RetryCount *int64 `json:"retry_count,omitempty"`
	// RUM application context associated with a step or sub-test.
	RumContext *SyntheticsTestResultRumContext `json:"rum_context,omitempty"`
	// Unix timestamp (ms) of when the step started.
	StartedAt *int64 `json:"started_at,omitempty"`
	// Status of the step (for example, `passed`, `failed`).
	Status *string `json:"status,omitempty"`
	// Information about a sub-step in a nested test execution.
	SubStep *SyntheticsTestResultSubStep `json:"sub_step,omitempty"`
	// Information about a sub-test played from a parent browser test.
	SubTest *SyntheticsTestResultSubTest `json:"sub_test,omitempty"`
	// Subtype of the step.
	Subtype *string `json:"subtype,omitempty"`
	// Browser tabs involved in the step.
	Tabs []SyntheticsTestResultTab `json:"tabs,omitempty"`
	// Timing breakdown of the step execution.
	Timings map[string]interface{} `json:"timings,omitempty"`
	// Whether the step was executed through a Synthetics tunnel.
	Tunnel *bool `json:"tunnel,omitempty"`
	// Type of the step (for example, `click`, `assertElementContent`, `runApiTest`).
	Type *string `json:"type,omitempty"`
	// URL associated with the step (for navigation steps).
	Url *string `json:"url,omitempty"`
	// Step value. Its type depends on the step type.
	Value interface{} `json:"value,omitempty"`
	// Variables captured during a test step.
	Variables *SyntheticsTestResultVariables `json:"variables,omitempty"`
	// Web vitals metrics captured during the step.
	VitalsMetrics []SyntheticsTestResultVitalsMetrics `json:"vitals_metrics,omitempty"`
	// Warnings emitted during the step.
	Warnings []SyntheticsTestResultWarning `json:"warnings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultStep instantiates a new SyntheticsTestResultStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultStep() *SyntheticsTestResultStep {
	this := SyntheticsTestResultStep{}
	return &this
}

// NewSyntheticsTestResultStepWithDefaults instantiates a new SyntheticsTestResultStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultStepWithDefaults() *SyntheticsTestResultStep {
	this := SyntheticsTestResultStep{}
	return &this
}

// GetAllowFailure returns the AllowFailure field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetAllowFailure() bool {
	if o == nil || o.AllowFailure == nil {
		var ret bool
		return ret
	}
	return *o.AllowFailure
}

// GetAllowFailureOk returns a tuple with the AllowFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetAllowFailureOk() (*bool, bool) {
	if o == nil || o.AllowFailure == nil {
		return nil, false
	}
	return o.AllowFailure, true
}

// HasAllowFailure returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasAllowFailure() bool {
	return o != nil && o.AllowFailure != nil
}

// SetAllowFailure gets a reference to the given bool and assigns it to the AllowFailure field.
func (o *SyntheticsTestResultStep) SetAllowFailure(v bool) {
	o.AllowFailure = &v
}

// GetApiTest returns the ApiTest field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetApiTest() map[string]interface{} {
	if o == nil || o.ApiTest == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ApiTest
}

// GetApiTestOk returns a tuple with the ApiTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetApiTestOk() (*map[string]interface{}, bool) {
	if o == nil || o.ApiTest == nil {
		return nil, false
	}
	return &o.ApiTest, true
}

// HasApiTest returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasApiTest() bool {
	return o != nil && o.ApiTest != nil
}

// SetApiTest gets a reference to the given map[string]interface{} and assigns it to the ApiTest field.
func (o *SyntheticsTestResultStep) SetApiTest(v map[string]interface{}) {
	o.ApiTest = v
}

// GetAssertionResult returns the AssertionResult field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetAssertionResult() SyntheticsTestResultStepAssertionResult {
	if o == nil || o.AssertionResult == nil {
		var ret SyntheticsTestResultStepAssertionResult
		return ret
	}
	return *o.AssertionResult
}

// GetAssertionResultOk returns a tuple with the AssertionResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetAssertionResultOk() (*SyntheticsTestResultStepAssertionResult, bool) {
	if o == nil || o.AssertionResult == nil {
		return nil, false
	}
	return o.AssertionResult, true
}

// HasAssertionResult returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasAssertionResult() bool {
	return o != nil && o.AssertionResult != nil
}

// SetAssertionResult gets a reference to the given SyntheticsTestResultStepAssertionResult and assigns it to the AssertionResult field.
func (o *SyntheticsTestResultStep) SetAssertionResult(v SyntheticsTestResultStepAssertionResult) {
	o.AssertionResult = &v
}

// GetAssertions returns the Assertions field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetAssertions() []SyntheticsTestResultAssertionResult {
	if o == nil || o.Assertions == nil {
		var ret []SyntheticsTestResultAssertionResult
		return ret
	}
	return o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetAssertionsOk() (*[]SyntheticsTestResultAssertionResult, bool) {
	if o == nil || o.Assertions == nil {
		return nil, false
	}
	return &o.Assertions, true
}

// HasAssertions returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasAssertions() bool {
	return o != nil && o.Assertions != nil
}

// SetAssertions gets a reference to the given []SyntheticsTestResultAssertionResult and assigns it to the Assertions field.
func (o *SyntheticsTestResultStep) SetAssertions(v []SyntheticsTestResultAssertionResult) {
	o.Assertions = v
}

// GetBlockedRequestsUrls returns the BlockedRequestsUrls field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetBlockedRequestsUrls() []string {
	if o == nil || o.BlockedRequestsUrls == nil {
		var ret []string
		return ret
	}
	return o.BlockedRequestsUrls
}

// GetBlockedRequestsUrlsOk returns a tuple with the BlockedRequestsUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetBlockedRequestsUrlsOk() (*[]string, bool) {
	if o == nil || o.BlockedRequestsUrls == nil {
		return nil, false
	}
	return &o.BlockedRequestsUrls, true
}

// HasBlockedRequestsUrls returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasBlockedRequestsUrls() bool {
	return o != nil && o.BlockedRequestsUrls != nil
}

// SetBlockedRequestsUrls gets a reference to the given []string and assigns it to the BlockedRequestsUrls field.
func (o *SyntheticsTestResultStep) SetBlockedRequestsUrls(v []string) {
	o.BlockedRequestsUrls = v
}

// GetBounds returns the Bounds field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetBounds() SyntheticsTestResultBounds {
	if o == nil || o.Bounds == nil {
		var ret SyntheticsTestResultBounds
		return ret
	}
	return *o.Bounds
}

// GetBoundsOk returns a tuple with the Bounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetBoundsOk() (*SyntheticsTestResultBounds, bool) {
	if o == nil || o.Bounds == nil {
		return nil, false
	}
	return o.Bounds, true
}

// HasBounds returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasBounds() bool {
	return o != nil && o.Bounds != nil
}

// SetBounds gets a reference to the given SyntheticsTestResultBounds and assigns it to the Bounds field.
func (o *SyntheticsTestResultStep) SetBounds(v SyntheticsTestResultBounds) {
	o.Bounds = &v
}

// GetBrowserErrors returns the BrowserErrors field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetBrowserErrors() []SyntheticsTestResultBrowserError {
	if o == nil || o.BrowserErrors == nil {
		var ret []SyntheticsTestResultBrowserError
		return ret
	}
	return o.BrowserErrors
}

// GetBrowserErrorsOk returns a tuple with the BrowserErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetBrowserErrorsOk() (*[]SyntheticsTestResultBrowserError, bool) {
	if o == nil || o.BrowserErrors == nil {
		return nil, false
	}
	return &o.BrowserErrors, true
}

// HasBrowserErrors returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasBrowserErrors() bool {
	return o != nil && o.BrowserErrors != nil
}

// SetBrowserErrors gets a reference to the given []SyntheticsTestResultBrowserError and assigns it to the BrowserErrors field.
func (o *SyntheticsTestResultStep) SetBrowserErrors(v []SyntheticsTestResultBrowserError) {
	o.BrowserErrors = v
}

// GetBucketKeys returns the BucketKeys field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetBucketKeys() SyntheticsTestResultBucketKeys {
	if o == nil || o.BucketKeys == nil {
		var ret SyntheticsTestResultBucketKeys
		return ret
	}
	return *o.BucketKeys
}

// GetBucketKeysOk returns a tuple with the BucketKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetBucketKeysOk() (*SyntheticsTestResultBucketKeys, bool) {
	if o == nil || o.BucketKeys == nil {
		return nil, false
	}
	return o.BucketKeys, true
}

// HasBucketKeys returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasBucketKeys() bool {
	return o != nil && o.BucketKeys != nil
}

// SetBucketKeys gets a reference to the given SyntheticsTestResultBucketKeys and assigns it to the BucketKeys field.
func (o *SyntheticsTestResultStep) SetBucketKeys(v SyntheticsTestResultBucketKeys) {
	o.BucketKeys = &v
}

// GetCdnResources returns the CdnResources field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetCdnResources() []SyntheticsTestResultCdnResource {
	if o == nil || o.CdnResources == nil {
		var ret []SyntheticsTestResultCdnResource
		return ret
	}
	return o.CdnResources
}

// GetCdnResourcesOk returns a tuple with the CdnResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetCdnResourcesOk() (*[]SyntheticsTestResultCdnResource, bool) {
	if o == nil || o.CdnResources == nil {
		return nil, false
	}
	return &o.CdnResources, true
}

// HasCdnResources returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasCdnResources() bool {
	return o != nil && o.CdnResources != nil
}

// SetCdnResources gets a reference to the given []SyntheticsTestResultCdnResource and assigns it to the CdnResources field.
func (o *SyntheticsTestResultStep) SetCdnResources(v []SyntheticsTestResultCdnResource) {
	o.CdnResources = v
}

// GetClickType returns the ClickType field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetClickType() string {
	if o == nil || o.ClickType == nil {
		var ret string
		return ret
	}
	return *o.ClickType
}

// GetClickTypeOk returns a tuple with the ClickType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetClickTypeOk() (*string, bool) {
	if o == nil || o.ClickType == nil {
		return nil, false
	}
	return o.ClickType, true
}

// HasClickType returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasClickType() bool {
	return o != nil && o.ClickType != nil
}

// SetClickType gets a reference to the given string and assigns it to the ClickType field.
func (o *SyntheticsTestResultStep) SetClickType(v string) {
	o.ClickType = &v
}

// GetCompressedJsonDescriptor returns the CompressedJsonDescriptor field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetCompressedJsonDescriptor() string {
	if o == nil || o.CompressedJsonDescriptor == nil {
		var ret string
		return ret
	}
	return *o.CompressedJsonDescriptor
}

// GetCompressedJsonDescriptorOk returns a tuple with the CompressedJsonDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetCompressedJsonDescriptorOk() (*string, bool) {
	if o == nil || o.CompressedJsonDescriptor == nil {
		return nil, false
	}
	return o.CompressedJsonDescriptor, true
}

// HasCompressedJsonDescriptor returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasCompressedJsonDescriptor() bool {
	return o != nil && o.CompressedJsonDescriptor != nil
}

// SetCompressedJsonDescriptor gets a reference to the given string and assigns it to the CompressedJsonDescriptor field.
func (o *SyntheticsTestResultStep) SetCompressedJsonDescriptor(v string) {
	o.CompressedJsonDescriptor = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *SyntheticsTestResultStep) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SyntheticsTestResultStep) SetDescription(v string) {
	o.Description = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetDuration() float64 {
	if o == nil || o.Duration == nil {
		var ret float64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetDurationOk() (*float64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given float64 and assigns it to the Duration field.
func (o *SyntheticsTestResultStep) SetDuration(v float64) {
	o.Duration = &v
}

// GetElementDescription returns the ElementDescription field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetElementDescription() string {
	if o == nil || o.ElementDescription == nil {
		var ret string
		return ret
	}
	return *o.ElementDescription
}

// GetElementDescriptionOk returns a tuple with the ElementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetElementDescriptionOk() (*string, bool) {
	if o == nil || o.ElementDescription == nil {
		return nil, false
	}
	return o.ElementDescription, true
}

// HasElementDescription returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasElementDescription() bool {
	return o != nil && o.ElementDescription != nil
}

// SetElementDescription gets a reference to the given string and assigns it to the ElementDescription field.
func (o *SyntheticsTestResultStep) SetElementDescription(v string) {
	o.ElementDescription = &v
}

// GetElementUpdates returns the ElementUpdates field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetElementUpdates() SyntheticsTestResultStepElementUpdates {
	if o == nil || o.ElementUpdates == nil {
		var ret SyntheticsTestResultStepElementUpdates
		return ret
	}
	return *o.ElementUpdates
}

// GetElementUpdatesOk returns a tuple with the ElementUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetElementUpdatesOk() (*SyntheticsTestResultStepElementUpdates, bool) {
	if o == nil || o.ElementUpdates == nil {
		return nil, false
	}
	return o.ElementUpdates, true
}

// HasElementUpdates returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasElementUpdates() bool {
	return o != nil && o.ElementUpdates != nil
}

// SetElementUpdates gets a reference to the given SyntheticsTestResultStepElementUpdates and assigns it to the ElementUpdates field.
func (o *SyntheticsTestResultStep) SetElementUpdates(v SyntheticsTestResultStepElementUpdates) {
	o.ElementUpdates = &v
}

// GetExtractedValue returns the ExtractedValue field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetExtractedValue() SyntheticsTestResultVariable {
	if o == nil || o.ExtractedValue == nil {
		var ret SyntheticsTestResultVariable
		return ret
	}
	return *o.ExtractedValue
}

// GetExtractedValueOk returns a tuple with the ExtractedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetExtractedValueOk() (*SyntheticsTestResultVariable, bool) {
	if o == nil || o.ExtractedValue == nil {
		return nil, false
	}
	return o.ExtractedValue, true
}

// HasExtractedValue returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasExtractedValue() bool {
	return o != nil && o.ExtractedValue != nil
}

// SetExtractedValue gets a reference to the given SyntheticsTestResultVariable and assigns it to the ExtractedValue field.
func (o *SyntheticsTestResultStep) SetExtractedValue(v SyntheticsTestResultVariable) {
	o.ExtractedValue = &v
}

// GetFailure returns the Failure field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetFailure() SyntheticsTestResultFailure {
	if o == nil || o.Failure == nil {
		var ret SyntheticsTestResultFailure
		return ret
	}
	return *o.Failure
}

// GetFailureOk returns a tuple with the Failure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetFailureOk() (*SyntheticsTestResultFailure, bool) {
	if o == nil || o.Failure == nil {
		return nil, false
	}
	return o.Failure, true
}

// HasFailure returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasFailure() bool {
	return o != nil && o.Failure != nil
}

// SetFailure gets a reference to the given SyntheticsTestResultFailure and assigns it to the Failure field.
func (o *SyntheticsTestResultStep) SetFailure(v SyntheticsTestResultFailure) {
	o.Failure = &v
}

// GetHttpResults returns the HttpResults field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetHttpResults() []SyntheticsTestResultAssertionResult {
	if o == nil || o.HttpResults == nil {
		var ret []SyntheticsTestResultAssertionResult
		return ret
	}
	return o.HttpResults
}

// GetHttpResultsOk returns a tuple with the HttpResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetHttpResultsOk() (*[]SyntheticsTestResultAssertionResult, bool) {
	if o == nil || o.HttpResults == nil {
		return nil, false
	}
	return &o.HttpResults, true
}

// HasHttpResults returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasHttpResults() bool {
	return o != nil && o.HttpResults != nil
}

// SetHttpResults gets a reference to the given []SyntheticsTestResultAssertionResult and assigns it to the HttpResults field.
func (o *SyntheticsTestResultStep) SetHttpResults(v []SyntheticsTestResultAssertionResult) {
	o.HttpResults = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsTestResultStep) SetId(v string) {
	o.Id = &v
}

// GetIsCritical returns the IsCritical field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetIsCritical() bool {
	if o == nil || o.IsCritical == nil {
		var ret bool
		return ret
	}
	return *o.IsCritical
}

// GetIsCriticalOk returns a tuple with the IsCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetIsCriticalOk() (*bool, bool) {
	if o == nil || o.IsCritical == nil {
		return nil, false
	}
	return o.IsCritical, true
}

// HasIsCritical returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasIsCritical() bool {
	return o != nil && o.IsCritical != nil
}

// SetIsCritical gets a reference to the given bool and assigns it to the IsCritical field.
func (o *SyntheticsTestResultStep) SetIsCritical(v bool) {
	o.IsCritical = &v
}

// GetJavascriptCustomAssertionCode returns the JavascriptCustomAssertionCode field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetJavascriptCustomAssertionCode() bool {
	if o == nil || o.JavascriptCustomAssertionCode == nil {
		var ret bool
		return ret
	}
	return *o.JavascriptCustomAssertionCode
}

// GetJavascriptCustomAssertionCodeOk returns a tuple with the JavascriptCustomAssertionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetJavascriptCustomAssertionCodeOk() (*bool, bool) {
	if o == nil || o.JavascriptCustomAssertionCode == nil {
		return nil, false
	}
	return o.JavascriptCustomAssertionCode, true
}

// HasJavascriptCustomAssertionCode returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasJavascriptCustomAssertionCode() bool {
	return o != nil && o.JavascriptCustomAssertionCode != nil
}

// SetJavascriptCustomAssertionCode gets a reference to the given bool and assigns it to the JavascriptCustomAssertionCode field.
func (o *SyntheticsTestResultStep) SetJavascriptCustomAssertionCode(v bool) {
	o.JavascriptCustomAssertionCode = &v
}

// GetLocateElementDuration returns the LocateElementDuration field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetLocateElementDuration() float64 {
	if o == nil || o.LocateElementDuration == nil {
		var ret float64
		return ret
	}
	return *o.LocateElementDuration
}

// GetLocateElementDurationOk returns a tuple with the LocateElementDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetLocateElementDurationOk() (*float64, bool) {
	if o == nil || o.LocateElementDuration == nil {
		return nil, false
	}
	return o.LocateElementDuration, true
}

// HasLocateElementDuration returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasLocateElementDuration() bool {
	return o != nil && o.LocateElementDuration != nil
}

// SetLocateElementDuration gets a reference to the given float64 and assigns it to the LocateElementDuration field.
func (o *SyntheticsTestResultStep) SetLocateElementDuration(v float64) {
	o.LocateElementDuration = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestResultStep) SetName(v string) {
	o.Name = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetRequest() SyntheticsTestResultRequestInfo {
	if o == nil || o.Request == nil {
		var ret SyntheticsTestResultRequestInfo
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetRequestOk() (*SyntheticsTestResultRequestInfo, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasRequest() bool {
	return o != nil && o.Request != nil
}

// SetRequest gets a reference to the given SyntheticsTestResultRequestInfo and assigns it to the Request field.
func (o *SyntheticsTestResultStep) SetRequest(v SyntheticsTestResultRequestInfo) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetResponse() SyntheticsTestResultResponseInfo {
	if o == nil || o.Response == nil {
		var ret SyntheticsTestResultResponseInfo
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetResponseOk() (*SyntheticsTestResultResponseInfo, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasResponse() bool {
	return o != nil && o.Response != nil
}

// SetResponse gets a reference to the given SyntheticsTestResultResponseInfo and assigns it to the Response field.
func (o *SyntheticsTestResultStep) SetResponse(v SyntheticsTestResultResponseInfo) {
	o.Response = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetRetries() []SyntheticsTestResultStep {
	if o == nil || o.Retries == nil {
		var ret []SyntheticsTestResultStep
		return ret
	}
	return o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetRetriesOk() (*[]SyntheticsTestResultStep, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return &o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasRetries() bool {
	return o != nil && o.Retries != nil
}

// SetRetries gets a reference to the given []SyntheticsTestResultStep and assigns it to the Retries field.
func (o *SyntheticsTestResultStep) SetRetries(v []SyntheticsTestResultStep) {
	o.Retries = v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetRetryCount() int64 {
	if o == nil || o.RetryCount == nil {
		var ret int64
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetRetryCountOk() (*int64, bool) {
	if o == nil || o.RetryCount == nil {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasRetryCount() bool {
	return o != nil && o.RetryCount != nil
}

// SetRetryCount gets a reference to the given int64 and assigns it to the RetryCount field.
func (o *SyntheticsTestResultStep) SetRetryCount(v int64) {
	o.RetryCount = &v
}

// GetRumContext returns the RumContext field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetRumContext() SyntheticsTestResultRumContext {
	if o == nil || o.RumContext == nil {
		var ret SyntheticsTestResultRumContext
		return ret
	}
	return *o.RumContext
}

// GetRumContextOk returns a tuple with the RumContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetRumContextOk() (*SyntheticsTestResultRumContext, bool) {
	if o == nil || o.RumContext == nil {
		return nil, false
	}
	return o.RumContext, true
}

// HasRumContext returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasRumContext() bool {
	return o != nil && o.RumContext != nil
}

// SetRumContext gets a reference to the given SyntheticsTestResultRumContext and assigns it to the RumContext field.
func (o *SyntheticsTestResultStep) SetRumContext(v SyntheticsTestResultRumContext) {
	o.RumContext = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetStartedAt() int64 {
	if o == nil || o.StartedAt == nil {
		var ret int64
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetStartedAtOk() (*int64, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given int64 and assigns it to the StartedAt field.
func (o *SyntheticsTestResultStep) SetStartedAt(v int64) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyntheticsTestResultStep) SetStatus(v string) {
	o.Status = &v
}

// GetSubStep returns the SubStep field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetSubStep() SyntheticsTestResultSubStep {
	if o == nil || o.SubStep == nil {
		var ret SyntheticsTestResultSubStep
		return ret
	}
	return *o.SubStep
}

// GetSubStepOk returns a tuple with the SubStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetSubStepOk() (*SyntheticsTestResultSubStep, bool) {
	if o == nil || o.SubStep == nil {
		return nil, false
	}
	return o.SubStep, true
}

// HasSubStep returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasSubStep() bool {
	return o != nil && o.SubStep != nil
}

// SetSubStep gets a reference to the given SyntheticsTestResultSubStep and assigns it to the SubStep field.
func (o *SyntheticsTestResultStep) SetSubStep(v SyntheticsTestResultSubStep) {
	o.SubStep = &v
}

// GetSubTest returns the SubTest field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetSubTest() SyntheticsTestResultSubTest {
	if o == nil || o.SubTest == nil {
		var ret SyntheticsTestResultSubTest
		return ret
	}
	return *o.SubTest
}

// GetSubTestOk returns a tuple with the SubTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetSubTestOk() (*SyntheticsTestResultSubTest, bool) {
	if o == nil || o.SubTest == nil {
		return nil, false
	}
	return o.SubTest, true
}

// HasSubTest returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasSubTest() bool {
	return o != nil && o.SubTest != nil
}

// SetSubTest gets a reference to the given SyntheticsTestResultSubTest and assigns it to the SubTest field.
func (o *SyntheticsTestResultStep) SetSubTest(v SyntheticsTestResultSubTest) {
	o.SubTest = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasSubtype() bool {
	return o != nil && o.Subtype != nil
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *SyntheticsTestResultStep) SetSubtype(v string) {
	o.Subtype = &v
}

// GetTabs returns the Tabs field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetTabs() []SyntheticsTestResultTab {
	if o == nil || o.Tabs == nil {
		var ret []SyntheticsTestResultTab
		return ret
	}
	return o.Tabs
}

// GetTabsOk returns a tuple with the Tabs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetTabsOk() (*[]SyntheticsTestResultTab, bool) {
	if o == nil || o.Tabs == nil {
		return nil, false
	}
	return &o.Tabs, true
}

// HasTabs returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasTabs() bool {
	return o != nil && o.Tabs != nil
}

// SetTabs gets a reference to the given []SyntheticsTestResultTab and assigns it to the Tabs field.
func (o *SyntheticsTestResultStep) SetTabs(v []SyntheticsTestResultTab) {
	o.Tabs = v
}

// GetTimings returns the Timings field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetTimings() map[string]interface{} {
	if o == nil || o.Timings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Timings
}

// GetTimingsOk returns a tuple with the Timings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetTimingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Timings == nil {
		return nil, false
	}
	return &o.Timings, true
}

// HasTimings returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasTimings() bool {
	return o != nil && o.Timings != nil
}

// SetTimings gets a reference to the given map[string]interface{} and assigns it to the Timings field.
func (o *SyntheticsTestResultStep) SetTimings(v map[string]interface{}) {
	o.Timings = v
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetTunnel() bool {
	if o == nil || o.Tunnel == nil {
		var ret bool
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetTunnelOk() (*bool, bool) {
	if o == nil || o.Tunnel == nil {
		return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasTunnel() bool {
	return o != nil && o.Tunnel != nil
}

// SetTunnel gets a reference to the given bool and assigns it to the Tunnel field.
func (o *SyntheticsTestResultStep) SetTunnel(v bool) {
	o.Tunnel = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SyntheticsTestResultStep) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SyntheticsTestResultStep) SetUrl(v string) {
	o.Url = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetValue() interface{} {
	if o == nil || o.Value == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *SyntheticsTestResultStep) SetValue(v interface{}) {
	o.Value = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetVariables() SyntheticsTestResultVariables {
	if o == nil || o.Variables == nil {
		var ret SyntheticsTestResultVariables
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetVariablesOk() (*SyntheticsTestResultVariables, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasVariables() bool {
	return o != nil && o.Variables != nil
}

// SetVariables gets a reference to the given SyntheticsTestResultVariables and assigns it to the Variables field.
func (o *SyntheticsTestResultStep) SetVariables(v SyntheticsTestResultVariables) {
	o.Variables = &v
}

// GetVitalsMetrics returns the VitalsMetrics field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetVitalsMetrics() []SyntheticsTestResultVitalsMetrics {
	if o == nil || o.VitalsMetrics == nil {
		var ret []SyntheticsTestResultVitalsMetrics
		return ret
	}
	return o.VitalsMetrics
}

// GetVitalsMetricsOk returns a tuple with the VitalsMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetVitalsMetricsOk() (*[]SyntheticsTestResultVitalsMetrics, bool) {
	if o == nil || o.VitalsMetrics == nil {
		return nil, false
	}
	return &o.VitalsMetrics, true
}

// HasVitalsMetrics returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasVitalsMetrics() bool {
	return o != nil && o.VitalsMetrics != nil
}

// SetVitalsMetrics gets a reference to the given []SyntheticsTestResultVitalsMetrics and assigns it to the VitalsMetrics field.
func (o *SyntheticsTestResultStep) SetVitalsMetrics(v []SyntheticsTestResultVitalsMetrics) {
	o.VitalsMetrics = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SyntheticsTestResultStep) GetWarnings() []SyntheticsTestResultWarning {
	if o == nil || o.Warnings == nil {
		var ret []SyntheticsTestResultWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStep) GetWarningsOk() (*[]SyntheticsTestResultWarning, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SyntheticsTestResultStep) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []SyntheticsTestResultWarning and assigns it to the Warnings field.
func (o *SyntheticsTestResultStep) SetWarnings(v []SyntheticsTestResultWarning) {
	o.Warnings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowFailure != nil {
		toSerialize["allow_failure"] = o.AllowFailure
	}
	if o.ApiTest != nil {
		toSerialize["api_test"] = o.ApiTest
	}
	if o.AssertionResult != nil {
		toSerialize["assertion_result"] = o.AssertionResult
	}
	if o.Assertions != nil {
		toSerialize["assertions"] = o.Assertions
	}
	if o.BlockedRequestsUrls != nil {
		toSerialize["blocked_requests_urls"] = o.BlockedRequestsUrls
	}
	if o.Bounds != nil {
		toSerialize["bounds"] = o.Bounds
	}
	if o.BrowserErrors != nil {
		toSerialize["browser_errors"] = o.BrowserErrors
	}
	if o.BucketKeys != nil {
		toSerialize["bucket_keys"] = o.BucketKeys
	}
	if o.CdnResources != nil {
		toSerialize["cdn_resources"] = o.CdnResources
	}
	if o.ClickType != nil {
		toSerialize["click_type"] = o.ClickType
	}
	if o.CompressedJsonDescriptor != nil {
		toSerialize["compressed_json_descriptor"] = o.CompressedJsonDescriptor
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.ElementDescription != nil {
		toSerialize["element_description"] = o.ElementDescription
	}
	if o.ElementUpdates != nil {
		toSerialize["element_updates"] = o.ElementUpdates
	}
	if o.ExtractedValue != nil {
		toSerialize["extracted_value"] = o.ExtractedValue
	}
	if o.Failure != nil {
		toSerialize["failure"] = o.Failure
	}
	if o.HttpResults != nil {
		toSerialize["http_results"] = o.HttpResults
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsCritical != nil {
		toSerialize["is_critical"] = o.IsCritical
	}
	if o.JavascriptCustomAssertionCode != nil {
		toSerialize["javascript_custom_assertion_code"] = o.JavascriptCustomAssertionCode
	}
	if o.LocateElementDuration != nil {
		toSerialize["locate_element_duration"] = o.LocateElementDuration
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}
	if o.RetryCount != nil {
		toSerialize["retry_count"] = o.RetryCount
	}
	if o.RumContext != nil {
		toSerialize["rum_context"] = o.RumContext
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SubStep != nil {
		toSerialize["sub_step"] = o.SubStep
	}
	if o.SubTest != nil {
		toSerialize["sub_test"] = o.SubTest
	}
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
	}
	if o.Tabs != nil {
		toSerialize["tabs"] = o.Tabs
	}
	if o.Timings != nil {
		toSerialize["timings"] = o.Timings
	}
	if o.Tunnel != nil {
		toSerialize["tunnel"] = o.Tunnel
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	if o.VitalsMetrics != nil {
		toSerialize["vitals_metrics"] = o.VitalsMetrics
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowFailure                  *bool                                    `json:"allow_failure,omitempty"`
		ApiTest                       map[string]interface{}                   `json:"api_test,omitempty"`
		AssertionResult               *SyntheticsTestResultStepAssertionResult `json:"assertion_result,omitempty"`
		Assertions                    []SyntheticsTestResultAssertionResult    `json:"assertions,omitempty"`
		BlockedRequestsUrls           []string                                 `json:"blocked_requests_urls,omitempty"`
		Bounds                        *SyntheticsTestResultBounds              `json:"bounds,omitempty"`
		BrowserErrors                 []SyntheticsTestResultBrowserError       `json:"browser_errors,omitempty"`
		BucketKeys                    *SyntheticsTestResultBucketKeys          `json:"bucket_keys,omitempty"`
		CdnResources                  []SyntheticsTestResultCdnResource        `json:"cdn_resources,omitempty"`
		ClickType                     *string                                  `json:"click_type,omitempty"`
		CompressedJsonDescriptor      *string                                  `json:"compressed_json_descriptor,omitempty"`
		Config                        map[string]interface{}                   `json:"config,omitempty"`
		Description                   *string                                  `json:"description,omitempty"`
		Duration                      *float64                                 `json:"duration,omitempty"`
		ElementDescription            *string                                  `json:"element_description,omitempty"`
		ElementUpdates                *SyntheticsTestResultStepElementUpdates  `json:"element_updates,omitempty"`
		ExtractedValue                *SyntheticsTestResultVariable            `json:"extracted_value,omitempty"`
		Failure                       *SyntheticsTestResultFailure             `json:"failure,omitempty"`
		HttpResults                   []SyntheticsTestResultAssertionResult    `json:"http_results,omitempty"`
		Id                            *string                                  `json:"id,omitempty"`
		IsCritical                    *bool                                    `json:"is_critical,omitempty"`
		JavascriptCustomAssertionCode *bool                                    `json:"javascript_custom_assertion_code,omitempty"`
		LocateElementDuration         *float64                                 `json:"locate_element_duration,omitempty"`
		Name                          *string                                  `json:"name,omitempty"`
		Request                       *SyntheticsTestResultRequestInfo         `json:"request,omitempty"`
		Response                      *SyntheticsTestResultResponseInfo        `json:"response,omitempty"`
		Retries                       []SyntheticsTestResultStep               `json:"retries,omitempty"`
		RetryCount                    *int64                                   `json:"retry_count,omitempty"`
		RumContext                    *SyntheticsTestResultRumContext          `json:"rum_context,omitempty"`
		StartedAt                     *int64                                   `json:"started_at,omitempty"`
		Status                        *string                                  `json:"status,omitempty"`
		SubStep                       *SyntheticsTestResultSubStep             `json:"sub_step,omitempty"`
		SubTest                       *SyntheticsTestResultSubTest             `json:"sub_test,omitempty"`
		Subtype                       *string                                  `json:"subtype,omitempty"`
		Tabs                          []SyntheticsTestResultTab                `json:"tabs,omitempty"`
		Timings                       map[string]interface{}                   `json:"timings,omitempty"`
		Tunnel                        *bool                                    `json:"tunnel,omitempty"`
		Type                          *string                                  `json:"type,omitempty"`
		Url                           *string                                  `json:"url,omitempty"`
		Value                         interface{}                              `json:"value,omitempty"`
		Variables                     *SyntheticsTestResultVariables           `json:"variables,omitempty"`
		VitalsMetrics                 []SyntheticsTestResultVitalsMetrics      `json:"vitals_metrics,omitempty"`
		Warnings                      []SyntheticsTestResultWarning            `json:"warnings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_failure", "api_test", "assertion_result", "assertions", "blocked_requests_urls", "bounds", "browser_errors", "bucket_keys", "cdn_resources", "click_type", "compressed_json_descriptor", "config", "description", "duration", "element_description", "element_updates", "extracted_value", "failure", "http_results", "id", "is_critical", "javascript_custom_assertion_code", "locate_element_duration", "name", "request", "response", "retries", "retry_count", "rum_context", "started_at", "status", "sub_step", "sub_test", "subtype", "tabs", "timings", "tunnel", "type", "url", "value", "variables", "vitals_metrics", "warnings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowFailure = all.AllowFailure
	o.ApiTest = all.ApiTest
	if all.AssertionResult != nil && all.AssertionResult.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AssertionResult = all.AssertionResult
	o.Assertions = all.Assertions
	o.BlockedRequestsUrls = all.BlockedRequestsUrls
	if all.Bounds != nil && all.Bounds.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Bounds = all.Bounds
	o.BrowserErrors = all.BrowserErrors
	if all.BucketKeys != nil && all.BucketKeys.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BucketKeys = all.BucketKeys
	o.CdnResources = all.CdnResources
	o.ClickType = all.ClickType
	o.CompressedJsonDescriptor = all.CompressedJsonDescriptor
	o.Config = all.Config
	o.Description = all.Description
	o.Duration = all.Duration
	o.ElementDescription = all.ElementDescription
	if all.ElementUpdates != nil && all.ElementUpdates.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElementUpdates = all.ElementUpdates
	if all.ExtractedValue != nil && all.ExtractedValue.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ExtractedValue = all.ExtractedValue
	if all.Failure != nil && all.Failure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Failure = all.Failure
	o.HttpResults = all.HttpResults
	o.Id = all.Id
	o.IsCritical = all.IsCritical
	o.JavascriptCustomAssertionCode = all.JavascriptCustomAssertionCode
	o.LocateElementDuration = all.LocateElementDuration
	o.Name = all.Name
	if all.Request != nil && all.Request.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Request = all.Request
	if all.Response != nil && all.Response.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Response = all.Response
	o.Retries = all.Retries
	o.RetryCount = all.RetryCount
	if all.RumContext != nil && all.RumContext.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RumContext = all.RumContext
	o.StartedAt = all.StartedAt
	o.Status = all.Status
	if all.SubStep != nil && all.SubStep.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SubStep = all.SubStep
	if all.SubTest != nil && all.SubTest.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SubTest = all.SubTest
	o.Subtype = all.Subtype
	o.Tabs = all.Tabs
	o.Timings = all.Timings
	o.Tunnel = all.Tunnel
	o.Type = all.Type
	o.Url = all.Url
	o.Value = all.Value
	if all.Variables != nil && all.Variables.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Variables = all.Variables
	o.VitalsMetrics = all.VitalsMetrics
	o.Warnings = all.Warnings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
