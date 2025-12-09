// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail Google Cloud Platform storage access configuration.
type PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail struct {
	// The relative file path from the GCS bucket root to the CSV file.
	FilePath *string `json:"file_path,omitempty"`
	// GCP bucket containing the CSV file.
	GcpBucketName *string `json:"gcp_bucket_name,omitempty"`
	// GCP project ID where the bucket is located.
	GcpProjectId *string `json:"gcp_project_id,omitempty"`
	// Service account email with read permissions for the GCS bucket.
	GcpServiceAccountEmail *string `json:"gcp_service_account_email,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail instantiates a new PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail() *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail {
	this := PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail{}
	return &this
}

// NewPatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetailWithDefaults instantiates a new PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetailWithDefaults() *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail {
	this := PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail{}
	return &this
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) HasFilePath() bool {
	return o != nil && o.FilePath != nil
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) SetFilePath(v string) {
	o.FilePath = &v
}

// GetGcpBucketName returns the GcpBucketName field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpBucketName() string {
	if o == nil || o.GcpBucketName == nil {
		var ret string
		return ret
	}
	return *o.GcpBucketName
}

// GetGcpBucketNameOk returns a tuple with the GcpBucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpBucketNameOk() (*string, bool) {
	if o == nil || o.GcpBucketName == nil {
		return nil, false
	}
	return o.GcpBucketName, true
}

// HasGcpBucketName returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) HasGcpBucketName() bool {
	return o != nil && o.GcpBucketName != nil
}

// SetGcpBucketName gets a reference to the given string and assigns it to the GcpBucketName field.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) SetGcpBucketName(v string) {
	o.GcpBucketName = &v
}

// GetGcpProjectId returns the GcpProjectId field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpProjectId() string {
	if o == nil || o.GcpProjectId == nil {
		var ret string
		return ret
	}
	return *o.GcpProjectId
}

// GetGcpProjectIdOk returns a tuple with the GcpProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpProjectIdOk() (*string, bool) {
	if o == nil || o.GcpProjectId == nil {
		return nil, false
	}
	return o.GcpProjectId, true
}

// HasGcpProjectId returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) HasGcpProjectId() bool {
	return o != nil && o.GcpProjectId != nil
}

// SetGcpProjectId gets a reference to the given string and assigns it to the GcpProjectId field.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) SetGcpProjectId(v string) {
	o.GcpProjectId = &v
}

// GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpServiceAccountEmail() string {
	if o == nil || o.GcpServiceAccountEmail == nil {
		var ret string
		return ret
	}
	return *o.GcpServiceAccountEmail
}

// GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) GetGcpServiceAccountEmailOk() (*string, bool) {
	if o == nil || o.GcpServiceAccountEmail == nil {
		return nil, false
	}
	return o.GcpServiceAccountEmail, true
}

// HasGcpServiceAccountEmail returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) HasGcpServiceAccountEmail() bool {
	return o != nil && o.GcpServiceAccountEmail != nil
}

// SetGcpServiceAccountEmail gets a reference to the given string and assigns it to the GcpServiceAccountEmail field.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) SetGcpServiceAccountEmail(v string) {
	o.GcpServiceAccountEmail = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FilePath != nil {
		toSerialize["file_path"] = o.FilePath
	}
	if o.GcpBucketName != nil {
		toSerialize["gcp_bucket_name"] = o.GcpBucketName
	}
	if o.GcpProjectId != nil {
		toSerialize["gcp_project_id"] = o.GcpProjectId
	}
	if o.GcpServiceAccountEmail != nil {
		toSerialize["gcp_service_account_email"] = o.GcpServiceAccountEmail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FilePath               *string `json:"file_path,omitempty"`
		GcpBucketName          *string `json:"gcp_bucket_name,omitempty"`
		GcpProjectId           *string `json:"gcp_project_id,omitempty"`
		GcpServiceAccountEmail *string `json:"gcp_service_account_email,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"file_path", "gcp_bucket_name", "gcp_project_id", "gcp_service_account_email"})
	} else {
		return err
	}
	o.FilePath = all.FilePath
	o.GcpBucketName = all.GcpBucketName
	o.GcpProjectId = all.GcpProjectId
	o.GcpServiceAccountEmail = all.GcpServiceAccountEmail

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
