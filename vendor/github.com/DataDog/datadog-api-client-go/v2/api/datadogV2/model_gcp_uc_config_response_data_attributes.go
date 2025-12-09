// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GcpUcConfigResponseDataAttributes The definition of `GcpUcConfigResponseDataAttributes` object.
type GcpUcConfigResponseDataAttributes struct {
	// The `attributes` `account_id`.
	AccountId *string `json:"account_id,omitempty"`
	// The `attributes` `bucket_name`.
	BucketName *string `json:"bucket_name,omitempty"`
	// The `attributes` `created_at`.
	CreatedAt *string `json:"created_at,omitempty"`
	// The `attributes` `dataset`.
	Dataset *string `json:"dataset,omitempty"`
	// The `attributes` `error_messages`.
	ErrorMessages datadog.NullableList[string] `json:"error_messages,omitempty"`
	// The `attributes` `export_prefix`.
	ExportPrefix *string `json:"export_prefix,omitempty"`
	// The `attributes` `export_project_name`.
	ExportProjectName *string `json:"export_project_name,omitempty"`
	// The `attributes` `months`.
	Months *int64 `json:"months,omitempty"`
	// The `attributes` `project_id`.
	ProjectId *string `json:"project_id,omitempty"`
	// The `attributes` `service_account`.
	ServiceAccount *string `json:"service_account,omitempty"`
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

// NewGcpUcConfigResponseDataAttributes instantiates a new GcpUcConfigResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGcpUcConfigResponseDataAttributes() *GcpUcConfigResponseDataAttributes {
	this := GcpUcConfigResponseDataAttributes{}
	return &this
}

// NewGcpUcConfigResponseDataAttributesWithDefaults instantiates a new GcpUcConfigResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGcpUcConfigResponseDataAttributesWithDefaults() *GcpUcConfigResponseDataAttributes {
	this := GcpUcConfigResponseDataAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *GcpUcConfigResponseDataAttributes) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasBucketName() bool {
	return o != nil && o.BucketName != nil
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *GcpUcConfigResponseDataAttributes) SetBucketName(v string) {
	o.BucketName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GcpUcConfigResponseDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDataset returns the Dataset field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetDataset() string {
	if o == nil || o.Dataset == nil {
		var ret string
		return ret
	}
	return *o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetDatasetOk() (*string, bool) {
	if o == nil || o.Dataset == nil {
		return nil, false
	}
	return o.Dataset, true
}

// HasDataset returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasDataset() bool {
	return o != nil && o.Dataset != nil
}

// SetDataset gets a reference to the given string and assigns it to the Dataset field.
func (o *GcpUcConfigResponseDataAttributes) SetDataset(v string) {
	o.Dataset = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GcpUcConfigResponseDataAttributes) GetErrorMessages() []string {
	if o == nil || o.ErrorMessages.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ErrorMessages.Get()
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GcpUcConfigResponseDataAttributes) GetErrorMessagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessages.Get(), o.ErrorMessages.IsSet()
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasErrorMessages() bool {
	return o != nil && o.ErrorMessages.IsSet()
}

// SetErrorMessages gets a reference to the given datadog.NullableList[string] and assigns it to the ErrorMessages field.
func (o *GcpUcConfigResponseDataAttributes) SetErrorMessages(v []string) {
	o.ErrorMessages.Set(&v)
}

// SetErrorMessagesNil sets the value for ErrorMessages to be an explicit nil.
func (o *GcpUcConfigResponseDataAttributes) SetErrorMessagesNil() {
	o.ErrorMessages.Set(nil)
}

// UnsetErrorMessages ensures that no value is present for ErrorMessages, not even an explicit nil.
func (o *GcpUcConfigResponseDataAttributes) UnsetErrorMessages() {
	o.ErrorMessages.Unset()
}

// GetExportPrefix returns the ExportPrefix field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetExportPrefix() string {
	if o == nil || o.ExportPrefix == nil {
		var ret string
		return ret
	}
	return *o.ExportPrefix
}

// GetExportPrefixOk returns a tuple with the ExportPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetExportPrefixOk() (*string, bool) {
	if o == nil || o.ExportPrefix == nil {
		return nil, false
	}
	return o.ExportPrefix, true
}

// HasExportPrefix returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasExportPrefix() bool {
	return o != nil && o.ExportPrefix != nil
}

// SetExportPrefix gets a reference to the given string and assigns it to the ExportPrefix field.
func (o *GcpUcConfigResponseDataAttributes) SetExportPrefix(v string) {
	o.ExportPrefix = &v
}

// GetExportProjectName returns the ExportProjectName field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetExportProjectName() string {
	if o == nil || o.ExportProjectName == nil {
		var ret string
		return ret
	}
	return *o.ExportProjectName
}

// GetExportProjectNameOk returns a tuple with the ExportProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetExportProjectNameOk() (*string, bool) {
	if o == nil || o.ExportProjectName == nil {
		return nil, false
	}
	return o.ExportProjectName, true
}

// HasExportProjectName returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasExportProjectName() bool {
	return o != nil && o.ExportProjectName != nil
}

// SetExportProjectName gets a reference to the given string and assigns it to the ExportProjectName field.
func (o *GcpUcConfigResponseDataAttributes) SetExportProjectName(v string) {
	o.ExportProjectName = &v
}

// GetMonths returns the Months field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetMonths() int64 {
	if o == nil || o.Months == nil {
		var ret int64
		return ret
	}
	return *o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetMonthsOk() (*int64, bool) {
	if o == nil || o.Months == nil {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasMonths() bool {
	return o != nil && o.Months != nil
}

// SetMonths gets a reference to the given int64 and assigns it to the Months field.
func (o *GcpUcConfigResponseDataAttributes) SetMonths(v int64) {
	o.Months = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasProjectId() bool {
	return o != nil && o.ProjectId != nil
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GcpUcConfigResponseDataAttributes) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetServiceAccount() string {
	if o == nil || o.ServiceAccount == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetServiceAccountOk() (*string, bool) {
	if o == nil || o.ServiceAccount == nil {
		return nil, false
	}
	return o.ServiceAccount, true
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasServiceAccount() bool {
	return o != nil && o.ServiceAccount != nil
}

// SetServiceAccount gets a reference to the given string and assigns it to the ServiceAccount field.
func (o *GcpUcConfigResponseDataAttributes) SetServiceAccount(v string) {
	o.ServiceAccount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GcpUcConfigResponseDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetStatusUpdatedAt returns the StatusUpdatedAt field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetStatusUpdatedAt() string {
	if o == nil || o.StatusUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.StatusUpdatedAt
}

// GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetStatusUpdatedAtOk() (*string, bool) {
	if o == nil || o.StatusUpdatedAt == nil {
		return nil, false
	}
	return o.StatusUpdatedAt, true
}

// HasStatusUpdatedAt returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasStatusUpdatedAt() bool {
	return o != nil && o.StatusUpdatedAt != nil
}

// SetStatusUpdatedAt gets a reference to the given string and assigns it to the StatusUpdatedAt field.
func (o *GcpUcConfigResponseDataAttributes) SetStatusUpdatedAt(v string) {
	o.StatusUpdatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GcpUcConfigResponseDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpUcConfigResponseDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GcpUcConfigResponseDataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GcpUcConfigResponseDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GcpUcConfigResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.BucketName != nil {
		toSerialize["bucket_name"] = o.BucketName
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Dataset != nil {
		toSerialize["dataset"] = o.Dataset
	}
	if o.ErrorMessages.IsSet() {
		toSerialize["error_messages"] = o.ErrorMessages.Get()
	}
	if o.ExportPrefix != nil {
		toSerialize["export_prefix"] = o.ExportPrefix
	}
	if o.ExportProjectName != nil {
		toSerialize["export_project_name"] = o.ExportProjectName
	}
	if o.Months != nil {
		toSerialize["months"] = o.Months
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.ServiceAccount != nil {
		toSerialize["service_account"] = o.ServiceAccount
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
func (o *GcpUcConfigResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId         *string                      `json:"account_id,omitempty"`
		BucketName        *string                      `json:"bucket_name,omitempty"`
		CreatedAt         *string                      `json:"created_at,omitempty"`
		Dataset           *string                      `json:"dataset,omitempty"`
		ErrorMessages     datadog.NullableList[string] `json:"error_messages,omitempty"`
		ExportPrefix      *string                      `json:"export_prefix,omitempty"`
		ExportProjectName *string                      `json:"export_project_name,omitempty"`
		Months            *int64                       `json:"months,omitempty"`
		ProjectId         *string                      `json:"project_id,omitempty"`
		ServiceAccount    *string                      `json:"service_account,omitempty"`
		Status            *string                      `json:"status,omitempty"`
		StatusUpdatedAt   *string                      `json:"status_updated_at,omitempty"`
		UpdatedAt         *string                      `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "bucket_name", "created_at", "dataset", "error_messages", "export_prefix", "export_project_name", "months", "project_id", "service_account", "status", "status_updated_at", "updated_at"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.BucketName = all.BucketName
	o.CreatedAt = all.CreatedAt
	o.Dataset = all.Dataset
	o.ErrorMessages = all.ErrorMessages
	o.ExportPrefix = all.ExportPrefix
	o.ExportProjectName = all.ExportProjectName
	o.Months = all.Months
	o.ProjectId = all.ProjectId
	o.ServiceAccount = all.ServiceAccount
	o.Status = all.Status
	o.StatusUpdatedAt = all.StatusUpdatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
