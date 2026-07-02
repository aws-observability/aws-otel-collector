// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatDelegatedUserResponse Response containing a Google Chat delegated user.
type GoogleChatDelegatedUserResponse struct {
	// Google Chat delegated user data from a response.
	Data GoogleChatDelegatedUserData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleChatDelegatedUserResponse instantiates a new GoogleChatDelegatedUserResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleChatDelegatedUserResponse(data GoogleChatDelegatedUserData) *GoogleChatDelegatedUserResponse {
	this := GoogleChatDelegatedUserResponse{}
	this.Data = data
	return &this
}

// NewGoogleChatDelegatedUserResponseWithDefaults instantiates a new GoogleChatDelegatedUserResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleChatDelegatedUserResponseWithDefaults() *GoogleChatDelegatedUserResponse {
	this := GoogleChatDelegatedUserResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *GoogleChatDelegatedUserResponse) GetData() GoogleChatDelegatedUserData {
	if o == nil {
		var ret GoogleChatDelegatedUserData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GoogleChatDelegatedUserResponse) GetDataOk() (*GoogleChatDelegatedUserData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *GoogleChatDelegatedUserResponse) SetData(v GoogleChatDelegatedUserData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleChatDelegatedUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleChatDelegatedUserResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *GoogleChatDelegatedUserData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
