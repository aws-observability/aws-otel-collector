// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultFailure Details about the failure of a Synthetic test.
type SyntheticsTestResultFailure struct {
	// Error code for the failure.
	Code *string `json:"code,omitempty"`
	// Internal error code used for debugging.
	InternalCode *string `json:"internal_code,omitempty"`
	// Internal error message used for debugging.
	InternalMessage *string `json:"internal_message,omitempty"`
	// Error message for the failure.
	Message *string `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultFailure instantiates a new SyntheticsTestResultFailure object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultFailure() *SyntheticsTestResultFailure {
	this := SyntheticsTestResultFailure{}
	return &this
}

// NewSyntheticsTestResultFailureWithDefaults instantiates a new SyntheticsTestResultFailure object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultFailureWithDefaults() *SyntheticsTestResultFailure {
	this := SyntheticsTestResultFailure{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SyntheticsTestResultFailure) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultFailure) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SyntheticsTestResultFailure) HasCode() bool {
	return o != nil && o.Code != nil
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SyntheticsTestResultFailure) SetCode(v string) {
	o.Code = &v
}

// GetInternalCode returns the InternalCode field value if set, zero value otherwise.
func (o *SyntheticsTestResultFailure) GetInternalCode() string {
	if o == nil || o.InternalCode == nil {
		var ret string
		return ret
	}
	return *o.InternalCode
}

// GetInternalCodeOk returns a tuple with the InternalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultFailure) GetInternalCodeOk() (*string, bool) {
	if o == nil || o.InternalCode == nil {
		return nil, false
	}
	return o.InternalCode, true
}

// HasInternalCode returns a boolean if a field has been set.
func (o *SyntheticsTestResultFailure) HasInternalCode() bool {
	return o != nil && o.InternalCode != nil
}

// SetInternalCode gets a reference to the given string and assigns it to the InternalCode field.
func (o *SyntheticsTestResultFailure) SetInternalCode(v string) {
	o.InternalCode = &v
}

// GetInternalMessage returns the InternalMessage field value if set, zero value otherwise.
func (o *SyntheticsTestResultFailure) GetInternalMessage() string {
	if o == nil || o.InternalMessage == nil {
		var ret string
		return ret
	}
	return *o.InternalMessage
}

// GetInternalMessageOk returns a tuple with the InternalMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultFailure) GetInternalMessageOk() (*string, bool) {
	if o == nil || o.InternalMessage == nil {
		return nil, false
	}
	return o.InternalMessage, true
}

// HasInternalMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultFailure) HasInternalMessage() bool {
	return o != nil && o.InternalMessage != nil
}

// SetInternalMessage gets a reference to the given string and assigns it to the InternalMessage field.
func (o *SyntheticsTestResultFailure) SetInternalMessage(v string) {
	o.InternalMessage = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsTestResultFailure) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultFailure) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultFailure) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SyntheticsTestResultFailure) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.InternalCode != nil {
		toSerialize["internal_code"] = o.InternalCode
	}
	if o.InternalMessage != nil {
		toSerialize["internal_message"] = o.InternalMessage
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultFailure) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code            *string `json:"code,omitempty"`
		InternalCode    *string `json:"internal_code,omitempty"`
		InternalMessage *string `json:"internal_message,omitempty"`
		Message         *string `json:"message,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code", "internal_code", "internal_message", "message"})
	} else {
		return err
	}
	o.Code = all.Code
	o.InternalCode = all.InternalCode
	o.InternalMessage = all.InternalMessage
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
