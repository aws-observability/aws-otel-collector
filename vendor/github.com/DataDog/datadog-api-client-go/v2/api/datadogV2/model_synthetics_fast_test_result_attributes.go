// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsFastTestResultAttributes Attributes of the fast test result.
type SyntheticsFastTestResultAttributes struct {
	// Device information for the test result (browser and mobile tests).
	Device *SyntheticsTestResultDevice `json:"device,omitempty"`
	// Location information for a Synthetic test result.
	Location *SyntheticsTestResultLocation `json:"location,omitempty"`
	// Detailed result data for the fast test run. The exact shape of nested fields
	// (`request`, `response`, `assertions`, etc.) depends on the test subtype.
	Result *SyntheticsFastTestResultDetail `json:"result,omitempty"`
	// Subtype of the Synthetic test that produced this result.
	TestSubType *SyntheticsFastTestSubType `json:"test_sub_type,omitempty"`
	// Type of the Synthetic fast test that produced this result.
	TestType *SyntheticsFastTestType `json:"test_type,omitempty"`
	// Version of the test at the time the fast test was triggered.
	TestVersion *int64 `json:"test_version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsFastTestResultAttributes instantiates a new SyntheticsFastTestResultAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsFastTestResultAttributes() *SyntheticsFastTestResultAttributes {
	this := SyntheticsFastTestResultAttributes{}
	return &this
}

// NewSyntheticsFastTestResultAttributesWithDefaults instantiates a new SyntheticsFastTestResultAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsFastTestResultAttributesWithDefaults() *SyntheticsFastTestResultAttributes {
	this := SyntheticsFastTestResultAttributes{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultAttributes) GetDevice() SyntheticsTestResultDevice {
	if o == nil || o.Device == nil {
		var ret SyntheticsTestResultDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultAttributes) GetDeviceOk() (*SyntheticsTestResultDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultAttributes) HasDevice() bool {
	return o != nil && o.Device != nil
}

// SetDevice gets a reference to the given SyntheticsTestResultDevice and assigns it to the Device field.
func (o *SyntheticsFastTestResultAttributes) SetDevice(v SyntheticsTestResultDevice) {
	o.Device = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultAttributes) GetLocation() SyntheticsTestResultLocation {
	if o == nil || o.Location == nil {
		var ret SyntheticsTestResultLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultAttributes) GetLocationOk() (*SyntheticsTestResultLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultAttributes) HasLocation() bool {
	return o != nil && o.Location != nil
}

// SetLocation gets a reference to the given SyntheticsTestResultLocation and assigns it to the Location field.
func (o *SyntheticsFastTestResultAttributes) SetLocation(v SyntheticsTestResultLocation) {
	o.Location = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultAttributes) GetResult() SyntheticsFastTestResultDetail {
	if o == nil || o.Result == nil {
		var ret SyntheticsFastTestResultDetail
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultAttributes) GetResultOk() (*SyntheticsFastTestResultDetail, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultAttributes) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given SyntheticsFastTestResultDetail and assigns it to the Result field.
func (o *SyntheticsFastTestResultAttributes) SetResult(v SyntheticsFastTestResultDetail) {
	o.Result = &v
}

// GetTestSubType returns the TestSubType field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultAttributes) GetTestSubType() SyntheticsFastTestSubType {
	if o == nil || o.TestSubType == nil {
		var ret SyntheticsFastTestSubType
		return ret
	}
	return *o.TestSubType
}

// GetTestSubTypeOk returns a tuple with the TestSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultAttributes) GetTestSubTypeOk() (*SyntheticsFastTestSubType, bool) {
	if o == nil || o.TestSubType == nil {
		return nil, false
	}
	return o.TestSubType, true
}

// HasTestSubType returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultAttributes) HasTestSubType() bool {
	return o != nil && o.TestSubType != nil
}

// SetTestSubType gets a reference to the given SyntheticsFastTestSubType and assigns it to the TestSubType field.
func (o *SyntheticsFastTestResultAttributes) SetTestSubType(v SyntheticsFastTestSubType) {
	o.TestSubType = &v
}

// GetTestType returns the TestType field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultAttributes) GetTestType() SyntheticsFastTestType {
	if o == nil || o.TestType == nil {
		var ret SyntheticsFastTestType
		return ret
	}
	return *o.TestType
}

// GetTestTypeOk returns a tuple with the TestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultAttributes) GetTestTypeOk() (*SyntheticsFastTestType, bool) {
	if o == nil || o.TestType == nil {
		return nil, false
	}
	return o.TestType, true
}

// HasTestType returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultAttributes) HasTestType() bool {
	return o != nil && o.TestType != nil
}

// SetTestType gets a reference to the given SyntheticsFastTestType and assigns it to the TestType field.
func (o *SyntheticsFastTestResultAttributes) SetTestType(v SyntheticsFastTestType) {
	o.TestType = &v
}

// GetTestVersion returns the TestVersion field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultAttributes) GetTestVersion() int64 {
	if o == nil || o.TestVersion == nil {
		var ret int64
		return ret
	}
	return *o.TestVersion
}

// GetTestVersionOk returns a tuple with the TestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultAttributes) GetTestVersionOk() (*int64, bool) {
	if o == nil || o.TestVersion == nil {
		return nil, false
	}
	return o.TestVersion, true
}

// HasTestVersion returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultAttributes) HasTestVersion() bool {
	return o != nil && o.TestVersion != nil
}

// SetTestVersion gets a reference to the given int64 and assigns it to the TestVersion field.
func (o *SyntheticsFastTestResultAttributes) SetTestVersion(v int64) {
	o.TestVersion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsFastTestResultAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
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
	if o.TestVersion != nil {
		toSerialize["test_version"] = o.TestVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsFastTestResultAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Device      *SyntheticsTestResultDevice     `json:"device,omitempty"`
		Location    *SyntheticsTestResultLocation   `json:"location,omitempty"`
		Result      *SyntheticsFastTestResultDetail `json:"result,omitempty"`
		TestSubType *SyntheticsFastTestSubType      `json:"test_sub_type,omitempty"`
		TestType    *SyntheticsFastTestType         `json:"test_type,omitempty"`
		TestVersion *int64                          `json:"test_version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"device", "location", "result", "test_sub_type", "test_type", "test_version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Device != nil && all.Device.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Device = all.Device
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
	o.TestVersion = all.TestVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
