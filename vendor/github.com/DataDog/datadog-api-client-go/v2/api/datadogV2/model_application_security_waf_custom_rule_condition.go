// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleCondition One condition of the WAF Custom Rule.
type ApplicationSecurityWafCustomRuleCondition struct {
	// Operator to use for the WAF Condition.
	Operator ApplicationSecurityWafCustomRuleConditionOperator `json:"operator"`
	// The scope of the WAF custom rule.
	Parameters ApplicationSecurityWafCustomRuleConditionParameters `json:"parameters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityWafCustomRuleCondition instantiates a new ApplicationSecurityWafCustomRuleCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafCustomRuleCondition(operator ApplicationSecurityWafCustomRuleConditionOperator, parameters ApplicationSecurityWafCustomRuleConditionParameters) *ApplicationSecurityWafCustomRuleCondition {
	this := ApplicationSecurityWafCustomRuleCondition{}
	this.Operator = operator
	this.Parameters = parameters
	return &this
}

// NewApplicationSecurityWafCustomRuleConditionWithDefaults instantiates a new ApplicationSecurityWafCustomRuleCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafCustomRuleConditionWithDefaults() *ApplicationSecurityWafCustomRuleCondition {
	this := ApplicationSecurityWafCustomRuleCondition{}
	return &this
}

// GetOperator returns the Operator field value.
func (o *ApplicationSecurityWafCustomRuleCondition) GetOperator() ApplicationSecurityWafCustomRuleConditionOperator {
	if o == nil {
		var ret ApplicationSecurityWafCustomRuleConditionOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleCondition) GetOperatorOk() (*ApplicationSecurityWafCustomRuleConditionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *ApplicationSecurityWafCustomRuleCondition) SetOperator(v ApplicationSecurityWafCustomRuleConditionOperator) {
	o.Operator = v
}

// GetParameters returns the Parameters field value.
func (o *ApplicationSecurityWafCustomRuleCondition) GetParameters() ApplicationSecurityWafCustomRuleConditionParameters {
	if o == nil {
		var ret ApplicationSecurityWafCustomRuleConditionParameters
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleCondition) GetParametersOk() (*ApplicationSecurityWafCustomRuleConditionParameters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value.
func (o *ApplicationSecurityWafCustomRuleCondition) SetParameters(v ApplicationSecurityWafCustomRuleConditionParameters) {
	o.Parameters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafCustomRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["operator"] = o.Operator
	toSerialize["parameters"] = o.Parameters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafCustomRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Operator   *ApplicationSecurityWafCustomRuleConditionOperator   `json:"operator"`
		Parameters *ApplicationSecurityWafCustomRuleConditionParameters `json:"parameters"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Parameters == nil {
		return fmt.Errorf("required field parameters missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"operator", "parameters"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	if all.Parameters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Parameters = *all.Parameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
