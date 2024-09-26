// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BillConfig Bill config.
type BillConfig struct {
	// The name of the configured Azure Export.
	ExportName string `json:"export_name"`
	// The path where the Azure Export is saved.
	ExportPath string `json:"export_path"`
	// The name of the storage account where the Azure Export is saved.
	StorageAccount string `json:"storage_account"`
	// The name of the storage container where the Azure Export is saved.
	StorageContainer string `json:"storage_container"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBillConfig instantiates a new BillConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBillConfig(exportName string, exportPath string, storageAccount string, storageContainer string) *BillConfig {
	this := BillConfig{}
	this.ExportName = exportName
	this.ExportPath = exportPath
	this.StorageAccount = storageAccount
	this.StorageContainer = storageContainer
	return &this
}

// NewBillConfigWithDefaults instantiates a new BillConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBillConfigWithDefaults() *BillConfig {
	this := BillConfig{}
	return &this
}

// GetExportName returns the ExportName field value.
func (o *BillConfig) GetExportName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExportName
}

// GetExportNameOk returns a tuple with the ExportName field value
// and a boolean to check if the value has been set.
func (o *BillConfig) GetExportNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportName, true
}

// SetExportName sets field value.
func (o *BillConfig) SetExportName(v string) {
	o.ExportName = v
}

// GetExportPath returns the ExportPath field value.
func (o *BillConfig) GetExportPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExportPath
}

// GetExportPathOk returns a tuple with the ExportPath field value
// and a boolean to check if the value has been set.
func (o *BillConfig) GetExportPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportPath, true
}

// SetExportPath sets field value.
func (o *BillConfig) SetExportPath(v string) {
	o.ExportPath = v
}

// GetStorageAccount returns the StorageAccount field value.
func (o *BillConfig) GetStorageAccount() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageAccount
}

// GetStorageAccountOk returns a tuple with the StorageAccount field value
// and a boolean to check if the value has been set.
func (o *BillConfig) GetStorageAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageAccount, true
}

// SetStorageAccount sets field value.
func (o *BillConfig) SetStorageAccount(v string) {
	o.StorageAccount = v
}

// GetStorageContainer returns the StorageContainer field value.
func (o *BillConfig) GetStorageContainer() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value
// and a boolean to check if the value has been set.
func (o *BillConfig) GetStorageContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageContainer, true
}

// SetStorageContainer sets field value.
func (o *BillConfig) SetStorageContainer(v string) {
	o.StorageContainer = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BillConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["export_name"] = o.ExportName
	toSerialize["export_path"] = o.ExportPath
	toSerialize["storage_account"] = o.StorageAccount
	toSerialize["storage_container"] = o.StorageContainer

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BillConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExportName       *string `json:"export_name"`
		ExportPath       *string `json:"export_path"`
		StorageAccount   *string `json:"storage_account"`
		StorageContainer *string `json:"storage_container"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExportName == nil {
		return fmt.Errorf("required field export_name missing")
	}
	if all.ExportPath == nil {
		return fmt.Errorf("required field export_path missing")
	}
	if all.StorageAccount == nil {
		return fmt.Errorf("required field storage_account missing")
	}
	if all.StorageContainer == nil {
		return fmt.Errorf("required field storage_container missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"export_name", "export_path", "storage_account", "storage_container"})
	} else {
		return err
	}
	o.ExportName = *all.ExportName
	o.ExportPath = *all.ExportPath
	o.StorageAccount = *all.StorageAccount
	o.StorageContainer = *all.StorageContainer

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
