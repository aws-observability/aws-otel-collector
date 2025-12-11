// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleUpdateAttributes Update an existing Cloud Workload Security Agent rule
type CloudWorkloadSecurityAgentRuleUpdateAttributes struct {
	// The array of actions the rule can perform if triggered
	Actions []CloudWorkloadSecurityAgentRuleAction `json:"actions,omitempty"`
	// Constrain the rule to specific versions of the Datadog Agent
	AgentVersion *string `json:"agent_version,omitempty"`
	// The blocking policies that the rule belongs to
	Blocking []string `json:"blocking,omitempty"`
	// The description of the Agent rule
	Description *string `json:"description,omitempty"`
	// The disabled policies that the rule belongs to
	Disabled []string `json:"disabled,omitempty"`
	// Whether the Agent rule is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The SECL expression of the Agent rule
	Expression *string `json:"expression,omitempty"`
	// The monitoring policies that the rule belongs to
	Monitoring []string `json:"monitoring,omitempty"`
	// The ID of the policy where the Agent rule is saved
	PolicyId *string `json:"policy_id,omitempty"`
	// The list of product tags associated with the rule
	ProductTags []string `json:"product_tags,omitempty"`
	// Whether the rule is silent.
	Silent *bool `json:"silent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleUpdateAttributes instantiates a new CloudWorkloadSecurityAgentRuleUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleUpdateAttributes() *CloudWorkloadSecurityAgentRuleUpdateAttributes {
	this := CloudWorkloadSecurityAgentRuleUpdateAttributes{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleUpdateAttributesWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleUpdateAttributesWithDefaults() *CloudWorkloadSecurityAgentRuleUpdateAttributes {
	this := CloudWorkloadSecurityAgentRuleUpdateAttributes{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetActions() []CloudWorkloadSecurityAgentRuleAction {
	if o == nil {
		var ret []CloudWorkloadSecurityAgentRuleAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetActionsOk() (*[]CloudWorkloadSecurityAgentRuleAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return &o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given []CloudWorkloadSecurityAgentRuleAction and assigns it to the Actions field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetActions(v []CloudWorkloadSecurityAgentRuleAction) {
	o.Actions = v
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetAgentVersion() string {
	if o == nil || o.AgentVersion == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetAgentVersionOk() (*string, bool) {
	if o == nil || o.AgentVersion == nil {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasAgentVersion() bool {
	return o != nil && o.AgentVersion != nil
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetBlocking returns the Blocking field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetBlocking() []string {
	if o == nil || o.Blocking == nil {
		var ret []string
		return ret
	}
	return o.Blocking
}

// GetBlockingOk returns a tuple with the Blocking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetBlockingOk() (*[]string, bool) {
	if o == nil || o.Blocking == nil {
		return nil, false
	}
	return &o.Blocking, true
}

// HasBlocking returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasBlocking() bool {
	return o != nil && o.Blocking != nil
}

// SetBlocking gets a reference to the given []string and assigns it to the Blocking field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetBlocking(v []string) {
	o.Blocking = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetDisabled() []string {
	if o == nil || o.Disabled == nil {
		var ret []string
		return ret
	}
	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetDisabledOk() (*[]string, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given []string and assigns it to the Disabled field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetDisabled(v []string) {
	o.Disabled = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasExpression() bool {
	return o != nil && o.Expression != nil
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetExpression(v string) {
	o.Expression = &v
}

// GetMonitoring returns the Monitoring field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetMonitoring() []string {
	if o == nil || o.Monitoring == nil {
		var ret []string
		return ret
	}
	return o.Monitoring
}

// GetMonitoringOk returns a tuple with the Monitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetMonitoringOk() (*[]string, bool) {
	if o == nil || o.Monitoring == nil {
		return nil, false
	}
	return &o.Monitoring, true
}

// HasMonitoring returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasMonitoring() bool {
	return o != nil && o.Monitoring != nil
}

// SetMonitoring gets a reference to the given []string and assigns it to the Monitoring field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetMonitoring(v []string) {
	o.Monitoring = v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasPolicyId() bool {
	return o != nil && o.PolicyId != nil
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetProductTags returns the ProductTags field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetProductTags() []string {
	if o == nil || o.ProductTags == nil {
		var ret []string
		return ret
	}
	return o.ProductTags
}

// GetProductTagsOk returns a tuple with the ProductTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetProductTagsOk() (*[]string, bool) {
	if o == nil || o.ProductTags == nil {
		return nil, false
	}
	return &o.ProductTags, true
}

// HasProductTags returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasProductTags() bool {
	return o != nil && o.ProductTags != nil
}

// SetProductTags gets a reference to the given []string and assigns it to the ProductTags field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetProductTags(v []string) {
	o.ProductTags = v
}

// GetSilent returns the Silent field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetSilent() bool {
	if o == nil || o.Silent == nil {
		var ret bool
		return ret
	}
	return *o.Silent
}

// GetSilentOk returns a tuple with the Silent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetSilentOk() (*bool, bool) {
	if o == nil || o.Silent == nil {
		return nil, false
	}
	return o.Silent, true
}

// HasSilent returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasSilent() bool {
	return o != nil && o.Silent != nil
}

// SetSilent gets a reference to the given bool and assigns it to the Silent field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetSilent(v bool) {
	o.Silent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.AgentVersion != nil {
		toSerialize["agent_version"] = o.AgentVersion
	}
	if o.Blocking != nil {
		toSerialize["blocking"] = o.Blocking
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.Monitoring != nil {
		toSerialize["monitoring"] = o.Monitoring
	}
	if o.PolicyId != nil {
		toSerialize["policy_id"] = o.PolicyId
	}
	if o.ProductTags != nil {
		toSerialize["product_tags"] = o.ProductTags
	}
	if o.Silent != nil {
		toSerialize["silent"] = o.Silent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actions      []CloudWorkloadSecurityAgentRuleAction `json:"actions,omitempty"`
		AgentVersion *string                                `json:"agent_version,omitempty"`
		Blocking     []string                               `json:"blocking,omitempty"`
		Description  *string                                `json:"description,omitempty"`
		Disabled     []string                               `json:"disabled,omitempty"`
		Enabled      *bool                                  `json:"enabled,omitempty"`
		Expression   *string                                `json:"expression,omitempty"`
		Monitoring   []string                               `json:"monitoring,omitempty"`
		PolicyId     *string                                `json:"policy_id,omitempty"`
		ProductTags  []string                               `json:"product_tags,omitempty"`
		Silent       *bool                                  `json:"silent,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actions", "agent_version", "blocking", "description", "disabled", "enabled", "expression", "monitoring", "policy_id", "product_tags", "silent"})
	} else {
		return err
	}
	o.Actions = all.Actions
	o.AgentVersion = all.AgentVersion
	o.Blocking = all.Blocking
	o.Description = all.Description
	o.Disabled = all.Disabled
	o.Enabled = all.Enabled
	o.Expression = all.Expression
	o.Monitoring = all.Monitoring
	o.PolicyId = all.PolicyId
	o.ProductTags = all.ProductTags
	o.Silent = all.Silent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
