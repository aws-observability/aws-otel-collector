// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoleCreateData Data related to the creation of a role.
type RoleCreateData struct {
	// Attributes of the created role.
	Attributes RoleCreateAttributes `json:"attributes"`
	// Relationships of the role object.
	Relationships *RoleRelationships `json:"relationships,omitempty"`
	// Roles type.
	Type *RolesType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRoleCreateData instantiates a new RoleCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRoleCreateData(attributes RoleCreateAttributes) *RoleCreateData {
	this := RoleCreateData{}
	this.Attributes = attributes
	var typeVar RolesType = ROLESTYPE_ROLES
	this.Type = &typeVar
	return &this
}

// NewRoleCreateDataWithDefaults instantiates a new RoleCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRoleCreateDataWithDefaults() *RoleCreateData {
	this := RoleCreateData{}
	var typeVar RolesType = ROLESTYPE_ROLES
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *RoleCreateData) GetAttributes() RoleCreateAttributes {
	if o == nil {
		var ret RoleCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RoleCreateData) GetAttributesOk() (*RoleCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *RoleCreateData) SetAttributes(v RoleCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *RoleCreateData) GetRelationships() RoleRelationships {
	if o == nil || o.Relationships == nil {
		var ret RoleRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCreateData) GetRelationshipsOk() (*RoleRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *RoleCreateData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given RoleRelationships and assigns it to the Relationships field.
func (o *RoleCreateData) SetRelationships(v RoleRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleCreateData) GetType() RolesType {
	if o == nil || o.Type == nil {
		var ret RolesType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCreateData) GetTypeOk() (*RolesType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleCreateData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given RolesType and assigns it to the Type field.
func (o *RoleCreateData) SetType(v RolesType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RoleCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RoleCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *RoleCreateAttributes `json:"attributes"`
		Relationships *RoleRelationships    `json:"relationships,omitempty"`
		Type          *RolesType            `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
