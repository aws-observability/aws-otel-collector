// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceAccessToken Datadog access token.
type ServiceAccessToken struct {
	// Attributes of an access token.
	Attributes *ServiceAccessTokenAttributes `json:"attributes,omitempty"`
	// ID of the access token.
	Id *string `json:"id,omitempty"`
	// Resources related to the access token.
	Relationships *ServiceAccessTokenRelationships `json:"relationships,omitempty"`
	// Service access tokens resource type.
	Type *ServiceAccessTokensType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceAccessToken instantiates a new ServiceAccessToken object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceAccessToken() *ServiceAccessToken {
	this := ServiceAccessToken{}
	var typeVar ServiceAccessTokensType = SERVICEACCESSTOKENSTYPE_SERVICE_ACCESS_TOKENS
	this.Type = &typeVar
	return &this
}

// NewServiceAccessTokenWithDefaults instantiates a new ServiceAccessToken object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceAccessTokenWithDefaults() *ServiceAccessToken {
	this := ServiceAccessToken{}
	var typeVar ServiceAccessTokensType = SERVICEACCESSTOKENSTYPE_SERVICE_ACCESS_TOKENS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ServiceAccessToken) GetAttributes() ServiceAccessTokenAttributes {
	if o == nil || o.Attributes == nil {
		var ret ServiceAccessTokenAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessToken) GetAttributesOk() (*ServiceAccessTokenAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ServiceAccessToken) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ServiceAccessTokenAttributes and assigns it to the Attributes field.
func (o *ServiceAccessToken) SetAttributes(v ServiceAccessTokenAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceAccessToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessToken) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceAccessToken) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceAccessToken) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ServiceAccessToken) GetRelationships() ServiceAccessTokenRelationships {
	if o == nil || o.Relationships == nil {
		var ret ServiceAccessTokenRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessToken) GetRelationshipsOk() (*ServiceAccessTokenRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ServiceAccessToken) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ServiceAccessTokenRelationships and assigns it to the Relationships field.
func (o *ServiceAccessToken) SetRelationships(v ServiceAccessTokenRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceAccessToken) GetType() ServiceAccessTokensType {
	if o == nil || o.Type == nil {
		var ret ServiceAccessTokensType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessToken) GetTypeOk() (*ServiceAccessTokensType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceAccessToken) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ServiceAccessTokensType and assigns it to the Type field.
func (o *ServiceAccessToken) SetType(v ServiceAccessTokensType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceAccessToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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
func (o *ServiceAccessToken) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ServiceAccessTokenAttributes    `json:"attributes,omitempty"`
		Id            *string                          `json:"id,omitempty"`
		Relationships *ServiceAccessTokenRelationships `json:"relationships,omitempty"`
		Type          *ServiceAccessTokensType         `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.Id = all.Id
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
