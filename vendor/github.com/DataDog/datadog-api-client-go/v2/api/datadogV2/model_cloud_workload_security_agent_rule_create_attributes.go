// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleCreateAttributes Create a new Cloud Workload Security Agent rule.
type CloudWorkloadSecurityAgentRuleCreateAttributes struct {
	// The array of actions the rule can perform if triggered
	Actions []CloudWorkloadSecurityAgentRuleAction `json:"actions,omitempty"`
	// Constrain the rule to specific versions of the Datadog Agent.
	AgentVersion *string `json:"agent_version,omitempty"`
	// The blocking policies that the rule belongs to.
	Blocking []string `json:"blocking,omitempty"`
	// The description of the Agent rule.
	Description *string `json:"description,omitempty"`
	// The disabled policies that the rule belongs to.
	Disabled []string `json:"disabled,omitempty"`
	// Whether the Agent rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The SECL expression of the Agent rule.
	Expression string `json:"expression"`
	// The platforms the Agent rule is supported on.
	Filters []string `json:"filters,omitempty"`
	// The monitoring policies that the rule belongs to.
	Monitoring []string `json:"monitoring,omitempty"`
	// The name of the Agent rule.
	Name string `json:"name"`
	// The ID of the policy where the Agent rule is saved.
	PolicyId *string `json:"policy_id,omitempty"`
	// The list of product tags associated with the rule.
	ProductTags []string `json:"product_tags,omitempty"`
	// Whether the rule is silent.
	Silent *bool `json:"silent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleCreateAttributes instantiates a new CloudWorkloadSecurityAgentRuleCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleCreateAttributes(expression string, name string) *CloudWorkloadSecurityAgentRuleCreateAttributes {
	this := CloudWorkloadSecurityAgentRuleCreateAttributes{}
	this.Expression = expression
	this.Name = name
	return &this
}

// NewCloudWorkloadSecurityAgentRuleCreateAttributesWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleCreateAttributesWithDefaults() *CloudWorkloadSecurityAgentRuleCreateAttributes {
	this := CloudWorkloadSecurityAgentRuleCreateAttributes{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetActions() []CloudWorkloadSecurityAgentRuleAction {
	if o == nil {
		var ret []CloudWorkloadSecurityAgentRuleAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetActionsOk() (*[]CloudWorkloadSecurityAgentRuleAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return &o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given []CloudWorkloadSecurityAgentRuleAction and assigns it to the Actions field.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetActions(v []CloudWorkloadSecurityAgentRuleAction) {
	o.Actions = v
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetAgentVersion() string {
	if o == nil || o.AgentVersion == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetAgentVersionOk() (*string, bool) {
	if o == nil || o.AgentVersion == nil {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasAgentVersion() bool {
	return o != nil && o.AgentVersion != nil
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetBlocking returns the Blocking field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetBlocking() []string {
	if o == nil || o.Blocking == nil {
		var ret []string
		return ret
	}
	return o.Blocking
}

// GetBlockingOk returns a tuple with the Blocking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetBlockingOk() (*[]string, bool) {
	if o == nil || o.Blocking == nil {
		return nil, false
	}
	return &o.Blocking, true
}

// HasBlocking returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasBlocking() bool {
	return o != nil && o.Blocking != nil
}

// SetBlocking gets a reference to the given []string and assigns it to the Blocking field.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetBlocking(v []string) {
	o.Blocking = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetDisabled() []string {
	if o == nil || o.Disabled == nil {
		var ret []string
		return ret
	}
	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetDisabledOk() (*[]string, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given []string and assigns it to the Disabled field.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetDisabled(v []string) {
	o.Disabled = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExpression returns the Expression field value.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetExpression(v string) {
	o.Expression = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetFilters() []string {
	if o == nil || o.Filters == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetFiltersOk() (*[]string, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given []string and assigns it to the Filters field.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetFilters(v []string) {
	o.Filters = v
}

// GetMonitoring returns the Monitoring field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetMonitoring() []string {
	if o == nil || o.Monitoring == nil {
		var ret []string
		return ret
	}
	return o.Monitoring
}

// GetMonitoringOk returns a tuple with the Monitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetMonitoringOk() (*[]string, bool) {
	if o == nil || o.Monitoring == nil {
		return nil, false
	}
	return &o.Monitoring, true
}

// HasMonitoring returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasMonitoring() bool {
	return o != nil && o.Monitoring != nil
}

// SetMonitoring gets a reference to the given []string and assigns it to the Monitoring field.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetMonitoring(v []string) {
	o.Monitoring = v
}

// GetName returns the Name field value.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasPolicyId() bool {
	return o != nil && o.PolicyId != nil
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetProductTags returns the ProductTags field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetProductTags() []string {
	if o == nil || o.ProductTags == nil {
		var ret []string
		return ret
	}
	return o.ProductTags
}

// GetProductTagsOk returns a tuple with the ProductTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetProductTagsOk() (*[]string, bool) {
	if o == nil || o.ProductTags == nil {
		return nil, false
	}
	return &o.ProductTags, true
}

// HasProductTags returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasProductTags() bool {
	return o != nil && o.ProductTags != nil
}

// SetProductTags gets a reference to the given []string and assigns it to the ProductTags field.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetProductTags(v []string) {
	o.ProductTags = v
}

// GetSilent returns the Silent field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetSilent() bool {
	if o == nil || o.Silent == nil {
		var ret bool
		return ret
	}
	return *o.Silent
}

// GetSilentOk returns a tuple with the Silent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetSilentOk() (*bool, bool) {
	if o == nil || o.Silent == nil {
		return nil, false
	}
	return o.Silent, true
}

// HasSilent returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasSilent() bool {
	return o != nil && o.Silent != nil
}

// SetSilent gets a reference to the given bool and assigns it to the Silent field.
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetSilent(v bool) {
	o.Silent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleCreateAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["expression"] = o.Expression
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Monitoring != nil {
		toSerialize["monitoring"] = o.Monitoring
	}
	toSerialize["name"] = o.Name
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
func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actions      []CloudWorkloadSecurityAgentRuleAction `json:"actions,omitempty"`
		AgentVersion *string                                `json:"agent_version,omitempty"`
		Blocking     []string                               `json:"blocking,omitempty"`
		Description  *string                                `json:"description,omitempty"`
		Disabled     []string                               `json:"disabled,omitempty"`
		Enabled      *bool                                  `json:"enabled,omitempty"`
		Expression   *string                                `json:"expression"`
		Filters      []string                               `json:"filters,omitempty"`
		Monitoring   []string                               `json:"monitoring,omitempty"`
		Name         *string                                `json:"name"`
		PolicyId     *string                                `json:"policy_id,omitempty"`
		ProductTags  []string                               `json:"product_tags,omitempty"`
		Silent       *bool                                  `json:"silent,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actions", "agent_version", "blocking", "description", "disabled", "enabled", "expression", "filters", "monitoring", "name", "policy_id", "product_tags", "silent"})
	} else {
		return err
	}
	o.Actions = all.Actions
	o.AgentVersion = all.AgentVersion
	o.Blocking = all.Blocking
	o.Description = all.Description
	o.Disabled = all.Disabled
	o.Enabled = all.Enabled
	o.Expression = *all.Expression
	o.Filters = all.Filters
	o.Monitoring = all.Monitoring
	o.Name = *all.Name
	o.PolicyId = all.PolicyId
	o.ProductTags = all.ProductTags
	o.Silent = all.Silent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
