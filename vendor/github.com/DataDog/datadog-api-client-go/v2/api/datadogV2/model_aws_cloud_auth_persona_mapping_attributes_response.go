// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCloudAuthPersonaMappingAttributesResponse Attributes for AWS cloud authentication persona mapping response
type AWSCloudAuthPersonaMappingAttributesResponse struct {
	// Datadog account identifier (email or handle) mapped to the AWS principal
	AccountIdentifier string `json:"account_identifier"`
	// Datadog account UUID
	AccountUuid string `json:"account_uuid"`
	// AWS IAM ARN pattern to match for authentication
	ArnPattern string `json:"arn_pattern"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSCloudAuthPersonaMappingAttributesResponse instantiates a new AWSCloudAuthPersonaMappingAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSCloudAuthPersonaMappingAttributesResponse(accountIdentifier string, accountUuid string, arnPattern string) *AWSCloudAuthPersonaMappingAttributesResponse {
	this := AWSCloudAuthPersonaMappingAttributesResponse{}
	this.AccountIdentifier = accountIdentifier
	this.AccountUuid = accountUuid
	this.ArnPattern = arnPattern
	return &this
}

// NewAWSCloudAuthPersonaMappingAttributesResponseWithDefaults instantiates a new AWSCloudAuthPersonaMappingAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSCloudAuthPersonaMappingAttributesResponseWithDefaults() *AWSCloudAuthPersonaMappingAttributesResponse {
	this := AWSCloudAuthPersonaMappingAttributesResponse{}
	return &this
}

// GetAccountIdentifier returns the AccountIdentifier field value.
func (o *AWSCloudAuthPersonaMappingAttributesResponse) GetAccountIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *AWSCloudAuthPersonaMappingAttributesResponse) GetAccountIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentifier, true
}

// SetAccountIdentifier sets field value.
func (o *AWSCloudAuthPersonaMappingAttributesResponse) SetAccountIdentifier(v string) {
	o.AccountIdentifier = v
}

// GetAccountUuid returns the AccountUuid field value.
func (o *AWSCloudAuthPersonaMappingAttributesResponse) GetAccountUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountUuid
}

// GetAccountUuidOk returns a tuple with the AccountUuid field value
// and a boolean to check if the value has been set.
func (o *AWSCloudAuthPersonaMappingAttributesResponse) GetAccountUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountUuid, true
}

// SetAccountUuid sets field value.
func (o *AWSCloudAuthPersonaMappingAttributesResponse) SetAccountUuid(v string) {
	o.AccountUuid = v
}

// GetArnPattern returns the ArnPattern field value.
func (o *AWSCloudAuthPersonaMappingAttributesResponse) GetArnPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ArnPattern
}

// GetArnPatternOk returns a tuple with the ArnPattern field value
// and a boolean to check if the value has been set.
func (o *AWSCloudAuthPersonaMappingAttributesResponse) GetArnPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArnPattern, true
}

// SetArnPattern sets field value.
func (o *AWSCloudAuthPersonaMappingAttributesResponse) SetArnPattern(v string) {
	o.ArnPattern = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSCloudAuthPersonaMappingAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_identifier"] = o.AccountIdentifier
	toSerialize["account_uuid"] = o.AccountUuid
	toSerialize["arn_pattern"] = o.ArnPattern

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSCloudAuthPersonaMappingAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountIdentifier *string `json:"account_identifier"`
		AccountUuid       *string `json:"account_uuid"`
		ArnPattern        *string `json:"arn_pattern"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountIdentifier == nil {
		return fmt.Errorf("required field account_identifier missing")
	}
	if all.AccountUuid == nil {
		return fmt.Errorf("required field account_uuid missing")
	}
	if all.ArnPattern == nil {
		return fmt.Errorf("required field arn_pattern missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_identifier", "account_uuid", "arn_pattern"})
	} else {
		return err
	}
	o.AccountIdentifier = *all.AccountIdentifier
	o.AccountUuid = *all.AccountUuid
	o.ArnPattern = *all.ArnPattern

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
