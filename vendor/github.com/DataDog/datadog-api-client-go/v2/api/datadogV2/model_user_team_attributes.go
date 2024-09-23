// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeamAttributes Team membership attributes
type UserTeamAttributes struct {
	// The mechanism responsible for provisioning the team relationship.
	// Possible values: null for added by a user, "service_account" if added by a service account, and "saml_mapping" if provisioned via SAML mapping.
	ProvisionedBy datadog.NullableString `json:"provisioned_by,omitempty"`
	// UUID of the User or Service Account who provisioned this team membership, or null if provisioned via SAML mapping.
	ProvisionedById datadog.NullableString `json:"provisioned_by_id,omitempty"`
	// The user's role within the team
	Role NullableUserTeamRole `json:"role,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
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

// GetProvisionedBy returns the ProvisionedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTeamAttributes) GetProvisionedBy() string {
	if o == nil || o.ProvisionedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProvisionedBy.Get()
}

// GetProvisionedByOk returns a tuple with the ProvisionedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UserTeamAttributes) GetProvisionedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvisionedBy.Get(), o.ProvisionedBy.IsSet()
}

// HasProvisionedBy returns a boolean if a field has been set.
func (o *UserTeamAttributes) HasProvisionedBy() bool {
	return o != nil && o.ProvisionedBy.IsSet()
}

// SetProvisionedBy gets a reference to the given datadog.NullableString and assigns it to the ProvisionedBy field.
func (o *UserTeamAttributes) SetProvisionedBy(v string) {
	o.ProvisionedBy.Set(&v)
}

// SetProvisionedByNil sets the value for ProvisionedBy to be an explicit nil.
func (o *UserTeamAttributes) SetProvisionedByNil() {
	o.ProvisionedBy.Set(nil)
}

// UnsetProvisionedBy ensures that no value is present for ProvisionedBy, not even an explicit nil.
func (o *UserTeamAttributes) UnsetProvisionedBy() {
	o.ProvisionedBy.Unset()
}

// GetProvisionedById returns the ProvisionedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTeamAttributes) GetProvisionedById() string {
	if o == nil || o.ProvisionedById.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProvisionedById.Get()
}

// GetProvisionedByIdOk returns a tuple with the ProvisionedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UserTeamAttributes) GetProvisionedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvisionedById.Get(), o.ProvisionedById.IsSet()
}

// HasProvisionedById returns a boolean if a field has been set.
func (o *UserTeamAttributes) HasProvisionedById() bool {
	return o != nil && o.ProvisionedById.IsSet()
}

// SetProvisionedById gets a reference to the given datadog.NullableString and assigns it to the ProvisionedById field.
func (o *UserTeamAttributes) SetProvisionedById(v string) {
	o.ProvisionedById.Set(&v)
}

// SetProvisionedByIdNil sets the value for ProvisionedById to be an explicit nil.
func (o *UserTeamAttributes) SetProvisionedByIdNil() {
	o.ProvisionedById.Set(nil)
}

// UnsetProvisionedById ensures that no value is present for ProvisionedById, not even an explicit nil.
func (o *UserTeamAttributes) UnsetProvisionedById() {
	o.ProvisionedById.Unset()
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
	if o.ProvisionedBy.IsSet() {
		toSerialize["provisioned_by"] = o.ProvisionedBy.Get()
	}
	if o.ProvisionedById.IsSet() {
		toSerialize["provisioned_by_id"] = o.ProvisionedById.Get()
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
		ProvisionedBy   datadog.NullableString `json:"provisioned_by,omitempty"`
		ProvisionedById datadog.NullableString `json:"provisioned_by_id,omitempty"`
		Role            NullableUserTeamRole   `json:"role,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"provisioned_by", "provisioned_by_id", "role"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ProvisionedBy = all.ProvisionedBy
	o.ProvisionedById = all.ProvisionedById
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
