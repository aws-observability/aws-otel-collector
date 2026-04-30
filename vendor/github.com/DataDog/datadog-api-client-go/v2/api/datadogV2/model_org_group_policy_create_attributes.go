// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyCreateAttributes Attributes for creating an org group policy. If `policy_type` or `enforcement_tier` are not provided, they default to `org_config` and `DEFAULT` respectively.
type OrgGroupPolicyCreateAttributes struct {
	// The policy content as key-value pairs.
	Content map[string]interface{} `json:"content"`
	// The enforcement tier of the policy. `DEFAULT` means the policy is set but member orgs may mutate it. `ENFORCE` means the policy is strictly controlled and mutations are blocked for affected orgs. `DELEGATE` means each member org controls its own value.
	EnforcementTier *OrgGroupPolicyEnforcementTier `json:"enforcement_tier,omitempty"`
	// The name of the policy.
	PolicyName string `json:"policy_name"`
	// The type of the policy. Only `org_config` is supported, indicating a policy backed by an organization configuration setting.
	PolicyType *OrgGroupPolicyPolicyType `json:"policy_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyCreateAttributes instantiates a new OrgGroupPolicyCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyCreateAttributes(content map[string]interface{}, policyName string) *OrgGroupPolicyCreateAttributes {
	this := OrgGroupPolicyCreateAttributes{}
	this.Content = content
	var enforcementTier OrgGroupPolicyEnforcementTier = ORGGROUPPOLICYENFORCEMENTTIER_DEFAULT
	this.EnforcementTier = &enforcementTier
	this.PolicyName = policyName
	var policyType OrgGroupPolicyPolicyType = ORGGROUPPOLICYPOLICYTYPE_ORG_CONFIG
	this.PolicyType = &policyType
	return &this
}

// NewOrgGroupPolicyCreateAttributesWithDefaults instantiates a new OrgGroupPolicyCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyCreateAttributesWithDefaults() *OrgGroupPolicyCreateAttributes {
	this := OrgGroupPolicyCreateAttributes{}
	var enforcementTier OrgGroupPolicyEnforcementTier = ORGGROUPPOLICYENFORCEMENTTIER_DEFAULT
	this.EnforcementTier = &enforcementTier
	var policyType OrgGroupPolicyPolicyType = ORGGROUPPOLICYPOLICYTYPE_ORG_CONFIG
	this.PolicyType = &policyType
	return &this
}

// GetContent returns the Content field value.
func (o *OrgGroupPolicyCreateAttributes) GetContent() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyCreateAttributes) GetContentOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *OrgGroupPolicyCreateAttributes) SetContent(v map[string]interface{}) {
	o.Content = v
}

// GetEnforcementTier returns the EnforcementTier field value if set, zero value otherwise.
func (o *OrgGroupPolicyCreateAttributes) GetEnforcementTier() OrgGroupPolicyEnforcementTier {
	if o == nil || o.EnforcementTier == nil {
		var ret OrgGroupPolicyEnforcementTier
		return ret
	}
	return *o.EnforcementTier
}

// GetEnforcementTierOk returns a tuple with the EnforcementTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyCreateAttributes) GetEnforcementTierOk() (*OrgGroupPolicyEnforcementTier, bool) {
	if o == nil || o.EnforcementTier == nil {
		return nil, false
	}
	return o.EnforcementTier, true
}

// HasEnforcementTier returns a boolean if a field has been set.
func (o *OrgGroupPolicyCreateAttributes) HasEnforcementTier() bool {
	return o != nil && o.EnforcementTier != nil
}

// SetEnforcementTier gets a reference to the given OrgGroupPolicyEnforcementTier and assigns it to the EnforcementTier field.
func (o *OrgGroupPolicyCreateAttributes) SetEnforcementTier(v OrgGroupPolicyEnforcementTier) {
	o.EnforcementTier = &v
}

// GetPolicyName returns the PolicyName field value.
func (o *OrgGroupPolicyCreateAttributes) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyCreateAttributes) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value.
func (o *OrgGroupPolicyCreateAttributes) SetPolicyName(v string) {
	o.PolicyName = v
}

// GetPolicyType returns the PolicyType field value if set, zero value otherwise.
func (o *OrgGroupPolicyCreateAttributes) GetPolicyType() OrgGroupPolicyPolicyType {
	if o == nil || o.PolicyType == nil {
		var ret OrgGroupPolicyPolicyType
		return ret
	}
	return *o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyCreateAttributes) GetPolicyTypeOk() (*OrgGroupPolicyPolicyType, bool) {
	if o == nil || o.PolicyType == nil {
		return nil, false
	}
	return o.PolicyType, true
}

// HasPolicyType returns a boolean if a field has been set.
func (o *OrgGroupPolicyCreateAttributes) HasPolicyType() bool {
	return o != nil && o.PolicyType != nil
}

// SetPolicyType gets a reference to the given OrgGroupPolicyPolicyType and assigns it to the PolicyType field.
func (o *OrgGroupPolicyCreateAttributes) SetPolicyType(v OrgGroupPolicyPolicyType) {
	o.PolicyType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["content"] = o.Content
	if o.EnforcementTier != nil {
		toSerialize["enforcement_tier"] = o.EnforcementTier
	}
	toSerialize["policy_name"] = o.PolicyName
	if o.PolicyType != nil {
		toSerialize["policy_type"] = o.PolicyType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content         *map[string]interface{}        `json:"content"`
		EnforcementTier *OrgGroupPolicyEnforcementTier `json:"enforcement_tier,omitempty"`
		PolicyName      *string                        `json:"policy_name"`
		PolicyType      *OrgGroupPolicyPolicyType      `json:"policy_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.PolicyName == nil {
		return fmt.Errorf("required field policy_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "enforcement_tier", "policy_name", "policy_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Content = *all.Content
	if all.EnforcementTier != nil && !all.EnforcementTier.IsValid() {
		hasInvalidField = true
	} else {
		o.EnforcementTier = all.EnforcementTier
	}
	o.PolicyName = *all.PolicyName
	if all.PolicyType != nil && !all.PolicyType.IsValid() {
		hasInvalidField = true
	} else {
		o.PolicyType = all.PolicyType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
