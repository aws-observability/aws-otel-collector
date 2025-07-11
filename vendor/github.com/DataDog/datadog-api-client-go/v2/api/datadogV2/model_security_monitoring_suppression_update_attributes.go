// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSuppressionUpdateAttributes The suppression rule properties to be updated.
type SecurityMonitoringSuppressionUpdateAttributes struct {
	// An exclusion query on the input data of the security rules, which could be logs, Agent events, or other types of data based on the security rule. Events matching this query are ignored by any detection rules referenced in the suppression rule.
	DataExclusionQuery *string `json:"data_exclusion_query,omitempty"`
	// A description for the suppression rule.
	Description *string `json:"description,omitempty"`
	// Whether the suppression rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// A Unix millisecond timestamp giving an expiration date for the suppression rule. After this date, it won't suppress signals anymore. If unset, the expiration date of the suppression rule is left untouched. If set to `null`, the expiration date is removed.
	ExpirationDate datadog.NullableInt64 `json:"expiration_date,omitempty"`
	// The name of the suppression rule.
	Name *string `json:"name,omitempty"`
	// The rule query of the suppression rule, with the same syntax as the search bar for detection rules.
	RuleQuery *string `json:"rule_query,omitempty"`
	// A Unix millisecond timestamp giving the start date for the suppression rule. After this date, it starts suppressing signals. If unset, the start date of the suppression rule is left untouched. If set to `null`, the start date is removed.
	StartDate datadog.NullableInt64 `json:"start_date,omitempty"`
	// The suppression query of the suppression rule. If a signal matches this query, it is suppressed and not triggered. Same syntax as the queries to search signals in the signal explorer.
	SuppressionQuery *string `json:"suppression_query,omitempty"`
	// The current version of the suppression. This is optional, but it can help prevent concurrent modifications.
	Version *int32 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSuppressionUpdateAttributes instantiates a new SecurityMonitoringSuppressionUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSuppressionUpdateAttributes() *SecurityMonitoringSuppressionUpdateAttributes {
	this := SecurityMonitoringSuppressionUpdateAttributes{}
	return &this
}

// NewSecurityMonitoringSuppressionUpdateAttributesWithDefaults instantiates a new SecurityMonitoringSuppressionUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSuppressionUpdateAttributesWithDefaults() *SecurityMonitoringSuppressionUpdateAttributes {
	this := SecurityMonitoringSuppressionUpdateAttributes{}
	return &this
}

// GetDataExclusionQuery returns the DataExclusionQuery field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetDataExclusionQuery() string {
	if o == nil || o.DataExclusionQuery == nil {
		var ret string
		return ret
	}
	return *o.DataExclusionQuery
}

// GetDataExclusionQueryOk returns a tuple with the DataExclusionQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetDataExclusionQueryOk() (*string, bool) {
	if o == nil || o.DataExclusionQuery == nil {
		return nil, false
	}
	return o.DataExclusionQuery, true
}

// HasDataExclusionQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) HasDataExclusionQuery() bool {
	return o != nil && o.DataExclusionQuery != nil
}

// SetDataExclusionQuery gets a reference to the given string and assigns it to the DataExclusionQuery field.
func (o *SecurityMonitoringSuppressionUpdateAttributes) SetDataExclusionQuery(v string) {
	o.DataExclusionQuery = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityMonitoringSuppressionUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SecurityMonitoringSuppressionUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetExpirationDate() int64 {
	if o == nil || o.ExpirationDate.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetExpirationDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate.IsSet()
}

// SetExpirationDate gets a reference to the given datadog.NullableInt64 and assigns it to the ExpirationDate field.
func (o *SecurityMonitoringSuppressionUpdateAttributes) SetExpirationDate(v int64) {
	o.ExpirationDate.Set(&v)
}

// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil.
func (o *SecurityMonitoringSuppressionUpdateAttributes) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil.
func (o *SecurityMonitoringSuppressionUpdateAttributes) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringSuppressionUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetRuleQuery returns the RuleQuery field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetRuleQuery() string {
	if o == nil || o.RuleQuery == nil {
		var ret string
		return ret
	}
	return *o.RuleQuery
}

// GetRuleQueryOk returns a tuple with the RuleQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetRuleQueryOk() (*string, bool) {
	if o == nil || o.RuleQuery == nil {
		return nil, false
	}
	return o.RuleQuery, true
}

// HasRuleQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) HasRuleQuery() bool {
	return o != nil && o.RuleQuery != nil
}

// SetRuleQuery gets a reference to the given string and assigns it to the RuleQuery field.
func (o *SecurityMonitoringSuppressionUpdateAttributes) SetRuleQuery(v string) {
	o.RuleQuery = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetStartDate() int64 {
	if o == nil || o.StartDate.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetStartDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) HasStartDate() bool {
	return o != nil && o.StartDate.IsSet()
}

// SetStartDate gets a reference to the given datadog.NullableInt64 and assigns it to the StartDate field.
func (o *SecurityMonitoringSuppressionUpdateAttributes) SetStartDate(v int64) {
	o.StartDate.Set(&v)
}

// SetStartDateNil sets the value for StartDate to be an explicit nil.
func (o *SecurityMonitoringSuppressionUpdateAttributes) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil.
func (o *SecurityMonitoringSuppressionUpdateAttributes) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetSuppressionQuery returns the SuppressionQuery field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetSuppressionQuery() string {
	if o == nil || o.SuppressionQuery == nil {
		var ret string
		return ret
	}
	return *o.SuppressionQuery
}

// GetSuppressionQueryOk returns a tuple with the SuppressionQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetSuppressionQueryOk() (*string, bool) {
	if o == nil || o.SuppressionQuery == nil {
		return nil, false
	}
	return o.SuppressionQuery, true
}

// HasSuppressionQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) HasSuppressionQuery() bool {
	return o != nil && o.SuppressionQuery != nil
}

// SetSuppressionQuery gets a reference to the given string and assigns it to the SuppressionQuery field.
func (o *SecurityMonitoringSuppressionUpdateAttributes) SetSuppressionQuery(v string) {
	o.SuppressionQuery = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionUpdateAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SecurityMonitoringSuppressionUpdateAttributes) SetVersion(v int32) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSuppressionUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DataExclusionQuery != nil {
		toSerialize["data_exclusion_query"] = o.DataExclusionQuery
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExpirationDate.IsSet() {
		toSerialize["expiration_date"] = o.ExpirationDate.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RuleQuery != nil {
		toSerialize["rule_query"] = o.RuleQuery
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.SuppressionQuery != nil {
		toSerialize["suppression_query"] = o.SuppressionQuery
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
func (o *SecurityMonitoringSuppressionUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataExclusionQuery *string               `json:"data_exclusion_query,omitempty"`
		Description        *string               `json:"description,omitempty"`
		Enabled            *bool                 `json:"enabled,omitempty"`
		ExpirationDate     datadog.NullableInt64 `json:"expiration_date,omitempty"`
		Name               *string               `json:"name,omitempty"`
		RuleQuery          *string               `json:"rule_query,omitempty"`
		StartDate          datadog.NullableInt64 `json:"start_date,omitempty"`
		SuppressionQuery   *string               `json:"suppression_query,omitempty"`
		Version            *int32                `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_exclusion_query", "description", "enabled", "expiration_date", "name", "rule_query", "start_date", "suppression_query", "version"})
	} else {
		return err
	}
	o.DataExclusionQuery = all.DataExclusionQuery
	o.Description = all.Description
	o.Enabled = all.Enabled
	o.ExpirationDate = all.ExpirationDate
	o.Name = all.Name
	o.RuleQuery = all.RuleQuery
	o.StartDate = all.StartDate
	o.SuppressionQuery = all.SuppressionQuery
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
