// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableResultV2DataAttributesFileMetadataCloudStorage File metadata for reference tables created by cloud storage.
type TableResultV2DataAttributesFileMetadataCloudStorage struct {
	// Cloud storage access configuration for the reference table data file.
	AccessDetails TableResultV2DataAttributesFileMetadataOneOfAccessDetails `json:"access_details"`
	// The error message returned from the sync.
	ErrorMessage *string `json:"error_message,omitempty"`
	// The number of rows that failed to sync.
	ErrorRowCount *int64 `json:"error_row_count,omitempty"`
	// The type of error that occurred during file processing. This field provides high-level error categories for easier troubleshooting and is only present when there are errors.
	ErrorType *TableResultV2DataAttributesFileMetadataCloudStorageErrorType `json:"error_type,omitempty"`
	// Whether this table is synced automatically.
	SyncEnabled *bool `json:"sync_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableResultV2DataAttributesFileMetadataCloudStorage instantiates a new TableResultV2DataAttributesFileMetadataCloudStorage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableResultV2DataAttributesFileMetadataCloudStorage(accessDetails TableResultV2DataAttributesFileMetadataOneOfAccessDetails) *TableResultV2DataAttributesFileMetadataCloudStorage {
	this := TableResultV2DataAttributesFileMetadataCloudStorage{}
	this.AccessDetails = accessDetails
	return &this
}

// NewTableResultV2DataAttributesFileMetadataCloudStorageWithDefaults instantiates a new TableResultV2DataAttributesFileMetadataCloudStorage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableResultV2DataAttributesFileMetadataCloudStorageWithDefaults() *TableResultV2DataAttributesFileMetadataCloudStorage {
	this := TableResultV2DataAttributesFileMetadataCloudStorage{}
	return &this
}

// GetAccessDetails returns the AccessDetails field value.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) GetAccessDetails() TableResultV2DataAttributesFileMetadataOneOfAccessDetails {
	if o == nil {
		var ret TableResultV2DataAttributesFileMetadataOneOfAccessDetails
		return ret
	}
	return o.AccessDetails
}

// GetAccessDetailsOk returns a tuple with the AccessDetails field value
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) GetAccessDetailsOk() (*TableResultV2DataAttributesFileMetadataOneOfAccessDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessDetails, true
}

// SetAccessDetails sets field value.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) SetAccessDetails(v TableResultV2DataAttributesFileMetadataOneOfAccessDetails) {
	o.AccessDetails = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorRowCount returns the ErrorRowCount field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) GetErrorRowCount() int64 {
	if o == nil || o.ErrorRowCount == nil {
		var ret int64
		return ret
	}
	return *o.ErrorRowCount
}

// GetErrorRowCountOk returns a tuple with the ErrorRowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) GetErrorRowCountOk() (*int64, bool) {
	if o == nil || o.ErrorRowCount == nil {
		return nil, false
	}
	return o.ErrorRowCount, true
}

// HasErrorRowCount returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) HasErrorRowCount() bool {
	return o != nil && o.ErrorRowCount != nil
}

// SetErrorRowCount gets a reference to the given int64 and assigns it to the ErrorRowCount field.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) SetErrorRowCount(v int64) {
	o.ErrorRowCount = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) GetErrorType() TableResultV2DataAttributesFileMetadataCloudStorageErrorType {
	if o == nil || o.ErrorType == nil {
		var ret TableResultV2DataAttributesFileMetadataCloudStorageErrorType
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) GetErrorTypeOk() (*TableResultV2DataAttributesFileMetadataCloudStorageErrorType, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) HasErrorType() bool {
	return o != nil && o.ErrorType != nil
}

// SetErrorType gets a reference to the given TableResultV2DataAttributesFileMetadataCloudStorageErrorType and assigns it to the ErrorType field.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) SetErrorType(v TableResultV2DataAttributesFileMetadataCloudStorageErrorType) {
	o.ErrorType = &v
}

// GetSyncEnabled returns the SyncEnabled field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) GetSyncEnabled() bool {
	if o == nil || o.SyncEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SyncEnabled
}

// GetSyncEnabledOk returns a tuple with the SyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) GetSyncEnabledOk() (*bool, bool) {
	if o == nil || o.SyncEnabled == nil {
		return nil, false
	}
	return o.SyncEnabled, true
}

// HasSyncEnabled returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) HasSyncEnabled() bool {
	return o != nil && o.SyncEnabled != nil
}

// SetSyncEnabled gets a reference to the given bool and assigns it to the SyncEnabled field.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) SetSyncEnabled(v bool) {
	o.SyncEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableResultV2DataAttributesFileMetadataCloudStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["access_details"] = o.AccessDetails
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.ErrorRowCount != nil {
		toSerialize["error_row_count"] = o.ErrorRowCount
	}
	if o.ErrorType != nil {
		toSerialize["error_type"] = o.ErrorType
	}
	if o.SyncEnabled != nil {
		toSerialize["sync_enabled"] = o.SyncEnabled
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableResultV2DataAttributesFileMetadataCloudStorage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessDetails *TableResultV2DataAttributesFileMetadataOneOfAccessDetails    `json:"access_details"`
		ErrorMessage  *string                                                       `json:"error_message,omitempty"`
		ErrorRowCount *int64                                                        `json:"error_row_count,omitempty"`
		ErrorType     *TableResultV2DataAttributesFileMetadataCloudStorageErrorType `json:"error_type,omitempty"`
		SyncEnabled   *bool                                                         `json:"sync_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccessDetails == nil {
		return fmt.Errorf("required field access_details missing")
	}

	hasInvalidField := false
	if all.AccessDetails.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AccessDetails = *all.AccessDetails
	o.ErrorMessage = all.ErrorMessage
	o.ErrorRowCount = all.ErrorRowCount
	if all.ErrorType != nil && !all.ErrorType.IsValid() {
		hasInvalidField = true
	} else {
		o.ErrorType = all.ErrorType
	}
	o.SyncEnabled = all.SyncEnabled

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
