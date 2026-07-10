// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleBasedViewAttributes Attributes of the rule-based view.
type RuleBasedViewAttributes struct {
	// Total number of rules in the view.
	Count int64 `json:"count"`
	// List of rules in the rule-based view.
	Rules []RuleBasedViewRule `json:"rules"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRuleBasedViewAttributes instantiates a new RuleBasedViewAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRuleBasedViewAttributes(count int64, rules []RuleBasedViewRule) *RuleBasedViewAttributes {
	this := RuleBasedViewAttributes{}
	this.Count = count
	this.Rules = rules
	return &this
}

// NewRuleBasedViewAttributesWithDefaults instantiates a new RuleBasedViewAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRuleBasedViewAttributesWithDefaults() *RuleBasedViewAttributes {
	this := RuleBasedViewAttributes{}
	return &this
}

// GetCount returns the Count field value.
func (o *RuleBasedViewAttributes) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewAttributes) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *RuleBasedViewAttributes) SetCount(v int64) {
	o.Count = v
}

// GetRules returns the Rules field value.
func (o *RuleBasedViewAttributes) GetRules() []RuleBasedViewRule {
	if o == nil {
		var ret []RuleBasedViewRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewAttributes) GetRulesOk() (*[]RuleBasedViewRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *RuleBasedViewAttributes) SetRules(v []RuleBasedViewRule) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RuleBasedViewAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count
	toSerialize["rules"] = o.Rules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RuleBasedViewAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *int64               `json:"count"`
		Rules *[]RuleBasedViewRule `json:"rules"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "rules"})
	} else {
		return err
	}
	o.Count = *all.Count
	o.Rules = *all.Rules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
