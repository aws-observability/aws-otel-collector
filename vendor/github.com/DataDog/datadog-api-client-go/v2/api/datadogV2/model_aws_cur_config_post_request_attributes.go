// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsCURConfigPostRequestAttributes Attributes for AWS CUR config Post Request.
type AwsCURConfigPostRequestAttributes struct {
	// The account filtering configuration.
	AccountFilters *AccountFilteringConfig `json:"account_filters,omitempty"`
	// The AWS account ID.
	AccountId string `json:"account_id"`
	// The AWS bucket name used to store the Cost and Usage Report.
	BucketName string `json:"bucket_name"`
	// The region the bucket is located in.
	BucketRegion *string `json:"bucket_region,omitempty"`
	// The month of the report.
	Months *int32 `json:"months,omitempty"`
	// The name of the Cost and Usage Report.
	ReportName string `json:"report_name"`
	// The report prefix used for the Cost and Usage Report.
	ReportPrefix string `json:"report_prefix"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsCURConfigPostRequestAttributes instantiates a new AwsCURConfigPostRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsCURConfigPostRequestAttributes(accountId string, bucketName string, reportName string, reportPrefix string) *AwsCURConfigPostRequestAttributes {
	this := AwsCURConfigPostRequestAttributes{}
	this.AccountId = accountId
	this.BucketName = bucketName
	this.ReportName = reportName
	this.ReportPrefix = reportPrefix
	return &this
}

// NewAwsCURConfigPostRequestAttributesWithDefaults instantiates a new AwsCURConfigPostRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsCURConfigPostRequestAttributesWithDefaults() *AwsCURConfigPostRequestAttributes {
	this := AwsCURConfigPostRequestAttributes{}
	return &this
}

// GetAccountFilters returns the AccountFilters field value if set, zero value otherwise.
func (o *AwsCURConfigPostRequestAttributes) GetAccountFilters() AccountFilteringConfig {
	if o == nil || o.AccountFilters == nil {
		var ret AccountFilteringConfig
		return ret
	}
	return *o.AccountFilters
}

// GetAccountFiltersOk returns a tuple with the AccountFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCURConfigPostRequestAttributes) GetAccountFiltersOk() (*AccountFilteringConfig, bool) {
	if o == nil || o.AccountFilters == nil {
		return nil, false
	}
	return o.AccountFilters, true
}

// HasAccountFilters returns a boolean if a field has been set.
func (o *AwsCURConfigPostRequestAttributes) HasAccountFilters() bool {
	return o != nil && o.AccountFilters != nil
}

// SetAccountFilters gets a reference to the given AccountFilteringConfig and assigns it to the AccountFilters field.
func (o *AwsCURConfigPostRequestAttributes) SetAccountFilters(v AccountFilteringConfig) {
	o.AccountFilters = &v
}

// GetAccountId returns the AccountId field value.
func (o *AwsCURConfigPostRequestAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigPostRequestAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *AwsCURConfigPostRequestAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetBucketName returns the BucketName field value.
func (o *AwsCURConfigPostRequestAttributes) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigPostRequestAttributes) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value.
func (o *AwsCURConfigPostRequestAttributes) SetBucketName(v string) {
	o.BucketName = v
}

// GetBucketRegion returns the BucketRegion field value if set, zero value otherwise.
func (o *AwsCURConfigPostRequestAttributes) GetBucketRegion() string {
	if o == nil || o.BucketRegion == nil {
		var ret string
		return ret
	}
	return *o.BucketRegion
}

// GetBucketRegionOk returns a tuple with the BucketRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCURConfigPostRequestAttributes) GetBucketRegionOk() (*string, bool) {
	if o == nil || o.BucketRegion == nil {
		return nil, false
	}
	return o.BucketRegion, true
}

// HasBucketRegion returns a boolean if a field has been set.
func (o *AwsCURConfigPostRequestAttributes) HasBucketRegion() bool {
	return o != nil && o.BucketRegion != nil
}

// SetBucketRegion gets a reference to the given string and assigns it to the BucketRegion field.
func (o *AwsCURConfigPostRequestAttributes) SetBucketRegion(v string) {
	o.BucketRegion = &v
}

// GetMonths returns the Months field value if set, zero value otherwise.
func (o *AwsCURConfigPostRequestAttributes) GetMonths() int32 {
	if o == nil || o.Months == nil {
		var ret int32
		return ret
	}
	return *o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCURConfigPostRequestAttributes) GetMonthsOk() (*int32, bool) {
	if o == nil || o.Months == nil {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *AwsCURConfigPostRequestAttributes) HasMonths() bool {
	return o != nil && o.Months != nil
}

// SetMonths gets a reference to the given int32 and assigns it to the Months field.
func (o *AwsCURConfigPostRequestAttributes) SetMonths(v int32) {
	o.Months = &v
}

// GetReportName returns the ReportName field value.
func (o *AwsCURConfigPostRequestAttributes) GetReportName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReportName
}

// GetReportNameOk returns a tuple with the ReportName field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigPostRequestAttributes) GetReportNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportName, true
}

// SetReportName sets field value.
func (o *AwsCURConfigPostRequestAttributes) SetReportName(v string) {
	o.ReportName = v
}

// GetReportPrefix returns the ReportPrefix field value.
func (o *AwsCURConfigPostRequestAttributes) GetReportPrefix() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReportPrefix
}

// GetReportPrefixOk returns a tuple with the ReportPrefix field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigPostRequestAttributes) GetReportPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportPrefix, true
}

// SetReportPrefix sets field value.
func (o *AwsCURConfigPostRequestAttributes) SetReportPrefix(v string) {
	o.ReportPrefix = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsCURConfigPostRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountFilters != nil {
		toSerialize["account_filters"] = o.AccountFilters
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["bucket_name"] = o.BucketName
	if o.BucketRegion != nil {
		toSerialize["bucket_region"] = o.BucketRegion
	}
	if o.Months != nil {
		toSerialize["months"] = o.Months
	}
	toSerialize["report_name"] = o.ReportName
	toSerialize["report_prefix"] = o.ReportPrefix

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AwsCURConfigPostRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountFilters *AccountFilteringConfig `json:"account_filters,omitempty"`
		AccountId      *string                 `json:"account_id"`
		BucketName     *string                 `json:"bucket_name"`
		BucketRegion   *string                 `json:"bucket_region,omitempty"`
		Months         *int32                  `json:"months,omitempty"`
		ReportName     *string                 `json:"report_name"`
		ReportPrefix   *string                 `json:"report_prefix"`
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
	if all.ReportName == nil {
		return fmt.Errorf("required field report_name missing")
	}
	if all.ReportPrefix == nil {
		return fmt.Errorf("required field report_prefix missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_filters", "account_id", "bucket_name", "bucket_region", "months", "report_name", "report_prefix"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AccountFilters != nil && all.AccountFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AccountFilters = all.AccountFilters
	o.AccountId = *all.AccountId
	o.BucketName = *all.BucketName
	o.BucketRegion = all.BucketRegion
	o.Months = all.Months
	o.ReportName = *all.ReportName
	o.ReportPrefix = *all.ReportPrefix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
