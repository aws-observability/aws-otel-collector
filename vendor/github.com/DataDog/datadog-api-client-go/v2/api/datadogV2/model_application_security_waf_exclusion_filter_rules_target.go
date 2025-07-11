// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafExclusionFilterRulesTarget Target WAF rules based either on an identifier or tags.
type ApplicationSecurityWafExclusionFilterRulesTarget struct {
	// Target a single WAF rule based on its identifier.
	RuleId *string `json:"rule_id,omitempty"`
	// Target multiple WAF rules based on their tags.
	Tags *ApplicationSecurityWafExclusionFilterRulesTargetTags `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityWafExclusionFilterRulesTarget instantiates a new ApplicationSecurityWafExclusionFilterRulesTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafExclusionFilterRulesTarget() *ApplicationSecurityWafExclusionFilterRulesTarget {
	this := ApplicationSecurityWafExclusionFilterRulesTarget{}
	return &this
}

// NewApplicationSecurityWafExclusionFilterRulesTargetWithDefaults instantiates a new ApplicationSecurityWafExclusionFilterRulesTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafExclusionFilterRulesTargetWithDefaults() *ApplicationSecurityWafExclusionFilterRulesTarget {
	this := ApplicationSecurityWafExclusionFilterRulesTarget{}
	return &this
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterRulesTarget) GetRuleId() string {
	if o == nil || o.RuleId == nil {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterRulesTarget) GetRuleIdOk() (*string, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterRulesTarget) HasRuleId() bool {
	return o != nil && o.RuleId != nil
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *ApplicationSecurityWafExclusionFilterRulesTarget) SetRuleId(v string) {
	o.RuleId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterRulesTarget) GetTags() ApplicationSecurityWafExclusionFilterRulesTargetTags {
	if o == nil || o.Tags == nil {
		var ret ApplicationSecurityWafExclusionFilterRulesTargetTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterRulesTarget) GetTagsOk() (*ApplicationSecurityWafExclusionFilterRulesTargetTags, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterRulesTarget) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given ApplicationSecurityWafExclusionFilterRulesTargetTags and assigns it to the Tags field.
func (o *ApplicationSecurityWafExclusionFilterRulesTarget) SetTags(v ApplicationSecurityWafExclusionFilterRulesTargetTags) {
	o.Tags = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafExclusionFilterRulesTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RuleId != nil {
		toSerialize["rule_id"] = o.RuleId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafExclusionFilterRulesTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RuleId *string                                               `json:"rule_id,omitempty"`
		Tags   *ApplicationSecurityWafExclusionFilterRulesTargetTags `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rule_id", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.RuleId = all.RuleId
	if all.Tags != nil && all.Tags.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
