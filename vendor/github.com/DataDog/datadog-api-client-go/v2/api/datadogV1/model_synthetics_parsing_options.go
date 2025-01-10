// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsParsingOptions Parsing options for variables to extract.
type SyntheticsParsingOptions struct {
	// When type is `http_header` or `grpc_metadata`, name of the header or metadatum to extract.
	Field *string `json:"field,omitempty"`
	// Name of the variable to extract.
	Name *string `json:"name,omitempty"`
	// Details of the parser to use for the global variable.
	Parser *SyntheticsVariableParser `json:"parser,omitempty"`
	// Determines whether or not the extracted value will be obfuscated.
	Secure *bool `json:"secure,omitempty"`
	// Property of the Synthetic Test Response to extract into a local variable.
	Type *SyntheticsLocalVariableParsingOptionsType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsParsingOptions instantiates a new SyntheticsParsingOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsParsingOptions() *SyntheticsParsingOptions {
	this := SyntheticsParsingOptions{}
	return &this
}

// NewSyntheticsParsingOptionsWithDefaults instantiates a new SyntheticsParsingOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsParsingOptionsWithDefaults() *SyntheticsParsingOptions {
	this := SyntheticsParsingOptions{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *SyntheticsParsingOptions) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsParsingOptions) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *SyntheticsParsingOptions) HasField() bool {
	return o != nil && o.Field != nil
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *SyntheticsParsingOptions) SetField(v string) {
	o.Field = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsParsingOptions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsParsingOptions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsParsingOptions) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsParsingOptions) SetName(v string) {
	o.Name = &v
}

// GetParser returns the Parser field value if set, zero value otherwise.
func (o *SyntheticsParsingOptions) GetParser() SyntheticsVariableParser {
	if o == nil || o.Parser == nil {
		var ret SyntheticsVariableParser
		return ret
	}
	return *o.Parser
}

// GetParserOk returns a tuple with the Parser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsParsingOptions) GetParserOk() (*SyntheticsVariableParser, bool) {
	if o == nil || o.Parser == nil {
		return nil, false
	}
	return o.Parser, true
}

// HasParser returns a boolean if a field has been set.
func (o *SyntheticsParsingOptions) HasParser() bool {
	return o != nil && o.Parser != nil
}

// SetParser gets a reference to the given SyntheticsVariableParser and assigns it to the Parser field.
func (o *SyntheticsParsingOptions) SetParser(v SyntheticsVariableParser) {
	o.Parser = &v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *SyntheticsParsingOptions) GetSecure() bool {
	if o == nil || o.Secure == nil {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsParsingOptions) GetSecureOk() (*bool, bool) {
	if o == nil || o.Secure == nil {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *SyntheticsParsingOptions) HasSecure() bool {
	return o != nil && o.Secure != nil
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *SyntheticsParsingOptions) SetSecure(v bool) {
	o.Secure = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsParsingOptions) GetType() SyntheticsLocalVariableParsingOptionsType {
	if o == nil || o.Type == nil {
		var ret SyntheticsLocalVariableParsingOptionsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsParsingOptions) GetTypeOk() (*SyntheticsLocalVariableParsingOptionsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsParsingOptions) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SyntheticsLocalVariableParsingOptionsType and assigns it to the Type field.
func (o *SyntheticsParsingOptions) SetType(v SyntheticsLocalVariableParsingOptionsType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsParsingOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Parser != nil {
		toSerialize["parser"] = o.Parser
	}
	if o.Secure != nil {
		toSerialize["secure"] = o.Secure
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
func (o *SyntheticsParsingOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field  *string                                    `json:"field,omitempty"`
		Name   *string                                    `json:"name,omitempty"`
		Parser *SyntheticsVariableParser                  `json:"parser,omitempty"`
		Secure *bool                                      `json:"secure,omitempty"`
		Type   *SyntheticsLocalVariableParsingOptionsType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field", "name", "parser", "secure", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Field = all.Field
	o.Name = all.Name
	if all.Parser != nil && all.Parser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Parser = all.Parser
	o.Secure = all.Secure
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
