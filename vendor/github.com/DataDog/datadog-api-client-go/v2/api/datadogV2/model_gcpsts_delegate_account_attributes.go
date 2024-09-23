// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPSTSDelegateAccountAttributes Your delegate account attributes.
type GCPSTSDelegateAccountAttributes struct {
	// Your organization's Datadog principal email address.
	DelegateAccountEmail *string `json:"delegate_account_email,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPSTSDelegateAccountAttributes instantiates a new GCPSTSDelegateAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPSTSDelegateAccountAttributes() *GCPSTSDelegateAccountAttributes {
	this := GCPSTSDelegateAccountAttributes{}
	return &this
}

// NewGCPSTSDelegateAccountAttributesWithDefaults instantiates a new GCPSTSDelegateAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPSTSDelegateAccountAttributesWithDefaults() *GCPSTSDelegateAccountAttributes {
	this := GCPSTSDelegateAccountAttributes{}
	return &this
}

// GetDelegateAccountEmail returns the DelegateAccountEmail field value if set, zero value otherwise.
func (o *GCPSTSDelegateAccountAttributes) GetDelegateAccountEmail() string {
	if o == nil || o.DelegateAccountEmail == nil {
		var ret string
		return ret
	}
	return *o.DelegateAccountEmail
}

// GetDelegateAccountEmailOk returns a tuple with the DelegateAccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSDelegateAccountAttributes) GetDelegateAccountEmailOk() (*string, bool) {
	if o == nil || o.DelegateAccountEmail == nil {
		return nil, false
	}
	return o.DelegateAccountEmail, true
}

// HasDelegateAccountEmail returns a boolean if a field has been set.
func (o *GCPSTSDelegateAccountAttributes) HasDelegateAccountEmail() bool {
	return o != nil && o.DelegateAccountEmail != nil
}

// SetDelegateAccountEmail gets a reference to the given string and assigns it to the DelegateAccountEmail field.
func (o *GCPSTSDelegateAccountAttributes) SetDelegateAccountEmail(v string) {
	o.DelegateAccountEmail = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPSTSDelegateAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DelegateAccountEmail != nil {
		toSerialize["delegate_account_email"] = o.DelegateAccountEmail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPSTSDelegateAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DelegateAccountEmail *string `json:"delegate_account_email,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"delegate_account_email"})
	} else {
		return err
	}
	o.DelegateAccountEmail = all.DelegateAccountEmail

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
