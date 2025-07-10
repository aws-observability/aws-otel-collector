// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleVersions A rule version with a list of updates.
type RuleVersions struct {
	// A list of changes.
	Changes []RuleVersionUpdate `json:"changes,omitempty"`
	// Create a new rule.
	Rule *SecurityMonitoringRuleResponse `json:"rule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRuleVersions instantiates a new RuleVersions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRuleVersions() *RuleVersions {
	this := RuleVersions{}
	return &this
}

// NewRuleVersionsWithDefaults instantiates a new RuleVersions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRuleVersionsWithDefaults() *RuleVersions {
	this := RuleVersions{}
	return &this
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *RuleVersions) GetChanges() []RuleVersionUpdate {
	if o == nil || o.Changes == nil {
		var ret []RuleVersionUpdate
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleVersions) GetChangesOk() (*[]RuleVersionUpdate, bool) {
	if o == nil || o.Changes == nil {
		return nil, false
	}
	return &o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *RuleVersions) HasChanges() bool {
	return o != nil && o.Changes != nil
}

// SetChanges gets a reference to the given []RuleVersionUpdate and assigns it to the Changes field.
func (o *RuleVersions) SetChanges(v []RuleVersionUpdate) {
	o.Changes = v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *RuleVersions) GetRule() SecurityMonitoringRuleResponse {
	if o == nil || o.Rule == nil {
		var ret SecurityMonitoringRuleResponse
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleVersions) GetRuleOk() (*SecurityMonitoringRuleResponse, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *RuleVersions) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given SecurityMonitoringRuleResponse and assigns it to the Rule field.
func (o *RuleVersions) SetRule(v SecurityMonitoringRuleResponse) {
	o.Rule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RuleVersions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RuleVersions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Changes []RuleVersionUpdate             `json:"changes,omitempty"`
		Rule    *SecurityMonitoringRuleResponse `json:"rule,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"changes", "rule"})
	} else {
		return err
	}
	o.Changes = all.Changes
	o.Rule = all.Rule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
