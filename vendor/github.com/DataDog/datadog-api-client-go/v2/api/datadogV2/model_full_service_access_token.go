// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FullServiceAccessToken Datadog access token, including the token key.
type FullServiceAccessToken struct {
	// Attributes of a full access token, including the token key.
	Attributes *FullServiceAccessTokenAttributes `json:"attributes,omitempty"`
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

// NewFullServiceAccessToken instantiates a new FullServiceAccessToken object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFullServiceAccessToken() *FullServiceAccessToken {
	this := FullServiceAccessToken{}
	var typeVar ServiceAccessTokensType = SERVICEACCESSTOKENSTYPE_SERVICE_ACCESS_TOKENS
	this.Type = &typeVar
	return &this
}

// NewFullServiceAccessTokenWithDefaults instantiates a new FullServiceAccessToken object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFullServiceAccessTokenWithDefaults() *FullServiceAccessToken {
	this := FullServiceAccessToken{}
	var typeVar ServiceAccessTokensType = SERVICEACCESSTOKENSTYPE_SERVICE_ACCESS_TOKENS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FullServiceAccessToken) GetAttributes() FullServiceAccessTokenAttributes {
	if o == nil || o.Attributes == nil {
		var ret FullServiceAccessTokenAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullServiceAccessToken) GetAttributesOk() (*FullServiceAccessTokenAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FullServiceAccessToken) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given FullServiceAccessTokenAttributes and assigns it to the Attributes field.
func (o *FullServiceAccessToken) SetAttributes(v FullServiceAccessTokenAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullServiceAccessToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullServiceAccessToken) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullServiceAccessToken) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FullServiceAccessToken) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FullServiceAccessToken) GetRelationships() ServiceAccessTokenRelationships {
	if o == nil || o.Relationships == nil {
		var ret ServiceAccessTokenRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullServiceAccessToken) GetRelationshipsOk() (*ServiceAccessTokenRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FullServiceAccessToken) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ServiceAccessTokenRelationships and assigns it to the Relationships field.
func (o *FullServiceAccessToken) SetRelationships(v ServiceAccessTokenRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FullServiceAccessToken) GetType() ServiceAccessTokensType {
	if o == nil || o.Type == nil {
		var ret ServiceAccessTokensType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullServiceAccessToken) GetTypeOk() (*ServiceAccessTokensType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FullServiceAccessToken) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ServiceAccessTokensType and assigns it to the Type field.
func (o *FullServiceAccessToken) SetType(v ServiceAccessTokensType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FullServiceAccessToken) MarshalJSON() ([]byte, error) {
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
func (o *FullServiceAccessToken) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *FullServiceAccessTokenAttributes `json:"attributes,omitempty"`
		Id            *string                           `json:"id,omitempty"`
		Relationships *ServiceAccessTokenRelationships  `json:"relationships,omitempty"`
		Type          *ServiceAccessTokensType          `json:"type,omitempty"`
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
