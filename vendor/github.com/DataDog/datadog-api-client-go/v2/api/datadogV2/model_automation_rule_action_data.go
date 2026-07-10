// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleActionData Configuration for the action to execute, dependent on the action type.
type AutomationRuleActionData struct {
	// The type of AI agent to assign. Required when the action type is `assign_agent`.
	AgentType *string `json:"agent_type,omitempty"`
	// The identifier of the AI agent to assign to the case. Required when the action type is `assign_agent`.
	AssignedAgentId *string `json:"assigned_agent_id,omitempty"`
	// The handle of the Datadog workflow to execute. Required when the action type is `execute_workflow`.
	Handle *string `json:"handle,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutomationRuleActionData instantiates a new AutomationRuleActionData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutomationRuleActionData() *AutomationRuleActionData {
	this := AutomationRuleActionData{}
	return &this
}

// NewAutomationRuleActionDataWithDefaults instantiates a new AutomationRuleActionData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutomationRuleActionDataWithDefaults() *AutomationRuleActionData {
	this := AutomationRuleActionData{}
	return &this
}

// GetAgentType returns the AgentType field value if set, zero value otherwise.
func (o *AutomationRuleActionData) GetAgentType() string {
	if o == nil || o.AgentType == nil {
		var ret string
		return ret
	}
	return *o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleActionData) GetAgentTypeOk() (*string, bool) {
	if o == nil || o.AgentType == nil {
		return nil, false
	}
	return o.AgentType, true
}

// HasAgentType returns a boolean if a field has been set.
func (o *AutomationRuleActionData) HasAgentType() bool {
	return o != nil && o.AgentType != nil
}

// SetAgentType gets a reference to the given string and assigns it to the AgentType field.
func (o *AutomationRuleActionData) SetAgentType(v string) {
	o.AgentType = &v
}

// GetAssignedAgentId returns the AssignedAgentId field value if set, zero value otherwise.
func (o *AutomationRuleActionData) GetAssignedAgentId() string {
	if o == nil || o.AssignedAgentId == nil {
		var ret string
		return ret
	}
	return *o.AssignedAgentId
}

// GetAssignedAgentIdOk returns a tuple with the AssignedAgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleActionData) GetAssignedAgentIdOk() (*string, bool) {
	if o == nil || o.AssignedAgentId == nil {
		return nil, false
	}
	return o.AssignedAgentId, true
}

// HasAssignedAgentId returns a boolean if a field has been set.
func (o *AutomationRuleActionData) HasAssignedAgentId() bool {
	return o != nil && o.AssignedAgentId != nil
}

// SetAssignedAgentId gets a reference to the given string and assigns it to the AssignedAgentId field.
func (o *AutomationRuleActionData) SetAssignedAgentId(v string) {
	o.AssignedAgentId = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *AutomationRuleActionData) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleActionData) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *AutomationRuleActionData) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *AutomationRuleActionData) SetHandle(v string) {
	o.Handle = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutomationRuleActionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgentType != nil {
		toSerialize["agent_type"] = o.AgentType
	}
	if o.AssignedAgentId != nil {
		toSerialize["assigned_agent_id"] = o.AssignedAgentId
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutomationRuleActionData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentType       *string `json:"agent_type,omitempty"`
		AssignedAgentId *string `json:"assigned_agent_id,omitempty"`
		Handle          *string `json:"handle,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_type", "assigned_agent_id", "handle"})
	} else {
		return err
	}
	o.AgentType = all.AgentType
	o.AssignedAgentId = all.AssignedAgentId
	o.Handle = all.Handle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
