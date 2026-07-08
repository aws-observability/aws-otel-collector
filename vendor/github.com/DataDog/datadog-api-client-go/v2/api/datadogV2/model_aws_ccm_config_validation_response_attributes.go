// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCcmConfigValidationResponseAttributes Attributes for an AWS CCM config validation response.
type AWSCcmConfigValidationResponseAttributes struct {
	// Your AWS Account ID without dashes.
	AccountId string `json:"account_id"`
	// List of validation issues found for the Cost and Usage Report (CUR) 2.0 configuration. Empty when the configuration is valid.
	Issues []AWSCcmConfigValidationIssue `json:"issues"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSCcmConfigValidationResponseAttributes instantiates a new AWSCcmConfigValidationResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSCcmConfigValidationResponseAttributes(accountId string, issues []AWSCcmConfigValidationIssue) *AWSCcmConfigValidationResponseAttributes {
	this := AWSCcmConfigValidationResponseAttributes{}
	this.AccountId = accountId
	this.Issues = issues
	return &this
}

// NewAWSCcmConfigValidationResponseAttributesWithDefaults instantiates a new AWSCcmConfigValidationResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSCcmConfigValidationResponseAttributesWithDefaults() *AWSCcmConfigValidationResponseAttributes {
	this := AWSCcmConfigValidationResponseAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *AWSCcmConfigValidationResponseAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AWSCcmConfigValidationResponseAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *AWSCcmConfigValidationResponseAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetIssues returns the Issues field value.
func (o *AWSCcmConfigValidationResponseAttributes) GetIssues() []AWSCcmConfigValidationIssue {
	if o == nil {
		var ret []AWSCcmConfigValidationIssue
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value
// and a boolean to check if the value has been set.
func (o *AWSCcmConfigValidationResponseAttributes) GetIssuesOk() (*[]AWSCcmConfigValidationIssue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issues, true
}

// SetIssues sets field value.
func (o *AWSCcmConfigValidationResponseAttributes) SetIssues(v []AWSCcmConfigValidationIssue) {
	o.Issues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSCcmConfigValidationResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["issues"] = o.Issues

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSCcmConfigValidationResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId *string                        `json:"account_id"`
		Issues    *[]AWSCcmConfigValidationIssue `json:"issues"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.Issues == nil {
		return fmt.Errorf("required field issues missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "issues"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.Issues = *all.Issues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
