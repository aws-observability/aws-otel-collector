// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPUsageCostConfigAttributes Attributes for a Google Cloud Usage Cost config.
type GCPUsageCostConfigAttributes struct {
	// The Google Cloud account ID.
	AccountId string `json:"account_id"`
	// The Google Cloud bucket name used to store the Usage Cost export.
	BucketName string `json:"bucket_name"`
	// The timestamp when the Google Cloud Usage Cost config was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// The export dataset name used for the Google Cloud Usage Cost Report.
	Dataset string `json:"dataset"`
	// The error messages for the Google Cloud Usage Cost config.
	ErrorMessages datadog.NullableList[string] `json:"error_messages,omitempty"`
	// The export prefix used for the Google Cloud Usage Cost Report.
	ExportPrefix string `json:"export_prefix"`
	// The name of the Google Cloud Usage Cost Report.
	ExportProjectName string `json:"export_project_name"`
	// The number of months the report has been backfilled.
	// Deprecated
	Months *int32 `json:"months,omitempty"`
	// The `project_id` of the Google Cloud Usage Cost report.
	ProjectId *string `json:"project_id,omitempty"`
	// The unique Google Cloud service account email.
	ServiceAccount string `json:"service_account"`
	// The status of the Google Cloud Usage Cost config.
	Status string `json:"status"`
	// The timestamp when the Google Cloud Usage Cost config status was updated.
	StatusUpdatedAt *string `json:"status_updated_at,omitempty"`
	// The timestamp when the Google Cloud Usage Cost config status was updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPUsageCostConfigAttributes instantiates a new GCPUsageCostConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPUsageCostConfigAttributes(accountId string, bucketName string, dataset string, exportPrefix string, exportProjectName string, serviceAccount string, status string) *GCPUsageCostConfigAttributes {
	this := GCPUsageCostConfigAttributes{}
	this.AccountId = accountId
	this.BucketName = bucketName
	this.Dataset = dataset
	this.ExportPrefix = exportPrefix
	this.ExportProjectName = exportProjectName
	this.ServiceAccount = serviceAccount
	this.Status = status
	return &this
}

// NewGCPUsageCostConfigAttributesWithDefaults instantiates a new GCPUsageCostConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPUsageCostConfigAttributesWithDefaults() *GCPUsageCostConfigAttributes {
	this := GCPUsageCostConfigAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *GCPUsageCostConfigAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *GCPUsageCostConfigAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetBucketName returns the BucketName field value.
func (o *GCPUsageCostConfigAttributes) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigAttributes) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value.
func (o *GCPUsageCostConfigAttributes) SetBucketName(v string) {
	o.BucketName = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GCPUsageCostConfigAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GCPUsageCostConfigAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GCPUsageCostConfigAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDataset returns the Dataset field value.
func (o *GCPUsageCostConfigAttributes) GetDataset() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigAttributes) GetDatasetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dataset, true
}

// SetDataset sets field value.
func (o *GCPUsageCostConfigAttributes) SetDataset(v string) {
	o.Dataset = v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GCPUsageCostConfigAttributes) GetErrorMessages() []string {
	if o == nil || o.ErrorMessages.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ErrorMessages.Get()
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GCPUsageCostConfigAttributes) GetErrorMessagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessages.Get(), o.ErrorMessages.IsSet()
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *GCPUsageCostConfigAttributes) HasErrorMessages() bool {
	return o != nil && o.ErrorMessages.IsSet()
}

// SetErrorMessages gets a reference to the given datadog.NullableList[string] and assigns it to the ErrorMessages field.
func (o *GCPUsageCostConfigAttributes) SetErrorMessages(v []string) {
	o.ErrorMessages.Set(&v)
}

// SetErrorMessagesNil sets the value for ErrorMessages to be an explicit nil.
func (o *GCPUsageCostConfigAttributes) SetErrorMessagesNil() {
	o.ErrorMessages.Set(nil)
}

// UnsetErrorMessages ensures that no value is present for ErrorMessages, not even an explicit nil.
func (o *GCPUsageCostConfigAttributes) UnsetErrorMessages() {
	o.ErrorMessages.Unset()
}

// GetExportPrefix returns the ExportPrefix field value.
func (o *GCPUsageCostConfigAttributes) GetExportPrefix() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExportPrefix
}

// GetExportPrefixOk returns a tuple with the ExportPrefix field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigAttributes) GetExportPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportPrefix, true
}

// SetExportPrefix sets field value.
func (o *GCPUsageCostConfigAttributes) SetExportPrefix(v string) {
	o.ExportPrefix = v
}

// GetExportProjectName returns the ExportProjectName field value.
func (o *GCPUsageCostConfigAttributes) GetExportProjectName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExportProjectName
}

// GetExportProjectNameOk returns a tuple with the ExportProjectName field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigAttributes) GetExportProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportProjectName, true
}

// SetExportProjectName sets field value.
func (o *GCPUsageCostConfigAttributes) SetExportProjectName(v string) {
	o.ExportProjectName = v
}

// GetMonths returns the Months field value if set, zero value otherwise.
// Deprecated
func (o *GCPUsageCostConfigAttributes) GetMonths() int32 {
	if o == nil || o.Months == nil {
		var ret int32
		return ret
	}
	return *o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GCPUsageCostConfigAttributes) GetMonthsOk() (*int32, bool) {
	if o == nil || o.Months == nil {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *GCPUsageCostConfigAttributes) HasMonths() bool {
	return o != nil && o.Months != nil
}

// SetMonths gets a reference to the given int32 and assigns it to the Months field.
// Deprecated
func (o *GCPUsageCostConfigAttributes) SetMonths(v int32) {
	o.Months = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GCPUsageCostConfigAttributes) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigAttributes) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GCPUsageCostConfigAttributes) HasProjectId() bool {
	return o != nil && o.ProjectId != nil
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GCPUsageCostConfigAttributes) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetServiceAccount returns the ServiceAccount field value.
func (o *GCPUsageCostConfigAttributes) GetServiceAccount() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigAttributes) GetServiceAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccount, true
}

// SetServiceAccount sets field value.
func (o *GCPUsageCostConfigAttributes) SetServiceAccount(v string) {
	o.ServiceAccount = v
}

// GetStatus returns the Status field value.
func (o *GCPUsageCostConfigAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *GCPUsageCostConfigAttributes) SetStatus(v string) {
	o.Status = v
}

// GetStatusUpdatedAt returns the StatusUpdatedAt field value if set, zero value otherwise.
func (o *GCPUsageCostConfigAttributes) GetStatusUpdatedAt() string {
	if o == nil || o.StatusUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.StatusUpdatedAt
}

// GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigAttributes) GetStatusUpdatedAtOk() (*string, bool) {
	if o == nil || o.StatusUpdatedAt == nil {
		return nil, false
	}
	return o.StatusUpdatedAt, true
}

// HasStatusUpdatedAt returns a boolean if a field has been set.
func (o *GCPUsageCostConfigAttributes) HasStatusUpdatedAt() bool {
	return o != nil && o.StatusUpdatedAt != nil
}

// SetStatusUpdatedAt gets a reference to the given string and assigns it to the StatusUpdatedAt field.
func (o *GCPUsageCostConfigAttributes) SetStatusUpdatedAt(v string) {
	o.StatusUpdatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GCPUsageCostConfigAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GCPUsageCostConfigAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GCPUsageCostConfigAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPUsageCostConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["bucket_name"] = o.BucketName
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["dataset"] = o.Dataset
	if o.ErrorMessages.IsSet() {
		toSerialize["error_messages"] = o.ErrorMessages.Get()
	}
	toSerialize["export_prefix"] = o.ExportPrefix
	toSerialize["export_project_name"] = o.ExportProjectName
	if o.Months != nil {
		toSerialize["months"] = o.Months
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	toSerialize["service_account"] = o.ServiceAccount
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
func (o *GCPUsageCostConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId         *string                      `json:"account_id"`
		BucketName        *string                      `json:"bucket_name"`
		CreatedAt         *string                      `json:"created_at,omitempty"`
		Dataset           *string                      `json:"dataset"`
		ErrorMessages     datadog.NullableList[string] `json:"error_messages,omitempty"`
		ExportPrefix      *string                      `json:"export_prefix"`
		ExportProjectName *string                      `json:"export_project_name"`
		Months            *int32                       `json:"months,omitempty"`
		ProjectId         *string                      `json:"project_id,omitempty"`
		ServiceAccount    *string                      `json:"service_account"`
		Status            *string                      `json:"status"`
		StatusUpdatedAt   *string                      `json:"status_updated_at,omitempty"`
		UpdatedAt         *string                      `json:"updated_at,omitempty"`
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
	if all.Dataset == nil {
		return fmt.Errorf("required field dataset missing")
	}
	if all.ExportPrefix == nil {
		return fmt.Errorf("required field export_prefix missing")
	}
	if all.ExportProjectName == nil {
		return fmt.Errorf("required field export_project_name missing")
	}
	if all.ServiceAccount == nil {
		return fmt.Errorf("required field service_account missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "bucket_name", "created_at", "dataset", "error_messages", "export_prefix", "export_project_name", "months", "project_id", "service_account", "status", "status_updated_at", "updated_at"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.BucketName = *all.BucketName
	o.CreatedAt = all.CreatedAt
	o.Dataset = *all.Dataset
	o.ErrorMessages = all.ErrorMessages
	o.ExportPrefix = *all.ExportPrefix
	o.ExportProjectName = *all.ExportProjectName
	o.Months = all.Months
	o.ProjectId = all.ProjectId
	o.ServiceAccount = *all.ServiceAccount
	o.Status = *all.Status
	o.StatusUpdatedAt = all.StatusUpdatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
