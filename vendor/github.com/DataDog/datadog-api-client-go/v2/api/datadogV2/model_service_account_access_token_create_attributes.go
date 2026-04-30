// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceAccountAccessTokenCreateAttributes Attributes used to create a service account access token.
type ServiceAccountAccessTokenCreateAttributes struct {
	// Expiration date of the access token. Optional for service account tokens.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// Name of the access token.
	Name string `json:"name"`
	// Array of scopes to grant the access token.
	Scopes []string `json:"scopes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceAccountAccessTokenCreateAttributes instantiates a new ServiceAccountAccessTokenCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceAccountAccessTokenCreateAttributes(name string, scopes []string) *ServiceAccountAccessTokenCreateAttributes {
	this := ServiceAccountAccessTokenCreateAttributes{}
	this.Name = name
	this.Scopes = scopes
	return &this
}

// NewServiceAccountAccessTokenCreateAttributesWithDefaults instantiates a new ServiceAccountAccessTokenCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceAccountAccessTokenCreateAttributesWithDefaults() *ServiceAccountAccessTokenCreateAttributes {
	this := ServiceAccountAccessTokenCreateAttributes{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ServiceAccountAccessTokenCreateAttributes) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessTokenCreateAttributes) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ServiceAccountAccessTokenCreateAttributes) HasExpiresAt() bool {
	return o != nil && o.ExpiresAt != nil
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *ServiceAccountAccessTokenCreateAttributes) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetName returns the Name field value.
func (o *ServiceAccountAccessTokenCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessTokenCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ServiceAccountAccessTokenCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetScopes returns the Scopes field value.
func (o *ServiceAccountAccessTokenCreateAttributes) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessTokenCreateAttributes) GetScopesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value.
func (o *ServiceAccountAccessTokenCreateAttributes) SetScopes(v []string) {
	o.Scopes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceAccountAccessTokenCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExpiresAt != nil {
		if o.ExpiresAt.Nanosecond() == 0 {
			toSerialize["expires_at"] = o.ExpiresAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["expires_at"] = o.ExpiresAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["name"] = o.Name
	toSerialize["scopes"] = o.Scopes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceAccountAccessTokenCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExpiresAt *time.Time `json:"expires_at,omitempty"`
		Name      *string    `json:"name"`
		Scopes    *[]string  `json:"scopes"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.ExpiresAt = all.ExpiresAt
	o.Name = *all.Name
	o.Scopes = *all.Scopes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
