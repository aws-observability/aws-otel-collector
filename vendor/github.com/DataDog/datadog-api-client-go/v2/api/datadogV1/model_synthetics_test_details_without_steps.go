// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestDetailsWithoutSteps Object containing details about your Synthetic test, without test steps.
type SyntheticsTestDetailsWithoutSteps struct {
	// Configuration object for a Synthetic test.
	Config *SyntheticsTestConfig `json:"config,omitempty"`
	// Object describing the creator of the shared element.
	Creator *Creator `json:"creator,omitempty"`
	// Array of locations used to run the test.
	Locations []string `json:"locations,omitempty"`
	// Notification message associated with the test.
	Message *string `json:"message,omitempty"`
	// The associated monitor ID.
	MonitorId *int64 `json:"monitor_id,omitempty"`
	// Name of the test.
	Name *string `json:"name,omitempty"`
	// Object describing the extra options for a Synthetic test.
	Options *SyntheticsTestOptions `json:"options,omitempty"`
	// The test public ID.
	PublicId *string `json:"public_id,omitempty"`
	// Define whether you want to start (`live`) or pause (`paused`) a
	// Synthetic test.
	Status *SyntheticsTestPauseStatus `json:"status,omitempty"`
	// The subtype of the Synthetic API test, `http`, `ssl`, `tcp`,
	// `dns`, `icmp`, `udp`, `websocket`, `grpc` or `multi`.
	Subtype *SyntheticsTestDetailsSubType `json:"subtype,omitempty"`
	// Array of tags attached to the test.
	Tags []string `json:"tags,omitempty"`
	// Type of the Synthetic test.
	Type *SyntheticsTestDetailsType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestDetailsWithoutSteps instantiates a new SyntheticsTestDetailsWithoutSteps object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestDetailsWithoutSteps() *SyntheticsTestDetailsWithoutSteps {
	this := SyntheticsTestDetailsWithoutSteps{}
	return &this
}

// NewSyntheticsTestDetailsWithoutStepsWithDefaults instantiates a new SyntheticsTestDetailsWithoutSteps object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestDetailsWithoutStepsWithDefaults() *SyntheticsTestDetailsWithoutSteps {
	this := SyntheticsTestDetailsWithoutSteps{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetConfig() SyntheticsTestConfig {
	if o == nil || o.Config == nil {
		var ret SyntheticsTestConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetConfigOk() (*SyntheticsTestConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given SyntheticsTestConfig and assigns it to the Config field.
func (o *SyntheticsTestDetailsWithoutSteps) SetConfig(v SyntheticsTestConfig) {
	o.Config = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetCreator() Creator {
	if o == nil || o.Creator == nil {
		var ret Creator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetCreatorOk() (*Creator, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given Creator and assigns it to the Creator field.
func (o *SyntheticsTestDetailsWithoutSteps) SetCreator(v Creator) {
	o.Creator = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetLocations() []string {
	if o == nil || o.Locations == nil {
		var ret []string
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetLocationsOk() (*[]string, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return &o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasLocations() bool {
	return o != nil && o.Locations != nil
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *SyntheticsTestDetailsWithoutSteps) SetLocations(v []string) {
	o.Locations = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SyntheticsTestDetailsWithoutSteps) SetMessage(v string) {
	o.Message = &v
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetMonitorId() int64 {
	if o == nil || o.MonitorId == nil {
		var ret int64
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetMonitorIdOk() (*int64, bool) {
	if o == nil || o.MonitorId == nil {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasMonitorId() bool {
	return o != nil && o.MonitorId != nil
}

// SetMonitorId gets a reference to the given int64 and assigns it to the MonitorId field.
func (o *SyntheticsTestDetailsWithoutSteps) SetMonitorId(v int64) {
	o.MonitorId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestDetailsWithoutSteps) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetOptions() SyntheticsTestOptions {
	if o == nil || o.Options == nil {
		var ret SyntheticsTestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetOptionsOk() (*SyntheticsTestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given SyntheticsTestOptions and assigns it to the Options field.
func (o *SyntheticsTestDetailsWithoutSteps) SetOptions(v SyntheticsTestOptions) {
	o.Options = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *SyntheticsTestDetailsWithoutSteps) SetPublicId(v string) {
	o.PublicId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetStatus() SyntheticsTestPauseStatus {
	if o == nil || o.Status == nil {
		var ret SyntheticsTestPauseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetStatusOk() (*SyntheticsTestPauseStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SyntheticsTestPauseStatus and assigns it to the Status field.
func (o *SyntheticsTestDetailsWithoutSteps) SetStatus(v SyntheticsTestPauseStatus) {
	o.Status = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetSubtype() SyntheticsTestDetailsSubType {
	if o == nil || o.Subtype == nil {
		var ret SyntheticsTestDetailsSubType
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetSubtypeOk() (*SyntheticsTestDetailsSubType, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasSubtype() bool {
	return o != nil && o.Subtype != nil
}

// SetSubtype gets a reference to the given SyntheticsTestDetailsSubType and assigns it to the Subtype field.
func (o *SyntheticsTestDetailsWithoutSteps) SetSubtype(v SyntheticsTestDetailsSubType) {
	o.Subtype = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SyntheticsTestDetailsWithoutSteps) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestDetailsWithoutSteps) GetType() SyntheticsTestDetailsType {
	if o == nil || o.Type == nil {
		var ret SyntheticsTestDetailsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestDetailsWithoutSteps) GetTypeOk() (*SyntheticsTestDetailsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestDetailsWithoutSteps) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SyntheticsTestDetailsType and assigns it to the Type field.
func (o *SyntheticsTestDetailsWithoutSteps) SetType(v SyntheticsTestDetailsType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestDetailsWithoutSteps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.MonitorId != nil {
		toSerialize["monitor_id"] = o.MonitorId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestDetailsWithoutSteps) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config    *SyntheticsTestConfig         `json:"config,omitempty"`
		Creator   *Creator                      `json:"creator,omitempty"`
		Locations []string                      `json:"locations,omitempty"`
		Message   *string                       `json:"message,omitempty"`
		MonitorId *int64                        `json:"monitor_id,omitempty"`
		Name      *string                       `json:"name,omitempty"`
		Options   *SyntheticsTestOptions        `json:"options,omitempty"`
		PublicId  *string                       `json:"public_id,omitempty"`
		Status    *SyntheticsTestPauseStatus    `json:"status,omitempty"`
		Subtype   *SyntheticsTestDetailsSubType `json:"subtype,omitempty"`
		Tags      []string                      `json:"tags,omitempty"`
		Type      *SyntheticsTestDetailsType    `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config", "creator", "locations", "message", "monitor_id", "name", "options", "public_id", "status", "subtype", "tags", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Config != nil && all.Config.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Config = all.Config
	if all.Creator != nil && all.Creator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Creator = all.Creator
	o.Locations = all.Locations
	o.Message = all.Message
	o.MonitorId = all.MonitorId
	o.Name = all.Name
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
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
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
