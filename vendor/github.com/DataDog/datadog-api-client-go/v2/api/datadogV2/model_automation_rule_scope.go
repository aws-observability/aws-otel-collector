// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleScope Defines the scope of findings to which the automation rule applies.
type AutomationRuleScope struct {
	// The list of security finding types that the automation rule applies to.
	FindingTypes []SecurityFindingType `json:"finding_types"`
	// A search query to further filter the findings matched by this rule. The `@workflow.*` namespace and `@status` fields are not permitted. For a reference of available fields, see the [Security Findings schema documentation](https://docs.datadoghq.com/security/guide/findings-schema/).
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutomationRuleScope instantiates a new AutomationRuleScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutomationRuleScope(findingTypes []SecurityFindingType) *AutomationRuleScope {
	this := AutomationRuleScope{}
	this.FindingTypes = findingTypes
	return &this
}

// NewAutomationRuleScopeWithDefaults instantiates a new AutomationRuleScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutomationRuleScopeWithDefaults() *AutomationRuleScope {
	this := AutomationRuleScope{}
	return &this
}

// GetFindingTypes returns the FindingTypes field value.
func (o *AutomationRuleScope) GetFindingTypes() []SecurityFindingType {
	if o == nil {
		var ret []SecurityFindingType
		return ret
	}
	return o.FindingTypes
}

// GetFindingTypesOk returns a tuple with the FindingTypes field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleScope) GetFindingTypesOk() (*[]SecurityFindingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FindingTypes, true
}

// SetFindingTypes sets field value.
func (o *AutomationRuleScope) SetFindingTypes(v []SecurityFindingType) {
	o.FindingTypes = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *AutomationRuleScope) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleScope) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *AutomationRuleScope) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *AutomationRuleScope) SetQuery(v string) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutomationRuleScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["finding_types"] = o.FindingTypes
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutomationRuleScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FindingTypes *[]SecurityFindingType `json:"finding_types"`
		Query        *string                `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FindingTypes == nil {
		return fmt.Errorf("required field finding_types missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"finding_types", "query"})
	} else {
		return err
	}
	o.FindingTypes = *all.FindingTypes
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
