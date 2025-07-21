// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineParseGrokProcessorRuleSupportRule The Grok helper rule referenced in the parsing rules.
type ObservabilityPipelineParseGrokProcessorRuleSupportRule struct {
	// The name of the Grok helper rule.
	Name string `json:"name"`
	// The definition of the Grok helper rule.
	Rule string `json:"rule"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineParseGrokProcessorRuleSupportRule instantiates a new ObservabilityPipelineParseGrokProcessorRuleSupportRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineParseGrokProcessorRuleSupportRule(name string, rule string) *ObservabilityPipelineParseGrokProcessorRuleSupportRule {
	this := ObservabilityPipelineParseGrokProcessorRuleSupportRule{}
	this.Name = name
	this.Rule = rule
	return &this
}

// NewObservabilityPipelineParseGrokProcessorRuleSupportRuleWithDefaults instantiates a new ObservabilityPipelineParseGrokProcessorRuleSupportRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineParseGrokProcessorRuleSupportRuleWithDefaults() *ObservabilityPipelineParseGrokProcessorRuleSupportRule {
	this := ObservabilityPipelineParseGrokProcessorRuleSupportRule{}
	return &this
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineParseGrokProcessorRuleSupportRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessorRuleSupportRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineParseGrokProcessorRuleSupportRule) SetName(v string) {
	o.Name = v
}

// GetRule returns the Rule field value.
func (o *ObservabilityPipelineParseGrokProcessorRuleSupportRule) GetRule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessorRuleSupportRule) GetRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *ObservabilityPipelineParseGrokProcessorRuleSupportRule) SetRule(v string) {
	o.Rule = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineParseGrokProcessorRuleSupportRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["rule"] = o.Rule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineParseGrokProcessorRuleSupportRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name"`
		Rule *string `json:"rule"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Rule == nil {
		return fmt.Errorf("required field rule missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "rule"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Rule = *all.Rule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
