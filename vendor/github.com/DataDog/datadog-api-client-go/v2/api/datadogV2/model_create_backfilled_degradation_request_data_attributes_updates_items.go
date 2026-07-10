// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateBackfilledDegradationRequestDataAttributesUpdatesItems A backfilled degradation update entry.
type CreateBackfilledDegradationRequestDataAttributesUpdatesItems struct {
	// The components affected.
	ComponentsAffected []CreateDegradationRequestDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
	// A description of the update.
	Description *string `json:"description,omitempty"`
	// Timestamp of when the update occurred.
	StartedAt time.Time `json:"started_at"`
	// The status of the degradation.
	Status CreateDegradationRequestDataAttributesStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateBackfilledDegradationRequestDataAttributesUpdatesItems instantiates a new CreateBackfilledDegradationRequestDataAttributesUpdatesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateBackfilledDegradationRequestDataAttributesUpdatesItems(startedAt time.Time, status CreateDegradationRequestDataAttributesStatus) *CreateBackfilledDegradationRequestDataAttributesUpdatesItems {
	this := CreateBackfilledDegradationRequestDataAttributesUpdatesItems{}
	this.StartedAt = startedAt
	this.Status = status
	return &this
}

// NewCreateBackfilledDegradationRequestDataAttributesUpdatesItemsWithDefaults instantiates a new CreateBackfilledDegradationRequestDataAttributesUpdatesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateBackfilledDegradationRequestDataAttributesUpdatesItemsWithDefaults() *CreateBackfilledDegradationRequestDataAttributesUpdatesItems {
	this := CreateBackfilledDegradationRequestDataAttributesUpdatesItems{}
	return &this
}

// GetComponentsAffected returns the ComponentsAffected field value if set, zero value otherwise.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) GetComponentsAffected() []CreateDegradationRequestDataAttributesComponentsAffectedItems {
	if o == nil || o.ComponentsAffected == nil {
		var ret []CreateDegradationRequestDataAttributesComponentsAffectedItems
		return ret
	}
	return o.ComponentsAffected
}

// GetComponentsAffectedOk returns a tuple with the ComponentsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) GetComponentsAffectedOk() (*[]CreateDegradationRequestDataAttributesComponentsAffectedItems, bool) {
	if o == nil || o.ComponentsAffected == nil {
		return nil, false
	}
	return &o.ComponentsAffected, true
}

// HasComponentsAffected returns a boolean if a field has been set.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) HasComponentsAffected() bool {
	return o != nil && o.ComponentsAffected != nil
}

// SetComponentsAffected gets a reference to the given []CreateDegradationRequestDataAttributesComponentsAffectedItems and assigns it to the ComponentsAffected field.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) SetComponentsAffected(v []CreateDegradationRequestDataAttributesComponentsAffectedItems) {
	o.ComponentsAffected = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) SetDescription(v string) {
	o.Description = &v
}

// GetStartedAt returns the StartedAt field value.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetStatus returns the Status field value.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) GetStatus() CreateDegradationRequestDataAttributesStatus {
	if o == nil {
		var ret CreateDegradationRequestDataAttributesStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) GetStatusOk() (*CreateDegradationRequestDataAttributesStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) SetStatus(v CreateDegradationRequestDataAttributesStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateBackfilledDegradationRequestDataAttributesUpdatesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComponentsAffected != nil {
		toSerialize["components_affected"] = o.ComponentsAffected
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
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
func (o *CreateBackfilledDegradationRequestDataAttributesUpdatesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentsAffected []CreateDegradationRequestDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
		Description        *string                                                         `json:"description,omitempty"`
		StartedAt          *time.Time                                                      `json:"started_at"`
		Status             *CreateDegradationRequestDataAttributesStatus                   `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.ComponentsAffected = all.ComponentsAffected
	o.Description = all.Description
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
