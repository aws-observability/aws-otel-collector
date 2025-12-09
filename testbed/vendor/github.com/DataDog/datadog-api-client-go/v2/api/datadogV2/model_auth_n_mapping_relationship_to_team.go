// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthNMappingRelationshipToTeam Relationship of AuthN Mapping to a Team.
type AuthNMappingRelationshipToTeam struct {
	// Relationship to team.
	Team RelationshipToTeam `json:"team"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAuthNMappingRelationshipToTeam instantiates a new AuthNMappingRelationshipToTeam object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthNMappingRelationshipToTeam(team RelationshipToTeam) *AuthNMappingRelationshipToTeam {
	this := AuthNMappingRelationshipToTeam{}
	this.Team = team
	return &this
}

// NewAuthNMappingRelationshipToTeamWithDefaults instantiates a new AuthNMappingRelationshipToTeam object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAuthNMappingRelationshipToTeamWithDefaults() *AuthNMappingRelationshipToTeam {
	this := AuthNMappingRelationshipToTeam{}
	return &this
}

// GetTeam returns the Team field value.
func (o *AuthNMappingRelationshipToTeam) GetTeam() RelationshipToTeam {
	if o == nil {
		var ret RelationshipToTeam
		return ret
	}
	return o.Team
}

// GetTeamOk returns a tuple with the Team field value
// and a boolean to check if the value has been set.
func (o *AuthNMappingRelationshipToTeam) GetTeamOk() (*RelationshipToTeam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Team, true
}

// SetTeam sets field value.
func (o *AuthNMappingRelationshipToTeam) SetTeam(v RelationshipToTeam) {
	o.Team = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AuthNMappingRelationshipToTeam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["team"] = o.Team

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AuthNMappingRelationshipToTeam) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Team *RelationshipToTeam `json:"team"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Team == nil {
		return fmt.Errorf("required field team missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"team"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Team.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Team = *all.Team

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
