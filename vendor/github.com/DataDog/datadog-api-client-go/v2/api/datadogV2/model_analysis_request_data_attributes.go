// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisRequestDataAttributes The attributes of the analysis request, containing the source code and rules to apply.
type AnalysisRequestDataAttributes struct {
	// The base64-encoded source code to analyze.
	Code string `json:"code"`
	// The encoding of the source code file (must be `utf-8`).
	FileEncoding string `json:"file_encoding"`
	// The name of the file being analyzed.
	Filename string `json:"filename"`
	// The programming language of the source code.
	Language string `json:"language"`
	// The list of static analysis rules to apply during analysis.
	Rules []AnalysisRequestRule `json:"rules"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnalysisRequestDataAttributes instantiates a new AnalysisRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisRequestDataAttributes(code string, fileEncoding string, filename string, language string, rules []AnalysisRequestRule) *AnalysisRequestDataAttributes {
	this := AnalysisRequestDataAttributes{}
	this.Code = code
	this.FileEncoding = fileEncoding
	this.Filename = filename
	this.Language = language
	this.Rules = rules
	return &this
}

// NewAnalysisRequestDataAttributesWithDefaults instantiates a new AnalysisRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisRequestDataAttributesWithDefaults() *AnalysisRequestDataAttributes {
	this := AnalysisRequestDataAttributes{}
	return &this
}

// GetCode returns the Code field value.
func (o *AnalysisRequestDataAttributes) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestDataAttributes) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *AnalysisRequestDataAttributes) SetCode(v string) {
	o.Code = v
}

// GetFileEncoding returns the FileEncoding field value.
func (o *AnalysisRequestDataAttributes) GetFileEncoding() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FileEncoding
}

// GetFileEncodingOk returns a tuple with the FileEncoding field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestDataAttributes) GetFileEncodingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileEncoding, true
}

// SetFileEncoding sets field value.
func (o *AnalysisRequestDataAttributes) SetFileEncoding(v string) {
	o.FileEncoding = v
}

// GetFilename returns the Filename field value.
func (o *AnalysisRequestDataAttributes) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestDataAttributes) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value.
func (o *AnalysisRequestDataAttributes) SetFilename(v string) {
	o.Filename = v
}

// GetLanguage returns the Language field value.
func (o *AnalysisRequestDataAttributes) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestDataAttributes) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value.
func (o *AnalysisRequestDataAttributes) SetLanguage(v string) {
	o.Language = v
}

// GetRules returns the Rules field value.
func (o *AnalysisRequestDataAttributes) GetRules() []AnalysisRequestRule {
	if o == nil {
		var ret []AnalysisRequestRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestDataAttributes) GetRulesOk() (*[]AnalysisRequestRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *AnalysisRequestDataAttributes) SetRules(v []AnalysisRequestRule) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	toSerialize["file_encoding"] = o.FileEncoding
	toSerialize["filename"] = o.Filename
	toSerialize["language"] = o.Language
	toSerialize["rules"] = o.Rules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code         *string                `json:"code"`
		FileEncoding *string                `json:"file_encoding"`
		Filename     *string                `json:"filename"`
		Language     *string                `json:"language"`
		Rules        *[]AnalysisRequestRule `json:"rules"`
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
	if all.Filename == nil {
		return fmt.Errorf("required field filename missing")
	}
	if all.Language == nil {
		return fmt.Errorf("required field language missing")
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code", "file_encoding", "filename", "language", "rules"})
	} else {
		return err
	}
	o.Code = *all.Code
	o.FileEncoding = *all.FileEncoding
	o.Filename = *all.Filename
	o.Language = *all.Language
	o.Rules = *all.Rules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
