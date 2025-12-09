// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamRelationships Resources related to a team
type TeamRelationships struct {
	// Relationship between a team and a team link
	TeamLinks *RelationshipToTeamLinks `json:"team_links,omitempty"`
	// Relationship between a user team permission and a team
	UserTeamPermissions *RelationshipToUserTeamPermission `json:"user_team_permissions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamRelationships instantiates a new TeamRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamRelationships() *TeamRelationships {
	this := TeamRelationships{}
	return &this
}

// NewTeamRelationshipsWithDefaults instantiates a new TeamRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamRelationshipsWithDefaults() *TeamRelationships {
	this := TeamRelationships{}
	return &this
}

// GetTeamLinks returns the TeamLinks field value if set, zero value otherwise.
func (o *TeamRelationships) GetTeamLinks() RelationshipToTeamLinks {
	if o == nil || o.TeamLinks == nil {
		var ret RelationshipToTeamLinks
		return ret
	}
	return *o.TeamLinks
}

// GetTeamLinksOk returns a tuple with the TeamLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRelationships) GetTeamLinksOk() (*RelationshipToTeamLinks, bool) {
	if o == nil || o.TeamLinks == nil {
		return nil, false
	}
	return o.TeamLinks, true
}

// HasTeamLinks returns a boolean if a field has been set.
func (o *TeamRelationships) HasTeamLinks() bool {
	return o != nil && o.TeamLinks != nil
}

// SetTeamLinks gets a reference to the given RelationshipToTeamLinks and assigns it to the TeamLinks field.
func (o *TeamRelationships) SetTeamLinks(v RelationshipToTeamLinks) {
	o.TeamLinks = &v
}

// GetUserTeamPermissions returns the UserTeamPermissions field value if set, zero value otherwise.
func (o *TeamRelationships) GetUserTeamPermissions() RelationshipToUserTeamPermission {
	if o == nil || o.UserTeamPermissions == nil {
		var ret RelationshipToUserTeamPermission
		return ret
	}
	return *o.UserTeamPermissions
}

// GetUserTeamPermissionsOk returns a tuple with the UserTeamPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRelationships) GetUserTeamPermissionsOk() (*RelationshipToUserTeamPermission, bool) {
	if o == nil || o.UserTeamPermissions == nil {
		return nil, false
	}
	return o.UserTeamPermissions, true
}

// HasUserTeamPermissions returns a boolean if a field has been set.
func (o *TeamRelationships) HasUserTeamPermissions() bool {
	return o != nil && o.UserTeamPermissions != nil
}

// SetUserTeamPermissions gets a reference to the given RelationshipToUserTeamPermission and assigns it to the UserTeamPermissions field.
func (o *TeamRelationships) SetUserTeamPermissions(v RelationshipToUserTeamPermission) {
	o.UserTeamPermissions = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TeamLinks != nil {
		toSerialize["team_links"] = o.TeamLinks
	}
	if o.UserTeamPermissions != nil {
		toSerialize["user_team_permissions"] = o.UserTeamPermissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TeamLinks           *RelationshipToTeamLinks          `json:"team_links,omitempty"`
		UserTeamPermissions *RelationshipToUserTeamPermission `json:"user_team_permissions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"team_links", "user_team_permissions"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.TeamLinks != nil && all.TeamLinks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TeamLinks = all.TeamLinks
	if all.UserTeamPermissions != nil && all.UserTeamPermissions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UserTeamPermissions = all.UserTeamPermissions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
