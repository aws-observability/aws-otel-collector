// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionJavascript A JavaScript assertion.
type SyntheticsAssertionJavascript struct {
	// The JavaScript code that performs the assertions.
	Code string `json:"code"`
	// Type of the assertion.
	Type SyntheticsAssertionJavascriptType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsAssertionJavascript instantiates a new SyntheticsAssertionJavascript object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsAssertionJavascript(code string, typeVar SyntheticsAssertionJavascriptType) *SyntheticsAssertionJavascript {
	this := SyntheticsAssertionJavascript{}
	this.Code = code
	this.Type = typeVar
	return &this
}

// NewSyntheticsAssertionJavascriptWithDefaults instantiates a new SyntheticsAssertionJavascript object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsAssertionJavascriptWithDefaults() *SyntheticsAssertionJavascript {
	this := SyntheticsAssertionJavascript{}
	return &this
}

// GetCode returns the Code field value.
func (o *SyntheticsAssertionJavascript) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionJavascript) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *SyntheticsAssertionJavascript) SetCode(v string) {
	o.Code = v
}

// GetType returns the Type field value.
func (o *SyntheticsAssertionJavascript) GetType() SyntheticsAssertionJavascriptType {
	if o == nil {
		var ret SyntheticsAssertionJavascriptType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertionJavascript) GetTypeOk() (*SyntheticsAssertionJavascriptType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsAssertionJavascript) SetType(v SyntheticsAssertionJavascriptType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsAssertionJavascript) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsAssertionJavascript) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code *string                            `json:"code"`
		Type *SyntheticsAssertionJavascriptType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Code = *all.Code
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
