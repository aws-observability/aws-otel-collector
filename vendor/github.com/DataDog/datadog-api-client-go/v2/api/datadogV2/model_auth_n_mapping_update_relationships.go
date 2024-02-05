// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthNMappingUpdateRelationships Relationship of AuthN Mapping update object to Role.
type AuthNMappingUpdateRelationships struct {
	// Relationship to role.
	Role *RelationshipToRole `json:"role,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAuthNMappingUpdateRelationships instantiates a new AuthNMappingUpdateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthNMappingUpdateRelationships() *AuthNMappingUpdateRelationships {
	this := AuthNMappingUpdateRelationships{}
	return &this
}

// NewAuthNMappingUpdateRelationshipsWithDefaults instantiates a new AuthNMappingUpdateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAuthNMappingUpdateRelationshipsWithDefaults() *AuthNMappingUpdateRelationships {
	this := AuthNMappingUpdateRelationships{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AuthNMappingUpdateRelationships) GetRole() RelationshipToRole {
	if o == nil || o.Role == nil {
		var ret RelationshipToRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingUpdateRelationships) GetRoleOk() (*RelationshipToRole, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AuthNMappingUpdateRelationships) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given RelationshipToRole and assigns it to the Role field.
func (o *AuthNMappingUpdateRelationships) SetRole(v RelationshipToRole) {
	o.Role = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AuthNMappingUpdateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AuthNMappingUpdateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role *RelationshipToRole `json:"role,omitempty"`
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
	if all.Role != nil && all.Role.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Role = all.Role

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
