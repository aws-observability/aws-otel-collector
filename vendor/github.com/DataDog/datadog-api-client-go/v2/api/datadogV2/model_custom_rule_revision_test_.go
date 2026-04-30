// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRuleRevisionTest A test case associated with a custom rule revision, used to validate rule behavior.
type CustomRuleRevisionTest struct {
	// Expected violation count
	AnnotationCount int64 `json:"annotation_count"`
	// Test code
	Code string `json:"code"`
	// Test filename
	Filename string `json:"filename"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomRuleRevisionTest instantiates a new CustomRuleRevisionTest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomRuleRevisionTest(annotationCount int64, code string, filename string) *CustomRuleRevisionTest {
	this := CustomRuleRevisionTest{}
	this.AnnotationCount = annotationCount
	this.Code = code
	this.Filename = filename
	return &this
}

// NewCustomRuleRevisionTestWithDefaults instantiates a new CustomRuleRevisionTest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomRuleRevisionTestWithDefaults() *CustomRuleRevisionTest {
	this := CustomRuleRevisionTest{}
	return &this
}

// GetAnnotationCount returns the AnnotationCount field value.
func (o *CustomRuleRevisionTest) GetAnnotationCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AnnotationCount
}

// GetAnnotationCountOk returns a tuple with the AnnotationCount field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionTest) GetAnnotationCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnnotationCount, true
}

// SetAnnotationCount sets field value.
func (o *CustomRuleRevisionTest) SetAnnotationCount(v int64) {
	o.AnnotationCount = v
}

// GetCode returns the Code field value.
func (o *CustomRuleRevisionTest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionTest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *CustomRuleRevisionTest) SetCode(v string) {
	o.Code = v
}

// GetFilename returns the Filename field value.
func (o *CustomRuleRevisionTest) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionTest) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value.
func (o *CustomRuleRevisionTest) SetFilename(v string) {
	o.Filename = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomRuleRevisionTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotation_count"] = o.AnnotationCount
	toSerialize["code"] = o.Code
	toSerialize["filename"] = o.Filename

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomRuleRevisionTest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotationCount *int64  `json:"annotation_count"`
		Code            *string `json:"code"`
		Filename        *string `json:"filename"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AnnotationCount == nil {
		return fmt.Errorf("required field annotation_count missing")
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.Filename == nil {
		return fmt.Errorf("required field filename missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotation_count", "code", "filename"})
	} else {
		return err
	}
	o.AnnotationCount = *all.AnnotationCount
	o.Code = *all.Code
	o.Filename = *all.Filename

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
