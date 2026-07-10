// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatDelegatedUserAttributes Google Chat delegated user attributes.
type GoogleChatDelegatedUserAttributes struct {
	// The delegated user's display name.
	DisplayName *string `json:"display_name,omitempty"`
	// The delegated user's email address.
	Email *string `json:"email,omitempty"`
	// The list of features enabled for the delegated user.
	Features []string `json:"features,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleChatDelegatedUserAttributes instantiates a new GoogleChatDelegatedUserAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleChatDelegatedUserAttributes() *GoogleChatDelegatedUserAttributes {
	this := GoogleChatDelegatedUserAttributes{}
	return &this
}

// NewGoogleChatDelegatedUserAttributesWithDefaults instantiates a new GoogleChatDelegatedUserAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleChatDelegatedUserAttributesWithDefaults() *GoogleChatDelegatedUserAttributes {
	this := GoogleChatDelegatedUserAttributes{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GoogleChatDelegatedUserAttributes) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatDelegatedUserAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GoogleChatDelegatedUserAttributes) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GoogleChatDelegatedUserAttributes) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GoogleChatDelegatedUserAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatDelegatedUserAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GoogleChatDelegatedUserAttributes) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GoogleChatDelegatedUserAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *GoogleChatDelegatedUserAttributes) GetFeatures() []string {
	if o == nil || o.Features == nil {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatDelegatedUserAttributes) GetFeaturesOk() (*[]string, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return &o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *GoogleChatDelegatedUserAttributes) HasFeatures() bool {
	return o != nil && o.Features != nil
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *GoogleChatDelegatedUserAttributes) SetFeatures(v []string) {
	o.Features = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleChatDelegatedUserAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleChatDelegatedUserAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string  `json:"display_name,omitempty"`
		Email       *string  `json:"email,omitempty"`
		Features    []string `json:"features,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "email", "features"})
	} else {
		return err
	}
	o.DisplayName = all.DisplayName
	o.Email = all.Email
	o.Features = all.Features

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
