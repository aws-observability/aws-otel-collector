// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PersonalAccessTokenCreateAttributes Attributes used to create a personal access token.
type PersonalAccessTokenCreateAttributes struct {
	// Expiration date of the personal access token. Must be at least 24 hours in the future.
	ExpiresAt time.Time `json:"expires_at"`
	// Name of the personal access token.
	Name string `json:"name"`
	// Array of scopes to grant the personal access token.
	Scopes []string `json:"scopes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPersonalAccessTokenCreateAttributes instantiates a new PersonalAccessTokenCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPersonalAccessTokenCreateAttributes(expiresAt time.Time, name string, scopes []string) *PersonalAccessTokenCreateAttributes {
	this := PersonalAccessTokenCreateAttributes{}
	this.ExpiresAt = expiresAt
	this.Name = name
	this.Scopes = scopes
	return &this
}

// NewPersonalAccessTokenCreateAttributesWithDefaults instantiates a new PersonalAccessTokenCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPersonalAccessTokenCreateAttributesWithDefaults() *PersonalAccessTokenCreateAttributes {
	this := PersonalAccessTokenCreateAttributes{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value.
func (o *PersonalAccessTokenCreateAttributes) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenCreateAttributes) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value.
func (o *PersonalAccessTokenCreateAttributes) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetName returns the Name field value.
func (o *PersonalAccessTokenCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PersonalAccessTokenCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetScopes returns the Scopes field value.
func (o *PersonalAccessTokenCreateAttributes) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenCreateAttributes) GetScopesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value.
func (o *PersonalAccessTokenCreateAttributes) SetScopes(v []string) {
	o.Scopes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PersonalAccessTokenCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExpiresAt.Nanosecond() == 0 {
		toSerialize["expires_at"] = o.ExpiresAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["expires_at"] = o.ExpiresAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["scopes"] = o.Scopes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PersonalAccessTokenCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExpiresAt *time.Time `json:"expires_at"`
		Name      *string    `json:"name"`
		Scopes    *[]string  `json:"scopes"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExpiresAt == nil {
		return fmt.Errorf("required field expires_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Scopes == nil {
		return fmt.Errorf("required field scopes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expires_at", "name", "scopes"})
	} else {
		return err
	}
	o.ExpiresAt = *all.ExpiresAt
	o.Name = *all.Name
	o.Scopes = *all.Scopes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
