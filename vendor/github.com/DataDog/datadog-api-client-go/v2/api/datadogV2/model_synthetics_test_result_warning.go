// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultWarning A warning captured during a browser test step.
type SyntheticsTestResultWarning struct {
	// Bounds of elements related to the warning.
	ElementBounds []SyntheticsTestResultBounds `json:"element_bounds,omitempty"`
	// Warning message.
	Message *string `json:"message,omitempty"`
	// Type of the warning.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultWarning instantiates a new SyntheticsTestResultWarning object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultWarning() *SyntheticsTestResultWarning {
	this := SyntheticsTestResultWarning{}
	return &this
}

// NewSyntheticsTestResultWarningWithDefaults instantiates a new SyntheticsTestResultWarning object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultWarningWithDefaults() *SyntheticsTestResultWarning {
	this := SyntheticsTestResultWarning{}
	return &this
}

// GetElementBounds returns the ElementBounds field value if set, zero value otherwise.
func (o *SyntheticsTestResultWarning) GetElementBounds() []SyntheticsTestResultBounds {
	if o == nil || o.ElementBounds == nil {
		var ret []SyntheticsTestResultBounds
		return ret
	}
	return o.ElementBounds
}

// GetElementBoundsOk returns a tuple with the ElementBounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultWarning) GetElementBoundsOk() (*[]SyntheticsTestResultBounds, bool) {
	if o == nil || o.ElementBounds == nil {
		return nil, false
	}
	return &o.ElementBounds, true
}

// HasElementBounds returns a boolean if a field has been set.
func (o *SyntheticsTestResultWarning) HasElementBounds() bool {
	return o != nil && o.ElementBounds != nil
}

// SetElementBounds gets a reference to the given []SyntheticsTestResultBounds and assigns it to the ElementBounds field.
func (o *SyntheticsTestResultWarning) SetElementBounds(v []SyntheticsTestResultBounds) {
	o.ElementBounds = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsTestResultWarning) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultWarning) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultWarning) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SyntheticsTestResultWarning) SetMessage(v string) {
	o.Message = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestResultWarning) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultWarning) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestResultWarning) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SyntheticsTestResultWarning) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultWarning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ElementBounds != nil {
		toSerialize["element_bounds"] = o.ElementBounds
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultWarning) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ElementBounds []SyntheticsTestResultBounds `json:"element_bounds,omitempty"`
		Message       *string                      `json:"message,omitempty"`
		Type          *string                      `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"element_bounds", "message", "type"})
	} else {
		return err
	}
	o.ElementBounds = all.ElementBounds
	o.Message = all.Message
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
