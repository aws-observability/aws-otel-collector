// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleTestRequest Test the rule queries of a rule (rule property is ignored when applied to an existing rule)
type SecurityMonitoringRuleTestRequest struct {
	// Test a rule.
	Rule *SecurityMonitoringRuleTestPayload `json:"rule,omitempty"`
	// Data payloads used to test rules query with the expected result.
	RuleQueryPayloads []SecurityMonitoringRuleQueryPayload `json:"ruleQueryPayloads,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleTestRequest instantiates a new SecurityMonitoringRuleTestRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleTestRequest() *SecurityMonitoringRuleTestRequest {
	this := SecurityMonitoringRuleTestRequest{}
	return &this
}

// NewSecurityMonitoringRuleTestRequestWithDefaults instantiates a new SecurityMonitoringRuleTestRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleTestRequestWithDefaults() *SecurityMonitoringRuleTestRequest {
	this := SecurityMonitoringRuleTestRequest{}
	return &this
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleTestRequest) GetRule() SecurityMonitoringRuleTestPayload {
	if o == nil || o.Rule == nil {
		var ret SecurityMonitoringRuleTestPayload
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleTestRequest) GetRuleOk() (*SecurityMonitoringRuleTestPayload, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleTestRequest) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given SecurityMonitoringRuleTestPayload and assigns it to the Rule field.
func (o *SecurityMonitoringRuleTestRequest) SetRule(v SecurityMonitoringRuleTestPayload) {
	o.Rule = &v
}

// GetRuleQueryPayloads returns the RuleQueryPayloads field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleTestRequest) GetRuleQueryPayloads() []SecurityMonitoringRuleQueryPayload {
	if o == nil || o.RuleQueryPayloads == nil {
		var ret []SecurityMonitoringRuleQueryPayload
		return ret
	}
	return o.RuleQueryPayloads
}

// GetRuleQueryPayloadsOk returns a tuple with the RuleQueryPayloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleTestRequest) GetRuleQueryPayloadsOk() (*[]SecurityMonitoringRuleQueryPayload, bool) {
	if o == nil || o.RuleQueryPayloads == nil {
		return nil, false
	}
	return &o.RuleQueryPayloads, true
}

// HasRuleQueryPayloads returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleTestRequest) HasRuleQueryPayloads() bool {
	return o != nil && o.RuleQueryPayloads != nil
}

// SetRuleQueryPayloads gets a reference to the given []SecurityMonitoringRuleQueryPayload and assigns it to the RuleQueryPayloads field.
func (o *SecurityMonitoringRuleTestRequest) SetRuleQueryPayloads(v []SecurityMonitoringRuleQueryPayload) {
	o.RuleQueryPayloads = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}
	if o.RuleQueryPayloads != nil {
		toSerialize["ruleQueryPayloads"] = o.RuleQueryPayloads
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleTestRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rule              *SecurityMonitoringRuleTestPayload   `json:"rule,omitempty"`
		RuleQueryPayloads []SecurityMonitoringRuleQueryPayload `json:"ruleQueryPayloads,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rule", "ruleQueryPayloads"})
	} else {
		return err
	}
	o.Rule = all.Rule
	o.RuleQueryPayloads = all.RuleQueryPayloads

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
