// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisResponseDataAttributes The attributes of the analysis response, containing rule results and any top-level errors.
type AnalysisResponseDataAttributes struct {
	// Top-level error messages encountered during the analysis operation.
	Errors []string `json:"errors"`
	// The list of results for each static analysis rule applied during analysis.
	RuleResponses []AnalysisRuleResponse `json:"rule_responses"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnalysisResponseDataAttributes instantiates a new AnalysisResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisResponseDataAttributes(errors []string, ruleResponses []AnalysisRuleResponse) *AnalysisResponseDataAttributes {
	this := AnalysisResponseDataAttributes{}
	this.Errors = errors
	this.RuleResponses = ruleResponses
	return &this
}

// NewAnalysisResponseDataAttributesWithDefaults instantiates a new AnalysisResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisResponseDataAttributesWithDefaults() *AnalysisResponseDataAttributes {
	this := AnalysisResponseDataAttributes{}
	return &this
}

// GetErrors returns the Errors field value.
func (o *AnalysisResponseDataAttributes) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *AnalysisResponseDataAttributes) GetErrorsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value.
func (o *AnalysisResponseDataAttributes) SetErrors(v []string) {
	o.Errors = v
}

// GetRuleResponses returns the RuleResponses field value.
func (o *AnalysisResponseDataAttributes) GetRuleResponses() []AnalysisRuleResponse {
	if o == nil {
		var ret []AnalysisRuleResponse
		return ret
	}
	return o.RuleResponses
}

// GetRuleResponsesOk returns a tuple with the RuleResponses field value
// and a boolean to check if the value has been set.
func (o *AnalysisResponseDataAttributes) GetRuleResponsesOk() (*[]AnalysisRuleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleResponses, true
}

// SetRuleResponses sets field value.
func (o *AnalysisResponseDataAttributes) SetRuleResponses(v []AnalysisRuleResponse) {
	o.RuleResponses = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["errors"] = o.Errors
	toSerialize["rule_responses"] = o.RuleResponses

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Errors        *[]string               `json:"errors"`
		RuleResponses *[]AnalysisRuleResponse `json:"rule_responses"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Errors == nil {
		return fmt.Errorf("required field errors missing")
	}
	if all.RuleResponses == nil {
		return fmt.Errorf("required field rule_responses missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"errors", "rule_responses"})
	} else {
		return err
	}
	o.Errors = *all.Errors
	o.RuleResponses = *all.RuleResponses

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
