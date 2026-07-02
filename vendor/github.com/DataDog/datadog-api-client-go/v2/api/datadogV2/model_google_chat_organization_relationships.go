// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatOrganizationRelationships Google Chat organization relationships.
type GoogleChatOrganizationRelationships struct {
	// The delegated user relationship.
	DelegatedUser *GoogleChatOrganizationRelationshipsDelegatedUser `json:"delegated_user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleChatOrganizationRelationships instantiates a new GoogleChatOrganizationRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleChatOrganizationRelationships() *GoogleChatOrganizationRelationships {
	this := GoogleChatOrganizationRelationships{}
	return &this
}

// NewGoogleChatOrganizationRelationshipsWithDefaults instantiates a new GoogleChatOrganizationRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleChatOrganizationRelationshipsWithDefaults() *GoogleChatOrganizationRelationships {
	this := GoogleChatOrganizationRelationships{}
	return &this
}

// GetDelegatedUser returns the DelegatedUser field value if set, zero value otherwise.
func (o *GoogleChatOrganizationRelationships) GetDelegatedUser() GoogleChatOrganizationRelationshipsDelegatedUser {
	if o == nil || o.DelegatedUser == nil {
		var ret GoogleChatOrganizationRelationshipsDelegatedUser
		return ret
	}
	return *o.DelegatedUser
}

// GetDelegatedUserOk returns a tuple with the DelegatedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatOrganizationRelationships) GetDelegatedUserOk() (*GoogleChatOrganizationRelationshipsDelegatedUser, bool) {
	if o == nil || o.DelegatedUser == nil {
		return nil, false
	}
	return o.DelegatedUser, true
}

// HasDelegatedUser returns a boolean if a field has been set.
func (o *GoogleChatOrganizationRelationships) HasDelegatedUser() bool {
	return o != nil && o.DelegatedUser != nil
}

// SetDelegatedUser gets a reference to the given GoogleChatOrganizationRelationshipsDelegatedUser and assigns it to the DelegatedUser field.
func (o *GoogleChatOrganizationRelationships) SetDelegatedUser(v GoogleChatOrganizationRelationshipsDelegatedUser) {
	o.DelegatedUser = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleChatOrganizationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DelegatedUser != nil {
		toSerialize["delegated_user"] = o.DelegatedUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleChatOrganizationRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DelegatedUser *GoogleChatOrganizationRelationshipsDelegatedUser `json:"delegated_user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"delegated_user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DelegatedUser != nil && all.DelegatedUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DelegatedUser = all.DelegatedUser

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
