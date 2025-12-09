// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail Azure Blob Storage access configuration.
type CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail struct {
	// Azure service principal (application) client ID with permissions to read from the container.
	AzureClientId string `json:"azure_client_id"`
	// Azure Blob Storage container containing the CSV file.
	AzureContainerName string `json:"azure_container_name"`
	// Azure storage account where the container is located.
	AzureStorageAccountName string `json:"azure_storage_account_name"`
	// Azure Active Directory tenant ID.
	AzureTenantId string `json:"azure_tenant_id"`
	// The relative file path from the Azure container root to the CSV file.
	FilePath string `json:"file_path"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail instantiates a new CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail(azureClientId string, azureContainerName string, azureStorageAccountName string, azureTenantId string, filePath string) *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail {
	this := CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail{}
	this.AzureClientId = azureClientId
	this.AzureContainerName = azureContainerName
	this.AzureStorageAccountName = azureStorageAccountName
	this.AzureTenantId = azureTenantId
	this.FilePath = filePath
	return &this
}

// NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetailWithDefaults instantiates a new CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetailWithDefaults() *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail {
	this := CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail{}
	return &this
}

// GetAzureClientId returns the AzureClientId field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureClientId
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureClientId, true
}

// SetAzureClientId sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) SetAzureClientId(v string) {
	o.AzureClientId = v
}

// GetAzureContainerName returns the AzureContainerName field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureContainerName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureContainerName
}

// GetAzureContainerNameOk returns a tuple with the AzureContainerName field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureContainerName, true
}

// SetAzureContainerName sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) SetAzureContainerName(v string) {
	o.AzureContainerName = v
}

// GetAzureStorageAccountName returns the AzureStorageAccountName field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureStorageAccountName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureStorageAccountName
}

// GetAzureStorageAccountNameOk returns a tuple with the AzureStorageAccountName field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureStorageAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureStorageAccountName, true
}

// SetAzureStorageAccountName sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) SetAzureStorageAccountName(v string) {
	o.AzureStorageAccountName = v
}

// GetAzureTenantId returns the AzureTenantId field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureTenantId, true
}

// SetAzureTenantId sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) SetAzureTenantId(v string) {
	o.AzureTenantId = v
}

// GetFilePath returns the FilePath field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetFilePath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilePath, true
}

// SetFilePath sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) SetFilePath(v string) {
	o.FilePath = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["azure_client_id"] = o.AzureClientId
	toSerialize["azure_container_name"] = o.AzureContainerName
	toSerialize["azure_storage_account_name"] = o.AzureStorageAccountName
	toSerialize["azure_tenant_id"] = o.AzureTenantId
	toSerialize["file_path"] = o.FilePath

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AzureClientId           *string `json:"azure_client_id"`
		AzureContainerName      *string `json:"azure_container_name"`
		AzureStorageAccountName *string `json:"azure_storage_account_name"`
		AzureTenantId           *string `json:"azure_tenant_id"`
		FilePath                *string `json:"file_path"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AzureClientId == nil {
		return fmt.Errorf("required field azure_client_id missing")
	}
	if all.AzureContainerName == nil {
		return fmt.Errorf("required field azure_container_name missing")
	}
	if all.AzureStorageAccountName == nil {
		return fmt.Errorf("required field azure_storage_account_name missing")
	}
	if all.AzureTenantId == nil {
		return fmt.Errorf("required field azure_tenant_id missing")
	}
	if all.FilePath == nil {
		return fmt.Errorf("required field file_path missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"azure_client_id", "azure_container_name", "azure_storage_account_name", "azure_tenant_id", "file_path"})
	} else {
		return err
	}
	o.AzureClientId = *all.AzureClientId
	o.AzureContainerName = *all.AzureContainerName
	o.AzureStorageAccountName = *all.AzureStorageAccountName
	o.AzureTenantId = *all.AzureTenantId
	o.FilePath = *all.FilePath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
