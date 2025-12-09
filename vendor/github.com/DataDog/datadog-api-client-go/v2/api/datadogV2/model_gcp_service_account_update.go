// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPServiceAccountUpdate The definition of the `GCPServiceAccount` object.
type GCPServiceAccountUpdate struct {
	// The `GCPServiceAccountUpdate` `private_key`.
	PrivateKey *string `json:"private_key,omitempty"`
	// The `GCPServiceAccountUpdate` `service_account_email`.
	ServiceAccountEmail *string `json:"service_account_email,omitempty"`
	// The definition of the `GCPServiceAccount` object.
	Type GCPServiceAccountCredentialType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPServiceAccountUpdate instantiates a new GCPServiceAccountUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPServiceAccountUpdate(typeVar GCPServiceAccountCredentialType) *GCPServiceAccountUpdate {
	this := GCPServiceAccountUpdate{}
	this.Type = typeVar
	return &this
}

// NewGCPServiceAccountUpdateWithDefaults instantiates a new GCPServiceAccountUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPServiceAccountUpdateWithDefaults() *GCPServiceAccountUpdate {
	this := GCPServiceAccountUpdate{}
	return &this
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *GCPServiceAccountUpdate) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPServiceAccountUpdate) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *GCPServiceAccountUpdate) HasPrivateKey() bool {
	return o != nil && o.PrivateKey != nil
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *GCPServiceAccountUpdate) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetServiceAccountEmail returns the ServiceAccountEmail field value if set, zero value otherwise.
func (o *GCPServiceAccountUpdate) GetServiceAccountEmail() string {
	if o == nil || o.ServiceAccountEmail == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccountEmail
}

// GetServiceAccountEmailOk returns a tuple with the ServiceAccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPServiceAccountUpdate) GetServiceAccountEmailOk() (*string, bool) {
	if o == nil || o.ServiceAccountEmail == nil {
		return nil, false
	}
	return o.ServiceAccountEmail, true
}

// HasServiceAccountEmail returns a boolean if a field has been set.
func (o *GCPServiceAccountUpdate) HasServiceAccountEmail() bool {
	return o != nil && o.ServiceAccountEmail != nil
}

// SetServiceAccountEmail gets a reference to the given string and assigns it to the ServiceAccountEmail field.
func (o *GCPServiceAccountUpdate) SetServiceAccountEmail(v string) {
	o.ServiceAccountEmail = &v
}

// GetType returns the Type field value.
func (o *GCPServiceAccountUpdate) GetType() GCPServiceAccountCredentialType {
	if o == nil {
		var ret GCPServiceAccountCredentialType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GCPServiceAccountUpdate) GetTypeOk() (*GCPServiceAccountCredentialType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GCPServiceAccountUpdate) SetType(v GCPServiceAccountCredentialType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPServiceAccountUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	if o.ServiceAccountEmail != nil {
		toSerialize["service_account_email"] = o.ServiceAccountEmail
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPServiceAccountUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PrivateKey          *string                          `json:"private_key,omitempty"`
		ServiceAccountEmail *string                          `json:"service_account_email,omitempty"`
		Type                *GCPServiceAccountCredentialType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.PrivateKey = all.PrivateKey
	o.ServiceAccountEmail = all.ServiceAccountEmail
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
