// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabArtifactInfo Information about a project-level artifact file.
type ModelLabArtifactInfo struct {
	// The full artifact path relative to the project's artifact root.
	ArtifactPath string `json:"artifact_path"`
	// The date and time the artifact was created.
	CreatedAt time.Time `json:"created_at"`
	// The size of the file in bytes.
	FileSize datadog.NullableInt64 `json:"file_size,omitempty"`
	// The filename of the artifact.
	Filename string `json:"filename"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabArtifactInfo instantiates a new ModelLabArtifactInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabArtifactInfo(artifactPath string, createdAt time.Time, filename string) *ModelLabArtifactInfo {
	this := ModelLabArtifactInfo{}
	this.ArtifactPath = artifactPath
	this.CreatedAt = createdAt
	this.Filename = filename
	return &this
}

// NewModelLabArtifactInfoWithDefaults instantiates a new ModelLabArtifactInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabArtifactInfoWithDefaults() *ModelLabArtifactInfo {
	this := ModelLabArtifactInfo{}
	return &this
}

// GetArtifactPath returns the ArtifactPath field value.
func (o *ModelLabArtifactInfo) GetArtifactPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ArtifactPath
}

// GetArtifactPathOk returns a tuple with the ArtifactPath field value
// and a boolean to check if the value has been set.
func (o *ModelLabArtifactInfo) GetArtifactPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactPath, true
}

// SetArtifactPath sets field value.
func (o *ModelLabArtifactInfo) SetArtifactPath(v string) {
	o.ArtifactPath = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ModelLabArtifactInfo) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelLabArtifactInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ModelLabArtifactInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabArtifactInfo) GetFileSize() int64 {
	if o == nil || o.FileSize.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FileSize.Get()
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabArtifactInfo) GetFileSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileSize.Get(), o.FileSize.IsSet()
}

// HasFileSize returns a boolean if a field has been set.
func (o *ModelLabArtifactInfo) HasFileSize() bool {
	return o != nil && o.FileSize.IsSet()
}

// SetFileSize gets a reference to the given datadog.NullableInt64 and assigns it to the FileSize field.
func (o *ModelLabArtifactInfo) SetFileSize(v int64) {
	o.FileSize.Set(&v)
}

// SetFileSizeNil sets the value for FileSize to be an explicit nil.
func (o *ModelLabArtifactInfo) SetFileSizeNil() {
	o.FileSize.Set(nil)
}

// UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil.
func (o *ModelLabArtifactInfo) UnsetFileSize() {
	o.FileSize.Unset()
}

// GetFilename returns the Filename field value.
func (o *ModelLabArtifactInfo) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *ModelLabArtifactInfo) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value.
func (o *ModelLabArtifactInfo) SetFilename(v string) {
	o.Filename = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabArtifactInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["artifact_path"] = o.ArtifactPath
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.FileSize.IsSet() {
		toSerialize["file_size"] = o.FileSize.Get()
	}
	toSerialize["filename"] = o.Filename

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabArtifactInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArtifactPath *string               `json:"artifact_path"`
		CreatedAt    *time.Time            `json:"created_at"`
		FileSize     datadog.NullableInt64 `json:"file_size,omitempty"`
		Filename     *string               `json:"filename"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ArtifactPath == nil {
		return fmt.Errorf("required field artifact_path missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Filename == nil {
		return fmt.Errorf("required field filename missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"artifact_path", "created_at", "file_size", "filename"})
	} else {
		return err
	}
	o.ArtifactPath = *all.ArtifactPath
	o.CreatedAt = *all.CreatedAt
	o.FileSize = all.FileSize
	o.Filename = *all.Filename

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
