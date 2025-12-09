// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableResultV2DataAttributesFileMetadataLocalFile File metadata for reference tables created by upload. Note that upload_id is only returned in the immediate create/replace response and is not available in subsequent GET requests.
type TableResultV2DataAttributesFileMetadataLocalFile struct {
	// The error message returned from the creation/update.
	ErrorMessage *string `json:"error_message,omitempty"`
	// The number of rows that failed to create/update.
	ErrorRowCount *int64 `json:"error_row_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableResultV2DataAttributesFileMetadataLocalFile instantiates a new TableResultV2DataAttributesFileMetadataLocalFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableResultV2DataAttributesFileMetadataLocalFile() *TableResultV2DataAttributesFileMetadataLocalFile {
	this := TableResultV2DataAttributesFileMetadataLocalFile{}
	return &this
}

// NewTableResultV2DataAttributesFileMetadataLocalFileWithDefaults instantiates a new TableResultV2DataAttributesFileMetadataLocalFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableResultV2DataAttributesFileMetadataLocalFileWithDefaults() *TableResultV2DataAttributesFileMetadataLocalFile {
	this := TableResultV2DataAttributesFileMetadataLocalFile{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataLocalFile) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataLocalFile) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataLocalFile) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *TableResultV2DataAttributesFileMetadataLocalFile) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorRowCount returns the ErrorRowCount field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataLocalFile) GetErrorRowCount() int64 {
	if o == nil || o.ErrorRowCount == nil {
		var ret int64
		return ret
	}
	return *o.ErrorRowCount
}

// GetErrorRowCountOk returns a tuple with the ErrorRowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataLocalFile) GetErrorRowCountOk() (*int64, bool) {
	if o == nil || o.ErrorRowCount == nil {
		return nil, false
	}
	return o.ErrorRowCount, true
}

// HasErrorRowCount returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataLocalFile) HasErrorRowCount() bool {
	return o != nil && o.ErrorRowCount != nil
}

// SetErrorRowCount gets a reference to the given int64 and assigns it to the ErrorRowCount field.
func (o *TableResultV2DataAttributesFileMetadataLocalFile) SetErrorRowCount(v int64) {
	o.ErrorRowCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableResultV2DataAttributesFileMetadataLocalFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.ErrorRowCount != nil {
		toSerialize["error_row_count"] = o.ErrorRowCount
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableResultV2DataAttributesFileMetadataLocalFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorMessage  *string `json:"error_message,omitempty"`
		ErrorRowCount *int64  `json:"error_row_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.ErrorMessage = all.ErrorMessage
	o.ErrorRowCount = all.ErrorRowCount

	return nil
}
