// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineParseGrokProcessorRuleMatchRule Defines a Grok parsing rule, which extracts structured fields from log content using named Grok patterns.
// Each rule must have a unique name and a valid Datadog Grok pattern that will be applied to the source field.
type ObservabilityPipelineParseGrokProcessorRuleMatchRule struct {
	// The name of the rule.
	Name string `json:"name"`
	// The definition of the Grok rule.
	Rule string `json:"rule"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineParseGrokProcessorRuleMatchRule instantiates a new ObservabilityPipelineParseGrokProcessorRuleMatchRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineParseGrokProcessorRuleMatchRule(name string, rule string) *ObservabilityPipelineParseGrokProcessorRuleMatchRule {
	this := ObservabilityPipelineParseGrokProcessorRuleMatchRule{}
	this.Name = name
	this.Rule = rule
	return &this
}

// NewObservabilityPipelineParseGrokProcessorRuleMatchRuleWithDefaults instantiates a new ObservabilityPipelineParseGrokProcessorRuleMatchRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineParseGrokProcessorRuleMatchRuleWithDefaults() *ObservabilityPipelineParseGrokProcessorRuleMatchRule {
	this := ObservabilityPipelineParseGrokProcessorRuleMatchRule{}
	return &this
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineParseGrokProcessorRuleMatchRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessorRuleMatchRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineParseGrokProcessorRuleMatchRule) SetName(v string) {
	o.Name = v
}

// GetRule returns the Rule field value.
func (o *ObservabilityPipelineParseGrokProcessorRuleMatchRule) GetRule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessorRuleMatchRule) GetRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *ObservabilityPipelineParseGrokProcessorRuleMatchRule) SetRule(v string) {
	o.Rule = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineParseGrokProcessorRuleMatchRule) MarshalJSON() ([]byte, error) {
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
func (o *ObservabilityPipelineParseGrokProcessorRuleMatchRule) UnmarshalJSON(bytes []byte) (err error) {
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
