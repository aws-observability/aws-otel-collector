// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPServiceAccount The definition of the `GCPServiceAccount` object.
type GCPServiceAccount struct {
	// The `GCPServiceAccount` `private_key`.
	PrivateKey string `json:"private_key"`
	// The `GCPServiceAccount` `service_account_email`.
	ServiceAccountEmail string `json:"service_account_email"`
	// The definition of the `GCPServiceAccount` object.
	Type GCPServiceAccountCredentialType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPServiceAccount instantiates a new GCPServiceAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPServiceAccount(privateKey string, serviceAccountEmail string, typeVar GCPServiceAccountCredentialType) *GCPServiceAccount {
	this := GCPServiceAccount{}
	this.PrivateKey = privateKey
	this.ServiceAccountEmail = serviceAccountEmail
	this.Type = typeVar
	return &this
}

// NewGCPServiceAccountWithDefaults instantiates a new GCPServiceAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPServiceAccountWithDefaults() *GCPServiceAccount {
	this := GCPServiceAccount{}
	return &this
}

// GetPrivateKey returns the PrivateKey field value.
func (o *GCPServiceAccount) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *GCPServiceAccount) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value.
func (o *GCPServiceAccount) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetServiceAccountEmail returns the ServiceAccountEmail field value.
func (o *GCPServiceAccount) GetServiceAccountEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceAccountEmail
}

// GetServiceAccountEmailOk returns a tuple with the ServiceAccountEmail field value
// and a boolean to check if the value has been set.
func (o *GCPServiceAccount) GetServiceAccountEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccountEmail, true
}

// SetServiceAccountEmail sets field value.
func (o *GCPServiceAccount) SetServiceAccountEmail(v string) {
	o.ServiceAccountEmail = v
}

// GetType returns the Type field value.
func (o *GCPServiceAccount) GetType() GCPServiceAccountCredentialType {
	if o == nil {
		var ret GCPServiceAccountCredentialType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GCPServiceAccount) GetTypeOk() (*GCPServiceAccountCredentialType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GCPServiceAccount) SetType(v GCPServiceAccountCredentialType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["service_account_email"] = o.ServiceAccountEmail
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPServiceAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PrivateKey          *string                          `json:"private_key"`
		ServiceAccountEmail *string                          `json:"service_account_email"`
		Type                *GCPServiceAccountCredentialType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PrivateKey == nil {
		return fmt.Errorf("required field private_key missing")
	}
	if all.ServiceAccountEmail == nil {
		return fmt.Errorf("required field service_account_email missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"private_key", "service_account_email", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.PrivateKey = *all.PrivateKey
	o.ServiceAccountEmail = *all.ServiceAccountEmail
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
