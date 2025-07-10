// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamRoutingRulesRequestDataAttributes Represents the attributes of a request to update or create team routing rules.
type TeamRoutingRulesRequestDataAttributes struct {
	// A list of routing rule items that define how incoming pages should be handled.
	Rules []TeamRoutingRulesRequestRule `json:"rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamRoutingRulesRequestDataAttributes instantiates a new TeamRoutingRulesRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamRoutingRulesRequestDataAttributes() *TeamRoutingRulesRequestDataAttributes {
	this := TeamRoutingRulesRequestDataAttributes{}
	return &this
}

// NewTeamRoutingRulesRequestDataAttributesWithDefaults instantiates a new TeamRoutingRulesRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamRoutingRulesRequestDataAttributesWithDefaults() *TeamRoutingRulesRequestDataAttributes {
	this := TeamRoutingRulesRequestDataAttributes{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *TeamRoutingRulesRequestDataAttributes) GetRules() []TeamRoutingRulesRequestRule {
	if o == nil || o.Rules == nil {
		var ret []TeamRoutingRulesRequestRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRoutingRulesRequestDataAttributes) GetRulesOk() (*[]TeamRoutingRulesRequestRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *TeamRoutingRulesRequestDataAttributes) HasRules() bool {
	return o != nil && o.Rules != nil
}

// SetRules gets a reference to the given []TeamRoutingRulesRequestRule and assigns it to the Rules field.
func (o *TeamRoutingRulesRequestDataAttributes) SetRules(v []TeamRoutingRulesRequestRule) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamRoutingRulesRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamRoutingRulesRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rules []TeamRoutingRulesRequestRule `json:"rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rules"})
	} else {
		return err
	}
	o.Rules = all.Rules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
