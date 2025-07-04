// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSuppressionCreateAttributes Object containing the attributes of the suppression rule to be created.
type SecurityMonitoringSuppressionCreateAttributes struct {
	// An exclusion query on the input data of the security rules, which could be logs, Agent events, or other types of data based on the security rule. Events matching this query are ignored by any detection rules referenced in the suppression rule.
	DataExclusionQuery *string `json:"data_exclusion_query,omitempty"`
	// A description for the suppression rule.
	Description *string `json:"description,omitempty"`
	// Whether the suppression rule is enabled.
	Enabled bool `json:"enabled"`
	// A Unix millisecond timestamp giving an expiration date for the suppression rule. After this date, it won't suppress signals anymore.
	ExpirationDate *int64 `json:"expiration_date,omitempty"`
	// The name of the suppression rule.
	Name string `json:"name"`
	// The rule query of the suppression rule, with the same syntax as the search bar for detection rules.
	RuleQuery string `json:"rule_query"`
	// A Unix millisecond timestamp giving the start date for the suppression rule. After this date, it starts suppressing signals.
	StartDate *int64 `json:"start_date,omitempty"`
	// The suppression query of the suppression rule. If a signal matches this query, it is suppressed and is not triggered. It uses the same syntax as the queries to search signals in the Signals Explorer.
	SuppressionQuery *string `json:"suppression_query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSuppressionCreateAttributes instantiates a new SecurityMonitoringSuppressionCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSuppressionCreateAttributes(enabled bool, name string, ruleQuery string) *SecurityMonitoringSuppressionCreateAttributes {
	this := SecurityMonitoringSuppressionCreateAttributes{}
	this.Enabled = enabled
	this.Name = name
	this.RuleQuery = ruleQuery
	return &this
}

// NewSecurityMonitoringSuppressionCreateAttributesWithDefaults instantiates a new SecurityMonitoringSuppressionCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSuppressionCreateAttributesWithDefaults() *SecurityMonitoringSuppressionCreateAttributes {
	this := SecurityMonitoringSuppressionCreateAttributes{}
	return &this
}

// GetDataExclusionQuery returns the DataExclusionQuery field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetDataExclusionQuery() string {
	if o == nil || o.DataExclusionQuery == nil {
		var ret string
		return ret
	}
	return *o.DataExclusionQuery
}

// GetDataExclusionQueryOk returns a tuple with the DataExclusionQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetDataExclusionQueryOk() (*string, bool) {
	if o == nil || o.DataExclusionQuery == nil {
		return nil, false
	}
	return o.DataExclusionQuery, true
}

// HasDataExclusionQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) HasDataExclusionQuery() bool {
	return o != nil && o.DataExclusionQuery != nil
}

// SetDataExclusionQuery gets a reference to the given string and assigns it to the DataExclusionQuery field.
func (o *SecurityMonitoringSuppressionCreateAttributes) SetDataExclusionQuery(v string) {
	o.DataExclusionQuery = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityMonitoringSuppressionCreateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *SecurityMonitoringSuppressionCreateAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetExpirationDate() int64 {
	if o == nil || o.ExpirationDate == nil {
		var ret int64
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetExpirationDateOk() (*int64, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate != nil
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *SecurityMonitoringSuppressionCreateAttributes) SetExpirationDate(v int64) {
	o.ExpirationDate = &v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringSuppressionCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetRuleQuery returns the RuleQuery field value.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetRuleQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleQuery
}

// GetRuleQueryOk returns a tuple with the RuleQuery field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetRuleQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleQuery, true
}

// SetRuleQuery sets field value.
func (o *SecurityMonitoringSuppressionCreateAttributes) SetRuleQuery(v string) {
	o.RuleQuery = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetStartDate() int64 {
	if o == nil || o.StartDate == nil {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetStartDateOk() (*int64, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *SecurityMonitoringSuppressionCreateAttributes) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetSuppressionQuery returns the SuppressionQuery field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetSuppressionQuery() string {
	if o == nil || o.SuppressionQuery == nil {
		var ret string
		return ret
	}
	return *o.SuppressionQuery
}

// GetSuppressionQueryOk returns a tuple with the SuppressionQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) GetSuppressionQueryOk() (*string, bool) {
	if o == nil || o.SuppressionQuery == nil {
		return nil, false
	}
	return o.SuppressionQuery, true
}

// HasSuppressionQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionCreateAttributes) HasSuppressionQuery() bool {
	return o != nil && o.SuppressionQuery != nil
}

// SetSuppressionQuery gets a reference to the given string and assigns it to the SuppressionQuery field.
func (o *SecurityMonitoringSuppressionCreateAttributes) SetSuppressionQuery(v string) {
	o.SuppressionQuery = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSuppressionCreateAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["enabled"] = o.Enabled
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	toSerialize["name"] = o.Name
	toSerialize["rule_query"] = o.RuleQuery
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.SuppressionQuery != nil {
		toSerialize["suppression_query"] = o.SuppressionQuery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSuppressionCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataExclusionQuery *string `json:"data_exclusion_query,omitempty"`
		Description        *string `json:"description,omitempty"`
		Enabled            *bool   `json:"enabled"`
		ExpirationDate     *int64  `json:"expiration_date,omitempty"`
		Name               *string `json:"name"`
		RuleQuery          *string `json:"rule_query"`
		StartDate          *int64  `json:"start_date,omitempty"`
		SuppressionQuery   *string `json:"suppression_query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.RuleQuery == nil {
		return fmt.Errorf("required field rule_query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_exclusion_query", "description", "enabled", "expiration_date", "name", "rule_query", "start_date", "suppression_query"})
	} else {
		return err
	}
	o.DataExclusionQuery = all.DataExclusionQuery
	o.Description = all.Description
	o.Enabled = *all.Enabled
	o.ExpirationDate = all.ExpirationDate
	o.Name = *all.Name
	o.RuleQuery = *all.RuleQuery
	o.StartDate = all.StartDate
	o.SuppressionQuery = all.SuppressionQuery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
