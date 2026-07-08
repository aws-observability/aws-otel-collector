// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyAttributes The attributes of a tag policy resource.
type TagPolicyAttributes struct {
	// The RFC 3339 timestamp at which the policy was created.
	CreatedAt time.Time `json:"created_at"`
	// The identifier of the user who created the policy.
	CreatedBy string `json:"created_by"`
	// The RFC 3339 timestamp at which the policy was soft-deleted. `null` if the policy has not been deleted. Only present when `include_deleted=true` is requested.
	DeletedAt datadog.NullableTime `json:"deleted_at,omitempty"`
	// The identifier of the user who soft-deleted the policy. `null` if the policy has not been deleted.
	DeletedBy datadog.NullableString `json:"deleted_by,omitempty"`
	// Whether the policy is currently enforced.
	Enabled bool `json:"enabled"`
	// The RFC 3339 timestamp at which the policy was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The identifier of the user who last modified the policy.
	ModifiedBy string `json:"modified_by"`
	// When `true`, the policy matches tag values that do NOT match any of the supplied patterns.
	Negated bool `json:"negated"`
	// Human-readable name for the tag policy.
	PolicyName string `json:"policy_name"`
	// How the policy is enforced. `blocking` rejects telemetry that violates the policy.
	// `surfacing` only highlights non-compliant telemetry without blocking it.
	PolicyType TagPolicyType `json:"policy_type"`
	// When `true`, telemetry without this tag is treated as a violation.
	Required bool `json:"required"`
	// The scope the policy applies within.
	Scope string `json:"scope"`
	// The telemetry source that a tag policy applies to.
	Source TagPolicySource `json:"source"`
	// The tag key that the policy governs.
	TagKey string `json:"tag_key"`
	// The patterns that valid values for the tag key must match.
	TagValuePatterns []string `json:"tag_value_patterns"`
	// A monotonically increasing version counter that is incremented on each update.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagPolicyAttributes instantiates a new TagPolicyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagPolicyAttributes(createdAt time.Time, createdBy string, enabled bool, modifiedAt time.Time, modifiedBy string, negated bool, policyName string, policyType TagPolicyType, required bool, scope string, source TagPolicySource, tagKey string, tagValuePatterns []string, version int64) *TagPolicyAttributes {
	this := TagPolicyAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Enabled = enabled
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Negated = negated
	this.PolicyName = policyName
	this.PolicyType = policyType
	this.Required = required
	this.Scope = scope
	this.Source = source
	this.TagKey = tagKey
	this.TagValuePatterns = tagValuePatterns
	this.Version = version
	return &this
}

// NewTagPolicyAttributesWithDefaults instantiates a new TagPolicyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagPolicyAttributesWithDefaults() *TagPolicyAttributes {
	this := TagPolicyAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *TagPolicyAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *TagPolicyAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *TagPolicyAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *TagPolicyAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagPolicyAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TagPolicyAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TagPolicyAttributes) HasDeletedAt() bool {
	return o != nil && o.DeletedAt.IsSet()
}

// SetDeletedAt gets a reference to the given datadog.NullableTime and assigns it to the DeletedAt field.
func (o *TagPolicyAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil.
func (o *TagPolicyAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil.
func (o *TagPolicyAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagPolicyAttributes) GetDeletedBy() string {
	if o == nil || o.DeletedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeletedBy.Get()
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TagPolicyAttributes) GetDeletedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedBy.Get(), o.DeletedBy.IsSet()
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *TagPolicyAttributes) HasDeletedBy() bool {
	return o != nil && o.DeletedBy.IsSet()
}

// SetDeletedBy gets a reference to the given datadog.NullableString and assigns it to the DeletedBy field.
func (o *TagPolicyAttributes) SetDeletedBy(v string) {
	o.DeletedBy.Set(&v)
}

// SetDeletedByNil sets the value for DeletedBy to be an explicit nil.
func (o *TagPolicyAttributes) SetDeletedByNil() {
	o.DeletedBy.Set(nil)
}

// UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil.
func (o *TagPolicyAttributes) UnsetDeletedBy() {
	o.DeletedBy.Unset()
}

// GetEnabled returns the Enabled field value.
func (o *TagPolicyAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *TagPolicyAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *TagPolicyAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *TagPolicyAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *TagPolicyAttributes) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *TagPolicyAttributes) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetNegated returns the Negated field value.
func (o *TagPolicyAttributes) GetNegated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetNegatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Negated, true
}

// SetNegated sets field value.
func (o *TagPolicyAttributes) SetNegated(v bool) {
	o.Negated = v
}

// GetPolicyName returns the PolicyName field value.
func (o *TagPolicyAttributes) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value.
func (o *TagPolicyAttributes) SetPolicyName(v string) {
	o.PolicyName = v
}

// GetPolicyType returns the PolicyType field value.
func (o *TagPolicyAttributes) GetPolicyType() TagPolicyType {
	if o == nil {
		var ret TagPolicyType
		return ret
	}
	return o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetPolicyTypeOk() (*TagPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyType, true
}

// SetPolicyType sets field value.
func (o *TagPolicyAttributes) SetPolicyType(v TagPolicyType) {
	o.PolicyType = v
}

// GetRequired returns the Required field value.
func (o *TagPolicyAttributes) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value.
func (o *TagPolicyAttributes) SetRequired(v bool) {
	o.Required = v
}

// GetScope returns the Scope field value.
func (o *TagPolicyAttributes) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *TagPolicyAttributes) SetScope(v string) {
	o.Scope = v
}

// GetSource returns the Source field value.
func (o *TagPolicyAttributes) GetSource() TagPolicySource {
	if o == nil {
		var ret TagPolicySource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetSourceOk() (*TagPolicySource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *TagPolicyAttributes) SetSource(v TagPolicySource) {
	o.Source = v
}

// GetTagKey returns the TagKey field value.
func (o *TagPolicyAttributes) GetTagKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKey, true
}

// SetTagKey sets field value.
func (o *TagPolicyAttributes) SetTagKey(v string) {
	o.TagKey = v
}

// GetTagValuePatterns returns the TagValuePatterns field value.
func (o *TagPolicyAttributes) GetTagValuePatterns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagValuePatterns
}

// GetTagValuePatternsOk returns a tuple with the TagValuePatterns field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetTagValuePatternsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagValuePatterns, true
}

// SetTagValuePatterns sets field value.
func (o *TagPolicyAttributes) SetTagValuePatterns(v []string) {
	o.TagValuePatterns = v
}

// GetVersion returns the Version field value.
func (o *TagPolicyAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *TagPolicyAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagPolicyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.DeletedBy.IsSet() {
		toSerialize["deleted_by"] = o.DeletedBy.Get()
	}
	toSerialize["enabled"] = o.Enabled
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["negated"] = o.Negated
	toSerialize["policy_name"] = o.PolicyName
	toSerialize["policy_type"] = o.PolicyType
	toSerialize["required"] = o.Required
	toSerialize["scope"] = o.Scope
	toSerialize["source"] = o.Source
	toSerialize["tag_key"] = o.TagKey
	toSerialize["tag_value_patterns"] = o.TagValuePatterns
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagPolicyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt        *time.Time             `json:"created_at"`
		CreatedBy        *string                `json:"created_by"`
		DeletedAt        datadog.NullableTime   `json:"deleted_at,omitempty"`
		DeletedBy        datadog.NullableString `json:"deleted_by,omitempty"`
		Enabled          *bool                  `json:"enabled"`
		ModifiedAt       *time.Time             `json:"modified_at"`
		ModifiedBy       *string                `json:"modified_by"`
		Negated          *bool                  `json:"negated"`
		PolicyName       *string                `json:"policy_name"`
		PolicyType       *TagPolicyType         `json:"policy_type"`
		Required         *bool                  `json:"required"`
		Scope            *string                `json:"scope"`
		Source           *TagPolicySource       `json:"source"`
		TagKey           *string                `json:"tag_key"`
		TagValuePatterns *[]string              `json:"tag_value_patterns"`
		Version          *int64                 `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	if all.Negated == nil {
		return fmt.Errorf("required field negated missing")
	}
	if all.PolicyName == nil {
		return fmt.Errorf("required field policy_name missing")
	}
	if all.PolicyType == nil {
		return fmt.Errorf("required field policy_type missing")
	}
	if all.Required == nil {
		return fmt.Errorf("required field required missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.TagKey == nil {
		return fmt.Errorf("required field tag_key missing")
	}
	if all.TagValuePatterns == nil {
		return fmt.Errorf("required field tag_value_patterns missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "deleted_at", "deleted_by", "enabled", "modified_at", "modified_by", "negated", "policy_name", "policy_type", "required", "scope", "source", "tag_key", "tag_value_patterns", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.DeletedAt = all.DeletedAt
	o.DeletedBy = all.DeletedBy
	o.Enabled = *all.Enabled
	o.ModifiedAt = *all.ModifiedAt
	o.ModifiedBy = *all.ModifiedBy
	o.Negated = *all.Negated
	o.PolicyName = *all.PolicyName
	if !all.PolicyType.IsValid() {
		hasInvalidField = true
	} else {
		o.PolicyType = *all.PolicyType
	}
	o.Required = *all.Required
	o.Scope = *all.Scope
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.TagKey = *all.TagKey
	o.TagValuePatterns = *all.TagValuePatterns
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
