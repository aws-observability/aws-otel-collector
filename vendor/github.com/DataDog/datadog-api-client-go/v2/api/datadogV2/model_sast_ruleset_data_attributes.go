// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SastRulesetDataAttributes The attributes of a SAST ruleset, including its name, description, and rules.
type SastRulesetDataAttributes struct {
	// A detailed description of the ruleset's purpose and the types of issues it targets.
	Description string `json:"description"`
	// The unique name of the ruleset.
	Name string `json:"name"`
	// The list of static analysis rules included in this ruleset.
	Rules []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems `json:"rules"`
	// A brief summary of the ruleset, suitable for display in listings.
	ShortDescription string `json:"short_description"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSastRulesetDataAttributes instantiates a new SastRulesetDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSastRulesetDataAttributes(description string, name string, rules []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems, shortDescription string) *SastRulesetDataAttributes {
	this := SastRulesetDataAttributes{}
	this.Description = description
	this.Name = name
	this.Rules = rules
	this.ShortDescription = shortDescription
	return &this
}

// NewSastRulesetDataAttributesWithDefaults instantiates a new SastRulesetDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSastRulesetDataAttributesWithDefaults() *SastRulesetDataAttributes {
	this := SastRulesetDataAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *SastRulesetDataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SastRulesetDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *SastRulesetDataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value.
func (o *SastRulesetDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SastRulesetDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SastRulesetDataAttributes) SetName(v string) {
	o.Name = v
}

// GetRules returns the Rules field value.
func (o *SastRulesetDataAttributes) GetRules() []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems {
	if o == nil {
		var ret []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *SastRulesetDataAttributes) GetRulesOk() (*[]GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *SastRulesetDataAttributes) SetRules(v []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) {
	o.Rules = v
}

// GetShortDescription returns the ShortDescription field value.
func (o *SastRulesetDataAttributes) GetShortDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
func (o *SastRulesetDataAttributes) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortDescription, true
}

// SetShortDescription sets field value.
func (o *SastRulesetDataAttributes) SetShortDescription(v string) {
	o.ShortDescription = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SastRulesetDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["rules"] = o.Rules
	toSerialize["short_description"] = o.ShortDescription

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SastRulesetDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description      *string                                                             `json:"description"`
		Name             *string                                                             `json:"name"`
		Rules            *[]GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems `json:"rules"`
		ShortDescription *string                                                             `json:"short_description"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	if all.ShortDescription == nil {
		return fmt.Errorf("required field short_description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "rules", "short_description"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.Name = *all.Name
	o.Rules = *all.Rules
	o.ShortDescription = *all.ShortDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
