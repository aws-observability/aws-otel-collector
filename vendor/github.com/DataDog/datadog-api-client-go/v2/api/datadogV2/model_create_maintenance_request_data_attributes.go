// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateMaintenanceRequestDataAttributes The supported attributes for creating a maintenance.
type CreateMaintenanceRequestDataAttributes struct {
	// Timestamp of when the maintenance was completed.
	CompletedDate time.Time `json:"completed_date"`
	// The description shown when the maintenance is completed.
	CompletedDescription string `json:"completed_description"`
	// The components affected by the maintenance.
	ComponentsAffected []CreateMaintenanceRequestDataAttributesComponentsAffectedItems `json:"components_affected"`
	// The description shown while the maintenance is in progress.
	InProgressDescription string `json:"in_progress_description"`
	// The description shown when the maintenance is scheduled.
	ScheduledDescription string `json:"scheduled_description"`
	// Timestamp of when the maintenance is scheduled to start.
	StartDate time.Time `json:"start_date"`
	// The title of the maintenance.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateMaintenanceRequestDataAttributes instantiates a new CreateMaintenanceRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateMaintenanceRequestDataAttributes(completedDate time.Time, completedDescription string, componentsAffected []CreateMaintenanceRequestDataAttributesComponentsAffectedItems, inProgressDescription string, scheduledDescription string, startDate time.Time, title string) *CreateMaintenanceRequestDataAttributes {
	this := CreateMaintenanceRequestDataAttributes{}
	this.CompletedDate = completedDate
	this.CompletedDescription = completedDescription
	this.ComponentsAffected = componentsAffected
	this.InProgressDescription = inProgressDescription
	this.ScheduledDescription = scheduledDescription
	this.StartDate = startDate
	this.Title = title
	return &this
}

// NewCreateMaintenanceRequestDataAttributesWithDefaults instantiates a new CreateMaintenanceRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateMaintenanceRequestDataAttributesWithDefaults() *CreateMaintenanceRequestDataAttributes {
	this := CreateMaintenanceRequestDataAttributes{}
	return &this
}

// GetCompletedDate returns the CompletedDate field value.
func (o *CreateMaintenanceRequestDataAttributes) GetCompletedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CompletedDate
}

// GetCompletedDateOk returns a tuple with the CompletedDate field value
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceRequestDataAttributes) GetCompletedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedDate, true
}

// SetCompletedDate sets field value.
func (o *CreateMaintenanceRequestDataAttributes) SetCompletedDate(v time.Time) {
	o.CompletedDate = v
}

// GetCompletedDescription returns the CompletedDescription field value.
func (o *CreateMaintenanceRequestDataAttributes) GetCompletedDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CompletedDescription
}

// GetCompletedDescriptionOk returns a tuple with the CompletedDescription field value
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceRequestDataAttributes) GetCompletedDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedDescription, true
}

// SetCompletedDescription sets field value.
func (o *CreateMaintenanceRequestDataAttributes) SetCompletedDescription(v string) {
	o.CompletedDescription = v
}

// GetComponentsAffected returns the ComponentsAffected field value.
func (o *CreateMaintenanceRequestDataAttributes) GetComponentsAffected() []CreateMaintenanceRequestDataAttributesComponentsAffectedItems {
	if o == nil {
		var ret []CreateMaintenanceRequestDataAttributesComponentsAffectedItems
		return ret
	}
	return o.ComponentsAffected
}

// GetComponentsAffectedOk returns a tuple with the ComponentsAffected field value
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceRequestDataAttributes) GetComponentsAffectedOk() (*[]CreateMaintenanceRequestDataAttributesComponentsAffectedItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentsAffected, true
}

// SetComponentsAffected sets field value.
func (o *CreateMaintenanceRequestDataAttributes) SetComponentsAffected(v []CreateMaintenanceRequestDataAttributesComponentsAffectedItems) {
	o.ComponentsAffected = v
}

// GetInProgressDescription returns the InProgressDescription field value.
func (o *CreateMaintenanceRequestDataAttributes) GetInProgressDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InProgressDescription
}

// GetInProgressDescriptionOk returns a tuple with the InProgressDescription field value
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceRequestDataAttributes) GetInProgressDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InProgressDescription, true
}

// SetInProgressDescription sets field value.
func (o *CreateMaintenanceRequestDataAttributes) SetInProgressDescription(v string) {
	o.InProgressDescription = v
}

// GetScheduledDescription returns the ScheduledDescription field value.
func (o *CreateMaintenanceRequestDataAttributes) GetScheduledDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ScheduledDescription
}

// GetScheduledDescriptionOk returns a tuple with the ScheduledDescription field value
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceRequestDataAttributes) GetScheduledDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledDescription, true
}

// SetScheduledDescription sets field value.
func (o *CreateMaintenanceRequestDataAttributes) SetScheduledDescription(v string) {
	o.ScheduledDescription = v
}

// GetStartDate returns the StartDate field value.
func (o *CreateMaintenanceRequestDataAttributes) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceRequestDataAttributes) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value.
func (o *CreateMaintenanceRequestDataAttributes) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetTitle returns the Title field value.
func (o *CreateMaintenanceRequestDataAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceRequestDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *CreateMaintenanceRequestDataAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateMaintenanceRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompletedDate.Nanosecond() == 0 {
		toSerialize["completed_date"] = o.CompletedDate.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["completed_date"] = o.CompletedDate.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["completed_description"] = o.CompletedDescription
	toSerialize["components_affected"] = o.ComponentsAffected
	toSerialize["in_progress_description"] = o.InProgressDescription
	toSerialize["scheduled_description"] = o.ScheduledDescription
	if o.StartDate.Nanosecond() == 0 {
		toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateMaintenanceRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompletedDate         *time.Time                                                       `json:"completed_date"`
		CompletedDescription  *string                                                          `json:"completed_description"`
		ComponentsAffected    *[]CreateMaintenanceRequestDataAttributesComponentsAffectedItems `json:"components_affected"`
		InProgressDescription *string                                                          `json:"in_progress_description"`
		ScheduledDescription  *string                                                          `json:"scheduled_description"`
		StartDate             *time.Time                                                       `json:"start_date"`
		Title                 *string                                                          `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CompletedDate == nil {
		return fmt.Errorf("required field completed_date missing")
	}
	if all.CompletedDescription == nil {
		return fmt.Errorf("required field completed_description missing")
	}
	if all.ComponentsAffected == nil {
		return fmt.Errorf("required field components_affected missing")
	}
	if all.InProgressDescription == nil {
		return fmt.Errorf("required field in_progress_description missing")
	}
	if all.ScheduledDescription == nil {
		return fmt.Errorf("required field scheduled_description missing")
	}
	if all.StartDate == nil {
		return fmt.Errorf("required field start_date missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completed_date", "completed_description", "components_affected", "in_progress_description", "scheduled_description", "start_date", "title"})
	} else {
		return err
	}
	o.CompletedDate = *all.CompletedDate
	o.CompletedDescription = *all.CompletedDescription
	o.ComponentsAffected = *all.ComponentsAffected
	o.InProgressDescription = *all.InProgressDescription
	o.ScheduledDescription = *all.ScheduledDescription
	o.StartDate = *all.StartDate
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
