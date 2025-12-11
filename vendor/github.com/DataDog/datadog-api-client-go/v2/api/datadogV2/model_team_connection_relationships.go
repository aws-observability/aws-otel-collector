// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamConnectionRelationships Relationships of the team connection.
type TeamConnectionRelationships struct {
	// Reference to a team from an external system.
	ConnectedTeam *ConnectedTeamRef `json:"connected_team,omitempty"`
	// Reference to a Datadog team.
	Team *TeamRef `json:"team,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamConnectionRelationships instantiates a new TeamConnectionRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamConnectionRelationships() *TeamConnectionRelationships {
	this := TeamConnectionRelationships{}
	return &this
}

// NewTeamConnectionRelationshipsWithDefaults instantiates a new TeamConnectionRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamConnectionRelationshipsWithDefaults() *TeamConnectionRelationships {
	this := TeamConnectionRelationships{}
	return &this
}

// GetConnectedTeam returns the ConnectedTeam field value if set, zero value otherwise.
func (o *TeamConnectionRelationships) GetConnectedTeam() ConnectedTeamRef {
	if o == nil || o.ConnectedTeam == nil {
		var ret ConnectedTeamRef
		return ret
	}
	return *o.ConnectedTeam
}

// GetConnectedTeamOk returns a tuple with the ConnectedTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamConnectionRelationships) GetConnectedTeamOk() (*ConnectedTeamRef, bool) {
	if o == nil || o.ConnectedTeam == nil {
		return nil, false
	}
	return o.ConnectedTeam, true
}

// HasConnectedTeam returns a boolean if a field has been set.
func (o *TeamConnectionRelationships) HasConnectedTeam() bool {
	return o != nil && o.ConnectedTeam != nil
}

// SetConnectedTeam gets a reference to the given ConnectedTeamRef and assigns it to the ConnectedTeam field.
func (o *TeamConnectionRelationships) SetConnectedTeam(v ConnectedTeamRef) {
	o.ConnectedTeam = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *TeamConnectionRelationships) GetTeam() TeamRef {
	if o == nil || o.Team == nil {
		var ret TeamRef
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamConnectionRelationships) GetTeamOk() (*TeamRef, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *TeamConnectionRelationships) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given TeamRef and assigns it to the Team field.
func (o *TeamConnectionRelationships) SetTeam(v TeamRef) {
	o.Team = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamConnectionRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConnectedTeam != nil {
		toSerialize["connected_team"] = o.ConnectedTeam
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
func (o *TeamConnectionRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectedTeam *ConnectedTeamRef `json:"connected_team,omitempty"`
		Team          *TeamRef          `json:"team,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connected_team", "team"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ConnectedTeam != nil && all.ConnectedTeam.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConnectedTeam = all.ConnectedTeam
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
