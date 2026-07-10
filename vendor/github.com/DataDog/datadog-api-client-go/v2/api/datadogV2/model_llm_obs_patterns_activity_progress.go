// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsActivityProgress Progress information for a single step of a patterns run.
type LLMObsPatternsActivityProgress struct {
	// Name of the step.
	Name string `json:"name"`
	// Timestamp when the step started. Null if the step has not started.
	StartedAt datadog.NullableTime `json:"started_at,omitempty"`
	// Status of the step.
	Status string `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsActivityProgress instantiates a new LLMObsPatternsActivityProgress object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsActivityProgress(name string, status string) *LLMObsPatternsActivityProgress {
	this := LLMObsPatternsActivityProgress{}
	this.Name = name
	this.Status = status
	return &this
}

// NewLLMObsPatternsActivityProgressWithDefaults instantiates a new LLMObsPatternsActivityProgress object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsActivityProgressWithDefaults() *LLMObsPatternsActivityProgress {
	this := LLMObsPatternsActivityProgress{}
	return &this
}

// GetName returns the Name field value.
func (o *LLMObsPatternsActivityProgress) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsActivityProgress) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsPatternsActivityProgress) SetName(v string) {
	o.Name = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsPatternsActivityProgress) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsPatternsActivityProgress) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *LLMObsPatternsActivityProgress) HasStartedAt() bool {
	return o != nil && o.StartedAt.IsSet()
}

// SetStartedAt gets a reference to the given datadog.NullableTime and assigns it to the StartedAt field.
func (o *LLMObsPatternsActivityProgress) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// SetStartedAtNil sets the value for StartedAt to be an explicit nil.
func (o *LLMObsPatternsActivityProgress) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil.
func (o *LLMObsPatternsActivityProgress) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetStatus returns the Status field value.
func (o *LLMObsPatternsActivityProgress) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsActivityProgress) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *LLMObsPatternsActivityProgress) SetStatus(v string) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsActivityProgress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.StartedAt.IsSet() {
		toSerialize["started_at"] = o.StartedAt.Get()
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsActivityProgress) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string              `json:"name"`
		StartedAt datadog.NullableTime `json:"started_at,omitempty"`
		Status    *string              `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "started_at", "status"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.StartedAt = all.StartedAt
	o.Status = *all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
