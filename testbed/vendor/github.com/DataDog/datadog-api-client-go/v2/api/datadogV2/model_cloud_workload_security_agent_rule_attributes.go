// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleAttributes A Cloud Workload Security Agent rule returned by the API
type CloudWorkloadSecurityAgentRuleAttributes struct {
	// The array of actions the rule can perform if triggered
	Actions []CloudWorkloadSecurityAgentRuleAction `json:"actions,omitempty"`
	// The version of the Agent
	AgentConstraint *string `json:"agentConstraint,omitempty"`
	// The blocking policies that the rule belongs to
	Blocking []string `json:"blocking,omitempty"`
	// The category of the Agent rule
	Category *string `json:"category,omitempty"`
	// The ID of the user who created the rule
	CreationAuthorUuId *string `json:"creationAuthorUuId,omitempty"`
	// When the Agent rule was created, timestamp in milliseconds
	CreationDate *int64 `json:"creationDate,omitempty"`
	// The attributes of the user who created the Agent rule
	Creator *CloudWorkloadSecurityAgentRuleCreatorAttributes `json:"creator,omitempty"`
	// Whether the rule is included by default
	DefaultRule *bool `json:"defaultRule,omitempty"`
	// The description of the Agent rule
	Description *string `json:"description,omitempty"`
	// The disabled policies that the rule belongs to
	Disabled []string `json:"disabled,omitempty"`
	// Whether the Agent rule is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The SECL expression of the Agent rule
	Expression *string `json:"expression,omitempty"`
	// The platforms the Agent rule is supported on
	Filters []string `json:"filters,omitempty"`
	// The monitoring policies that the rule belongs to
	Monitoring []string `json:"monitoring,omitempty"`
	// The name of the Agent rule
	Name *string `json:"name,omitempty"`
	// The list of product tags associated with the rule
	ProductTags []string `json:"product_tags,omitempty"`
	// Whether the rule is silent.
	Silent *bool `json:"silent,omitempty"`
	// The ID of the user who updated the rule
	UpdateAuthorUuId *string `json:"updateAuthorUuId,omitempty"`
	// Timestamp in milliseconds when the Agent rule was last updated
	UpdateDate *int64 `json:"updateDate,omitempty"`
	// When the Agent rule was last updated, timestamp in milliseconds
	UpdatedAt *int64 `json:"updatedAt,omitempty"`
	// The attributes of the user who last updated the Agent rule
	Updater *CloudWorkloadSecurityAgentRuleUpdaterAttributes `json:"updater,omitempty"`
	// The version of the Agent rule
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleAttributes instantiates a new CloudWorkloadSecurityAgentRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleAttributes() *CloudWorkloadSecurityAgentRuleAttributes {
	this := CloudWorkloadSecurityAgentRuleAttributes{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleAttributesWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleAttributesWithDefaults() *CloudWorkloadSecurityAgentRuleAttributes {
	this := CloudWorkloadSecurityAgentRuleAttributes{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetActions() []CloudWorkloadSecurityAgentRuleAction {
	if o == nil {
		var ret []CloudWorkloadSecurityAgentRuleAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetActionsOk() (*[]CloudWorkloadSecurityAgentRuleAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return &o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given []CloudWorkloadSecurityAgentRuleAction and assigns it to the Actions field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetActions(v []CloudWorkloadSecurityAgentRuleAction) {
	o.Actions = v
}

// GetAgentConstraint returns the AgentConstraint field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetAgentConstraint() string {
	if o == nil || o.AgentConstraint == nil {
		var ret string
		return ret
	}
	return *o.AgentConstraint
}

// GetAgentConstraintOk returns a tuple with the AgentConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetAgentConstraintOk() (*string, bool) {
	if o == nil || o.AgentConstraint == nil {
		return nil, false
	}
	return o.AgentConstraint, true
}

// HasAgentConstraint returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasAgentConstraint() bool {
	return o != nil && o.AgentConstraint != nil
}

// SetAgentConstraint gets a reference to the given string and assigns it to the AgentConstraint field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetAgentConstraint(v string) {
	o.AgentConstraint = &v
}

// GetBlocking returns the Blocking field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetBlocking() []string {
	if o == nil || o.Blocking == nil {
		var ret []string
		return ret
	}
	return o.Blocking
}

// GetBlockingOk returns a tuple with the Blocking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetBlockingOk() (*[]string, bool) {
	if o == nil || o.Blocking == nil {
		return nil, false
	}
	return &o.Blocking, true
}

// HasBlocking returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasBlocking() bool {
	return o != nil && o.Blocking != nil
}

// SetBlocking gets a reference to the given []string and assigns it to the Blocking field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetBlocking(v []string) {
	o.Blocking = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetCategory(v string) {
	o.Category = &v
}

// GetCreationAuthorUuId returns the CreationAuthorUuId field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCreationAuthorUuId() string {
	if o == nil || o.CreationAuthorUuId == nil {
		var ret string
		return ret
	}
	return *o.CreationAuthorUuId
}

// GetCreationAuthorUuIdOk returns a tuple with the CreationAuthorUuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCreationAuthorUuIdOk() (*string, bool) {
	if o == nil || o.CreationAuthorUuId == nil {
		return nil, false
	}
	return o.CreationAuthorUuId, true
}

// HasCreationAuthorUuId returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasCreationAuthorUuId() bool {
	return o != nil && o.CreationAuthorUuId != nil
}

// SetCreationAuthorUuId gets a reference to the given string and assigns it to the CreationAuthorUuId field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetCreationAuthorUuId(v string) {
	o.CreationAuthorUuId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCreationDate() int64 {
	if o == nil || o.CreationDate == nil {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCreationDateOk() (*int64, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasCreationDate() bool {
	return o != nil && o.CreationDate != nil
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCreator() CloudWorkloadSecurityAgentRuleCreatorAttributes {
	if o == nil || o.Creator == nil {
		var ret CloudWorkloadSecurityAgentRuleCreatorAttributes
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCreatorOk() (*CloudWorkloadSecurityAgentRuleCreatorAttributes, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given CloudWorkloadSecurityAgentRuleCreatorAttributes and assigns it to the Creator field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetCreator(v CloudWorkloadSecurityAgentRuleCreatorAttributes) {
	o.Creator = &v
}

// GetDefaultRule returns the DefaultRule field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetDefaultRule() bool {
	if o == nil || o.DefaultRule == nil {
		var ret bool
		return ret
	}
	return *o.DefaultRule
}

// GetDefaultRuleOk returns a tuple with the DefaultRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetDefaultRuleOk() (*bool, bool) {
	if o == nil || o.DefaultRule == nil {
		return nil, false
	}
	return o.DefaultRule, true
}

// HasDefaultRule returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasDefaultRule() bool {
	return o != nil && o.DefaultRule != nil
}

// SetDefaultRule gets a reference to the given bool and assigns it to the DefaultRule field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetDefaultRule(v bool) {
	o.DefaultRule = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetDisabled() []string {
	if o == nil || o.Disabled == nil {
		var ret []string
		return ret
	}
	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetDisabledOk() (*[]string, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given []string and assigns it to the Disabled field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetDisabled(v []string) {
	o.Disabled = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasExpression() bool {
	return o != nil && o.Expression != nil
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetExpression(v string) {
	o.Expression = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetFilters() []string {
	if o == nil || o.Filters == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetFiltersOk() (*[]string, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given []string and assigns it to the Filters field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetFilters(v []string) {
	o.Filters = v
}

// GetMonitoring returns the Monitoring field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetMonitoring() []string {
	if o == nil || o.Monitoring == nil {
		var ret []string
		return ret
	}
	return o.Monitoring
}

// GetMonitoringOk returns a tuple with the Monitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetMonitoringOk() (*[]string, bool) {
	if o == nil || o.Monitoring == nil {
		return nil, false
	}
	return &o.Monitoring, true
}

// HasMonitoring returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasMonitoring() bool {
	return o != nil && o.Monitoring != nil
}

// SetMonitoring gets a reference to the given []string and assigns it to the Monitoring field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetMonitoring(v []string) {
	o.Monitoring = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetName(v string) {
	o.Name = &v
}

// GetProductTags returns the ProductTags field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetProductTags() []string {
	if o == nil || o.ProductTags == nil {
		var ret []string
		return ret
	}
	return o.ProductTags
}

// GetProductTagsOk returns a tuple with the ProductTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetProductTagsOk() (*[]string, bool) {
	if o == nil || o.ProductTags == nil {
		return nil, false
	}
	return &o.ProductTags, true
}

// HasProductTags returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasProductTags() bool {
	return o != nil && o.ProductTags != nil
}

// SetProductTags gets a reference to the given []string and assigns it to the ProductTags field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetProductTags(v []string) {
	o.ProductTags = v
}

// GetSilent returns the Silent field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetSilent() bool {
	if o == nil || o.Silent == nil {
		var ret bool
		return ret
	}
	return *o.Silent
}

// GetSilentOk returns a tuple with the Silent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetSilentOk() (*bool, bool) {
	if o == nil || o.Silent == nil {
		return nil, false
	}
	return o.Silent, true
}

// HasSilent returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasSilent() bool {
	return o != nil && o.Silent != nil
}

// SetSilent gets a reference to the given bool and assigns it to the Silent field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetSilent(v bool) {
	o.Silent = &v
}

// GetUpdateAuthorUuId returns the UpdateAuthorUuId field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdateAuthorUuId() string {
	if o == nil || o.UpdateAuthorUuId == nil {
		var ret string
		return ret
	}
	return *o.UpdateAuthorUuId
}

// GetUpdateAuthorUuIdOk returns a tuple with the UpdateAuthorUuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdateAuthorUuIdOk() (*string, bool) {
	if o == nil || o.UpdateAuthorUuId == nil {
		return nil, false
	}
	return o.UpdateAuthorUuId, true
}

// HasUpdateAuthorUuId returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasUpdateAuthorUuId() bool {
	return o != nil && o.UpdateAuthorUuId != nil
}

// SetUpdateAuthorUuId gets a reference to the given string and assigns it to the UpdateAuthorUuId field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetUpdateAuthorUuId(v string) {
	o.UpdateAuthorUuId = &v
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdateDate() int64 {
	if o == nil || o.UpdateDate == nil {
		var ret int64
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdateDateOk() (*int64, bool) {
	if o == nil || o.UpdateDate == nil {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasUpdateDate() bool {
	return o != nil && o.UpdateDate != nil
}

// SetUpdateDate gets a reference to the given int64 and assigns it to the UpdateDate field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetUpdateDate(v int64) {
	o.UpdateDate = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdater returns the Updater field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdater() CloudWorkloadSecurityAgentRuleUpdaterAttributes {
	if o == nil || o.Updater == nil {
		var ret CloudWorkloadSecurityAgentRuleUpdaterAttributes
		return ret
	}
	return *o.Updater
}

// GetUpdaterOk returns a tuple with the Updater field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdaterOk() (*CloudWorkloadSecurityAgentRuleUpdaterAttributes, bool) {
	if o == nil || o.Updater == nil {
		return nil, false
	}
	return o.Updater, true
}

// HasUpdater returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasUpdater() bool {
	return o != nil && o.Updater != nil
}

// SetUpdater gets a reference to the given CloudWorkloadSecurityAgentRuleUpdaterAttributes and assigns it to the Updater field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetUpdater(v CloudWorkloadSecurityAgentRuleUpdaterAttributes) {
	o.Updater = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *CloudWorkloadSecurityAgentRuleAttributes) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.AgentConstraint != nil {
		toSerialize["agentConstraint"] = o.AgentConstraint
	}
	if o.Blocking != nil {
		toSerialize["blocking"] = o.Blocking
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.CreationAuthorUuId != nil {
		toSerialize["creationAuthorUuId"] = o.CreationAuthorUuId
	}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.DefaultRule != nil {
		toSerialize["defaultRule"] = o.DefaultRule
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
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Monitoring != nil {
		toSerialize["monitoring"] = o.Monitoring
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProductTags != nil {
		toSerialize["product_tags"] = o.ProductTags
	}
	if o.Silent != nil {
		toSerialize["silent"] = o.Silent
	}
	if o.UpdateAuthorUuId != nil {
		toSerialize["updateAuthorUuId"] = o.UpdateAuthorUuId
	}
	if o.UpdateDate != nil {
		toSerialize["updateDate"] = o.UpdateDate
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Updater != nil {
		toSerialize["updater"] = o.Updater
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
func (o *CloudWorkloadSecurityAgentRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actions            []CloudWorkloadSecurityAgentRuleAction           `json:"actions,omitempty"`
		AgentConstraint    *string                                          `json:"agentConstraint,omitempty"`
		Blocking           []string                                         `json:"blocking,omitempty"`
		Category           *string                                          `json:"category,omitempty"`
		CreationAuthorUuId *string                                          `json:"creationAuthorUuId,omitempty"`
		CreationDate       *int64                                           `json:"creationDate,omitempty"`
		Creator            *CloudWorkloadSecurityAgentRuleCreatorAttributes `json:"creator,omitempty"`
		DefaultRule        *bool                                            `json:"defaultRule,omitempty"`
		Description        *string                                          `json:"description,omitempty"`
		Disabled           []string                                         `json:"disabled,omitempty"`
		Enabled            *bool                                            `json:"enabled,omitempty"`
		Expression         *string                                          `json:"expression,omitempty"`
		Filters            []string                                         `json:"filters,omitempty"`
		Monitoring         []string                                         `json:"monitoring,omitempty"`
		Name               *string                                          `json:"name,omitempty"`
		ProductTags        []string                                         `json:"product_tags,omitempty"`
		Silent             *bool                                            `json:"silent,omitempty"`
		UpdateAuthorUuId   *string                                          `json:"updateAuthorUuId,omitempty"`
		UpdateDate         *int64                                           `json:"updateDate,omitempty"`
		UpdatedAt          *int64                                           `json:"updatedAt,omitempty"`
		Updater            *CloudWorkloadSecurityAgentRuleUpdaterAttributes `json:"updater,omitempty"`
		Version            *int64                                           `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actions", "agentConstraint", "blocking", "category", "creationAuthorUuId", "creationDate", "creator", "defaultRule", "description", "disabled", "enabled", "expression", "filters", "monitoring", "name", "product_tags", "silent", "updateAuthorUuId", "updateDate", "updatedAt", "updater", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Actions = all.Actions
	o.AgentConstraint = all.AgentConstraint
	o.Blocking = all.Blocking
	o.Category = all.Category
	o.CreationAuthorUuId = all.CreationAuthorUuId
	o.CreationDate = all.CreationDate
	if all.Creator != nil && all.Creator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Creator = all.Creator
	o.DefaultRule = all.DefaultRule
	o.Description = all.Description
	o.Disabled = all.Disabled
	o.Enabled = all.Enabled
	o.Expression = all.Expression
	o.Filters = all.Filters
	o.Monitoring = all.Monitoring
	o.Name = all.Name
	o.ProductTags = all.ProductTags
	o.Silent = all.Silent
	o.UpdateAuthorUuId = all.UpdateAuthorUuId
	o.UpdateDate = all.UpdateDate
	o.UpdatedAt = all.UpdatedAt
	if all.Updater != nil && all.Updater.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Updater = all.Updater
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
