// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsRunSummary Summary of an LLM Observability patterns run.
type LLMObsPatternsRunSummary struct {
	// Timestamp when the run completed. Null if the run has not completed.
	CompletedAt datadog.NullableTime `json:"completed_at,omitempty"`
	// Snapshot of the configuration used for a patterns run.
	ConfigSnapshot *LLMObsPatternsConfigSnapshot `json:"config_snapshot,omitempty"`
	// Timestamp when the run was created.
	CreatedAt time.Time `json:"created_at"`
	// Unique identifier of the run.
	Id string `json:"id"`
	// Status of the run.
	Status string `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsRunSummary instantiates a new LLMObsPatternsRunSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsRunSummary(createdAt time.Time, id string, status string) *LLMObsPatternsRunSummary {
	this := LLMObsPatternsRunSummary{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Status = status
	return &this
}

// NewLLMObsPatternsRunSummaryWithDefaults instantiates a new LLMObsPatternsRunSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsRunSummaryWithDefaults() *LLMObsPatternsRunSummary {
	this := LLMObsPatternsRunSummary{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsPatternsRunSummary) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsPatternsRunSummary) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *LLMObsPatternsRunSummary) HasCompletedAt() bool {
	return o != nil && o.CompletedAt.IsSet()
}

// SetCompletedAt gets a reference to the given datadog.NullableTime and assigns it to the CompletedAt field.
func (o *LLMObsPatternsRunSummary) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}

// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil.
func (o *LLMObsPatternsRunSummary) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil.
func (o *LLMObsPatternsRunSummary) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetConfigSnapshot returns the ConfigSnapshot field value if set, zero value otherwise.
func (o *LLMObsPatternsRunSummary) GetConfigSnapshot() LLMObsPatternsConfigSnapshot {
	if o == nil || o.ConfigSnapshot == nil {
		var ret LLMObsPatternsConfigSnapshot
		return ret
	}
	return *o.ConfigSnapshot
}

// GetConfigSnapshotOk returns a tuple with the ConfigSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsRunSummary) GetConfigSnapshotOk() (*LLMObsPatternsConfigSnapshot, bool) {
	if o == nil || o.ConfigSnapshot == nil {
		return nil, false
	}
	return o.ConfigSnapshot, true
}

// HasConfigSnapshot returns a boolean if a field has been set.
func (o *LLMObsPatternsRunSummary) HasConfigSnapshot() bool {
	return o != nil && o.ConfigSnapshot != nil
}

// SetConfigSnapshot gets a reference to the given LLMObsPatternsConfigSnapshot and assigns it to the ConfigSnapshot field.
func (o *LLMObsPatternsRunSummary) SetConfigSnapshot(v LLMObsPatternsConfigSnapshot) {
	o.ConfigSnapshot = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsPatternsRunSummary) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsRunSummary) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsPatternsRunSummary) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value.
func (o *LLMObsPatternsRunSummary) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsRunSummary) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsPatternsRunSummary) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value.
func (o *LLMObsPatternsRunSummary) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsRunSummary) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *LLMObsPatternsRunSummary) SetStatus(v string) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsRunSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if o.ConfigSnapshot != nil {
		toSerialize["config_snapshot"] = o.ConfigSnapshot
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsRunSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompletedAt    datadog.NullableTime          `json:"completed_at,omitempty"`
		ConfigSnapshot *LLMObsPatternsConfigSnapshot `json:"config_snapshot,omitempty"`
		CreatedAt      *time.Time                    `json:"created_at"`
		Id             *string                       `json:"id"`
		Status         *string                       `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completed_at", "config_snapshot", "created_at", "id", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CompletedAt = all.CompletedAt
	if all.ConfigSnapshot != nil && all.ConfigSnapshot.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConfigSnapshot = all.ConfigSnapshot
	o.CreatedAt = *all.CreatedAt
	o.Id = *all.Id
	o.Status = *all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
