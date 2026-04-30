// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkTest Object containing details about a Network Path test.
type SyntheticsNetworkTest struct {
	// Configuration object for a Network Path test.
	Config SyntheticsNetworkTestConfig `json:"config"`
	// Array of locations used to run the test. Network Path tests can be run from managed locations to test public endpoints,
	// or from a [Datadog Agent](https://docs.datadoghq.com/synthetics/network_path_tests/#agent-configuration) to test private environments.
	Locations []string `json:"locations"`
	// Notification message associated with the test.
	Message string `json:"message"`
	// The associated monitor ID.
	MonitorId *int64 `json:"monitor_id,omitempty"`
	// Name of the test.
	Name string `json:"name"`
	// Object describing the extra options for a Synthetic test.
	Options SyntheticsTestOptions `json:"options"`
	// The public ID for the test.
	PublicId *string `json:"public_id,omitempty"`
	// Define whether you want to start (`live`) or pause (`paused`) a
	// Synthetic test.
	Status *SyntheticsTestPauseStatus `json:"status,omitempty"`
	// Subtype of the Synthetic Network Path test: `tcp`, `udp`, or `icmp`.
	Subtype *SyntheticsNetworkTestSubType `json:"subtype,omitempty"`
	// Array of tags attached to the test.
	Tags []string `json:"tags,omitempty"`
	// Type of the Synthetic test, `network`.
	Type SyntheticsNetworkTestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsNetworkTest instantiates a new SyntheticsNetworkTest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsNetworkTest(config SyntheticsNetworkTestConfig, locations []string, message string, name string, options SyntheticsTestOptions, typeVar SyntheticsNetworkTestType) *SyntheticsNetworkTest {
	this := SyntheticsNetworkTest{}
	this.Config = config
	this.Locations = locations
	this.Message = message
	this.Name = name
	this.Options = options
	this.Type = typeVar
	return &this
}

// NewSyntheticsNetworkTestWithDefaults instantiates a new SyntheticsNetworkTest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsNetworkTestWithDefaults() *SyntheticsNetworkTest {
	this := SyntheticsNetworkTest{}
	var typeVar SyntheticsNetworkTestType = SYNTHETICSNETWORKTESTTYPE_NETWORK
	this.Type = typeVar
	return &this
}

// GetConfig returns the Config field value.
func (o *SyntheticsNetworkTest) GetConfig() SyntheticsNetworkTestConfig {
	if o == nil {
		var ret SyntheticsNetworkTestConfig
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTest) GetConfigOk() (*SyntheticsNetworkTestConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value.
func (o *SyntheticsNetworkTest) SetConfig(v SyntheticsNetworkTestConfig) {
	o.Config = v
}

// GetLocations returns the Locations field value.
func (o *SyntheticsNetworkTest) GetLocations() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTest) GetLocationsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locations, true
}

// SetLocations sets field value.
func (o *SyntheticsNetworkTest) SetLocations(v []string) {
	o.Locations = v
}

// GetMessage returns the Message field value.
func (o *SyntheticsNetworkTest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *SyntheticsNetworkTest) SetMessage(v string) {
	o.Message = v
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *SyntheticsNetworkTest) GetMonitorId() int64 {
	if o == nil || o.MonitorId == nil {
		var ret int64
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTest) GetMonitorIdOk() (*int64, bool) {
	if o == nil || o.MonitorId == nil {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *SyntheticsNetworkTest) HasMonitorId() bool {
	return o != nil && o.MonitorId != nil
}

// SetMonitorId gets a reference to the given int64 and assigns it to the MonitorId field.
func (o *SyntheticsNetworkTest) SetMonitorId(v int64) {
	o.MonitorId = &v
}

// GetName returns the Name field value.
func (o *SyntheticsNetworkTest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsNetworkTest) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value.
func (o *SyntheticsNetworkTest) GetOptions() SyntheticsTestOptions {
	if o == nil {
		var ret SyntheticsTestOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTest) GetOptionsOk() (*SyntheticsTestOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *SyntheticsNetworkTest) SetOptions(v SyntheticsTestOptions) {
	o.Options = v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *SyntheticsNetworkTest) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTest) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *SyntheticsNetworkTest) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *SyntheticsNetworkTest) SetPublicId(v string) {
	o.PublicId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsNetworkTest) GetStatus() SyntheticsTestPauseStatus {
	if o == nil || o.Status == nil {
		var ret SyntheticsTestPauseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTest) GetStatusOk() (*SyntheticsTestPauseStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsNetworkTest) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SyntheticsTestPauseStatus and assigns it to the Status field.
func (o *SyntheticsNetworkTest) SetStatus(v SyntheticsTestPauseStatus) {
	o.Status = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *SyntheticsNetworkTest) GetSubtype() SyntheticsNetworkTestSubType {
	if o == nil || o.Subtype == nil {
		var ret SyntheticsNetworkTestSubType
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTest) GetSubtypeOk() (*SyntheticsNetworkTestSubType, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *SyntheticsNetworkTest) HasSubtype() bool {
	return o != nil && o.Subtype != nil
}

// SetSubtype gets a reference to the given SyntheticsNetworkTestSubType and assigns it to the Subtype field.
func (o *SyntheticsNetworkTest) SetSubtype(v SyntheticsNetworkTestSubType) {
	o.Subtype = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SyntheticsNetworkTest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTest) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SyntheticsNetworkTest) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SyntheticsNetworkTest) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value.
func (o *SyntheticsNetworkTest) GetType() SyntheticsNetworkTestType {
	if o == nil {
		var ret SyntheticsNetworkTestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTest) GetTypeOk() (*SyntheticsNetworkTestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsNetworkTest) SetType(v SyntheticsNetworkTestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsNetworkTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["config"] = o.Config
	toSerialize["locations"] = o.Locations
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
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
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
func (o *SyntheticsNetworkTest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config    *SyntheticsNetworkTestConfig  `json:"config"`
		Locations *[]string                     `json:"locations"`
		Message   *string                       `json:"message"`
		MonitorId *int64                        `json:"monitor_id,omitempty"`
		Name      *string                       `json:"name"`
		Options   *SyntheticsTestOptions        `json:"options"`
		PublicId  *string                       `json:"public_id,omitempty"`
		Status    *SyntheticsTestPauseStatus    `json:"status,omitempty"`
		Subtype   *SyntheticsNetworkTestSubType `json:"subtype,omitempty"`
		Tags      []string                      `json:"tags,omitempty"`
		Type      *SyntheticsNetworkTestType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Config == nil {
		return fmt.Errorf("required field config missing")
	}
	if all.Locations == nil {
		return fmt.Errorf("required field locations missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"config", "locations", "message", "monitor_id", "name", "options", "public_id", "status", "subtype", "tags", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Config.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Config = *all.Config
	o.Locations = *all.Locations
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
	if all.Subtype != nil && !all.Subtype.IsValid() {
		hasInvalidField = true
	} else {
		o.Subtype = all.Subtype
	}
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
