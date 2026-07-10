// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityPolicyUpdateAttributes Update a WAF policy.
type ApplicationSecurityPolicyUpdateAttributes struct {
	// Description of the WAF policy.
	Description string `json:"description"`
	// Make this policy the default policy. The default policy is applied to
	// every service not specifically assigned to another policy.
	IsDefault bool `json:"isDefault"`
	// The name of the WAF policy.
	Name string `json:"name"`
	// Presets enabled on this policy.
	ProtectionPresets []string `json:"protectionPresets"`
	// Rule overrides applied by the policy.
	Rules []ApplicationSecurityPolicyRuleOverride `json:"rules"`
	// Deprecated: Ruleset overrides. Use `protectionPresets` instead.
	// Deprecated
	Rulesets []ApplicationSecurityPolicyRulesetOverride `json:"rulesets,omitempty"`
	// The scope of the WAF policy.
	Scope []ApplicationSecurityPolicyScope `json:"scope"`
	// Version of the WAF ruleset maintained by Datadog used by this policy. 0 is the default value.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityPolicyUpdateAttributes instantiates a new ApplicationSecurityPolicyUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityPolicyUpdateAttributes(description string, isDefault bool, name string, protectionPresets []string, rules []ApplicationSecurityPolicyRuleOverride, scope []ApplicationSecurityPolicyScope, version int64) *ApplicationSecurityPolicyUpdateAttributes {
	this := ApplicationSecurityPolicyUpdateAttributes{}
	this.Description = description
	this.IsDefault = isDefault
	this.Name = name
	this.ProtectionPresets = protectionPresets
	this.Rules = rules
	this.Scope = scope
	this.Version = version
	return &this
}

// NewApplicationSecurityPolicyUpdateAttributesWithDefaults instantiates a new ApplicationSecurityPolicyUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityPolicyUpdateAttributesWithDefaults() *ApplicationSecurityPolicyUpdateAttributes {
	this := ApplicationSecurityPolicyUpdateAttributes{}
	var version int64 = 0
	this.Version = version
	return &this
}

// GetDescription returns the Description field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) SetDescription(v string) {
	o.Description = v
}

// GetIsDefault returns the IsDefault field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetName returns the Name field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) SetName(v string) {
	o.Name = v
}

// GetProtectionPresets returns the ProtectionPresets field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetProtectionPresets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProtectionPresets
}

// GetProtectionPresetsOk returns a tuple with the ProtectionPresets field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetProtectionPresetsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtectionPresets, true
}

// SetProtectionPresets sets field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) SetProtectionPresets(v []string) {
	o.ProtectionPresets = v
}

// GetRules returns the Rules field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetRules() []ApplicationSecurityPolicyRuleOverride {
	if o == nil {
		var ret []ApplicationSecurityPolicyRuleOverride
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetRulesOk() (*[]ApplicationSecurityPolicyRuleOverride, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) SetRules(v []ApplicationSecurityPolicyRuleOverride) {
	o.Rules = v
}

// GetRulesets returns the Rulesets field value if set, zero value otherwise.
// Deprecated
func (o *ApplicationSecurityPolicyUpdateAttributes) GetRulesets() []ApplicationSecurityPolicyRulesetOverride {
	if o == nil || o.Rulesets == nil {
		var ret []ApplicationSecurityPolicyRulesetOverride
		return ret
	}
	return o.Rulesets
}

// GetRulesetsOk returns a tuple with the Rulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplicationSecurityPolicyUpdateAttributes) GetRulesetsOk() (*[]ApplicationSecurityPolicyRulesetOverride, bool) {
	if o == nil || o.Rulesets == nil {
		return nil, false
	}
	return &o.Rulesets, true
}

// HasRulesets returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyUpdateAttributes) HasRulesets() bool {
	return o != nil && o.Rulesets != nil
}

// SetRulesets gets a reference to the given []ApplicationSecurityPolicyRulesetOverride and assigns it to the Rulesets field.
// Deprecated
func (o *ApplicationSecurityPolicyUpdateAttributes) SetRulesets(v []ApplicationSecurityPolicyRulesetOverride) {
	o.Rulesets = v
}

// GetScope returns the Scope field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetScope() []ApplicationSecurityPolicyScope {
	if o == nil {
		var ret []ApplicationSecurityPolicyScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetScopeOk() (*[]ApplicationSecurityPolicyScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) SetScope(v []ApplicationSecurityPolicyScope) {
	o.Scope = v
}

// GetVersion returns the Version field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyUpdateAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *ApplicationSecurityPolicyUpdateAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityPolicyUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["isDefault"] = o.IsDefault
	toSerialize["name"] = o.Name
	toSerialize["protectionPresets"] = o.ProtectionPresets
	toSerialize["rules"] = o.Rules
	if o.Rulesets != nil {
		toSerialize["rulesets"] = o.Rulesets
	}
	toSerialize["scope"] = o.Scope
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityPolicyUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description       *string                                    `json:"description"`
		IsDefault         *bool                                      `json:"isDefault"`
		Name              *string                                    `json:"name"`
		ProtectionPresets *[]string                                  `json:"protectionPresets"`
		Rules             *[]ApplicationSecurityPolicyRuleOverride   `json:"rules"`
		Rulesets          []ApplicationSecurityPolicyRulesetOverride `json:"rulesets,omitempty"`
		Scope             *[]ApplicationSecurityPolicyScope          `json:"scope"`
		Version           *int64                                     `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.IsDefault == nil {
		return fmt.Errorf("required field isDefault missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ProtectionPresets == nil {
		return fmt.Errorf("required field protectionPresets missing")
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "isDefault", "name", "protectionPresets", "rules", "rulesets", "scope", "version"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.IsDefault = *all.IsDefault
	o.Name = *all.Name
	o.ProtectionPresets = *all.ProtectionPresets
	o.Rules = *all.Rules
	o.Rulesets = all.Rulesets
	o.Scope = *all.Scope
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
