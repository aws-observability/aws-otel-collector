// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthNMappingRelationshipToRole Relationship of AuthN Mapping to a Role.
type AuthNMappingRelationshipToRole struct {
	// Relationship to role.
	Role RelationshipToRole `json:"role"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAuthNMappingRelationshipToRole instantiates a new AuthNMappingRelationshipToRole object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthNMappingRelationshipToRole(role RelationshipToRole) *AuthNMappingRelationshipToRole {
	this := AuthNMappingRelationshipToRole{}
	this.Role = role
	return &this
}

// NewAuthNMappingRelationshipToRoleWithDefaults instantiates a new AuthNMappingRelationshipToRole object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAuthNMappingRelationshipToRoleWithDefaults() *AuthNMappingRelationshipToRole {
	this := AuthNMappingRelationshipToRole{}
	return &this
}

// GetRole returns the Role field value.
func (o *AuthNMappingRelationshipToRole) GetRole() RelationshipToRole {
	if o == nil {
		var ret RelationshipToRole
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *AuthNMappingRelationshipToRole) GetRoleOk() (*RelationshipToRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *AuthNMappingRelationshipToRole) SetRole(v RelationshipToRole) {
	o.Role = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AuthNMappingRelationshipToRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AuthNMappingRelationshipToRole) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role *RelationshipToRole `json:"role"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"role"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Role.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Role = *all.Role

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
