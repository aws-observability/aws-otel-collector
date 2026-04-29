// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsSuite Object containing details about a Synthetic suite.
type SyntheticsSuite struct {
	// Notification message associated with the suite.
	Message *string `json:"message,omitempty"`
	// The associated monitor ID.
	MonitorId *int64 `json:"monitor_id,omitempty"`
	// Name of the suite.
	Name string `json:"name"`
	// Object describing the extra options for a Synthetic suite.
	Options SyntheticsSuiteOptions `json:"options"`
	// The public ID for the test.
	PublicId *string `json:"public_id,omitempty"`
	// Array of tags attached to the suite.
	Tags []string `json:"tags,omitempty"`
	// Array of Synthetic tests included in the suite.
	Tests []SyntheticsSuiteTest `json:"tests"`
	// Type of the Synthetic suite, `suite`.
	Type SyntheticsSuiteType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsSuite instantiates a new SyntheticsSuite object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsSuite(name string, options SyntheticsSuiteOptions, tests []SyntheticsSuiteTest, typeVar SyntheticsSuiteType) *SyntheticsSuite {
	this := SyntheticsSuite{}
	this.Name = name
	this.Options = options
	this.Tests = tests
	this.Type = typeVar
	return &this
}

// NewSyntheticsSuiteWithDefaults instantiates a new SyntheticsSuite object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsSuiteWithDefaults() *SyntheticsSuite {
	this := SyntheticsSuite{}
	var typeVar SyntheticsSuiteType = SYNTHETICSSUITETYPE_SUITE
	this.Type = typeVar
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsSuite) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSuite) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsSuite) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SyntheticsSuite) SetMessage(v string) {
	o.Message = &v
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *SyntheticsSuite) GetMonitorId() int64 {
	if o == nil || o.MonitorId == nil {
		var ret int64
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSuite) GetMonitorIdOk() (*int64, bool) {
	if o == nil || o.MonitorId == nil {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *SyntheticsSuite) HasMonitorId() bool {
	return o != nil && o.MonitorId != nil
}

// SetMonitorId gets a reference to the given int64 and assigns it to the MonitorId field.
func (o *SyntheticsSuite) SetMonitorId(v int64) {
	o.MonitorId = &v
}

// GetName returns the Name field value.
func (o *SyntheticsSuite) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsSuite) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsSuite) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value.
func (o *SyntheticsSuite) GetOptions() SyntheticsSuiteOptions {
	if o == nil {
		var ret SyntheticsSuiteOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *SyntheticsSuite) GetOptionsOk() (*SyntheticsSuiteOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *SyntheticsSuite) SetOptions(v SyntheticsSuiteOptions) {
	o.Options = v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *SyntheticsSuite) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSuite) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *SyntheticsSuite) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *SyntheticsSuite) SetPublicId(v string) {
	o.PublicId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SyntheticsSuite) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSuite) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SyntheticsSuite) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SyntheticsSuite) SetTags(v []string) {
	o.Tags = v
}

// GetTests returns the Tests field value.
func (o *SyntheticsSuite) GetTests() []SyntheticsSuiteTest {
	if o == nil {
		var ret []SyntheticsSuiteTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value
// and a boolean to check if the value has been set.
func (o *SyntheticsSuite) GetTestsOk() (*[]SyntheticsSuiteTest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tests, true
}

// SetTests sets field value.
func (o *SyntheticsSuite) SetTests(v []SyntheticsSuiteTest) {
	o.Tests = v
}

// GetType returns the Type field value.
func (o *SyntheticsSuite) GetType() SyntheticsSuiteType {
	if o == nil {
		var ret SyntheticsSuiteType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsSuite) GetTypeOk() (*SyntheticsSuiteType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsSuite) SetType(v SyntheticsSuiteType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsSuite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.MonitorId != nil {
		toSerialize["monitor_id"] = o.MonitorId
	}
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["tests"] = o.Tests
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsSuite) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message   *string                 `json:"message,omitempty"`
		MonitorId *int64                  `json:"monitor_id,omitempty"`
		Name      *string                 `json:"name"`
		Options   *SyntheticsSuiteOptions `json:"options"`
		PublicId  *string                 `json:"public_id,omitempty"`
		Tags      []string                `json:"tags,omitempty"`
		Tests     *[]SyntheticsSuiteTest  `json:"tests"`
		Type      *SyntheticsSuiteType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Options == nil {
		return fmt.Errorf("required field options missing")
	}
	if all.Tests == nil {
		return fmt.Errorf("required field tests missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "monitor_id", "name", "options", "public_id", "tags", "tests", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Message = all.Message
	o.MonitorId = all.MonitorId
	o.Name = *all.Name
	if all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = *all.Options
	o.PublicId = all.PublicId
	o.Tags = all.Tags
	o.Tests = *all.Tests
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
