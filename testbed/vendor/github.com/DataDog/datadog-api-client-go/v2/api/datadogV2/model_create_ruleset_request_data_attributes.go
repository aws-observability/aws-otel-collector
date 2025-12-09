// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateRulesetRequestDataAttributes The definition of `CreateRulesetRequestDataAttributes` object.
type CreateRulesetRequestDataAttributes struct {
	// The `attributes` `enabled`.
	Enabled *bool `json:"enabled,omitempty"`
	// The `attributes` `rules`.
	Rules []CreateRulesetRequestDataAttributesRulesItems `json:"rules"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateRulesetRequestDataAttributes instantiates a new CreateRulesetRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateRulesetRequestDataAttributes(rules []CreateRulesetRequestDataAttributesRulesItems) *CreateRulesetRequestDataAttributes {
	this := CreateRulesetRequestDataAttributes{}
	this.Rules = rules
	return &this
}

// NewCreateRulesetRequestDataAttributesWithDefaults instantiates a new CreateRulesetRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateRulesetRequestDataAttributesWithDefaults() *CreateRulesetRequestDataAttributes {
	this := CreateRulesetRequestDataAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateRulesetRequestDataAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRulesetRequestDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateRulesetRequestDataAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateRulesetRequestDataAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRules returns the Rules field value.
func (o *CreateRulesetRequestDataAttributes) GetRules() []CreateRulesetRequestDataAttributesRulesItems {
	if o == nil {
		var ret []CreateRulesetRequestDataAttributesRulesItems
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *CreateRulesetRequestDataAttributes) GetRulesOk() (*[]CreateRulesetRequestDataAttributesRulesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *CreateRulesetRequestDataAttributes) SetRules(v []CreateRulesetRequestDataAttributesRulesItems) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateRulesetRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["rules"] = o.Rules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateRulesetRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool                                           `json:"enabled,omitempty"`
		Rules   *[]CreateRulesetRequestDataAttributesRulesItems `json:"rules"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "rules"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.Rules = *all.Rules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
