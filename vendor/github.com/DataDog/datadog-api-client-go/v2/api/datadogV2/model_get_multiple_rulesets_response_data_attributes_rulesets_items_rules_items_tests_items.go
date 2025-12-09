// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems
type GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems struct {
	//
	AnnotationCount *int64 `json:"annotation_count,omitempty"`
	//
	Code *string `json:"code,omitempty"`
	//
	Filename *string `json:"filename,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems() *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems{}
	return &this
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItemsWithDefaults instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItemsWithDefaults() *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems{}
	return &this
}

// GetAnnotationCount returns the AnnotationCount field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) GetAnnotationCount() int64 {
	if o == nil || o.AnnotationCount == nil {
		var ret int64
		return ret
	}
	return *o.AnnotationCount
}

// GetAnnotationCountOk returns a tuple with the AnnotationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) GetAnnotationCountOk() (*int64, bool) {
	if o == nil || o.AnnotationCount == nil {
		return nil, false
	}
	return o.AnnotationCount, true
}

// HasAnnotationCount returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) HasAnnotationCount() bool {
	return o != nil && o.AnnotationCount != nil
}

// SetAnnotationCount gets a reference to the given int64 and assigns it to the AnnotationCount field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) SetAnnotationCount(v int64) {
	o.AnnotationCount = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) HasCode() bool {
	return o != nil && o.Code != nil
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) SetCode(v string) {
	o.Code = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) HasFilename() bool {
	return o != nil && o.Filename != nil
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) SetFilename(v string) {
	o.Filename = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnnotationCount != nil {
		toSerialize["annotation_count"] = o.AnnotationCount
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotationCount *int64  `json:"annotation_count,omitempty"`
		Code            *string `json:"code,omitempty"`
		Filename        *string `json:"filename,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotation_count", "code", "filename"})
	} else {
		return err
	}
	o.AnnotationCount = all.AnnotationCount
	o.Code = all.Code
	o.Filename = all.Filename

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
