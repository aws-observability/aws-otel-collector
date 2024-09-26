// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppCIError Contains information of the CI error.
type CIAppCIError struct {
	// Error category used to differentiate between issues related to the developer or provider environments.
	Domain *CIAppCIErrorDomain `json:"domain,omitempty"`
	// Error message.
	Message datadog.NullableString `json:"message,omitempty"`
	// The stack trace of the reported errors.
	Stack datadog.NullableString `json:"stack,omitempty"`
	// Short description of the error type.
	Type datadog.NullableString `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCIAppCIError instantiates a new CIAppCIError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppCIError() *CIAppCIError {
	this := CIAppCIError{}
	return &this
}

// NewCIAppCIErrorWithDefaults instantiates a new CIAppCIError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppCIErrorWithDefaults() *CIAppCIError {
	this := CIAppCIError{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CIAppCIError) GetDomain() CIAppCIErrorDomain {
	if o == nil || o.Domain == nil {
		var ret CIAppCIErrorDomain
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppCIError) GetDomainOk() (*CIAppCIErrorDomain, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CIAppCIError) HasDomain() bool {
	return o != nil && o.Domain != nil
}

// SetDomain gets a reference to the given CIAppCIErrorDomain and assigns it to the Domain field.
func (o *CIAppCIError) SetDomain(v CIAppCIErrorDomain) {
	o.Domain = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppCIError) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppCIError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *CIAppCIError) HasMessage() bool {
	return o != nil && o.Message.IsSet()
}

// SetMessage gets a reference to the given datadog.NullableString and assigns it to the Message field.
func (o *CIAppCIError) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil.
func (o *CIAppCIError) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil.
func (o *CIAppCIError) UnsetMessage() {
	o.Message.Unset()
}

// GetStack returns the Stack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppCIError) GetStack() string {
	if o == nil || o.Stack.Get() == nil {
		var ret string
		return ret
	}
	return *o.Stack.Get()
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppCIError) GetStackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stack.Get(), o.Stack.IsSet()
}

// HasStack returns a boolean if a field has been set.
func (o *CIAppCIError) HasStack() bool {
	return o != nil && o.Stack.IsSet()
}

// SetStack gets a reference to the given datadog.NullableString and assigns it to the Stack field.
func (o *CIAppCIError) SetStack(v string) {
	o.Stack.Set(&v)
}

// SetStackNil sets the value for Stack to be an explicit nil.
func (o *CIAppCIError) SetStackNil() {
	o.Stack.Set(nil)
}

// UnsetStack ensures that no value is present for Stack, not even an explicit nil.
func (o *CIAppCIError) UnsetStack() {
	o.Stack.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppCIError) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppCIError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CIAppCIError) HasType() bool {
	return o != nil && o.Type.IsSet()
}

// SetType gets a reference to the given datadog.NullableString and assigns it to the Type field.
func (o *CIAppCIError) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil.
func (o *CIAppCIError) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil.
func (o *CIAppCIError) UnsetType() {
	o.Type.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppCIError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Stack.IsSet() {
		toSerialize["stack"] = o.Stack.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppCIError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Domain  *CIAppCIErrorDomain    `json:"domain,omitempty"`
		Message datadog.NullableString `json:"message,omitempty"`
		Stack   datadog.NullableString `json:"stack,omitempty"`
		Type    datadog.NullableString `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domain", "message", "stack", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Domain != nil && !all.Domain.IsValid() {
		hasInvalidField = true
	} else {
		o.Domain = all.Domain
	}
	o.Message = all.Message
	o.Stack = all.Stack
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableCIAppCIError handles when a null is used for CIAppCIError.
type NullableCIAppCIError struct {
	value *CIAppCIError
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppCIError) Get() *CIAppCIError {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppCIError) Set(val *CIAppCIError) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppCIError) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableCIAppCIError) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppCIError initializes the struct as if Set has been called.
func NewNullableCIAppCIError(val *CIAppCIError) *NullableCIAppCIError {
	return &NullableCIAppCIError{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppCIError) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppCIError) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
