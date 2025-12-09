// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationshipToRule Scorecard create rule response relationship.
type RelationshipToRule struct {
	// Relationship data for a rule.
	Scorecard *RelationshipToRuleData `json:"scorecard,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRelationshipToRule instantiates a new RelationshipToRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRelationshipToRule() *RelationshipToRule {
	this := RelationshipToRule{}
	return &this
}

// NewRelationshipToRuleWithDefaults instantiates a new RelationshipToRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRelationshipToRuleWithDefaults() *RelationshipToRule {
	this := RelationshipToRule{}
	return &this
}

// GetScorecard returns the Scorecard field value if set, zero value otherwise.
func (o *RelationshipToRule) GetScorecard() RelationshipToRuleData {
	if o == nil || o.Scorecard == nil {
		var ret RelationshipToRuleData
		return ret
	}
	return *o.Scorecard
}

// GetScorecardOk returns a tuple with the Scorecard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipToRule) GetScorecardOk() (*RelationshipToRuleData, bool) {
	if o == nil || o.Scorecard == nil {
		return nil, false
	}
	return o.Scorecard, true
}

// HasScorecard returns a boolean if a field has been set.
func (o *RelationshipToRule) HasScorecard() bool {
	return o != nil && o.Scorecard != nil
}

// SetScorecard gets a reference to the given RelationshipToRuleData and assigns it to the Scorecard field.
func (o *RelationshipToRule) SetScorecard(v RelationshipToRuleData) {
	o.Scorecard = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RelationshipToRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Scorecard != nil {
		toSerialize["scorecard"] = o.Scorecard
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RelationshipToRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Scorecard *RelationshipToRuleData `json:"scorecard,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"scorecard"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Scorecard != nil && all.Scorecard.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scorecard = all.Scorecard

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
