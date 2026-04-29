// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExposureRolloutStepRequest Rollout step request payload.
type ExposureRolloutStepRequest struct {
	// The exposure ratio for this step.
	ExposureRatio float64 `json:"exposure_ratio"`
	// Logical index grouping related steps.
	GroupedStepIndex int64 `json:"grouped_step_index"`
	// The unique identifier of the progression step.
	Id *uuid.UUID `json:"id,omitempty"`
	// Step duration in milliseconds.
	IntervalMs datadog.NullableInt64 `json:"interval_ms,omitempty"`
	// Whether this step represents a pause record.
	IsPauseRecord bool `json:"is_pause_record"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExposureRolloutStepRequest instantiates a new ExposureRolloutStepRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExposureRolloutStepRequest(exposureRatio float64, groupedStepIndex int64, isPauseRecord bool) *ExposureRolloutStepRequest {
	this := ExposureRolloutStepRequest{}
	this.ExposureRatio = exposureRatio
	this.GroupedStepIndex = groupedStepIndex
	this.IsPauseRecord = isPauseRecord
	return &this
}

// NewExposureRolloutStepRequestWithDefaults instantiates a new ExposureRolloutStepRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExposureRolloutStepRequestWithDefaults() *ExposureRolloutStepRequest {
	this := ExposureRolloutStepRequest{}
	return &this
}

// GetExposureRatio returns the ExposureRatio field value.
func (o *ExposureRolloutStepRequest) GetExposureRatio() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ExposureRatio
}

// GetExposureRatioOk returns a tuple with the ExposureRatio field value
// and a boolean to check if the value has been set.
func (o *ExposureRolloutStepRequest) GetExposureRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExposureRatio, true
}

// SetExposureRatio sets field value.
func (o *ExposureRolloutStepRequest) SetExposureRatio(v float64) {
	o.ExposureRatio = v
}

// GetGroupedStepIndex returns the GroupedStepIndex field value.
func (o *ExposureRolloutStepRequest) GetGroupedStepIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.GroupedStepIndex
}

// GetGroupedStepIndexOk returns a tuple with the GroupedStepIndex field value
// and a boolean to check if the value has been set.
func (o *ExposureRolloutStepRequest) GetGroupedStepIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupedStepIndex, true
}

// SetGroupedStepIndex sets field value.
func (o *ExposureRolloutStepRequest) SetGroupedStepIndex(v int64) {
	o.GroupedStepIndex = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExposureRolloutStepRequest) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExposureRolloutStepRequest) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExposureRolloutStepRequest) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *ExposureRolloutStepRequest) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetIntervalMs returns the IntervalMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExposureRolloutStepRequest) GetIntervalMs() int64 {
	if o == nil || o.IntervalMs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IntervalMs.Get()
}

// GetIntervalMsOk returns a tuple with the IntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExposureRolloutStepRequest) GetIntervalMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalMs.Get(), o.IntervalMs.IsSet()
}

// HasIntervalMs returns a boolean if a field has been set.
func (o *ExposureRolloutStepRequest) HasIntervalMs() bool {
	return o != nil && o.IntervalMs.IsSet()
}

// SetIntervalMs gets a reference to the given datadog.NullableInt64 and assigns it to the IntervalMs field.
func (o *ExposureRolloutStepRequest) SetIntervalMs(v int64) {
	o.IntervalMs.Set(&v)
}

// SetIntervalMsNil sets the value for IntervalMs to be an explicit nil.
func (o *ExposureRolloutStepRequest) SetIntervalMsNil() {
	o.IntervalMs.Set(nil)
}

// UnsetIntervalMs ensures that no value is present for IntervalMs, not even an explicit nil.
func (o *ExposureRolloutStepRequest) UnsetIntervalMs() {
	o.IntervalMs.Unset()
}

// GetIsPauseRecord returns the IsPauseRecord field value.
func (o *ExposureRolloutStepRequest) GetIsPauseRecord() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPauseRecord
}

// GetIsPauseRecordOk returns a tuple with the IsPauseRecord field value
// and a boolean to check if the value has been set.
func (o *ExposureRolloutStepRequest) GetIsPauseRecordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPauseRecord, true
}

// SetIsPauseRecord sets field value.
func (o *ExposureRolloutStepRequest) SetIsPauseRecord(v bool) {
	o.IsPauseRecord = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExposureRolloutStepRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["exposure_ratio"] = o.ExposureRatio
	toSerialize["grouped_step_index"] = o.GroupedStepIndex
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IntervalMs.IsSet() {
		toSerialize["interval_ms"] = o.IntervalMs.Get()
	}
	toSerialize["is_pause_record"] = o.IsPauseRecord

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExposureRolloutStepRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExposureRatio    *float64              `json:"exposure_ratio"`
		GroupedStepIndex *int64                `json:"grouped_step_index"`
		Id               *uuid.UUID            `json:"id,omitempty"`
		IntervalMs       datadog.NullableInt64 `json:"interval_ms,omitempty"`
		IsPauseRecord    *bool                 `json:"is_pause_record"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExposureRatio == nil {
		return fmt.Errorf("required field exposure_ratio missing")
	}
	if all.GroupedStepIndex == nil {
		return fmt.Errorf("required field grouped_step_index missing")
	}
	if all.IsPauseRecord == nil {
		return fmt.Errorf("required field is_pause_record missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exposure_ratio", "grouped_step_index", "id", "interval_ms", "is_pause_record"})
	} else {
		return err
	}
	o.ExposureRatio = *all.ExposureRatio
	o.GroupedStepIndex = *all.GroupedStepIndex
	o.Id = all.Id
	o.IntervalMs = all.IntervalMs
	o.IsPauseRecord = *all.IsPauseRecord

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
