// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// User User object returned by the API.
type User struct {
	// Attributes of user object returned by the API.
	Attributes *UserAttributes `json:"attributes,omitempty"`
	// ID of the user.
	Id *string `json:"id,omitempty"`
	// Relationships of the user object returned by the API.
	Relationships *UserResponseRelationships `json:"relationships,omitempty"`
	// Users resource type.
	Type *UsersType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUser instantiates a new User object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUser() *User {
	this := User{}
	var typeVar UsersType = USERSTYPE_USERS
	this.Type = &typeVar
	return &this
}

// NewUserWithDefaults instantiates a new User object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserWithDefaults() *User {
	this := User{}
	var typeVar UsersType = USERSTYPE_USERS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *User) GetAttributes() UserAttributes {
	if o == nil || o.Attributes == nil {
		var ret UserAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAttributesOk() (*UserAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *User) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UserAttributes and assigns it to the Attributes field.
func (o *User) SetAttributes(v UserAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *User) GetRelationships() UserResponseRelationships {
	if o == nil || o.Relationships == nil {
		var ret UserResponseRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRelationshipsOk() (*UserResponseRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *User) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given UserResponseRelationships and assigns it to the Relationships field.
func (o *User) SetRelationships(v UserResponseRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *User) GetType() UsersType {
	if o == nil || o.Type == nil {
		var ret UsersType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTypeOk() (*UsersType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *User) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given UsersType and assigns it to the Type field.
func (o *User) SetType(v UsersType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
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
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *User) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *UserAttributes            `json:"attributes,omitempty"`
		Id            *string                    `json:"id,omitempty"`
		Relationships *UserResponseRelationships `json:"relationships,omitempty"`
		Type          *UsersType                 `json:"type,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}