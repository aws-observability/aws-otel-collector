// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAuthConfigKeys AWS Authentication config to integrate your account using an access key pair.
type AWSAuthConfigKeys struct {
	// AWS Access Key ID.
	AccessKeyId string `json:"access_key_id"`
	// AWS Secret Access Key.
	SecretAccessKey *string `json:"secret_access_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSAuthConfigKeys instantiates a new AWSAuthConfigKeys object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAuthConfigKeys(accessKeyId string) *AWSAuthConfigKeys {
	this := AWSAuthConfigKeys{}
	this.AccessKeyId = accessKeyId
	return &this
}

// NewAWSAuthConfigKeysWithDefaults instantiates a new AWSAuthConfigKeys object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAuthConfigKeysWithDefaults() *AWSAuthConfigKeys {
	this := AWSAuthConfigKeys{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value.
func (o *AWSAuthConfigKeys) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *AWSAuthConfigKeys) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value.
func (o *AWSAuthConfigKeys) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *AWSAuthConfigKeys) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAuthConfigKeys) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKey == nil {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *AWSAuthConfigKeys) HasSecretAccessKey() bool {
	return o != nil && o.SecretAccessKey != nil
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *AWSAuthConfigKeys) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAuthConfigKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["access_key_id"] = o.AccessKeyId
	if o.SecretAccessKey != nil {
		toSerialize["secret_access_key"] = o.SecretAccessKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSAuthConfigKeys) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessKeyId     *string `json:"access_key_id"`
		SecretAccessKey *string `json:"secret_access_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccessKeyId == nil {
		return fmt.Errorf("required field access_key_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_key_id", "secret_access_key"})
	} else {
		return err
	}
	o.AccessKeyId = *all.AccessKeyId
	o.SecretAccessKey = all.SecretAccessKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
