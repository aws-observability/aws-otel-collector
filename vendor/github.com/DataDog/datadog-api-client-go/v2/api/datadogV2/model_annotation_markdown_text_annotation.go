// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnnotationMarkdownTextAnnotation The definition of `AnnotationMarkdownTextAnnotation` object.
type AnnotationMarkdownTextAnnotation struct {
	// The `markdownTextAnnotation` `text`.
	Text *string `json:"text,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnnotationMarkdownTextAnnotation instantiates a new AnnotationMarkdownTextAnnotation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnnotationMarkdownTextAnnotation() *AnnotationMarkdownTextAnnotation {
	this := AnnotationMarkdownTextAnnotation{}
	return &this
}

// NewAnnotationMarkdownTextAnnotationWithDefaults instantiates a new AnnotationMarkdownTextAnnotation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnnotationMarkdownTextAnnotationWithDefaults() *AnnotationMarkdownTextAnnotation {
	this := AnnotationMarkdownTextAnnotation{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AnnotationMarkdownTextAnnotation) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationMarkdownTextAnnotation) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AnnotationMarkdownTextAnnotation) HasText() bool {
	return o != nil && o.Text != nil
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AnnotationMarkdownTextAnnotation) SetText(v string) {
	o.Text = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnnotationMarkdownTextAnnotation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnnotationMarkdownTextAnnotation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Text *string `json:"text,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"text"})
	} else {
		return err
	}
	o.Text = all.Text

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
