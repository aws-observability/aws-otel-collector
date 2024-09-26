// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBasicAuthDigest Object to handle digest authentication when performing the test.
type SyntheticsBasicAuthDigest struct {
	// Password to use for the digest authentication.
	Password string `json:"password"`
	// The type of basic authentication to use when performing the test.
	Type SyntheticsBasicAuthDigestType `json:"type"`
	// Username to use for the digest authentication.
	Username string `json:"username"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsBasicAuthDigest instantiates a new SyntheticsBasicAuthDigest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsBasicAuthDigest(password string, typeVar SyntheticsBasicAuthDigestType, username string) *SyntheticsBasicAuthDigest {
	this := SyntheticsBasicAuthDigest{}
	this.Password = password
	this.Type = typeVar
	this.Username = username
	return &this
}

// NewSyntheticsBasicAuthDigestWithDefaults instantiates a new SyntheticsBasicAuthDigest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsBasicAuthDigestWithDefaults() *SyntheticsBasicAuthDigest {
	this := SyntheticsBasicAuthDigest{}
	var typeVar SyntheticsBasicAuthDigestType = SYNTHETICSBASICAUTHDIGESTTYPE_DIGEST
	this.Type = typeVar
	return &this
}

// GetPassword returns the Password field value.
func (o *SyntheticsBasicAuthDigest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthDigest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *SyntheticsBasicAuthDigest) SetPassword(v string) {
	o.Password = v
}

// GetType returns the Type field value.
func (o *SyntheticsBasicAuthDigest) GetType() SyntheticsBasicAuthDigestType {
	if o == nil {
		var ret SyntheticsBasicAuthDigestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthDigest) GetTypeOk() (*SyntheticsBasicAuthDigestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsBasicAuthDigest) SetType(v SyntheticsBasicAuthDigestType) {
	o.Type = v
}

// GetUsername returns the Username field value.
func (o *SyntheticsBasicAuthDigest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthDigest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *SyntheticsBasicAuthDigest) SetUsername(v string) {
	o.Username = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsBasicAuthDigest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["password"] = o.Password
	toSerialize["type"] = o.Type
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsBasicAuthDigest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Password *string                        `json:"password"`
		Type     *SyntheticsBasicAuthDigestType `json:"type"`
		Username *string                        `json:"username"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Username == nil {
		return fmt.Errorf("required field username missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"password", "type", "username"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Password = *all.Password
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Username = *all.Username

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
