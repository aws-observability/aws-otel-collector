// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagEnvironment Environment-specific settings for a feature flag.
type FeatureFlagEnvironment struct {
	// Allocation metadata for this environment.
	Allocations map[string]interface{} `json:"allocations,omitempty"`
	// The allocation key used for the default variant.
	DefaultAllocationKey *string `json:"default_allocation_key,omitempty"`
	// The ID of the default variant for this environment.
	DefaultVariantId datadog.NullableString `json:"default_variant_id,omitempty"`
	// The ID of the environment.
	EnvironmentId uuid.UUID `json:"environment_id"`
	// The name of the environment.
	EnvironmentName *string `json:"environment_name,omitempty"`
	// Queries that target this environment.
	EnvironmentQueries []string `json:"environment_queries,omitempty"`
	// Indicates whether the environment is production.
	IsProduction *bool `json:"is_production,omitempty"`
	// The allocation key used for the override variant.
	OverrideAllocationKey *string `json:"override_allocation_key,omitempty"`
	// The ID of the override variant for this environment.
	OverrideVariantId datadog.NullableString `json:"override_variant_id,omitempty"`
	// Pending suggestion identifier, if approval is required.
	PendingSuggestionId datadog.NullableString `json:"pending_suggestion_id,omitempty"`
	// Indicates whether feature flag changes require approval in this environment.
	RequireFeatureFlagApproval *bool `json:"require_feature_flag_approval,omitempty"`
	// Rollout percentage for this environment.
	RolloutPercentage *int64 `json:"rollout_percentage,omitempty"`
	// Environment targeting rules for this feature flag.
	Rules []map[string]interface{} `json:"rules,omitempty"`
	// The status of a feature flag in an environment.
	Status FeatureFlagStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagEnvironment instantiates a new FeatureFlagEnvironment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagEnvironment(environmentId uuid.UUID, status FeatureFlagStatus) *FeatureFlagEnvironment {
	this := FeatureFlagEnvironment{}
	this.EnvironmentId = environmentId
	this.Status = status
	return &this
}

// NewFeatureFlagEnvironmentWithDefaults instantiates a new FeatureFlagEnvironment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagEnvironmentWithDefaults() *FeatureFlagEnvironment {
	this := FeatureFlagEnvironment{}
	return &this
}

// GetAllocations returns the Allocations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagEnvironment) GetAllocations() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Allocations
}

// GetAllocationsOk returns a tuple with the Allocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagEnvironment) GetAllocationsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Allocations == nil {
		return nil, false
	}
	return &o.Allocations, true
}

// HasAllocations returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasAllocations() bool {
	return o != nil && o.Allocations != nil
}

// SetAllocations gets a reference to the given map[string]interface{} and assigns it to the Allocations field.
func (o *FeatureFlagEnvironment) SetAllocations(v map[string]interface{}) {
	o.Allocations = v
}

// GetDefaultAllocationKey returns the DefaultAllocationKey field value if set, zero value otherwise.
func (o *FeatureFlagEnvironment) GetDefaultAllocationKey() string {
	if o == nil || o.DefaultAllocationKey == nil {
		var ret string
		return ret
	}
	return *o.DefaultAllocationKey
}

// GetDefaultAllocationKeyOk returns a tuple with the DefaultAllocationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironment) GetDefaultAllocationKeyOk() (*string, bool) {
	if o == nil || o.DefaultAllocationKey == nil {
		return nil, false
	}
	return o.DefaultAllocationKey, true
}

// HasDefaultAllocationKey returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasDefaultAllocationKey() bool {
	return o != nil && o.DefaultAllocationKey != nil
}

// SetDefaultAllocationKey gets a reference to the given string and assigns it to the DefaultAllocationKey field.
func (o *FeatureFlagEnvironment) SetDefaultAllocationKey(v string) {
	o.DefaultAllocationKey = &v
}

// GetDefaultVariantId returns the DefaultVariantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagEnvironment) GetDefaultVariantId() string {
	if o == nil || o.DefaultVariantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultVariantId.Get()
}

// GetDefaultVariantIdOk returns a tuple with the DefaultVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagEnvironment) GetDefaultVariantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultVariantId.Get(), o.DefaultVariantId.IsSet()
}

// HasDefaultVariantId returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasDefaultVariantId() bool {
	return o != nil && o.DefaultVariantId.IsSet()
}

// SetDefaultVariantId gets a reference to the given datadog.NullableString and assigns it to the DefaultVariantId field.
func (o *FeatureFlagEnvironment) SetDefaultVariantId(v string) {
	o.DefaultVariantId.Set(&v)
}

// SetDefaultVariantIdNil sets the value for DefaultVariantId to be an explicit nil.
func (o *FeatureFlagEnvironment) SetDefaultVariantIdNil() {
	o.DefaultVariantId.Set(nil)
}

// UnsetDefaultVariantId ensures that no value is present for DefaultVariantId, not even an explicit nil.
func (o *FeatureFlagEnvironment) UnsetDefaultVariantId() {
	o.DefaultVariantId.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value.
func (o *FeatureFlagEnvironment) GetEnvironmentId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironment) GetEnvironmentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value.
func (o *FeatureFlagEnvironment) SetEnvironmentId(v uuid.UUID) {
	o.EnvironmentId = v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *FeatureFlagEnvironment) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironment) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *FeatureFlagEnvironment) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetEnvironmentQueries returns the EnvironmentQueries field value if set, zero value otherwise.
func (o *FeatureFlagEnvironment) GetEnvironmentQueries() []string {
	if o == nil || o.EnvironmentQueries == nil {
		var ret []string
		return ret
	}
	return o.EnvironmentQueries
}

// GetEnvironmentQueriesOk returns a tuple with the EnvironmentQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironment) GetEnvironmentQueriesOk() (*[]string, bool) {
	if o == nil || o.EnvironmentQueries == nil {
		return nil, false
	}
	return &o.EnvironmentQueries, true
}

// HasEnvironmentQueries returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasEnvironmentQueries() bool {
	return o != nil && o.EnvironmentQueries != nil
}

// SetEnvironmentQueries gets a reference to the given []string and assigns it to the EnvironmentQueries field.
func (o *FeatureFlagEnvironment) SetEnvironmentQueries(v []string) {
	o.EnvironmentQueries = v
}

// GetIsProduction returns the IsProduction field value if set, zero value otherwise.
func (o *FeatureFlagEnvironment) GetIsProduction() bool {
	if o == nil || o.IsProduction == nil {
		var ret bool
		return ret
	}
	return *o.IsProduction
}

// GetIsProductionOk returns a tuple with the IsProduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironment) GetIsProductionOk() (*bool, bool) {
	if o == nil || o.IsProduction == nil {
		return nil, false
	}
	return o.IsProduction, true
}

// HasIsProduction returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasIsProduction() bool {
	return o != nil && o.IsProduction != nil
}

// SetIsProduction gets a reference to the given bool and assigns it to the IsProduction field.
func (o *FeatureFlagEnvironment) SetIsProduction(v bool) {
	o.IsProduction = &v
}

// GetOverrideAllocationKey returns the OverrideAllocationKey field value if set, zero value otherwise.
func (o *FeatureFlagEnvironment) GetOverrideAllocationKey() string {
	if o == nil || o.OverrideAllocationKey == nil {
		var ret string
		return ret
	}
	return *o.OverrideAllocationKey
}

// GetOverrideAllocationKeyOk returns a tuple with the OverrideAllocationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironment) GetOverrideAllocationKeyOk() (*string, bool) {
	if o == nil || o.OverrideAllocationKey == nil {
		return nil, false
	}
	return o.OverrideAllocationKey, true
}

// HasOverrideAllocationKey returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasOverrideAllocationKey() bool {
	return o != nil && o.OverrideAllocationKey != nil
}

// SetOverrideAllocationKey gets a reference to the given string and assigns it to the OverrideAllocationKey field.
func (o *FeatureFlagEnvironment) SetOverrideAllocationKey(v string) {
	o.OverrideAllocationKey = &v
}

// GetOverrideVariantId returns the OverrideVariantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagEnvironment) GetOverrideVariantId() string {
	if o == nil || o.OverrideVariantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverrideVariantId.Get()
}

// GetOverrideVariantIdOk returns a tuple with the OverrideVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagEnvironment) GetOverrideVariantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverrideVariantId.Get(), o.OverrideVariantId.IsSet()
}

// HasOverrideVariantId returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasOverrideVariantId() bool {
	return o != nil && o.OverrideVariantId.IsSet()
}

// SetOverrideVariantId gets a reference to the given datadog.NullableString and assigns it to the OverrideVariantId field.
func (o *FeatureFlagEnvironment) SetOverrideVariantId(v string) {
	o.OverrideVariantId.Set(&v)
}

// SetOverrideVariantIdNil sets the value for OverrideVariantId to be an explicit nil.
func (o *FeatureFlagEnvironment) SetOverrideVariantIdNil() {
	o.OverrideVariantId.Set(nil)
}

// UnsetOverrideVariantId ensures that no value is present for OverrideVariantId, not even an explicit nil.
func (o *FeatureFlagEnvironment) UnsetOverrideVariantId() {
	o.OverrideVariantId.Unset()
}

// GetPendingSuggestionId returns the PendingSuggestionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagEnvironment) GetPendingSuggestionId() string {
	if o == nil || o.PendingSuggestionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PendingSuggestionId.Get()
}

// GetPendingSuggestionIdOk returns a tuple with the PendingSuggestionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagEnvironment) GetPendingSuggestionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingSuggestionId.Get(), o.PendingSuggestionId.IsSet()
}

// HasPendingSuggestionId returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasPendingSuggestionId() bool {
	return o != nil && o.PendingSuggestionId.IsSet()
}

// SetPendingSuggestionId gets a reference to the given datadog.NullableString and assigns it to the PendingSuggestionId field.
func (o *FeatureFlagEnvironment) SetPendingSuggestionId(v string) {
	o.PendingSuggestionId.Set(&v)
}

// SetPendingSuggestionIdNil sets the value for PendingSuggestionId to be an explicit nil.
func (o *FeatureFlagEnvironment) SetPendingSuggestionIdNil() {
	o.PendingSuggestionId.Set(nil)
}

// UnsetPendingSuggestionId ensures that no value is present for PendingSuggestionId, not even an explicit nil.
func (o *FeatureFlagEnvironment) UnsetPendingSuggestionId() {
	o.PendingSuggestionId.Unset()
}

// GetRequireFeatureFlagApproval returns the RequireFeatureFlagApproval field value if set, zero value otherwise.
func (o *FeatureFlagEnvironment) GetRequireFeatureFlagApproval() bool {
	if o == nil || o.RequireFeatureFlagApproval == nil {
		var ret bool
		return ret
	}
	return *o.RequireFeatureFlagApproval
}

// GetRequireFeatureFlagApprovalOk returns a tuple with the RequireFeatureFlagApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironment) GetRequireFeatureFlagApprovalOk() (*bool, bool) {
	if o == nil || o.RequireFeatureFlagApproval == nil {
		return nil, false
	}
	return o.RequireFeatureFlagApproval, true
}

// HasRequireFeatureFlagApproval returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasRequireFeatureFlagApproval() bool {
	return o != nil && o.RequireFeatureFlagApproval != nil
}

// SetRequireFeatureFlagApproval gets a reference to the given bool and assigns it to the RequireFeatureFlagApproval field.
func (o *FeatureFlagEnvironment) SetRequireFeatureFlagApproval(v bool) {
	o.RequireFeatureFlagApproval = &v
}

// GetRolloutPercentage returns the RolloutPercentage field value if set, zero value otherwise.
func (o *FeatureFlagEnvironment) GetRolloutPercentage() int64 {
	if o == nil || o.RolloutPercentage == nil {
		var ret int64
		return ret
	}
	return *o.RolloutPercentage
}

// GetRolloutPercentageOk returns a tuple with the RolloutPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironment) GetRolloutPercentageOk() (*int64, bool) {
	if o == nil || o.RolloutPercentage == nil {
		return nil, false
	}
	return o.RolloutPercentage, true
}

// HasRolloutPercentage returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasRolloutPercentage() bool {
	return o != nil && o.RolloutPercentage != nil
}

// SetRolloutPercentage gets a reference to the given int64 and assigns it to the RolloutPercentage field.
func (o *FeatureFlagEnvironment) SetRolloutPercentage(v int64) {
	o.RolloutPercentage = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *FeatureFlagEnvironment) GetRules() []map[string]interface{} {
	if o == nil || o.Rules == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironment) GetRulesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *FeatureFlagEnvironment) HasRules() bool {
	return o != nil && o.Rules != nil
}

// SetRules gets a reference to the given []map[string]interface{} and assigns it to the Rules field.
func (o *FeatureFlagEnvironment) SetRules(v []map[string]interface{}) {
	o.Rules = v
}

// GetStatus returns the Status field value.
func (o *FeatureFlagEnvironment) GetStatus() FeatureFlagStatus {
	if o == nil {
		var ret FeatureFlagStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironment) GetStatusOk() (*FeatureFlagStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *FeatureFlagEnvironment) SetStatus(v FeatureFlagStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Allocations != nil {
		toSerialize["allocations"] = o.Allocations
	}
	if o.DefaultAllocationKey != nil {
		toSerialize["default_allocation_key"] = o.DefaultAllocationKey
	}
	if o.DefaultVariantId.IsSet() {
		toSerialize["default_variant_id"] = o.DefaultVariantId.Get()
	}
	toSerialize["environment_id"] = o.EnvironmentId
	if o.EnvironmentName != nil {
		toSerialize["environment_name"] = o.EnvironmentName
	}
	if o.EnvironmentQueries != nil {
		toSerialize["environment_queries"] = o.EnvironmentQueries
	}
	if o.IsProduction != nil {
		toSerialize["is_production"] = o.IsProduction
	}
	if o.OverrideAllocationKey != nil {
		toSerialize["override_allocation_key"] = o.OverrideAllocationKey
	}
	if o.OverrideVariantId.IsSet() {
		toSerialize["override_variant_id"] = o.OverrideVariantId.Get()
	}
	if o.PendingSuggestionId.IsSet() {
		toSerialize["pending_suggestion_id"] = o.PendingSuggestionId.Get()
	}
	if o.RequireFeatureFlagApproval != nil {
		toSerialize["require_feature_flag_approval"] = o.RequireFeatureFlagApproval
	}
	if o.RolloutPercentage != nil {
		toSerialize["rollout_percentage"] = o.RolloutPercentage
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagEnvironment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Allocations                map[string]interface{}   `json:"allocations,omitempty"`
		DefaultAllocationKey       *string                  `json:"default_allocation_key,omitempty"`
		DefaultVariantId           datadog.NullableString   `json:"default_variant_id,omitempty"`
		EnvironmentId              *uuid.UUID               `json:"environment_id"`
		EnvironmentName            *string                  `json:"environment_name,omitempty"`
		EnvironmentQueries         []string                 `json:"environment_queries,omitempty"`
		IsProduction               *bool                    `json:"is_production,omitempty"`
		OverrideAllocationKey      *string                  `json:"override_allocation_key,omitempty"`
		OverrideVariantId          datadog.NullableString   `json:"override_variant_id,omitempty"`
		PendingSuggestionId        datadog.NullableString   `json:"pending_suggestion_id,omitempty"`
		RequireFeatureFlagApproval *bool                    `json:"require_feature_flag_approval,omitempty"`
		RolloutPercentage          *int64                   `json:"rollout_percentage,omitempty"`
		Rules                      []map[string]interface{} `json:"rules,omitempty"`
		Status                     *FeatureFlagStatus       `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EnvironmentId == nil {
		return fmt.Errorf("required field environment_id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allocations", "default_allocation_key", "default_variant_id", "environment_id", "environment_name", "environment_queries", "is_production", "override_allocation_key", "override_variant_id", "pending_suggestion_id", "require_feature_flag_approval", "rollout_percentage", "rules", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Allocations = all.Allocations
	o.DefaultAllocationKey = all.DefaultAllocationKey
	o.DefaultVariantId = all.DefaultVariantId
	o.EnvironmentId = *all.EnvironmentId
	o.EnvironmentName = all.EnvironmentName
	o.EnvironmentQueries = all.EnvironmentQueries
	o.IsProduction = all.IsProduction
	o.OverrideAllocationKey = all.OverrideAllocationKey
	o.OverrideVariantId = all.OverrideVariantId
	o.PendingSuggestionId = all.PendingSuggestionId
	o.RequireFeatureFlagApproval = all.RequireFeatureFlagApproval
	o.RolloutPercentage = all.RolloutPercentage
	o.Rules = all.Rules
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
