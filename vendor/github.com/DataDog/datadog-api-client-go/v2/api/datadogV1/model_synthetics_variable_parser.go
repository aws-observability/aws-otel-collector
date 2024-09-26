// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsVariableParser Details of the parser to use for the global variable.
type SyntheticsVariableParser struct {
	// Type of parser for a Synthetic global variable from a synthetics test.
	Type SyntheticsGlobalVariableParserType `json:"type"`
	// Regex or JSON path used for the parser. Not used with type `raw`.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsVariableParser instantiates a new SyntheticsVariableParser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsVariableParser(typeVar SyntheticsGlobalVariableParserType) *SyntheticsVariableParser {
	this := SyntheticsVariableParser{}
	this.Type = typeVar
	return &this
}

// NewSyntheticsVariableParserWithDefaults instantiates a new SyntheticsVariableParser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsVariableParserWithDefaults() *SyntheticsVariableParser {
	this := SyntheticsVariableParser{}
	return &this
}

// GetType returns the Type field value.
func (o *SyntheticsVariableParser) GetType() SyntheticsGlobalVariableParserType {
	if o == nil {
		var ret SyntheticsGlobalVariableParserType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsVariableParser) GetTypeOk() (*SyntheticsGlobalVariableParserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsVariableParser) SetType(v SyntheticsGlobalVariableParserType) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SyntheticsVariableParser) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsVariableParser) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SyntheticsVariableParser) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SyntheticsVariableParser) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsVariableParser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsVariableParser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *SyntheticsGlobalVariableParserType `json:"type"`
		Value *string                             `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
