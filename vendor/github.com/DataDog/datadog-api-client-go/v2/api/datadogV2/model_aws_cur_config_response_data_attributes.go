// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsCurConfigResponseDataAttributes The definition of `AwsCurConfigResponseDataAttributes` object.
type AwsCurConfigResponseDataAttributes struct {
	// The definition of `AwsCurConfigResponseDataAttributesAccountFilters` object.
	AccountFilters *AwsCurConfigResponseDataAttributesAccountFilters `json:"account_filters,omitempty"`
	// The `attributes` `account_id`.
	AccountId *string `json:"account_id,omitempty"`
	// The `attributes` `bucket_name`.
	BucketName *string `json:"bucket_name,omitempty"`
	// The `attributes` `bucket_region`.
	BucketRegion *string `json:"bucket_region,omitempty"`
	// The `attributes` `created_at`.
	CreatedAt *string `json:"created_at,omitempty"`
	// The `attributes` `error_messages`.
	ErrorMessages datadog.NullableList[string] `json:"error_messages,omitempty"`
	// The `attributes` `months`.
	Months *int64 `json:"months,omitempty"`
	// The `attributes` `report_name`.
	ReportName *string `json:"report_name,omitempty"`
	// The `attributes` `report_prefix`.
	ReportPrefix *string `json:"report_prefix,omitempty"`
	// The `attributes` `status`.
	Status *string `json:"status,omitempty"`
	// The `attributes` `status_updated_at`.
	StatusUpdatedAt *string `json:"status_updated_at,omitempty"`
	// The `attributes` `updated_at`.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsCurConfigResponseDataAttributes instantiates a new AwsCurConfigResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsCurConfigResponseDataAttributes() *AwsCurConfigResponseDataAttributes {
	this := AwsCurConfigResponseDataAttributes{}
	return &this
}

// NewAwsCurConfigResponseDataAttributesWithDefaults instantiates a new AwsCurConfigResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsCurConfigResponseDataAttributesWithDefaults() *AwsCurConfigResponseDataAttributes {
	this := AwsCurConfigResponseDataAttributes{}
	return &this
}

// GetAccountFilters returns the AccountFilters field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributes) GetAccountFilters() AwsCurConfigResponseDataAttributesAccountFilters {
	if o == nil || o.AccountFilters == nil {
		var ret AwsCurConfigResponseDataAttributesAccountFilters
		return ret
	}
	return *o.AccountFilters
}

// GetAccountFiltersOk returns a tuple with the AccountFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributes) GetAccountFiltersOk() (*AwsCurConfigResponseDataAttributesAccountFilters, bool) {
	if o == nil || o.AccountFilters == nil {
		return nil, false
	}
	return o.AccountFilters, true
}

// HasAccountFilters returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasAccountFilters() bool {
	return o != nil && o.AccountFilters != nil
}

// SetAccountFilters gets a reference to the given AwsCurConfigResponseDataAttributesAccountFilters and assigns it to the AccountFilters field.
func (o *AwsCurConfigResponseDataAttributes) SetAccountFilters(v AwsCurConfigResponseDataAttributesAccountFilters) {
	o.AccountFilters = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributes) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AwsCurConfigResponseDataAttributes) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributes) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributes) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasBucketName() bool {
	return o != nil && o.BucketName != nil
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *AwsCurConfigResponseDataAttributes) SetBucketName(v string) {
	o.BucketName = &v
}

// GetBucketRegion returns the BucketRegion field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributes) GetBucketRegion() string {
	if o == nil || o.BucketRegion == nil {
		var ret string
		return ret
	}
	return *o.BucketRegion
}

// GetBucketRegionOk returns a tuple with the BucketRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributes) GetBucketRegionOk() (*string, bool) {
	if o == nil || o.BucketRegion == nil {
		return nil, false
	}
	return o.BucketRegion, true
}

// HasBucketRegion returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasBucketRegion() bool {
	return o != nil && o.BucketRegion != nil
}

// SetBucketRegion gets a reference to the given string and assigns it to the BucketRegion field.
func (o *AwsCurConfigResponseDataAttributes) SetBucketRegion(v string) {
	o.BucketRegion = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AwsCurConfigResponseDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsCurConfigResponseDataAttributes) GetErrorMessages() []string {
	if o == nil || o.ErrorMessages.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ErrorMessages.Get()
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AwsCurConfigResponseDataAttributes) GetErrorMessagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessages.Get(), o.ErrorMessages.IsSet()
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasErrorMessages() bool {
	return o != nil && o.ErrorMessages.IsSet()
}

// SetErrorMessages gets a reference to the given datadog.NullableList[string] and assigns it to the ErrorMessages field.
func (o *AwsCurConfigResponseDataAttributes) SetErrorMessages(v []string) {
	o.ErrorMessages.Set(&v)
}

// SetErrorMessagesNil sets the value for ErrorMessages to be an explicit nil.
func (o *AwsCurConfigResponseDataAttributes) SetErrorMessagesNil() {
	o.ErrorMessages.Set(nil)
}

// UnsetErrorMessages ensures that no value is present for ErrorMessages, not even an explicit nil.
func (o *AwsCurConfigResponseDataAttributes) UnsetErrorMessages() {
	o.ErrorMessages.Unset()
}

// GetMonths returns the Months field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributes) GetMonths() int64 {
	if o == nil || o.Months == nil {
		var ret int64
		return ret
	}
	return *o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributes) GetMonthsOk() (*int64, bool) {
	if o == nil || o.Months == nil {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasMonths() bool {
	return o != nil && o.Months != nil
}

// SetMonths gets a reference to the given int64 and assigns it to the Months field.
func (o *AwsCurConfigResponseDataAttributes) SetMonths(v int64) {
	o.Months = &v
}

// GetReportName returns the ReportName field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributes) GetReportName() string {
	if o == nil || o.ReportName == nil {
		var ret string
		return ret
	}
	return *o.ReportName
}

// GetReportNameOk returns a tuple with the ReportName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributes) GetReportNameOk() (*string, bool) {
	if o == nil || o.ReportName == nil {
		return nil, false
	}
	return o.ReportName, true
}

// HasReportName returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasReportName() bool {
	return o != nil && o.ReportName != nil
}

// SetReportName gets a reference to the given string and assigns it to the ReportName field.
func (o *AwsCurConfigResponseDataAttributes) SetReportName(v string) {
	o.ReportName = &v
}

// GetReportPrefix returns the ReportPrefix field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributes) GetReportPrefix() string {
	if o == nil || o.ReportPrefix == nil {
		var ret string
		return ret
	}
	return *o.ReportPrefix
}

// GetReportPrefixOk returns a tuple with the ReportPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributes) GetReportPrefixOk() (*string, bool) {
	if o == nil || o.ReportPrefix == nil {
		return nil, false
	}
	return o.ReportPrefix, true
}

// HasReportPrefix returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasReportPrefix() bool {
	return o != nil && o.ReportPrefix != nil
}

// SetReportPrefix gets a reference to the given string and assigns it to the ReportPrefix field.
func (o *AwsCurConfigResponseDataAttributes) SetReportPrefix(v string) {
	o.ReportPrefix = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AwsCurConfigResponseDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetStatusUpdatedAt returns the StatusUpdatedAt field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributes) GetStatusUpdatedAt() string {
	if o == nil || o.StatusUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.StatusUpdatedAt
}

// GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributes) GetStatusUpdatedAtOk() (*string, bool) {
	if o == nil || o.StatusUpdatedAt == nil {
		return nil, false
	}
	return o.StatusUpdatedAt, true
}

// HasStatusUpdatedAt returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasStatusUpdatedAt() bool {
	return o != nil && o.StatusUpdatedAt != nil
}

// SetStatusUpdatedAt gets a reference to the given string and assigns it to the StatusUpdatedAt field.
func (o *AwsCurConfigResponseDataAttributes) SetStatusUpdatedAt(v string) {
	o.StatusUpdatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AwsCurConfigResponseDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsCurConfigResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountFilters != nil {
		toSerialize["account_filters"] = o.AccountFilters
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.BucketName != nil {
		toSerialize["bucket_name"] = o.BucketName
	}
	if o.BucketRegion != nil {
		toSerialize["bucket_region"] = o.BucketRegion
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ErrorMessages.IsSet() {
		toSerialize["error_messages"] = o.ErrorMessages.Get()
	}
	if o.Months != nil {
		toSerialize["months"] = o.Months
	}
	if o.ReportName != nil {
		toSerialize["report_name"] = o.ReportName
	}
	if o.ReportPrefix != nil {
		toSerialize["report_prefix"] = o.ReportPrefix
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
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
func (o *AwsCurConfigResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountFilters  *AwsCurConfigResponseDataAttributesAccountFilters `json:"account_filters,omitempty"`
		AccountId       *string                                           `json:"account_id,omitempty"`
		BucketName      *string                                           `json:"bucket_name,omitempty"`
		BucketRegion    *string                                           `json:"bucket_region,omitempty"`
		CreatedAt       *string                                           `json:"created_at,omitempty"`
		ErrorMessages   datadog.NullableList[string]                      `json:"error_messages,omitempty"`
		Months          *int64                                            `json:"months,omitempty"`
		ReportName      *string                                           `json:"report_name,omitempty"`
		ReportPrefix    *string                                           `json:"report_prefix,omitempty"`
		Status          *string                                           `json:"status,omitempty"`
		StatusUpdatedAt *string                                           `json:"status_updated_at,omitempty"`
		UpdatedAt       *string                                           `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.AccountId = all.AccountId
	o.BucketName = all.BucketName
	o.BucketRegion = all.BucketRegion
	o.CreatedAt = all.CreatedAt
	o.ErrorMessages = all.ErrorMessages
	o.Months = all.Months
	o.ReportName = all.ReportName
	o.ReportPrefix = all.ReportPrefix
	o.Status = all.Status
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
