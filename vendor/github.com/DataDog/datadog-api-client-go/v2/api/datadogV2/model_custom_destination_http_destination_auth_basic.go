// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationHttpDestinationAuthBasic Basic access authentication.
type CustomDestinationHttpDestinationAuthBasic struct {
	// The password of the authentication. This field is not returned by the API.
	Password string `json:"password"`
	// Type of the basic access authentication.
	Type CustomDestinationHttpDestinationAuthBasicType `json:"type"`
	// The username of the authentication. This field is not returned by the API.
	Username string `json:"username"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationHttpDestinationAuthBasic instantiates a new CustomDestinationHttpDestinationAuthBasic object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationHttpDestinationAuthBasic(password string, typeVar CustomDestinationHttpDestinationAuthBasicType, username string) *CustomDestinationHttpDestinationAuthBasic {
	this := CustomDestinationHttpDestinationAuthBasic{}
	this.Password = password
	this.Type = typeVar
	this.Username = username
	return &this
}

// NewCustomDestinationHttpDestinationAuthBasicWithDefaults instantiates a new CustomDestinationHttpDestinationAuthBasic object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationHttpDestinationAuthBasicWithDefaults() *CustomDestinationHttpDestinationAuthBasic {
	this := CustomDestinationHttpDestinationAuthBasic{}
	var typeVar CustomDestinationHttpDestinationAuthBasicType = CUSTOMDESTINATIONHTTPDESTINATIONAUTHBASICTYPE_BASIC
	this.Type = typeVar
	return &this
}

// GetPassword returns the Password field value.
func (o *CustomDestinationHttpDestinationAuthBasic) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationHttpDestinationAuthBasic) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *CustomDestinationHttpDestinationAuthBasic) SetPassword(v string) {
	o.Password = v
}

// GetType returns the Type field value.
func (o *CustomDestinationHttpDestinationAuthBasic) GetType() CustomDestinationHttpDestinationAuthBasicType {
	if o == nil {
		var ret CustomDestinationHttpDestinationAuthBasicType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationHttpDestinationAuthBasic) GetTypeOk() (*CustomDestinationHttpDestinationAuthBasicType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationHttpDestinationAuthBasic) SetType(v CustomDestinationHttpDestinationAuthBasicType) {
	o.Type = v
}

// GetUsername returns the Username field value.
func (o *CustomDestinationHttpDestinationAuthBasic) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationHttpDestinationAuthBasic) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *CustomDestinationHttpDestinationAuthBasic) SetUsername(v string) {
	o.Username = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationHttpDestinationAuthBasic) MarshalJSON() ([]byte, error) {
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
func (o *CustomDestinationHttpDestinationAuthBasic) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Password *string                                        `json:"password"`
		Type     *CustomDestinationHttpDestinationAuthBasicType `json:"type"`
		Username *string                                        `json:"username"`
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
