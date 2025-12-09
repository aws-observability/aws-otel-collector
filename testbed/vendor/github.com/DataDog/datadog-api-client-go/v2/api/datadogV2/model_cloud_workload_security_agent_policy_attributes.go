// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentPolicyAttributes A Cloud Workload Security Agent policy returned by the API
type CloudWorkloadSecurityAgentPolicyAttributes struct {
	// The number of rules with the blocking feature in this policy
	BlockingRulesCount *int32 `json:"blockingRulesCount,omitempty"`
	// Whether the policy is managed by Datadog
	DatadogManaged *bool `json:"datadogManaged,omitempty"`
	// The description of the policy
	Description *string `json:"description,omitempty"`
	// The number of rules that are disabled in this policy
	DisabledRulesCount *int32 `json:"disabledRulesCount,omitempty"`
	// Whether the Agent policy is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The host tags defining where this policy is deployed
	HostTags []string `json:"hostTags,omitempty"`
	// The host tags defining where this policy is deployed, the inner values are linked with AND, the outer values are linked with OR
	HostTagsLists [][]string `json:"hostTagsLists,omitempty"`
	// The number of rules in the monitoring state in this policy
	MonitoringRulesCount *int32 `json:"monitoringRulesCount,omitempty"`
	// The name of the policy
	Name *string `json:"name,omitempty"`
	// Whether the policy is pinned
	Pinned *bool `json:"pinned,omitempty"`
	// The version of the policy
	PolicyVersion *string `json:"policyVersion,omitempty"`
	// The priority of the policy
	Priority *int64 `json:"priority,omitempty"`
	// The number of rules in this policy
	RuleCount *int32 `json:"ruleCount,omitempty"`
	// Timestamp in milliseconds when the policy was last updated
	UpdateDate *int64 `json:"updateDate,omitempty"`
	// When the policy was last updated, timestamp in milliseconds
	UpdatedAt *int64 `json:"updatedAt,omitempty"`
	// The attributes of the user who last updated the policy
	Updater *CloudWorkloadSecurityAgentPolicyUpdaterAttributes `json:"updater,omitempty"`
	// The versions of the policy
	Versions []CloudWorkloadSecurityAgentPolicyVersion `json:"versions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentPolicyAttributes instantiates a new CloudWorkloadSecurityAgentPolicyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentPolicyAttributes() *CloudWorkloadSecurityAgentPolicyAttributes {
	this := CloudWorkloadSecurityAgentPolicyAttributes{}
	return &this
}

// NewCloudWorkloadSecurityAgentPolicyAttributesWithDefaults instantiates a new CloudWorkloadSecurityAgentPolicyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentPolicyAttributesWithDefaults() *CloudWorkloadSecurityAgentPolicyAttributes {
	this := CloudWorkloadSecurityAgentPolicyAttributes{}
	return &this
}

// GetBlockingRulesCount returns the BlockingRulesCount field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetBlockingRulesCount() int32 {
	if o == nil || o.BlockingRulesCount == nil {
		var ret int32
		return ret
	}
	return *o.BlockingRulesCount
}

// GetBlockingRulesCountOk returns a tuple with the BlockingRulesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetBlockingRulesCountOk() (*int32, bool) {
	if o == nil || o.BlockingRulesCount == nil {
		return nil, false
	}
	return o.BlockingRulesCount, true
}

// HasBlockingRulesCount returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasBlockingRulesCount() bool {
	return o != nil && o.BlockingRulesCount != nil
}

// SetBlockingRulesCount gets a reference to the given int32 and assigns it to the BlockingRulesCount field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetBlockingRulesCount(v int32) {
	o.BlockingRulesCount = &v
}

// GetDatadogManaged returns the DatadogManaged field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetDatadogManaged() bool {
	if o == nil || o.DatadogManaged == nil {
		var ret bool
		return ret
	}
	return *o.DatadogManaged
}

// GetDatadogManagedOk returns a tuple with the DatadogManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetDatadogManagedOk() (*bool, bool) {
	if o == nil || o.DatadogManaged == nil {
		return nil, false
	}
	return o.DatadogManaged, true
}

// HasDatadogManaged returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasDatadogManaged() bool {
	return o != nil && o.DatadogManaged != nil
}

// SetDatadogManaged gets a reference to the given bool and assigns it to the DatadogManaged field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetDatadogManaged(v bool) {
	o.DatadogManaged = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDisabledRulesCount returns the DisabledRulesCount field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetDisabledRulesCount() int32 {
	if o == nil || o.DisabledRulesCount == nil {
		var ret int32
		return ret
	}
	return *o.DisabledRulesCount
}

// GetDisabledRulesCountOk returns a tuple with the DisabledRulesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetDisabledRulesCountOk() (*int32, bool) {
	if o == nil || o.DisabledRulesCount == nil {
		return nil, false
	}
	return o.DisabledRulesCount, true
}

// HasDisabledRulesCount returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasDisabledRulesCount() bool {
	return o != nil && o.DisabledRulesCount != nil
}

// SetDisabledRulesCount gets a reference to the given int32 and assigns it to the DisabledRulesCount field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetDisabledRulesCount(v int32) {
	o.DisabledRulesCount = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHostTags returns the HostTags field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetHostTags() []string {
	if o == nil || o.HostTags == nil {
		var ret []string
		return ret
	}
	return o.HostTags
}

// GetHostTagsOk returns a tuple with the HostTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetHostTagsOk() (*[]string, bool) {
	if o == nil || o.HostTags == nil {
		return nil, false
	}
	return &o.HostTags, true
}

// HasHostTags returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasHostTags() bool {
	return o != nil && o.HostTags != nil
}

// SetHostTags gets a reference to the given []string and assigns it to the HostTags field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetHostTags(v []string) {
	o.HostTags = v
}

// GetHostTagsLists returns the HostTagsLists field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetHostTagsLists() [][]string {
	if o == nil || o.HostTagsLists == nil {
		var ret [][]string
		return ret
	}
	return o.HostTagsLists
}

// GetHostTagsListsOk returns a tuple with the HostTagsLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetHostTagsListsOk() (*[][]string, bool) {
	if o == nil || o.HostTagsLists == nil {
		return nil, false
	}
	return &o.HostTagsLists, true
}

// HasHostTagsLists returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasHostTagsLists() bool {
	return o != nil && o.HostTagsLists != nil
}

// SetHostTagsLists gets a reference to the given [][]string and assigns it to the HostTagsLists field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetHostTagsLists(v [][]string) {
	o.HostTagsLists = v
}

// GetMonitoringRulesCount returns the MonitoringRulesCount field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetMonitoringRulesCount() int32 {
	if o == nil || o.MonitoringRulesCount == nil {
		var ret int32
		return ret
	}
	return *o.MonitoringRulesCount
}

// GetMonitoringRulesCountOk returns a tuple with the MonitoringRulesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetMonitoringRulesCountOk() (*int32, bool) {
	if o == nil || o.MonitoringRulesCount == nil {
		return nil, false
	}
	return o.MonitoringRulesCount, true
}

// HasMonitoringRulesCount returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasMonitoringRulesCount() bool {
	return o != nil && o.MonitoringRulesCount != nil
}

// SetMonitoringRulesCount gets a reference to the given int32 and assigns it to the MonitoringRulesCount field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetMonitoringRulesCount(v int32) {
	o.MonitoringRulesCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetName(v string) {
	o.Name = &v
}

// GetPinned returns the Pinned field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetPinned() bool {
	if o == nil || o.Pinned == nil {
		var ret bool
		return ret
	}
	return *o.Pinned
}

// GetPinnedOk returns a tuple with the Pinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetPinnedOk() (*bool, bool) {
	if o == nil || o.Pinned == nil {
		return nil, false
	}
	return o.Pinned, true
}

// HasPinned returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasPinned() bool {
	return o != nil && o.Pinned != nil
}

// SetPinned gets a reference to the given bool and assigns it to the Pinned field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetPinned(v bool) {
	o.Pinned = &v
}

// GetPolicyVersion returns the PolicyVersion field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetPolicyVersion() string {
	if o == nil || o.PolicyVersion == nil {
		var ret string
		return ret
	}
	return *o.PolicyVersion
}

// GetPolicyVersionOk returns a tuple with the PolicyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetPolicyVersionOk() (*string, bool) {
	if o == nil || o.PolicyVersion == nil {
		return nil, false
	}
	return o.PolicyVersion, true
}

// HasPolicyVersion returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasPolicyVersion() bool {
	return o != nil && o.PolicyVersion != nil
}

// SetPolicyVersion gets a reference to the given string and assigns it to the PolicyVersion field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetPolicyVersion(v string) {
	o.PolicyVersion = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetPriority(v int64) {
	o.Priority = &v
}

// GetRuleCount returns the RuleCount field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetRuleCount() int32 {
	if o == nil || o.RuleCount == nil {
		var ret int32
		return ret
	}
	return *o.RuleCount
}

// GetRuleCountOk returns a tuple with the RuleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetRuleCountOk() (*int32, bool) {
	if o == nil || o.RuleCount == nil {
		return nil, false
	}
	return o.RuleCount, true
}

// HasRuleCount returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasRuleCount() bool {
	return o != nil && o.RuleCount != nil
}

// SetRuleCount gets a reference to the given int32 and assigns it to the RuleCount field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetRuleCount(v int32) {
	o.RuleCount = &v
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetUpdateDate() int64 {
	if o == nil || o.UpdateDate == nil {
		var ret int64
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetUpdateDateOk() (*int64, bool) {
	if o == nil || o.UpdateDate == nil {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasUpdateDate() bool {
	return o != nil && o.UpdateDate != nil
}

// SetUpdateDate gets a reference to the given int64 and assigns it to the UpdateDate field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetUpdateDate(v int64) {
	o.UpdateDate = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdater returns the Updater field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetUpdater() CloudWorkloadSecurityAgentPolicyUpdaterAttributes {
	if o == nil || o.Updater == nil {
		var ret CloudWorkloadSecurityAgentPolicyUpdaterAttributes
		return ret
	}
	return *o.Updater
}

// GetUpdaterOk returns a tuple with the Updater field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetUpdaterOk() (*CloudWorkloadSecurityAgentPolicyUpdaterAttributes, bool) {
	if o == nil || o.Updater == nil {
		return nil, false
	}
	return o.Updater, true
}

// HasUpdater returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasUpdater() bool {
	return o != nil && o.Updater != nil
}

// SetUpdater gets a reference to the given CloudWorkloadSecurityAgentPolicyUpdaterAttributes and assigns it to the Updater field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetUpdater(v CloudWorkloadSecurityAgentPolicyUpdaterAttributes) {
	o.Updater = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetVersions() []CloudWorkloadSecurityAgentPolicyVersion {
	if o == nil || o.Versions == nil {
		var ret []CloudWorkloadSecurityAgentPolicyVersion
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) GetVersionsOk() (*[]CloudWorkloadSecurityAgentPolicyVersion, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return &o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) HasVersions() bool {
	return o != nil && o.Versions != nil
}

// SetVersions gets a reference to the given []CloudWorkloadSecurityAgentPolicyVersion and assigns it to the Versions field.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) SetVersions(v []CloudWorkloadSecurityAgentPolicyVersion) {
	o.Versions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentPolicyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BlockingRulesCount != nil {
		toSerialize["blockingRulesCount"] = o.BlockingRulesCount
	}
	if o.DatadogManaged != nil {
		toSerialize["datadogManaged"] = o.DatadogManaged
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisabledRulesCount != nil {
		toSerialize["disabledRulesCount"] = o.DisabledRulesCount
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.HostTags != nil {
		toSerialize["hostTags"] = o.HostTags
	}
	if o.HostTagsLists != nil {
		toSerialize["hostTagsLists"] = o.HostTagsLists
	}
	if o.MonitoringRulesCount != nil {
		toSerialize["monitoringRulesCount"] = o.MonitoringRulesCount
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Pinned != nil {
		toSerialize["pinned"] = o.Pinned
	}
	if o.PolicyVersion != nil {
		toSerialize["policyVersion"] = o.PolicyVersion
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.RuleCount != nil {
		toSerialize["ruleCount"] = o.RuleCount
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
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentPolicyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BlockingRulesCount   *int32                                             `json:"blockingRulesCount,omitempty"`
		DatadogManaged       *bool                                              `json:"datadogManaged,omitempty"`
		Description          *string                                            `json:"description,omitempty"`
		DisabledRulesCount   *int32                                             `json:"disabledRulesCount,omitempty"`
		Enabled              *bool                                              `json:"enabled,omitempty"`
		HostTags             []string                                           `json:"hostTags,omitempty"`
		HostTagsLists        [][]string                                         `json:"hostTagsLists,omitempty"`
		MonitoringRulesCount *int32                                             `json:"monitoringRulesCount,omitempty"`
		Name                 *string                                            `json:"name,omitempty"`
		Pinned               *bool                                              `json:"pinned,omitempty"`
		PolicyVersion        *string                                            `json:"policyVersion,omitempty"`
		Priority             *int64                                             `json:"priority,omitempty"`
		RuleCount            *int32                                             `json:"ruleCount,omitempty"`
		UpdateDate           *int64                                             `json:"updateDate,omitempty"`
		UpdatedAt            *int64                                             `json:"updatedAt,omitempty"`
		Updater              *CloudWorkloadSecurityAgentPolicyUpdaterAttributes `json:"updater,omitempty"`
		Versions             []CloudWorkloadSecurityAgentPolicyVersion          `json:"versions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"blockingRulesCount", "datadogManaged", "description", "disabledRulesCount", "enabled", "hostTags", "hostTagsLists", "monitoringRulesCount", "name", "pinned", "policyVersion", "priority", "ruleCount", "updateDate", "updatedAt", "updater", "versions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BlockingRulesCount = all.BlockingRulesCount
	o.DatadogManaged = all.DatadogManaged
	o.Description = all.Description
	o.DisabledRulesCount = all.DisabledRulesCount
	o.Enabled = all.Enabled
	o.HostTags = all.HostTags
	o.HostTagsLists = all.HostTagsLists
	o.MonitoringRulesCount = all.MonitoringRulesCount
	o.Name = all.Name
	o.Pinned = all.Pinned
	o.PolicyVersion = all.PolicyVersion
	o.Priority = all.Priority
	o.RuleCount = all.RuleCount
	o.UpdateDate = all.UpdateDate
	o.UpdatedAt = all.UpdatedAt
	if all.Updater != nil && all.Updater.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Updater = all.Updater
	o.Versions = all.Versions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
