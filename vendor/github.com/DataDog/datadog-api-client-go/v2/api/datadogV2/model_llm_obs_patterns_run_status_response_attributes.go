// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsRunStatusResponseAttributes Attributes of an LLM Observability patterns run status.
type LLMObsPatternsRunStatusResponseAttributes struct {
	// Timestamp when the run was created.
	CreatedAt time.Time `json:"created_at"`
	// List of step-by-step progress entries for a patterns run.
	Progress []LLMObsPatternsActivityProgress `json:"progress"`
	// Overall status of the run.
	Status string `json:"status"`
	// The current step of the run.
	Step string `json:"step"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsRunStatusResponseAttributes instantiates a new LLMObsPatternsRunStatusResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsRunStatusResponseAttributes(createdAt time.Time, progress []LLMObsPatternsActivityProgress, status string, step string) *LLMObsPatternsRunStatusResponseAttributes {
	this := LLMObsPatternsRunStatusResponseAttributes{}
	this.CreatedAt = createdAt
	this.Progress = progress
	this.Status = status
	this.Step = step
	return &this
}

// NewLLMObsPatternsRunStatusResponseAttributesWithDefaults instantiates a new LLMObsPatternsRunStatusResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsRunStatusResponseAttributesWithDefaults() *LLMObsPatternsRunStatusResponseAttributes {
	this := LLMObsPatternsRunStatusResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsPatternsRunStatusResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsRunStatusResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsPatternsRunStatusResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetProgress returns the Progress field value.
func (o *LLMObsPatternsRunStatusResponseAttributes) GetProgress() []LLMObsPatternsActivityProgress {
	if o == nil {
		var ret []LLMObsPatternsActivityProgress
		return ret
	}
	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsRunStatusResponseAttributes) GetProgressOk() (*[]LLMObsPatternsActivityProgress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value.
func (o *LLMObsPatternsRunStatusResponseAttributes) SetProgress(v []LLMObsPatternsActivityProgress) {
	o.Progress = v
}

// GetStatus returns the Status field value.
func (o *LLMObsPatternsRunStatusResponseAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsRunStatusResponseAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *LLMObsPatternsRunStatusResponseAttributes) SetStatus(v string) {
	o.Status = v
}

// GetStep returns the Step field value.
func (o *LLMObsPatternsRunStatusResponseAttributes) GetStep() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Step
}

// GetStepOk returns a tuple with the Step field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsRunStatusResponseAttributes) GetStepOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Step, true
}

// SetStep sets field value.
func (o *LLMObsPatternsRunStatusResponseAttributes) SetStep(v string) {
	o.Step = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsRunStatusResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["progress"] = o.Progress
	toSerialize["status"] = o.Status
	toSerialize["step"] = o.Step

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsRunStatusResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time                        `json:"created_at"`
		Progress  *[]LLMObsPatternsActivityProgress `json:"progress"`
		Status    *string                           `json:"status"`
		Step      *string                           `json:"step"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Progress == nil {
		return fmt.Errorf("required field progress missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Step == nil {
		return fmt.Errorf("required field step missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "progress", "status", "step"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.Progress = *all.Progress
	o.Status = *all.Status
	o.Step = *all.Step

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
