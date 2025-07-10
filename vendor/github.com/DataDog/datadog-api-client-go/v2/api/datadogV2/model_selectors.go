// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Selectors Selectors are used to filter security issues for which notifications should be generated.
// Users can specify rule severities, rule types, a query to filter security issues on tags and attributes, and the trigger source.
// Only the trigger_source field is required.
type Selectors struct {
	// The query is composed of one or several key:value pairs, which can be used to filter security issues on tags and attributes.
	Query *string `json:"query,omitempty"`
	// Security rule types used as filters in security rules.
	RuleTypes []RuleTypesItems `json:"rule_types,omitempty"`
	// The security rules severities to consider.
	Severities []RuleSeverity `json:"severities,omitempty"`
	// The type of security issues on which the rule applies. Notification rules based on security signals need to use the trigger source "security_signals",
	// while notification rules based on security vulnerabilities need to use the trigger source "security_findings".
	TriggerSource TriggerSource `json:"trigger_source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSelectors instantiates a new Selectors object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSelectors(triggerSource TriggerSource) *Selectors {
	this := Selectors{}
	this.TriggerSource = triggerSource
	return &this
}

// NewSelectorsWithDefaults instantiates a new Selectors object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSelectorsWithDefaults() *Selectors {
	this := Selectors{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *Selectors) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selectors) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *Selectors) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *Selectors) SetQuery(v string) {
	o.Query = &v
}

// GetRuleTypes returns the RuleTypes field value if set, zero value otherwise.
func (o *Selectors) GetRuleTypes() []RuleTypesItems {
	if o == nil || o.RuleTypes == nil {
		var ret []RuleTypesItems
		return ret
	}
	return o.RuleTypes
}

// GetRuleTypesOk returns a tuple with the RuleTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selectors) GetRuleTypesOk() (*[]RuleTypesItems, bool) {
	if o == nil || o.RuleTypes == nil {
		return nil, false
	}
	return &o.RuleTypes, true
}

// HasRuleTypes returns a boolean if a field has been set.
func (o *Selectors) HasRuleTypes() bool {
	return o != nil && o.RuleTypes != nil
}

// SetRuleTypes gets a reference to the given []RuleTypesItems and assigns it to the RuleTypes field.
func (o *Selectors) SetRuleTypes(v []RuleTypesItems) {
	o.RuleTypes = v
}

// GetSeverities returns the Severities field value if set, zero value otherwise.
func (o *Selectors) GetSeverities() []RuleSeverity {
	if o == nil || o.Severities == nil {
		var ret []RuleSeverity
		return ret
	}
	return o.Severities
}

// GetSeveritiesOk returns a tuple with the Severities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selectors) GetSeveritiesOk() (*[]RuleSeverity, bool) {
	if o == nil || o.Severities == nil {
		return nil, false
	}
	return &o.Severities, true
}

// HasSeverities returns a boolean if a field has been set.
func (o *Selectors) HasSeverities() bool {
	return o != nil && o.Severities != nil
}

// SetSeverities gets a reference to the given []RuleSeverity and assigns it to the Severities field.
func (o *Selectors) SetSeverities(v []RuleSeverity) {
	o.Severities = v
}

// GetTriggerSource returns the TriggerSource field value.
func (o *Selectors) GetTriggerSource() TriggerSource {
	if o == nil {
		var ret TriggerSource
		return ret
	}
	return o.TriggerSource
}

// GetTriggerSourceOk returns a tuple with the TriggerSource field value
// and a boolean to check if the value has been set.
func (o *Selectors) GetTriggerSourceOk() (*TriggerSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerSource, true
}

// SetTriggerSource sets field value.
func (o *Selectors) SetTriggerSource(v TriggerSource) {
	o.TriggerSource = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Selectors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.RuleTypes != nil {
		toSerialize["rule_types"] = o.RuleTypes
	}
	if o.Severities != nil {
		toSerialize["severities"] = o.Severities
	}
	toSerialize["trigger_source"] = o.TriggerSource

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Selectors) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query         *string          `json:"query,omitempty"`
		RuleTypes     []RuleTypesItems `json:"rule_types,omitempty"`
		Severities    []RuleSeverity   `json:"severities,omitempty"`
		TriggerSource *TriggerSource   `json:"trigger_source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TriggerSource == nil {
		return fmt.Errorf("required field trigger_source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"query", "rule_types", "severities", "trigger_source"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Query = all.Query
	o.RuleTypes = all.RuleTypes
	o.Severities = all.Severities
	if !all.TriggerSource.IsValid() {
		hasInvalidField = true
	} else {
		o.TriggerSource = *all.TriggerSource
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
