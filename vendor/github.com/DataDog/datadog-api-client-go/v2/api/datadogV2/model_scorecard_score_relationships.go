// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScorecardScoreRelationships Relationships for a scorecard score, depending on the aggregation type.
type ScorecardScoreRelationships struct {
	// A relationship item for a score.
	Entity *ScorecardScoreRelationshipItem `json:"entity,omitempty"`
	// A relationship item for a score.
	Rule *ScorecardScoreRelationshipItem `json:"rule,omitempty"`
	// A relationship item for a score.
	Scorecard *ScorecardScoreRelationshipItem `json:"scorecard,omitempty"`
	// A relationship item for a score.
	Service *ScorecardScoreRelationshipItem `json:"service,omitempty"`
	// A relationship item for a score.
	Team *ScorecardScoreRelationshipItem `json:"team,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScorecardScoreRelationships instantiates a new ScorecardScoreRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScorecardScoreRelationships() *ScorecardScoreRelationships {
	this := ScorecardScoreRelationships{}
	return &this
}

// NewScorecardScoreRelationshipsWithDefaults instantiates a new ScorecardScoreRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScorecardScoreRelationshipsWithDefaults() *ScorecardScoreRelationships {
	this := ScorecardScoreRelationships{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *ScorecardScoreRelationships) GetEntity() ScorecardScoreRelationshipItem {
	if o == nil || o.Entity == nil {
		var ret ScorecardScoreRelationshipItem
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreRelationships) GetEntityOk() (*ScorecardScoreRelationshipItem, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *ScorecardScoreRelationships) HasEntity() bool {
	return o != nil && o.Entity != nil
}

// SetEntity gets a reference to the given ScorecardScoreRelationshipItem and assigns it to the Entity field.
func (o *ScorecardScoreRelationships) SetEntity(v ScorecardScoreRelationshipItem) {
	o.Entity = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *ScorecardScoreRelationships) GetRule() ScorecardScoreRelationshipItem {
	if o == nil || o.Rule == nil {
		var ret ScorecardScoreRelationshipItem
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreRelationships) GetRuleOk() (*ScorecardScoreRelationshipItem, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *ScorecardScoreRelationships) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given ScorecardScoreRelationshipItem and assigns it to the Rule field.
func (o *ScorecardScoreRelationships) SetRule(v ScorecardScoreRelationshipItem) {
	o.Rule = &v
}

// GetScorecard returns the Scorecard field value if set, zero value otherwise.
func (o *ScorecardScoreRelationships) GetScorecard() ScorecardScoreRelationshipItem {
	if o == nil || o.Scorecard == nil {
		var ret ScorecardScoreRelationshipItem
		return ret
	}
	return *o.Scorecard
}

// GetScorecardOk returns a tuple with the Scorecard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreRelationships) GetScorecardOk() (*ScorecardScoreRelationshipItem, bool) {
	if o == nil || o.Scorecard == nil {
		return nil, false
	}
	return o.Scorecard, true
}

// HasScorecard returns a boolean if a field has been set.
func (o *ScorecardScoreRelationships) HasScorecard() bool {
	return o != nil && o.Scorecard != nil
}

// SetScorecard gets a reference to the given ScorecardScoreRelationshipItem and assigns it to the Scorecard field.
func (o *ScorecardScoreRelationships) SetScorecard(v ScorecardScoreRelationshipItem) {
	o.Scorecard = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ScorecardScoreRelationships) GetService() ScorecardScoreRelationshipItem {
	if o == nil || o.Service == nil {
		var ret ScorecardScoreRelationshipItem
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreRelationships) GetServiceOk() (*ScorecardScoreRelationshipItem, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ScorecardScoreRelationships) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given ScorecardScoreRelationshipItem and assigns it to the Service field.
func (o *ScorecardScoreRelationships) SetService(v ScorecardScoreRelationshipItem) {
	o.Service = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *ScorecardScoreRelationships) GetTeam() ScorecardScoreRelationshipItem {
	if o == nil || o.Team == nil {
		var ret ScorecardScoreRelationshipItem
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreRelationships) GetTeamOk() (*ScorecardScoreRelationshipItem, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *ScorecardScoreRelationships) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given ScorecardScoreRelationshipItem and assigns it to the Team field.
func (o *ScorecardScoreRelationships) SetTeam(v ScorecardScoreRelationshipItem) {
	o.Team = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScorecardScoreRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}
	if o.Scorecard != nil {
		toSerialize["scorecard"] = o.Scorecard
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Team != nil {
		toSerialize["team"] = o.Team
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScorecardScoreRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Entity    *ScorecardScoreRelationshipItem `json:"entity,omitempty"`
		Rule      *ScorecardScoreRelationshipItem `json:"rule,omitempty"`
		Scorecard *ScorecardScoreRelationshipItem `json:"scorecard,omitempty"`
		Service   *ScorecardScoreRelationshipItem `json:"service,omitempty"`
		Team      *ScorecardScoreRelationshipItem `json:"team,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entity", "rule", "scorecard", "service", "team"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Entity != nil && all.Entity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Entity = all.Entity
	if all.Rule != nil && all.Rule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rule = all.Rule
	if all.Scorecard != nil && all.Scorecard.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scorecard = all.Scorecard
	if all.Service != nil && all.Service.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Service = all.Service
	if all.Team != nil && all.Team.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Team = all.Team

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
