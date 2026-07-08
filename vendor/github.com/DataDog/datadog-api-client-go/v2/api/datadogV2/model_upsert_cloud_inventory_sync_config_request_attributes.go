// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertCloudInventorySyncConfigRequestAttributes Settings for the cloud provider specified in `data.id`. Include only the matching provider object (`aws`, `gcp`, or `azure`).
type UpsertCloudInventorySyncConfigRequestAttributes struct {
	// AWS settings for the S3 bucket Storage Management reads inventory reports from.
	Aws *CloudInventorySyncConfigAWSRequestAttributes `json:"aws,omitempty"`
	// Azure settings for the storage account and container with inventory data.
	Azure *CloudInventorySyncConfigAzureRequestAttributes `json:"azure,omitempty"`
	// GCP settings for buckets involved in inventory reporting.
	Gcp *CloudInventorySyncConfigGCPRequestAttributes `json:"gcp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpsertCloudInventorySyncConfigRequestAttributes instantiates a new UpsertCloudInventorySyncConfigRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpsertCloudInventorySyncConfigRequestAttributes() *UpsertCloudInventorySyncConfigRequestAttributes {
	this := UpsertCloudInventorySyncConfigRequestAttributes{}
	return &this
}

// NewUpsertCloudInventorySyncConfigRequestAttributesWithDefaults instantiates a new UpsertCloudInventorySyncConfigRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpsertCloudInventorySyncConfigRequestAttributesWithDefaults() *UpsertCloudInventorySyncConfigRequestAttributes {
	this := UpsertCloudInventorySyncConfigRequestAttributes{}
	return &this
}

// GetAws returns the Aws field value if set, zero value otherwise.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) GetAws() CloudInventorySyncConfigAWSRequestAttributes {
	if o == nil || o.Aws == nil {
		var ret CloudInventorySyncConfigAWSRequestAttributes
		return ret
	}
	return *o.Aws
}

// GetAwsOk returns a tuple with the Aws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) GetAwsOk() (*CloudInventorySyncConfigAWSRequestAttributes, bool) {
	if o == nil || o.Aws == nil {
		return nil, false
	}
	return o.Aws, true
}

// HasAws returns a boolean if a field has been set.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) HasAws() bool {
	return o != nil && o.Aws != nil
}

// SetAws gets a reference to the given CloudInventorySyncConfigAWSRequestAttributes and assigns it to the Aws field.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) SetAws(v CloudInventorySyncConfigAWSRequestAttributes) {
	o.Aws = &v
}

// GetAzure returns the Azure field value if set, zero value otherwise.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) GetAzure() CloudInventorySyncConfigAzureRequestAttributes {
	if o == nil || o.Azure == nil {
		var ret CloudInventorySyncConfigAzureRequestAttributes
		return ret
	}
	return *o.Azure
}

// GetAzureOk returns a tuple with the Azure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) GetAzureOk() (*CloudInventorySyncConfigAzureRequestAttributes, bool) {
	if o == nil || o.Azure == nil {
		return nil, false
	}
	return o.Azure, true
}

// HasAzure returns a boolean if a field has been set.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) HasAzure() bool {
	return o != nil && o.Azure != nil
}

// SetAzure gets a reference to the given CloudInventorySyncConfigAzureRequestAttributes and assigns it to the Azure field.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) SetAzure(v CloudInventorySyncConfigAzureRequestAttributes) {
	o.Azure = &v
}

// GetGcp returns the Gcp field value if set, zero value otherwise.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) GetGcp() CloudInventorySyncConfigGCPRequestAttributes {
	if o == nil || o.Gcp == nil {
		var ret CloudInventorySyncConfigGCPRequestAttributes
		return ret
	}
	return *o.Gcp
}

// GetGcpOk returns a tuple with the Gcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) GetGcpOk() (*CloudInventorySyncConfigGCPRequestAttributes, bool) {
	if o == nil || o.Gcp == nil {
		return nil, false
	}
	return o.Gcp, true
}

// HasGcp returns a boolean if a field has been set.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) HasGcp() bool {
	return o != nil && o.Gcp != nil
}

// SetGcp gets a reference to the given CloudInventorySyncConfigGCPRequestAttributes and assigns it to the Gcp field.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) SetGcp(v CloudInventorySyncConfigGCPRequestAttributes) {
	o.Gcp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpsertCloudInventorySyncConfigRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aws != nil {
		toSerialize["aws"] = o.Aws
	}
	if o.Azure != nil {
		toSerialize["azure"] = o.Azure
	}
	if o.Gcp != nil {
		toSerialize["gcp"] = o.Gcp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpsertCloudInventorySyncConfigRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aws   *CloudInventorySyncConfigAWSRequestAttributes   `json:"aws,omitempty"`
		Azure *CloudInventorySyncConfigAzureRequestAttributes `json:"azure,omitempty"`
		Gcp   *CloudInventorySyncConfigGCPRequestAttributes   `json:"gcp,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aws", "azure", "gcp"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aws != nil && all.Aws.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Aws = all.Aws
	if all.Azure != nil && all.Azure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Azure = all.Azure
	if all.Gcp != nil && all.Gcp.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Gcp = all.Gcp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
