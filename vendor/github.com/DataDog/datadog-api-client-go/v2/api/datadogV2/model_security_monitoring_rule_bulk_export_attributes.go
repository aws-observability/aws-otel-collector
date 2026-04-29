// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleBulkExportAttributes Attributes for bulk exporting security monitoring rules.
type SecurityMonitoringRuleBulkExportAttributes struct {
	// List of rule IDs to export. Each rule will be included in the resulting ZIP file
	// as a separate JSON file.
	RuleIds []string `json:"ruleIds"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleBulkExportAttributes instantiates a new SecurityMonitoringRuleBulkExportAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleBulkExportAttributes(ruleIds []string) *SecurityMonitoringRuleBulkExportAttributes {
	this := SecurityMonitoringRuleBulkExportAttributes{}
	this.RuleIds = ruleIds
	return &this
}

// NewSecurityMonitoringRuleBulkExportAttributesWithDefaults instantiates a new SecurityMonitoringRuleBulkExportAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleBulkExportAttributesWithDefaults() *SecurityMonitoringRuleBulkExportAttributes {
	this := SecurityMonitoringRuleBulkExportAttributes{}
	return &this
}

// GetRuleIds returns the RuleIds field value.
func (o *SecurityMonitoringRuleBulkExportAttributes) GetRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleBulkExportAttributes) GetRuleIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleIds, true
}

// SetRuleIds sets field value.
func (o *SecurityMonitoringRuleBulkExportAttributes) SetRuleIds(v []string) {
	o.RuleIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleBulkExportAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ruleIds"] = o.RuleIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleBulkExportAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RuleIds *[]string `json:"ruleIds"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RuleIds == nil {
		return fmt.Errorf("required field ruleIds missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ruleIds"})
	} else {
		return err
	}
	o.RuleIds = *all.RuleIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
