// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthNMappingRelationships All relationships associated with AuthN Mapping.
type AuthNMappingRelationships struct {
	// Relationship to role.
	Role *RelationshipToRole `json:"role,omitempty"`
	// AuthN Mapping relationship to SAML Assertion Attribute.
	SamlAssertionAttribute *RelationshipToSAMLAssertionAttribute `json:"saml_assertion_attribute,omitempty"`
	// Relationship to team.
	Team *RelationshipToTeam `json:"team,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAuthNMappingRelationships instantiates a new AuthNMappingRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthNMappingRelationships() *AuthNMappingRelationships {
	this := AuthNMappingRelationships{}
	return &this
}

// NewAuthNMappingRelationshipsWithDefaults instantiates a new AuthNMappingRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAuthNMappingRelationshipsWithDefaults() *AuthNMappingRelationships {
	this := AuthNMappingRelationships{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AuthNMappingRelationships) GetRole() RelationshipToRole {
	if o == nil || o.Role == nil {
		var ret RelationshipToRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingRelationships) GetRoleOk() (*RelationshipToRole, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AuthNMappingRelationships) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given RelationshipToRole and assigns it to the Role field.
func (o *AuthNMappingRelationships) SetRole(v RelationshipToRole) {
	o.Role = &v
}

// GetSamlAssertionAttribute returns the SamlAssertionAttribute field value if set, zero value otherwise.
func (o *AuthNMappingRelationships) GetSamlAssertionAttribute() RelationshipToSAMLAssertionAttribute {
	if o == nil || o.SamlAssertionAttribute == nil {
		var ret RelationshipToSAMLAssertionAttribute
		return ret
	}
	return *o.SamlAssertionAttribute
}

// GetSamlAssertionAttributeOk returns a tuple with the SamlAssertionAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingRelationships) GetSamlAssertionAttributeOk() (*RelationshipToSAMLAssertionAttribute, bool) {
	if o == nil || o.SamlAssertionAttribute == nil {
		return nil, false
	}
	return o.SamlAssertionAttribute, true
}

// HasSamlAssertionAttribute returns a boolean if a field has been set.
func (o *AuthNMappingRelationships) HasSamlAssertionAttribute() bool {
	return o != nil && o.SamlAssertionAttribute != nil
}

// SetSamlAssertionAttribute gets a reference to the given RelationshipToSAMLAssertionAttribute and assigns it to the SamlAssertionAttribute field.
func (o *AuthNMappingRelationships) SetSamlAssertionAttribute(v RelationshipToSAMLAssertionAttribute) {
	o.SamlAssertionAttribute = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *AuthNMappingRelationships) GetTeam() RelationshipToTeam {
	if o == nil || o.Team == nil {
		var ret RelationshipToTeam
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingRelationships) GetTeamOk() (*RelationshipToTeam, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *AuthNMappingRelationships) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given RelationshipToTeam and assigns it to the Team field.
func (o *AuthNMappingRelationships) SetTeam(v RelationshipToTeam) {
	o.Team = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AuthNMappingRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.SamlAssertionAttribute != nil {
		toSerialize["saml_assertion_attribute"] = o.SamlAssertionAttribute
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
func (o *AuthNMappingRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role                   *RelationshipToRole                   `json:"role,omitempty"`
		SamlAssertionAttribute *RelationshipToSAMLAssertionAttribute `json:"saml_assertion_attribute,omitempty"`
		Team                   *RelationshipToTeam                   `json:"team,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"role", "saml_assertion_attribute", "team"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Role != nil && all.Role.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Role = all.Role
	if all.SamlAssertionAttribute != nil && all.SamlAssertionAttribute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SamlAssertionAttribute = all.SamlAssertionAttribute
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
