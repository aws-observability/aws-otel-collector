// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineParseGrokProcessorRule A Grok parsing rule used in the `parse_grok` processor. Each rule defines how to extract structured fields
// from a specific log field using Grok patterns.
type ObservabilityPipelineParseGrokProcessorRule struct {
	// A list of Grok parsing rules that define how to extract fields from the source field.
	// Each rule must contain a name and a valid Grok pattern.
	MatchRules []ObservabilityPipelineParseGrokProcessorRuleMatchRule `json:"match_rules"`
	// The name of the field in the log event to apply the Grok rules to.
	Source string `json:"source"`
	// A list of Grok helper rules that can be referenced by the parsing rules.
	SupportRules []ObservabilityPipelineParseGrokProcessorRuleSupportRule `json:"support_rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineParseGrokProcessorRule instantiates a new ObservabilityPipelineParseGrokProcessorRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineParseGrokProcessorRule(matchRules []ObservabilityPipelineParseGrokProcessorRuleMatchRule, source string) *ObservabilityPipelineParseGrokProcessorRule {
	this := ObservabilityPipelineParseGrokProcessorRule{}
	this.MatchRules = matchRules
	this.Source = source
	return &this
}

// NewObservabilityPipelineParseGrokProcessorRuleWithDefaults instantiates a new ObservabilityPipelineParseGrokProcessorRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineParseGrokProcessorRuleWithDefaults() *ObservabilityPipelineParseGrokProcessorRule {
	this := ObservabilityPipelineParseGrokProcessorRule{}
	return &this
}

// GetMatchRules returns the MatchRules field value.
func (o *ObservabilityPipelineParseGrokProcessorRule) GetMatchRules() []ObservabilityPipelineParseGrokProcessorRuleMatchRule {
	if o == nil {
		var ret []ObservabilityPipelineParseGrokProcessorRuleMatchRule
		return ret
	}
	return o.MatchRules
}

// GetMatchRulesOk returns a tuple with the MatchRules field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessorRule) GetMatchRulesOk() (*[]ObservabilityPipelineParseGrokProcessorRuleMatchRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchRules, true
}

// SetMatchRules sets field value.
func (o *ObservabilityPipelineParseGrokProcessorRule) SetMatchRules(v []ObservabilityPipelineParseGrokProcessorRuleMatchRule) {
	o.MatchRules = v
}

// GetSource returns the Source field value.
func (o *ObservabilityPipelineParseGrokProcessorRule) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessorRule) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *ObservabilityPipelineParseGrokProcessorRule) SetSource(v string) {
	o.Source = v
}

// GetSupportRules returns the SupportRules field value if set, zero value otherwise.
func (o *ObservabilityPipelineParseGrokProcessorRule) GetSupportRules() []ObservabilityPipelineParseGrokProcessorRuleSupportRule {
	if o == nil || o.SupportRules == nil {
		var ret []ObservabilityPipelineParseGrokProcessorRuleSupportRule
		return ret
	}
	return o.SupportRules
}

// GetSupportRulesOk returns a tuple with the SupportRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessorRule) GetSupportRulesOk() (*[]ObservabilityPipelineParseGrokProcessorRuleSupportRule, bool) {
	if o == nil || o.SupportRules == nil {
		return nil, false
	}
	return &o.SupportRules, true
}

// HasSupportRules returns a boolean if a field has been set.
func (o *ObservabilityPipelineParseGrokProcessorRule) HasSupportRules() bool {
	return o != nil && o.SupportRules != nil
}

// SetSupportRules gets a reference to the given []ObservabilityPipelineParseGrokProcessorRuleSupportRule and assigns it to the SupportRules field.
func (o *ObservabilityPipelineParseGrokProcessorRule) SetSupportRules(v []ObservabilityPipelineParseGrokProcessorRuleSupportRule) {
	o.SupportRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineParseGrokProcessorRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["match_rules"] = o.MatchRules
	toSerialize["source"] = o.Source
	if o.SupportRules != nil {
		toSerialize["support_rules"] = o.SupportRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineParseGrokProcessorRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MatchRules   *[]ObservabilityPipelineParseGrokProcessorRuleMatchRule  `json:"match_rules"`
		Source       *string                                                  `json:"source"`
		SupportRules []ObservabilityPipelineParseGrokProcessorRuleSupportRule `json:"support_rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MatchRules == nil {
		return fmt.Errorf("required field match_rules missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"match_rules", "source", "support_rules"})
	} else {
		return err
	}
	o.MatchRules = *all.MatchRules
	o.Source = *all.Source
	o.SupportRules = all.SupportRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
