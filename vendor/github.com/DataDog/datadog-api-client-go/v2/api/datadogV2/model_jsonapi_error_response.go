// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JSONAPIErrorResponse API error response.
type JSONAPIErrorResponse struct {
	// A list of errors.
	Errors []JSONAPIErrorItem `json:"errors"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJSONAPIErrorResponse instantiates a new JSONAPIErrorResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJSONAPIErrorResponse(errors []JSONAPIErrorItem) *JSONAPIErrorResponse {
	this := JSONAPIErrorResponse{}
	this.Errors = errors
	return &this
}

// NewJSONAPIErrorResponseWithDefaults instantiates a new JSONAPIErrorResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJSONAPIErrorResponseWithDefaults() *JSONAPIErrorResponse {
	this := JSONAPIErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value.
func (o *JSONAPIErrorResponse) GetErrors() []JSONAPIErrorItem {
	if o == nil {
		var ret []JSONAPIErrorItem
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *JSONAPIErrorResponse) GetErrorsOk() (*[]JSONAPIErrorItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value.
func (o *JSONAPIErrorResponse) SetErrors(v []JSONAPIErrorItem) {
	o.Errors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JSONAPIErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["errors"] = o.Errors

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JSONAPIErrorResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Errors *[]JSONAPIErrorItem `json:"errors"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Errors == nil {
		return fmt.Errorf("required field errors missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"errors"})
	} else {
		return err
	}
	o.Errors = *all.Errors

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
