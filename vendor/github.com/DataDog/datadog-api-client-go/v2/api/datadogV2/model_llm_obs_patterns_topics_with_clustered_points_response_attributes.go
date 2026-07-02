// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsTopicsWithClusteredPointsResponseAttributes Attributes of an LLM Observability patterns topics-with-clustered-points response.
type LLMObsPatternsTopicsWithClusteredPointsResponseAttributes struct {
	// Timestamp when the run completed. Null if the run has not completed.
	CompletedAt datadog.NullableTime `json:"completed_at,omitempty"`
	// Identifier of the configuration that produced the run.
	ConfigId string `json:"config_id"`
	// Snapshot of the configuration used for a patterns run.
	ConfigSnapshot *LLMObsPatternsConfigSnapshot `json:"config_snapshot,omitempty"`
	// Timestamp when the run was created.
	CreatedAt time.Time `json:"created_at"`
	// Identifier of the run that completed immediately before this one. Empty if none.
	PreviousRunId string `json:"previous_run_id"`
	// Identifier of the run that produced the topics.
	RunId string `json:"run_id"`
	// List of discovered topics with their clustered points.
	Topics []LLMObsPatternsTopicWithClusteredPoints `json:"topics"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsTopicsWithClusteredPointsResponseAttributes instantiates a new LLMObsPatternsTopicsWithClusteredPointsResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsTopicsWithClusteredPointsResponseAttributes(configId string, createdAt time.Time, previousRunId string, runId string, topics []LLMObsPatternsTopicWithClusteredPoints) *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes {
	this := LLMObsPatternsTopicsWithClusteredPointsResponseAttributes{}
	this.ConfigId = configId
	this.CreatedAt = createdAt
	this.PreviousRunId = previousRunId
	this.RunId = runId
	this.Topics = topics
	return &this
}

// NewLLMObsPatternsTopicsWithClusteredPointsResponseAttributesWithDefaults instantiates a new LLMObsPatternsTopicsWithClusteredPointsResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsTopicsWithClusteredPointsResponseAttributesWithDefaults() *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes {
	this := LLMObsPatternsTopicsWithClusteredPointsResponseAttributes{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) HasCompletedAt() bool {
	return o != nil && o.CompletedAt.IsSet()
}

// SetCompletedAt gets a reference to the given datadog.NullableTime and assigns it to the CompletedAt field.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}

// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetConfigId returns the ConfigId field value.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetConfigId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigId, true
}

// SetConfigId sets field value.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) SetConfigId(v string) {
	o.ConfigId = v
}

// GetConfigSnapshot returns the ConfigSnapshot field value if set, zero value otherwise.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetConfigSnapshot() LLMObsPatternsConfigSnapshot {
	if o == nil || o.ConfigSnapshot == nil {
		var ret LLMObsPatternsConfigSnapshot
		return ret
	}
	return *o.ConfigSnapshot
}

// GetConfigSnapshotOk returns a tuple with the ConfigSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetConfigSnapshotOk() (*LLMObsPatternsConfigSnapshot, bool) {
	if o == nil || o.ConfigSnapshot == nil {
		return nil, false
	}
	return o.ConfigSnapshot, true
}

// HasConfigSnapshot returns a boolean if a field has been set.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) HasConfigSnapshot() bool {
	return o != nil && o.ConfigSnapshot != nil
}

// SetConfigSnapshot gets a reference to the given LLMObsPatternsConfigSnapshot and assigns it to the ConfigSnapshot field.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) SetConfigSnapshot(v LLMObsPatternsConfigSnapshot) {
	o.ConfigSnapshot = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetPreviousRunId returns the PreviousRunId field value.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetPreviousRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PreviousRunId
}

// GetPreviousRunIdOk returns a tuple with the PreviousRunId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetPreviousRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousRunId, true
}

// SetPreviousRunId sets field value.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) SetPreviousRunId(v string) {
	o.PreviousRunId = v
}

// GetRunId returns the RunId field value.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) SetRunId(v string) {
	o.RunId = v
}

// GetTopics returns the Topics field value.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetTopics() []LLMObsPatternsTopicWithClusteredPoints {
	if o == nil {
		var ret []LLMObsPatternsTopicWithClusteredPoints
		return ret
	}
	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) GetTopicsOk() (*[]LLMObsPatternsTopicWithClusteredPoints, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topics, true
}

// SetTopics sets field value.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) SetTopics(v []LLMObsPatternsTopicWithClusteredPoints) {
	o.Topics = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	toSerialize["config_id"] = o.ConfigId
	if o.ConfigSnapshot != nil {
		toSerialize["config_snapshot"] = o.ConfigSnapshot
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["previous_run_id"] = o.PreviousRunId
	toSerialize["run_id"] = o.RunId
	toSerialize["topics"] = o.Topics

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsTopicsWithClusteredPointsResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompletedAt    datadog.NullableTime                      `json:"completed_at,omitempty"`
		ConfigId       *string                                   `json:"config_id"`
		ConfigSnapshot *LLMObsPatternsConfigSnapshot             `json:"config_snapshot,omitempty"`
		CreatedAt      *time.Time                                `json:"created_at"`
		PreviousRunId  *string                                   `json:"previous_run_id"`
		RunId          *string                                   `json:"run_id"`
		Topics         *[]LLMObsPatternsTopicWithClusteredPoints `json:"topics"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfigId == nil {
		return fmt.Errorf("required field config_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.PreviousRunId == nil {
		return fmt.Errorf("required field previous_run_id missing")
	}
	if all.RunId == nil {
		return fmt.Errorf("required field run_id missing")
	}
	if all.Topics == nil {
		return fmt.Errorf("required field topics missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completed_at", "config_id", "config_snapshot", "created_at", "previous_run_id", "run_id", "topics"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CompletedAt = all.CompletedAt
	o.ConfigId = *all.ConfigId
	if all.ConfigSnapshot != nil && all.ConfigSnapshot.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConfigSnapshot = all.ConfigSnapshot
	o.CreatedAt = *all.CreatedAt
	o.PreviousRunId = *all.PreviousRunId
	o.RunId = *all.RunId
	o.Topics = *all.Topics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
