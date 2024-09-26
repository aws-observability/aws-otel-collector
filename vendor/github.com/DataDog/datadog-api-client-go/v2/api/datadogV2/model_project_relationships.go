// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectRelationships Project relationships
type ProjectRelationships struct {
	// Relationship between a team and a team link
	MemberTeam *RelationshipToTeamLinks `json:"member_team,omitempty"`
	// Relationship to users.
	MemberUser *UsersRelationship `json:"member_user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectRelationships instantiates a new ProjectRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectRelationships() *ProjectRelationships {
	this := ProjectRelationships{}
	return &this
}

// NewProjectRelationshipsWithDefaults instantiates a new ProjectRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectRelationshipsWithDefaults() *ProjectRelationships {
	this := ProjectRelationships{}
	return &this
}

// GetMemberTeam returns the MemberTeam field value if set, zero value otherwise.
func (o *ProjectRelationships) GetMemberTeam() RelationshipToTeamLinks {
	if o == nil || o.MemberTeam == nil {
		var ret RelationshipToTeamLinks
		return ret
	}
	return *o.MemberTeam
}

// GetMemberTeamOk returns a tuple with the MemberTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRelationships) GetMemberTeamOk() (*RelationshipToTeamLinks, bool) {
	if o == nil || o.MemberTeam == nil {
		return nil, false
	}
	return o.MemberTeam, true
}

// HasMemberTeam returns a boolean if a field has been set.
func (o *ProjectRelationships) HasMemberTeam() bool {
	return o != nil && o.MemberTeam != nil
}

// SetMemberTeam gets a reference to the given RelationshipToTeamLinks and assigns it to the MemberTeam field.
func (o *ProjectRelationships) SetMemberTeam(v RelationshipToTeamLinks) {
	o.MemberTeam = &v
}

// GetMemberUser returns the MemberUser field value if set, zero value otherwise.
func (o *ProjectRelationships) GetMemberUser() UsersRelationship {
	if o == nil || o.MemberUser == nil {
		var ret UsersRelationship
		return ret
	}
	return *o.MemberUser
}

// GetMemberUserOk returns a tuple with the MemberUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRelationships) GetMemberUserOk() (*UsersRelationship, bool) {
	if o == nil || o.MemberUser == nil {
		return nil, false
	}
	return o.MemberUser, true
}

// HasMemberUser returns a boolean if a field has been set.
func (o *ProjectRelationships) HasMemberUser() bool {
	return o != nil && o.MemberUser != nil
}

// SetMemberUser gets a reference to the given UsersRelationship and assigns it to the MemberUser field.
func (o *ProjectRelationships) SetMemberUser(v UsersRelationship) {
	o.MemberUser = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MemberTeam != nil {
		toSerialize["member_team"] = o.MemberTeam
	}
	if o.MemberUser != nil {
		toSerialize["member_user"] = o.MemberUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProjectRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MemberTeam *RelationshipToTeamLinks `json:"member_team,omitempty"`
		MemberUser *UsersRelationship       `json:"member_user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"member_team", "member_user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.MemberTeam != nil && all.MemberTeam.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MemberTeam = all.MemberTeam
	if all.MemberUser != nil && all.MemberUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MemberUser = all.MemberUser

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
