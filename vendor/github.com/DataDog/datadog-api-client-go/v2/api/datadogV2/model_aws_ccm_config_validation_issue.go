// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCcmConfigValidationIssue A single validation issue found while validating an AWS Cost and Usage Report (CUR) 2.0 configuration.
type AWSCcmConfigValidationIssue struct {
	// Identifies the specific reason a Cost and Usage Report (CUR) 2.0 configuration failed validation.
	Code AWSCcmConfigValidationIssueCode `json:"code"`
	// Human-readable description of the validation issue.
	Description string `json:"description"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSCcmConfigValidationIssue instantiates a new AWSCcmConfigValidationIssue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSCcmConfigValidationIssue(code AWSCcmConfigValidationIssueCode, description string) *AWSCcmConfigValidationIssue {
	this := AWSCcmConfigValidationIssue{}
	this.Code = code
	this.Description = description
	return &this
}

// NewAWSCcmConfigValidationIssueWithDefaults instantiates a new AWSCcmConfigValidationIssue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSCcmConfigValidationIssueWithDefaults() *AWSCcmConfigValidationIssue {
	this := AWSCcmConfigValidationIssue{}
	return &this
}

// GetCode returns the Code field value.
func (o *AWSCcmConfigValidationIssue) GetCode() AWSCcmConfigValidationIssueCode {
	if o == nil {
		var ret AWSCcmConfigValidationIssueCode
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AWSCcmConfigValidationIssue) GetCodeOk() (*AWSCcmConfigValidationIssueCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *AWSCcmConfigValidationIssue) SetCode(v AWSCcmConfigValidationIssueCode) {
	o.Code = v
}

// GetDescription returns the Description field value.
func (o *AWSCcmConfigValidationIssue) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AWSCcmConfigValidationIssue) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *AWSCcmConfigValidationIssue) SetDescription(v string) {
	o.Description = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSCcmConfigValidationIssue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSCcmConfigValidationIssue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code        *AWSCcmConfigValidationIssueCode `json:"code"`
		Description *string                          `json:"description"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code", "description"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Code.IsValid() {
		hasInvalidField = true
	} else {
		o.Code = *all.Code
	}
	o.Description = *all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
