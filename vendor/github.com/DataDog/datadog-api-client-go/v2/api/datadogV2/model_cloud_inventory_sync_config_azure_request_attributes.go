// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudInventorySyncConfigAzureRequestAttributes Azure settings for the storage account and container with inventory data.
type CloudInventorySyncConfigAzureRequestAttributes struct {
	// Azure AD application (client) ID used for access.
	ClientId string `json:"client_id"`
	// Blob container name.
	Container string `json:"container"`
	// Resource group containing the storage account.
	ResourceGroup string `json:"resource_group"`
	// Storage account name.
	StorageAccount string `json:"storage_account"`
	// Azure subscription ID.
	SubscriptionId string `json:"subscription_id"`
	// Azure AD tenant ID.
	TenantId string `json:"tenant_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudInventorySyncConfigAzureRequestAttributes instantiates a new CloudInventorySyncConfigAzureRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudInventorySyncConfigAzureRequestAttributes(clientId string, container string, resourceGroup string, storageAccount string, subscriptionId string, tenantId string) *CloudInventorySyncConfigAzureRequestAttributes {
	this := CloudInventorySyncConfigAzureRequestAttributes{}
	this.ClientId = clientId
	this.Container = container
	this.ResourceGroup = resourceGroup
	this.StorageAccount = storageAccount
	this.SubscriptionId = subscriptionId
	this.TenantId = tenantId
	return &this
}

// NewCloudInventorySyncConfigAzureRequestAttributesWithDefaults instantiates a new CloudInventorySyncConfigAzureRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudInventorySyncConfigAzureRequestAttributesWithDefaults() *CloudInventorySyncConfigAzureRequestAttributes {
	this := CloudInventorySyncConfigAzureRequestAttributes{}
	return &this
}

// GetClientId returns the ClientId field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) SetClientId(v string) {
	o.ClientId = v
}

// GetContainer returns the Container field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetContainer() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) SetContainer(v string) {
	o.Container = v
}

// GetResourceGroup returns the ResourceGroup field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetResourceGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetResourceGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroup, true
}

// SetResourceGroup sets field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) SetResourceGroup(v string) {
	o.ResourceGroup = v
}

// GetStorageAccount returns the StorageAccount field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetStorageAccount() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageAccount
}

// GetStorageAccountOk returns a tuple with the StorageAccount field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetStorageAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageAccount, true
}

// SetStorageAccount sets field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) SetStorageAccount(v string) {
	o.StorageAccount = v
}

// GetSubscriptionId returns the SubscriptionId field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetTenantId returns the TenantId field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAzureRequestAttributes) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value.
func (o *CloudInventorySyncConfigAzureRequestAttributes) SetTenantId(v string) {
	o.TenantId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudInventorySyncConfigAzureRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_id"] = o.ClientId
	toSerialize["container"] = o.Container
	toSerialize["resource_group"] = o.ResourceGroup
	toSerialize["storage_account"] = o.StorageAccount
	toSerialize["subscription_id"] = o.SubscriptionId
	toSerialize["tenant_id"] = o.TenantId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudInventorySyncConfigAzureRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientId       *string `json:"client_id"`
		Container      *string `json:"container"`
		ResourceGroup  *string `json:"resource_group"`
		StorageAccount *string `json:"storage_account"`
		SubscriptionId *string `json:"subscription_id"`
		TenantId       *string `json:"tenant_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	if all.Container == nil {
		return fmt.Errorf("required field container missing")
	}
	if all.ResourceGroup == nil {
		return fmt.Errorf("required field resource_group missing")
	}
	if all.StorageAccount == nil {
		return fmt.Errorf("required field storage_account missing")
	}
	if all.SubscriptionId == nil {
		return fmt.Errorf("required field subscription_id missing")
	}
	if all.TenantId == nil {
		return fmt.Errorf("required field tenant_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_id", "container", "resource_group", "storage_account", "subscription_id", "tenant_id"})
	} else {
		return err
	}
	o.ClientId = *all.ClientId
	o.Container = *all.Container
	o.ResourceGroup = *all.ResourceGroup
	o.StorageAccount = *all.StorageAccount
	o.SubscriptionId = *all.SubscriptionId
	o.TenantId = *all.TenantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
