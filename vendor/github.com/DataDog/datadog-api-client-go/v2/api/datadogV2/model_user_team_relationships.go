// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeamRelationships Relationship between membership and a user
type UserTeamRelationships struct {
	// Relationship between team membership and user
	User *RelationshipToUserTeamUser `json:"user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
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
		return json.Marshal(o.UnparsedObject)
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserTeamRelationships) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		User *RelationshipToUserTeamUser `json:"user,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"user"})
	} else {
		return err
	}
	if all.User != nil && all.User.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.User = all.User
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
