// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsCURConfigAttributes Attributes for An AWS CUR config.
type AwsCURConfigAttributes struct {
	// The account filtering configuration.
	AccountFilters *AccountFilteringConfig `json:"account_filters,omitempty"`
	// The AWS account ID.
	AccountId string `json:"account_id"`
	// The AWS bucket name used to store the Cost and Usage Report.
	BucketName string `json:"bucket_name"`
	// The region the bucket is located in.
	BucketRegion string `json:"bucket_region"`
	// The timestamp when the AWS CUR config was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// The error messages for the AWS CUR config.
	ErrorMessages datadog.NullableList[string] `json:"error_messages,omitempty"`
	// The number of months the report has been backfilled.
	// Deprecated
	Months *int32 `json:"months,omitempty"`
	// The name of the Cost and Usage Report.
	ReportName string `json:"report_name"`
	// The report prefix used for the Cost and Usage Report.
	ReportPrefix string `json:"report_prefix"`
	// The status of the AWS CUR.
	Status string `json:"status"`
	// The timestamp when the AWS CUR config status was updated.
	StatusUpdatedAt *string `json:"status_updated_at,omitempty"`
	// The timestamp when the AWS CUR config status was updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsCURConfigAttributes instantiates a new AwsCURConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsCURConfigAttributes(accountId string, bucketName string, bucketRegion string, reportName string, reportPrefix string, status string) *AwsCURConfigAttributes {
	this := AwsCURConfigAttributes{}
	this.AccountId = accountId
	this.BucketName = bucketName
	this.BucketRegion = bucketRegion
	this.ReportName = reportName
	this.ReportPrefix = reportPrefix
	this.Status = status
	return &this
}

// NewAwsCURConfigAttributesWithDefaults instantiates a new AwsCURConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsCURConfigAttributesWithDefaults() *AwsCURConfigAttributes {
	this := AwsCURConfigAttributes{}
	return &this
}

// GetAccountFilters returns the AccountFilters field value if set, zero value otherwise.
func (o *AwsCURConfigAttributes) GetAccountFilters() AccountFilteringConfig {
	if o == nil || o.AccountFilters == nil {
		var ret AccountFilteringConfig
		return ret
	}
	return *o.AccountFilters
}

// GetAccountFiltersOk returns a tuple with the AccountFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCURConfigAttributes) GetAccountFiltersOk() (*AccountFilteringConfig, bool) {
	if o == nil || o.AccountFilters == nil {
		return nil, false
	}
	return o.AccountFilters, true
}

// HasAccountFilters returns a boolean if a field has been set.
func (o *AwsCURConfigAttributes) HasAccountFilters() bool {
	return o != nil && o.AccountFilters != nil
}

// SetAccountFilters gets a reference to the given AccountFilteringConfig and assigns it to the AccountFilters field.
func (o *AwsCURConfigAttributes) SetAccountFilters(v AccountFilteringConfig) {
	o.AccountFilters = &v
}

// GetAccountId returns the AccountId field value.
func (o *AwsCURConfigAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *AwsCURConfigAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetBucketName returns the BucketName field value.
func (o *AwsCURConfigAttributes) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigAttributes) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value.
func (o *AwsCURConfigAttributes) SetBucketName(v string) {
	o.BucketName = v
}

// GetBucketRegion returns the BucketRegion field value.
func (o *AwsCURConfigAttributes) GetBucketRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketRegion
}

// GetBucketRegionOk returns a tuple with the BucketRegion field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigAttributes) GetBucketRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketRegion, true
}

// SetBucketRegion sets field value.
func (o *AwsCURConfigAttributes) SetBucketRegion(v string) {
	o.BucketRegion = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AwsCURConfigAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCURConfigAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AwsCURConfigAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AwsCURConfigAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsCURConfigAttributes) GetErrorMessages() []string {
	if o == nil || o.ErrorMessages.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ErrorMessages.Get()
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AwsCURConfigAttributes) GetErrorMessagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessages.Get(), o.ErrorMessages.IsSet()
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *AwsCURConfigAttributes) HasErrorMessages() bool {
	return o != nil && o.ErrorMessages.IsSet()
}

// SetErrorMessages gets a reference to the given datadog.NullableList[string] and assigns it to the ErrorMessages field.
func (o *AwsCURConfigAttributes) SetErrorMessages(v []string) {
	o.ErrorMessages.Set(&v)
}

// SetErrorMessagesNil sets the value for ErrorMessages to be an explicit nil.
func (o *AwsCURConfigAttributes) SetErrorMessagesNil() {
	o.ErrorMessages.Set(nil)
}

// UnsetErrorMessages ensures that no value is present for ErrorMessages, not even an explicit nil.
func (o *AwsCURConfigAttributes) UnsetErrorMessages() {
	o.ErrorMessages.Unset()
}

// GetMonths returns the Months field value if set, zero value otherwise.
// Deprecated
func (o *AwsCURConfigAttributes) GetMonths() int32 {
	if o == nil || o.Months == nil {
		var ret int32
		return ret
	}
	return *o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AwsCURConfigAttributes) GetMonthsOk() (*int32, bool) {
	if o == nil || o.Months == nil {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *AwsCURConfigAttributes) HasMonths() bool {
	return o != nil && o.Months != nil
}

// SetMonths gets a reference to the given int32 and assigns it to the Months field.
// Deprecated
func (o *AwsCURConfigAttributes) SetMonths(v int32) {
	o.Months = &v
}

// GetReportName returns the ReportName field value.
func (o *AwsCURConfigAttributes) GetReportName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReportName
}

// GetReportNameOk returns a tuple with the ReportName field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigAttributes) GetReportNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportName, true
}

// SetReportName sets field value.
func (o *AwsCURConfigAttributes) SetReportName(v string) {
	o.ReportName = v
}

// GetReportPrefix returns the ReportPrefix field value.
func (o *AwsCURConfigAttributes) GetReportPrefix() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReportPrefix
}

// GetReportPrefixOk returns a tuple with the ReportPrefix field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigAttributes) GetReportPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportPrefix, true
}

// SetReportPrefix sets field value.
func (o *AwsCURConfigAttributes) SetReportPrefix(v string) {
	o.ReportPrefix = v
}

// GetStatus returns the Status field value.
func (o *AwsCURConfigAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AwsCURConfigAttributes) SetStatus(v string) {
	o.Status = v
}

// GetStatusUpdatedAt returns the StatusUpdatedAt field value if set, zero value otherwise.
func (o *AwsCURConfigAttributes) GetStatusUpdatedAt() string {
	if o == nil || o.StatusUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.StatusUpdatedAt
}

// GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCURConfigAttributes) GetStatusUpdatedAtOk() (*string, bool) {
	if o == nil || o.StatusUpdatedAt == nil {
		return nil, false
	}
	return o.StatusUpdatedAt, true
}

// HasStatusUpdatedAt returns a boolean if a field has been set.
func (o *AwsCURConfigAttributes) HasStatusUpdatedAt() bool {
	return o != nil && o.StatusUpdatedAt != nil
}

// SetStatusUpdatedAt gets a reference to the given string and assigns it to the StatusUpdatedAt field.
func (o *AwsCURConfigAttributes) SetStatusUpdatedAt(v string) {
	o.StatusUpdatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AwsCURConfigAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCURConfigAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AwsCURConfigAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AwsCURConfigAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsCURConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountFilters != nil {
		toSerialize["account_filters"] = o.AccountFilters
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["bucket_name"] = o.BucketName
	toSerialize["bucket_region"] = o.BucketRegion
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ErrorMessages.IsSet() {
		toSerialize["error_messages"] = o.ErrorMessages.Get()
	}
	if o.Months != nil {
		toSerialize["months"] = o.Months
	}
	toSerialize["report_name"] = o.ReportName
	toSerialize["report_prefix"] = o.ReportPrefix
	toSerialize["status"] = o.Status
	if o.StatusUpdatedAt != nil {
		toSerialize["status_updated_at"] = o.StatusUpdatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AwsCURConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountFilters  *AccountFilteringConfig      `json:"account_filters,omitempty"`
		AccountId       *string                      `json:"account_id"`
		BucketName      *string                      `json:"bucket_name"`
		BucketRegion    *string                      `json:"bucket_region"`
		CreatedAt       *string                      `json:"created_at,omitempty"`
		ErrorMessages   datadog.NullableList[string] `json:"error_messages,omitempty"`
		Months          *int32                       `json:"months,omitempty"`
		ReportName      *string                      `json:"report_name"`
		ReportPrefix    *string                      `json:"report_prefix"`
		Status          *string                      `json:"status"`
		StatusUpdatedAt *string                      `json:"status_updated_at,omitempty"`
		UpdatedAt       *string                      `json:"updated_at,omitempty"`
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
	if all.ReportPrefix == nil {
		return fmt.Errorf("required field report_prefix missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_filters", "account_id", "bucket_name", "bucket_region", "created_at", "error_messages", "months", "report_name", "report_prefix", "status", "status_updated_at", "updated_at"})
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
	o.BucketRegion = *all.BucketRegion
	o.CreatedAt = all.CreatedAt
	o.ErrorMessages = all.ErrorMessages
	o.Months = all.Months
	o.ReportName = *all.ReportName
	o.ReportPrefix = *all.ReportPrefix
	o.Status = *all.Status
	o.StatusUpdatedAt = all.StatusUpdatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
