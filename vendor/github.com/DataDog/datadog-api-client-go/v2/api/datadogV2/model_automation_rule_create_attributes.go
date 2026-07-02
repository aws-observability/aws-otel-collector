// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleCreateAttributes Attributes required to create an automation rule.
type AutomationRuleCreateAttributes struct {
	// Defines what happens when the rule triggers. Combines an action type with action-specific configuration data.
	Action AutomationRuleAction `json:"action"`
	// Name of the automation rule.
	Name string `json:"name"`
	// Whether the automation rule is active. Enabled rules trigger on matching case events; disabled rules are inactive but preserve their configuration.
	State *CaseAutomationRuleState `json:"state,omitempty"`
	// Defines when the rule activates. Combines a trigger type (the case event to listen for) with optional trigger data (conditions that narrow when the trigger fires).
	Trigger AutomationRuleTrigger `json:"trigger"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutomationRuleCreateAttributes instantiates a new AutomationRuleCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutomationRuleCreateAttributes(action AutomationRuleAction, name string, trigger AutomationRuleTrigger) *AutomationRuleCreateAttributes {
	this := AutomationRuleCreateAttributes{}
	this.Action = action
	this.Name = name
	this.Trigger = trigger
	return &this
}

// NewAutomationRuleCreateAttributesWithDefaults instantiates a new AutomationRuleCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutomationRuleCreateAttributesWithDefaults() *AutomationRuleCreateAttributes {
	this := AutomationRuleCreateAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *AutomationRuleCreateAttributes) GetAction() AutomationRuleAction {
	if o == nil {
		var ret AutomationRuleAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleCreateAttributes) GetActionOk() (*AutomationRuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *AutomationRuleCreateAttributes) SetAction(v AutomationRuleAction) {
	o.Action = v
}

// GetName returns the Name field value.
func (o *AutomationRuleCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AutomationRuleCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AutomationRuleCreateAttributes) GetState() CaseAutomationRuleState {
	if o == nil || o.State == nil {
		var ret CaseAutomationRuleState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleCreateAttributes) GetStateOk() (*CaseAutomationRuleState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AutomationRuleCreateAttributes) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given CaseAutomationRuleState and assigns it to the State field.
func (o *AutomationRuleCreateAttributes) SetState(v CaseAutomationRuleState) {
	o.State = &v
}

// GetTrigger returns the Trigger field value.
func (o *AutomationRuleCreateAttributes) GetTrigger() AutomationRuleTrigger {
	if o == nil {
		var ret AutomationRuleTrigger
		return ret
	}
	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleCreateAttributes) GetTriggerOk() (*AutomationRuleTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value.
func (o *AutomationRuleCreateAttributes) SetTrigger(v AutomationRuleTrigger) {
	o.Trigger = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutomationRuleCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["name"] = o.Name
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	toSerialize["trigger"] = o.Trigger

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutomationRuleCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action  *AutomationRuleAction    `json:"action"`
		Name    *string                  `json:"name"`
		State   *CaseAutomationRuleState `json:"state,omitempty"`
		Trigger *AutomationRuleTrigger   `json:"trigger"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Trigger == nil {
		return fmt.Errorf("required field trigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "name", "state", "trigger"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Action.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Action = *all.Action
	o.Name = *all.Name
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
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
