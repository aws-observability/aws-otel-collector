// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleAction The definition of `ApplicationSecurityWafCustomRuleAction` object.
type ApplicationSecurityWafCustomRuleAction struct {
	// Override the default action to take when the WAF custom rule would block.
	Action *ApplicationSecurityWafCustomRuleActionAction `json:"action,omitempty"`
	// The definition of `ApplicationSecurityWafCustomRuleActionParameters` object.
	Parameters *ApplicationSecurityWafCustomRuleActionParameters `json:"parameters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityWafCustomRuleAction instantiates a new ApplicationSecurityWafCustomRuleAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafCustomRuleAction() *ApplicationSecurityWafCustomRuleAction {
	this := ApplicationSecurityWafCustomRuleAction{}
	var action ApplicationSecurityWafCustomRuleActionAction = APPLICATIONSECURITYWAFCUSTOMRULEACTIONACTION_BLOCK_REQUEST
	this.Action = &action
	return &this
}

// NewApplicationSecurityWafCustomRuleActionWithDefaults instantiates a new ApplicationSecurityWafCustomRuleAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafCustomRuleActionWithDefaults() *ApplicationSecurityWafCustomRuleAction {
	this := ApplicationSecurityWafCustomRuleAction{}
	var action ApplicationSecurityWafCustomRuleActionAction = APPLICATIONSECURITYWAFCUSTOMRULEACTIONACTION_BLOCK_REQUEST
	this.Action = &action
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleAction) GetAction() ApplicationSecurityWafCustomRuleActionAction {
	if o == nil || o.Action == nil {
		var ret ApplicationSecurityWafCustomRuleActionAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleAction) GetActionOk() (*ApplicationSecurityWafCustomRuleActionAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleAction) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given ApplicationSecurityWafCustomRuleActionAction and assigns it to the Action field.
func (o *ApplicationSecurityWafCustomRuleAction) SetAction(v ApplicationSecurityWafCustomRuleActionAction) {
	o.Action = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleAction) GetParameters() ApplicationSecurityWafCustomRuleActionParameters {
	if o == nil || o.Parameters == nil {
		var ret ApplicationSecurityWafCustomRuleActionParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleAction) GetParametersOk() (*ApplicationSecurityWafCustomRuleActionParameters, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleAction) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given ApplicationSecurityWafCustomRuleActionParameters and assigns it to the Parameters field.
func (o *ApplicationSecurityWafCustomRuleAction) SetParameters(v ApplicationSecurityWafCustomRuleActionParameters) {
	o.Parameters = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafCustomRuleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafCustomRuleAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action     *ApplicationSecurityWafCustomRuleActionAction     `json:"action,omitempty"`
		Parameters *ApplicationSecurityWafCustomRuleActionParameters `json:"parameters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "parameters"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Action != nil && !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = all.Action
	}
	if all.Parameters != nil && all.Parameters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Parameters = all.Parameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
