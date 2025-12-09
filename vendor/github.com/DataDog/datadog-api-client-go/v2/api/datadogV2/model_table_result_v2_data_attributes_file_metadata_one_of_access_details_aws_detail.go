// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail Amazon Web Services S3 storage access configuration.
type TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail struct {
	// AWS account ID where the S3 bucket is located.
	AwsAccountId *string `json:"aws_account_id,omitempty"`
	// S3 bucket containing the CSV file.
	AwsBucketName *string `json:"aws_bucket_name,omitempty"`
	// The relative file path from the S3 bucket root to the CSV file.
	FilePath *string `json:"file_path,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail instantiates a new TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail() *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail {
	this := TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail{}
	return &this
}

// NewTableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetailWithDefaults instantiates a new TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetailWithDefaults() *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail {
	this := TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail{}
	return &this
}

// GetAwsAccountId returns the AwsAccountId field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetAwsAccountId() string {
	if o == nil || o.AwsAccountId == nil {
		var ret string
		return ret
	}
	return *o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetAwsAccountIdOk() (*string, bool) {
	if o == nil || o.AwsAccountId == nil {
		return nil, false
	}
	return o.AwsAccountId, true
}

// HasAwsAccountId returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) HasAwsAccountId() bool {
	return o != nil && o.AwsAccountId != nil
}

// SetAwsAccountId gets a reference to the given string and assigns it to the AwsAccountId field.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) SetAwsAccountId(v string) {
	o.AwsAccountId = &v
}

// GetAwsBucketName returns the AwsBucketName field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetAwsBucketName() string {
	if o == nil || o.AwsBucketName == nil {
		var ret string
		return ret
	}
	return *o.AwsBucketName
}

// GetAwsBucketNameOk returns a tuple with the AwsBucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetAwsBucketNameOk() (*string, bool) {
	if o == nil || o.AwsBucketName == nil {
		return nil, false
	}
	return o.AwsBucketName, true
}

// HasAwsBucketName returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) HasAwsBucketName() bool {
	return o != nil && o.AwsBucketName != nil
}

// SetAwsBucketName gets a reference to the given string and assigns it to the AwsBucketName field.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) SetAwsBucketName(v string) {
	o.AwsBucketName = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) HasFilePath() bool {
	return o != nil && o.FilePath != nil
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) SetFilePath(v string) {
	o.FilePath = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AwsAccountId != nil {
		toSerialize["aws_account_id"] = o.AwsAccountId
	}
	if o.AwsBucketName != nil {
		toSerialize["aws_bucket_name"] = o.AwsBucketName
	}
	if o.FilePath != nil {
		toSerialize["file_path"] = o.FilePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AwsAccountId  *string `json:"aws_account_id,omitempty"`
		AwsBucketName *string `json:"aws_bucket_name,omitempty"`
		FilePath      *string `json:"file_path,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aws_account_id", "aws_bucket_name", "file_path"})
	} else {
		return err
	}
	o.AwsAccountId = all.AwsAccountId
	o.AwsBucketName = all.AwsBucketName
	o.FilePath = all.FilePath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
