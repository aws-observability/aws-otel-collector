// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleOutcomeRelationships The JSON:API relationship to a scorecard rule.
type RuleOutcomeRelationships struct {
	// The JSON:API relationship to a scorecard outcome.
	Rule *RelationshipToOutcome `json:"rule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRuleOutcomeRelationships instantiates a new RuleOutcomeRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRuleOutcomeRelationships() *RuleOutcomeRelationships {
	this := RuleOutcomeRelationships{}
	return &this
}

// NewRuleOutcomeRelationshipsWithDefaults instantiates a new RuleOutcomeRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRuleOutcomeRelationshipsWithDefaults() *RuleOutcomeRelationships {
	this := RuleOutcomeRelationships{}
	return &this
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *RuleOutcomeRelationships) GetRule() RelationshipToOutcome {
	if o == nil || o.Rule == nil {
		var ret RelationshipToOutcome
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleOutcomeRelationships) GetRuleOk() (*RelationshipToOutcome, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *RuleOutcomeRelationships) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given RelationshipToOutcome and assigns it to the Rule field.
func (o *RuleOutcomeRelationships) SetRule(v RelationshipToOutcome) {
	o.Rule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RuleOutcomeRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *RuleOutcomeRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rule *RelationshipToOutcome `json:"rule,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Rule != nil && all.Rule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rule = all.Rule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
