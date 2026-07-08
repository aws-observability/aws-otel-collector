// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyUpdateAttributes Mutable attributes of a tag policy. Each field is optional; omitting a field leaves its
// current value unchanged. The `source` of a policy cannot be changed.
type TagPolicyUpdateAttributes struct {
	// Whether the policy is currently enforced.
	Enabled *bool `json:"enabled,omitempty"`
	// When `true`, the policy matches tag values that do NOT match any of the supplied patterns.
	Negated *bool `json:"negated,omitempty"`
	// Human-readable name for the tag policy.
	PolicyName *string `json:"policy_name,omitempty"`
	// How the policy is enforced. `blocking` rejects telemetry that violates the policy.
	// `surfacing` only highlights non-compliant telemetry without blocking it.
	PolicyType *TagPolicyType `json:"policy_type,omitempty"`
	// When `true`, telemetry without this tag is treated as a violation.
	Required *bool `json:"required,omitempty"`
	// The scope the policy applies within.
	Scope *string `json:"scope,omitempty"`
	// The tag key that the policy governs.
	TagKey *string `json:"tag_key,omitempty"`
	// One or more patterns that valid values for the tag key must match.
	TagValuePatterns []string `json:"tag_value_patterns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagPolicyUpdateAttributes instantiates a new TagPolicyUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagPolicyUpdateAttributes() *TagPolicyUpdateAttributes {
	this := TagPolicyUpdateAttributes{}
	return &this
}

// NewTagPolicyUpdateAttributesWithDefaults instantiates a new TagPolicyUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagPolicyUpdateAttributesWithDefaults() *TagPolicyUpdateAttributes {
	this := TagPolicyUpdateAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TagPolicyUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TagPolicyUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TagPolicyUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNegated returns the Negated field value if set, zero value otherwise.
func (o *TagPolicyUpdateAttributes) GetNegated() bool {
	if o == nil || o.Negated == nil {
		var ret bool
		return ret
	}
	return *o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyUpdateAttributes) GetNegatedOk() (*bool, bool) {
	if o == nil || o.Negated == nil {
		return nil, false
	}
	return o.Negated, true
}

// HasNegated returns a boolean if a field has been set.
func (o *TagPolicyUpdateAttributes) HasNegated() bool {
	return o != nil && o.Negated != nil
}

// SetNegated gets a reference to the given bool and assigns it to the Negated field.
func (o *TagPolicyUpdateAttributes) SetNegated(v bool) {
	o.Negated = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *TagPolicyUpdateAttributes) GetPolicyName() string {
	if o == nil || o.PolicyName == nil {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyUpdateAttributes) GetPolicyNameOk() (*string, bool) {
	if o == nil || o.PolicyName == nil {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *TagPolicyUpdateAttributes) HasPolicyName() bool {
	return o != nil && o.PolicyName != nil
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *TagPolicyUpdateAttributes) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetPolicyType returns the PolicyType field value if set, zero value otherwise.
func (o *TagPolicyUpdateAttributes) GetPolicyType() TagPolicyType {
	if o == nil || o.PolicyType == nil {
		var ret TagPolicyType
		return ret
	}
	return *o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyUpdateAttributes) GetPolicyTypeOk() (*TagPolicyType, bool) {
	if o == nil || o.PolicyType == nil {
		return nil, false
	}
	return o.PolicyType, true
}

// HasPolicyType returns a boolean if a field has been set.
func (o *TagPolicyUpdateAttributes) HasPolicyType() bool {
	return o != nil && o.PolicyType != nil
}

// SetPolicyType gets a reference to the given TagPolicyType and assigns it to the PolicyType field.
func (o *TagPolicyUpdateAttributes) SetPolicyType(v TagPolicyType) {
	o.PolicyType = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *TagPolicyUpdateAttributes) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyUpdateAttributes) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *TagPolicyUpdateAttributes) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *TagPolicyUpdateAttributes) SetRequired(v bool) {
	o.Required = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *TagPolicyUpdateAttributes) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyUpdateAttributes) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *TagPolicyUpdateAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *TagPolicyUpdateAttributes) SetScope(v string) {
	o.Scope = &v
}

// GetTagKey returns the TagKey field value if set, zero value otherwise.
func (o *TagPolicyUpdateAttributes) GetTagKey() string {
	if o == nil || o.TagKey == nil {
		var ret string
		return ret
	}
	return *o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyUpdateAttributes) GetTagKeyOk() (*string, bool) {
	if o == nil || o.TagKey == nil {
		return nil, false
	}
	return o.TagKey, true
}

// HasTagKey returns a boolean if a field has been set.
func (o *TagPolicyUpdateAttributes) HasTagKey() bool {
	return o != nil && o.TagKey != nil
}

// SetTagKey gets a reference to the given string and assigns it to the TagKey field.
func (o *TagPolicyUpdateAttributes) SetTagKey(v string) {
	o.TagKey = &v
}

// GetTagValuePatterns returns the TagValuePatterns field value if set, zero value otherwise.
func (o *TagPolicyUpdateAttributes) GetTagValuePatterns() []string {
	if o == nil || o.TagValuePatterns == nil {
		var ret []string
		return ret
	}
	return o.TagValuePatterns
}

// GetTagValuePatternsOk returns a tuple with the TagValuePatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyUpdateAttributes) GetTagValuePatternsOk() (*[]string, bool) {
	if o == nil || o.TagValuePatterns == nil {
		return nil, false
	}
	return &o.TagValuePatterns, true
}

// HasTagValuePatterns returns a boolean if a field has been set.
func (o *TagPolicyUpdateAttributes) HasTagValuePatterns() bool {
	return o != nil && o.TagValuePatterns != nil
}

// SetTagValuePatterns gets a reference to the given []string and assigns it to the TagValuePatterns field.
func (o *TagPolicyUpdateAttributes) SetTagValuePatterns(v []string) {
	o.TagValuePatterns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagPolicyUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Negated != nil {
		toSerialize["negated"] = o.Negated
	}
	if o.PolicyName != nil {
		toSerialize["policy_name"] = o.PolicyName
	}
	if o.PolicyType != nil {
		toSerialize["policy_type"] = o.PolicyType
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.TagKey != nil {
		toSerialize["tag_key"] = o.TagKey
	}
	if o.TagValuePatterns != nil {
		toSerialize["tag_value_patterns"] = o.TagValuePatterns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagPolicyUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled          *bool          `json:"enabled,omitempty"`
		Negated          *bool          `json:"negated,omitempty"`
		PolicyName       *string        `json:"policy_name,omitempty"`
		PolicyType       *TagPolicyType `json:"policy_type,omitempty"`
		Required         *bool          `json:"required,omitempty"`
		Scope            *string        `json:"scope,omitempty"`
		TagKey           *string        `json:"tag_key,omitempty"`
		TagValuePatterns []string       `json:"tag_value_patterns,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "negated", "policy_name", "policy_type", "required", "scope", "tag_key", "tag_value_patterns"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	o.Negated = all.Negated
	o.PolicyName = all.PolicyName
	if all.PolicyType != nil && !all.PolicyType.IsValid() {
		hasInvalidField = true
	} else {
		o.PolicyType = all.PolicyType
	}
	o.Required = all.Required
	o.Scope = all.Scope
	o.TagKey = all.TagKey
	o.TagValuePatterns = all.TagValuePatterns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
