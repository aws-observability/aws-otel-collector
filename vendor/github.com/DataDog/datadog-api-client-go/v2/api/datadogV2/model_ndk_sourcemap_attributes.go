// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NDKSourcemapAttributes Attributes of an Android NDK symbol file.
type NDKSourcemapAttributes struct {
	// The target CPU architecture.
	Arch *string `json:"arch,omitempty"`
	// The build identifier (UUID format).
	BuildId *string `json:"build_id,omitempty"`
	// The timestamp when the symbol file was created.
	CreatedAt time.Time `json:"created_at"`
	// The NDK library file name.
	FileName *string `json:"file_name,omitempty"`
	// The type of source map.
	Mapkind string `json:"mapkind"`
	// The size of the symbol file in bytes.
	Size int64 `json:"size"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNDKSourcemapAttributes instantiates a new NDKSourcemapAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNDKSourcemapAttributes(createdAt time.Time, mapkind string, size int64) *NDKSourcemapAttributes {
	this := NDKSourcemapAttributes{}
	this.CreatedAt = createdAt
	this.Mapkind = mapkind
	this.Size = size
	return &this
}

// NewNDKSourcemapAttributesWithDefaults instantiates a new NDKSourcemapAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNDKSourcemapAttributesWithDefaults() *NDKSourcemapAttributes {
	this := NDKSourcemapAttributes{}
	return &this
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *NDKSourcemapAttributes) GetArch() string {
	if o == nil || o.Arch == nil {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDKSourcemapAttributes) GetArchOk() (*string, bool) {
	if o == nil || o.Arch == nil {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *NDKSourcemapAttributes) HasArch() bool {
	return o != nil && o.Arch != nil
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *NDKSourcemapAttributes) SetArch(v string) {
	o.Arch = &v
}

// GetBuildId returns the BuildId field value if set, zero value otherwise.
func (o *NDKSourcemapAttributes) GetBuildId() string {
	if o == nil || o.BuildId == nil {
		var ret string
		return ret
	}
	return *o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDKSourcemapAttributes) GetBuildIdOk() (*string, bool) {
	if o == nil || o.BuildId == nil {
		return nil, false
	}
	return o.BuildId, true
}

// HasBuildId returns a boolean if a field has been set.
func (o *NDKSourcemapAttributes) HasBuildId() bool {
	return o != nil && o.BuildId != nil
}

// SetBuildId gets a reference to the given string and assigns it to the BuildId field.
func (o *NDKSourcemapAttributes) SetBuildId(v string) {
	o.BuildId = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *NDKSourcemapAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NDKSourcemapAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *NDKSourcemapAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *NDKSourcemapAttributes) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDKSourcemapAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *NDKSourcemapAttributes) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *NDKSourcemapAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetMapkind returns the Mapkind field value.
func (o *NDKSourcemapAttributes) GetMapkind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mapkind
}

// GetMapkindOk returns a tuple with the Mapkind field value
// and a boolean to check if the value has been set.
func (o *NDKSourcemapAttributes) GetMapkindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapkind, true
}

// SetMapkind sets field value.
func (o *NDKSourcemapAttributes) SetMapkind(v string) {
	o.Mapkind = v
}

// GetSize returns the Size field value.
func (o *NDKSourcemapAttributes) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *NDKSourcemapAttributes) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *NDKSourcemapAttributes) SetSize(v int64) {
	o.Size = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NDKSourcemapAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Arch != nil {
		toSerialize["arch"] = o.Arch
	}
	if o.BuildId != nil {
		toSerialize["build_id"] = o.BuildId
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.FileName != nil {
		toSerialize["file_name"] = o.FileName
	}
	toSerialize["mapkind"] = o.Mapkind
	toSerialize["size"] = o.Size

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NDKSourcemapAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arch      *string    `json:"arch,omitempty"`
		BuildId   *string    `json:"build_id,omitempty"`
		CreatedAt *time.Time `json:"created_at"`
		FileName  *string    `json:"file_name,omitempty"`
		Mapkind   *string    `json:"mapkind"`
		Size      *int64     `json:"size"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Mapkind == nil {
		return fmt.Errorf("required field mapkind missing")
	}
	if all.Size == nil {
		return fmt.Errorf("required field size missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arch", "build_id", "created_at", "file_name", "mapkind", "size"})
	} else {
		return err
	}
	o.Arch = all.Arch
	o.BuildId = all.BuildId
	o.CreatedAt = *all.CreatedAt
	o.FileName = all.FileName
	o.Mapkind = *all.Mapkind
	o.Size = *all.Size

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
