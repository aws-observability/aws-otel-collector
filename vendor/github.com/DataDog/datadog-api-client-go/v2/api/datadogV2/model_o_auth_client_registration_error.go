// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OAuthClientRegistrationError Error payload returned by OAuth2 dynamic client registration as defined by RFC 7591.
type OAuthClientRegistrationError struct {
	// Single ASCII error code per RFC 7591, such as `invalid_request` or `invalid_client_metadata`.
	Error string `json:"error"`
	// Human-readable description of the error.
	ErrorDescription string `json:"error_description"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOAuthClientRegistrationError instantiates a new OAuthClientRegistrationError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOAuthClientRegistrationError(error string, errorDescription string) *OAuthClientRegistrationError {
	this := OAuthClientRegistrationError{}
	this.Error = error
	this.ErrorDescription = errorDescription
	return &this
}

// NewOAuthClientRegistrationErrorWithDefaults instantiates a new OAuthClientRegistrationError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOAuthClientRegistrationErrorWithDefaults() *OAuthClientRegistrationError {
	this := OAuthClientRegistrationError{}
	return &this
}

// GetError returns the Error field value.
func (o *OAuthClientRegistrationError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value.
func (o *OAuthClientRegistrationError) SetError(v string) {
	o.Error = v
}

// GetErrorDescription returns the ErrorDescription field value.
func (o *OAuthClientRegistrationError) GetErrorDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationError) GetErrorDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorDescription, true
}

// SetErrorDescription sets field value.
func (o *OAuthClientRegistrationError) SetErrorDescription(v string) {
	o.ErrorDescription = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OAuthClientRegistrationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["error"] = o.Error
	toSerialize["error_description"] = o.ErrorDescription

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OAuthClientRegistrationError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error            *string `json:"error"`
		ErrorDescription *string `json:"error_description"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Error == nil {
		return fmt.Errorf("required field error missing")
	}
	if all.ErrorDescription == nil {
		return fmt.Errorf("required field error_description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error", "error_description"})
	} else {
		return err
	}
	o.Error = *all.Error
	o.ErrorDescription = *all.ErrorDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
