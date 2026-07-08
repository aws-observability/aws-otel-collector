// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleAction Defines what happens when the rule triggers. Combines an action type with action-specific configuration data.
type AutomationRuleAction struct {
	// Configuration for the action to execute, dependent on the action type.
	Data AutomationRuleActionData `json:"data"`
	// The type of automated action to perform when the rule triggers. `execute_workflow` runs a Datadog workflow; `assign_agent` assigns an AI agent to the case.
	Type AutomationRuleActionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutomationRuleAction instantiates a new AutomationRuleAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutomationRuleAction(data AutomationRuleActionData, typeVar AutomationRuleActionType) *AutomationRuleAction {
	this := AutomationRuleAction{}
	this.Data = data
	this.Type = typeVar
	return &this
}

// NewAutomationRuleActionWithDefaults instantiates a new AutomationRuleAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutomationRuleActionWithDefaults() *AutomationRuleAction {
	this := AutomationRuleAction{}
	return &this
}

// GetData returns the Data field value.
func (o *AutomationRuleAction) GetData() AutomationRuleActionData {
	if o == nil {
		var ret AutomationRuleActionData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleAction) GetDataOk() (*AutomationRuleActionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *AutomationRuleAction) SetData(v AutomationRuleActionData) {
	o.Data = v
}

// GetType returns the Type field value.
func (o *AutomationRuleAction) GetType() AutomationRuleActionType {
	if o == nil {
		var ret AutomationRuleActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleAction) GetTypeOk() (*AutomationRuleActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AutomationRuleAction) SetType(v AutomationRuleActionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutomationRuleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutomationRuleAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *AutomationRuleActionData `json:"data"`
		Type *AutomationRuleActionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
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
