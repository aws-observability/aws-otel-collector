// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserInvitationResponseData Object of a user invitation returned by the API.
type UserInvitationResponseData struct {
	// Attributes of a user invitation.
	Attributes *UserInvitationDataAttributes `json:"attributes,omitempty"`
	// ID of the user invitation.
	Id *string `json:"id,omitempty"`
	// Relationships data for user invitation.
	Relationships *UserInvitationRelationships `json:"relationships,omitempty"`
	// User invitations type.
	Type *UserInvitationsType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserInvitationResponseData instantiates a new UserInvitationResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserInvitationResponseData() *UserInvitationResponseData {
	this := UserInvitationResponseData{}
	var typeVar UserInvitationsType = USERINVITATIONSTYPE_USER_INVITATIONS
	this.Type = &typeVar
	return &this
}

// NewUserInvitationResponseDataWithDefaults instantiates a new UserInvitationResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserInvitationResponseDataWithDefaults() *UserInvitationResponseData {
	this := UserInvitationResponseData{}
	var typeVar UserInvitationsType = USERINVITATIONSTYPE_USER_INVITATIONS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UserInvitationResponseData) GetAttributes() UserInvitationDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret UserInvitationDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationResponseData) GetAttributesOk() (*UserInvitationDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UserInvitationResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UserInvitationDataAttributes and assigns it to the Attributes field.
func (o *UserInvitationResponseData) SetAttributes(v UserInvitationDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserInvitationResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserInvitationResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserInvitationResponseData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *UserInvitationResponseData) GetRelationships() UserInvitationRelationships {
	if o == nil || o.Relationships == nil {
		var ret UserInvitationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationResponseData) GetRelationshipsOk() (*UserInvitationRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *UserInvitationResponseData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given UserInvitationRelationships and assigns it to the Relationships field.
func (o *UserInvitationResponseData) SetRelationships(v UserInvitationRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserInvitationResponseData) GetType() UserInvitationsType {
	if o == nil || o.Type == nil {
		var ret UserInvitationsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationResponseData) GetTypeOk() (*UserInvitationsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserInvitationResponseData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given UserInvitationsType and assigns it to the Type field.
func (o *UserInvitationResponseData) SetType(v UserInvitationsType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserInvitationResponseData) MarshalJSON() ([]byte, error) {
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
func (o *UserInvitationResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *UserInvitationDataAttributes `json:"attributes,omitempty"`
		Id            *string                       `json:"id,omitempty"`
		Relationships *UserInvitationRelationships  `json:"relationships,omitempty"`
		Type          *UserInvitationsType          `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
