// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UCConfigPairDataAttributesConfigsItems The definition of `UCConfigPairDataAttributesConfigsItems` object.
type UCConfigPairDataAttributesConfigsItems struct {
	// The `items` `account_id`.
	AccountId *string `json:"account_id,omitempty"`
	// The `items` `client_id`.
	ClientId *string `json:"client_id,omitempty"`
	// The `items` `created_at`.
	CreatedAt *string `json:"created_at,omitempty"`
	// The `items` `dataset_type`.
	DatasetType *string `json:"dataset_type,omitempty"`
	// The `items` `error_messages`.
	ErrorMessages datadog.NullableList[string] `json:"error_messages,omitempty"`
	// The `items` `export_name`.
	ExportName *string `json:"export_name,omitempty"`
	// The `items` `export_path`.
	ExportPath *string `json:"export_path,omitempty"`
	// The `items` `id`.
	Id *string `json:"id,omitempty"`
	// The `items` `months`.
	Months *int64 `json:"months,omitempty"`
	// The `items` `scope`.
	Scope *string `json:"scope,omitempty"`
	// The `items` `status`.
	Status *string `json:"status,omitempty"`
	// The `items` `status_updated_at`.
	StatusUpdatedAt *string `json:"status_updated_at,omitempty"`
	// The `items` `storage_account`.
	StorageAccount *string `json:"storage_account,omitempty"`
	// The `items` `storage_container`.
	StorageContainer *string `json:"storage_container,omitempty"`
	// The `items` `updated_at`.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUCConfigPairDataAttributesConfigsItems instantiates a new UCConfigPairDataAttributesConfigsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUCConfigPairDataAttributesConfigsItems() *UCConfigPairDataAttributesConfigsItems {
	this := UCConfigPairDataAttributesConfigsItems{}
	return &this
}

// NewUCConfigPairDataAttributesConfigsItemsWithDefaults instantiates a new UCConfigPairDataAttributesConfigsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUCConfigPairDataAttributesConfigsItemsWithDefaults() *UCConfigPairDataAttributesConfigsItems {
	this := UCConfigPairDataAttributesConfigsItems{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *UCConfigPairDataAttributesConfigsItems) SetAccountId(v string) {
	o.AccountId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasClientId() bool {
	return o != nil && o.ClientId != nil
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *UCConfigPairDataAttributesConfigsItems) SetClientId(v string) {
	o.ClientId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *UCConfigPairDataAttributesConfigsItems) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDatasetType returns the DatasetType field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetDatasetType() string {
	if o == nil || o.DatasetType == nil {
		var ret string
		return ret
	}
	return *o.DatasetType
}

// GetDatasetTypeOk returns a tuple with the DatasetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetDatasetTypeOk() (*string, bool) {
	if o == nil || o.DatasetType == nil {
		return nil, false
	}
	return o.DatasetType, true
}

// HasDatasetType returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasDatasetType() bool {
	return o != nil && o.DatasetType != nil
}

// SetDatasetType gets a reference to the given string and assigns it to the DatasetType field.
func (o *UCConfigPairDataAttributesConfigsItems) SetDatasetType(v string) {
	o.DatasetType = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UCConfigPairDataAttributesConfigsItems) GetErrorMessages() []string {
	if o == nil || o.ErrorMessages.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ErrorMessages.Get()
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UCConfigPairDataAttributesConfigsItems) GetErrorMessagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessages.Get(), o.ErrorMessages.IsSet()
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasErrorMessages() bool {
	return o != nil && o.ErrorMessages.IsSet()
}

// SetErrorMessages gets a reference to the given datadog.NullableList[string] and assigns it to the ErrorMessages field.
func (o *UCConfigPairDataAttributesConfigsItems) SetErrorMessages(v []string) {
	o.ErrorMessages.Set(&v)
}

// SetErrorMessagesNil sets the value for ErrorMessages to be an explicit nil.
func (o *UCConfigPairDataAttributesConfigsItems) SetErrorMessagesNil() {
	o.ErrorMessages.Set(nil)
}

// UnsetErrorMessages ensures that no value is present for ErrorMessages, not even an explicit nil.
func (o *UCConfigPairDataAttributesConfigsItems) UnsetErrorMessages() {
	o.ErrorMessages.Unset()
}

// GetExportName returns the ExportName field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetExportName() string {
	if o == nil || o.ExportName == nil {
		var ret string
		return ret
	}
	return *o.ExportName
}

// GetExportNameOk returns a tuple with the ExportName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetExportNameOk() (*string, bool) {
	if o == nil || o.ExportName == nil {
		return nil, false
	}
	return o.ExportName, true
}

// HasExportName returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasExportName() bool {
	return o != nil && o.ExportName != nil
}

// SetExportName gets a reference to the given string and assigns it to the ExportName field.
func (o *UCConfigPairDataAttributesConfigsItems) SetExportName(v string) {
	o.ExportName = &v
}

// GetExportPath returns the ExportPath field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetExportPath() string {
	if o == nil || o.ExportPath == nil {
		var ret string
		return ret
	}
	return *o.ExportPath
}

// GetExportPathOk returns a tuple with the ExportPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetExportPathOk() (*string, bool) {
	if o == nil || o.ExportPath == nil {
		return nil, false
	}
	return o.ExportPath, true
}

// HasExportPath returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasExportPath() bool {
	return o != nil && o.ExportPath != nil
}

// SetExportPath gets a reference to the given string and assigns it to the ExportPath field.
func (o *UCConfigPairDataAttributesConfigsItems) SetExportPath(v string) {
	o.ExportPath = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UCConfigPairDataAttributesConfigsItems) SetId(v string) {
	o.Id = &v
}

// GetMonths returns the Months field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetMonths() int64 {
	if o == nil || o.Months == nil {
		var ret int64
		return ret
	}
	return *o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetMonthsOk() (*int64, bool) {
	if o == nil || o.Months == nil {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasMonths() bool {
	return o != nil && o.Months != nil
}

// SetMonths gets a reference to the given int64 and assigns it to the Months field.
func (o *UCConfigPairDataAttributesConfigsItems) SetMonths(v int64) {
	o.Months = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *UCConfigPairDataAttributesConfigsItems) SetScope(v string) {
	o.Scope = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UCConfigPairDataAttributesConfigsItems) SetStatus(v string) {
	o.Status = &v
}

// GetStatusUpdatedAt returns the StatusUpdatedAt field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetStatusUpdatedAt() string {
	if o == nil || o.StatusUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.StatusUpdatedAt
}

// GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetStatusUpdatedAtOk() (*string, bool) {
	if o == nil || o.StatusUpdatedAt == nil {
		return nil, false
	}
	return o.StatusUpdatedAt, true
}

// HasStatusUpdatedAt returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasStatusUpdatedAt() bool {
	return o != nil && o.StatusUpdatedAt != nil
}

// SetStatusUpdatedAt gets a reference to the given string and assigns it to the StatusUpdatedAt field.
func (o *UCConfigPairDataAttributesConfigsItems) SetStatusUpdatedAt(v string) {
	o.StatusUpdatedAt = &v
}

// GetStorageAccount returns the StorageAccount field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetStorageAccount() string {
	if o == nil || o.StorageAccount == nil {
		var ret string
		return ret
	}
	return *o.StorageAccount
}

// GetStorageAccountOk returns a tuple with the StorageAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetStorageAccountOk() (*string, bool) {
	if o == nil || o.StorageAccount == nil {
		return nil, false
	}
	return o.StorageAccount, true
}

// HasStorageAccount returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasStorageAccount() bool {
	return o != nil && o.StorageAccount != nil
}

// SetStorageAccount gets a reference to the given string and assigns it to the StorageAccount field.
func (o *UCConfigPairDataAttributesConfigsItems) SetStorageAccount(v string) {
	o.StorageAccount = &v
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetStorageContainer() string {
	if o == nil || o.StorageContainer == nil {
		var ret string
		return ret
	}
	return *o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetStorageContainerOk() (*string, bool) {
	if o == nil || o.StorageContainer == nil {
		return nil, false
	}
	return o.StorageContainer, true
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasStorageContainer() bool {
	return o != nil && o.StorageContainer != nil
}

// SetStorageContainer gets a reference to the given string and assigns it to the StorageContainer field.
func (o *UCConfigPairDataAttributesConfigsItems) SetStorageContainer(v string) {
	o.StorageContainer = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UCConfigPairDataAttributesConfigsItems) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairDataAttributesConfigsItems) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UCConfigPairDataAttributesConfigsItems) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *UCConfigPairDataAttributesConfigsItems) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UCConfigPairDataAttributesConfigsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DatasetType != nil {
		toSerialize["dataset_type"] = o.DatasetType
	}
	if o.ErrorMessages.IsSet() {
		toSerialize["error_messages"] = o.ErrorMessages.Get()
	}
	if o.ExportName != nil {
		toSerialize["export_name"] = o.ExportName
	}
	if o.ExportPath != nil {
		toSerialize["export_path"] = o.ExportPath
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Months != nil {
		toSerialize["months"] = o.Months
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusUpdatedAt != nil {
		toSerialize["status_updated_at"] = o.StatusUpdatedAt
	}
	if o.StorageAccount != nil {
		toSerialize["storage_account"] = o.StorageAccount
	}
	if o.StorageContainer != nil {
		toSerialize["storage_container"] = o.StorageContainer
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
func (o *UCConfigPairDataAttributesConfigsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId        *string                      `json:"account_id,omitempty"`
		ClientId         *string                      `json:"client_id,omitempty"`
		CreatedAt        *string                      `json:"created_at,omitempty"`
		DatasetType      *string                      `json:"dataset_type,omitempty"`
		ErrorMessages    datadog.NullableList[string] `json:"error_messages,omitempty"`
		ExportName       *string                      `json:"export_name,omitempty"`
		ExportPath       *string                      `json:"export_path,omitempty"`
		Id               *string                      `json:"id,omitempty"`
		Months           *int64                       `json:"months,omitempty"`
		Scope            *string                      `json:"scope,omitempty"`
		Status           *string                      `json:"status,omitempty"`
		StatusUpdatedAt  *string                      `json:"status_updated_at,omitempty"`
		StorageAccount   *string                      `json:"storage_account,omitempty"`
		StorageContainer *string                      `json:"storage_container,omitempty"`
		UpdatedAt        *string                      `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "client_id", "created_at", "dataset_type", "error_messages", "export_name", "export_path", "id", "months", "scope", "status", "status_updated_at", "storage_account", "storage_container", "updated_at"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.ClientId = all.ClientId
	o.CreatedAt = all.CreatedAt
	o.DatasetType = all.DatasetType
	o.ErrorMessages = all.ErrorMessages
	o.ExportName = all.ExportName
	o.ExportPath = all.ExportPath
	o.Id = all.Id
	o.Months = all.Months
	o.Scope = all.Scope
	o.Status = all.Status
	o.StatusUpdatedAt = all.StatusUpdatedAt
	o.StorageAccount = all.StorageAccount
	o.StorageContainer = all.StorageContainer
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
