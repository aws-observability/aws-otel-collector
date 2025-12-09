// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationElasticsearchDestinationAuth Basic access authentication.
type CustomDestinationElasticsearchDestinationAuth struct {
	// The password of the authentication. This field is not returned by the API.
	Password string `json:"password"`
	// The username of the authentication. This field is not returned by the API.
	Username string `json:"username"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationElasticsearchDestinationAuth instantiates a new CustomDestinationElasticsearchDestinationAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationElasticsearchDestinationAuth(password string, username string) *CustomDestinationElasticsearchDestinationAuth {
	this := CustomDestinationElasticsearchDestinationAuth{}
	this.Password = password
	this.Username = username
	return &this
}

// NewCustomDestinationElasticsearchDestinationAuthWithDefaults instantiates a new CustomDestinationElasticsearchDestinationAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationElasticsearchDestinationAuthWithDefaults() *CustomDestinationElasticsearchDestinationAuth {
	this := CustomDestinationElasticsearchDestinationAuth{}
	return &this
}

// GetPassword returns the Password field value.
func (o *CustomDestinationElasticsearchDestinationAuth) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationElasticsearchDestinationAuth) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *CustomDestinationElasticsearchDestinationAuth) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value.
func (o *CustomDestinationElasticsearchDestinationAuth) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationElasticsearchDestinationAuth) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *CustomDestinationElasticsearchDestinationAuth) SetUsername(v string) {
	o.Username = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationElasticsearchDestinationAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationElasticsearchDestinationAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Password *string `json:"password"`
		Username *string `json:"username"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	if all.Username == nil {
		return fmt.Errorf("required field username missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"password", "username"})
	} else {
		return err
	}
	o.Password = *all.Password
	o.Username = *all.Username

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
