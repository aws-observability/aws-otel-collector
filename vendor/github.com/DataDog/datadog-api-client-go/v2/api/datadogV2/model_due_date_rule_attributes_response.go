// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DueDateRuleAttributesResponse Attributes of a due date rule returned by the API.
type DueDateRuleAttributesResponse struct {
	// The action to take when the due date rule matches a finding.
	Action DueDateRuleAction `json:"action"`
	// The Unix timestamp in milliseconds when the rule was created.
	CreatedAt int64 `json:"created_at"`
	// The user or Datadog system who created the rule.
	CreatedBy AutomationRuleCreatedBy `json:"created_by"`
	// Whether the due date rule is enabled.
	Enabled bool `json:"enabled"`
	// The Unix timestamp in milliseconds when the rule was last modified.
	ModifiedAt int64 `json:"modified_at"`
	// The user or Datadog system who last modified the rule.
	ModifiedBy AutomationRuleModifiedBy `json:"modified_by"`
	// The name of the due date rule.
	Name string `json:"name"`
	// Defines the scope of findings to which the automation rule applies.
	Rule AutomationRuleScope `json:"rule"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDueDateRuleAttributesResponse instantiates a new DueDateRuleAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDueDateRuleAttributesResponse(action DueDateRuleAction, createdAt int64, createdBy AutomationRuleCreatedBy, enabled bool, modifiedAt int64, modifiedBy AutomationRuleModifiedBy, name string, rule AutomationRuleScope) *DueDateRuleAttributesResponse {
	this := DueDateRuleAttributesResponse{}
	this.Action = action
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Enabled = enabled
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Name = name
	this.Rule = rule
	return &this
}

// NewDueDateRuleAttributesResponseWithDefaults instantiates a new DueDateRuleAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDueDateRuleAttributesResponseWithDefaults() *DueDateRuleAttributesResponse {
	this := DueDateRuleAttributesResponse{}
	return &this
}

// GetAction returns the Action field value.
func (o *DueDateRuleAttributesResponse) GetAction() DueDateRuleAction {
	if o == nil {
		var ret DueDateRuleAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *DueDateRuleAttributesResponse) GetActionOk() (*DueDateRuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *DueDateRuleAttributesResponse) SetAction(v DueDateRuleAction) {
	o.Action = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DueDateRuleAttributesResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DueDateRuleAttributesResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DueDateRuleAttributesResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *DueDateRuleAttributesResponse) GetCreatedBy() AutomationRuleCreatedBy {
	if o == nil {
		var ret AutomationRuleCreatedBy
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *DueDateRuleAttributesResponse) GetCreatedByOk() (*AutomationRuleCreatedBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *DueDateRuleAttributesResponse) SetCreatedBy(v AutomationRuleCreatedBy) {
	o.CreatedBy = v
}

// GetEnabled returns the Enabled field value.
func (o *DueDateRuleAttributesResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DueDateRuleAttributesResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *DueDateRuleAttributesResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *DueDateRuleAttributesResponse) GetModifiedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *DueDateRuleAttributesResponse) GetModifiedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *DueDateRuleAttributesResponse) SetModifiedAt(v int64) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *DueDateRuleAttributesResponse) GetModifiedBy() AutomationRuleModifiedBy {
	if o == nil {
		var ret AutomationRuleModifiedBy
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *DueDateRuleAttributesResponse) GetModifiedByOk() (*AutomationRuleModifiedBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *DueDateRuleAttributesResponse) SetModifiedBy(v AutomationRuleModifiedBy) {
	o.ModifiedBy = v
}

// GetName returns the Name field value.
func (o *DueDateRuleAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DueDateRuleAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DueDateRuleAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetRule returns the Rule field value.
func (o *DueDateRuleAttributesResponse) GetRule() AutomationRuleScope {
	if o == nil {
		var ret AutomationRuleScope
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *DueDateRuleAttributesResponse) GetRuleOk() (*AutomationRuleScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *DueDateRuleAttributesResponse) SetRule(v AutomationRuleScope) {
	o.Rule = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DueDateRuleAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["enabled"] = o.Enabled
	toSerialize["modified_at"] = o.ModifiedAt
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["name"] = o.Name
	toSerialize["rule"] = o.Rule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DueDateRuleAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action     *DueDateRuleAction        `json:"action"`
		CreatedAt  *int64                    `json:"created_at"`
		CreatedBy  *AutomationRuleCreatedBy  `json:"created_by"`
		Enabled    *bool                     `json:"enabled"`
		ModifiedAt *int64                    `json:"modified_at"`
		ModifiedBy *AutomationRuleModifiedBy `json:"modified_by"`
		Name       *string                   `json:"name"`
		Rule       *AutomationRuleScope      `json:"rule"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Rule == nil {
		return fmt.Errorf("required field rule missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "created_at", "created_by", "enabled", "modified_at", "modified_by", "name", "rule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Action.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Action = *all.Action
	o.CreatedAt = *all.CreatedAt
	if all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = *all.CreatedBy
	o.Enabled = *all.Enabled
	o.ModifiedAt = *all.ModifiedAt
	if all.ModifiedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ModifiedBy = *all.ModifiedBy
	o.Name = *all.Name
	if all.Rule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rule = *all.Rule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
