// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleBasedViewRule A compliance rule along with its evaluation statistics and framework mappings.
type RuleBasedViewRule struct {
	// List of compliance framework mappings associated with the rule.
	ComplianceFrameworks []RuleBasedViewComplianceFramework `json:"compliance_frameworks"`
	// Whether the rule is enabled.
	Enabled bool `json:"enabled"`
	// Unique identifier of the rule.
	Id string `json:"id"`
	// Human-readable name of the rule.
	Name string `json:"name"`
	// List of resource attribute names exposed by the rule.
	ResourceAttributes []string `json:"resourceAttributes"`
	// Resource category targeted by the rule.
	ResourceCategory string `json:"resourceCategory"`
	// Resource type targeted by the rule.
	ResourceType string `json:"resourceType"`
	// Counts of findings for the rule, grouped by their evaluation status.
	Stats RuleBasedViewRuleStats `json:"stats"`
	// Severity associated with the rule (for example, `info`, `low`, `medium`, `high`, or `critical`).
	Status string `json:"status"`
	// List of tags attached to the rule.
	Tags []string `json:"tags"`
	// The category of the security rule.
	Type RuleBasedViewRuleCategory `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRuleBasedViewRule instantiates a new RuleBasedViewRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRuleBasedViewRule(complianceFrameworks []RuleBasedViewComplianceFramework, enabled bool, id string, name string, resourceAttributes []string, resourceCategory string, resourceType string, stats RuleBasedViewRuleStats, status string, tags []string, typeVar RuleBasedViewRuleCategory) *RuleBasedViewRule {
	this := RuleBasedViewRule{}
	this.ComplianceFrameworks = complianceFrameworks
	this.Enabled = enabled
	this.Id = id
	this.Name = name
	this.ResourceAttributes = resourceAttributes
	this.ResourceCategory = resourceCategory
	this.ResourceType = resourceType
	this.Stats = stats
	this.Status = status
	this.Tags = tags
	this.Type = typeVar
	return &this
}

// NewRuleBasedViewRuleWithDefaults instantiates a new RuleBasedViewRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRuleBasedViewRuleWithDefaults() *RuleBasedViewRule {
	this := RuleBasedViewRule{}
	return &this
}

// GetComplianceFrameworks returns the ComplianceFrameworks field value.
func (o *RuleBasedViewRule) GetComplianceFrameworks() []RuleBasedViewComplianceFramework {
	if o == nil {
		var ret []RuleBasedViewComplianceFramework
		return ret
	}
	return o.ComplianceFrameworks
}

// GetComplianceFrameworksOk returns a tuple with the ComplianceFrameworks field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRule) GetComplianceFrameworksOk() (*[]RuleBasedViewComplianceFramework, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceFrameworks, true
}

// SetComplianceFrameworks sets field value.
func (o *RuleBasedViewRule) SetComplianceFrameworks(v []RuleBasedViewComplianceFramework) {
	o.ComplianceFrameworks = v
}

// GetEnabled returns the Enabled field value.
func (o *RuleBasedViewRule) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRule) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *RuleBasedViewRule) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *RuleBasedViewRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RuleBasedViewRule) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *RuleBasedViewRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RuleBasedViewRule) SetName(v string) {
	o.Name = v
}

// GetResourceAttributes returns the ResourceAttributes field value.
func (o *RuleBasedViewRule) GetResourceAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ResourceAttributes
}

// GetResourceAttributesOk returns a tuple with the ResourceAttributes field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRule) GetResourceAttributesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceAttributes, true
}

// SetResourceAttributes sets field value.
func (o *RuleBasedViewRule) SetResourceAttributes(v []string) {
	o.ResourceAttributes = v
}

// GetResourceCategory returns the ResourceCategory field value.
func (o *RuleBasedViewRule) GetResourceCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceCategory
}

// GetResourceCategoryOk returns a tuple with the ResourceCategory field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRule) GetResourceCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceCategory, true
}

// SetResourceCategory sets field value.
func (o *RuleBasedViewRule) SetResourceCategory(v string) {
	o.ResourceCategory = v
}

// GetResourceType returns the ResourceType field value.
func (o *RuleBasedViewRule) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRule) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *RuleBasedViewRule) SetResourceType(v string) {
	o.ResourceType = v
}

// GetStats returns the Stats field value.
func (o *RuleBasedViewRule) GetStats() RuleBasedViewRuleStats {
	if o == nil {
		var ret RuleBasedViewRuleStats
		return ret
	}
	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRule) GetStatsOk() (*RuleBasedViewRuleStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value.
func (o *RuleBasedViewRule) SetStats(v RuleBasedViewRuleStats) {
	o.Stats = v
}

// GetStatus returns the Status field value.
func (o *RuleBasedViewRule) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRule) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *RuleBasedViewRule) SetStatus(v string) {
	o.Status = v
}

// GetTags returns the Tags field value.
func (o *RuleBasedViewRule) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRule) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *RuleBasedViewRule) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value.
func (o *RuleBasedViewRule) GetType() RuleBasedViewRuleCategory {
	if o == nil {
		var ret RuleBasedViewRuleCategory
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRule) GetTypeOk() (*RuleBasedViewRuleCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RuleBasedViewRule) SetType(v RuleBasedViewRuleCategory) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RuleBasedViewRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compliance_frameworks"] = o.ComplianceFrameworks
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["resourceAttributes"] = o.ResourceAttributes
	toSerialize["resourceCategory"] = o.ResourceCategory
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["stats"] = o.Stats
	toSerialize["status"] = o.Status
	toSerialize["tags"] = o.Tags
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RuleBasedViewRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComplianceFrameworks *[]RuleBasedViewComplianceFramework `json:"compliance_frameworks"`
		Enabled              *bool                               `json:"enabled"`
		Id                   *string                             `json:"id"`
		Name                 *string                             `json:"name"`
		ResourceAttributes   *[]string                           `json:"resourceAttributes"`
		ResourceCategory     *string                             `json:"resourceCategory"`
		ResourceType         *string                             `json:"resourceType"`
		Stats                *RuleBasedViewRuleStats             `json:"stats"`
		Status               *string                             `json:"status"`
		Tags                 *[]string                           `json:"tags"`
		Type                 *RuleBasedViewRuleCategory          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ComplianceFrameworks == nil {
		return fmt.Errorf("required field compliance_frameworks missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ResourceAttributes == nil {
		return fmt.Errorf("required field resourceAttributes missing")
	}
	if all.ResourceCategory == nil {
		return fmt.Errorf("required field resourceCategory missing")
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resourceType missing")
	}
	if all.Stats == nil {
		return fmt.Errorf("required field stats missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compliance_frameworks", "enabled", "id", "name", "resourceAttributes", "resourceCategory", "resourceType", "stats", "status", "tags", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComplianceFrameworks = *all.ComplianceFrameworks
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Name = *all.Name
	o.ResourceAttributes = *all.ResourceAttributes
	o.ResourceCategory = *all.ResourceCategory
	o.ResourceType = *all.ResourceType
	if all.Stats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Stats = *all.Stats
	o.Status = *all.Status
	o.Tags = *all.Tags
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
