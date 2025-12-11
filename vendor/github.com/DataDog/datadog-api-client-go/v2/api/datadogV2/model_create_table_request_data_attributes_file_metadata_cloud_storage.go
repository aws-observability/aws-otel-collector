// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTableRequestDataAttributesFileMetadataCloudStorage Cloud storage file metadata for create requests. Both access_details and sync_enabled are required.
type CreateTableRequestDataAttributesFileMetadataCloudStorage struct {
	// Cloud storage access configuration for the reference table data file.
	AccessDetails CreateTableRequestDataAttributesFileMetadataOneOfAccessDetails `json:"access_details"`
	// Whether this table is synced automatically.
	SyncEnabled bool `json:"sync_enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCreateTableRequestDataAttributesFileMetadataCloudStorage instantiates a new CreateTableRequestDataAttributesFileMetadataCloudStorage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTableRequestDataAttributesFileMetadataCloudStorage(accessDetails CreateTableRequestDataAttributesFileMetadataOneOfAccessDetails, syncEnabled bool) *CreateTableRequestDataAttributesFileMetadataCloudStorage {
	this := CreateTableRequestDataAttributesFileMetadataCloudStorage{}
	this.AccessDetails = accessDetails
	this.SyncEnabled = syncEnabled
	return &this
}

// NewCreateTableRequestDataAttributesFileMetadataCloudStorageWithDefaults instantiates a new CreateTableRequestDataAttributesFileMetadataCloudStorage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTableRequestDataAttributesFileMetadataCloudStorageWithDefaults() *CreateTableRequestDataAttributesFileMetadataCloudStorage {
	this := CreateTableRequestDataAttributesFileMetadataCloudStorage{}
	return &this
}

// GetAccessDetails returns the AccessDetails field value.
func (o *CreateTableRequestDataAttributesFileMetadataCloudStorage) GetAccessDetails() CreateTableRequestDataAttributesFileMetadataOneOfAccessDetails {
	if o == nil {
		var ret CreateTableRequestDataAttributesFileMetadataOneOfAccessDetails
		return ret
	}
	return o.AccessDetails
}

// GetAccessDetailsOk returns a tuple with the AccessDetails field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataCloudStorage) GetAccessDetailsOk() (*CreateTableRequestDataAttributesFileMetadataOneOfAccessDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessDetails, true
}

// SetAccessDetails sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataCloudStorage) SetAccessDetails(v CreateTableRequestDataAttributesFileMetadataOneOfAccessDetails) {
	o.AccessDetails = v
}

// GetSyncEnabled returns the SyncEnabled field value.
func (o *CreateTableRequestDataAttributesFileMetadataCloudStorage) GetSyncEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SyncEnabled
}

// GetSyncEnabledOk returns a tuple with the SyncEnabled field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataCloudStorage) GetSyncEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncEnabled, true
}

// SetSyncEnabled sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataCloudStorage) SetSyncEnabled(v bool) {
	o.SyncEnabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTableRequestDataAttributesFileMetadataCloudStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["access_details"] = o.AccessDetails
	toSerialize["sync_enabled"] = o.SyncEnabled
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateTableRequestDataAttributesFileMetadataCloudStorage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessDetails *CreateTableRequestDataAttributesFileMetadataOneOfAccessDetails `json:"access_details"`
		SyncEnabled   *bool                                                           `json:"sync_enabled"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccessDetails == nil {
		return fmt.Errorf("required field access_details missing")
	}
	if all.SyncEnabled == nil {
		return fmt.Errorf("required field sync_enabled missing")
	}

	hasInvalidField := false
	if all.AccessDetails.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AccessDetails = *all.AccessDetails
	o.SyncEnabled = *all.SyncEnabled

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
