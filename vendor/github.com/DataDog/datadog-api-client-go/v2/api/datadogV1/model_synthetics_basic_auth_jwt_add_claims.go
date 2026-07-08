// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBasicAuthJWTAddClaims Standard JWT claims to automatically inject.
type SyntheticsBasicAuthJWTAddClaims struct {
	// Whether to inject the `exp` (expiration) claim.
	Exp *bool `json:"exp,omitempty"`
	// Whether to inject the `iat` (issued at) claim.
	Iat *bool `json:"iat,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsBasicAuthJWTAddClaims instantiates a new SyntheticsBasicAuthJWTAddClaims object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsBasicAuthJWTAddClaims() *SyntheticsBasicAuthJWTAddClaims {
	this := SyntheticsBasicAuthJWTAddClaims{}
	return &this
}

// NewSyntheticsBasicAuthJWTAddClaimsWithDefaults instantiates a new SyntheticsBasicAuthJWTAddClaims object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsBasicAuthJWTAddClaimsWithDefaults() *SyntheticsBasicAuthJWTAddClaims {
	this := SyntheticsBasicAuthJWTAddClaims{}
	return &this
}

// GetExp returns the Exp field value if set, zero value otherwise.
func (o *SyntheticsBasicAuthJWTAddClaims) GetExp() bool {
	if o == nil || o.Exp == nil {
		var ret bool
		return ret
	}
	return *o.Exp
}

// GetExpOk returns a tuple with the Exp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthJWTAddClaims) GetExpOk() (*bool, bool) {
	if o == nil || o.Exp == nil {
		return nil, false
	}
	return o.Exp, true
}

// HasExp returns a boolean if a field has been set.
func (o *SyntheticsBasicAuthJWTAddClaims) HasExp() bool {
	return o != nil && o.Exp != nil
}

// SetExp gets a reference to the given bool and assigns it to the Exp field.
func (o *SyntheticsBasicAuthJWTAddClaims) SetExp(v bool) {
	o.Exp = &v
}

// GetIat returns the Iat field value if set, zero value otherwise.
func (o *SyntheticsBasicAuthJWTAddClaims) GetIat() bool {
	if o == nil || o.Iat == nil {
		var ret bool
		return ret
	}
	return *o.Iat
}

// GetIatOk returns a tuple with the Iat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthJWTAddClaims) GetIatOk() (*bool, bool) {
	if o == nil || o.Iat == nil {
		return nil, false
	}
	return o.Iat, true
}

// HasIat returns a boolean if a field has been set.
func (o *SyntheticsBasicAuthJWTAddClaims) HasIat() bool {
	return o != nil && o.Iat != nil
}

// SetIat gets a reference to the given bool and assigns it to the Iat field.
func (o *SyntheticsBasicAuthJWTAddClaims) SetIat(v bool) {
	o.Iat = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsBasicAuthJWTAddClaims) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Exp != nil {
		toSerialize["exp"] = o.Exp
	}
	if o.Iat != nil {
		toSerialize["iat"] = o.Iat
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsBasicAuthJWTAddClaims) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Exp *bool `json:"exp,omitempty"`
		Iat *bool `json:"iat,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exp", "iat"})
	} else {
		return err
	}
	o.Exp = all.Exp
	o.Iat = all.Iat

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
