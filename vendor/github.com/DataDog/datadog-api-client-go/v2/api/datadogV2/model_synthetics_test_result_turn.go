// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultTurn A turn in a goal-based browser test, grouping steps and reasoning.
type SyntheticsTestResultTurn struct {
	// Storage bucket keys for artifacts produced during a step or test.
	BucketKeys *SyntheticsTestResultBucketKeys `json:"bucket_keys,omitempty"`
	// Name of the turn.
	Name *string `json:"name,omitempty"`
	// Agent reasoning produced for this turn.
	Reasoning *string `json:"reasoning,omitempty"`
	// Status of the turn (for example, `passed`, `failed`).
	Status *string `json:"status,omitempty"`
	// Steps executed during the turn.
	Steps []SyntheticsTestResultTurnStep `json:"steps,omitempty"`
	// Unix timestamp (ms) of when the turn finished.
	TurnFinishedAt *int64 `json:"turn_finished_at,omitempty"`
	// Unix timestamp (ms) of when the turn started.
	TurnStartedAt *int64 `json:"turn_started_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultTurn instantiates a new SyntheticsTestResultTurn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultTurn() *SyntheticsTestResultTurn {
	this := SyntheticsTestResultTurn{}
	return &this
}

// NewSyntheticsTestResultTurnWithDefaults instantiates a new SyntheticsTestResultTurn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultTurnWithDefaults() *SyntheticsTestResultTurn {
	this := SyntheticsTestResultTurn{}
	return &this
}

// GetBucketKeys returns the BucketKeys field value if set, zero value otherwise.
func (o *SyntheticsTestResultTurn) GetBucketKeys() SyntheticsTestResultBucketKeys {
	if o == nil || o.BucketKeys == nil {
		var ret SyntheticsTestResultBucketKeys
		return ret
	}
	return *o.BucketKeys
}

// GetBucketKeysOk returns a tuple with the BucketKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTurn) GetBucketKeysOk() (*SyntheticsTestResultBucketKeys, bool) {
	if o == nil || o.BucketKeys == nil {
		return nil, false
	}
	return o.BucketKeys, true
}

// HasBucketKeys returns a boolean if a field has been set.
func (o *SyntheticsTestResultTurn) HasBucketKeys() bool {
	return o != nil && o.BucketKeys != nil
}

// SetBucketKeys gets a reference to the given SyntheticsTestResultBucketKeys and assigns it to the BucketKeys field.
func (o *SyntheticsTestResultTurn) SetBucketKeys(v SyntheticsTestResultBucketKeys) {
	o.BucketKeys = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestResultTurn) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTurn) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestResultTurn) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestResultTurn) SetName(v string) {
	o.Name = &v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *SyntheticsTestResultTurn) GetReasoning() string {
	if o == nil || o.Reasoning == nil {
		var ret string
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTurn) GetReasoningOk() (*string, bool) {
	if o == nil || o.Reasoning == nil {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *SyntheticsTestResultTurn) HasReasoning() bool {
	return o != nil && o.Reasoning != nil
}

// SetReasoning gets a reference to the given string and assigns it to the Reasoning field.
func (o *SyntheticsTestResultTurn) SetReasoning(v string) {
	o.Reasoning = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsTestResultTurn) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTurn) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsTestResultTurn) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyntheticsTestResultTurn) SetStatus(v string) {
	o.Status = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *SyntheticsTestResultTurn) GetSteps() []SyntheticsTestResultTurnStep {
	if o == nil || o.Steps == nil {
		var ret []SyntheticsTestResultTurnStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTurn) GetStepsOk() (*[]SyntheticsTestResultTurnStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *SyntheticsTestResultTurn) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []SyntheticsTestResultTurnStep and assigns it to the Steps field.
func (o *SyntheticsTestResultTurn) SetSteps(v []SyntheticsTestResultTurnStep) {
	o.Steps = v
}

// GetTurnFinishedAt returns the TurnFinishedAt field value if set, zero value otherwise.
func (o *SyntheticsTestResultTurn) GetTurnFinishedAt() int64 {
	if o == nil || o.TurnFinishedAt == nil {
		var ret int64
		return ret
	}
	return *o.TurnFinishedAt
}

// GetTurnFinishedAtOk returns a tuple with the TurnFinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTurn) GetTurnFinishedAtOk() (*int64, bool) {
	if o == nil || o.TurnFinishedAt == nil {
		return nil, false
	}
	return o.TurnFinishedAt, true
}

// HasTurnFinishedAt returns a boolean if a field has been set.
func (o *SyntheticsTestResultTurn) HasTurnFinishedAt() bool {
	return o != nil && o.TurnFinishedAt != nil
}

// SetTurnFinishedAt gets a reference to the given int64 and assigns it to the TurnFinishedAt field.
func (o *SyntheticsTestResultTurn) SetTurnFinishedAt(v int64) {
	o.TurnFinishedAt = &v
}

// GetTurnStartedAt returns the TurnStartedAt field value if set, zero value otherwise.
func (o *SyntheticsTestResultTurn) GetTurnStartedAt() int64 {
	if o == nil || o.TurnStartedAt == nil {
		var ret int64
		return ret
	}
	return *o.TurnStartedAt
}

// GetTurnStartedAtOk returns a tuple with the TurnStartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTurn) GetTurnStartedAtOk() (*int64, bool) {
	if o == nil || o.TurnStartedAt == nil {
		return nil, false
	}
	return o.TurnStartedAt, true
}

// HasTurnStartedAt returns a boolean if a field has been set.
func (o *SyntheticsTestResultTurn) HasTurnStartedAt() bool {
	return o != nil && o.TurnStartedAt != nil
}

// SetTurnStartedAt gets a reference to the given int64 and assigns it to the TurnStartedAt field.
func (o *SyntheticsTestResultTurn) SetTurnStartedAt(v int64) {
	o.TurnStartedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultTurn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BucketKeys != nil {
		toSerialize["bucket_keys"] = o.BucketKeys
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Reasoning != nil {
		toSerialize["reasoning"] = o.Reasoning
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.TurnFinishedAt != nil {
		toSerialize["turn_finished_at"] = o.TurnFinishedAt
	}
	if o.TurnStartedAt != nil {
		toSerialize["turn_started_at"] = o.TurnStartedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultTurn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketKeys     *SyntheticsTestResultBucketKeys `json:"bucket_keys,omitempty"`
		Name           *string                         `json:"name,omitempty"`
		Reasoning      *string                         `json:"reasoning,omitempty"`
		Status         *string                         `json:"status,omitempty"`
		Steps          []SyntheticsTestResultTurnStep  `json:"steps,omitempty"`
		TurnFinishedAt *int64                          `json:"turn_finished_at,omitempty"`
		TurnStartedAt  *int64                          `json:"turn_started_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucket_keys", "name", "reasoning", "status", "steps", "turn_finished_at", "turn_started_at"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.BucketKeys != nil && all.BucketKeys.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BucketKeys = all.BucketKeys
	o.Name = all.Name
	o.Reasoning = all.Reasoning
	o.Status = all.Status
	o.Steps = all.Steps
	o.TurnFinishedAt = all.TurnFinishedAt
	o.TurnStartedAt = all.TurnStartedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
