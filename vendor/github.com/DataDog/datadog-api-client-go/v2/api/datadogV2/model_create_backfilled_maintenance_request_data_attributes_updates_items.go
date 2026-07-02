// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems A backfilled maintenance update entry.
type CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems struct {
	// The components affected.
	ComponentsAffected []CreateMaintenanceRequestDataAttributesComponentsAffectedItems `json:"components_affected"`
	// A description of the update.
	Description string `json:"description"`
	// Timestamp of when the update occurred.
	StartedAt time.Time `json:"started_at"`
	// The status of a maintenance update.
	Status CreateMaintenanceRequestDataAttributesUpdatesItemsStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateBackfilledMaintenanceRequestDataAttributesUpdatesItems instantiates a new CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateBackfilledMaintenanceRequestDataAttributesUpdatesItems(componentsAffected []CreateMaintenanceRequestDataAttributesComponentsAffectedItems, description string, startedAt time.Time, status CreateMaintenanceRequestDataAttributesUpdatesItemsStatus) *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems {
	this := CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems{}
	this.ComponentsAffected = componentsAffected
	this.Description = description
	this.StartedAt = startedAt
	this.Status = status
	return &this
}

// NewCreateBackfilledMaintenanceRequestDataAttributesUpdatesItemsWithDefaults instantiates a new CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateBackfilledMaintenanceRequestDataAttributesUpdatesItemsWithDefaults() *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems {
	this := CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems{}
	return &this
}

// GetComponentsAffected returns the ComponentsAffected field value.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) GetComponentsAffected() []CreateMaintenanceRequestDataAttributesComponentsAffectedItems {
	if o == nil {
		var ret []CreateMaintenanceRequestDataAttributesComponentsAffectedItems
		return ret
	}
	return o.ComponentsAffected
}

// GetComponentsAffectedOk returns a tuple with the ComponentsAffected field value
// and a boolean to check if the value has been set.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) GetComponentsAffectedOk() (*[]CreateMaintenanceRequestDataAttributesComponentsAffectedItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentsAffected, true
}

// SetComponentsAffected sets field value.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) SetComponentsAffected(v []CreateMaintenanceRequestDataAttributesComponentsAffectedItems) {
	o.ComponentsAffected = v
}

// GetDescription returns the Description field value.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) SetDescription(v string) {
	o.Description = v
}

// GetStartedAt returns the StartedAt field value.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetStatus returns the Status field value.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) GetStatus() CreateMaintenanceRequestDataAttributesUpdatesItemsStatus {
	if o == nil {
		var ret CreateMaintenanceRequestDataAttributesUpdatesItemsStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) GetStatusOk() (*CreateMaintenanceRequestDataAttributesUpdatesItemsStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) SetStatus(v CreateMaintenanceRequestDataAttributesUpdatesItemsStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["components_affected"] = o.ComponentsAffected
	toSerialize["description"] = o.Description
	if o.StartedAt.Nanosecond() == 0 {
		toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateBackfilledMaintenanceRequestDataAttributesUpdatesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentsAffected *[]CreateMaintenanceRequestDataAttributesComponentsAffectedItems `json:"components_affected"`
		Description        *string                                                          `json:"description"`
		StartedAt          *time.Time                                                       `json:"started_at"`
		Status             *CreateMaintenanceRequestDataAttributesUpdatesItemsStatus        `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ComponentsAffected == nil {
		return fmt.Errorf("required field components_affected missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.StartedAt == nil {
		return fmt.Errorf("required field started_at missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components_affected", "description", "started_at", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComponentsAffected = *all.ComponentsAffected
	o.Description = *all.Description
	o.StartedAt = *all.StartedAt
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
