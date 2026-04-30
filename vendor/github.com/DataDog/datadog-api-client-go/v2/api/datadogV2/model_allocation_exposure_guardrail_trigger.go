// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AllocationExposureGuardrailTrigger Guardrail trigger details for a progressive rollout.
type AllocationExposureGuardrailTrigger struct {
	// The progressive rollout ID this trigger belongs to.
	AllocationExposureScheduleId uuid.UUID `json:"allocation_exposure_schedule_id"`
	// The timestamp when this trigger was created.
	CreatedAt time.Time `json:"created_at"`
	// The variant ID that triggered this event.
	FlaggingVariantId uuid.UUID `json:"flagging_variant_id"`
	// The unique identifier of the guardrail trigger.
	Id uuid.UUID `json:"id"`
	// The metric ID associated with the trigger.
	MetricId string `json:"metric_id"`
	// The action that was triggered.
	TriggeredAction string `json:"triggered_action"`
	// The timestamp when this trigger was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAllocationExposureGuardrailTrigger instantiates a new AllocationExposureGuardrailTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAllocationExposureGuardrailTrigger(allocationExposureScheduleId uuid.UUID, createdAt time.Time, flaggingVariantId uuid.UUID, id uuid.UUID, metricId string, triggeredAction string, updatedAt time.Time) *AllocationExposureGuardrailTrigger {
	this := AllocationExposureGuardrailTrigger{}
	this.AllocationExposureScheduleId = allocationExposureScheduleId
	this.CreatedAt = createdAt
	this.FlaggingVariantId = flaggingVariantId
	this.Id = id
	this.MetricId = metricId
	this.TriggeredAction = triggeredAction
	this.UpdatedAt = updatedAt
	return &this
}

// NewAllocationExposureGuardrailTriggerWithDefaults instantiates a new AllocationExposureGuardrailTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAllocationExposureGuardrailTriggerWithDefaults() *AllocationExposureGuardrailTrigger {
	this := AllocationExposureGuardrailTrigger{}
	return &this
}

// GetAllocationExposureScheduleId returns the AllocationExposureScheduleId field value.
func (o *AllocationExposureGuardrailTrigger) GetAllocationExposureScheduleId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.AllocationExposureScheduleId
}

// GetAllocationExposureScheduleIdOk returns a tuple with the AllocationExposureScheduleId field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureGuardrailTrigger) GetAllocationExposureScheduleIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocationExposureScheduleId, true
}

// SetAllocationExposureScheduleId sets field value.
func (o *AllocationExposureGuardrailTrigger) SetAllocationExposureScheduleId(v uuid.UUID) {
	o.AllocationExposureScheduleId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AllocationExposureGuardrailTrigger) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureGuardrailTrigger) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AllocationExposureGuardrailTrigger) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFlaggingVariantId returns the FlaggingVariantId field value.
func (o *AllocationExposureGuardrailTrigger) GetFlaggingVariantId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FlaggingVariantId
}

// GetFlaggingVariantIdOk returns a tuple with the FlaggingVariantId field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureGuardrailTrigger) GetFlaggingVariantIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlaggingVariantId, true
}

// SetFlaggingVariantId sets field value.
func (o *AllocationExposureGuardrailTrigger) SetFlaggingVariantId(v uuid.UUID) {
	o.FlaggingVariantId = v
}

// GetId returns the Id field value.
func (o *AllocationExposureGuardrailTrigger) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureGuardrailTrigger) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AllocationExposureGuardrailTrigger) SetId(v uuid.UUID) {
	o.Id = v
}

// GetMetricId returns the MetricId field value.
func (o *AllocationExposureGuardrailTrigger) GetMetricId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureGuardrailTrigger) GetMetricIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value.
func (o *AllocationExposureGuardrailTrigger) SetMetricId(v string) {
	o.MetricId = v
}

// GetTriggeredAction returns the TriggeredAction field value.
func (o *AllocationExposureGuardrailTrigger) GetTriggeredAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TriggeredAction
}

// GetTriggeredActionOk returns a tuple with the TriggeredAction field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureGuardrailTrigger) GetTriggeredActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggeredAction, true
}

// SetTriggeredAction sets field value.
func (o *AllocationExposureGuardrailTrigger) SetTriggeredAction(v string) {
	o.TriggeredAction = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AllocationExposureGuardrailTrigger) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureGuardrailTrigger) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AllocationExposureGuardrailTrigger) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AllocationExposureGuardrailTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["allocation_exposure_schedule_id"] = o.AllocationExposureScheduleId
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["flagging_variant_id"] = o.FlaggingVariantId
	toSerialize["id"] = o.Id
	toSerialize["metric_id"] = o.MetricId
	toSerialize["triggered_action"] = o.TriggeredAction
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AllocationExposureGuardrailTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllocationExposureScheduleId *uuid.UUID `json:"allocation_exposure_schedule_id"`
		CreatedAt                    *time.Time `json:"created_at"`
		FlaggingVariantId            *uuid.UUID `json:"flagging_variant_id"`
		Id                           *uuid.UUID `json:"id"`
		MetricId                     *string    `json:"metric_id"`
		TriggeredAction              *string    `json:"triggered_action"`
		UpdatedAt                    *time.Time `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AllocationExposureScheduleId == nil {
		return fmt.Errorf("required field allocation_exposure_schedule_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.FlaggingVariantId == nil {
		return fmt.Errorf("required field flagging_variant_id missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.MetricId == nil {
		return fmt.Errorf("required field metric_id missing")
	}
	if all.TriggeredAction == nil {
		return fmt.Errorf("required field triggered_action missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allocation_exposure_schedule_id", "created_at", "flagging_variant_id", "id", "metric_id", "triggered_action", "updated_at"})
	} else {
		return err
	}
	o.AllocationExposureScheduleId = *all.AllocationExposureScheduleId
	o.CreatedAt = *all.CreatedAt
	o.FlaggingVariantId = *all.FlaggingVariantId
	o.Id = *all.Id
	o.MetricId = *all.MetricId
	o.TriggeredAction = *all.TriggeredAction
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
