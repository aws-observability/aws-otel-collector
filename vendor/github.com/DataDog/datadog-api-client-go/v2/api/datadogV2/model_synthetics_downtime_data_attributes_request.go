// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsDowntimeDataAttributesRequest Attributes for creating or updating a Synthetics downtime.
type SyntheticsDowntimeDataAttributesRequest struct {
	// An optional description of the downtime.
	Description *string `json:"description,omitempty"`
	// Whether the downtime is enabled.
	IsEnabled bool `json:"isEnabled"`
	// The name of the downtime.
	Name string `json:"name"`
	// List of tags associated with a Synthetics downtime.
	Tags []string `json:"tags,omitempty"`
	// List of Synthetics test public IDs associated with a downtime.
	TestIds []string `json:"testIds"`
	// List of time slots for a Synthetics downtime create or update request.
	TimeSlots []SyntheticsDowntimeTimeSlotRequest `json:"timeSlots"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsDowntimeDataAttributesRequest instantiates a new SyntheticsDowntimeDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsDowntimeDataAttributesRequest(isEnabled bool, name string, testIds []string, timeSlots []SyntheticsDowntimeTimeSlotRequest) *SyntheticsDowntimeDataAttributesRequest {
	this := SyntheticsDowntimeDataAttributesRequest{}
	this.IsEnabled = isEnabled
	this.Name = name
	this.TestIds = testIds
	this.TimeSlots = timeSlots
	return &this
}

// NewSyntheticsDowntimeDataAttributesRequestWithDefaults instantiates a new SyntheticsDowntimeDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsDowntimeDataAttributesRequestWithDefaults() *SyntheticsDowntimeDataAttributesRequest {
	this := SyntheticsDowntimeDataAttributesRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SyntheticsDowntimeDataAttributesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SyntheticsDowntimeDataAttributesRequest) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SyntheticsDowntimeDataAttributesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsEnabled returns the IsEnabled field value.
func (o *SyntheticsDowntimeDataAttributesRequest) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value.
func (o *SyntheticsDowntimeDataAttributesRequest) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetName returns the Name field value.
func (o *SyntheticsDowntimeDataAttributesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsDowntimeDataAttributesRequest) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SyntheticsDowntimeDataAttributesRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesRequest) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SyntheticsDowntimeDataAttributesRequest) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SyntheticsDowntimeDataAttributesRequest) SetTags(v []string) {
	o.Tags = v
}

// GetTestIds returns the TestIds field value.
func (o *SyntheticsDowntimeDataAttributesRequest) GetTestIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TestIds
}

// GetTestIdsOk returns a tuple with the TestIds field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesRequest) GetTestIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestIds, true
}

// SetTestIds sets field value.
func (o *SyntheticsDowntimeDataAttributesRequest) SetTestIds(v []string) {
	o.TestIds = v
}

// GetTimeSlots returns the TimeSlots field value.
func (o *SyntheticsDowntimeDataAttributesRequest) GetTimeSlots() []SyntheticsDowntimeTimeSlotRequest {
	if o == nil {
		var ret []SyntheticsDowntimeTimeSlotRequest
		return ret
	}
	return o.TimeSlots
}

// GetTimeSlotsOk returns a tuple with the TimeSlots field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesRequest) GetTimeSlotsOk() (*[]SyntheticsDowntimeTimeSlotRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeSlots, true
}

// SetTimeSlots sets field value.
func (o *SyntheticsDowntimeDataAttributesRequest) SetTimeSlots(v []SyntheticsDowntimeTimeSlotRequest) {
	o.TimeSlots = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsDowntimeDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["name"] = o.Name
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["testIds"] = o.TestIds
	toSerialize["timeSlots"] = o.TimeSlots

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsDowntimeDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                              `json:"description,omitempty"`
		IsEnabled   *bool                                `json:"isEnabled"`
		Name        *string                              `json:"name"`
		Tags        []string                             `json:"tags,omitempty"`
		TestIds     *[]string                            `json:"testIds"`
		TimeSlots   *[]SyntheticsDowntimeTimeSlotRequest `json:"timeSlots"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsEnabled == nil {
		return fmt.Errorf("required field isEnabled missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.TestIds == nil {
		return fmt.Errorf("required field testIds missing")
	}
	if all.TimeSlots == nil {
		return fmt.Errorf("required field timeSlots missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "isEnabled", "name", "tags", "testIds", "timeSlots"})
	} else {
		return err
	}
	o.Description = all.Description
	o.IsEnabled = *all.IsEnabled
	o.Name = *all.Name
	o.Tags = all.Tags
	o.TestIds = *all.TestIds
	o.TimeSlots = *all.TimeSlots

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
