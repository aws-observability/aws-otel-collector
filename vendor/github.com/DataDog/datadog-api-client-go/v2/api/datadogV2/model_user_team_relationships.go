// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeamRelationships Relationship between membership and a user
type UserTeamRelationships struct {
	// Relationship between team membership and team
	Team *RelationshipToUserTeamTeam `json:"team,omitempty"`
	// Relationship between team membership and user
	User *RelationshipToUserTeamUser `json:"user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserTeamRelationships instantiates a new UserTeamRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserTeamRelationships() *UserTeamRelationships {
	this := UserTeamRelationships{}
	return &this
}

// NewUserTeamRelationshipsWithDefaults instantiates a new UserTeamRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserTeamRelationshipsWithDefaults() *UserTeamRelationships {
	this := UserTeamRelationships{}
	return &this
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *UserTeamRelationships) GetTeam() RelationshipToUserTeamTeam {
	if o == nil || o.Team == nil {
		var ret RelationshipToUserTeamTeam
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeamRelationships) GetTeamOk() (*RelationshipToUserTeamTeam, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *UserTeamRelationships) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given RelationshipToUserTeamTeam and assigns it to the Team field.
func (o *UserTeamRelationships) SetTeam(v RelationshipToUserTeamTeam) {
	o.Team = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserTeamRelationships) GetUser() RelationshipToUserTeamUser {
	if o == nil || o.User == nil {
		var ret RelationshipToUserTeamUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeamRelationships) GetUserOk() (*RelationshipToUserTeamUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserTeamRelationships) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given RelationshipToUserTeamUser and assigns it to the User field.
func (o *UserTeamRelationships) SetUser(v RelationshipToUserTeamUser) {
	o.User = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserTeamRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Team != nil {
		toSerialize["team"] = o.Team
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserTeamRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Team *RelationshipToUserTeamTeam `json:"team,omitempty"`
		User *RelationshipToUserTeamUser `json:"user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"team", "user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Team != nil && all.Team.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Team = all.Team
	if all.User != nil && all.User.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.User = all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
