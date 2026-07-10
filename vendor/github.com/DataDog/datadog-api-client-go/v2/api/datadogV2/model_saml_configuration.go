// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SAMLConfiguration A SAML configuration object.
type SAMLConfiguration struct {
	// Attributes of a SAML configuration.
	Attributes *SAMLConfigurationAttributes `json:"attributes,omitempty"`
	// The UUID of the SAML configuration.
	Id string `json:"id"`
	// Relationships of a SAML configuration.
	Relationships *SAMLConfigurationRelationships `json:"relationships,omitempty"`
	// SAML configurations resource type.
	Type SAMLConfigurationsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSAMLConfiguration instantiates a new SAMLConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSAMLConfiguration(id string, typeVar SAMLConfigurationsType) *SAMLConfiguration {
	this := SAMLConfiguration{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewSAMLConfigurationWithDefaults instantiates a new SAMLConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSAMLConfigurationWithDefaults() *SAMLConfiguration {
	this := SAMLConfiguration{}
	var typeVar SAMLConfigurationsType = SAMLCONFIGURATIONSTYPE_SAML_CONFIGURATIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SAMLConfiguration) GetAttributes() SAMLConfigurationAttributes {
	if o == nil || o.Attributes == nil {
		var ret SAMLConfigurationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfiguration) GetAttributesOk() (*SAMLConfigurationAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SAMLConfiguration) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given SAMLConfigurationAttributes and assigns it to the Attributes field.
func (o *SAMLConfiguration) SetAttributes(v SAMLConfigurationAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *SAMLConfiguration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SAMLConfiguration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SAMLConfiguration) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SAMLConfiguration) GetRelationships() SAMLConfigurationRelationships {
	if o == nil || o.Relationships == nil {
		var ret SAMLConfigurationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfiguration) GetRelationshipsOk() (*SAMLConfigurationRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SAMLConfiguration) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given SAMLConfigurationRelationships and assigns it to the Relationships field.
func (o *SAMLConfiguration) SetRelationships(v SAMLConfigurationRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *SAMLConfiguration) GetType() SAMLConfigurationsType {
	if o == nil {
		var ret SAMLConfigurationsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SAMLConfiguration) GetTypeOk() (*SAMLConfigurationsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SAMLConfiguration) SetType(v SAMLConfigurationsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SAMLConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
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
func (o *SAMLConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *SAMLConfigurationAttributes    `json:"attributes,omitempty"`
		Id            *string                         `json:"id"`
		Relationships *SAMLConfigurationRelationships `json:"relationships,omitempty"`
		Type          *SAMLConfigurationsType         `json:"type"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = *all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
