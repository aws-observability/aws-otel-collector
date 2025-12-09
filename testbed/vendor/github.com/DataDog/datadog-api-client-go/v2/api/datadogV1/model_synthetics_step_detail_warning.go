// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsStepDetailWarning Object collecting warnings for a given step.
type SyntheticsStepDetailWarning struct {
	// Message for the warning.
	Message string `json:"message"`
	// User locator used.
	Type SyntheticsWarningType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsStepDetailWarning instantiates a new SyntheticsStepDetailWarning object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsStepDetailWarning(message string, typeVar SyntheticsWarningType) *SyntheticsStepDetailWarning {
	this := SyntheticsStepDetailWarning{}
	this.Message = message
	this.Type = typeVar
	return &this
}

// NewSyntheticsStepDetailWarningWithDefaults instantiates a new SyntheticsStepDetailWarning object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsStepDetailWarningWithDefaults() *SyntheticsStepDetailWarning {
	this := SyntheticsStepDetailWarning{}
	return &this
}

// GetMessage returns the Message field value.
func (o *SyntheticsStepDetailWarning) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetailWarning) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *SyntheticsStepDetailWarning) SetMessage(v string) {
	o.Message = v
}

// GetType returns the Type field value.
func (o *SyntheticsStepDetailWarning) GetType() SyntheticsWarningType {
	if o == nil {
		var ret SyntheticsWarningType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetailWarning) GetTypeOk() (*SyntheticsWarningType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsStepDetailWarning) SetType(v SyntheticsWarningType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsStepDetailWarning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["message"] = o.Message
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsStepDetailWarning) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message *string                `json:"message"`
		Type    *SyntheticsWarningType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Message = *all.Message
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
