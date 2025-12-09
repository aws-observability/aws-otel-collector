// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateUploadRequestDataAttributes Upload configuration specifying how data is uploaded by the user, and properties of the table to associate the upload with.
type CreateUploadRequestDataAttributes struct {
	// The CSV file headers that define the schema fields, provided in the same order as the columns in the uploaded file.
	Headers []string `json:"headers"`
	// Number of parts to split the file into for multipart upload.
	PartCount int32 `json:"part_count"`
	// The size of each part in the upload in bytes. All parts except the last one must be at least 5,000,000 bytes.
	PartSize int64 `json:"part_size"`
	// Name of the table to associate with this upload.
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateUploadRequestDataAttributes instantiates a new CreateUploadRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateUploadRequestDataAttributes(headers []string, partCount int32, partSize int64, tableName string) *CreateUploadRequestDataAttributes {
	this := CreateUploadRequestDataAttributes{}
	this.Headers = headers
	this.PartCount = partCount
	this.PartSize = partSize
	this.TableName = tableName
	return &this
}

// NewCreateUploadRequestDataAttributesWithDefaults instantiates a new CreateUploadRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateUploadRequestDataAttributesWithDefaults() *CreateUploadRequestDataAttributes {
	this := CreateUploadRequestDataAttributes{}
	return &this
}

// GetHeaders returns the Headers field value.
func (o *CreateUploadRequestDataAttributes) GetHeaders() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *CreateUploadRequestDataAttributes) GetHeadersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value.
func (o *CreateUploadRequestDataAttributes) SetHeaders(v []string) {
	o.Headers = v
}

// GetPartCount returns the PartCount field value.
func (o *CreateUploadRequestDataAttributes) GetPartCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.PartCount
}

// GetPartCountOk returns a tuple with the PartCount field value
// and a boolean to check if the value has been set.
func (o *CreateUploadRequestDataAttributes) GetPartCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartCount, true
}

// SetPartCount sets field value.
func (o *CreateUploadRequestDataAttributes) SetPartCount(v int32) {
	o.PartCount = v
}

// GetPartSize returns the PartSize field value.
func (o *CreateUploadRequestDataAttributes) GetPartSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PartSize
}

// GetPartSizeOk returns a tuple with the PartSize field value
// and a boolean to check if the value has been set.
func (o *CreateUploadRequestDataAttributes) GetPartSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartSize, true
}

// SetPartSize sets field value.
func (o *CreateUploadRequestDataAttributes) SetPartSize(v int64) {
	o.PartSize = v
}

// GetTableName returns the TableName field value.
func (o *CreateUploadRequestDataAttributes) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *CreateUploadRequestDataAttributes) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *CreateUploadRequestDataAttributes) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateUploadRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["headers"] = o.Headers
	toSerialize["part_count"] = o.PartCount
	toSerialize["part_size"] = o.PartSize
	toSerialize["table_name"] = o.TableName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateUploadRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Headers   *[]string `json:"headers"`
		PartCount *int32    `json:"part_count"`
		PartSize  *int64    `json:"part_size"`
		TableName *string   `json:"table_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Headers == nil {
		return fmt.Errorf("required field headers missing")
	}
	if all.PartCount == nil {
		return fmt.Errorf("required field part_count missing")
	}
	if all.PartSize == nil {
		return fmt.Errorf("required field part_size missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field table_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"headers", "part_count", "part_size", "table_name"})
	} else {
		return err
	}
	o.Headers = *all.Headers
	o.PartCount = *all.PartCount
	o.PartSize = *all.PartSize
	o.TableName = *all.TableName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
