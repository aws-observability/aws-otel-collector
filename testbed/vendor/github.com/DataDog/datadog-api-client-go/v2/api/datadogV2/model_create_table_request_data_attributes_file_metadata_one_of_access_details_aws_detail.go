// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail Amazon Web Services S3 storage access configuration.
type CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail struct {
	// AWS account ID where the S3 bucket is located.
	AwsAccountId string `json:"aws_account_id"`
	// S3 bucket containing the CSV file.
	AwsBucketName string `json:"aws_bucket_name"`
	// The relative file path from the S3 bucket root to the CSV file.
	FilePath string `json:"file_path"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail instantiates a new CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail(awsAccountId string, awsBucketName string, filePath string) *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail {
	this := CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail{}
	this.AwsAccountId = awsAccountId
	this.AwsBucketName = awsBucketName
	this.FilePath = filePath
	return &this
}

// NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetailWithDefaults instantiates a new CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetailWithDefaults() *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail {
	this := CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail{}
	return &this
}

// GetAwsAccountId returns the AwsAccountId field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetAwsAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetAwsAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountId, true
}

// SetAwsAccountId sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) SetAwsAccountId(v string) {
	o.AwsAccountId = v
}

// GetAwsBucketName returns the AwsBucketName field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetAwsBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AwsBucketName
}

// GetAwsBucketNameOk returns a tuple with the AwsBucketName field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetAwsBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsBucketName, true
}

// SetAwsBucketName sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) SetAwsBucketName(v string) {
	o.AwsBucketName = v
}

// GetFilePath returns the FilePath field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetFilePath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) GetFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilePath, true
}

// SetFilePath sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) SetFilePath(v string) {
	o.FilePath = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aws_account_id"] = o.AwsAccountId
	toSerialize["aws_bucket_name"] = o.AwsBucketName
	toSerialize["file_path"] = o.FilePath

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AwsAccountId  *string `json:"aws_account_id"`
		AwsBucketName *string `json:"aws_bucket_name"`
		FilePath      *string `json:"file_path"`
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
	if all.FilePath == nil {
		return fmt.Errorf("required field file_path missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aws_account_id", "aws_bucket_name", "file_path"})
	} else {
		return err
	}
	o.AwsAccountId = *all.AwsAccountId
	o.AwsBucketName = *all.AwsBucketName
	o.FilePath = *all.FilePath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
