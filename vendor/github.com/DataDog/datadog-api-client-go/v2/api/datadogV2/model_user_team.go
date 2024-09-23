// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeam A user's relationship with a team
type UserTeam struct {
	// Team membership attributes
	Attributes *UserTeamAttributes `json:"attributes,omitempty"`
	// The ID of a user's relationship with a team
	Id string `json:"id"`
	// Relationship between membership and a user
	Relationships *UserTeamRelationships `json:"relationships,omitempty"`
	// Team membership type
	Type UserTeamType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserTeam instantiates a new UserTeam object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserTeam(id string, typeVar UserTeamType) *UserTeam {
	this := UserTeam{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewUserTeamWithDefaults instantiates a new UserTeam object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserTeamWithDefaults() *UserTeam {
	this := UserTeam{}
	var typeVar UserTeamType = USERTEAMTYPE_TEAM_MEMBERSHIPS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UserTeam) GetAttributes() UserTeamAttributes {
	if o == nil || o.Attributes == nil {
		var ret UserTeamAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeam) GetAttributesOk() (*UserTeamAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UserTeam) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UserTeamAttributes and assigns it to the Attributes field.
func (o *UserTeam) SetAttributes(v UserTeamAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *UserTeam) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserTeam) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *UserTeam) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *UserTeam) GetRelationships() UserTeamRelationships {
	if o == nil || o.Relationships == nil {
		var ret UserTeamRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeam) GetRelationshipsOk() (*UserTeamRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *UserTeam) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given UserTeamRelationships and assigns it to the Relationships field.
func (o *UserTeam) SetRelationships(v UserTeamRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *UserTeam) GetType() UserTeamType {
	if o == nil {
		var ret UserTeamType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserTeam) GetTypeOk() (*UserTeamType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UserTeam) SetType(v UserTeamType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserTeam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserTeam) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *UserTeamAttributes    `json:"attributes,omitempty"`
		Id            *string                `json:"id"`
		Relationships *UserTeamRelationships `json:"relationships,omitempty"`
		Type          *UserTeamType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = *all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
