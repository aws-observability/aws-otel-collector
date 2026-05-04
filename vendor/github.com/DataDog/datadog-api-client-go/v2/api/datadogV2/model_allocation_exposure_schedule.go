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

// AllocationExposureSchedule Progressive release details for a targeting rule allocation.
type AllocationExposureSchedule struct {
	// The absolute UTC start time for this schedule.
	AbsoluteStartTime datadog.NullableTime `json:"absolute_start_time,omitempty"`
	// The targeting rule allocation ID this progressive rollout belongs to.
	AllocationId uuid.UUID `json:"allocation_id"`
	// The control variant ID used for experiment comparisons.
	ControlVariantId datadog.NullableString `json:"control_variant_id,omitempty"`
	// The timestamp when the schedule was created.
	CreatedAt time.Time `json:"created_at"`
	// Last guardrail action triggered for this schedule.
	GuardrailTriggeredAction datadog.NullableString `json:"guardrail_triggered_action,omitempty"`
	// Guardrail trigger records for this schedule.
	GuardrailTriggers []AllocationExposureGuardrailTrigger `json:"guardrail_triggers"`
	// The unique identifier of the progressive rollout.
	Id *uuid.UUID `json:"id,omitempty"`
	// Applied progression options for a progressive rollout.
	RolloutOptions RolloutOptions `json:"rollout_options"`
	// Ordered progression steps for exposure.
	RolloutSteps []AllocationExposureRolloutStep `json:"rollout_steps"`
	// The timestamp when the schedule was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAllocationExposureSchedule instantiates a new AllocationExposureSchedule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAllocationExposureSchedule(allocationId uuid.UUID, createdAt time.Time, guardrailTriggers []AllocationExposureGuardrailTrigger, rolloutOptions RolloutOptions, rolloutSteps []AllocationExposureRolloutStep, updatedAt time.Time) *AllocationExposureSchedule {
	this := AllocationExposureSchedule{}
	this.AllocationId = allocationId
	this.CreatedAt = createdAt
	this.GuardrailTriggers = guardrailTriggers
	this.RolloutOptions = rolloutOptions
	this.RolloutSteps = rolloutSteps
	this.UpdatedAt = updatedAt
	return &this
}

// NewAllocationExposureScheduleWithDefaults instantiates a new AllocationExposureSchedule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAllocationExposureScheduleWithDefaults() *AllocationExposureSchedule {
	this := AllocationExposureSchedule{}
	return &this
}

// GetAbsoluteStartTime returns the AbsoluteStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllocationExposureSchedule) GetAbsoluteStartTime() time.Time {
	if o == nil || o.AbsoluteStartTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AbsoluteStartTime.Get()
}

// GetAbsoluteStartTimeOk returns a tuple with the AbsoluteStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AllocationExposureSchedule) GetAbsoluteStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbsoluteStartTime.Get(), o.AbsoluteStartTime.IsSet()
}

// HasAbsoluteStartTime returns a boolean if a field has been set.
func (o *AllocationExposureSchedule) HasAbsoluteStartTime() bool {
	return o != nil && o.AbsoluteStartTime.IsSet()
}

// SetAbsoluteStartTime gets a reference to the given datadog.NullableTime and assigns it to the AbsoluteStartTime field.
func (o *AllocationExposureSchedule) SetAbsoluteStartTime(v time.Time) {
	o.AbsoluteStartTime.Set(&v)
}

// SetAbsoluteStartTimeNil sets the value for AbsoluteStartTime to be an explicit nil.
func (o *AllocationExposureSchedule) SetAbsoluteStartTimeNil() {
	o.AbsoluteStartTime.Set(nil)
}

// UnsetAbsoluteStartTime ensures that no value is present for AbsoluteStartTime, not even an explicit nil.
func (o *AllocationExposureSchedule) UnsetAbsoluteStartTime() {
	o.AbsoluteStartTime.Unset()
}

// GetAllocationId returns the AllocationId field value.
func (o *AllocationExposureSchedule) GetAllocationId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.AllocationId
}

// GetAllocationIdOk returns a tuple with the AllocationId field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureSchedule) GetAllocationIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocationId, true
}

// SetAllocationId sets field value.
func (o *AllocationExposureSchedule) SetAllocationId(v uuid.UUID) {
	o.AllocationId = v
}

// GetControlVariantId returns the ControlVariantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllocationExposureSchedule) GetControlVariantId() string {
	if o == nil || o.ControlVariantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ControlVariantId.Get()
}

// GetControlVariantIdOk returns a tuple with the ControlVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AllocationExposureSchedule) GetControlVariantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ControlVariantId.Get(), o.ControlVariantId.IsSet()
}

// HasControlVariantId returns a boolean if a field has been set.
func (o *AllocationExposureSchedule) HasControlVariantId() bool {
	return o != nil && o.ControlVariantId.IsSet()
}

// SetControlVariantId gets a reference to the given datadog.NullableString and assigns it to the ControlVariantId field.
func (o *AllocationExposureSchedule) SetControlVariantId(v string) {
	o.ControlVariantId.Set(&v)
}

// SetControlVariantIdNil sets the value for ControlVariantId to be an explicit nil.
func (o *AllocationExposureSchedule) SetControlVariantIdNil() {
	o.ControlVariantId.Set(nil)
}

// UnsetControlVariantId ensures that no value is present for ControlVariantId, not even an explicit nil.
func (o *AllocationExposureSchedule) UnsetControlVariantId() {
	o.ControlVariantId.Unset()
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AllocationExposureSchedule) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureSchedule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AllocationExposureSchedule) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetGuardrailTriggeredAction returns the GuardrailTriggeredAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllocationExposureSchedule) GetGuardrailTriggeredAction() string {
	if o == nil || o.GuardrailTriggeredAction.Get() == nil {
		var ret string
		return ret
	}
	return *o.GuardrailTriggeredAction.Get()
}

// GetGuardrailTriggeredActionOk returns a tuple with the GuardrailTriggeredAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AllocationExposureSchedule) GetGuardrailTriggeredActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuardrailTriggeredAction.Get(), o.GuardrailTriggeredAction.IsSet()
}

// HasGuardrailTriggeredAction returns a boolean if a field has been set.
func (o *AllocationExposureSchedule) HasGuardrailTriggeredAction() bool {
	return o != nil && o.GuardrailTriggeredAction.IsSet()
}

// SetGuardrailTriggeredAction gets a reference to the given datadog.NullableString and assigns it to the GuardrailTriggeredAction field.
func (o *AllocationExposureSchedule) SetGuardrailTriggeredAction(v string) {
	o.GuardrailTriggeredAction.Set(&v)
}

// SetGuardrailTriggeredActionNil sets the value for GuardrailTriggeredAction to be an explicit nil.
func (o *AllocationExposureSchedule) SetGuardrailTriggeredActionNil() {
	o.GuardrailTriggeredAction.Set(nil)
}

// UnsetGuardrailTriggeredAction ensures that no value is present for GuardrailTriggeredAction, not even an explicit nil.
func (o *AllocationExposureSchedule) UnsetGuardrailTriggeredAction() {
	o.GuardrailTriggeredAction.Unset()
}

// GetGuardrailTriggers returns the GuardrailTriggers field value.
func (o *AllocationExposureSchedule) GetGuardrailTriggers() []AllocationExposureGuardrailTrigger {
	if o == nil {
		var ret []AllocationExposureGuardrailTrigger
		return ret
	}
	return o.GuardrailTriggers
}

// GetGuardrailTriggersOk returns a tuple with the GuardrailTriggers field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureSchedule) GetGuardrailTriggersOk() (*[]AllocationExposureGuardrailTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuardrailTriggers, true
}

// SetGuardrailTriggers sets field value.
func (o *AllocationExposureSchedule) SetGuardrailTriggers(v []AllocationExposureGuardrailTrigger) {
	o.GuardrailTriggers = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AllocationExposureSchedule) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationExposureSchedule) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AllocationExposureSchedule) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *AllocationExposureSchedule) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetRolloutOptions returns the RolloutOptions field value.
func (o *AllocationExposureSchedule) GetRolloutOptions() RolloutOptions {
	if o == nil {
		var ret RolloutOptions
		return ret
	}
	return o.RolloutOptions
}

// GetRolloutOptionsOk returns a tuple with the RolloutOptions field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureSchedule) GetRolloutOptionsOk() (*RolloutOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RolloutOptions, true
}

// SetRolloutOptions sets field value.
func (o *AllocationExposureSchedule) SetRolloutOptions(v RolloutOptions) {
	o.RolloutOptions = v
}

// GetRolloutSteps returns the RolloutSteps field value.
func (o *AllocationExposureSchedule) GetRolloutSteps() []AllocationExposureRolloutStep {
	if o == nil {
		var ret []AllocationExposureRolloutStep
		return ret
	}
	return o.RolloutSteps
}

// GetRolloutStepsOk returns a tuple with the RolloutSteps field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureSchedule) GetRolloutStepsOk() (*[]AllocationExposureRolloutStep, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RolloutSteps, true
}

// SetRolloutSteps sets field value.
func (o *AllocationExposureSchedule) SetRolloutSteps(v []AllocationExposureRolloutStep) {
	o.RolloutSteps = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AllocationExposureSchedule) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AllocationExposureSchedule) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AllocationExposureSchedule) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AllocationExposureSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AbsoluteStartTime.IsSet() {
		toSerialize["absolute_start_time"] = o.AbsoluteStartTime.Get()
	}
	toSerialize["allocation_id"] = o.AllocationId
	if o.ControlVariantId.IsSet() {
		toSerialize["control_variant_id"] = o.ControlVariantId.Get()
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.GuardrailTriggeredAction.IsSet() {
		toSerialize["guardrail_triggered_action"] = o.GuardrailTriggeredAction.Get()
	}
	toSerialize["guardrail_triggers"] = o.GuardrailTriggers
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["rollout_options"] = o.RolloutOptions
	toSerialize["rollout_steps"] = o.RolloutSteps
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
func (o *AllocationExposureSchedule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AbsoluteStartTime        datadog.NullableTime                  `json:"absolute_start_time,omitempty"`
		AllocationId             *uuid.UUID                            `json:"allocation_id"`
		ControlVariantId         datadog.NullableString                `json:"control_variant_id,omitempty"`
		CreatedAt                *time.Time                            `json:"created_at"`
		GuardrailTriggeredAction datadog.NullableString                `json:"guardrail_triggered_action,omitempty"`
		GuardrailTriggers        *[]AllocationExposureGuardrailTrigger `json:"guardrail_triggers"`
		Id                       *uuid.UUID                            `json:"id,omitempty"`
		RolloutOptions           *RolloutOptions                       `json:"rollout_options"`
		RolloutSteps             *[]AllocationExposureRolloutStep      `json:"rollout_steps"`
		UpdatedAt                *time.Time                            `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AllocationId == nil {
		return fmt.Errorf("required field allocation_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.GuardrailTriggers == nil {
		return fmt.Errorf("required field guardrail_triggers missing")
	}
	if all.RolloutOptions == nil {
		return fmt.Errorf("required field rollout_options missing")
	}
	if all.RolloutSteps == nil {
		return fmt.Errorf("required field rollout_steps missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"absolute_start_time", "allocation_id", "control_variant_id", "created_at", "guardrail_triggered_action", "guardrail_triggers", "id", "rollout_options", "rollout_steps", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AbsoluteStartTime = all.AbsoluteStartTime
	o.AllocationId = *all.AllocationId
	o.ControlVariantId = all.ControlVariantId
	o.CreatedAt = *all.CreatedAt
	o.GuardrailTriggeredAction = all.GuardrailTriggeredAction
	o.GuardrailTriggers = *all.GuardrailTriggers
	o.Id = all.Id
	if all.RolloutOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RolloutOptions = *all.RolloutOptions
	o.RolloutSteps = *all.RolloutSteps
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
