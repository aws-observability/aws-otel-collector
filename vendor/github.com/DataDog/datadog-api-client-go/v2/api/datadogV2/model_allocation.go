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

// Allocation Targeting rule (allocation) details for a feature flag environment.
type Allocation struct {
	// The timestamp when the targeting rule allocation was created.
	CreatedAt time.Time `json:"created_at"`
	// Environment IDs associated with this targeting rule allocation.
	EnvironmentIds []uuid.UUID `json:"environment_ids"`
	// The experiment ID linked to this targeting rule allocation.
	ExperimentId datadog.NullableString `json:"experiment_id,omitempty"`
	// Progressive release details for a targeting rule allocation.
	ExposureSchedule *AllocationExposureSchedule `json:"exposure_schedule,omitempty"`
	// Guardrail metrics associated with this targeting rule allocation.
	GuardrailMetrics []GuardrailMetric `json:"guardrail_metrics"`
	// The unique identifier of the targeting rule allocation.
	Id *uuid.UUID `json:"id,omitempty"`
	// The unique key of the targeting rule allocation.
	Key string `json:"key"`
	// The display name of the targeting rule.
	Name string `json:"name"`
	// Sort order position within the environment.
	OrderPosition int64 `json:"order_position"`
	// Conditions associated with this targeting rule allocation.
	TargetingRules []TargetingRule `json:"targeting_rules"`
	// The type of targeting rule (called allocation in the API model).
	Type AllocationType `json:"type"`
	// The timestamp when the targeting rule allocation was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// Weighted variant assignments for this targeting rule allocation.
	VariantWeights []VariantWeight `json:"variant_weights"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAllocation instantiates a new Allocation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAllocation(createdAt time.Time, environmentIds []uuid.UUID, guardrailMetrics []GuardrailMetric, key string, name string, orderPosition int64, targetingRules []TargetingRule, typeVar AllocationType, updatedAt time.Time, variantWeights []VariantWeight) *Allocation {
	this := Allocation{}
	this.CreatedAt = createdAt
	this.EnvironmentIds = environmentIds
	this.GuardrailMetrics = guardrailMetrics
	this.Key = key
	this.Name = name
	this.OrderPosition = orderPosition
	this.TargetingRules = targetingRules
	this.Type = typeVar
	this.UpdatedAt = updatedAt
	this.VariantWeights = variantWeights
	return &this
}

// NewAllocationWithDefaults instantiates a new Allocation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAllocationWithDefaults() *Allocation {
	this := Allocation{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *Allocation) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Allocation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Allocation) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvironmentIds returns the EnvironmentIds field value.
func (o *Allocation) GetEnvironmentIds() []uuid.UUID {
	if o == nil {
		var ret []uuid.UUID
		return ret
	}
	return o.EnvironmentIds
}

// GetEnvironmentIdsOk returns a tuple with the EnvironmentIds field value
// and a boolean to check if the value has been set.
func (o *Allocation) GetEnvironmentIdsOk() (*[]uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentIds, true
}

// SetEnvironmentIds sets field value.
func (o *Allocation) SetEnvironmentIds(v []uuid.UUID) {
	o.EnvironmentIds = v
}

// GetExperimentId returns the ExperimentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Allocation) GetExperimentId() string {
	if o == nil || o.ExperimentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExperimentId.Get()
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Allocation) GetExperimentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExperimentId.Get(), o.ExperimentId.IsSet()
}

// HasExperimentId returns a boolean if a field has been set.
func (o *Allocation) HasExperimentId() bool {
	return o != nil && o.ExperimentId.IsSet()
}

// SetExperimentId gets a reference to the given datadog.NullableString and assigns it to the ExperimentId field.
func (o *Allocation) SetExperimentId(v string) {
	o.ExperimentId.Set(&v)
}

// SetExperimentIdNil sets the value for ExperimentId to be an explicit nil.
func (o *Allocation) SetExperimentIdNil() {
	o.ExperimentId.Set(nil)
}

// UnsetExperimentId ensures that no value is present for ExperimentId, not even an explicit nil.
func (o *Allocation) UnsetExperimentId() {
	o.ExperimentId.Unset()
}

// GetExposureSchedule returns the ExposureSchedule field value if set, zero value otherwise.
func (o *Allocation) GetExposureSchedule() AllocationExposureSchedule {
	if o == nil || o.ExposureSchedule == nil {
		var ret AllocationExposureSchedule
		return ret
	}
	return *o.ExposureSchedule
}

// GetExposureScheduleOk returns a tuple with the ExposureSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allocation) GetExposureScheduleOk() (*AllocationExposureSchedule, bool) {
	if o == nil || o.ExposureSchedule == nil {
		return nil, false
	}
	return o.ExposureSchedule, true
}

// HasExposureSchedule returns a boolean if a field has been set.
func (o *Allocation) HasExposureSchedule() bool {
	return o != nil && o.ExposureSchedule != nil
}

// SetExposureSchedule gets a reference to the given AllocationExposureSchedule and assigns it to the ExposureSchedule field.
func (o *Allocation) SetExposureSchedule(v AllocationExposureSchedule) {
	o.ExposureSchedule = &v
}

// GetGuardrailMetrics returns the GuardrailMetrics field value.
func (o *Allocation) GetGuardrailMetrics() []GuardrailMetric {
	if o == nil {
		var ret []GuardrailMetric
		return ret
	}
	return o.GuardrailMetrics
}

// GetGuardrailMetricsOk returns a tuple with the GuardrailMetrics field value
// and a boolean to check if the value has been set.
func (o *Allocation) GetGuardrailMetricsOk() (*[]GuardrailMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuardrailMetrics, true
}

// SetGuardrailMetrics sets field value.
func (o *Allocation) SetGuardrailMetrics(v []GuardrailMetric) {
	o.GuardrailMetrics = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Allocation) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allocation) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Allocation) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *Allocation) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetKey returns the Key field value.
func (o *Allocation) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Allocation) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *Allocation) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value.
func (o *Allocation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Allocation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Allocation) SetName(v string) {
	o.Name = v
}

// GetOrderPosition returns the OrderPosition field value.
func (o *Allocation) GetOrderPosition() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrderPosition
}

// GetOrderPositionOk returns a tuple with the OrderPosition field value
// and a boolean to check if the value has been set.
func (o *Allocation) GetOrderPositionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderPosition, true
}

// SetOrderPosition sets field value.
func (o *Allocation) SetOrderPosition(v int64) {
	o.OrderPosition = v
}

// GetTargetingRules returns the TargetingRules field value.
func (o *Allocation) GetTargetingRules() []TargetingRule {
	if o == nil {
		var ret []TargetingRule
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value
// and a boolean to check if the value has been set.
func (o *Allocation) GetTargetingRulesOk() (*[]TargetingRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// SetTargetingRules sets field value.
func (o *Allocation) SetTargetingRules(v []TargetingRule) {
	o.TargetingRules = v
}

// GetType returns the Type field value.
func (o *Allocation) GetType() AllocationType {
	if o == nil {
		var ret AllocationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Allocation) GetTypeOk() (*AllocationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Allocation) SetType(v AllocationType) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *Allocation) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Allocation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *Allocation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVariantWeights returns the VariantWeights field value.
func (o *Allocation) GetVariantWeights() []VariantWeight {
	if o == nil {
		var ret []VariantWeight
		return ret
	}
	return o.VariantWeights
}

// GetVariantWeightsOk returns a tuple with the VariantWeights field value
// and a boolean to check if the value has been set.
func (o *Allocation) GetVariantWeightsOk() (*[]VariantWeight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariantWeights, true
}

// SetVariantWeights sets field value.
func (o *Allocation) SetVariantWeights(v []VariantWeight) {
	o.VariantWeights = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Allocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["environment_ids"] = o.EnvironmentIds
	if o.ExperimentId.IsSet() {
		toSerialize["experiment_id"] = o.ExperimentId.Get()
	}
	if o.ExposureSchedule != nil {
		toSerialize["exposure_schedule"] = o.ExposureSchedule
	}
	toSerialize["guardrail_metrics"] = o.GuardrailMetrics
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["order_position"] = o.OrderPosition
	toSerialize["targeting_rules"] = o.TargetingRules
	toSerialize["type"] = o.Type
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["variant_weights"] = o.VariantWeights

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Allocation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt        *time.Time                  `json:"created_at"`
		EnvironmentIds   *[]uuid.UUID                `json:"environment_ids"`
		ExperimentId     datadog.NullableString      `json:"experiment_id,omitempty"`
		ExposureSchedule *AllocationExposureSchedule `json:"exposure_schedule,omitempty"`
		GuardrailMetrics *[]GuardrailMetric          `json:"guardrail_metrics"`
		Id               *uuid.UUID                  `json:"id,omitempty"`
		Key              *string                     `json:"key"`
		Name             *string                     `json:"name"`
		OrderPosition    *int64                      `json:"order_position"`
		TargetingRules   *[]TargetingRule            `json:"targeting_rules"`
		Type             *AllocationType             `json:"type"`
		UpdatedAt        *time.Time                  `json:"updated_at"`
		VariantWeights   *[]VariantWeight            `json:"variant_weights"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.EnvironmentIds == nil {
		return fmt.Errorf("required field environment_ids missing")
	}
	if all.GuardrailMetrics == nil {
		return fmt.Errorf("required field guardrail_metrics missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrderPosition == nil {
		return fmt.Errorf("required field order_position missing")
	}
	if all.TargetingRules == nil {
		return fmt.Errorf("required field targeting_rules missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.VariantWeights == nil {
		return fmt.Errorf("required field variant_weights missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "environment_ids", "experiment_id", "exposure_schedule", "guardrail_metrics", "id", "key", "name", "order_position", "targeting_rules", "type", "updated_at", "variant_weights"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.EnvironmentIds = *all.EnvironmentIds
	o.ExperimentId = all.ExperimentId
	if all.ExposureSchedule != nil && all.ExposureSchedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ExposureSchedule = all.ExposureSchedule
	o.GuardrailMetrics = *all.GuardrailMetrics
	o.Id = all.Id
	o.Key = *all.Key
	o.Name = *all.Name
	o.OrderPosition = *all.OrderPosition
	o.TargetingRules = *all.TargetingRules
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.UpdatedAt = *all.UpdatedAt
	o.VariantWeights = *all.VariantWeights

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
