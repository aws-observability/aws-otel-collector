// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DefaultRulesetsPerLanguageDataAttributes The attributes of the default rulesets per language response, containing the list of default ruleset names.
type DefaultRulesetsPerLanguageDataAttributes struct {
	// The list of default ruleset names for the specified programming language.
	Rulesets []string `json:"rulesets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDefaultRulesetsPerLanguageDataAttributes instantiates a new DefaultRulesetsPerLanguageDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDefaultRulesetsPerLanguageDataAttributes(rulesets []string) *DefaultRulesetsPerLanguageDataAttributes {
	this := DefaultRulesetsPerLanguageDataAttributes{}
	this.Rulesets = rulesets
	return &this
}

// NewDefaultRulesetsPerLanguageDataAttributesWithDefaults instantiates a new DefaultRulesetsPerLanguageDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDefaultRulesetsPerLanguageDataAttributesWithDefaults() *DefaultRulesetsPerLanguageDataAttributes {
	this := DefaultRulesetsPerLanguageDataAttributes{}
	return &this
}

// GetRulesets returns the Rulesets field value.
func (o *DefaultRulesetsPerLanguageDataAttributes) GetRulesets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Rulesets
}

// GetRulesetsOk returns a tuple with the Rulesets field value
// and a boolean to check if the value has been set.
func (o *DefaultRulesetsPerLanguageDataAttributes) GetRulesetsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rulesets, true
}

// SetRulesets sets field value.
func (o *DefaultRulesetsPerLanguageDataAttributes) SetRulesets(v []string) {
	o.Rulesets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DefaultRulesetsPerLanguageDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rulesets"] = o.Rulesets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DefaultRulesetsPerLanguageDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rulesets *[]string `json:"rulesets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rulesets == nil {
		return fmt.Errorf("required field rulesets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rulesets"})
	} else {
		return err
	}
	o.Rulesets = *all.Rulesets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
