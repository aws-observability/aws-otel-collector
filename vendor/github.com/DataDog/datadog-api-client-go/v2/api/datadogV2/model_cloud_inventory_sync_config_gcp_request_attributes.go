// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudInventorySyncConfigGCPRequestAttributes GCP settings for buckets involved in inventory reporting.
type CloudInventorySyncConfigGCPRequestAttributes struct {
	// GCS bucket name where Datadog reads inventory reports.
	DestinationBucketName string `json:"destination_bucket_name"`
	// GCP project ID for the inventory destination bucket.
	ProjectId string `json:"project_id"`
	// Service account email used to read the destination bucket.
	ServiceAccountEmail string `json:"service_account_email"`
	// GCS bucket name that inventory reports are generated for.
	SourceBucketName string `json:"source_bucket_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudInventorySyncConfigGCPRequestAttributes instantiates a new CloudInventorySyncConfigGCPRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudInventorySyncConfigGCPRequestAttributes(destinationBucketName string, projectId string, serviceAccountEmail string, sourceBucketName string) *CloudInventorySyncConfigGCPRequestAttributes {
	this := CloudInventorySyncConfigGCPRequestAttributes{}
	this.DestinationBucketName = destinationBucketName
	this.ProjectId = projectId
	this.ServiceAccountEmail = serviceAccountEmail
	this.SourceBucketName = sourceBucketName
	return &this
}

// NewCloudInventorySyncConfigGCPRequestAttributesWithDefaults instantiates a new CloudInventorySyncConfigGCPRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudInventorySyncConfigGCPRequestAttributesWithDefaults() *CloudInventorySyncConfigGCPRequestAttributes {
	this := CloudInventorySyncConfigGCPRequestAttributes{}
	return &this
}

// GetDestinationBucketName returns the DestinationBucketName field value.
func (o *CloudInventorySyncConfigGCPRequestAttributes) GetDestinationBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DestinationBucketName
}

// GetDestinationBucketNameOk returns a tuple with the DestinationBucketName field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigGCPRequestAttributes) GetDestinationBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationBucketName, true
}

// SetDestinationBucketName sets field value.
func (o *CloudInventorySyncConfigGCPRequestAttributes) SetDestinationBucketName(v string) {
	o.DestinationBucketName = v
}

// GetProjectId returns the ProjectId field value.
func (o *CloudInventorySyncConfigGCPRequestAttributes) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigGCPRequestAttributes) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *CloudInventorySyncConfigGCPRequestAttributes) SetProjectId(v string) {
	o.ProjectId = v
}

// GetServiceAccountEmail returns the ServiceAccountEmail field value.
func (o *CloudInventorySyncConfigGCPRequestAttributes) GetServiceAccountEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceAccountEmail
}

// GetServiceAccountEmailOk returns a tuple with the ServiceAccountEmail field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigGCPRequestAttributes) GetServiceAccountEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccountEmail, true
}

// SetServiceAccountEmail sets field value.
func (o *CloudInventorySyncConfigGCPRequestAttributes) SetServiceAccountEmail(v string) {
	o.ServiceAccountEmail = v
}

// GetSourceBucketName returns the SourceBucketName field value.
func (o *CloudInventorySyncConfigGCPRequestAttributes) GetSourceBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceBucketName
}

// GetSourceBucketNameOk returns a tuple with the SourceBucketName field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigGCPRequestAttributes) GetSourceBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceBucketName, true
}

// SetSourceBucketName sets field value.
func (o *CloudInventorySyncConfigGCPRequestAttributes) SetSourceBucketName(v string) {
	o.SourceBucketName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudInventorySyncConfigGCPRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["destination_bucket_name"] = o.DestinationBucketName
	toSerialize["project_id"] = o.ProjectId
	toSerialize["service_account_email"] = o.ServiceAccountEmail
	toSerialize["source_bucket_name"] = o.SourceBucketName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudInventorySyncConfigGCPRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DestinationBucketName *string `json:"destination_bucket_name"`
		ProjectId             *string `json:"project_id"`
		ServiceAccountEmail   *string `json:"service_account_email"`
		SourceBucketName      *string `json:"source_bucket_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DestinationBucketName == nil {
		return fmt.Errorf("required field destination_bucket_name missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.ServiceAccountEmail == nil {
		return fmt.Errorf("required field service_account_email missing")
	}
	if all.SourceBucketName == nil {
		return fmt.Errorf("required field source_bucket_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"destination_bucket_name", "project_id", "service_account_email", "source_bucket_name"})
	} else {
		return err
	}
	o.DestinationBucketName = *all.DestinationBucketName
	o.ProjectId = *all.ProjectId
	o.ServiceAccountEmail = *all.ServiceAccountEmail
	o.SourceBucketName = *all.SourceBucketName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
