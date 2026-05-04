// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnonymizeUserError Error encountered when anonymizing a specific user.
type AnonymizeUserError struct {
	// Error message describing why anonymization failed.
	Error string `json:"error"`
	// UUID of the user that failed to be anonymized.
	UserId string `json:"user_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnonymizeUserError instantiates a new AnonymizeUserError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnonymizeUserError(error string, userId string) *AnonymizeUserError {
	this := AnonymizeUserError{}
	this.Error = error
	this.UserId = userId
	return &this
}

// NewAnonymizeUserErrorWithDefaults instantiates a new AnonymizeUserError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnonymizeUserErrorWithDefaults() *AnonymizeUserError {
	this := AnonymizeUserError{}
	return &this
}

// GetError returns the Error field value.
func (o *AnonymizeUserError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AnonymizeUserError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value.
func (o *AnonymizeUserError) SetError(v string) {
	o.Error = v
}

// GetUserId returns the UserId field value.
func (o *AnonymizeUserError) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AnonymizeUserError) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *AnonymizeUserError) SetUserId(v string) {
	o.UserId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnonymizeUserError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["error"] = o.Error
	toSerialize["user_id"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnonymizeUserError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error  *string `json:"error"`
		UserId *string `json:"user_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Error == nil {
		return fmt.Errorf("required field error missing")
	}
	if all.UserId == nil {
		return fmt.Errorf("required field user_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error", "user_id"})
	} else {
		return err
	}
	o.Error = *all.Error
	o.UserId = *all.UserId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
