// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringCriticalAssetCreateAttributes Object containing the attributes of the critical asset to be created.
type SecurityMonitoringCriticalAssetCreateAttributes struct {
	// Whether the critical asset is enabled. Defaults to `true` if not specified.
	Enabled *bool `json:"enabled,omitempty"`
	// The query for the critical asset. It uses the same syntax as the queries to search signals in the Signals Explorer.
	Query string `json:"query"`
	// The rule query of the critical asset, with the same syntax as the search bar for detection rules. This determines which rules this critical asset will apply to.
	RuleQuery string `json:"rule_query"`
	// Severity associated with this critical asset. Either an explicit severity can be set, or the severity can be increased or decreased, or the severity can be left unchanged (no-op).
	Severity SecurityMonitoringCriticalAssetSeverity `json:"severity"`
	// List of tags associated with the critical asset.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringCriticalAssetCreateAttributes instantiates a new SecurityMonitoringCriticalAssetCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringCriticalAssetCreateAttributes(query string, ruleQuery string, severity SecurityMonitoringCriticalAssetSeverity) *SecurityMonitoringCriticalAssetCreateAttributes {
	this := SecurityMonitoringCriticalAssetCreateAttributes{}
	var enabled bool = true
	this.Enabled = &enabled
	this.Query = query
	this.RuleQuery = ruleQuery
	this.Severity = severity
	return &this
}

// NewSecurityMonitoringCriticalAssetCreateAttributesWithDefaults instantiates a new SecurityMonitoringCriticalAssetCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringCriticalAssetCreateAttributesWithDefaults() *SecurityMonitoringCriticalAssetCreateAttributes {
	this := SecurityMonitoringCriticalAssetCreateAttributes{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetQuery returns the Query field value.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) SetQuery(v string) {
	o.Query = v
}

// GetRuleQuery returns the RuleQuery field value.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) GetRuleQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleQuery
}

// GetRuleQueryOk returns a tuple with the RuleQuery field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) GetRuleQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleQuery, true
}

// SetRuleQuery sets field value.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) SetRuleQuery(v string) {
	o.RuleQuery = v
}

// GetSeverity returns the Severity field value.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) GetSeverity() SecurityMonitoringCriticalAssetSeverity {
	if o == nil {
		var ret SecurityMonitoringCriticalAssetSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) GetSeverityOk() (*SecurityMonitoringCriticalAssetSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) SetSeverity(v SecurityMonitoringCriticalAssetSeverity) {
	o.Severity = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringCriticalAssetCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["query"] = o.Query
	toSerialize["rule_query"] = o.RuleQuery
	toSerialize["severity"] = o.Severity
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringCriticalAssetCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled   *bool                                    `json:"enabled,omitempty"`
		Query     *string                                  `json:"query"`
		RuleQuery *string                                  `json:"rule_query"`
		Severity  *SecurityMonitoringCriticalAssetSeverity `json:"severity"`
		Tags      []string                                 `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.RuleQuery == nil {
		return fmt.Errorf("required field rule_query missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "query", "rule_query", "severity", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	o.Query = *all.Query
	o.RuleQuery = *all.RuleQuery
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
