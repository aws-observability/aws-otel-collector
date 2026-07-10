// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SAMLConfigurationRelationships Relationships of a SAML configuration.
type SAMLConfigurationRelationships struct {
	// Relationship to roles.
	DefaultRoles *RelationshipToRoles `json:"default_roles,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSAMLConfigurationRelationships instantiates a new SAMLConfigurationRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSAMLConfigurationRelationships() *SAMLConfigurationRelationships {
	this := SAMLConfigurationRelationships{}
	return &this
}

// NewSAMLConfigurationRelationshipsWithDefaults instantiates a new SAMLConfigurationRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSAMLConfigurationRelationshipsWithDefaults() *SAMLConfigurationRelationships {
	this := SAMLConfigurationRelationships{}
	return &this
}

// GetDefaultRoles returns the DefaultRoles field value if set, zero value otherwise.
func (o *SAMLConfigurationRelationships) GetDefaultRoles() RelationshipToRoles {
	if o == nil || o.DefaultRoles == nil {
		var ret RelationshipToRoles
		return ret
	}
	return *o.DefaultRoles
}

// GetDefaultRolesOk returns a tuple with the DefaultRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigurationRelationships) GetDefaultRolesOk() (*RelationshipToRoles, bool) {
	if o == nil || o.DefaultRoles == nil {
		return nil, false
	}
	return o.DefaultRoles, true
}

// HasDefaultRoles returns a boolean if a field has been set.
func (o *SAMLConfigurationRelationships) HasDefaultRoles() bool {
	return o != nil && o.DefaultRoles != nil
}

// SetDefaultRoles gets a reference to the given RelationshipToRoles and assigns it to the DefaultRoles field.
func (o *SAMLConfigurationRelationships) SetDefaultRoles(v RelationshipToRoles) {
	o.DefaultRoles = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SAMLConfigurationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DefaultRoles != nil {
		toSerialize["default_roles"] = o.DefaultRoles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SAMLConfigurationRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultRoles *RelationshipToRoles `json:"default_roles,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default_roles"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DefaultRoles != nil && all.DefaultRoles.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DefaultRoles = all.DefaultRoles

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
