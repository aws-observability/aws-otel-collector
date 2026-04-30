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

// ExposureScheduleRequest Progressive release request payload.
type ExposureScheduleRequest struct {
	// The absolute UTC start time for this schedule.
	AbsoluteStartTime datadog.NullableTime `json:"absolute_start_time,omitempty"`
	// The control variant ID used for experiment comparisons.
	ControlVariantId datadog.NullableString `json:"control_variant_id,omitempty"`
	// The control variant key used during creation workflows.
	ControlVariantKey datadog.NullableString `json:"control_variant_key,omitempty"`
	// The unique identifier of the progressive rollout.
	Id *uuid.UUID `json:"id,omitempty"`
	// Rollout options request payload.
	RolloutOptions RolloutOptionsRequest `json:"rollout_options"`
	// Ordered progression steps for exposure.
	RolloutSteps []ExposureRolloutStepRequest `json:"rollout_steps"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExposureScheduleRequest instantiates a new ExposureScheduleRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExposureScheduleRequest(rolloutOptions RolloutOptionsRequest, rolloutSteps []ExposureRolloutStepRequest) *ExposureScheduleRequest {
	this := ExposureScheduleRequest{}
	this.RolloutOptions = rolloutOptions
	this.RolloutSteps = rolloutSteps
	return &this
}

// NewExposureScheduleRequestWithDefaults instantiates a new ExposureScheduleRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExposureScheduleRequestWithDefaults() *ExposureScheduleRequest {
	this := ExposureScheduleRequest{}
	return &this
}

// GetAbsoluteStartTime returns the AbsoluteStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExposureScheduleRequest) GetAbsoluteStartTime() time.Time {
	if o == nil || o.AbsoluteStartTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AbsoluteStartTime.Get()
}

// GetAbsoluteStartTimeOk returns a tuple with the AbsoluteStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExposureScheduleRequest) GetAbsoluteStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbsoluteStartTime.Get(), o.AbsoluteStartTime.IsSet()
}

// HasAbsoluteStartTime returns a boolean if a field has been set.
func (o *ExposureScheduleRequest) HasAbsoluteStartTime() bool {
	return o != nil && o.AbsoluteStartTime.IsSet()
}

// SetAbsoluteStartTime gets a reference to the given datadog.NullableTime and assigns it to the AbsoluteStartTime field.
func (o *ExposureScheduleRequest) SetAbsoluteStartTime(v time.Time) {
	o.AbsoluteStartTime.Set(&v)
}

// SetAbsoluteStartTimeNil sets the value for AbsoluteStartTime to be an explicit nil.
func (o *ExposureScheduleRequest) SetAbsoluteStartTimeNil() {
	o.AbsoluteStartTime.Set(nil)
}

// UnsetAbsoluteStartTime ensures that no value is present for AbsoluteStartTime, not even an explicit nil.
func (o *ExposureScheduleRequest) UnsetAbsoluteStartTime() {
	o.AbsoluteStartTime.Unset()
}

// GetControlVariantId returns the ControlVariantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExposureScheduleRequest) GetControlVariantId() string {
	if o == nil || o.ControlVariantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ControlVariantId.Get()
}

// GetControlVariantIdOk returns a tuple with the ControlVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExposureScheduleRequest) GetControlVariantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ControlVariantId.Get(), o.ControlVariantId.IsSet()
}

// HasControlVariantId returns a boolean if a field has been set.
func (o *ExposureScheduleRequest) HasControlVariantId() bool {
	return o != nil && o.ControlVariantId.IsSet()
}

// SetControlVariantId gets a reference to the given datadog.NullableString and assigns it to the ControlVariantId field.
func (o *ExposureScheduleRequest) SetControlVariantId(v string) {
	o.ControlVariantId.Set(&v)
}

// SetControlVariantIdNil sets the value for ControlVariantId to be an explicit nil.
func (o *ExposureScheduleRequest) SetControlVariantIdNil() {
	o.ControlVariantId.Set(nil)
}

// UnsetControlVariantId ensures that no value is present for ControlVariantId, not even an explicit nil.
func (o *ExposureScheduleRequest) UnsetControlVariantId() {
	o.ControlVariantId.Unset()
}

// GetControlVariantKey returns the ControlVariantKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExposureScheduleRequest) GetControlVariantKey() string {
	if o == nil || o.ControlVariantKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.ControlVariantKey.Get()
}

// GetControlVariantKeyOk returns a tuple with the ControlVariantKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExposureScheduleRequest) GetControlVariantKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ControlVariantKey.Get(), o.ControlVariantKey.IsSet()
}

// HasControlVariantKey returns a boolean if a field has been set.
func (o *ExposureScheduleRequest) HasControlVariantKey() bool {
	return o != nil && o.ControlVariantKey.IsSet()
}

// SetControlVariantKey gets a reference to the given datadog.NullableString and assigns it to the ControlVariantKey field.
func (o *ExposureScheduleRequest) SetControlVariantKey(v string) {
	o.ControlVariantKey.Set(&v)
}

// SetControlVariantKeyNil sets the value for ControlVariantKey to be an explicit nil.
func (o *ExposureScheduleRequest) SetControlVariantKeyNil() {
	o.ControlVariantKey.Set(nil)
}

// UnsetControlVariantKey ensures that no value is present for ControlVariantKey, not even an explicit nil.
func (o *ExposureScheduleRequest) UnsetControlVariantKey() {
	o.ControlVariantKey.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExposureScheduleRequest) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExposureScheduleRequest) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExposureScheduleRequest) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *ExposureScheduleRequest) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetRolloutOptions returns the RolloutOptions field value.
func (o *ExposureScheduleRequest) GetRolloutOptions() RolloutOptionsRequest {
	if o == nil {
		var ret RolloutOptionsRequest
		return ret
	}
	return o.RolloutOptions
}

// GetRolloutOptionsOk returns a tuple with the RolloutOptions field value
// and a boolean to check if the value has been set.
func (o *ExposureScheduleRequest) GetRolloutOptionsOk() (*RolloutOptionsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RolloutOptions, true
}

// SetRolloutOptions sets field value.
func (o *ExposureScheduleRequest) SetRolloutOptions(v RolloutOptionsRequest) {
	o.RolloutOptions = v
}

// GetRolloutSteps returns the RolloutSteps field value.
func (o *ExposureScheduleRequest) GetRolloutSteps() []ExposureRolloutStepRequest {
	if o == nil {
		var ret []ExposureRolloutStepRequest
		return ret
	}
	return o.RolloutSteps
}

// GetRolloutStepsOk returns a tuple with the RolloutSteps field value
// and a boolean to check if the value has been set.
func (o *ExposureScheduleRequest) GetRolloutStepsOk() (*[]ExposureRolloutStepRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RolloutSteps, true
}

// SetRolloutSteps sets field value.
func (o *ExposureScheduleRequest) SetRolloutSteps(v []ExposureRolloutStepRequest) {
	o.RolloutSteps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExposureScheduleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AbsoluteStartTime.IsSet() {
		toSerialize["absolute_start_time"] = o.AbsoluteStartTime.Get()
	}
	if o.ControlVariantId.IsSet() {
		toSerialize["control_variant_id"] = o.ControlVariantId.Get()
	}
	if o.ControlVariantKey.IsSet() {
		toSerialize["control_variant_key"] = o.ControlVariantKey.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["rollout_options"] = o.RolloutOptions
	toSerialize["rollout_steps"] = o.RolloutSteps

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExposureScheduleRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AbsoluteStartTime datadog.NullableTime          `json:"absolute_start_time,omitempty"`
		ControlVariantId  datadog.NullableString        `json:"control_variant_id,omitempty"`
		ControlVariantKey datadog.NullableString        `json:"control_variant_key,omitempty"`
		Id                *uuid.UUID                    `json:"id,omitempty"`
		RolloutOptions    *RolloutOptionsRequest        `json:"rollout_options"`
		RolloutSteps      *[]ExposureRolloutStepRequest `json:"rollout_steps"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RolloutOptions == nil {
		return fmt.Errorf("required field rollout_options missing")
	}
	if all.RolloutSteps == nil {
		return fmt.Errorf("required field rollout_steps missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"absolute_start_time", "control_variant_id", "control_variant_key", "id", "rollout_options", "rollout_steps"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AbsoluteStartTime = all.AbsoluteStartTime
	o.ControlVariantId = all.ControlVariantId
	o.ControlVariantKey = all.ControlVariantKey
	o.Id = all.Id
	if all.RolloutOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RolloutOptions = *all.RolloutOptions
	o.RolloutSteps = *all.RolloutSteps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
