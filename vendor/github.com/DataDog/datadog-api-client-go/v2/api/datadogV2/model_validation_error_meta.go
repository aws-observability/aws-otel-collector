// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ValidationErrorMeta Describes additional metadata for validation errors, including field names and error messages.
type ValidationErrorMeta struct {
	// The field name that caused the error.
	Field *string `json:"field,omitempty"`
	// The ID of the component in which the error occurred.
	Id *string `json:"id,omitempty"`
	// The detailed error message.
	Message string `json:"message"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewValidationErrorMeta instantiates a new ValidationErrorMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewValidationErrorMeta(message string) *ValidationErrorMeta {
	this := ValidationErrorMeta{}
	this.Message = message
	return &this
}

// NewValidationErrorMetaWithDefaults instantiates a new ValidationErrorMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewValidationErrorMetaWithDefaults() *ValidationErrorMeta {
	this := ValidationErrorMeta{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ValidationErrorMeta) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorMeta) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ValidationErrorMeta) HasField() bool {
	return o != nil && o.Field != nil
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ValidationErrorMeta) SetField(v string) {
	o.Field = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ValidationErrorMeta) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorMeta) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ValidationErrorMeta) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ValidationErrorMeta) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value.
func (o *ValidationErrorMeta) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ValidationErrorMeta) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *ValidationErrorMeta) SetMessage(v string) {
	o.Message = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ValidationErrorMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["message"] = o.Message

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ValidationErrorMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field   *string `json:"field,omitempty"`
		Id      *string `json:"id,omitempty"`
		Message *string `json:"message"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field", "id", "message"})
	} else {
		return err
	}
	o.Field = all.Field
	o.Id = all.Id
	o.Message = *all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
