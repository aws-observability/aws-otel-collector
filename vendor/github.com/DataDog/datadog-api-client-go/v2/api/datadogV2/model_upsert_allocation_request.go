// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertAllocationRequest Request to create or update a targeting rule (allocation) for a feature flag environment.
type UpsertAllocationRequest struct {
	// The experiment ID for experiment-linked allocations.
	ExperimentId datadog.NullableString `json:"experiment_id,omitempty"`
	// Progressive release request payload.
	ExposureSchedule *ExposureScheduleRequest `json:"exposure_schedule,omitempty"`
	// Guardrail metrics used to monitor and auto-pause or abort.
	GuardrailMetrics []GuardrailMetricRequest `json:"guardrail_metrics,omitempty"`
	// The unique identifier of the targeting rule allocation.
	Id *uuid.UUID `json:"id,omitempty"`
	// The unique key of the targeting rule allocation.
	Key string `json:"key"`
	// The display name of the targeting rule.
	Name string `json:"name"`
	// Targeting rules that determine audience eligibility.
	TargetingRules []TargetingRuleRequest `json:"targeting_rules,omitempty"`
	// The type of targeting rule (called allocation in the API model).
	Type AllocationType `json:"type"`
	// Variant distribution weights.
	VariantWeights []VariantWeightRequest `json:"variant_weights,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpsertAllocationRequest instantiates a new UpsertAllocationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpsertAllocationRequest(key string, name string, typeVar AllocationType) *UpsertAllocationRequest {
	this := UpsertAllocationRequest{}
	this.Key = key
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewUpsertAllocationRequestWithDefaults instantiates a new UpsertAllocationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpsertAllocationRequestWithDefaults() *UpsertAllocationRequest {
	this := UpsertAllocationRequest{}
	return &this
}

// GetExperimentId returns the ExperimentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertAllocationRequest) GetExperimentId() string {
	if o == nil || o.ExperimentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExperimentId.Get()
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UpsertAllocationRequest) GetExperimentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExperimentId.Get(), o.ExperimentId.IsSet()
}

// HasExperimentId returns a boolean if a field has been set.
func (o *UpsertAllocationRequest) HasExperimentId() bool {
	return o != nil && o.ExperimentId.IsSet()
}

// SetExperimentId gets a reference to the given datadog.NullableString and assigns it to the ExperimentId field.
func (o *UpsertAllocationRequest) SetExperimentId(v string) {
	o.ExperimentId.Set(&v)
}

// SetExperimentIdNil sets the value for ExperimentId to be an explicit nil.
func (o *UpsertAllocationRequest) SetExperimentIdNil() {
	o.ExperimentId.Set(nil)
}

// UnsetExperimentId ensures that no value is present for ExperimentId, not even an explicit nil.
func (o *UpsertAllocationRequest) UnsetExperimentId() {
	o.ExperimentId.Unset()
}

// GetExposureSchedule returns the ExposureSchedule field value if set, zero value otherwise.
func (o *UpsertAllocationRequest) GetExposureSchedule() ExposureScheduleRequest {
	if o == nil || o.ExposureSchedule == nil {
		var ret ExposureScheduleRequest
		return ret
	}
	return *o.ExposureSchedule
}

// GetExposureScheduleOk returns a tuple with the ExposureSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertAllocationRequest) GetExposureScheduleOk() (*ExposureScheduleRequest, bool) {
	if o == nil || o.ExposureSchedule == nil {
		return nil, false
	}
	return o.ExposureSchedule, true
}

// HasExposureSchedule returns a boolean if a field has been set.
func (o *UpsertAllocationRequest) HasExposureSchedule() bool {
	return o != nil && o.ExposureSchedule != nil
}

// SetExposureSchedule gets a reference to the given ExposureScheduleRequest and assigns it to the ExposureSchedule field.
func (o *UpsertAllocationRequest) SetExposureSchedule(v ExposureScheduleRequest) {
	o.ExposureSchedule = &v
}

// GetGuardrailMetrics returns the GuardrailMetrics field value if set, zero value otherwise.
func (o *UpsertAllocationRequest) GetGuardrailMetrics() []GuardrailMetricRequest {
	if o == nil || o.GuardrailMetrics == nil {
		var ret []GuardrailMetricRequest
		return ret
	}
	return o.GuardrailMetrics
}

// GetGuardrailMetricsOk returns a tuple with the GuardrailMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertAllocationRequest) GetGuardrailMetricsOk() (*[]GuardrailMetricRequest, bool) {
	if o == nil || o.GuardrailMetrics == nil {
		return nil, false
	}
	return &o.GuardrailMetrics, true
}

// HasGuardrailMetrics returns a boolean if a field has been set.
func (o *UpsertAllocationRequest) HasGuardrailMetrics() bool {
	return o != nil && o.GuardrailMetrics != nil
}

// SetGuardrailMetrics gets a reference to the given []GuardrailMetricRequest and assigns it to the GuardrailMetrics field.
func (o *UpsertAllocationRequest) SetGuardrailMetrics(v []GuardrailMetricRequest) {
	o.GuardrailMetrics = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpsertAllocationRequest) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertAllocationRequest) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpsertAllocationRequest) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *UpsertAllocationRequest) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetKey returns the Key field value.
func (o *UpsertAllocationRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *UpsertAllocationRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *UpsertAllocationRequest) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value.
func (o *UpsertAllocationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpsertAllocationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *UpsertAllocationRequest) SetName(v string) {
	o.Name = v
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise.
func (o *UpsertAllocationRequest) GetTargetingRules() []TargetingRuleRequest {
	if o == nil || o.TargetingRules == nil {
		var ret []TargetingRuleRequest
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertAllocationRequest) GetTargetingRulesOk() (*[]TargetingRuleRequest, bool) {
	if o == nil || o.TargetingRules == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *UpsertAllocationRequest) HasTargetingRules() bool {
	return o != nil && o.TargetingRules != nil
}

// SetTargetingRules gets a reference to the given []TargetingRuleRequest and assigns it to the TargetingRules field.
func (o *UpsertAllocationRequest) SetTargetingRules(v []TargetingRuleRequest) {
	o.TargetingRules = v
}

// GetType returns the Type field value.
func (o *UpsertAllocationRequest) GetType() AllocationType {
	if o == nil {
		var ret AllocationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpsertAllocationRequest) GetTypeOk() (*AllocationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UpsertAllocationRequest) SetType(v AllocationType) {
	o.Type = v
}

// GetVariantWeights returns the VariantWeights field value if set, zero value otherwise.
func (o *UpsertAllocationRequest) GetVariantWeights() []VariantWeightRequest {
	if o == nil || o.VariantWeights == nil {
		var ret []VariantWeightRequest
		return ret
	}
	return o.VariantWeights
}

// GetVariantWeightsOk returns a tuple with the VariantWeights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertAllocationRequest) GetVariantWeightsOk() (*[]VariantWeightRequest, bool) {
	if o == nil || o.VariantWeights == nil {
		return nil, false
	}
	return &o.VariantWeights, true
}

// HasVariantWeights returns a boolean if a field has been set.
func (o *UpsertAllocationRequest) HasVariantWeights() bool {
	return o != nil && o.VariantWeights != nil
}

// SetVariantWeights gets a reference to the given []VariantWeightRequest and assigns it to the VariantWeights field.
func (o *UpsertAllocationRequest) SetVariantWeights(v []VariantWeightRequest) {
	o.VariantWeights = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpsertAllocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExperimentId.IsSet() {
		toSerialize["experiment_id"] = o.ExperimentId.Get()
	}
	if o.ExposureSchedule != nil {
		toSerialize["exposure_schedule"] = o.ExposureSchedule
	}
	if o.GuardrailMetrics != nil {
		toSerialize["guardrail_metrics"] = o.GuardrailMetrics
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	if o.TargetingRules != nil {
		toSerialize["targeting_rules"] = o.TargetingRules
	}
	toSerialize["type"] = o.Type
	if o.VariantWeights != nil {
		toSerialize["variant_weights"] = o.VariantWeights
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpsertAllocationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExperimentId     datadog.NullableString   `json:"experiment_id,omitempty"`
		ExposureSchedule *ExposureScheduleRequest `json:"exposure_schedule,omitempty"`
		GuardrailMetrics []GuardrailMetricRequest `json:"guardrail_metrics,omitempty"`
		Id               *uuid.UUID               `json:"id,omitempty"`
		Key              *string                  `json:"key"`
		Name             *string                  `json:"name"`
		TargetingRules   []TargetingRuleRequest   `json:"targeting_rules,omitempty"`
		Type             *AllocationType          `json:"type"`
		VariantWeights   []VariantWeightRequest   `json:"variant_weights,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"experiment_id", "exposure_schedule", "guardrail_metrics", "id", "key", "name", "targeting_rules", "type", "variant_weights"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ExperimentId = all.ExperimentId
	if all.ExposureSchedule != nil && all.ExposureSchedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ExposureSchedule = all.ExposureSchedule
	o.GuardrailMetrics = all.GuardrailMetrics
	o.Id = all.Id
	o.Key = *all.Key
	o.Name = *all.Name
	o.TargetingRules = all.TargetingRules
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.VariantWeights = all.VariantWeights

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
