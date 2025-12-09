// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail Google Cloud Platform storage access configuration.
type CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail struct {
	// The relative file path from the GCS bucket root to the CSV file.
	FilePath string `json:"file_path"`
	// GCP bucket containing the CSV file.
	GcpBucketName string `json:"gcp_bucket_name"`
	// GCP project ID where the bucket is located.
	GcpProjectId string `json:"gcp_project_id"`
	// Service account email with read permissions for the GCS bucket.
	GcpServiceAccountEmail string `json:"gcp_service_account_email"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail instantiates a new CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail(filePath string, gcpBucketName string, gcpProjectId string, gcpServiceAccountEmail string) *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail {
	this := CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail{}
	this.FilePath = filePath
	this.GcpBucketName = gcpBucketName
	this.GcpProjectId = gcpProjectId
	this.GcpServiceAccountEmail = gcpServiceAccountEmail
	return &this
}

// NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetailWithDefaults instantiates a new CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetailWithDefaults() *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail {
	this := CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail{}
	return &this
}

// GetFilePath returns the FilePath field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetFilePath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilePath, true
}

// SetFilePath sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) SetFilePath(v string) {
	o.FilePath = v
}

// GetGcpBucketName returns the GcpBucketName field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GcpBucketName
}

// GetGcpBucketNameOk returns a tuple with the GcpBucketName field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpBucketName, true
}

// SetGcpBucketName sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) SetGcpBucketName(v string) {
	o.GcpBucketName = v
}

// GetGcpProjectId returns the GcpProjectId field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GcpProjectId
}

// GetGcpProjectIdOk returns a tuple with the GcpProjectId field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpProjectId, true
}

// SetGcpProjectId sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) SetGcpProjectId(v string) {
	o.GcpProjectId = v
}

// GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpServiceAccountEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GcpServiceAccountEmail
}

// GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpServiceAccountEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpServiceAccountEmail, true
}

// SetGcpServiceAccountEmail sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) SetGcpServiceAccountEmail(v string) {
	o.GcpServiceAccountEmail = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["file_path"] = o.FilePath
	toSerialize["gcp_bucket_name"] = o.GcpBucketName
	toSerialize["gcp_project_id"] = o.GcpProjectId
	toSerialize["gcp_service_account_email"] = o.GcpServiceAccountEmail

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FilePath               *string `json:"file_path"`
		GcpBucketName          *string `json:"gcp_bucket_name"`
		GcpProjectId           *string `json:"gcp_project_id"`
		GcpServiceAccountEmail *string `json:"gcp_service_account_email"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FilePath == nil {
		return fmt.Errorf("required field file_path missing")
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"file_path", "gcp_bucket_name", "gcp_project_id", "gcp_service_account_email"})
	} else {
		return err
	}
	o.FilePath = *all.FilePath
	o.GcpBucketName = *all.GcpBucketName
	o.GcpProjectId = *all.GcpProjectId
	o.GcpServiceAccountEmail = *all.GcpServiceAccountEmail

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
