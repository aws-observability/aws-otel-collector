// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeamPermission A user's permissions for a given team
type UserTeamPermission struct {
	// User team permission attributes
	Attributes *UserTeamPermissionAttributes `json:"attributes,omitempty"`
	// The user team permission's identifier
	Id string `json:"id"`
	// User team permission type
	Type UserTeamPermissionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserTeamPermission instantiates a new UserTeamPermission object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserTeamPermission(id string, typeVar UserTeamPermissionType) *UserTeamPermission {
	this := UserTeamPermission{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewUserTeamPermissionWithDefaults instantiates a new UserTeamPermission object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserTeamPermissionWithDefaults() *UserTeamPermission {
	this := UserTeamPermission{}
	var typeVar UserTeamPermissionType = USERTEAMPERMISSIONTYPE_USER_TEAM_PERMISSIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UserTeamPermission) GetAttributes() UserTeamPermissionAttributes {
	if o == nil || o.Attributes == nil {
		var ret UserTeamPermissionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeamPermission) GetAttributesOk() (*UserTeamPermissionAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UserTeamPermission) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UserTeamPermissionAttributes and assigns it to the Attributes field.
func (o *UserTeamPermission) SetAttributes(v UserTeamPermissionAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *UserTeamPermission) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserTeamPermission) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *UserTeamPermission) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *UserTeamPermission) GetType() UserTeamPermissionType {
	if o == nil {
		var ret UserTeamPermissionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserTeamPermission) GetTypeOk() (*UserTeamPermissionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UserTeamPermission) SetType(v UserTeamPermissionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserTeamPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserTeamPermission) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UserTeamPermissionAttributes `json:"attributes,omitempty"`
		Id         *string                       `json:"id"`
		Type       *UserTeamPermissionType       `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = *all.Id
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
