// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudInventorySyncConfigAWSRequestAttributes AWS settings for the S3 bucket Storage Management reads inventory reports from.
type CloudInventorySyncConfigAWSRequestAttributes struct {
	// AWS account ID that owns the inventory bucket.
	AwsAccountId string `json:"aws_account_id"`
	// Name of the S3 bucket containing inventory files.
	DestinationBucketName string `json:"destination_bucket_name"`
	// AWS Region of the inventory bucket.
	DestinationBucketRegion string `json:"destination_bucket_region"`
	// Object key prefix where inventory reports are written. Omit or set to `/` when reports are written at the bucket root.
	DestinationPrefix *string `json:"destination_prefix,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudInventorySyncConfigAWSRequestAttributes instantiates a new CloudInventorySyncConfigAWSRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudInventorySyncConfigAWSRequestAttributes(awsAccountId string, destinationBucketName string, destinationBucketRegion string) *CloudInventorySyncConfigAWSRequestAttributes {
	this := CloudInventorySyncConfigAWSRequestAttributes{}
	this.AwsAccountId = awsAccountId
	this.DestinationBucketName = destinationBucketName
	this.DestinationBucketRegion = destinationBucketRegion
	return &this
}

// NewCloudInventorySyncConfigAWSRequestAttributesWithDefaults instantiates a new CloudInventorySyncConfigAWSRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudInventorySyncConfigAWSRequestAttributesWithDefaults() *CloudInventorySyncConfigAWSRequestAttributes {
	this := CloudInventorySyncConfigAWSRequestAttributes{}
	return &this
}

// GetAwsAccountId returns the AwsAccountId field value.
func (o *CloudInventorySyncConfigAWSRequestAttributes) GetAwsAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAWSRequestAttributes) GetAwsAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountId, true
}

// SetAwsAccountId sets field value.
func (o *CloudInventorySyncConfigAWSRequestAttributes) SetAwsAccountId(v string) {
	o.AwsAccountId = v
}

// GetDestinationBucketName returns the DestinationBucketName field value.
func (o *CloudInventorySyncConfigAWSRequestAttributes) GetDestinationBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DestinationBucketName
}

// GetDestinationBucketNameOk returns a tuple with the DestinationBucketName field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAWSRequestAttributes) GetDestinationBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationBucketName, true
}

// SetDestinationBucketName sets field value.
func (o *CloudInventorySyncConfigAWSRequestAttributes) SetDestinationBucketName(v string) {
	o.DestinationBucketName = v
}

// GetDestinationBucketRegion returns the DestinationBucketRegion field value.
func (o *CloudInventorySyncConfigAWSRequestAttributes) GetDestinationBucketRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DestinationBucketRegion
}

// GetDestinationBucketRegionOk returns a tuple with the DestinationBucketRegion field value
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAWSRequestAttributes) GetDestinationBucketRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationBucketRegion, true
}

// SetDestinationBucketRegion sets field value.
func (o *CloudInventorySyncConfigAWSRequestAttributes) SetDestinationBucketRegion(v string) {
	o.DestinationBucketRegion = v
}

// GetDestinationPrefix returns the DestinationPrefix field value if set, zero value otherwise.
func (o *CloudInventorySyncConfigAWSRequestAttributes) GetDestinationPrefix() string {
	if o == nil || o.DestinationPrefix == nil {
		var ret string
		return ret
	}
	return *o.DestinationPrefix
}

// GetDestinationPrefixOk returns a tuple with the DestinationPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInventorySyncConfigAWSRequestAttributes) GetDestinationPrefixOk() (*string, bool) {
	if o == nil || o.DestinationPrefix == nil {
		return nil, false
	}
	return o.DestinationPrefix, true
}

// HasDestinationPrefix returns a boolean if a field has been set.
func (o *CloudInventorySyncConfigAWSRequestAttributes) HasDestinationPrefix() bool {
	return o != nil && o.DestinationPrefix != nil
}

// SetDestinationPrefix gets a reference to the given string and assigns it to the DestinationPrefix field.
func (o *CloudInventorySyncConfigAWSRequestAttributes) SetDestinationPrefix(v string) {
	o.DestinationPrefix = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudInventorySyncConfigAWSRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aws_account_id"] = o.AwsAccountId
	toSerialize["destination_bucket_name"] = o.DestinationBucketName
	toSerialize["destination_bucket_region"] = o.DestinationBucketRegion
	if o.DestinationPrefix != nil {
		toSerialize["destination_prefix"] = o.DestinationPrefix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudInventorySyncConfigAWSRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AwsAccountId            *string `json:"aws_account_id"`
		DestinationBucketName   *string `json:"destination_bucket_name"`
		DestinationBucketRegion *string `json:"destination_bucket_region"`
		DestinationPrefix       *string `json:"destination_prefix,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AwsAccountId == nil {
		return fmt.Errorf("required field aws_account_id missing")
	}
	if all.DestinationBucketName == nil {
		return fmt.Errorf("required field destination_bucket_name missing")
	}
	if all.DestinationBucketRegion == nil {
		return fmt.Errorf("required field destination_bucket_region missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aws_account_id", "destination_bucket_name", "destination_bucket_region", "destination_prefix"})
	} else {
		return err
	}
	o.AwsAccountId = *all.AwsAccountId
	o.DestinationBucketName = *all.DestinationBucketName
	o.DestinationBucketRegion = *all.DestinationBucketRegion
	o.DestinationPrefix = all.DestinationPrefix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
