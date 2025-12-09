// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTest Object containing details about a Synthetic mobile test.
type SyntheticsMobileTest struct {
	// Configuration object for a Synthetic mobile test.
	Config SyntheticsMobileTestConfig `json:"config"`
	// Array with the different device IDs used to run the test.
	DeviceIds []string `json:"device_ids,omitempty"`
	// Notification message associated with the test.
	Message string `json:"message"`
	// The associated monitor ID.
	MonitorId *int64 `json:"monitor_id,omitempty"`
	// Name of the test.
	Name string `json:"name"`
	// Object describing the extra options for a Synthetic test.
	Options SyntheticsMobileTestOptions `json:"options"`
	// The public ID of the test.
	PublicId *string `json:"public_id,omitempty"`
	// Define whether you want to start (`live`) or pause (`paused`) a
	// Synthetic test.
	Status *SyntheticsTestPauseStatus `json:"status,omitempty"`
	// Array of steps for the test.
	Steps []SyntheticsMobileStep `json:"steps,omitempty"`
	// Array of tags attached to the test.
	Tags []string `json:"tags,omitempty"`
	// Type of the Synthetic test, `mobile`.
	Type SyntheticsMobileTestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileTest instantiates a new SyntheticsMobileTest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileTest(config SyntheticsMobileTestConfig, message string, name string, options SyntheticsMobileTestOptions, typeVar SyntheticsMobileTestType) *SyntheticsMobileTest {
	this := SyntheticsMobileTest{}
	this.Config = config
	this.Message = message
	this.Name = name
	this.Options = options
	this.Type = typeVar
	return &this
}

// NewSyntheticsMobileTestWithDefaults instantiates a new SyntheticsMobileTest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileTestWithDefaults() *SyntheticsMobileTest {
	this := SyntheticsMobileTest{}
	var typeVar SyntheticsMobileTestType = SYNTHETICSMOBILETESTTYPE_MOBILE
	this.Type = typeVar
	return &this
}

// GetConfig returns the Config field value.
func (o *SyntheticsMobileTest) GetConfig() SyntheticsMobileTestConfig {
	if o == nil {
		var ret SyntheticsMobileTestConfig
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTest) GetConfigOk() (*SyntheticsMobileTestConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value.
func (o *SyntheticsMobileTest) SetConfig(v SyntheticsMobileTestConfig) {
	o.Config = v
}

// GetDeviceIds returns the DeviceIds field value if set, zero value otherwise.
func (o *SyntheticsMobileTest) GetDeviceIds() []string {
	if o == nil || o.DeviceIds == nil {
		var ret []string
		return ret
	}
	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTest) GetDeviceIdsOk() (*[]string, bool) {
	if o == nil || o.DeviceIds == nil {
		return nil, false
	}
	return &o.DeviceIds, true
}

// HasDeviceIds returns a boolean if a field has been set.
func (o *SyntheticsMobileTest) HasDeviceIds() bool {
	return o != nil && o.DeviceIds != nil
}

// SetDeviceIds gets a reference to the given []string and assigns it to the DeviceIds field.
func (o *SyntheticsMobileTest) SetDeviceIds(v []string) {
	o.DeviceIds = v
}

// GetMessage returns the Message field value.
func (o *SyntheticsMobileTest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *SyntheticsMobileTest) SetMessage(v string) {
	o.Message = v
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *SyntheticsMobileTest) GetMonitorId() int64 {
	if o == nil || o.MonitorId == nil {
		var ret int64
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTest) GetMonitorIdOk() (*int64, bool) {
	if o == nil || o.MonitorId == nil {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *SyntheticsMobileTest) HasMonitorId() bool {
	return o != nil && o.MonitorId != nil
}

// SetMonitorId gets a reference to the given int64 and assigns it to the MonitorId field.
func (o *SyntheticsMobileTest) SetMonitorId(v int64) {
	o.MonitorId = &v
}

// GetName returns the Name field value.
func (o *SyntheticsMobileTest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsMobileTest) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value.
func (o *SyntheticsMobileTest) GetOptions() SyntheticsMobileTestOptions {
	if o == nil {
		var ret SyntheticsMobileTestOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTest) GetOptionsOk() (*SyntheticsMobileTestOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *SyntheticsMobileTest) SetOptions(v SyntheticsMobileTestOptions) {
	o.Options = v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *SyntheticsMobileTest) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTest) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *SyntheticsMobileTest) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *SyntheticsMobileTest) SetPublicId(v string) {
	o.PublicId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsMobileTest) GetStatus() SyntheticsTestPauseStatus {
	if o == nil || o.Status == nil {
		var ret SyntheticsTestPauseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTest) GetStatusOk() (*SyntheticsTestPauseStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsMobileTest) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SyntheticsTestPauseStatus and assigns it to the Status field.
func (o *SyntheticsMobileTest) SetStatus(v SyntheticsTestPauseStatus) {
	o.Status = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *SyntheticsMobileTest) GetSteps() []SyntheticsMobileStep {
	if o == nil || o.Steps == nil {
		var ret []SyntheticsMobileStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTest) GetStepsOk() (*[]SyntheticsMobileStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *SyntheticsMobileTest) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []SyntheticsMobileStep and assigns it to the Steps field.
func (o *SyntheticsMobileTest) SetSteps(v []SyntheticsMobileStep) {
	o.Steps = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SyntheticsMobileTest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTest) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SyntheticsMobileTest) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SyntheticsMobileTest) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value.
func (o *SyntheticsMobileTest) GetType() SyntheticsMobileTestType {
	if o == nil {
		var ret SyntheticsMobileTestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTest) GetTypeOk() (*SyntheticsMobileTestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsMobileTest) SetType(v SyntheticsMobileTestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["config"] = o.Config
	if o.DeviceIds != nil {
		toSerialize["device_ids"] = o.DeviceIds
	}
	toSerialize["message"] = o.Message
	if o.MonitorId != nil {
		toSerialize["monitor_id"] = o.MonitorId
	}
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileTest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config    *SyntheticsMobileTestConfig  `json:"config"`
		DeviceIds []string                     `json:"device_ids,omitempty"`
		Message   *string                      `json:"message"`
		MonitorId *int64                       `json:"monitor_id,omitempty"`
		Name      *string                      `json:"name"`
		Options   *SyntheticsMobileTestOptions `json:"options"`
		PublicId  *string                      `json:"public_id,omitempty"`
		Status    *SyntheticsTestPauseStatus   `json:"status,omitempty"`
		Steps     []SyntheticsMobileStep       `json:"steps,omitempty"`
		Tags      []string                     `json:"tags,omitempty"`
		Type      *SyntheticsMobileTestType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Config == nil {
		return fmt.Errorf("required field config missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Options == nil {
		return fmt.Errorf("required field options missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config", "device_ids", "message", "monitor_id", "name", "options", "public_id", "status", "steps", "tags", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Config.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Config = *all.Config
	o.DeviceIds = all.DeviceIds
	o.Message = *all.Message
	o.MonitorId = all.MonitorId
	o.Name = *all.Name
	if all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = *all.Options
	o.PublicId = all.PublicId
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Steps = all.Steps
	o.Tags = all.Tags
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
