// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagEnvironmentListItem Environment-specific settings for a feature flag in list responses.
type FeatureFlagEnvironmentListItem struct {
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
	// The status of a feature flag in an environment.
	Status FeatureFlagStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagEnvironmentListItem instantiates a new FeatureFlagEnvironmentListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagEnvironmentListItem(environmentId uuid.UUID, status FeatureFlagStatus) *FeatureFlagEnvironmentListItem {
	this := FeatureFlagEnvironmentListItem{}
	this.EnvironmentId = environmentId
	this.Status = status
	return &this
}

// NewFeatureFlagEnvironmentListItemWithDefaults instantiates a new FeatureFlagEnvironmentListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagEnvironmentListItemWithDefaults() *FeatureFlagEnvironmentListItem {
	this := FeatureFlagEnvironmentListItem{}
	return &this
}

// GetDefaultAllocationKey returns the DefaultAllocationKey field value if set, zero value otherwise.
func (o *FeatureFlagEnvironmentListItem) GetDefaultAllocationKey() string {
	if o == nil || o.DefaultAllocationKey == nil {
		var ret string
		return ret
	}
	return *o.DefaultAllocationKey
}

// GetDefaultAllocationKeyOk returns a tuple with the DefaultAllocationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironmentListItem) GetDefaultAllocationKeyOk() (*string, bool) {
	if o == nil || o.DefaultAllocationKey == nil {
		return nil, false
	}
	return o.DefaultAllocationKey, true
}

// HasDefaultAllocationKey returns a boolean if a field has been set.
func (o *FeatureFlagEnvironmentListItem) HasDefaultAllocationKey() bool {
	return o != nil && o.DefaultAllocationKey != nil
}

// SetDefaultAllocationKey gets a reference to the given string and assigns it to the DefaultAllocationKey field.
func (o *FeatureFlagEnvironmentListItem) SetDefaultAllocationKey(v string) {
	o.DefaultAllocationKey = &v
}

// GetDefaultVariantId returns the DefaultVariantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagEnvironmentListItem) GetDefaultVariantId() string {
	if o == nil || o.DefaultVariantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultVariantId.Get()
}

// GetDefaultVariantIdOk returns a tuple with the DefaultVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagEnvironmentListItem) GetDefaultVariantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultVariantId.Get(), o.DefaultVariantId.IsSet()
}

// HasDefaultVariantId returns a boolean if a field has been set.
func (o *FeatureFlagEnvironmentListItem) HasDefaultVariantId() bool {
	return o != nil && o.DefaultVariantId.IsSet()
}

// SetDefaultVariantId gets a reference to the given datadog.NullableString and assigns it to the DefaultVariantId field.
func (o *FeatureFlagEnvironmentListItem) SetDefaultVariantId(v string) {
	o.DefaultVariantId.Set(&v)
}

// SetDefaultVariantIdNil sets the value for DefaultVariantId to be an explicit nil.
func (o *FeatureFlagEnvironmentListItem) SetDefaultVariantIdNil() {
	o.DefaultVariantId.Set(nil)
}

// UnsetDefaultVariantId ensures that no value is present for DefaultVariantId, not even an explicit nil.
func (o *FeatureFlagEnvironmentListItem) UnsetDefaultVariantId() {
	o.DefaultVariantId.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value.
func (o *FeatureFlagEnvironmentListItem) GetEnvironmentId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironmentListItem) GetEnvironmentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value.
func (o *FeatureFlagEnvironmentListItem) SetEnvironmentId(v uuid.UUID) {
	o.EnvironmentId = v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *FeatureFlagEnvironmentListItem) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironmentListItem) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *FeatureFlagEnvironmentListItem) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *FeatureFlagEnvironmentListItem) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetEnvironmentQueries returns the EnvironmentQueries field value if set, zero value otherwise.
func (o *FeatureFlagEnvironmentListItem) GetEnvironmentQueries() []string {
	if o == nil || o.EnvironmentQueries == nil {
		var ret []string
		return ret
	}
	return o.EnvironmentQueries
}

// GetEnvironmentQueriesOk returns a tuple with the EnvironmentQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironmentListItem) GetEnvironmentQueriesOk() (*[]string, bool) {
	if o == nil || o.EnvironmentQueries == nil {
		return nil, false
	}
	return &o.EnvironmentQueries, true
}

// HasEnvironmentQueries returns a boolean if a field has been set.
func (o *FeatureFlagEnvironmentListItem) HasEnvironmentQueries() bool {
	return o != nil && o.EnvironmentQueries != nil
}

// SetEnvironmentQueries gets a reference to the given []string and assigns it to the EnvironmentQueries field.
func (o *FeatureFlagEnvironmentListItem) SetEnvironmentQueries(v []string) {
	o.EnvironmentQueries = v
}

// GetIsProduction returns the IsProduction field value if set, zero value otherwise.
func (o *FeatureFlagEnvironmentListItem) GetIsProduction() bool {
	if o == nil || o.IsProduction == nil {
		var ret bool
		return ret
	}
	return *o.IsProduction
}

// GetIsProductionOk returns a tuple with the IsProduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironmentListItem) GetIsProductionOk() (*bool, bool) {
	if o == nil || o.IsProduction == nil {
		return nil, false
	}
	return o.IsProduction, true
}

// HasIsProduction returns a boolean if a field has been set.
func (o *FeatureFlagEnvironmentListItem) HasIsProduction() bool {
	return o != nil && o.IsProduction != nil
}

// SetIsProduction gets a reference to the given bool and assigns it to the IsProduction field.
func (o *FeatureFlagEnvironmentListItem) SetIsProduction(v bool) {
	o.IsProduction = &v
}

// GetOverrideAllocationKey returns the OverrideAllocationKey field value if set, zero value otherwise.
func (o *FeatureFlagEnvironmentListItem) GetOverrideAllocationKey() string {
	if o == nil || o.OverrideAllocationKey == nil {
		var ret string
		return ret
	}
	return *o.OverrideAllocationKey
}

// GetOverrideAllocationKeyOk returns a tuple with the OverrideAllocationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironmentListItem) GetOverrideAllocationKeyOk() (*string, bool) {
	if o == nil || o.OverrideAllocationKey == nil {
		return nil, false
	}
	return o.OverrideAllocationKey, true
}

// HasOverrideAllocationKey returns a boolean if a field has been set.
func (o *FeatureFlagEnvironmentListItem) HasOverrideAllocationKey() bool {
	return o != nil && o.OverrideAllocationKey != nil
}

// SetOverrideAllocationKey gets a reference to the given string and assigns it to the OverrideAllocationKey field.
func (o *FeatureFlagEnvironmentListItem) SetOverrideAllocationKey(v string) {
	o.OverrideAllocationKey = &v
}

// GetOverrideVariantId returns the OverrideVariantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagEnvironmentListItem) GetOverrideVariantId() string {
	if o == nil || o.OverrideVariantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverrideVariantId.Get()
}

// GetOverrideVariantIdOk returns a tuple with the OverrideVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagEnvironmentListItem) GetOverrideVariantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverrideVariantId.Get(), o.OverrideVariantId.IsSet()
}

// HasOverrideVariantId returns a boolean if a field has been set.
func (o *FeatureFlagEnvironmentListItem) HasOverrideVariantId() bool {
	return o != nil && o.OverrideVariantId.IsSet()
}

// SetOverrideVariantId gets a reference to the given datadog.NullableString and assigns it to the OverrideVariantId field.
func (o *FeatureFlagEnvironmentListItem) SetOverrideVariantId(v string) {
	o.OverrideVariantId.Set(&v)
}

// SetOverrideVariantIdNil sets the value for OverrideVariantId to be an explicit nil.
func (o *FeatureFlagEnvironmentListItem) SetOverrideVariantIdNil() {
	o.OverrideVariantId.Set(nil)
}

// UnsetOverrideVariantId ensures that no value is present for OverrideVariantId, not even an explicit nil.
func (o *FeatureFlagEnvironmentListItem) UnsetOverrideVariantId() {
	o.OverrideVariantId.Unset()
}

// GetPendingSuggestionId returns the PendingSuggestionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagEnvironmentListItem) GetPendingSuggestionId() string {
	if o == nil || o.PendingSuggestionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PendingSuggestionId.Get()
}

// GetPendingSuggestionIdOk returns a tuple with the PendingSuggestionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagEnvironmentListItem) GetPendingSuggestionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingSuggestionId.Get(), o.PendingSuggestionId.IsSet()
}

// HasPendingSuggestionId returns a boolean if a field has been set.
func (o *FeatureFlagEnvironmentListItem) HasPendingSuggestionId() bool {
	return o != nil && o.PendingSuggestionId.IsSet()
}

// SetPendingSuggestionId gets a reference to the given datadog.NullableString and assigns it to the PendingSuggestionId field.
func (o *FeatureFlagEnvironmentListItem) SetPendingSuggestionId(v string) {
	o.PendingSuggestionId.Set(&v)
}

// SetPendingSuggestionIdNil sets the value for PendingSuggestionId to be an explicit nil.
func (o *FeatureFlagEnvironmentListItem) SetPendingSuggestionIdNil() {
	o.PendingSuggestionId.Set(nil)
}

// UnsetPendingSuggestionId ensures that no value is present for PendingSuggestionId, not even an explicit nil.
func (o *FeatureFlagEnvironmentListItem) UnsetPendingSuggestionId() {
	o.PendingSuggestionId.Unset()
}

// GetRequireFeatureFlagApproval returns the RequireFeatureFlagApproval field value if set, zero value otherwise.
func (o *FeatureFlagEnvironmentListItem) GetRequireFeatureFlagApproval() bool {
	if o == nil || o.RequireFeatureFlagApproval == nil {
		var ret bool
		return ret
	}
	return *o.RequireFeatureFlagApproval
}

// GetRequireFeatureFlagApprovalOk returns a tuple with the RequireFeatureFlagApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironmentListItem) GetRequireFeatureFlagApprovalOk() (*bool, bool) {
	if o == nil || o.RequireFeatureFlagApproval == nil {
		return nil, false
	}
	return o.RequireFeatureFlagApproval, true
}

// HasRequireFeatureFlagApproval returns a boolean if a field has been set.
func (o *FeatureFlagEnvironmentListItem) HasRequireFeatureFlagApproval() bool {
	return o != nil && o.RequireFeatureFlagApproval != nil
}

// SetRequireFeatureFlagApproval gets a reference to the given bool and assigns it to the RequireFeatureFlagApproval field.
func (o *FeatureFlagEnvironmentListItem) SetRequireFeatureFlagApproval(v bool) {
	o.RequireFeatureFlagApproval = &v
}

// GetStatus returns the Status field value.
func (o *FeatureFlagEnvironmentListItem) GetStatus() FeatureFlagStatus {
	if o == nil {
		var ret FeatureFlagStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagEnvironmentListItem) GetStatusOk() (*FeatureFlagStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *FeatureFlagEnvironmentListItem) SetStatus(v FeatureFlagStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagEnvironmentListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagEnvironmentListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultAllocationKey       *string                `json:"default_allocation_key,omitempty"`
		DefaultVariantId           datadog.NullableString `json:"default_variant_id,omitempty"`
		EnvironmentId              *uuid.UUID             `json:"environment_id"`
		EnvironmentName            *string                `json:"environment_name,omitempty"`
		EnvironmentQueries         []string               `json:"environment_queries,omitempty"`
		IsProduction               *bool                  `json:"is_production,omitempty"`
		OverrideAllocationKey      *string                `json:"override_allocation_key,omitempty"`
		OverrideVariantId          datadog.NullableString `json:"override_variant_id,omitempty"`
		PendingSuggestionId        datadog.NullableString `json:"pending_suggestion_id,omitempty"`
		RequireFeatureFlagApproval *bool                  `json:"require_feature_flag_approval,omitempty"`
		Status                     *FeatureFlagStatus     `json:"status"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default_allocation_key", "default_variant_id", "environment_id", "environment_name", "environment_queries", "is_production", "override_allocation_key", "override_variant_id", "pending_suggestion_id", "require_feature_flag_approval", "status"})
	} else {
		return err
	}

	hasInvalidField := false
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
