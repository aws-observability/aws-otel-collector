// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchMaintenanceRequestDataAttributes The supported attributes for updating a maintenance.
type PatchMaintenanceRequestDataAttributes struct {
	// Timestamp of when the maintenance was completed.
	CompletedDate *time.Time `json:"completed_date,omitempty"`
	// The description shown when the maintenance is completed.
	CompletedDescription *string `json:"completed_description,omitempty"`
	// The components affected by the maintenance.
	ComponentsAffected []PatchMaintenanceRequestDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
	// The description shown while the maintenance is in progress.
	InProgressDescription *string `json:"in_progress_description,omitempty"`
	// The description shown when the maintenance is scheduled.
	ScheduledDescription *string `json:"scheduled_description,omitempty"`
	// Timestamp of when the maintenance is scheduled to start.
	StartDate *time.Time `json:"start_date,omitempty"`
	// The status of the maintenance.
	Status *MaintenanceDataAttributesStatus `json:"status,omitempty"`
	// The title of the maintenance.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchMaintenanceRequestDataAttributes instantiates a new PatchMaintenanceRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchMaintenanceRequestDataAttributes() *PatchMaintenanceRequestDataAttributes {
	this := PatchMaintenanceRequestDataAttributes{}
	return &this
}

// NewPatchMaintenanceRequestDataAttributesWithDefaults instantiates a new PatchMaintenanceRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchMaintenanceRequestDataAttributesWithDefaults() *PatchMaintenanceRequestDataAttributes {
	this := PatchMaintenanceRequestDataAttributes{}
	return &this
}

// GetCompletedDate returns the CompletedDate field value if set, zero value otherwise.
func (o *PatchMaintenanceRequestDataAttributes) GetCompletedDate() time.Time {
	if o == nil || o.CompletedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedDate
}

// GetCompletedDateOk returns a tuple with the CompletedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMaintenanceRequestDataAttributes) GetCompletedDateOk() (*time.Time, bool) {
	if o == nil || o.CompletedDate == nil {
		return nil, false
	}
	return o.CompletedDate, true
}

// HasCompletedDate returns a boolean if a field has been set.
func (o *PatchMaintenanceRequestDataAttributes) HasCompletedDate() bool {
	return o != nil && o.CompletedDate != nil
}

// SetCompletedDate gets a reference to the given time.Time and assigns it to the CompletedDate field.
func (o *PatchMaintenanceRequestDataAttributes) SetCompletedDate(v time.Time) {
	o.CompletedDate = &v
}

// GetCompletedDescription returns the CompletedDescription field value if set, zero value otherwise.
func (o *PatchMaintenanceRequestDataAttributes) GetCompletedDescription() string {
	if o == nil || o.CompletedDescription == nil {
		var ret string
		return ret
	}
	return *o.CompletedDescription
}

// GetCompletedDescriptionOk returns a tuple with the CompletedDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMaintenanceRequestDataAttributes) GetCompletedDescriptionOk() (*string, bool) {
	if o == nil || o.CompletedDescription == nil {
		return nil, false
	}
	return o.CompletedDescription, true
}

// HasCompletedDescription returns a boolean if a field has been set.
func (o *PatchMaintenanceRequestDataAttributes) HasCompletedDescription() bool {
	return o != nil && o.CompletedDescription != nil
}

// SetCompletedDescription gets a reference to the given string and assigns it to the CompletedDescription field.
func (o *PatchMaintenanceRequestDataAttributes) SetCompletedDescription(v string) {
	o.CompletedDescription = &v
}

// GetComponentsAffected returns the ComponentsAffected field value if set, zero value otherwise.
func (o *PatchMaintenanceRequestDataAttributes) GetComponentsAffected() []PatchMaintenanceRequestDataAttributesComponentsAffectedItems {
	if o == nil || o.ComponentsAffected == nil {
		var ret []PatchMaintenanceRequestDataAttributesComponentsAffectedItems
		return ret
	}
	return o.ComponentsAffected
}

// GetComponentsAffectedOk returns a tuple with the ComponentsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMaintenanceRequestDataAttributes) GetComponentsAffectedOk() (*[]PatchMaintenanceRequestDataAttributesComponentsAffectedItems, bool) {
	if o == nil || o.ComponentsAffected == nil {
		return nil, false
	}
	return &o.ComponentsAffected, true
}

// HasComponentsAffected returns a boolean if a field has been set.
func (o *PatchMaintenanceRequestDataAttributes) HasComponentsAffected() bool {
	return o != nil && o.ComponentsAffected != nil
}

// SetComponentsAffected gets a reference to the given []PatchMaintenanceRequestDataAttributesComponentsAffectedItems and assigns it to the ComponentsAffected field.
func (o *PatchMaintenanceRequestDataAttributes) SetComponentsAffected(v []PatchMaintenanceRequestDataAttributesComponentsAffectedItems) {
	o.ComponentsAffected = v
}

// GetInProgressDescription returns the InProgressDescription field value if set, zero value otherwise.
func (o *PatchMaintenanceRequestDataAttributes) GetInProgressDescription() string {
	if o == nil || o.InProgressDescription == nil {
		var ret string
		return ret
	}
	return *o.InProgressDescription
}

// GetInProgressDescriptionOk returns a tuple with the InProgressDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMaintenanceRequestDataAttributes) GetInProgressDescriptionOk() (*string, bool) {
	if o == nil || o.InProgressDescription == nil {
		return nil, false
	}
	return o.InProgressDescription, true
}

// HasInProgressDescription returns a boolean if a field has been set.
func (o *PatchMaintenanceRequestDataAttributes) HasInProgressDescription() bool {
	return o != nil && o.InProgressDescription != nil
}

// SetInProgressDescription gets a reference to the given string and assigns it to the InProgressDescription field.
func (o *PatchMaintenanceRequestDataAttributes) SetInProgressDescription(v string) {
	o.InProgressDescription = &v
}

// GetScheduledDescription returns the ScheduledDescription field value if set, zero value otherwise.
func (o *PatchMaintenanceRequestDataAttributes) GetScheduledDescription() string {
	if o == nil || o.ScheduledDescription == nil {
		var ret string
		return ret
	}
	return *o.ScheduledDescription
}

// GetScheduledDescriptionOk returns a tuple with the ScheduledDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMaintenanceRequestDataAttributes) GetScheduledDescriptionOk() (*string, bool) {
	if o == nil || o.ScheduledDescription == nil {
		return nil, false
	}
	return o.ScheduledDescription, true
}

// HasScheduledDescription returns a boolean if a field has been set.
func (o *PatchMaintenanceRequestDataAttributes) HasScheduledDescription() bool {
	return o != nil && o.ScheduledDescription != nil
}

// SetScheduledDescription gets a reference to the given string and assigns it to the ScheduledDescription field.
func (o *PatchMaintenanceRequestDataAttributes) SetScheduledDescription(v string) {
	o.ScheduledDescription = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *PatchMaintenanceRequestDataAttributes) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMaintenanceRequestDataAttributes) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *PatchMaintenanceRequestDataAttributes) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *PatchMaintenanceRequestDataAttributes) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchMaintenanceRequestDataAttributes) GetStatus() MaintenanceDataAttributesStatus {
	if o == nil || o.Status == nil {
		var ret MaintenanceDataAttributesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMaintenanceRequestDataAttributes) GetStatusOk() (*MaintenanceDataAttributesStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchMaintenanceRequestDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given MaintenanceDataAttributesStatus and assigns it to the Status field.
func (o *PatchMaintenanceRequestDataAttributes) SetStatus(v MaintenanceDataAttributesStatus) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PatchMaintenanceRequestDataAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMaintenanceRequestDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PatchMaintenanceRequestDataAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PatchMaintenanceRequestDataAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchMaintenanceRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompletedDate != nil {
		if o.CompletedDate.Nanosecond() == 0 {
			toSerialize["completed_date"] = o.CompletedDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["completed_date"] = o.CompletedDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CompletedDescription != nil {
		toSerialize["completed_description"] = o.CompletedDescription
	}
	if o.ComponentsAffected != nil {
		toSerialize["components_affected"] = o.ComponentsAffected
	}
	if o.InProgressDescription != nil {
		toSerialize["in_progress_description"] = o.InProgressDescription
	}
	if o.ScheduledDescription != nil {
		toSerialize["scheduled_description"] = o.ScheduledDescription
	}
	if o.StartDate != nil {
		if o.StartDate.Nanosecond() == 0 {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchMaintenanceRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompletedDate         *time.Time                                                     `json:"completed_date,omitempty"`
		CompletedDescription  *string                                                        `json:"completed_description,omitempty"`
		ComponentsAffected    []PatchMaintenanceRequestDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
		InProgressDescription *string                                                        `json:"in_progress_description,omitempty"`
		ScheduledDescription  *string                                                        `json:"scheduled_description,omitempty"`
		StartDate             *time.Time                                                     `json:"start_date,omitempty"`
		Status                *MaintenanceDataAttributesStatus                               `json:"status,omitempty"`
		Title                 *string                                                        `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completed_date", "completed_description", "components_affected", "in_progress_description", "scheduled_description", "start_date", "status", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CompletedDate = all.CompletedDate
	o.CompletedDescription = all.CompletedDescription
	o.ComponentsAffected = all.ComponentsAffected
	o.InProgressDescription = all.InProgressDescription
	o.ScheduledDescription = all.ScheduledDescription
	o.StartDate = all.StartDate
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
