// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTenancyConfigDataAttributesAuthCredentials OCI API signing key credentials used to authenticate the Datadog integration with the OCI tenancy.
type CreateTenancyConfigDataAttributesAuthCredentials struct {
	// The fingerprint of the OCI API signing key used for authentication.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// The PEM-encoded private key corresponding to the OCI API signing key fingerprint.
	PrivateKey string `json:"private_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateTenancyConfigDataAttributesAuthCredentials instantiates a new CreateTenancyConfigDataAttributesAuthCredentials object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTenancyConfigDataAttributesAuthCredentials(privateKey string) *CreateTenancyConfigDataAttributesAuthCredentials {
	this := CreateTenancyConfigDataAttributesAuthCredentials{}
	this.PrivateKey = privateKey
	return &this
}

// NewCreateTenancyConfigDataAttributesAuthCredentialsWithDefaults instantiates a new CreateTenancyConfigDataAttributesAuthCredentials object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTenancyConfigDataAttributesAuthCredentialsWithDefaults() *CreateTenancyConfigDataAttributesAuthCredentials {
	this := CreateTenancyConfigDataAttributesAuthCredentials{}
	return &this
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *CreateTenancyConfigDataAttributesAuthCredentials) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributesAuthCredentials) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CreateTenancyConfigDataAttributesAuthCredentials) HasFingerprint() bool {
	return o != nil && o.Fingerprint != nil
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *CreateTenancyConfigDataAttributesAuthCredentials) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetPrivateKey returns the PrivateKey field value.
func (o *CreateTenancyConfigDataAttributesAuthCredentials) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributesAuthCredentials) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value.
func (o *CreateTenancyConfigDataAttributesAuthCredentials) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTenancyConfigDataAttributesAuthCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	toSerialize["private_key"] = o.PrivateKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateTenancyConfigDataAttributesAuthCredentials) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fingerprint *string `json:"fingerprint,omitempty"`
		PrivateKey  *string `json:"private_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PrivateKey == nil {
		return fmt.Errorf("required field private_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fingerprint", "private_key"})
	} else {
		return err
	}
	o.Fingerprint = all.Fingerprint
	o.PrivateKey = *all.PrivateKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
