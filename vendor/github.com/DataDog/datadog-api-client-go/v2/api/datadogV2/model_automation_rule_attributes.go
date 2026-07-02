// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleAttributes Core attributes of an automation rule, including its name, trigger condition, action to execute, and current state.
type AutomationRuleAttributes struct {
	// Defines what happens when the rule triggers. Combines an action type with action-specific configuration data.
	Action AutomationRuleAction `json:"action"`
	// Timestamp when the automation rule was created.
	CreatedAt time.Time `json:"created_at"`
	// Timestamp when the automation rule was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// A human-readable name for the automation rule, used to identify the rule in the UI and API responses.
	Name string `json:"name"`
	// Whether the automation rule is active. Enabled rules trigger on matching case events; disabled rules are inactive but preserve their configuration.
	State CaseAutomationRuleState `json:"state"`
	// Defines when the rule activates. Combines a trigger type (the case event to listen for) with optional trigger data (conditions that narrow when the trigger fires).
	Trigger AutomationRuleTrigger `json:"trigger"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutomationRuleAttributes instantiates a new AutomationRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutomationRuleAttributes(action AutomationRuleAction, createdAt time.Time, name string, state CaseAutomationRuleState, trigger AutomationRuleTrigger) *AutomationRuleAttributes {
	this := AutomationRuleAttributes{}
	this.Action = action
	this.CreatedAt = createdAt
	this.Name = name
	this.State = state
	this.Trigger = trigger
	return &this
}

// NewAutomationRuleAttributesWithDefaults instantiates a new AutomationRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutomationRuleAttributesWithDefaults() *AutomationRuleAttributes {
	this := AutomationRuleAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *AutomationRuleAttributes) GetAction() AutomationRuleAction {
	if o == nil {
		var ret AutomationRuleAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleAttributes) GetActionOk() (*AutomationRuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *AutomationRuleAttributes) SetAction(v AutomationRuleAction) {
	o.Action = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AutomationRuleAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AutomationRuleAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *AutomationRuleAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *AutomationRuleAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *AutomationRuleAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value.
func (o *AutomationRuleAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AutomationRuleAttributes) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value.
func (o *AutomationRuleAttributes) GetState() CaseAutomationRuleState {
	if o == nil {
		var ret CaseAutomationRuleState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleAttributes) GetStateOk() (*CaseAutomationRuleState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *AutomationRuleAttributes) SetState(v CaseAutomationRuleState) {
	o.State = v
}

// GetTrigger returns the Trigger field value.
func (o *AutomationRuleAttributes) GetTrigger() AutomationRuleTrigger {
	if o == nil {
		var ret AutomationRuleTrigger
		return ret
	}
	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleAttributes) GetTriggerOk() (*AutomationRuleTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value.
func (o *AutomationRuleAttributes) SetTrigger(v AutomationRuleTrigger) {
	o.Trigger = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutomationRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["trigger"] = o.Trigger

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutomationRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action     *AutomationRuleAction    `json:"action"`
		CreatedAt  *time.Time               `json:"created_at"`
		ModifiedAt *time.Time               `json:"modified_at,omitempty"`
		Name       *string                  `json:"name"`
		State      *CaseAutomationRuleState `json:"state"`
		Trigger    *AutomationRuleTrigger   `json:"trigger"`
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
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.Trigger == nil {
		return fmt.Errorf("required field trigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "created_at", "modified_at", "name", "state", "trigger"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Action.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Action = *all.Action
	o.CreatedAt = *all.CreatedAt
	o.ModifiedAt = all.ModifiedAt
	o.Name = *all.Name
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	if all.Trigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Trigger = *all.Trigger

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
