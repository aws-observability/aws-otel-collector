// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPUsageCostConfigPostRequestAttributes Attributes for Google Cloud Usage Cost config post request.
type GCPUsageCostConfigPostRequestAttributes struct {
	// The Google Cloud account ID.
	BillingAccountId string `json:"billing_account_id"`
	// The Google Cloud bucket name used to store the Usage Cost export.
	BucketName string `json:"bucket_name"`
	// The export dataset name used for the Google Cloud Usage Cost report.
	ExportDatasetName string `json:"export_dataset_name"`
	// The export prefix used for the Google Cloud Usage Cost report.
	ExportPrefix *string `json:"export_prefix,omitempty"`
	// The name of the Google Cloud Usage Cost report.
	ExportProjectName string `json:"export_project_name"`
	// The unique Google Cloud service account email.
	ServiceAccount string `json:"service_account"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPUsageCostConfigPostRequestAttributes instantiates a new GCPUsageCostConfigPostRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPUsageCostConfigPostRequestAttributes(billingAccountId string, bucketName string, exportDatasetName string, exportProjectName string, serviceAccount string) *GCPUsageCostConfigPostRequestAttributes {
	this := GCPUsageCostConfigPostRequestAttributes{}
	this.BillingAccountId = billingAccountId
	this.BucketName = bucketName
	this.ExportDatasetName = exportDatasetName
	this.ExportProjectName = exportProjectName
	this.ServiceAccount = serviceAccount
	return &this
}

// NewGCPUsageCostConfigPostRequestAttributesWithDefaults instantiates a new GCPUsageCostConfigPostRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPUsageCostConfigPostRequestAttributesWithDefaults() *GCPUsageCostConfigPostRequestAttributes {
	this := GCPUsageCostConfigPostRequestAttributes{}
	return &this
}

// GetBillingAccountId returns the BillingAccountId field value.
func (o *GCPUsageCostConfigPostRequestAttributes) GetBillingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BillingAccountId
}

// GetBillingAccountIdOk returns a tuple with the BillingAccountId field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetBillingAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingAccountId, true
}

// SetBillingAccountId sets field value.
func (o *GCPUsageCostConfigPostRequestAttributes) SetBillingAccountId(v string) {
	o.BillingAccountId = v
}

// GetBucketName returns the BucketName field value.
func (o *GCPUsageCostConfigPostRequestAttributes) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value.
func (o *GCPUsageCostConfigPostRequestAttributes) SetBucketName(v string) {
	o.BucketName = v
}

// GetExportDatasetName returns the ExportDatasetName field value.
func (o *GCPUsageCostConfigPostRequestAttributes) GetExportDatasetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExportDatasetName
}

// GetExportDatasetNameOk returns a tuple with the ExportDatasetName field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetExportDatasetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportDatasetName, true
}

// SetExportDatasetName sets field value.
func (o *GCPUsageCostConfigPostRequestAttributes) SetExportDatasetName(v string) {
	o.ExportDatasetName = v
}

// GetExportPrefix returns the ExportPrefix field value if set, zero value otherwise.
func (o *GCPUsageCostConfigPostRequestAttributes) GetExportPrefix() string {
	if o == nil || o.ExportPrefix == nil {
		var ret string
		return ret
	}
	return *o.ExportPrefix
}

// GetExportPrefixOk returns a tuple with the ExportPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetExportPrefixOk() (*string, bool) {
	if o == nil || o.ExportPrefix == nil {
		return nil, false
	}
	return o.ExportPrefix, true
}

// HasExportPrefix returns a boolean if a field has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) HasExportPrefix() bool {
	return o != nil && o.ExportPrefix != nil
}

// SetExportPrefix gets a reference to the given string and assigns it to the ExportPrefix field.
func (o *GCPUsageCostConfigPostRequestAttributes) SetExportPrefix(v string) {
	o.ExportPrefix = &v
}

// GetExportProjectName returns the ExportProjectName field value.
func (o *GCPUsageCostConfigPostRequestAttributes) GetExportProjectName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExportProjectName
}

// GetExportProjectNameOk returns a tuple with the ExportProjectName field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetExportProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportProjectName, true
}

// SetExportProjectName sets field value.
func (o *GCPUsageCostConfigPostRequestAttributes) SetExportProjectName(v string) {
	o.ExportProjectName = v
}

// GetServiceAccount returns the ServiceAccount field value.
func (o *GCPUsageCostConfigPostRequestAttributes) GetServiceAccount() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetServiceAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccount, true
}

// SetServiceAccount sets field value.
func (o *GCPUsageCostConfigPostRequestAttributes) SetServiceAccount(v string) {
	o.ServiceAccount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPUsageCostConfigPostRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["billing_account_id"] = o.BillingAccountId
	toSerialize["bucket_name"] = o.BucketName
	toSerialize["export_dataset_name"] = o.ExportDatasetName
	if o.ExportPrefix != nil {
		toSerialize["export_prefix"] = o.ExportPrefix
	}
	toSerialize["export_project_name"] = o.ExportProjectName
	toSerialize["service_account"] = o.ServiceAccount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPUsageCostConfigPostRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BillingAccountId  *string `json:"billing_account_id"`
		BucketName        *string `json:"bucket_name"`
		ExportDatasetName *string `json:"export_dataset_name"`
		ExportPrefix      *string `json:"export_prefix,omitempty"`
		ExportProjectName *string `json:"export_project_name"`
		ServiceAccount    *string `json:"service_account"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BillingAccountId == nil {
		return fmt.Errorf("required field billing_account_id missing")
	}
	if all.BucketName == nil {
		return fmt.Errorf("required field bucket_name missing")
	}
	if all.ExportDatasetName == nil {
		return fmt.Errorf("required field export_dataset_name missing")
	}
	if all.ExportProjectName == nil {
		return fmt.Errorf("required field export_project_name missing")
	}
	if all.ServiceAccount == nil {
		return fmt.Errorf("required field service_account missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"billing_account_id", "bucket_name", "export_dataset_name", "export_prefix", "export_project_name", "service_account"})
	} else {
		return err
	}
	o.BillingAccountId = *all.BillingAccountId
	o.BucketName = *all.BucketName
	o.ExportDatasetName = *all.ExportDatasetName
	o.ExportPrefix = all.ExportPrefix
	o.ExportProjectName = *all.ExportProjectName
	o.ServiceAccount = *all.ServiceAccount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
