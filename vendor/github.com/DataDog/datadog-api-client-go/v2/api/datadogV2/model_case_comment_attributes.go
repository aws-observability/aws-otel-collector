// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseCommentAttributes Case comment attributes
type CaseCommentAttributes struct {
	// The `CaseCommentAttributes` `message`.
	Comment string `json:"comment"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseCommentAttributes instantiates a new CaseCommentAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseCommentAttributes(comment string) *CaseCommentAttributes {
	this := CaseCommentAttributes{}
	this.Comment = comment
	return &this
}

// NewCaseCommentAttributesWithDefaults instantiates a new CaseCommentAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseCommentAttributesWithDefaults() *CaseCommentAttributes {
	this := CaseCommentAttributes{}
	return &this
}

// GetComment returns the Comment field value.
func (o *CaseCommentAttributes) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *CaseCommentAttributes) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value.
func (o *CaseCommentAttributes) SetComment(v string) {
	o.Comment = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseCommentAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["comment"] = o.Comment

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseCommentAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Comment *string `json:"comment"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Comment == nil {
		return fmt.Errorf("required field comment missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"comment"})
	} else {
		return err
	}
	o.Comment = *all.Comment

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
