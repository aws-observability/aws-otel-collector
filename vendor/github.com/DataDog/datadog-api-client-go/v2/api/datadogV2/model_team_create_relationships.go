// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamCreateRelationships Relationships formed with the team on creation
type TeamCreateRelationships struct {
	// Relationship to users.
	Users *RelationshipToUsers `json:"users,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewTeamCreateRelationships instantiates a new TeamCreateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamCreateRelationships() *TeamCreateRelationships {
	this := TeamCreateRelationships{}
	return &this
}

// NewTeamCreateRelationshipsWithDefaults instantiates a new TeamCreateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamCreateRelationshipsWithDefaults() *TeamCreateRelationships {
	this := TeamCreateRelationships{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *TeamCreateRelationships) GetUsers() RelationshipToUsers {
	if o == nil || o.Users == nil {
		var ret RelationshipToUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamCreateRelationships) GetUsersOk() (*RelationshipToUsers, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *TeamCreateRelationships) HasUsers() bool {
	return o != nil && o.Users != nil
}

// SetUsers gets a reference to the given RelationshipToUsers and assigns it to the Users field.
func (o *TeamCreateRelationships) SetUsers(v RelationshipToUsers) {
	o.Users = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamCreateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Users *RelationshipToUsers `json:"users,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"users"})
	} else {
		return err
	}
	if all.Users != nil && all.Users.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Users = all.Users
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
