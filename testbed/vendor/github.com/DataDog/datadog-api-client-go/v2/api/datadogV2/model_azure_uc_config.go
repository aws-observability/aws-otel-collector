// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureUCConfig Azure config.
type AzureUCConfig struct {
	// The tenant ID of the Azure account.
	AccountId string `json:"account_id"`
	// The client ID of the Azure account.
	ClientId string `json:"client_id"`
	// The timestamp when the Azure config was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// The dataset type of the Azure config.
	DatasetType string `json:"dataset_type"`
	// The error messages for the Azure config.
	ErrorMessages datadog.NullableList[string] `json:"error_messages,omitempty"`
	// The name of the configured Azure Export.
	ExportName string `json:"export_name"`
	// The path where the Azure Export is saved.
	ExportPath string `json:"export_path"`
	// The ID of the Azure config.
	Id *string `json:"id,omitempty"`
	// The number of months the report has been backfilled.
	// Deprecated
	Months *int32 `json:"months,omitempty"`
	// The scope of your observed subscription.
	Scope string `json:"scope"`
	// The status of the Azure config.
	Status string `json:"status"`
	// The timestamp when the Azure config status was last updated.
	StatusUpdatedAt *string `json:"status_updated_at,omitempty"`
	// The name of the storage account where the Azure Export is saved.
	StorageAccount string `json:"storage_account"`
	// The name of the storage container where the Azure Export is saved.
	StorageContainer string `json:"storage_container"`
	// The timestamp when the Azure config was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAzureUCConfig instantiates a new AzureUCConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureUCConfig(accountId string, clientId string, datasetType string, exportName string, exportPath string, scope string, status string, storageAccount string, storageContainer string) *AzureUCConfig {
	this := AzureUCConfig{}
	this.AccountId = accountId
	this.ClientId = clientId
	this.DatasetType = datasetType
	this.ExportName = exportName
	this.ExportPath = exportPath
	this.Scope = scope
	this.Status = status
	this.StorageAccount = storageAccount
	this.StorageContainer = storageContainer
	return &this
}

// NewAzureUCConfigWithDefaults instantiates a new AzureUCConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureUCConfigWithDefaults() *AzureUCConfig {
	this := AzureUCConfig{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *AzureUCConfig) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *AzureUCConfig) SetAccountId(v string) {
	o.AccountId = v
}

// GetClientId returns the ClientId field value.
func (o *AzureUCConfig) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *AzureUCConfig) SetClientId(v string) {
	o.ClientId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AzureUCConfig) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AzureUCConfig) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AzureUCConfig) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDatasetType returns the DatasetType field value.
func (o *AzureUCConfig) GetDatasetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetType
}

// GetDatasetTypeOk returns a tuple with the DatasetType field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetDatasetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetType, true
}

// SetDatasetType sets field value.
func (o *AzureUCConfig) SetDatasetType(v string) {
	o.DatasetType = v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureUCConfig) GetErrorMessages() []string {
	if o == nil || o.ErrorMessages.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ErrorMessages.Get()
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AzureUCConfig) GetErrorMessagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessages.Get(), o.ErrorMessages.IsSet()
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *AzureUCConfig) HasErrorMessages() bool {
	return o != nil && o.ErrorMessages.IsSet()
}

// SetErrorMessages gets a reference to the given datadog.NullableList[string] and assigns it to the ErrorMessages field.
func (o *AzureUCConfig) SetErrorMessages(v []string) {
	o.ErrorMessages.Set(&v)
}

// SetErrorMessagesNil sets the value for ErrorMessages to be an explicit nil.
func (o *AzureUCConfig) SetErrorMessagesNil() {
	o.ErrorMessages.Set(nil)
}

// UnsetErrorMessages ensures that no value is present for ErrorMessages, not even an explicit nil.
func (o *AzureUCConfig) UnsetErrorMessages() {
	o.ErrorMessages.Unset()
}

// GetExportName returns the ExportName field value.
func (o *AzureUCConfig) GetExportName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExportName
}

// GetExportNameOk returns a tuple with the ExportName field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetExportNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportName, true
}

// SetExportName sets field value.
func (o *AzureUCConfig) SetExportName(v string) {
	o.ExportName = v
}

// GetExportPath returns the ExportPath field value.
func (o *AzureUCConfig) GetExportPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExportPath
}

// GetExportPathOk returns a tuple with the ExportPath field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetExportPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportPath, true
}

// SetExportPath sets field value.
func (o *AzureUCConfig) SetExportPath(v string) {
	o.ExportPath = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureUCConfig) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureUCConfig) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureUCConfig) SetId(v string) {
	o.Id = &v
}

// GetMonths returns the Months field value if set, zero value otherwise.
// Deprecated
func (o *AzureUCConfig) GetMonths() int32 {
	if o == nil || o.Months == nil {
		var ret int32
		return ret
	}
	return *o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AzureUCConfig) GetMonthsOk() (*int32, bool) {
	if o == nil || o.Months == nil {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *AzureUCConfig) HasMonths() bool {
	return o != nil && o.Months != nil
}

// SetMonths gets a reference to the given int32 and assigns it to the Months field.
// Deprecated
func (o *AzureUCConfig) SetMonths(v int32) {
	o.Months = &v
}

// GetScope returns the Scope field value.
func (o *AzureUCConfig) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *AzureUCConfig) SetScope(v string) {
	o.Scope = v
}

// GetStatus returns the Status field value.
func (o *AzureUCConfig) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AzureUCConfig) SetStatus(v string) {
	o.Status = v
}

// GetStatusUpdatedAt returns the StatusUpdatedAt field value if set, zero value otherwise.
func (o *AzureUCConfig) GetStatusUpdatedAt() string {
	if o == nil || o.StatusUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.StatusUpdatedAt
}

// GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetStatusUpdatedAtOk() (*string, bool) {
	if o == nil || o.StatusUpdatedAt == nil {
		return nil, false
	}
	return o.StatusUpdatedAt, true
}

// HasStatusUpdatedAt returns a boolean if a field has been set.
func (o *AzureUCConfig) HasStatusUpdatedAt() bool {
	return o != nil && o.StatusUpdatedAt != nil
}

// SetStatusUpdatedAt gets a reference to the given string and assigns it to the StatusUpdatedAt field.
func (o *AzureUCConfig) SetStatusUpdatedAt(v string) {
	o.StatusUpdatedAt = &v
}

// GetStorageAccount returns the StorageAccount field value.
func (o *AzureUCConfig) GetStorageAccount() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageAccount
}

// GetStorageAccountOk returns a tuple with the StorageAccount field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetStorageAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageAccount, true
}

// SetStorageAccount sets field value.
func (o *AzureUCConfig) SetStorageAccount(v string) {
	o.StorageAccount = v
}

// GetStorageContainer returns the StorageContainer field value.
func (o *AzureUCConfig) GetStorageContainer() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetStorageContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageContainer, true
}

// SetStorageContainer sets field value.
func (o *AzureUCConfig) SetStorageContainer(v string) {
	o.StorageContainer = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AzureUCConfig) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureUCConfig) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AzureUCConfig) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AzureUCConfig) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureUCConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["client_id"] = o.ClientId
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["dataset_type"] = o.DatasetType
	if o.ErrorMessages.IsSet() {
		toSerialize["error_messages"] = o.ErrorMessages.Get()
	}
	toSerialize["export_name"] = o.ExportName
	toSerialize["export_path"] = o.ExportPath
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Months != nil {
		toSerialize["months"] = o.Months
	}
	toSerialize["scope"] = o.Scope
	toSerialize["status"] = o.Status
	if o.StatusUpdatedAt != nil {
		toSerialize["status_updated_at"] = o.StatusUpdatedAt
	}
	toSerialize["storage_account"] = o.StorageAccount
	toSerialize["storage_container"] = o.StorageContainer
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureUCConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId        *string                      `json:"account_id"`
		ClientId         *string                      `json:"client_id"`
		CreatedAt        *string                      `json:"created_at,omitempty"`
		DatasetType      *string                      `json:"dataset_type"`
		ErrorMessages    datadog.NullableList[string] `json:"error_messages,omitempty"`
		ExportName       *string                      `json:"export_name"`
		ExportPath       *string                      `json:"export_path"`
		Id               *string                      `json:"id,omitempty"`
		Months           *int32                       `json:"months,omitempty"`
		Scope            *string                      `json:"scope"`
		Status           *string                      `json:"status"`
		StatusUpdatedAt  *string                      `json:"status_updated_at,omitempty"`
		StorageAccount   *string                      `json:"storage_account"`
		StorageContainer *string                      `json:"storage_container"`
		UpdatedAt        *string                      `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	if all.DatasetType == nil {
		return fmt.Errorf("required field dataset_type missing")
	}
	if all.ExportName == nil {
		return fmt.Errorf("required field export_name missing")
	}
	if all.ExportPath == nil {
		return fmt.Errorf("required field export_path missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.StorageAccount == nil {
		return fmt.Errorf("required field storage_account missing")
	}
	if all.StorageContainer == nil {
		return fmt.Errorf("required field storage_container missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "client_id", "created_at", "dataset_type", "error_messages", "export_name", "export_path", "id", "months", "scope", "status", "status_updated_at", "storage_account", "storage_container", "updated_at"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.ClientId = *all.ClientId
	o.CreatedAt = all.CreatedAt
	o.DatasetType = *all.DatasetType
	o.ErrorMessages = all.ErrorMessages
	o.ExportName = *all.ExportName
	o.ExportPath = *all.ExportPath
	o.Id = all.Id
	o.Months = all.Months
	o.Scope = *all.Scope
	o.Status = *all.Status
	o.StatusUpdatedAt = all.StatusUpdatedAt
	o.StorageAccount = *all.StorageAccount
	o.StorageContainer = *all.StorageContainer
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
