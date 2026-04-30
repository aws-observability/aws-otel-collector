// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringCriticalAssetUpdateAttributes The critical asset properties to be updated.
type SecurityMonitoringCriticalAssetUpdateAttributes struct {
	// Whether the critical asset is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The query for the critical asset. It uses the same syntax as the queries to search signals in the Signals Explorer.
	Query *string `json:"query,omitempty"`
	// The rule query of the critical asset, with the same syntax as the search bar for detection rules. This determines which rules this critical asset will apply to.
	RuleQuery *string `json:"rule_query,omitempty"`
	// Severity associated with this critical asset. Either an explicit severity can be set, or the severity can be increased or decreased, or the severity can be left unchanged (no-op).
	Severity *SecurityMonitoringCriticalAssetSeverity `json:"severity,omitempty"`
	// List of tags associated with the critical asset.
	Tags []string `json:"tags,omitempty"`
	// The version of the critical asset being updated. Used for optimistic locking to prevent concurrent modifications.
	Version *int32 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringCriticalAssetUpdateAttributes instantiates a new SecurityMonitoringCriticalAssetUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringCriticalAssetUpdateAttributes() *SecurityMonitoringCriticalAssetUpdateAttributes {
	this := SecurityMonitoringCriticalAssetUpdateAttributes{}
	return &this
}

// NewSecurityMonitoringCriticalAssetUpdateAttributesWithDefaults instantiates a new SecurityMonitoringCriticalAssetUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringCriticalAssetUpdateAttributesWithDefaults() *SecurityMonitoringCriticalAssetUpdateAttributes {
	this := SecurityMonitoringCriticalAssetUpdateAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetRuleQuery returns the RuleQuery field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetRuleQuery() string {
	if o == nil || o.RuleQuery == nil {
		var ret string
		return ret
	}
	return *o.RuleQuery
}

// GetRuleQueryOk returns a tuple with the RuleQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetRuleQueryOk() (*string, bool) {
	if o == nil || o.RuleQuery == nil {
		return nil, false
	}
	return o.RuleQuery, true
}

// HasRuleQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) HasRuleQuery() bool {
	return o != nil && o.RuleQuery != nil
}

// SetRuleQuery gets a reference to the given string and assigns it to the RuleQuery field.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) SetRuleQuery(v string) {
	o.RuleQuery = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetSeverity() SecurityMonitoringCriticalAssetSeverity {
	if o == nil || o.Severity == nil {
		var ret SecurityMonitoringCriticalAssetSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetSeverityOk() (*SecurityMonitoringCriticalAssetSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given SecurityMonitoringCriticalAssetSeverity and assigns it to the Severity field.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) SetSeverity(v SecurityMonitoringCriticalAssetSeverity) {
	o.Severity = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) SetVersion(v int32) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringCriticalAssetUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.RuleQuery != nil {
		toSerialize["rule_query"] = o.RuleQuery
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringCriticalAssetUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled   *bool                                    `json:"enabled,omitempty"`
		Query     *string                                  `json:"query,omitempty"`
		RuleQuery *string                                  `json:"rule_query,omitempty"`
		Severity  *SecurityMonitoringCriticalAssetSeverity `json:"severity,omitempty"`
		Tags      []string                                 `json:"tags,omitempty"`
		Version   *int32                                   `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "query", "rule_query", "severity", "tags", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	o.Query = all.Query
	o.RuleQuery = all.RuleQuery
	if all.Severity != nil && !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = all.Severity
	}
	o.Tags = all.Tags
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
