// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultSummaryAttributes Attributes of a Synthetic test result summary.
type SyntheticsTestResultSummaryAttributes struct {
	// Device information for the test result (browser and mobile tests).
	Device *SyntheticsTestResultDevice `json:"device,omitempty"`
	// Execution details for a Synthetic test result.
	ExecutionInfo *SyntheticsTestResultExecutionInfo `json:"execution_info,omitempty"`
	// Timestamp of when the test finished (in milliseconds).
	FinishedAt *int64 `json:"finished_at,omitempty"`
	// Location information for a Synthetic test result.
	Location *SyntheticsTestResultLocation `json:"location,omitempty"`
	// The type of run for a Synthetic test result.
	RunType *SyntheticsTestResultRunType `json:"run_type,omitempty"`
	// Timestamp of when the test started (in milliseconds).
	StartedAt *int64 `json:"started_at,omitempty"`
	// Status of a Synthetic test result.
	Status *SyntheticsTestResultStatus `json:"status,omitempty"`
	// Step execution summary for a Synthetic test result.
	StepsInfo *SyntheticsTestResultStepsInfo `json:"steps_info,omitempty"`
	// Subtype of the Synthetic test that produced this result.
	TestSubType *SyntheticsTestSubType `json:"test_sub_type,omitempty"`
	// Type of the Synthetic test that produced this result.
	TestType *SyntheticsTestType `json:"test_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultSummaryAttributes instantiates a new SyntheticsTestResultSummaryAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultSummaryAttributes() *SyntheticsTestResultSummaryAttributes {
	this := SyntheticsTestResultSummaryAttributes{}
	return &this
}

// NewSyntheticsTestResultSummaryAttributesWithDefaults instantiates a new SyntheticsTestResultSummaryAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultSummaryAttributesWithDefaults() *SyntheticsTestResultSummaryAttributes {
	this := SyntheticsTestResultSummaryAttributes{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *SyntheticsTestResultSummaryAttributes) GetDevice() SyntheticsTestResultDevice {
	if o == nil || o.Device == nil {
		var ret SyntheticsTestResultDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSummaryAttributes) GetDeviceOk() (*SyntheticsTestResultDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *SyntheticsTestResultSummaryAttributes) HasDevice() bool {
	return o != nil && o.Device != nil
}

// SetDevice gets a reference to the given SyntheticsTestResultDevice and assigns it to the Device field.
func (o *SyntheticsTestResultSummaryAttributes) SetDevice(v SyntheticsTestResultDevice) {
	o.Device = &v
}

// GetExecutionInfo returns the ExecutionInfo field value if set, zero value otherwise.
func (o *SyntheticsTestResultSummaryAttributes) GetExecutionInfo() SyntheticsTestResultExecutionInfo {
	if o == nil || o.ExecutionInfo == nil {
		var ret SyntheticsTestResultExecutionInfo
		return ret
	}
	return *o.ExecutionInfo
}

// GetExecutionInfoOk returns a tuple with the ExecutionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSummaryAttributes) GetExecutionInfoOk() (*SyntheticsTestResultExecutionInfo, bool) {
	if o == nil || o.ExecutionInfo == nil {
		return nil, false
	}
	return o.ExecutionInfo, true
}

// HasExecutionInfo returns a boolean if a field has been set.
func (o *SyntheticsTestResultSummaryAttributes) HasExecutionInfo() bool {
	return o != nil && o.ExecutionInfo != nil
}

// SetExecutionInfo gets a reference to the given SyntheticsTestResultExecutionInfo and assigns it to the ExecutionInfo field.
func (o *SyntheticsTestResultSummaryAttributes) SetExecutionInfo(v SyntheticsTestResultExecutionInfo) {
	o.ExecutionInfo = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *SyntheticsTestResultSummaryAttributes) GetFinishedAt() int64 {
	if o == nil || o.FinishedAt == nil {
		var ret int64
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSummaryAttributes) GetFinishedAtOk() (*int64, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *SyntheticsTestResultSummaryAttributes) HasFinishedAt() bool {
	return o != nil && o.FinishedAt != nil
}

// SetFinishedAt gets a reference to the given int64 and assigns it to the FinishedAt field.
func (o *SyntheticsTestResultSummaryAttributes) SetFinishedAt(v int64) {
	o.FinishedAt = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SyntheticsTestResultSummaryAttributes) GetLocation() SyntheticsTestResultLocation {
	if o == nil || o.Location == nil {
		var ret SyntheticsTestResultLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSummaryAttributes) GetLocationOk() (*SyntheticsTestResultLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SyntheticsTestResultSummaryAttributes) HasLocation() bool {
	return o != nil && o.Location != nil
}

// SetLocation gets a reference to the given SyntheticsTestResultLocation and assigns it to the Location field.
func (o *SyntheticsTestResultSummaryAttributes) SetLocation(v SyntheticsTestResultLocation) {
	o.Location = &v
}

// GetRunType returns the RunType field value if set, zero value otherwise.
func (o *SyntheticsTestResultSummaryAttributes) GetRunType() SyntheticsTestResultRunType {
	if o == nil || o.RunType == nil {
		var ret SyntheticsTestResultRunType
		return ret
	}
	return *o.RunType
}

// GetRunTypeOk returns a tuple with the RunType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSummaryAttributes) GetRunTypeOk() (*SyntheticsTestResultRunType, bool) {
	if o == nil || o.RunType == nil {
		return nil, false
	}
	return o.RunType, true
}

// HasRunType returns a boolean if a field has been set.
func (o *SyntheticsTestResultSummaryAttributes) HasRunType() bool {
	return o != nil && o.RunType != nil
}

// SetRunType gets a reference to the given SyntheticsTestResultRunType and assigns it to the RunType field.
func (o *SyntheticsTestResultSummaryAttributes) SetRunType(v SyntheticsTestResultRunType) {
	o.RunType = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *SyntheticsTestResultSummaryAttributes) GetStartedAt() int64 {
	if o == nil || o.StartedAt == nil {
		var ret int64
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSummaryAttributes) GetStartedAtOk() (*int64, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *SyntheticsTestResultSummaryAttributes) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given int64 and assigns it to the StartedAt field.
func (o *SyntheticsTestResultSummaryAttributes) SetStartedAt(v int64) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsTestResultSummaryAttributes) GetStatus() SyntheticsTestResultStatus {
	if o == nil || o.Status == nil {
		var ret SyntheticsTestResultStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSummaryAttributes) GetStatusOk() (*SyntheticsTestResultStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsTestResultSummaryAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SyntheticsTestResultStatus and assigns it to the Status field.
func (o *SyntheticsTestResultSummaryAttributes) SetStatus(v SyntheticsTestResultStatus) {
	o.Status = &v
}

// GetStepsInfo returns the StepsInfo field value if set, zero value otherwise.
func (o *SyntheticsTestResultSummaryAttributes) GetStepsInfo() SyntheticsTestResultStepsInfo {
	if o == nil || o.StepsInfo == nil {
		var ret SyntheticsTestResultStepsInfo
		return ret
	}
	return *o.StepsInfo
}

// GetStepsInfoOk returns a tuple with the StepsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSummaryAttributes) GetStepsInfoOk() (*SyntheticsTestResultStepsInfo, bool) {
	if o == nil || o.StepsInfo == nil {
		return nil, false
	}
	return o.StepsInfo, true
}

// HasStepsInfo returns a boolean if a field has been set.
func (o *SyntheticsTestResultSummaryAttributes) HasStepsInfo() bool {
	return o != nil && o.StepsInfo != nil
}

// SetStepsInfo gets a reference to the given SyntheticsTestResultStepsInfo and assigns it to the StepsInfo field.
func (o *SyntheticsTestResultSummaryAttributes) SetStepsInfo(v SyntheticsTestResultStepsInfo) {
	o.StepsInfo = &v
}

// GetTestSubType returns the TestSubType field value if set, zero value otherwise.
func (o *SyntheticsTestResultSummaryAttributes) GetTestSubType() SyntheticsTestSubType {
	if o == nil || o.TestSubType == nil {
		var ret SyntheticsTestSubType
		return ret
	}
	return *o.TestSubType
}

// GetTestSubTypeOk returns a tuple with the TestSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSummaryAttributes) GetTestSubTypeOk() (*SyntheticsTestSubType, bool) {
	if o == nil || o.TestSubType == nil {
		return nil, false
	}
	return o.TestSubType, true
}

// HasTestSubType returns a boolean if a field has been set.
func (o *SyntheticsTestResultSummaryAttributes) HasTestSubType() bool {
	return o != nil && o.TestSubType != nil
}

// SetTestSubType gets a reference to the given SyntheticsTestSubType and assigns it to the TestSubType field.
func (o *SyntheticsTestResultSummaryAttributes) SetTestSubType(v SyntheticsTestSubType) {
	o.TestSubType = &v
}

// GetTestType returns the TestType field value if set, zero value otherwise.
func (o *SyntheticsTestResultSummaryAttributes) GetTestType() SyntheticsTestType {
	if o == nil || o.TestType == nil {
		var ret SyntheticsTestType
		return ret
	}
	return *o.TestType
}

// GetTestTypeOk returns a tuple with the TestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultSummaryAttributes) GetTestTypeOk() (*SyntheticsTestType, bool) {
	if o == nil || o.TestType == nil {
		return nil, false
	}
	return o.TestType, true
}

// HasTestType returns a boolean if a field has been set.
func (o *SyntheticsTestResultSummaryAttributes) HasTestType() bool {
	return o != nil && o.TestType != nil
}

// SetTestType gets a reference to the given SyntheticsTestType and assigns it to the TestType field.
func (o *SyntheticsTestResultSummaryAttributes) SetTestType(v SyntheticsTestType) {
	o.TestType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultSummaryAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.ExecutionInfo != nil {
		toSerialize["execution_info"] = o.ExecutionInfo
	}
	if o.FinishedAt != nil {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.RunType != nil {
		toSerialize["run_type"] = o.RunType
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StepsInfo != nil {
		toSerialize["steps_info"] = o.StepsInfo
	}
	if o.TestSubType != nil {
		toSerialize["test_sub_type"] = o.TestSubType
	}
	if o.TestType != nil {
		toSerialize["test_type"] = o.TestType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultSummaryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Device        *SyntheticsTestResultDevice        `json:"device,omitempty"`
		ExecutionInfo *SyntheticsTestResultExecutionInfo `json:"execution_info,omitempty"`
		FinishedAt    *int64                             `json:"finished_at,omitempty"`
		Location      *SyntheticsTestResultLocation      `json:"location,omitempty"`
		RunType       *SyntheticsTestResultRunType       `json:"run_type,omitempty"`
		StartedAt     *int64                             `json:"started_at,omitempty"`
		Status        *SyntheticsTestResultStatus        `json:"status,omitempty"`
		StepsInfo     *SyntheticsTestResultStepsInfo     `json:"steps_info,omitempty"`
		TestSubType   *SyntheticsTestSubType             `json:"test_sub_type,omitempty"`
		TestType      *SyntheticsTestType                `json:"test_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"device", "execution_info", "finished_at", "location", "run_type", "started_at", "status", "steps_info", "test_sub_type", "test_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Device != nil && all.Device.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Device = all.Device
	if all.ExecutionInfo != nil && all.ExecutionInfo.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ExecutionInfo = all.ExecutionInfo
	o.FinishedAt = all.FinishedAt
	if all.Location != nil && all.Location.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Location = all.Location
	if all.RunType != nil && !all.RunType.IsValid() {
		hasInvalidField = true
	} else {
		o.RunType = all.RunType
	}
	o.StartedAt = all.StartedAt
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	if all.StepsInfo != nil && all.StepsInfo.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StepsInfo = all.StepsInfo
	if all.TestSubType != nil && !all.TestSubType.IsValid() {
		hasInvalidField = true
	} else {
		o.TestSubType = all.TestSubType
	}
	if all.TestType != nil && !all.TestType.IsValid() {
		hasInvalidField = true
	} else {
		o.TestType = all.TestType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
