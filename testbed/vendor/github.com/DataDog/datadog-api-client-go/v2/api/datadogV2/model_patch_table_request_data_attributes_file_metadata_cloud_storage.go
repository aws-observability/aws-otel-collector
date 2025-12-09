// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchTableRequestDataAttributesFileMetadataCloudStorage Cloud storage file metadata for patch requests. Allows partial updates of access_details and sync_enabled.
type PatchTableRequestDataAttributesFileMetadataCloudStorage struct {
	// Cloud storage access configuration for the reference table data file.
	AccessDetails *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails `json:"access_details,omitempty"`
	// Whether this table is synced automatically.
	SyncEnabled *bool `json:"sync_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewPatchTableRequestDataAttributesFileMetadataCloudStorage instantiates a new PatchTableRequestDataAttributesFileMetadataCloudStorage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchTableRequestDataAttributesFileMetadataCloudStorage() *PatchTableRequestDataAttributesFileMetadataCloudStorage {
	this := PatchTableRequestDataAttributesFileMetadataCloudStorage{}
	return &this
}

// NewPatchTableRequestDataAttributesFileMetadataCloudStorageWithDefaults instantiates a new PatchTableRequestDataAttributesFileMetadataCloudStorage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchTableRequestDataAttributesFileMetadataCloudStorageWithDefaults() *PatchTableRequestDataAttributesFileMetadataCloudStorage {
	this := PatchTableRequestDataAttributesFileMetadataCloudStorage{}
	return &this
}

// GetAccessDetails returns the AccessDetails field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributesFileMetadataCloudStorage) GetAccessDetails() PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails {
	if o == nil || o.AccessDetails == nil {
		var ret PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails
		return ret
	}
	return *o.AccessDetails
}

// GetAccessDetailsOk returns a tuple with the AccessDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributesFileMetadataCloudStorage) GetAccessDetailsOk() (*PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails, bool) {
	if o == nil || o.AccessDetails == nil {
		return nil, false
	}
	return o.AccessDetails, true
}

// HasAccessDetails returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributesFileMetadataCloudStorage) HasAccessDetails() bool {
	return o != nil && o.AccessDetails != nil
}

// SetAccessDetails gets a reference to the given PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails and assigns it to the AccessDetails field.
func (o *PatchTableRequestDataAttributesFileMetadataCloudStorage) SetAccessDetails(v PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) {
	o.AccessDetails = &v
}

// GetSyncEnabled returns the SyncEnabled field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributesFileMetadataCloudStorage) GetSyncEnabled() bool {
	if o == nil || o.SyncEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SyncEnabled
}

// GetSyncEnabledOk returns a tuple with the SyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributesFileMetadataCloudStorage) GetSyncEnabledOk() (*bool, bool) {
	if o == nil || o.SyncEnabled == nil {
		return nil, false
	}
	return o.SyncEnabled, true
}

// HasSyncEnabled returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributesFileMetadataCloudStorage) HasSyncEnabled() bool {
	return o != nil && o.SyncEnabled != nil
}

// SetSyncEnabled gets a reference to the given bool and assigns it to the SyncEnabled field.
func (o *PatchTableRequestDataAttributesFileMetadataCloudStorage) SetSyncEnabled(v bool) {
	o.SyncEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchTableRequestDataAttributesFileMetadataCloudStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccessDetails != nil {
		toSerialize["access_details"] = o.AccessDetails
	}
	if o.SyncEnabled != nil {
		toSerialize["sync_enabled"] = o.SyncEnabled
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchTableRequestDataAttributesFileMetadataCloudStorage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessDetails *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails `json:"access_details,omitempty"`
		SyncEnabled   *bool                                                          `json:"sync_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.AccessDetails != nil && all.AccessDetails.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AccessDetails = all.AccessDetails
	o.SyncEnabled = all.SyncEnabled

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
