// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultBucketKeys Storage bucket keys for artifacts produced during a step or test.
type SyntheticsTestResultBucketKeys struct {
	// Key for the screenshot captured after the step (goal-based tests).
	AfterStepScreenshot *string `json:"after_step_screenshot,omitempty"`
	// Key for the screenshot captured after the turn (goal-based tests).
	AfterTurnScreenshot *string `json:"after_turn_screenshot,omitempty"`
	// Key for miscellaneous artifacts.
	Artifacts *string `json:"artifacts,omitempty"`
	// Key for the screenshot captured before the step (goal-based tests).
	BeforeStepScreenshot *string `json:"before_step_screenshot,omitempty"`
	// Key for the screenshot captured before the turn (goal-based tests).
	BeforeTurnScreenshot *string `json:"before_turn_screenshot,omitempty"`
	// Key for a captured crash report.
	CrashReport *string `json:"crash_report,omitempty"`
	// Key for captured device logs.
	DeviceLogs *string `json:"device_logs,omitempty"`
	// Keys for email message payloads captured by the step.
	EmailMessages []string `json:"email_messages,omitempty"`
	// Key for the captured screenshot.
	Screenshot *string `json:"screenshot,omitempty"`
	// Key for the captured DOM snapshot.
	Snapshot *string `json:"snapshot,omitempty"`
	// Key for the page source or element source.
	Source *string `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultBucketKeys instantiates a new SyntheticsTestResultBucketKeys object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultBucketKeys() *SyntheticsTestResultBucketKeys {
	this := SyntheticsTestResultBucketKeys{}
	return &this
}

// NewSyntheticsTestResultBucketKeysWithDefaults instantiates a new SyntheticsTestResultBucketKeys object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultBucketKeysWithDefaults() *SyntheticsTestResultBucketKeys {
	this := SyntheticsTestResultBucketKeys{}
	return &this
}

// GetAfterStepScreenshot returns the AfterStepScreenshot field value if set, zero value otherwise.
func (o *SyntheticsTestResultBucketKeys) GetAfterStepScreenshot() string {
	if o == nil || o.AfterStepScreenshot == nil {
		var ret string
		return ret
	}
	return *o.AfterStepScreenshot
}

// GetAfterStepScreenshotOk returns a tuple with the AfterStepScreenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBucketKeys) GetAfterStepScreenshotOk() (*string, bool) {
	if o == nil || o.AfterStepScreenshot == nil {
		return nil, false
	}
	return o.AfterStepScreenshot, true
}

// HasAfterStepScreenshot returns a boolean if a field has been set.
func (o *SyntheticsTestResultBucketKeys) HasAfterStepScreenshot() bool {
	return o != nil && o.AfterStepScreenshot != nil
}

// SetAfterStepScreenshot gets a reference to the given string and assigns it to the AfterStepScreenshot field.
func (o *SyntheticsTestResultBucketKeys) SetAfterStepScreenshot(v string) {
	o.AfterStepScreenshot = &v
}

// GetAfterTurnScreenshot returns the AfterTurnScreenshot field value if set, zero value otherwise.
func (o *SyntheticsTestResultBucketKeys) GetAfterTurnScreenshot() string {
	if o == nil || o.AfterTurnScreenshot == nil {
		var ret string
		return ret
	}
	return *o.AfterTurnScreenshot
}

// GetAfterTurnScreenshotOk returns a tuple with the AfterTurnScreenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBucketKeys) GetAfterTurnScreenshotOk() (*string, bool) {
	if o == nil || o.AfterTurnScreenshot == nil {
		return nil, false
	}
	return o.AfterTurnScreenshot, true
}

// HasAfterTurnScreenshot returns a boolean if a field has been set.
func (o *SyntheticsTestResultBucketKeys) HasAfterTurnScreenshot() bool {
	return o != nil && o.AfterTurnScreenshot != nil
}

// SetAfterTurnScreenshot gets a reference to the given string and assigns it to the AfterTurnScreenshot field.
func (o *SyntheticsTestResultBucketKeys) SetAfterTurnScreenshot(v string) {
	o.AfterTurnScreenshot = &v
}

// GetArtifacts returns the Artifacts field value if set, zero value otherwise.
func (o *SyntheticsTestResultBucketKeys) GetArtifacts() string {
	if o == nil || o.Artifacts == nil {
		var ret string
		return ret
	}
	return *o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBucketKeys) GetArtifactsOk() (*string, bool) {
	if o == nil || o.Artifacts == nil {
		return nil, false
	}
	return o.Artifacts, true
}

// HasArtifacts returns a boolean if a field has been set.
func (o *SyntheticsTestResultBucketKeys) HasArtifacts() bool {
	return o != nil && o.Artifacts != nil
}

// SetArtifacts gets a reference to the given string and assigns it to the Artifacts field.
func (o *SyntheticsTestResultBucketKeys) SetArtifacts(v string) {
	o.Artifacts = &v
}

// GetBeforeStepScreenshot returns the BeforeStepScreenshot field value if set, zero value otherwise.
func (o *SyntheticsTestResultBucketKeys) GetBeforeStepScreenshot() string {
	if o == nil || o.BeforeStepScreenshot == nil {
		var ret string
		return ret
	}
	return *o.BeforeStepScreenshot
}

// GetBeforeStepScreenshotOk returns a tuple with the BeforeStepScreenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBucketKeys) GetBeforeStepScreenshotOk() (*string, bool) {
	if o == nil || o.BeforeStepScreenshot == nil {
		return nil, false
	}
	return o.BeforeStepScreenshot, true
}

// HasBeforeStepScreenshot returns a boolean if a field has been set.
func (o *SyntheticsTestResultBucketKeys) HasBeforeStepScreenshot() bool {
	return o != nil && o.BeforeStepScreenshot != nil
}

// SetBeforeStepScreenshot gets a reference to the given string and assigns it to the BeforeStepScreenshot field.
func (o *SyntheticsTestResultBucketKeys) SetBeforeStepScreenshot(v string) {
	o.BeforeStepScreenshot = &v
}

// GetBeforeTurnScreenshot returns the BeforeTurnScreenshot field value if set, zero value otherwise.
func (o *SyntheticsTestResultBucketKeys) GetBeforeTurnScreenshot() string {
	if o == nil || o.BeforeTurnScreenshot == nil {
		var ret string
		return ret
	}
	return *o.BeforeTurnScreenshot
}

// GetBeforeTurnScreenshotOk returns a tuple with the BeforeTurnScreenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBucketKeys) GetBeforeTurnScreenshotOk() (*string, bool) {
	if o == nil || o.BeforeTurnScreenshot == nil {
		return nil, false
	}
	return o.BeforeTurnScreenshot, true
}

// HasBeforeTurnScreenshot returns a boolean if a field has been set.
func (o *SyntheticsTestResultBucketKeys) HasBeforeTurnScreenshot() bool {
	return o != nil && o.BeforeTurnScreenshot != nil
}

// SetBeforeTurnScreenshot gets a reference to the given string and assigns it to the BeforeTurnScreenshot field.
func (o *SyntheticsTestResultBucketKeys) SetBeforeTurnScreenshot(v string) {
	o.BeforeTurnScreenshot = &v
}

// GetCrashReport returns the CrashReport field value if set, zero value otherwise.
func (o *SyntheticsTestResultBucketKeys) GetCrashReport() string {
	if o == nil || o.CrashReport == nil {
		var ret string
		return ret
	}
	return *o.CrashReport
}

// GetCrashReportOk returns a tuple with the CrashReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBucketKeys) GetCrashReportOk() (*string, bool) {
	if o == nil || o.CrashReport == nil {
		return nil, false
	}
	return o.CrashReport, true
}

// HasCrashReport returns a boolean if a field has been set.
func (o *SyntheticsTestResultBucketKeys) HasCrashReport() bool {
	return o != nil && o.CrashReport != nil
}

// SetCrashReport gets a reference to the given string and assigns it to the CrashReport field.
func (o *SyntheticsTestResultBucketKeys) SetCrashReport(v string) {
	o.CrashReport = &v
}

// GetDeviceLogs returns the DeviceLogs field value if set, zero value otherwise.
func (o *SyntheticsTestResultBucketKeys) GetDeviceLogs() string {
	if o == nil || o.DeviceLogs == nil {
		var ret string
		return ret
	}
	return *o.DeviceLogs
}

// GetDeviceLogsOk returns a tuple with the DeviceLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBucketKeys) GetDeviceLogsOk() (*string, bool) {
	if o == nil || o.DeviceLogs == nil {
		return nil, false
	}
	return o.DeviceLogs, true
}

// HasDeviceLogs returns a boolean if a field has been set.
func (o *SyntheticsTestResultBucketKeys) HasDeviceLogs() bool {
	return o != nil && o.DeviceLogs != nil
}

// SetDeviceLogs gets a reference to the given string and assigns it to the DeviceLogs field.
func (o *SyntheticsTestResultBucketKeys) SetDeviceLogs(v string) {
	o.DeviceLogs = &v
}

// GetEmailMessages returns the EmailMessages field value if set, zero value otherwise.
func (o *SyntheticsTestResultBucketKeys) GetEmailMessages() []string {
	if o == nil || o.EmailMessages == nil {
		var ret []string
		return ret
	}
	return o.EmailMessages
}

// GetEmailMessagesOk returns a tuple with the EmailMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBucketKeys) GetEmailMessagesOk() (*[]string, bool) {
	if o == nil || o.EmailMessages == nil {
		return nil, false
	}
	return &o.EmailMessages, true
}

// HasEmailMessages returns a boolean if a field has been set.
func (o *SyntheticsTestResultBucketKeys) HasEmailMessages() bool {
	return o != nil && o.EmailMessages != nil
}

// SetEmailMessages gets a reference to the given []string and assigns it to the EmailMessages field.
func (o *SyntheticsTestResultBucketKeys) SetEmailMessages(v []string) {
	o.EmailMessages = v
}

// GetScreenshot returns the Screenshot field value if set, zero value otherwise.
func (o *SyntheticsTestResultBucketKeys) GetScreenshot() string {
	if o == nil || o.Screenshot == nil {
		var ret string
		return ret
	}
	return *o.Screenshot
}

// GetScreenshotOk returns a tuple with the Screenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBucketKeys) GetScreenshotOk() (*string, bool) {
	if o == nil || o.Screenshot == nil {
		return nil, false
	}
	return o.Screenshot, true
}

// HasScreenshot returns a boolean if a field has been set.
func (o *SyntheticsTestResultBucketKeys) HasScreenshot() bool {
	return o != nil && o.Screenshot != nil
}

// SetScreenshot gets a reference to the given string and assigns it to the Screenshot field.
func (o *SyntheticsTestResultBucketKeys) SetScreenshot(v string) {
	o.Screenshot = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *SyntheticsTestResultBucketKeys) GetSnapshot() string {
	if o == nil || o.Snapshot == nil {
		var ret string
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBucketKeys) GetSnapshotOk() (*string, bool) {
	if o == nil || o.Snapshot == nil {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *SyntheticsTestResultBucketKeys) HasSnapshot() bool {
	return o != nil && o.Snapshot != nil
}

// SetSnapshot gets a reference to the given string and assigns it to the Snapshot field.
func (o *SyntheticsTestResultBucketKeys) SetSnapshot(v string) {
	o.Snapshot = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SyntheticsTestResultBucketKeys) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBucketKeys) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SyntheticsTestResultBucketKeys) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *SyntheticsTestResultBucketKeys) SetSource(v string) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultBucketKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AfterStepScreenshot != nil {
		toSerialize["after_step_screenshot"] = o.AfterStepScreenshot
	}
	if o.AfterTurnScreenshot != nil {
		toSerialize["after_turn_screenshot"] = o.AfterTurnScreenshot
	}
	if o.Artifacts != nil {
		toSerialize["artifacts"] = o.Artifacts
	}
	if o.BeforeStepScreenshot != nil {
		toSerialize["before_step_screenshot"] = o.BeforeStepScreenshot
	}
	if o.BeforeTurnScreenshot != nil {
		toSerialize["before_turn_screenshot"] = o.BeforeTurnScreenshot
	}
	if o.CrashReport != nil {
		toSerialize["crash_report"] = o.CrashReport
	}
	if o.DeviceLogs != nil {
		toSerialize["device_logs"] = o.DeviceLogs
	}
	if o.EmailMessages != nil {
		toSerialize["email_messages"] = o.EmailMessages
	}
	if o.Screenshot != nil {
		toSerialize["screenshot"] = o.Screenshot
	}
	if o.Snapshot != nil {
		toSerialize["snapshot"] = o.Snapshot
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultBucketKeys) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AfterStepScreenshot  *string  `json:"after_step_screenshot,omitempty"`
		AfterTurnScreenshot  *string  `json:"after_turn_screenshot,omitempty"`
		Artifacts            *string  `json:"artifacts,omitempty"`
		BeforeStepScreenshot *string  `json:"before_step_screenshot,omitempty"`
		BeforeTurnScreenshot *string  `json:"before_turn_screenshot,omitempty"`
		CrashReport          *string  `json:"crash_report,omitempty"`
		DeviceLogs           *string  `json:"device_logs,omitempty"`
		EmailMessages        []string `json:"email_messages,omitempty"`
		Screenshot           *string  `json:"screenshot,omitempty"`
		Snapshot             *string  `json:"snapshot,omitempty"`
		Source               *string  `json:"source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"after_step_screenshot", "after_turn_screenshot", "artifacts", "before_step_screenshot", "before_turn_screenshot", "crash_report", "device_logs", "email_messages", "screenshot", "snapshot", "source"})
	} else {
		return err
	}
	o.AfterStepScreenshot = all.AfterStepScreenshot
	o.AfterTurnScreenshot = all.AfterTurnScreenshot
	o.Artifacts = all.Artifacts
	o.BeforeStepScreenshot = all.BeforeStepScreenshot
	o.BeforeTurnScreenshot = all.BeforeTurnScreenshot
	o.CrashReport = all.CrashReport
	o.DeviceLogs = all.DeviceLogs
	o.EmailMessages = all.EmailMessages
	o.Screenshot = all.Screenshot
	o.Snapshot = all.Snapshot
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
