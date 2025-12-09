// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthNMappingCreateData Data for creating an AuthN Mapping.
type AuthNMappingCreateData struct {
	// Key/Value pair of attributes used for create request.
	Attributes *AuthNMappingCreateAttributes `json:"attributes,omitempty"`
	// Relationship of AuthN Mapping create object to a Role or Team.
	Relationships *AuthNMappingCreateRelationships `json:"relationships,omitempty"`
	// AuthN Mappings resource type.
	Type AuthNMappingsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAuthNMappingCreateData instantiates a new AuthNMappingCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthNMappingCreateData(typeVar AuthNMappingsType) *AuthNMappingCreateData {
	this := AuthNMappingCreateData{}
	this.Type = typeVar
	return &this
}

// NewAuthNMappingCreateDataWithDefaults instantiates a new AuthNMappingCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAuthNMappingCreateDataWithDefaults() *AuthNMappingCreateData {
	this := AuthNMappingCreateData{}
	var typeVar AuthNMappingsType = AUTHNMAPPINGSTYPE_AUTHN_MAPPINGS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AuthNMappingCreateData) GetAttributes() AuthNMappingCreateAttributes {
	if o == nil || o.Attributes == nil {
		var ret AuthNMappingCreateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingCreateData) GetAttributesOk() (*AuthNMappingCreateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AuthNMappingCreateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given AuthNMappingCreateAttributes and assigns it to the Attributes field.
func (o *AuthNMappingCreateData) SetAttributes(v AuthNMappingCreateAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AuthNMappingCreateData) GetRelationships() AuthNMappingCreateRelationships {
	if o == nil || o.Relationships == nil {
		var ret AuthNMappingCreateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingCreateData) GetRelationshipsOk() (*AuthNMappingCreateRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AuthNMappingCreateData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given AuthNMappingCreateRelationships and assigns it to the Relationships field.
func (o *AuthNMappingCreateData) SetRelationships(v AuthNMappingCreateRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *AuthNMappingCreateData) GetType() AuthNMappingsType {
	if o == nil {
		var ret AuthNMappingsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthNMappingCreateData) GetTypeOk() (*AuthNMappingsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AuthNMappingCreateData) SetType(v AuthNMappingsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AuthNMappingCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
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
func (o *AuthNMappingCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *AuthNMappingCreateAttributes    `json:"attributes,omitempty"`
		Relationships *AuthNMappingCreateRelationships `json:"relationships,omitempty"`
		Type          *AuthNMappingsType               `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
