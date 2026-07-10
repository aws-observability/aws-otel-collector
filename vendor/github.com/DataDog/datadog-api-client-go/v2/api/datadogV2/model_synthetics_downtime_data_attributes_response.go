// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsDowntimeDataAttributesResponse Attributes of a Synthetics downtime response object.
type SyntheticsDowntimeDataAttributesResponse struct {
	// The timestamp when the downtime was created.
	CreatedAt time.Time `json:"createdAt"`
	// The UUID of the user who created the downtime.
	CreatedBy string `json:"createdBy"`
	// The display name of the user who created the downtime.
	CreatedByName string `json:"createdByName"`
	// The description of the downtime.
	Description string `json:"description"`
	// Whether the downtime is enabled.
	IsEnabled bool `json:"isEnabled"`
	// The name of the downtime.
	Name string `json:"name"`
	// List of tags associated with a Synthetics downtime.
	Tags []string `json:"tags"`
	// List of Synthetics test public IDs associated with a downtime.
	TestIds []string `json:"testIds"`
	// List of time slots in a Synthetics downtime response.
	TimeSlots []SyntheticsDowntimeTimeSlotResponse `json:"timeSlots"`
	// The timestamp when the downtime was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// The UUID of the user who last updated the downtime.
	UpdatedBy string `json:"updatedBy"`
	// The display name of the user who last updated the downtime.
	UpdatedByName string `json:"updatedByName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsDowntimeDataAttributesResponse instantiates a new SyntheticsDowntimeDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsDowntimeDataAttributesResponse(createdAt time.Time, createdBy string, createdByName string, description string, isEnabled bool, name string, tags []string, testIds []string, timeSlots []SyntheticsDowntimeTimeSlotResponse, updatedAt time.Time, updatedBy string, updatedByName string) *SyntheticsDowntimeDataAttributesResponse {
	this := SyntheticsDowntimeDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.CreatedByName = createdByName
	this.Description = description
	this.IsEnabled = isEnabled
	this.Name = name
	this.Tags = tags
	this.TestIds = testIds
	this.TimeSlots = timeSlots
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	this.UpdatedByName = updatedByName
	return &this
}

// NewSyntheticsDowntimeDataAttributesResponseWithDefaults instantiates a new SyntheticsDowntimeDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsDowntimeDataAttributesResponseWithDefaults() *SyntheticsDowntimeDataAttributesResponse {
	this := SyntheticsDowntimeDataAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedByName returns the CreatedByName field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetCreatedByName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetCreatedByNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByName, true
}

// SetCreatedByName sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetCreatedByName(v string) {
	o.CreatedByName = v
}

// GetDescription returns the Description field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetDescription(v string) {
	o.Description = v
}

// GetIsEnabled returns the IsEnabled field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetName returns the Name field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetTags(v []string) {
	o.Tags = v
}

// GetTestIds returns the TestIds field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetTestIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TestIds
}

// GetTestIdsOk returns a tuple with the TestIds field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetTestIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestIds, true
}

// SetTestIds sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetTestIds(v []string) {
	o.TestIds = v
}

// GetTimeSlots returns the TimeSlots field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetTimeSlots() []SyntheticsDowntimeTimeSlotResponse {
	if o == nil {
		var ret []SyntheticsDowntimeTimeSlotResponse
		return ret
	}
	return o.TimeSlots
}

// GetTimeSlotsOk returns a tuple with the TimeSlots field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetTimeSlotsOk() (*[]SyntheticsDowntimeTimeSlotResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeSlots, true
}

// SetTimeSlots sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetTimeSlots(v []SyntheticsDowntimeTimeSlotResponse) {
	o.TimeSlots = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// GetUpdatedByName returns the UpdatedByName field value.
func (o *SyntheticsDowntimeDataAttributesResponse) GetUpdatedByName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedByName
}

// GetUpdatedByNameOk returns a tuple with the UpdatedByName field value
// and a boolean to check if the value has been set.
func (o *SyntheticsDowntimeDataAttributesResponse) GetUpdatedByNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedByName, true
}

// SetUpdatedByName sets field value.
func (o *SyntheticsDowntimeDataAttributesResponse) SetUpdatedByName(v string) {
	o.UpdatedByName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsDowntimeDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["createdByName"] = o.CreatedByName
	toSerialize["description"] = o.Description
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["name"] = o.Name
	toSerialize["tags"] = o.Tags
	toSerialize["testIds"] = o.TestIds
	toSerialize["timeSlots"] = o.TimeSlots
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["updatedBy"] = o.UpdatedBy
	toSerialize["updatedByName"] = o.UpdatedByName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsDowntimeDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt     *time.Time                            `json:"createdAt"`
		CreatedBy     *string                               `json:"createdBy"`
		CreatedByName *string                               `json:"createdByName"`
		Description   *string                               `json:"description"`
		IsEnabled     *bool                                 `json:"isEnabled"`
		Name          *string                               `json:"name"`
		Tags          *[]string                             `json:"tags"`
		TestIds       *[]string                             `json:"testIds"`
		TimeSlots     *[]SyntheticsDowntimeTimeSlotResponse `json:"timeSlots"`
		UpdatedAt     *time.Time                            `json:"updatedAt"`
		UpdatedBy     *string                               `json:"updatedBy"`
		UpdatedByName *string                               `json:"updatedByName"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field createdBy missing")
	}
	if all.CreatedByName == nil {
		return fmt.Errorf("required field createdByName missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.IsEnabled == nil {
		return fmt.Errorf("required field isEnabled missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.TestIds == nil {
		return fmt.Errorf("required field testIds missing")
	}
	if all.TimeSlots == nil {
		return fmt.Errorf("required field timeSlots missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	if all.UpdatedBy == nil {
		return fmt.Errorf("required field updatedBy missing")
	}
	if all.UpdatedByName == nil {
		return fmt.Errorf("required field updatedByName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"createdAt", "createdBy", "createdByName", "description", "isEnabled", "name", "tags", "testIds", "timeSlots", "updatedAt", "updatedBy", "updatedByName"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.CreatedByName = *all.CreatedByName
	o.Description = *all.Description
	o.IsEnabled = *all.IsEnabled
	o.Name = *all.Name
	o.Tags = *all.Tags
	o.TestIds = *all.TestIds
	o.TimeSlots = *all.TimeSlots
	o.UpdatedAt = *all.UpdatedAt
	o.UpdatedBy = *all.UpdatedBy
	o.UpdatedByName = *all.UpdatedByName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
