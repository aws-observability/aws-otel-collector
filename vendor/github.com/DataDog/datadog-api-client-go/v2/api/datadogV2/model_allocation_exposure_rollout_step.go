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

// AllocationExposureRolloutStep Exposure progression step details.
type AllocationExposureRolloutStep struct {
	// The progressive rollout ID this step belongs to.
	AllocationExposureScheduleId uuid.UUID `json:"allocation_exposure_schedule_id"`
	// The timestamp when the progression step was created.
	CreatedAt time.Time `json:"created_at"`
	// The exposure ratio for this step.
	ExposureRatio float64 `json:"exposure_ratio"`
	// Logical index grouping related steps.
	GroupedStepIndex int64 `json:"grouped_step_index"`
	// The unique identifier of the progression step.
	Id uuid.UUID `json:"id"`
	// Step duration in milliseconds.
	IntervalMs datadog.NullableInt64 `json:"interval_ms,omitempty"`
	// Whether this step represents a pause record.
	IsPauseRecord bool `json:"is_pause_record"`
	// Sort order for the progression step.
	OrderPosition int64 `json:"order_position"`
	// The timestamp when the progression step was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAllocationExposureRolloutStep instantiates a new AllocationExposureRolloutStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAllocationExposureRolloutStep(allocationExposureScheduleId uuid.UUID, createdAt time.Time, exposureRatio float64, groupedStepIndex int64, id uuid.UUID, isPauseRecord bool, orderPosition int64, updatedAt time.Time) *AllocationExposureRolloutStep {
	this := AllocationExposureRolloutStep{}
	this.AllocationExposureScheduleId = allocationExposureScheduleId
	this.CreatedAt = createdAt
	this.ExposureRatio = exposureRatio
	this.GroupedStepIndex = groupedStepIndex
	this.Id = id
	this.IsPauseRecord = isPauseRecord
	this.OrderPosition = orderPosition
	this.UpdatedAt = updatedAt
	return &this
}

// NewAllocationExposureRolloutStepWithDefaults instantiates a new AllocationExposureRolloutStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAllocationExposureRolloutStepWithDefaults() *AllocationExposureRolloutStep {
	this := AllocationExposureRolloutStep{}
	return &this
}

// GetAllocationExposureScheduleId returns the AllocationExposureScheduleId field value.
func (o *AllocationExposureRolloutStep) GetAllocationExposureScheduleId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.AllocationExposureScheduleId
}

// GetAllocationExposureScheduleIdOk returns a tuple with the AllocationExposureScheduleId field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureRolloutStep) GetAllocationExposureScheduleIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocationExposureScheduleId, true
}

// SetAllocationExposureScheduleId sets field value.
func (o *AllocationExposureRolloutStep) SetAllocationExposureScheduleId(v uuid.UUID) {
	o.AllocationExposureScheduleId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AllocationExposureRolloutStep) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureRolloutStep) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AllocationExposureRolloutStep) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExposureRatio returns the ExposureRatio field value.
func (o *AllocationExposureRolloutStep) GetExposureRatio() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ExposureRatio
}

// GetExposureRatioOk returns a tuple with the ExposureRatio field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureRolloutStep) GetExposureRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExposureRatio, true
}

// SetExposureRatio sets field value.
func (o *AllocationExposureRolloutStep) SetExposureRatio(v float64) {
	o.ExposureRatio = v
}

// GetGroupedStepIndex returns the GroupedStepIndex field value.
func (o *AllocationExposureRolloutStep) GetGroupedStepIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.GroupedStepIndex
}

// GetGroupedStepIndexOk returns a tuple with the GroupedStepIndex field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureRolloutStep) GetGroupedStepIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupedStepIndex, true
}

// SetGroupedStepIndex sets field value.
func (o *AllocationExposureRolloutStep) SetGroupedStepIndex(v int64) {
	o.GroupedStepIndex = v
}

// GetId returns the Id field value.
func (o *AllocationExposureRolloutStep) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureRolloutStep) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AllocationExposureRolloutStep) SetId(v uuid.UUID) {
	o.Id = v
}

// GetIntervalMs returns the IntervalMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllocationExposureRolloutStep) GetIntervalMs() int64 {
	if o == nil || o.IntervalMs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IntervalMs.Get()
}

// GetIntervalMsOk returns a tuple with the IntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AllocationExposureRolloutStep) GetIntervalMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalMs.Get(), o.IntervalMs.IsSet()
}

// HasIntervalMs returns a boolean if a field has been set.
func (o *AllocationExposureRolloutStep) HasIntervalMs() bool {
	return o != nil && o.IntervalMs.IsSet()
}

// SetIntervalMs gets a reference to the given datadog.NullableInt64 and assigns it to the IntervalMs field.
func (o *AllocationExposureRolloutStep) SetIntervalMs(v int64) {
	o.IntervalMs.Set(&v)
}

// SetIntervalMsNil sets the value for IntervalMs to be an explicit nil.
func (o *AllocationExposureRolloutStep) SetIntervalMsNil() {
	o.IntervalMs.Set(nil)
}

// UnsetIntervalMs ensures that no value is present for IntervalMs, not even an explicit nil.
func (o *AllocationExposureRolloutStep) UnsetIntervalMs() {
	o.IntervalMs.Unset()
}

// GetIsPauseRecord returns the IsPauseRecord field value.
func (o *AllocationExposureRolloutStep) GetIsPauseRecord() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPauseRecord
}

// GetIsPauseRecordOk returns a tuple with the IsPauseRecord field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureRolloutStep) GetIsPauseRecordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPauseRecord, true
}

// SetIsPauseRecord sets field value.
func (o *AllocationExposureRolloutStep) SetIsPauseRecord(v bool) {
	o.IsPauseRecord = v
}

// GetOrderPosition returns the OrderPosition field value.
func (o *AllocationExposureRolloutStep) GetOrderPosition() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrderPosition
}

// GetOrderPositionOk returns a tuple with the OrderPosition field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureRolloutStep) GetOrderPositionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderPosition, true
}

// SetOrderPosition sets field value.
func (o *AllocationExposureRolloutStep) SetOrderPosition(v int64) {
	o.OrderPosition = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AllocationExposureRolloutStep) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureRolloutStep) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AllocationExposureRolloutStep) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AllocationExposureRolloutStep) MarshalJSON() ([]byte, error) {
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
	toSerialize["exposure_ratio"] = o.ExposureRatio
	toSerialize["grouped_step_index"] = o.GroupedStepIndex
	toSerialize["id"] = o.Id
	if o.IntervalMs.IsSet() {
		toSerialize["interval_ms"] = o.IntervalMs.Get()
	}
	toSerialize["is_pause_record"] = o.IsPauseRecord
	toSerialize["order_position"] = o.OrderPosition
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
func (o *AllocationExposureRolloutStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllocationExposureScheduleId *uuid.UUID            `json:"allocation_exposure_schedule_id"`
		CreatedAt                    *time.Time            `json:"created_at"`
		ExposureRatio                *float64              `json:"exposure_ratio"`
		GroupedStepIndex             *int64                `json:"grouped_step_index"`
		Id                           *uuid.UUID            `json:"id"`
		IntervalMs                   datadog.NullableInt64 `json:"interval_ms,omitempty"`
		IsPauseRecord                *bool                 `json:"is_pause_record"`
		OrderPosition                *int64                `json:"order_position"`
		UpdatedAt                    *time.Time            `json:"updated_at"`
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
	if all.ExposureRatio == nil {
		return fmt.Errorf("required field exposure_ratio missing")
	}
	if all.GroupedStepIndex == nil {
		return fmt.Errorf("required field grouped_step_index missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.IsPauseRecord == nil {
		return fmt.Errorf("required field is_pause_record missing")
	}
	if all.OrderPosition == nil {
		return fmt.Errorf("required field order_position missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allocation_exposure_schedule_id", "created_at", "exposure_ratio", "grouped_step_index", "id", "interval_ms", "is_pause_record", "order_position", "updated_at"})
	} else {
		return err
	}
	o.AllocationExposureScheduleId = *all.AllocationExposureScheduleId
	o.CreatedAt = *all.CreatedAt
	o.ExposureRatio = *all.ExposureRatio
	o.GroupedStepIndex = *all.GroupedStepIndex
	o.Id = *all.Id
	o.IntervalMs = all.IntervalMs
	o.IsPauseRecord = *all.IsPauseRecord
	o.OrderPosition = *all.OrderPosition
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
