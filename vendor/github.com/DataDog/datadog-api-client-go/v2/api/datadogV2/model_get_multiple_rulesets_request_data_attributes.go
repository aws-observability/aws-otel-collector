// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsRequestDataAttributes
type GetMultipleRulesetsRequestDataAttributes struct {
	//
	IncludeTestingRules *bool `json:"include_testing_rules,omitempty"`
	//
	IncludeTests *bool `json:"include_tests,omitempty"`
	//
	Rulesets []string `json:"rulesets,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetMultipleRulesetsRequestDataAttributes instantiates a new GetMultipleRulesetsRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetMultipleRulesetsRequestDataAttributes() *GetMultipleRulesetsRequestDataAttributes {
	this := GetMultipleRulesetsRequestDataAttributes{}
	return &this
}

// NewGetMultipleRulesetsRequestDataAttributesWithDefaults instantiates a new GetMultipleRulesetsRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetMultipleRulesetsRequestDataAttributesWithDefaults() *GetMultipleRulesetsRequestDataAttributes {
	this := GetMultipleRulesetsRequestDataAttributes{}
	return &this
}

// GetIncludeTestingRules returns the IncludeTestingRules field value if set, zero value otherwise.
func (o *GetMultipleRulesetsRequestDataAttributes) GetIncludeTestingRules() bool {
	if o == nil || o.IncludeTestingRules == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTestingRules
}

// GetIncludeTestingRulesOk returns a tuple with the IncludeTestingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsRequestDataAttributes) GetIncludeTestingRulesOk() (*bool, bool) {
	if o == nil || o.IncludeTestingRules == nil {
		return nil, false
	}
	return o.IncludeTestingRules, true
}

// HasIncludeTestingRules returns a boolean if a field has been set.
func (o *GetMultipleRulesetsRequestDataAttributes) HasIncludeTestingRules() bool {
	return o != nil && o.IncludeTestingRules != nil
}

// SetIncludeTestingRules gets a reference to the given bool and assigns it to the IncludeTestingRules field.
func (o *GetMultipleRulesetsRequestDataAttributes) SetIncludeTestingRules(v bool) {
	o.IncludeTestingRules = &v
}

// GetIncludeTests returns the IncludeTests field value if set, zero value otherwise.
func (o *GetMultipleRulesetsRequestDataAttributes) GetIncludeTests() bool {
	if o == nil || o.IncludeTests == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTests
}

// GetIncludeTestsOk returns a tuple with the IncludeTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsRequestDataAttributes) GetIncludeTestsOk() (*bool, bool) {
	if o == nil || o.IncludeTests == nil {
		return nil, false
	}
	return o.IncludeTests, true
}

// HasIncludeTests returns a boolean if a field has been set.
func (o *GetMultipleRulesetsRequestDataAttributes) HasIncludeTests() bool {
	return o != nil && o.IncludeTests != nil
}

// SetIncludeTests gets a reference to the given bool and assigns it to the IncludeTests field.
func (o *GetMultipleRulesetsRequestDataAttributes) SetIncludeTests(v bool) {
	o.IncludeTests = &v
}

// GetRulesets returns the Rulesets field value if set, zero value otherwise.
func (o *GetMultipleRulesetsRequestDataAttributes) GetRulesets() []string {
	if o == nil || o.Rulesets == nil {
		var ret []string
		return ret
	}
	return o.Rulesets
}

// GetRulesetsOk returns a tuple with the Rulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsRequestDataAttributes) GetRulesetsOk() (*[]string, bool) {
	if o == nil || o.Rulesets == nil {
		return nil, false
	}
	return &o.Rulesets, true
}

// HasRulesets returns a boolean if a field has been set.
func (o *GetMultipleRulesetsRequestDataAttributes) HasRulesets() bool {
	return o != nil && o.Rulesets != nil
}

// SetRulesets gets a reference to the given []string and assigns it to the Rulesets field.
func (o *GetMultipleRulesetsRequestDataAttributes) SetRulesets(v []string) {
	o.Rulesets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetMultipleRulesetsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncludeTestingRules != nil {
		toSerialize["include_testing_rules"] = o.IncludeTestingRules
	}
	if o.IncludeTests != nil {
		toSerialize["include_tests"] = o.IncludeTests
	}
	if o.Rulesets != nil {
		toSerialize["rulesets"] = o.Rulesets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetMultipleRulesetsRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeTestingRules *bool    `json:"include_testing_rules,omitempty"`
		IncludeTests        *bool    `json:"include_tests,omitempty"`
		Rulesets            []string `json:"rulesets,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_testing_rules", "include_tests", "rulesets"})
	} else {
		return err
	}
	o.IncludeTestingRules = all.IncludeTestingRules
	o.IncludeTests = all.IncludeTests
	o.Rulesets = all.Rulesets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
