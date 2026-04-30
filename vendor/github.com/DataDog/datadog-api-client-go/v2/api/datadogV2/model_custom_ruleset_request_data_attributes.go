// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRulesetRequestDataAttributes Attributes for creating or updating a custom ruleset.
type CustomRulesetRequestDataAttributes struct {
	// Base64-encoded full description
	Description *string `json:"description,omitempty"`
	// Ruleset name
	Name *string `json:"name,omitempty"`
	// Rules in the ruleset
	Rules []CustomRule `json:"rules,omitempty"`
	// Base64-encoded short description
	ShortDescription *string `json:"short_description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomRulesetRequestDataAttributes instantiates a new CustomRulesetRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomRulesetRequestDataAttributes() *CustomRulesetRequestDataAttributes {
	this := CustomRulesetRequestDataAttributes{}
	return &this
}

// NewCustomRulesetRequestDataAttributesWithDefaults instantiates a new CustomRulesetRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomRulesetRequestDataAttributesWithDefaults() *CustomRulesetRequestDataAttributes {
	this := CustomRulesetRequestDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomRulesetRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRulesetRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomRulesetRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomRulesetRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomRulesetRequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRulesetRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomRulesetRequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomRulesetRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomRulesetRequestDataAttributes) GetRules() []CustomRule {
	if o == nil {
		var ret []CustomRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRulesetRequestDataAttributes) GetRulesOk() (*[]CustomRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *CustomRulesetRequestDataAttributes) HasRules() bool {
	return o != nil && o.Rules != nil
}

// SetRules gets a reference to the given []CustomRule and assigns it to the Rules field.
func (o *CustomRulesetRequestDataAttributes) SetRules(v []CustomRule) {
	o.Rules = v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *CustomRulesetRequestDataAttributes) GetShortDescription() string {
	if o == nil || o.ShortDescription == nil {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRulesetRequestDataAttributes) GetShortDescriptionOk() (*string, bool) {
	if o == nil || o.ShortDescription == nil {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *CustomRulesetRequestDataAttributes) HasShortDescription() bool {
	return o != nil && o.ShortDescription != nil
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *CustomRulesetRequestDataAttributes) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomRulesetRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.ShortDescription != nil {
		toSerialize["short_description"] = o.ShortDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomRulesetRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description      *string      `json:"description,omitempty"`
		Name             *string      `json:"name,omitempty"`
		Rules            []CustomRule `json:"rules,omitempty"`
		ShortDescription *string      `json:"short_description,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "rules", "short_description"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Name = all.Name
	o.Rules = all.Rules
	o.ShortDescription = all.ShortDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
