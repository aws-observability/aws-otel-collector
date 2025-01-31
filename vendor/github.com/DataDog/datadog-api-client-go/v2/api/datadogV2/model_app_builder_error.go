// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppBuilderError The definition of `AppBuilderError` object.
type AppBuilderError struct {
	// The `AppBuilderError` `errors`.
	Errors []AppBuilderErrorErrorsItems `json:"errors,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAppBuilderError instantiates a new AppBuilderError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAppBuilderError() *AppBuilderError {
	this := AppBuilderError{}
	return &this
}

// NewAppBuilderErrorWithDefaults instantiates a new AppBuilderError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAppBuilderErrorWithDefaults() *AppBuilderError {
	this := AppBuilderError{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AppBuilderError) GetErrors() []AppBuilderErrorErrorsItems {
	if o == nil || o.Errors == nil {
		var ret []AppBuilderErrorErrorsItems
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppBuilderError) GetErrorsOk() (*[]AppBuilderErrorErrorsItems, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AppBuilderError) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given []AppBuilderErrorErrorsItems and assigns it to the Errors field.
func (o *AppBuilderError) SetErrors(v []AppBuilderErrorErrorsItems) {
	o.Errors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AppBuilderError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AppBuilderError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Errors []AppBuilderErrorErrorsItems `json:"errors,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"errors"})
	} else {
		return err
	}
	o.Errors = all.Errors

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
