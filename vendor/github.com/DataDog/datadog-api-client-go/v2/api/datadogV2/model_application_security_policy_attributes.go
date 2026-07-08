// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityPolicyAttributes A WAF policy.
type ApplicationSecurityPolicyAttributes struct {
	// Description of the WAF policy.
	Description string `json:"description"`
	// Make this policy the default policy. The default policy is applied to
	// every service not specifically assigned to another policy.
	IsDefault *bool `json:"isDefault,omitempty"`
	// The name of the WAF policy.
	Name string `json:"name"`
	// Presets enabled on this policy.
	ProtectionPresets []string `json:"protectionPresets,omitempty"`
	// Rule overrides applied by the policy.
	Rules []ApplicationSecurityPolicyRuleOverride `json:"rules,omitempty"`
	// Deprecated: Ruleset overrides. Use `protectionPresets` instead.
	// Deprecated
	Rulesets []ApplicationSecurityPolicyRulesetOverride `json:"rulesets,omitempty"`
	// The scope of the WAF policy.
	Scope []ApplicationSecurityPolicyScope `json:"scope,omitempty"`
	// Version of the WAF ruleset maintained by Datadog used by this policy. 0 is the default value.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityPolicyAttributes instantiates a new ApplicationSecurityPolicyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityPolicyAttributes(description string, name string) *ApplicationSecurityPolicyAttributes {
	this := ApplicationSecurityPolicyAttributes{}
	this.Description = description
	this.Name = name
	var version int64 = 0
	this.Version = &version
	return &this
}

// NewApplicationSecurityPolicyAttributesWithDefaults instantiates a new ApplicationSecurityPolicyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityPolicyAttributesWithDefaults() *ApplicationSecurityPolicyAttributes {
	this := ApplicationSecurityPolicyAttributes{}
	var version int64 = 0
	this.Version = &version
	return &this
}

// GetDescription returns the Description field value.
func (o *ApplicationSecurityPolicyAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ApplicationSecurityPolicyAttributes) SetDescription(v string) {
	o.Description = v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ApplicationSecurityPolicyAttributes) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyAttributes) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyAttributes) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ApplicationSecurityPolicyAttributes) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetName returns the Name field value.
func (o *ApplicationSecurityPolicyAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ApplicationSecurityPolicyAttributes) SetName(v string) {
	o.Name = v
}

// GetProtectionPresets returns the ProtectionPresets field value if set, zero value otherwise.
func (o *ApplicationSecurityPolicyAttributes) GetProtectionPresets() []string {
	if o == nil || o.ProtectionPresets == nil {
		var ret []string
		return ret
	}
	return o.ProtectionPresets
}

// GetProtectionPresetsOk returns a tuple with the ProtectionPresets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyAttributes) GetProtectionPresetsOk() (*[]string, bool) {
	if o == nil || o.ProtectionPresets == nil {
		return nil, false
	}
	return &o.ProtectionPresets, true
}

// HasProtectionPresets returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyAttributes) HasProtectionPresets() bool {
	return o != nil && o.ProtectionPresets != nil
}

// SetProtectionPresets gets a reference to the given []string and assigns it to the ProtectionPresets field.
func (o *ApplicationSecurityPolicyAttributes) SetProtectionPresets(v []string) {
	o.ProtectionPresets = v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ApplicationSecurityPolicyAttributes) GetRules() []ApplicationSecurityPolicyRuleOverride {
	if o == nil || o.Rules == nil {
		var ret []ApplicationSecurityPolicyRuleOverride
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyAttributes) GetRulesOk() (*[]ApplicationSecurityPolicyRuleOverride, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyAttributes) HasRules() bool {
	return o != nil && o.Rules != nil
}

// SetRules gets a reference to the given []ApplicationSecurityPolicyRuleOverride and assigns it to the Rules field.
func (o *ApplicationSecurityPolicyAttributes) SetRules(v []ApplicationSecurityPolicyRuleOverride) {
	o.Rules = v
}

// GetRulesets returns the Rulesets field value if set, zero value otherwise.
// Deprecated
func (o *ApplicationSecurityPolicyAttributes) GetRulesets() []ApplicationSecurityPolicyRulesetOverride {
	if o == nil || o.Rulesets == nil {
		var ret []ApplicationSecurityPolicyRulesetOverride
		return ret
	}
	return o.Rulesets
}

// GetRulesetsOk returns a tuple with the Rulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplicationSecurityPolicyAttributes) GetRulesetsOk() (*[]ApplicationSecurityPolicyRulesetOverride, bool) {
	if o == nil || o.Rulesets == nil {
		return nil, false
	}
	return &o.Rulesets, true
}

// HasRulesets returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyAttributes) HasRulesets() bool {
	return o != nil && o.Rulesets != nil
}

// SetRulesets gets a reference to the given []ApplicationSecurityPolicyRulesetOverride and assigns it to the Rulesets field.
// Deprecated
func (o *ApplicationSecurityPolicyAttributes) SetRulesets(v []ApplicationSecurityPolicyRulesetOverride) {
	o.Rulesets = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ApplicationSecurityPolicyAttributes) GetScope() []ApplicationSecurityPolicyScope {
	if o == nil || o.Scope == nil {
		var ret []ApplicationSecurityPolicyScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyAttributes) GetScopeOk() (*[]ApplicationSecurityPolicyScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given []ApplicationSecurityPolicyScope and assigns it to the Scope field.
func (o *ApplicationSecurityPolicyAttributes) SetScope(v []ApplicationSecurityPolicyScope) {
	o.Scope = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplicationSecurityPolicyAttributes) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyAttributes) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *ApplicationSecurityPolicyAttributes) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityPolicyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	toSerialize["name"] = o.Name
	if o.ProtectionPresets != nil {
		toSerialize["protectionPresets"] = o.ProtectionPresets
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Rulesets != nil {
		toSerialize["rulesets"] = o.Rulesets
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityPolicyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description       *string                                    `json:"description"`
		IsDefault         *bool                                      `json:"isDefault,omitempty"`
		Name              *string                                    `json:"name"`
		ProtectionPresets []string                                   `json:"protectionPresets,omitempty"`
		Rules             []ApplicationSecurityPolicyRuleOverride    `json:"rules,omitempty"`
		Rulesets          []ApplicationSecurityPolicyRulesetOverride `json:"rulesets,omitempty"`
		Scope             []ApplicationSecurityPolicyScope           `json:"scope,omitempty"`
		Version           *int64                                     `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "isDefault", "name", "protectionPresets", "rules", "rulesets", "scope", "version"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.IsDefault = all.IsDefault
	o.Name = *all.Name
	o.ProtectionPresets = all.ProtectionPresets
	o.Rules = all.Rules
	o.Rulesets = all.Rulesets
	o.Scope = all.Scope
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
