// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsStepDetail Object describing a step for a Synthetic test.
type SyntheticsStepDetail struct {
	// Whether or not the step was allowed to fail.
	AllowFailure *bool `json:"allowFailure,omitempty"`
	// Array of errors collected for a browser test.
	BrowserErrors []SyntheticsBrowserError `json:"browserErrors,omitempty"`
	// Type of assertion to apply in an API test.
	CheckType *SyntheticsCheckType `json:"checkType,omitempty"`
	// Description of the test.
	Description *string `json:"description,omitempty"`
	// Total duration in millisecond of the test.
	Duration *float64 `json:"duration,omitempty"`
	// Error returned by the test.
	Error *string `json:"error,omitempty"`
	// The browser test failure details.
	Failure *SyntheticsBrowserTestResultFailure `json:"failure,omitempty"`
	// Navigate between different tabs for your browser test.
	PlayingTab *SyntheticsPlayingTab `json:"playingTab,omitempty"`
	// Whether or not screenshots where collected by the test.
	ScreenshotBucketKey *bool `json:"screenshotBucketKey,omitempty"`
	// Whether or not to skip this step.
	Skipped *bool `json:"skipped,omitempty"`
	// Whether or not snapshots where collected by the test.
	SnapshotBucketKey *bool `json:"snapshotBucketKey,omitempty"`
	// The step ID.
	StepId *int64 `json:"stepId,omitempty"`
	// If this step includes a sub-test.
	// [Subtests documentation](https://docs.datadoghq.com/synthetics/browser_tests/advanced_options/#subtests).
	SubTestStepDetails []SyntheticsStepDetail `json:"subTestStepDetails,omitempty"`
	// Time before starting the step.
	TimeToInteractive *float64 `json:"timeToInteractive,omitempty"`
	// Step type used in your Synthetic test.
	Type *SyntheticsStepType `json:"type,omitempty"`
	// URL to perform the step against.
	Url *string `json:"url,omitempty"`
	// Value for the step.
	Value interface{} `json:"value,omitempty"`
	// Array of Core Web Vitals metrics for the step.
	VitalsMetrics []SyntheticsCoreWebVitals `json:"vitalsMetrics,omitempty"`
	// Warning collected that didn't failed the step.
	Warnings []SyntheticsStepDetailWarning `json:"warnings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsStepDetail instantiates a new SyntheticsStepDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsStepDetail() *SyntheticsStepDetail {
	this := SyntheticsStepDetail{}
	return &this
}

// NewSyntheticsStepDetailWithDefaults instantiates a new SyntheticsStepDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsStepDetailWithDefaults() *SyntheticsStepDetail {
	this := SyntheticsStepDetail{}
	return &this
}

// GetAllowFailure returns the AllowFailure field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetAllowFailure() bool {
	if o == nil || o.AllowFailure == nil {
		var ret bool
		return ret
	}
	return *o.AllowFailure
}

// GetAllowFailureOk returns a tuple with the AllowFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetAllowFailureOk() (*bool, bool) {
	if o == nil || o.AllowFailure == nil {
		return nil, false
	}
	return o.AllowFailure, true
}

// HasAllowFailure returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasAllowFailure() bool {
	return o != nil && o.AllowFailure != nil
}

// SetAllowFailure gets a reference to the given bool and assigns it to the AllowFailure field.
func (o *SyntheticsStepDetail) SetAllowFailure(v bool) {
	o.AllowFailure = &v
}

// GetBrowserErrors returns the BrowserErrors field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetBrowserErrors() []SyntheticsBrowserError {
	if o == nil || o.BrowserErrors == nil {
		var ret []SyntheticsBrowserError
		return ret
	}
	return o.BrowserErrors
}

// GetBrowserErrorsOk returns a tuple with the BrowserErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetBrowserErrorsOk() (*[]SyntheticsBrowserError, bool) {
	if o == nil || o.BrowserErrors == nil {
		return nil, false
	}
	return &o.BrowserErrors, true
}

// HasBrowserErrors returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasBrowserErrors() bool {
	return o != nil && o.BrowserErrors != nil
}

// SetBrowserErrors gets a reference to the given []SyntheticsBrowserError and assigns it to the BrowserErrors field.
func (o *SyntheticsStepDetail) SetBrowserErrors(v []SyntheticsBrowserError) {
	o.BrowserErrors = v
}

// GetCheckType returns the CheckType field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetCheckType() SyntheticsCheckType {
	if o == nil || o.CheckType == nil {
		var ret SyntheticsCheckType
		return ret
	}
	return *o.CheckType
}

// GetCheckTypeOk returns a tuple with the CheckType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetCheckTypeOk() (*SyntheticsCheckType, bool) {
	if o == nil || o.CheckType == nil {
		return nil, false
	}
	return o.CheckType, true
}

// HasCheckType returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasCheckType() bool {
	return o != nil && o.CheckType != nil
}

// SetCheckType gets a reference to the given SyntheticsCheckType and assigns it to the CheckType field.
func (o *SyntheticsStepDetail) SetCheckType(v SyntheticsCheckType) {
	o.CheckType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SyntheticsStepDetail) SetDescription(v string) {
	o.Description = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetDuration() float64 {
	if o == nil || o.Duration == nil {
		var ret float64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetDurationOk() (*float64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given float64 and assigns it to the Duration field.
func (o *SyntheticsStepDetail) SetDuration(v float64) {
	o.Duration = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *SyntheticsStepDetail) SetError(v string) {
	o.Error = &v
}

// GetFailure returns the Failure field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetFailure() SyntheticsBrowserTestResultFailure {
	if o == nil || o.Failure == nil {
		var ret SyntheticsBrowserTestResultFailure
		return ret
	}
	return *o.Failure
}

// GetFailureOk returns a tuple with the Failure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetFailureOk() (*SyntheticsBrowserTestResultFailure, bool) {
	if o == nil || o.Failure == nil {
		return nil, false
	}
	return o.Failure, true
}

// HasFailure returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasFailure() bool {
	return o != nil && o.Failure != nil
}

// SetFailure gets a reference to the given SyntheticsBrowserTestResultFailure and assigns it to the Failure field.
func (o *SyntheticsStepDetail) SetFailure(v SyntheticsBrowserTestResultFailure) {
	o.Failure = &v
}

// GetPlayingTab returns the PlayingTab field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetPlayingTab() SyntheticsPlayingTab {
	if o == nil || o.PlayingTab == nil {
		var ret SyntheticsPlayingTab
		return ret
	}
	return *o.PlayingTab
}

// GetPlayingTabOk returns a tuple with the PlayingTab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetPlayingTabOk() (*SyntheticsPlayingTab, bool) {
	if o == nil || o.PlayingTab == nil {
		return nil, false
	}
	return o.PlayingTab, true
}

// HasPlayingTab returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasPlayingTab() bool {
	return o != nil && o.PlayingTab != nil
}

// SetPlayingTab gets a reference to the given SyntheticsPlayingTab and assigns it to the PlayingTab field.
func (o *SyntheticsStepDetail) SetPlayingTab(v SyntheticsPlayingTab) {
	o.PlayingTab = &v
}

// GetScreenshotBucketKey returns the ScreenshotBucketKey field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetScreenshotBucketKey() bool {
	if o == nil || o.ScreenshotBucketKey == nil {
		var ret bool
		return ret
	}
	return *o.ScreenshotBucketKey
}

// GetScreenshotBucketKeyOk returns a tuple with the ScreenshotBucketKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetScreenshotBucketKeyOk() (*bool, bool) {
	if o == nil || o.ScreenshotBucketKey == nil {
		return nil, false
	}
	return o.ScreenshotBucketKey, true
}

// HasScreenshotBucketKey returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasScreenshotBucketKey() bool {
	return o != nil && o.ScreenshotBucketKey != nil
}

// SetScreenshotBucketKey gets a reference to the given bool and assigns it to the ScreenshotBucketKey field.
func (o *SyntheticsStepDetail) SetScreenshotBucketKey(v bool) {
	o.ScreenshotBucketKey = &v
}

// GetSkipped returns the Skipped field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetSkipped() bool {
	if o == nil || o.Skipped == nil {
		var ret bool
		return ret
	}
	return *o.Skipped
}

// GetSkippedOk returns a tuple with the Skipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetSkippedOk() (*bool, bool) {
	if o == nil || o.Skipped == nil {
		return nil, false
	}
	return o.Skipped, true
}

// HasSkipped returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasSkipped() bool {
	return o != nil && o.Skipped != nil
}

// SetSkipped gets a reference to the given bool and assigns it to the Skipped field.
func (o *SyntheticsStepDetail) SetSkipped(v bool) {
	o.Skipped = &v
}

// GetSnapshotBucketKey returns the SnapshotBucketKey field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetSnapshotBucketKey() bool {
	if o == nil || o.SnapshotBucketKey == nil {
		var ret bool
		return ret
	}
	return *o.SnapshotBucketKey
}

// GetSnapshotBucketKeyOk returns a tuple with the SnapshotBucketKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetSnapshotBucketKeyOk() (*bool, bool) {
	if o == nil || o.SnapshotBucketKey == nil {
		return nil, false
	}
	return o.SnapshotBucketKey, true
}

// HasSnapshotBucketKey returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasSnapshotBucketKey() bool {
	return o != nil && o.SnapshotBucketKey != nil
}

// SetSnapshotBucketKey gets a reference to the given bool and assigns it to the SnapshotBucketKey field.
func (o *SyntheticsStepDetail) SetSnapshotBucketKey(v bool) {
	o.SnapshotBucketKey = &v
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetStepId() int64 {
	if o == nil || o.StepId == nil {
		var ret int64
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetStepIdOk() (*int64, bool) {
	if o == nil || o.StepId == nil {
		return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasStepId() bool {
	return o != nil && o.StepId != nil
}

// SetStepId gets a reference to the given int64 and assigns it to the StepId field.
func (o *SyntheticsStepDetail) SetStepId(v int64) {
	o.StepId = &v
}

// GetSubTestStepDetails returns the SubTestStepDetails field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetSubTestStepDetails() []SyntheticsStepDetail {
	if o == nil || o.SubTestStepDetails == nil {
		var ret []SyntheticsStepDetail
		return ret
	}
	return o.SubTestStepDetails
}

// GetSubTestStepDetailsOk returns a tuple with the SubTestStepDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetSubTestStepDetailsOk() (*[]SyntheticsStepDetail, bool) {
	if o == nil || o.SubTestStepDetails == nil {
		return nil, false
	}
	return &o.SubTestStepDetails, true
}

// HasSubTestStepDetails returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasSubTestStepDetails() bool {
	return o != nil && o.SubTestStepDetails != nil
}

// SetSubTestStepDetails gets a reference to the given []SyntheticsStepDetail and assigns it to the SubTestStepDetails field.
func (o *SyntheticsStepDetail) SetSubTestStepDetails(v []SyntheticsStepDetail) {
	o.SubTestStepDetails = v
}

// GetTimeToInteractive returns the TimeToInteractive field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetTimeToInteractive() float64 {
	if o == nil || o.TimeToInteractive == nil {
		var ret float64
		return ret
	}
	return *o.TimeToInteractive
}

// GetTimeToInteractiveOk returns a tuple with the TimeToInteractive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetTimeToInteractiveOk() (*float64, bool) {
	if o == nil || o.TimeToInteractive == nil {
		return nil, false
	}
	return o.TimeToInteractive, true
}

// HasTimeToInteractive returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasTimeToInteractive() bool {
	return o != nil && o.TimeToInteractive != nil
}

// SetTimeToInteractive gets a reference to the given float64 and assigns it to the TimeToInteractive field.
func (o *SyntheticsStepDetail) SetTimeToInteractive(v float64) {
	o.TimeToInteractive = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetType() SyntheticsStepType {
	if o == nil || o.Type == nil {
		var ret SyntheticsStepType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetTypeOk() (*SyntheticsStepType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SyntheticsStepType and assigns it to the Type field.
func (o *SyntheticsStepDetail) SetType(v SyntheticsStepType) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SyntheticsStepDetail) SetUrl(v string) {
	o.Url = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetValue() interface{} {
	if o == nil || o.Value == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *SyntheticsStepDetail) SetValue(v interface{}) {
	o.Value = v
}

// GetVitalsMetrics returns the VitalsMetrics field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetVitalsMetrics() []SyntheticsCoreWebVitals {
	if o == nil || o.VitalsMetrics == nil {
		var ret []SyntheticsCoreWebVitals
		return ret
	}
	return o.VitalsMetrics
}

// GetVitalsMetricsOk returns a tuple with the VitalsMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetVitalsMetricsOk() (*[]SyntheticsCoreWebVitals, bool) {
	if o == nil || o.VitalsMetrics == nil {
		return nil, false
	}
	return &o.VitalsMetrics, true
}

// HasVitalsMetrics returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasVitalsMetrics() bool {
	return o != nil && o.VitalsMetrics != nil
}

// SetVitalsMetrics gets a reference to the given []SyntheticsCoreWebVitals and assigns it to the VitalsMetrics field.
func (o *SyntheticsStepDetail) SetVitalsMetrics(v []SyntheticsCoreWebVitals) {
	o.VitalsMetrics = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SyntheticsStepDetail) GetWarnings() []SyntheticsStepDetailWarning {
	if o == nil || o.Warnings == nil {
		var ret []SyntheticsStepDetailWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetail) GetWarningsOk() (*[]SyntheticsStepDetailWarning, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SyntheticsStepDetail) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []SyntheticsStepDetailWarning and assigns it to the Warnings field.
func (o *SyntheticsStepDetail) SetWarnings(v []SyntheticsStepDetailWarning) {
	o.Warnings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsStepDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowFailure != nil {
		toSerialize["allowFailure"] = o.AllowFailure
	}
	if o.BrowserErrors != nil {
		toSerialize["browserErrors"] = o.BrowserErrors
	}
	if o.CheckType != nil {
		toSerialize["checkType"] = o.CheckType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Failure != nil {
		toSerialize["failure"] = o.Failure
	}
	if o.PlayingTab != nil {
		toSerialize["playingTab"] = o.PlayingTab
	}
	if o.ScreenshotBucketKey != nil {
		toSerialize["screenshotBucketKey"] = o.ScreenshotBucketKey
	}
	if o.Skipped != nil {
		toSerialize["skipped"] = o.Skipped
	}
	if o.SnapshotBucketKey != nil {
		toSerialize["snapshotBucketKey"] = o.SnapshotBucketKey
	}
	if o.StepId != nil {
		toSerialize["stepId"] = o.StepId
	}
	if o.SubTestStepDetails != nil {
		toSerialize["subTestStepDetails"] = o.SubTestStepDetails
	}
	if o.TimeToInteractive != nil {
		toSerialize["timeToInteractive"] = o.TimeToInteractive
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
	if o.VitalsMetrics != nil {
		toSerialize["vitalsMetrics"] = o.VitalsMetrics
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
func (o *SyntheticsStepDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowFailure        *bool                               `json:"allowFailure,omitempty"`
		BrowserErrors       []SyntheticsBrowserError            `json:"browserErrors,omitempty"`
		CheckType           *SyntheticsCheckType                `json:"checkType,omitempty"`
		Description         *string                             `json:"description,omitempty"`
		Duration            *float64                            `json:"duration,omitempty"`
		Error               *string                             `json:"error,omitempty"`
		Failure             *SyntheticsBrowserTestResultFailure `json:"failure,omitempty"`
		PlayingTab          *SyntheticsPlayingTab               `json:"playingTab,omitempty"`
		ScreenshotBucketKey *bool                               `json:"screenshotBucketKey,omitempty"`
		Skipped             *bool                               `json:"skipped,omitempty"`
		SnapshotBucketKey   *bool                               `json:"snapshotBucketKey,omitempty"`
		StepId              *int64                              `json:"stepId,omitempty"`
		SubTestStepDetails  []SyntheticsStepDetail              `json:"subTestStepDetails,omitempty"`
		TimeToInteractive   *float64                            `json:"timeToInteractive,omitempty"`
		Type                *SyntheticsStepType                 `json:"type,omitempty"`
		Url                 *string                             `json:"url,omitempty"`
		Value               interface{}                         `json:"value,omitempty"`
		VitalsMetrics       []SyntheticsCoreWebVitals           `json:"vitalsMetrics,omitempty"`
		Warnings            []SyntheticsStepDetailWarning       `json:"warnings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowFailure", "browserErrors", "checkType", "description", "duration", "error", "failure", "playingTab", "screenshotBucketKey", "skipped", "snapshotBucketKey", "stepId", "subTestStepDetails", "timeToInteractive", "type", "url", "value", "vitalsMetrics", "warnings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowFailure = all.AllowFailure
	o.BrowserErrors = all.BrowserErrors
	if all.CheckType != nil && !all.CheckType.IsValid() {
		hasInvalidField = true
	} else {
		o.CheckType = all.CheckType
	}
	o.Description = all.Description
	o.Duration = all.Duration
	o.Error = all.Error
	if all.Failure != nil && all.Failure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Failure = all.Failure
	if all.PlayingTab != nil && !all.PlayingTab.IsValid() {
		hasInvalidField = true
	} else {
		o.PlayingTab = all.PlayingTab
	}
	o.ScreenshotBucketKey = all.ScreenshotBucketKey
	o.Skipped = all.Skipped
	o.SnapshotBucketKey = all.SnapshotBucketKey
	o.StepId = all.StepId
	o.SubTestStepDetails = all.SubTestStepDetails
	o.TimeToInteractive = all.TimeToInteractive
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Url = all.Url
	o.Value = all.Value
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
