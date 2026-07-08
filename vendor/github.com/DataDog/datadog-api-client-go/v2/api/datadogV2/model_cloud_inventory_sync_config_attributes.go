// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudInventorySyncConfigAttributes Attributes for a Storage Management configuration. Fields other than `id` may be empty in the response immediately after a create or update; subsequent reads return the full configuration.
type CloudInventorySyncConfigAttributes struct {
	// AWS account ID for the inventory bucket.
	AwsAccountId string `json:"aws_account_id"`
	// AWS S3 bucket name for inventory files.
	AwsBucketName string `json:"aws_bucket_name"`
	// AWS Region for the inventory bucket.
	AwsRegion string `json:"aws_region"`
	// Azure AD application (client) ID.
	AzureClientId string `json:"azure_client_id"`
	// Azure blob container name.
	AzureContainerName string `json:"azure_container_name"`
	// Azure storage account name.
	AzureStorageAccountName string `json:"azure_storage_account_name"`
	// Azure AD tenant ID.
	AzureTenantId string `json:"azure_tenant_id"`
	// Cloud provider for this sync configuration (`aws`, `gcp`, or `azure`). For requests, must match the provider block supplied under `attributes`.
	CloudProvider CloudInventoryCloudProviderId `json:"cloud_provider"`
	// Human-readable error detail when sync is unhealthy.
	Error string `json:"error"`
	// Machine-readable error code when sync is unhealthy.
	ErrorCode string `json:"error_code"`
	// GCS bucket name for inventory files Datadog reads.
	GcpBucketName string `json:"gcp_bucket_name"`
	// GCP project ID.
	GcpProjectId string `json:"gcp_project_id"`
	// Service account email for bucket access.
	GcpServiceAccountEmail string `json:"gcp_service_account_email"`
	// Object key prefix where inventory reports are written. Returns `/` when reports are written at the bucket root.
	Prefix string `json:"prefix"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudInventorySyncConfigAttributes instantiates a new CloudInventorySyncConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudInventorySyncConfigAttributes(awsAccountId string, awsBucketName string, awsRegion string, azureClientId string, azureContainerName string, azureStorageAccountName string, azureTenantId string, cloudProvider CloudInventoryCloudProviderId, error string, errorCode string, gcpBucketName string, gcpProjectId string, gcpServiceAccountEmail string, prefix string) *CloudInventorySyncConfigAttributes {
	this := CloudInventorySyncConfigAttributes{}
	this.AwsAccountId = awsAccountId
	this.AwsBucketName = awsBucketName
	this.AwsRegion = awsRegion
	this.AzureClientId = azureClientId
	this.AzureContainerName = azureContainerName
	this.AzureStorageAccountName = azureStorageAccountName
	this.AzureTenantId = azureTenantId
	this.CloudProvider = cloudProvider
	this.Error = error
	this.ErrorCode = errorCode
	this.GcpBucketName = gcpBucketName
	this.GcpProjectId = gcpProjectId
	this.GcpServiceAccountEmail = gcpServiceAccountEmail
	this.Prefix = prefix
	return &this
}

// NewCloudInventorySyncConfigAttributesWithDefaults instantiates a new CloudInventorySyncConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudInventorySyncConfigAttributesWithDefaults() *CloudInventorySyncConfigAttributes {
	this := CloudInventorySyncConfigAttributes{}
	return &this
}

// GetAwsAccountId returns the AwsAccountId field value.
func (o *CloudInventorySyncConfigAttributes) GetAwsAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetAwsAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountId, true
}

// SetAwsAccountId sets field value.
func (o *CloudInventorySyncConfigAttributes) SetAwsAccountId(v string) {
	o.AwsAccountId = v
}

// GetAwsBucketName returns the AwsBucketName field value.
func (o *CloudInventorySyncConfigAttributes) GetAwsBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AwsBucketName
}

// GetAwsBucketNameOk returns a tuple with the AwsBucketName field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetAwsBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsBucketName, true
}

// SetAwsBucketName sets field value.
func (o *CloudInventorySyncConfigAttributes) SetAwsBucketName(v string) {
	o.AwsBucketName = v
}

// GetAwsRegion returns the AwsRegion field value.
func (o *CloudInventorySyncConfigAttributes) GetAwsRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetAwsRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsRegion, true
}

// SetAwsRegion sets field value.
func (o *CloudInventorySyncConfigAttributes) SetAwsRegion(v string) {
	o.AwsRegion = v
}

// GetAzureClientId returns the AzureClientId field value.
func (o *CloudInventorySyncConfigAttributes) GetAzureClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureClientId
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetAzureClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureClientId, true
}

// SetAzureClientId sets field value.
func (o *CloudInventorySyncConfigAttributes) SetAzureClientId(v string) {
	o.AzureClientId = v
}

// GetAzureContainerName returns the AzureContainerName field value.
func (o *CloudInventorySyncConfigAttributes) GetAzureContainerName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureContainerName
}

// GetAzureContainerNameOk returns a tuple with the AzureContainerName field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetAzureContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureContainerName, true
}

// SetAzureContainerName sets field value.
func (o *CloudInventorySyncConfigAttributes) SetAzureContainerName(v string) {
	o.AzureContainerName = v
}

// GetAzureStorageAccountName returns the AzureStorageAccountName field value.
func (o *CloudInventorySyncConfigAttributes) GetAzureStorageAccountName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureStorageAccountName
}

// GetAzureStorageAccountNameOk returns a tuple with the AzureStorageAccountName field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetAzureStorageAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureStorageAccountName, true
}

// SetAzureStorageAccountName sets field value.
func (o *CloudInventorySyncConfigAttributes) SetAzureStorageAccountName(v string) {
	o.AzureStorageAccountName = v
}

// GetAzureTenantId returns the AzureTenantId field value.
func (o *CloudInventorySyncConfigAttributes) GetAzureTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetAzureTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureTenantId, true
}

// SetAzureTenantId sets field value.
func (o *CloudInventorySyncConfigAttributes) SetAzureTenantId(v string) {
	o.AzureTenantId = v
}

// GetCloudProvider returns the CloudProvider field value.
func (o *CloudInventorySyncConfigAttributes) GetCloudProvider() CloudInventoryCloudProviderId {
	if o == nil {
		var ret CloudInventoryCloudProviderId
		return ret
	}
	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetCloudProviderOk() (*CloudInventoryCloudProviderId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value.
func (o *CloudInventorySyncConfigAttributes) SetCloudProvider(v CloudInventoryCloudProviderId) {
	o.CloudProvider = v
}

// GetError returns the Error field value.
func (o *CloudInventorySyncConfigAttributes) GetError() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value.
func (o *CloudInventorySyncConfigAttributes) SetError(v string) {
	o.Error = v
}

// GetErrorCode returns the ErrorCode field value.
func (o *CloudInventorySyncConfigAttributes) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value.
func (o *CloudInventorySyncConfigAttributes) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetGcpBucketName returns the GcpBucketName field value.
func (o *CloudInventorySyncConfigAttributes) GetGcpBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GcpBucketName
}

// GetGcpBucketNameOk returns a tuple with the GcpBucketName field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetGcpBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpBucketName, true
}

// SetGcpBucketName sets field value.
func (o *CloudInventorySyncConfigAttributes) SetGcpBucketName(v string) {
	o.GcpBucketName = v
}

// GetGcpProjectId returns the GcpProjectId field value.
func (o *CloudInventorySyncConfigAttributes) GetGcpProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GcpProjectId
}

// GetGcpProjectIdOk returns a tuple with the GcpProjectId field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetGcpProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpProjectId, true
}

// SetGcpProjectId sets field value.
func (o *CloudInventorySyncConfigAttributes) SetGcpProjectId(v string) {
	o.GcpProjectId = v
}

// GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field value.
func (o *CloudInventorySyncConfigAttributes) GetGcpServiceAccountEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GcpServiceAccountEmail
}

// GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetGcpServiceAccountEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpServiceAccountEmail, true
}

// SetGcpServiceAccountEmail sets field value.
func (o *CloudInventorySyncConfigAttributes) SetGcpServiceAccountEmail(v string) {
	o.GcpServiceAccountEmail = v
}

// GetPrefix returns the Prefix field value.
func (o *CloudInventorySyncConfigAttributes) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAttributes) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value.
func (o *CloudInventorySyncConfigAttributes) SetPrefix(v string) {
	o.Prefix = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudInventorySyncConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aws_account_id"] = o.AwsAccountId
	toSerialize["aws_bucket_name"] = o.AwsBucketName
	toSerialize["aws_region"] = o.AwsRegion
	toSerialize["azure_client_id"] = o.AzureClientId
	toSerialize["azure_container_name"] = o.AzureContainerName
	toSerialize["azure_storage_account_name"] = o.AzureStorageAccountName
	toSerialize["azure_tenant_id"] = o.AzureTenantId
	toSerialize["cloud_provider"] = o.CloudProvider
	toSerialize["error"] = o.Error
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["gcp_bucket_name"] = o.GcpBucketName
	toSerialize["gcp_project_id"] = o.GcpProjectId
	toSerialize["gcp_service_account_email"] = o.GcpServiceAccountEmail
	toSerialize["prefix"] = o.Prefix

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudInventorySyncConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AwsAccountId            *string                        `json:"aws_account_id"`
		AwsBucketName           *string                        `json:"aws_bucket_name"`
		AwsRegion               *string                        `json:"aws_region"`
		AzureClientId           *string                        `json:"azure_client_id"`
		AzureContainerName      *string                        `json:"azure_container_name"`
		AzureStorageAccountName *string                        `json:"azure_storage_account_name"`
		AzureTenantId           *string                        `json:"azure_tenant_id"`
		CloudProvider           *CloudInventoryCloudProviderId `json:"cloud_provider"`
		Error                   *string                        `json:"error"`
		ErrorCode               *string                        `json:"error_code"`
		GcpBucketName           *string                        `json:"gcp_bucket_name"`
		GcpProjectId            *string                        `json:"gcp_project_id"`
		GcpServiceAccountEmail  *string                        `json:"gcp_service_account_email"`
		Prefix                  *string                        `json:"prefix"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AwsAccountId == nil {
		return fmt.Errorf("required field aws_account_id missing")
	}
	if all.AwsBucketName == nil {
		return fmt.Errorf("required field aws_bucket_name missing")
	}
	if all.AwsRegion == nil {
		return fmt.Errorf("required field aws_region missing")
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
	if all.CloudProvider == nil {
		return fmt.Errorf("required field cloud_provider missing")
	}
	if all.Error == nil {
		return fmt.Errorf("required field error missing")
	}
	if all.ErrorCode == nil {
		return fmt.Errorf("required field error_code missing")
	}
	if all.GcpBucketName == nil {
		return fmt.Errorf("required field gcp_bucket_name missing")
	}
	if all.GcpProjectId == nil {
		return fmt.Errorf("required field gcp_project_id missing")
	}
	if all.GcpServiceAccountEmail == nil {
		return fmt.Errorf("required field gcp_service_account_email missing")
	}
	if all.Prefix == nil {
		return fmt.Errorf("required field prefix missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aws_account_id", "aws_bucket_name", "aws_region", "azure_client_id", "azure_container_name", "azure_storage_account_name", "azure_tenant_id", "cloud_provider", "error", "error_code", "gcp_bucket_name", "gcp_project_id", "gcp_service_account_email", "prefix"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AwsAccountId = *all.AwsAccountId
	o.AwsBucketName = *all.AwsBucketName
	o.AwsRegion = *all.AwsRegion
	o.AzureClientId = *all.AzureClientId
	o.AzureContainerName = *all.AzureContainerName
	o.AzureStorageAccountName = *all.AzureStorageAccountName
	o.AzureTenantId = *all.AzureTenantId
	if !all.CloudProvider.IsValid() {
		hasInvalidField = true
	} else {
		o.CloudProvider = *all.CloudProvider
	}
	o.Error = *all.Error
	o.ErrorCode = *all.ErrorCode
	o.GcpBucketName = *all.GcpBucketName
	o.GcpProjectId = *all.GcpProjectId
	o.GcpServiceAccountEmail = *all.GcpServiceAccountEmail
	o.Prefix = *all.Prefix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
