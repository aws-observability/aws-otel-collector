// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabArtifactObjectInfo Information about an artifact file or directory within a run.
type ModelLabArtifactObjectInfo struct {
	// The size of the file in bytes.
	FileSize datadog.NullableInt64 `json:"file_size,omitempty"`
	// Whether this artifact entry is a directory.
	IsDir bool `json:"is_dir"`
	// The path of the artifact relative to the run's artifact root.
	Path string `json:"path"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabArtifactObjectInfo instantiates a new ModelLabArtifactObjectInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabArtifactObjectInfo(isDir bool, path string) *ModelLabArtifactObjectInfo {
	this := ModelLabArtifactObjectInfo{}
	this.IsDir = isDir
	this.Path = path
	return &this
}

// NewModelLabArtifactObjectInfoWithDefaults instantiates a new ModelLabArtifactObjectInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabArtifactObjectInfoWithDefaults() *ModelLabArtifactObjectInfo {
	this := ModelLabArtifactObjectInfo{}
	return &this
}

// GetFileSize returns the FileSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabArtifactObjectInfo) GetFileSize() int64 {
	if o == nil || o.FileSize.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FileSize.Get()
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabArtifactObjectInfo) GetFileSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileSize.Get(), o.FileSize.IsSet()
}

// HasFileSize returns a boolean if a field has been set.
func (o *ModelLabArtifactObjectInfo) HasFileSize() bool {
	return o != nil && o.FileSize.IsSet()
}

// SetFileSize gets a reference to the given datadog.NullableInt64 and assigns it to the FileSize field.
func (o *ModelLabArtifactObjectInfo) SetFileSize(v int64) {
	o.FileSize.Set(&v)
}

// SetFileSizeNil sets the value for FileSize to be an explicit nil.
func (o *ModelLabArtifactObjectInfo) SetFileSizeNil() {
	o.FileSize.Set(nil)
}

// UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil.
func (o *ModelLabArtifactObjectInfo) UnsetFileSize() {
	o.FileSize.Unset()
}

// GetIsDir returns the IsDir field value.
func (o *ModelLabArtifactObjectInfo) GetIsDir() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDir
}

// GetIsDirOk returns a tuple with the IsDir field value
// and a boolean to check if the value has been set.
func (o *ModelLabArtifactObjectInfo) GetIsDirOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDir, true
}

// SetIsDir sets field value.
func (o *ModelLabArtifactObjectInfo) SetIsDir(v bool) {
	o.IsDir = v
}

// GetPath returns the Path field value.
func (o *ModelLabArtifactObjectInfo) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ModelLabArtifactObjectInfo) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *ModelLabArtifactObjectInfo) SetPath(v string) {
	o.Path = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabArtifactObjectInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FileSize.IsSet() {
		toSerialize["file_size"] = o.FileSize.Get()
	}
	toSerialize["is_dir"] = o.IsDir
	toSerialize["path"] = o.Path

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabArtifactObjectInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FileSize datadog.NullableInt64 `json:"file_size,omitempty"`
		IsDir    *bool                 `json:"is_dir"`
		Path     *string               `json:"path"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsDir == nil {
		return fmt.Errorf("required field is_dir missing")
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"file_size", "is_dir", "path"})
	} else {
		return err
	}
	o.FileSize = all.FileSize
	o.IsDir = *all.IsDir
	o.Path = *all.Path

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
