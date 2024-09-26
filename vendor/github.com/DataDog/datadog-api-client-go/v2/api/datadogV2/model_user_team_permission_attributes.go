// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeamPermissionAttributes User team permission attributes
type UserTeamPermissionAttributes struct {
	// Object of team permission actions and boolean values that a logged in user can perform on this team.
	Permissions interface{} `json:"permissions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserTeamPermissionAttributes instantiates a new UserTeamPermissionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserTeamPermissionAttributes() *UserTeamPermissionAttributes {
	this := UserTeamPermissionAttributes{}
	return &this
}

// NewUserTeamPermissionAttributesWithDefaults instantiates a new UserTeamPermissionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserTeamPermissionAttributesWithDefaults() *UserTeamPermissionAttributes {
	this := UserTeamPermissionAttributes{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UserTeamPermissionAttributes) GetPermissions() interface{} {
	if o == nil || o.Permissions == nil {
		var ret interface{}
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeamPermissionAttributes) GetPermissionsOk() (*interface{}, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UserTeamPermissionAttributes) HasPermissions() bool {
	return o != nil && o.Permissions != nil
}

// SetPermissions gets a reference to the given interface{} and assigns it to the Permissions field.
func (o *UserTeamPermissionAttributes) SetPermissions(v interface{}) {
	o.Permissions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserTeamPermissionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserTeamPermissionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Permissions interface{} `json:"permissions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"permissions"})
	} else {
		return err
	}
	o.Permissions = all.Permissions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
