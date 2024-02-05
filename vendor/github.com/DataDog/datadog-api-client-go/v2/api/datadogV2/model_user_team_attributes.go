// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeamAttributes Team membership attributes
type UserTeamAttributes struct {
	// The user's role within the team
	Role NullableUserTeamRole `json:"role,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUserTeamAttributes instantiates a new UserTeamAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserTeamAttributes() *UserTeamAttributes {
	this := UserTeamAttributes{}
	return &this
}

// NewUserTeamAttributesWithDefaults instantiates a new UserTeamAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserTeamAttributesWithDefaults() *UserTeamAttributes {
	this := UserTeamAttributes{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTeamAttributes) GetRole() UserTeamRole {
	if o == nil || o.Role.Get() == nil {
		var ret UserTeamRole
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UserTeamAttributes) GetRoleOk() (*UserTeamRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *UserTeamAttributes) HasRole() bool {
	return o != nil && o.Role.IsSet()
}

// SetRole gets a reference to the given NullableUserTeamRole and assigns it to the Role field.
func (o *UserTeamAttributes) SetRole(v UserTeamRole) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil.
func (o *UserTeamAttributes) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil.
func (o *UserTeamAttributes) UnsetRole() {
	o.Role.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o UserTeamAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserTeamAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role NullableUserTeamRole `json:"role,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"role"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Role.Get() != nil && !all.Role.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.Role = all.Role
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
