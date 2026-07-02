// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetAstRequestDataAttributes The attributes of the get-AST request, containing the source code to parse.
type GetAstRequestDataAttributes struct {
	// The base64-encoded source code to parse into an abstract syntax tree.
	Code string `json:"code"`
	// The encoding of the source code file (must be utf-8).
	FileEncoding string `json:"file_encoding"`
	// The programming language of the source code to parse.
	Language string `json:"language"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetAstRequestDataAttributes instantiates a new GetAstRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetAstRequestDataAttributes(code string, fileEncoding string, language string) *GetAstRequestDataAttributes {
	this := GetAstRequestDataAttributes{}
	this.Code = code
	this.FileEncoding = fileEncoding
	this.Language = language
	return &this
}

// NewGetAstRequestDataAttributesWithDefaults instantiates a new GetAstRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetAstRequestDataAttributesWithDefaults() *GetAstRequestDataAttributes {
	this := GetAstRequestDataAttributes{}
	return &this
}

// GetCode returns the Code field value.
func (o *GetAstRequestDataAttributes) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GetAstRequestDataAttributes) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *GetAstRequestDataAttributes) SetCode(v string) {
	o.Code = v
}

// GetFileEncoding returns the FileEncoding field value.
func (o *GetAstRequestDataAttributes) GetFileEncoding() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FileEncoding
}

// GetFileEncodingOk returns a tuple with the FileEncoding field value
// and a boolean to check if the value has been set.
func (o *GetAstRequestDataAttributes) GetFileEncodingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileEncoding, true
}

// SetFileEncoding sets field value.
func (o *GetAstRequestDataAttributes) SetFileEncoding(v string) {
	o.FileEncoding = v
}

// GetLanguage returns the Language field value.
func (o *GetAstRequestDataAttributes) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *GetAstRequestDataAttributes) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value.
func (o *GetAstRequestDataAttributes) SetLanguage(v string) {
	o.Language = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetAstRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	toSerialize["file_encoding"] = o.FileEncoding
	toSerialize["language"] = o.Language

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetAstRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code         *string `json:"code"`
		FileEncoding *string `json:"file_encoding"`
		Language     *string `json:"language"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.FileEncoding == nil {
		return fmt.Errorf("required field file_encoding missing")
	}
	if all.Language == nil {
		return fmt.Errorf("required field language missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code", "file_encoding", "language"})
	} else {
		return err
	}
	o.Code = *all.Code
	o.FileEncoding = *all.FileEncoding
	o.Language = *all.Language

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
