// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOHistoryResponseErrorWithType An object describing the error with error type and error message.
type SLOHistoryResponseErrorWithType struct {
	// A message with more details about the error.
	ErrorMessage string `json:"error_message"`
	// Type of the error.
	ErrorType string `json:"error_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSLOHistoryResponseErrorWithType instantiates a new SLOHistoryResponseErrorWithType object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOHistoryResponseErrorWithType(errorMessage string, errorType string) *SLOHistoryResponseErrorWithType {
	this := SLOHistoryResponseErrorWithType{}
	this.ErrorMessage = errorMessage
	this.ErrorType = errorType
	return &this
}

// NewSLOHistoryResponseErrorWithTypeWithDefaults instantiates a new SLOHistoryResponseErrorWithType object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOHistoryResponseErrorWithTypeWithDefaults() *SLOHistoryResponseErrorWithType {
	this := SLOHistoryResponseErrorWithType{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value.
func (o *SLOHistoryResponseErrorWithType) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *SLOHistoryResponseErrorWithType) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value.
func (o *SLOHistoryResponseErrorWithType) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

// GetErrorType returns the ErrorType field value.
func (o *SLOHistoryResponseErrorWithType) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SLOHistoryResponseErrorWithType) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value.
func (o *SLOHistoryResponseErrorWithType) SetErrorType(v string) {
	o.ErrorType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOHistoryResponseErrorWithType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["error_message"] = o.ErrorMessage
	toSerialize["error_type"] = o.ErrorType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOHistoryResponseErrorWithType) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorMessage *string `json:"error_message"`
		ErrorType    *string `json:"error_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ErrorMessage == nil {
		return fmt.Errorf("required field error_message missing")
	}
	if all.ErrorType == nil {
		return fmt.Errorf("required field error_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error_message", "error_type"})
	} else {
		return err
	}
	o.ErrorMessage = *all.ErrorMessage
	o.ErrorType = *all.ErrorType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
