// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCcmConfigValidationRequestAttributes Attributes for an AWS CCM config validation request.
type AWSCcmConfigValidationRequestAttributes struct {
	// Your AWS Account ID without dashes.
	AccountId string `json:"account_id"`
	// Name of the S3 bucket where the Cost and Usage Report is stored.
	BucketName string `json:"bucket_name"`
	// AWS region of the S3 bucket.
	BucketRegion string `json:"bucket_region"`
	// Name of the Cost and Usage Report.
	ReportName string `json:"report_name"`
	// S3 prefix where the Cost and Usage Report is stored.
	ReportPrefix *string `json:"report_prefix,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSCcmConfigValidationRequestAttributes instantiates a new AWSCcmConfigValidationRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSCcmConfigValidationRequestAttributes(accountId string, bucketName string, bucketRegion string, reportName string) *AWSCcmConfigValidationRequestAttributes {
	this := AWSCcmConfigValidationRequestAttributes{}
	this.AccountId = accountId
	this.BucketName = bucketName
	this.BucketRegion = bucketRegion
	this.ReportName = reportName
	return &this
}

// NewAWSCcmConfigValidationRequestAttributesWithDefaults instantiates a new AWSCcmConfigValidationRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSCcmConfigValidationRequestAttributesWithDefaults() *AWSCcmConfigValidationRequestAttributes {
	this := AWSCcmConfigValidationRequestAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *AWSCcmConfigValidationRequestAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AWSCcmConfigValidationRequestAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *AWSCcmConfigValidationRequestAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetBucketName returns the BucketName field value.
func (o *AWSCcmConfigValidationRequestAttributes) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *AWSCcmConfigValidationRequestAttributes) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value.
func (o *AWSCcmConfigValidationRequestAttributes) SetBucketName(v string) {
	o.BucketName = v
}

// GetBucketRegion returns the BucketRegion field value.
func (o *AWSCcmConfigValidationRequestAttributes) GetBucketRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketRegion
}

// GetBucketRegionOk returns a tuple with the BucketRegion field value
// and a boolean to check if the value has been set.
func (o *AWSCcmConfigValidationRequestAttributes) GetBucketRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketRegion, true
}

// SetBucketRegion sets field value.
func (o *AWSCcmConfigValidationRequestAttributes) SetBucketRegion(v string) {
	o.BucketRegion = v
}

// GetReportName returns the ReportName field value.
func (o *AWSCcmConfigValidationRequestAttributes) GetReportName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReportName
}

// GetReportNameOk returns a tuple with the ReportName field value
// and a boolean to check if the value has been set.
func (o *AWSCcmConfigValidationRequestAttributes) GetReportNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportName, true
}

// SetReportName sets field value.
func (o *AWSCcmConfigValidationRequestAttributes) SetReportName(v string) {
	o.ReportName = v
}

// GetReportPrefix returns the ReportPrefix field value if set, zero value otherwise.
func (o *AWSCcmConfigValidationRequestAttributes) GetReportPrefix() string {
	if o == nil || o.ReportPrefix == nil {
		var ret string
		return ret
	}
	return *o.ReportPrefix
}

// GetReportPrefixOk returns a tuple with the ReportPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCcmConfigValidationRequestAttributes) GetReportPrefixOk() (*string, bool) {
	if o == nil || o.ReportPrefix == nil {
		return nil, false
	}
	return o.ReportPrefix, true
}

// HasReportPrefix returns a boolean if a field has been set.
func (o *AWSCcmConfigValidationRequestAttributes) HasReportPrefix() bool {
	return o != nil && o.ReportPrefix != nil
}

// SetReportPrefix gets a reference to the given string and assigns it to the ReportPrefix field.
func (o *AWSCcmConfigValidationRequestAttributes) SetReportPrefix(v string) {
	o.ReportPrefix = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSCcmConfigValidationRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["bucket_name"] = o.BucketName
	toSerialize["bucket_region"] = o.BucketRegion
	toSerialize["report_name"] = o.ReportName
	if o.ReportPrefix != nil {
		toSerialize["report_prefix"] = o.ReportPrefix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSCcmConfigValidationRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId    *string `json:"account_id"`
		BucketName   *string `json:"bucket_name"`
		BucketRegion *string `json:"bucket_region"`
		ReportName   *string `json:"report_name"`
		ReportPrefix *string `json:"report_prefix,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.BucketName == nil {
		return fmt.Errorf("required field bucket_name missing")
	}
	if all.BucketRegion == nil {
		return fmt.Errorf("required field bucket_region missing")
	}
	if all.ReportName == nil {
		return fmt.Errorf("required field report_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "bucket_name", "bucket_region", "report_name", "report_prefix"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.BucketName = *all.BucketName
	o.BucketRegion = *all.BucketRegion
	o.ReportName = *all.ReportName
	o.ReportPrefix = all.ReportPrefix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
