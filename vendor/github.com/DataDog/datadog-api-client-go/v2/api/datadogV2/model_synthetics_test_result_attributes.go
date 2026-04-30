// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultAttributes Attributes of a Synthetic test result.
type SyntheticsTestResultAttributes struct {
	// Batch information for the test result.
	Batch *SyntheticsTestResultBatch `json:"batch,omitempty"`
	// CI information associated with the test result.
	Ci *SyntheticsTestResultCI `json:"ci,omitempty"`
	// Device information for the test result (browser and mobile tests).
	Device *SyntheticsTestResultDevice `json:"device,omitempty"`
	// Git information associated with the test result.
	Git *SyntheticsTestResultGit `json:"git,omitempty"`
	// Location information for a Synthetic test result.
	Location *SyntheticsTestResultLocation `json:"location,omitempty"`
	// Full result details for a Synthetic test execution.
	Result *SyntheticsTestResultDetail `json:"result,omitempty"`
	// Subtype of the Synthetic test that produced this result.
	TestSubType *SyntheticsTestSubType `json:"test_sub_type,omitempty"`
	// Type of the Synthetic test that produced this result.
	TestType *SyntheticsTestType `json:"test_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultAttributes instantiates a new SyntheticsTestResultAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultAttributes() *SyntheticsTestResultAttributes {
	this := SyntheticsTestResultAttributes{}
	return &this
}

// NewSyntheticsTestResultAttributesWithDefaults instantiates a new SyntheticsTestResultAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultAttributesWithDefaults() *SyntheticsTestResultAttributes {
	this := SyntheticsTestResultAttributes{}
	return &this
}

// GetBatch returns the Batch field value if set, zero value otherwise.
func (o *SyntheticsTestResultAttributes) GetBatch() SyntheticsTestResultBatch {
	if o == nil || o.Batch == nil {
		var ret SyntheticsTestResultBatch
		return ret
	}
	return *o.Batch
}

// GetBatchOk returns a tuple with the Batch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAttributes) GetBatchOk() (*SyntheticsTestResultBatch, bool) {
	if o == nil || o.Batch == nil {
		return nil, false
	}
	return o.Batch, true
}

// HasBatch returns a boolean if a field has been set.
func (o *SyntheticsTestResultAttributes) HasBatch() bool {
	return o != nil && o.Batch != nil
}

// SetBatch gets a reference to the given SyntheticsTestResultBatch and assigns it to the Batch field.
func (o *SyntheticsTestResultAttributes) SetBatch(v SyntheticsTestResultBatch) {
	o.Batch = &v
}

// GetCi returns the Ci field value if set, zero value otherwise.
func (o *SyntheticsTestResultAttributes) GetCi() SyntheticsTestResultCI {
	if o == nil || o.Ci == nil {
		var ret SyntheticsTestResultCI
		return ret
	}
	return *o.Ci
}

// GetCiOk returns a tuple with the Ci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAttributes) GetCiOk() (*SyntheticsTestResultCI, bool) {
	if o == nil || o.Ci == nil {
		return nil, false
	}
	return o.Ci, true
}

// HasCi returns a boolean if a field has been set.
func (o *SyntheticsTestResultAttributes) HasCi() bool {
	return o != nil && o.Ci != nil
}

// SetCi gets a reference to the given SyntheticsTestResultCI and assigns it to the Ci field.
func (o *SyntheticsTestResultAttributes) SetCi(v SyntheticsTestResultCI) {
	o.Ci = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *SyntheticsTestResultAttributes) GetDevice() SyntheticsTestResultDevice {
	if o == nil || o.Device == nil {
		var ret SyntheticsTestResultDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAttributes) GetDeviceOk() (*SyntheticsTestResultDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *SyntheticsTestResultAttributes) HasDevice() bool {
	return o != nil && o.Device != nil
}

// SetDevice gets a reference to the given SyntheticsTestResultDevice and assigns it to the Device field.
func (o *SyntheticsTestResultAttributes) SetDevice(v SyntheticsTestResultDevice) {
	o.Device = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *SyntheticsTestResultAttributes) GetGit() SyntheticsTestResultGit {
	if o == nil || o.Git == nil {
		var ret SyntheticsTestResultGit
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAttributes) GetGitOk() (*SyntheticsTestResultGit, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *SyntheticsTestResultAttributes) HasGit() bool {
	return o != nil && o.Git != nil
}

// SetGit gets a reference to the given SyntheticsTestResultGit and assigns it to the Git field.
func (o *SyntheticsTestResultAttributes) SetGit(v SyntheticsTestResultGit) {
	o.Git = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SyntheticsTestResultAttributes) GetLocation() SyntheticsTestResultLocation {
	if o == nil || o.Location == nil {
		var ret SyntheticsTestResultLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAttributes) GetLocationOk() (*SyntheticsTestResultLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SyntheticsTestResultAttributes) HasLocation() bool {
	return o != nil && o.Location != nil
}

// SetLocation gets a reference to the given SyntheticsTestResultLocation and assigns it to the Location field.
func (o *SyntheticsTestResultAttributes) SetLocation(v SyntheticsTestResultLocation) {
	o.Location = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SyntheticsTestResultAttributes) GetResult() SyntheticsTestResultDetail {
	if o == nil || o.Result == nil {
		var ret SyntheticsTestResultDetail
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAttributes) GetResultOk() (*SyntheticsTestResultDetail, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SyntheticsTestResultAttributes) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given SyntheticsTestResultDetail and assigns it to the Result field.
func (o *SyntheticsTestResultAttributes) SetResult(v SyntheticsTestResultDetail) {
	o.Result = &v
}

// GetTestSubType returns the TestSubType field value if set, zero value otherwise.
func (o *SyntheticsTestResultAttributes) GetTestSubType() SyntheticsTestSubType {
	if o == nil || o.TestSubType == nil {
		var ret SyntheticsTestSubType
		return ret
	}
	return *o.TestSubType
}

// GetTestSubTypeOk returns a tuple with the TestSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAttributes) GetTestSubTypeOk() (*SyntheticsTestSubType, bool) {
	if o == nil || o.TestSubType == nil {
		return nil, false
	}
	return o.TestSubType, true
}

// HasTestSubType returns a boolean if a field has been set.
func (o *SyntheticsTestResultAttributes) HasTestSubType() bool {
	return o != nil && o.TestSubType != nil
}

// SetTestSubType gets a reference to the given SyntheticsTestSubType and assigns it to the TestSubType field.
func (o *SyntheticsTestResultAttributes) SetTestSubType(v SyntheticsTestSubType) {
	o.TestSubType = &v
}

// GetTestType returns the TestType field value if set, zero value otherwise.
func (o *SyntheticsTestResultAttributes) GetTestType() SyntheticsTestType {
	if o == nil || o.TestType == nil {
		var ret SyntheticsTestType
		return ret
	}
	return *o.TestType
}

// GetTestTypeOk returns a tuple with the TestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultAttributes) GetTestTypeOk() (*SyntheticsTestType, bool) {
	if o == nil || o.TestType == nil {
		return nil, false
	}
	return o.TestType, true
}

// HasTestType returns a boolean if a field has been set.
func (o *SyntheticsTestResultAttributes) HasTestType() bool {
	return o != nil && o.TestType != nil
}

// SetTestType gets a reference to the given SyntheticsTestType and assigns it to the TestType field.
func (o *SyntheticsTestResultAttributes) SetTestType(v SyntheticsTestType) {
	o.TestType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Batch != nil {
		toSerialize["batch"] = o.Batch
	}
	if o.Ci != nil {
		toSerialize["ci"] = o.Ci
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
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
func (o *SyntheticsTestResultAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Batch       *SyntheticsTestResultBatch    `json:"batch,omitempty"`
		Ci          *SyntheticsTestResultCI       `json:"ci,omitempty"`
		Device      *SyntheticsTestResultDevice   `json:"device,omitempty"`
		Git         *SyntheticsTestResultGit      `json:"git,omitempty"`
		Location    *SyntheticsTestResultLocation `json:"location,omitempty"`
		Result      *SyntheticsTestResultDetail   `json:"result,omitempty"`
		TestSubType *SyntheticsTestSubType        `json:"test_sub_type,omitempty"`
		TestType    *SyntheticsTestType           `json:"test_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"batch", "ci", "device", "git", "location", "result", "test_sub_type", "test_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Batch != nil && all.Batch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Batch = all.Batch
	if all.Ci != nil && all.Ci.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Ci = all.Ci
	if all.Device != nil && all.Device.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Device = all.Device
	if all.Git != nil && all.Git.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Git = all.Git
	if all.Location != nil && all.Location.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Location = all.Location
	if all.Result != nil && all.Result.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Result = all.Result
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
