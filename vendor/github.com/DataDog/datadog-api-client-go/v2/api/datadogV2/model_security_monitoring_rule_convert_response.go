// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleConvertResponse Result of the convert rule request containing Terraform content.
type SecurityMonitoringRuleConvertResponse struct {
	// the ID of the rule.
	RuleId *string `json:"ruleId,omitempty"`
	// Terraform string as a result of converting the rule from JSON.
	TerraformContent *string `json:"terraformContent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleConvertResponse instantiates a new SecurityMonitoringRuleConvertResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleConvertResponse() *SecurityMonitoringRuleConvertResponse {
	this := SecurityMonitoringRuleConvertResponse{}
	return &this
}

// NewSecurityMonitoringRuleConvertResponseWithDefaults instantiates a new SecurityMonitoringRuleConvertResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleConvertResponseWithDefaults() *SecurityMonitoringRuleConvertResponse {
	this := SecurityMonitoringRuleConvertResponse{}
	return &this
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleConvertResponse) GetRuleId() string {
	if o == nil || o.RuleId == nil {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleConvertResponse) GetRuleIdOk() (*string, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleConvertResponse) HasRuleId() bool {
	return o != nil && o.RuleId != nil
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *SecurityMonitoringRuleConvertResponse) SetRuleId(v string) {
	o.RuleId = &v
}

// GetTerraformContent returns the TerraformContent field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleConvertResponse) GetTerraformContent() string {
	if o == nil || o.TerraformContent == nil {
		var ret string
		return ret
	}
	return *o.TerraformContent
}

// GetTerraformContentOk returns a tuple with the TerraformContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleConvertResponse) GetTerraformContentOk() (*string, bool) {
	if o == nil || o.TerraformContent == nil {
		return nil, false
	}
	return o.TerraformContent, true
}

// HasTerraformContent returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleConvertResponse) HasTerraformContent() bool {
	return o != nil && o.TerraformContent != nil
}

// SetTerraformContent gets a reference to the given string and assigns it to the TerraformContent field.
func (o *SecurityMonitoringRuleConvertResponse) SetTerraformContent(v string) {
	o.TerraformContent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleConvertResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RuleId != nil {
		toSerialize["ruleId"] = o.RuleId
	}
	if o.TerraformContent != nil {
		toSerialize["terraformContent"] = o.TerraformContent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleConvertResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RuleId           *string `json:"ruleId,omitempty"`
		TerraformContent *string `json:"terraformContent,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ruleId", "terraformContent"})
	} else {
		return err
	}
	o.RuleId = all.RuleId
	o.TerraformContent = all.TerraformContent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
